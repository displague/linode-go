# OAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The OAuth Client ID.  This is used to identify the client, and is a publicly-known value (it is not a secret).  | [optional] [readonly] 
**RedirectUri** | Pointer to **string** | The location a successful log in from &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;https://login.linode.com\&quot;&gt;https://login.linode.com&lt;/a&gt; should be redirected to for this client.  The receiver of this redirect should be ready to accept an OAuth exchange code and finish the OAuth exchange.  | [optional] 
**Label** | Pointer to **string** | The name of this application.  This will be presented to users when they are asked to grant it access to their Account.  | [optional] 
**Status** | Pointer to **string** | The status of this application.  &#x60;active&#x60; by default.  | [optional] [readonly] 
**Secret** | Pointer to **string** | The OAuth Client secret, used in the OAuth exchange.  This is returned as &#x60;&lt;REDACTED&gt;&#x60; except when an OAuth Client is created or its secret is reset.  This is a secret, and should not be shared or disclosed publicly.  | [optional] [readonly] 
**ThumbnailUrl** | Pointer to **NullableString** | The URL where this client&#39;s thumbnail may be viewed, or &#x60;null&#x60; if this client does not have a thumbnail set.  | [optional] [readonly] 
**Public** | Pointer to **bool** | If this is a public or private OAuth Client.  Public clients have a slightly different authentication workflow than private clients.  See the &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;https://oauth.net/2/\&quot;&gt;OAuth spec&lt;/a&gt; for more details.  | [optional] [readonly] 

## Methods

### NewOAuthClient

`func NewOAuthClient() *OAuthClient`

NewOAuthClient instantiates a new OAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientWithDefaults

`func NewOAuthClientWithDefaults() *OAuthClient`

NewOAuthClientWithDefaults instantiates a new OAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuthClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuthClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuthClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuthClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OAuthClient) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OAuthClient) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OAuthClient) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OAuthClient) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetLabel

`func (o *OAuthClient) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OAuthClient) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OAuthClient) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OAuthClient) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *OAuthClient) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuthClient) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuthClient) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuthClient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecret

`func (o *OAuthClient) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OAuthClient) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OAuthClient) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OAuthClient) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *OAuthClient) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *OAuthClient) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *OAuthClient) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *OAuthClient) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *OAuthClient) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *OAuthClient) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetPublic

`func (o *OAuthClient) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *OAuthClient) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *OAuthClient) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *OAuthClient) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


