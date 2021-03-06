/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBcsSecrets implements BcsSecretInterface
type FakeBcsSecrets struct {
	Fake *FakeBkbcsV2
	ns   string
}

var bcssecretsResource = schema.GroupVersionResource{Group: "bkbcs.tencent.com", Version: "v2", Resource: "bcssecrets"}

var bcssecretsKind = schema.GroupVersionKind{Group: "bkbcs.tencent.com", Version: "v2", Kind: "BcsSecret"}

// Get takes name of the bcsSecret, and returns the corresponding bcsSecret object, and an error if there is any.
func (c *FakeBcsSecrets) Get(name string, options v1.GetOptions) (result *v2.BcsSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bcssecretsResource, c.ns, name), &v2.BcsSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsSecret), err
}

// List takes label and field selectors, and returns the list of BcsSecrets that match those selectors.
func (c *FakeBcsSecrets) List(opts v1.ListOptions) (result *v2.BcsSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bcssecretsResource, bcssecretsKind, c.ns, opts), &v2.BcsSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.BcsSecretList{ListMeta: obj.(*v2.BcsSecretList).ListMeta}
	for _, item := range obj.(*v2.BcsSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bcsSecrets.
func (c *FakeBcsSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bcssecretsResource, c.ns, opts))

}

// Create takes the representation of a bcsSecret and creates it.  Returns the server's representation of the bcsSecret, and an error, if there is any.
func (c *FakeBcsSecrets) Create(bcsSecret *v2.BcsSecret) (result *v2.BcsSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bcssecretsResource, c.ns, bcsSecret), &v2.BcsSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsSecret), err
}

// Update takes the representation of a bcsSecret and updates it. Returns the server's representation of the bcsSecret, and an error, if there is any.
func (c *FakeBcsSecrets) Update(bcsSecret *v2.BcsSecret) (result *v2.BcsSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bcssecretsResource, c.ns, bcsSecret), &v2.BcsSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBcsSecrets) UpdateStatus(bcsSecret *v2.BcsSecret) (*v2.BcsSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bcssecretsResource, "status", c.ns, bcsSecret), &v2.BcsSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsSecret), err
}

// Delete takes name of the bcsSecret and deletes it. Returns an error if one occurs.
func (c *FakeBcsSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bcssecretsResource, c.ns, name), &v2.BcsSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBcsSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bcssecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.BcsSecretList{})
	return err
}

// Patch applies the patch and returns the patched bcsSecret.
func (c *FakeBcsSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bcssecretsResource, c.ns, name, data, subresources...), &v2.BcsSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsSecret), err
}
