# OpenAPI\Client\SubscriptionApi

All URIs are relative to http://localhost, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**apiSubscriptionsGetCollection()**](SubscriptionApi.md#apiSubscriptionsGetCollection) | **GET** /api-platform/subscriptions | Retrieves the collection of Subscription resources. |
| [**apiSubscriptionsIdDelete()**](SubscriptionApi.md#apiSubscriptionsIdDelete) | **DELETE** /api-platform/subscriptions/{id} | Removes the Subscription resource. |
| [**apiSubscriptionsIdGet()**](SubscriptionApi.md#apiSubscriptionsIdGet) | **GET** /api-platform/subscriptions/{id} | Retrieves a Subscription resource. |
| [**apiSubscriptionsIdPatch()**](SubscriptionApi.md#apiSubscriptionsIdPatch) | **PATCH** /api-platform/subscriptions/{id} | Updates the Subscription resource. |
| [**apiSubscriptionsPost()**](SubscriptionApi.md#apiSubscriptionsPost) | **POST** /api-platform/subscriptions | Creates a Subscription resource. |


## `apiSubscriptionsGetCollection()`

```php
apiSubscriptionsGetCollection($page, $follower_login): \OpenAPI\Client\Model\ApiSubscriptionsGetCollection200Response
```

Retrieves the collection of Subscription resources.

Retrieves the collection of Subscription resources.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\SubscriptionApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$page = 1; // int | The collection page number
$follower_login = 'follower_login_example'; // string | 

try {
    $result = $apiInstance->apiSubscriptionsGetCollection($page, $follower_login);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionApi->apiSubscriptionsGetCollection: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **page** | **int**| The collection page number | [optional] [default to 1] |
| **follower_login** | **string**|  | [optional] |

### Return type

[**\OpenAPI\Client\Model\ApiSubscriptionsGetCollection200Response**](../Model/ApiSubscriptionsGetCollection200Response.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiSubscriptionsIdDelete()`

```php
apiSubscriptionsIdDelete($id)
```

Removes the Subscription resource.

Removes the Subscription resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\SubscriptionApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | Subscription identifier

try {
    $apiInstance->apiSubscriptionsIdDelete($id);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionApi->apiSubscriptionsIdDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| Subscription identifier | |

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

## `apiSubscriptionsIdGet()`

```php
apiSubscriptionsIdGet($id): \OpenAPI\Client\Model\SubscriptionJsonld
```

Retrieves a Subscription resource.

Retrieves a Subscription resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\SubscriptionApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | Subscription identifier

try {
    $result = $apiInstance->apiSubscriptionsIdGet($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionApi->apiSubscriptionsIdGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| Subscription identifier | |

### Return type

[**\OpenAPI\Client\Model\SubscriptionJsonld**](../Model/SubscriptionJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiSubscriptionsIdPatch()`

```php
apiSubscriptionsIdPatch($id, $subscription): \OpenAPI\Client\Model\SubscriptionJsonld
```

Updates the Subscription resource.

Updates the Subscription resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\SubscriptionApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 'id_example'; // string | Subscription identifier
$subscription = new \OpenAPI\Client\Model\Subscription(); // \OpenAPI\Client\Model\Subscription | The updated Subscription resource

try {
    $result = $apiInstance->apiSubscriptionsIdPatch($id, $subscription);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionApi->apiSubscriptionsIdPatch: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **id** | **string**| Subscription identifier | |
| **subscription** | [**\OpenAPI\Client\Model\Subscription**](../Model/Subscription.md)| The updated Subscription resource | |

### Return type

[**\OpenAPI\Client\Model\SubscriptionJsonld**](../Model/SubscriptionJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: `application/merge-patch+json`
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `apiSubscriptionsPost()`

```php
apiSubscriptionsPost($subscription_jsonld): \OpenAPI\Client\Model\SubscriptionJsonld
```

Creates a Subscription resource.

Creates a Subscription resource.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: JWT
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\SubscriptionApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$subscription_jsonld = new \OpenAPI\Client\Model\SubscriptionJsonld(); // \OpenAPI\Client\Model\SubscriptionJsonld | The new Subscription resource

try {
    $result = $apiInstance->apiSubscriptionsPost($subscription_jsonld);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionApi->apiSubscriptionsPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **subscription_jsonld** | [**\OpenAPI\Client\Model\SubscriptionJsonld**](../Model/SubscriptionJsonld.md)| The new Subscription resource | |

### Return type

[**\OpenAPI\Client\Model\SubscriptionJsonld**](../Model/SubscriptionJsonld.md)

### Authorization

[JWT](../../README.md#JWT)

### HTTP request headers

- **Content-Type**: `application/ld+json`
- **Accept**: `application/ld+json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
