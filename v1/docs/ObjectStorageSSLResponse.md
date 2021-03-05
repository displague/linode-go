# ObjectStorageSSLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssl** | Pointer to **bool** | A boolean indicating if this Bucket has a corresponding TLS/SSL certificate that was uploaded by an Account user.  | [optional] [readonly] 

## Methods

### NewObjectStorageSSLResponse

`func NewObjectStorageSSLResponse() *ObjectStorageSSLResponse`

NewObjectStorageSSLResponse instantiates a new ObjectStorageSSLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageSSLResponseWithDefaults

`func NewObjectStorageSSLResponseWithDefaults() *ObjectStorageSSLResponse`

NewObjectStorageSSLResponseWithDefaults instantiates a new ObjectStorageSSLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsl

`func (o *ObjectStorageSSLResponse) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *ObjectStorageSSLResponse) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *ObjectStorageSSLResponse) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *ObjectStorageSSLResponse) HasSsl() bool`

HasSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


