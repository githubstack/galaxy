/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "git.code.oa.com/gaia/tapp-controller/pkg/apis/tappcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTApps implements TAppInterface
type FakeTApps struct {
	Fake *FakeTappcontrollerV1alpha1
	ns   string
}

var tappsResource = schema.GroupVersionResource{Group: "tappcontroller.k8s.io", Version: "v1alpha1", Resource: "tapps"}

var tappsKind = schema.GroupVersionKind{Group: "tappcontroller.k8s.io", Version: "v1alpha1", Kind: "TApp"}

// Get takes name of the tApp, and returns the corresponding tApp object, and an error if there is any.
func (c *FakeTApps) Get(name string, options v1.GetOptions) (result *v1alpha1.TApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tappsResource, c.ns, name), &v1alpha1.TApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TApp), err
}

// List takes label and field selectors, and returns the list of TApps that match those selectors.
func (c *FakeTApps) List(opts v1.ListOptions) (result *v1alpha1.TAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tappsResource, tappsKind, c.ns, opts), &v1alpha1.TAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TAppList{}
	for _, item := range obj.(*v1alpha1.TAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tApps.
func (c *FakeTApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tappsResource, c.ns, opts))

}

// Create takes the representation of a tApp and creates it.  Returns the server's representation of the tApp, and an error, if there is any.
func (c *FakeTApps) Create(tApp *v1alpha1.TApp) (result *v1alpha1.TApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tappsResource, c.ns, tApp), &v1alpha1.TApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TApp), err
}

// Update takes the representation of a tApp and updates it. Returns the server's representation of the tApp, and an error, if there is any.
func (c *FakeTApps) Update(tApp *v1alpha1.TApp) (result *v1alpha1.TApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tappsResource, c.ns, tApp), &v1alpha1.TApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TApp), err
}

// Delete takes name of the tApp and deletes it. Returns an error if one occurs.
func (c *FakeTApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tappsResource, c.ns, name), &v1alpha1.TApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TAppList{})
	return err
}

// Patch applies the patch and returns the patched tApp.
func (c *FakeTApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tappsResource, c.ns, name, data, subresources...), &v1alpha1.TApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TApp), err
}