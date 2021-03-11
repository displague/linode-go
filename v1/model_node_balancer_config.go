/*
 * Linode API
 *
 * ## Introduction The Linode API provides the ability to programmatically manage the full range of Linode products and services.  This reference is designed to assist application developers and system administrators.  Each endpoint includes descriptions, request syntax, and examples using standard HTTP requests. Response data is returned in JSON format.   This document was generated from our OpenAPI Specification.  See the <a target=\"_top\" href=\"https://www.openapis.org\">OpenAPI website</a> for more information.  <a target=\"_top\" href=\"/docs/api/openapi.yaml\">Download the Linode OpenAPI Specification</a>.   ## Changelog  <a target=\"_top\" href=\"https://developers.linode.com/changelog\">View our Changelog</a> to see release notes on all changes made to our API.  ## Access and Authentication  Some endpoints are publicly accessible without requiring authentication. All endpoints affecting your Account, however, require either a Personal Access Token or OAuth authentication (when using third-party applications).  ### Personal Access Token  The easiest way to access the API is with a Personal Access Token (PAT) generated from the <a target=\"_top\" href=\"https://cloud.linode.com/profile/tokens\">Linode Cloud Manager</a> or the [Create Personal Access Token](/docs/api/profile/#personal-access-token-create) endpoint.  All scopes for the OAuth security model ([defined below](/docs/api/profile/#oauth)) apply to this security model as well.  #### Authentication  | Security Scheme Type: | HTTP | |-----------------------|------| | **HTTP Authorization Scheme** | bearer |  ### OAuth If you only need to access the Linode API for personal use, we recommend that you create a [personal access token](/docs/api/#personal-access-token). If you're designing an application that can authenticate with an arbitrary Linode user, then you should use the OAuth 2.0 workflows presented in this section.  For a more detailed example of an OAuth 2.0 implementation, see our guide on [How to Create an OAuth App with the Linode Python API Library](/docs/platform/api/how-to-create-an-oauth-app-with-the-linode-python-api-library/#oauth-2-authentication-exchange).  Before you implement OAuth in your application, you first need to create an OAuth client. You can do this [with the Linode API](/docs/api/account/#oauth-client-create) or [via the Cloud Manager](https://cloud.linode.com/profile/clients):    - When creating the client, you'll supply a `label` and a `redirect_uri` (referred to as the Callback URL in the Cloud Manager).   - The response from this endpoint will give you a `client_id` and a `secret`.   - Clients can be public or private, and are private by default. You can choose to make the client public when it is created.     - A private client is used with applications which can securely store the client secret (that is, the secret returned to you when you first created the client). For example, an application running on a secured server that only the developer has access to would use a private OAuth client. This is also called a confidential client in some OAuth documentation.     - A public client is used with applications where the client secret is not guaranteed to be secure. For example, a native app running on a user's computer may not be able to keep the client secret safe, as a user could potentially inspect the source of the application. So, native apps or apps that run in a user's browser should use a public client.     - Public and private clients follow different workflows, as described below.  #### OAuth Workflow  The OAuth workflow is a series of exchanges between your third-party app and Linode. The workflow is used to authenticate a user before an application can start making API calls on the user's behalf.  Notes:  - With respect to the diagram in [section 1.2 of RFC 6749](https://tools.ietf.org/html/rfc6749#section-1.2), login.linode.com (referred to in this section as the *login server*) is the Resource Owner and the Authorization Server; api.linode.com (referred to here as the *api server*) is the Resource Server. - The OAuth spec refers to the private and public workflows listed below as the [authorization code flow](https://tools.ietf.org/html/rfc6749#section-1.3.1) and [implicit flow](https://tools.ietf.org/html/rfc6749#section-1.3.2).  | PRIVATE WORKFLOW | PUBLIC WORKFLOW | |------------------|------------------| | 1.  The user visits the application's website and is directed to login with Linode. | 1.  The user visits the application's website and is directed to login with Linode. | | 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. | 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. | | 3.  The user logs into the login server with their username and password. | 3.  The user logs into the login server with their username and password. | | 4.  The login server redirects the user to the specificed redirect URL with a temporary authorization `code` (exchange code) in the URL. | 4.  The login server redirects the user back to your application with an OAuth `access_token` embedded in the redirect URL's hash. This is temporary and expires in two hours. No `refresh_token` is issued. Therefore, once the `access_token` expires, a new one will need to be issued by having the user log in again. | | 5.  The application issues a POST request (*see below*) to the login server with the exchange code, `client_id`, and the client application's `client_secret`. | | | 6.  The login server responds to the client application with a new OAuth `access_token` and `refresh_token`. The `access_token` is set to expire in two hours. | | | 7.  The `refresh_token` can be used by contacting the login server with the `client_id`, `client_secret`, `grant_type`, and `refresh_token` to get a new OAuth `access_token` and `refresh_token`. The new `access_token` is good for another two hours, and the new `refresh_token`, can be used to extend the session again by this same method. | |  #### OAuth Private Workflow - Additional Details  The following information expands on steps 5 through 7 of the private workflow:  Once the user has logged into Linode and you have received an exchange code, you will need to trade that exchange code for an `access_token` and `refresh_token`. You do this by making an HTTP POST request to the following address:  ``` https://login.linode.com/oauth/token ```  Make this request as `application/x-www-form-urlencoded` or as `multipart/form-data` and include the following parameters in the POST body:  | PARAMETER | DESCRIPTION | |-----------|-------------| | grant_type | The grant type you're using for renewal.  Currently only the string \"refresh_token\" is accepted. | | client_id | Your app's client ID. | | client_secret | Your app's client secret. | | code | The code you just received from the redirect. |  You'll get a response like this:  ```json {   \"scope\": \"linodes:read_write\",   \"access_token\": \"03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c\"   \"token_type\": \"bearer\",   \"expires_in\": 7200, } ```  Included in the reponse is an `access_token`. With this token, you can proceed to make authenticated HTTP requests to the API by adding this header to each request:  ``` Authorization: Bearer 03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c ```  #### OAuth Reference  | Security Scheme Type | OAuth 2.0 | |-----------------------|--------| | **Authorization URL** | https://login.linode.com/oauth/authorize | | **Token URL** | https://login.linode.com/oauth/token | | **Scopes** | <ul><li>`account:read_only` - Allows access to GET information about your Account.</li><li>`account:read_write` - Allows access to all endpoints related to your Account.</li><li>`domains:read_only` - Allows access to GET Domains on your Account.</li><li>`domains:read_write` - Allows access to all Domain endpoints.</li><li>`events:read_only` - Allows access to GET your Events.</li><li>`events:read_write` - Allows access to all endpoints related to your Events.</li><li>`firewall:read_only` - Allows access to GET information about your Firewalls.</li><li>`firewall:read_write` - Allows access to all Firewall endpoints.</li><li>`images:read_only` - Allows access to GET your Images.</li><li>`images:read_write` - Allows access to all endpoints related to your Images.</li><li>`ips:read_only` - Allows access to GET your ips.</li><li>`ips:read_write` - Allows access to all endpoints related to your ips.</li><li>`linodes:read_only` - Allows access to GET Linodes on your Account.</li><li>`linodes:read_write` - Allow access to all endpoints related to your Linodes.</li><li>`lke:read_only` - Allows access to GET LKE Clusters on your Account.</li><li>`lke:read_write` - Allows access to all endpoints related to LKE Clusters on your Account.</li><li>`longview:read_only` - Allows access to GET your Longview Clients.</li><li>`longview:read_write` - Allows access to all endpoints related to your Longview Clients.</li><li>`maintenance:read_only` - Allows access to GET information about Maintenance on your account.</li><li>`nodebalancers:read_only` - Allows access to GET NodeBalancers on your Account.</li><li>`nodebalancers:read_write` - Allows access to all NodeBalancer endpoints.</li><li>`object_storage:read_only` - Allows access to GET information related to your Object Storage.</li><li>`object_storage:read_write` - Allows access to all Object Storage endpoints.</li><li>`stackscripts:read_only` - Allows access to GET your StackScripts.</li><li>`stackscripts:read_write` - Allows access to all endpoints related to your StackScripts.</li><li>`volumes:read_only` - Allows access to GET your Volumes.</li><li>`volumes:read_write` - Allows access to all endpoints related to your Volumes.</li></ul><br/>|  ## Requests  Requests must be made over HTTPS to ensure transactions are encrypted. The following Request methods are supported:  | METHOD | USAGE | |--------|-------| | GET    | Retrieves data about collections and individual resources. | | POST   | For collections, creates a new resource of that type. Also used to perform actions on action endpoints. | | PUT    | Updates an existing resource. | | DELETE | Deletes a resource. This is a destructive action. |   ## Responses  Actions will return one following HTTP response status codes:  | STATUS  | DESCRIPTION | |---------|-------------| | 200 OK  | The request was successful. | | 204 No Content | The server successfully fulfilled the request and there is no additional content to send. | | 400 Bad Request | You submitted an invalid request (missing parameters, etc.). | | 401 Unauthorized | You failed to authenticate for this resource. | | 403 Forbidden | You are authenticated, but don't have permission to do this. | | 404 Not Found | The resource you're requesting does not exist. | | 429 Too Many Requests | You've hit a rate limit. | | 500 Internal Server Error | Please [open a Support Ticket](/docs/api/support/#support-ticket-open). |  ## Errors  Success is indicated via <a href=\"https://en.wikipedia.org/wiki/List_of_HTTP_status_codes\" target=\"_top\">Standard HTTP status codes</a>. `2xx` codes indicate success, `4xx` codes indicate a request error, and `5xx` errors indicate a server error. A request error might be an invalid input, a required parameter being omitted, or a malformed request. A server error means something went wrong processing your request. If this occurs, please [open a Support Ticket](/docs/api/support/#support-ticket-open) and let us know. Though errors are logged and we work quickly to resolve issues, opening a ticket and providing us with reproducable steps and data is always helpful.  The `errors` field is an array of the things that went wrong with your request. We will try to include as many of the problems in the response as possible, but it's conceivable that fixing these errors and resubmitting may result in new errors coming back once we are able to get further along in the process of handling your request.   Within each error object, the `field` parameter will be included if the error pertains to a specific field in the JSON you've submitted. This will be omitted if there is no relevant field. The `reason` is a human-readable explanation of the error, and will always be included.  ## Pagination  Resource lists are always paginated. The response will look similar to this:  ```json {     \"data\": [ ... ],     \"page\": 1,     \"pages\": 3,     \"results\": 300 } ```  * Pages start at 1. You may retrieve a specific page of results by adding `?page=x` to your URL (for example, `?page=4`). If the value of `page` exceeds `2^64/page_size`, the last possible page will be returned.   * Page sizes default to 100, and can be set to return between 25 and 500. Page size can be set using `?page_size=x`.  ## Filtering and Sorting  Collections are searchable by fields they include, marked in the spec as `x-linode-filterable: true`. Filters are passed in the `X-Filter` header and are formatted as JSON objects. Here is a request call for Linode Types in our \"standard\" class:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"class\": \"standard\"     }' ```  The filter object's keys are the keys of the object you're filtering, and the values are accepted values. You can add multiple filters by including more than one key. For example, filtering for \"standard\" Linode Types that offer one vcpu:  ```Shell  curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"class\": \"standard\",       \"vcpus\": 1     }' ```  In the above example, both filters are combined with an \"and\" operation. However, if you wanted either Types with one vcpu or Types in our \"standard\" class, you can add an operator:   ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"+or\": [         { \"vcpus\": 1 },         { \"class\": \"standard\" }       ]     }' ```  Each filter in the `+or` array is its own filter object, and all conditions in it are combined with an \"and\" operation as they were in the previous example.  Other operators are also available. Operators are keys of a Filter JSON object. Their value must be of the appropriate type, and they are evaluated as described below:  | OPERATOR | TYPE   | DESCRIPTION                       | |----------|--------|-----------------------------------| | +and     | array  | All conditions must be true.       | | +or      | array  | One condition must be true.        | | +gt      | number | Value must be greater than number. | | +gte     | number | Value must be greater than or equal to number. | | +lt      | number | Value must be less than number. | | +lte     | number | Value must be less than or equal to number. | | +contains | string | Given string must be in the value. | | +neq      | string | Does not equal the value.          | | +order_by | string | Attribute to order the results by - must be filterable. | | +order    | string | Either \"asc\" or \"desc\". Defaults to \"asc\". Requires `+order_by`. |  For example, filtering for [Linode Types](/docs/api/linode-types/) that offer memory equal to or higher than 61440:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"memory\": {         \"+gte\": 61440       }     }' ```  You can combine and nest operators to construct arbitrarily-complex queries. For example, give me all [Linode Types](/docs/api/linode-types/) which are either `standard` or `highmem` class, or have between 12 and 20 vcpus:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"+or\": [         {           \"+or\": [             {               \"class\": \"standard\"             },             {               \"class\": \"highmem\"             }           ]         },         {           \"+and\": [             {               \"vcpus\": {                 \"+gte\": 12               }             },             {               \"vcpus\": {                 \"+lte\": 20               }             }           ]         }       ]     }' ```  ## Rate Limiting  With the Linode API, you can make up to 1,600 general API requests every two minutes per user as determined by IP adddress or by OAuth token. Additionally, there are endpoint specfic limits defined below.  **Note:** There may be rate limiting applied at other levels outside of the API, for example, at the load balancer.  `/stats` endpoints have their own dedicated limits of 100 requests per minute per user. These endpoints are:  * [View Linode Statistics](/docs/api/linode-instances/#linode-statistics-view) * [View Linode Statistics (year/month)](/docs/api/linode-instances/#statistics-yearmonth-view) * [View NodeBalancer Statistics](/docs/api/nodebalancers/#nodebalancer-statistics-view) * [List Managed Stats](/docs/api/managed/#managed-stats-list)  Object Storage endpoints have a dedicated limit of 750 requests per second per user. The Object Storage endpoints are:  * [Object Storage Endpoints](/docs/api/object-storage/)  Opening Support Tickets has a dedicated limit of 2 requests per minute per user. That endpoint is:  * [Open Support Ticket](/docs/api/support/#support-ticket-open)  Accepting Entity Transfers has a dedicated limit of 2 requests per minute per user. That endpoint is:  * [Entity Transfer Accept](/docs/api/account/#entity-transfer-accept)  ## CLI (Command Line Interface)  The <a href=\"https://github.com/linode/linode-cli\" target=\"_top\">Linode CLI</a> allows you to easily work with the API using intuitive and simple syntax. It requires a [Personal Access Token](/docs/api/#personal-access-token) for authentication, and gives you access to all of the features and functionality of the Linode API that are documented here with CLI examples.  Endpoints that do not have CLI examples are currently unavailable through the CLI, but can be accessed via other methods such as Shell commands and other third-party applications. 
 *
 * API version: 4.85.0
 * Contact: support@linode.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// NodeBalancerConfig A NodeBalancer config represents the configuration of this NodeBalancer on a single port.  For example, a NodeBalancer Config on port 80 would typically represent how this NodeBalancer response to HTTP requests.  NodeBalancer configs have a list of backends, called \"nodes,\" that they forward requests between based on their configuration. 
type NodeBalancerConfig struct {
	// This config's unique ID
	Id *int32 `json:"id,omitempty"`
	// The port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example).  While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. 
	Port *int32 `json:"port,omitempty"`
	// The protocol this port is configured to serve. * If using `http` or `tcp` protocol, `ssl_cert` and `ssl_key` are not supported. * If using `https` protocol, `ssl_cert` and `ssl_key` are required. 
	Protocol *string `json:"protocol,omitempty"`
	// What algorithm this NodeBalancer should use for routing traffic to backends. 
	Algorithm *string `json:"algorithm,omitempty"`
	// Controls how session stickiness is handled on this port. * If set to `none` connections will always be assigned a backend based on the algorithm configured. * If set to `table` sessions from the same remote address will be routed to the same   backend.  * For HTTP or HTTPS clients, `http_cookie` allows sessions to be   routed to the same backend based on a cookie set by the NodeBalancer. 
	Stickiness *string `json:"stickiness,omitempty"`
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. * If `none` no check is performed. * `connection` requires only a connection to the backend to succeed. * `http` and `http_body` rely on the backend serving HTTP, and that   the response returned matches what is expected. 
	Check *string `json:"check,omitempty"`
	// How often, in seconds, to check that backends are up and serving requests. 
	CheckInterval *int32 `json:"check_interval,omitempty"`
	// How long, in seconds, to wait for a check attempt before considering it failed. 
	CheckTimeout *int32 `json:"check_timeout,omitempty"`
	// How many times to attempt a check before considering a backend to be down. 
	CheckAttempts *int32 `json:"check_attempts,omitempty"`
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down. 
	CheckPath *string `json:"check_path,omitempty"`
	// This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down. 
	CheckBody *string `json:"check_body,omitempty"`
	// If true, any response from this backend with a `5xx` status code will be enough for it to be considered unhealthy and taken out of rotation. 
	CheckPassive *bool `json:"check_passive,omitempty"`
	// ProxyProtocol is a TCP extension that sends initial TCP connection information such as source/destination IPs and ports to backend devices. This information would be lost otherwise. Backend devices must be configured to work with ProxyProtocol if enabled.  * If ommited, or set to `none`, the NodeBalancer doesn't send any auxilary data over TCP connections. This is the default. * If set to `v1`, the human-readable header format (Version 1) is used. * If set to `v2`, the binary header format (Version 2) is used. 
	ProxyProtocol *string `json:"proxy_protocol,omitempty"`
	// What ciphers to use for SSL connections served by this NodeBalancer.  * `legacy` is considered insecure and should only be used if necessary. 
	CipherSuite *string `json:"cipher_suite,omitempty"`
	// The ID for the NodeBalancer this config belongs to. 
	NodebalancerId *int32 `json:"nodebalancer_id,omitempty"`
	// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig. 
	SslCommonname *string `json:"ssl_commonname,omitempty"`
	// The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig. 
	SslFingerprint *string `json:"ssl_fingerprint,omitempty"`
	// The PEM-formatted public SSL certificate (or the combined PEM-formatted SSL certificate and Certificate Authority chain) that should be served on this NodeBalancerConfig's port.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, `<REDACTED>` will be printed where the field appears.  The read-only `ssl_commonname` and `ssl_fingerprint` fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig. 
	SslCert NullableString `json:"ssl_cert,omitempty"`
	// The PEM-formatted private key for the SSL certificate set in the `ssl_cert` field.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, `<REDACTED>` will be printed where the field appears.  The read-only `ssl_commonname` and `ssl_fingerprint` fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig. 
	SslKey NullableString `json:"ssl_key,omitempty"`
	NodesStatus *NodeBalancerConfigNodesStatus `json:"nodes_status,omitempty"`
}

// NewNodeBalancerConfig instantiates a new NodeBalancerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBalancerConfig() *NodeBalancerConfig {
	this := NodeBalancerConfig{}
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol
	return &this
}

// NewNodeBalancerConfigWithDefaults instantiates a new NodeBalancerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBalancerConfigWithDefaults() *NodeBalancerConfig {
	this := NodeBalancerConfig{}
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NodeBalancerConfig) SetId(v int32) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NodeBalancerConfig) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NodeBalancerConfig) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *NodeBalancerConfig) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetStickiness returns the Stickiness field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetStickiness() string {
	if o == nil || o.Stickiness == nil {
		var ret string
		return ret
	}
	return *o.Stickiness
}

// GetStickinessOk returns a tuple with the Stickiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetStickinessOk() (*string, bool) {
	if o == nil || o.Stickiness == nil {
		return nil, false
	}
	return o.Stickiness, true
}

// HasStickiness returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasStickiness() bool {
	if o != nil && o.Stickiness != nil {
		return true
	}

	return false
}

// SetStickiness gets a reference to the given string and assigns it to the Stickiness field.
func (o *NodeBalancerConfig) SetStickiness(v string) {
	o.Stickiness = &v
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheck() string {
	if o == nil || o.Check == nil {
		var ret string
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckOk() (*string, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheck() bool {
	if o != nil && o.Check != nil {
		return true
	}

	return false
}

// SetCheck gets a reference to the given string and assigns it to the Check field.
func (o *NodeBalancerConfig) SetCheck(v string) {
	o.Check = &v
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckInterval() int32 {
	if o == nil || o.CheckInterval == nil {
		var ret int32
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckIntervalOk() (*int32, bool) {
	if o == nil || o.CheckInterval == nil {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckInterval() bool {
	if o != nil && o.CheckInterval != nil {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given int32 and assigns it to the CheckInterval field.
func (o *NodeBalancerConfig) SetCheckInterval(v int32) {
	o.CheckInterval = &v
}

// GetCheckTimeout returns the CheckTimeout field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckTimeout() int32 {
	if o == nil || o.CheckTimeout == nil {
		var ret int32
		return ret
	}
	return *o.CheckTimeout
}

// GetCheckTimeoutOk returns a tuple with the CheckTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckTimeoutOk() (*int32, bool) {
	if o == nil || o.CheckTimeout == nil {
		return nil, false
	}
	return o.CheckTimeout, true
}

// HasCheckTimeout returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckTimeout() bool {
	if o != nil && o.CheckTimeout != nil {
		return true
	}

	return false
}

// SetCheckTimeout gets a reference to the given int32 and assigns it to the CheckTimeout field.
func (o *NodeBalancerConfig) SetCheckTimeout(v int32) {
	o.CheckTimeout = &v
}

// GetCheckAttempts returns the CheckAttempts field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckAttempts() int32 {
	if o == nil || o.CheckAttempts == nil {
		var ret int32
		return ret
	}
	return *o.CheckAttempts
}

// GetCheckAttemptsOk returns a tuple with the CheckAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckAttemptsOk() (*int32, bool) {
	if o == nil || o.CheckAttempts == nil {
		return nil, false
	}
	return o.CheckAttempts, true
}

// HasCheckAttempts returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckAttempts() bool {
	if o != nil && o.CheckAttempts != nil {
		return true
	}

	return false
}

// SetCheckAttempts gets a reference to the given int32 and assigns it to the CheckAttempts field.
func (o *NodeBalancerConfig) SetCheckAttempts(v int32) {
	o.CheckAttempts = &v
}

// GetCheckPath returns the CheckPath field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckPath() string {
	if o == nil || o.CheckPath == nil {
		var ret string
		return ret
	}
	return *o.CheckPath
}

// GetCheckPathOk returns a tuple with the CheckPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckPathOk() (*string, bool) {
	if o == nil || o.CheckPath == nil {
		return nil, false
	}
	return o.CheckPath, true
}

// HasCheckPath returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckPath() bool {
	if o != nil && o.CheckPath != nil {
		return true
	}

	return false
}

// SetCheckPath gets a reference to the given string and assigns it to the CheckPath field.
func (o *NodeBalancerConfig) SetCheckPath(v string) {
	o.CheckPath = &v
}

// GetCheckBody returns the CheckBody field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckBody() string {
	if o == nil || o.CheckBody == nil {
		var ret string
		return ret
	}
	return *o.CheckBody
}

// GetCheckBodyOk returns a tuple with the CheckBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckBodyOk() (*string, bool) {
	if o == nil || o.CheckBody == nil {
		return nil, false
	}
	return o.CheckBody, true
}

// HasCheckBody returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckBody() bool {
	if o != nil && o.CheckBody != nil {
		return true
	}

	return false
}

// SetCheckBody gets a reference to the given string and assigns it to the CheckBody field.
func (o *NodeBalancerConfig) SetCheckBody(v string) {
	o.CheckBody = &v
}

// GetCheckPassive returns the CheckPassive field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCheckPassive() bool {
	if o == nil || o.CheckPassive == nil {
		var ret bool
		return ret
	}
	return *o.CheckPassive
}

// GetCheckPassiveOk returns a tuple with the CheckPassive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCheckPassiveOk() (*bool, bool) {
	if o == nil || o.CheckPassive == nil {
		return nil, false
	}
	return o.CheckPassive, true
}

// HasCheckPassive returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCheckPassive() bool {
	if o != nil && o.CheckPassive != nil {
		return true
	}

	return false
}

// SetCheckPassive gets a reference to the given bool and assigns it to the CheckPassive field.
func (o *NodeBalancerConfig) SetCheckPassive(v bool) {
	o.CheckPassive = &v
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetProxyProtocol() string {
	if o == nil || o.ProxyProtocol == nil {
		var ret string
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetProxyProtocolOk() (*string, bool) {
	if o == nil || o.ProxyProtocol == nil {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// HasProxyProtocol returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasProxyProtocol() bool {
	if o != nil && o.ProxyProtocol != nil {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given string and assigns it to the ProxyProtocol field.
func (o *NodeBalancerConfig) SetProxyProtocol(v string) {
	o.ProxyProtocol = &v
}

// GetCipherSuite returns the CipherSuite field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetCipherSuite() string {
	if o == nil || o.CipherSuite == nil {
		var ret string
		return ret
	}
	return *o.CipherSuite
}

// GetCipherSuiteOk returns a tuple with the CipherSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetCipherSuiteOk() (*string, bool) {
	if o == nil || o.CipherSuite == nil {
		return nil, false
	}
	return o.CipherSuite, true
}

// HasCipherSuite returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasCipherSuite() bool {
	if o != nil && o.CipherSuite != nil {
		return true
	}

	return false
}

// SetCipherSuite gets a reference to the given string and assigns it to the CipherSuite field.
func (o *NodeBalancerConfig) SetCipherSuite(v string) {
	o.CipherSuite = &v
}

// GetNodebalancerId returns the NodebalancerId field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetNodebalancerId() int32 {
	if o == nil || o.NodebalancerId == nil {
		var ret int32
		return ret
	}
	return *o.NodebalancerId
}

// GetNodebalancerIdOk returns a tuple with the NodebalancerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetNodebalancerIdOk() (*int32, bool) {
	if o == nil || o.NodebalancerId == nil {
		return nil, false
	}
	return o.NodebalancerId, true
}

// HasNodebalancerId returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasNodebalancerId() bool {
	if o != nil && o.NodebalancerId != nil {
		return true
	}

	return false
}

// SetNodebalancerId gets a reference to the given int32 and assigns it to the NodebalancerId field.
func (o *NodeBalancerConfig) SetNodebalancerId(v int32) {
	o.NodebalancerId = &v
}

// GetSslCommonname returns the SslCommonname field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetSslCommonname() string {
	if o == nil || o.SslCommonname == nil {
		var ret string
		return ret
	}
	return *o.SslCommonname
}

// GetSslCommonnameOk returns a tuple with the SslCommonname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetSslCommonnameOk() (*string, bool) {
	if o == nil || o.SslCommonname == nil {
		return nil, false
	}
	return o.SslCommonname, true
}

// HasSslCommonname returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasSslCommonname() bool {
	if o != nil && o.SslCommonname != nil {
		return true
	}

	return false
}

// SetSslCommonname gets a reference to the given string and assigns it to the SslCommonname field.
func (o *NodeBalancerConfig) SetSslCommonname(v string) {
	o.SslCommonname = &v
}

// GetSslFingerprint returns the SslFingerprint field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetSslFingerprint() string {
	if o == nil || o.SslFingerprint == nil {
		var ret string
		return ret
	}
	return *o.SslFingerprint
}

// GetSslFingerprintOk returns a tuple with the SslFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetSslFingerprintOk() (*string, bool) {
	if o == nil || o.SslFingerprint == nil {
		return nil, false
	}
	return o.SslFingerprint, true
}

// HasSslFingerprint returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasSslFingerprint() bool {
	if o != nil && o.SslFingerprint != nil {
		return true
	}

	return false
}

// SetSslFingerprint gets a reference to the given string and assigns it to the SslFingerprint field.
func (o *NodeBalancerConfig) SetSslFingerprint(v string) {
	o.SslFingerprint = &v
}

// GetSslCert returns the SslCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeBalancerConfig) GetSslCert() string {
	if o == nil || o.SslCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslCert.Get()
}

// GetSslCertOk returns a tuple with the SslCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeBalancerConfig) GetSslCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslCert.Get(), o.SslCert.IsSet()
}

// HasSslCert returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasSslCert() bool {
	if o != nil && o.SslCert.IsSet() {
		return true
	}

	return false
}

// SetSslCert gets a reference to the given NullableString and assigns it to the SslCert field.
func (o *NodeBalancerConfig) SetSslCert(v string) {
	o.SslCert.Set(&v)
}
// SetSslCertNil sets the value for SslCert to be an explicit nil
func (o *NodeBalancerConfig) SetSslCertNil() {
	o.SslCert.Set(nil)
}

// UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
func (o *NodeBalancerConfig) UnsetSslCert() {
	o.SslCert.Unset()
}

// GetSslKey returns the SslKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeBalancerConfig) GetSslKey() string {
	if o == nil || o.SslKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SslKey.Get()
}

// GetSslKeyOk returns a tuple with the SslKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeBalancerConfig) GetSslKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SslKey.Get(), o.SslKey.IsSet()
}

// HasSslKey returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasSslKey() bool {
	if o != nil && o.SslKey.IsSet() {
		return true
	}

	return false
}

// SetSslKey gets a reference to the given NullableString and assigns it to the SslKey field.
func (o *NodeBalancerConfig) SetSslKey(v string) {
	o.SslKey.Set(&v)
}
// SetSslKeyNil sets the value for SslKey to be an explicit nil
func (o *NodeBalancerConfig) SetSslKeyNil() {
	o.SslKey.Set(nil)
}

// UnsetSslKey ensures that no value is present for SslKey, not even an explicit nil
func (o *NodeBalancerConfig) UnsetSslKey() {
	o.SslKey.Unset()
}

// GetNodesStatus returns the NodesStatus field value if set, zero value otherwise.
func (o *NodeBalancerConfig) GetNodesStatus() NodeBalancerConfigNodesStatus {
	if o == nil || o.NodesStatus == nil {
		var ret NodeBalancerConfigNodesStatus
		return ret
	}
	return *o.NodesStatus
}

// GetNodesStatusOk returns a tuple with the NodesStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancerConfig) GetNodesStatusOk() (*NodeBalancerConfigNodesStatus, bool) {
	if o == nil || o.NodesStatus == nil {
		return nil, false
	}
	return o.NodesStatus, true
}

// HasNodesStatus returns a boolean if a field has been set.
func (o *NodeBalancerConfig) HasNodesStatus() bool {
	if o != nil && o.NodesStatus != nil {
		return true
	}

	return false
}

// SetNodesStatus gets a reference to the given NodeBalancerConfigNodesStatus and assigns it to the NodesStatus field.
func (o *NodeBalancerConfig) SetNodesStatus(v NodeBalancerConfigNodesStatus) {
	o.NodesStatus = &v
}

func (o NodeBalancerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Stickiness != nil {
		toSerialize["stickiness"] = o.Stickiness
	}
	if o.Check != nil {
		toSerialize["check"] = o.Check
	}
	if o.CheckInterval != nil {
		toSerialize["check_interval"] = o.CheckInterval
	}
	if o.CheckTimeout != nil {
		toSerialize["check_timeout"] = o.CheckTimeout
	}
	if o.CheckAttempts != nil {
		toSerialize["check_attempts"] = o.CheckAttempts
	}
	if o.CheckPath != nil {
		toSerialize["check_path"] = o.CheckPath
	}
	if o.CheckBody != nil {
		toSerialize["check_body"] = o.CheckBody
	}
	if o.CheckPassive != nil {
		toSerialize["check_passive"] = o.CheckPassive
	}
	if o.ProxyProtocol != nil {
		toSerialize["proxy_protocol"] = o.ProxyProtocol
	}
	if o.CipherSuite != nil {
		toSerialize["cipher_suite"] = o.CipherSuite
	}
	if o.NodebalancerId != nil {
		toSerialize["nodebalancer_id"] = o.NodebalancerId
	}
	if o.SslCommonname != nil {
		toSerialize["ssl_commonname"] = o.SslCommonname
	}
	if o.SslFingerprint != nil {
		toSerialize["ssl_fingerprint"] = o.SslFingerprint
	}
	if o.SslCert.IsSet() {
		toSerialize["ssl_cert"] = o.SslCert.Get()
	}
	if o.SslKey.IsSet() {
		toSerialize["ssl_key"] = o.SslKey.Get()
	}
	if o.NodesStatus != nil {
		toSerialize["nodes_status"] = o.NodesStatus
	}
	return json.Marshal(toSerialize)
}

type NullableNodeBalancerConfig struct {
	value *NodeBalancerConfig
	isSet bool
}

func (v NullableNodeBalancerConfig) Get() *NodeBalancerConfig {
	return v.value
}

func (v *NullableNodeBalancerConfig) Set(val *NodeBalancerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBalancerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBalancerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBalancerConfig(val *NodeBalancerConfig) *NullableNodeBalancerConfig {
	return &NullableNodeBalancerConfig{value: val, isSet: true}
}

func (v NullableNodeBalancerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBalancerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


