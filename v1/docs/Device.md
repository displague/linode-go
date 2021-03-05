# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | Pointer to **int32** | The Disk ID, or &#x60;null&#x60; if a Volume is assigned to this slot. | [optional] 
**VolumeId** | Pointer to **int32** | The Volume ID, or &#x60;null&#x60; if a Disk is assigned to this slot. | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskId

`func (o *Device) GetDiskId() int32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *Device) GetDiskIdOk() (*int32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *Device) SetDiskId(v int32)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *Device) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetVolumeId

`func (o *Device) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Device) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *Device) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *Device) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


