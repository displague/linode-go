# PersonalAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This token&#39;s unique ID, which can be used to revoke it.  | [optional] [readonly] 
**Scopes** | Pointer to **string** | The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;https://github.com/linode/linode-cli\&quot;&gt;Linode CLI&lt;/a&gt;, require tokens with access to &#x60;*&#x60;. Tokens with more restrictive scopes are generally more secure.  | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date and time this token was created.  | [optional] [readonly] 
**Label** | Pointer to **string** | This token&#39;s label.  This is for display purposes only, but can be used to more easily track what you&#39;re using each token for.  | [optional] 
**Token** | Pointer to **string** | The token used to access the API.  When the token is created, the full token is returned here.  Otherwise, only the first 16 characters are returned.  | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | When this token will expire.  Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated.  Tokens may be created with \&quot;null\&quot; as their expiry and will never expire unless revoked.  | [optional] [readonly] 

## Methods

### NewPersonalAccessToken

`func NewPersonalAccessToken() *PersonalAccessToken`

NewPersonalAccessToken instantiates a new PersonalAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalAccessTokenWithDefaults

`func NewPersonalAccessTokenWithDefaults() *PersonalAccessToken`

NewPersonalAccessTokenWithDefaults instantiates a new PersonalAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonalAccessToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonalAccessToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonalAccessToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PersonalAccessToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScopes

`func (o *PersonalAccessToken) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PersonalAccessToken) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PersonalAccessToken) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PersonalAccessToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetCreated

`func (o *PersonalAccessToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PersonalAccessToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PersonalAccessToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PersonalAccessToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLabel

`func (o *PersonalAccessToken) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PersonalAccessToken) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PersonalAccessToken) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PersonalAccessToken) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetToken

`func (o *PersonalAccessToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PersonalAccessToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PersonalAccessToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PersonalAccessToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpiry

`func (o *PersonalAccessToken) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PersonalAccessToken) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PersonalAccessToken) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PersonalAccessToken) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


