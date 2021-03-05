# \RegionsApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegion**](RegionsApi.md#GetRegion) | **Get** /regions/{regionId} | Region View
[**GetRegions**](RegionsApi.md#GetRegions) | **Get** /regions | Regions List



## GetRegion

> Region GetRegion(ctx, regionId).Execute()

Region View



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
    regionId := "regionId_example" // string | ID of the Region to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionsApi.GetRegion(context.Background(), regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegion`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string** | ID of the Region to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Region**](Region.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegions

> InlineResponse20055 GetRegions(ctx).Execute()

Regions List



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
    resp, r, err := api_client.RegionsApi.GetRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.GetRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegions`: InlineResponse20055
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


### Return type

[**InlineResponse20055**](InlineResponse20055.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

