# \DomainsApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDomain**](DomainsApi.md#CloneDomain) | **Post** /domains/{domainId}/clone | Domain Clone
[**CreateDomain**](DomainsApi.md#CreateDomain) | **Post** /domains | Domain Create
[**CreateDomainRecord**](DomainsApi.md#CreateDomainRecord) | **Post** /domains/{domainId}/records | Domain Record Create
[**DeleteDomain**](DomainsApi.md#DeleteDomain) | **Delete** /domains/{domainId} | Domain Delete
[**DeleteDomainRecord**](DomainsApi.md#DeleteDomainRecord) | **Delete** /domains/{domainId}/records/{recordId} | Domain Record Delete
[**GetDomain**](DomainsApi.md#GetDomain) | **Get** /domains/{domainId} | Domain View
[**GetDomainRecord**](DomainsApi.md#GetDomainRecord) | **Get** /domains/{domainId}/records/{recordId} | Domain Record View
[**GetDomainRecords**](DomainsApi.md#GetDomainRecords) | **Get** /domains/{domainId}/records | Domain Records List
[**GetDomains**](DomainsApi.md#GetDomains) | **Get** /domains | Domains List
[**ImportDomain**](DomainsApi.md#ImportDomain) | **Post** /domains/import | Domain Import
[**UpdateDomain**](DomainsApi.md#UpdateDomain) | **Put** /domains/{domainId} | Domain Update
[**UpdateDomainRecord**](DomainsApi.md#UpdateDomainRecord) | **Put** /domains/{domainId}/records/{recordId} | Domain Record Update



## CloneDomain

> Domain CloneDomain(ctx, domainId).InlineObject1(inlineObject1).Execute()

Domain Clone



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
    domainId := "domainId_example" // string | ID of the Domain to clone.
    inlineObject1 := *openapiclient.NewInlineObject1("example.com") // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.CloneDomain(context.Background(), domainId).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CloneDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDomain`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CloneDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | ID of the Domain to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomain

> Domain CreateDomain(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Domain Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the domain you are registering.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.CreateDomain(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the domain you are registering. | 

### Return type

[**Domain**](Domain.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomainRecord

> DomainRecord CreateDomainRecord(ctx, domainId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Domain Record Create



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
    domainId := int32(56) // int32 | The ID of the Domain we are accessing Records for.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the new Record to add. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.CreateDomainRecord(context.Background(), domainId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CreateDomainRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainRecord`: DomainRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CreateDomainRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain we are accessing Records for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the new Record to add.  | 

### Return type

[**DomainRecord**](DomainRecord.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> map[string]interface{} DeleteDomain(ctx, domainId).Execute()

Domain Delete



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
    domainId := int32(56) // int32 | The ID of the Domain to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.DeleteDomain(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomain`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


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


## DeleteDomainRecord

> map[string]interface{} DeleteDomainRecord(ctx, domainId, recordId).Execute()

Domain Record Delete



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
    domainId := int32(56) // int32 | The ID of the Domain whose Record you are accessing.
    recordId := int32(56) // int32 | The ID of the Record you are accessing.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.DeleteDomainRecord(context.Background(), domainId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DeleteDomainRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomainRecord`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.DeleteDomainRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain whose Record you are accessing. | 
**recordId** | **int32** | The ID of the Record you are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRecordRequest struct via the builder pattern


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


## GetDomain

> Domain GetDomain(ctx, domainId).Execute()

Domain View



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
    domainId := int32(56) // int32 | The ID of the Domain to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomain(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainRecord

> DomainRecord GetDomainRecord(ctx, domainId, recordId).Execute()

Domain Record View



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
    domainId := int32(56) // int32 | The ID of the Domain whose Record you are accessing.
    recordId := int32(56) // int32 | The ID of the Record you are accessing.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomainRecord(context.Background(), domainId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomainRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainRecord`: DomainRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomainRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain whose Record you are accessing. | 
**recordId** | **int32** | The ID of the Record you are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DomainRecord**](DomainRecord.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainRecords

> InlineResponse20013 GetDomainRecords(ctx, domainId).Page(page).PageSize(pageSize).Execute()

Domain Records List



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
    domainId := int32(56) // int32 | The ID of the Domain we are accessing Records for.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomainRecords(context.Background(), domainId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomainRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainRecords`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomainRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain we are accessing Records for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> InlineResponse20012 GetDomains(ctx).Page(page).PageSize(pageSize).Execute()

Domains List



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
    resp, r, err := api_client.DomainsApi.GetDomains(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomains`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDomain

> Domain ImportDomain(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Domain Import



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the Domain to import. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.ImportDomain(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.ImportDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportDomain`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.ImportDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the Domain to import. | 

### Return type

[**Domain**](Domain.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomain

> Domain UpdateDomain(ctx, domainId).Domain(domain).Execute()

Domain Update



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
    domainId := int32(56) // int32 | The ID of the Domain to access.
    domain := *openapiclient.NewDomain() // Domain | The Domain information to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.UpdateDomain(context.Background(), domainId).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UpdateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomain`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | [**Domain**](Domain.md) | The Domain information to update. | 

### Return type

[**Domain**](Domain.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainRecord

> DomainRecord UpdateDomainRecord(ctx, domainId, recordId).DomainRecord(domainRecord).Execute()

Domain Record Update



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
    domainId := int32(56) // int32 | The ID of the Domain whose Record you are accessing.
    recordId := int32(56) // int32 | The ID of the Record you are accessing.
    domainRecord := *openapiclient.NewDomainRecord() // DomainRecord | The values to change.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.UpdateDomainRecord(context.Background(), domainId, recordId).DomainRecord(domainRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UpdateDomainRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainRecord`: DomainRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.UpdateDomainRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | The ID of the Domain whose Record you are accessing. | 
**recordId** | **int32** | The ID of the Record you are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainRecord** | [**DomainRecord**](DomainRecord.md) | The values to change. | 

### Return type

[**DomainRecord**](DomainRecord.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

