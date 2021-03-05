# LongviewSubscriptionPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hourly** | Pointer to **float32** | The hourly price, in US dollars, for this Subscription tier.  | [optional] [readonly] 
**Monthly** | Pointer to **float32** | The maximum monthly price in US Dollars for this Subscription tier. You will never be charged more than this amount per month for this subscription.  | [optional] [readonly] 

## Methods

### NewLongviewSubscriptionPrice

`func NewLongviewSubscriptionPrice() *LongviewSubscriptionPrice`

NewLongviewSubscriptionPrice instantiates a new LongviewSubscriptionPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongviewSubscriptionPriceWithDefaults

`func NewLongviewSubscriptionPriceWithDefaults() *LongviewSubscriptionPrice`

NewLongviewSubscriptionPriceWithDefaults instantiates a new LongviewSubscriptionPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourly

`func (o *LongviewSubscriptionPrice) GetHourly() float32`

GetHourly returns the Hourly field if non-nil, zero value otherwise.

### GetHourlyOk

`func (o *LongviewSubscriptionPrice) GetHourlyOk() (*float32, bool)`

GetHourlyOk returns a tuple with the Hourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourly

`func (o *LongviewSubscriptionPrice) SetHourly(v float32)`

SetHourly sets Hourly field to given value.

### HasHourly

`func (o *LongviewSubscriptionPrice) HasHourly() bool`

HasHourly returns a boolean if a field has been set.

### GetMonthly

`func (o *LongviewSubscriptionPrice) GetMonthly() float32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *LongviewSubscriptionPrice) GetMonthlyOk() (*float32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *LongviewSubscriptionPrice) SetMonthly(v float32)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *LongviewSubscriptionPrice) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


