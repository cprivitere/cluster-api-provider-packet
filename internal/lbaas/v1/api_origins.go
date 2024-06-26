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

// OriginsApiService OriginsApi service
type OriginsApiService service

type ApiDeleteLoadBalancerOriginRequest struct {
	ctx                  context.Context
	ApiService           *OriginsApiService
	loadBalancerOriginID string
}

func (r ApiDeleteLoadBalancerOriginRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLoadBalancerOriginExecute(r)
}

/*
DeleteLoadBalancerOrigin Delete a load balancer origin.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param loadBalancerOriginID ID of the load balancer origin
	@return ApiDeleteLoadBalancerOriginRequest
*/
func (a *OriginsApiService) DeleteLoadBalancerOrigin(ctx context.Context, loadBalancerOriginID string) ApiDeleteLoadBalancerOriginRequest {
	return ApiDeleteLoadBalancerOriginRequest{
		ApiService:           a,
		ctx:                  ctx,
		loadBalancerOriginID: loadBalancerOriginID,
	}
}

// Execute executes the request
func (a *OriginsApiService) DeleteLoadBalancerOriginExecute(r ApiDeleteLoadBalancerOriginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginsApiService.DeleteLoadBalancerOrigin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/loadbalancers/pools/origins/{loadBalancerOriginID}"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerOriginID"+"}", url.PathEscape(parameterValueToString(r.loadBalancerOriginID, "loadBalancerOriginID")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetLoadBalancerOriginRequest struct {
	ctx                  context.Context
	ApiService           *OriginsApiService
	loadBalancerOriginID string
}

func (r ApiGetLoadBalancerOriginRequest) Execute() (*LoadBalancerPoolOrigin, *http.Response, error) {
	return r.ApiService.GetLoadBalancerOriginExecute(r)
}

/*
GetLoadBalancerOrigin Gets a load balancer origin by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param loadBalancerOriginID ID of the load balancer origin
	@return ApiGetLoadBalancerOriginRequest
*/
func (a *OriginsApiService) GetLoadBalancerOrigin(ctx context.Context, loadBalancerOriginID string) ApiGetLoadBalancerOriginRequest {
	return ApiGetLoadBalancerOriginRequest{
		ApiService:           a,
		ctx:                  ctx,
		loadBalancerOriginID: loadBalancerOriginID,
	}
}

// Execute executes the request
//
//	@return LoadBalancerPoolOrigin
func (a *OriginsApiService) GetLoadBalancerOriginExecute(r ApiGetLoadBalancerOriginRequest) (*LoadBalancerPoolOrigin, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LoadBalancerPoolOrigin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginsApiService.GetLoadBalancerOrigin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/loadbalancers/pools/origins/{loadBalancerOriginID}"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerOriginID"+"}", url.PathEscape(parameterValueToString(r.loadBalancerOriginID, "loadBalancerOriginID")), -1)

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

type ApiUpdateLoadBalancerOriginRequest struct {
	ctx                          context.Context
	ApiService                   *OriginsApiService
	loadBalancerOriginID         string
	loadBalancerPoolOriginUpdate *LoadBalancerPoolOriginUpdate
}

func (r ApiUpdateLoadBalancerOriginRequest) LoadBalancerPoolOriginUpdate(loadBalancerPoolOriginUpdate LoadBalancerPoolOriginUpdate) ApiUpdateLoadBalancerOriginRequest {
	r.loadBalancerPoolOriginUpdate = &loadBalancerPoolOriginUpdate
	return r
}

func (r ApiUpdateLoadBalancerOriginRequest) Execute() (*LoadBalancerPoolOrigin, *http.Response, error) {
	return r.ApiService.UpdateLoadBalancerOriginExecute(r)
}

/*
UpdateLoadBalancerOrigin Update a load balancer origin.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param loadBalancerOriginID ID of the load balancer origin
	@return ApiUpdateLoadBalancerOriginRequest
*/
func (a *OriginsApiService) UpdateLoadBalancerOrigin(ctx context.Context, loadBalancerOriginID string) ApiUpdateLoadBalancerOriginRequest {
	return ApiUpdateLoadBalancerOriginRequest{
		ApiService:           a,
		ctx:                  ctx,
		loadBalancerOriginID: loadBalancerOriginID,
	}
}

// Execute executes the request
//
//	@return LoadBalancerPoolOrigin
func (a *OriginsApiService) UpdateLoadBalancerOriginExecute(r ApiUpdateLoadBalancerOriginRequest) (*LoadBalancerPoolOrigin, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LoadBalancerPoolOrigin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OriginsApiService.UpdateLoadBalancerOrigin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/loadbalancers/pools/origins/{loadBalancerOriginID}"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerOriginID"+"}", url.PathEscape(parameterValueToString(r.loadBalancerOriginID, "loadBalancerOriginID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.loadBalancerPoolOriginUpdate == nil {
		return localVarReturnValue, nil, reportError("loadBalancerPoolOriginUpdate is required and must be specified")
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
	localVarPostBody = r.loadBalancerPoolOriginUpdate
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
