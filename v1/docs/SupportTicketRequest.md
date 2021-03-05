# SupportTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The full details of the issue or question.  | 
**DomainId** | Pointer to **int32** | The ID of the Domain this ticket is regarding, if relevant.  | [optional] 
**LinodeId** | Pointer to **int32** | The ID of the Linode this ticket is regarding, if relevant.  | [optional] 
**LongviewclientId** | Pointer to **int32** | The ID of the Longview client this ticket is regarding, if relevant.  | [optional] 
**NodebalancerId** | Pointer to **int32** | The ID of the NodeBalancer this ticket is regarding, if relevant.  | [optional] 
**Summary** | **string** | The summary or title for this SupportTicket.  | 
**ManagedIssue** | Pointer to **bool** | Designates if this ticket is related to a [Managed service](https://www.linode.com/products/managed/). If &#x60;true&#x60;, the following constraints will apply: * No ID attributes (i.e. &#x60;linode_id&#x60;, &#x60;domain_id&#x60;, etc.) should be provided with this request. * Your account must have a [Managed service enabled](/docs/api/managed/#managed-service-enable).  | [optional] 
**VolumeId** | Pointer to **int32** | The ID of the Volume this ticket is regarding, if relevant.  | [optional] 

## Methods

### NewSupportTicketRequest

`func NewSupportTicketRequest(description string, summary string, ) *SupportTicketRequest`

NewSupportTicketRequest instantiates a new SupportTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketRequestWithDefaults

`func NewSupportTicketRequestWithDefaults() *SupportTicketRequest`

NewSupportTicketRequestWithDefaults instantiates a new SupportTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SupportTicketRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDomainId

`func (o *SupportTicketRequest) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *SupportTicketRequest) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *SupportTicketRequest) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *SupportTicketRequest) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetLinodeId

`func (o *SupportTicketRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *SupportTicketRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *SupportTicketRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *SupportTicketRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetLongviewclientId

`func (o *SupportTicketRequest) GetLongviewclientId() int32`

GetLongviewclientId returns the LongviewclientId field if non-nil, zero value otherwise.

### GetLongviewclientIdOk

`func (o *SupportTicketRequest) GetLongviewclientIdOk() (*int32, bool)`

GetLongviewclientIdOk returns a tuple with the LongviewclientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewclientId

`func (o *SupportTicketRequest) SetLongviewclientId(v int32)`

SetLongviewclientId sets LongviewclientId field to given value.

### HasLongviewclientId

`func (o *SupportTicketRequest) HasLongviewclientId() bool`

HasLongviewclientId returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *SupportTicketRequest) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *SupportTicketRequest) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *SupportTicketRequest) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *SupportTicketRequest) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.

### GetSummary

`func (o *SupportTicketRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SupportTicketRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SupportTicketRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetManagedIssue

`func (o *SupportTicketRequest) GetManagedIssue() bool`

GetManagedIssue returns the ManagedIssue field if non-nil, zero value otherwise.

### GetManagedIssueOk

`func (o *SupportTicketRequest) GetManagedIssueOk() (*bool, bool)`

GetManagedIssueOk returns a tuple with the ManagedIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedIssue

`func (o *SupportTicketRequest) SetManagedIssue(v bool)`

SetManagedIssue sets ManagedIssue field to given value.

### HasManagedIssue

`func (o *SupportTicketRequest) HasManagedIssue() bool`

HasManagedIssue returns a boolean if a field has been set.

### GetVolumeId

`func (o *SupportTicketRequest) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SupportTicketRequest) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SupportTicketRequest) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *SupportTicketRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


