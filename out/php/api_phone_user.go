/*
My App

This is an awesome app!

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PhoneUserAPIService PhoneUserAPI service
type PhoneUserAPIService service

type ApiApiPhoneUsersGetCollectionRequest struct {
	ctx context.Context
	ApiService *PhoneUserAPIService
	page *int32
	login *string
	orderLogin *string
}

// The collection page number
func (r ApiApiPhoneUsersGetCollectionRequest) Page(page int32) ApiApiPhoneUsersGetCollectionRequest {
	r.page = &page
	return r
}

// 
func (r ApiApiPhoneUsersGetCollectionRequest) Login(login string) ApiApiPhoneUsersGetCollectionRequest {
	r.login = &login
	return r
}

// 
func (r ApiApiPhoneUsersGetCollectionRequest) OrderLogin(orderLogin string) ApiApiPhoneUsersGetCollectionRequest {
	r.orderLogin = &orderLogin
	return r
}

func (r ApiApiPhoneUsersGetCollectionRequest) Execute() (*ApiPhoneUsersGetCollection200Response, *http.Response, error) {
	return r.ApiService.ApiPhoneUsersGetCollectionExecute(r)
}

/*
ApiPhoneUsersGetCollection Retrieves the collection of PhoneUser resources.

Retrieves the collection of PhoneUser resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiPhoneUsersGetCollectionRequest
*/
func (a *PhoneUserAPIService) ApiPhoneUsersGetCollection(ctx context.Context) ApiApiPhoneUsersGetCollectionRequest {
	return ApiApiPhoneUsersGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiPhoneUsersGetCollection200Response
func (a *PhoneUserAPIService) ApiPhoneUsersGetCollectionExecute(r ApiApiPhoneUsersGetCollectionRequest) (*ApiPhoneUsersGetCollection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiPhoneUsersGetCollection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PhoneUserAPIService.ApiPhoneUsersGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api-platform/phone_users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.login != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "login", r.login, "form", "")
	}
	if r.orderLogin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order[login]", r.orderLogin, "form", "")
	} else {
		var defaultValue string = "asc"
		r.orderLogin = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiPhoneUsersIdDeleteRequest struct {
	ctx context.Context
	ApiService *PhoneUserAPIService
	id string
}

func (r ApiApiPhoneUsersIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiPhoneUsersIdDeleteExecute(r)
}

/*
ApiPhoneUsersIdDelete Removes the PhoneUser resource.

Removes the PhoneUser resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id PhoneUser identifier
 @return ApiApiPhoneUsersIdDeleteRequest
*/
func (a *PhoneUserAPIService) ApiPhoneUsersIdDelete(ctx context.Context, id string) ApiApiPhoneUsersIdDeleteRequest {
	return ApiApiPhoneUsersIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PhoneUserAPIService) ApiPhoneUsersIdDeleteExecute(r ApiApiPhoneUsersIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PhoneUserAPIService.ApiPhoneUsersIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api-platform/phone_users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiPhoneUsersIdGetRequest struct {
	ctx context.Context
	ApiService *PhoneUserAPIService
	id string
}

func (r ApiApiPhoneUsersIdGetRequest) Execute() (*PhoneUserJsonld, *http.Response, error) {
	return r.ApiService.ApiPhoneUsersIdGetExecute(r)
}

/*
ApiPhoneUsersIdGet Retrieves a PhoneUser resource.

Retrieves a PhoneUser resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id PhoneUser identifier
 @return ApiApiPhoneUsersIdGetRequest
*/
func (a *PhoneUserAPIService) ApiPhoneUsersIdGet(ctx context.Context, id string) ApiApiPhoneUsersIdGetRequest {
	return ApiApiPhoneUsersIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PhoneUserJsonld
func (a *PhoneUserAPIService) ApiPhoneUsersIdGetExecute(r ApiApiPhoneUsersIdGetRequest) (*PhoneUserJsonld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PhoneUserJsonld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PhoneUserAPIService.ApiPhoneUsersIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api-platform/phone_users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiPhoneUsersIdPatchRequest struct {
	ctx context.Context
	ApiService *PhoneUserAPIService
	id string
	phoneUser *PhoneUser
}

// The updated PhoneUser resource
func (r ApiApiPhoneUsersIdPatchRequest) PhoneUser(phoneUser PhoneUser) ApiApiPhoneUsersIdPatchRequest {
	r.phoneUser = &phoneUser
	return r
}

func (r ApiApiPhoneUsersIdPatchRequest) Execute() (*PhoneUserJsonld, *http.Response, error) {
	return r.ApiService.ApiPhoneUsersIdPatchExecute(r)
}

/*
ApiPhoneUsersIdPatch Updates the PhoneUser resource.

Updates the PhoneUser resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id PhoneUser identifier
 @return ApiApiPhoneUsersIdPatchRequest
*/
func (a *PhoneUserAPIService) ApiPhoneUsersIdPatch(ctx context.Context, id string) ApiApiPhoneUsersIdPatchRequest {
	return ApiApiPhoneUsersIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PhoneUserJsonld
func (a *PhoneUserAPIService) ApiPhoneUsersIdPatchExecute(r ApiApiPhoneUsersIdPatchRequest) (*PhoneUserJsonld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PhoneUserJsonld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PhoneUserAPIService.ApiPhoneUsersIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api-platform/phone_users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.phoneUser == nil {
		return localVarReturnValue, nil, reportError("phoneUser is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.phoneUser
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiPhoneUsersPostRequest struct {
	ctx context.Context
	ApiService *PhoneUserAPIService
	phoneUserJsonld *PhoneUserJsonld
}

// The new PhoneUser resource
func (r ApiApiPhoneUsersPostRequest) PhoneUserJsonld(phoneUserJsonld PhoneUserJsonld) ApiApiPhoneUsersPostRequest {
	r.phoneUserJsonld = &phoneUserJsonld
	return r
}

func (r ApiApiPhoneUsersPostRequest) Execute() (*PhoneUserJsonld, *http.Response, error) {
	return r.ApiService.ApiPhoneUsersPostExecute(r)
}

/*
ApiPhoneUsersPost Creates a PhoneUser resource.

Creates a PhoneUser resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiPhoneUsersPostRequest
*/
func (a *PhoneUserAPIService) ApiPhoneUsersPost(ctx context.Context) ApiApiPhoneUsersPostRequest {
	return ApiApiPhoneUsersPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PhoneUserJsonld
func (a *PhoneUserAPIService) ApiPhoneUsersPostExecute(r ApiApiPhoneUsersPostRequest) (*PhoneUserJsonld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PhoneUserJsonld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PhoneUserAPIService.ApiPhoneUsersPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api-platform/phone_users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.phoneUserJsonld == nil {
		return localVarReturnValue, nil, reportError("phoneUserJsonld is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/ld+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.phoneUserJsonld
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
