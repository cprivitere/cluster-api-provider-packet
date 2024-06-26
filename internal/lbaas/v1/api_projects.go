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

/*
Load Balancer Management API

Load Balancer Management API is an API for managing load balancers.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ProjectsApiService ProjectsApi service
type ProjectsApiService service

type ApiCreateLoadBalancerRequest struct {
	ctx                context.Context
	ApiService         *ProjectsApiService
	projectID          string
	loadBalancerCreate *LoadBalancerCreate
}

func (r ApiCreateLoadBalancerRequest) LoadBalancerCreate(loadBalancerCreate LoadBalancerCreate) ApiCreateLoadBalancerRequest {
	r.loadBalancerCreate = &loadBalancerCreate
	return r
}

func (r ApiCreateLoadBalancerRequest) Execute() (*ResourceCreatedResponse, *http.Response, error) {
	return r.ApiService.CreateLoadBalancerExecute(r)
}

/*
CreateLoadBalancer Create a load balancer for project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectID ID of project
	@return ApiCreateLoadBalancerRequest
*/
func (a *ProjectsApiService) CreateLoadBalancer(ctx context.Context, projectID string) ApiCreateLoadBalancerRequest {
	return ApiCreateLoadBalancerRequest{
		ApiService: a,
		ctx:        ctx,
		projectID:  projectID,
	}
}

// Execute executes the request
//
//	@return ResourceCreatedResponse
func (a *ProjectsApiService) CreateLoadBalancerExecute(r ApiCreateLoadBalancerRequest) (*ResourceCreatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceCreatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreateLoadBalancer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectID}/loadbalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"projectID"+"}", url.PathEscape(parameterValueToString(r.projectID, "projectID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.loadBalancerCreate == nil {
		return localVarReturnValue, nil, reportError("loadBalancerCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.loadBalancerCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePoolRequest struct {
	ctx                    context.Context
	ApiService             *ProjectsApiService
	projectID              string
	loadBalancerPoolCreate *LoadBalancerPoolCreate
}

func (r ApiCreatePoolRequest) LoadBalancerPoolCreate(loadBalancerPoolCreate LoadBalancerPoolCreate) ApiCreatePoolRequest {
	r.loadBalancerPoolCreate = &loadBalancerPoolCreate
	return r
}

func (r ApiCreatePoolRequest) Execute() (*ResourceCreatedResponse, *http.Response, error) {
	return r.ApiService.CreatePoolExecute(r)
}

/*
CreatePool Create a load balancer pool for project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectID ID of project
	@return ApiCreatePoolRequest
*/
func (a *ProjectsApiService) CreatePool(ctx context.Context, projectID string) ApiCreatePoolRequest {
	return ApiCreatePoolRequest{
		ApiService: a,
		ctx:        ctx,
		projectID:  projectID,
	}
}

// Execute executes the request
//
//	@return ResourceCreatedResponse
func (a *ProjectsApiService) CreatePoolExecute(r ApiCreatePoolRequest) (*ResourceCreatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceCreatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreatePool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectID}/loadbalancers/pools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectID"+"}", url.PathEscape(parameterValueToString(r.projectID, "projectID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.loadBalancerPoolCreate == nil {
		return localVarReturnValue, nil, reportError("loadBalancerPoolCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.loadBalancerPoolCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListLoadBalancersRequest struct {
	ctx        context.Context
	ApiService *ProjectsApiService
	projectID  string
}

func (r ApiListLoadBalancersRequest) Execute() (*LoadBalancerCollection, *http.Response, error) {
	return r.ApiService.ListLoadBalancersExecute(r)
}

/*
ListLoadBalancers Gets the load balancers for a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectID ID of project
	@return ApiListLoadBalancersRequest
*/
func (a *ProjectsApiService) ListLoadBalancers(ctx context.Context, projectID string) ApiListLoadBalancersRequest {
	return ApiListLoadBalancersRequest{
		ApiService: a,
		ctx:        ctx,
		projectID:  projectID,
	}
}

// Execute executes the request
//
//	@return LoadBalancerCollection
func (a *ProjectsApiService) ListLoadBalancersExecute(r ApiListLoadBalancersRequest) (*LoadBalancerCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LoadBalancerCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListLoadBalancers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectID}/loadbalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"projectID"+"}", url.PathEscape(parameterValueToString(r.projectID, "projectID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPoolsRequest struct {
	ctx        context.Context
	ApiService *ProjectsApiService
	projectID  string
}

func (r ApiListPoolsRequest) Execute() (*LoadBalancerPoolCollection, *http.Response, error) {
	return r.ApiService.ListPoolsExecute(r)
}

/*
ListPools Gets the pools for a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectID ID of project
	@return ApiListPoolsRequest
*/
func (a *ProjectsApiService) ListPools(ctx context.Context, projectID string) ApiListPoolsRequest {
	return ApiListPoolsRequest{
		ApiService: a,
		ctx:        ctx,
		projectID:  projectID,
	}
}

// Execute executes the request
//
//	@return LoadBalancerPoolCollection
func (a *ProjectsApiService) ListPoolsExecute(r ApiListPoolsRequest) (*LoadBalancerPoolCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LoadBalancerPoolCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListPools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectID}/loadbalancers/pools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectID"+"}", url.PathEscape(parameterValueToString(r.projectID, "projectID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
