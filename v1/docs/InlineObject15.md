# InlineObject15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The new Tag.  | 
**Linodes** | Pointer to **[]int32** | A list of Linode IDs to apply the new Tag to.  You must be allowed to &#x60;read_write&#x60; all of the requested Linodes, or the Tag will not be created and an error will be returned.  | [optional] 
**Domains** | Pointer to **[]int32** | A list of Domain IDs to apply the new Tag to.  You must be allowed to &#x60;read_write&#x60; all of the requested Domains, or the Tag will not be created and an error will be returned.  | [optional] 
**Volumes** | Pointer to **[]int32** | A list of Volume IDs to apply the new Tag to.  You must be allowed to &#x60;read_write&#x60; all of the requested Volumes, or the Tag will not be created and an error will be returned.  | [optional] 
**Nodebalancers** | Pointer to **[]int32** | A list of NodeBalancer IDs to apply the new Tag to. You must be allowed to &#x60;read_write&#x60; all of the requested NodeBalancers, or the Tag will not be created and an error will be returned.  | [optional] 

## Methods

### NewInlineObject15

`func NewInlineObject15(label string, ) *InlineObject15`

NewInlineObject15 instantiates a new InlineObject15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject15WithDefaults

`func NewInlineObject15WithDefaults() *InlineObject15`

NewInlineObject15WithDefaults instantiates a new InlineObject15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InlineObject15) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject15) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject15) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLinodes

`func (o *InlineObject15) GetLinodes() []int32`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *InlineObject15) GetLinodesOk() (*[]int32, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *InlineObject15) SetLinodes(v []int32)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *InlineObject15) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.

### GetDomains

`func (o *InlineObject15) GetDomains() []int32`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InlineObject15) GetDomainsOk() (*[]int32, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InlineObject15) SetDomains(v []int32)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *InlineObject15) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetVolumes

`func (o *InlineObject15) GetVolumes() []int32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InlineObject15) GetVolumesOk() (*[]int32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InlineObject15) SetVolumes(v []int32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InlineObject15) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNodebalancers

`func (o *InlineObject15) GetNodebalancers() []int32`

GetNodebalancers returns the Nodebalancers field if non-nil, zero value otherwise.

### GetNodebalancersOk

`func (o *InlineObject15) GetNodebalancersOk() (*[]int32, bool)`

GetNodebalancersOk returns a tuple with the Nodebalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancers

`func (o *InlineObject15) SetNodebalancers(v []int32)`

SetNodebalancers sets Nodebalancers field to given value.

### HasNodebalancers

`func (o *InlineObject15) HasNodebalancers() bool`

HasNodebalancers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


