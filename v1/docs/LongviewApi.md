# \LongviewApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLongviewClient**](LongviewApi.md#CreateLongviewClient) | **Post** /longview/clients | Longview Client Create
[**DeleteLongviewClient**](LongviewApi.md#DeleteLongviewClient) | **Delete** /longview/clients/{clientId} | Longview Client Delete
[**GetLongviewClient**](LongviewApi.md#GetLongviewClient) | **Get** /longview/clients/{clientId} | Longview Client View
[**GetLongviewClients**](LongviewApi.md#GetLongviewClients) | **Get** /longview/clients | Longview Clients List
[**GetLongviewPlan**](LongviewApi.md#GetLongviewPlan) | **Get** /longview/plan | Longview Plan View
[**GetLongviewSubscription**](LongviewApi.md#GetLongviewSubscription) | **Get** /longview/subscriptions/{subscriptionId} | Longview Subscription View
[**GetLongviewSubscriptions**](LongviewApi.md#GetLongviewSubscriptions) | **Get** /longview/subscriptions | Longview Subscriptions List
[**UpdateLongviewClient**](LongviewApi.md#UpdateLongviewClient) | **Put** /longview/clients/{clientId} | Longview Client Update
[**UpdateLongviewPlan**](LongviewApi.md#UpdateLongviewPlan) | **Put** /longview/plan | Longview Plan Update



## CreateLongviewClient

> LongviewClient CreateLongviewClient(ctx).LongviewClient(longviewClient).Execute()

Longview Client Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    longviewClient := *openapiclient.NewLongviewClient() // LongviewClient | Information about the LongviewClient to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.CreateLongviewClient(context.Background()).LongviewClient(longviewClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.CreateLongviewClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLongviewClient`: LongviewClient
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.CreateLongviewClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **longviewClient** | [**LongviewClient**](LongviewClient.md) | Information about the LongviewClient to create. | 

### Return type

[**LongviewClient**](LongviewClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLongviewClient

> map[string]interface{} DeleteLongviewClient(ctx, clientId).Execute()

Longview Client Delete



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := int32(56) // int32 | The Longview Client ID to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.DeleteLongviewClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.DeleteLongviewClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLongviewClient`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.DeleteLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewClient

> LongviewClient GetLongviewClient(ctx, clientId).Execute()

Longview Client View



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := int32(56) // int32 | The Longview Client ID to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.GetLongviewClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.GetLongviewClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLongviewClient`: LongviewClient
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.GetLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LongviewClient**](LongviewClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewClients

> InlineResponse20031 GetLongviewClients(ctx).Page(page).PageSize(pageSize).Execute()

Longview Clients List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.GetLongviewClients(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.GetLongviewClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLongviewClients`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.GetLongviewClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewPlan

> LongviewSubscription GetLongviewPlan(ctx).Execute()

Longview Plan View



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.GetLongviewPlan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.GetLongviewPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLongviewPlan`: LongviewSubscription
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.GetLongviewPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewPlanRequest struct via the builder pattern


### Return type

[**LongviewSubscription**](LongviewSubscription.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewSubscription

> LongviewSubscription GetLongviewSubscription(ctx, subscriptionId).Execute()

Longview Subscription View



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | The Longview Subscription to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.GetLongviewSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.GetLongviewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLongviewSubscription`: LongviewSubscription
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.GetLongviewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The Longview Subscription to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LongviewSubscription**](LongviewSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewSubscriptions

> InlineResponse20032 GetLongviewSubscriptions(ctx).Page(page).PageSize(pageSize).Execute()

Longview Subscriptions List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.GetLongviewSubscriptions(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.GetLongviewSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLongviewSubscriptions`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.GetLongviewSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLongviewClient

> LongviewClient UpdateLongviewClient(ctx, clientId).LongviewClient(longviewClient).Execute()

Longview Client Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := int32(56) // int32 | The Longview Client ID to access.
    longviewClient := *openapiclient.NewLongviewClient() // LongviewClient | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.UpdateLongviewClient(context.Background(), clientId).LongviewClient(longviewClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.UpdateLongviewClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLongviewClient`: LongviewClient
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.UpdateLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **longviewClient** | [**LongviewClient**](LongviewClient.md) | The fields to update. | 

### Return type

[**LongviewClient**](LongviewClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLongviewPlan

> LongviewSubscription UpdateLongviewPlan(ctx).LongviewPlan(longviewPlan).Execute()

Longview Plan Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    longviewPlan := *openapiclient.NewLongviewPlan() // LongviewPlan | Update your Longview subscription plan.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LongviewApi.UpdateLongviewPlan(context.Background()).LongviewPlan(longviewPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LongviewApi.UpdateLongviewPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLongviewPlan`: LongviewSubscription
    fmt.Fprintf(os.Stdout, "Response from `LongviewApi.UpdateLongviewPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLongviewPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **longviewPlan** | [**LongviewPlan**](LongviewPlan.md) | Update your Longview subscription plan. | 

### Return type

[**LongviewSubscription**](LongviewSubscription.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

