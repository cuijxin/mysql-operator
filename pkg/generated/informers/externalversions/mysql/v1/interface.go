// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/cuijxin/mysql-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// MySQLBackups returns a MySQLBackupInformer.
	MySQLBackups() MySQLBackupInformer
	// MySQLBackupSchedules returns a MySQLBackupScheduleInformer.
	MySQLBackupSchedules() MySQLBackupScheduleInformer
	// MySQLClusters returns a MySQLClusterInformer.
	MySQLClusters() MySQLClusterInformer
	// MySQLRestores returns a MySQLRestoreInformer.
	MySQLRestores() MySQLRestoreInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// MySQLBackups returns a MySQLBackupInformer.
func (v *version) MySQLBackups() MySQLBackupInformer {
	return &mySQLBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MySQLBackupSchedules returns a MySQLBackupScheduleInformer.
func (v *version) MySQLBackupSchedules() MySQLBackupScheduleInformer {
	return &mySQLBackupScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MySQLClusters returns a MySQLClusterInformer.
func (v *version) MySQLClusters() MySQLClusterInformer {
	return &mySQLClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MySQLRestores returns a MySQLRestoreInformer.
func (v *version) MySQLRestores() MySQLRestoreInformer {
	return &mySQLRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
