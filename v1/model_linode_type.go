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

// LinodeType Returns collection of Linode types, including pricing and specifications for each type. These are used when [creating](/docs/api/linode-instances/#linode-create) or [resizing](/docs/api/linode-instances/#linode-resize) Linodes. 
type LinodeType struct {
	// The ID representing the Linode Type.
	Id *string `json:"id,omitempty"`
	// The Linode Type's label is for display purposes only. 
	Label *string `json:"label,omitempty"`
	// The Disk size, in MB, of the Linode Type. 
	Disk *int32 `json:"disk,omitempty"`
	// The class of the Linode Type. We currently offer five classes of Linodes:    * nanode - Nanode instances are good for low-duty workloads,     where performance isn't critical. **Note:** As of June 16th, 2020, Nanodes became     1 GB Linodes in the Cloud Manager, however, the API, the CLI, and billing will     continue to refer to these instances as Nanodes.   * standard - Standard Shared instances are good for medium-duty workloads and     are a good mix of performance, resources, and price. **Note:** As of June 16th, 2020,     Standard Linodes in the Cloud Manager became Shared Linodes, however, the API, the CLI, and     billing will continue to refer to these instances as Standard Linodes.   * dedicated - Dedicated CPU instances are good for full-duty workloads     where consistent performance is important.   * gpu - Linodes with dedicated NVIDIA Quadro &reg; RTX 6000 GPUs accelerate highly     specialized applications such as machine learning, AI, and video transcoding.   * highmem - High Memory instances favor RAM over other resources, and can be     good for memory hungry use cases like caching and in-memory databases.     All High Memory plans contain dedicated CPU cores. 
	Class *string `json:"class,omitempty"`
	Price *LinodeTypePrice `json:"price,omitempty"`
	Addons *LinodeTypeAddons `json:"addons,omitempty"`
	// The Mbits outbound bandwidth allocation. 
	NetworkOut *int32 `json:"network_out,omitempty"`
	// Amount of RAM included in this Linode Type. 
	Memory *int32 `json:"memory,omitempty"`
	// The Linode Type that a [mutate](/docs/api/linode-instances/#linode-upgrade) will upgrade to for a Linode of this type.  If \"null\", a Linode of this type may not mutate. 
	Successor NullableString `json:"successor,omitempty"`
	// The monthly outbound transfer amount, in MB. 
	Transfer *int32 `json:"transfer,omitempty"`
	// The number of VCPU cores this Linode Type offers. 
	Vcpus *int32 `json:"vcpus,omitempty"`
	// The number of GPUs this Linode Type offers. 
	Gpus *int32 `json:"gpus,omitempty"`
}

// NewLinodeType instantiates a new LinodeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinodeType() *LinodeType {
	this := LinodeType{}
	return &this
}

// NewLinodeTypeWithDefaults instantiates a new LinodeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinodeTypeWithDefaults() *LinodeType {
	this := LinodeType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LinodeType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinodeType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LinodeType) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LinodeType) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LinodeType) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *LinodeType) SetLabel(v string) {
	o.Label = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *LinodeType) GetDisk() int32 {
	if o == nil || o.Disk == nil {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetDiskOk() (*int32, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *LinodeType) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *LinodeType) SetDisk(v int32) {
	o.Disk = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *LinodeType) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *LinodeType) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *LinodeType) SetClass(v string) {
	o.Class = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *LinodeType) GetPrice() LinodeTypePrice {
	if o == nil || o.Price == nil {
		var ret LinodeTypePrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetPriceOk() (*LinodeTypePrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *LinodeType) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given LinodeTypePrice and assigns it to the Price field.
func (o *LinodeType) SetPrice(v LinodeTypePrice) {
	o.Price = &v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *LinodeType) GetAddons() LinodeTypeAddons {
	if o == nil || o.Addons == nil {
		var ret LinodeTypeAddons
		return ret
	}
	return *o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetAddonsOk() (*LinodeTypeAddons, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *LinodeType) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given LinodeTypeAddons and assigns it to the Addons field.
func (o *LinodeType) SetAddons(v LinodeTypeAddons) {
	o.Addons = &v
}

// GetNetworkOut returns the NetworkOut field value if set, zero value otherwise.
func (o *LinodeType) GetNetworkOut() int32 {
	if o == nil || o.NetworkOut == nil {
		var ret int32
		return ret
	}
	return *o.NetworkOut
}

// GetNetworkOutOk returns a tuple with the NetworkOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetNetworkOutOk() (*int32, bool) {
	if o == nil || o.NetworkOut == nil {
		return nil, false
	}
	return o.NetworkOut, true
}

// HasNetworkOut returns a boolean if a field has been set.
func (o *LinodeType) HasNetworkOut() bool {
	if o != nil && o.NetworkOut != nil {
		return true
	}

	return false
}

// SetNetworkOut gets a reference to the given int32 and assigns it to the NetworkOut field.
func (o *LinodeType) SetNetworkOut(v int32) {
	o.NetworkOut = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *LinodeType) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *LinodeType) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *LinodeType) SetMemory(v int32) {
	o.Memory = &v
}

// GetSuccessor returns the Successor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinodeType) GetSuccessor() string {
	if o == nil || o.Successor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Successor.Get()
}

// GetSuccessorOk returns a tuple with the Successor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinodeType) GetSuccessorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Successor.Get(), o.Successor.IsSet()
}

// HasSuccessor returns a boolean if a field has been set.
func (o *LinodeType) HasSuccessor() bool {
	if o != nil && o.Successor.IsSet() {
		return true
	}

	return false
}

// SetSuccessor gets a reference to the given NullableString and assigns it to the Successor field.
func (o *LinodeType) SetSuccessor(v string) {
	o.Successor.Set(&v)
}
// SetSuccessorNil sets the value for Successor to be an explicit nil
func (o *LinodeType) SetSuccessorNil() {
	o.Successor.Set(nil)
}

// UnsetSuccessor ensures that no value is present for Successor, not even an explicit nil
func (o *LinodeType) UnsetSuccessor() {
	o.Successor.Unset()
}

// GetTransfer returns the Transfer field value if set, zero value otherwise.
func (o *LinodeType) GetTransfer() int32 {
	if o == nil || o.Transfer == nil {
		var ret int32
		return ret
	}
	return *o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetTransferOk() (*int32, bool) {
	if o == nil || o.Transfer == nil {
		return nil, false
	}
	return o.Transfer, true
}

// HasTransfer returns a boolean if a field has been set.
func (o *LinodeType) HasTransfer() bool {
	if o != nil && o.Transfer != nil {
		return true
	}

	return false
}

// SetTransfer gets a reference to the given int32 and assigns it to the Transfer field.
func (o *LinodeType) SetTransfer(v int32) {
	o.Transfer = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise.
func (o *LinodeType) GetVcpus() int32 {
	if o == nil || o.Vcpus == nil {
		var ret int32
		return ret
	}
	return *o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetVcpusOk() (*int32, bool) {
	if o == nil || o.Vcpus == nil {
		return nil, false
	}
	return o.Vcpus, true
}

// HasVcpus returns a boolean if a field has been set.
func (o *LinodeType) HasVcpus() bool {
	if o != nil && o.Vcpus != nil {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given int32 and assigns it to the Vcpus field.
func (o *LinodeType) SetVcpus(v int32) {
	o.Vcpus = &v
}

// GetGpus returns the Gpus field value if set, zero value otherwise.
func (o *LinodeType) GetGpus() int32 {
	if o == nil || o.Gpus == nil {
		var ret int32
		return ret
	}
	return *o.Gpus
}

// GetGpusOk returns a tuple with the Gpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeType) GetGpusOk() (*int32, bool) {
	if o == nil || o.Gpus == nil {
		return nil, false
	}
	return o.Gpus, true
}

// HasGpus returns a boolean if a field has been set.
func (o *LinodeType) HasGpus() bool {
	if o != nil && o.Gpus != nil {
		return true
	}

	return false
}

// SetGpus gets a reference to the given int32 and assigns it to the Gpus field.
func (o *LinodeType) SetGpus(v int32) {
	o.Gpus = &v
}

func (o LinodeType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Addons != nil {
		toSerialize["addons"] = o.Addons
	}
	if o.NetworkOut != nil {
		toSerialize["network_out"] = o.NetworkOut
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Successor.IsSet() {
		toSerialize["successor"] = o.Successor.Get()
	}
	if o.Transfer != nil {
		toSerialize["transfer"] = o.Transfer
	}
	if o.Vcpus != nil {
		toSerialize["vcpus"] = o.Vcpus
	}
	if o.Gpus != nil {
		toSerialize["gpus"] = o.Gpus
	}
	return json.Marshal(toSerialize)
}

type NullableLinodeType struct {
	value *LinodeType
	isSet bool
}

func (v NullableLinodeType) Get() *LinodeType {
	return v.value
}

func (v *NullableLinodeType) Set(val *LinodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinodeType(val *LinodeType) *NullableLinodeType {
	return &NullableLinodeType{value: val, isSet: true}
}

func (v NullableLinodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


