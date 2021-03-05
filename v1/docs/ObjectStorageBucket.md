# ObjectStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this bucket was created. | [optional] 
**Cluster** | Pointer to **string** | The ID of the Object Storage Cluster this bucket is in. | [optional] 
**Label** | Pointer to **string** | The name of this bucket. | [optional] 
**Hostname** | Pointer to **string** | The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public.  | [optional] 
**Size** | Pointer to **int32** | The size of the bucket in bytes. | [optional] 
**Objects** | Pointer to **int32** | The number of objects stored in this bucket.  | [optional] 

## Methods

### NewObjectStorageBucket

`func NewObjectStorageBucket() *ObjectStorageBucket`

NewObjectStorageBucket instantiates a new ObjectStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketWithDefaults

`func NewObjectStorageBucketWithDefaults() *ObjectStorageBucket`

NewObjectStorageBucketWithDefaults instantiates a new ObjectStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ObjectStorageBucket) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ObjectStorageBucket) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ObjectStorageBucket) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ObjectStorageBucket) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageBucket) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageBucket) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageBucket) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageBucket) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetLabel

`func (o *ObjectStorageBucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectStorageBucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectStorageBucket) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ObjectStorageBucket) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHostname

`func (o *ObjectStorageBucket) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ObjectStorageBucket) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ObjectStorageBucket) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ObjectStorageBucket) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetSize

`func (o *ObjectStorageBucket) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ObjectStorageBucket) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ObjectStorageBucket) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ObjectStorageBucket) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetObjects

`func (o *ObjectStorageBucket) GetObjects() int32`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectStorageBucket) GetObjectsOk() (*int32, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectStorageBucket) SetObjects(v int32)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectStorageBucket) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


