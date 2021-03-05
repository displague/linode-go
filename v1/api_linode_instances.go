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

// LinodeInstancesApiService LinodeInstancesApi service
type LinodeInstancesApiService service

type ApiAddLinodeConfigRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	linodeConfig *LinodeConfig
}

func (r ApiAddLinodeConfigRequest) LinodeConfig(linodeConfig LinodeConfig) ApiAddLinodeConfigRequest {
	r.linodeConfig = &linodeConfig
	return r
}

func (r ApiAddLinodeConfigRequest) Execute() (LinodeConfig, *_nethttp.Response, error) {
	return r.ApiService.AddLinodeConfigExecute(r)
}

/*
 * AddLinodeConfig Configuration Profile Create
 * Adds a new Configuration profile to a Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up Configuration profiles for.
 * @return ApiAddLinodeConfigRequest
 */
func (a *LinodeInstancesApiService) AddLinodeConfig(ctx _context.Context, linodeId int32) ApiAddLinodeConfigRequest {
	return ApiAddLinodeConfigRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return LinodeConfig
 */
func (a *LinodeInstancesApiService) AddLinodeConfigExecute(r ApiAddLinodeConfigRequest) (LinodeConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LinodeConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.AddLinodeConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/configs"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.linodeConfig == nil {
		return localVarReturnValue, nil, reportError("linodeConfig is required and must be specified")
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
	localVarPostBody = r.linodeConfig
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

type ApiAddLinodeDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskRequest *DiskRequest
}

func (r ApiAddLinodeDiskRequest) DiskRequest(diskRequest DiskRequest) ApiAddLinodeDiskRequest {
	r.diskRequest = &diskRequest
	return r
}

func (r ApiAddLinodeDiskRequest) Execute() (Disk, *_nethttp.Response, error) {
	return r.ApiService.AddLinodeDiskExecute(r)
}

/*
 * AddLinodeDisk Disk Create
 * Adds a new Disk to a Linode. You can optionally create a Disk from an Image (see [/images](/docs/api/images/#images-list) for a list of available public images, or use one of your own), and optionally provide a StackScript to deploy with this Disk.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiAddLinodeDiskRequest
 */
func (a *LinodeInstancesApiService) AddLinodeDisk(ctx _context.Context, linodeId int32) ApiAddLinodeDiskRequest {
	return ApiAddLinodeDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Disk
 */
func (a *LinodeInstancesApiService) AddLinodeDiskExecute(r ApiAddLinodeDiskRequest) (Disk, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Disk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.AddLinodeDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.diskRequest == nil {
		return localVarReturnValue, nil, reportError("diskRequest is required and must be specified")
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
	localVarPostBody = r.diskRequest
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

type ApiAddLinodeIPRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiAddLinodeIPRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiAddLinodeIPRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiAddLinodeIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.AddLinodeIPExecute(r)
}

/*
 * AddLinodeIP IPv4 Address Allocate
 * Allocates a public or private IPv4 address to a Linode. Public IP Addresses, after the one included with each Linode, incur an additional monthly charge. If you need an additional public IP Address you must request one - please [open a support ticket](/docs/api/support/#support-ticket-open). You may not add more than one private IPv4 address to a single Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiAddLinodeIPRequest
 */
func (a *LinodeInstancesApiService) AddLinodeIP(ctx _context.Context, linodeId int32) ApiAddLinodeIPRequest {
	return ApiAddLinodeIPRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *LinodeInstancesApiService) AddLinodeIPExecute(r ApiAddLinodeIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.AddLinodeIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/ips"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiBootLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject5 *InlineObject5
}

func (r ApiBootLinodeInstanceRequest) InlineObject5(inlineObject5 InlineObject5) ApiBootLinodeInstanceRequest {
	r.inlineObject5 = &inlineObject5
	return r
}

func (r ApiBootLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.BootLinodeInstanceExecute(r)
}

/*
 * BootLinodeInstance Linode Boot
 * Boots a Linode you have permission to modify. If no parameters are given, a Config profile
will be chosen for this boot based on the following criteria:

* If there is only one Config profile for this Linode, it will be used.
* If there is more than one Config profile, the last booted config will be used.
* If there is more than one Config profile and none were the last to be booted (because the
  Linode was never booted or the last booted config was deleted) an error will be returned.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to boot.
 * @return ApiBootLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) BootLinodeInstance(ctx _context.Context, linodeId int32) ApiBootLinodeInstanceRequest {
	return ApiBootLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) BootLinodeInstanceExecute(r ApiBootLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.BootLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/boot"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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
	localVarPostBody = r.inlineObject5
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

type ApiCancelBackupsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiCancelBackupsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CancelBackupsExecute(r)
}

/*
 * CancelBackups Backups Cancel
 * Cancels the Backup service on the given Linode. Deletes all of this Linode's existing backups forever.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to cancel backup service for.
 * @return ApiCancelBackupsRequest
 */
func (a *LinodeInstancesApiService) CancelBackups(ctx _context.Context, linodeId int32) ApiCancelBackupsRequest {
	return ApiCancelBackupsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) CancelBackupsExecute(r ApiCancelBackupsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.CancelBackups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiCloneLinodeDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
}


func (r ApiCloneLinodeDiskRequest) Execute() (Disk, *_nethttp.Response, error) {
	return r.ApiService.CloneLinodeDiskExecute(r)
}

/*
 * CloneLinodeDisk Disk Clone
 * Copies a disk, byte-for-byte, into a new Disk belonging to the same Linode. The Linode must have enough storage space available to accept a new Disk of the same size as this one or this operation will fail.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to clone.
 * @return ApiCloneLinodeDiskRequest
 */
func (a *LinodeInstancesApiService) CloneLinodeDisk(ctx _context.Context, linodeId int32, diskId int32) ApiCloneLinodeDiskRequest {
	return ApiCloneLinodeDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return Disk
 */
func (a *LinodeInstancesApiService) CloneLinodeDiskExecute(r ApiCloneLinodeDiskRequest) (Disk, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Disk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.CloneLinodeDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

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

type ApiCloneLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject6 *InlineObject6
}

func (r ApiCloneLinodeInstanceRequest) InlineObject6(inlineObject6 InlineObject6) ApiCloneLinodeInstanceRequest {
	r.inlineObject6 = &inlineObject6
	return r
}

func (r ApiCloneLinodeInstanceRequest) Execute() (Linode, *_nethttp.Response, error) {
	return r.ApiService.CloneLinodeInstanceExecute(r)
}

/*
 * CloneLinodeInstance Linode Clone
 * You can clone your Linode's existing Disks or Configuration profiles to
another Linode on your Account. In order for this request to complete
successfully, your User must have the `add_linodes` grant. Cloning to a
new Linode will incur a charge on your Account.

If cloning to an existing Linode, any actions currently running or
queued must be completed first before you can clone to it.

Up to five clone operations from any given source Linode can be run concurrently.
If more concurrent clones are attempted, an HTTP 400 error will be
returned by this endpoint.

Any [tags](/docs/api/tags/#tags-list) existing on the source Linode will be cloned to the target Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to clone.
 * @return ApiCloneLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) CloneLinodeInstance(ctx _context.Context, linodeId int32) ApiCloneLinodeInstanceRequest {
	return ApiCloneLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Linode
 */
func (a *LinodeInstancesApiService) CloneLinodeInstanceExecute(r ApiCloneLinodeInstanceRequest) (Linode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Linode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.CloneLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject6 == nil {
		return localVarReturnValue, nil, reportError("inlineObject6 is required and must be specified")
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
	localVarPostBody = r.inlineObject6
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

type ApiCreateLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	inlineObject2 *InlineObject2
}

func (r ApiCreateLinodeInstanceRequest) InlineObject2(inlineObject2 InlineObject2) ApiCreateLinodeInstanceRequest {
	r.inlineObject2 = &inlineObject2
	return r
}

func (r ApiCreateLinodeInstanceRequest) Execute() (Linode, *_nethttp.Response, error) {
	return r.ApiService.CreateLinodeInstanceExecute(r)
}

/*
 * CreateLinodeInstance Linode Create
 * Creates a Linode Instance on your Account. In order for this
request to complete successfully, your User must have the `add_linodes` grant. Creating a
new Linode will incur a charge on your Account.

Linodes can be created using one of the available Types. See
[GET /linode/types](/docs/api/linode-types/#types-list) to get more
information about each Type's specs and cost.

Linodes can be created in any one of our available
[Regions](/docs/api/regions/#regions-list) for a list
of available Regions you can deploy your Linode in.

In an effort to fight spam, Linode restricts outbound connections on ports 25, 465, and 587
on all Linodes for new accounts created after November 5th, 2019. For more information,
see [Sending Email on Linode](/docs/email/running-a-mail-server/#sending-email-on-linode).

Linodes can be created in a number of ways:

* Using a Linode Linux Distribution image or an Image you created based on another Linode.
  * The Linode will be `running` after it completes `provisioning`.
  * A default config with two Disks, one being a 512 swap disk, is created.
    * `swap_size` can be used to customize the swap disk size.
  * Requires a `root_pass` be supplied to use for the root User's Account.
  * It is recommended to supply SSH keys for the root User using the `authorized_keys` field.
  * You may also supply a list of usernames via the `authorized_users` field.
    * These users must have an SSH Key associated with your Profile first. See [/profile/sshkeys](/docs/api/profile/#ssh-key-add) for more information.

* Using a StackScript.
  * See [/linode/stackscripts](/docs/api/stackscripts/#stackscripts-list) for
    a list of available StackScripts.
  * The Linode will be `running` after it completes `provisioning`.
  * Requires a compatible Image to be supplied.
    * See [/linode/stackscript/{stackscriptId}](/docs/api/stackscripts/#stackscript-view) for compatible Images.
  * Requires a `root_pass` be supplied to use for the root User's Account.
  * It is recommended to supply SSH keys for the root User using the `authorized_keys` field.
  * You may also supply a list of usernames via the `authorized_users` field.
    * These users must have an SSH Key associated with your Profile first. See [/profile/sshkeys](/docs/api/profile/#ssh-key-add) for more information.

* Using one of your other Linode's backups.
  * You must create a Linode large enough to accommodate the Backup's size.
  * The Disks and Config will match that of the Linode that was backed up.
  * The `root_pass` will match that of the Linode that was backed up.

* Create an empty Linode.
  * The Linode will remain `offline` and must be manually started.
    * See [POST /linode/instances/{linodeId}/boot](/docs/api/linode-instances/#linode-boot).
  * Disks and Configs must be created manually.
  * This is only recommended for advanced use cases.


**Important**: You must be an unrestricted User in order to add or modify
tags on Linodes.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) CreateLinodeInstance(ctx _context.Context) ApiCreateLinodeInstanceRequest {
	return ApiCreateLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Linode
 */
func (a *LinodeInstancesApiService) CreateLinodeInstanceExecute(r ApiCreateLinodeInstanceRequest) (Linode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Linode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.CreateLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject2 == nil {
		return localVarReturnValue, nil, reportError("inlineObject2 is required and must be specified")
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
	localVarPostBody = r.inlineObject2
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

type ApiCreateSnapshotRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject3 *InlineObject3
}

func (r ApiCreateSnapshotRequest) InlineObject3(inlineObject3 InlineObject3) ApiCreateSnapshotRequest {
	r.inlineObject3 = &inlineObject3
	return r
}

func (r ApiCreateSnapshotRequest) Execute() (Backup, *_nethttp.Response, error) {
	return r.ApiService.CreateSnapshotExecute(r)
}

/*
 * CreateSnapshot Snapshot Create
 * Creates a snapshot Backup of a Linode.
** If you already have a snapshot of this Linode, this is a destructive action. The previous snapshot will be deleted.**

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode the backups belong to.
 * @return ApiCreateSnapshotRequest
 */
func (a *LinodeInstancesApiService) CreateSnapshot(ctx _context.Context, linodeId int32) ApiCreateSnapshotRequest {
	return ApiCreateSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Backup
 */
func (a *LinodeInstancesApiService) CreateSnapshotExecute(r ApiCreateSnapshotRequest) (Backup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Backup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.CreateSnapshot")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject3 == nil {
		return localVarReturnValue, nil, reportError("inlineObject3 is required and must be specified")
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
	localVarPostBody = r.inlineObject3
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

type ApiDeleteDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
}


func (r ApiDeleteDiskRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteDiskExecute(r)
}

/*
 * DeleteDisk Disk Delete
 * Deletes a Disk you have permission to `read_write`.

**Deleting a Disk is a destructive action and cannot be undone.**

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to look up.
 * @return ApiDeleteDiskRequest
 */
func (a *LinodeInstancesApiService) DeleteDisk(ctx _context.Context, linodeId int32, diskId int32) ApiDeleteDiskRequest {
	return ApiDeleteDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) DeleteDiskExecute(r ApiDeleteDiskRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.DeleteDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

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

type ApiDeleteLinodeConfigRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	configId int32
}


func (r ApiDeleteLinodeConfigRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteLinodeConfigExecute(r)
}

/*
 * DeleteLinodeConfig Configuration Profile Delete
 * Deletes the specified Configuration profile from the specified Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode whose Configuration profile to look up.
 * @param configId The ID of the Configuration profile to look up.
 * @return ApiDeleteLinodeConfigRequest
 */
func (a *LinodeInstancesApiService) DeleteLinodeConfig(ctx _context.Context, linodeId int32, configId int32) ApiDeleteLinodeConfigRequest {
	return ApiDeleteLinodeConfigRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		configId: configId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) DeleteLinodeConfigExecute(r ApiDeleteLinodeConfigRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.DeleteLinodeConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", _neturl.PathEscape(parameterToString(r.configId, "")), -1)

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

type ApiDeleteLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiDeleteLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteLinodeInstanceExecute(r)
}

/*
 * DeleteLinodeInstance Linode Delete
 * Deletes a Linode you have permission to `read_write`.

**Deleting a Linode is a destructive action and cannot be undone.**

Additionally, deleting a Linode:

  * Gives up any IP addresses the Linode was assigned.
  * Deletes all Disks, Backups, Configs, etc.
  * Stops billing for the Linode and its associated services. You will be billed for time used
    within the billing period the Linode was active.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up
 * @return ApiDeleteLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) DeleteLinodeInstance(ctx _context.Context, linodeId int32) ApiDeleteLinodeInstanceRequest {
	return ApiDeleteLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) DeleteLinodeInstanceExecute(r ApiDeleteLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.DeleteLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiEnableBackupsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiEnableBackupsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.EnableBackupsExecute(r)
}

/*
 * EnableBackups Backups Enable
 * Enables backups for the specified Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to enable backup service for.
 * @return ApiEnableBackupsRequest
 */
func (a *LinodeInstancesApiService) EnableBackups(ctx _context.Context, linodeId int32) ApiEnableBackupsRequest {
	return ApiEnableBackupsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) EnableBackupsExecute(r ApiEnableBackupsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.EnableBackups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetBackupRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	backupId int32
}


func (r ApiGetBackupRequest) Execute() (Backup, *_nethttp.Response, error) {
	return r.ApiService.GetBackupExecute(r)
}

/*
 * GetBackup Backup View
 * Returns information about a Backup.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode the Backup belongs to.
 * @param backupId The ID of the Backup to look up.
 * @return ApiGetBackupRequest
 */
func (a *LinodeInstancesApiService) GetBackup(ctx _context.Context, linodeId int32, backupId int32) ApiGetBackupRequest {
	return ApiGetBackupRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		backupId: backupId,
	}
}

/*
 * Execute executes the request
 * @return Backup
 */
func (a *LinodeInstancesApiService) GetBackupExecute(r ApiGetBackupRequest) (Backup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Backup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetBackup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups/{backupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", _neturl.PathEscape(parameterToString(r.backupId, "")), -1)

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

type ApiGetBackupsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiGetBackupsRequest) Execute() (InlineResponse20016, *_nethttp.Response, error) {
	return r.ApiService.GetBackupsExecute(r)
}

/*
 * GetBackups Backups List
 * Returns information about this Linode's available backups.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode the backups belong to.
 * @return ApiGetBackupsRequest
 */
func (a *LinodeInstancesApiService) GetBackups(ctx _context.Context, linodeId int32) ApiGetBackupsRequest {
	return ApiGetBackupsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20016
 */
func (a *LinodeInstancesApiService) GetBackupsExecute(r ApiGetBackupsRequest) (InlineResponse20016, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetBackups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetKernelRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	kernelId string
}


func (r ApiGetKernelRequest) Execute() (Kernel, *_nethttp.Response, error) {
	return r.ApiService.GetKernelExecute(r)
}

/*
 * GetKernel Kernel View
 * Returns information about a single Kernel.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param kernelId ID of the Kernel to look up.
 * @return ApiGetKernelRequest
 */
func (a *LinodeInstancesApiService) GetKernel(ctx _context.Context, kernelId string) ApiGetKernelRequest {
	return ApiGetKernelRequest{
		ApiService: a,
		ctx: ctx,
		kernelId: kernelId,
	}
}

/*
 * Execute executes the request
 * @return Kernel
 */
func (a *LinodeInstancesApiService) GetKernelExecute(r ApiGetKernelRequest) (Kernel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Kernel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetKernel")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/kernels/{kernelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"kernelId"+"}", _neturl.PathEscape(parameterToString(r.kernelId, "")), -1)

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

type ApiGetKernelsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	page *int32
	pageSize *int32
}

func (r ApiGetKernelsRequest) Page(page int32) ApiGetKernelsRequest {
	r.page = &page
	return r
}
func (r ApiGetKernelsRequest) PageSize(pageSize int32) ApiGetKernelsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetKernelsRequest) Execute() (InlineResponse20020, *_nethttp.Response, error) {
	return r.ApiService.GetKernelsExecute(r)
}

/*
 * GetKernels Kernels List
 * Lists available Kernels.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetKernelsRequest
 */
func (a *LinodeInstancesApiService) GetKernels(ctx _context.Context) ApiGetKernelsRequest {
	return ApiGetKernelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20020
 */
func (a *LinodeInstancesApiService) GetKernelsExecute(r ApiGetKernelsRequest) (InlineResponse20020, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20020
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetKernels")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/kernels"

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

type ApiGetLinodeConfigRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	configId int32
}


func (r ApiGetLinodeConfigRequest) Execute() (LinodeConfig, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeConfigExecute(r)
}

/*
 * GetLinodeConfig Configuration Profile View
 * Returns information about a specific Configuration profile.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode whose Configuration profile to look up.
 * @param configId The ID of the Configuration profile to look up.
 * @return ApiGetLinodeConfigRequest
 */
func (a *LinodeInstancesApiService) GetLinodeConfig(ctx _context.Context, linodeId int32, configId int32) ApiGetLinodeConfigRequest {
	return ApiGetLinodeConfigRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		configId: configId,
	}
}

/*
 * Execute executes the request
 * @return LinodeConfig
 */
func (a *LinodeInstancesApiService) GetLinodeConfigExecute(r ApiGetLinodeConfigRequest) (LinodeConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LinodeConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", _neturl.PathEscape(parameterToString(r.configId, "")), -1)

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

type ApiGetLinodeConfigsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	page *int32
	pageSize *int32
}

func (r ApiGetLinodeConfigsRequest) Page(page int32) ApiGetLinodeConfigsRequest {
	r.page = &page
	return r
}
func (r ApiGetLinodeConfigsRequest) PageSize(pageSize int32) ApiGetLinodeConfigsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetLinodeConfigsRequest) Execute() (InlineResponse20017, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeConfigsExecute(r)
}

/*
 * GetLinodeConfigs Configuration Profiles List
 * Lists Configuration profiles associated with a Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up Configuration profiles for.
 * @return ApiGetLinodeConfigsRequest
 */
func (a *LinodeInstancesApiService) GetLinodeConfigs(ctx _context.Context, linodeId int32) ApiGetLinodeConfigsRequest {
	return ApiGetLinodeConfigsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20017
 */
func (a *LinodeInstancesApiService) GetLinodeConfigsExecute(r ApiGetLinodeConfigsRequest) (InlineResponse20017, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20017
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeConfigs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/configs"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
}


func (r ApiGetLinodeDiskRequest) Execute() (Disk, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeDiskExecute(r)
}

/*
 * GetLinodeDisk Disk View
 * View Disk information for a Disk associated with this Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to look up.
 * @return ApiGetLinodeDiskRequest
 */
func (a *LinodeInstancesApiService) GetLinodeDisk(ctx _context.Context, linodeId int32, diskId int32) ApiGetLinodeDiskRequest {
	return ApiGetLinodeDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return Disk
 */
func (a *LinodeInstancesApiService) GetLinodeDiskExecute(r ApiGetLinodeDiskRequest) (Disk, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Disk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

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

type ApiGetLinodeDisksRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	page *int32
	pageSize *int32
}

func (r ApiGetLinodeDisksRequest) Page(page int32) ApiGetLinodeDisksRequest {
	r.page = &page
	return r
}
func (r ApiGetLinodeDisksRequest) PageSize(pageSize int32) ApiGetLinodeDisksRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetLinodeDisksRequest) Execute() (InlineResponse20018, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeDisksExecute(r)
}

/*
 * GetLinodeDisks Disks List
 * View Disk information for Disks associated with this Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeDisksRequest
 */
func (a *LinodeInstancesApiService) GetLinodeDisks(ctx _context.Context, linodeId int32) ApiGetLinodeDisksRequest {
	return ApiGetLinodeDisksRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20018
 */
func (a *LinodeInstancesApiService) GetLinodeDisksExecute(r ApiGetLinodeDisksRequest) (InlineResponse20018, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20018
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeDisks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeFirewallsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	page *int32
	pageSize *int32
}

func (r ApiGetLinodeFirewallsRequest) Page(page int32) ApiGetLinodeFirewallsRequest {
	r.page = &page
	return r
}
func (r ApiGetLinodeFirewallsRequest) PageSize(pageSize int32) ApiGetLinodeFirewallsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetLinodeFirewallsRequest) Execute() (InlineResponse20019, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeFirewallsExecute(r)
}

/*
 * GetLinodeFirewalls Firewalls List
 * View Firewall information for Firewalls associated with this Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeFirewallsRequest
 */
func (a *LinodeInstancesApiService) GetLinodeFirewalls(ctx _context.Context, linodeId int32) ApiGetLinodeFirewallsRequest {
	return ApiGetLinodeFirewallsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20019
 */
func (a *LinodeInstancesApiService) GetLinodeFirewallsExecute(r ApiGetLinodeFirewallsRequest) (InlineResponse20019, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20019
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeFirewalls")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/firewalls"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeIPRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	address string
}


func (r ApiGetLinodeIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeIPExecute(r)
}

/*
 * GetLinodeIP IP Address View
 * View information about the specified IP address associated with the specified Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to look up.
 * @param address The IP address to look up.
 * @return ApiGetLinodeIPRequest
 */
func (a *LinodeInstancesApiService) GetLinodeIP(ctx _context.Context, linodeId int32, address string) ApiGetLinodeIPRequest {
	return ApiGetLinodeIPRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		address: address,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *LinodeInstancesApiService) GetLinodeIPExecute(r ApiGetLinodeIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/ips/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
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

type ApiGetLinodeIPsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiGetLinodeIPsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeIPsExecute(r)
}

/*
 * GetLinodeIPs Networking Information List
 * Returns networking information for a single Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeIPsRequest
 */
func (a *LinodeInstancesApiService) GetLinodeIPs(ctx _context.Context, linodeId int32) ApiGetLinodeIPsRequest {
	return ApiGetLinodeIPsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) GetLinodeIPsExecute(r ApiGetLinodeIPsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeIPs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/ips"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiGetLinodeInstanceRequest) Execute() (Linode, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeInstanceExecute(r)
}

/*
 * GetLinodeInstance Linode View
 * Get a specific Linode by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up
 * @return ApiGetLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) GetLinodeInstance(ctx _context.Context, linodeId int32) ApiGetLinodeInstanceRequest {
	return ApiGetLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Linode
 */
func (a *LinodeInstancesApiService) GetLinodeInstanceExecute(r ApiGetLinodeInstanceRequest) (Linode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Linode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeInstancesRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	page *int32
	pageSize *int32
}

func (r ApiGetLinodeInstancesRequest) Page(page int32) ApiGetLinodeInstancesRequest {
	r.page = &page
	return r
}
func (r ApiGetLinodeInstancesRequest) PageSize(pageSize int32) ApiGetLinodeInstancesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetLinodeInstancesRequest) Execute() (InlineResponse20015, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeInstancesExecute(r)
}

/*
 * GetLinodeInstances Linodes List
 * Returns a paginated list of Linodes you have permission to view.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetLinodeInstancesRequest
 */
func (a *LinodeInstancesApiService) GetLinodeInstances(ctx _context.Context) ApiGetLinodeInstancesRequest {
	return ApiGetLinodeInstancesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20015
 */
func (a *LinodeInstancesApiService) GetLinodeInstancesExecute(r ApiGetLinodeInstancesRequest) (InlineResponse20015, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20015
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeInstances")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances"

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

type ApiGetLinodeStatsRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiGetLinodeStatsRequest) Execute() (LinodeStats, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeStatsExecute(r)
}

/*
 * GetLinodeStats Linode Statistics View
 * Returns CPU, IO, IPv4, and IPv6 statistics for your Linode for the past 24 hours.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeStatsRequest
 */
func (a *LinodeInstancesApiService) GetLinodeStats(ctx _context.Context, linodeId int32) ApiGetLinodeStatsRequest {
	return ApiGetLinodeStatsRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return LinodeStats
 */
func (a *LinodeInstancesApiService) GetLinodeStatsExecute(r ApiGetLinodeStatsRequest) (LinodeStats, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LinodeStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/stats"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeStatsByYearMonthRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	year int32
	month int32
}


func (r ApiGetLinodeStatsByYearMonthRequest) Execute() (LinodeStats, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeStatsByYearMonthExecute(r)
}

/*
 * GetLinodeStatsByYearMonth Statistics View (year/month)
 * Returns statistics for a specific month. The year/month values must be either a date in the past, or the current month. If the current month, statistics will be retrieved for the past 30 days.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param year Numeric value representing the year to look up.
 * @param month Numeric value representing the month to look up.
 * @return ApiGetLinodeStatsByYearMonthRequest
 */
func (a *LinodeInstancesApiService) GetLinodeStatsByYearMonth(ctx _context.Context, linodeId int32, year int32, month int32) ApiGetLinodeStatsByYearMonthRequest {
	return ApiGetLinodeStatsByYearMonthRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		year: year,
		month: month,
	}
}

/*
 * Execute executes the request
 * @return LinodeStats
 */
func (a *LinodeInstancesApiService) GetLinodeStatsByYearMonthExecute(r ApiGetLinodeStatsByYearMonthRequest) (LinodeStats, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LinodeStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeStatsByYearMonth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/stats/{year}/{month}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"year"+"}", _neturl.PathEscape(parameterToString(r.year, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"month"+"}", _neturl.PathEscape(parameterToString(r.month, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.month < 1 {
		return localVarReturnValue, nil, reportError("month must be greater than 1")
	}
	if r.month > 12 {
		return localVarReturnValue, nil, reportError("month must be less than 12")
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

type ApiGetLinodeTransferRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiGetLinodeTransferRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeTransferExecute(r)
}

/*
 * GetLinodeTransfer Network Transfer View
 * Returns a Linode's network transfer pool statistics for the current month.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeTransferRequest
 */
func (a *LinodeInstancesApiService) GetLinodeTransfer(ctx _context.Context, linodeId int32) ApiGetLinodeTransferRequest {
	return ApiGetLinodeTransferRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) GetLinodeTransferExecute(r ApiGetLinodeTransferRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeTransfer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/transfer"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiGetLinodeTransferByYearMonthRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	year int32
	month int32
}


func (r ApiGetLinodeTransferByYearMonthRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeTransferByYearMonthExecute(r)
}

/*
 * GetLinodeTransferByYearMonth Network Transfer View (year/month)
 * Returns a Linode's network transfer statistics for a specific month. The year/month values must be either a date in the past, or the current month.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param year Numeric value representing the year to look up.
 * @param month Numeric value representing the month to look up.
 * @return ApiGetLinodeTransferByYearMonthRequest
 */
func (a *LinodeInstancesApiService) GetLinodeTransferByYearMonth(ctx _context.Context, linodeId int32, year int32, month int32) ApiGetLinodeTransferByYearMonthRequest {
	return ApiGetLinodeTransferByYearMonthRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		year: year,
		month: month,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) GetLinodeTransferByYearMonthExecute(r ApiGetLinodeTransferByYearMonthRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeTransferByYearMonth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/transfer/{year}/{month}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"year"+"}", _neturl.PathEscape(parameterToString(r.year, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"month"+"}", _neturl.PathEscape(parameterToString(r.month, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.year < 2000 {
		return localVarReturnValue, nil, reportError("year must be greater than 2000")
	}
	if r.year > 2037 {
		return localVarReturnValue, nil, reportError("year must be less than 2037")
	}
	if r.month < 1 {
		return localVarReturnValue, nil, reportError("month must be greater than 1")
	}
	if r.month > 12 {
		return localVarReturnValue, nil, reportError("month must be less than 12")
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

type ApiGetLinodeVolumesRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	page *int32
	pageSize *int32
}

func (r ApiGetLinodeVolumesRequest) Page(page int32) ApiGetLinodeVolumesRequest {
	r.page = &page
	return r
}
func (r ApiGetLinodeVolumesRequest) PageSize(pageSize int32) ApiGetLinodeVolumesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetLinodeVolumesRequest) Execute() (InlineResponse20021, *_nethttp.Response, error) {
	return r.ApiService.GetLinodeVolumesExecute(r)
}

/*
 * GetLinodeVolumes Linode's Volumes List
 * View Block Storage Volumes attached to this Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @return ApiGetLinodeVolumesRequest
 */
func (a *LinodeInstancesApiService) GetLinodeVolumes(ctx _context.Context, linodeId int32) ApiGetLinodeVolumesRequest {
	return ApiGetLinodeVolumesRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20021
 */
func (a *LinodeInstancesApiService) GetLinodeVolumesExecute(r ApiGetLinodeVolumesRequest) (InlineResponse20021, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20021
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.GetLinodeVolumes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/volumes"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiMigrateLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiMigrateLinodeInstanceRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiMigrateLinodeInstanceRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiMigrateLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.MigrateLinodeInstanceExecute(r)
}

/*
 * MigrateLinodeInstance DC Migration/Pending Host Migration Initiate
 * Initiate a pending host migration that has been scheduled by Linode or initiate a cross data center (DC) migration.  A list of pending migrations, if any, can be accessed from [GET /account/notifications](/docs/api/account/#notifications-list). When the migration begins, your Linode will be shutdown if not already off. If the migration initiated the shutdown, it will reboot the Linode when completed.

To initiate a cross DC migration, you must pass a `region` parameter to the request body specifying the target data center region. You can view a list of all available regions and their feature capabilities from [GET /regions](/docs/api/regions/#regions-list). If your Linode has a DC migration already queued or you have initiated a previously scheduled migration, you will not be able to initiate a DC migration until it has completed.

**Note:** Next Generation Network (NGN) data centers do not support IPv6 `/116` pools or IP Failover. If you have these features enabled on your Linode and attempt to migrate to an NGN data center, the migration will not initiate. If a Linode cannot be migrated because of an incompatibility, you will be prompted to select a different data center or contact support.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to migrate.
 * @return ApiMigrateLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) MigrateLinodeInstance(ctx _context.Context, linodeId int32) ApiMigrateLinodeInstanceRequest {
	return ApiMigrateLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) MigrateLinodeInstanceExecute(r ApiMigrateLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.MigrateLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/migrate"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiMutateLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject7 *InlineObject7
}

func (r ApiMutateLinodeInstanceRequest) InlineObject7(inlineObject7 InlineObject7) ApiMutateLinodeInstanceRequest {
	r.inlineObject7 = &inlineObject7
	return r
}

func (r ApiMutateLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.MutateLinodeInstanceExecute(r)
}

/*
 * MutateLinodeInstance Linode Upgrade
 * Linodes created with now-deprecated Types are entitled to a free upgrade to the next generation. A mutating Linode will be allocated any new resources the upgraded Type provides, and will be subsequently restarted if it was currently running.
If any actions are currently running or queued, those actions must be completed first before you can initiate a mutate.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to mutate.
 * @return ApiMutateLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) MutateLinodeInstance(ctx _context.Context, linodeId int32) ApiMutateLinodeInstanceRequest {
	return ApiMutateLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) MutateLinodeInstanceExecute(r ApiMutateLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.MutateLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/mutate"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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
	localVarPostBody = r.inlineObject7
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

type ApiRebootLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiRebootLinodeInstanceRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiRebootLinodeInstanceRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiRebootLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.RebootLinodeInstanceExecute(r)
}

/*
 * RebootLinodeInstance Linode Reboot
 * Reboots a Linode you have permission to modify. If any actions are currently running or queued, those actions must be completed first before you can initiate a reboot.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the linode to reboot.
 * @return ApiRebootLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) RebootLinodeInstance(ctx _context.Context, linodeId int32) ApiRebootLinodeInstanceRequest {
	return ApiRebootLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) RebootLinodeInstanceExecute(r ApiRebootLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.RebootLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/reboot"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiRebuildLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiRebuildLinodeInstanceRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiRebuildLinodeInstanceRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiRebuildLinodeInstanceRequest) Execute() (Linode, *_nethttp.Response, error) {
	return r.ApiService.RebuildLinodeInstanceExecute(r)
}

/*
 * RebuildLinodeInstance Linode Rebuild
 * Rebuilds a Linode you have the `read_write` permission to modify.
A rebuild will first shut down the Linode, delete all disks and configs on the Linode, and then deploy a new `image` to the Linode with the given attributes. Additionally:

  * Requires an `image` be supplied.
  * Requires a `root_pass` be supplied to use for the root User's Account.
  * It is recommended to supply SSH keys for the root User using the
    `authorized_keys` field.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to rebuild.
 * @return ApiRebuildLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) RebuildLinodeInstance(ctx _context.Context, linodeId int32) ApiRebuildLinodeInstanceRequest {
	return ApiRebuildLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Linode
 */
func (a *LinodeInstancesApiService) RebuildLinodeInstanceExecute(r ApiRebuildLinodeInstanceRequest) (Linode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Linode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.RebuildLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/rebuild"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiRemoveLinodeIPRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	address string
}


func (r ApiRemoveLinodeIPRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.RemoveLinodeIPExecute(r)
}

/*
 * RemoveLinodeIP IPv4 Address Delete
 * Deletes a public IPv4 address associated with this Linode. This will fail if it is the Linode's last remaining public IPv4 address. Private IPv4 addresses cannot be removed via this endpoint.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to look up.
 * @param address The IP address to look up.
 * @return ApiRemoveLinodeIPRequest
 */
func (a *LinodeInstancesApiService) RemoveLinodeIP(ctx _context.Context, linodeId int32, address string) ApiRemoveLinodeIPRequest {
	return ApiRemoveLinodeIPRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		address: address,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) RemoveLinodeIPExecute(r ApiRemoveLinodeIPRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.RemoveLinodeIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/ips/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
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

type ApiRescueLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject8 *InlineObject8
}

func (r ApiRescueLinodeInstanceRequest) InlineObject8(inlineObject8 InlineObject8) ApiRescueLinodeInstanceRequest {
	r.inlineObject8 = &inlineObject8
	return r
}

func (r ApiRescueLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.RescueLinodeInstanceExecute(r)
}

/*
 * RescueLinodeInstance Linode Boot into Rescue Mode
 * Rescue Mode is a safe environment for performing many system recovery and disk management tasks. Rescue Mode is based on the Finnix recovery distribution, a self-contained and bootable Linux distribution. You can also use Rescue Mode for tasks other than disaster recovery, such as formatting disks to use different filesystems, copying data between disks, and downloading files from a disk via SSH and SFTP.
* Note that "sdh" is reserved and unavailable during rescue.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to rescue.
 * @return ApiRescueLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) RescueLinodeInstance(ctx _context.Context, linodeId int32) ApiRescueLinodeInstanceRequest {
	return ApiRescueLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) RescueLinodeInstanceExecute(r ApiRescueLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.RescueLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/rescue"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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
	localVarPostBody = r.inlineObject8
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

type ApiResetDiskPasswordRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiResetDiskPasswordRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiResetDiskPasswordRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiResetDiskPasswordRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ResetDiskPasswordExecute(r)
}

/*
 * ResetDiskPassword Disk Root Password Reset
 * Resets the password of a Disk you have permission to `read_write`.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to look up.
 * @return ApiResetDiskPasswordRequest
 */
func (a *LinodeInstancesApiService) ResetDiskPassword(ctx _context.Context, linodeId int32, diskId int32) ApiResetDiskPasswordRequest {
	return ApiResetDiskPasswordRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) ResetDiskPasswordExecute(r ApiResetDiskPasswordRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.ResetDiskPassword")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}/password"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

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

type ApiResetLinodePasswordRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiResetLinodePasswordRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiResetLinodePasswordRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiResetLinodePasswordRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ResetLinodePasswordExecute(r)
}

/*
 * ResetLinodePassword Linode Root Password Reset
 * Resets the root password for this Linode.
* Your Linode must be [shut down](/docs/api/linode-instances/#linode-shut-down) for a password reset to complete.
* If your Linode has more than one disk (not counting its swap disk), use the [Reset Disk Root Password](/docs/api/linode-instances/#disk-root-password-reset) endpoint to update a specific disk's root password.
* A `password_reset` event is generated when a root password reset is successful.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode for which to reset its root password.
 * @return ApiResetLinodePasswordRequest
 */
func (a *LinodeInstancesApiService) ResetLinodePassword(ctx _context.Context, linodeId int32) ApiResetLinodePasswordRequest {
	return ApiResetLinodePasswordRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) ResetLinodePasswordExecute(r ApiResetLinodePasswordRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.ResetLinodePassword")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/password"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiResizeDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiResizeDiskRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiResizeDiskRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiResizeDiskRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ResizeDiskExecute(r)
}

/*
 * ResizeDisk Disk Resize
 * Resizes a Disk you have permission to `read_write`.
The Linode this Disk is attached to must be shut down for resizing to take effect.
If you are resizing the Disk to a smaller size, it cannot be made smaller than what is required by the total size of the files current on the Disk. The Disk must not be in use. If the Disk is in use, the request will succeed but the resize will ultimately fail.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to look up.
 * @return ApiResizeDiskRequest
 */
func (a *LinodeInstancesApiService) ResizeDisk(ctx _context.Context, linodeId int32, diskId int32) ApiResizeDiskRequest {
	return ApiResizeDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) ResizeDiskExecute(r ApiResizeDiskRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.ResizeDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}/resize"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

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

type ApiResizeLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	inlineObject9 *InlineObject9
}

func (r ApiResizeLinodeInstanceRequest) InlineObject9(inlineObject9 InlineObject9) ApiResizeLinodeInstanceRequest {
	r.inlineObject9 = &inlineObject9
	return r
}

func (r ApiResizeLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ResizeLinodeInstanceExecute(r)
}

/*
 * ResizeLinodeInstance Linode Resize
 * Resizes a Linode you have the `read_write` permission to a different Type. If any actions are currently running or queued, those actions must be completed first before you can initiate a resize. Additionally, the following criteria must be met in order to resize a Linode:

  * The Linode must not have a pending migration.
  * Your Account cannot have an outstanding balance.
  * The Linode must not have more disk allocation than the new Type allows.
    * In that situation, you must first delete or resize the disk to be smaller.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to resize.
 * @return ApiResizeLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) ResizeLinodeInstance(ctx _context.Context, linodeId int32) ApiResizeLinodeInstanceRequest {
	return ApiResizeLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) ResizeLinodeInstanceExecute(r ApiResizeLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.ResizeLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/resize"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject9 == nil {
		return localVarReturnValue, nil, reportError("inlineObject9 is required and must be specified")
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
	localVarPostBody = r.inlineObject9
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

type ApiRestoreBackupRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	backupId int32
	inlineObject4 *InlineObject4
}

func (r ApiRestoreBackupRequest) InlineObject4(inlineObject4 InlineObject4) ApiRestoreBackupRequest {
	r.inlineObject4 = &inlineObject4
	return r
}

func (r ApiRestoreBackupRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.RestoreBackupExecute(r)
}

/*
 * RestoreBackup Backup Restore
 * Restores a Linode's Backup to the specified Linode.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode that the Backup belongs to.
 * @param backupId The ID of the Backup to restore.
 * @return ApiRestoreBackupRequest
 */
func (a *LinodeInstancesApiService) RestoreBackup(ctx _context.Context, linodeId int32, backupId int32) ApiRestoreBackupRequest {
	return ApiRestoreBackupRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		backupId: backupId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) RestoreBackupExecute(r ApiRestoreBackupRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.RestoreBackup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/backups/{backupId}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", _neturl.PathEscape(parameterToString(r.backupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject4 == nil {
		return localVarReturnValue, nil, reportError("inlineObject4 is required and must be specified")
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
	localVarPostBody = r.inlineObject4
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

type ApiShutdownLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
}


func (r ApiShutdownLinodeInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ShutdownLinodeInstanceExecute(r)
}

/*
 * ShutdownLinodeInstance Linode Shut Down
 * Shuts down a Linode you have permission to modify. If any actions are currently running or queued, those actions must be completed first before you can initiate a shutdown.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to shutdown.
 * @return ApiShutdownLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) ShutdownLinodeInstance(ctx _context.Context, linodeId int32) ApiShutdownLinodeInstanceRequest {
	return ApiShutdownLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *LinodeInstancesApiService) ShutdownLinodeInstanceExecute(r ApiShutdownLinodeInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.ShutdownLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/shutdown"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

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

type ApiUpdateDiskRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	diskId int32
	disk *Disk
}

func (r ApiUpdateDiskRequest) Disk(disk Disk) ApiUpdateDiskRequest {
	r.disk = &disk
	return r
}

func (r ApiUpdateDiskRequest) Execute() (Disk, *_nethttp.Response, error) {
	return r.ApiService.UpdateDiskExecute(r)
}

/*
 * UpdateDisk Disk Update
 * Updates a Disk that you have permission to `read_write`.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up.
 * @param diskId ID of the Disk to look up.
 * @return ApiUpdateDiskRequest
 */
func (a *LinodeInstancesApiService) UpdateDisk(ctx _context.Context, linodeId int32, diskId int32) ApiUpdateDiskRequest {
	return ApiUpdateDiskRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		diskId: diskId,
	}
}

/*
 * Execute executes the request
 * @return Disk
 */
func (a *LinodeInstancesApiService) UpdateDiskExecute(r ApiUpdateDiskRequest) (Disk, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Disk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.UpdateDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/disks/{diskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"diskId"+"}", _neturl.PathEscape(parameterToString(r.diskId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.disk == nil {
		return localVarReturnValue, nil, reportError("disk is required and must be specified")
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
	localVarPostBody = r.disk
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

type ApiUpdateLinodeConfigRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	configId int32
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiUpdateLinodeConfigRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiUpdateLinodeConfigRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiUpdateLinodeConfigRequest) Execute() (LinodeConfig, *_nethttp.Response, error) {
	return r.ApiService.UpdateLinodeConfigExecute(r)
}

/*
 * UpdateLinodeConfig Configuration Profile Update
 * Updates a Configuration profile.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode whose Configuration profile to look up.
 * @param configId The ID of the Configuration profile to look up.
 * @return ApiUpdateLinodeConfigRequest
 */
func (a *LinodeInstancesApiService) UpdateLinodeConfig(ctx _context.Context, linodeId int32, configId int32) ApiUpdateLinodeConfigRequest {
	return ApiUpdateLinodeConfigRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		configId: configId,
	}
}

/*
 * Execute executes the request
 * @return LinodeConfig
 */
func (a *LinodeInstancesApiService) UpdateLinodeConfigExecute(r ApiUpdateLinodeConfigRequest) (LinodeConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LinodeConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.UpdateLinodeConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/configs/{configId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", _neturl.PathEscape(parameterToString(r.configId, "")), -1)

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

type ApiUpdateLinodeIPRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	address string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiUpdateLinodeIPRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiUpdateLinodeIPRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiUpdateLinodeIPRequest) Execute() (IPAddress, *_nethttp.Response, error) {
	return r.ApiService.UpdateLinodeIPExecute(r)
}

/*
 * UpdateLinodeIP IP Address Update
 * Updates a particular IP Address associated with this Linode. Only allows setting/resetting reverse DNS.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId The ID of the Linode to look up.
 * @param address The IP address to look up.
 * @return ApiUpdateLinodeIPRequest
 */
func (a *LinodeInstancesApiService) UpdateLinodeIP(ctx _context.Context, linodeId int32, address string) ApiUpdateLinodeIPRequest {
	return ApiUpdateLinodeIPRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
		address: address,
	}
}

/*
 * Execute executes the request
 * @return IPAddress
 */
func (a *LinodeInstancesApiService) UpdateLinodeIPExecute(r ApiUpdateLinodeIPRequest) (IPAddress, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IPAddress
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.UpdateLinodeIP")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}/ips/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

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

type ApiUpdateLinodeInstanceRequest struct {
	ctx _context.Context
	ApiService *LinodeInstancesApiService
	linodeId int32
	linode *Linode
}

func (r ApiUpdateLinodeInstanceRequest) Linode(linode Linode) ApiUpdateLinodeInstanceRequest {
	r.linode = &linode
	return r
}

func (r ApiUpdateLinodeInstanceRequest) Execute() (Linode, *_nethttp.Response, error) {
	return r.ApiService.UpdateLinodeInstanceExecute(r)
}

/*
 * UpdateLinodeInstance Linode Update
 * Updates a Linode that you have permission to `read_write`.

**Important**: You must be an unrestricted User in order to add or modify tags on Linodes.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param linodeId ID of the Linode to look up
 * @return ApiUpdateLinodeInstanceRequest
 */
func (a *LinodeInstancesApiService) UpdateLinodeInstance(ctx _context.Context, linodeId int32) ApiUpdateLinodeInstanceRequest {
	return ApiUpdateLinodeInstanceRequest{
		ApiService: a,
		ctx: ctx,
		linodeId: linodeId,
	}
}

/*
 * Execute executes the request
 * @return Linode
 */
func (a *LinodeInstancesApiService) UpdateLinodeInstanceExecute(r ApiUpdateLinodeInstanceRequest) (Linode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Linode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeInstancesApiService.UpdateLinodeInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/linode/instances/{linodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"linodeId"+"}", _neturl.PathEscape(parameterToString(r.linodeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.linode == nil {
		return localVarReturnValue, nil, reportError("linode is required and must be specified")
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
	localVarPostBody = r.linode
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
