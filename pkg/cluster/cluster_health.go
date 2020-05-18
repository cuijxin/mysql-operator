package cluster

import (
	"context"
	"errors"

	"github.com/cuijxin/mysql-operator/pkg/cluster/innodb"
	"github.com/cuijxin/mysql-operator/pkg/util/mysqlsh"
	"github.com/golang/glog"
	utilexec "k8s.io/utils/exec"
)

// CheckNodeInCluster checks whether or not the local MySQL instance is a member
// of an InnoDB cluster.
func CheckNodeInCluster(ctx context.Context) error {
	instance, err := NewLocalInstance()
	if err != nil {
		return err
	}
	mysh := mysqlsh.New(utilexec.New(), instance.GetShellURI())
	clusterStatus, err := mysh.GetClusterStatus(ctx)
	if err != nil {
		return err
	}
	if clusterStatus.GetInstanceStatus(instance.Name()) != innodb.InstanceStatusOnline {
		return errors.New("database still requires management")
	}
	return nil
}

// GetClusterStatus returns a JSON string representing the status of the InnoDb
// MySQL cluster. TODO: Remove me.
func GetClusterStatus(ctx context.Context) (*innodb.ClusterStatus, error) {
	pod, err := NewLocalInstance()
	if err != nil {
		glog.Errorf("Failed to get the pod details: %+v", err)
		return nil, err
	}

	mysh := mysqlsh.New(utilexec.New(), pod.GetShellURI())
	clusterStatus, err := mysh.GetClusterStatus(ctx)
	if err != nil {
		glog.V(4).Info("Failed to get the cluster status")
		return nil, err
	}
	return clusterStatus, nil
}
