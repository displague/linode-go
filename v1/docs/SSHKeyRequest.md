# SSHKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A label for the key.  | [optional] 
**SshKey** | Pointer to **string** | The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.  | [optional] 

## Methods

### NewSSHKeyRequest

`func NewSSHKeyRequest() *SSHKeyRequest`

NewSSHKeyRequest instantiates a new SSHKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyRequestWithDefaults

`func NewSSHKeyRequestWithDefaults() *SSHKeyRequest`

NewSSHKeyRequestWithDefaults instantiates a new SSHKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SSHKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SSHKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SSHKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SSHKeyRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSshKey

`func (o *SSHKeyRequest) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *SSHKeyRequest) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *SSHKeyRequest) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *SSHKeyRequest) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


