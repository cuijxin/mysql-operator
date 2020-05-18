package innodb

import (
	"fmt"
	"net"
)

// DefaultClusterName is the default name assigned to InnoDB clusters created by
// the MySQL operator.
const DefaultClusterName = "MySQLCluster"

// MySQLDBPort is port on which MySQL listens for client connections.
const MySQLDBPort = 3306

// InstanceStatus denotes the status of a MySQL Instance.
type InstanceStatus string

// Instance statuses.
const (
	InstanceStatusOnline     InstanceStatus = "ONLINE"
	InstanceStatusMissing                   = "(MISSING)"
	InstanceStatusRecovering                = "RECOVERING"
	InstanceStatusNotFound                  = ""
	InstanceStatusUnknown                   = "UNKNOWN"
)

// instanceState denotes the state of a MySQL Instance.
type instanceState string

// Instance states.
const (
	instanceStateOk    instanceState = "ok"
	instanceStateError               = "error"
)

// instanceReason denotes the reason for the state of a MySQL Instance.
type instanceReason string

// Instance reasons.
const (
	instanceReasonRecoverable instanceReason = "recoverable"
)

// InstanceMode denotes the mode of a MySQL Instance.
type InstanceMode string

// Instance modes.
const (
	ReadWrite InstanceMode = "R/W"
	ReadOnly               = "R/O"
)

// Instance represents an individual MySQL instance in an InnoDB cluster.
type Instance struct {
	Address string         `json:"address"`
	Mode    InstanceMode   `json:"mode"`
	Role    string         `json:"role"`
	Status  InstanceStatus `json:"status"`
}

// InstanceState represents the state of a MySQL instance with respect to an
// InnoDB cluster.
type InstanceState struct {
	Reason instanceReason `json:"reason"`
	State  instanceState  `json:"state"`
}

// ReplicaSet holds the server instances which belong to an InnoDB
// cluster.
type ReplicaSet struct {
	Name       string               `json:"name"`
	Primary    string               `json:"primary"`
	Status     string               `json:"status"`
	StatusText string               `json:"statusText"`
	Topology   map[string]*Instance `json:"topology"`
}

// DeepCopy takes a deep copy of a ReplicaSet object.
func (rs *ReplicaSet) DeepCopy() *ReplicaSet {
	new := new(ReplicaSet)
	*new = *rs
	for k := range rs.Topology {
		new.Topology[k] = rs.Topology[k].DeepCopy()
	}
	return new
}

// ClusterStatus represents the status of an InnoDB cluster
type ClusterStatus struct {
	ClusterName       string     `json:"clusterName"`
	DefaultReplicaSet ReplicaSet `json:"defaultReplicaSet"`
}

// GetInstanceStatus returns the InstanceStatus of the given instance.
func (s *ClusterStatus) GetInstanceStatus(name string) InstanceStatus {
	if s.DefaultReplicaSet.Topology == nil {
		return InstanceStatusNotFound
	}
	if is, ok := s.DefaultReplicaSet.Topology[fmt.Sprintf("%s:%d", name, MySQLDBPort)]; ok {
		return is.Status
	}
	return InstanceStatusNotFound
}

// GetPrimary returns a primary in the given cluster.
func (s *ClusterStatus) GetPrimaryAddr() (string, error) {
	if s.DefaultReplicaSet.Primary != "" {
		// Single-primary mode.
		return s.DefaultReplicaSet.Primary, nil
	}
	for _, instance := range s.DefaultReplicaSet.Topology {
		// Multi-primary mode.
		if instance.Mode == ReadWrite {
			return instance.Address, nil
		}
	}
	return "", fmt.Errorf("unable to find primary for cluster: %s", s.ClusterName)
}

// DeepCopy takes a deep copy of a ClusterStatus object.
func (s *ClusterStatus) DeepCopy() *ClusterStatus {
	new := new(ClusterStatus)
	*new = *s
	new.DefaultReplicaSet = *s.DefaultReplicaSet.DeepCopy()
	return new
}

// Name returns the dns name of the Instance.
func (i *Instance) Name() string {
	name, _, _ := net.SplitHostPort(i.Address)
	return name
}

// DeepCopy takes a deep copy of an Instance object.
func (i *Instance) DeepCopy() *Instance {
	new := new(Instance)
	*new = *i
	new.Status = i.Status
	return new
}

// CanRejoinCluster returns true if the instance can rejoin the InnoDB cluster.
func (s *InstanceState) CanRejoinCluster() bool {
	return s.State == instanceStateOk && s.Reason == instanceReasonRecoverable
}

// RequiresClearBinaryLogs returns true if the instance needs to clear its binary logs.
func (s *InstanceState) RequiresClearBinaryLogs() bool {
	return s.State == instanceStateError
}
