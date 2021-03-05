# ObjectStorageKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This keypair&#39;s unique ID | [optional] [readonly] 
**Label** | Pointer to **string** | The label given to this key. For display purposes only. | [optional] 
**AccessKey** | Pointer to **string** | This keypair&#39;s access key. This is not secret. | [optional] [readonly] 
**SecretKey** | Pointer to **string** | This keypair&#39;s secret key. Only returned on key creation. | [optional] [readonly] 
**Limited** | Pointer to **bool** | Whether or not this key is a limited access key. Will return &#x60;false&#x60; if this key grants full access to all buckets on the user&#39;s account. | [optional] [readonly] 
**BucketAccess** | Pointer to [**[]ObjectStorageKeyBucketAccess**](ObjectStorageKeyBucketAccess.md) | Defines this key as a Limited Access Key. Limited Access Keys restrict this Object Storage key&#39;s access to only the bucket(s) declared in this array and define their bucket-level permissions.     Limited Access Keys can:    * [list all buckets](/docs/api/object-storage/#object-storage-buckets-list) available on this Account, but cannot perform any actions on a bucket unless it has access to the bucket.     * [create new buckets](/docs/api/object-storage/#object-storage-bucket-create), but do not have any access to the buckets it creates, unless explicitly given access to them.     **Note:** You can create an Object Storage Limited Access Key without access to any buckets.   This is achieved by sending a request with an empty &#x60;bucket_access&#x60; array.     **Note:** If this field is omitted, a regular unlimited access key is issued.  | [optional] 

## Methods

### NewObjectStorageKey

`func NewObjectStorageKey() *ObjectStorageKey`

NewObjectStorageKey instantiates a new ObjectStorageKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageKeyWithDefaults

`func NewObjectStorageKeyWithDefaults() *ObjectStorageKey`

NewObjectStorageKeyWithDefaults instantiates a new ObjectStorageKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectStorageKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ObjectStorageKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectStorageKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectStorageKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ObjectStorageKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAccessKey

`func (o *ObjectStorageKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ObjectStorageKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ObjectStorageKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ObjectStorageKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ObjectStorageKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ObjectStorageKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ObjectStorageKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ObjectStorageKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetLimited

`func (o *ObjectStorageKey) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *ObjectStorageKey) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *ObjectStorageKey) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *ObjectStorageKey) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetBucketAccess

`func (o *ObjectStorageKey) GetBucketAccess() []ObjectStorageKeyBucketAccess`

GetBucketAccess returns the BucketAccess field if non-nil, zero value otherwise.

### GetBucketAccessOk

`func (o *ObjectStorageKey) GetBucketAccessOk() (*[]ObjectStorageKeyBucketAccess, bool)`

GetBucketAccessOk returns a tuple with the BucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccess

`func (o *ObjectStorageKey) SetBucketAccess(v []ObjectStorageKeyBucketAccess)`

SetBucketAccess sets BucketAccess field to given value.

### HasBucketAccess

`func (o *ObjectStorageKey) HasBucketAccess() bool`

HasBucketAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


