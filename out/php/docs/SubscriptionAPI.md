# \SubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSubscriptionsGetCollection**](SubscriptionAPI.md#ApiSubscriptionsGetCollection) | **Get** /api-platform/subscriptions | Retrieves the collection of Subscription resources.
[**ApiSubscriptionsIdDelete**](SubscriptionAPI.md#ApiSubscriptionsIdDelete) | **Delete** /api-platform/subscriptions/{id} | Removes the Subscription resource.
[**ApiSubscriptionsIdGet**](SubscriptionAPI.md#ApiSubscriptionsIdGet) | **Get** /api-platform/subscriptions/{id} | Retrieves a Subscription resource.
[**ApiSubscriptionsIdPatch**](SubscriptionAPI.md#ApiSubscriptionsIdPatch) | **Patch** /api-platform/subscriptions/{id} | Updates the Subscription resource.
[**ApiSubscriptionsPost**](SubscriptionAPI.md#ApiSubscriptionsPost) | **Post** /api-platform/subscriptions | Creates a Subscription resource.



## ApiSubscriptionsGetCollection

> ApiSubscriptionsGetCollection200Response ApiSubscriptionsGetCollection(ctx).Page(page).FollowerLogin(followerLogin).Execute()

Retrieves the collection of Subscription resources.



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
	followerLogin := "followerLogin_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.ApiSubscriptionsGetCollection(context.Background()).Page(page).FollowerLogin(followerLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ApiSubscriptionsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsGetCollection`: ApiSubscriptionsGetCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ApiSubscriptionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **followerLogin** | **string** |  | 

### Return type

[**ApiSubscriptionsGetCollection200Response**](ApiSubscriptionsGetCollection200Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubscriptionsIdDelete

> ApiSubscriptionsIdDelete(ctx, id).Execute()

Removes the Subscription resource.



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
	id := "id_example" // string | Subscription identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionAPI.ApiSubscriptionsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ApiSubscriptionsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsIdDeleteRequest struct via the builder pattern


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


## ApiSubscriptionsIdGet

> SubscriptionJsonld ApiSubscriptionsIdGet(ctx, id).Execute()

Retrieves a Subscription resource.



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
	id := "id_example" // string | Subscription identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.ApiSubscriptionsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ApiSubscriptionsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsIdGet`: SubscriptionJsonld
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ApiSubscriptionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionJsonld**](SubscriptionJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubscriptionsIdPatch

> SubscriptionJsonld ApiSubscriptionsIdPatch(ctx, id).Subscription(subscription).Execute()

Updates the Subscription resource.



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
	id := "id_example" // string | Subscription identifier
	subscription := *openapiclient.NewSubscription() // Subscription | The updated Subscription resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.ApiSubscriptionsIdPatch(context.Background(), id).Subscription(subscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ApiSubscriptionsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsIdPatch`: SubscriptionJsonld
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ApiSubscriptionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscription** | [**Subscription**](Subscription.md) | The updated Subscription resource | 

### Return type

[**SubscriptionJsonld**](SubscriptionJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubscriptionsPost

> SubscriptionJsonld ApiSubscriptionsPost(ctx).SubscriptionJsonld(subscriptionJsonld).Execute()

Creates a Subscription resource.



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
	subscriptionJsonld := *openapiclient.NewSubscriptionJsonld() // SubscriptionJsonld | The new Subscription resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.ApiSubscriptionsPost(context.Background()).SubscriptionJsonld(subscriptionJsonld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ApiSubscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsPost`: SubscriptionJsonld
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ApiSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionJsonld** | [**SubscriptionJsonld**](SubscriptionJsonld.md) | The new Subscription resource | 

### Return type

[**SubscriptionJsonld**](SubscriptionJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

