# StackScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of this StackScript. | [optional] [readonly] 
**Username** | Pointer to **string** | The User who created the StackScript.  | [optional] [readonly] 
**UserGravatarId** | Pointer to **string** | The Gravatar ID for the User who created the StackScript.  | [optional] [readonly] 
**Label** | Pointer to **string** | The StackScript&#39;s label is for display purposes only.  | [optional] 
**Description** | Pointer to **string** | A description for the StackScript.  | [optional] 
**Images** | Pointer to **[]string** | An array of Image IDs. These are the images that can be deployed with this Stackscript.  | [optional] 
**DeploymentsTotal** | Pointer to **int32** | The total number of times this StackScript has been deployed.  | [optional] [readonly] 
**DeploymentsActive** | Pointer to **int32** | Count of currently active, deployed Linodes created from this StackScript.  | [optional] [readonly] 
**IsPublic** | Pointer to **bool** | This determines whether other users can use your StackScript. **Once a StackScript is made public, it cannot be made private.**  | [optional] 
**Created** | Pointer to **time.Time** | The date this StackScript was created.  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The date this StackScript was last updated.  | [optional] [readonly] 
**RevNote** | Pointer to **string** | This field allows you to add notes for the set of revisions made to this StackScript.  | [optional] 
**Script** | Pointer to **string** | The script to execute when provisioning a new Linode with this StackScript.  | [optional] 
**UserDefinedFields** | Pointer to [**[]UserDefinedField**](UserDefinedField.md) | This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment. See &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;/docs/platform/stackscripts/#variables-and-udfs\&quot;&gt;Variables and UDFs&lt;/a&gt; for more information.  | [optional] [readonly] 

## Methods

### NewStackScript

`func NewStackScript() *StackScript`

NewStackScript instantiates a new StackScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackScriptWithDefaults

`func NewStackScriptWithDefaults() *StackScript`

NewStackScriptWithDefaults instantiates a new StackScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackScript) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackScript) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackScript) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StackScript) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *StackScript) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StackScript) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StackScript) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StackScript) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUserGravatarId

`func (o *StackScript) GetUserGravatarId() string`

GetUserGravatarId returns the UserGravatarId field if non-nil, zero value otherwise.

### GetUserGravatarIdOk

`func (o *StackScript) GetUserGravatarIdOk() (*string, bool)`

GetUserGravatarIdOk returns a tuple with the UserGravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGravatarId

`func (o *StackScript) SetUserGravatarId(v string)`

SetUserGravatarId sets UserGravatarId field to given value.

### HasUserGravatarId

`func (o *StackScript) HasUserGravatarId() bool`

HasUserGravatarId returns a boolean if a field has been set.

### GetLabel

`func (o *StackScript) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StackScript) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StackScript) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StackScript) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *StackScript) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StackScript) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StackScript) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StackScript) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImages

`func (o *StackScript) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *StackScript) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *StackScript) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *StackScript) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetDeploymentsTotal

`func (o *StackScript) GetDeploymentsTotal() int32`

GetDeploymentsTotal returns the DeploymentsTotal field if non-nil, zero value otherwise.

### GetDeploymentsTotalOk

`func (o *StackScript) GetDeploymentsTotalOk() (*int32, bool)`

GetDeploymentsTotalOk returns a tuple with the DeploymentsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsTotal

`func (o *StackScript) SetDeploymentsTotal(v int32)`

SetDeploymentsTotal sets DeploymentsTotal field to given value.

### HasDeploymentsTotal

`func (o *StackScript) HasDeploymentsTotal() bool`

HasDeploymentsTotal returns a boolean if a field has been set.

### GetDeploymentsActive

`func (o *StackScript) GetDeploymentsActive() int32`

GetDeploymentsActive returns the DeploymentsActive field if non-nil, zero value otherwise.

### GetDeploymentsActiveOk

`func (o *StackScript) GetDeploymentsActiveOk() (*int32, bool)`

GetDeploymentsActiveOk returns a tuple with the DeploymentsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsActive

`func (o *StackScript) SetDeploymentsActive(v int32)`

SetDeploymentsActive sets DeploymentsActive field to given value.

### HasDeploymentsActive

`func (o *StackScript) HasDeploymentsActive() bool`

HasDeploymentsActive returns a boolean if a field has been set.

### GetIsPublic

`func (o *StackScript) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *StackScript) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *StackScript) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *StackScript) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetCreated

`func (o *StackScript) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StackScript) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StackScript) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StackScript) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *StackScript) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StackScript) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *StackScript) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *StackScript) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetRevNote

`func (o *StackScript) GetRevNote() string`

GetRevNote returns the RevNote field if non-nil, zero value otherwise.

### GetRevNoteOk

`func (o *StackScript) GetRevNoteOk() (*string, bool)`

GetRevNoteOk returns a tuple with the RevNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevNote

`func (o *StackScript) SetRevNote(v string)`

SetRevNote sets RevNote field to given value.

### HasRevNote

`func (o *StackScript) HasRevNote() bool`

HasRevNote returns a boolean if a field has been set.

### GetScript

`func (o *StackScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *StackScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *StackScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *StackScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *StackScript) GetUserDefinedFields() []UserDefinedField`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *StackScript) GetUserDefinedFieldsOk() (*[]UserDefinedField, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *StackScript) SetUserDefinedFields(v []UserDefinedField)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *StackScript) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


