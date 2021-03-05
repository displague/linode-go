# \StackScriptsApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStackScript**](StackScriptsApi.md#AddStackScript) | **Post** /linode/stackscripts | StackScript Create
[**DeleteStackScript**](StackScriptsApi.md#DeleteStackScript) | **Delete** /linode/stackscripts/{stackscriptId} | StackScript Delete
[**GetStackScript**](StackScriptsApi.md#GetStackScript) | **Get** /linode/stackscripts/{stackscriptId} | StackScript View
[**GetStackScripts**](StackScriptsApi.md#GetStackScripts) | **Get** /linode/stackscripts | StackScripts List
[**UpdateStackScript**](StackScriptsApi.md#UpdateStackScript) | **Put** /linode/stackscripts/{stackscriptId} | StackScript Update



## AddStackScript

> StackScript AddStackScript(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

StackScript Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The properties to set for the new StackScript.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackScriptsApi.AddStackScript(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackScriptsApi.AddStackScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStackScript`: StackScript
    fmt.Fprintf(os.Stdout, "Response from `StackScriptsApi.AddStackScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The properties to set for the new StackScript. | 

### Return type

[**StackScript**](StackScript.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStackScript

> map[string]interface{} DeleteStackScript(ctx, stackscriptId).Execute()

StackScript Delete



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
    stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackScriptsApi.DeleteStackScript(context.Background(), stackscriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackScriptsApi.DeleteStackScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStackScript`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StackScriptsApi.DeleteStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackScriptRequest struct via the builder pattern


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


## GetStackScript

> StackScript GetStackScript(ctx, stackscriptId).Execute()

StackScript View



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
    stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackScriptsApi.GetStackScript(context.Background(), stackscriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackScriptsApi.GetStackScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackScript`: StackScript
    fmt.Fprintf(os.Stdout, "Response from `StackScriptsApi.GetStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackScript**](StackScript.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackScripts

> InlineResponse20022 GetStackScripts(ctx).Page(page).PageSize(pageSize).Execute()

StackScripts List



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
    resp, r, err := api_client.StackScriptsApi.GetStackScripts(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackScriptsApi.GetStackScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackScripts`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `StackScriptsApi.GetStackScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStackScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackScript

> StackScript UpdateStackScript(ctx, stackscriptId).StackScript(stackScript).Execute()

StackScript Update



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
    stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.
    stackScript := *openapiclient.NewStackScript() // StackScript | The fields to update. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackScriptsApi.UpdateStackScript(context.Background(), stackscriptId).StackScript(stackScript).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackScriptsApi.UpdateStackScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStackScript`: StackScript
    fmt.Fprintf(os.Stdout, "Response from `StackScriptsApi.UpdateStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackScript** | [**StackScript**](StackScript.md) | The fields to update. | 

### Return type

[**StackScript**](StackScript.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

