# Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Firewall&#39;s unique ID.  | [optional] [readonly] 
**Label** | Pointer to **string** | The Firewall&#39;s label, for display purposes only. Firewall labels have the following constraints:    * Must begin and end with an alphanumeric character.   * May only consist of alphanumeric characters, dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;).   * Cannot have two dashes (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row.   * Must be between 3 and 32 characters.   * Must be unique.  | [optional] 
**Created** | Pointer to **time.Time** | When this Firewall was created.  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Firewall was last updated.  | [optional] [readonly] 
**Status** | Pointer to **string** | The status of this Firewall.    * When a Firewall is first created its status is &#x60;enabled&#x60;.   * Use the [Update Firewall](/docs/api/networking/#firewall-update) endpoint to set a Firewall&#39;s status to &#x60;enbaled&#x60; or &#x60;disabled&#x60;.   * Use the [Delete Firewall](/docs/api/networking/#firewall-delete) endpoint to delete a Firewall.  | [optional] 
**Rules** | Pointer to [**FirewallRules**](FirewallRules.md) |  | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object. Tags are for organizational purposes only.  | [optional] 

## Methods

### NewFirewall

`func NewFirewall() *Firewall`

NewFirewall instantiates a new Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallWithDefaults

`func NewFirewallWithDefaults() *Firewall`

NewFirewallWithDefaults instantiates a new Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Firewall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Firewall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Firewall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Firewall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Firewall) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Firewall) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Firewall) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Firewall) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreated

`func (o *Firewall) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Firewall) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Firewall) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Firewall) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Firewall) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Firewall) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Firewall) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Firewall) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *Firewall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Firewall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Firewall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Firewall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRules

`func (o *Firewall) GetRules() FirewallRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Firewall) GetRulesOk() (*FirewallRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Firewall) SetRules(v FirewallRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Firewall) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *Firewall) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Firewall) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Firewall) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Firewall) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


