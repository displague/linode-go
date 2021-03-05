# ObjectStorageKeyBucketAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | The Object Storage cluster where a bucket to which the key is granting access is hosted. | [optional] 
**BucketName** | Pointer to **string** | The unique label of the bucket to which the key will grant limited access. | [optional] 
**Permissions** | Pointer to **string** | This Limited Access Key&#39;s permissions for the selected bucket. | [optional] 

## Methods

### NewObjectStorageKeyBucketAccess

`func NewObjectStorageKeyBucketAccess() *ObjectStorageKeyBucketAccess`

NewObjectStorageKeyBucketAccess instantiates a new ObjectStorageKeyBucketAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageKeyBucketAccessWithDefaults

`func NewObjectStorageKeyBucketAccessWithDefaults() *ObjectStorageKeyBucketAccess`

NewObjectStorageKeyBucketAccessWithDefaults instantiates a new ObjectStorageKeyBucketAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ObjectStorageKeyBucketAccess) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageKeyBucketAccess) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageKeyBucketAccess) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageKeyBucketAccess) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetBucketName

`func (o *ObjectStorageKeyBucketAccess) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ObjectStorageKeyBucketAccess) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ObjectStorageKeyBucketAccess) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ObjectStorageKeyBucketAccess) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPermissions

`func (o *ObjectStorageKeyBucketAccess) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ObjectStorageKeyBucketAccess) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ObjectStorageKeyBucketAccess) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ObjectStorageKeyBucketAccess) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


