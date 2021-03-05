# PayPalExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayerId** | **string** | The PayerID returned by PayPal during the transaction authorization process.  | 
**PaymentId** | **string** | The PaymentID returned from [POST /account/payments/paypal](/docs/api/account/#paypal-payment-stage) that has been approved with PayPal.  | 

## Methods

### NewPayPalExecute

`func NewPayPalExecute(payerId string, paymentId string, ) *PayPalExecute`

NewPayPalExecute instantiates a new PayPalExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayPalExecuteWithDefaults

`func NewPayPalExecuteWithDefaults() *PayPalExecute`

NewPayPalExecuteWithDefaults instantiates a new PayPalExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayerId

`func (o *PayPalExecute) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PayPalExecute) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PayPalExecute) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetPaymentId

`func (o *PayPalExecute) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PayPalExecute) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PayPalExecute) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


