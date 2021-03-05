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

// DiskRequest Disk object request.
type DiskRequest struct {
	Size int32 `json:"size"`
	Label Label `json:"label"`
	Filesystem *Filesystem `json:"filesystem,omitempty"`
	// If true, this Disk is read-only. 
	ReadOnly *bool `json:"read_only,omitempty"`
	// An Image ID to deploy the Disk from. Official Linode Images start with `linode/ `, while your Images start with `private/`. See [/images](/docs/api/images/#images-list) for more information on the Images available for you to use. 
	Image *string `json:"image,omitempty"`
	// A list of public SSH keys that will be automatically appended to the root user's `~/.ssh/authorized_keys` file. 
	AuthorizedKeys *[]string `json:"authorized_keys,omitempty"`
	// A list of usernames that will have their SSH keys, if any, automatically appended to the root user's `~/.ssh/authorized_keys` file. 
	AuthorizedUsers *[]string `json:"authorized_users,omitempty"`
	// This will set the root user's password on the newly-created Linode. Linode passwords have the following constraints:    * Must meet a password strength score requirement that is calculated internally by the API.     If the strength requirement is not met, you will receive a `Password does not meet strength requirement` error. 
	RootPass *string `json:"root_pass,omitempty"`
	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible `image` is required to use a StackScript. To get a list of available StackScript and their permitted Images see [/stackscripts](/docs/api/stackscripts/#stackscripts-list). This field cannot be used when deploying from a Backup or a private Image. 
	StackscriptId *int32 `json:"stackscript_id,omitempty"`
	// This field is required only if the StackScript being deployed requires input data from the User for successful completion. See <a target=\"_top\" href=\"/docs/platform/stackscripts/#variables-and-udfs\">Variables and UDFs</a> for more details. This field is required to be valid JSON. 
	StackscriptData *map[string]interface{} `json:"stackscript_data,omitempty"`
}

// NewDiskRequest instantiates a new DiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskRequest(size int32, label Label) *DiskRequest {
	this := DiskRequest{}
	this.Size = size
	this.Label = label
	return &this
}

// NewDiskRequestWithDefaults instantiates a new DiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskRequestWithDefaults() *DiskRequest {
	this := DiskRequest{}
	return &this
}

// GetSize returns the Size field value
func (o *DiskRequest) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *DiskRequest) SetSize(v int32) {
	o.Size = v
}

// GetLabel returns the Label field value
func (o *DiskRequest) GetLabel() Label {
	if o == nil {
		var ret Label
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetLabelOk() (*Label, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DiskRequest) SetLabel(v Label) {
	o.Label = v
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *DiskRequest) GetFilesystem() Filesystem {
	if o == nil || o.Filesystem == nil {
		var ret Filesystem
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetFilesystemOk() (*Filesystem, bool) {
	if o == nil || o.Filesystem == nil {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *DiskRequest) HasFilesystem() bool {
	if o != nil && o.Filesystem != nil {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given Filesystem and assigns it to the Filesystem field.
func (o *DiskRequest) SetFilesystem(v Filesystem) {
	o.Filesystem = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *DiskRequest) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *DiskRequest) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *DiskRequest) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DiskRequest) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DiskRequest) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DiskRequest) SetImage(v string) {
	o.Image = &v
}

// GetAuthorizedKeys returns the AuthorizedKeys field value if set, zero value otherwise.
func (o *DiskRequest) GetAuthorizedKeys() []string {
	if o == nil || o.AuthorizedKeys == nil {
		var ret []string
		return ret
	}
	return *o.AuthorizedKeys
}

// GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetAuthorizedKeysOk() (*[]string, bool) {
	if o == nil || o.AuthorizedKeys == nil {
		return nil, false
	}
	return o.AuthorizedKeys, true
}

// HasAuthorizedKeys returns a boolean if a field has been set.
func (o *DiskRequest) HasAuthorizedKeys() bool {
	if o != nil && o.AuthorizedKeys != nil {
		return true
	}

	return false
}

// SetAuthorizedKeys gets a reference to the given []string and assigns it to the AuthorizedKeys field.
func (o *DiskRequest) SetAuthorizedKeys(v []string) {
	o.AuthorizedKeys = &v
}

// GetAuthorizedUsers returns the AuthorizedUsers field value if set, zero value otherwise.
func (o *DiskRequest) GetAuthorizedUsers() []string {
	if o == nil || o.AuthorizedUsers == nil {
		var ret []string
		return ret
	}
	return *o.AuthorizedUsers
}

// GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetAuthorizedUsersOk() (*[]string, bool) {
	if o == nil || o.AuthorizedUsers == nil {
		return nil, false
	}
	return o.AuthorizedUsers, true
}

// HasAuthorizedUsers returns a boolean if a field has been set.
func (o *DiskRequest) HasAuthorizedUsers() bool {
	if o != nil && o.AuthorizedUsers != nil {
		return true
	}

	return false
}

// SetAuthorizedUsers gets a reference to the given []string and assigns it to the AuthorizedUsers field.
func (o *DiskRequest) SetAuthorizedUsers(v []string) {
	o.AuthorizedUsers = &v
}

// GetRootPass returns the RootPass field value if set, zero value otherwise.
func (o *DiskRequest) GetRootPass() string {
	if o == nil || o.RootPass == nil {
		var ret string
		return ret
	}
	return *o.RootPass
}

// GetRootPassOk returns a tuple with the RootPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetRootPassOk() (*string, bool) {
	if o == nil || o.RootPass == nil {
		return nil, false
	}
	return o.RootPass, true
}

// HasRootPass returns a boolean if a field has been set.
func (o *DiskRequest) HasRootPass() bool {
	if o != nil && o.RootPass != nil {
		return true
	}

	return false
}

// SetRootPass gets a reference to the given string and assigns it to the RootPass field.
func (o *DiskRequest) SetRootPass(v string) {
	o.RootPass = &v
}

// GetStackscriptId returns the StackscriptId field value if set, zero value otherwise.
func (o *DiskRequest) GetStackscriptId() int32 {
	if o == nil || o.StackscriptId == nil {
		var ret int32
		return ret
	}
	return *o.StackscriptId
}

// GetStackscriptIdOk returns a tuple with the StackscriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetStackscriptIdOk() (*int32, bool) {
	if o == nil || o.StackscriptId == nil {
		return nil, false
	}
	return o.StackscriptId, true
}

// HasStackscriptId returns a boolean if a field has been set.
func (o *DiskRequest) HasStackscriptId() bool {
	if o != nil && o.StackscriptId != nil {
		return true
	}

	return false
}

// SetStackscriptId gets a reference to the given int32 and assigns it to the StackscriptId field.
func (o *DiskRequest) SetStackscriptId(v int32) {
	o.StackscriptId = &v
}

// GetStackscriptData returns the StackscriptData field value if set, zero value otherwise.
func (o *DiskRequest) GetStackscriptData() map[string]interface{} {
	if o == nil || o.StackscriptData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StackscriptData
}

// GetStackscriptDataOk returns a tuple with the StackscriptData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskRequest) GetStackscriptDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.StackscriptData == nil {
		return nil, false
	}
	return o.StackscriptData, true
}

// HasStackscriptData returns a boolean if a field has been set.
func (o *DiskRequest) HasStackscriptData() bool {
	if o != nil && o.StackscriptData != nil {
		return true
	}

	return false
}

// SetStackscriptData gets a reference to the given map[string]interface{} and assigns it to the StackscriptData field.
func (o *DiskRequest) SetStackscriptData(v map[string]interface{}) {
	o.StackscriptData = &v
}

func (o DiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if o.Filesystem != nil {
		toSerialize["filesystem"] = o.Filesystem
	}
	if o.ReadOnly != nil {
		toSerialize["read_only"] = o.ReadOnly
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.AuthorizedKeys != nil {
		toSerialize["authorized_keys"] = o.AuthorizedKeys
	}
	if o.AuthorizedUsers != nil {
		toSerialize["authorized_users"] = o.AuthorizedUsers
	}
	if o.RootPass != nil {
		toSerialize["root_pass"] = o.RootPass
	}
	if o.StackscriptId != nil {
		toSerialize["stackscript_id"] = o.StackscriptId
	}
	if o.StackscriptData != nil {
		toSerialize["stackscript_data"] = o.StackscriptData
	}
	return json.Marshal(toSerialize)
}

type NullableDiskRequest struct {
	value *DiskRequest
	isSet bool
}

func (v NullableDiskRequest) Get() *DiskRequest {
	return v.value
}

func (v *NullableDiskRequest) Set(val *DiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskRequest(val *DiskRequest) *NullableDiskRequest {
	return &NullableDiskRequest{value: val, isSet: true}
}

func (v NullableDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


