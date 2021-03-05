# RegionResolvers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to **string** | The IPv4 addresses for this region&#39;s DNS resolvers, separated by commas.  | [optional] [readonly] 
**Ipv6** | Pointer to **string** | The IPv6 addresses for this region&#39;s DNS resolvers, separated by commas.  | [optional] [readonly] 

## Methods

### NewRegionResolvers

`func NewRegionResolvers() *RegionResolvers`

NewRegionResolvers instantiates a new RegionResolvers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionResolversWithDefaults

`func NewRegionResolversWithDefaults() *RegionResolvers`

NewRegionResolversWithDefaults instantiates a new RegionResolvers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *RegionResolvers) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RegionResolvers) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RegionResolvers) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RegionResolvers) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RegionResolvers) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RegionResolvers) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RegionResolvers) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RegionResolvers) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


