# \ProfileApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSSHKey**](ProfileApi.md#AddSSHKey) | **Post** /profile/sshkeys | SSH Key Add
[**CreatePersonalAccessToken**](ProfileApi.md#CreatePersonalAccessToken) | **Post** /profile/tokens | Personal Access Token Create
[**DeletePersonalAccessToken**](ProfileApi.md#DeletePersonalAccessToken) | **Delete** /profile/tokens/{tokenId} | Personal Access Token Revoke
[**DeleteProfileApp**](ProfileApi.md#DeleteProfileApp) | **Delete** /profile/apps/{appId} | App Access Revoke
[**DeleteSSHKey**](ProfileApi.md#DeleteSSHKey) | **Delete** /profile/sshkeys/{sshKeyId} | SSH Key Delete
[**GetDevices**](ProfileApi.md#GetDevices) | **Get** /profile/devices | Trusted Devices List
[**GetPersonalAccessToken**](ProfileApi.md#GetPersonalAccessToken) | **Get** /profile/tokens/{tokenId} | Personal Access Token View
[**GetPersonalAccessTokens**](ProfileApi.md#GetPersonalAccessTokens) | **Get** /profile/tokens | Personal Access Tokens List
[**GetProfile**](ProfileApi.md#GetProfile) | **Get** /profile | Profile View
[**GetProfileApp**](ProfileApi.md#GetProfileApp) | **Get** /profile/apps/{appId} | Authorized App View
[**GetProfileApps**](ProfileApi.md#GetProfileApps) | **Get** /profile/apps | Authorized Apps List
[**GetProfileGrants**](ProfileApi.md#GetProfileGrants) | **Get** /profile/grants | Grants List
[**GetProfileLogin**](ProfileApi.md#GetProfileLogin) | **Get** /profile/logins/{loginId} | Login View
[**GetProfileLogins**](ProfileApi.md#GetProfileLogins) | **Get** /profile/logins | Logins List
[**GetSSHKey**](ProfileApi.md#GetSSHKey) | **Get** /profile/sshkeys/{sshKeyId} | SSH Key View
[**GetSSHKeys**](ProfileApi.md#GetSSHKeys) | **Get** /profile/sshkeys | SSH Keys List
[**GetTrustedDevice**](ProfileApi.md#GetTrustedDevice) | **Get** /profile/devices/{deviceId} | Trusted Device View
[**GetUserPreferences**](ProfileApi.md#GetUserPreferences) | **Get** /profile/preferences | User Preferences View
[**RevokeTrustedDevice**](ProfileApi.md#RevokeTrustedDevice) | **Delete** /profile/devices/{deviceId} | Trusted Device Revoke
[**TfaConfirm**](ProfileApi.md#TfaConfirm) | **Post** /profile/tfa-enable-confirm | Two Factor Authentication Confirm/Enable
[**TfaDisable**](ProfileApi.md#TfaDisable) | **Post** /profile/tfa-disable | Two Factor Authentication Disable
[**TfaEnable**](ProfileApi.md#TfaEnable) | **Post** /profile/tfa-enable | Two Factor Secret Create
[**UpdatePersonalAccessToken**](ProfileApi.md#UpdatePersonalAccessToken) | **Put** /profile/tokens/{tokenId} | Personal Access Token Update
[**UpdateProfile**](ProfileApi.md#UpdateProfile) | **Put** /profile | Profile Update
[**UpdateSSHKey**](ProfileApi.md#UpdateSSHKey) | **Put** /profile/sshkeys/{sshKeyId} | SSH Key Update
[**UpdateUserPreferences**](ProfileApi.md#UpdateUserPreferences) | **Put** /profile/preferences | User Preferences Update



## AddSSHKey

> SSHKey AddSSHKey(ctx).SSHKeyRequest(sSHKeyRequest).Execute()

SSH Key Add



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
    sSHKeyRequest := *openapiclient.NewSSHKeyRequest() // SSHKeyRequest | Add SSH Key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.AddSSHKey(context.Background()).SSHKeyRequest(sSHKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.AddSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.AddSSHKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHKeyRequest** | [**SSHKeyRequest**](SSHKeyRequest.md) | Add SSH Key | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePersonalAccessToken

> PersonalAccessToken CreatePersonalAccessToken(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Personal Access Token Create



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the requested token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.CreatePersonalAccessToken(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.CreatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersonalAccessToken`: PersonalAccessToken
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.CreatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Information about the requested token. | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonalAccessToken

> map[string]interface{} DeletePersonalAccessToken(ctx, tokenId).Execute()

Personal Access Token Revoke



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
    tokenId := int32(56) // int32 | The ID of the token to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.DeletePersonalAccessToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.DeletePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonalAccessToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.DeletePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | The ID of the token to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalAccessTokenRequest struct via the builder pattern


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


## DeleteProfileApp

> map[string]interface{} DeleteProfileApp(ctx, appId).Execute()

App Access Revoke



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
    appId := int32(56) // int32 | The authorized app ID to manage.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.DeleteProfileApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.DeleteProfileApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProfileApp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.DeleteProfileApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The authorized app ID to manage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileAppRequest struct via the builder pattern


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


## DeleteSSHKey

> map[string]interface{} DeleteSSHKey(ctx, sshKeyId).Execute()

SSH Key Delete



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
    sshKeyId := int32(56) // int32 | The ID of the SSHKey

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.DeleteSSHKey(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.DeleteSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSSHKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.DeleteSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | The ID of the SSHKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHKeyRequest struct via the builder pattern


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


## GetDevices

> InlineResponse20053 GetDevices(ctx).Execute()

Trusted Devices List



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
    resp, r, err := api_client.ProfileApi.GetDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


### Return type

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalAccessToken

> PersonalAccessToken GetPersonalAccessToken(ctx, tokenId).Execute()

Personal Access Token View



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
    tokenId := int32(56) // int32 | The ID of the token to access.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.GetPersonalAccessToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetPersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalAccessToken`: PersonalAccessToken
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | The ID of the token to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalAccessTokens

> InlineResponse20052 GetPersonalAccessTokens(ctx).Execute()

Personal Access Tokens List



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
    resp, r, err := api_client.ProfileApi.GetPersonalAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetPersonalAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalAccessTokens`: InlineResponse20052
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetPersonalAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalAccessTokensRequest struct via the builder pattern


### Return type

[**InlineResponse20052**](InlineResponse20052.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> Profile GetProfile(ctx).Execute()

Profile View



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
    resp, r, err := api_client.ProfileApi.GetProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: Profile
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**Profile**](Profile.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileApp

> AuthorizedApp GetProfileApp(ctx, appId).Execute()

Authorized App View



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
    appId := int32(56) // int32 | The authorized app ID to manage.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.GetProfileApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileApp`: AuthorizedApp
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The authorized app ID to manage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizedApp**](AuthorizedApp.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileApps

> InlineResponse20051 GetProfileApps(ctx).Page(page).PageSize(pageSize).Execute()

Authorized Apps List



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
    resp, r, err := api_client.ProfileApi.GetProfileApps(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileApps`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileGrants

> GrantsResponse GetProfileGrants(ctx).Execute()

Grants List



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
    resp, r, err := api_client.ProfileApi.GetProfileGrants(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileGrants`: GrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileGrants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileGrantsRequest struct via the builder pattern


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


## GetProfileLogin

> Login GetProfileLogin(ctx, loginId).Execute()

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
    resp, r, err := api_client.ProfileApi.GetProfileLogin(context.Background(), loginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileLogin`: Login
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginId** | **int32** | The ID of the login object to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileLoginRequest struct via the builder pattern


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


## GetProfileLogins

> InlineResponse2005 GetProfileLogins(ctx).Execute()

Logins List



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
    resp, r, err := api_client.ProfileApi.GetProfileLogins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetProfileLogins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileLogins`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetProfileLogins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileLoginsRequest struct via the builder pattern


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


## GetSSHKey

> SSHKey GetSSHKey(ctx, sshKeyId).Execute()

SSH Key View



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
    sshKeyId := int32(56) // int32 | The ID of the SSHKey

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.GetSSHKey(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | The ID of the SSHKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHKeys

> InlineResponse20054 GetSSHKeys(ctx).Page(page).PageSize(pageSize).Execute()

SSH Keys List



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
    resp, r, err := api_client.ProfileApi.GetSSHKeys(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKeys`: InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetSSHKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20054**](InlineResponse20054.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustedDevice

> TrustedDevice GetTrustedDevice(ctx, deviceId).Execute()

Trusted Device View



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
    deviceId := int32(56) // int32 | The ID of the TrustedDevice

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.GetTrustedDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetTrustedDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedDevice`: TrustedDevice
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetTrustedDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | The ID of the TrustedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedDevice**](TrustedDevice.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPreferences

> map[string]interface{} GetUserPreferences(ctx).Execute()

User Preferences View



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
    resp, r, err := api_client.ProfileApi.GetUserPreferences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetUserPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPreferences`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetUserPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferencesRequest struct via the builder pattern


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


## RevokeTrustedDevice

> map[string]interface{} RevokeTrustedDevice(ctx, deviceId).Execute()

Trusted Device Revoke



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
    deviceId := int32(56) // int32 | The ID of the TrustedDevice

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.RevokeTrustedDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.RevokeTrustedDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeTrustedDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.RevokeTrustedDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32** | The ID of the TrustedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTrustedDeviceRequest struct via the builder pattern


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


## TfaConfirm

> map[string]interface{} TfaConfirm(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Two Factor Authentication Confirm/Enable



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The Two Factor code you generated with your Two Factor secret.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.TfaConfirm(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.TfaConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TfaConfirm`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.TfaConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The Two Factor code you generated with your Two Factor secret. | 

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


## TfaDisable

> map[string]interface{} TfaDisable(ctx).Execute()

Two Factor Authentication Disable



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
    resp, r, err := api_client.ProfileApi.TfaDisable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.TfaDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TfaDisable`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.TfaDisable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTfaDisableRequest struct via the builder pattern


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


## TfaEnable

> map[string]interface{} TfaEnable(ctx).Execute()

Two Factor Secret Create



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
    resp, r, err := api_client.ProfileApi.TfaEnable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.TfaEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TfaEnable`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.TfaEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTfaEnableRequest struct via the builder pattern


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


## UpdatePersonalAccessToken

> PersonalAccessToken UpdatePersonalAccessToken(ctx, tokenId).PersonalAccessToken(personalAccessToken).Execute()

Personal Access Token Update



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
    tokenId := int32(56) // int32 | The ID of the token to access.
    personalAccessToken := *openapiclient.NewPersonalAccessToken() // PersonalAccessToken | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.UpdatePersonalAccessToken(context.Background(), tokenId).PersonalAccessToken(personalAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonalAccessToken`: PersonalAccessToken
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | The ID of the token to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personalAccessToken** | [**PersonalAccessToken**](PersonalAccessToken.md) | The fields to update. | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> Profile UpdateProfile(ctx).Profile(profile).Execute()

Profile Update



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
    profile := *openapiclient.NewProfile() // Profile | The fields to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.UpdateProfile(context.Background()).Profile(profile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: Profile
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | [**Profile**](Profile.md) | The fields to update. | 

### Return type

[**Profile**](Profile.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHKey

> SSHKey UpdateSSHKey(ctx, sshKeyId).SSHKey(sSHKey).Execute()

SSH Key Update



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
    sshKeyId := int32(56) // int32 | The ID of the SSHKey
    sSHKey := *openapiclient.NewSSHKey() // SSHKey | The fields to update. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.UpdateSSHKey(context.Background(), sshKeyId).SSHKey(sSHKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | The ID of the SSHKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHKey** | [**SSHKey**](SSHKey.md) | The fields to update.  | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPreferences

> map[string]interface{} UpdateUserPreferences(ctx).Body(body).Execute()

User Preferences Update



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
    body := map[string]interface{}(Object) // map[string]interface{} | The user preferences to update or store. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProfileApi.UpdateUserPreferences(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateUserPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPreferences`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateUserPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The user preferences to update or store.  | 

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

