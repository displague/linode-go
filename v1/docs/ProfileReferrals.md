# ProfileReferrals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Your referral code.  If others use this when signing up for Linode, you will receive account credit.  | [optional] [readonly] 
**Url** | Pointer to **string** | Your referral url, used to direct others to sign up for Linode with your referral code.  | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of users who have signed up with your referral code.  | [optional] [readonly] 
**Completed** | Pointer to **int32** | The number of completed signups with your referral code.  | [optional] [readonly] 
**Pending** | Pointer to **int32** | The number of pending signups with your referral code.  You will not receive credit for these signups until they are completed.  | [optional] [readonly] 
**Credit** | Pointer to **int32** | The amount of account credit in US Dollars issued to you through the referral program.  | [optional] [readonly] 

## Methods

### NewProfileReferrals

`func NewProfileReferrals() *ProfileReferrals`

NewProfileReferrals instantiates a new ProfileReferrals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileReferralsWithDefaults

`func NewProfileReferralsWithDefaults() *ProfileReferrals`

NewProfileReferralsWithDefaults instantiates a new ProfileReferrals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ProfileReferrals) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProfileReferrals) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProfileReferrals) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProfileReferrals) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUrl

`func (o *ProfileReferrals) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProfileReferrals) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProfileReferrals) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProfileReferrals) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTotal

`func (o *ProfileReferrals) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProfileReferrals) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProfileReferrals) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProfileReferrals) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCompleted

`func (o *ProfileReferrals) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ProfileReferrals) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ProfileReferrals) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ProfileReferrals) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetPending

`func (o *ProfileReferrals) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ProfileReferrals) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ProfileReferrals) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ProfileReferrals) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetCredit

`func (o *ProfileReferrals) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *ProfileReferrals) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *ProfileReferrals) SetCredit(v int32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *ProfileReferrals) HasCredit() bool`

HasCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


