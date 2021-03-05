# GrantsResponseGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddLinodes** | Pointer to **bool** | If true, this User may create Linodes. | [optional] 
**AddLongview** | Pointer to **bool** | If true, this User may create Longview clients and view the current plan. | [optional] 
**LongviewSubscription** | Pointer to **bool** | If true, this User may manage the Account&#39;s Longview subscription. | [optional] 
**AccountAccess** | Pointer to **NullableString** | The level of access this User has to Account-level actions, like billing information. A restricted User will never be able to manage users.  | [optional] 
**CancelAccount** | Pointer to **bool** | If true, this User may cancel the entire Account. | [optional] 
**AddDomains** | Pointer to **bool** | If true, this User may add Domains. | [optional] 
**AddStackscripts** | Pointer to **bool** | If true, this User may add StackScripts. | [optional] 
**AddNodebalancers** | Pointer to **bool** | If true, this User may add NodeBalancers. | [optional] 
**AddImages** | Pointer to **bool** | If true, this User may add Images. | [optional] 
**AddVolumes** | Pointer to **bool** | If true, this User may add Volumes. | [optional] 

## Methods

### NewGrantsResponseGlobal

`func NewGrantsResponseGlobal() *GrantsResponseGlobal`

NewGrantsResponseGlobal instantiates a new GrantsResponseGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantsResponseGlobalWithDefaults

`func NewGrantsResponseGlobalWithDefaults() *GrantsResponseGlobal`

NewGrantsResponseGlobalWithDefaults instantiates a new GrantsResponseGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddLinodes

`func (o *GrantsResponseGlobal) GetAddLinodes() bool`

GetAddLinodes returns the AddLinodes field if non-nil, zero value otherwise.

### GetAddLinodesOk

`func (o *GrantsResponseGlobal) GetAddLinodesOk() (*bool, bool)`

GetAddLinodesOk returns a tuple with the AddLinodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLinodes

`func (o *GrantsResponseGlobal) SetAddLinodes(v bool)`

SetAddLinodes sets AddLinodes field to given value.

### HasAddLinodes

`func (o *GrantsResponseGlobal) HasAddLinodes() bool`

HasAddLinodes returns a boolean if a field has been set.

### GetAddLongview

`func (o *GrantsResponseGlobal) GetAddLongview() bool`

GetAddLongview returns the AddLongview field if non-nil, zero value otherwise.

### GetAddLongviewOk

`func (o *GrantsResponseGlobal) GetAddLongviewOk() (*bool, bool)`

GetAddLongviewOk returns a tuple with the AddLongview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLongview

`func (o *GrantsResponseGlobal) SetAddLongview(v bool)`

SetAddLongview sets AddLongview field to given value.

### HasAddLongview

`func (o *GrantsResponseGlobal) HasAddLongview() bool`

HasAddLongview returns a boolean if a field has been set.

### GetLongviewSubscription

`func (o *GrantsResponseGlobal) GetLongviewSubscription() bool`

GetLongviewSubscription returns the LongviewSubscription field if non-nil, zero value otherwise.

### GetLongviewSubscriptionOk

`func (o *GrantsResponseGlobal) GetLongviewSubscriptionOk() (*bool, bool)`

GetLongviewSubscriptionOk returns a tuple with the LongviewSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewSubscription

`func (o *GrantsResponseGlobal) SetLongviewSubscription(v bool)`

SetLongviewSubscription sets LongviewSubscription field to given value.

### HasLongviewSubscription

`func (o *GrantsResponseGlobal) HasLongviewSubscription() bool`

HasLongviewSubscription returns a boolean if a field has been set.

### GetAccountAccess

`func (o *GrantsResponseGlobal) GetAccountAccess() string`

GetAccountAccess returns the AccountAccess field if non-nil, zero value otherwise.

### GetAccountAccessOk

`func (o *GrantsResponseGlobal) GetAccountAccessOk() (*string, bool)`

GetAccountAccessOk returns a tuple with the AccountAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAccess

`func (o *GrantsResponseGlobal) SetAccountAccess(v string)`

SetAccountAccess sets AccountAccess field to given value.

### HasAccountAccess

`func (o *GrantsResponseGlobal) HasAccountAccess() bool`

HasAccountAccess returns a boolean if a field has been set.

### SetAccountAccessNil

`func (o *GrantsResponseGlobal) SetAccountAccessNil(b bool)`

 SetAccountAccessNil sets the value for AccountAccess to be an explicit nil

### UnsetAccountAccess
`func (o *GrantsResponseGlobal) UnsetAccountAccess()`

UnsetAccountAccess ensures that no value is present for AccountAccess, not even an explicit nil
### GetCancelAccount

`func (o *GrantsResponseGlobal) GetCancelAccount() bool`

GetCancelAccount returns the CancelAccount field if non-nil, zero value otherwise.

### GetCancelAccountOk

`func (o *GrantsResponseGlobal) GetCancelAccountOk() (*bool, bool)`

GetCancelAccountOk returns a tuple with the CancelAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAccount

`func (o *GrantsResponseGlobal) SetCancelAccount(v bool)`

SetCancelAccount sets CancelAccount field to given value.

### HasCancelAccount

`func (o *GrantsResponseGlobal) HasCancelAccount() bool`

HasCancelAccount returns a boolean if a field has been set.

### GetAddDomains

`func (o *GrantsResponseGlobal) GetAddDomains() bool`

GetAddDomains returns the AddDomains field if non-nil, zero value otherwise.

### GetAddDomainsOk

`func (o *GrantsResponseGlobal) GetAddDomainsOk() (*bool, bool)`

GetAddDomainsOk returns a tuple with the AddDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDomains

`func (o *GrantsResponseGlobal) SetAddDomains(v bool)`

SetAddDomains sets AddDomains field to given value.

### HasAddDomains

`func (o *GrantsResponseGlobal) HasAddDomains() bool`

HasAddDomains returns a boolean if a field has been set.

### GetAddStackscripts

`func (o *GrantsResponseGlobal) GetAddStackscripts() bool`

GetAddStackscripts returns the AddStackscripts field if non-nil, zero value otherwise.

### GetAddStackscriptsOk

`func (o *GrantsResponseGlobal) GetAddStackscriptsOk() (*bool, bool)`

GetAddStackscriptsOk returns a tuple with the AddStackscripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddStackscripts

`func (o *GrantsResponseGlobal) SetAddStackscripts(v bool)`

SetAddStackscripts sets AddStackscripts field to given value.

### HasAddStackscripts

`func (o *GrantsResponseGlobal) HasAddStackscripts() bool`

HasAddStackscripts returns a boolean if a field has been set.

### GetAddNodebalancers

`func (o *GrantsResponseGlobal) GetAddNodebalancers() bool`

GetAddNodebalancers returns the AddNodebalancers field if non-nil, zero value otherwise.

### GetAddNodebalancersOk

`func (o *GrantsResponseGlobal) GetAddNodebalancersOk() (*bool, bool)`

GetAddNodebalancersOk returns a tuple with the AddNodebalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNodebalancers

`func (o *GrantsResponseGlobal) SetAddNodebalancers(v bool)`

SetAddNodebalancers sets AddNodebalancers field to given value.

### HasAddNodebalancers

`func (o *GrantsResponseGlobal) HasAddNodebalancers() bool`

HasAddNodebalancers returns a boolean if a field has been set.

### GetAddImages

`func (o *GrantsResponseGlobal) GetAddImages() bool`

GetAddImages returns the AddImages field if non-nil, zero value otherwise.

### GetAddImagesOk

`func (o *GrantsResponseGlobal) GetAddImagesOk() (*bool, bool)`

GetAddImagesOk returns a tuple with the AddImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddImages

`func (o *GrantsResponseGlobal) SetAddImages(v bool)`

SetAddImages sets AddImages field to given value.

### HasAddImages

`func (o *GrantsResponseGlobal) HasAddImages() bool`

HasAddImages returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GrantsResponseGlobal) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GrantsResponseGlobal) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GrantsResponseGlobal) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GrantsResponseGlobal) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


