# InlineObject11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | [**Label**](Label.md) |  | 
**Rules** | [**Rules**](Rules.md) |  | 
**Tags** | Pointer to [**Tags**](Tags.md) |  | [optional] 
**Devices** | Pointer to [**NetworkingFirewallsDevices**](NetworkingFirewallsDevices.md) |  | [optional] 

## Methods

### NewInlineObject11

`func NewInlineObject11(label Label, rules Rules, ) *InlineObject11`

NewInlineObject11 instantiates a new InlineObject11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11WithDefaults

`func NewInlineObject11WithDefaults() *InlineObject11`

NewInlineObject11WithDefaults instantiates a new InlineObject11 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InlineObject11) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject11) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject11) SetLabel(v Label)`

SetLabel sets Label field to given value.


### GetRules

`func (o *InlineObject11) GetRules() Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject11) GetRulesOk() (*Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject11) SetRules(v Rules)`

SetRules sets Rules field to given value.


### GetTags

`func (o *InlineObject11) GetTags() Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject11) GetTagsOk() (*Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject11) SetTags(v Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject11) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDevices

`func (o *InlineObject11) GetDevices() NetworkingFirewallsDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject11) GetDevicesOk() (*NetworkingFirewallsDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject11) SetDevices(v NetworkingFirewallsDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineObject11) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


