# InlineObject6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | This is the Region where the Linode will be deployed. To view all available Regions you can deploy to see [/regions](/docs/api/regions/#regions-list). * Region can only be provided when cloning to a new Linode.  | 
**Type** | **string** | A Linode&#39;s Type determines what resources are available to it, including disk space, memory, and virtual cpus. The amounts available to a specific Linode are returned as &#x60;specs&#x60; on the Linode object.  To view all available Linode Types you can deploy with see [/linode/types](/docs/api/linode-types/#types-list).  * Type can only be provided when cloning to a new Linode.  | 
**LinodeId** | Pointer to **int32** | If an existing Linode is to be the target for the clone, the ID of that Linode. The existing Linode must have enough resources to accept the clone.  | [optional] 
**Label** | Pointer to **string** | The label to assign this Linode when cloning to a new Linode. * Can only be provided when cloning to a new Linode. * Defaults to \&quot;linode\&quot;.  | [optional] 
**Group** | Pointer to **string** | A label used to group Linodes for display. Linodes are not required to have a group.  | [optional] 
**BackupsEnabled** | Pointer to **bool** | If this field is set to &#x60;true&#x60;, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. Pricing is included in the response from [/linodes/types](/docs/api/linode-types/#types-list).  * Can only be included when cloning to a new Linode.  | [optional] 
**Disks** | Pointer to **[]int32** | An array of disk IDs. * If the &#x60;disks&#x60; parameter **is not provided**, then **no extra disks will be cloned** from the source Linode. All disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter will still be cloned. * **If an empty array is provided** for the &#x60;disks&#x60; parameter, then **no extra disks will be cloned** from the source Linode. All disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter will still be cloned. * **If a non-empty array is provided** for the &#x60;disks&#x60; parameter, then **the disks specified in the array will be cloned** from the source Linode, in addition to any disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter.  | [optional] 
**Configs** | Pointer to **[]int32** | An array of configuration profile IDs. * If the &#x60;configs&#x60; parameter **is not provided**, then **all configuration profiles and their associated disks will be cloned** from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will also be cloned. * **If an empty array is provided** for the &#x60;configs&#x60; parameter, then **no configuration profiles (nor their associated disks) will be cloned** from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will still be cloned. * **If a non-empty array is provided** for the &#x60;configs&#x60; parameter, then **the configuration profiles specified in the array (and their associated disks) will be cloned** from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will also be cloned.  | [optional] 

## Methods

### NewInlineObject6

`func NewInlineObject6(region string, type_ string, ) *InlineObject6`

NewInlineObject6 instantiates a new InlineObject6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject6WithDefaults

`func NewInlineObject6WithDefaults() *InlineObject6`

NewInlineObject6WithDefaults instantiates a new InlineObject6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *InlineObject6) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineObject6) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineObject6) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetType

`func (o *InlineObject6) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject6) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject6) SetType(v string)`

SetType sets Type field to given value.


### GetLinodeId

`func (o *InlineObject6) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *InlineObject6) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *InlineObject6) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *InlineObject6) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetLabel

`func (o *InlineObject6) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject6) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject6) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InlineObject6) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroup

`func (o *InlineObject6) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineObject6) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineObject6) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineObject6) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBackupsEnabled

`func (o *InlineObject6) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *InlineObject6) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *InlineObject6) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *InlineObject6) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetDisks

`func (o *InlineObject6) GetDisks() []int32`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *InlineObject6) GetDisksOk() (*[]int32, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *InlineObject6) SetDisks(v []int32)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *InlineObject6) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetConfigs

`func (o *InlineObject6) GetConfigs() []int32`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *InlineObject6) GetConfigsOk() (*[]int32, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *InlineObject6) SetConfigs(v []int32)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *InlineObject6) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


