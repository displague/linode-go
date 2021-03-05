# \SupportApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseTicket**](SupportApi.md#CloseTicket) | **Post** /support/tickets/{ticketId}/close | Support Ticket Close
[**CreateTicket**](SupportApi.md#CreateTicket) | **Post** /support/tickets | Support Ticket Open
[**CreateTicketAttachment**](SupportApi.md#CreateTicketAttachment) | **Post** /support/tickets/{ticketId}/attachments | Ticket Attachment Create
[**CreateTicketReply**](SupportApi.md#CreateTicketReply) | **Post** /support/tickets/{ticketId}/replies | Reply Create
[**GetTicket**](SupportApi.md#GetTicket) | **Get** /support/tickets/{ticketId} | Support Ticket View
[**GetTicketReplies**](SupportApi.md#GetTicketReplies) | **Get** /support/tickets/{ticketId}/replies | Replies List
[**GetTickets**](SupportApi.md#GetTickets) | **Get** /support/tickets | Support Tickets List



## CloseTicket

> map[string]interface{} CloseTicket(ctx, ticketId).Execute()

Support Ticket Close



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
    ticketId := int32(56) // int32 | The ID of the Support Ticket.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.CloseTicket(context.Background(), ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.CloseTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseTicket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.CloseTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | The ID of the Support Ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseTicketRequest struct via the builder pattern


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


## CreateTicket

> SupportTicket CreateTicket(ctx).SupportTicketRequest(supportTicketRequest).Execute()

Support Ticket Open



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
    supportTicketRequest := *openapiclient.NewSupportTicketRequest("I'm having trouble setting the root password on my Linode. I tried following the instructions but something is not working and I'm not sure what I'm doing wrong. Can you please help me figure out how I can reset it?
", "Having trouble resetting root password on my Linode
") // SupportTicketRequest | Open a Support Ticket. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.CreateTicket(context.Background()).SupportTicketRequest(supportTicketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.CreateTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTicket`: SupportTicket
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.CreateTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportTicketRequest** | [**SupportTicketRequest**](SupportTicketRequest.md) | Open a Support Ticket. | 

### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTicketAttachment

> map[string]interface{} CreateTicketAttachment(ctx, ticketId).File(file).Execute()

Ticket Attachment Create



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
    ticketId := int32(56) // int32 | The ID of the Support Ticket.
    file := "file_example" // string | The local, absolute path to the file you want to attach to your Support Ticket. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.CreateTicketAttachment(context.Background(), ticketId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.CreateTicketAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTicketAttachment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.CreateTicketAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | The ID of the Support Ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **string** | The local, absolute path to the file you want to attach to your Support Ticket.  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTicketReply

> SupportTicketReply CreateTicketReply(ctx, ticketId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Reply Create



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
    ticketId := int32(56) // int32 | The ID of the Support Ticket.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Add a reply.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.CreateTicketReply(context.Background(), ticketId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.CreateTicketReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTicketReply`: SupportTicketReply
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.CreateTicketReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | The ID of the Support Ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Add a reply. | 

### Return type

[**SupportTicketReply**](SupportTicketReply.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicket

> SupportTicket GetTicket(ctx, ticketId).Execute()

Support Ticket View



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
    ticketId := int32(56) // int32 | The ID of the Support Ticket.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.GetTicket(context.Background(), ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.GetTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicket`: SupportTicket
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.GetTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | The ID of the Support Ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupportTicket**](SupportTicket.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketReplies

> InlineResponse20057 GetTicketReplies(ctx, ticketId).Page(page).PageSize(pageSize).Execute()

Replies List



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
    ticketId := int32(56) // int32 | The ID of the Support Ticket.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportApi.GetTicketReplies(context.Background(), ticketId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.GetTicketReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicketReplies`: InlineResponse20057
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.GetTicketReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32** | The ID of the Support Ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20057**](InlineResponse20057.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickets

> InlineResponse20056 GetTickets(ctx).Page(page).PageSize(pageSize).Execute()

Support Tickets List



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
    resp, r, err := api_client.SupportApi.GetTickets(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.GetTickets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTickets`: InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.GetTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

