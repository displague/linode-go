# AccountCreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastFour** | Pointer to **string** | The last four digits of the credit card associated with this Account.  | [optional] 
**Expiry** | Pointer to **string** | The expiration month and year of the credit card. | [optional] 

## Methods

### NewAccountCreditCard

`func NewAccountCreditCard() *AccountCreditCard`

NewAccountCreditCard instantiates a new AccountCreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreditCardWithDefaults

`func NewAccountCreditCardWithDefaults() *AccountCreditCard`

NewAccountCreditCardWithDefaults instantiates a new AccountCreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastFour

`func (o *AccountCreditCard) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *AccountCreditCard) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *AccountCreditCard) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *AccountCreditCard) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetExpiry

`func (o *AccountCreditCard) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AccountCreditCard) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AccountCreditCard) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AccountCreditCard) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


