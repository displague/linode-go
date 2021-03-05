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
	"time"
)

// Linode struct for Linode
type Linode struct {
	// The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned. Linode labels have the following constraints:    * Must start with an alpha character.   * May only consist of alphanumeric characters, dashes (`-`), underscores (`_`) or periods (`.`).   * Cannot have two dashes (`--`), underscores (`__`) or periods (`..`) in a row. 
	Label interface{} `json:"label,omitempty"`
	// This is the location where the Linode was deployed. A Linode's region can only be changed by initiating a [cross data center migration](/docs/api/linode-instances/#dc-migrationpending-host-migration-initiate). 
	Region interface{} `json:"region,omitempty"`
	Image NullableImage `json:"image,omitempty"`
	// This is the [Linode Type](/docs/api/linode-types/#types-list) that this Linode was deployed with. To change a Linode's Type, use [POST /linode/instances/{linodeId}/resize](/docs/api/linode-instances/#linode-resize). 
	Type interface{} `json:"type,omitempty"`
	// A deprecated property denoting a group label for this Linode. 
	Group *string `json:"group,omitempty"`
	// An array of tags applied to this object.  Tags are for organizational purposes only. 
	Tags *[]string `json:"tags,omitempty"`
	// This Linode's ID which must be provided for all operations impacting this Linode. 
	Id *int32 `json:"id,omitempty"`
	// A brief description of this Linode's current state. This field may change without direct action from you. For example, when a Linode goes into maintenance mode its status will display \"stopped\". 
	Status *string `json:"status,omitempty"`
	// The virtualization software powering this Linode. 
	Hypervisor *string `json:"hypervisor,omitempty"`
	// When this Linode was created.
	Created *time.Time `json:"created,omitempty"`
	// When this Linode was last updated.
	Updated *time.Time `json:"updated,omitempty"`
	// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to [open a support ticket](/docs/api/support/#support-ticket-open) to get additional IPv4 addresses.  IPv4 addresses may be reassigned between your Linodes, or shared with other Linodes. See the [/networking](/docs/api/networking/) endpoints for details. 
	Ipv4 *[]string `json:"ipv4,omitempty"`
	// This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared. If the Linode has not been assigned an IPv6 address, the return value will be `null`. 
	Ipv6 *string `json:"ipv6,omitempty"`
	// Information about the resources available to this Linode.
	Specs *map[string]interface{} `json:"specs,omitempty"`
	Alerts *map[string]interface{} `json:"alerts,omitempty"`
	// Information about this Linode's backups status. For information about available backups, see [/linode/instances/{linodeId}/backups](/docs/api/linode-instances/#backups-list). 
	Backups *map[string]interface{} `json:"backups,omitempty"`
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes. 
	WatchdogEnabled *bool `json:"watchdog_enabled,omitempty"`
}

// NewLinode instantiates a new Linode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinode() *Linode {
	this := Linode{}
	return &this
}

// NewLinodeWithDefaults instantiates a new Linode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinodeWithDefaults() *Linode {
	this := Linode{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode) GetLabel() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode) GetLabelOk() (*interface{}, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return &o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Linode) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given interface{} and assigns it to the Label field.
func (o *Linode) SetLabel(v interface{}) {
	o.Label = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode) GetRegion() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode) GetRegionOk() (*interface{}, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return &o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Linode) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given interface{} and assigns it to the Region field.
func (o *Linode) SetRegion(v interface{}) {
	o.Region = v
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode) GetImage() Image {
	if o == nil || o.Image.Get() == nil {
		var ret Image
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode) GetImageOk() (*Image, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *Linode) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableImage and assigns it to the Image field.
func (o *Linode) SetImage(v Image) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *Linode) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *Linode) UnsetImage() {
	o.Image.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode) GetType() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Linode) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *Linode) SetType(v interface{}) {
	o.Type = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Linode) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Linode) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *Linode) SetGroup(v string) {
	o.Group = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Linode) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Linode) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Linode) SetTags(v []string) {
	o.Tags = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Linode) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Linode) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Linode) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Linode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Linode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Linode) SetStatus(v string) {
	o.Status = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *Linode) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *Linode) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *Linode) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Linode) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Linode) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Linode) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Linode) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Linode) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Linode) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *Linode) GetIpv4() []string {
	if o == nil || o.Ipv4 == nil {
		var ret []string
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetIpv4Ok() (*[]string, bool) {
	if o == nil || o.Ipv4 == nil {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *Linode) HasIpv4() bool {
	if o != nil && o.Ipv4 != nil {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given []string and assigns it to the Ipv4 field.
func (o *Linode) SetIpv4(v []string) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *Linode) GetIpv6() string {
	if o == nil || o.Ipv6 == nil {
		var ret string
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetIpv6Ok() (*string, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *Linode) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given string and assigns it to the Ipv6 field.
func (o *Linode) SetIpv6(v string) {
	o.Ipv6 = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *Linode) GetSpecs() map[string]interface{} {
	if o == nil || o.Specs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetSpecsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Specs == nil {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *Linode) HasSpecs() bool {
	if o != nil && o.Specs != nil {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given map[string]interface{} and assigns it to the Specs field.
func (o *Linode) SetSpecs(v map[string]interface{}) {
	o.Specs = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *Linode) GetAlerts() map[string]interface{} {
	if o == nil || o.Alerts == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetAlertsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *Linode) HasAlerts() bool {
	if o != nil && o.Alerts != nil {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given map[string]interface{} and assigns it to the Alerts field.
func (o *Linode) SetAlerts(v map[string]interface{}) {
	o.Alerts = &v
}

// GetBackups returns the Backups field value if set, zero value otherwise.
func (o *Linode) GetBackups() map[string]interface{} {
	if o == nil || o.Backups == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetBackupsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Backups == nil {
		return nil, false
	}
	return o.Backups, true
}

// HasBackups returns a boolean if a field has been set.
func (o *Linode) HasBackups() bool {
	if o != nil && o.Backups != nil {
		return true
	}

	return false
}

// SetBackups gets a reference to the given map[string]interface{} and assigns it to the Backups field.
func (o *Linode) SetBackups(v map[string]interface{}) {
	o.Backups = &v
}

// GetWatchdogEnabled returns the WatchdogEnabled field value if set, zero value otherwise.
func (o *Linode) GetWatchdogEnabled() bool {
	if o == nil || o.WatchdogEnabled == nil {
		var ret bool
		return ret
	}
	return *o.WatchdogEnabled
}

// GetWatchdogEnabledOk returns a tuple with the WatchdogEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode) GetWatchdogEnabledOk() (*bool, bool) {
	if o == nil || o.WatchdogEnabled == nil {
		return nil, false
	}
	return o.WatchdogEnabled, true
}

// HasWatchdogEnabled returns a boolean if a field has been set.
func (o *Linode) HasWatchdogEnabled() bool {
	if o != nil && o.WatchdogEnabled != nil {
		return true
	}

	return false
}

// SetWatchdogEnabled gets a reference to the given bool and assigns it to the WatchdogEnabled field.
func (o *Linode) SetWatchdogEnabled(v bool) {
	o.WatchdogEnabled = &v
}

func (o Linode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Hypervisor != nil {
		toSerialize["hypervisor"] = o.Hypervisor
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Ipv4 != nil {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	if o.Specs != nil {
		toSerialize["specs"] = o.Specs
	}
	if o.Alerts != nil {
		toSerialize["alerts"] = o.Alerts
	}
	if o.Backups != nil {
		toSerialize["backups"] = o.Backups
	}
	if o.WatchdogEnabled != nil {
		toSerialize["watchdog_enabled"] = o.WatchdogEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableLinode struct {
	value *Linode
	isSet bool
}

func (v NullableLinode) Get() *Linode {
	return v.value
}

func (v *NullableLinode) Set(val *Linode) {
	v.value = val
	v.isSet = true
}

func (v NullableLinode) IsSet() bool {
	return v.isSet
}

func (v *NullableLinode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinode(val *Linode) *NullableLinode {
	return &NullableLinode{value: val, isSet: true}
}

func (v NullableLinode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


