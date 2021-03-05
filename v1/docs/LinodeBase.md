# LinodeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Linode&#39;s label is for display purposes only. If no label is provided for a Linode, a default will be assigned. Linode labels have the following constraints:    * Must begin and end with an alphanumeric character.   * May only consist of alphanumeric characters, dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;).   * Cannot have two dashes (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row.  | [optional] 
**Group** | Pointer to **string** | A deprecated property denoting a group label for this Linode.  | [optional] 
**Type** | Pointer to [**Id**](Id.md) |  | [optional] 
**Region** | Pointer to [**Id**](Id.md) |  | [optional] 

## Methods

### NewLinodeBase

`func NewLinodeBase() *LinodeBase`

NewLinodeBase instantiates a new LinodeBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeBaseWithDefaults

`func NewLinodeBaseWithDefaults() *LinodeBase`

NewLinodeBaseWithDefaults instantiates a new LinodeBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LinodeBase) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinodeBase) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinodeBase) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LinodeBase) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroup

`func (o *LinodeBase) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LinodeBase) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LinodeBase) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *LinodeBase) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetType

`func (o *LinodeBase) GetType() Id`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinodeBase) GetTypeOk() (*Id, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinodeBase) SetType(v Id)`

SetType sets Type field to given value.

### HasType

`func (o *LinodeBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *LinodeBase) GetRegion() Id`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LinodeBase) GetRegionOk() (*Id, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LinodeBase) SetRegion(v Id)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LinodeBase) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


