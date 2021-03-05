# FirewallDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Device&#39;s unique ID  | [optional] 
**Created** | Pointer to **time.Time** | When this Device was created.  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Device was last updated.  | [optional] [readonly] 
**Entity** | Pointer to [**FirewallDevicesEntity**](FirewallDevicesEntity.md) |  | [optional] 

## Methods

### NewFirewallDevices

`func NewFirewallDevices() *FirewallDevices`

NewFirewallDevices instantiates a new FirewallDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallDevicesWithDefaults

`func NewFirewallDevicesWithDefaults() *FirewallDevices`

NewFirewallDevicesWithDefaults instantiates a new FirewallDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallDevices) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallDevices) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallDevices) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallDevices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *FirewallDevices) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FirewallDevices) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FirewallDevices) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FirewallDevices) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FirewallDevices) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FirewallDevices) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FirewallDevices) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FirewallDevices) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetEntity

`func (o *FirewallDevices) GetEntity() FirewallDevicesEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *FirewallDevices) GetEntityOk() (*FirewallDevicesEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *FirewallDevices) SetEntity(v FirewallDevicesEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *FirewallDevices) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


