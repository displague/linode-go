# LinodeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **[][]float32** | Percentage of CPU used.  | [optional] 
**Io** | Pointer to [**LinodeStatsIo**](LinodeStatsIo.md) |  | [optional] 
**Netv4** | Pointer to [**LinodeStatsNetv4**](LinodeStatsNetv4.md) |  | [optional] 
**Netv6** | Pointer to [**LinodeStatsNetv6**](LinodeStatsNetv6.md) |  | [optional] 
**Title** | Pointer to **string** | The title for this data set. | [optional] 

## Methods

### NewLinodeStats

`func NewLinodeStats() *LinodeStats`

NewLinodeStats instantiates a new LinodeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeStatsWithDefaults

`func NewLinodeStatsWithDefaults() *LinodeStats`

NewLinodeStatsWithDefaults instantiates a new LinodeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *LinodeStats) GetCpu() [][]float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *LinodeStats) GetCpuOk() (*[][]float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *LinodeStats) SetCpu(v [][]float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *LinodeStats) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetIo

`func (o *LinodeStats) GetIo() LinodeStatsIo`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *LinodeStats) GetIoOk() (*LinodeStatsIo, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *LinodeStats) SetIo(v LinodeStatsIo)`

SetIo sets Io field to given value.

### HasIo

`func (o *LinodeStats) HasIo() bool`

HasIo returns a boolean if a field has been set.

### GetNetv4

`func (o *LinodeStats) GetNetv4() LinodeStatsNetv4`

GetNetv4 returns the Netv4 field if non-nil, zero value otherwise.

### GetNetv4Ok

`func (o *LinodeStats) GetNetv4Ok() (*LinodeStatsNetv4, bool)`

GetNetv4Ok returns a tuple with the Netv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetv4

`func (o *LinodeStats) SetNetv4(v LinodeStatsNetv4)`

SetNetv4 sets Netv4 field to given value.

### HasNetv4

`func (o *LinodeStats) HasNetv4() bool`

HasNetv4 returns a boolean if a field has been set.

### GetNetv6

`func (o *LinodeStats) GetNetv6() LinodeStatsNetv6`

GetNetv6 returns the Netv6 field if non-nil, zero value otherwise.

### GetNetv6Ok

`func (o *LinodeStats) GetNetv6Ok() (*LinodeStatsNetv6, bool)`

GetNetv6Ok returns a tuple with the Netv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetv6

`func (o *LinodeStats) SetNetv6(v LinodeStatsNetv6)`

SetNetv6 sets Netv6 field to given value.

### HasNetv6

`func (o *LinodeStats) HasNetv6() bool`

HasNetv6 returns a boolean if a field has been set.

### GetTitle

`func (o *LinodeStats) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinodeStats) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinodeStats) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LinodeStats) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


