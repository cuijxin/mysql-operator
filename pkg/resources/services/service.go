package service

import (
	api "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	"github.com/cuijxin/mysql-operator/pkg/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewForCluster will return a new headless Kubernetes service for a MySQL cluster
func NewForCluster(cluster *api.MySQLCluster) *v1.Service {
	mysqlPort := v1.ServicePort{Port: 3306}
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{constants.MySQLClusterLabel: cluster.Name},
			Name:      cluster.Name,
			Namespace: cluster.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    api.MySQLClusterCRDResourceKind,
				}),
			},
			Annotations: map[string]string{
				"service.alpha.kubernetes.io/tolerate-unready-endpoints": "true",
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{mysqlPort},
			Selector: map[string]string{
				constants.MySQLClusterLabel: cluster.Name,
			},
			ClusterIP: v1.ClusterIPNone,
		},
	}

	return svc
}
