/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package exampleservicev1 : Operations and models for the ExampleServiceV1 service
package exampleservicev1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.com/powellquiring/my-sdk/common"
)

// ExampleServiceV1 : No description provided (generated by Openapi Generator
// https://github.com/openapitools/openapi-generator)
//
// Version: 1.0.0
type ExampleServiceV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "http://cloud.ibm.com/mysdk/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "example_service"

// ExampleServiceV1Options : Service options
type ExampleServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewExampleServiceV1UsingExternalConfig : constructs an instance of ExampleServiceV1 with passed in options and external configuration.
func NewExampleServiceV1UsingExternalConfig(options *ExampleServiceV1Options) (exampleService *ExampleServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	exampleService, err = NewExampleServiceV1(options)
	if err != nil {
		return
	}

	err = exampleService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = exampleService.Service.SetServiceURL(options.URL)
	}
	return
}

// NewExampleServiceV1 : constructs an instance of ExampleServiceV1 with passed in options.
func NewExampleServiceV1(options *ExampleServiceV1Options) (service *ExampleServiceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &ExampleServiceV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (exampleService *ExampleServiceV1) SetServiceURL(url string) error {
	return exampleService.Service.SetServiceURL(url)
}

// ListResources : List all resources
func (exampleService *ExampleServiceV1) ListResources(listResourcesOptions *ListResourcesOptions) (result *Resources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listResourcesOptions, "listResourcesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resources"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(exampleService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourcesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "ListResources")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listResourcesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listResourcesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = exampleService.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResources(m)
		response.Result = result
	}

	return
}

// CreateResource : Create a resource
func (exampleService *ExampleServiceV1) CreateResource(createResourceOptions *CreateResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceOptions, "createResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceOptions, "createResourceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resources"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(exampleService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "CreateResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createResourceOptions.ResourceID != nil {
		body["resource_id"] = createResourceOptions.ResourceID
	}
	if createResourceOptions.Name != nil {
		body["name"] = createResourceOptions.Name
	}
	if createResourceOptions.Tag != nil {
		body["tag"] = createResourceOptions.Tag
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = exampleService.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResource(m)
		response.Result = result
	}

	return
}

// GetResource : Info for a specific resource
func (exampleService *ExampleServiceV1) GetResource(getResourceOptions *GetResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceOptions, "getResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceOptions, "getResourceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resources"}
	pathParameters := []string{*getResourceOptions.ResourceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(exampleService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("example_service", "V1", "GetResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = exampleService.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResource(m)
		response.Result = result
	}

	return
}

// CreateResourceOptions : The CreateResource options.
type CreateResourceOptions struct {
	// The id of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A tag value for the resource.
	Tag *string `json:"tag,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateResourceOptions : Instantiate CreateResourceOptions
func (*ExampleServiceV1) NewCreateResourceOptions(resourceID string, name string) *CreateResourceOptions {
	return &CreateResourceOptions{
		ResourceID: core.StringPtr(resourceID),
		Name: core.StringPtr(name),
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *CreateResourceOptions) SetResourceID(resourceID string) *CreateResourceOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateResourceOptions) SetName(name string) *CreateResourceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTag : Allow user to set Tag
func (options *CreateResourceOptions) SetTag(tag string) *CreateResourceOptions {
	options.Tag = core.StringPtr(tag)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceOptions) SetHeaders(param map[string]string) *CreateResourceOptions {
	options.Headers = param
	return options
}

// GetResourceOptions : The GetResource options.
type GetResourceOptions struct {
	// The id of the resource to retrieve.
	ResourceID *string `json:"resource_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceOptions : Instantiate GetResourceOptions
func (*ExampleServiceV1) NewGetResourceOptions(resourceID string) *GetResourceOptions {
	return &GetResourceOptions{
		ResourceID: core.StringPtr(resourceID),
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *GetResourceOptions) SetResourceID(resourceID string) *GetResourceOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceOptions) SetHeaders(param map[string]string) *GetResourceOptions {
	options.Headers = param
	return options
}

// ListResourcesOptions : The ListResources options.
type ListResourcesOptions struct {
	// How many items to return at one time (max 100).
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourcesOptions : Instantiate ListResourcesOptions
func (*ExampleServiceV1) NewListResourcesOptions() *ListResourcesOptions {
	return &ListResourcesOptions{}
}

// SetLimit : Allow user to set Limit
func (options *ListResourcesOptions) SetLimit(limit int64) *ListResourcesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourcesOptions) SetHeaders(param map[string]string) *ListResourcesOptions {
	options.Headers = param
	return options
}

// Resource : A resource.
type Resource struct {
	// The id of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A tag value for the resource.
	Tag *string `json:"tag,omitempty"`

	// This is a read only string.
	ReadOnly *string `json:"read_only,omitempty"`
}


// NewResource : Instantiate Resource (Generic Model Constructor)
func (*ExampleServiceV1) NewResource(resourceID string, name string) (model *Resource, err error) {
	model = &Resource{
		ResourceID: core.StringPtr(resourceID),
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResource constructs an instance of Resource from the specified map.
func UnmarshalResource(m map[string]interface{}) (result *Resource, err error) {
	obj := new(Resource)
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Tag, err = core.UnmarshalString(m, "tag")
	if err != nil {
		return
	}
	obj.ReadOnly, err = core.UnmarshalString(m, "read_only")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceSlice unmarshals a slice of Resource instances from the specified list of maps.
func UnmarshalResourceSlice(s []interface{}) (slice []Resource, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resource'")
			return
		}
		obj, e := UnmarshalResource(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceAsProperty unmarshals an instance of Resource that is stored as a property
// within the specified map.
func UnmarshalResourceAsProperty(m map[string]interface{}, propertyName string) (result *Resource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resource'", propertyName)
			return
		}
		result, err = UnmarshalResource(objMap)
	}
	return
}

// UnmarshalResourceSliceAsProperty unmarshals a slice of Resource instances that are stored as a property
// within the specified map.
func UnmarshalResourceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resource'", propertyName)
			return
		}
		slice, err = UnmarshalResourceSlice(vSlice)
	}
	return
}

// Resources : List of resources.
type Resources struct {
	// Offset value for this portion of the resource list.
	Offset *int64 `json:"offset,omitempty"`

	// Limit value specified or defaulted in the list_resources request.
	Limit *int64 `json:"limit,omitempty"`

	// A list of resources.
	Resources []Resource `json:"resources,omitempty"`
}


// UnmarshalResources constructs an instance of Resources from the specified map.
func UnmarshalResources(m map[string]interface{}) (result *Resources, err error) {
	obj := new(Resources)
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalResourceSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourcesSlice unmarshals a slice of Resources instances from the specified list of maps.
func UnmarshalResourcesSlice(s []interface{}) (slice []Resources, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resources'")
			return
		}
		obj, e := UnmarshalResources(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourcesAsProperty unmarshals an instance of Resources that is stored as a property
// within the specified map.
func UnmarshalResourcesAsProperty(m map[string]interface{}, propertyName string) (result *Resources, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resources'", propertyName)
			return
		}
		result, err = UnmarshalResources(objMap)
	}
	return
}

// UnmarshalResourcesSliceAsProperty unmarshals a slice of Resources instances that are stored as a property
// within the specified map.
func UnmarshalResourcesSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resources, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resources'", propertyName)
			return
		}
		slice, err = UnmarshalResourcesSlice(vSlice)
	}
	return
}
