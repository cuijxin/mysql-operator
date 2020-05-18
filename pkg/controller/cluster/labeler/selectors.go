package labeler

import (
	labels "k8s.io/apimachinery/pkg/labels"
	selection "k8s.io/apimachinery/pkg/selection"

	constants "github.com/cuijxin/mysql-operator/pkg/constants"
)

// PrimarySelector returns a label selector that selects only primaries of a
// MySQLCluster's Pods.
func PrimarySelector(name string) labels.Selector {
	return labels.SelectorFromSet(labels.Set{
		constants.MySQLClusterLabel:     name,
		constants.LabelMySQLClusterRole: constants.MySQLClusterRolePrimary,
	})
}

// SecondarySelector returns a label selector that selects only secondaries of a
// MySQLCluster's Pods.
func SecondarySelector(name string) labels.Selector {
	return labels.SelectorFromSet(labels.Set{
		constants.MySQLClusterLabel:     name,
		constants.LabelMySQLClusterRole: constants.MySQLClusterRoleSecondary,
	})
}

// NonPrimarySelector returns a label selector that selects all Pods excluding
// primaries of a MySQLCluster.
func NonPrimarySelector(name string) labels.Selector {
	s := labels.SelectorFromSet(labels.Set{constants.MySQLClusterLabel: name})
	requirement, _ := labels.NewRequirement(constants.LabelMySQLClusterRole, selection.NotIn, []string{constants.MySQLClusterRolePrimary})
	return s.Add(*requirement)
}

// HasRoleSelector returns a label selector that selects all Pods for a
// MySQLCluster that have been labeled as having a role.
func HasRoleSelector(name string) labels.Selector {
	s := labels.SelectorFromSet(labels.Set{constants.MySQLClusterLabel: name})
	requirement, _ := labels.NewRequirement(constants.LabelMySQLClusterRole, selection.Exists, []string{})
	return s.Add(*requirement)
}
