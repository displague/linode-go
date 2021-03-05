# DiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**Label** | [**Label**](Label.md) |  | 
**Filesystem** | Pointer to [**Filesystem**](Filesystem.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** | If true, this Disk is read-only.  | [optional] 
**Image** | Pointer to **string** | An Image ID to deploy the Disk from. Official Linode Images start with &#x60;linode/ &#x60;, while your Images start with &#x60;private/&#x60;. See [/images](/docs/api/images/#images-list) for more information on the Images available for you to use.  | [optional] 
**AuthorizedKeys** | Pointer to **[]string** | A list of public SSH keys that will be automatically appended to the root user&#39;s &#x60;~/.ssh/authorized_keys&#x60; file.  | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | A list of usernames that will have their SSH keys, if any, automatically appended to the root user&#39;s &#x60;~/.ssh/authorized_keys&#x60; file.  | [optional] 
**RootPass** | Pointer to **string** | This will set the root user&#39;s password on the newly-created Linode. Linode passwords have the following constraints:    * Must meet a password strength score requirement that is calculated internally by the API.     If the strength requirement is not met, you will receive a &#x60;Password does not meet strength requirement&#x60; error.  | [optional] 
**StackscriptId** | Pointer to **int32** | A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible &#x60;image&#x60; is required to use a StackScript. To get a list of available StackScript and their permitted Images see [/stackscripts](/docs/api/stackscripts/#stackscripts-list). This field cannot be used when deploying from a Backup or a private Image.  | [optional] 
**StackscriptData** | Pointer to **map[string]interface{}** | This field is required only if the StackScript being deployed requires input data from the User for successful completion. See &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;/docs/platform/stackscripts/#variables-and-udfs\&quot;&gt;Variables and UDFs&lt;/a&gt; for more details. This field is required to be valid JSON.  | [optional] 

## Methods

### NewDiskRequest

`func NewDiskRequest(size int32, label Label, ) *DiskRequest`

NewDiskRequest instantiates a new DiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRequestWithDefaults

`func NewDiskRequestWithDefaults() *DiskRequest`

NewDiskRequestWithDefaults instantiates a new DiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *DiskRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DiskRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DiskRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLabel

`func (o *DiskRequest) GetLabel() Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DiskRequest) GetLabelOk() (*Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DiskRequest) SetLabel(v Label)`

SetLabel sets Label field to given value.


### GetFilesystem

`func (o *DiskRequest) GetFilesystem() Filesystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *DiskRequest) GetFilesystemOk() (*Filesystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *DiskRequest) SetFilesystem(v Filesystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *DiskRequest) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetReadOnly

`func (o *DiskRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DiskRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DiskRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *DiskRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetImage

`func (o *DiskRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DiskRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DiskRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DiskRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetAuthorizedKeys

`func (o *DiskRequest) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *DiskRequest) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *DiskRequest) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *DiskRequest) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *DiskRequest) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *DiskRequest) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *DiskRequest) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *DiskRequest) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### GetRootPass

`func (o *DiskRequest) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *DiskRequest) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *DiskRequest) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.

### HasRootPass

`func (o *DiskRequest) HasRootPass() bool`

HasRootPass returns a boolean if a field has been set.

### GetStackscriptId

`func (o *DiskRequest) GetStackscriptId() int32`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *DiskRequest) GetStackscriptIdOk() (*int32, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *DiskRequest) SetStackscriptId(v int32)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *DiskRequest) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.

### GetStackscriptData

`func (o *DiskRequest) GetStackscriptData() map[string]interface{}`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *DiskRequest) GetStackscriptDataOk() (*map[string]interface{}, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *DiskRequest) SetStackscriptData(v map[string]interface{})`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *DiskRequest) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


