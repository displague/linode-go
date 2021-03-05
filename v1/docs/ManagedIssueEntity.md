# ManagedIssueEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This ticket&#39;s ID  | [optional] [readonly] 
**Type** | Pointer to **string** | The type of entity this is. In this case, it is always a Ticket.  | [optional] [readonly] 
**Label** | Pointer to **string** | The summary for this Ticket.  | [optional] [readonly] 
**Url** | Pointer to **string** | The relative URL where you can access this Ticket.  | [optional] [readonly] 

## Methods

### NewManagedIssueEntity

`func NewManagedIssueEntity() *ManagedIssueEntity`

NewManagedIssueEntity instantiates a new ManagedIssueEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedIssueEntityWithDefaults

`func NewManagedIssueEntityWithDefaults() *ManagedIssueEntity`

NewManagedIssueEntityWithDefaults instantiates a new ManagedIssueEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedIssueEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedIssueEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedIssueEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedIssueEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ManagedIssueEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedIssueEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedIssueEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagedIssueEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *ManagedIssueEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ManagedIssueEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ManagedIssueEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ManagedIssueEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUrl

`func (o *ManagedIssueEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ManagedIssueEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ManagedIssueEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ManagedIssueEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


