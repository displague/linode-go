# SupportTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the Support Ticket.  | [optional] [readonly] 
**Attachments** | Pointer to **[]string** | A list of filenames representing attached files associated with this Ticket.  | [optional] [readonly] 
**Closed** | Pointer to **NullableTime** | The date and time this Ticket was closed.  | [optional] [readonly] 
**Closable** | Pointer to **bool** | Whether the Support Ticket may be closed.  | [optional] 
**Description** | Pointer to **string** | The full details of the issue or question.  | [optional] [readonly] 
**Entity** | Pointer to [**NullableSupportTicketEntity**](SupportTicketEntity.md) |  | [optional] 
**GravatarId** | Pointer to **string** | The Gravatar ID of the User who opened this Ticket.  | [optional] [readonly] 
**Opened** | Pointer to **time.Time** | The date and time this Ticket was created.  | [optional] [readonly] 
**OpenedBy** | Pointer to **string** | The User who opened this Ticket.  | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of this Ticket. | [optional] [readonly] 
**Summary** | Pointer to **string** | The summary or title for this Ticket.  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The date and time this Ticket was last updated.  | [optional] [readonly] 
**UpdatedBy** | Pointer to **NullableString** | The User who last updated this Ticket.  | [optional] [readonly] 

## Methods

### NewSupportTicket

`func NewSupportTicket() *SupportTicket`

NewSupportTicket instantiates a new SupportTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketWithDefaults

`func NewSupportTicketWithDefaults() *SupportTicket`

NewSupportTicketWithDefaults instantiates a new SupportTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttachments

`func (o *SupportTicket) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SupportTicket) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SupportTicket) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SupportTicket) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetClosed

`func (o *SupportTicket) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *SupportTicket) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *SupportTicket) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *SupportTicket) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### SetClosedNil

`func (o *SupportTicket) SetClosedNil(b bool)`

 SetClosedNil sets the value for Closed to be an explicit nil

### UnsetClosed
`func (o *SupportTicket) UnsetClosed()`

UnsetClosed ensures that no value is present for Closed, not even an explicit nil
### GetClosable

`func (o *SupportTicket) GetClosable() bool`

GetClosable returns the Closable field if non-nil, zero value otherwise.

### GetClosableOk

`func (o *SupportTicket) GetClosableOk() (*bool, bool)`

GetClosableOk returns a tuple with the Closable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosable

`func (o *SupportTicket) SetClosable(v bool)`

SetClosable sets Closable field to given value.

### HasClosable

`func (o *SupportTicket) HasClosable() bool`

HasClosable returns a boolean if a field has been set.

### GetDescription

`func (o *SupportTicket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntity

`func (o *SupportTicket) GetEntity() SupportTicketEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SupportTicket) GetEntityOk() (*SupportTicketEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SupportTicket) SetEntity(v SupportTicketEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SupportTicket) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *SupportTicket) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *SupportTicket) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil
### GetGravatarId

`func (o *SupportTicket) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *SupportTicket) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *SupportTicket) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *SupportTicket) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### GetOpened

`func (o *SupportTicket) GetOpened() time.Time`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *SupportTicket) GetOpenedOk() (*time.Time, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *SupportTicket) SetOpened(v time.Time)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *SupportTicket) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetOpenedBy

`func (o *SupportTicket) GetOpenedBy() string`

GetOpenedBy returns the OpenedBy field if non-nil, zero value otherwise.

### GetOpenedByOk

`func (o *SupportTicket) GetOpenedByOk() (*string, bool)`

GetOpenedByOk returns a tuple with the OpenedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedBy

`func (o *SupportTicket) SetOpenedBy(v string)`

SetOpenedBy sets OpenedBy field to given value.

### HasOpenedBy

`func (o *SupportTicket) HasOpenedBy() bool`

HasOpenedBy returns a boolean if a field has been set.

### GetStatus

`func (o *SupportTicket) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportTicket) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportTicket) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportTicket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *SupportTicket) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SupportTicket) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SupportTicket) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SupportTicket) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdated

`func (o *SupportTicket) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SupportTicket) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SupportTicket) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SupportTicket) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SupportTicket) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SupportTicket) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SupportTicket) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SupportTicket) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *SupportTicket) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *SupportTicket) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


