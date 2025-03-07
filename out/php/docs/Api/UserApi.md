# OpenAPI\Client\UserApi

All URIs are relative to http://localhost, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**apiUsersPost()**](UserApi.md#apiUsersPost) | **POST** /api-platform/users | Creates a User resource. |


## `apiUsersPost()`

```php
apiUsersPost($user_create_user_dto_jsonld): \OpenAPI\Client\Model\UserCreatedUserDTOJsonld
```

Creates a User resource.

Creates a User resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\UserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_create_user_dto_jsonld = new \OpenAPI\Client\Model\UserCreateUserDTOJsonld(); // \OpenAPI\Client\Model\UserCreateUserDTOJsonld | The new User resource

try {
    $result = $apiInstance->apiUsersPost($user_create_user_dto_jsonld);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UserApi->apiUsersPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **user_create_user_dto_jsonld** | [**\OpenAPI\Client\Model\UserCreateUserDTOJsonld**](../Model/UserCreateUserDTOJsonld.md)| The new User resource | |

### Return type

[**\OpenAPI\Client\Model\UserCreatedUserDTOJsonld**](../Model/UserCreatedUserDTOJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: `application/ld+json`
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
