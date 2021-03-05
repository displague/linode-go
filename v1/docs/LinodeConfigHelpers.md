# LinodeConfigHelpers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedbDisabled** | Pointer to **bool** | Disables updatedb cron job to avoid disk thrashing. | [optional] 
**Distro** | Pointer to **bool** | Helps maintain correct inittab/upstart console device. | [optional] 
**ModulesDep** | Pointer to **bool** | Creates a modules dependency file for the Kernel you run. | [optional] 
**Network** | Pointer to **bool** | Automatically configures static networking. | [optional] 
**DevtmpfsAutomount** | Pointer to **bool** | Populates the /dev directory early during boot without udev.  Defaults to false.  | [optional] 

## Methods

### NewLinodeConfigHelpers

`func NewLinodeConfigHelpers() *LinodeConfigHelpers`

NewLinodeConfigHelpers instantiates a new LinodeConfigHelpers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeConfigHelpersWithDefaults

`func NewLinodeConfigHelpersWithDefaults() *LinodeConfigHelpers`

NewLinodeConfigHelpersWithDefaults instantiates a new LinodeConfigHelpers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedbDisabled

`func (o *LinodeConfigHelpers) GetUpdatedbDisabled() bool`

GetUpdatedbDisabled returns the UpdatedbDisabled field if non-nil, zero value otherwise.

### GetUpdatedbDisabledOk

`func (o *LinodeConfigHelpers) GetUpdatedbDisabledOk() (*bool, bool)`

GetUpdatedbDisabledOk returns a tuple with the UpdatedbDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedbDisabled

`func (o *LinodeConfigHelpers) SetUpdatedbDisabled(v bool)`

SetUpdatedbDisabled sets UpdatedbDisabled field to given value.

### HasUpdatedbDisabled

`func (o *LinodeConfigHelpers) HasUpdatedbDisabled() bool`

HasUpdatedbDisabled returns a boolean if a field has been set.

### GetDistro

`func (o *LinodeConfigHelpers) GetDistro() bool`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *LinodeConfigHelpers) GetDistroOk() (*bool, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *LinodeConfigHelpers) SetDistro(v bool)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *LinodeConfigHelpers) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetModulesDep

`func (o *LinodeConfigHelpers) GetModulesDep() bool`

GetModulesDep returns the ModulesDep field if non-nil, zero value otherwise.

### GetModulesDepOk

`func (o *LinodeConfigHelpers) GetModulesDepOk() (*bool, bool)`

GetModulesDepOk returns a tuple with the ModulesDep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesDep

`func (o *LinodeConfigHelpers) SetModulesDep(v bool)`

SetModulesDep sets ModulesDep field to given value.

### HasModulesDep

`func (o *LinodeConfigHelpers) HasModulesDep() bool`

HasModulesDep returns a boolean if a field has been set.

### GetNetwork

`func (o *LinodeConfigHelpers) GetNetwork() bool`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *LinodeConfigHelpers) GetNetworkOk() (*bool, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *LinodeConfigHelpers) SetNetwork(v bool)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *LinodeConfigHelpers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDevtmpfsAutomount

`func (o *LinodeConfigHelpers) GetDevtmpfsAutomount() bool`

GetDevtmpfsAutomount returns the DevtmpfsAutomount field if non-nil, zero value otherwise.

### GetDevtmpfsAutomountOk

`func (o *LinodeConfigHelpers) GetDevtmpfsAutomountOk() (*bool, bool)`

GetDevtmpfsAutomountOk returns a tuple with the DevtmpfsAutomount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevtmpfsAutomount

`func (o *LinodeConfigHelpers) SetDevtmpfsAutomount(v bool)`

SetDevtmpfsAutomount sets DevtmpfsAutomount field to given value.

### HasDevtmpfsAutomount

`func (o *LinodeConfigHelpers) HasDevtmpfsAutomount() bool`

HasDevtmpfsAutomount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


