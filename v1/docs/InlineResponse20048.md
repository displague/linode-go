# InlineResponse20048

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The Access Control Level of the bucket, as a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly.  | [optional] 
**AclXml** | Pointer to **string** | The full XML of the object&#39;s ACL policy.  | [optional] 

## Methods

### NewInlineResponse20048

`func NewInlineResponse20048() *InlineResponse20048`

NewInlineResponse20048 instantiates a new InlineResponse20048 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20048WithDefaults

`func NewInlineResponse20048WithDefaults() *InlineResponse20048`

NewInlineResponse20048WithDefaults instantiates a new InlineResponse20048 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *InlineResponse20048) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *InlineResponse20048) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *InlineResponse20048) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *InlineResponse20048) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAclXml

`func (o *InlineResponse20048) GetAclXml() string`

GetAclXml returns the AclXml field if non-nil, zero value otherwise.

### GetAclXmlOk

`func (o *InlineResponse20048) GetAclXmlOk() (*string, bool)`

GetAclXmlOk returns a tuple with the AclXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclXml

`func (o *InlineResponse20048) SetAclXml(v string)`

SetAclXml sets AclXml field to given value.

### HasAclXml

`func (o *InlineResponse20048) HasAclXml() bool`

HasAclXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


