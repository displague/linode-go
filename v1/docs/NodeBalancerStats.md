# NodeBalancerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NodeBalancerStatsData**](NodeBalancerStatsData.md) |  | [optional] 
**Title** | Pointer to **string** | The title for the statistics generated in this response.  | [optional] 

## Methods

### NewNodeBalancerStats

`func NewNodeBalancerStats() *NodeBalancerStats`

NewNodeBalancerStats instantiates a new NodeBalancerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerStatsWithDefaults

`func NewNodeBalancerStatsWithDefaults() *NodeBalancerStats`

NewNodeBalancerStatsWithDefaults instantiates a new NodeBalancerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NodeBalancerStats) GetData() NodeBalancerStatsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NodeBalancerStats) GetDataOk() (*NodeBalancerStatsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NodeBalancerStats) SetData(v NodeBalancerStatsData)`

SetData sets Data field to given value.

### HasData

`func (o *NodeBalancerStats) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTitle

`func (o *NodeBalancerStats) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NodeBalancerStats) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NodeBalancerStats) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NodeBalancerStats) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


