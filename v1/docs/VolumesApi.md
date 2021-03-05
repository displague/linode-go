# \VolumesApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolume**](VolumesApi.md#AttachVolume) | **Post** /volumes/{volumeId}/attach | Volume Attach
[**CloneVolume**](VolumesApi.md#CloneVolume) | **Post** /volumes/{volumeId}/clone | Volume Clone
[**CreateVolume**](VolumesApi.md#CreateVolume) | **Post** /volumes | Volume Create
[**DeleteVolume**](VolumesApi.md#DeleteVolume) | **Delete** /volumes/{volumeId} | Volume Delete
[**DetachVolume**](VolumesApi.md#DetachVolume) | **Post** /volumes/{volumeId}/detach | Volume Detach
[**GetVolume**](VolumesApi.md#GetVolume) | **Get** /volumes/{volumeId} | Volume View
[**GetVolumes**](VolumesApi.md#GetVolumes) | **Get** /volumes | Volumes List
[**ResizeVolume**](VolumesApi.md#ResizeVolume) | **Post** /volumes/{volumeId}/resize | Volume Resize
[**UpdateVolume**](VolumesApi.md#UpdateVolume) | **Put** /volumes/{volumeId} | Volume Update



## AttachVolume

> Volume AttachVolume(ctx, volumeId).InlineObject17(inlineObject17).Execute()

Volume Attach



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
    volumeId := int32(56) // int32 | ID of the Volume to attach.
    inlineObject17 := *openapiclient.NewInlineObject17("TODO") // InlineObject17 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.AttachVolume(context.Background(), volumeId).InlineObject17(inlineObject17).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.AttachVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.AttachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to attach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject17** | [**InlineObject17**](InlineObject17.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneVolume

> Volume CloneVolume(ctx, volumeId).InlineObject18(inlineObject18).Execute()

Volume Clone



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
    volumeId := int32(56) // int32 | ID of the Volume to clone.
    inlineObject18 := *openapiclient.NewInlineObject18("TODO") // InlineObject18 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CloneVolume(context.Background(), volumeId).InlineObject18(inlineObject18).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CloneVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CloneVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject18** | [**InlineObject18**](InlineObject18.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolume

> Volume CreateVolume(ctx).InlineObject16(inlineObject16).Execute()

Volume Create



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
    inlineObject16 := *openapiclient.NewInlineObject16() // InlineObject16 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.CreateVolume(context.Background()).InlineObject16(inlineObject16).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject16** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> map[string]interface{} DeleteVolume(ctx, volumeId).Execute()

Volume Delete



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
    volumeId := int32(56) // int32 | ID of the Volume to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DeleteVolume(context.Background(), volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolume`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## DetachVolume

> map[string]interface{} DetachVolume(ctx, volumeId).Execute()

Volume Detach



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
    volumeId := int32(56) // int32 | ID of the Volume to detach.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DetachVolume(context.Background(), volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DetachVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachVolume`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DetachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to detach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumeRequest struct via the builder pattern


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


## GetVolume

> Volume GetVolume(ctx, volumeId).Page(page).PageSize(pageSize).Execute()

Volume View



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
    volumeId := int32(56) // int32 | ID of the Volume to look up.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.GetVolume(context.Background(), volumeId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumes

> InlineResponse20021 GetVolumes(ctx).Page(page).PageSize(pageSize).Execute()

Volumes List



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
    resp, r, err := api_client.VolumesApi.GetVolumes(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumes`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVolume

> map[string]interface{} ResizeVolume(ctx, volumeId).InlineObject19(inlineObject19).Execute()

Volume Resize



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
    volumeId := int32(56) // int32 | ID of the Volume to resize.
    inlineObject19 := *openapiclient.NewInlineObject19("TODO") // InlineObject19 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.ResizeVolume(context.Background(), volumeId).InlineObject19(inlineObject19).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.ResizeVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeVolume`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.ResizeVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to resize. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject19** | [**InlineObject19**](InlineObject19.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> Volume UpdateVolume(ctx, volumeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Volume Update



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
    volumeId := int32(56) // int32 | ID of the Volume to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | If any updated field fails to pass validation, the Volume will not be updated. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.UpdateVolume(context.Background(), volumeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.UpdateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | If any updated field fails to pass validation, the Volume will not be updated.  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

