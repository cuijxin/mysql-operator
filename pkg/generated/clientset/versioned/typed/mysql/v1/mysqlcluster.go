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

package v1

import (
	"context"
	"time"

	v1 "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	scheme "github.com/cuijxin/mysql-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MySQLClustersGetter has a method to return a MySQLClusterInterface.
// A group's client should implement this interface.
type MySQLClustersGetter interface {
	MySQLClusters(namespace string) MySQLClusterInterface
}

// MySQLClusterInterface has methods to work with MySQLCluster resources.
type MySQLClusterInterface interface {
	Create(ctx context.Context, mySQLCluster *v1.MySQLCluster, opts metav1.CreateOptions) (*v1.MySQLCluster, error)
	Update(ctx context.Context, mySQLCluster *v1.MySQLCluster, opts metav1.UpdateOptions) (*v1.MySQLCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MySQLCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MySQLClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MySQLCluster, err error)
	MySQLClusterExpansion
}

// mySQLClusters implements MySQLClusterInterface
type mySQLClusters struct {
	client rest.Interface
	ns     string
}

// newMySQLClusters returns a MySQLClusters
func newMySQLClusters(c *MysqlV1Client, namespace string) *mySQLClusters {
	return &mySQLClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mySQLCluster, and returns the corresponding mySQLCluster object, and an error if there is any.
func (c *mySQLClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MySQLCluster, err error) {
	result = &v1.MySQLCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysql5clusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MySQLClusters that match those selectors.
func (c *mySQLClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MySQLClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MySQLClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysql5clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mySQLClusters.
func (c *mySQLClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysql5clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mySQLCluster and creates it.  Returns the server's representation of the mySQLCluster, and an error, if there is any.
func (c *mySQLClusters) Create(ctx context.Context, mySQLCluster *v1.MySQLCluster, opts metav1.CreateOptions) (result *v1.MySQLCluster, err error) {
	result = &v1.MySQLCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysql5clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mySQLCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mySQLCluster and updates it. Returns the server's representation of the mySQLCluster, and an error, if there is any.
func (c *mySQLClusters) Update(ctx context.Context, mySQLCluster *v1.MySQLCluster, opts metav1.UpdateOptions) (result *v1.MySQLCluster, err error) {
	result = &v1.MySQLCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysql5clusters").
		Name(mySQLCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mySQLCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mySQLCluster and deletes it. Returns an error if one occurs.
func (c *mySQLClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysql5clusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mySQLClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysql5clusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mySQLCluster.
func (c *mySQLClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MySQLCluster, err error) {
	result = &v1.MySQLCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysql5clusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
