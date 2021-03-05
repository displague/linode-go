# InlineObject16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The Region to deploy this Volume in. This is only required if a linode_id is not given.  | [optional] 
**LinodeId** | Pointer to **int32** | The Linode this volume should be attached to upon creation. If not given, the volume will be created without an attachment.  | [optional] 
**Size** | Pointer to **int32** | The initial size of this volume, in GB.  Be aware that volumes may only be resized up after creation.  | [optional] [default to 20]
**Label** | Pointer to **string** | The Volume&#39;s label, which is also used in the &#x60;filesystem_path&#x60; of the resulting volume.  | [optional] 
**ConfigId** | Pointer to **int32** | When creating a Volume attached to a Linode, the ID of the Linode Config to include the new Volume in. This Config must belong to the Linode referenced by &#x60;linode_id&#x60;. Must _not_ be provided if &#x60;linode_id&#x60; is not sent. If a &#x60;linode_id&#x60; is sent without a &#x60;config_id&#x60;, the volume will be attached:    * to the Linode&#39;s only config if it only has one config.   * to the Linode&#39;s last used config, if possible.  If no config can be selected for attachment, an error will be returned.  | [optional] 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only.  | [optional] 

## Methods

### NewInlineObject16

`func NewInlineObject16() *InlineObject16`

NewInlineObject16 instantiates a new InlineObject16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16WithDefaults

`func NewInlineObject16WithDefaults() *InlineObject16`

NewInlineObject16WithDefaults instantiates a new InlineObject16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *InlineObject16) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineObject16) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineObject16) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineObject16) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLinodeId

`func (o *InlineObject16) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *InlineObject16) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *InlineObject16) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *InlineObject16) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetSize

`func (o *InlineObject16) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InlineObject16) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InlineObject16) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *InlineObject16) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLabel

`func (o *InlineObject16) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject16) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject16) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InlineObject16) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetConfigId

`func (o *InlineObject16) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *InlineObject16) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *InlineObject16) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *InlineObject16) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject16) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject16) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject16) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject16) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


