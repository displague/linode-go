# ManagedCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Credential&#39;s unique ID.  | [optional] [readonly] 
**Label** | Pointer to **string** | The unique label for this Credential. This is for display purposes only.  | [optional] 
**LastDecrypted** | Pointer to **time.Time** | The date this Credential was last decrypted by a member of Linode special forces.  | [optional] [readonly] 

## Methods

### NewManagedCredential

`func NewManagedCredential() *ManagedCredential`

NewManagedCredential instantiates a new ManagedCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCredentialWithDefaults

`func NewManagedCredentialWithDefaults() *ManagedCredential`

NewManagedCredentialWithDefaults instantiates a new ManagedCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedCredential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedCredential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedCredential) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ManagedCredential) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ManagedCredential) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ManagedCredential) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ManagedCredential) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastDecrypted

`func (o *ManagedCredential) GetLastDecrypted() time.Time`

GetLastDecrypted returns the LastDecrypted field if non-nil, zero value otherwise.

### GetLastDecryptedOk

`func (o *ManagedCredential) GetLastDecryptedOk() (*time.Time, bool)`

GetLastDecryptedOk returns a tuple with the LastDecrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDecrypted

`func (o *ManagedCredential) SetLastDecrypted(v time.Time)`

SetLastDecrypted sets LastDecrypted field to given value.

### HasLastDecrypted

`func (o *ManagedCredential) HasLastDecrypted() bool`

HasLastDecrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


