# ManagedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Issue&#39;s unique ID.  | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Issue was created. Issues are created in response to issues detected with Managed Services, so this is also when the Issue was detected.  | [optional] [readonly] 
**Services** | Pointer to **[]int32** | An array of Managed Service IDs that were affected by this Issue.  | [optional] [readonly] 
**Entity** | Pointer to [**ManagedIssueEntity**](ManagedIssueEntity.md) |  | [optional] 

## Methods

### NewManagedIssue

`func NewManagedIssue() *ManagedIssue`

NewManagedIssue instantiates a new ManagedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedIssueWithDefaults

`func NewManagedIssueWithDefaults() *ManagedIssue`

NewManagedIssueWithDefaults instantiates a new ManagedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedIssue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *ManagedIssue) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ManagedIssue) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ManagedIssue) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ManagedIssue) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetServices

`func (o *ManagedIssue) GetServices() []int32`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ManagedIssue) GetServicesOk() (*[]int32, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ManagedIssue) SetServices(v []int32)`

SetServices sets Services field to given value.

### HasServices

`func (o *ManagedIssue) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEntity

`func (o *ManagedIssue) GetEntity() ManagedIssueEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ManagedIssue) GetEntityOk() (*ManagedIssueEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ManagedIssue) SetEntity(v ManagedIssueEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ManagedIssue) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


