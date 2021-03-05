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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// NetworkingApiService NetworkingApi service
type NetworkingApiService service

type ApiAllocateIPRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiAllocateIPRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiAllocateIPRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiAllocateIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.AllocateIPExecute(r)
}

/*
 * AllocateIP IP Address Allocate
 * Allocates a new IPv4 Address on your Account. The Linode must be configured to support additional addresses - please [open a support ticket](/docs/api/support/#support-ticket-open) requesting additional addresses before attempting allocation.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAllocateIPRequest
 */
func (a *NetworkingApiService) AllocateIP(ctx _context.Context) ApiAllocateIPRequest {
	return ApiAllocateIPRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *NetworkingApiService) AllocateIPExecute(r ApiAllocateIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.AllocateIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ips"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.uNKNOWNBASETYPE == nil {
		return localVarReturnValue, nil, reportError("uNKNOWNBASETYPE is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAssignIPsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiAssignIPsRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiAssignIPsRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiAssignIPsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.AssignIPsExecute(r)
}

/*
 * AssignIPs Linodes Assign IPs
 * Assign multiple IPs to multiple Linodes in one Region. This allows swapping, shuffling, or otherwise reorganizing IPv4 Addresses to your Linodes.  When the assignment is finished, all Linodes must end up with at least one public IPv4 and no more than one private IPv4.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAssignIPsRequest
 */
func (a *NetworkingApiService) AssignIPs(ctx _context.Context) ApiAssignIPsRequest {
	return ApiAssignIPsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *NetworkingApiService) AssignIPsExecute(r ApiAssignIPsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.AssignIPs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ipv4/assign"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.uNKNOWNBASETYPE == nil {
		return localVarReturnValue, nil, reportError("uNKNOWNBASETYPE is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateFirewallDeviceRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiCreateFirewallDeviceRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiCreateFirewallDeviceRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiCreateFirewallDeviceRequest) Execute() (FirewallDevices, *_nethttp.Response, error) {
	return r.ApiService.CreateFirewallDeviceExecute(r)
}

/*
 * CreateFirewallDevice Firewall Device Create
 * Creates a Firewall Device, which assigns a Firewall to a Linode service (referred to
as the Device's `entity`). Currently, only Devices with an entity of type `linode` are accepted.
A Firewall can be assigned to multiple Linode services, and up to five active Firewalls can
be assigned to a single Linode service. Additional disabled Firewalls can be
assigned to a service, but they cannot be enabled if five other active Firewalls
are already assigned to the same service.

Creating a Firewall Device will apply the Rules from a Firewall to a Linode service.
A `firewall_device_add` Event is generated when the Firewall Device is added successfully.

**Note:** When a Firewall is assigned to a Linode and you attempt
to [migrate the Linode to a data center](/docs/api/linode-instances/#dc-migrationpending-host-migration-initiate) that does not support Cloud Firewalls, the migration will fail.
Use the [List Regions](/docs/api/regions/#regions-list) endpoint to view a list of a data center's capabilities.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiCreateFirewallDeviceRequest
 */
func (a *NetworkingApiService) CreateFirewallDevice(ctx _context.Context, firewallId int32) ApiCreateFirewallDeviceRequest {
	return ApiCreateFirewallDeviceRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return FirewallDevices
 */
func (a *NetworkingApiService) CreateFirewallDeviceExecute(r ApiCreateFirewallDeviceRequest) (FirewallDevices, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FirewallDevices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.CreateFirewallDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateFirewallsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	inlineObject11 *InlineObject11
}

func (r ApiCreateFirewallsRequest) InlineObject11(inlineObject11 InlineObject11) ApiCreateFirewallsRequest {
	r.inlineObject11 = &inlineObject11
	return r
}

func (r ApiCreateFirewallsRequest) Execute() (Firewall, *_nethttp.Response, error) {
	return r.ApiService.CreateFirewallsExecute(r)
}

/*
 * CreateFirewalls Firewall Create
 * Creates a Firewall to filter network traffic. Use the `rules` property to
create inbound and outbound access rules. Use the `devices` property to assign the
Firewall to a Linode service. Currently, Firewalls can only be assigned to Linode
instances.

A Firewall can be assigned to multiple Linode services, and up to five active Firewalls
can be assigned to a single Linode service.

A `firewall_create` Event is generated when this endpoint returns successfully.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateFirewallsRequest
 */
func (a *NetworkingApiService) CreateFirewalls(ctx _context.Context) ApiCreateFirewallsRequest {
	return ApiCreateFirewallsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Firewall
 */
func (a *NetworkingApiService) CreateFirewallsExecute(r ApiCreateFirewallsRequest) (Firewall, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Firewall
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.CreateFirewalls")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inlineObject11
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteFirewallRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
}


func (r ApiDeleteFirewallRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteFirewallExecute(r)
}

/*
 * DeleteFirewall Firewall Delete
 * Delete a Firewall resource by its ID. This will remove all of the Firewall's Rules
from any Linode services that the Firewall was assigned to.

A `firewall_delete` Event is generated when this endpoint returns successfully.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiDeleteFirewallRequest
 */
func (a *NetworkingApiService) DeleteFirewall(ctx _context.Context, firewallId int32) ApiDeleteFirewallRequest {
	return ApiDeleteFirewallRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *NetworkingApiService) DeleteFirewallExecute(r ApiDeleteFirewallRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.DeleteFirewall")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteFirewallDeviceRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	deviceId int32
}


func (r ApiDeleteFirewallDeviceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteFirewallDeviceExecute(r)
}

/*
 * DeleteFirewallDevice Firewall Device Delete
 * Removes a Firewall Device, which removes a Firewall from the Linode service it was
assigned to by the Device. This will remove all of the Firewall's Rules from the Linode
service. If any other Firewalls have been assigned to the Linode service, then those Rules
will remain in effect.

A `firewall_device_remove` Event is generated when the Firewall Device is removed successfully.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @param deviceId ID of the Firewall Device to access. 
 * @return ApiDeleteFirewallDeviceRequest
 */
func (a *NetworkingApiService) DeleteFirewallDevice(ctx _context.Context, firewallId int32, deviceId int32) ApiDeleteFirewallDeviceRequest {
	return ApiDeleteFirewallDeviceRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
		deviceId: deviceId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *NetworkingApiService) DeleteFirewallDeviceExecute(r ApiDeleteFirewallDeviceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.DeleteFirewallDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(r.deviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFirewallRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
}


func (r ApiGetFirewallRequest) Execute() (Firewall, *_nethttp.Response, error) {
	return r.ApiService.GetFirewallExecute(r)
}

/*
 * GetFirewall Firewall View
 * Get a specific Firewall resource by its ID. The Firewall's Devices will not be
returned in the response. Instead, use the
[List Firewall Devices](/docs/api/networking/#firewall-devices-list)
endpoint to review them.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiGetFirewallRequest
 */
func (a *NetworkingApiService) GetFirewall(ctx _context.Context, firewallId int32) ApiGetFirewallRequest {
	return ApiGetFirewallRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return Firewall
 */
func (a *NetworkingApiService) GetFirewallExecute(r ApiGetFirewallRequest) (Firewall, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Firewall
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetFirewall")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFirewallDeviceRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	deviceId int32
}


func (r ApiGetFirewallDeviceRequest) Execute() (FirewallDevices, *_nethttp.Response, error) {
	return r.ApiService.GetFirewallDeviceExecute(r)
}

/*
 * GetFirewallDevice Firewall Device View
 * Returns information for a Firewall Device, which assigns a Firewall
to a Linode service (referred to as the Device's `entity`). Currently,
only Devices with an entity of type `linode` are accepted.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @param deviceId ID of the Firewall Device to access. 
 * @return ApiGetFirewallDeviceRequest
 */
func (a *NetworkingApiService) GetFirewallDevice(ctx _context.Context, firewallId int32, deviceId int32) ApiGetFirewallDeviceRequest {
	return ApiGetFirewallDeviceRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
		deviceId: deviceId,
	}
}

/*
 * Execute executes the request
 * @return FirewallDevices
 */
func (a *NetworkingApiService) GetFirewallDeviceExecute(r ApiGetFirewallDeviceRequest) (FirewallDevices, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FirewallDevices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetFirewallDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(r.deviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFirewallDevicesRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	page *int32
	pageSize *int32
}

func (r ApiGetFirewallDevicesRequest) Page(page int32) ApiGetFirewallDevicesRequest {
	r.page = &page
	return r
}
func (r ApiGetFirewallDevicesRequest) PageSize(pageSize int32) ApiGetFirewallDevicesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetFirewallDevicesRequest) Execute() (InlineResponse20043, *_nethttp.Response, error) {
	return r.ApiService.GetFirewallDevicesExecute(r)
}

/*
 * GetFirewallDevices Firewall Devices List
 * Returns a paginated list of a Firewall's Devices. A Firewall Device assigns a
Firewall to a Linode service (referred to as the Device's `entity`). Currently,
only Devices with an entity of type `linode` are accepted.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiGetFirewallDevicesRequest
 */
func (a *NetworkingApiService) GetFirewallDevices(ctx _context.Context, firewallId int32) ApiGetFirewallDevicesRequest {
	return ApiGetFirewallDevicesRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20043
 */
func (a *NetworkingApiService) GetFirewallDevicesExecute(r ApiGetFirewallDevicesRequest) (InlineResponse20043, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20043
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetFirewallDevices")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFirewallRulesRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
}


func (r ApiGetFirewallRulesRequest) Execute() (Rules, *_nethttp.Response, error) {
	return r.ApiService.GetFirewallRulesExecute(r)
}

/*
 * GetFirewallRules Firewall Rules List
 * Returns the inbound and outbound Rules for a Firewall.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiGetFirewallRulesRequest
 */
func (a *NetworkingApiService) GetFirewallRules(ctx _context.Context, firewallId int32) ApiGetFirewallRulesRequest {
	return ApiGetFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return Rules
 */
func (a *NetworkingApiService) GetFirewallRulesExecute(r ApiGetFirewallRulesRequest) (Rules, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Rules
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFirewallsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	page *int32
	pageSize *int32
}

func (r ApiGetFirewallsRequest) Page(page int32) ApiGetFirewallsRequest {
	r.page = &page
	return r
}
func (r ApiGetFirewallsRequest) PageSize(pageSize int32) ApiGetFirewallsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetFirewallsRequest) Execute() (InlineResponse20019, *_nethttp.Response, error) {
	return r.ApiService.GetFirewallsExecute(r)
}

/*
 * GetFirewalls Firewalls List
 * Returns a paginated list of your Firewalls.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetFirewallsRequest
 */
func (a *NetworkingApiService) GetFirewalls(ctx _context.Context) ApiGetFirewallsRequest {
	return ApiGetFirewallsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20019
 */
func (a *NetworkingApiService) GetFirewallsExecute(r ApiGetFirewallsRequest) (InlineResponse20019, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20019
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetFirewalls")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIPRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	address string
}


func (r ApiGetIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.GetIPExecute(r)
}

/*
 * GetIP IP Address View
 * Returns information about a single IP Address on your Account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param address The address to operate on.
 * @return ApiGetIPRequest
 */
func (a *NetworkingApiService) GetIP(ctx _context.Context, address string) ApiGetIPRequest {
	return ApiGetIPRequest{
		ApiService: a,
		ctx: ctx,
		address: address,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *NetworkingApiService) GetIPExecute(r ApiGetIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ips/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIPsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
}


func (r ApiGetIPsRequest) Execute() (InlineResponse20040, *_nethttp.Response, error) {
	return r.ApiService.GetIPsExecute(r)
}

/*
 * GetIPs IP Addresses List
 * Returns a paginated list of IP Addresses on your Account, excluding private addresses.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetIPsRequest
 */
func (a *NetworkingApiService) GetIPs(ctx _context.Context) ApiGetIPsRequest {
	return ApiGetIPsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20040
 */
func (a *NetworkingApiService) GetIPsExecute(r ApiGetIPsRequest) (InlineResponse20040, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20040
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetIPs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ips"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIPv6PoolsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	page *int32
	pageSize *int32
}

func (r ApiGetIPv6PoolsRequest) Page(page int32) ApiGetIPv6PoolsRequest {
	r.page = &page
	return r
}
func (r ApiGetIPv6PoolsRequest) PageSize(pageSize int32) ApiGetIPv6PoolsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetIPv6PoolsRequest) Execute() (InlineResponse20041, *_nethttp.Response, error) {
	return r.ApiService.GetIPv6PoolsExecute(r)
}

/*
 * GetIPv6Pools IPv6 Pools List
 * Displays the IPv6 pools on your Account. A pool of IPv6 addresses are routed to all of your Linodes in a single [Region](/docs/api/regions/#regions-list). Any Linode on your Account may bring up any address in this pool at any time, with no external configuration required.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetIPv6PoolsRequest
 */
func (a *NetworkingApiService) GetIPv6Pools(ctx _context.Context) ApiGetIPv6PoolsRequest {
	return ApiGetIPv6PoolsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20041
 */
func (a *NetworkingApiService) GetIPv6PoolsExecute(r ApiGetIPv6PoolsRequest) (InlineResponse20041, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20041
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetIPv6Pools")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ipv6/pools"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIPv6RangesRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	page *int32
	pageSize *int32
}

func (r ApiGetIPv6RangesRequest) Page(page int32) ApiGetIPv6RangesRequest {
	r.page = &page
	return r
}
func (r ApiGetIPv6RangesRequest) PageSize(pageSize int32) ApiGetIPv6RangesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetIPv6RangesRequest) Execute() (InlineResponse20042, *_nethttp.Response, error) {
	return r.ApiService.GetIPv6RangesExecute(r)
}

/*
 * GetIPv6Ranges IPv6 Ranges List
 * Displays the IPv6 ranges on your Account.


  * An IPv6 range is a `/64` block of IPv6 addresses routed to a single Linode in a given [Region](/docs/api/regions/#regions-list).

  * Your Linode is responsible for routing individual addresses in the range, or handling traffic for all the addresses in the range.

  * You must [open a support ticket](/docs/api/support/#support-ticket-open) to request a `/64` block of IPv6 addresses to be added to your account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetIPv6RangesRequest
 */
func (a *NetworkingApiService) GetIPv6Ranges(ctx _context.Context) ApiGetIPv6RangesRequest {
	return ApiGetIPv6RangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20042
 */
func (a *NetworkingApiService) GetIPv6RangesExecute(r ApiGetIPv6RangesRequest) (InlineResponse20042, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20042
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.GetIPv6Ranges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ipv6/ranges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiShareIPsRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiShareIPsRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiShareIPsRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiShareIPsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ShareIPsExecute(r)
}

/*
 * ShareIPs IP Sharing Configure
 * Configure shared IPs.  A shared IP may be brought up on a Linode other than the one it lists in its response.  This can be used to allow one Linode to begin serving requests should another become unresponsive.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiShareIPsRequest
 */
func (a *NetworkingApiService) ShareIPs(ctx _context.Context) ApiShareIPsRequest {
	return ApiShareIPsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *NetworkingApiService) ShareIPsExecute(r ApiShareIPsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.ShareIPs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ipv4/share"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.uNKNOWNBASETYPE == nil {
		return localVarReturnValue, nil, reportError("uNKNOWNBASETYPE is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateFirewallRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	inlineObject12 *InlineObject12
}

func (r ApiUpdateFirewallRequest) InlineObject12(inlineObject12 InlineObject12) ApiUpdateFirewallRequest {
	r.inlineObject12 = &inlineObject12
	return r
}

func (r ApiUpdateFirewallRequest) Execute() (Firewall, *_nethttp.Response, error) {
	return r.ApiService.UpdateFirewallExecute(r)
}

/*
 * UpdateFirewall Firewall Update
 * Updates information for a Firewall. Some parts of a Firewall's configuration cannot
be manipulated by this endpoint:

- A Firewall's Devices cannot be set with this endpoint. Instead, use the
[Create Firewall Device](/docs/api/networking/#firewall-device-create)
and [Delete Firewall Device](/docs/api/networking/#firewall-device-delete)
endpoints to assign and remove this Firewall from Linode services.

- A Firewall's Rules cannot be changed with this endpoint. Instead, use the
[Update Firewall Rules](/docs/api/networking/#firewall-rules-update)
endpoint to update your Rules.

- A Firewall's status can be set to `enabled` or `disabled` by this endpoint, but it cannot be
set to `deleted`. Instead, use the
[Delete Firewall](/docs/api/networking/#firewall-delete)
endpoint to delete a Firewall.

If a Firewall's status is changed with this endpoint, a corresponding `firewall_enable` or
`firewall_disable` Event will be generated.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiUpdateFirewallRequest
 */
func (a *NetworkingApiService) UpdateFirewall(ctx _context.Context, firewallId int32) ApiUpdateFirewallRequest {
	return ApiUpdateFirewallRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return Firewall
 */
func (a *NetworkingApiService) UpdateFirewallExecute(r ApiUpdateFirewallRequest) (Firewall, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Firewall
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.UpdateFirewall")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inlineObject12
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateFirewallRulesRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	firewallId int32
	body *Rules
}

func (r ApiUpdateFirewallRulesRequest) Body(body Rules) ApiUpdateFirewallRulesRequest {
	r.body = &body
	return r
}

func (r ApiUpdateFirewallRulesRequest) Execute() (Rules, *_nethttp.Response, error) {
	return r.ApiService.UpdateFirewallRulesExecute(r)
}

/*
 * UpdateFirewallRules Firewall Rules Update
 * Updates the inbound and outbound Rules for a Firewall. Using this endpoint will
replace all of a Firewall's ruleset with the Rules specified in your request.

This endpoint is in **beta**.


* Gain access to [Linode Cloud Firewall](https://www.linode.com/products/firewall/) by signing up for our [Greenlight Beta program](https://www.linode.com/green-light/#sign-up-form).
* During the beta, Cloud Firewall is not available in every [data center region](/docs/api/regions). For the current list of availability, see the [Cloud Firewall Product Documentation](https://www.linode.com/docs/products/networking/cloud-firewall/).
* Please make sure to prepend all requests with
`/v4beta` instead of `/v4`, and be aware that this endpoint may receive breaking
updates in the future.  This notice will be removed when this endpoint is out of
beta.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firewallId ID of the Firewall to access. 
 * @return ApiUpdateFirewallRulesRequest
 */
func (a *NetworkingApiService) UpdateFirewallRules(ctx _context.Context, firewallId int32) ApiUpdateFirewallRulesRequest {
	return ApiUpdateFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		firewallId: firewallId,
	}
}

/*
 * Execute executes the request
 * @return Rules
 */
func (a *NetworkingApiService) UpdateFirewallRulesExecute(r ApiUpdateFirewallRulesRequest) (Rules, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Rules
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.UpdateFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/firewalls/{firewallId}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"firewallId"+"}", _neturl.PathEscape(parameterToString(r.firewallId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateIPRequest struct {
	ctx _context.Context
	ApiService *NetworkingApiService
	address string
	iPAddress *IPAddress
}

func (r ApiUpdateIPRequest) IPAddress(iPAddress IPAddress) ApiUpdateIPRequest {
	r.iPAddress = &iPAddress
	return r
}

func (r ApiUpdateIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.UpdateIPExecute(r)
}

/*
 * UpdateIP IP Address RDNS Update
 * Sets RDNS on an IP Address. Forward DNS must already be set up for reverse DNS to be applied. If you set the RDNS to `null` for public IPv4 addresses, it will be reset to the default _members.linode.com_ RDNS value.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param address The address to operate on.
 * @return ApiUpdateIPRequest
 */
func (a *NetworkingApiService) UpdateIP(ctx _context.Context, address string) ApiUpdateIPRequest {
	return ApiUpdateIPRequest{
		ApiService: a,
		ctx: ctx,
		address: address,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *NetworkingApiService) UpdateIPExecute(r ApiUpdateIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkingApiService.UpdateIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/ips/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.iPAddress == nil {
		return localVarReturnValue, nil, reportError("iPAddress is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.iPAddress
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v InlineResponseDefault
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
