# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The ID representing the Linode Type. | 
**AllowAutoDiskResize** | Pointer to **bool** | Automatically resize disks when resizing a Linode. When resizing down to a smaller plan your Linode&#39;s data must fit within the smaller disk size.  | [optional] [default to true]

## Methods

### NewInlineObject9

`func NewInlineObject9(type_ string, ) *InlineObject9`

NewInlineObject9 instantiates a new InlineObject9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject9WithDefaults

`func NewInlineObject9WithDefaults() *InlineObject9`

NewInlineObject9WithDefaults instantiates a new InlineObject9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject9) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject9) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject9) SetType(v string)`

SetType sets Type field to given value.


### GetAllowAutoDiskResize

`func (o *InlineObject9) GetAllowAutoDiskResize() bool`

GetAllowAutoDiskResize returns the AllowAutoDiskResize field if non-nil, zero value otherwise.

### GetAllowAutoDiskResizeOk

`func (o *InlineObject9) GetAllowAutoDiskResizeOk() (*bool, bool)`

GetAllowAutoDiskResizeOk returns a tuple with the AllowAutoDiskResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoDiskResize

`func (o *InlineObject9) SetAllowAutoDiskResize(v bool)`

SetAllowAutoDiskResize sets AllowAutoDiskResize field to given value.

### HasAllowAutoDiskResize

`func (o *InlineObject9) HasAllowAutoDiskResize() bool`

HasAllowAutoDiskResize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


