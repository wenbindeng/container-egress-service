/*
Copyright 2021 The Kube-OVN CES Controller Authors.

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
	"context"
	"time"

	v1alpha1 "github.com/kubeovn/ces-controller/pkg/apis/kubeovn.io/v1alpha1"
	scheme "github.com/kubeovn/ces-controller/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NamespaceEgressRulesGetter has a method to return a NamespaceEgressRuleInterface.
// A group's client should implement this interface.
type NamespaceEgressRulesGetter interface {
	NamespaceEgressRules(namespace string) NamespaceEgressRuleInterface
}

// NamespaceEgressRuleInterface has methods to work with NamespaceEgressRule resources.
type NamespaceEgressRuleInterface interface {
	Create(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.CreateOptions) (*v1alpha1.NamespaceEgressRule, error)
	Update(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.UpdateOptions) (*v1alpha1.NamespaceEgressRule, error)
	UpdateStatus(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.UpdateOptions) (*v1alpha1.NamespaceEgressRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NamespaceEgressRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NamespaceEgressRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NamespaceEgressRule, err error)
	NamespaceEgressRuleExpansion
}

// namespaceEgressRules implements NamespaceEgressRuleInterface
type namespaceEgressRules struct {
	client rest.Interface
	ns     string
}

// newNamespaceEgressRules returns a NamespaceEgressRules
func newNamespaceEgressRules(c *KubeovnV1alpha1Client, namespace string) *namespaceEgressRules {
	return &namespaceEgressRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the namespaceEgressRule, and returns the corresponding namespaceEgressRule object, and an error if there is any.
func (c *namespaceEgressRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NamespaceEgressRule, err error) {
	result = &v1alpha1.NamespaceEgressRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NamespaceEgressRules that match those selectors.
func (c *namespaceEgressRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NamespaceEgressRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NamespaceEgressRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested namespaceEgressRules.
func (c *namespaceEgressRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a namespaceEgressRule and creates it.  Returns the server's representation of the namespaceEgressRule, and an error, if there is any.
func (c *namespaceEgressRules) Create(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.CreateOptions) (result *v1alpha1.NamespaceEgressRule, err error) {
	result = &v1alpha1.NamespaceEgressRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceEgressRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a namespaceEgressRule and updates it. Returns the server's representation of the namespaceEgressRule, and an error, if there is any.
func (c *namespaceEgressRules) Update(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.UpdateOptions) (result *v1alpha1.NamespaceEgressRule, err error) {
	result = &v1alpha1.NamespaceEgressRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		Name(namespaceEgressRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceEgressRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *namespaceEgressRules) UpdateStatus(ctx context.Context, namespaceEgressRule *v1alpha1.NamespaceEgressRule, opts v1.UpdateOptions) (result *v1alpha1.NamespaceEgressRule, err error) {
	result = &v1alpha1.NamespaceEgressRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		Name(namespaceEgressRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceEgressRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the namespaceEgressRule and deletes it. Returns an error if one occurs.
func (c *namespaceEgressRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *namespaceEgressRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namespaceegressrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched namespaceEgressRule.
func (c *namespaceEgressRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NamespaceEgressRule, err error) {
	result = &v1alpha1.NamespaceEgressRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("namespaceegressrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}