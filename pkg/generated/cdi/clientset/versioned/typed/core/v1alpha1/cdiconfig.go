/*
Copyright 2020 The vm import Authors.

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

	scheme "github.com/kubevirt/vm-import-operator/pkg/generated/cdi/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

// CDIConfigsGetter has a method to return a CDIConfigInterface.
// A group's client should implement this interface.
type CDIConfigsGetter interface {
	CDIConfigs() CDIConfigInterface
}

// CDIConfigInterface has methods to work with CDIConfig resources.
type CDIConfigInterface interface {
	Create(*v1alpha1.CDIConfig) (*v1alpha1.CDIConfig, error)
	Update(*v1alpha1.CDIConfig) (*v1alpha1.CDIConfig, error)
	UpdateStatus(*v1alpha1.CDIConfig) (*v1alpha1.CDIConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CDIConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.CDIConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CDIConfig, err error)
	CDIConfigExpansion
}

// cDIConfigs implements CDIConfigInterface
type cDIConfigs struct {
	client rest.Interface
}

// newCDIConfigs returns a CDIConfigs
func newCDIConfigs(c *CdiV1alpha1Client) *cDIConfigs {
	return &cDIConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the cDIConfig, and returns the corresponding cDIConfig object, and an error if there is any.
func (c *cDIConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CDIConfig, err error) {
	result = &v1alpha1.CDIConfig{}
	err = c.client.Get().
		Resource("cdiconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CDIConfigs that match those selectors.
func (c *cDIConfigs) List(opts v1.ListOptions) (result *v1alpha1.CDIConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CDIConfigList{}
	err = c.client.Get().
		Resource("cdiconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cDIConfigs.
func (c *cDIConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cdiconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cDIConfig and creates it.  Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *cDIConfigs) Create(cDIConfig *v1alpha1.CDIConfig) (result *v1alpha1.CDIConfig, err error) {
	result = &v1alpha1.CDIConfig{}
	err = c.client.Post().
		Resource("cdiconfigs").
		Body(cDIConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cDIConfig and updates it. Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *cDIConfigs) Update(cDIConfig *v1alpha1.CDIConfig) (result *v1alpha1.CDIConfig, err error) {
	result = &v1alpha1.CDIConfig{}
	err = c.client.Put().
		Resource("cdiconfigs").
		Name(cDIConfig.Name).
		Body(cDIConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cDIConfigs) UpdateStatus(cDIConfig *v1alpha1.CDIConfig) (result *v1alpha1.CDIConfig, err error) {
	result = &v1alpha1.CDIConfig{}
	err = c.client.Put().
		Resource("cdiconfigs").
		Name(cDIConfig.Name).
		SubResource("status").
		Body(cDIConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the cDIConfig and deletes it. Returns an error if one occurs.
func (c *cDIConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cdiconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cDIConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cdiconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cDIConfig.
func (c *cDIConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CDIConfig, err error) {
	result = &v1alpha1.CDIConfig{}
	err = c.client.Patch(pt).
		Resource("cdiconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
