# SupportTicketReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date and time this Ticket reply was created.  | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The User who submitted this reply.  | [optional] [readonly] 
**Description** | Pointer to **string** | The body of this Support Ticket reply.  | [optional] [readonly] 
**FromLinode** | Pointer to **bool** | If set to true, this reply came from a Linode employee.  | [optional] [readonly] 
**GravatarId** | Pointer to **string** | The Gravatar ID of the User who created this reply.  | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this Support Ticket reply.  | [optional] [readonly] 

## Methods

### NewSupportTicketReply

`func NewSupportTicketReply() *SupportTicketReply`

NewSupportTicketReply instantiates a new SupportTicketReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketReplyWithDefaults

`func NewSupportTicketReplyWithDefaults() *SupportTicketReply`

NewSupportTicketReplyWithDefaults instantiates a new SupportTicketReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SupportTicketReply) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SupportTicketReply) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SupportTicketReply) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SupportTicketReply) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SupportTicketReply) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SupportTicketReply) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SupportTicketReply) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SupportTicketReply) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *SupportTicketReply) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketReply) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketReply) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicketReply) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromLinode

`func (o *SupportTicketReply) GetFromLinode() bool`

GetFromLinode returns the FromLinode field if non-nil, zero value otherwise.

### GetFromLinodeOk

`func (o *SupportTicketReply) GetFromLinodeOk() (*bool, bool)`

GetFromLinodeOk returns a tuple with the FromLinode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLinode

`func (o *SupportTicketReply) SetFromLinode(v bool)`

SetFromLinode sets FromLinode field to given value.

### HasFromLinode

`func (o *SupportTicketReply) HasFromLinode() bool`

HasFromLinode returns a boolean if a field has been set.

### GetGravatarId

`func (o *SupportTicketReply) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *SupportTicketReply) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *SupportTicketReply) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *SupportTicketReply) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### GetId

`func (o *SupportTicketReply) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketReply) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketReply) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketReply) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


