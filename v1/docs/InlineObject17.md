# InlineObject17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinodeId** | [**LinodeId**](LinodeId.md) |  | 
**ConfigId** | Pointer to **int32** | The ID of the Linode Config to include this Volume in. Must belong to the Linode referenced by &#x60;linode_id&#x60;. If not given, the last booted Config will be chosen.  | [optional] 
**PersistAcrossBoots** | Pointer to **bool** | Defaults to true, if false is provided, the Volume will not be attached to the Linode Config. In this case more than 8 Volumes may be attached to a Linode if a Linode has 16GB of RAM or more. The number of volumes that can be attached is equal to the number of GB of RAM that the Linode has, up to a maximum of 64. &#x60;config_id&#x60; should not be passed if this is set to false and linode_id must be passed. The Linode must be running.  | [optional] 

## Methods

### NewInlineObject17

`func NewInlineObject17(linodeId LinodeId, ) *InlineObject17`

NewInlineObject17 instantiates a new InlineObject17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject17WithDefaults

`func NewInlineObject17WithDefaults() *InlineObject17`

NewInlineObject17WithDefaults instantiates a new InlineObject17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodeId

`func (o *InlineObject17) GetLinodeId() LinodeId`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *InlineObject17) GetLinodeIdOk() (*LinodeId, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *InlineObject17) SetLinodeId(v LinodeId)`

SetLinodeId sets LinodeId field to given value.


### GetConfigId

`func (o *InlineObject17) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *InlineObject17) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *InlineObject17) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *InlineObject17) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetPersistAcrossBoots

`func (o *InlineObject17) GetPersistAcrossBoots() bool`

GetPersistAcrossBoots returns the PersistAcrossBoots field if non-nil, zero value otherwise.

### GetPersistAcrossBootsOk

`func (o *InlineObject17) GetPersistAcrossBootsOk() (*bool, bool)`

GetPersistAcrossBootsOk returns a tuple with the PersistAcrossBoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistAcrossBoots

`func (o *InlineObject17) SetPersistAcrossBoots(v bool)`

SetPersistAcrossBoots sets PersistAcrossBoots field to given value.

### HasPersistAcrossBoots

`func (o *InlineObject17) HasPersistAcrossBoots() bool`

HasPersistAcrossBoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


