# \ImagesApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImagesApi.md#CreateImage) | **Post** /images | Image Create
[**DeleteImage**](ImagesApi.md#DeleteImage) | **Delete** /images/{imageId} | Image Delete
[**GetImage**](ImagesApi.md#GetImage) | **Get** /images/{imageId} | Image View
[**GetImages**](ImagesApi.md#GetImages) | **Get** /images | Images List
[**UpdateImage**](ImagesApi.md#UpdateImage) | **Put** /images/{imageId} | Image Update



## CreateImage

> ImagePrivate CreateImage(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Image Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the Image to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.CreateImage(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImage`: ImagePrivate
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the Image to create. | 

### Return type

[**ImagePrivate**](ImagePrivate.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> map[string]interface{} DeleteImage(ctx, imageId).Execute()

Image Delete



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
    imageId := "imageId_example" // string | ID of the Image to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.DeleteImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteImage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID of the Image to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## GetImage

> ImagePrivate GetImage(ctx, imageId).Execute()

Image View



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
    imageId := "imageId_example" // string | ID of the Image to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.GetImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImage`: ImagePrivate
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID of the Image to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImagePrivate**](ImagePrivate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> InlineResponse20014 GetImages(ctx).Page(page).PageSize(pageSize).Execute()

Images List



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
    resp, r, err := api_client.ImagesApi.GetImages(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImages`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> ImagePrivate UpdateImage(ctx, imageId).ImagePrivate(imagePrivate).Execute()

Image Update



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
    imageId := "imageId_example" // string | ID of the Image to look up.
    imagePrivate := *openapiclient.NewImagePrivate() // ImagePrivate | The fields to update. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.UpdateImage(context.Background(), imageId).ImagePrivate(imagePrivate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.UpdateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImage`: ImagePrivate
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID of the Image to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imagePrivate** | [**ImagePrivate**](ImagePrivate.md) | The fields to update.  | 

### Return type

[**ImagePrivate**](ImagePrivate.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

