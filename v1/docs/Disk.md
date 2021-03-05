# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Disk&#39;s ID which must be provided for all operations impacting this Disk.  | [optional] [readonly] 
**Label** | Pointer to **string** | The Disk&#39;s label is for display purposes only.  | [optional] 
**Status** | Pointer to **string** | A brief description of this Disk&#39;s current state. This field may change without direct action from you, as a result of operations performed to the Disk or the Linode containing the Disk.  | [optional] [readonly] 
**Size** | Pointer to **int32** | The size of the Disk in MB. | [optional] [readonly] 
**Filesystem** | Pointer to **string** | The Disk filesystem can be one of:    * raw - No filesystem, just a raw binary stream.   * swap - Linux swap area.   * ext3 - The ext3 journaling filesystem for Linux.   * ext4 - The ext4 journaling filesystem for Linux.   * initrd - initrd (uncompressed initrd, ext2, max 32 MB).  | [optional] 
**Created** | Pointer to **time.Time** | When this Disk was created. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Disk was last updated. | [optional] [readonly] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Disk) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Disk) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Disk) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Disk) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *Disk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Disk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Disk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Disk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSize

`func (o *Disk) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Disk) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Disk) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Disk) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFilesystem

`func (o *Disk) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *Disk) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *Disk) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *Disk) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetCreated

`func (o *Disk) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Disk) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Disk) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Disk) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Disk) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Disk) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Disk) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Disk) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


