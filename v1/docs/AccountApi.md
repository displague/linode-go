# \AccountApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptEntityTransfer**](AccountApi.md#AcceptEntityTransfer) | **Post** /account/entity-transfers/{token}/accept | Entity Transfer Accept
[**CancelAccount**](AccountApi.md#CancelAccount) | **Post** /account/cancel | Account Cancel
[**CreateClient**](AccountApi.md#CreateClient) | **Post** /account/oauth-clients | OAuth Client Create
[**CreateCreditCard**](AccountApi.md#CreateCreditCard) | **Post** /account/credit-card | Credit Card Add/Edit
[**CreateEntityTransfer**](AccountApi.md#CreateEntityTransfer) | **Post** /account/entity-transfers | Entity Transfer Create
[**CreatePayPalPayment**](AccountApi.md#CreatePayPalPayment) | **Post** /account/payments/paypal | PayPal Payment Stage
[**CreatePayment**](AccountApi.md#CreatePayment) | **Post** /account/payments | Payment Make
[**CreateUser**](AccountApi.md#CreateUser) | **Post** /account/users | User Create
[**DeleteClient**](AccountApi.md#DeleteClient) | **Delete** /account/oauth-clients/{clientId} | OAuth Client Delete
[**DeleteEntityTransfer**](AccountApi.md#DeleteEntityTransfer) | **Delete** /account/entity-transfers/{token} | Entity Transfer Cancel
[**DeleteUser**](AccountApi.md#DeleteUser) | **Delete** /account/users/{username} | User Delete
[**EnableAccountManged**](AccountApi.md#EnableAccountManged) | **Post** /account/settings/managed-enable | Linode Managed Enable
[**EventRead**](AccountApi.md#EventRead) | **Post** /account/events/{eventId}/read | Event Mark as Read
[**EventSeen**](AccountApi.md#EventSeen) | **Post** /account/events/{eventId}/seen | Event Mark as Seen
[**ExecutePayPalPayment**](AccountApi.md#ExecutePayPalPayment) | **Post** /account/payments/paypal/execute | Staged/Approved PayPal Payment Execute
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /account | Account View
[**GetAccountLogin**](AccountApi.md#GetAccountLogin) | **Get** /account/logins/{loginId} | Login View
[**GetAccountLogins**](AccountApi.md#GetAccountLogins) | **Get** /account/logins | User Logins List All
[**GetAccountSettings**](AccountApi.md#GetAccountSettings) | **Get** /account/settings | Account Settings View
[**GetClient**](AccountApi.md#GetClient) | **Get** /account/oauth-clients/{clientId} | OAuth Client View
[**GetClientThumbnail**](AccountApi.md#GetClientThumbnail) | **Get** /account/oauth-clients/{clientId}/thumbnail | OAuth Client Thumbnail View
[**GetClients**](AccountApi.md#GetClients) | **Get** /account/oauth-clients | OAuth Clients List
[**GetEntityTransfer**](AccountApi.md#GetEntityTransfer) | **Get** /account/entity-transfers/{token} | Entity Transfer View
[**GetEntityTransfers**](AccountApi.md#GetEntityTransfers) | **Get** /account/entity-transfers | Entity Transfers List
[**GetEvent**](AccountApi.md#GetEvent) | **Get** /account/events/{eventId} | Event View
[**GetEvents**](AccountApi.md#GetEvents) | **Get** /account/events | Events List
[**GetInvoice**](AccountApi.md#GetInvoice) | **Get** /account/invoices/{invoiceId} | Invoice View
[**GetInvoiceItems**](AccountApi.md#GetInvoiceItems) | **Get** /account/invoices/{invoiceId}/items | Invoice Items List
[**GetInvoices**](AccountApi.md#GetInvoices) | **Get** /account/invoices | Invoices List
[**GetMaintenance**](AccountApi.md#GetMaintenance) | **Get** /account/maintenance | Maintenance List
[**GetNotifications**](AccountApi.md#GetNotifications) | **Get** /account/notifications | Notifications List
[**GetPayment**](AccountApi.md#GetPayment) | **Get** /account/payments/{paymentId} | Payment View
[**GetPayments**](AccountApi.md#GetPayments) | **Get** /account/payments | Payments List
[**GetTransfer**](AccountApi.md#GetTransfer) | **Get** /account/transfer | Network Utilization View
[**GetUser**](AccountApi.md#GetUser) | **Get** /account/users/{username} | User View
[**GetUserGrants**](AccountApi.md#GetUserGrants) | **Get** /account/users/{username}/grants | User&#39;s Grants View
[**GetUsers**](AccountApi.md#GetUsers) | **Get** /account/users | Users List
[**ResetClientSecret**](AccountApi.md#ResetClientSecret) | **Post** /account/oauth-clients/{clientId}/reset-secret | OAuth Client Secret Reset
[**SetClientThumbnail**](AccountApi.md#SetClientThumbnail) | **Put** /account/oauth-clients/{clientId}/thumbnail | OAuth Client Thumbnail Update
[**UpdateAccount**](AccountApi.md#UpdateAccount) | **Put** /account | Account Update
[**UpdateAccountSettings**](AccountApi.md#UpdateAccountSettings) | **Put** /account/settings | Account Settings Update
[**UpdateClient**](AccountApi.md#UpdateClient) | **Put** /account/oauth-clients/{clientId} | OAuth Client Update
[**UpdateUser**](AccountApi.md#UpdateUser) | **Put** /account/users/{username} | User Update
[**UpdateUserGrants**](AccountApi.md#UpdateUserGrants) | **Put** /account/users/{username}/grants | User&#39;s Grants Update



## AcceptEntityTransfer

> map[string]interface{} AcceptEntityTransfer(ctx, token).Execute()

Entity Transfer Accept



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
    token := TODO // string | The UUID of the Entity Transfer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AcceptEntityTransfer(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AcceptEntityTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptEntityTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AcceptEntityTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | [**string**](.md) | The UUID of the Entity Transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptEntityTransferRequest struct via the builder pattern


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


## CancelAccount

> InlineResponse200 CancelAccount(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Account Cancel



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Supply a comment stating the reason that you are cancelling your account. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CancelAccount(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CancelAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelAccount`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CancelAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Supply a comment stating the reason that you are cancelling your account.  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClient

> OAuthClient CreateClient(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

OAuth Client Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the OAuth Client to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreateClient(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClient`: OAuthClient
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the OAuth Client to create. | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditCard

> map[string]interface{} CreateCreditCard(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Credit Card Add/Edit



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Update the credit card information associated with your Account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreateCreditCard(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateCreditCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreditCard`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateCreditCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Update the credit card information associated with your Account. | 

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


## CreateEntityTransfer

> EntityTransfer CreateEntityTransfer(ctx).InlineObject(inlineObject).Execute()

Entity Transfer Create



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
    inlineObject := *openapiclient.NewInlineObject("TODO") // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreateEntityTransfer(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateEntityTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntityTransfer`: EntityTransfer
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateEntityTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**EntityTransfer**](EntityTransfer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayPalPayment

> InlineResponse20010 CreatePayPalPayment(ctx).PayPal(payPal).Execute()

PayPal Payment Stage



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
    payPal := *openapiclient.NewPayPal("https://example.org", "https://example.org", "120.50") // PayPal | The amount of the Payment to submit via PayPal. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreatePayPalPayment(context.Background()).PayPal(payPal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreatePayPalPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayPalPayment`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreatePayPalPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayPalPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payPal** | [**PayPal**](PayPal.md) | The amount of the Payment to submit via PayPal.  | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayment

> Payment CreatePayment(ctx).PaymentRequest(paymentRequest).Execute()

Payment Make



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
    paymentRequest := *openapiclient.NewPaymentRequest("120.50") // PaymentRequest | Information about the Payment you are making.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreatePayment(context.Background()).PaymentRequest(paymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayment`: Payment
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequest** | [**PaymentRequest**](PaymentRequest.md) | Information about the Payment you are making. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

User Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the User to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.CreateUser(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the User to create. | 

### Return type

[**User**](User.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClient

> map[string]interface{} DeleteClient(ctx, clientId).Execute()

OAuth Client Delete



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.DeleteClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClient`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


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


## DeleteEntityTransfer

> map[string]interface{} DeleteEntityTransfer(ctx, token).Execute()

Entity Transfer Cancel



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
    token := TODO // string | The UUID of the Entity Transfer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.DeleteEntityTransfer(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteEntityTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntityTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteEntityTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | [**string**](.md) | The UUID of the Entity Transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityTransferRequest struct via the builder pattern


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


## DeleteUser

> map[string]interface{} DeleteUser(ctx, username).Execute()

User Delete



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
    username := "username_example" // string | The username to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.DeleteUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## EnableAccountManged

> map[string]interface{} EnableAccountManged(ctx).Execute()

Linode Managed Enable



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
    resp, r, err := api_client.AccountApi.EnableAccountManged(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.EnableAccountManged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAccountManged`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.EnableAccountManged`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAccountMangedRequest struct via the builder pattern


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


## EventRead

> map[string]interface{} EventRead(ctx, eventId).Execute()

Event Mark as Read



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
    eventId := int32(56) // int32 | The ID of the Event to designate as read.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.EventRead(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.EventRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventRead`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.EventRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int32** | The ID of the Event to designate as read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventReadRequest struct via the builder pattern


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


## EventSeen

> map[string]interface{} EventSeen(ctx, eventId).Execute()

Event Mark as Seen



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
    eventId := int32(56) // int32 | The ID of the Event to designate as seen.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.EventSeen(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.EventSeen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventSeen`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.EventSeen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int32** | The ID of the Event to designate as seen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventSeenRequest struct via the builder pattern


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


## ExecutePayPalPayment

> map[string]interface{} ExecutePayPalPayment(ctx).PayPalExecute(payPalExecute).Execute()

Staged/Approved PayPal Payment Execute



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
    payPalExecute := *openapiclient.NewPayPalExecute("ABCDEFGHIJKLM", "PAY-1234567890ABCDEFGHIJKLMN") // PayPalExecute | The details of the Payment to execute. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ExecutePayPalPayment(context.Background()).PayPalExecute(payPalExecute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ExecutePayPalPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutePayPalPayment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ExecutePayPalPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutePayPalPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payPalExecute** | [**PayPalExecute**](PayPalExecute.md) | The details of the Payment to execute.  | 

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


## GetAccount

> Account GetAccount(ctx).Execute()

Account View



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
    resp, r, err := api_client.AccountApi.GetAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLogin

> Login GetAccountLogin(ctx, loginId).Execute()

Login View



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
    loginId := int32(56) // int32 | The ID of the login object to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetAccountLogin(context.Background(), loginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountLogin`: Login
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginId** | **int32** | The ID of the login object to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Login**](Login.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLogins

> InlineResponse2005 GetAccountLogins(ctx).Execute()

User Logins List All



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
    resp, r, err := api_client.AccountApi.GetAccountLogins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountLogins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountLogins`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountLogins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLoginsRequest struct via the builder pattern


### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSettings

> AccountSettings GetAccountSettings(ctx).Execute()

Account Settings View



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
    resp, r, err := api_client.AccountApi.GetAccountSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountSettings`: AccountSettings
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSettingsRequest struct via the builder pattern


### Return type

[**AccountSettings**](AccountSettings.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClient

> OAuthClient GetClient(ctx, clientId).Execute()

OAuth Client View



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClient`: OAuthClient
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientThumbnail

> *os.File GetClientThumbnail(ctx, clientId).Execute()

OAuth Client Thumbnail View



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetClientThumbnail(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetClientThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientThumbnail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetClientThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> InlineResponse2008 GetClients(ctx).Page(page).PageSize(pageSize).Execute()

OAuth Clients List



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
    resp, r, err := api_client.AccountApi.GetClients(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClients`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityTransfer

> EntityTransfer GetEntityTransfer(ctx, token).Execute()

Entity Transfer View



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
    token := TODO // string | The UUID of the Entity Transfer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetEntityTransfer(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetEntityTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityTransfer`: EntityTransfer
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetEntityTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | [**string**](.md) | The UUID of the Entity Transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityTransfer**](EntityTransfer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityTransfers

> InlineResponse2001 GetEntityTransfers(ctx).Page(page).PageSize(pageSize).Execute()

Entity Transfers List



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
    resp, r, err := api_client.AccountApi.GetEntityTransfers(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetEntityTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityTransfers`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetEntityTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> Event GetEvent(ctx, eventId).Execute()

Event View



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
    eventId := int32(56) // int32 | The ID of the Event.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetEvent(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: Event
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int32** | The ID of the Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> InlineResponse2002 GetEvents(ctx).Page(page).PageSize(pageSize).Execute()

Events List



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
    resp, r, err := api_client.AccountApi.GetEvents(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Invoice View



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
    invoiceId := int32(56) // int32 | The ID of the Invoice.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | The ID of the Invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceItems

> InlineResponse2004 GetInvoiceItems(ctx, invoiceId).Page(page).PageSize(pageSize).Execute()

Invoice Items List



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
    invoiceId := int32(56) // int32 | The ID of the Invoice.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetInvoiceItems(context.Background(), invoiceId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetInvoiceItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceItems`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetInvoiceItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | The ID of the Invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> InlineResponse2003 GetInvoices(ctx).Page(page).PageSize(pageSize).Execute()

Invoices List



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
    resp, r, err := api_client.AccountApi.GetInvoices(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenance

> InlineResponse2006 GetMaintenance(ctx).Execute()

Maintenance List



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
    resp, r, err := api_client.AccountApi.GetMaintenance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenance`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetMaintenance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceRequest struct via the builder pattern


### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> InlineResponse2007 GetNotifications(ctx).Execute()

Notifications List



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
    resp, r, err := api_client.AccountApi.GetNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayment

> Payment GetPayment(ctx, paymentId).Execute()

Payment View



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
    paymentId := int32(56) // int32 | The ID of the Payment to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetPayment(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayment`: Payment
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **int32** | The ID of the Payment to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Payment**](Payment.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> InlineResponse2009 GetPayments(ctx).Page(page).PageSize(pageSize).Execute()

Payments List



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
    resp, r, err := api_client.AccountApi.GetPayments(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayments`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfer

> Transfer GetTransfer(ctx).Execute()

Network Utilization View



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
    resp, r, err := api_client.AccountApi.GetTransfer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfer`: Transfer
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetTransfer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRequest struct via the builder pattern


### Return type

[**Transfer**](Transfer.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, username).Execute()

User View



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
    username := "username_example" // string | The username to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGrants

> GrantsResponse GetUserGrants(ctx, username).Execute()

User's Grants View



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
    username := "username_example" // string | The username to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.GetUserGrants(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetUserGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserGrants`: GrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetUserGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GrantsResponse**](GrantsResponse.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse20011 GetUsers(ctx).Page(page).PageSize(pageSize).Execute()

Users List



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
    resp, r, err := api_client.AccountApi.GetUsers(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetClientSecret

> OAuthClient ResetClientSecret(ctx, clientId).Execute()

OAuth Client Secret Reset



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.ResetClientSecret(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ResetClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetClientSecret`: OAuthClient
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ResetClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClientThumbnail

> map[string]interface{} SetClientThumbnail(ctx, clientId).Body(body).Execute()

OAuth Client Thumbnail Update



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.
    body := os.NewFile(1234, "some_file") // *os.File | The image to set as the thumbnail.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.SetClientThumbnail(context.Background(), clientId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.SetClientThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetClientThumbnail`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.SetClientThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClientThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | The image to set as the thumbnail. | 

### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: image/png
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Account UpdateAccount(ctx).Account(account).Execute()

Account Update



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
    account := *openapiclient.NewAccount() // Account | Update contact and billing information.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateAccount(context.Background()).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**Account**](Account.md) | Update contact and billing information. | 

### Return type

[**Account**](Account.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSettings

> AccountSettings UpdateAccountSettings(ctx).AccountSettings(accountSettings).Execute()

Account Settings Update



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
    accountSettings := *openapiclient.NewAccountSettings() // AccountSettings | Update Account settings information.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateAccountSettings(context.Background()).AccountSettings(accountSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateAccountSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountSettings`: AccountSettings
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateAccountSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountSettings** | [**AccountSettings**](AccountSettings.md) | Update Account settings information. | 

### Return type

[**AccountSettings**](AccountSettings.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClient

> OAuthClient UpdateClient(ctx, clientId).OAuthClient(oAuthClient).Execute()

OAuth Client Update



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
    clientId := "clientId_example" // string | The OAuth Client ID to look up.
    oAuthClient := *openapiclient.NewOAuthClient() // OAuthClient | The fields to update. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateClient(context.Background(), clientId).OAuthClient(oAuthClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClient`: OAuthClient
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth Client ID to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthClient** | [**OAuthClient**](OAuthClient.md) | The fields to update. | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, username).User(user).Execute()

User Update



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
    username := "username_example" // string | The username to look up.
    user := *openapiclient.NewUser() // User | The information to update. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateUser(context.Background(), username).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | The information to update. | 

### Return type

[**User**](User.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGrants

> GrantsResponse UpdateUserGrants(ctx, username).GrantsResponse(grantsResponse).Execute()

User's Grants Update



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
    username := "username_example" // string | The username to look up.
    grantsResponse := *openapiclient.NewGrantsResponse() // GrantsResponse | The grants to update. Omitted grants will be left unchanged.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.UpdateUserGrants(context.Background(), username).GrantsResponse(grantsResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.UpdateUserGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserGrants`: GrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.UpdateUserGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantsResponse** | [**GrantsResponse**](GrantsResponse.md) | The grants to update. Omitted grants will be left unchanged. | 

### Return type

[**GrantsResponse**](GrantsResponse.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

