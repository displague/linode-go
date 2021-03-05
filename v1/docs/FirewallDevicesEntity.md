# FirewallDevicesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The entity&#39;s ID | [optional] 
**Type** | Pointer to **string** | The entity&#39;s type. | [optional] 
**Label** | Pointer to **string** | The entity&#39;s label. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL you can use to access this entity.  | [optional] [readonly] 

## Methods

### NewFirewallDevicesEntity

`func NewFirewallDevicesEntity() *FirewallDevicesEntity`

NewFirewallDevicesEntity instantiates a new FirewallDevicesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallDevicesEntityWithDefaults

`func NewFirewallDevicesEntityWithDefaults() *FirewallDevicesEntity`

NewFirewallDevicesEntityWithDefaults instantiates a new FirewallDevicesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallDevicesEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallDevicesEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallDevicesEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallDevicesEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *FirewallDevicesEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallDevicesEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallDevicesEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallDevicesEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *FirewallDevicesEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FirewallDevicesEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FirewallDevicesEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FirewallDevicesEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUrl

`func (o *FirewallDevicesEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FirewallDevicesEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FirewallDevicesEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FirewallDevicesEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


