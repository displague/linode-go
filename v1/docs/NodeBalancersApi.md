# \NodeBalancersApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeBalancer**](NodeBalancersApi.md#CreateNodeBalancer) | **Post** /nodebalancers | NodeBalancer Create
[**CreateNodeBalancerConfig**](NodeBalancersApi.md#CreateNodeBalancerConfig) | **Post** /nodebalancers/{nodeBalancerId}/configs | Config Create
[**CreateNodeBalancerNode**](NodeBalancersApi.md#CreateNodeBalancerNode) | **Post** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | Node Create
[**DeleteNodeBalancer**](NodeBalancersApi.md#DeleteNodeBalancer) | **Delete** /nodebalancers/{nodeBalancerId} | NodeBalancer Delete
[**DeleteNodeBalancerConfig**](NodeBalancersApi.md#DeleteNodeBalancerConfig) | **Delete** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config Delete
[**DeleteNodeBalancerConfigNode**](NodeBalancersApi.md#DeleteNodeBalancerConfigNode) | **Delete** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node Delete
[**GetNodeBalancer**](NodeBalancersApi.md#GetNodeBalancer) | **Get** /nodebalancers/{nodeBalancerId} | NodeBalancer View
[**GetNodeBalancerConfig**](NodeBalancersApi.md#GetNodeBalancerConfig) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config View
[**GetNodeBalancerConfigNodes**](NodeBalancersApi.md#GetNodeBalancerConfigNodes) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | Nodes List
[**GetNodeBalancerConfigs**](NodeBalancersApi.md#GetNodeBalancerConfigs) | **Get** /nodebalancers/{nodeBalancerId}/configs | Configs List
[**GetNodeBalancerNode**](NodeBalancersApi.md#GetNodeBalancerNode) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node View
[**GetNodeBalancers**](NodeBalancersApi.md#GetNodeBalancers) | **Get** /nodebalancers | NodeBalancers List
[**NodebalancersNodeBalancerIdStatsGet**](NodeBalancersApi.md#NodebalancersNodeBalancerIdStatsGet) | **Get** /nodebalancers/{nodeBalancerId}/stats | NodeBalancer Statistics View
[**RebuildNodeBalancerConfig**](NodeBalancersApi.md#RebuildNodeBalancerConfig) | **Post** /nodebalancers/{nodeBalancerId}/configs/{configId}/rebuild | Config Rebuild
[**UpdateNodeBalancer**](NodeBalancersApi.md#UpdateNodeBalancer) | **Put** /nodebalancers/{nodeBalancerId} | NodeBalancer Update
[**UpdateNodeBalancerConfig**](NodeBalancersApi.md#UpdateNodeBalancerConfig) | **Put** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config Update
[**UpdateNodeBalancerNode**](NodeBalancersApi.md#UpdateNodeBalancerNode) | **Put** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node Update



## CreateNodeBalancer

> NodeBalancer CreateNodeBalancer(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

NodeBalancer Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the NodeBalancer to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.CreateNodeBalancer(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.CreateNodeBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodeBalancer`: NodeBalancer
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.CreateNodeBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the NodeBalancer to create. | 

### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodeBalancerConfig

> NodeBalancerConfig CreateNodeBalancerConfig(ctx, nodeBalancerId).NodeBalancerConfig(nodeBalancerConfig).Execute()

Config Create



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    nodeBalancerConfig := *openapiclient.NewNodeBalancerConfig() // NodeBalancerConfig | Information about the port to configure. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.CreateNodeBalancerConfig(context.Background(), nodeBalancerId).NodeBalancerConfig(nodeBalancerConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.CreateNodeBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodeBalancerConfig`: NodeBalancerConfig
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.CreateNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBalancerConfig** | [**NodeBalancerConfig**](NodeBalancerConfig.md) | Information about the port to configure. | 

### Return type

[**NodeBalancerConfig**](NodeBalancerConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodeBalancerNode

> NodeBalancerNode CreateNodeBalancerNode(ctx, nodeBalancerId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Node Create



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the NodeBalancer config to access.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the Node to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.CreateNodeBalancerNode(context.Background(), nodeBalancerId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.CreateNodeBalancerNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodeBalancerNode`: NodeBalancerNode
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.CreateNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the NodeBalancer config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the Node to create. | 

### Return type

[**NodeBalancerNode**](NodeBalancerNode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodeBalancer

> map[string]interface{} DeleteNodeBalancer(ctx, nodeBalancerId).Execute()

NodeBalancer Delete



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.DeleteNodeBalancer(context.Background(), nodeBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.DeleteNodeBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNodeBalancer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.DeleteNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerRequest struct via the builder pattern


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


## DeleteNodeBalancerConfig

> map[string]interface{} DeleteNodeBalancerConfig(ctx, nodeBalancerId, configId).Execute()

Config Delete



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the config to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.DeleteNodeBalancerConfig(context.Background(), nodeBalancerId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.DeleteNodeBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNodeBalancerConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.DeleteNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerConfigRequest struct via the builder pattern


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


## DeleteNodeBalancerConfigNode

> map[string]interface{} DeleteNodeBalancerConfigNode(ctx, nodeBalancerId, configId, nodeId).Execute()

Node Delete



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the Config to access
    nodeId := int32(56) // int32 | The ID of the Node to access

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.DeleteNodeBalancerConfigNode(context.Background(), nodeBalancerId, configId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.DeleteNodeBalancerConfigNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNodeBalancerConfigNode`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.DeleteNodeBalancerConfigNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access | 
**nodeId** | **int32** | The ID of the Node to access | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerConfigNodeRequest struct via the builder pattern


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


## GetNodeBalancer

> NodeBalancer GetNodeBalancer(ctx, nodeBalancerId).Execute()

NodeBalancer View



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancer(context.Background(), nodeBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancer`: NodeBalancer
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfig

> NodeBalancerConfig GetNodeBalancerConfig(ctx, nodeBalancerId, configId).Execute()

Config View



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the config to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancerConfig(context.Background(), nodeBalancerId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancerConfig`: NodeBalancerConfig
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeBalancerConfig**](NodeBalancerConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfigNodes

> InlineResponse20046 GetNodeBalancerConfigNodes(ctx, nodeBalancerId, configId).Page(page).PageSize(pageSize).Execute()

Nodes List



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the NodeBalancer config to access.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancerConfigNodes(context.Background(), nodeBalancerId, configId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancerConfigNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancerConfigNodes`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancerConfigNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the NodeBalancer config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfigs

> InlineResponse20045 GetNodeBalancerConfigs(ctx, nodeBalancerId).Page(page).PageSize(pageSize).Execute()

Configs List



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancerConfigs(context.Background(), nodeBalancerId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancerConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancerConfigs`: InlineResponse20045
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancerConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20045**](InlineResponse20045.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerNode

> NodeBalancerNode GetNodeBalancerNode(ctx, nodeBalancerId, configId, nodeId).Execute()

Node View



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the Config to access
    nodeId := int32(56) // int32 | The ID of the Node to access

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancerNode(context.Background(), nodeBalancerId, configId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancerNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancerNode`: NodeBalancerNode
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access | 
**nodeId** | **int32** | The ID of the Node to access | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NodeBalancerNode**](NodeBalancerNode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancers

> InlineResponse20044 GetNodeBalancers(ctx).Page(page).PageSize(pageSize).Execute()

NodeBalancers List



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
    resp, r, err := api_client.NodeBalancersApi.GetNodeBalancers(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.GetNodeBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeBalancers`: InlineResponse20044
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.GetNodeBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20044**](InlineResponse20044.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodebalancersNodeBalancerIdStatsGet

> NodeBalancerStats NodebalancersNodeBalancerIdStatsGet(ctx, nodeBalancerId).Execute()

NodeBalancer Statistics View



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.NodebalancersNodeBalancerIdStatsGet(context.Background(), nodeBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.NodebalancersNodeBalancerIdStatsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NodebalancersNodeBalancerIdStatsGet`: NodeBalancerStats
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.NodebalancersNodeBalancerIdStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodebalancersNodeBalancerIdStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeBalancerStats**](NodeBalancerStats.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildNodeBalancerConfig

> NodeBalancer RebuildNodeBalancerConfig(ctx, nodeBalancerId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Config Rebuild



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the Config to access.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the NodeBalancer Config to rebuild. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.RebuildNodeBalancerConfig(context.Background(), nodeBalancerId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.RebuildNodeBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebuildNodeBalancerConfig`: NodeBalancer
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.RebuildNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the NodeBalancer Config to rebuild.  | 

### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeBalancer

> NodeBalancer UpdateNodeBalancer(ctx, nodeBalancerId).NodeBalancer(nodeBalancer).Execute()

NodeBalancer Update



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    nodeBalancer := *openapiclient.NewNodeBalancer() // NodeBalancer | The information to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.UpdateNodeBalancer(context.Background(), nodeBalancerId).NodeBalancer(nodeBalancer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.UpdateNodeBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeBalancer`: NodeBalancer
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.UpdateNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBalancer** | [**NodeBalancer**](NodeBalancer.md) | The information to update. | 

### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeBalancerConfig

> NodeBalancerConfig UpdateNodeBalancerConfig(ctx, nodeBalancerId, configId).NodeBalancerConfig(nodeBalancerConfig).Execute()

Config Update



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the config to access.
    nodeBalancerConfig := *openapiclient.NewNodeBalancerConfig() // NodeBalancerConfig | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.UpdateNodeBalancerConfig(context.Background(), nodeBalancerId, configId).NodeBalancerConfig(nodeBalancerConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.UpdateNodeBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeBalancerConfig`: NodeBalancerConfig
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.UpdateNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeBalancerConfig** | [**NodeBalancerConfig**](NodeBalancerConfig.md) | The fields to update. | 

### Return type

[**NodeBalancerConfig**](NodeBalancerConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeBalancerNode

> NodeBalancerNode UpdateNodeBalancerNode(ctx, nodeBalancerId, configId, nodeId).NodeBalancerNode(nodeBalancerNode).Execute()

Node Update



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
    nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
    configId := int32(56) // int32 | The ID of the Config to access
    nodeId := int32(56) // int32 | The ID of the Node to access
    nodeBalancerNode := *openapiclient.NewNodeBalancerNode() // NodeBalancerNode | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeBalancersApi.UpdateNodeBalancerNode(context.Background(), nodeBalancerId, configId, nodeId).NodeBalancerNode(nodeBalancerNode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersApi.UpdateNodeBalancerNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeBalancerNode`: NodeBalancerNode
    fmt.Fprintf(os.Stdout, "Response from `NodeBalancersApi.UpdateNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access | 
**nodeId** | **int32** | The ID of the Node to access | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeBalancerNode** | [**NodeBalancerNode**](NodeBalancerNode.md) | The fields to update. | 

### Return type

[**NodeBalancerNode**](NodeBalancerNode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

