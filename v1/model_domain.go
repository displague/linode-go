/*
 * Linode API
 *
 * ## Introduction The Linode API provides the ability to programmatically manage the full range of Linode products and services.  This reference is designed to assist application developers and system administrators.  Each endpoint includes descriptions, request syntax, and examples using standard HTTP requests. Response data is returned in JSON format.   This document was generated from our OpenAPI Specification.  See the <a target=\"_top\" href=\"https://www.openapis.org\">OpenAPI website</a> for more information.  <a target=\"_top\" href=\"/docs/api/openapi.yaml\">Download the Linode OpenAPI Specification</a>.   ## Changelog  <a target=\"_top\" href=\"https://developers.linode.com/changelog\">View our Changelog</a> to see release notes on all changes made to our API.  ## Access and Authentication  Some endpoints are publicly accessible without requiring authentication. All endpoints affecting your Account, however, require either a Personal Access Token or OAuth authentication (when using third-party applications).  ### Personal Access Token  The easiest way to access the API is with a Personal Access Token (PAT) generated from the <a target=\"_top\" href=\"https://cloud.linode.com/profile/tokens\">Linode Cloud Manager</a> or the [Create Personal Access Token](/docs/api/profile/#personal-access-token-create) endpoint.  All scopes for the OAuth security model ([defined below](/docs/api/profile/#oauth)) apply to this security model as well.  #### Authentication  | Security Scheme Type: | HTTP | |-----------------------|------| | **HTTP Authorization Scheme** | bearer |  ### OAuth If you only need to access the Linode API for personal use, we recommend that you create a [personal access token](/docs/api/#personal-access-token). If you're designing an application that can authenticate with an arbitrary Linode user, then you should use the OAuth 2.0 workflows presented in this section.  For a more detailed example of an OAuth 2.0 implementation, see our guide on [How to Create an OAuth App with the Linode Python API Library](/docs/platform/api/how-to-create-an-oauth-app-with-the-linode-python-api-library/#oauth-2-authentication-exchange).  Before you implement OAuth in your application, you first need to create an OAuth client. You can do this [with the Linode API](/docs/api/account/#oauth-client-create) or [via the Cloud Manager](https://cloud.linode.com/profile/clients):    - When creating the client, you'll supply a `label` and a `redirect_uri` (referred to as the Callback URL in the Cloud Manager).   - The response from this endpoint will give you a `client_id` and a `secret`.   - Clients can be public or private, and are private by default. You can choose to make the client public when it is created.     - A private client is used with applications which can securely store the client secret (that is, the secret returned to you when you first created the client). For example, an application running on a secured server that only the developer has access to would use a private OAuth client. This is also called a confidential client in some OAuth documentation.     - A public client is used with applications where the client secret is not guaranteed to be secure. For example, a native app running on a user's computer may not be able to keep the client secret safe, as a user could potentially inspect the source of the application. So, native apps or apps that run in a user's browser should use a public client.     - Public and private clients follow different workflows, as described below.  #### OAuth Workflow  The OAuth workflow is a series of exchanges between your third-party app and Linode. The workflow is used to authenticate a user before an application can start making API calls on the user's behalf.  Notes:  - With respect to the diagram in [section 1.2 of RFC 6749](https://tools.ietf.org/html/rfc6749#section-1.2), login.linode.com (referred to in this section as the *login server*) is the Resource Owner and the Authorization Server; api.linode.com (referred to here as the *api server*) is the Resource Server. - The OAuth spec refers to the private and public workflows listed below as the [authorization code flow](https://tools.ietf.org/html/rfc6749#section-1.3.1) and [implicit flow](https://tools.ietf.org/html/rfc6749#section-1.3.2).  | PRIVATE WORKFLOW | PUBLIC WORKFLOW | |------------------|------------------| | 1.  The user visits the application's website and is directed to login with Linode. | 1.  The user visits the application's website and is directed to login with Linode. | | 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. | 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. | | 3.  The user logs into the login server with their username and password. | 3.  The user logs into the login server with their username and password. | | 4.  The login server redirects the user to the specificed redirect URL with a temporary authorization `code` (exchange code) in the URL. | 4.  The login server redirects the user back to your application with an OAuth `access_token` embedded in the redirect URL's hash. This is temporary and expires in two hours. No `refresh_token` is issued. Therefore, once the `access_token` expires, a new one will need to be issued by having the user log in again. | | 5.  The application issues a POST request (*see below*) to the login server with the exchange code, `client_id`, and the client application's `client_secret`. | | | 6.  The login server responds to the client application with a new OAuth `access_token` and `refresh_token`. The `access_token` is set to expire in two hours. | | | 7.  The `refresh_token` can be used by contacting the login server with the `client_id`, `client_secret`, `grant_type`, and `refresh_token` to get a new OAuth `access_token` and `refresh_token`. The new `access_token` is good for another two hours, and the new `refresh_token`, can be used to extend the session again by this same method. | |  #### OAuth Private Workflow - Additional Details  The following information expands on steps 5 through 7 of the private workflow:  Once the user has logged into Linode and you have received an exchange code, you will need to trade that exchange code for an `access_token` and `refresh_token`. You do this by making an HTTP POST request to the following address:  ``` https://login.linode.com/oauth/token ```  Make this request as `application/x-www-form-urlencoded` or as `multipart/form-data` and include the following parameters in the POST body:  | PARAMETER | DESCRIPTION | |-----------|-------------| | grant_type | The grant type you're using for renewal.  Currently only the string \"refresh_token\" is accepted. | | client_id | Your app's client ID. | | client_secret | Your app's client secret. | | code | The code you just received from the redirect. |  You'll get a response like this:  ```json {   \"scope\": \"linodes:read_write\",   \"access_token\": \"03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c\"   \"token_type\": \"bearer\",   \"expires_in\": 7200, } ```  Included in the reponse is an `access_token`. With this token, you can proceed to make authenticated HTTP requests to the API by adding this header to each request:  ``` Authorization: Bearer 03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c ```  #### OAuth Reference  | Security Scheme Type | OAuth 2.0 | |-----------------------|--------| | **Authorization URL** | https://login.linode.com/oauth/authorize | | **Token URL** | https://login.linode.com/oauth/token | | **Scopes** | <ul><li>`account:read_only` - Allows access to GET information about your Account.</li><li>`account:read_write` - Allows access to all endpoints related to your Account.</li><li>`domains:read_only` - Allows access to GET Domains on your Account.</li><li>`domains:read_write` - Allows access to all Domain endpoints.</li><li>`events:read_only` - Allows access to GET your Events.</li><li>`events:read_write` - Allows access to all endpoints related to your Events.</li><li>`firewall:read_only` - Allows access to GET information about your Firewalls.</li><li>`firewall:read_write` - Allows access to all Firewall endpoints.</li><li>`images:read_only` - Allows access to GET your Images.</li><li>`images:read_write` - Allows access to all endpoints related to your Images.</li><li>`ips:read_only` - Allows access to GET your ips.</li><li>`ips:read_write` - Allows access to all endpoints related to your ips.</li><li>`linodes:read_only` - Allows access to GET Linodes on your Account.</li><li>`linodes:read_write` - Allow access to all endpoints related to your Linodes.</li><li>`lke:read_only` - Allows access to GET LKE Clusters on your Account.</li><li>`lke:read_write` - Allows access to all endpoints related to LKE Clusters on your Account.</li><li>`longview:read_only` - Allows access to GET your Longview Clients.</li><li>`longview:read_write` - Allows access to all endpoints related to your Longview Clients.</li><li>`maintenance:read_only` - Allows access to GET information about Maintenance on your account.</li><li>`nodebalancers:read_only` - Allows access to GET NodeBalancers on your Account.</li><li>`nodebalancers:read_write` - Allows access to all NodeBalancer endpoints.</li><li>`object_storage:read_only` - Allows access to GET information related to your Object Storage.</li><li>`object_storage:read_write` - Allows access to all Object Storage endpoints.</li><li>`stackscripts:read_only` - Allows access to GET your StackScripts.</li><li>`stackscripts:read_write` - Allows access to all endpoints related to your StackScripts.</li><li>`volumes:read_only` - Allows access to GET your Volumes.</li><li>`volumes:read_write` - Allows access to all endpoints related to your Volumes.</li></ul><br/>|  ## Requests  Requests must be made over HTTPS to ensure transactions are encrypted. The following Request methods are supported:  | METHOD | USAGE | |--------|-------| | GET    | Retrieves data about collections and individual resources. | | POST   | For collections, creates a new resource of that type. Also used to perform actions on action endpoints. | | PUT    | Updates an existing resource. | | DELETE | Deletes a resource. This is a destructive action. |   ## Responses  Actions will return one following HTTP response status codes:  | STATUS  | DESCRIPTION | |---------|-------------| | 200 OK  | The request was successful. | | 204 No Content | The server successfully fulfilled the request and there is no additional content to send. | | 400 Bad Request | You submitted an invalid request (missing parameters, etc.). | | 401 Unauthorized | You failed to authenticate for this resource. | | 403 Forbidden | You are authenticated, but don't have permission to do this. | | 404 Not Found | The resource you're requesting does not exist. | | 429 Too Many Requests | You've hit a rate limit. | | 500 Internal Server Error | Please [open a Support Ticket](/docs/api/support/#support-ticket-open). |  ## Errors  Success is indicated via <a href=\"https://en.wikipedia.org/wiki/List_of_HTTP_status_codes\" target=\"_top\">Standard HTTP status codes</a>. `2xx` codes indicate success, `4xx` codes indicate a request error, and `5xx` errors indicate a server error. A request error might be an invalid input, a required parameter being omitted, or a malformed request. A server error means something went wrong processing your request. If this occurs, please [open a Support Ticket](/docs/api/support/#support-ticket-open) and let us know. Though errors are logged and we work quickly to resolve issues, opening a ticket and providing us with reproducable steps and data is always helpful.  The `errors` field is an array of the things that went wrong with your request. We will try to include as many of the problems in the response as possible, but it's conceivable that fixing these errors and resubmitting may result in new errors coming back once we are able to get further along in the process of handling your request.   Within each error object, the `field` parameter will be included if the error pertains to a specific field in the JSON you've submitted. This will be omitted if there is no relevant field. The `reason` is a human-readable explanation of the error, and will always be included.  ## Pagination  Resource lists are always paginated. The response will look similar to this:  ```json {     \"data\": [ ... ],     \"page\": 1,     \"pages\": 3,     \"results\": 300 } ```  * Pages start at 1. You may retrieve a specific page of results by adding `?page=x` to your URL (for example, `?page=4`). If the value of `page` exceeds `2^64/page_size`, the last possible page will be returned.   * Page sizes default to 100, and can be set to return between 25 and 500. Page size can be set using `?page_size=x`.  ## Filtering and Sorting  Collections are searchable by fields they include, marked in the spec as `x-linode-filterable: true`. Filters are passed in the `X-Filter` header and are formatted as JSON objects. Here is a request call for Linode Types in our \"standard\" class:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"class\": \"standard\"     }' ```  The filter object's keys are the keys of the object you're filtering, and the values are accepted values. You can add multiple filters by including more than one key. For example, filtering for \"standard\" Linode Types that offer one vcpu:  ```Shell  curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"class\": \"standard\",       \"vcpus\": 1     }' ```  In the above example, both filters are combined with an \"and\" operation. However, if you wanted either Types with one vcpu or Types in our \"standard\" class, you can add an operator:   ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"+or\": [         { \"vcpus\": 1 },         { \"class\": \"standard\" }       ]     }' ```  Each filter in the `+or` array is its own filter object, and all conditions in it are combined with an \"and\" operation as they were in the previous example.  Other operators are also available. Operators are keys of a Filter JSON object. Their value must be of the appropriate type, and they are evaluated as described below:  | OPERATOR | TYPE   | DESCRIPTION                       | |----------|--------|-----------------------------------| | +and     | array  | All conditions must be true.       | | +or      | array  | One condition must be true.        | | +gt      | number | Value must be greater than number. | | +gte     | number | Value must be greater than or equal to number. | | +lt      | number | Value must be less than number. | | +lte     | number | Value must be less than or equal to number. | | +contains | string | Given string must be in the value. | | +neq      | string | Does not equal the value.          | | +order_by | string | Attribute to order the results by - must be filterable. | | +order    | string | Either \"asc\" or \"desc\". Defaults to \"asc\". Requires `+order_by`. |  For example, filtering for [Linode Types](/docs/api/linode-types/) that offer memory equal to or higher than 61440:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"memory\": {         \"+gte\": 61440       }     }' ```  You can combine and nest operators to construct arbitrarily-complex queries. For example, give me all [Linode Types](/docs/api/linode-types/) which are either `standard` or `highmem` class, or have between 12 and 20 vcpus:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H '     X-Filter: {       \"+or\": [         {           \"+or\": [             {               \"class\": \"standard\"             },             {               \"class\": \"highmem\"             }           ]         },         {           \"+and\": [             {               \"vcpus\": {                 \"+gte\": 12               }             },             {               \"vcpus\": {                 \"+lte\": 20               }             }           ]         }       ]     }' ```  ## Rate Limiting  With the Linode API, you can make up to 1,600 general API requests every two minutes per user as determined by IP adddress or by OAuth token. Additionally, there are endpoint specfic limits defined below.  **Note:** There may be rate limiting applied at other levels outside of the API, for example, at the load balancer.  `/stats` endpoints have their own dedicated limits of 100 requests per minute per user. These endpoints are:  * [View Linode Statistics](/docs/api/linode-instances/#linode-statistics-view) * [View Linode Statistics (year/month)](/docs/api/linode-instances/#statistics-yearmonth-view) * [View NodeBalancer Statistics](/docs/api/nodebalancers/#nodebalancer-statistics-view) * [List Managed Stats](/docs/api/managed/#managed-stats-list)  Object Storage endpoints have a dedicated limit of 750 requests per second per user. The Object Storage endpoints are:  * [Object Storage Endpoints](/docs/api/object-storage/)  Opening Support Tickets has a dedicated limit of 2 requests per minute per user. That endpoint is:  * [Open Support Ticket](/docs/api/support/#support-ticket-open)  Accepting Entity Transfers has a dedicated limit of 2 requests per minute per user. That endpoint is:  * [Entity Transfer Accept](/docs/api/account/#entity-transfer-accept)  ## CLI (Command Line Interface)  The <a href=\"https://github.com/linode/linode-cli\" target=\"_top\">Linode CLI</a> allows you to easily work with the API using intuitive and simple syntax. It requires a [Personal Access Token](/docs/api/#personal-access-token) for authentication, and gives you access to all of the features and functionality of the Linode API that are documented here with CLI examples.  Endpoints that do not have CLI examples are currently unavailable through the CLI, but can be accessed via other methods such as Shell commands and other third-party applications. 
 *
 * API version: 4.85.0
 * Contact: support@linode.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Domain A domain zonefile in our DNS system.  You must own the domain name and tell your registrar to use Linode's nameservers in order for a domain in our system to be treated as authoritative. 
type Domain struct {
	// This Domain's unique ID
	Id *int32 `json:"id,omitempty"`
	// Whether this Domain represents the authoritative source of information for the domain it describes (\"master\"), or whether it is a read-only copy of a master (\"slave\"). 
	Type *string `json:"type,omitempty"`
	// The domain this Domain represents. Domain labels cannot be longer than 63 characters and must conform to [RFC1035](https://tools.ietf.org/html/rfc1035). Domains must be unique on Linode's platform, including across different Linode accounts; there cannot be two Domains representing the same domain. 
	Domain *string `json:"domain,omitempty"`
	// The group this Domain belongs to.  This is for display purposes only. 
	Group *string `json:"group,omitempty"`
	// Used to control whether this Domain is currently being rendered. 
	Status *string `json:"status,omitempty"`
	// A description for this Domain. This is for display purposes only. 
	Description *string `json:"description,omitempty"`
	// Start of Authority email address. This is required for `type` master Domains. 
	SoaEmail *string `json:"soa_email,omitempty"`
	// The interval, in seconds, at which a failed refresh should be retried.  * Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  * Any other value is rounded up to the nearest valid value.  * A value of 0 is equivalent to the default value of 14400. 
	RetrySec *int32 `json:"retry_sec,omitempty"`
	// The IP addresses representing the master DNS for this Domain. At least one value is required for `type` slave Domains. 
	MasterIps *[]string `json:"master_ips,omitempty"`
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it. 
	AxfrIps *[]string `json:"axfr_ips,omitempty"`
	// The amount of time in seconds that may pass before this Domain is no longer authoritative.  * Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  * Any other value is rounded up to the nearest valid value.  * A value of 0 is equivalent to the default value of 1209600. 
	ExpireSec *int32 `json:"expire_sec,omitempty"`
	// The amount of time in seconds before this Domain should be refreshed.  * Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  * Any other value is rounded up to the nearest valid value.  * A value of 0 is equivalent to the default value of 14400. 
	RefreshSec *int32 `json:"refresh_sec,omitempty"`
	// \"Time to Live\" - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. * Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. * Any other value is rounded up to the nearest valid value. * A value of 0 is equivalent to the default value of 86400. 
	TtlSec *int32 `json:"ttl_sec,omitempty"`
	// An array of tags applied to this object.  Tags are for organizational purposes only. 
	Tags *[]string `json:"tags,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	var status string = "active"
	this.Status = &status
	var retrySec int32 = 0
	this.RetrySec = &retrySec
	var expireSec int32 = 0
	this.ExpireSec = &expireSec
	var refreshSec int32 = 0
	this.RefreshSec = &refreshSec
	var ttlSec int32 = 0
	this.TtlSec = &ttlSec
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	var status string = "active"
	this.Status = &status
	var retrySec int32 = 0
	this.RetrySec = &retrySec
	var expireSec int32 = 0
	this.ExpireSec = &expireSec
	var refreshSec int32 = 0
	this.RefreshSec = &refreshSec
	var ttlSec int32 = 0
	this.TtlSec = &ttlSec
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Domain) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Domain) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Domain) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Domain) SetType(v string) {
	o.Type = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Domain) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Domain) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Domain) SetDomain(v string) {
	o.Domain = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Domain) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Domain) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *Domain) SetGroup(v string) {
	o.Group = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Domain) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Domain) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Domain) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Domain) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Domain) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Domain) SetDescription(v string) {
	o.Description = &v
}

// GetSoaEmail returns the SoaEmail field value if set, zero value otherwise.
func (o *Domain) GetSoaEmail() string {
	if o == nil || o.SoaEmail == nil {
		var ret string
		return ret
	}
	return *o.SoaEmail
}

// GetSoaEmailOk returns a tuple with the SoaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSoaEmailOk() (*string, bool) {
	if o == nil || o.SoaEmail == nil {
		return nil, false
	}
	return o.SoaEmail, true
}

// HasSoaEmail returns a boolean if a field has been set.
func (o *Domain) HasSoaEmail() bool {
	if o != nil && o.SoaEmail != nil {
		return true
	}

	return false
}

// SetSoaEmail gets a reference to the given string and assigns it to the SoaEmail field.
func (o *Domain) SetSoaEmail(v string) {
	o.SoaEmail = &v
}

// GetRetrySec returns the RetrySec field value if set, zero value otherwise.
func (o *Domain) GetRetrySec() int32 {
	if o == nil || o.RetrySec == nil {
		var ret int32
		return ret
	}
	return *o.RetrySec
}

// GetRetrySecOk returns a tuple with the RetrySec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRetrySecOk() (*int32, bool) {
	if o == nil || o.RetrySec == nil {
		return nil, false
	}
	return o.RetrySec, true
}

// HasRetrySec returns a boolean if a field has been set.
func (o *Domain) HasRetrySec() bool {
	if o != nil && o.RetrySec != nil {
		return true
	}

	return false
}

// SetRetrySec gets a reference to the given int32 and assigns it to the RetrySec field.
func (o *Domain) SetRetrySec(v int32) {
	o.RetrySec = &v
}

// GetMasterIps returns the MasterIps field value if set, zero value otherwise.
func (o *Domain) GetMasterIps() []string {
	if o == nil || o.MasterIps == nil {
		var ret []string
		return ret
	}
	return *o.MasterIps
}

// GetMasterIpsOk returns a tuple with the MasterIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetMasterIpsOk() (*[]string, bool) {
	if o == nil || o.MasterIps == nil {
		return nil, false
	}
	return o.MasterIps, true
}

// HasMasterIps returns a boolean if a field has been set.
func (o *Domain) HasMasterIps() bool {
	if o != nil && o.MasterIps != nil {
		return true
	}

	return false
}

// SetMasterIps gets a reference to the given []string and assigns it to the MasterIps field.
func (o *Domain) SetMasterIps(v []string) {
	o.MasterIps = &v
}

// GetAxfrIps returns the AxfrIps field value if set, zero value otherwise.
func (o *Domain) GetAxfrIps() []string {
	if o == nil || o.AxfrIps == nil {
		var ret []string
		return ret
	}
	return *o.AxfrIps
}

// GetAxfrIpsOk returns a tuple with the AxfrIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetAxfrIpsOk() (*[]string, bool) {
	if o == nil || o.AxfrIps == nil {
		return nil, false
	}
	return o.AxfrIps, true
}

// HasAxfrIps returns a boolean if a field has been set.
func (o *Domain) HasAxfrIps() bool {
	if o != nil && o.AxfrIps != nil {
		return true
	}

	return false
}

// SetAxfrIps gets a reference to the given []string and assigns it to the AxfrIps field.
func (o *Domain) SetAxfrIps(v []string) {
	o.AxfrIps = &v
}

// GetExpireSec returns the ExpireSec field value if set, zero value otherwise.
func (o *Domain) GetExpireSec() int32 {
	if o == nil || o.ExpireSec == nil {
		var ret int32
		return ret
	}
	return *o.ExpireSec
}

// GetExpireSecOk returns a tuple with the ExpireSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetExpireSecOk() (*int32, bool) {
	if o == nil || o.ExpireSec == nil {
		return nil, false
	}
	return o.ExpireSec, true
}

// HasExpireSec returns a boolean if a field has been set.
func (o *Domain) HasExpireSec() bool {
	if o != nil && o.ExpireSec != nil {
		return true
	}

	return false
}

// SetExpireSec gets a reference to the given int32 and assigns it to the ExpireSec field.
func (o *Domain) SetExpireSec(v int32) {
	o.ExpireSec = &v
}

// GetRefreshSec returns the RefreshSec field value if set, zero value otherwise.
func (o *Domain) GetRefreshSec() int32 {
	if o == nil || o.RefreshSec == nil {
		var ret int32
		return ret
	}
	return *o.RefreshSec
}

// GetRefreshSecOk returns a tuple with the RefreshSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRefreshSecOk() (*int32, bool) {
	if o == nil || o.RefreshSec == nil {
		return nil, false
	}
	return o.RefreshSec, true
}

// HasRefreshSec returns a boolean if a field has been set.
func (o *Domain) HasRefreshSec() bool {
	if o != nil && o.RefreshSec != nil {
		return true
	}

	return false
}

// SetRefreshSec gets a reference to the given int32 and assigns it to the RefreshSec field.
func (o *Domain) SetRefreshSec(v int32) {
	o.RefreshSec = &v
}

// GetTtlSec returns the TtlSec field value if set, zero value otherwise.
func (o *Domain) GetTtlSec() int32 {
	if o == nil || o.TtlSec == nil {
		var ret int32
		return ret
	}
	return *o.TtlSec
}

// GetTtlSecOk returns a tuple with the TtlSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTtlSecOk() (*int32, bool) {
	if o == nil || o.TtlSec == nil {
		return nil, false
	}
	return o.TtlSec, true
}

// HasTtlSec returns a boolean if a field has been set.
func (o *Domain) HasTtlSec() bool {
	if o != nil && o.TtlSec != nil {
		return true
	}

	return false
}

// SetTtlSec gets a reference to the given int32 and assigns it to the TtlSec field.
func (o *Domain) SetTtlSec(v int32) {
	o.TtlSec = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Domain) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Domain) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Domain) SetTags(v []string) {
	o.Tags = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.SoaEmail != nil {
		toSerialize["soa_email"] = o.SoaEmail
	}
	if o.RetrySec != nil {
		toSerialize["retry_sec"] = o.RetrySec
	}
	if o.MasterIps != nil {
		toSerialize["master_ips"] = o.MasterIps
	}
	if o.AxfrIps != nil {
		toSerialize["axfr_ips"] = o.AxfrIps
	}
	if o.ExpireSec != nil {
		toSerialize["expire_sec"] = o.ExpireSec
	}
	if o.RefreshSec != nil {
		toSerialize["refresh_sec"] = o.RefreshSec
	}
	if o.TtlSec != nil {
		toSerialize["ttl_sec"] = o.TtlSec
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


