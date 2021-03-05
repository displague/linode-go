# LKECluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Kubernetes cluster&#39;s unique ID. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Kubernetes cluster was created. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Kubernetes cluster was updated. | [optional] [readonly] 
**Label** | Pointer to **string** | This Kubernetes cluster&#39;s unique label for display purposes only. Labels have the following constraints:    * UTF-8 characters will be returned by the API using escape     sequences of their Unicode code points. For example, the     Japanese character *か* is 3 bytes in UTF-8 (&#x60;0xE382AB&#x60;). Its     Unicode code point is 2 bytes (&#x60;0x30AB&#x60;). APIv4 supports this     character and the API will return it as the escape sequence     using six 1 byte characters which represent 2 bytes of Unicode     code point (&#x60;\&quot;\\u30ab\&quot;&#x60;).   * 4 byte UTF-8 characters are not supported.   * If the label is entirely composed of UTF-8 characters, the API     response will return the code points using up to 193 1 byte     characters.  | [optional] 
**Region** | Pointer to **string** | This Kubernetes cluster&#39;s location. | [optional] 
**K8sVersion** | Pointer to **string** | The desired Kubernetes version for this Kubernetes cluster in the format of &amp;lt;major&amp;gt;.&amp;lt;minor&amp;gt;, and the latest supported patch version will be deployed.  | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to the Kubernetes cluster. Tags are for organizational purposes only.  | [optional] 

## Methods

### NewLKECluster

`func NewLKECluster() *LKECluster`

NewLKECluster instantiates a new LKECluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLKEClusterWithDefaults

`func NewLKEClusterWithDefaults() *LKECluster`

NewLKEClusterWithDefaults instantiates a new LKECluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LKECluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LKECluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LKECluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LKECluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *LKECluster) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LKECluster) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LKECluster) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LKECluster) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *LKECluster) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LKECluster) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LKECluster) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *LKECluster) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetLabel

`func (o *LKECluster) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LKECluster) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LKECluster) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LKECluster) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *LKECluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LKECluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LKECluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LKECluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetK8sVersion

`func (o *LKECluster) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *LKECluster) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *LKECluster) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *LKECluster) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetTags

`func (o *LKECluster) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LKECluster) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LKECluster) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LKECluster) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


