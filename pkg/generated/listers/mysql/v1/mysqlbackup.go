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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MySQLBackupLister helps list MySQLBackups.
type MySQLBackupLister interface {
	// List lists all MySQLBackups in the indexer.
	List(selector labels.Selector) (ret []*v1.MySQLBackup, err error)
	// MySQLBackups returns an object that can list and get MySQLBackups.
	MySQLBackups(namespace string) MySQLBackupNamespaceLister
	MySQLBackupListerExpansion
}

// mySQLBackupLister implements the MySQLBackupLister interface.
type mySQLBackupLister struct {
	indexer cache.Indexer
}

// NewMySQLBackupLister returns a new MySQLBackupLister.
func NewMySQLBackupLister(indexer cache.Indexer) MySQLBackupLister {
	return &mySQLBackupLister{indexer: indexer}
}

// List lists all MySQLBackups in the indexer.
func (s *mySQLBackupLister) List(selector labels.Selector) (ret []*v1.MySQLBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackup))
	})
	return ret, err
}

// MySQLBackups returns an object that can list and get MySQLBackups.
func (s *mySQLBackupLister) MySQLBackups(namespace string) MySQLBackupNamespaceLister {
	return mySQLBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MySQLBackupNamespaceLister helps list and get MySQLBackups.
type MySQLBackupNamespaceLister interface {
	// List lists all MySQLBackups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MySQLBackup, err error)
	// Get retrieves the MySQLBackup from the indexer for a given namespace and name.
	Get(name string) (*v1.MySQLBackup, error)
	MySQLBackupNamespaceListerExpansion
}

// mySQLBackupNamespaceLister implements the MySQLBackupNamespaceLister
// interface.
type mySQLBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MySQLBackups in the indexer for a given namespace.
func (s mySQLBackupNamespaceLister) List(selector labels.Selector) (ret []*v1.MySQLBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackup))
	})
	return ret, err
}

// Get retrieves the MySQLBackup from the indexer for a given namespace and name.
func (s mySQLBackupNamespaceLister) Get(name string) (*v1.MySQLBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mysql5backup"), name)
	}
	return obj.(*v1.MySQLBackup), nil
}
