/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "git.code.oa.com/gaiastack/galaxy/pkg/ipam/apis/floatip/v1alpha1"
	scheme "git.code.oa.com/gaiastack/galaxy/pkg/ipam/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FloatIpsGetter has a method to return a FloatIpInterface.
// A group's client should implement this interface.
type FloatIpsGetter interface {
	FloatIps(namespace string) FloatIpInterface
}

// FloatIpInterface has methods to work with FloatIp resources.
type FloatIpInterface interface {
	Create(*v1alpha1.FloatIp) (*v1alpha1.FloatIp, error)
	Update(*v1alpha1.FloatIp) (*v1alpha1.FloatIp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FloatIp, error)
	List(opts v1.ListOptions) (*v1alpha1.FloatIpList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatIp, err error)
	FloatIpExpansion
}

// floatIps implements FloatIpInterface
type floatIps struct {
	client rest.Interface
	ns     string
}

// newFloatIps returns a FloatIps
func newFloatIps(c *GalaxyV1alpha1Client, namespace string) *floatIps {
	return &floatIps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the floatIp, and returns the corresponding floatIp object, and an error if there is any.
func (c *floatIps) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatIp, err error) {
	result = &v1alpha1.FloatIp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("floatips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FloatIps that match those selectors.
func (c *floatIps) List(opts v1.ListOptions) (result *v1alpha1.FloatIpList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FloatIpList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("floatips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested floatIps.
func (c *floatIps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("floatips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a floatIp and creates it.  Returns the server's representation of the floatIp, and an error, if there is any.
func (c *floatIps) Create(floatIp *v1alpha1.FloatIp) (result *v1alpha1.FloatIp, err error) {
	result = &v1alpha1.FloatIp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("floatips").
		Body(floatIp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a floatIp and updates it. Returns the server's representation of the floatIp, and an error, if there is any.
func (c *floatIps) Update(floatIp *v1alpha1.FloatIp) (result *v1alpha1.FloatIp, err error) {
	result = &v1alpha1.FloatIp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("floatips").
		Name(floatIp.Name).
		Body(floatIp).
		Do().
		Into(result)
	return
}

// Delete takes name of the floatIp and deletes it. Returns an error if one occurs.
func (c *floatIps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("floatips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *floatIps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("floatips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched floatIp.
func (c *floatIps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatIp, err error) {
	result = &v1alpha1.FloatIp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("floatips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}