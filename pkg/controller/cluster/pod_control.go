package cluster

import (
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"

	"github.com/cuijxin/mysql-operator/pkg/constants"
	"github.com/cuijxin/mysql-operator/pkg/controllers/util"
	"github.com/cuijxin/mysql-operator/pkg/resources/statefulsets"
)

// PodControlInterface defines the interface that the
// MySQLClusterController uses to create, update, and delete mysql pods. It
// is implemented as an interface to enable testing.
type PodControlInterface interface {
	PatchPod(old *v1.Pod, new *v1.Pod) error
}

type realPodControl struct {
	client    kubernetes.Interface
	podLister corelisters.PodLister
}

// NewRealPodControl creates a concrete implementation of the
// PodControlInterface.
func NewRealPodControl(client kubernetes.Interface, podLister corelisters.PodLister) PodControlInterface {
	return &realPodControl{client: client, podLister: podLister}
}

func (rpc *realPodControl) PatchPod(old *v1.Pod, new *v1.Pod) error {
	_, err := util.PatchPod(rpc.client, old, new)
	return err
}

// updatePodToOperatorVersion sets the specified MySQLOperator version on:
//   1. The Pod operator version label.
//   2. The MySQLAgent container image version
func updatePodToOperatorVersion(pod *v1.Pod, mysqlAgentImage, version string) *v1.Pod {
	targetContainer := statefulsets.MySQLAgentName
	newAgentImage := fmt.Sprintf("%s:%s", mysqlAgentImage, version)
	newAgentImage = strings.TrimRight(newAgentImage, ":")
	pod.Labels[constants.MySQLOperatorVersionLabel] = version
	for idx, container := range pod.Spec.Containers {
		if container.Name == targetContainer {
			pod.Spec.Containers[idx].Image = newAgentImage
			break
		}
	}
	return pod
}
