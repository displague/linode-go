# \ObjectStorageApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelObjectStorage**](ObjectStorageApi.md#CancelObjectStorage) | **Post** /object-storage/cancel | Object Storage Cancel
[**CreateObjectStorageBucket**](ObjectStorageApi.md#CreateObjectStorageBucket) | **Post** /object-storage/buckets | Object Storage Bucket Create
[**CreateObjectStorageKeys**](ObjectStorageApi.md#CreateObjectStorageKeys) | **Post** /object-storage/keys | Object Storage Key Create
[**CreateObjectStorageObjectURL**](ObjectStorageApi.md#CreateObjectStorageObjectURL) | **Post** /object-storage/buckets/{clusterId}/{bucket}/object-url | Object Storage Object URL Create
[**CreateObjectStorageSSL**](ObjectStorageApi.md#CreateObjectStorageSSL) | **Post** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert Upload
[**DeleteObjectStorageBucket**](ObjectStorageApi.md#DeleteObjectStorageBucket) | **Delete** /object-storage/buckets/{clusterId}/{bucket} | Object Storage Bucket Remove
[**DeleteObjectStorageKey**](ObjectStorageApi.md#DeleteObjectStorageKey) | **Delete** /object-storage/keys/{keyId} | Object Storage Key Revoke
[**DeleteObjectStorageSSL**](ObjectStorageApi.md#DeleteObjectStorageSSL) | **Delete** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert Delete
[**GetObjectStorageBucket**](ObjectStorageApi.md#GetObjectStorageBucket) | **Get** /object-storage/buckets/{clusterId}/{bucket} | Object Storage Bucket View
[**GetObjectStorageBucketContent**](ObjectStorageApi.md#GetObjectStorageBucketContent) | **Get** /object-storage/buckets/{clusterId}/{bucket}/object-list | Object Storage Bucket Contents List
[**GetObjectStorageBucketinCluster**](ObjectStorageApi.md#GetObjectStorageBucketinCluster) | **Get** /object-storage/buckets/{clusterId} | Object Storage Buckets in Cluster List
[**GetObjectStorageBuckets**](ObjectStorageApi.md#GetObjectStorageBuckets) | **Get** /object-storage/buckets | Object Storage Buckets List
[**GetObjectStorageCluster**](ObjectStorageApi.md#GetObjectStorageCluster) | **Get** /object-storage/clusters/{clusterId} | Cluster View
[**GetObjectStorageClusters**](ObjectStorageApi.md#GetObjectStorageClusters) | **Get** /object-storage/clusters | Clusters List
[**GetObjectStorageKey**](ObjectStorageApi.md#GetObjectStorageKey) | **Get** /object-storage/keys/{keyId} | Object Storage Key View
[**GetObjectStorageKeys**](ObjectStorageApi.md#GetObjectStorageKeys) | **Get** /object-storage/keys | Object Storage Keys List
[**GetObjectStorageSSL**](ObjectStorageApi.md#GetObjectStorageSSL) | **Get** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert View
[**GetObjectStorageTransfer**](ObjectStorageApi.md#GetObjectStorageTransfer) | **Get** /object-storage/transfer | Object Storage Transfer View
[**ModifyObjectStorageBucketAccess**](ObjectStorageApi.md#ModifyObjectStorageBucketAccess) | **Post** /object-storage/buckets/{clusterId}/{bucket}/access | Object Storage Bucket Access Modify
[**UpdateObjectStorageBucketACL**](ObjectStorageApi.md#UpdateObjectStorageBucketACL) | **Put** /object-storage/buckets/{clusterId}/{bucket}/object-acl | Object Storage Object ACL Config Update
[**UpdateObjectStorageBucketAccess**](ObjectStorageApi.md#UpdateObjectStorageBucketAccess) | **Put** /object-storage/buckets/{clusterId}/{bucket}/access | Object Storage Bucket Access Update
[**UpdateObjectStorageKey**](ObjectStorageApi.md#UpdateObjectStorageKey) | **Put** /object-storage/keys/{keyId} | Object Storage Key Update
[**ViewObjectStorageBucketACL**](ObjectStorageApi.md#ViewObjectStorageBucketACL) | **Get** /object-storage/buckets/{clusterId}/{bucket}/object-acl | Object Storage Object ACL Config View



## CancelObjectStorage

> map[string]interface{} CancelObjectStorage(ctx).Execute()

Object Storage Cancel



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
    resp, r, err := api_client.ObjectStorageApi.CancelObjectStorage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.CancelObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelObjectStorage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.CancelObjectStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancelObjectStorageRequest struct via the builder pattern


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


## CreateObjectStorageBucket

> ObjectStorageBucket CreateObjectStorageBucket(ctx).InlineObject13(inlineObject13).Execute()

Object Storage Bucket Create



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
    inlineObject13 := *openapiclient.NewInlineObject13("example-bucket", "us-east-1") // InlineObject13 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.CreateObjectStorageBucket(context.Background()).InlineObject13(inlineObject13).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.CreateObjectStorageBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectStorageBucket`: ObjectStorageBucket
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.CreateObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject13** | [**InlineObject13**](InlineObject13.md) |  | 

### Return type

[**ObjectStorageBucket**](ObjectStorageBucket.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageKeys

> ObjectStorageKey CreateObjectStorageKeys(ctx).ObjectStorageKey(objectStorageKey).Execute()

Object Storage Key Create



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
    objectStorageKey := *openapiclient.NewObjectStorageKey() // ObjectStorageKey | The label of the key to create. This is used to identify the created key.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.CreateObjectStorageKeys(context.Background()).ObjectStorageKey(objectStorageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.CreateObjectStorageKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectStorageKeys`: ObjectStorageKey
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.CreateObjectStorageKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectStorageKey** | [**ObjectStorageKey**](ObjectStorageKey.md) | The label of the key to create. This is used to identify the created key.  | 

### Return type

[**ObjectStorageKey**](ObjectStorageKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageObjectURL

> map[string]interface{} CreateObjectStorageObjectURL(ctx, clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Object Storage Object URL Create



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the request to sign. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.CreateObjectStorageObjectURL(context.Background(), clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.CreateObjectStorageObjectURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectStorageObjectURL`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.CreateObjectStorageObjectURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageObjectURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the request to sign. | 

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


## CreateObjectStorageSSL

> ObjectStorageSSLResponse CreateObjectStorageSSL(ctx, clusterId, bucket).ObjectStorageSSL(objectStorageSSL).Execute()

Object Storage TLS/SSL Cert Upload



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    objectStorageSSL := *openapiclient.NewObjectStorageSSL(""-----BEGIN CERTIFICATE----- MIIFTTCCAzWgAwIBAgIURwtqMl ... -----END CERTIFICATE-----"
", ""-----BEGIN PRIVATE KEY----- MIIEvgIBADANBgkqhkiG9w0BAQE ... -----END PRIVATE KEY-----"
") // ObjectStorageSSL | Upload this TLS/SSL certificate with its corresponding secret key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.CreateObjectStorageSSL(context.Background(), clusterId, bucket).ObjectStorageSSL(objectStorageSSL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.CreateObjectStorageSSL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectStorageSSL`: ObjectStorageSSLResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.CreateObjectStorageSSL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageSSLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **objectStorageSSL** | [**ObjectStorageSSL**](ObjectStorageSSL.md) | Upload this TLS/SSL certificate with its corresponding secret key. | 

### Return type

[**ObjectStorageSSLResponse**](ObjectStorageSSLResponse.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageBucket

> map[string]interface{} DeleteObjectStorageBucket(ctx, clusterId, bucket).Execute()

Object Storage Bucket Remove



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.DeleteObjectStorageBucket(context.Background(), clusterId, bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.DeleteObjectStorageBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObjectStorageBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.DeleteObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageBucketRequest struct via the builder pattern


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


## DeleteObjectStorageKey

> map[string]interface{} DeleteObjectStorageKey(ctx, keyId).Execute()

Object Storage Key Revoke



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
    keyId := int32(56) // int32 | The key to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.DeleteObjectStorageKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.DeleteObjectStorageKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObjectStorageKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.DeleteObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageKeyRequest struct via the builder pattern


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


## DeleteObjectStorageSSL

> map[string]interface{} DeleteObjectStorageSSL(ctx, clusterId, bucket).Execute()

Object Storage TLS/SSL Cert Delete



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.DeleteObjectStorageSSL(context.Background(), clusterId, bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.DeleteObjectStorageSSL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObjectStorageSSL`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.DeleteObjectStorageSSL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageSSLRequest struct via the builder pattern


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


## GetObjectStorageBucket

> ObjectStorageBucket GetObjectStorageBucket(ctx, clusterId, bucket).Execute()

Object Storage Bucket View



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageBucket(context.Background(), clusterId, bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageBucket`: ObjectStorageBucket
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectStorageBucket**](ObjectStorageBucket.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketContent

> map[string]interface{} GetObjectStorageBucketContent(ctx, clusterId, bucket).Marker(marker).Delimiter(delimiter).Prefix(prefix).PageSize(pageSize).Execute()

Object Storage Bucket Contents List



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    marker := "marker_example" // string | The \"marker\" for this request, which can be used to paginate through large buckets. Its value should be the value of the `next_marker` property returned with the last page. Listing bucket contents *does not* support arbitrary page access. See the `next_marker` property in the responses section for more details.  (optional)
    delimiter := "delimiter_example" // string | The delimiter for object names; if given, object names will be returned up to the first occurrence of this character. This is most commonly used with the `/` character to allow bucket transversal in a manner similar to a filesystem, however any delimiter may be used. Use in conjunction with `prefix` to see object names past the first occurrence of the delimiter.  (optional)
    prefix := "prefix_example" // string | Filters objects returned to only those whose name start with the given prefix. Commonly used in conjunction with `delimiter` to allow transversal of bucket contents in a manner similar to a filesystem.  (optional)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageBucketContent(context.Background(), clusterId, bucket).Marker(marker).Delimiter(delimiter).Prefix(prefix).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageBucketContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageBucketContent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageBucketContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marker** | **string** | The \&quot;marker\&quot; for this request, which can be used to paginate through large buckets. Its value should be the value of the &#x60;next_marker&#x60; property returned with the last page. Listing bucket contents *does not* support arbitrary page access. See the &#x60;next_marker&#x60; property in the responses section for more details.  | 
 **delimiter** | **string** | The delimiter for object names; if given, object names will be returned up to the first occurrence of this character. This is most commonly used with the &#x60;/&#x60; character to allow bucket transversal in a manner similar to a filesystem, however any delimiter may be used. Use in conjunction with &#x60;prefix&#x60; to see object names past the first occurrence of the delimiter.  | 
 **prefix** | **string** | Filters objects returned to only those whose name start with the given prefix. Commonly used in conjunction with &#x60;delimiter&#x60; to allow transversal of bucket contents in a manner similar to a filesystem.  | 
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


## GetObjectStorageBucketinCluster

> InlineResponse20047 GetObjectStorageBucketinCluster(ctx, clusterId).Execute()

Object Storage Buckets in Cluster List



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageBucketinCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageBucketinCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageBucketinCluster`: InlineResponse20047
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageBucketinCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketinClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBuckets

> InlineResponse20047 GetObjectStorageBuckets(ctx).Execute()

Object Storage Buckets List



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
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageBuckets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageBuckets`: InlineResponse20047
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageBuckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketsRequest struct via the builder pattern


### Return type

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageCluster

> ObjectStorageCluster GetObjectStorageCluster(ctx, clusterId).Execute()

Cluster View



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
    clusterId := "clusterId_example" // string | The ID of the Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageCluster`: ObjectStorageCluster
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageCluster**](ObjectStorageCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageClusters

> InlineResponse20049 GetObjectStorageClusters(ctx).Execute()

Clusters List



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
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageClusters`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageClustersRequest struct via the builder pattern


### Return type

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageKey

> ObjectStorageKey GetObjectStorageKey(ctx, keyId).Execute()

Object Storage Key View



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
    keyId := int32(56) // int32 | The key to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageKey`: ObjectStorageKey
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageKey**](ObjectStorageKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageKeys

> InlineResponse20050 GetObjectStorageKeys(ctx).Execute()

Object Storage Keys List



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
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageKeys`: InlineResponse20050
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageKeysRequest struct via the builder pattern


### Return type

[**InlineResponse20050**](InlineResponse20050.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageSSL

> ObjectStorageSSLResponse GetObjectStorageSSL(ctx, clusterId, bucket).Execute()

Object Storage TLS/SSL Cert View



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageSSL(context.Background(), clusterId, bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageSSL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageSSL`: ObjectStorageSSLResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageSSL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageSSLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectStorageSSLResponse**](ObjectStorageSSLResponse.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageTransfer

> map[string]interface{} GetObjectStorageTransfer(ctx).Execute()

Object Storage Transfer View



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
    resp, r, err := api_client.ObjectStorageApi.GetObjectStorageTransfer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.GetObjectStorageTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.GetObjectStorageTransfer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageTransferRequest struct via the builder pattern


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


## ModifyObjectStorageBucketAccess

> map[string]interface{} ModifyObjectStorageBucketAccess(ctx, clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Object Storage Bucket Access Modify



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The changes to make to the bucket's access controls. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.ModifyObjectStorageBucketAccess(context.Background(), clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.ModifyObjectStorageBucketAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyObjectStorageBucketAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.ModifyObjectStorageBucketAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyObjectStorageBucketAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The changes to make to the bucket&#39;s access controls. | 

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


## UpdateObjectStorageBucketACL

> InlineResponse20048 UpdateObjectStorageBucketACL(ctx, clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Object Storage Object ACL Config Update



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The changes to make to this Object's access controls. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.UpdateObjectStorageBucketACL(context.Background(), clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.UpdateObjectStorageBucketACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectStorageBucketACL`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.UpdateObjectStorageBucketACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageBucketACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The changes to make to this Object&#39;s access controls. | 

### Return type

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageBucketAccess

> map[string]interface{} UpdateObjectStorageBucketAccess(ctx, clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Object Storage Bucket Access Update



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The changes to make to the bucket's access controls. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.UpdateObjectStorageBucketAccess(context.Background(), clusterId, bucket).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.UpdateObjectStorageBucketAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectStorageBucketAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.UpdateObjectStorageBucketAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageBucketAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The changes to make to the bucket&#39;s access controls. | 

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


## UpdateObjectStorageKey

> ObjectStorageKey UpdateObjectStorageKey(ctx, keyId).InlineObject14(inlineObject14).Execute()

Object Storage Key Update



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
    keyId := int32(56) // int32 | The key to look up.
    inlineObject14 := *openapiclient.NewInlineObject14() // InlineObject14 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.UpdateObjectStorageKey(context.Background(), keyId).InlineObject14(inlineObject14).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.UpdateObjectStorageKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectStorageKey`: ObjectStorageKey
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.UpdateObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject14** | [**InlineObject14**](InlineObject14.md) |  | 

### Return type

[**ObjectStorageKey**](ObjectStorageKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewObjectStorageBucketACL

> InlineResponse20048 ViewObjectStorageBucketACL(ctx, clusterId, bucket).Name(name).Execute()

Object Storage Object ACL Config View



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
    clusterId := "clusterId_example" // string | The ID of the cluster this bucket exists in.
    bucket := "bucket_example" // string | The bucket name.
    name := "name_example" // string | The `name` of the object for which to retrieve its Access Control List (ACL). Use the [Object Storage Bucket Contents List](/docs/api/object-storage/#object-storage-bucket-contents-list) endpoint to access all object names in a bucket. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStorageApi.ViewObjectStorageBucketACL(context.Background(), clusterId, bucket).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageApi.ViewObjectStorageBucketACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewObjectStorageBucketACL`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `ObjectStorageApi.ViewObjectStorageBucketACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster this bucket exists in. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewObjectStorageBucketACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | The &#x60;name&#x60; of the object for which to retrieve its Access Control List (ACL). Use the [Object Storage Bucket Contents List](/docs/api/object-storage/#object-storage-bucket-contents-list) endpoint to access all object names in a bucket.  | 

### Return type

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

