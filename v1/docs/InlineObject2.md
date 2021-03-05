# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | Pointer to **int32** | A Backup ID from another Linode&#39;s available backups. Your User must have &#x60;read_write&#x60; access to that Linode, the Backup must have a &#x60;status&#x60; of &#x60;successful&#x60;, and the Linode must be deployed to the same &#x60;region&#x60; as the Backup. See [/linode/instances/{linodeId}/backups](/docs/api/linode-instances/#backups-list) for a Linode&#39;s available backups.  This field and the &#x60;image&#x60; field are mutually exclusive.  | [optional] 
**BackupsEnabled** | Pointer to **bool** | If this field is set to &#x60;true&#x60;, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.  This option is always treated as &#x60;true&#x60; if the account-wide &#x60;backups_enabled&#x60; setting is &#x60;true&#x60;.  See [account settings](/docs/api/account/#account-settings-view) for more information.  Backup pricing is included in the response from [/linodes/types](/docs/api/linode-types/#types-list)  | [optional] 
**SwapSize** | Pointer to **int32** | When deploying from an Image, this field is optional, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.  | [optional] [default to 512]
**Type** | **string** | The [Linode Type](/docs/api/linode-types/#types-list) of the Linode you are creating.  | 
**Region** | **string** | The [Region](/docs/api/regions/#regions-list) where the Linode will be located.  | 
**Image** | Pointer to **string** | An Image ID to deploy the Disk from. Official Linode Images start with &#x60;linode/ &#x60;, while your Images start with &#x60;private/&#x60;. See [/images](/docs/api/images/) for more information on the Images available for you to use.  | [optional] 
**RootPass** | Pointer to [**RootPass**](RootPass.md) |  | [optional] 
**AuthorizedKeys** | Pointer to **[]string** | A list of SSH public keys to deploy for the root user on the newly-created Linode.  Only accepted if &#x60;image&#x60; is provided.  | [optional] 
**StackscriptId** | Pointer to **int32** | The StackScript to deploy to the newly-created Linode.  If provided, \&quot;image\&quot; must also be provided, and must be an Image that is compatible with this StackScript.  | [optional] 
**StackscriptData** | Pointer to **map[string]interface{}** | An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if &#x60;stackscript_id&#x60; is given.  The required values depend on the StackScript being deployed.  | [optional] 
**Booted** | Pointer to **bool** | Whether to boot this Linode after the deploy is complete. Defaults to true if &#x60;image&#x60; is provided. Not accepted if not deploying from an Image.  | [optional] 
**Label** | Pointer to [**Label**](Label.md) |  | [optional] 
**Tags** | Pointer to [**Tags**](Tags.md) |  | [optional] 
**Group** | Pointer to [**Group**](Group.md) |  | [optional] 
**PrivateIp** | Pointer to **bool** | If true, the created Linode will have private networking enabled.  | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users &#x60;~/.ssh/authorized_keys&#x60; file automatically.  | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(type_ string, region string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *InlineObject2) GetBackupId() int32`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *InlineObject2) GetBackupIdOk() (*int32, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *InlineObject2) SetBackupId(v int32)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *InlineObject2) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackupsEnabled

`func (o *InlineObject2) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *InlineObject2) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *InlineObject2) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *InlineObject2) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetSwapSize

`func (o *InlineObject2) GetSwapSize() int32`

GetSwapSize returns the SwapSize field if non-nil, zero value otherwise.

### GetSwapSizeOk

`func (o *InlineObject2) GetSwapSizeOk() (*int32, bool)`

GetSwapSizeOk returns a tuple with the SwapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapSize

`func (o *InlineObject2) SetSwapSize(v int32)`

SetSwapSize sets SwapSize field to given value.

### HasSwapSize

`func (o *InlineObject2) HasSwapSize() bool`

HasSwapSize returns a boolean if a field has been set.

### GetType

`func (o *InlineObject2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject2) SetType(v string)`

SetType sets Type field to given value.


### GetRegion

`func (o *InlineObject2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineObject2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineObject2) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetImage

`func (o *InlineObject2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InlineObject2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InlineObject2) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *InlineObject2) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRootPass

`func (o *InlineObject2) GetRootPass() RootPass`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *InlineObject2) GetRootPassOk() (*RootPass, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *InlineObject2) SetRootPass(v RootPass)`

SetRootPass sets RootPass field to given value.

### HasRootPass

`func (o *InlineObject2) HasRootPass() bool`

HasRootPass returns a boolean if a field has been set.

### GetAuthorizedKeys

`func (o *InlineObject2) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *InlineObject2) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *InlineObject2) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *InlineObject2) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetStackscriptId

`func (o *InlineObject2) GetStackscriptId() int32`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *InlineObject2) GetStackscriptIdOk() (*int32, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *InlineObject2) SetStackscriptId(v int32)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *InlineObject2) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.

### GetStackscriptData

`func (o *InlineObject2) GetStackscriptData() map[string]interface{}`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *InlineObject2) GetStackscriptDataOk() (*map[string]interface{}, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *InlineObject2) SetStackscriptData(v map[string]interface{})`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *InlineObject2) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.

### GetBooted

`func (o *InlineObject2) GetBooted() bool`

GetBooted returns the Booted field if non-nil, zero value otherwise.

### GetBootedOk

`func (o *InlineObject2) GetBootedOk() (*bool, bool)`

GetBootedOk returns a tuple with the Booted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooted

`func (o *InlineObject2) SetBooted(v bool)`

SetBooted sets Booted field to given value.

### HasBooted

`func (o *InlineObject2) HasBooted() bool`

HasBooted returns a boolean if a field has been set.

### GetLabel

`func (o *InlineObject2) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject2) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject2) SetLabel(v Label)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InlineObject2) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject2) GetTags() Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject2) GetTagsOk() (*Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject2) SetTags(v Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGroup

`func (o *InlineObject2) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineObject2) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineObject2) SetGroup(v Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineObject2) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPrivateIp

`func (o *InlineObject2) GetPrivateIp() bool`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *InlineObject2) GetPrivateIpOk() (*bool, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *InlineObject2) SetPrivateIp(v bool)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *InlineObject2) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *InlineObject2) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *InlineObject2) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *InlineObject2) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *InlineObject2) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


