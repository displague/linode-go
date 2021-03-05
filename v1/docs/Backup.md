# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of this Backup. | [optional] [readonly] 
**Type** | Pointer to **string** | This indicates whether the Backup is an automatic Backup or manual snapshot taken by the User at a specific point in time.  | [optional] [readonly] 
**Status** | Pointer to **string** | The current state of a specific Backup. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date the Backup was taken. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The date the Backup was most recently updated. | [optional] [readonly] 
**Finished** | Pointer to **time.Time** | The date the Backup completed. | [optional] [readonly] 
**Label** | Pointer to **NullableString** | A label for Backups that are of type &#x60;snapshot&#x60;. | [optional] 
**Configs** | Pointer to **[]string** | A list of the labels of the Configuration profiles that are part of the Backup.  | [optional] [readonly] 
**Disks** | Pointer to [**[]BackupDisks**](BackupDisks.md) | A list of the disks that are part of the Backup.  | [optional] [readonly] 

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Backup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Backup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Backup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Backup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Backup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Backup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Backup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Backup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Backup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *Backup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Backup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Backup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Backup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Backup) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Backup) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Backup) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Backup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetFinished

`func (o *Backup) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *Backup) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *Backup) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *Backup) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetLabel

`func (o *Backup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Backup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Backup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Backup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Backup) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Backup) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetConfigs

`func (o *Backup) GetConfigs() []string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *Backup) GetConfigsOk() (*[]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *Backup) SetConfigs(v []string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *Backup) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDisks

`func (o *Backup) GetDisks() []BackupDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Backup) GetDisksOk() (*[]BackupDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Backup) SetDisks(v []BackupDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Backup) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


