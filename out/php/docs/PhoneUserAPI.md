# \PhoneUserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPhoneUsersGetCollection**](PhoneUserAPI.md#ApiPhoneUsersGetCollection) | **Get** /api-platform/phone_users | Retrieves the collection of PhoneUser resources.
[**ApiPhoneUsersIdDelete**](PhoneUserAPI.md#ApiPhoneUsersIdDelete) | **Delete** /api-platform/phone_users/{id} | Removes the PhoneUser resource.
[**ApiPhoneUsersIdGet**](PhoneUserAPI.md#ApiPhoneUsersIdGet) | **Get** /api-platform/phone_users/{id} | Retrieves a PhoneUser resource.
[**ApiPhoneUsersIdPatch**](PhoneUserAPI.md#ApiPhoneUsersIdPatch) | **Patch** /api-platform/phone_users/{id} | Updates the PhoneUser resource.
[**ApiPhoneUsersPost**](PhoneUserAPI.md#ApiPhoneUsersPost) | **Post** /api-platform/phone_users | Creates a PhoneUser resource.



## ApiPhoneUsersGetCollection

> ApiPhoneUsersGetCollection200Response ApiPhoneUsersGetCollection(ctx).Page(page).Login(login).OrderLogin(orderLogin).Execute()

Retrieves the collection of PhoneUser resources.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | The collection page number (optional) (default to 1)
	login := "login_example" // string |  (optional)
	orderLogin := "orderLogin_example" // string |  (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneUserAPI.ApiPhoneUsersGetCollection(context.Background()).Page(page).Login(login).OrderLogin(orderLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneUserAPI.ApiPhoneUsersGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPhoneUsersGetCollection`: ApiPhoneUsersGetCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `PhoneUserAPI.ApiPhoneUsersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPhoneUsersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **login** | **string** |  | 
 **orderLogin** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**ApiPhoneUsersGetCollection200Response**](ApiPhoneUsersGetCollection200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPhoneUsersIdDelete

> ApiPhoneUsersIdDelete(ctx, id).Execute()

Removes the PhoneUser resource.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | PhoneUser identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PhoneUserAPI.ApiPhoneUsersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneUserAPI.ApiPhoneUsersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | PhoneUser identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPhoneUsersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPhoneUsersIdGet

> PhoneUserJsonld ApiPhoneUsersIdGet(ctx, id).Execute()

Retrieves a PhoneUser resource.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | PhoneUser identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneUserAPI.ApiPhoneUsersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneUserAPI.ApiPhoneUsersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPhoneUsersIdGet`: PhoneUserJsonld
	fmt.Fprintf(os.Stdout, "Response from `PhoneUserAPI.ApiPhoneUsersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | PhoneUser identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPhoneUsersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneUserJsonld**](PhoneUserJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPhoneUsersIdPatch

> PhoneUserJsonld ApiPhoneUsersIdPatch(ctx, id).PhoneUser(phoneUser).Execute()

Updates the PhoneUser resource.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | PhoneUser identifier
	phoneUser := *openapiclient.NewPhoneUser() // PhoneUser | The updated PhoneUser resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneUserAPI.ApiPhoneUsersIdPatch(context.Background(), id).PhoneUser(phoneUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneUserAPI.ApiPhoneUsersIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPhoneUsersIdPatch`: PhoneUserJsonld
	fmt.Fprintf(os.Stdout, "Response from `PhoneUserAPI.ApiPhoneUsersIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | PhoneUser identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPhoneUsersIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneUser** | [**PhoneUser**](PhoneUser.md) | The updated PhoneUser resource | 

### Return type

[**PhoneUserJsonld**](PhoneUserJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPhoneUsersPost

> PhoneUserJsonld ApiPhoneUsersPost(ctx).PhoneUserJsonld(phoneUserJsonld).Execute()

Creates a PhoneUser resource.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	phoneUserJsonld := *openapiclient.NewPhoneUserJsonld() // PhoneUserJsonld | The new PhoneUser resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneUserAPI.ApiPhoneUsersPost(context.Background()).PhoneUserJsonld(phoneUserJsonld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneUserAPI.ApiPhoneUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPhoneUsersPost`: PhoneUserJsonld
	fmt.Fprintf(os.Stdout, "Response from `PhoneUserAPI.ApiPhoneUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPhoneUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneUserJsonld** | [**PhoneUserJsonld**](PhoneUserJsonld.md) | The new PhoneUser resource | 

### Return type

[**PhoneUserJsonld**](PhoneUserJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

