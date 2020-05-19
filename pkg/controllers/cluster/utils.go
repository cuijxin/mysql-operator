package cluster

import (
	"strings"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	api "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	"github.com/cuijxin/mysql-operator/pkg/constants"
)

// SelectorForCluster creates a labels.Selector to match a given clusters
// associated resources.
func SelectorForCluster(c *api.MySQLCluster) labels.Selector {
	return labels.SelectorFromSet(labels.Set{constants.MySQLClusterLabel: c.Name})
}

// SelectorForClusterOperatorVersion creates a labels.Selector to match a given clusters
// associated resources MySQLOperatorVersionLabel.
func SelectorForClusterOperatorVersion(operatorVersion string) labels.Selector {
	return labels.SelectorFromSet(labels.Set{constants.MySQLOperatorVersionLabel: operatorVersion})
}

func requiresMySQLAgentStatefulSetUpgrade(ss *apps.StatefulSet, targetContainer string, operatorVersion string) bool {
	if !SelectorForClusterOperatorVersion(operatorVersion).Matches(labels.Set(ss.Labels)) {
		return true
	}
	for _, container := range ss.Spec.Template.Spec.Containers {
		if container.Name == targetContainer {
			parts := strings.Split(container.Image, ":")
			version := parts[len(parts)-1]
			return version != operatorVersion
		}
	}
	return false
}

func requiresMySQLAgentPodUpgrade(pod *v1.Pod, targetContainer string, operatorVersion string) bool {
	if !SelectorForClusterOperatorVersion(operatorVersion).Matches(labels.Set(pod.Labels)) {
		return true
	}
	for _, container := range pod.Spec.Containers {
		if container.Name == targetContainer {
			parts := strings.Split(container.Image, ":")
			version := parts[len(parts)-1]
			return version != operatorVersion
		}
	}
	return false
}

// canUpgradeMySQLAgent checks that pod can actually be updated (e.g. there no backups currently taking place).
// TODO: Implement.
func canUpgradeMySQLAgent(pod *v1.Pod) bool {
	return true
}
