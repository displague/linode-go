# NodeBalancerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This config&#39;s unique ID | [optional] [readonly] 
**Port** | Pointer to **int32** | The port this Config is for. These values must be unique across configs on a single NodeBalancer (you can&#39;t have two configs for port 80, for example).  While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.  | [optional] 
**Protocol** | Pointer to **string** | The protocol this port is configured to serve. * If using &#x60;http&#x60; or &#x60;tcp&#x60; protocol, &#x60;ssl_cert&#x60; and &#x60;ssl_key&#x60; are not supported. * If using &#x60;https&#x60; protocol, &#x60;ssl_cert&#x60; and &#x60;ssl_key&#x60; are required.  | [optional] 
**Algorithm** | Pointer to **string** | What algorithm this NodeBalancer should use for routing traffic to backends.  | [optional] 
**Stickiness** | Pointer to **string** | Controls how session stickiness is handled on this port. * If set to &#x60;none&#x60; connections will always be assigned a backend based on the algorithm configured. * If set to &#x60;table&#x60; sessions from the same remote address will be routed to the same   backend.  * For HTTP or HTTPS clients, &#x60;http_cookie&#x60; allows sessions to be   routed to the same backend based on a cookie set by the NodeBalancer.  | [optional] 
**Check** | Pointer to **string** | The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. * If &#x60;none&#x60; no check is performed. * &#x60;connection&#x60; requires only a connection to the backend to succeed. * &#x60;http&#x60; and &#x60;http_body&#x60; rely on the backend serving HTTP, and that   the response returned matches what is expected.  | [optional] 
**CheckInterval** | Pointer to **int32** | How often, in seconds, to check that backends are up and serving requests.  | [optional] 
**CheckTimeout** | Pointer to **int32** | How long, in seconds, to wait for a check attempt before considering it failed.  | [optional] 
**CheckAttempts** | Pointer to **int32** | How many times to attempt a check before considering a backend to be down.  | [optional] 
**CheckPath** | Pointer to **string** | The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.  | [optional] 
**CheckBody** | Pointer to **string** | This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down.  | [optional] 
**CheckPassive** | Pointer to **bool** | If true, any response from this backend with a &#x60;5xx&#x60; status code will be enough for it to be considered unhealthy and taken out of rotation.  | [optional] 
**ProxyProtocol** | Pointer to **string** | ProxyProtocol is a TCP extension that sends initial TCP connection information such as source/destination IPs and ports to backend devices. This information would be lost otherwise. Backend devices must be configured to work with ProxyProtocol if enabled.  * If ommited, or set to &#x60;none&#x60;, the NodeBalancer doesn&#39;t send any auxilary data over TCP connections. This is the default. * If set to &#x60;v1&#x60;, the human-readable header format (Version 1) is used. * If set to &#x60;v2&#x60;, the binary header format (Version 2) is used.  | [optional] [default to "none"]
**CipherSuite** | Pointer to **string** | What ciphers to use for SSL connections served by this NodeBalancer.  * &#x60;legacy&#x60; is considered insecure and should only be used if necessary.  | [optional] 
**NodebalancerId** | Pointer to **int32** | The ID for the NodeBalancer this config belongs to.  | [optional] [readonly] 
**SslCommonname** | Pointer to **string** | The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.  | [optional] [readonly] 
**SslFingerprint** | Pointer to **string** | The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.  | [optional] [readonly] 
**SslCert** | Pointer to **NullableString** | The PEM-formatted public SSL certificate (or the combined PEM-formatted SSL certificate and Certificate Authority chain) that should be served on this NodeBalancerConfig&#39;s port.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, &#x60;&lt;REDACTED&gt;&#x60; will be printed where the field appears.  The read-only &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig.  | [optional] 
**SslKey** | Pointer to **NullableString** | The PEM-formatted private key for the SSL certificate set in the &#x60;ssl_cert&#x60; field.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, &#x60;&lt;REDACTED&gt;&#x60; will be printed where the field appears.  The read-only &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig.  | [optional] 
**NodesStatus** | Pointer to [**NodeBalancerConfigNodesStatus**](NodeBalancerConfigNodesStatus.md) |  | [optional] 

## Methods

### NewNodeBalancerConfig

`func NewNodeBalancerConfig() *NodeBalancerConfig`

NewNodeBalancerConfig instantiates a new NodeBalancerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerConfigWithDefaults

`func NewNodeBalancerConfigWithDefaults() *NodeBalancerConfig`

NewNodeBalancerConfigWithDefaults instantiates a new NodeBalancerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeBalancerConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeBalancerConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeBalancerConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeBalancerConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *NodeBalancerConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NodeBalancerConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NodeBalancerConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NodeBalancerConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *NodeBalancerConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NodeBalancerConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NodeBalancerConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NodeBalancerConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAlgorithm

`func (o *NodeBalancerConfig) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NodeBalancerConfig) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NodeBalancerConfig) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *NodeBalancerConfig) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetStickiness

`func (o *NodeBalancerConfig) GetStickiness() string`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *NodeBalancerConfig) GetStickinessOk() (*string, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *NodeBalancerConfig) SetStickiness(v string)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *NodeBalancerConfig) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetCheck

`func (o *NodeBalancerConfig) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *NodeBalancerConfig) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *NodeBalancerConfig) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *NodeBalancerConfig) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckInterval

`func (o *NodeBalancerConfig) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *NodeBalancerConfig) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *NodeBalancerConfig) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *NodeBalancerConfig) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetCheckTimeout

`func (o *NodeBalancerConfig) GetCheckTimeout() int32`

GetCheckTimeout returns the CheckTimeout field if non-nil, zero value otherwise.

### GetCheckTimeoutOk

`func (o *NodeBalancerConfig) GetCheckTimeoutOk() (*int32, bool)`

GetCheckTimeoutOk returns a tuple with the CheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTimeout

`func (o *NodeBalancerConfig) SetCheckTimeout(v int32)`

SetCheckTimeout sets CheckTimeout field to given value.

### HasCheckTimeout

`func (o *NodeBalancerConfig) HasCheckTimeout() bool`

HasCheckTimeout returns a boolean if a field has been set.

### GetCheckAttempts

`func (o *NodeBalancerConfig) GetCheckAttempts() int32`

GetCheckAttempts returns the CheckAttempts field if non-nil, zero value otherwise.

### GetCheckAttemptsOk

`func (o *NodeBalancerConfig) GetCheckAttemptsOk() (*int32, bool)`

GetCheckAttemptsOk returns a tuple with the CheckAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAttempts

`func (o *NodeBalancerConfig) SetCheckAttempts(v int32)`

SetCheckAttempts sets CheckAttempts field to given value.

### HasCheckAttempts

`func (o *NodeBalancerConfig) HasCheckAttempts() bool`

HasCheckAttempts returns a boolean if a field has been set.

### GetCheckPath

`func (o *NodeBalancerConfig) GetCheckPath() string`

GetCheckPath returns the CheckPath field if non-nil, zero value otherwise.

### GetCheckPathOk

`func (o *NodeBalancerConfig) GetCheckPathOk() (*string, bool)`

GetCheckPathOk returns a tuple with the CheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPath

`func (o *NodeBalancerConfig) SetCheckPath(v string)`

SetCheckPath sets CheckPath field to given value.

### HasCheckPath

`func (o *NodeBalancerConfig) HasCheckPath() bool`

HasCheckPath returns a boolean if a field has been set.

### GetCheckBody

`func (o *NodeBalancerConfig) GetCheckBody() string`

GetCheckBody returns the CheckBody field if non-nil, zero value otherwise.

### GetCheckBodyOk

`func (o *NodeBalancerConfig) GetCheckBodyOk() (*string, bool)`

GetCheckBodyOk returns a tuple with the CheckBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBody

`func (o *NodeBalancerConfig) SetCheckBody(v string)`

SetCheckBody sets CheckBody field to given value.

### HasCheckBody

`func (o *NodeBalancerConfig) HasCheckBody() bool`

HasCheckBody returns a boolean if a field has been set.

### GetCheckPassive

`func (o *NodeBalancerConfig) GetCheckPassive() bool`

GetCheckPassive returns the CheckPassive field if non-nil, zero value otherwise.

### GetCheckPassiveOk

`func (o *NodeBalancerConfig) GetCheckPassiveOk() (*bool, bool)`

GetCheckPassiveOk returns a tuple with the CheckPassive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassive

`func (o *NodeBalancerConfig) SetCheckPassive(v bool)`

SetCheckPassive sets CheckPassive field to given value.

### HasCheckPassive

`func (o *NodeBalancerConfig) HasCheckPassive() bool`

HasCheckPassive returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *NodeBalancerConfig) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *NodeBalancerConfig) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *NodeBalancerConfig) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *NodeBalancerConfig) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetCipherSuite

`func (o *NodeBalancerConfig) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *NodeBalancerConfig) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *NodeBalancerConfig) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *NodeBalancerConfig) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *NodeBalancerConfig) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *NodeBalancerConfig) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *NodeBalancerConfig) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *NodeBalancerConfig) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.

### GetSslCommonname

`func (o *NodeBalancerConfig) GetSslCommonname() string`

GetSslCommonname returns the SslCommonname field if non-nil, zero value otherwise.

### GetSslCommonnameOk

`func (o *NodeBalancerConfig) GetSslCommonnameOk() (*string, bool)`

GetSslCommonnameOk returns a tuple with the SslCommonname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCommonname

`func (o *NodeBalancerConfig) SetSslCommonname(v string)`

SetSslCommonname sets SslCommonname field to given value.

### HasSslCommonname

`func (o *NodeBalancerConfig) HasSslCommonname() bool`

HasSslCommonname returns a boolean if a field has been set.

### GetSslFingerprint

`func (o *NodeBalancerConfig) GetSslFingerprint() string`

GetSslFingerprint returns the SslFingerprint field if non-nil, zero value otherwise.

### GetSslFingerprintOk

`func (o *NodeBalancerConfig) GetSslFingerprintOk() (*string, bool)`

GetSslFingerprintOk returns a tuple with the SslFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslFingerprint

`func (o *NodeBalancerConfig) SetSslFingerprint(v string)`

SetSslFingerprint sets SslFingerprint field to given value.

### HasSslFingerprint

`func (o *NodeBalancerConfig) HasSslFingerprint() bool`

HasSslFingerprint returns a boolean if a field has been set.

### GetSslCert

`func (o *NodeBalancerConfig) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *NodeBalancerConfig) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *NodeBalancerConfig) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *NodeBalancerConfig) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *NodeBalancerConfig) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *NodeBalancerConfig) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetSslKey

`func (o *NodeBalancerConfig) GetSslKey() string`

GetSslKey returns the SslKey field if non-nil, zero value otherwise.

### GetSslKeyOk

`func (o *NodeBalancerConfig) GetSslKeyOk() (*string, bool)`

GetSslKeyOk returns a tuple with the SslKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslKey

`func (o *NodeBalancerConfig) SetSslKey(v string)`

SetSslKey sets SslKey field to given value.

### HasSslKey

`func (o *NodeBalancerConfig) HasSslKey() bool`

HasSslKey returns a boolean if a field has been set.

### SetSslKeyNil

`func (o *NodeBalancerConfig) SetSslKeyNil(b bool)`

 SetSslKeyNil sets the value for SslKey to be an explicit nil

### UnsetSslKey
`func (o *NodeBalancerConfig) UnsetSslKey()`

UnsetSslKey ensures that no value is present for SslKey, not even an explicit nil
### GetNodesStatus

`func (o *NodeBalancerConfig) GetNodesStatus() NodeBalancerConfigNodesStatus`

GetNodesStatus returns the NodesStatus field if non-nil, zero value otherwise.

### GetNodesStatusOk

`func (o *NodeBalancerConfig) GetNodesStatusOk() (*NodeBalancerConfigNodesStatus, bool)`

GetNodesStatusOk returns a tuple with the NodesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesStatus

`func (o *NodeBalancerConfig) SetNodesStatus(v NodeBalancerConfigNodesStatus)`

SetNodesStatus sets NodesStatus field to given value.

### HasNodesStatus

`func (o *NodeBalancerConfig) HasNodesStatus() bool`

HasNodesStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


