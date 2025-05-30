/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKontainerDrivers implements KontainerDriverInterface
type FakeKontainerDrivers struct {
	Fake *FakeManagementV3
}

var kontainerdriversResource = v3.SchemeGroupVersion.WithResource("kontainerdrivers")

var kontainerdriversKind = v3.SchemeGroupVersion.WithKind("KontainerDriver")

// Get takes name of the kontainerDriver, and returns the corresponding kontainerDriver object, and an error if there is any.
func (c *FakeKontainerDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.KontainerDriver, err error) {
	emptyResult := &v3.KontainerDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(kontainerdriversResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.KontainerDriver), err
}

// List takes label and field selectors, and returns the list of KontainerDrivers that match those selectors.
func (c *FakeKontainerDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v3.KontainerDriverList, err error) {
	emptyResult := &v3.KontainerDriverList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(kontainerdriversResource, kontainerdriversKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.KontainerDriverList{ListMeta: obj.(*v3.KontainerDriverList).ListMeta}
	for _, item := range obj.(*v3.KontainerDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kontainerDrivers.
func (c *FakeKontainerDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(kontainerdriversResource, opts))
}

// Create takes the representation of a kontainerDriver and creates it.  Returns the server's representation of the kontainerDriver, and an error, if there is any.
func (c *FakeKontainerDrivers) Create(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.CreateOptions) (result *v3.KontainerDriver, err error) {
	emptyResult := &v3.KontainerDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(kontainerdriversResource, kontainerDriver, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.KontainerDriver), err
}

// Update takes the representation of a kontainerDriver and updates it. Returns the server's representation of the kontainerDriver, and an error, if there is any.
func (c *FakeKontainerDrivers) Update(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (result *v3.KontainerDriver, err error) {
	emptyResult := &v3.KontainerDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(kontainerdriversResource, kontainerDriver, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.KontainerDriver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKontainerDrivers) UpdateStatus(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (result *v3.KontainerDriver, err error) {
	emptyResult := &v3.KontainerDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(kontainerdriversResource, "status", kontainerDriver, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.KontainerDriver), err
}

// Delete takes name of the kontainerDriver and deletes it. Returns an error if one occurs.
func (c *FakeKontainerDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kontainerdriversResource, name, opts), &v3.KontainerDriver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKontainerDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(kontainerdriversResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.KontainerDriverList{})
	return err
}

// Patch applies the patch and returns the patched kontainerDriver.
func (c *FakeKontainerDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.KontainerDriver, err error) {
	emptyResult := &v3.KontainerDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(kontainerdriversResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.KontainerDriver), err
}
