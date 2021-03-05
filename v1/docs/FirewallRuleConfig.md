# FirewallRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The type of network traffic to allow.  | 
**Ports** | Pointer to **string** | A string representing the port or ports on which traffic will be allowed:  - The string may be a single port, a range of ports, or a comma-separated list of single ports and port ranges. A space is permitted following each comma. - A range of ports is inclusive of the start and end values for the range. The end value of the range must be greater than the start value. - Ports must be within 1 and 65535, and may not contain any leading zeroes. For example, port \&quot;080\&quot; is not allowed. - Ports may not be specified if a rule&#39;s protocol is &#x60;ICMP&#x60;. At least one port must be specified if a rule&#39;s protocol is &#x60;TCP&#x60; or &#x60;UDP&#x60;. - The ports string can have up to 15 *pieces*, where a single port is treated as one piece, and a port range is treated as two pieces. For example, the string \&quot;22-24, 80, 443\&quot; has four pieces.  | [optional] 
**Addresses** | Pointer to [**FirewallRuleConfigAddresses**](FirewallRuleConfigAddresses.md) |  | [optional] 

## Methods

### NewFirewallRuleConfig

`func NewFirewallRuleConfig(protocol string, ) *FirewallRuleConfig`

NewFirewallRuleConfig instantiates a new FirewallRuleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleConfigWithDefaults

`func NewFirewallRuleConfigWithDefaults() *FirewallRuleConfig`

NewFirewallRuleConfigWithDefaults instantiates a new FirewallRuleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *FirewallRuleConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPorts

`func (o *FirewallRuleConfig) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FirewallRuleConfig) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FirewallRuleConfig) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FirewallRuleConfig) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetAddresses

`func (o *FirewallRuleConfig) GetAddresses() FirewallRuleConfigAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *FirewallRuleConfig) GetAddressesOk() (*FirewallRuleConfigAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *FirewallRuleConfig) SetAddresses(v FirewallRuleConfigAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *FirewallRuleConfig) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


