/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package fake

import (
	"context"

	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQuotas implements QuotaInterface
type FakeQuotas struct {
	Fake *FakeCore
	ns   string
}

var quotasResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "", Resource: "quotas"}

var quotasKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "", Kind: "Quota"}

// Get takes name of the quota, and returns the corresponding quota object, and an error if there is any.
func (c *FakeQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *core.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(quotasResource, c.ns, name), &core.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Quota), err
}

// List takes label and field selectors, and returns the list of Quotas that match those selectors.
func (c *FakeQuotas) List(ctx context.Context, opts v1.ListOptions) (result *core.QuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(quotasResource, quotasKind, c.ns, opts), &core.QuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.QuotaList{ListMeta: obj.(*core.QuotaList).ListMeta}
	for _, item := range obj.(*core.QuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested quotas.
func (c *FakeQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(quotasResource, c.ns, opts))

}

// Create takes the representation of a quota and creates it.  Returns the server's representation of the quota, and an error, if there is any.
func (c *FakeQuotas) Create(ctx context.Context, quota *core.Quota, opts v1.CreateOptions) (result *core.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(quotasResource, c.ns, quota), &core.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Quota), err
}

// Update takes the representation of a quota and updates it. Returns the server's representation of the quota, and an error, if there is any.
func (c *FakeQuotas) Update(ctx context.Context, quota *core.Quota, opts v1.UpdateOptions) (result *core.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(quotasResource, c.ns, quota), &core.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Quota), err
}

// Delete takes name of the quota and deletes it. Returns an error if one occurs.
func (c *FakeQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(quotasResource, c.ns, name), &core.Quota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(quotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &core.QuotaList{})
	return err
}

// Patch applies the patch and returns the patched quota.
func (c *FakeQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *core.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(quotasResource, c.ns, name, pt, data, subresources...), &core.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Quota), err
}
