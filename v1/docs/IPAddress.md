# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address.  | [optional] [readonly] 
**Gateway** | Pointer to **NullableString** | The default gateway for this address.  | [optional] [readonly] 
**SubnetMask** | Pointer to **string** | The mask that separates host bits from network bits for this address.  | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The number of bits set in the subnet mask.  | [optional] [readonly] 
**Type** | Pointer to **string** | The type of address this is.  | [optional] [readonly] 
**Public** | Pointer to **bool** | Whether this is a public or private IP address.  | [optional] [readonly] 
**Rdns** | Pointer to **string** | The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.  | [optional] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this address currently belongs to. For IPv4 addresses, this is by default the Linode that this address was assigned to on creation, and these addresses my be moved using the [/networking/ipv4/assign](/docs/api/networking/#ips-to-linodes-assign) endpoint. For SLAAC and link-local addresses, this value may not be changed.  | [optional] [readonly] 
**Region** | Pointer to **string** | The Region this IP address resides in.  | [optional] [readonly] 

## Methods

### NewIPAddress

`func NewIPAddress() *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *IPAddress) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IPAddress) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IPAddress) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IPAddress) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *IPAddress) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *IPAddress) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSubnetMask

`func (o *IPAddress) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *IPAddress) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *IPAddress) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *IPAddress) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetPrefix

`func (o *IPAddress) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IPAddress) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IPAddress) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IPAddress) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *IPAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublic

`func (o *IPAddress) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPAddress) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPAddress) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPAddress) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRdns

`func (o *IPAddress) GetRdns() string`

GetRdns returns the Rdns field if non-nil, zero value otherwise.

### GetRdnsOk

`func (o *IPAddress) GetRdnsOk() (*string, bool)`

GetRdnsOk returns a tuple with the Rdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdns

`func (o *IPAddress) SetRdns(v string)`

SetRdns sets Rdns field to given value.

### HasRdns

`func (o *IPAddress) HasRdns() bool`

HasRdns returns a boolean if a field has been set.

### GetLinodeId

`func (o *IPAddress) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *IPAddress) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *IPAddress) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *IPAddress) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetRegion

`func (o *IPAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IPAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IPAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IPAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


