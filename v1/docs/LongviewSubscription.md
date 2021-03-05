# LongviewSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of this Subscription tier.  | [optional] [readonly] 
**Price** | Pointer to [**LongviewSubscriptionPrice**](LongviewSubscriptionPrice.md) |  | [optional] 
**Label** | Pointer to **string** | A display name for this Subscription tier.  | [optional] [readonly] 
**ClientsIncluded** | Pointer to **int32** | The number of Longview Clients that may be created with this Subscription tier.  | [optional] [readonly] 

## Methods

### NewLongviewSubscription

`func NewLongviewSubscription() *LongviewSubscription`

NewLongviewSubscription instantiates a new LongviewSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongviewSubscriptionWithDefaults

`func NewLongviewSubscriptionWithDefaults() *LongviewSubscription`

NewLongviewSubscriptionWithDefaults instantiates a new LongviewSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LongviewSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LongviewSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LongviewSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LongviewSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *LongviewSubscription) GetPrice() LongviewSubscriptionPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LongviewSubscription) GetPriceOk() (*LongviewSubscriptionPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LongviewSubscription) SetPrice(v LongviewSubscriptionPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LongviewSubscription) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLabel

`func (o *LongviewSubscription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LongviewSubscription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LongviewSubscription) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LongviewSubscription) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetClientsIncluded

`func (o *LongviewSubscription) GetClientsIncluded() int32`

GetClientsIncluded returns the ClientsIncluded field if non-nil, zero value otherwise.

### GetClientsIncludedOk

`func (o *LongviewSubscription) GetClientsIncludedOk() (*int32, bool)`

GetClientsIncludedOk returns a tuple with the ClientsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsIncluded

`func (o *LongviewSubscription) SetClientsIncluded(v int32)`

SetClientsIncluded sets ClientsIncluded field to given value.

### HasClientsIncluded

`func (o *LongviewSubscription) HasClientsIncluded() bool`

HasClientsIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


