# LinodeStatsNetv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[][]float32** | Input stats for IPv6, measured in bits/s (bits/second). | [optional] 
**Out** | Pointer to **[][]float32** | Output stats for IPv6, measured in bits/s (bits/second). | [optional] 
**PrivateIn** | Pointer to **[][]float32** | Private IPv6 input statistics, measured in bits/s (bits/second). | [optional] 
**PrivateOut** | Pointer to **[][]float32** | Private IPv6 output statistics, measured in bits/s (bits/second). | [optional] 

## Methods

### NewLinodeStatsNetv6

`func NewLinodeStatsNetv6() *LinodeStatsNetv6`

NewLinodeStatsNetv6 instantiates a new LinodeStatsNetv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeStatsNetv6WithDefaults

`func NewLinodeStatsNetv6WithDefaults() *LinodeStatsNetv6`

NewLinodeStatsNetv6WithDefaults instantiates a new LinodeStatsNetv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *LinodeStatsNetv6) GetIn() [][]float32`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *LinodeStatsNetv6) GetInOk() (*[][]float32, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *LinodeStatsNetv6) SetIn(v [][]float32)`

SetIn sets In field to given value.

### HasIn

`func (o *LinodeStatsNetv6) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetOut

`func (o *LinodeStatsNetv6) GetOut() [][]float32`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *LinodeStatsNetv6) GetOutOk() (*[][]float32, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *LinodeStatsNetv6) SetOut(v [][]float32)`

SetOut sets Out field to given value.

### HasOut

`func (o *LinodeStatsNetv6) HasOut() bool`

HasOut returns a boolean if a field has been set.

### GetPrivateIn

`func (o *LinodeStatsNetv6) GetPrivateIn() [][]float32`

GetPrivateIn returns the PrivateIn field if non-nil, zero value otherwise.

### GetPrivateInOk

`func (o *LinodeStatsNetv6) GetPrivateInOk() (*[][]float32, bool)`

GetPrivateInOk returns a tuple with the PrivateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIn

`func (o *LinodeStatsNetv6) SetPrivateIn(v [][]float32)`

SetPrivateIn sets PrivateIn field to given value.

### HasPrivateIn

`func (o *LinodeStatsNetv6) HasPrivateIn() bool`

HasPrivateIn returns a boolean if a field has been set.

### GetPrivateOut

`func (o *LinodeStatsNetv6) GetPrivateOut() [][]float32`

GetPrivateOut returns the PrivateOut field if non-nil, zero value otherwise.

### GetPrivateOutOk

`func (o *LinodeStatsNetv6) GetPrivateOutOk() (*[][]float32, bool)`

GetPrivateOutOk returns a tuple with the PrivateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOut

`func (o *LinodeStatsNetv6) SetPrivateOut(v [][]float32)`

SetPrivateOut sets PrivateOut field to given value.

### HasPrivateOut

`func (o *LinodeStatsNetv6) HasPrivateOut() bool`

HasPrivateOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


