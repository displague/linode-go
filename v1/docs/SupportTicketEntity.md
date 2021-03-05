# SupportTicketEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for this Ticket&#39;s entity.  | [optional] [readonly] 
**Label** | Pointer to **string** | The current label of this entity.  | [optional] [readonly] 
**Type** | Pointer to **string** | The type of entity this is related to.  | [optional] [readonly] 
**Url** | Pointer to **string** | The URL where you can access the object this event is for. If a relative URL, it is relative to the domain you retrieved the entity from.  | [optional] [readonly] 

## Methods

### NewSupportTicketEntity

`func NewSupportTicketEntity() *SupportTicketEntity`

NewSupportTicketEntity instantiates a new SupportTicketEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketEntityWithDefaults

`func NewSupportTicketEntityWithDefaults() *SupportTicketEntity`

NewSupportTicketEntityWithDefaults instantiates a new SupportTicketEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SupportTicketEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SupportTicketEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SupportTicketEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SupportTicketEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *SupportTicketEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportTicketEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportTicketEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SupportTicketEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *SupportTicketEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupportTicketEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupportTicketEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupportTicketEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


