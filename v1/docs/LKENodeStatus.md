# LKENodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Node&#39;s ID.  | [optional] 
**InstanceId** | Pointer to **string** | The Linode&#39;s ID. When no Linode is currently provisioned for this Node, this will be null.  | [optional] 
**Status** | Pointer to **string** | The Node&#39;s status as it pertains to being a Kubernetes node.  | [optional] 

## Methods

### NewLKENodeStatus

`func NewLKENodeStatus() *LKENodeStatus`

NewLKENodeStatus instantiates a new LKENodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLKENodeStatusWithDefaults

`func NewLKENodeStatusWithDefaults() *LKENodeStatus`

NewLKENodeStatusWithDefaults instantiates a new LKENodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LKENodeStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LKENodeStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LKENodeStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LKENodeStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *LKENodeStatus) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *LKENodeStatus) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *LKENodeStatus) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *LKENodeStatus) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *LKENodeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LKENodeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LKENodeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LKENodeStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


