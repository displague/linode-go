# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | Your unique ID in our system. This value will never change, and can safely be used to identify your User.  | [optional] [readonly] 
**Username** | Pointer to **string** | Your username, used for logging in to our system.  | [optional] [readonly] 
**Email** | Pointer to **string** | Your email address.  This address will be used for communication with Linode as necessary.  | [optional] 
**Timezone** | Pointer to **string** | The timezone you prefer to see times in. This is not used by the API directly. It is provided for the benefit of clients such as the Linode Cloud Manager and other clients built on the API. All times returned by the API are in UTC.  | [optional] 
**EmailNotifications** | Pointer to **bool** | If true, you will receive email notifications about account activity.  If false, you may still receive business-critical communications through email.  | [optional] 
**Referrals** | Pointer to [**ProfileReferrals**](ProfileReferrals.md) |  | [optional] 
**IpWhitelistEnabled** | Pointer to **bool** | If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.  If you disable this setting, you will not be able to re-enable it.  | [optional] 
**LishAuthMethod** | Pointer to **string** | The authentication methods that are allowed when connecting to [the Linode Shell (Lish)](/docs/platform/manager/using-the-linode-shell-lish/). * &#x60;keys_only&#x60; is the most secure if you intend to use Lish. * &#x60;disabled&#x60; is recommended if you do not intend to use Lish at all. * If this account&#39;s Cloud Manager authentication type is set to a Third-Party Authentication method, &#x60;password_keys&#x60; cannot be used as your Lish authentication method. To view this account&#39;s Cloud Manager &#x60;authentication_type&#x60; field, send a request to the [View Profile](/docs/api/profile/#profile-view) endpoint.  | [optional] 
**AuthorizedKeys** | Pointer to **[]string** | The list of SSH Keys authorized to use Lish for your User. This value is ignored if &#x60;lish_auth_method&#x60; is \&quot;disabled.\&quot;  | [optional] 
**TwoFactorAuth** | Pointer to **bool** | If true, logins from untrusted computers will require Two Factor Authentication.  See [/profile/tfa-enable](/docs/api/profile/#two-factor-secret-create) to enable Two Factor Authentication.  | [optional] 
**Restricted** | Pointer to **bool** | If true, your User has restrictions on what can be accessed on your Account. To get details on what entities/actions you can access/perform, see [/profile/grants](/docs/api/profile/#grants-list).  | [optional] 
**AuthenticationType** | Pointer to **string** | This account&#39;s Cloud Manager authentication type. Authentication types are chosen through Cloud Manager and authorized when logging into your account. These authentication types are either the user&#39;s password (in conjunction with their username), or the name of their indentity provider such as GitHub. For example, if a user:  - Has never used Third-Party Authentication, their authentication type will be &#x60;password&#x60;. - Is using Third-Party Authentication, their authentication type will be the name of their Identity Provider (eg. &#x60;github&#x60;). - Has used Third-Party Authentication and has since revoked it, their authentication type will be &#x60;password&#x60;.   **Note:** This functionality may not yet be available in Cloud Manager. See the [Cloud Manager Changelog](/changelog/cloud-manager/) for the latest updates.  | [optional] [readonly] 

## Methods

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Profile) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Profile) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Profile) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Profile) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *Profile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Profile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Profile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Profile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *Profile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Profile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Profile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Profile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTimezone

`func (o *Profile) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Profile) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Profile) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Profile) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *Profile) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *Profile) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *Profile) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *Profile) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetReferrals

`func (o *Profile) GetReferrals() ProfileReferrals`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *Profile) GetReferralsOk() (*ProfileReferrals, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *Profile) SetReferrals(v ProfileReferrals)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *Profile) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetIpWhitelistEnabled

`func (o *Profile) GetIpWhitelistEnabled() bool`

GetIpWhitelistEnabled returns the IpWhitelistEnabled field if non-nil, zero value otherwise.

### GetIpWhitelistEnabledOk

`func (o *Profile) GetIpWhitelistEnabledOk() (*bool, bool)`

GetIpWhitelistEnabledOk returns a tuple with the IpWhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelistEnabled

`func (o *Profile) SetIpWhitelistEnabled(v bool)`

SetIpWhitelistEnabled sets IpWhitelistEnabled field to given value.

### HasIpWhitelistEnabled

`func (o *Profile) HasIpWhitelistEnabled() bool`

HasIpWhitelistEnabled returns a boolean if a field has been set.

### GetLishAuthMethod

`func (o *Profile) GetLishAuthMethod() string`

GetLishAuthMethod returns the LishAuthMethod field if non-nil, zero value otherwise.

### GetLishAuthMethodOk

`func (o *Profile) GetLishAuthMethodOk() (*string, bool)`

GetLishAuthMethodOk returns a tuple with the LishAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLishAuthMethod

`func (o *Profile) SetLishAuthMethod(v string)`

SetLishAuthMethod sets LishAuthMethod field to given value.

### HasLishAuthMethod

`func (o *Profile) HasLishAuthMethod() bool`

HasLishAuthMethod returns a boolean if a field has been set.

### GetAuthorizedKeys

`func (o *Profile) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *Profile) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *Profile) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *Profile) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### SetAuthorizedKeysNil

`func (o *Profile) SetAuthorizedKeysNil(b bool)`

 SetAuthorizedKeysNil sets the value for AuthorizedKeys to be an explicit nil

### UnsetAuthorizedKeys
`func (o *Profile) UnsetAuthorizedKeys()`

UnsetAuthorizedKeys ensures that no value is present for AuthorizedKeys, not even an explicit nil
### GetTwoFactorAuth

`func (o *Profile) GetTwoFactorAuth() bool`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *Profile) GetTwoFactorAuthOk() (*bool, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *Profile) SetTwoFactorAuth(v bool)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *Profile) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetRestricted

`func (o *Profile) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *Profile) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *Profile) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *Profile) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *Profile) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Profile) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *Profile) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *Profile) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


