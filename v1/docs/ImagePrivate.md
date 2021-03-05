# ImagePrivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of this Image. | [optional] [readonly] 
**Label** | Pointer to **string** | A short description of the Image. Labels cannot contain special characters.  | [optional] 
**Created** | Pointer to **time.Time** | When this Image was created. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Image was last updated. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The name of the User who created this Image.  | [optional] [readonly] 
**Deprecated** | Pointer to **bool** | Whether or not this Image is deprecated. Will only be True for deprecated public Images.  | [optional] [readonly] 
**Description** | Pointer to **string** | A detailed description of this Image. | [optional] 
**IsPublic** | Pointer to **bool** | True if the Image is public. | [optional] [readonly] 
**Size** | Pointer to **int32** | The minimum size this Image needs to deploy. Size is in MB.  | [optional] [readonly] 
**Type** | Pointer to **string** | How the Image was created. \&quot;Manual\&quot; Images can be created at any time. \&quot;Automatic\&quot; images are created automatically from a deleted Linode.  | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | Only Images created automatically (from a deleted Linode; type&#x3D;automatic) will expire.  | [optional] [readonly] 
**Eol** | Pointer to **time.Time** | The date of the image&#39;s planned end of life. Some images, like custom private images, will not have an end of life date. In that case this field will be &#x60;None&#x60;.  | [optional] [readonly] 
**Vendor** | Pointer to **NullableString** | The upstream distribution vendor. &#x60;None&#x60; for private Images.  | [optional] [readonly] 

## Methods

### NewImagePrivate

`func NewImagePrivate() *ImagePrivate`

NewImagePrivate instantiates a new ImagePrivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePrivateWithDefaults

`func NewImagePrivateWithDefaults() *ImagePrivate`

NewImagePrivateWithDefaults instantiates a new ImagePrivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImagePrivate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImagePrivate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImagePrivate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImagePrivate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ImagePrivate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ImagePrivate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ImagePrivate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ImagePrivate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreated

`func (o *ImagePrivate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImagePrivate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImagePrivate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ImagePrivate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ImagePrivate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ImagePrivate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ImagePrivate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ImagePrivate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ImagePrivate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ImagePrivate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ImagePrivate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ImagePrivate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeprecated

`func (o *ImagePrivate) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ImagePrivate) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ImagePrivate) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ImagePrivate) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *ImagePrivate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImagePrivate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImagePrivate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImagePrivate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *ImagePrivate) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ImagePrivate) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ImagePrivate) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ImagePrivate) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSize

`func (o *ImagePrivate) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImagePrivate) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImagePrivate) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImagePrivate) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ImagePrivate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImagePrivate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImagePrivate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImagePrivate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpiry

`func (o *ImagePrivate) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ImagePrivate) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ImagePrivate) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ImagePrivate) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetEol

`func (o *ImagePrivate) GetEol() time.Time`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *ImagePrivate) GetEolOk() (*time.Time, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *ImagePrivate) SetEol(v time.Time)`

SetEol sets Eol field to given value.

### HasEol

`func (o *ImagePrivate) HasEol() bool`

HasEol returns a boolean if a field has been set.

### GetVendor

`func (o *ImagePrivate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ImagePrivate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ImagePrivate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ImagePrivate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *ImagePrivate) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *ImagePrivate) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


