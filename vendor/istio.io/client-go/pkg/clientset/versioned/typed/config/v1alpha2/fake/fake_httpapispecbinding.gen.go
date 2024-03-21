// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPAPISpecBindings implements HTTPAPISpecBindingInterface
type FakeHTTPAPISpecBindings struct {
	Fake *FakeConfigV1alpha2
	ns   string
}

var httpapispecbindingsResource = schema.GroupVersionResource{Group: "config.istio.io", Version: "v1alpha2", Resource: "httpapispecbindings"}

var httpapispecbindingsKind = schema.GroupVersionKind{Group: "config.istio.io", Version: "v1alpha2", Kind: "HTTPAPISpecBinding"}

// Get takes name of the hTTPAPISpecBinding, and returns the corresponding hTTPAPISpecBinding object, and an error if there is any.
func (c *FakeHTTPAPISpecBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httpapispecbindingsResource, c.ns, name), &v1alpha2.HTTPAPISpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpecBinding), err
}

// List takes label and field selectors, and returns the list of HTTPAPISpecBindings that match those selectors.
func (c *FakeHTTPAPISpecBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httpapispecbindingsResource, httpapispecbindingsKind, c.ns, opts), &v1alpha2.HTTPAPISpecBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HTTPAPISpecBindingList{ListMeta: obj.(*v1alpha2.HTTPAPISpecBindingList).ListMeta}
	for _, item := range obj.(*v1alpha2.HTTPAPISpecBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecBindings.
func (c *FakeHTTPAPISpecBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httpapispecbindingsResource, c.ns, opts))

}

// Create takes the representation of a hTTPAPISpecBinding and creates it.  Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *FakeHTTPAPISpecBindings) Create(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.CreateOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httpapispecbindingsResource, c.ns, hTTPAPISpecBinding), &v1alpha2.HTTPAPISpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpecBinding), err
}

// Update takes the representation of a hTTPAPISpecBinding and updates it. Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *FakeHTTPAPISpecBindings) Update(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.UpdateOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httpapispecbindingsResource, c.ns, hTTPAPISpecBinding), &v1alpha2.HTTPAPISpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpecBinding), err
}

// Delete takes name of the hTTPAPISpecBinding and deletes it. Returns an error if one occurs.
func (c *FakeHTTPAPISpecBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httpapispecbindingsResource, c.ns, name), &v1alpha2.HTTPAPISpecBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPAPISpecBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httpapispecbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.HTTPAPISpecBindingList{})
	return err
}

// Patch applies the patch and returns the patched hTTPAPISpecBinding.
func (c *FakeHTTPAPISpecBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httpapispecbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha2.HTTPAPISpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpecBinding), err
}
