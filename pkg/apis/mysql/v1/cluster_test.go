package v1

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestDefaultReplicas(t *testing.T) {
	cluster := &MySQLCluster{}
	cluster.EnsureDefaults()

	if cluster.Spec.Replicas != defaultReplicas {
		t.Errorf("Expected default replicas to be %d but got %d", defaultReplicas, cluster.Spec.Replicas)
	}
}

func TestDefaultVersion(t *testing.T) {
	cluster := &MySQLCluster{}
	cluster.EnsureDefaults()

	if cluster.Spec.Version != defaultVersion {
		t.Errorf("Expected default version to be %s but got %s", defaultVersion, cluster.Spec.Version)
	}
}

func TestRequiresConfigMount(t *testing.T) {
	cluster := &MySQLCluster{}
	cluster.EnsureDefaults()

	if cluster.RequiresConfigMount() {
		t.Errorf("Cluster without configRef should not require a config mount")
	}

	cluster = &MySQLCluster{
		Spec: MySQLClusterSpec{
			ConfigRef: &corev1.LocalObjectReference{
				Name: "customconfig",
			},
		},
	}

	if !cluster.RequiresConfigMount() {
		t.Errorf("Cluster with configRef should require a config mount")
	}
}
