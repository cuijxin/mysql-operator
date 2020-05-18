package secrets

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
)

func TestGetRootPasswordSecretName(t *testing.T) {
	cluster := &api.MySQLCluster{
		ObjectMeta: metav1.ObjectMeta{Name: "example-cluster"},
		Spec: api.MySQLClusterSpec{},
	}

	actual := GetRootPasswordSecretName(cluster)

	if actual != "example-cluster-root-password" {
		t.Errorf("Expected example-cluster-root-password but got %s", actual)
	}
}