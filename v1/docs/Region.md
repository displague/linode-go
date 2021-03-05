# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of this Region. | [optional] [readonly] 
**Country** | Pointer to **string** | The country where this Region resides. | [optional] [readonly] 
**Capabilities** | Pointer to **[]string** | A list of capabilities of this region.  | [optional] [readonly] 
**Status** | Pointer to **string** | This region&#39;s current operational status.  | [optional] [readonly] 
**Resolvers** | Pointer to [**RegionResolvers**](RegionResolvers.md) |  | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountry

`func (o *Region) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Region) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Region) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Region) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCapabilities

`func (o *Region) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Region) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Region) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Region) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetStatus

`func (o *Region) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Region) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Region) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Region) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolvers

`func (o *Region) GetResolvers() RegionResolvers`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *Region) GetResolversOk() (*RegionResolvers, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *Region) SetResolvers(v RegionResolvers)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *Region) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


