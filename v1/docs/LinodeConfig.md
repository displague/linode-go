# LinodeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of this Config. | [optional] [readonly] 
**Kernel** | Pointer to **string** | A Kernel ID to boot a Linode with. Defaults to \&quot;linode/latest-64bit\&quot;. | [optional] 
**Comments** | Pointer to **NullableString** | Optional field for arbitrary User comments on this Config. | [optional] 
**MemoryLimit** | Pointer to **int32** | Defaults to the total RAM of the Linode.  | [optional] 
**RunLevel** | Pointer to **string** | Defines the state of your Linode after booting. Defaults to &#x60;default&#x60;.  | [optional] 
**VirtMode** | Pointer to **string** | Controls the virtualization mode. Defaults to &#x60;paravirt&#x60;. * &#x60;paravirt&#x60; is suitable for most cases. Linodes running in paravirt mode   share some qualities with the host, ultimately making it run faster since   there is less transition between it and the host. * &#x60;full_virt&#x60; affords more customization, but is slower because 100% of the VM   is virtualized.  | [optional] 
**Helpers** | Pointer to [**LinodeConfigHelpers**](LinodeConfigHelpers.md) |  | [optional] 
**Label** | **string** | The Config&#39;s label is for display purposes only.  | 
**Devices** | [**Devices**](Devices.md) |  | 
**RootDevice** | Pointer to **string** | The root device to boot. * If no value or an invalid value is provided, root device will default to &#x60;/dev/sda&#x60;. * If the device specified at the root device location is not mounted, the Linode will not boot until a device is mounted.  | [optional] 

## Methods

### NewLinodeConfig

`func NewLinodeConfig(label string, devices Devices, ) *LinodeConfig`

NewLinodeConfig instantiates a new LinodeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeConfigWithDefaults

`func NewLinodeConfigWithDefaults() *LinodeConfig`

NewLinodeConfigWithDefaults instantiates a new LinodeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinodeConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinodeConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinodeConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LinodeConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKernel

`func (o *LinodeConfig) GetKernel() string`

GetKernel returns the Kernel field if non-nil, zero value otherwise.

### GetKernelOk

`func (o *LinodeConfig) GetKernelOk() (*string, bool)`

GetKernelOk returns a tuple with the Kernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernel

`func (o *LinodeConfig) SetKernel(v string)`

SetKernel sets Kernel field to given value.

### HasKernel

`func (o *LinodeConfig) HasKernel() bool`

HasKernel returns a boolean if a field has been set.

### GetComments

`func (o *LinodeConfig) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *LinodeConfig) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *LinodeConfig) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *LinodeConfig) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *LinodeConfig) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *LinodeConfig) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetMemoryLimit

`func (o *LinodeConfig) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *LinodeConfig) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *LinodeConfig) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *LinodeConfig) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetRunLevel

`func (o *LinodeConfig) GetRunLevel() string`

GetRunLevel returns the RunLevel field if non-nil, zero value otherwise.

### GetRunLevelOk

`func (o *LinodeConfig) GetRunLevelOk() (*string, bool)`

GetRunLevelOk returns a tuple with the RunLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLevel

`func (o *LinodeConfig) SetRunLevel(v string)`

SetRunLevel sets RunLevel field to given value.

### HasRunLevel

`func (o *LinodeConfig) HasRunLevel() bool`

HasRunLevel returns a boolean if a field has been set.

### GetVirtMode

`func (o *LinodeConfig) GetVirtMode() string`

GetVirtMode returns the VirtMode field if non-nil, zero value otherwise.

### GetVirtModeOk

`func (o *LinodeConfig) GetVirtModeOk() (*string, bool)`

GetVirtModeOk returns a tuple with the VirtMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtMode

`func (o *LinodeConfig) SetVirtMode(v string)`

SetVirtMode sets VirtMode field to given value.

### HasVirtMode

`func (o *LinodeConfig) HasVirtMode() bool`

HasVirtMode returns a boolean if a field has been set.

### GetHelpers

`func (o *LinodeConfig) GetHelpers() LinodeConfigHelpers`

GetHelpers returns the Helpers field if non-nil, zero value otherwise.

### GetHelpersOk

`func (o *LinodeConfig) GetHelpersOk() (*LinodeConfigHelpers, bool)`

GetHelpersOk returns a tuple with the Helpers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpers

`func (o *LinodeConfig) SetHelpers(v LinodeConfigHelpers)`

SetHelpers sets Helpers field to given value.

### HasHelpers

`func (o *LinodeConfig) HasHelpers() bool`

HasHelpers returns a boolean if a field has been set.

### GetLabel

`func (o *LinodeConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinodeConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinodeConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDevices

`func (o *LinodeConfig) GetDevices() Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LinodeConfig) GetDevicesOk() (*Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LinodeConfig) SetDevices(v Devices)`

SetDevices sets Devices field to given value.


### GetRootDevice

`func (o *LinodeConfig) GetRootDevice() string`

GetRootDevice returns the RootDevice field if non-nil, zero value otherwise.

### GetRootDeviceOk

`func (o *LinodeConfig) GetRootDeviceOk() (*string, bool)`

GetRootDeviceOk returns a tuple with the RootDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDevice

`func (o *LinodeConfig) SetRootDevice(v string)`

SetRootDevice sets RootDevice field to given value.

### HasRootDevice

`func (o *LinodeConfig) HasRootDevice() bool`

HasRootDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


