# DomainRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Record&#39;s unique ID. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. For more information, see our guide on [DNS Records](/docs/guides/dns-records-an-introduction).  | [optional] 
**Name** | Pointer to **string** | The name of this Record. For requests, this property&#39;s actual usage and whether it is required depends on the type of record this represents:  &#x60;A&#x60; and &#x60;AAAA&#x60;: The hostname or FQDN of the Record.  &#x60;NS&#x60;: The subdomain, if any, to use with the Domain of the Record.  &#x60;MX&#x60;: The subdomain.  &#x60;CNAME&#x60;: The hostname. Must be unique. Required.  &#x60;TXT&#x60;: The hostname.  &#x60;SRV&#x60;: Unused. Use the &#x60;service&#x60; property to set the service name for this record.  &#x60;CAA&#x60;: The subdomain. Omit or enter an empty string (\&quot;\&quot;) to apply to the entire Domain.  &#x60;PTR&#x60;: See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](/docs/guides/configure-your-linode-for-reverse-dns).  | [optional] 
**Target** | Pointer to **string** | The target for this Record. For requests, this property&#39;s actual usage and whether it is required depends on the type of record this represents:  &#x60;A&#x60; and &#x60;AAAA&#x60;: The IP address. Use &#x60;[remote_addr]&#x60; to submit the IPv4 address of the request. Required.  &#x60;NS&#x60;: The name server. Must be a valid domain. Required.  &#x60;MX&#x60;: The mail server. Must be a valid domain. Required.  &#x60;CNAME&#x60;: The alias. Must be a valid domain. Required.  &#x60;TXT&#x60;: The value. Required.  &#x60;SRV&#x60;: The target domain or subdomain. If a subdomain is entered, it is automatically used with the Domain. To configure for a different domain, enter a valid FQDN. For example, the value &#x60;www&#x60; with a Domain for &#x60;example.com&#x60; results in a target set to &#x60;www.example.com&#x60;, whereas the value &#x60;sample.com&#x60; results in a target set to &#x60;sample.com&#x60;. Required.  &#x60;CAA&#x60;: The value. For &#x60;issue&#x60; or &#x60;issuewild&#x60; tags, the domain of your certificate issuer. For the &#x60;iodef&#x60; tag, a contact or submission URL (http or mailto).  &#x60;PTR&#x60;: See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](/docs/guides/configure-your-linode-for-reverse-dns).  With the exception of A, AAAA, and CAA records, this field accepts a trailing period.  | [optional] 
**Priority** | Pointer to **int32** | The priority of the target host for this Record. Lower values are preferred. Only valid and required for SRV record requests.  | [optional] 
**Weight** | Pointer to **int32** | The relative weight of this Record. Higher values are preferred. Only valid and required for SRV record requests.  | [optional] 
**Port** | Pointer to **int32** | The port this Record points to. Only valid and required for SRV record requests.  | [optional] 
**Service** | Pointer to **NullableString** | The name of the service. An underscore (_) is prepended and a period (.) is appended automatically to the submitted value for this property. Only valid and required for SRV record requests.  | [optional] 
**Protocol** | Pointer to **NullableString** | The protocol this Record&#39;s service communicates with. An underscore (_) is prepended automatically to the submitted value for this property. Only valid for SRV record requests.  | [optional] 
**TtlSec** | Pointer to **int32** | \&quot;Time to Live\&quot; - the amount of time in seconds that this Domain&#39;s records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.  | [optional] 
**Tag** | Pointer to **NullableString** | The tag portion of a CAA record. Only valid and required for CAA record requests.  | [optional] 
**Created** | Pointer to **time.Time** | When this Domain Record was created. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Domain Record was last updated. | [optional] [readonly] 

## Methods

### NewDomainRecord

`func NewDomainRecord() *DomainRecord`

NewDomainRecord instantiates a new DomainRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRecordWithDefaults

`func NewDomainRecordWithDefaults() *DomainRecord`

NewDomainRecordWithDefaults instantiates a new DomainRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainRecord) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DomainRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DomainRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DomainRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DomainRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *DomainRecord) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DomainRecord) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DomainRecord) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DomainRecord) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPriority

`func (o *DomainRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DomainRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DomainRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DomainRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *DomainRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DomainRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DomainRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DomainRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *DomainRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DomainRecord) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DomainRecord) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DomainRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetService

`func (o *DomainRecord) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DomainRecord) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DomainRecord) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DomainRecord) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *DomainRecord) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *DomainRecord) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetProtocol

`func (o *DomainRecord) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DomainRecord) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DomainRecord) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DomainRecord) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *DomainRecord) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *DomainRecord) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetTtlSec

`func (o *DomainRecord) GetTtlSec() int32`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *DomainRecord) GetTtlSecOk() (*int32, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *DomainRecord) SetTtlSec(v int32)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *DomainRecord) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.

### GetTag

`func (o *DomainRecord) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DomainRecord) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DomainRecord) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DomainRecord) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *DomainRecord) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *DomainRecord) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetCreated

`func (o *DomainRecord) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DomainRecord) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DomainRecord) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DomainRecord) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DomainRecord) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DomainRecord) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DomainRecord) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DomainRecord) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


