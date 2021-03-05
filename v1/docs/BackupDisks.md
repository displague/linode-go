# BackupDisks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** |  | [optional] 
**Filesystem** | Pointer to [**Filesystem**](Filesystem.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupDisks

`func NewBackupDisks() *BackupDisks`

NewBackupDisks instantiates a new BackupDisks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDisksWithDefaults

`func NewBackupDisksWithDefaults() *BackupDisks`

NewBackupDisksWithDefaults instantiates a new BackupDisks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *BackupDisks) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BackupDisks) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BackupDisks) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BackupDisks) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFilesystem

`func (o *BackupDisks) GetFilesystem() Filesystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *BackupDisks) GetFilesystemOk() (*Filesystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *BackupDisks) SetFilesystem(v Filesystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *BackupDisks) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetLabel

`func (o *BackupDisks) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BackupDisks) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BackupDisks) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BackupDisks) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


