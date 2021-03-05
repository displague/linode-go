# ObjectStorageCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for this cluster. | [optional] 
**Domain** | Pointer to **string** | The base URL for this cluster, used for connecting with third-party clients. | [optional] 
**Status** | Pointer to **string** | This cluster&#39;s status. | [optional] 
**Region** | Pointer to **string** | The region where this cluster is located. | [optional] 
**StaticSiteDomain** | Pointer to **string** | The base URL for this cluster used when hosting static sites. | [optional] 

## Methods

### NewObjectStorageCluster

`func NewObjectStorageCluster() *ObjectStorageCluster`

NewObjectStorageCluster instantiates a new ObjectStorageCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageClusterWithDefaults

`func NewObjectStorageClusterWithDefaults() *ObjectStorageCluster`

NewObjectStorageClusterWithDefaults instantiates a new ObjectStorageCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectStorageCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *ObjectStorageCluster) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ObjectStorageCluster) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ObjectStorageCluster) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ObjectStorageCluster) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegion

`func (o *ObjectStorageCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ObjectStorageCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ObjectStorageCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ObjectStorageCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStaticSiteDomain

`func (o *ObjectStorageCluster) GetStaticSiteDomain() string`

GetStaticSiteDomain returns the StaticSiteDomain field if non-nil, zero value otherwise.

### GetStaticSiteDomainOk

`func (o *ObjectStorageCluster) GetStaticSiteDomainOk() (*string, bool)`

GetStaticSiteDomainOk returns a tuple with the StaticSiteDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticSiteDomain

`func (o *ObjectStorageCluster) SetStaticSiteDomain(v string)`

SetStaticSiteDomain sets StaticSiteDomain field to given value.

### HasStaticSiteDomain

`func (o *ObjectStorageCluster) HasStaticSiteDomain() bool`

HasStaticSiteDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


