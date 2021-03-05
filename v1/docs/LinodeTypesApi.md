# \LinodeTypesApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLinodeType**](LinodeTypesApi.md#GetLinodeType) | **Get** /linode/types/{typeId} | Type View
[**GetLinodeTypes**](LinodeTypesApi.md#GetLinodeTypes) | **Get** /linode/types | Types List



## GetLinodeType

> LinodeType GetLinodeType(ctx, typeId).Execute()

Type View



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
    typeId := "typeId_example" // string | The ID of the Linode Type to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeTypesApi.GetLinodeType(context.Background(), typeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeTypesApi.GetLinodeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeType`: LinodeType
    fmt.Fprintf(os.Stdout, "Response from `LinodeTypesApi.GetLinodeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | The ID of the Linode Type to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinodeType**](LinodeType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeTypes

> InlineResponse20023 GetLinodeTypes(ctx).Execute()

Types List



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
    resp, r, err := api_client.LinodeTypesApi.GetLinodeTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeTypesApi.GetLinodeTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeTypes`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `LinodeTypesApi.GetLinodeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTypesRequest struct via the builder pattern


### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

