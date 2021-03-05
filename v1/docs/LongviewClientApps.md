# LongviewClientApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apache** | Pointer to **bool** | If True, the Apache Longview Client module is monitoring Apache on your server.  | [optional] [readonly] 
**Nginx** | Pointer to **bool** | If True, the Nginx Longview Client module is monitoring Nginx on your server.  | [optional] [readonly] 
**Mysql** | Pointer to **bool** | If True, the MySQL Longview Client modules is monitoring MySQL on your server.  | [optional] [readonly] 

## Methods

### NewLongviewClientApps

`func NewLongviewClientApps() *LongviewClientApps`

NewLongviewClientApps instantiates a new LongviewClientApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongviewClientAppsWithDefaults

`func NewLongviewClientAppsWithDefaults() *LongviewClientApps`

NewLongviewClientAppsWithDefaults instantiates a new LongviewClientApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApache

`func (o *LongviewClientApps) GetApache() bool`

GetApache returns the Apache field if non-nil, zero value otherwise.

### GetApacheOk

`func (o *LongviewClientApps) GetApacheOk() (*bool, bool)`

GetApacheOk returns a tuple with the Apache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApache

`func (o *LongviewClientApps) SetApache(v bool)`

SetApache sets Apache field to given value.

### HasApache

`func (o *LongviewClientApps) HasApache() bool`

HasApache returns a boolean if a field has been set.

### GetNginx

`func (o *LongviewClientApps) GetNginx() bool`

GetNginx returns the Nginx field if non-nil, zero value otherwise.

### GetNginxOk

`func (o *LongviewClientApps) GetNginxOk() (*bool, bool)`

GetNginxOk returns a tuple with the Nginx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginx

`func (o *LongviewClientApps) SetNginx(v bool)`

SetNginx sets Nginx field to given value.

### HasNginx

`func (o *LongviewClientApps) HasNginx() bool`

HasNginx returns a boolean if a field has been set.

### GetMysql

`func (o *LongviewClientApps) GetMysql() bool`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *LongviewClientApps) GetMysqlOk() (*bool, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *LongviewClientApps) SetMysql(v bool)`

SetMysql sets Mysql field to given value.

### HasMysql

`func (o *LongviewClientApps) HasMysql() bool`

HasMysql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


