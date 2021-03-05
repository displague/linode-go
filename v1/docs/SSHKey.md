# SSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identifier of an SSH Key object.  | [optional] [readonly] 
**Label** | Pointer to **string** | A label for the SSH Key.  | [optional] 
**SshKey** | Pointer to **string** | The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.  | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date this key was added.  | [optional] [readonly] 

## Methods

### NewSSHKey

`func NewSSHKey() *SSHKey`

NewSSHKey instantiates a new SSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyWithDefaults

`func NewSSHKeyWithDefaults() *SSHKey`

NewSSHKeyWithDefaults instantiates a new SSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSHKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SSHKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SSHKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SSHKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SSHKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SSHKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSshKey

`func (o *SSHKey) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *SSHKey) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *SSHKey) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *SSHKey) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetCreated

`func (o *SSHKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SSHKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SSHKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SSHKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


