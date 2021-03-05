# \TagsApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /tags | New Tag Create
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /tags/{label} | Tag Delete
[**GetTaggedObjects**](TagsApi.md#GetTaggedObjects) | **Get** /tags/{label} | Tagged Objects List
[**GetTags**](TagsApi.md#GetTags) | **Get** /tags | Tags List



## CreateTag

> Tag CreateTag(ctx).InlineObject15(inlineObject15).Execute()

New Tag Create



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
    inlineObject15 := *openapiclient.NewInlineObject15("example tag") // InlineObject15 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.CreateTag(context.Background()).InlineObject15(inlineObject15).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**InlineObject15**](InlineObject15.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> map[string]interface{} DeleteTag(ctx, label).Execute()

Tag Delete



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
    label := "label_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.DeleteTag(context.Background(), label).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTag`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetTaggedObjects

> map[string]interface{} GetTaggedObjects(ctx, label).Page(page).PageSize(pageSize).Execute()

Tagged Objects List



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
    label := "label_example" // string | 
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GetTaggedObjects(context.Background(), label).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTaggedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaggedObjects`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTaggedObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaggedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

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


## GetTags

> InlineResponse20058 GetTags(ctx).Page(page).PageSize(pageSize).Execute()

Tags List



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
    resp, r, err := api_client.TagsApi.GetTags(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: InlineResponse20058
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20058**](InlineResponse20058.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

