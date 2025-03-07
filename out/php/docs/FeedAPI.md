# \FeedAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetFeed**](FeedAPI.md#V1GetFeed) | **Get** /api/v1/get-feed/{id} | 



## V1GetFeed

> Response V1GetFeed(ctx, id).Count(count).Execute()





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
	id := int32(56) // int32 | Идентификатор пользователя
	count := int32(56) // int32 | Количество сообщений в ленте (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.V1GetFeed(context.Background(), id).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.V1GetFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetFeed`: Response
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.V1GetFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Идентификатор пользователя | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Количество сообщений в ленте | 

### Return type

[**Response**](Response.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

