# Linode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **interface{}** | The Linode&#39;s label is for display purposes only. If no label is provided for a Linode, a default will be assigned. Linode labels have the following constraints:    * Must start with an alpha character.   * May only consist of alphanumeric characters, dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;).   * Cannot have two dashes (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row.  | [optional] 
**Region** | Pointer to **interface{}** | This is the location where the Linode was deployed. A Linode&#39;s region can only be changed by initiating a [cross data center migration](/docs/api/linode-instances/#dc-migrationpending-host-migration-initiate).  | [optional] [readonly] 
**Image** | Pointer to [**NullableImage**](image.md) |  | [optional] [readonly] 
**Type** | Pointer to **interface{}** | This is the [Linode Type](/docs/api/linode-types/#types-list) that this Linode was deployed with. To change a Linode&#39;s Type, use [POST /linode/instances/{linodeId}/resize](/docs/api/linode-instances/#linode-resize).  | [optional] [readonly] 
**Group** | Pointer to **string** | A deprecated property denoting a group label for this Linode.  | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object.  Tags are for organizational purposes only.  | [optional] 
**Id** | Pointer to **int32** | This Linode&#39;s ID which must be provided for all operations impacting this Linode.  | [optional] [readonly] 
**Status** | Pointer to **string** | A brief description of this Linode&#39;s current state. This field may change without direct action from you. For example, when a Linode goes into maintenance mode its status will display \&quot;stopped\&quot;.  | [optional] [readonly] 
**Hypervisor** | Pointer to **string** | The virtualization software powering this Linode.  | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Linode was created. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Linode was last updated. | [optional] [readonly] 
**Ipv4** | Pointer to **[]string** | This Linode&#39;s IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to [open a support ticket](/docs/api/support/#support-ticket-open) to get additional IPv4 addresses.  IPv4 addresses may be reassigned between your Linodes, or shared with other Linodes. See the [/networking](/docs/api/networking/) endpoints for details.  | [optional] [readonly] 
**Ipv6** | Pointer to **string** | This Linode&#39;s IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared. If the Linode has not been assigned an IPv6 address, the return value will be &#x60;null&#x60;.  | [optional] [readonly] 
**Specs** | Pointer to **map[string]interface{}** | Information about the resources available to this Linode. | [optional] [readonly] 
**Alerts** | Pointer to **map[string]interface{}** |  | [optional] 
**Backups** | Pointer to **map[string]interface{}** | Information about this Linode&#39;s backups status. For information about available backups, see [/linode/instances/{linodeId}/backups](/docs/api/linode-instances/#backups-list).  | [optional] 
**WatchdogEnabled** | Pointer to **bool** | The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.  | [optional] 

## Methods

### NewLinode

`func NewLinode() *Linode`

NewLinode instantiates a new Linode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeWithDefaults

`func NewLinodeWithDefaults() *Linode`

NewLinodeWithDefaults instantiates a new Linode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Linode) GetLabel() interface{}`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Linode) GetLabelOk() (*interface{}, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Linode) SetLabel(v interface{})`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Linode) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Linode) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Linode) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetRegion

`func (o *Linode) GetRegion() interface{}`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Linode) GetRegionOk() (*interface{}, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Linode) SetRegion(v interface{})`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Linode) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Linode) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Linode) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetImage

`func (o *Linode) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Linode) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Linode) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *Linode) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *Linode) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *Linode) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetType

`func (o *Linode) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Linode) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Linode) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *Linode) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Linode) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Linode) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetGroup

`func (o *Linode) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Linode) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Linode) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Linode) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTags

`func (o *Linode) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Linode) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Linode) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Linode) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *Linode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Linode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Linode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Linode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Linode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Linode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Linode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Linode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHypervisor

`func (o *Linode) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *Linode) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *Linode) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *Linode) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetCreated

`func (o *Linode) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Linode) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Linode) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Linode) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Linode) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Linode) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Linode) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Linode) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetIpv4

`func (o *Linode) GetIpv4() []string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Linode) GetIpv4Ok() (*[]string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Linode) SetIpv4(v []string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *Linode) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *Linode) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Linode) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Linode) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Linode) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetSpecs

`func (o *Linode) GetSpecs() map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Linode) GetSpecsOk() (*map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Linode) SetSpecs(v map[string]interface{})`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Linode) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetAlerts

`func (o *Linode) GetAlerts() map[string]interface{}`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Linode) GetAlertsOk() (*map[string]interface{}, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Linode) SetAlerts(v map[string]interface{})`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *Linode) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetBackups

`func (o *Linode) GetBackups() map[string]interface{}`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *Linode) GetBackupsOk() (*map[string]interface{}, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *Linode) SetBackups(v map[string]interface{})`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *Linode) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetWatchdogEnabled

`func (o *Linode) GetWatchdogEnabled() bool`

GetWatchdogEnabled returns the WatchdogEnabled field if non-nil, zero value otherwise.

### GetWatchdogEnabledOk

`func (o *Linode) GetWatchdogEnabledOk() (*bool, bool)`

GetWatchdogEnabledOk returns a tuple with the WatchdogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdogEnabled

`func (o *Linode) SetWatchdogEnabled(v bool)`

SetWatchdogEnabled sets WatchdogEnabled field to given value.

### HasWatchdogEnabled

`func (o *Linode) HasWatchdogEnabled() bool`

HasWatchdogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


