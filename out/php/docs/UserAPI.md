# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersPost**](UserAPI.md#ApiUsersPost) | **Post** /api-platform/users | Creates a User resource.



## ApiUsersPost

> UserCreatedUserDTOJsonld ApiUsersPost(ctx).UserCreateUserDTOJsonld(userCreateUserDTOJsonld).Execute()

Creates a User resource.



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
	userCreateUserDTOJsonld := *openapiclient.NewUserCreateUserDTOJsonld("Login_example", "Password_example", int32(123), false) // UserCreateUserDTOJsonld | The new User resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ApiUsersPost(context.Background()).UserCreateUserDTOJsonld(userCreateUserDTOJsonld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersPost`: UserCreatedUserDTOJsonld
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ApiUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateUserDTOJsonld** | [**UserCreateUserDTOJsonld**](UserCreateUserDTOJsonld.md) | The new User resource | 

### Return type

[**UserCreatedUserDTOJsonld**](UserCreatedUserDTOJsonld.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

