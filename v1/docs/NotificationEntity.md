# NotificationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of the Notification&#39;s entity, based on the entity type.  | [optional] 
**Label** | Pointer to **string** | The current label for this Notification&#39;s entity.  | [optional] 
**Type** | Pointer to **string** | The type of entity this is related to. | [optional] 
**Url** | Pointer to **string** | The URL where you can access the object this Notification is for. If a relative URL, it is relative to the domain you retrieved the Notification from.  | [optional] 

## Methods

### NewNotificationEntity

`func NewNotificationEntity() *NotificationEntity`

NewNotificationEntity instantiates a new NotificationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEntityWithDefaults

`func NewNotificationEntityWithDefaults() *NotificationEntity`

NewNotificationEntityWithDefaults instantiates a new NotificationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *NotificationEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NotificationEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NotificationEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *NotificationEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *NotificationEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


