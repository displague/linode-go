# NodeBalancerStatsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to **[]float32** | An array of key/value pairs representing unix timestamp and reading for connections to this NodeBalancer.  | [optional] 
**Traffic** | Pointer to [**NodeBalancerStatsDataTraffic**](NodeBalancerStatsDataTraffic.md) |  | [optional] 

## Methods

### NewNodeBalancerStatsData

`func NewNodeBalancerStatsData() *NodeBalancerStatsData`

NewNodeBalancerStatsData instantiates a new NodeBalancerStatsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerStatsDataWithDefaults

`func NewNodeBalancerStatsDataWithDefaults() *NodeBalancerStatsData`

NewNodeBalancerStatsDataWithDefaults instantiates a new NodeBalancerStatsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *NodeBalancerStatsData) GetConnections() []float32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *NodeBalancerStatsData) GetConnectionsOk() (*[]float32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *NodeBalancerStatsData) SetConnections(v []float32)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *NodeBalancerStatsData) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetTraffic

`func (o *NodeBalancerStatsData) GetTraffic() NodeBalancerStatsDataTraffic`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *NodeBalancerStatsData) GetTrafficOk() (*NodeBalancerStatsDataTraffic, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *NodeBalancerStatsData) SetTraffic(v NodeBalancerStatsDataTraffic)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *NodeBalancerStatsData) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


