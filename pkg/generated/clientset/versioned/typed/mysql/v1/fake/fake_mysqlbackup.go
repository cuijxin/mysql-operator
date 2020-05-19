// Copyright 2020 Oracle and/or its affiliates. All rights reserved.
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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	mysqlv1 "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMySQLBackups implements MySQLBackupInterface
type FakeMySQLBackups struct {
	Fake *FakeMysql5V1
	ns   string
}

var mysqlbackupsResource = schema.GroupVersionResource{Group: "mysql5.xintelligent.com", Version: "v1", Resource: "mysqlbackups"}

var mysqlbackupsKind = schema.GroupVersionKind{Group: "mysql5.xintelligent.com", Version: "v1", Kind: "MySQLBackup"}

// Get takes name of the mySQLBackup, and returns the corresponding mySQLBackup object, and an error if there is any.
func (c *FakeMySQLBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *mysqlv1.MySQLBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlbackupsResource, c.ns, name), &mysqlv1.MySQLBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackup), err
}

// List takes label and field selectors, and returns the list of MySQLBackups that match those selectors.
func (c *FakeMySQLBackups) List(ctx context.Context, opts v1.ListOptions) (result *mysqlv1.MySQLBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlbackupsResource, mysqlbackupsKind, c.ns, opts), &mysqlv1.MySQLBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mysqlv1.MySQLBackupList{ListMeta: obj.(*mysqlv1.MySQLBackupList).ListMeta}
	for _, item := range obj.(*mysqlv1.MySQLBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLBackups.
func (c *FakeMySQLBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlbackupsResource, c.ns, opts))

}

// Create takes the representation of a mySQLBackup and creates it.  Returns the server's representation of the mySQLBackup, and an error, if there is any.
func (c *FakeMySQLBackups) Create(ctx context.Context, mySQLBackup *mysqlv1.MySQLBackup, opts v1.CreateOptions) (result *mysqlv1.MySQLBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlbackupsResource, c.ns, mySQLBackup), &mysqlv1.MySQLBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackup), err
}

// Update takes the representation of a mySQLBackup and updates it. Returns the server's representation of the mySQLBackup, and an error, if there is any.
func (c *FakeMySQLBackups) Update(ctx context.Context, mySQLBackup *mysqlv1.MySQLBackup, opts v1.UpdateOptions) (result *mysqlv1.MySQLBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlbackupsResource, c.ns, mySQLBackup), &mysqlv1.MySQLBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackup), err
}

// Delete takes name of the mySQLBackup and deletes it. Returns an error if one occurs.
func (c *FakeMySQLBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlbackupsResource, c.ns, name), &mysqlv1.MySQLBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlbackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &mysqlv1.MySQLBackupList{})
	return err
}

// Patch applies the patch and returns the patched mySQLBackup.
func (c *FakeMySQLBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *mysqlv1.MySQLBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlbackupsResource, c.ns, name, pt, data, subresources...), &mysqlv1.MySQLBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackup), err
}
