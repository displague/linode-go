# StatsDataAvailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**[]StatsData**](StatsData.md) | CPU usage stats from the last 24 hours. | [optional] 
**Disk** | Pointer to [**[]StatsData**](StatsData.md) | Disk usage stats from the last 24 hours. | [optional] 
**Swap** | Pointer to [**[]StatsData**](StatsData.md) | Swap usage stats from the last 24 hours. | [optional] 
**NetIn** | Pointer to [**[]StatsData**](StatsData.md) | Inbound network traffic stats from the last 24 hours. | [optional] 
**NetOut** | Pointer to [**[]StatsData**](StatsData.md) | Outbound network traffic stats from the last 24 hours. | [optional] 

## Methods

### NewStatsDataAvailable

`func NewStatsDataAvailable() *StatsDataAvailable`

NewStatsDataAvailable instantiates a new StatsDataAvailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsDataAvailableWithDefaults

`func NewStatsDataAvailableWithDefaults() *StatsDataAvailable`

NewStatsDataAvailableWithDefaults instantiates a new StatsDataAvailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *StatsDataAvailable) GetCpu() []StatsData`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *StatsDataAvailable) GetCpuOk() (*[]StatsData, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *StatsDataAvailable) SetCpu(v []StatsData)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *StatsDataAvailable) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *StatsDataAvailable) GetDisk() []StatsData`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *StatsDataAvailable) GetDiskOk() (*[]StatsData, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *StatsDataAvailable) SetDisk(v []StatsData)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *StatsDataAvailable) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetSwap

`func (o *StatsDataAvailable) GetSwap() []StatsData`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *StatsDataAvailable) GetSwapOk() (*[]StatsData, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *StatsDataAvailable) SetSwap(v []StatsData)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *StatsDataAvailable) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetNetIn

`func (o *StatsDataAvailable) GetNetIn() []StatsData`

GetNetIn returns the NetIn field if non-nil, zero value otherwise.

### GetNetInOk

`func (o *StatsDataAvailable) GetNetInOk() (*[]StatsData, bool)`

GetNetInOk returns a tuple with the NetIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIn

`func (o *StatsDataAvailable) SetNetIn(v []StatsData)`

SetNetIn sets NetIn field to given value.

### HasNetIn

`func (o *StatsDataAvailable) HasNetIn() bool`

HasNetIn returns a boolean if a field has been set.

### GetNetOut

`func (o *StatsDataAvailable) GetNetOut() []StatsData`

GetNetOut returns the NetOut field if non-nil, zero value otherwise.

### GetNetOutOk

`func (o *StatsDataAvailable) GetNetOutOk() (*[]StatsData, bool)`

GetNetOutOk returns a tuple with the NetOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOut

`func (o *StatsDataAvailable) SetNetOut(v []StatsData)`

SetNetOut sets NetOut field to given value.

### HasNetOut

`func (o *StatsDataAvailable) HasNetOut() bool`

HasNetOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


