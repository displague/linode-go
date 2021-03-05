# \NetworkingApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateIP**](NetworkingApi.md#AllocateIP) | **Post** /networking/ips | IP Address Allocate
[**AssignIPs**](NetworkingApi.md#AssignIPs) | **Post** /networking/ipv4/assign | Linodes Assign IPs
[**CreateFirewallDevice**](NetworkingApi.md#CreateFirewallDevice) | **Post** /networking/firewalls/{firewallId}/devices | Firewall Device Create
[**CreateFirewalls**](NetworkingApi.md#CreateFirewalls) | **Post** /networking/firewalls | Firewall Create
[**DeleteFirewall**](NetworkingApi.md#DeleteFirewall) | **Delete** /networking/firewalls/{firewallId} | Firewall Delete
[**DeleteFirewallDevice**](NetworkingApi.md#DeleteFirewallDevice) | **Delete** /networking/firewalls/{firewallId}/devices/{deviceId} | Firewall Device Delete
[**GetFirewall**](NetworkingApi.md#GetFirewall) | **Get** /networking/firewalls/{firewallId} | Firewall View
[**GetFirewallDevice**](NetworkingApi.md#GetFirewallDevice) | **Get** /networking/firewalls/{firewallId}/devices/{deviceId} | Firewall Device View
[**GetFirewallDevices**](NetworkingApi.md#GetFirewallDevices) | **Get** /networking/firewalls/{firewallId}/devices | Firewall Devices List
[**GetFirewallRules**](NetworkingApi.md#GetFirewallRules) | **Get** /networking/firewalls/{firewallId}/rules | Firewall Rules List
[**GetFirewalls**](NetworkingApi.md#GetFirewalls) | **Get** /networking/firewalls | Firewalls List
[**GetIP**](NetworkingApi.md#GetIP) | **Get** /networking/ips/{address} | IP Address View
[**GetIPs**](NetworkingApi.md#GetIPs) | **Get** /networking/ips | IP Addresses List
[**GetIPv6Pools**](NetworkingApi.md#GetIPv6Pools) | **Get** /networking/ipv6/pools | IPv6 Pools List
[**GetIPv6Ranges**](NetworkingApi.md#GetIPv6Ranges) | **Get** /networking/ipv6/ranges | IPv6 Ranges List
[**ShareIPs**](NetworkingApi.md#ShareIPs) | **Post** /networking/ipv4/share | IP Sharing Configure
[**UpdateFirewall**](NetworkingApi.md#UpdateFirewall) | **Put** /networking/firewalls/{firewallId} | Firewall Update
[**UpdateFirewallRules**](NetworkingApi.md#UpdateFirewallRules) | **Put** /networking/firewalls/{firewallId}/rules | Firewall Rules Update
[**UpdateIP**](NetworkingApi.md#UpdateIP) | **Put** /networking/ips/{address} | IP Address RDNS Update



## AllocateIP

> IPAddress AllocateIP(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

IP Address Allocate



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the address you are creating.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.AllocateIP(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.AllocateIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllocateIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.AllocateIP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the address you are creating. | 

### Return type

[**IPAddress**](IPAddress.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignIPs

> map[string]interface{} AssignIPs(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Linodes Assign IPs



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about what IPv4 address to assign, and to which Linode. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.AssignIPs(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.AssignIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignIPs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.AssignIPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about what IPv4 address to assign, and to which Linode.  | 

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


## CreateFirewallDevice

> FirewallDevices CreateFirewallDevice(ctx, firewallId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Firewall Device Create



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.CreateFirewallDevice(context.Background(), firewallId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.CreateFirewallDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirewallDevice`: FirewallDevices
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.CreateFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**FirewallDevices**](FirewallDevices.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewalls

> Firewall CreateFirewalls(ctx).InlineObject11(inlineObject11).Execute()

Firewall Create



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
    inlineObject11 := *openapiclient.NewInlineObject11("TODO", "TODO") // InlineObject11 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.CreateFirewalls(context.Background()).InlineObject11(inlineObject11).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.CreateFirewalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirewalls`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.CreateFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject11** | [**InlineObject11**](InlineObject11.md) |  | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewall

> map[string]interface{} DeleteFirewall(ctx, firewallId).Execute()

Firewall Delete



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.DeleteFirewall(context.Background(), firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.DeleteFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewall`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.DeleteFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern


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


## DeleteFirewallDevice

> map[string]interface{} DeleteFirewallDevice(ctx, firewallId, deviceId).Execute()

Firewall Device Delete



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    deviceId := int32(56) // int32 | ID of the Firewall Device to access. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.DeleteFirewallDevice(context.Background(), firewallId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.DeleteFirewallDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.DeleteFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 
**deviceId** | **int32** | ID of the Firewall Device to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallDeviceRequest struct via the builder pattern


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


## GetFirewall

> Firewall GetFirewall(ctx, firewallId).Execute()

Firewall View



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.GetFirewall(context.Background(), firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewall`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Firewall**](Firewall.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallDevice

> FirewallDevices GetFirewallDevice(ctx, firewallId, deviceId).Execute()

Firewall Device View



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    deviceId := int32(56) // int32 | ID of the Firewall Device to access. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.GetFirewallDevice(context.Background(), firewallId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetFirewallDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallDevice`: FirewallDevices
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 
**deviceId** | **int32** | ID of the Firewall Device to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirewallDevices**](FirewallDevices.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallDevices

> InlineResponse20043 GetFirewallDevices(ctx, firewallId).Page(page).PageSize(pageSize).Execute()

Firewall Devices List



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.GetFirewallDevices(context.Background(), firewallId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetFirewallDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallDevices`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetFirewallDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20043**](InlineResponse20043.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRules

> Rules GetFirewallRules(ctx, firewallId).Execute()

Firewall Rules List



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.GetFirewallRules(context.Background(), firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRules`: Rules
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rules**](Rules.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewalls

> InlineResponse20019 GetFirewalls(ctx).Page(page).PageSize(pageSize).Execute()

Firewalls List



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
    resp, r, err := api_client.NetworkingApi.GetFirewalls(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetFirewalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewalls`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIP

> IPAddress GetIP(ctx, address).Execute()

IP Address View



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
    address := "address_example" // string | The address to operate on.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.GetIP(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPAddress**](IPAddress.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPs

> InlineResponse20040 GetIPs(ctx).Execute()

IP Addresses List



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
    resp, r, err := api_client.NetworkingApi.GetIPs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPs`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetIPs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPsRequest struct via the builder pattern


### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPv6Pools

> InlineResponse20041 GetIPv6Pools(ctx).Page(page).PageSize(pageSize).Execute()

IPv6 Pools List



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
    resp, r, err := api_client.NetworkingApi.GetIPv6Pools(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetIPv6Pools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPv6Pools`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetIPv6Pools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIPv6PoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPv6Ranges

> InlineResponse20042 GetIPv6Ranges(ctx).Page(page).PageSize(pageSize).Execute()

IPv6 Ranges List



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
    resp, r, err := api_client.NetworkingApi.GetIPv6Ranges(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.GetIPv6Ranges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIPv6Ranges`: InlineResponse20042
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.GetIPv6Ranges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIPv6RangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20042**](InlineResponse20042.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareIPs

> map[string]interface{} ShareIPs(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

IP Sharing Configure



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about what IPs to share with which Linode.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.ShareIPs(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.ShareIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShareIPs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.ShareIPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about what IPs to share with which Linode. | 

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


## UpdateFirewall

> Firewall UpdateFirewall(ctx, firewallId).InlineObject12(inlineObject12).Execute()

Firewall Update



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    inlineObject12 := *openapiclient.NewInlineObject12() // InlineObject12 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.UpdateFirewall(context.Background(), firewallId).InlineObject12(inlineObject12).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.UpdateFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFirewall`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.UpdateFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject12** | [**InlineObject12**](InlineObject12.md) |  | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallRules

> Rules UpdateFirewallRules(ctx, firewallId).Body(body).Execute()

Firewall Rules Update



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
    firewallId := int32(56) // int32 | ID of the Firewall to access. 
    body := Rules(987) // Rules | The Firewall Rules information to update. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.UpdateFirewallRules(context.Background(), firewallId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.UpdateFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFirewallRules`: Rules
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.UpdateFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** | ID of the Firewall to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **Rules** | The Firewall Rules information to update. | 

### Return type

[**Rules**](Rules.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIP

> IPAddress UpdateIP(ctx, address).IPAddress(iPAddress).Execute()

IP Address RDNS Update



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
    address := "address_example" // string | The address to operate on.
    iPAddress := *openapiclient.NewIPAddress() // IPAddress | The information to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkingApi.UpdateIP(context.Background(), address).IPAddress(iPAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingApi.UpdateIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `NetworkingApi.UpdateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iPAddress** | [**IPAddress**](IPAddress.md) | The information to update. | 

### Return type

[**IPAddress**](IPAddress.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

