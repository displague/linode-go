# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** | The paypal-generated ID for this Payment. Used when authorizing the Payment in PayPal&#39;s interface.  | [optional] 
**CheckoutToken** | Pointer to **string** | The checkout token generated for this Payment.  | [optional] [readonly] 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010() *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *InlineResponse20010) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *InlineResponse20010) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *InlineResponse20010) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *InlineResponse20010) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetCheckoutToken

`func (o *InlineResponse20010) GetCheckoutToken() string`

GetCheckoutToken returns the CheckoutToken field if non-nil, zero value otherwise.

### GetCheckoutTokenOk

`func (o *InlineResponse20010) GetCheckoutTokenOk() (*string, bool)`

GetCheckoutTokenOk returns a tuple with the CheckoutToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutToken

`func (o *InlineResponse20010) SetCheckoutToken(v string)`

SetCheckoutToken sets CheckoutToken field to given value.

### HasCheckoutToken

`func (o *InlineResponse20010) HasCheckoutToken() bool`

HasCheckoutToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


