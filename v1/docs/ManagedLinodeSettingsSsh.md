# ManagedLinodeSettingsSsh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **bool** | If true, Linode special forces may access this Linode over ssh to respond to Issues.  | [optional] 
**User** | Pointer to **string** | The user Linode&#39;s special forces should use when accessing this Linode to respond to an issue.  | [optional] 
**Ip** | Pointer to **string** | The IP Linode special forces should use to access this Linode when responding to an Issue.  | [optional] 
**Port** | Pointer to **int32** | The port Linode special forces should use to access this Linode over ssh to respond to an Issue.  | [optional] 

## Methods

### NewManagedLinodeSettingsSsh

`func NewManagedLinodeSettingsSsh() *ManagedLinodeSettingsSsh`

NewManagedLinodeSettingsSsh instantiates a new ManagedLinodeSettingsSsh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedLinodeSettingsSshWithDefaults

`func NewManagedLinodeSettingsSshWithDefaults() *ManagedLinodeSettingsSsh`

NewManagedLinodeSettingsSshWithDefaults instantiates a new ManagedLinodeSettingsSsh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ManagedLinodeSettingsSsh) GetAccess() bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ManagedLinodeSettingsSsh) GetAccessOk() (*bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ManagedLinodeSettingsSsh) SetAccess(v bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ManagedLinodeSettingsSsh) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetUser

`func (o *ManagedLinodeSettingsSsh) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ManagedLinodeSettingsSsh) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ManagedLinodeSettingsSsh) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ManagedLinodeSettingsSsh) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIp

`func (o *ManagedLinodeSettingsSsh) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ManagedLinodeSettingsSsh) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ManagedLinodeSettingsSsh) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ManagedLinodeSettingsSsh) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *ManagedLinodeSettingsSsh) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagedLinodeSettingsSsh) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagedLinodeSettingsSsh) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ManagedLinodeSettingsSsh) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


