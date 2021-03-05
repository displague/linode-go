# \LinodeKubernetesEngineLKEApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLKECluster**](LinodeKubernetesEngineLKEApi.md#CreateLKECluster) | **Post** /lke/clusters | Kubernetes Cluster Create
[**DeleteLKECluster**](LinodeKubernetesEngineLKEApi.md#DeleteLKECluster) | **Delete** /lke/clusters/{clusterId} | Kubernetes Cluster Delete
[**DeleteLKENodePool**](LinodeKubernetesEngineLKEApi.md#DeleteLKENodePool) | **Delete** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool Delete
[**GetLKECluster**](LinodeKubernetesEngineLKEApi.md#GetLKECluster) | **Get** /lke/clusters/{clusterId} | Kubernetes Cluster View
[**GetLKEClusterAPIEndpoints**](LinodeKubernetesEngineLKEApi.md#GetLKEClusterAPIEndpoints) | **Get** /lke/clusters/{clusterId}/api-endpoints | Kubernetes API Endpoints List
[**GetLKEClusterKubeconfig**](LinodeKubernetesEngineLKEApi.md#GetLKEClusterKubeconfig) | **Get** /lke/clusters/{clusterId}/kubeconfig | Kubeconfig View
[**GetLKEClusterNode**](LinodeKubernetesEngineLKEApi.md#GetLKEClusterNode) | **Get** /lke/clusters/{clusterId}/nodes/{nodeId} | Node View
[**GetLKEClusterPools**](LinodeKubernetesEngineLKEApi.md#GetLKEClusterPools) | **Get** /lke/clusters/{clusterId}/pools | Node Pools List
[**GetLKEClusters**](LinodeKubernetesEngineLKEApi.md#GetLKEClusters) | **Get** /lke/clusters | Kubernetes Clusters List
[**GetLKENodePool**](LinodeKubernetesEngineLKEApi.md#GetLKENodePool) | **Get** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool View
[**GetLKEVersion**](LinodeKubernetesEngineLKEApi.md#GetLKEVersion) | **Get** /lke/versions/{version} | Kubernetes Version View
[**GetLKEVersions**](LinodeKubernetesEngineLKEApi.md#GetLKEVersions) | **Get** /lke/versions | Kubernetes Versions List
[**PostLKEClusterNodeRecycle**](LinodeKubernetesEngineLKEApi.md#PostLKEClusterNodeRecycle) | **Post** /lke/clusters/{clusterId}/nodes/{nodeId}/recycle | Node Recycle
[**PostLKEClusterPoolRecycle**](LinodeKubernetesEngineLKEApi.md#PostLKEClusterPoolRecycle) | **Post** /lke/clusters/{clusterId}/pools/{poolId}/recycle | Node Pool Recycle
[**PostLKEClusterPools**](LinodeKubernetesEngineLKEApi.md#PostLKEClusterPools) | **Post** /lke/clusters/{clusterId}/pools | Node Pool Create
[**PostLKEClusterRecycle**](LinodeKubernetesEngineLKEApi.md#PostLKEClusterRecycle) | **Post** /lke/clusters/{clusterId}/recycle | Cluster Nodes Recycle
[**PutLKECluster**](LinodeKubernetesEngineLKEApi.md#PutLKECluster) | **Put** /lke/clusters/{clusterId} | Kubernetes Cluster Update
[**PutLKENodePool**](LinodeKubernetesEngineLKEApi.md#PutLKENodePool) | **Put** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool Update



## CreateLKECluster

> LKECluster CreateLKECluster(ctx).InlineObject10(inlineObject10).Execute()

Kubernetes Cluster Create



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
    inlineObject10 := *openapiclient.NewInlineObject10("TODO", *openapiclient.NewRegion(), "TODO", []openapiclient.LKENodePoolRequestBody{*openapiclient.NewLKENodePoolRequestBody()}) // InlineObject10 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.CreateLKECluster(context.Background()).InlineObject10(inlineObject10).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.CreateLKECluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLKECluster`: LKECluster
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.CreateLKECluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLKEClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject10** | [**InlineObject10**](InlineObject10.md) |  | 

### Return type

[**LKECluster**](LKECluster.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLKECluster

> map[string]interface{} DeleteLKECluster(ctx, clusterId).Execute()

Kubernetes Cluster Delete



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.DeleteLKECluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.DeleteLKECluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLKECluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.DeleteLKECluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLKEClusterRequest struct via the builder pattern


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


## DeleteLKENodePool

> map[string]interface{} DeleteLKENodePool(ctx, clusterId, poolId).Execute()

Node Pool Delete



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
    poolId := int32(56) // int32 | ID of the Pool to look up

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.DeleteLKENodePool(context.Background(), clusterId, poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.DeleteLKENodePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLKENodePool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.DeleteLKENodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLKENodePoolRequest struct via the builder pattern


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


## GetLKECluster

> LKECluster GetLKECluster(ctx, clusterId).Execute()

Kubernetes Cluster View



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKECluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKECluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKECluster`: LKECluster
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKECluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LKECluster**](LKECluster.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEClusterAPIEndpoints

> InlineResponse20028 GetLKEClusterAPIEndpoints(ctx, clusterId).Execute()

Kubernetes API Endpoints List



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEClusterAPIEndpoints(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEClusterAPIEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEClusterAPIEndpoints`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEClusterAPIEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClusterAPIEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEClusterKubeconfig

> InlineResponse20029 GetLKEClusterKubeconfig(ctx, clusterId).Execute()

Kubeconfig View



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEClusterKubeconfig(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEClusterKubeconfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEClusterKubeconfig`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEClusterNode

> InlineResponse20027 GetLKEClusterNode(ctx, clusterId, nodeId).Execute()

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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster containing the Node.
    nodeId := "nodeId_example" // string | ID of the Node to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEClusterNode(context.Background(), clusterId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEClusterNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEClusterNode`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEClusterNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster containing the Node. | 
**nodeId** | **string** | ID of the Node to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClusterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEClusterPools

> InlineResponse20025 GetLKEClusterPools(ctx, clusterId).Execute()

Node Pools List



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEClusterPools(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEClusterPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEClusterPools`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEClusterPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClusterPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEClusters

> InlineResponse20024 GetLKEClusters(ctx).Execute()

Kubernetes Clusters List



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
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEClusters`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEClustersRequest struct via the builder pattern


### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKENodePool

> InlineResponse20026 GetLKENodePool(ctx, clusterId, poolId).Execute()

Node Pool View



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
    poolId := int32(56) // int32 | ID of the Pool to look up

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKENodePool(context.Background(), clusterId, poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKENodePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKENodePool`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKENodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKENodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEVersion

> LKEVersion GetLKEVersion(ctx, version).Execute()

Kubernetes Version View



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
    version := "version_example" // string | The LKE version to view.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEVersion(context.Background(), version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEVersion`: LKEVersion
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | The LKE version to view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LKEVersion**](LKEVersion.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLKEVersions

> InlineResponse20030 GetLKEVersions(ctx).Execute()

Kubernetes Versions List



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
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.GetLKEVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.GetLKEVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLKEVersions`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.GetLKEVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLKEVersionsRequest struct via the builder pattern


### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLKEClusterNodeRecycle

> map[string]interface{} PostLKEClusterNodeRecycle(ctx, clusterId, nodeId).Execute()

Node Recycle



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster containing the Node.
    nodeId := "nodeId_example" // string | ID of the Node to be recycled.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PostLKEClusterNodeRecycle(context.Background(), clusterId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PostLKEClusterNodeRecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLKEClusterNodeRecycle`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PostLKEClusterNodeRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster containing the Node. | 
**nodeId** | **string** | ID of the Node to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLKEClusterNodeRecycleRequest struct via the builder pattern


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


## PostLKEClusterPoolRecycle

> map[string]interface{} PostLKEClusterPoolRecycle(ctx, clusterId, poolId).Execute()

Node Pool Recycle



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster this Node Pool is attached to.
    poolId := int32(56) // int32 | ID of the Node Pool to be recycled.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PostLKEClusterPoolRecycle(context.Background(), clusterId, poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PostLKEClusterPoolRecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLKEClusterPoolRecycle`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PostLKEClusterPoolRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster this Node Pool is attached to. | 
**poolId** | **int32** | ID of the Node Pool to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLKEClusterPoolRecycleRequest struct via the builder pattern


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


## PostLKEClusterPools

> InlineResponse20026 PostLKEClusterPools(ctx, clusterId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Node Pool Create



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Configuration for the Node Pool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PostLKEClusterPools(context.Background(), clusterId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PostLKEClusterPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLKEClusterPools`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PostLKEClusterPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLKEClusterPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Configuration for the Node Pool | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLKEClusterRecycle

> map[string]interface{} PostLKEClusterRecycle(ctx, clusterId).Execute()

Cluster Nodes Recycle



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster which contains nodes to be recycled.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PostLKEClusterRecycle(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PostLKEClusterRecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLKEClusterRecycle`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PostLKEClusterRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster which contains nodes to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLKEClusterRecycleRequest struct via the builder pattern


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


## PutLKECluster

> map[string]interface{} PutLKECluster(ctx, clusterId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Kubernetes Cluster Update



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The fields to update the Kubernetes cluster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PutLKECluster(context.Background(), clusterId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PutLKECluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLKECluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PutLKECluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLKEClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The fields to update the Kubernetes cluster. | 

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


## PutLKENodePool

> InlineResponse20026 PutLKENodePool(ctx, clusterId, poolId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Node Pool Update



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
    clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
    poolId := int32(56) // int32 | ID of the Pool to look up
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The fields to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeKubernetesEngineLKEApi.PutLKENodePool(context.Background(), clusterId, poolId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEApi.PutLKENodePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLKENodePool`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEApi.PutLKENodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLKENodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The fields to update | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

