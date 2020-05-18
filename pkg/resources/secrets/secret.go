package secrets

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	api "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	"github.com/cuijxin/mysql-operator/pkg/constants"
)

// NewMysqlRootPassword returns a Kubernetes secret containing a
// generated MySQL root password.
func NewMysqlRootPassword(cluster *api.MySQLCluster) *v1.Secret {
	CreateSecret := RandomAlphanumericString(16)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{constants.MySQLClusterLabel: cluster.Name},
			Name:   GetRootPasswordSecretName(cluster),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    api.MySQLClusterCRDResourceKind,
				}),
			},
			Namespace: cluster.Namespace,
		},
		Data: map[string][]byte{"password": []byte(CreateSecret)},
	}
	return secret
}

// GetRootPasswordSecretName returns the root password secret name for the
// given mysql cluster.
func GetRootPasswordSecretName(cluster *api.MySQLCluster) string {
	return fmt.Sprintf("%s-root-password", cluster.Name)
}
