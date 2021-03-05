# InlineObject13

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The name for this bucket. Must be unique in the cluster you are creating the bucket in, or an error will be returned. Labels will be reserved only for the cluster that active buckets are created and stored in. If you want to reserve this bucket&#39;s label in another cluster, you must create a new bucket with the same label in the new cluster.  | 
**Cluster** | **string** | The ID of the Object Storage Cluster where this bucket should be created.  | 
**CorsEnabled** | Pointer to **bool** | If true, the bucket will be created with CORS enabled for all origins. For more fine-grained controls of CORS, use the S3 API directly.  | [optional] [default to false]
**Acl** | Pointer to **string** | The Access Control Level of the bucket using a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly.  | [optional] [default to "private"]

## Methods

### NewInlineObject13

`func NewInlineObject13(label string, cluster string, ) *InlineObject13`

NewInlineObject13 instantiates a new InlineObject13 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject13WithDefaults

`func NewInlineObject13WithDefaults() *InlineObject13`

NewInlineObject13WithDefaults instantiates a new InlineObject13 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InlineObject13) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineObject13) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineObject13) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCluster

`func (o *InlineObject13) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineObject13) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineObject13) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetCorsEnabled

`func (o *InlineObject13) GetCorsEnabled() bool`

GetCorsEnabled returns the CorsEnabled field if non-nil, zero value otherwise.

### GetCorsEnabledOk

`func (o *InlineObject13) GetCorsEnabledOk() (*bool, bool)`

GetCorsEnabledOk returns a tuple with the CorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsEnabled

`func (o *InlineObject13) SetCorsEnabled(v bool)`

SetCorsEnabled sets CorsEnabled field to given value.

### HasCorsEnabled

`func (o *InlineObject13) HasCorsEnabled() bool`

HasCorsEnabled returns a boolean if a field has been set.

### GetAcl

`func (o *InlineObject13) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *InlineObject13) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *InlineObject13) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *InlineObject13) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


