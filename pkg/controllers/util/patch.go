package util

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
)

// UpdateStatefulSet performs a direct update for the specified StatefulSet.
func UpdateStatefulSet(kubeClient kubernetes.Interface, newData *apps.StatefulSet) (*apps.StatefulSet, error) {
	result, err := kubeClient.AppsV1().StatefulSets(newData.Namespace).Update(context.TODO(), newData, metav1.UpdateOptions{})
	if err != nil {
		glog.Errorf("Failed to update StatefulSet: %v", err)
		return nil, err
	}

	return result, nil
}

// PatchStatefulSet performs a direct patch update for the specified StatefulSet.
func PatchStatefulSet(kubeClient kubernetes.Interface, oldData *apps.StatefulSet, newData *apps.StatefulSet) (*apps.StatefulSet, error) {
	originalJSON, err := json.Marshal(oldData)
	if err != nil {
		return nil, err
	}

	updatedJSON, err := json.Marshal(newData)
	if err != nil {
		return nil, err
	}

	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(
		originalJSON, updatedJSON, apps.StatefulSet{})
	if err != nil {
		return nil, err
	}
	glog.V(4).Infof("Patching StatefulSet %s/%s: %s", oldData.Name, oldData.ObjectMeta.Namespace, string(patchBytes))

	result, err := kubeClient.AppsV1().StatefulSets(oldData.Namespace).Patch(context.TODO(), oldData.Name, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{}, "")
	if err != nil {
		glog.Errorf("Failed to patch StatefulSet: %v", err)
		return nil, err
	}

	return result, nil
}

// UpdatePod performs a direct update for the specified Pod.
func UpdatePod(kubeClient kubernetes.Interface, newData *v1.Pod) (*v1.Pod, error) {
	result, err := kubeClient.CoreV1().Pods(newData.Namespace).Update(context.TODO(), newData, metav1.UpdateOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update pod")
	}

	return result, nil
}

// PatchPod perform a direct patch update for the specified Pod.
func PatchPod(kubeClient kubernetes.Interface, oldData *v1.Pod, newData *v1.Pod) (*v1.Pod, error) {
	currentPodJSON, err := json.Marshal(oldData)
	if err != nil {
		return nil, err
	}

	updatedPodJSON, err := json.Marshal(newData)
	if err != nil {
		return nil, err
	}

	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(currentPodJSON, updatedPodJSON, v1.Pod{})
	if err != nil {
		return nil, err
	}
	glog.V(4).Infof("Patching Pod %s/%s: %s", oldData.Name, oldData.Namespace, string(patchBytes))

	result, err := kubeClient.CoreV1().Pods(oldData.Namespace).Patch(context.TODO(), oldData.Name, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to patch pod")
	}

	return result, nil
}
