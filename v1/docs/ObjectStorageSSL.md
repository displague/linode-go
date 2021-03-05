# ObjectStorageSSL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Your Base64 encoded and PEM formatted SSL certificate.  | 
**PrivateKey** | **string** | The private key associated with this TLS/SSL certificate.  | 

## Methods

### NewObjectStorageSSL

`func NewObjectStorageSSL(certificate string, privateKey string, ) *ObjectStorageSSL`

NewObjectStorageSSL instantiates a new ObjectStorageSSL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageSSLWithDefaults

`func NewObjectStorageSSLWithDefaults() *ObjectStorageSSL`

NewObjectStorageSSLWithDefaults instantiates a new ObjectStorageSSL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ObjectStorageSSL) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ObjectStorageSSL) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ObjectStorageSSL) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *ObjectStorageSSL) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ObjectStorageSSL) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ObjectStorageSSL) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


