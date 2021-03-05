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

// ObjectStorageApiService ObjectStorageApi service
type ObjectStorageApiService service

type ApiCancelObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
}


func (r ApiCancelObjectStorageRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CancelObjectStorageExecute(r)
}

/*
 * CancelObjectStorage Object Storage Cancel
 * Cancel Object Storage on an Account. All buckets on the Account must be empty
before Object Storage can be cancelled:

- To delete a smaller number of objects, review the instructions in our
[How to Use Object Storage](/docs/platform/object-storage/how-to-use-object-storage/)
guide.
- To delete large amounts of objects, consult our guide on
[Lifecycle Policies](/docs/platform/object-storage/lifecycle-policies/).

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCancelObjectStorageRequest
 */
func (a *ObjectStorageApiService) CancelObjectStorage(ctx _context.Context) ApiCancelObjectStorageRequest {
	return ApiCancelObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) CancelObjectStorageExecute(r ApiCancelObjectStorageRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.CancelObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/cancel"

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

type ApiCreateObjectStorageBucketRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	inlineObject13 *InlineObject13
}

func (r ApiCreateObjectStorageBucketRequest) InlineObject13(inlineObject13 InlineObject13) ApiCreateObjectStorageBucketRequest {
	r.inlineObject13 = &inlineObject13
	return r
}

func (r ApiCreateObjectStorageBucketRequest) Execute() (ObjectStorageBucket, *_nethttp.Response, error) {
	return r.ApiService.CreateObjectStorageBucketExecute(r)
}

/*
 * CreateObjectStorageBucket Object Storage Bucket Create
 * Creates an Object Storage Bucket in the cluster specified. If the
bucket already exists and is owned by you, this endpoint will return a
`200` response with that bucket as if it had just been created.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#put-bucket) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateObjectStorageBucketRequest
 */
func (a *ObjectStorageApiService) CreateObjectStorageBucket(ctx _context.Context) ApiCreateObjectStorageBucketRequest {
	return ApiCreateObjectStorageBucketRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageBucket
 */
func (a *ObjectStorageApiService) CreateObjectStorageBucketExecute(r ApiCreateObjectStorageBucketRequest) (ObjectStorageBucket, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageBucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.CreateObjectStorageBucket")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets"

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
	localVarPostBody = r.inlineObject13
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

type ApiCreateObjectStorageKeysRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	objectStorageKey *ObjectStorageKey
}

func (r ApiCreateObjectStorageKeysRequest) ObjectStorageKey(objectStorageKey ObjectStorageKey) ApiCreateObjectStorageKeysRequest {
	r.objectStorageKey = &objectStorageKey
	return r
}

func (r ApiCreateObjectStorageKeysRequest) Execute() (ObjectStorageKey, *_nethttp.Response, error) {
	return r.ApiService.CreateObjectStorageKeysExecute(r)
}

/*
 * CreateObjectStorageKeys Object Storage Key Create
 * Provisions a new Object Storage Key on your account.

* To create a Limited Access Key with specific permissions, send a `bucket_access`
array.

* To create a Limited Access Key without access to any buckets, send an empty
`bucket_access` array.

* To create an Access Key with unlimited access to all clusters and all buckets,
omit the `bucket_access` array.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateObjectStorageKeysRequest
 */
func (a *ObjectStorageApiService) CreateObjectStorageKeys(ctx _context.Context) ApiCreateObjectStorageKeysRequest {
	return ApiCreateObjectStorageKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageKey
 */
func (a *ObjectStorageApiService) CreateObjectStorageKeysExecute(r ApiCreateObjectStorageKeysRequest) (ObjectStorageKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.CreateObjectStorageKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/keys"

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
	localVarPostBody = r.objectStorageKey
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

type ApiCreateObjectStorageObjectURLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiCreateObjectStorageObjectURLRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiCreateObjectStorageObjectURLRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiCreateObjectStorageObjectURLRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateObjectStorageObjectURLExecute(r)
}

/*
 * CreateObjectStorageObjectURL Object Storage Object URL Create
 * Creates a pre-signed URL to access a single Object in a bucket. This
can be used to share objects, and also to create/delete objects by using
the appropriate HTTP method in your request body's `method` parameter.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/)
directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiCreateObjectStorageObjectURLRequest
 */
func (a *ObjectStorageApiService) CreateObjectStorageObjectURL(ctx _context.Context, clusterId string, bucket string) ApiCreateObjectStorageObjectURLRequest {
	return ApiCreateObjectStorageObjectURLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) CreateObjectStorageObjectURLExecute(r ApiCreateObjectStorageObjectURLRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.CreateObjectStorageObjectURL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/object-url"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiCreateObjectStorageSSLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	objectStorageSSL *ObjectStorageSSL
}

func (r ApiCreateObjectStorageSSLRequest) ObjectStorageSSL(objectStorageSSL ObjectStorageSSL) ApiCreateObjectStorageSSLRequest {
	r.objectStorageSSL = &objectStorageSSL
	return r
}

func (r ApiCreateObjectStorageSSLRequest) Execute() (ObjectStorageSSLResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateObjectStorageSSLExecute(r)
}

/*
 * CreateObjectStorageSSL Object Storage TLS/SSL Cert Upload
 * Upload a TLS/SSL certificate and private key to be served when you visit your Object Storage bucket via HTTPS.
Your TLS/SSL certificate and private key are stored encrypted at rest.


To replace an expired certificate, [delete your current certificate](/docs/api/object-storage/#object-storage-tlsssl-cert-delete)
and upload a new one.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiCreateObjectStorageSSLRequest
 */
func (a *ObjectStorageApiService) CreateObjectStorageSSL(ctx _context.Context, clusterId string, bucket string) ApiCreateObjectStorageSSLRequest {
	return ApiCreateObjectStorageSSLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageSSLResponse
 */
func (a *ObjectStorageApiService) CreateObjectStorageSSLExecute(r ApiCreateObjectStorageSSLRequest) (ObjectStorageSSLResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageSSLResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.CreateObjectStorageSSL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/ssl"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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
	localVarPostBody = r.objectStorageSSL
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

type ApiDeleteObjectStorageBucketRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
}


func (r ApiDeleteObjectStorageBucketRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteObjectStorageBucketExecute(r)
}

/*
 * DeleteObjectStorageBucket Object Storage Bucket Remove
 * Removes a single bucket. While buckets containing objects _may_ be deleted by including the `force` option in the request, such operations will fail if the bucket contains too many objects. The recommended way to empty large buckets is to use the [S3 API to configure lifecycle policies](https://docs.ceph.com/en/latest/radosgw/bucketpolicy/#) that remove all objects, then delete the bucket.

This endpoint is available for convenience. It is recommended that instead you use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#delete-bucket) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiDeleteObjectStorageBucketRequest
 */
func (a *ObjectStorageApiService) DeleteObjectStorageBucket(ctx _context.Context, clusterId string, bucket string) ApiDeleteObjectStorageBucketRequest {
	return ApiDeleteObjectStorageBucketRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) DeleteObjectStorageBucketExecute(r ApiDeleteObjectStorageBucketRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.DeleteObjectStorageBucket")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiDeleteObjectStorageKeyRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	keyId int32
}


func (r ApiDeleteObjectStorageKeyRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteObjectStorageKeyExecute(r)
}

/*
 * DeleteObjectStorageKey Object Storage Key Revoke
 * Revokes an Object Storage Key.  This keypair will no longer be usable by third-party clients.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param keyId The key to look up.
 * @return ApiDeleteObjectStorageKeyRequest
 */
func (a *ObjectStorageApiService) DeleteObjectStorageKey(ctx _context.Context, keyId int32) ApiDeleteObjectStorageKeyRequest {
	return ApiDeleteObjectStorageKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) DeleteObjectStorageKeyExecute(r ApiDeleteObjectStorageKeyRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.DeleteObjectStorageKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

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

type ApiDeleteObjectStorageSSLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
}


func (r ApiDeleteObjectStorageSSLRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteObjectStorageSSLExecute(r)
}

/*
 * DeleteObjectStorageSSL Object Storage TLS/SSL Cert Delete
 * Deletes this Object Storage bucket's user uploaded TLS/SSL certificate and private key.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiDeleteObjectStorageSSLRequest
 */
func (a *ObjectStorageApiService) DeleteObjectStorageSSL(ctx _context.Context, clusterId string, bucket string) ApiDeleteObjectStorageSSLRequest {
	return ApiDeleteObjectStorageSSLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) DeleteObjectStorageSSLExecute(r ApiDeleteObjectStorageSSLRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.DeleteObjectStorageSSL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/ssl"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiGetObjectStorageBucketRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
}


func (r ApiGetObjectStorageBucketRequest) Execute() (ObjectStorageBucket, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageBucketExecute(r)
}

/*
 * GetObjectStorageBucket Object Storage Bucket View
 * Returns a single Object Storage Bucket.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#get-bucket) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiGetObjectStorageBucketRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageBucket(ctx _context.Context, clusterId string, bucket string) ApiGetObjectStorageBucketRequest {
	return ApiGetObjectStorageBucketRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageBucket
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketExecute(r ApiGetObjectStorageBucketRequest) (ObjectStorageBucket, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageBucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageBucket")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiGetObjectStorageBucketContentRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	marker *string
	delimiter *string
	prefix *string
	pageSize *int32
}

func (r ApiGetObjectStorageBucketContentRequest) Marker(marker string) ApiGetObjectStorageBucketContentRequest {
	r.marker = &marker
	return r
}
func (r ApiGetObjectStorageBucketContentRequest) Delimiter(delimiter string) ApiGetObjectStorageBucketContentRequest {
	r.delimiter = &delimiter
	return r
}
func (r ApiGetObjectStorageBucketContentRequest) Prefix(prefix string) ApiGetObjectStorageBucketContentRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetObjectStorageBucketContentRequest) PageSize(pageSize int32) ApiGetObjectStorageBucketContentRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetObjectStorageBucketContentRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageBucketContentExecute(r)
}

/*
 * GetObjectStorageBucketContent Object Storage Bucket Contents List
 * Returns the contents of a bucket. The contents are paginated using a `marker`,
which is the name of the last object on the previous page.  Objects may
be filtered by `prefix` and `delimiter` as well; see Query Parameters for more
information.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/objectops/#get-object) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiGetObjectStorageBucketContentRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketContent(ctx _context.Context, clusterId string, bucket string) ApiGetObjectStorageBucketContentRequest {
	return ApiGetObjectStorageBucketContentRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketContentExecute(r ApiGetObjectStorageBucketContentRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageBucketContent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/object-list"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.marker != nil {
		localVarQueryParams.Add("marker", parameterToString(*r.marker, ""))
	}
	if r.delimiter != nil {
		localVarQueryParams.Add("delimiter", parameterToString(*r.delimiter, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
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

type ApiGetObjectStorageBucketinClusterRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
}


func (r ApiGetObjectStorageBucketinClusterRequest) Execute() (InlineResponse20047, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageBucketinClusterExecute(r)
}

/*
 * GetObjectStorageBucketinCluster Object Storage Buckets in Cluster List
 * Returns a list of Buckets in this cluster belonging to this Account.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#get-bucket) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @return ApiGetObjectStorageBucketinClusterRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketinCluster(ctx _context.Context, clusterId string) ApiGetObjectStorageBucketinClusterRequest {
	return ApiGetObjectStorageBucketinClusterRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20047
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketinClusterExecute(r ApiGetObjectStorageBucketinClusterRequest) (InlineResponse20047, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20047
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageBucketinCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)

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

type ApiGetObjectStorageBucketsRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
}


func (r ApiGetObjectStorageBucketsRequest) Execute() (InlineResponse20047, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageBucketsExecute(r)
}

/*
 * GetObjectStorageBuckets Object Storage Buckets List
 * Returns a paginated list of all Object Storage Buckets that you own.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/serviceops/#list-buckets) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetObjectStorageBucketsRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageBuckets(ctx _context.Context) ApiGetObjectStorageBucketsRequest {
	return ApiGetObjectStorageBucketsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20047
 */
func (a *ObjectStorageApiService) GetObjectStorageBucketsExecute(r ApiGetObjectStorageBucketsRequest) (InlineResponse20047, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20047
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageBuckets")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets"

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

type ApiGetObjectStorageClusterRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
}


func (r ApiGetObjectStorageClusterRequest) Execute() (ObjectStorageCluster, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageClusterExecute(r)
}

/*
 * GetObjectStorageCluster Cluster View
 * Returns a single Object Storage Cluster.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the Cluster.
 * @return ApiGetObjectStorageClusterRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageCluster(ctx _context.Context, clusterId string) ApiGetObjectStorageClusterRequest {
	return ApiGetObjectStorageClusterRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageCluster
 */
func (a *ObjectStorageApiService) GetObjectStorageClusterExecute(r ApiGetObjectStorageClusterRequest) (ObjectStorageCluster, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageCluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/clusters/{clusterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)

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

type ApiGetObjectStorageClustersRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
}


func (r ApiGetObjectStorageClustersRequest) Execute() (InlineResponse20049, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageClustersExecute(r)
}

/*
 * GetObjectStorageClusters Clusters List
 * Returns a paginated list of Object Storage Clusters that are available for
use.  Users can connect to the clusters with third party clients to create buckets
and upload objects.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetObjectStorageClustersRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageClusters(ctx _context.Context) ApiGetObjectStorageClustersRequest {
	return ApiGetObjectStorageClustersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20049
 */
func (a *ObjectStorageApiService) GetObjectStorageClustersExecute(r ApiGetObjectStorageClustersRequest) (InlineResponse20049, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20049
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/clusters"

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

type ApiGetObjectStorageKeyRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	keyId int32
}


func (r ApiGetObjectStorageKeyRequest) Execute() (ObjectStorageKey, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageKeyExecute(r)
}

/*
 * GetObjectStorageKey Object Storage Key View
 * Returns a single Object Storage Key provisioned for your account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param keyId The key to look up.
 * @return ApiGetObjectStorageKeyRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageKey(ctx _context.Context, keyId int32) ApiGetObjectStorageKeyRequest {
	return ApiGetObjectStorageKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageKey
 */
func (a *ObjectStorageApiService) GetObjectStorageKeyExecute(r ApiGetObjectStorageKeyRequest) (ObjectStorageKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

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

type ApiGetObjectStorageKeysRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
}


func (r ApiGetObjectStorageKeysRequest) Execute() (InlineResponse20050, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageKeysExecute(r)
}

/*
 * GetObjectStorageKeys Object Storage Keys List
 * Returns a paginated list of Object Storage Keys for authenticating to
the Object Storage S3 API.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetObjectStorageKeysRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageKeys(ctx _context.Context) ApiGetObjectStorageKeysRequest {
	return ApiGetObjectStorageKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20050
 */
func (a *ObjectStorageApiService) GetObjectStorageKeysExecute(r ApiGetObjectStorageKeysRequest) (InlineResponse20050, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20050
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/keys"

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

type ApiGetObjectStorageSSLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
}


func (r ApiGetObjectStorageSSLRequest) Execute() (ObjectStorageSSLResponse, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageSSLExecute(r)
}

/*
 * GetObjectStorageSSL Object Storage TLS/SSL Cert View
 * Returns a boolean value indicating if this bucket has a corresponding TLS/SSL certificate that was
uploaded by an Account user.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiGetObjectStorageSSLRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageSSL(ctx _context.Context, clusterId string, bucket string) ApiGetObjectStorageSSLRequest {
	return ApiGetObjectStorageSSLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageSSLResponse
 */
func (a *ObjectStorageApiService) GetObjectStorageSSLExecute(r ApiGetObjectStorageSSLRequest) (ObjectStorageSSLResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageSSLResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageSSL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/ssl"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiGetObjectStorageTransferRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
}


func (r ApiGetObjectStorageTransferRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageTransferExecute(r)
}

/*
 * GetObjectStorageTransfer Object Storage Transfer View
 * The amount of outbound data transfer used by your account's Object Storage buckets.
Object Storage adds 1 terabyte of outbound data transfer to your data transfer pool.
See the [Object Storage Pricing and Limitations](/docs/guides/pricing-and-limitations/)
guide for details on Object Storage transfer quotas.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetObjectStorageTransferRequest
 */
func (a *ObjectStorageApiService) GetObjectStorageTransfer(ctx _context.Context) ApiGetObjectStorageTransferRequest {
	return ApiGetObjectStorageTransferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) GetObjectStorageTransferExecute(r ApiGetObjectStorageTransferRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.GetObjectStorageTransfer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/transfer"

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

type ApiModifyObjectStorageBucketAccessRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiModifyObjectStorageBucketAccessRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiModifyObjectStorageBucketAccessRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiModifyObjectStorageBucketAccessRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ModifyObjectStorageBucketAccessExecute(r)
}

/*
 * ModifyObjectStorageBucketAccess Object Storage Bucket Access Modify
 * Allows changing basic Cross-origin Resource Sharing (CORS) and Access Control Level (ACL) settings.
Only allows enabling/disabling CORS for all origins, and/or setting canned ACLs.


For more fine-grained control of both systems, please use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#put-bucket-acl) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiModifyObjectStorageBucketAccessRequest
 */
func (a *ObjectStorageApiService) ModifyObjectStorageBucketAccess(ctx _context.Context, clusterId string, bucket string) ApiModifyObjectStorageBucketAccessRequest {
	return ApiModifyObjectStorageBucketAccessRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) ModifyObjectStorageBucketAccessExecute(r ApiModifyObjectStorageBucketAccessRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.ModifyObjectStorageBucketAccess")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/access"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiUpdateObjectStorageBucketACLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiUpdateObjectStorageBucketACLRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiUpdateObjectStorageBucketACLRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiUpdateObjectStorageBucketACLRequest) Execute() (InlineResponse20048, *_nethttp.Response, error) {
	return r.ApiService.UpdateObjectStorageBucketACLExecute(r)
}

/*
 * UpdateObjectStorageBucketACL Object Storage Object ACL Config Update
 * Update an Object's configured Access Control List (ACL) in this Object Storage bucket.
ACLs define who can access your buckets and objects and specify the level of access
granted to those users.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/objectops/#set-object-acl) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiUpdateObjectStorageBucketACLRequest
 */
func (a *ObjectStorageApiService) UpdateObjectStorageBucketACL(ctx _context.Context, clusterId string, bucket string) ApiUpdateObjectStorageBucketACLRequest {
	return ApiUpdateObjectStorageBucketACLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20048
 */
func (a *ObjectStorageApiService) UpdateObjectStorageBucketACLExecute(r ApiUpdateObjectStorageBucketACLRequest) (InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20048
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.UpdateObjectStorageBucketACL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/object-acl"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiUpdateObjectStorageBucketAccessRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiUpdateObjectStorageBucketAccessRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiUpdateObjectStorageBucketAccessRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiUpdateObjectStorageBucketAccessRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateObjectStorageBucketAccessExecute(r)
}

/*
 * UpdateObjectStorageBucketAccess Object Storage Bucket Access Update
 * Allows changing basic Cross-origin Resource Sharing (CORS) and Access Control Level (ACL) settings.
Only allows enabling/disabling CORS for all origins, and/or setting canned ACLs.


For more fine-grained control of both systems, please use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/bucketops/#put-bucket-acl) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiUpdateObjectStorageBucketAccessRequest
 */
func (a *ObjectStorageApiService) UpdateObjectStorageBucketAccess(ctx _context.Context, clusterId string, bucket string) ApiUpdateObjectStorageBucketAccessRequest {
	return ApiUpdateObjectStorageBucketAccessRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ObjectStorageApiService) UpdateObjectStorageBucketAccessExecute(r ApiUpdateObjectStorageBucketAccessRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.UpdateObjectStorageBucketAccess")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/access"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

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

type ApiUpdateObjectStorageKeyRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	keyId int32
	inlineObject14 *InlineObject14
}

func (r ApiUpdateObjectStorageKeyRequest) InlineObject14(inlineObject14 InlineObject14) ApiUpdateObjectStorageKeyRequest {
	r.inlineObject14 = &inlineObject14
	return r
}

func (r ApiUpdateObjectStorageKeyRequest) Execute() (ObjectStorageKey, *_nethttp.Response, error) {
	return r.ApiService.UpdateObjectStorageKeyExecute(r)
}

/*
 * UpdateObjectStorageKey Object Storage Key Update
 * Updates an Object Storage Key on your account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param keyId The key to look up.
 * @return ApiUpdateObjectStorageKeyRequest
 */
func (a *ObjectStorageApiService) UpdateObjectStorageKey(ctx _context.Context, keyId int32) ApiUpdateObjectStorageKeyRequest {
	return ApiUpdateObjectStorageKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
	}
}

/*
 * Execute executes the request
 * @return ObjectStorageKey
 */
func (a *ObjectStorageApiService) UpdateObjectStorageKeyExecute(r ApiUpdateObjectStorageKeyRequest) (ObjectStorageKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStorageKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.UpdateObjectStorageKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

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
	localVarPostBody = r.inlineObject14
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

type ApiViewObjectStorageBucketACLRequest struct {
	ctx _context.Context
	ApiService *ObjectStorageApiService
	clusterId string
	bucket string
	name *string
}

func (r ApiViewObjectStorageBucketACLRequest) Name(name string) ApiViewObjectStorageBucketACLRequest {
	r.name = &name
	return r
}

func (r ApiViewObjectStorageBucketACLRequest) Execute() (InlineResponse20048, *_nethttp.Response, error) {
	return r.ApiService.ViewObjectStorageBucketACLExecute(r)
}

/*
 * ViewObjectStorageBucketACL Object Storage Object ACL Config View
 * View an Object’s configured Access Control List (ACL) in this Object Storage bucket.
ACLs define who can access your buckets and objects and specify the level of access
granted to those users.


This endpoint is available for convenience. It is recommended that instead you
use the more [fully-featured S3 API](https://docs.ceph.com/en/latest/radosgw/s3/objectops/#get-object-acl) directly.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The ID of the cluster this bucket exists in.
 * @param bucket The bucket name.
 * @return ApiViewObjectStorageBucketACLRequest
 */
func (a *ObjectStorageApiService) ViewObjectStorageBucketACL(ctx _context.Context, clusterId string, bucket string) ApiViewObjectStorageBucketACLRequest {
	return ApiViewObjectStorageBucketACLRequest{
		ApiService: a,
		ctx: ctx,
		clusterId: clusterId,
		bucket: bucket,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20048
 */
func (a *ObjectStorageApiService) ViewObjectStorageBucketACLExecute(r ApiViewObjectStorageBucketACLRequest) (InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20048
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStorageApiService.ViewObjectStorageBucketACL")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-storage/buckets/{clusterId}/{bucket}/object-acl"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterId"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket"+"}", _neturl.PathEscape(parameterToString(r.bucket, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
