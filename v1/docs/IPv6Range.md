# IPv6Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** | The IPv6 range of addresses in this pool.  | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The prefix length of the address, denoting how many addresses can be assigned from this range calculated as 2 &lt;sup&gt;128-prefix&lt;/sup&gt;.  | [optional] 
**Region** | Pointer to **string** | The region for this range of IPv6 addresses.  | [optional] [readonly] 
**RouteTarget** | Pointer to **NullableString** | The last address in this block of IPv6 addresses.  | [optional] 

## Methods

### NewIPv6Range

`func NewIPv6Range() *IPv6Range`

NewIPv6Range instantiates a new IPv6Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6RangeWithDefaults

`func NewIPv6RangeWithDefaults() *IPv6Range`

NewIPv6RangeWithDefaults instantiates a new IPv6Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *IPv6Range) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *IPv6Range) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *IPv6Range) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *IPv6Range) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetPrefix

`func (o *IPv6Range) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IPv6Range) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IPv6Range) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IPv6Range) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *IPv6Range) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IPv6Range) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IPv6Range) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IPv6Range) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRouteTarget

`func (o *IPv6Range) GetRouteTarget() string`

GetRouteTarget returns the RouteTarget field if non-nil, zero value otherwise.

### GetRouteTargetOk

`func (o *IPv6Range) GetRouteTargetOk() (*string, bool)`

GetRouteTargetOk returns a tuple with the RouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTarget

`func (o *IPv6Range) SetRouteTarget(v string)`

SetRouteTarget sets RouteTarget field to given value.

### HasRouteTarget

`func (o *IPv6Range) HasRouteTarget() bool`

HasRouteTarget returns a boolean if a field has been set.

### SetRouteTargetNil

`func (o *IPv6Range) SetRouteTargetNil(b bool)`

 SetRouteTargetNil sets the value for RouteTarget to be an explicit nil

### UnsetRouteTarget
`func (o *IPv6Range) UnsetRouteTarget()`

UnsetRouteTarget ensures that no value is present for RouteTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


