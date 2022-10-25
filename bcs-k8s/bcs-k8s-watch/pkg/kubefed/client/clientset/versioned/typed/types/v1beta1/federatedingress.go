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

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedIngressesGetter has a method to return a FederatedIngressInterface.
// A group's client should implement this interface.
type FederatedIngressesGetter interface {
	FederatedIngresses(namespace string) FederatedIngressInterface
}

// FederatedIngressInterface has methods to work with FederatedIngress resources.
type FederatedIngressInterface interface {
	Create(*v1beta1.FederatedIngress) (*v1beta1.FederatedIngress, error)
	Update(*v1beta1.FederatedIngress) (*v1beta1.FederatedIngress, error)
	UpdateStatus(*v1beta1.FederatedIngress) (*v1beta1.FederatedIngress, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedIngress, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedIngressList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedIngress, err error)
	FederatedIngressExpansion
}

// federatedIngresses implements FederatedIngressInterface
type federatedIngresses struct {
	client rest.Interface
	ns     string
}

// newFederatedIngresses returns a FederatedIngresses
func newFederatedIngresses(c *TypesV1beta1Client, namespace string) *federatedIngresses {
	return &federatedIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedIngress, and returns the corresponding federatedIngress object, and an error if there is any.
func (c *federatedIngresses) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedIngresses that match those selectors.
func (c *federatedIngresses) List(opts v1.ListOptions) (result *v1beta1.FederatedIngressList, err error) {
	result = &v1beta1.FederatedIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedIngresses.
func (c *federatedIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedIngress and creates it.  Returns the server's representation of the federatedIngress, and an error, if there is any.
func (c *federatedIngresses) Create(federatedIngress *v1beta1.FederatedIngress) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedingresses").
		Body(federatedIngress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedIngress and updates it. Returns the server's representation of the federatedIngress, and an error, if there is any.
func (c *federatedIngresses) Update(federatedIngress *v1beta1.FederatedIngress) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(federatedIngress.Name).
		Body(federatedIngress).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedIngresses) UpdateStatus(federatedIngress *v1beta1.FederatedIngress) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(federatedIngress.Name).
		SubResource("status").
		Body(federatedIngress).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedIngress and deletes it. Returns an error if one occurs.
func (c *federatedIngresses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedingresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedIngress.
func (c *federatedIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedIngress, err error) {
	result = &v1beta1.FederatedIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}