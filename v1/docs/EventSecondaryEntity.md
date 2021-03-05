# EventSecondaryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the object that is the secondary entity.  | [optional] 
**Label** | Pointer to **string** | The label of this object.  | [optional] 
**Type** | Pointer to **string** | The type of entity that is being referenced by the Event.  | [optional] [readonly] 
**Url** | Pointer to **string** | The URL where you can access the object this Event is for. If a relative URL, it is relative to the domain you retrieved the Event from.  | [optional] 

## Methods

### NewEventSecondaryEntity

`func NewEventSecondaryEntity() *EventSecondaryEntity`

NewEventSecondaryEntity instantiates a new EventSecondaryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSecondaryEntityWithDefaults

`func NewEventSecondaryEntityWithDefaults() *EventSecondaryEntity`

NewEventSecondaryEntityWithDefaults instantiates a new EventSecondaryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventSecondaryEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSecondaryEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSecondaryEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventSecondaryEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *EventSecondaryEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EventSecondaryEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EventSecondaryEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EventSecondaryEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *EventSecondaryEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSecondaryEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSecondaryEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventSecondaryEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *EventSecondaryEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventSecondaryEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventSecondaryEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventSecondaryEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


