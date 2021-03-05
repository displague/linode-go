# IPAddressV6LinkLocal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv6 link-local address.  | [optional] [readonly] 
**Gateway** | Pointer to **string** | The default gateway for this address.  | [optional] [readonly] 
**SubnetMask** | Pointer to **string** | The subnet mask.  | [optional] [readonly] 
**Prefix** | Pointer to **int32** | The network prefix.  | [optional] [readonly] 
**Type** | Pointer to **string** | The type of address this is.  | [optional] [readonly] 
**Public** | Pointer to **bool** | Whether this is a public or private IP address.  | [optional] [readonly] 
**Rdns** | Pointer to **string** | The reverse DNS assigned to this address.  | [optional] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this address currently belongs to.  | [optional] [readonly] 
**Region** | Pointer to **string** | The Region this address resides in.  | [optional] [readonly] 

## Methods

### NewIPAddressV6LinkLocal

`func NewIPAddressV6LinkLocal() *IPAddressV6LinkLocal`

NewIPAddressV6LinkLocal instantiates a new IPAddressV6LinkLocal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressV6LinkLocalWithDefaults

`func NewIPAddressV6LinkLocalWithDefaults() *IPAddressV6LinkLocal`

NewIPAddressV6LinkLocalWithDefaults instantiates a new IPAddressV6LinkLocal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPAddressV6LinkLocal) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddressV6LinkLocal) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddressV6LinkLocal) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPAddressV6LinkLocal) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *IPAddressV6LinkLocal) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IPAddressV6LinkLocal) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IPAddressV6LinkLocal) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IPAddressV6LinkLocal) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetSubnetMask

`func (o *IPAddressV6LinkLocal) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *IPAddressV6LinkLocal) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *IPAddressV6LinkLocal) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *IPAddressV6LinkLocal) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetPrefix

`func (o *IPAddressV6LinkLocal) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IPAddressV6LinkLocal) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IPAddressV6LinkLocal) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IPAddressV6LinkLocal) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *IPAddressV6LinkLocal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAddressV6LinkLocal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAddressV6LinkLocal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPAddressV6LinkLocal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublic

`func (o *IPAddressV6LinkLocal) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPAddressV6LinkLocal) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPAddressV6LinkLocal) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPAddressV6LinkLocal) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRdns

`func (o *IPAddressV6LinkLocal) GetRdns() string`

GetRdns returns the Rdns field if non-nil, zero value otherwise.

### GetRdnsOk

`func (o *IPAddressV6LinkLocal) GetRdnsOk() (*string, bool)`

GetRdnsOk returns a tuple with the Rdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdns

`func (o *IPAddressV6LinkLocal) SetRdns(v string)`

SetRdns sets Rdns field to given value.

### HasRdns

`func (o *IPAddressV6LinkLocal) HasRdns() bool`

HasRdns returns a boolean if a field has been set.

### GetLinodeId

`func (o *IPAddressV6LinkLocal) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *IPAddressV6LinkLocal) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *IPAddressV6LinkLocal) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *IPAddressV6LinkLocal) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetRegion

`func (o *IPAddressV6LinkLocal) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IPAddressV6LinkLocal) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IPAddressV6LinkLocal) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IPAddressV6LinkLocal) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


