# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvv** | Pointer to **string** | CVV (Card Verification Value) of the credit card to be used for the Payment.  | [optional] 
**Usd** | **string** | The amount in US Dollars of the Payment. The maximum credit card payment that can be made is $50,000 dollars.  | 

## Methods

### NewPaymentRequest

`func NewPaymentRequest(usd string, ) *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvv

`func (o *PaymentRequest) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *PaymentRequest) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *PaymentRequest) SetCvv(v string)`

SetCvv sets Cvv field to given value.

### HasCvv

`func (o *PaymentRequest) HasCvv() bool`

HasCvv returns a boolean if a field has been set.

### GetUsd

`func (o *PaymentRequest) GetUsd() string`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *PaymentRequest) GetUsdOk() (*string, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *PaymentRequest) SetUsd(v string)`

SetUsd sets Usd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


