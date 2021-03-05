# Kernel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of this Kernel. | [optional] [readonly] 
**Label** | Pointer to **string** | The friendly name of this Kernel. | [optional] [readonly] 
**Version** | Pointer to **string** | Linux Kernel version. | [optional] [readonly] 
**Kvm** | Pointer to **bool** | If this Kernel is suitable for KVM Linodes. | [optional] [readonly] 
**Xen** | Pointer to **bool** | If this Kernel is suitable for Xen Linodes. | [optional] [readonly] 
**Architecture** | Pointer to **string** | The architecture of this Kernel. | [optional] [readonly] 
**Pvops** | Pointer to **bool** | If this Kernel is suitable for paravirtualized operations. | [optional] [readonly] 
**Deprecated** | Pointer to **bool** | If this Kernel is marked as deprecated, this field has a value of true; otherwise, this field is false. | [optional] [readonly] 

## Methods

### NewKernel

`func NewKernel() *Kernel`

NewKernel instantiates a new Kernel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKernelWithDefaults

`func NewKernelWithDefaults() *Kernel`

NewKernelWithDefaults instantiates a new Kernel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kernel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kernel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kernel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kernel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Kernel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Kernel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Kernel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Kernel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetVersion

`func (o *Kernel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Kernel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Kernel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Kernel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKvm

`func (o *Kernel) GetKvm() bool`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *Kernel) GetKvmOk() (*bool, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *Kernel) SetKvm(v bool)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *Kernel) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetXen

`func (o *Kernel) GetXen() bool`

GetXen returns the Xen field if non-nil, zero value otherwise.

### GetXenOk

`func (o *Kernel) GetXenOk() (*bool, bool)`

GetXenOk returns a tuple with the Xen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXen

`func (o *Kernel) SetXen(v bool)`

SetXen sets Xen field to given value.

### HasXen

`func (o *Kernel) HasXen() bool`

HasXen returns a boolean if a field has been set.

### GetArchitecture

`func (o *Kernel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Kernel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Kernel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Kernel) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetPvops

`func (o *Kernel) GetPvops() bool`

GetPvops returns the Pvops field if non-nil, zero value otherwise.

### GetPvopsOk

`func (o *Kernel) GetPvopsOk() (*bool, bool)`

GetPvopsOk returns a tuple with the Pvops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvops

`func (o *Kernel) SetPvops(v bool)`

SetPvops sets Pvops field to given value.

### HasPvops

`func (o *Kernel) HasPvops() bool`

HasPvops returns a boolean if a field has been set.

### GetDeprecated

`func (o *Kernel) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Kernel) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Kernel) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Kernel) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


