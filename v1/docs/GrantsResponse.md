# GrantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**GrantsResponseGlobal**](GrantsResponseGlobal.md) |  | [optional] 
**Linode** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to Linodes on this Account. There will be one entry per Linode on the Account.  | [optional] 
**Domain** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to Domains on this Account. There will be one entry per Domain on the Account.  | [optional] 
**Nodebalancer** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to NodeBalancers on this Account. There will be one entry per NodeBalancer on the Account.  | [optional] 
**Image** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to Images on this Account. There will be one entry per Image on the Account.  | [optional] 
**Longview** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to Longview Clients on this Account. There will be one entry per Longview Client on the Account.  | [optional] 
**Stackscript** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to StackScripts on this Account.  There will be one entry per StackScript on the Account.  | [optional] 
**Volume** | Pointer to [**[]Grant**](Grant.md) | The grants this User has pertaining to Volumes on this Account. There will be one entry per Volume on the Account.  | [optional] 

## Methods

### NewGrantsResponse

`func NewGrantsResponse() *GrantsResponse`

NewGrantsResponse instantiates a new GrantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantsResponseWithDefaults

`func NewGrantsResponseWithDefaults() *GrantsResponse`

NewGrantsResponseWithDefaults instantiates a new GrantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *GrantsResponse) GetGlobal() GrantsResponseGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *GrantsResponse) GetGlobalOk() (*GrantsResponseGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *GrantsResponse) SetGlobal(v GrantsResponseGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *GrantsResponse) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetLinode

`func (o *GrantsResponse) GetLinode() []Grant`

GetLinode returns the Linode field if non-nil, zero value otherwise.

### GetLinodeOk

`func (o *GrantsResponse) GetLinodeOk() (*[]Grant, bool)`

GetLinodeOk returns a tuple with the Linode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinode

`func (o *GrantsResponse) SetLinode(v []Grant)`

SetLinode sets Linode field to given value.

### HasLinode

`func (o *GrantsResponse) HasLinode() bool`

HasLinode returns a boolean if a field has been set.

### GetDomain

`func (o *GrantsResponse) GetDomain() []Grant`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GrantsResponse) GetDomainOk() (*[]Grant, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GrantsResponse) SetDomain(v []Grant)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GrantsResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNodebalancer

`func (o *GrantsResponse) GetNodebalancer() []Grant`

GetNodebalancer returns the Nodebalancer field if non-nil, zero value otherwise.

### GetNodebalancerOk

`func (o *GrantsResponse) GetNodebalancerOk() (*[]Grant, bool)`

GetNodebalancerOk returns a tuple with the Nodebalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancer

`func (o *GrantsResponse) SetNodebalancer(v []Grant)`

SetNodebalancer sets Nodebalancer field to given value.

### HasNodebalancer

`func (o *GrantsResponse) HasNodebalancer() bool`

HasNodebalancer returns a boolean if a field has been set.

### GetImage

`func (o *GrantsResponse) GetImage() []Grant`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GrantsResponse) GetImageOk() (*[]Grant, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GrantsResponse) SetImage(v []Grant)`

SetImage sets Image field to given value.

### HasImage

`func (o *GrantsResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLongview

`func (o *GrantsResponse) GetLongview() []Grant`

GetLongview returns the Longview field if non-nil, zero value otherwise.

### GetLongviewOk

`func (o *GrantsResponse) GetLongviewOk() (*[]Grant, bool)`

GetLongviewOk returns a tuple with the Longview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongview

`func (o *GrantsResponse) SetLongview(v []Grant)`

SetLongview sets Longview field to given value.

### HasLongview

`func (o *GrantsResponse) HasLongview() bool`

HasLongview returns a boolean if a field has been set.

### GetStackscript

`func (o *GrantsResponse) GetStackscript() []Grant`

GetStackscript returns the Stackscript field if non-nil, zero value otherwise.

### GetStackscriptOk

`func (o *GrantsResponse) GetStackscriptOk() (*[]Grant, bool)`

GetStackscriptOk returns a tuple with the Stackscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscript

`func (o *GrantsResponse) SetStackscript(v []Grant)`

SetStackscript sets Stackscript field to given value.

### HasStackscript

`func (o *GrantsResponse) HasStackscript() bool`

HasStackscript returns a boolean if a field has been set.

### GetVolume

`func (o *GrantsResponse) GetVolume() []Grant`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GrantsResponse) GetVolumeOk() (*[]Grant, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GrantsResponse) SetVolume(v []Grant)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GrantsResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


