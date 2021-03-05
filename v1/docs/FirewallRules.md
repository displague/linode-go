# FirewallRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | [**[]FirewallRuleConfig**](FirewallRuleConfig.md) |  | 
**Outbound** | Pointer to [**[]FirewallRuleConfig**](FirewallRuleConfig.md) |  | [optional] 

## Methods

### NewFirewallRules

`func NewFirewallRules(inbound []FirewallRuleConfig, ) *FirewallRules`

NewFirewallRules instantiates a new FirewallRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRulesWithDefaults

`func NewFirewallRulesWithDefaults() *FirewallRules`

NewFirewallRulesWithDefaults instantiates a new FirewallRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *FirewallRules) GetInbound() []FirewallRuleConfig`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *FirewallRules) GetInboundOk() (*[]FirewallRuleConfig, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *FirewallRules) SetInbound(v []FirewallRuleConfig)`

SetInbound sets Inbound field to given value.


### GetOutbound

`func (o *FirewallRules) GetOutbound() []FirewallRuleConfig`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *FirewallRules) GetOutboundOk() (*[]FirewallRuleConfig, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *FirewallRules) SetOutbound(v []FirewallRuleConfig)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *FirewallRules) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


