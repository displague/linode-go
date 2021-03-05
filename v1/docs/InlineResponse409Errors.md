# InlineResponse409Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | A string explaining that the account could not be cancelled because there is an outstanding balance on the account that must be paid first.  | [optional] 

## Methods

### NewInlineResponse409Errors

`func NewInlineResponse409Errors() *InlineResponse409Errors`

NewInlineResponse409Errors instantiates a new InlineResponse409Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse409ErrorsWithDefaults

`func NewInlineResponse409ErrorsWithDefaults() *InlineResponse409Errors`

NewInlineResponse409ErrorsWithDefaults instantiates a new InlineResponse409Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *InlineResponse409Errors) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse409Errors) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse409Errors) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse409Errors) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


