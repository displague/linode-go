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

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// A Backup ID from another Linode's available backups. Your User must have `read_write` access to that Linode, the Backup must have a `status` of `successful`, and the Linode must be deployed to the same `region` as the Backup. See [/linode/instances/{linodeId}/backups](/docs/api/linode-instances/#backups-list) for a Linode's available backups.  This field and the `image` field are mutually exclusive. 
	BackupId *int32 `json:"backup_id,omitempty"`
	// If this field is set to `true`, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.  This option is always treated as `true` if the account-wide `backups_enabled` setting is `true`.  See [account settings](/docs/api/account/#account-settings-view) for more information.  Backup pricing is included in the response from [/linodes/types](/docs/api/linode-types/#types-list) 
	BackupsEnabled *bool `json:"backups_enabled,omitempty"`
	// When deploying from an Image, this field is optional, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode. 
	SwapSize *int32 `json:"swap_size,omitempty"`
	// The [Linode Type](/docs/api/linode-types/#types-list) of the Linode you are creating. 
	Type string `json:"type"`
	// The [Region](/docs/api/regions/#regions-list) where the Linode will be located. 
	Region string `json:"region"`
	// An Image ID to deploy the Disk from. Official Linode Images start with `linode/ `, while your Images start with `private/`. See [/images](/docs/api/images/) for more information on the Images available for you to use. 
	Image *string `json:"image,omitempty"`
	RootPass *RootPass `json:"root_pass,omitempty"`
	// A list of SSH public keys to deploy for the root user on the newly-created Linode.  Only accepted if `image` is provided. 
	AuthorizedKeys *[]string `json:"authorized_keys,omitempty"`
	// The StackScript to deploy to the newly-created Linode.  If provided, \"image\" must also be provided, and must be an Image that is compatible with this StackScript. 
	StackscriptId *int32 `json:"stackscript_id,omitempty"`
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if `stackscript_id` is given.  The required values depend on the StackScript being deployed. 
	StackscriptData *map[string]interface{} `json:"stackscript_data,omitempty"`
	// Whether to boot this Linode after the deploy is complete. Defaults to true if `image` is provided. Not accepted if not deploying from an Image. 
	Booted *bool `json:"booted,omitempty"`
	Label *Label `json:"label,omitempty"`
	Tags *Tags `json:"tags,omitempty"`
	Group *Group `json:"group,omitempty"`
	// If true, the created Linode will have private networking enabled. 
	PrivateIp *bool `json:"private_ip,omitempty"`
	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users `~/.ssh/authorized_keys` file automatically. 
	AuthorizedUsers *[]string `json:"authorized_users,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2(type_ string, region string) *InlineObject2 {
	this := InlineObject2{}
	var swapSize int32 = 512
	this.SwapSize = &swapSize
	this.Type = type_
	this.Region = region
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	var swapSize int32 = 512
	this.SwapSize = &swapSize
	return &this
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *InlineObject2) GetBackupId() int32 {
	if o == nil || o.BackupId == nil {
		var ret int32
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBackupIdOk() (*int32, bool) {
	if o == nil || o.BackupId == nil {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *InlineObject2) HasBackupId() bool {
	if o != nil && o.BackupId != nil {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given int32 and assigns it to the BackupId field.
func (o *InlineObject2) SetBackupId(v int32) {
	o.BackupId = &v
}

// GetBackupsEnabled returns the BackupsEnabled field value if set, zero value otherwise.
func (o *InlineObject2) GetBackupsEnabled() bool {
	if o == nil || o.BackupsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BackupsEnabled
}

// GetBackupsEnabledOk returns a tuple with the BackupsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBackupsEnabledOk() (*bool, bool) {
	if o == nil || o.BackupsEnabled == nil {
		return nil, false
	}
	return o.BackupsEnabled, true
}

// HasBackupsEnabled returns a boolean if a field has been set.
func (o *InlineObject2) HasBackupsEnabled() bool {
	if o != nil && o.BackupsEnabled != nil {
		return true
	}

	return false
}

// SetBackupsEnabled gets a reference to the given bool and assigns it to the BackupsEnabled field.
func (o *InlineObject2) SetBackupsEnabled(v bool) {
	o.BackupsEnabled = &v
}

// GetSwapSize returns the SwapSize field value if set, zero value otherwise.
func (o *InlineObject2) GetSwapSize() int32 {
	if o == nil || o.SwapSize == nil {
		var ret int32
		return ret
	}
	return *o.SwapSize
}

// GetSwapSizeOk returns a tuple with the SwapSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetSwapSizeOk() (*int32, bool) {
	if o == nil || o.SwapSize == nil {
		return nil, false
	}
	return o.SwapSize, true
}

// HasSwapSize returns a boolean if a field has been set.
func (o *InlineObject2) HasSwapSize() bool {
	if o != nil && o.SwapSize != nil {
		return true
	}

	return false
}

// SetSwapSize gets a reference to the given int32 and assigns it to the SwapSize field.
func (o *InlineObject2) SetSwapSize(v int32) {
	o.SwapSize = &v
}

// GetType returns the Type field value
func (o *InlineObject2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject2) SetType(v string) {
	o.Type = v
}

// GetRegion returns the Region field value
func (o *InlineObject2) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *InlineObject2) SetRegion(v string) {
	o.Region = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *InlineObject2) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *InlineObject2) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *InlineObject2) SetImage(v string) {
	o.Image = &v
}

// GetRootPass returns the RootPass field value if set, zero value otherwise.
func (o *InlineObject2) GetRootPass() RootPass {
	if o == nil || o.RootPass == nil {
		var ret RootPass
		return ret
	}
	return *o.RootPass
}

// GetRootPassOk returns a tuple with the RootPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetRootPassOk() (*RootPass, bool) {
	if o == nil || o.RootPass == nil {
		return nil, false
	}
	return o.RootPass, true
}

// HasRootPass returns a boolean if a field has been set.
func (o *InlineObject2) HasRootPass() bool {
	if o != nil && o.RootPass != nil {
		return true
	}

	return false
}

// SetRootPass gets a reference to the given RootPass and assigns it to the RootPass field.
func (o *InlineObject2) SetRootPass(v RootPass) {
	o.RootPass = &v
}

// GetAuthorizedKeys returns the AuthorizedKeys field value if set, zero value otherwise.
func (o *InlineObject2) GetAuthorizedKeys() []string {
	if o == nil || o.AuthorizedKeys == nil {
		var ret []string
		return ret
	}
	return *o.AuthorizedKeys
}

// GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetAuthorizedKeysOk() (*[]string, bool) {
	if o == nil || o.AuthorizedKeys == nil {
		return nil, false
	}
	return o.AuthorizedKeys, true
}

// HasAuthorizedKeys returns a boolean if a field has been set.
func (o *InlineObject2) HasAuthorizedKeys() bool {
	if o != nil && o.AuthorizedKeys != nil {
		return true
	}

	return false
}

// SetAuthorizedKeys gets a reference to the given []string and assigns it to the AuthorizedKeys field.
func (o *InlineObject2) SetAuthorizedKeys(v []string) {
	o.AuthorizedKeys = &v
}

// GetStackscriptId returns the StackscriptId field value if set, zero value otherwise.
func (o *InlineObject2) GetStackscriptId() int32 {
	if o == nil || o.StackscriptId == nil {
		var ret int32
		return ret
	}
	return *o.StackscriptId
}

// GetStackscriptIdOk returns a tuple with the StackscriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetStackscriptIdOk() (*int32, bool) {
	if o == nil || o.StackscriptId == nil {
		return nil, false
	}
	return o.StackscriptId, true
}

// HasStackscriptId returns a boolean if a field has been set.
func (o *InlineObject2) HasStackscriptId() bool {
	if o != nil && o.StackscriptId != nil {
		return true
	}

	return false
}

// SetStackscriptId gets a reference to the given int32 and assigns it to the StackscriptId field.
func (o *InlineObject2) SetStackscriptId(v int32) {
	o.StackscriptId = &v
}

// GetStackscriptData returns the StackscriptData field value if set, zero value otherwise.
func (o *InlineObject2) GetStackscriptData() map[string]interface{} {
	if o == nil || o.StackscriptData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StackscriptData
}

// GetStackscriptDataOk returns a tuple with the StackscriptData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetStackscriptDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.StackscriptData == nil {
		return nil, false
	}
	return o.StackscriptData, true
}

// HasStackscriptData returns a boolean if a field has been set.
func (o *InlineObject2) HasStackscriptData() bool {
	if o != nil && o.StackscriptData != nil {
		return true
	}

	return false
}

// SetStackscriptData gets a reference to the given map[string]interface{} and assigns it to the StackscriptData field.
func (o *InlineObject2) SetStackscriptData(v map[string]interface{}) {
	o.StackscriptData = &v
}

// GetBooted returns the Booted field value if set, zero value otherwise.
func (o *InlineObject2) GetBooted() bool {
	if o == nil || o.Booted == nil {
		var ret bool
		return ret
	}
	return *o.Booted
}

// GetBootedOk returns a tuple with the Booted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBootedOk() (*bool, bool) {
	if o == nil || o.Booted == nil {
		return nil, false
	}
	return o.Booted, true
}

// HasBooted returns a boolean if a field has been set.
func (o *InlineObject2) HasBooted() bool {
	if o != nil && o.Booted != nil {
		return true
	}

	return false
}

// SetBooted gets a reference to the given bool and assigns it to the Booted field.
func (o *InlineObject2) SetBooted(v bool) {
	o.Booted = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InlineObject2) GetLabel() Label {
	if o == nil || o.Label == nil {
		var ret Label
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetLabelOk() (*Label, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InlineObject2) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given Label and assigns it to the Label field.
func (o *InlineObject2) SetLabel(v Label) {
	o.Label = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject2) GetTags() Tags {
	if o == nil || o.Tags == nil {
		var ret Tags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetTagsOk() (*Tags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject2) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given Tags and assigns it to the Tags field.
func (o *InlineObject2) SetTags(v Tags) {
	o.Tags = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineObject2) GetGroup() Group {
	if o == nil || o.Group == nil {
		var ret Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetGroupOk() (*Group, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineObject2) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Group and assigns it to the Group field.
func (o *InlineObject2) SetGroup(v Group) {
	o.Group = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *InlineObject2) GetPrivateIp() bool {
	if o == nil || o.PrivateIp == nil {
		var ret bool
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetPrivateIpOk() (*bool, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *InlineObject2) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given bool and assigns it to the PrivateIp field.
func (o *InlineObject2) SetPrivateIp(v bool) {
	o.PrivateIp = &v
}

// GetAuthorizedUsers returns the AuthorizedUsers field value if set, zero value otherwise.
func (o *InlineObject2) GetAuthorizedUsers() []string {
	if o == nil || o.AuthorizedUsers == nil {
		var ret []string
		return ret
	}
	return *o.AuthorizedUsers
}

// GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetAuthorizedUsersOk() (*[]string, bool) {
	if o == nil || o.AuthorizedUsers == nil {
		return nil, false
	}
	return o.AuthorizedUsers, true
}

// HasAuthorizedUsers returns a boolean if a field has been set.
func (o *InlineObject2) HasAuthorizedUsers() bool {
	if o != nil && o.AuthorizedUsers != nil {
		return true
	}

	return false
}

// SetAuthorizedUsers gets a reference to the given []string and assigns it to the AuthorizedUsers field.
func (o *InlineObject2) SetAuthorizedUsers(v []string) {
	o.AuthorizedUsers = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupId != nil {
		toSerialize["backup_id"] = o.BackupId
	}
	if o.BackupsEnabled != nil {
		toSerialize["backups_enabled"] = o.BackupsEnabled
	}
	if o.SwapSize != nil {
		toSerialize["swap_size"] = o.SwapSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.RootPass != nil {
		toSerialize["root_pass"] = o.RootPass
	}
	if o.AuthorizedKeys != nil {
		toSerialize["authorized_keys"] = o.AuthorizedKeys
	}
	if o.StackscriptId != nil {
		toSerialize["stackscript_id"] = o.StackscriptId
	}
	if o.StackscriptData != nil {
		toSerialize["stackscript_data"] = o.StackscriptData
	}
	if o.Booted != nil {
		toSerialize["booted"] = o.Booted
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.PrivateIp != nil {
		toSerialize["private_ip"] = o.PrivateIp
	}
	if o.AuthorizedUsers != nil {
		toSerialize["authorized_users"] = o.AuthorizedUsers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


