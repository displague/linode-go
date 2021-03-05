# NodeBalancerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This node&#39;s unique ID. | [optional] [readonly] 
**Address** | Pointer to **string** | The private IP Address where this backend can be reached. This _must_ be a private IP address.  | [optional] 
**Label** | Pointer to **string** | The label for this node.  This is for display purposes only.  | [optional] 
**Status** | Pointer to **string** | The current status of this node, based on the configured checks of its NodeBalancer Config.  | [optional] [readonly] 
**Weight** | Pointer to **int32** | Used when picking a backend to serve a request and is not pinned to a single backend yet.  Nodes with a higher weight will receive more traffic.  | [optional] 
**Mode** | Pointer to **string** | The mode this NodeBalancer should use when sending traffic to this backend. * If set to &#x60;accept&#x60; this backend is accepting traffic. * If set to &#x60;reject&#x60; this backend will not receive traffic. * If set to &#x60;drain&#x60; this backend will not receive _new_ traffic, but connections already   pinned to it will continue to be routed to it.  * If set to &#x60;backup&#x60;, this backend will only receive traffic if all &#x60;accept&#x60; nodes   are down.  | [optional] 
**ConfigId** | Pointer to **int32** | The NodeBalancer Config ID that this Node belongs to.  | [optional] [readonly] 
**NodebalancerId** | Pointer to **int32** | The NodeBalancer ID that this Node belongs to.  | [optional] [readonly] 

## Methods

### NewNodeBalancerNode

`func NewNodeBalancerNode() *NodeBalancerNode`

NewNodeBalancerNode instantiates a new NodeBalancerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerNodeWithDefaults

`func NewNodeBalancerNodeWithDefaults() *NodeBalancerNode`

NewNodeBalancerNodeWithDefaults instantiates a new NodeBalancerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeBalancerNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeBalancerNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeBalancerNode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeBalancerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeBalancerNode) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeBalancerNode) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeBalancerNode) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeBalancerNode) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLabel

`func (o *NodeBalancerNode) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NodeBalancerNode) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NodeBalancerNode) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *NodeBalancerNode) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *NodeBalancerNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeBalancerNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeBalancerNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeBalancerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWeight

`func (o *NodeBalancerNode) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *NodeBalancerNode) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *NodeBalancerNode) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *NodeBalancerNode) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMode

`func (o *NodeBalancerNode) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NodeBalancerNode) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NodeBalancerNode) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NodeBalancerNode) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetConfigId

`func (o *NodeBalancerNode) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *NodeBalancerNode) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *NodeBalancerNode) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *NodeBalancerNode) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *NodeBalancerNode) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *NodeBalancerNode) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *NodeBalancerNode) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *NodeBalancerNode) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


