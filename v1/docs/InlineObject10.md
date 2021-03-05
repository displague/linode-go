# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | [**Label**](Label.md) |  | 
**Region** | [**Region**](Region.md) |  | 
**K8sVersion** | [**K8sVersion**](K8sVersion.md) |  | 
**Tags** | Pointer to [**Tags**](Tags.md) |  | [optional] 
**NodePools** | [**[]LKENodePoolRequestBody**](LKENodePoolRequestBody.md) |  | 

## Methods

### NewInlineObject10

`func NewInlineObject10(label Label, region Region, k8sVersion K8sVersion, nodePools []LKENodePoolRequestBody, ) *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InlineObject10) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject10) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject10) SetLabel(v Label)`

SetLabel sets Label field to given value.


### GetRegion

`func (o *InlineObject10) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineObject10) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineObject10) SetRegion(v Region)`

SetRegion sets Region field to given value.


### GetK8sVersion

`func (o *InlineObject10) GetK8sVersion() K8sVersion`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *InlineObject10) GetK8sVersionOk() (*K8sVersion, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *InlineObject10) SetK8sVersion(v K8sVersion)`

SetK8sVersion sets K8sVersion field to given value.


### GetTags

`func (o *InlineObject10) GetTags() Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject10) GetTagsOk() (*Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject10) SetTags(v Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject10) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNodePools

`func (o *InlineObject10) GetNodePools() []LKENodePoolRequestBody`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *InlineObject10) GetNodePoolsOk() (*[]LKENodePoolRequestBody, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *InlineObject10) SetNodePools(v []LKENodePoolRequestBody)`

SetNodePools sets NodePools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


