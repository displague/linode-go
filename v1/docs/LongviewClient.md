# LongviewClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Client&#39;s unique ID.  | [optional] [readonly] 
**Label** | Pointer to **string** | This Client&#39;s unique label. This is for display purposes only.  | [optional] 
**ApiKey** | Pointer to **string** | The API key for this Client, used when configuring the Longview Client application on your Linode.  | [optional] [readonly] 
**InstallCode** | Pointer to **string** | The install code for this Client, used when configuring the Longview Client application on your Linode.  | [optional] [readonly] 
**Apps** | Pointer to [**LongviewClientApps**](LongviewClientApps.md) |  | [optional] 
**Created** | Pointer to **time.Time** | When this Longview Client was created.  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Longview Client was last updated.  | [optional] [readonly] 

## Methods

### NewLongviewClient

`func NewLongviewClient() *LongviewClient`

NewLongviewClient instantiates a new LongviewClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongviewClientWithDefaults

`func NewLongviewClientWithDefaults() *LongviewClient`

NewLongviewClientWithDefaults instantiates a new LongviewClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LongviewClient) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LongviewClient) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LongviewClient) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LongviewClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *LongviewClient) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LongviewClient) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LongviewClient) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LongviewClient) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetApiKey

`func (o *LongviewClient) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *LongviewClient) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *LongviewClient) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *LongviewClient) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetInstallCode

`func (o *LongviewClient) GetInstallCode() string`

GetInstallCode returns the InstallCode field if non-nil, zero value otherwise.

### GetInstallCodeOk

`func (o *LongviewClient) GetInstallCodeOk() (*string, bool)`

GetInstallCodeOk returns a tuple with the InstallCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallCode

`func (o *LongviewClient) SetInstallCode(v string)`

SetInstallCode sets InstallCode field to given value.

### HasInstallCode

`func (o *LongviewClient) HasInstallCode() bool`

HasInstallCode returns a boolean if a field has been set.

### GetApps

`func (o *LongviewClient) GetApps() LongviewClientApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *LongviewClient) GetAppsOk() (*LongviewClientApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *LongviewClient) SetApps(v LongviewClientApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *LongviewClient) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCreated

`func (o *LongviewClient) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LongviewClient) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LongviewClient) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LongviewClient) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *LongviewClient) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LongviewClient) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LongviewClient) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *LongviewClient) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


