# ObjectStorageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this object or prefix.  | [optional] 
**Etag** | Pointer to **string** | An MD-5 hash of the object. &#x60;null&#x60; if this object represents a prefix.  | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time this object was last modified. &#x60;null&#x60; if this object represents a prefix.  | [optional] 
**Owner** | Pointer to **string** | The owner of this object, as a UUID. &#x60;null&#x60; if this object represents a prefix.  | [optional] 
**Size** | Pointer to **int32** | The size of this object, in bytes. &#x60;null&#x60; if this object represents a prefix.  | [optional] 
**NextMarker** | Pointer to **NullableString** | Returns the value you should pass to the &#x60;marker&#x60; query parameter to get the next page of objects. If there is no next page, &#x60;null&#x60; will be returned.  | [optional] 
**IsTruncated** | Pointer to **bool** | Designates if there is another page of bucket objects. | [optional] 

## Methods

### NewObjectStorageObject

`func NewObjectStorageObject() *ObjectStorageObject`

NewObjectStorageObject instantiates a new ObjectStorageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageObjectWithDefaults

`func NewObjectStorageObjectWithDefaults() *ObjectStorageObject`

NewObjectStorageObjectWithDefaults instantiates a new ObjectStorageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectStorageObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEtag

`func (o *ObjectStorageObject) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ObjectStorageObject) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ObjectStorageObject) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ObjectStorageObject) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetLastModified

`func (o *ObjectStorageObject) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ObjectStorageObject) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ObjectStorageObject) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ObjectStorageObject) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetOwner

`func (o *ObjectStorageObject) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ObjectStorageObject) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ObjectStorageObject) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ObjectStorageObject) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSize

`func (o *ObjectStorageObject) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ObjectStorageObject) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ObjectStorageObject) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ObjectStorageObject) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetNextMarker

`func (o *ObjectStorageObject) GetNextMarker() string`

GetNextMarker returns the NextMarker field if non-nil, zero value otherwise.

### GetNextMarkerOk

`func (o *ObjectStorageObject) GetNextMarkerOk() (*string, bool)`

GetNextMarkerOk returns a tuple with the NextMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMarker

`func (o *ObjectStorageObject) SetNextMarker(v string)`

SetNextMarker sets NextMarker field to given value.

### HasNextMarker

`func (o *ObjectStorageObject) HasNextMarker() bool`

HasNextMarker returns a boolean if a field has been set.

### SetNextMarkerNil

`func (o *ObjectStorageObject) SetNextMarkerNil(b bool)`

 SetNextMarkerNil sets the value for NextMarker to be an explicit nil

### UnsetNextMarker
`func (o *ObjectStorageObject) UnsetNextMarker()`

UnsetNextMarker ensures that no value is present for NextMarker, not even an explicit nil
### GetIsTruncated

`func (o *ObjectStorageObject) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ObjectStorageObject) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ObjectStorageObject) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *ObjectStorageObject) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


