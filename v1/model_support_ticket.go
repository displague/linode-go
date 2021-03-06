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
	"time"
)

// SupportTicket A Support Ticket opened on your Account. 
type SupportTicket struct {
	// The ID of the Support Ticket. 
	Id *int32 `json:"id,omitempty"`
	// A list of filenames representing attached files associated with this Ticket. 
	Attachments *[]string `json:"attachments,omitempty"`
	// The date and time this Ticket was closed. 
	Closed NullableTime `json:"closed,omitempty"`
	// Whether the Support Ticket may be closed. 
	Closable *bool `json:"closable,omitempty"`
	// The full details of the issue or question. 
	Description *string `json:"description,omitempty"`
	Entity NullableSupportTicketEntity `json:"entity,omitempty"`
	// The Gravatar ID of the User who opened this Ticket. 
	GravatarId *string `json:"gravatar_id,omitempty"`
	// The date and time this Ticket was created. 
	Opened *time.Time `json:"opened,omitempty"`
	// The User who opened this Ticket. 
	OpenedBy *string `json:"opened_by,omitempty"`
	// The current status of this Ticket.
	Status *string `json:"status,omitempty"`
	// The summary or title for this Ticket. 
	Summary *string `json:"summary,omitempty"`
	// The date and time this Ticket was last updated. 
	Updated *time.Time `json:"updated,omitempty"`
	// The User who last updated this Ticket. 
	UpdatedBy NullableString `json:"updated_by,omitempty"`
}

// NewSupportTicket instantiates a new SupportTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicket() *SupportTicket {
	this := SupportTicket{}
	return &this
}

// NewSupportTicketWithDefaults instantiates a new SupportTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketWithDefaults() *SupportTicket {
	this := SupportTicket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportTicket) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportTicket) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SupportTicket) SetId(v int32) {
	o.Id = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SupportTicket) GetAttachments() []string {
	if o == nil || o.Attachments == nil {
		var ret []string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetAttachmentsOk() (*[]string, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SupportTicket) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *SupportTicket) SetAttachments(v []string) {
	o.Attachments = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicket) GetClosed() time.Time {
	if o == nil || o.Closed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Closed.Get()
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicket) GetClosedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Closed.Get(), o.Closed.IsSet()
}

// HasClosed returns a boolean if a field has been set.
func (o *SupportTicket) HasClosed() bool {
	if o != nil && o.Closed.IsSet() {
		return true
	}

	return false
}

// SetClosed gets a reference to the given NullableTime and assigns it to the Closed field.
func (o *SupportTicket) SetClosed(v time.Time) {
	o.Closed.Set(&v)
}
// SetClosedNil sets the value for Closed to be an explicit nil
func (o *SupportTicket) SetClosedNil() {
	o.Closed.Set(nil)
}

// UnsetClosed ensures that no value is present for Closed, not even an explicit nil
func (o *SupportTicket) UnsetClosed() {
	o.Closed.Unset()
}

// GetClosable returns the Closable field value if set, zero value otherwise.
func (o *SupportTicket) GetClosable() bool {
	if o == nil || o.Closable == nil {
		var ret bool
		return ret
	}
	return *o.Closable
}

// GetClosableOk returns a tuple with the Closable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetClosableOk() (*bool, bool) {
	if o == nil || o.Closable == nil {
		return nil, false
	}
	return o.Closable, true
}

// HasClosable returns a boolean if a field has been set.
func (o *SupportTicket) HasClosable() bool {
	if o != nil && o.Closable != nil {
		return true
	}

	return false
}

// SetClosable gets a reference to the given bool and assigns it to the Closable field.
func (o *SupportTicket) SetClosable(v bool) {
	o.Closable = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SupportTicket) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SupportTicket) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SupportTicket) SetDescription(v string) {
	o.Description = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicket) GetEntity() SupportTicketEntity {
	if o == nil || o.Entity.Get() == nil {
		var ret SupportTicketEntity
		return ret
	}
	return *o.Entity.Get()
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicket) GetEntityOk() (*SupportTicketEntity, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Entity.Get(), o.Entity.IsSet()
}

// HasEntity returns a boolean if a field has been set.
func (o *SupportTicket) HasEntity() bool {
	if o != nil && o.Entity.IsSet() {
		return true
	}

	return false
}

// SetEntity gets a reference to the given NullableSupportTicketEntity and assigns it to the Entity field.
func (o *SupportTicket) SetEntity(v SupportTicketEntity) {
	o.Entity.Set(&v)
}
// SetEntityNil sets the value for Entity to be an explicit nil
func (o *SupportTicket) SetEntityNil() {
	o.Entity.Set(nil)
}

// UnsetEntity ensures that no value is present for Entity, not even an explicit nil
func (o *SupportTicket) UnsetEntity() {
	o.Entity.Unset()
}

// GetGravatarId returns the GravatarId field value if set, zero value otherwise.
func (o *SupportTicket) GetGravatarId() string {
	if o == nil || o.GravatarId == nil {
		var ret string
		return ret
	}
	return *o.GravatarId
}

// GetGravatarIdOk returns a tuple with the GravatarId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetGravatarIdOk() (*string, bool) {
	if o == nil || o.GravatarId == nil {
		return nil, false
	}
	return o.GravatarId, true
}

// HasGravatarId returns a boolean if a field has been set.
func (o *SupportTicket) HasGravatarId() bool {
	if o != nil && o.GravatarId != nil {
		return true
	}

	return false
}

// SetGravatarId gets a reference to the given string and assigns it to the GravatarId field.
func (o *SupportTicket) SetGravatarId(v string) {
	o.GravatarId = &v
}

// GetOpened returns the Opened field value if set, zero value otherwise.
func (o *SupportTicket) GetOpened() time.Time {
	if o == nil || o.Opened == nil {
		var ret time.Time
		return ret
	}
	return *o.Opened
}

// GetOpenedOk returns a tuple with the Opened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetOpenedOk() (*time.Time, bool) {
	if o == nil || o.Opened == nil {
		return nil, false
	}
	return o.Opened, true
}

// HasOpened returns a boolean if a field has been set.
func (o *SupportTicket) HasOpened() bool {
	if o != nil && o.Opened != nil {
		return true
	}

	return false
}

// SetOpened gets a reference to the given time.Time and assigns it to the Opened field.
func (o *SupportTicket) SetOpened(v time.Time) {
	o.Opened = &v
}

// GetOpenedBy returns the OpenedBy field value if set, zero value otherwise.
func (o *SupportTicket) GetOpenedBy() string {
	if o == nil || o.OpenedBy == nil {
		var ret string
		return ret
	}
	return *o.OpenedBy
}

// GetOpenedByOk returns a tuple with the OpenedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetOpenedByOk() (*string, bool) {
	if o == nil || o.OpenedBy == nil {
		return nil, false
	}
	return o.OpenedBy, true
}

// HasOpenedBy returns a boolean if a field has been set.
func (o *SupportTicket) HasOpenedBy() bool {
	if o != nil && o.OpenedBy != nil {
		return true
	}

	return false
}

// SetOpenedBy gets a reference to the given string and assigns it to the OpenedBy field.
func (o *SupportTicket) SetOpenedBy(v string) {
	o.OpenedBy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SupportTicket) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SupportTicket) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SupportTicket) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *SupportTicket) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *SupportTicket) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *SupportTicket) SetSummary(v string) {
	o.Summary = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *SupportTicket) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *SupportTicket) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *SupportTicket) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportTicket) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportTicket) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SupportTicket) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *SupportTicket) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *SupportTicket) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *SupportTicket) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

func (o SupportTicket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Closed.IsSet() {
		toSerialize["closed"] = o.Closed.Get()
	}
	if o.Closable != nil {
		toSerialize["closable"] = o.Closable
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entity.IsSet() {
		toSerialize["entity"] = o.Entity.Get()
	}
	if o.GravatarId != nil {
		toSerialize["gravatar_id"] = o.GravatarId
	}
	if o.Opened != nil {
		toSerialize["opened"] = o.Opened
	}
	if o.OpenedBy != nil {
		toSerialize["opened_by"] = o.OpenedBy
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSupportTicket struct {
	value *SupportTicket
	isSet bool
}

func (v NullableSupportTicket) Get() *SupportTicket {
	return v.value
}

func (v *NullableSupportTicket) Set(val *SupportTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicket(val *SupportTicket) *NullableSupportTicket {
	return &NullableSupportTicket{value: val, isSet: true}
}

func (v NullableSupportTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


