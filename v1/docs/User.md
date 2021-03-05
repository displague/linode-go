# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | This User&#39;s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts).  | [optional] 
**Email** | Pointer to **string** | The email address for this User, for account management communications, and may be used for other communications as configured.  | [optional] [readonly] 
**Restricted** | Pointer to **bool** | If true, this User must be granted access to perform actions or access entities on this Account. See [/account/users/{username}/grants](/docs/api/account/#users-grants-view) for details on how to configure grants for a restricted User.  | [optional] 
**SshKeys** | Pointer to **[]string** | A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the &#x60;authorized_users&#x60; field of a [create Linode](/docs/api/linode-instances/#linode-create), [rebuild Linode](/docs/api/linode-instances/#linode-rebuild), or [create Disk](/docs/api/linode-instances/#disk-create) request.  | [optional] 
**TfaEnabled** | Pointer to **bool** | A boolean value indicating if the User has Two Factor Authentication (TFA) enabled. See the Create Two Factor Secret ([/profile/tfa-enable](/docs/api/profile/#two-factor-secret-create)) endpoint to enable TFA.  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRestricted

`func (o *User) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *User) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *User) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *User) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetSshKeys

`func (o *User) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *User) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *User) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *User) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetTfaEnabled

`func (o *User) GetTfaEnabled() bool`

GetTfaEnabled returns the TfaEnabled field if non-nil, zero value otherwise.

### GetTfaEnabledOk

`func (o *User) GetTfaEnabledOk() (*bool, bool)`

GetTfaEnabledOk returns a tuple with the TfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaEnabled

`func (o *User) SetTfaEnabled(v bool)`

SetTfaEnabled sets TfaEnabled field to given value.

### HasTfaEnabled

`func (o *User) HasTfaEnabled() bool`

HasTfaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


