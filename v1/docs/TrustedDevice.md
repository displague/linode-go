# TrustedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for this TrustedDevice | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Remember Me session was started.  This corresponds to the time of login with the \&quot;Remember Me\&quot; box checked.  | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | When this TrustedDevice session expires.  Sessions typically last 30 days.  | [optional] [readonly] 
**UserAgent** | Pointer to **string** | The User Agent of the browser that created this TrustedDevice session.  | [optional] [readonly] 
**LastAuthenticated** | Pointer to **time.Time** | The last time this TrustedDevice was successfully used to authenticate to &lt;a target&#x3D;\&quot;_top\&quot; href&#x3D;\&quot;https://login.linode.com\&quot;&gt;login.linode.com&lt;/a&gt;.  | [optional] [readonly] 
**LastRemoteAddr** | Pointer to **string** | The last IP Address to successfully authenticate with this TrustedDevice.  | [optional] [readonly] 

## Methods

### NewTrustedDevice

`func NewTrustedDevice() *TrustedDevice`

NewTrustedDevice instantiates a new TrustedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedDeviceWithDefaults

`func NewTrustedDeviceWithDefaults() *TrustedDevice`

NewTrustedDeviceWithDefaults instantiates a new TrustedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedDevice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrustedDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *TrustedDevice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrustedDevice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrustedDevice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrustedDevice) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *TrustedDevice) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TrustedDevice) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TrustedDevice) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *TrustedDevice) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetUserAgent

`func (o *TrustedDevice) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *TrustedDevice) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *TrustedDevice) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *TrustedDevice) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetLastAuthenticated

`func (o *TrustedDevice) GetLastAuthenticated() time.Time`

GetLastAuthenticated returns the LastAuthenticated field if non-nil, zero value otherwise.

### GetLastAuthenticatedOk

`func (o *TrustedDevice) GetLastAuthenticatedOk() (*time.Time, bool)`

GetLastAuthenticatedOk returns a tuple with the LastAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthenticated

`func (o *TrustedDevice) SetLastAuthenticated(v time.Time)`

SetLastAuthenticated sets LastAuthenticated field to given value.

### HasLastAuthenticated

`func (o *TrustedDevice) HasLastAuthenticated() bool`

HasLastAuthenticated returns a boolean if a field has been set.

### GetLastRemoteAddr

`func (o *TrustedDevice) GetLastRemoteAddr() string`

GetLastRemoteAddr returns the LastRemoteAddr field if non-nil, zero value otherwise.

### GetLastRemoteAddrOk

`func (o *TrustedDevice) GetLastRemoteAddrOk() (*string, bool)`

GetLastRemoteAddrOk returns a tuple with the LastRemoteAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRemoteAddr

`func (o *TrustedDevice) SetLastRemoteAddr(v string)`

SetLastRemoteAddr sets LastRemoteAddr field to given value.

### HasLastRemoteAddr

`func (o *TrustedDevice) HasLastRemoteAddr() bool`

HasLastRemoteAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


