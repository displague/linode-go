# NodeBalancerConfigNodesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Up** | Pointer to **int32** | The number of backends considered to be \&quot;UP\&quot; and healthy, and that are serving requests.  | [optional] [readonly] 
**Down** | Pointer to **int32** | The number of backends considered to be \&quot;DOWN\&quot; and unhealthy.  These are not in rotation, and not serving requests.  | [optional] [readonly] 

## Methods

### NewNodeBalancerConfigNodesStatus

`func NewNodeBalancerConfigNodesStatus() *NodeBalancerConfigNodesStatus`

NewNodeBalancerConfigNodesStatus instantiates a new NodeBalancerConfigNodesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerConfigNodesStatusWithDefaults

`func NewNodeBalancerConfigNodesStatusWithDefaults() *NodeBalancerConfigNodesStatus`

NewNodeBalancerConfigNodesStatusWithDefaults instantiates a new NodeBalancerConfigNodesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUp

`func (o *NodeBalancerConfigNodesStatus) GetUp() int32`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *NodeBalancerConfigNodesStatus) GetUpOk() (*int32, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *NodeBalancerConfigNodesStatus) SetUp(v int32)`

SetUp sets Up field to given value.

### HasUp

`func (o *NodeBalancerConfigNodesStatus) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetDown

`func (o *NodeBalancerConfigNodesStatus) GetDown() int32`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *NodeBalancerConfigNodesStatus) GetDownOk() (*int32, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *NodeBalancerConfigNodesStatus) SetDown(v int32)`

SetDown sets Down field to given value.

### HasDown

`func (o *NodeBalancerConfigNodesStatus) HasDown() bool`

HasDown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


