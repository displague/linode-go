# ManagedContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Contact&#39;s unique ID.  | [optional] [readonly] 
**Name** | Pointer to **string** | The name of this Contact.  | [optional] 
**Email** | Pointer to **string** | The address to email this Contact to alert them of issues.  | [optional] 
**Phone** | Pointer to [**ManagedContactPhone**](ManagedContactPhone.md) |  | [optional] 
**Group** | Pointer to **NullableString** | A grouping for this Contact. This is for display purposes only.  | [optional] 
**Updated** | Pointer to **time.Time** | When this Contact was last updated.  | [optional] [readonly] 

## Methods

### NewManagedContact

`func NewManagedContact() *ManagedContact`

NewManagedContact instantiates a new ManagedContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedContactWithDefaults

`func NewManagedContactWithDefaults() *ManagedContact`

NewManagedContactWithDefaults instantiates a new ManagedContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagedContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ManagedContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ManagedContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ManagedContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ManagedContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ManagedContact) GetPhone() ManagedContactPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ManagedContact) GetPhoneOk() (*ManagedContactPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ManagedContact) SetPhone(v ManagedContactPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ManagedContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetGroup

`func (o *ManagedContact) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagedContact) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagedContact) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagedContact) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *ManagedContact) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *ManagedContact) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUpdated

`func (o *ManagedContact) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ManagedContact) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ManagedContact) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ManagedContact) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


