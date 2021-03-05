# MaintenanceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The label of the entity being affected by maintenance.  | [optional] 
**Id** | Pointer to **float32** | The id of the entity being affected by maintenance.  | [optional] 
**Type** | Pointer to **string** | The type of entity.  | [optional] 
**Url** | Pointer to **string** | The API endpoint prefix to use in combination with the entity id to find specific information about the entity.  | [optional] 

## Methods

### NewMaintenanceEntity

`func NewMaintenanceEntity() *MaintenanceEntity`

NewMaintenanceEntity instantiates a new MaintenanceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceEntityWithDefaults

`func NewMaintenanceEntityWithDefaults() *MaintenanceEntity`

NewMaintenanceEntityWithDefaults instantiates a new MaintenanceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *MaintenanceEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MaintenanceEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MaintenanceEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MaintenanceEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetId

`func (o *MaintenanceEntity) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceEntity) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceEntity) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MaintenanceEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaintenanceEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaintenanceEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MaintenanceEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *MaintenanceEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MaintenanceEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MaintenanceEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MaintenanceEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


