# UserDefinedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A human-readable label for the field that will serve as the input prompt for entering the value during deployment.  | [readonly] 
**Name** | **string** | The name of the field.  | [readonly] 
**Example** | **string** | An example value for the field.  | [readonly] 
**OneOf** | Pointer to **string** | A list of acceptable single values for the field.  | [optional] [readonly] 
**ManyOf** | Pointer to **string** | A list of acceptable values for the field in any quantity, combination or order.  | [optional] [readonly] 
**Default** | Pointer to **string** | The default value.  If not specified, this value will be used.  | [optional] [readonly] 

## Methods

### NewUserDefinedField

`func NewUserDefinedField(label string, name string, example string, ) *UserDefinedField`

NewUserDefinedField instantiates a new UserDefinedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldWithDefaults

`func NewUserDefinedFieldWithDefaults() *UserDefinedField`

NewUserDefinedFieldWithDefaults instantiates a new UserDefinedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UserDefinedField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserDefinedField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserDefinedField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *UserDefinedField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedField) SetName(v string)`

SetName sets Name field to given value.


### GetExample

`func (o *UserDefinedField) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *UserDefinedField) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *UserDefinedField) SetExample(v string)`

SetExample sets Example field to given value.


### GetOneOf

`func (o *UserDefinedField) GetOneOf() string`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *UserDefinedField) GetOneOfOk() (*string, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *UserDefinedField) SetOneOf(v string)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *UserDefinedField) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### GetManyOf

`func (o *UserDefinedField) GetManyOf() string`

GetManyOf returns the ManyOf field if non-nil, zero value otherwise.

### GetManyOfOk

`func (o *UserDefinedField) GetManyOfOk() (*string, bool)`

GetManyOfOk returns a tuple with the ManyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManyOf

`func (o *UserDefinedField) SetManyOf(v string)`

SetManyOf sets ManyOf field to given value.

### HasManyOf

`func (o *UserDefinedField) HasManyOf() bool`

HasManyOf returns a boolean if a field has been set.

### GetDefault

`func (o *UserDefinedField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UserDefinedField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UserDefinedField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UserDefinedField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


