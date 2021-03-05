# \ManagedApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedContact**](ManagedApi.md#CreateManagedContact) | **Post** /managed/contacts | Managed Contact Create
[**CreateManagedCredential**](ManagedApi.md#CreateManagedCredential) | **Post** /managed/credentials | Managed Credential Create
[**CreateManagedService**](ManagedApi.md#CreateManagedService) | **Post** /managed/services | Managed Service Create
[**DeleteManagedContact**](ManagedApi.md#DeleteManagedContact) | **Delete** /managed/contacts/{contactId} | Managed Contact Delete
[**DeleteManagedCredential**](ManagedApi.md#DeleteManagedCredential) | **Post** /managed/credentials/{credentialId}/revoke | Managed Credential Delete
[**DeleteManagedService**](ManagedApi.md#DeleteManagedService) | **Delete** /managed/services/{serviceId} | Managed Service Delete
[**DisableManagedService**](ManagedApi.md#DisableManagedService) | **Post** /managed/services/{serviceId}/disable | Managed Service Disable
[**EnableManagedService**](ManagedApi.md#EnableManagedService) | **Post** /managed/services/{serviceId}/enable | Managed Service Enable
[**GetManagedContact**](ManagedApi.md#GetManagedContact) | **Get** /managed/contacts/{contactId} | Managed Contact View
[**GetManagedContacts**](ManagedApi.md#GetManagedContacts) | **Get** /managed/contacts | Managed Contacts List
[**GetManagedCredential**](ManagedApi.md#GetManagedCredential) | **Get** /managed/credentials/{credentialId} | Managed Credential View
[**GetManagedCredentials**](ManagedApi.md#GetManagedCredentials) | **Get** /managed/credentials | Managed Credentials List
[**GetManagedIssue**](ManagedApi.md#GetManagedIssue) | **Get** /managed/issues/{issueId} | Managed Issue View
[**GetManagedIssues**](ManagedApi.md#GetManagedIssues) | **Get** /managed/issues | Managed Issues List
[**GetManagedLinodeSetting**](ManagedApi.md#GetManagedLinodeSetting) | **Get** /managed/linode-settings/{linodeId} | Linode&#39;s Managed Settings View
[**GetManagedLinodeSettings**](ManagedApi.md#GetManagedLinodeSettings) | **Get** /managed/linode-settings | Managed Linode Settings List
[**GetManagedService**](ManagedApi.md#GetManagedService) | **Get** /managed/services/{serviceId} | Managed Service View
[**GetManagedServices**](ManagedApi.md#GetManagedServices) | **Get** /managed/services | Managed Services List
[**GetManagedStats**](ManagedApi.md#GetManagedStats) | **Get** /managed/stats | Managed Stats List
[**UpdateManagedContact**](ManagedApi.md#UpdateManagedContact) | **Put** /managed/contacts/{contactId} | Managed Contact Update
[**UpdateManagedCredential**](ManagedApi.md#UpdateManagedCredential) | **Put** /managed/credentials/{credentialId} | Managed Credential Update
[**UpdateManagedCredentialUsernamePassword**](ManagedApi.md#UpdateManagedCredentialUsernamePassword) | **Post** /managed/credentials/{credentialId}/update | Managed Credential Username and Password Update
[**UpdateManagedLinodeSetting**](ManagedApi.md#UpdateManagedLinodeSetting) | **Put** /managed/linode-settings/{linodeId} | Linode&#39;s Managed Settings Update
[**UpdateManagedService**](ManagedApi.md#UpdateManagedService) | **Put** /managed/services/{serviceId} | Managed Service Update
[**ViewManagedSSHKey**](ManagedApi.md#ViewManagedSSHKey) | **Get** /managed/credentials/sshkey | Managed SSH Key View



## CreateManagedContact

> ManagedContact CreateManagedContact(ctx).ManagedContact(managedContact).Execute()

Managed Contact Create



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
    managedContact := *openapiclient.NewManagedContact() // ManagedContact | Information about the contact to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.CreateManagedContact(context.Background()).ManagedContact(managedContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.CreateManagedContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedContact`: ManagedContact
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.CreateManagedContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedContact** | [**ManagedContact**](ManagedContact.md) | Information about the contact to create. | 

### Return type

[**ManagedContact**](ManagedContact.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagedCredential

> ManagedCredential CreateManagedCredential(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Managed Credential Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the Credential to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.CreateManagedCredential(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.CreateManagedCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedCredential`: ManagedCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.CreateManagedCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the Credential to create. | 

### Return type

[**ManagedCredential**](ManagedCredential.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagedService

> ManagedService CreateManagedService(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Managed Service Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the service to monitor. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.CreateManagedService(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.CreateManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedService`: ManagedService
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.CreateManagedService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the service to monitor. | 

### Return type

[**ManagedService**](ManagedService.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedContact

> map[string]interface{} DeleteManagedContact(ctx, contactId).Execute()

Managed Contact Delete



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
    contactId := int32(56) // int32 | The ID of the contact to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.DeleteManagedContact(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.DeleteManagedContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteManagedContact`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.DeleteManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedContactRequest struct via the builder pattern


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


## DeleteManagedCredential

> map[string]interface{} DeleteManagedCredential(ctx, credentialId).Execute()

Managed Credential Delete



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
    credentialId := int32(56) // int32 | The ID of the Credential to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.DeleteManagedCredential(context.Background(), credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.DeleteManagedCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteManagedCredential`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.DeleteManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedCredentialRequest struct via the builder pattern


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


## DeleteManagedService

> map[string]interface{} DeleteManagedService(ctx, serviceId).Execute()

Managed Service Delete



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
    serviceId := int32(56) // int32 | The ID of the Managed Service to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.DeleteManagedService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.DeleteManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteManagedService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.DeleteManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedServiceRequest struct via the builder pattern


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


## DisableManagedService

> ManagedService DisableManagedService(ctx, serviceId).Execute()

Managed Service Disable



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
    serviceId := int32(56) // int32 | The ID of the Managed Service to disable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.DisableManagedService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.DisableManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableManagedService`: ManagedService
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.DisableManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The ID of the Managed Service to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedService**](ManagedService.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableManagedService

> ManagedService EnableManagedService(ctx, serviceId).Execute()

Managed Service Enable



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
    serviceId := int32(56) // int32 | The ID of the Managed Service to enable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.EnableManagedService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.EnableManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableManagedService`: ManagedService
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.EnableManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The ID of the Managed Service to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedService**](ManagedService.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedContact

> ManagedContact GetManagedContact(ctx, contactId).Execute()

Managed Contact View



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
    contactId := int32(56) // int32 | The ID of the contact to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.GetManagedContact(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedContact`: ManagedContact
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedContact**](ManagedContact.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedContacts

> InlineResponse20033 GetManagedContacts(ctx).Page(page).PageSize(pageSize).Execute()

Managed Contacts List



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
    resp, r, err := api_client.ManagedApi.GetManagedContacts(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedContacts`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCredential

> ManagedCredential GetManagedCredential(ctx, credentialId).Execute()

Managed Credential View



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
    credentialId := int32(56) // int32 | The ID of the Credential to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.GetManagedCredential(context.Background(), credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedCredential`: ManagedCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedCredential**](ManagedCredential.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCredentials

> InlineResponse20034 GetManagedCredentials(ctx).Page(page).PageSize(pageSize).Execute()

Managed Credentials List



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
    resp, r, err := api_client.ManagedApi.GetManagedCredentials(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedCredentials`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedIssue

> ManagedIssue GetManagedIssue(ctx, issueId).Execute()

Managed Issue View



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
    issueId := int32(56) // int32 | The Issue to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.GetManagedIssue(context.Background(), issueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedIssue`: ManagedIssue
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | The Issue to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedIssue**](ManagedIssue.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedIssues

> InlineResponse20036 GetManagedIssues(ctx).Page(page).PageSize(pageSize).Execute()

Managed Issues List



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
    resp, r, err := api_client.ManagedApi.GetManagedIssues(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedIssues`: InlineResponse20036
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20036**](InlineResponse20036.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedLinodeSetting

> ManagedLinodeSettings GetManagedLinodeSetting(ctx, linodeId).Execute()

Linode's Managed Settings View



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
    linodeId := int32(56) // int32 | The Linode ID whose settings we are accessing.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.GetManagedLinodeSetting(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedLinodeSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedLinodeSetting`: ManagedLinodeSettings
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedLinodeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The Linode ID whose settings we are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedLinodeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedLinodeSettings**](ManagedLinodeSettings.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedLinodeSettings

> InlineResponse20037 GetManagedLinodeSettings(ctx).Page(page).PageSize(pageSize).Execute()

Managed Linode Settings List



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
    resp, r, err := api_client.ManagedApi.GetManagedLinodeSettings(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedLinodeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedLinodeSettings`: InlineResponse20037
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedLinodeSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedLinodeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20037**](InlineResponse20037.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedService

> ManagedService GetManagedService(ctx, serviceId).Execute()

Managed Service View



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
    serviceId := int32(56) // int32 | The ID of the Managed Service to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.GetManagedService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedService`: ManagedService
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedService**](ManagedService.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedServices

> InlineResponse20038 GetManagedServices(ctx).Execute()

Managed Services List



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
    resp, r, err := api_client.ManagedApi.GetManagedServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedServices`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedServicesRequest struct via the builder pattern


### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedStats

> InlineResponse20039 GetManagedStats(ctx).Execute()

Managed Stats List



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
    resp, r, err := api_client.ManagedApi.GetManagedStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.GetManagedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagedStats`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.GetManagedStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedStatsRequest struct via the builder pattern


### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedContact

> ManagedContact UpdateManagedContact(ctx, contactId).ManagedContact(managedContact).Execute()

Managed Contact Update



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
    contactId := int32(56) // int32 | The ID of the contact to access.
    managedContact := *openapiclient.NewManagedContact() // ManagedContact | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.UpdateManagedContact(context.Background(), contactId).ManagedContact(managedContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.UpdateManagedContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedContact`: ManagedContact
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.UpdateManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedContact** | [**ManagedContact**](ManagedContact.md) | The fields to update. | 

### Return type

[**ManagedContact**](ManagedContact.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedCredential

> ManagedCredential UpdateManagedCredential(ctx, credentialId).ManagedCredential(managedCredential).Execute()

Managed Credential Update



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
    credentialId := int32(56) // int32 | The ID of the Credential to access.
    managedCredential := *openapiclient.NewManagedCredential() // ManagedCredential | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.UpdateManagedCredential(context.Background(), credentialId).ManagedCredential(managedCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.UpdateManagedCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedCredential`: ManagedCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.UpdateManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCredential** | [**ManagedCredential**](ManagedCredential.md) | The fields to update. | 

### Return type

[**ManagedCredential**](ManagedCredential.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedCredentialUsernamePassword

> map[string]interface{} UpdateManagedCredentialUsernamePassword(ctx, credentialId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Managed Credential Username and Password Update



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
    credentialId := int32(56) // int32 | The ID of the Credential to update.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The new username and password to assign to the Managed Credential.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.UpdateManagedCredentialUsernamePassword(context.Background(), credentialId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.UpdateManagedCredentialUsernamePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedCredentialUsernamePassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.UpdateManagedCredentialUsernamePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **int32** | The ID of the Credential to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedCredentialUsernamePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The new username and password to assign to the Managed Credential.  | 

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


## UpdateManagedLinodeSetting

> ManagedLinodeSettings UpdateManagedLinodeSetting(ctx, linodeId).ManagedLinodeSettings(managedLinodeSettings).Execute()

Linode's Managed Settings Update



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
    linodeId := int32(56) // int32 | The Linode ID whose settings we are accessing.
    managedLinodeSettings := *openapiclient.NewManagedLinodeSettings() // ManagedLinodeSettings | The settings to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.UpdateManagedLinodeSetting(context.Background(), linodeId).ManagedLinodeSettings(managedLinodeSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.UpdateManagedLinodeSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedLinodeSetting`: ManagedLinodeSettings
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.UpdateManagedLinodeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The Linode ID whose settings we are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedLinodeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedLinodeSettings** | [**ManagedLinodeSettings**](ManagedLinodeSettings.md) | The settings to update. | 

### Return type

[**ManagedLinodeSettings**](ManagedLinodeSettings.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagedService

> ManagedService UpdateManagedService(ctx, serviceId).ManagedService(managedService).Execute()

Managed Service Update



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
    serviceId := int32(56) // int32 | The ID of the Managed Service to access.
    managedService := *openapiclient.NewManagedService() // ManagedService | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagedApi.UpdateManagedService(context.Background(), serviceId).ManagedService(managedService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.UpdateManagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedService`: ManagedService
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.UpdateManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedService** | [**ManagedService**](ManagedService.md) | The fields to update. | 

### Return type

[**ManagedService**](ManagedService.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewManagedSSHKey

> InlineResponse20035 ViewManagedSSHKey(ctx).Execute()

Managed SSH Key View



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
    resp, r, err := api_client.ManagedApi.ViewManagedSSHKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ViewManagedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ViewManagedSSHKey`: InlineResponse20035
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ViewManagedSSHKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiViewManagedSSHKeyRequest struct via the builder pattern


### Return type

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

