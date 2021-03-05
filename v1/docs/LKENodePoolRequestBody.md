# LKENodePoolRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A Linode Type for all of the nodes in the Node Pool. | [optional] 
**Count** | Pointer to **int32** | The number of nodes in the Node Pool. | [optional] 
**Disks** | Pointer to [**[]LKENodePoolRequestBodyDisks**](LKENodePoolRequestBodyDisks.md) | **Note**: This field should be omitted except for special use cases. The disks specified here are partitions in *addition* to the primary partition and reduce the size of the primary partition, which can lead to stability problems for the Node.  This Node Pool&#39;s custom disk layout. Each item in this array will create a new disk partition for each node in this Node Pool.    * The custom disk layout is applied to each node in this Node Pool.   * The maximum number of custom disk partitions that can be configured is 7.   * Once the requested disk paritions are allocated, the remaining disk space is allocated to the node&#39;s boot disk.   * A Node Pool&#39;s custom disk layout is immutable over the lifetime of the Node Pool.  | [optional] 

## Methods

### NewLKENodePoolRequestBody

`func NewLKENodePoolRequestBody() *LKENodePoolRequestBody`

NewLKENodePoolRequestBody instantiates a new LKENodePoolRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLKENodePoolRequestBodyWithDefaults

`func NewLKENodePoolRequestBodyWithDefaults() *LKENodePoolRequestBody`

NewLKENodePoolRequestBodyWithDefaults instantiates a new LKENodePoolRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LKENodePoolRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LKENodePoolRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LKENodePoolRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LKENodePoolRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCount

`func (o *LKENodePoolRequestBody) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LKENodePoolRequestBody) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LKENodePoolRequestBody) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LKENodePoolRequestBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDisks

`func (o *LKENodePoolRequestBody) GetDisks() []LKENodePoolRequestBodyDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *LKENodePoolRequestBody) GetDisksOk() (*[]LKENodePoolRequestBodyDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *LKENodePoolRequestBody) SetDisks(v []LKENodePoolRequestBodyDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *LKENodePoolRequestBody) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


