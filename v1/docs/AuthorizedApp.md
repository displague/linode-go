# AuthorizedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This authorization&#39;s ID, used for revoking access.  | [optional] [readonly] 
**Label** | Pointer to **string** | The name of the application you&#39;ve authorized.  | [optional] [readonly] 
**ThumbnailUrl** | Pointer to **string** | The URL at which this app&#39;s thumbnail may be accessed.  | [optional] [readonly] 
**Scopes** | Pointer to **string** | The OAuth scopes this app was authorized with.  This defines what parts of your Account the app is allowed to access.  | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this app was authorized. | [optional] [readonly] 
**Expiry** | Pointer to **NullableTime** | When the app&#39;s access to your account expires. If &#x60;null&#x60;, the app&#39;s access must be revoked manually.  | [optional] [readonly] 
**Website** | Pointer to **string** | The website where you can get more information about this app.  | [optional] [readonly] 

## Methods

### NewAuthorizedApp

`func NewAuthorizedApp() *AuthorizedApp`

NewAuthorizedApp instantiates a new AuthorizedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizedAppWithDefaults

`func NewAuthorizedAppWithDefaults() *AuthorizedApp`

NewAuthorizedAppWithDefaults instantiates a new AuthorizedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizedApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizedApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizedApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizedApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AuthorizedApp) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AuthorizedApp) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AuthorizedApp) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AuthorizedApp) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *AuthorizedApp) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AuthorizedApp) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AuthorizedApp) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AuthorizedApp) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizedApp) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizedApp) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizedApp) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizedApp) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorizedApp) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorizedApp) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorizedApp) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorizedApp) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *AuthorizedApp) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AuthorizedApp) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AuthorizedApp) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AuthorizedApp) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *AuthorizedApp) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *AuthorizedApp) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetWebsite

`func (o *AuthorizedApp) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AuthorizedApp) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AuthorizedApp) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AuthorizedApp) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


