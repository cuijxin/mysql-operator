package app

import (
	"context"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	options "github.com/cuijxin/mysql-operator/cmd/mysql-operator/app/options"
	backupcontroller "github.com/cuijxin/mysql-operator/pkg/controllers/backup"
	backupschedule "github.com/cuijxin/mysql-operator/pkg/controllers/backup/schedule"
	cluster "github.com/cuijxin/mysql-operator/pkg/controllers/cluster"
	restorecontroller "github.com/cuijxin/mysql-operator/pkg/controllers/restore"
	mysqlop "github.com/cuijxin/mysql-operator/pkg/generated/clientset/versioned"
	informers "github.com/cuijxin/mysql-operator/pkg/generated/informers/externalversions"
	metrics "github.com/cuijxin/mysql-operator/pkg/util/metrics"
	signals "github.com/cuijxin/mysql-operator/pkg/util/signals"
)

const (
	metricsEndpoint = "0.0.0.0:8080"
)

// resyncPeriod computes the time interval a shared informer waits before
// resyncing with the api server.
func resyncPeriod(s *options.MySQLOperatorServer) func() time.Duration {
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(s.MinResyncPeriod.Nanoseconds()) * factor)
	}
}

// Run starts the mysql-operator controllers. This should never exit.
func Run(s *options.MySQLOperatorServer) error {
	kubeconfig, err := clientcmd.BuildConfigFromFlags(s.Master, s.KubeConfig)
	if err != nil {
		return err
	}

	// Initialise the operator metrics.
	metrics.RegisterPodName(s.Hostname)
	cluster.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(metricsEndpoint, nil)

	ctx, cancelFunc := context.WithCancel(context.Background())

	// Set up signals so we handle the first shutdown signal gracefully.
	signals.SetupSignalHandler(cancelFunc)

	kubeClient := kubernetes.NewForConfigOrDie(kubeconfig)
	mysqlopClient := mysqlop.NewForConfigOrDie(kubeconfig)

	serverVersion, err := kubeClient.Discovery().ServerVersion()
	if err != nil {
		glog.Fatalf("Failed to discover Kubernetes API server version: %v", err)
	}

	// Shared informers (non namespace specific).
	operatorInformerFactory := informers.NewFilteredSharedInformerFactory(mysqlopClient, resyncPeriod(s)(), s.Namespace, nil)
	kubeInformerFactory := kubeinformers.NewFilteredSharedInformerFactory(kubeClient, resyncPeriod(s)(), s.Namespace, nil)

	var wg sync.WaitGroup

	clusterController := cluster.NewController(
		*s,
		mysqlopClient,
		kubeClient,
		serverVersion,
		operatorInformerFactory.Mysql5().V1().MySQLClusters(),
		kubeInformerFactory.Apps().V1beta1().StatefulSets(),
		kubeInformerFactory.Core().V1().Pods(),
		kubeInformerFactory.Core().V1().Services(),
		kubeInformerFactory.Core().V1().ConfigMaps(),
		30*time.Second,
		s.Namespace,
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		clusterController.Run(ctx, 5)
	}()

	backupController := backupcontroller.NewOperatorController(
		kubeClient,
		mysqlopClient.Mysql5V1(),
		operatorInformerFactory.Mysql5().V1().MySQLBackups(),
		operatorInformerFactory.Mysql5().V1().MySQLClusters(),
		kubeInformerFactory.Core().V1().Pods(),
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		backupController.Run(ctx, 5)
	}()

	restoreController := restorecontroller.NewOperatorController(
		kubeClient,
		mysqlopClient.Mysql5V1(),
		operatorInformerFactory.Mysql5().V1().MySQLRestores(),
		operatorInformerFactory.Mysql5().V1().MySQLClusters(),
		operatorInformerFactory.Mysql5().V1().MySQLBackups(),
		kubeInformerFactory.Core().V1().Pods(),
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		restoreController.Run(ctx, 5)
	}()

	backupScheduleController := backupschedule.NewController(
		mysqlopClient,
		kubeClient,
		operatorInformerFactory.Mysql5().V1().MySQLBackupSchedules(),
		30*time.Second,
		s.Namespace,
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		backupScheduleController.Run(ctx, 1)
	}()

	// Shared informers have to be started after ALL controllers.
	go kubeInformerFactory.Start(ctx.Done())
	go operatorInformerFactory.Start(ctx.Done())

	<-ctx.Done()

	glog.Info("Waiting for all controllers to shut down gracefully")
	wg.Wait()

	return nil
}
