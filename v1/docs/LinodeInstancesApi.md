# \LinodeInstancesApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLinodeConfig**](LinodeInstancesApi.md#AddLinodeConfig) | **Post** /linode/instances/{linodeId}/configs | Configuration Profile Create
[**AddLinodeDisk**](LinodeInstancesApi.md#AddLinodeDisk) | **Post** /linode/instances/{linodeId}/disks | Disk Create
[**AddLinodeIP**](LinodeInstancesApi.md#AddLinodeIP) | **Post** /linode/instances/{linodeId}/ips | IPv4 Address Allocate
[**BootLinodeInstance**](LinodeInstancesApi.md#BootLinodeInstance) | **Post** /linode/instances/{linodeId}/boot | Linode Boot
[**CancelBackups**](LinodeInstancesApi.md#CancelBackups) | **Post** /linode/instances/{linodeId}/backups/cancel | Backups Cancel
[**CloneLinodeDisk**](LinodeInstancesApi.md#CloneLinodeDisk) | **Post** /linode/instances/{linodeId}/disks/{diskId}/clone | Disk Clone
[**CloneLinodeInstance**](LinodeInstancesApi.md#CloneLinodeInstance) | **Post** /linode/instances/{linodeId}/clone | Linode Clone
[**CreateLinodeInstance**](LinodeInstancesApi.md#CreateLinodeInstance) | **Post** /linode/instances | Linode Create
[**CreateSnapshot**](LinodeInstancesApi.md#CreateSnapshot) | **Post** /linode/instances/{linodeId}/backups | Snapshot Create
[**DeleteDisk**](LinodeInstancesApi.md#DeleteDisk) | **Delete** /linode/instances/{linodeId}/disks/{diskId} | Disk Delete
[**DeleteLinodeConfig**](LinodeInstancesApi.md#DeleteLinodeConfig) | **Delete** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile Delete
[**DeleteLinodeInstance**](LinodeInstancesApi.md#DeleteLinodeInstance) | **Delete** /linode/instances/{linodeId} | Linode Delete
[**EnableBackups**](LinodeInstancesApi.md#EnableBackups) | **Post** /linode/instances/{linodeId}/backups/enable | Backups Enable
[**GetBackup**](LinodeInstancesApi.md#GetBackup) | **Get** /linode/instances/{linodeId}/backups/{backupId} | Backup View
[**GetBackups**](LinodeInstancesApi.md#GetBackups) | **Get** /linode/instances/{linodeId}/backups | Backups List
[**GetKernel**](LinodeInstancesApi.md#GetKernel) | **Get** /linode/kernels/{kernelId} | Kernel View
[**GetKernels**](LinodeInstancesApi.md#GetKernels) | **Get** /linode/kernels | Kernels List
[**GetLinodeConfig**](LinodeInstancesApi.md#GetLinodeConfig) | **Get** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile View
[**GetLinodeConfigs**](LinodeInstancesApi.md#GetLinodeConfigs) | **Get** /linode/instances/{linodeId}/configs | Configuration Profiles List
[**GetLinodeDisk**](LinodeInstancesApi.md#GetLinodeDisk) | **Get** /linode/instances/{linodeId}/disks/{diskId} | Disk View
[**GetLinodeDisks**](LinodeInstancesApi.md#GetLinodeDisks) | **Get** /linode/instances/{linodeId}/disks | Disks List
[**GetLinodeFirewalls**](LinodeInstancesApi.md#GetLinodeFirewalls) | **Get** /linode/instances/{linodeId}/firewalls | Firewalls List
[**GetLinodeIP**](LinodeInstancesApi.md#GetLinodeIP) | **Get** /linode/instances/{linodeId}/ips/{address} | IP Address View
[**GetLinodeIPs**](LinodeInstancesApi.md#GetLinodeIPs) | **Get** /linode/instances/{linodeId}/ips | Networking Information List
[**GetLinodeInstance**](LinodeInstancesApi.md#GetLinodeInstance) | **Get** /linode/instances/{linodeId} | Linode View
[**GetLinodeInstances**](LinodeInstancesApi.md#GetLinodeInstances) | **Get** /linode/instances | Linodes List
[**GetLinodeStats**](LinodeInstancesApi.md#GetLinodeStats) | **Get** /linode/instances/{linodeId}/stats | Linode Statistics View
[**GetLinodeStatsByYearMonth**](LinodeInstancesApi.md#GetLinodeStatsByYearMonth) | **Get** /linode/instances/{linodeId}/stats/{year}/{month} | Statistics View (year/month)
[**GetLinodeTransfer**](LinodeInstancesApi.md#GetLinodeTransfer) | **Get** /linode/instances/{linodeId}/transfer | Network Transfer View
[**GetLinodeTransferByYearMonth**](LinodeInstancesApi.md#GetLinodeTransferByYearMonth) | **Get** /linode/instances/{linodeId}/transfer/{year}/{month} | Network Transfer View (year/month)
[**GetLinodeVolumes**](LinodeInstancesApi.md#GetLinodeVolumes) | **Get** /linode/instances/{linodeId}/volumes | Linode&#39;s Volumes List
[**MigrateLinodeInstance**](LinodeInstancesApi.md#MigrateLinodeInstance) | **Post** /linode/instances/{linodeId}/migrate | DC Migration/Pending Host Migration Initiate
[**MutateLinodeInstance**](LinodeInstancesApi.md#MutateLinodeInstance) | **Post** /linode/instances/{linodeId}/mutate | Linode Upgrade
[**RebootLinodeInstance**](LinodeInstancesApi.md#RebootLinodeInstance) | **Post** /linode/instances/{linodeId}/reboot | Linode Reboot
[**RebuildLinodeInstance**](LinodeInstancesApi.md#RebuildLinodeInstance) | **Post** /linode/instances/{linodeId}/rebuild | Linode Rebuild
[**RemoveLinodeIP**](LinodeInstancesApi.md#RemoveLinodeIP) | **Delete** /linode/instances/{linodeId}/ips/{address} | IPv4 Address Delete
[**RescueLinodeInstance**](LinodeInstancesApi.md#RescueLinodeInstance) | **Post** /linode/instances/{linodeId}/rescue | Linode Boot into Rescue Mode
[**ResetDiskPassword**](LinodeInstancesApi.md#ResetDiskPassword) | **Post** /linode/instances/{linodeId}/disks/{diskId}/password | Disk Root Password Reset
[**ResetLinodePassword**](LinodeInstancesApi.md#ResetLinodePassword) | **Post** /linode/instances/{linodeId}/password | Linode Root Password Reset
[**ResizeDisk**](LinodeInstancesApi.md#ResizeDisk) | **Post** /linode/instances/{linodeId}/disks/{diskId}/resize | Disk Resize
[**ResizeLinodeInstance**](LinodeInstancesApi.md#ResizeLinodeInstance) | **Post** /linode/instances/{linodeId}/resize | Linode Resize
[**RestoreBackup**](LinodeInstancesApi.md#RestoreBackup) | **Post** /linode/instances/{linodeId}/backups/{backupId}/restore | Backup Restore
[**ShutdownLinodeInstance**](LinodeInstancesApi.md#ShutdownLinodeInstance) | **Post** /linode/instances/{linodeId}/shutdown | Linode Shut Down
[**UpdateDisk**](LinodeInstancesApi.md#UpdateDisk) | **Put** /linode/instances/{linodeId}/disks/{diskId} | Disk Update
[**UpdateLinodeConfig**](LinodeInstancesApi.md#UpdateLinodeConfig) | **Put** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile Update
[**UpdateLinodeIP**](LinodeInstancesApi.md#UpdateLinodeIP) | **Put** /linode/instances/{linodeId}/ips/{address} | IP Address Update
[**UpdateLinodeInstance**](LinodeInstancesApi.md#UpdateLinodeInstance) | **Put** /linode/instances/{linodeId} | Linode Update



## AddLinodeConfig

> LinodeConfig AddLinodeConfig(ctx, linodeId).LinodeConfig(linodeConfig).Execute()

Configuration Profile Create



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
    linodeId := int32(56) // int32 | ID of the Linode to look up Configuration profiles for.
    linodeConfig := *openapiclient.NewLinodeConfig("My Config", *openapiclient.NewDevices()) // LinodeConfig | The parameters to set when creating the Configuration profile. This determines which kernel, devices, how much memory, etc. a Linode boots with. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.AddLinodeConfig(context.Background(), linodeId).LinodeConfig(linodeConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.AddLinodeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLinodeConfig`: LinodeConfig
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.AddLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up Configuration profiles for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linodeConfig** | [**LinodeConfig**](LinodeConfig.md) | The parameters to set when creating the Configuration profile. This determines which kernel, devices, how much memory, etc. a Linode boots with.  | 

### Return type

[**LinodeConfig**](LinodeConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLinodeDisk

> Disk AddLinodeDisk(ctx, linodeId).DiskRequest(diskRequest).Execute()

Disk Create



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskRequest := *openapiclient.NewDiskRequest(int32(48640), "TODO") // DiskRequest | The parameters to set when creating the Disk. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.AddLinodeDisk(context.Background(), linodeId).DiskRequest(diskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.AddLinodeDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLinodeDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.AddLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskRequest** | [**DiskRequest**](DiskRequest.md) | The parameters to set when creating the Disk.  | 

### Return type

[**Disk**](Disk.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLinodeIP

> IPAddress AddLinodeIP(ctx, linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

IPv4 Address Allocate



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Information about the address you are creating.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.AddLinodeIP(context.Background(), linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.AddLinodeIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLinodeIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.AddLinodeIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLinodeIPRequest struct via the builder pattern


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


## BootLinodeInstance

> map[string]interface{} BootLinodeInstance(ctx, linodeId).InlineObject5(inlineObject5).Execute()

Linode Boot



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
    linodeId := int32(56) // int32 | The ID of the Linode to boot.
    inlineObject5 := *openapiclient.NewInlineObject5() // InlineObject5 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.BootLinodeInstance(context.Background(), linodeId).InlineObject5(inlineObject5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.BootLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BootLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.BootLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to boot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBootLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject5** | [**InlineObject5**](InlineObject5.md) |  | 

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


## CancelBackups

> map[string]interface{} CancelBackups(ctx, linodeId).Execute()

Backups Cancel



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
    linodeId := int32(56) // int32 | The ID of the Linode to cancel backup service for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.CancelBackups(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.CancelBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelBackups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.CancelBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to cancel backup service for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBackupsRequest struct via the builder pattern


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


## CloneLinodeDisk

> Disk CloneLinodeDisk(ctx, linodeId, diskId).Execute()

Disk Clone



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to clone.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.CloneLinodeDisk(context.Background(), linodeId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.CloneLinodeDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneLinodeDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.CloneLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Disk**](Disk.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneLinodeInstance

> Linode CloneLinodeInstance(ctx, linodeId).InlineObject6(inlineObject6).Execute()

Linode Clone



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
    linodeId := int32(56) // int32 | ID of the Linode to clone.
    inlineObject6 := *openapiclient.NewInlineObject6("us-east", "g6-standard-2") // InlineObject6 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.CloneLinodeInstance(context.Background(), linodeId).InlineObject6(inlineObject6).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.CloneLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneLinodeInstance`: Linode
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.CloneLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject6** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**Linode**](Linode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinodeInstance

> Linode CreateLinodeInstance(ctx).InlineObject2(inlineObject2).Execute()

Linode Create



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
    inlineObject2 := *openapiclient.NewInlineObject2("g6-standard-2", "us-east") // InlineObject2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.CreateLinodeInstance(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.CreateLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLinodeInstance`: Linode
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.CreateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**Linode**](Linode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> Backup CreateSnapshot(ctx, linodeId).InlineObject3(inlineObject3).Execute()

Snapshot Create



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
    linodeId := int32(56) // int32 | The ID of the Linode the backups belong to.
    inlineObject3 := *openapiclient.NewInlineObject3("SnapshotLabel") // InlineObject3 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.CreateSnapshot(context.Background(), linodeId).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshot`: Backup
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode the backups belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

[**Backup**](Backup.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDisk

> map[string]interface{} DeleteDisk(ctx, linodeId, diskId).Execute()

Disk Delete



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.DeleteDisk(context.Background(), linodeId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.DeleteDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDisk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.DeleteDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiskRequest struct via the builder pattern


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


## DeleteLinodeConfig

> map[string]interface{} DeleteLinodeConfig(ctx, linodeId, configId).Execute()

Configuration Profile Delete



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
    linodeId := int32(56) // int32 | The ID of the Linode whose Configuration profile to look up.
    configId := int32(56) // int32 | The ID of the Configuration profile to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.DeleteLinodeConfig(context.Background(), linodeId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.DeleteLinodeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLinodeConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.DeleteLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode whose Configuration profile to look up. | 
**configId** | **int32** | The ID of the Configuration profile to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeConfigRequest struct via the builder pattern


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


## DeleteLinodeInstance

> map[string]interface{} DeleteLinodeInstance(ctx, linodeId).Execute()

Linode Delete



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
    linodeId := int32(56) // int32 | ID of the Linode to look up

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.DeleteLinodeInstance(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.DeleteLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.DeleteLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeInstanceRequest struct via the builder pattern


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


## EnableBackups

> map[string]interface{} EnableBackups(ctx, linodeId).Execute()

Backups Enable



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
    linodeId := int32(56) // int32 | The ID of the Linode to enable backup service for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.EnableBackups(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.EnableBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableBackups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.EnableBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to enable backup service for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableBackupsRequest struct via the builder pattern


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


## GetBackup

> Backup GetBackup(ctx, linodeId, backupId).Execute()

Backup View



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
    linodeId := int32(56) // int32 | The ID of the Linode the Backup belongs to.
    backupId := int32(56) // int32 | The ID of the Backup to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetBackup(context.Background(), linodeId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackup`: Backup
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode the Backup belongs to. | 
**backupId** | **int32** | The ID of the Backup to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Backup**](Backup.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackups

> InlineResponse20016 GetBackups(ctx, linodeId).Execute()

Backups List



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
    linodeId := int32(56) // int32 | The ID of the Linode the backups belong to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetBackups(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackups`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode the backups belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKernel

> Kernel GetKernel(ctx, kernelId).Execute()

Kernel View



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
    kernelId := "kernelId_example" // string | ID of the Kernel to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetKernel(context.Background(), kernelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetKernel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKernel`: Kernel
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetKernel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kernelId** | **string** | ID of the Kernel to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKernelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kernel**](Kernel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKernels

> InlineResponse20020 GetKernels(ctx).Page(page).PageSize(pageSize).Execute()

Kernels List



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
    resp, r, err := api_client.LinodeInstancesApi.GetKernels(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetKernels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKernels`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetKernels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKernelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfig

> LinodeConfig GetLinodeConfig(ctx, linodeId, configId).Execute()

Configuration Profile View



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
    linodeId := int32(56) // int32 | The ID of the Linode whose Configuration profile to look up.
    configId := int32(56) // int32 | The ID of the Configuration profile to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeConfig(context.Background(), linodeId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeConfig`: LinodeConfig
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode whose Configuration profile to look up. | 
**configId** | **int32** | The ID of the Configuration profile to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LinodeConfig**](LinodeConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfigs

> InlineResponse20017 GetLinodeConfigs(ctx, linodeId).Page(page).PageSize(pageSize).Execute()

Configuration Profiles List



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
    linodeId := int32(56) // int32 | ID of the Linode to look up Configuration profiles for.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeConfigs(context.Background(), linodeId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeConfigs`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up Configuration profiles for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeDisk

> Disk GetLinodeDisk(ctx, linodeId, diskId).Execute()

Disk View



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeDisk(context.Background(), linodeId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Disk**](Disk.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeDisks

> InlineResponse20018 GetLinodeDisks(ctx, linodeId).Page(page).PageSize(pageSize).Execute()

Disks List



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeDisks(context.Background(), linodeId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeDisks`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeFirewalls

> InlineResponse20019 GetLinodeFirewalls(ctx, linodeId).Page(page).PageSize(pageSize).Execute()

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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeFirewalls(context.Background(), linodeId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeFirewalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeFirewalls`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeFirewallsRequest struct via the builder pattern


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


## GetLinodeIP

> IPAddress GetLinodeIP(ctx, linodeId, address).Execute()

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
    linodeId := int32(56) // int32 | The ID of the Linode to look up.
    address := "address_example" // string | The IP address to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeIP(context.Background(), linodeId, address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to look up. | 
**address** | **string** | The IP address to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeIPRequest struct via the builder pattern


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


## GetLinodeIPs

> map[string]interface{} GetLinodeIPs(ctx, linodeId).Execute()

Networking Information List



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeIPs(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeIPs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeIPsRequest struct via the builder pattern


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


## GetLinodeInstance

> Linode GetLinodeInstance(ctx, linodeId).Execute()

Linode View



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
    linodeId := int32(56) // int32 | ID of the Linode to look up

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeInstance(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeInstance`: Linode
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Linode**](Linode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeInstances

> InlineResponse20015 GetLinodeInstances(ctx).Page(page).PageSize(pageSize).Execute()

Linodes List



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
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeInstances(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeInstances`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeStats

> LinodeStats GetLinodeStats(ctx, linodeId).Execute()

Linode Statistics View



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeStats(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeStats`: LinodeStats
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinodeStats**](LinodeStats.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeStatsByYearMonth

> LinodeStats GetLinodeStatsByYearMonth(ctx, linodeId, year, month).Execute()

Statistics View (year/month)



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    year := int32(56) // int32 | Numeric value representing the year to look up.
    month := int32(56) // int32 | Numeric value representing the month to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeStatsByYearMonth(context.Background(), linodeId, year, month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeStatsByYearMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeStatsByYearMonth`: LinodeStats
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeStatsByYearMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**year** | **int32** | Numeric value representing the year to look up. | 
**month** | **int32** | Numeric value representing the month to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeStatsByYearMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LinodeStats**](LinodeStats.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeTransfer

> map[string]interface{} GetLinodeTransfer(ctx, linodeId).Execute()

Network Transfer View



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeTransfer(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTransferRequest struct via the builder pattern


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


## GetLinodeTransferByYearMonth

> map[string]interface{} GetLinodeTransferByYearMonth(ctx, linodeId, year, month).Execute()

Network Transfer View (year/month)



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    year := int32(56) // int32 | Numeric value representing the year to look up.
    month := int32(56) // int32 | Numeric value representing the month to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeTransferByYearMonth(context.Background(), linodeId, year, month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeTransferByYearMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeTransferByYearMonth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeTransferByYearMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**year** | **int32** | Numeric value representing the year to look up. | 
**month** | **int32** | Numeric value representing the month to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTransferByYearMonthRequest struct via the builder pattern


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


## GetLinodeVolumes

> InlineResponse20021 GetLinodeVolumes(ctx, linodeId).Page(page).PageSize(pageSize).Execute()

Linode's Volumes List



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.GetLinodeVolumes(context.Background(), linodeId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.GetLinodeVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinodeVolumes`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.GetLinodeVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateLinodeInstance

> map[string]interface{} MigrateLinodeInstance(ctx, linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

DC Migration/Pending Host Migration Initiate



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
    linodeId := int32(56) // int32 | ID of the Linode to migrate.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.MigrateLinodeInstance(context.Background(), linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.MigrateLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.MigrateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to migrate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

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


## MutateLinodeInstance

> map[string]interface{} MutateLinodeInstance(ctx, linodeId).InlineObject7(inlineObject7).Execute()

Linode Upgrade



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
    linodeId := int32(56) // int32 | ID of the Linode to mutate.
    inlineObject7 := *openapiclient.NewInlineObject7() // InlineObject7 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.MutateLinodeInstance(context.Background(), linodeId).InlineObject7(inlineObject7).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.MutateLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutateLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.MutateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to mutate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject7** | [**InlineObject7**](InlineObject7.md) |  | 

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


## RebootLinodeInstance

> map[string]interface{} RebootLinodeInstance(ctx, linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Linode Reboot



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
    linodeId := int32(56) // int32 | ID of the linode to reboot.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | Optional reboot parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.RebootLinodeInstance(context.Background(), linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.RebootLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.RebootLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the linode to reboot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | Optional reboot parameters. | 

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


## RebuildLinodeInstance

> Linode RebuildLinodeInstance(ctx, linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Linode Rebuild



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
    linodeId := int32(56) // int32 | ID of the Linode to rebuild.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The requested state your Linode will be rebuilt into.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.RebuildLinodeInstance(context.Background(), linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.RebuildLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebuildLinodeInstance`: Linode
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.RebuildLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to rebuild. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The requested state your Linode will be rebuilt into. | 

### Return type

[**Linode**](Linode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLinodeIP

> map[string]interface{} RemoveLinodeIP(ctx, linodeId, address).Execute()

IPv4 Address Delete



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
    linodeId := int32(56) // int32 | The ID of the Linode to look up.
    address := "address_example" // string | The IP address to look up.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.RemoveLinodeIP(context.Background(), linodeId, address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.RemoveLinodeIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLinodeIP`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.RemoveLinodeIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to look up. | 
**address** | **string** | The IP address to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLinodeIPRequest struct via the builder pattern


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


## RescueLinodeInstance

> map[string]interface{} RescueLinodeInstance(ctx, linodeId).InlineObject8(inlineObject8).Execute()

Linode Boot into Rescue Mode



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
    linodeId := int32(56) // int32 | ID of the Linode to rescue.
    inlineObject8 := *openapiclient.NewInlineObject8() // InlineObject8 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.RescueLinodeInstance(context.Background(), linodeId).InlineObject8(inlineObject8).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.RescueLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RescueLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.RescueLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to rescue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescueLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**InlineObject8**](InlineObject8.md) |  | 

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


## ResetDiskPassword

> map[string]interface{} ResetDiskPassword(ctx, linodeId, diskId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Disk Root Password Reset



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The new password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.ResetDiskPassword(context.Background(), linodeId, diskId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.ResetDiskPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetDiskPassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.ResetDiskPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetDiskPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The new password. | 

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


## ResetLinodePassword

> map[string]interface{} ResetLinodePassword(ctx, linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Linode Root Password Reset



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
    linodeId := int32(56) // int32 | ID of the Linode for which to reset its root password.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | This Linode's new root password. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.ResetLinodePassword(context.Background(), linodeId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.ResetLinodePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetLinodePassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.ResetLinodePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode for which to reset its root password. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetLinodePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | This Linode&#39;s new root password. | 

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


## ResizeDisk

> map[string]interface{} ResizeDisk(ctx, linodeId, diskId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Disk Resize



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The new size of the Disk.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.ResizeDisk(context.Background(), linodeId, diskId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.ResizeDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeDisk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.ResizeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The new size of the Disk. | 

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


## ResizeLinodeInstance

> map[string]interface{} ResizeLinodeInstance(ctx, linodeId).InlineObject9(inlineObject9).Execute()

Linode Resize



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
    linodeId := int32(56) // int32 | ID of the Linode to resize.
    inlineObject9 := *openapiclient.NewInlineObject9("g6-standard-2") // InlineObject9 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.ResizeLinodeInstance(context.Background(), linodeId).InlineObject9(inlineObject9).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.ResizeLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.ResizeLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to resize. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject9** | [**InlineObject9**](InlineObject9.md) |  | 

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


## RestoreBackup

> map[string]interface{} RestoreBackup(ctx, linodeId, backupId).InlineObject4(inlineObject4).Execute()

Backup Restore



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
    linodeId := int32(56) // int32 | The ID of the Linode that the Backup belongs to.
    backupId := int32(56) // int32 | The ID of the Backup to restore.
    inlineObject4 := *openapiclient.NewInlineObject4(int32(234)) // InlineObject4 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.RestoreBackup(context.Background(), linodeId, backupId).InlineObject4(inlineObject4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.RestoreBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.RestoreBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode that the Backup belongs to. | 
**backupId** | **int32** | The ID of the Backup to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject4** | [**InlineObject4**](InlineObject4.md) |  | 

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


## ShutdownLinodeInstance

> map[string]interface{} ShutdownLinodeInstance(ctx, linodeId).Execute()

Linode Shut Down



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
    linodeId := int32(56) // int32 | ID of the Linode to shutdown.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.ShutdownLinodeInstance(context.Background(), linodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.ShutdownLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownLinodeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.ShutdownLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to shutdown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownLinodeInstanceRequest struct via the builder pattern


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


## UpdateDisk

> Disk UpdateDisk(ctx, linodeId, diskId).Disk(disk).Execute()

Disk Update



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
    linodeId := int32(56) // int32 | ID of the Linode to look up.
    diskId := int32(56) // int32 | ID of the Disk to look up.
    disk := *openapiclient.NewDisk() // Disk | Updates the parameters of a single Disk. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.UpdateDisk(context.Background(), linodeId, diskId).Disk(disk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.UpdateDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDisk`: Disk
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.UpdateDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disk** | [**Disk**](Disk.md) | Updates the parameters of a single Disk.  | 

### Return type

[**Disk**](Disk.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinodeConfig

> LinodeConfig UpdateLinodeConfig(ctx, linodeId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Configuration Profile Update



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
    linodeId := int32(56) // int32 | The ID of the Linode whose Configuration profile to look up.
    configId := int32(56) // int32 | The ID of the Configuration profile to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The Configuration profile parameters to modify.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.UpdateLinodeConfig(context.Background(), linodeId, configId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.UpdateLinodeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLinodeConfig`: LinodeConfig
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.UpdateLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode whose Configuration profile to look up. | 
**configId** | **int32** | The ID of the Configuration profile to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The Configuration profile parameters to modify. | 

### Return type

[**LinodeConfig**](LinodeConfig.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinodeIP

> IPAddress UpdateLinodeIP(ctx, linodeId, address).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

IP Address Update



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
    linodeId := int32(56) // int32 | The ID of the Linode to look up.
    address := "address_example" // string | The IP address to look up.
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE | The information to update for the IP address. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.UpdateLinodeIP(context.Background(), linodeId, address).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.UpdateLinodeIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLinodeIP`: IPAddress
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.UpdateLinodeIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | The ID of the Linode to look up. | 
**address** | **string** | The IP address to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinodeIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) | The information to update for the IP address. | 

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


## UpdateLinodeInstance

> Linode UpdateLinodeInstance(ctx, linodeId).Linode(linode).Execute()

Linode Update



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
    linodeId := int32(56) // int32 | ID of the Linode to look up
    linode := *openapiclient.NewLinode() // Linode | Any field that is not marked as `readOnly` may be updated. Fields that are marked `readOnly` will be ignored. If any updated field fails to pass validation, the Linode will not be updated. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinodeInstancesApi.UpdateLinodeInstance(context.Background(), linodeId).Linode(linode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesApi.UpdateLinodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLinodeInstance`: Linode
    fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesApi.UpdateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linodeId** | **int32** | ID of the Linode to look up | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linode** | [**Linode**](Linode.md) | Any field that is not marked as &#x60;readOnly&#x60; may be updated. Fields that are marked &#x60;readOnly&#x60; will be ignored. If any updated field fails to pass validation, the Linode will not be updated.  | 

### Return type

[**Linode**](Linode.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

