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

// FakeMySQLBackupSchedules implements MySQLBackupScheduleInterface
type FakeMySQLBackupSchedules struct {
	Fake *FakeMysqlV1
	ns   string
}

var mysqlbackupschedulesResource = schema.GroupVersionResource{Group: "mysql.oracle.com", Version: "v1", Resource: "mysqlbackupschedules"}

var mysqlbackupschedulesKind = schema.GroupVersionKind{Group: "mysql.oracle.com", Version: "v1", Kind: "MySQLBackupSchedule"}

// Get takes name of the mySQLBackupSchedule, and returns the corresponding mySQLBackupSchedule object, and an error if there is any.
func (c *FakeMySQLBackupSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *mysqlv1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlbackupschedulesResource, c.ns, name), &mysqlv1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackupSchedule), err
}

// List takes label and field selectors, and returns the list of MySQLBackupSchedules that match those selectors.
func (c *FakeMySQLBackupSchedules) List(ctx context.Context, opts v1.ListOptions) (result *mysqlv1.MySQLBackupScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlbackupschedulesResource, mysqlbackupschedulesKind, c.ns, opts), &mysqlv1.MySQLBackupScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mysqlv1.MySQLBackupScheduleList{ListMeta: obj.(*mysqlv1.MySQLBackupScheduleList).ListMeta}
	for _, item := range obj.(*mysqlv1.MySQLBackupScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLBackupSchedules.
func (c *FakeMySQLBackupSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlbackupschedulesResource, c.ns, opts))

}

// Create takes the representation of a mySQLBackupSchedule and creates it.  Returns the server's representation of the mySQLBackupSchedule, and an error, if there is any.
func (c *FakeMySQLBackupSchedules) Create(ctx context.Context, mySQLBackupSchedule *mysqlv1.MySQLBackupSchedule, opts v1.CreateOptions) (result *mysqlv1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlbackupschedulesResource, c.ns, mySQLBackupSchedule), &mysqlv1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackupSchedule), err
}

// Update takes the representation of a mySQLBackupSchedule and updates it. Returns the server's representation of the mySQLBackupSchedule, and an error, if there is any.
func (c *FakeMySQLBackupSchedules) Update(ctx context.Context, mySQLBackupSchedule *mysqlv1.MySQLBackupSchedule, opts v1.UpdateOptions) (result *mysqlv1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlbackupschedulesResource, c.ns, mySQLBackupSchedule), &mysqlv1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackupSchedule), err
}

// Delete takes name of the mySQLBackupSchedule and deletes it. Returns an error if one occurs.
func (c *FakeMySQLBackupSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlbackupschedulesResource, c.ns, name), &mysqlv1.MySQLBackupSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLBackupSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlbackupschedulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &mysqlv1.MySQLBackupScheduleList{})
	return err
}

// Patch applies the patch and returns the patched mySQLBackupSchedule.
func (c *FakeMySQLBackupSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *mysqlv1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlbackupschedulesResource, c.ns, name, pt, data, subresources...), &mysqlv1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlv1.MySQLBackupSchedule), err
}
