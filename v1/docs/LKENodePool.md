# LKENodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Node Pool&#39;s unique ID.  | [optional] 
**Nodes** | Pointer to [**[]LKENodeStatus**](LKENodeStatus.md) | Status information for the Nodes which are members of this Node Pool. If a Linode has not been provisioned for a given Node slot, the instance_id will be returned as null.  | [optional] 
**Type** | Pointer to **string** | A Linode Type for all of the nodes in the Node Pool. | [optional] 
**Count** | Pointer to **int32** | The number of nodes in the Node Pool. | [optional] 
**Disks** | Pointer to [**[]LKENodePoolRequestBodyDisks**](LKENodePoolRequestBodyDisks.md) | **Note**: This field should be omitted except for special use cases. The disks specified here are partitions in *addition* to the primary partition and reduce the size of the primary partition, which can lead to stability problems for the Node.  This Node Pool&#39;s custom disk layout. Each item in this array will create a new disk partition for each node in this Node Pool.    * The custom disk layout is applied to each node in this Node Pool.   * The maximum number of custom disk partitions that can be configured is 7.   * Once the requested disk paritions are allocated, the remaining disk space is allocated to the node&#39;s boot disk.   * A Node Pool&#39;s custom disk layout is immutable over the lifetime of the Node Pool.  | [optional] 

## Methods

### NewLKENodePool

`func NewLKENodePool() *LKENodePool`

NewLKENodePool instantiates a new LKENodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLKENodePoolWithDefaults

`func NewLKENodePoolWithDefaults() *LKENodePool`

NewLKENodePoolWithDefaults instantiates a new LKENodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LKENodePool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LKENodePool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LKENodePool) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LKENodePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodes

`func (o *LKENodePool) GetNodes() []LKENodeStatus`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *LKENodePool) GetNodesOk() (*[]LKENodeStatus, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *LKENodePool) SetNodes(v []LKENodeStatus)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *LKENodePool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetType

`func (o *LKENodePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LKENodePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LKENodePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LKENodePool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCount

`func (o *LKENodePool) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LKENodePool) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LKENodePool) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LKENodePool) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDisks

`func (o *LKENodePool) GetDisks() []LKENodePoolRequestBodyDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *LKENodePool) GetDisksOk() (*[]LKENodePoolRequestBodyDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *LKENodePool) SetDisks(v []LKENodePoolRequestBodyDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *LKENodePool) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


