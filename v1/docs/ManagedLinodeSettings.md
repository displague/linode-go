# ManagedLinodeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the Linode these Settings are for.  | [optional] [readonly] 
**Label** | Pointer to **string** | The label of the Linode these Settings are for.  | [optional] [readonly] 
**Group** | Pointer to **string** | The group of the Linode these Settings are for. This is for display purposes only.  | [optional] [readonly] 
**Ssh** | Pointer to [**ManagedLinodeSettingsSsh**](ManagedLinodeSettingsSsh.md) |  | [optional] 

## Methods

### NewManagedLinodeSettings

`func NewManagedLinodeSettings() *ManagedLinodeSettings`

NewManagedLinodeSettings instantiates a new ManagedLinodeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedLinodeSettingsWithDefaults

`func NewManagedLinodeSettingsWithDefaults() *ManagedLinodeSettings`

NewManagedLinodeSettingsWithDefaults instantiates a new ManagedLinodeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedLinodeSettings) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedLinodeSettings) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedLinodeSettings) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedLinodeSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ManagedLinodeSettings) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ManagedLinodeSettings) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ManagedLinodeSettings) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ManagedLinodeSettings) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroup

`func (o *ManagedLinodeSettings) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagedLinodeSettings) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagedLinodeSettings) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagedLinodeSettings) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSsh

`func (o *ManagedLinodeSettings) GetSsh() ManagedLinodeSettingsSsh`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *ManagedLinodeSettings) GetSshOk() (*ManagedLinodeSettingsSsh, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *ManagedLinodeSettings) SetSsh(v ManagedLinodeSettingsSsh)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *ManagedLinodeSettings) HasSsh() bool`

HasSsh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


