# LinodeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID representing the Linode Type. | [optional] [readonly] 
**Label** | Pointer to **string** | The Linode Type&#39;s label is for display purposes only.  | [optional] [readonly] 
**Disk** | Pointer to **int32** | The Disk size, in MB, of the Linode Type.  | [optional] [readonly] 
**Class** | Pointer to **string** | The class of the Linode Type. We currently offer five classes of Linodes:    * nanode - Nanode instances are good for low-duty workloads,     where performance isn&#39;t critical. **Note:** As of June 16th, 2020, Nanodes became     1 GB Linodes in the Cloud Manager, however, the API, the CLI, and billing will     continue to refer to these instances as Nanodes.   * standard - Standard Shared instances are good for medium-duty workloads and     are a good mix of performance, resources, and price. **Note:** As of June 16th, 2020,     Standard Linodes in the Cloud Manager became Shared Linodes, however, the API, the CLI, and     billing will continue to refer to these instances as Standard Linodes.   * dedicated - Dedicated CPU instances are good for full-duty workloads     where consistent performance is important.   * gpu - Linodes with dedicated NVIDIA Quadro &amp;reg; RTX 6000 GPUs accelerate highly     specialized applications such as machine learning, AI, and video transcoding.   * highmem - High Memory instances favor RAM over other resources, and can be     good for memory hungry use cases like caching and in-memory databases.     All High Memory plans contain dedicated CPU cores.  | [optional] [readonly] 
**Price** | Pointer to [**LinodeTypePrice**](LinodeTypePrice.md) |  | [optional] 
**Addons** | Pointer to [**LinodeTypeAddons**](LinodeTypeAddons.md) |  | [optional] 
**NetworkOut** | Pointer to **int32** | The Mbits outbound bandwidth allocation.  | [optional] [readonly] 
**Memory** | Pointer to **int32** | Amount of RAM included in this Linode Type.  | [optional] [readonly] 
**Successor** | Pointer to **NullableString** | The Linode Type that a [mutate](/docs/api/linode-instances/#linode-upgrade) will upgrade to for a Linode of this type.  If \&quot;null\&quot;, a Linode of this type may not mutate.  | [optional] [readonly] 
**Transfer** | Pointer to **int32** | The monthly outbound transfer amount, in MB.  | [optional] [readonly] 
**Vcpus** | Pointer to **int32** | The number of VCPU cores this Linode Type offers.  | [optional] [readonly] 
**Gpus** | Pointer to **int32** | The number of GPUs this Linode Type offers.  | [optional] [readonly] 

## Methods

### NewLinodeType

`func NewLinodeType() *LinodeType`

NewLinodeType instantiates a new LinodeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeTypeWithDefaults

`func NewLinodeTypeWithDefaults() *LinodeType`

NewLinodeTypeWithDefaults instantiates a new LinodeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinodeType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinodeType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinodeType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinodeType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *LinodeType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinodeType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinodeType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LinodeType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisk

`func (o *LinodeType) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *LinodeType) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *LinodeType) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *LinodeType) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetClass

`func (o *LinodeType) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *LinodeType) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *LinodeType) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *LinodeType) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetPrice

`func (o *LinodeType) GetPrice() LinodeTypePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LinodeType) GetPriceOk() (*LinodeTypePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LinodeType) SetPrice(v LinodeTypePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LinodeType) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAddons

`func (o *LinodeType) GetAddons() LinodeTypeAddons`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *LinodeType) GetAddonsOk() (*LinodeTypeAddons, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *LinodeType) SetAddons(v LinodeTypeAddons)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *LinodeType) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetNetworkOut

`func (o *LinodeType) GetNetworkOut() int32`

GetNetworkOut returns the NetworkOut field if non-nil, zero value otherwise.

### GetNetworkOutOk

`func (o *LinodeType) GetNetworkOutOk() (*int32, bool)`

GetNetworkOutOk returns a tuple with the NetworkOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOut

`func (o *LinodeType) SetNetworkOut(v int32)`

SetNetworkOut sets NetworkOut field to given value.

### HasNetworkOut

`func (o *LinodeType) HasNetworkOut() bool`

HasNetworkOut returns a boolean if a field has been set.

### GetMemory

`func (o *LinodeType) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *LinodeType) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *LinodeType) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *LinodeType) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSuccessor

`func (o *LinodeType) GetSuccessor() string`

GetSuccessor returns the Successor field if non-nil, zero value otherwise.

### GetSuccessorOk

`func (o *LinodeType) GetSuccessorOk() (*string, bool)`

GetSuccessorOk returns a tuple with the Successor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessor

`func (o *LinodeType) SetSuccessor(v string)`

SetSuccessor sets Successor field to given value.

### HasSuccessor

`func (o *LinodeType) HasSuccessor() bool`

HasSuccessor returns a boolean if a field has been set.

### SetSuccessorNil

`func (o *LinodeType) SetSuccessorNil(b bool)`

 SetSuccessorNil sets the value for Successor to be an explicit nil

### UnsetSuccessor
`func (o *LinodeType) UnsetSuccessor()`

UnsetSuccessor ensures that no value is present for Successor, not even an explicit nil
### GetTransfer

`func (o *LinodeType) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *LinodeType) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *LinodeType) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *LinodeType) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetVcpus

`func (o *LinodeType) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *LinodeType) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *LinodeType) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *LinodeType) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetGpus

`func (o *LinodeType) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *LinodeType) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *LinodeType) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *LinodeType) HasGpus() bool`

HasGpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


