# OpenAPI\Client\PhoneUserApi

All URIs are relative to http://localhost, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**apiPhoneUsersGetCollection()**](PhoneUserApi.md#apiPhoneUsersGetCollection) | **GET** /api-platform/phone_users | Retrieves the collection of PhoneUser resources. |
| [**apiPhoneUsersIdDelete()**](PhoneUserApi.md#apiPhoneUsersIdDelete) | **DELETE** /api-platform/phone_users/{id} | Removes the PhoneUser resource. |
| [**apiPhoneUsersIdGet()**](PhoneUserApi.md#apiPhoneUsersIdGet) | **GET** /api-platform/phone_users/{id} | Retrieves a PhoneUser resource. |
| [**apiPhoneUsersIdPatch()**](PhoneUserApi.md#apiPhoneUsersIdPatch) | **PATCH** /api-platform/phone_users/{id} | Updates the PhoneUser resource. |
| [**apiPhoneUsersPost()**](PhoneUserApi.md#apiPhoneUsersPost) | **POST** /api-platform/phone_users | Creates a PhoneUser resource. |


## `apiPhoneUsersGetCollection()`

```php
apiPhoneUsersGetCollection($page, $login, $order_login): \OpenAPI\Client\Model\ApiPhoneUsersGetCollection200Response
```

Retrieves the collection of PhoneUser resources.

Retrieves the collection of PhoneUser resources.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PhoneUserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$page = 1; // int | The collection page number
$login = 'login_example'; // string | 
$order_login = 'asc'; // string | 

try {
    $result = $apiInstance->apiPhoneUsersGetCollection($page, $login, $order_login);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhoneUserApi->apiPhoneUsersGetCollection: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **page** | **int**| The collection page number | [optional] [default to 1] |
| **login** | **string**|  | [optional] |
| **order_login** | **string**|  | [optional] [default to &#39;asc&#39;] |

### Return type

[**\OpenAPI\Client\Model\ApiPhoneUsersGetCollection200Response**](../Model/ApiPhoneUsersGetCollection200Response.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiPhoneUsersIdDelete()`

```php
apiPhoneUsersIdDelete($id)
```

Removes the PhoneUser resource.

Removes the PhoneUser resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PhoneUserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | PhoneUser identifier

try {
    $apiInstance->apiPhoneUsersIdDelete($id);
} catch (Exception $e) {
    echo 'Exception when calling PhoneUserApi->apiPhoneUsersIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| PhoneUser identifier | |

### Return type

void (empty response body)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiPhoneUsersIdGet()`

```php
apiPhoneUsersIdGet($id): \OpenAPI\Client\Model\PhoneUserJsonld
```

Retrieves a PhoneUser resource.

Retrieves a PhoneUser resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PhoneUserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | PhoneUser identifier

try {
    $result = $apiInstance->apiPhoneUsersIdGet($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhoneUserApi->apiPhoneUsersIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| PhoneUser identifier | |

### Return type

[**\OpenAPI\Client\Model\PhoneUserJsonld**](../Model/PhoneUserJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiPhoneUsersIdPatch()`

```php
apiPhoneUsersIdPatch($id, $phone_user): \OpenAPI\Client\Model\PhoneUserJsonld
```

Updates the PhoneUser resource.

Updates the PhoneUser resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PhoneUserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | PhoneUser identifier
$phone_user = new \OpenAPI\Client\Model\PhoneUser(); // \OpenAPI\Client\Model\PhoneUser | The updated PhoneUser resource

try {
    $result = $apiInstance->apiPhoneUsersIdPatch($id, $phone_user);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhoneUserApi->apiPhoneUsersIdPatch: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| PhoneUser identifier | |
| **phone_user** | [**\OpenAPI\Client\Model\PhoneUser**](../Model/PhoneUser.md)| The updated PhoneUser resource | |

### Return type

[**\OpenAPI\Client\Model\PhoneUserJsonld**](../Model/PhoneUserJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: `application/merge-patch+json`
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiPhoneUsersPost()`

```php
apiPhoneUsersPost($phone_user_jsonld): \OpenAPI\Client\Model\PhoneUserJsonld
```

Creates a PhoneUser resource.

Creates a PhoneUser resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PhoneUserApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$phone_user_jsonld = new \OpenAPI\Client\Model\PhoneUserJsonld(); // \OpenAPI\Client\Model\PhoneUserJsonld | The new PhoneUser resource

try {
    $result = $apiInstance->apiPhoneUsersPost($phone_user_jsonld);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhoneUserApi->apiPhoneUsersPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **phone_user_jsonld** | [**\OpenAPI\Client\Model\PhoneUserJsonld**](../Model/PhoneUserJsonld.md)| The new PhoneUser resource | |

### Return type

[**\OpenAPI\Client\Model\PhoneUserJsonld**](../Model/PhoneUserJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: `application/ld+json`
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
