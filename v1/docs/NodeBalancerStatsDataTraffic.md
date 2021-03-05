# NodeBalancerStatsDataTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]float32** | An array of key/value pairs representing unix timestamp and reading for inbound traffic.  | [optional] 
**Out** | Pointer to **[]float32** | An array of key/value pairs representing unix timestamp and reading for outbound traffic.  | [optional] 

## Methods

### NewNodeBalancerStatsDataTraffic

`func NewNodeBalancerStatsDataTraffic() *NodeBalancerStatsDataTraffic`

NewNodeBalancerStatsDataTraffic instantiates a new NodeBalancerStatsDataTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerStatsDataTrafficWithDefaults

`func NewNodeBalancerStatsDataTrafficWithDefaults() *NodeBalancerStatsDataTraffic`

NewNodeBalancerStatsDataTrafficWithDefaults instantiates a new NodeBalancerStatsDataTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *NodeBalancerStatsDataTraffic) GetIn() []float32`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *NodeBalancerStatsDataTraffic) GetInOk() (*[]float32, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *NodeBalancerStatsDataTraffic) SetIn(v []float32)`

SetIn sets In field to given value.

### HasIn

`func (o *NodeBalancerStatsDataTraffic) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetOut

`func (o *NodeBalancerStatsDataTraffic) GetOut() []float32`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *NodeBalancerStatsDataTraffic) GetOutOk() (*[]float32, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *NodeBalancerStatsDataTraffic) SetOut(v []float32)`

SetOut sets Out field to given value.

### HasOut

`func (o *NodeBalancerStatsDataTraffic) HasOut() bool`

HasOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


