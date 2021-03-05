# LinodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**RootPass** | Pointer to [**RootPass**](RootPass.md) |  | [optional] 
**AuthorizedKeys** | Pointer to [**AuthorizedKeys**](AuthorizedKeys.md) |  | [optional] 
**AuthorizedUsers** | Pointer to [**AuthorizedUsers**](AuthorizedUsers.md) |  | [optional] 
**StackscriptId** | Pointer to [**StackscriptId**](StackscriptId.md) |  | [optional] 
**StackscriptData** | Pointer to [**StackscriptData**](StackscriptData.md) |  | [optional] 
**Booted** | Pointer to **bool** | This field defaults to &#x60;true&#x60; if the Linode is created with an Image or from a Backup. If it is deployed from an Image or a Backup and you wish it to remain &#x60;offline&#x60; after deployment, set this to &#x60;false&#x60;.  | [optional] 

## Methods

### NewLinodeRequest

`func NewLinodeRequest() *LinodeRequest`

NewLinodeRequest instantiates a new LinodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeRequestWithDefaults

`func NewLinodeRequestWithDefaults() *LinodeRequest`

NewLinodeRequestWithDefaults instantiates a new LinodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *LinodeRequest) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *LinodeRequest) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *LinodeRequest) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *LinodeRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRootPass

`func (o *LinodeRequest) GetRootPass() RootPass`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *LinodeRequest) GetRootPassOk() (*RootPass, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *LinodeRequest) SetRootPass(v RootPass)`

SetRootPass sets RootPass field to given value.

### HasRootPass

`func (o *LinodeRequest) HasRootPass() bool`

HasRootPass returns a boolean if a field has been set.

### GetAuthorizedKeys

`func (o *LinodeRequest) GetAuthorizedKeys() AuthorizedKeys`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *LinodeRequest) GetAuthorizedKeysOk() (*AuthorizedKeys, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *LinodeRequest) SetAuthorizedKeys(v AuthorizedKeys)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *LinodeRequest) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *LinodeRequest) GetAuthorizedUsers() AuthorizedUsers`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *LinodeRequest) GetAuthorizedUsersOk() (*AuthorizedUsers, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *LinodeRequest) SetAuthorizedUsers(v AuthorizedUsers)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *LinodeRequest) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### GetStackscriptId

`func (o *LinodeRequest) GetStackscriptId() StackscriptId`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *LinodeRequest) GetStackscriptIdOk() (*StackscriptId, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *LinodeRequest) SetStackscriptId(v StackscriptId)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *LinodeRequest) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.

### GetStackscriptData

`func (o *LinodeRequest) GetStackscriptData() StackscriptData`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *LinodeRequest) GetStackscriptDataOk() (*StackscriptData, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *LinodeRequest) SetStackscriptData(v StackscriptData)`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *LinodeRequest) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.

### GetBooted

`func (o *LinodeRequest) GetBooted() bool`

GetBooted returns the Booted field if non-nil, zero value otherwise.

### GetBootedOk

`func (o *LinodeRequest) GetBootedOk() (*bool, bool)`

GetBootedOk returns a tuple with the Booted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooted

`func (o *LinodeRequest) SetBooted(v bool)`

SetBooted sets Booted field to given value.

### HasBooted

`func (o *LinodeRequest) HasBooted() bool`

HasBooted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


