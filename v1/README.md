# Go API client for v1

## Introduction
The Linode API provides the ability to programmatically manage the full
range of Linode products and services.

This reference is designed to assist application developers and system
administrators.  Each endpoint includes descriptions, request syntax, and
examples using standard HTTP requests. Response data is returned in JSON
format.


This document was generated from our OpenAPI Specification.  See the
<a target=\"_top\" href=\"https://www.openapis.org\">OpenAPI website</a> for more information.

<a target=\"_top\" href=\"/docs/api/openapi.yaml\">Download the Linode OpenAPI Specification</a>.


## Changelog

<a target=\"_top\" href=\"https://developers.linode.com/changelog\">View our Changelog</a> to see release
notes on all changes made to our API.

## Access and Authentication

Some endpoints are publicly accessible without requiring authentication.
All endpoints affecting your Account, however, require either a Personal
Access Token or OAuth authentication (when using third-party
applications).

### Personal Access Token

The easiest way to access the API is with a Personal Access Token (PAT)
generated from the
<a target=\"_top\" href=\"https://cloud.linode.com/profile/tokens\">Linode Cloud Manager</a> or
the [Create Personal Access Token](/docs/api/profile/#personal-access-token-create) endpoint.

All scopes for the OAuth security model ([defined below](/docs/api/profile/#oauth)) apply to this
security model as well.

#### Authentication

| Security Scheme Type: | HTTP |
|-----------------------|------|
| **HTTP Authorization Scheme** | bearer |

### OAuth
If you only need to access the Linode API for personal use,
we recommend that you create a [personal access token](/docs/api/#personal-access-token).
If you're designing an application that can authenticate with an arbitrary Linode user, then
you should use the OAuth 2.0 workflows presented in this section.

For a more detailed example of an OAuth 2.0 implementation, see our guide on [How to Create an OAuth App with the Linode Python API Library](/docs/platform/api/how-to-create-an-oauth-app-with-the-linode-python-api-library/#oauth-2-authentication-exchange).

Before you implement OAuth in your application, you first need to create an OAuth client. You can do this [with the Linode API](/docs/api/account/#oauth-client-create) or [via the Cloud Manager](https://cloud.linode.com/profile/clients):

  - When creating the client, you'll supply a `label` and a `redirect_uri` (referred to as the Callback URL in the Cloud Manager).
  - The response from this endpoint will give you a `client_id` and a `secret`.
  - Clients can be public or private, and are private by default. You can choose to make the client public when it is created.
    - A private client is used with applications which can securely store the client secret (that is, the secret returned to you when you first created the client). For example, an application running on a secured server that only the developer has access to would use a private OAuth client. This is also called a confidential client in some OAuth documentation.
    - A public client is used with applications where the client secret is not guaranteed to be secure. For example, a native app running on a user's computer may not be able to keep the client secret safe, as a user could potentially inspect the source of the application. So, native apps or apps that run in a user's browser should use a public client.
    - Public and private clients follow different workflows, as described below.

#### OAuth Workflow

The OAuth workflow is a series of exchanges between your third-party app and Linode. The workflow is used
to authenticate a user before an application can start making API calls on the user's behalf.

Notes:

- With respect to the diagram in [section 1.2 of RFC 6749](https://tools.ietf.org/html/rfc6749#section-1.2), login.linode.com (referred to in this section as the *login server*)
is the Resource Owner and the Authorization Server; api.linode.com (referred to here as the *api server*) is the Resource Server.
- The OAuth spec refers to the private and public workflows listed below as the [authorization code flow](https://tools.ietf.org/html/rfc6749#section-1.3.1) and [implicit flow](https://tools.ietf.org/html/rfc6749#section-1.3.2).

| PRIVATE WORKFLOW | PUBLIC WORKFLOW |
|------------------|------------------|
| 1.  The user visits the application's website and is directed to login with Linode. | 1.  The user visits the application's website and is directed to login with Linode. |
| 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. | 2.  Your application then redirects the user to Linode's [login server](https://login.linode.com) with the client application's `client_id` and requested OAuth `scope`, which should appear in the URL of the login page. |
| 3.  The user logs into the login server with their username and password. | 3.  The user logs into the login server with their username and password. |
| 4.  The login server redirects the user to the specificed redirect URL with a temporary authorization `code` (exchange code) in the URL. | 4.  The login server redirects the user back to your application with an OAuth `access_token` embedded in the redirect URL's hash. This is temporary and expires in two hours. No `refresh_token` is issued. Therefore, once the `access_token` expires, a new one will need to be issued by having the user log in again. |
| 5.  The application issues a POST request (*see below*) to the login server with the exchange code, `client_id`, and the client application's `client_secret`. | |
| 6.  The login server responds to the client application with a new OAuth `access_token` and `refresh_token`. The `access_token` is set to expire in two hours. | |
| 7.  The `refresh_token` can be used by contacting the login server with the `client_id`, `client_secret`, `grant_type`, and `refresh_token` to get a new OAuth `access_token` and `refresh_token`. The new `access_token` is good for another two hours, and the new `refresh_token`, can be used to extend the session again by this same method. | |

#### OAuth Private Workflow - Additional Details

The following information expands on steps 5 through 7 of the private workflow:

Once the user has logged into Linode and you have received an exchange code,
you will need to trade that exchange code for an `access_token` and `refresh_token`. You
do this by making an HTTP POST request to the following address:

```
https://login.linode.com/oauth/token
```

Make this request as `application/x-www-form-urlencoded` or as
`multipart/form-data` and include the following parameters in the POST body:

| PARAMETER | DESCRIPTION |
|-----------|-------------|
| grant_type | The grant type you're using for renewal.  Currently only the string \"refresh_token\" is accepted. |
| client_id | Your app's client ID. |
| client_secret | Your app's client secret. |
| code | The code you just received from the redirect. |

You'll get a response like this:

```json
{
  \"scope\": \"linodes:read_write\",
  \"access_token\": \"03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c\"
  \"token_type\": \"bearer\",
  \"expires_in\": 7200,
}
```

Included in the reponse is an `access_token`. With this token, you can proceed to make
authenticated HTTP requests to the API by adding this header to each request:

```
Authorization: Bearer 03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c
```

#### OAuth Reference

| Security Scheme Type | OAuth 2.0 |
|-----------------------|--------|
| **Authorization URL** | https://login.linode.com/oauth/authorize |
| **Token URL** | https://login.linode.com/oauth/token |
| **Scopes** | <ul><li>`account:read_only` - Allows access to GET information about your Account.</li><li>`account:read_write` - Allows access to all endpoints related to your Account.</li><li>`domains:read_only` - Allows access to GET Domains on your Account.</li><li>`domains:read_write` - Allows access to all Domain endpoints.</li><li>`events:read_only` - Allows access to GET your Events.</li><li>`events:read_write` - Allows access to all endpoints related to your Events.</li><li>`firewall:read_only` - Allows access to GET information about your Firewalls.</li><li>`firewall:read_write` - Allows access to all Firewall endpoints.</li><li>`images:read_only` - Allows access to GET your Images.</li><li>`images:read_write` - Allows access to all endpoints related to your Images.</li><li>`ips:read_only` - Allows access to GET your ips.</li><li>`ips:read_write` - Allows access to all endpoints related to your ips.</li><li>`linodes:read_only` - Allows access to GET Linodes on your Account.</li><li>`linodes:read_write` - Allow access to all endpoints related to your Linodes.</li><li>`lke:read_only` - Allows access to GET LKE Clusters on your Account.</li><li>`lke:read_write` - Allows access to all endpoints related to LKE Clusters on your Account.</li><li>`longview:read_only` - Allows access to GET your Longview Clients.</li><li>`longview:read_write` - Allows access to all endpoints related to your Longview Clients.</li><li>`maintenance:read_only` - Allows access to GET information about Maintenance on your account.</li><li>`nodebalancers:read_only` - Allows access to GET NodeBalancers on your Account.</li><li>`nodebalancers:read_write` - Allows access to all NodeBalancer endpoints.</li><li>`object_storage:read_only` - Allows access to GET information related to your Object Storage.</li><li>`object_storage:read_write` - Allows access to all Object Storage endpoints.</li><li>`stackscripts:read_only` - Allows access to GET your StackScripts.</li><li>`stackscripts:read_write` - Allows access to all endpoints related to your StackScripts.</li><li>`volumes:read_only` - Allows access to GET your Volumes.</li><li>`volumes:read_write` - Allows access to all endpoints related to your Volumes.</li></ul><br/>|

## Requests

Requests must be made over HTTPS to ensure transactions are encrypted. The
following Request methods are supported:

| METHOD | USAGE |
|--------|-------|
| GET    | Retrieves data about collections and individual resources. |
| POST   | For collections, creates a new resource of that type. Also used to perform actions on action endpoints. |
| PUT    | Updates an existing resource. |
| DELETE | Deletes a resource. This is a destructive action. |


## Responses

Actions will return one following HTTP response status codes:

| STATUS  | DESCRIPTION |
|---------|-------------|
| 200 OK  | The request was successful. |
| 204 No Content | The server successfully fulfilled the request and there is no additional content to send. |
| 400 Bad Request | You submitted an invalid request (missing parameters, etc.). |
| 401 Unauthorized | You failed to authenticate for this resource. |
| 403 Forbidden | You are authenticated, but don't have permission to do this. |
| 404 Not Found | The resource you're requesting does not exist. |
| 429 Too Many Requests | You've hit a rate limit. |
| 500 Internal Server Error | Please [open a Support Ticket](/docs/api/support/#support-ticket-open). |

## Errors

Success is indicated via <a href=\"https://en.wikipedia.org/wiki/List_of_HTTP_status_codes\" target=\"_top\">Standard HTTP status codes</a>.
`2xx` codes indicate success, `4xx` codes indicate a request error, and
`5xx` errors indicate a server error. A
request error might be an invalid input, a required parameter being omitted,
or a malformed request. A server error means something went wrong processing
your request. If this occurs, please
[open a Support Ticket](/docs/api/support/#support-ticket-open)
and let us know. Though errors are logged and we work quickly to resolve issues,
opening a ticket and providing us with reproducable steps and data is always helpful.

The `errors` field is an array of the things that went wrong with your request.
We will try to include as many of the problems in the response as possible,
but it's conceivable that fixing these errors and resubmitting may result in
new errors coming back once we are able to get further along in the process
of handling your request.


Within each error object, the `field` parameter will be included if the error
pertains to a specific field in the JSON you've submitted. This will be
omitted if there is no relevant field. The `reason` is a human-readable
explanation of the error, and will always be included.

## Pagination

Resource lists are always paginated. The response will look similar to this:

```json
{
    \"data\": [ ... ],
    \"page\": 1,
    \"pages\": 3,
    \"results\": 300
}
```

* Pages start at 1. You may retrieve a specific page of results by adding
`?page=x` to your URL (for example, `?page=4`). If the value of `page`
exceeds `2^64/page_size`, the last possible page will be returned.


* Page sizes default to 100,
and can be set to return between 25 and 500. Page size can be set using
`?page_size=x`.

## Filtering and Sorting

Collections are searchable by fields they include, marked in the spec as
`x-linode-filterable: true`. Filters are passed
in the `X-Filter` header and are formatted as JSON objects. Here is a request
call for Linode Types in our \"standard\" class:

```Shell
curl \"https://api.linode.com/v4/linode/types\" \\
  -H '
    X-Filter: {
      \"class\": \"standard\"
    }'
```

The filter object's keys are the keys of the object you're filtering,
and the values are accepted values. You can add multiple filters by
including more than one key. For example, filtering for \"standard\" Linode
Types that offer one vcpu:

```Shell
 curl \"https://api.linode.com/v4/linode/types\" \\
  -H '
    X-Filter: {
      \"class\": \"standard\",
      \"vcpus\": 1
    }'
```

In the above example, both filters are combined with an \"and\" operation.
However, if you wanted either Types with one vcpu or Types in our \"standard\"
class, you can add an operator:

 ```Shell
curl \"https://api.linode.com/v4/linode/types\" \\
  -H '
    X-Filter: {
      \"+or\": [
        { \"vcpus\": 1 },
        { \"class\": \"standard\" }
      ]
    }'
```

Each filter in the `+or` array is its own filter object, and all conditions
in it are combined with an \"and\" operation as they were in the previous example.

Other operators are also available. Operators are keys of a Filter JSON
object. Their value must be of the appropriate type, and they are evaluated
as described below:

| OPERATOR | TYPE   | DESCRIPTION                       |
|----------|--------|-----------------------------------|
| +and     | array  | All conditions must be true.       |
| +or      | array  | One condition must be true.        |
| +gt      | number | Value must be greater than number. |
| +gte     | number | Value must be greater than or equal to number. |
| +lt      | number | Value must be less than number. |
| +lte     | number | Value must be less than or equal to number. |
| +contains | string | Given string must be in the value. |
| +neq      | string | Does not equal the value.          |
| +order_by | string | Attribute to order the results by - must be filterable. |
| +order    | string | Either \"asc\" or \"desc\". Defaults to \"asc\". Requires `+order_by`. |

For example, filtering for [Linode Types](/docs/api/linode-types/)
that offer memory equal to or higher than 61440:

```Shell
curl \"https://api.linode.com/v4/linode/types\" \\
  -H '
    X-Filter: {
      \"memory\": {
        \"+gte\": 61440
      }
    }'
```

You can combine and nest operators to construct arbitrarily-complex queries.
For example, give me all [Linode Types](/docs/api/linode-types/)
which are either `standard` or `highmem` class, or
have between 12 and 20 vcpus:

```Shell
curl \"https://api.linode.com/v4/linode/types\" \\
  -H '
    X-Filter: {
      \"+or\": [
        {
          \"+or\": [
            {
              \"class\": \"standard\"
            },
            {
              \"class\": \"highmem\"
            }
          ]
        },
        {
          \"+and\": [
            {
              \"vcpus\": {
                \"+gte\": 12
              }
            },
            {
              \"vcpus\": {
                \"+lte\": 20
              }
            }
          ]
        }
      ]
    }'
```

## Rate Limiting

With the Linode API, you can make up to 1,600 general API requests every two minutes per user as
determined by IP adddress or by OAuth token. Additionally, there are endpoint specfic limits defined below.

**Note:** There may be rate limiting applied at other levels outside of the API, for example, at the load balancer.

`/stats` endpoints have their own dedicated limits of 100 requests per minute per user.
These endpoints are:

* [View Linode Statistics](/docs/api/linode-instances/#linode-statistics-view)
* [View Linode Statistics (year/month)](/docs/api/linode-instances/#statistics-yearmonth-view)
* [View NodeBalancer Statistics](/docs/api/nodebalancers/#nodebalancer-statistics-view)
* [List Managed Stats](/docs/api/managed/#managed-stats-list)

Object Storage endpoints have a dedicated limit of 750 requests per second per user.
The Object Storage endpoints are:

* [Object Storage Endpoints](/docs/api/object-storage/)

Opening Support Tickets has a dedicated limit of 2 requests per minute per user.
That endpoint is:

* [Open Support Ticket](/docs/api/support/#support-ticket-open)

Accepting Entity Transfers has a dedicated limit of 2 requests per minute per user.
That endpoint is:

* [Entity Transfer Accept](/docs/api/account/#entity-transfer-accept)

## CLI (Command Line Interface)

The <a href=\"https://github.com/linode/linode-cli\" target=\"_top\">Linode CLI</a> allows you to easily
work with the API using intuitive and simple syntax. It requires a
[Personal Access Token](/docs/api/#personal-access-token)
for authentication, and gives you access to all of the features and functionality
of the Linode API that are documented here with CLI examples.

Endpoints that do not have CLI examples are currently unavailable through the CLI, but
can be accessed via other methods such as Shell commands and other third-party applications.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.85.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.linode.com/support/](https://www.linode.com/support/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v1"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.linode.com/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**AcceptEntityTransfer**](docs/AccountApi.md#acceptentitytransfer) | **Post** /account/entity-transfers/{token}/accept | Entity Transfer Accept
*AccountApi* | [**CancelAccount**](docs/AccountApi.md#cancelaccount) | **Post** /account/cancel | Account Cancel
*AccountApi* | [**CreateClient**](docs/AccountApi.md#createclient) | **Post** /account/oauth-clients | OAuth Client Create
*AccountApi* | [**CreateCreditCard**](docs/AccountApi.md#createcreditcard) | **Post** /account/credit-card | Credit Card Add/Edit
*AccountApi* | [**CreateEntityTransfer**](docs/AccountApi.md#createentitytransfer) | **Post** /account/entity-transfers | Entity Transfer Create
*AccountApi* | [**CreatePayPalPayment**](docs/AccountApi.md#createpaypalpayment) | **Post** /account/payments/paypal | PayPal Payment Stage
*AccountApi* | [**CreatePayment**](docs/AccountApi.md#createpayment) | **Post** /account/payments | Payment Make
*AccountApi* | [**CreateUser**](docs/AccountApi.md#createuser) | **Post** /account/users | User Create
*AccountApi* | [**DeleteClient**](docs/AccountApi.md#deleteclient) | **Delete** /account/oauth-clients/{clientId} | OAuth Client Delete
*AccountApi* | [**DeleteEntityTransfer**](docs/AccountApi.md#deleteentitytransfer) | **Delete** /account/entity-transfers/{token} | Entity Transfer Cancel
*AccountApi* | [**DeleteUser**](docs/AccountApi.md#deleteuser) | **Delete** /account/users/{username} | User Delete
*AccountApi* | [**EnableAccountManged**](docs/AccountApi.md#enableaccountmanged) | **Post** /account/settings/managed-enable | Linode Managed Enable
*AccountApi* | [**EventRead**](docs/AccountApi.md#eventread) | **Post** /account/events/{eventId}/read | Event Mark as Read
*AccountApi* | [**EventSeen**](docs/AccountApi.md#eventseen) | **Post** /account/events/{eventId}/seen | Event Mark as Seen
*AccountApi* | [**ExecutePayPalPayment**](docs/AccountApi.md#executepaypalpayment) | **Post** /account/payments/paypal/execute | Staged/Approved PayPal Payment Execute
*AccountApi* | [**GetAccount**](docs/AccountApi.md#getaccount) | **Get** /account | Account View
*AccountApi* | [**GetAccountLogin**](docs/AccountApi.md#getaccountlogin) | **Get** /account/logins/{loginId} | Login View
*AccountApi* | [**GetAccountLogins**](docs/AccountApi.md#getaccountlogins) | **Get** /account/logins | User Logins List All
*AccountApi* | [**GetAccountSettings**](docs/AccountApi.md#getaccountsettings) | **Get** /account/settings | Account Settings View
*AccountApi* | [**GetClient**](docs/AccountApi.md#getclient) | **Get** /account/oauth-clients/{clientId} | OAuth Client View
*AccountApi* | [**GetClientThumbnail**](docs/AccountApi.md#getclientthumbnail) | **Get** /account/oauth-clients/{clientId}/thumbnail | OAuth Client Thumbnail View
*AccountApi* | [**GetClients**](docs/AccountApi.md#getclients) | **Get** /account/oauth-clients | OAuth Clients List
*AccountApi* | [**GetEntityTransfer**](docs/AccountApi.md#getentitytransfer) | **Get** /account/entity-transfers/{token} | Entity Transfer View
*AccountApi* | [**GetEntityTransfers**](docs/AccountApi.md#getentitytransfers) | **Get** /account/entity-transfers | Entity Transfers List
*AccountApi* | [**GetEvent**](docs/AccountApi.md#getevent) | **Get** /account/events/{eventId} | Event View
*AccountApi* | [**GetEvents**](docs/AccountApi.md#getevents) | **Get** /account/events | Events List
*AccountApi* | [**GetInvoice**](docs/AccountApi.md#getinvoice) | **Get** /account/invoices/{invoiceId} | Invoice View
*AccountApi* | [**GetInvoiceItems**](docs/AccountApi.md#getinvoiceitems) | **Get** /account/invoices/{invoiceId}/items | Invoice Items List
*AccountApi* | [**GetInvoices**](docs/AccountApi.md#getinvoices) | **Get** /account/invoices | Invoices List
*AccountApi* | [**GetMaintenance**](docs/AccountApi.md#getmaintenance) | **Get** /account/maintenance | Maintenance List
*AccountApi* | [**GetNotifications**](docs/AccountApi.md#getnotifications) | **Get** /account/notifications | Notifications List
*AccountApi* | [**GetPayment**](docs/AccountApi.md#getpayment) | **Get** /account/payments/{paymentId} | Payment View
*AccountApi* | [**GetPayments**](docs/AccountApi.md#getpayments) | **Get** /account/payments | Payments List
*AccountApi* | [**GetTransfer**](docs/AccountApi.md#gettransfer) | **Get** /account/transfer | Network Utilization View
*AccountApi* | [**GetUser**](docs/AccountApi.md#getuser) | **Get** /account/users/{username} | User View
*AccountApi* | [**GetUserGrants**](docs/AccountApi.md#getusergrants) | **Get** /account/users/{username}/grants | User&#39;s Grants View
*AccountApi* | [**GetUsers**](docs/AccountApi.md#getusers) | **Get** /account/users | Users List
*AccountApi* | [**ResetClientSecret**](docs/AccountApi.md#resetclientsecret) | **Post** /account/oauth-clients/{clientId}/reset-secret | OAuth Client Secret Reset
*AccountApi* | [**SetClientThumbnail**](docs/AccountApi.md#setclientthumbnail) | **Put** /account/oauth-clients/{clientId}/thumbnail | OAuth Client Thumbnail Update
*AccountApi* | [**UpdateAccount**](docs/AccountApi.md#updateaccount) | **Put** /account | Account Update
*AccountApi* | [**UpdateAccountSettings**](docs/AccountApi.md#updateaccountsettings) | **Put** /account/settings | Account Settings Update
*AccountApi* | [**UpdateClient**](docs/AccountApi.md#updateclient) | **Put** /account/oauth-clients/{clientId} | OAuth Client Update
*AccountApi* | [**UpdateUser**](docs/AccountApi.md#updateuser) | **Put** /account/users/{username} | User Update
*AccountApi* | [**UpdateUserGrants**](docs/AccountApi.md#updateusergrants) | **Put** /account/users/{username}/grants | User&#39;s Grants Update
*DomainsApi* | [**CloneDomain**](docs/DomainsApi.md#clonedomain) | **Post** /domains/{domainId}/clone | Domain Clone
*DomainsApi* | [**CreateDomain**](docs/DomainsApi.md#createdomain) | **Post** /domains | Domain Create
*DomainsApi* | [**CreateDomainRecord**](docs/DomainsApi.md#createdomainrecord) | **Post** /domains/{domainId}/records | Domain Record Create
*DomainsApi* | [**DeleteDomain**](docs/DomainsApi.md#deletedomain) | **Delete** /domains/{domainId} | Domain Delete
*DomainsApi* | [**DeleteDomainRecord**](docs/DomainsApi.md#deletedomainrecord) | **Delete** /domains/{domainId}/records/{recordId} | Domain Record Delete
*DomainsApi* | [**GetDomain**](docs/DomainsApi.md#getdomain) | **Get** /domains/{domainId} | Domain View
*DomainsApi* | [**GetDomainRecord**](docs/DomainsApi.md#getdomainrecord) | **Get** /domains/{domainId}/records/{recordId} | Domain Record View
*DomainsApi* | [**GetDomainRecords**](docs/DomainsApi.md#getdomainrecords) | **Get** /domains/{domainId}/records | Domain Records List
*DomainsApi* | [**GetDomains**](docs/DomainsApi.md#getdomains) | **Get** /domains | Domains List
*DomainsApi* | [**ImportDomain**](docs/DomainsApi.md#importdomain) | **Post** /domains/import | Domain Import
*DomainsApi* | [**UpdateDomain**](docs/DomainsApi.md#updatedomain) | **Put** /domains/{domainId} | Domain Update
*DomainsApi* | [**UpdateDomainRecord**](docs/DomainsApi.md#updatedomainrecord) | **Put** /domains/{domainId}/records/{recordId} | Domain Record Update
*ImagesApi* | [**CreateImage**](docs/ImagesApi.md#createimage) | **Post** /images | Image Create
*ImagesApi* | [**DeleteImage**](docs/ImagesApi.md#deleteimage) | **Delete** /images/{imageId} | Image Delete
*ImagesApi* | [**GetImage**](docs/ImagesApi.md#getimage) | **Get** /images/{imageId} | Image View
*ImagesApi* | [**GetImages**](docs/ImagesApi.md#getimages) | **Get** /images | Images List
*ImagesApi* | [**UpdateImage**](docs/ImagesApi.md#updateimage) | **Put** /images/{imageId} | Image Update
*LinodeInstancesApi* | [**AddLinodeConfig**](docs/LinodeInstancesApi.md#addlinodeconfig) | **Post** /linode/instances/{linodeId}/configs | Configuration Profile Create
*LinodeInstancesApi* | [**AddLinodeDisk**](docs/LinodeInstancesApi.md#addlinodedisk) | **Post** /linode/instances/{linodeId}/disks | Disk Create
*LinodeInstancesApi* | [**AddLinodeIP**](docs/LinodeInstancesApi.md#addlinodeip) | **Post** /linode/instances/{linodeId}/ips | IPv4 Address Allocate
*LinodeInstancesApi* | [**BootLinodeInstance**](docs/LinodeInstancesApi.md#bootlinodeinstance) | **Post** /linode/instances/{linodeId}/boot | Linode Boot
*LinodeInstancesApi* | [**CancelBackups**](docs/LinodeInstancesApi.md#cancelbackups) | **Post** /linode/instances/{linodeId}/backups/cancel | Backups Cancel
*LinodeInstancesApi* | [**CloneLinodeDisk**](docs/LinodeInstancesApi.md#clonelinodedisk) | **Post** /linode/instances/{linodeId}/disks/{diskId}/clone | Disk Clone
*LinodeInstancesApi* | [**CloneLinodeInstance**](docs/LinodeInstancesApi.md#clonelinodeinstance) | **Post** /linode/instances/{linodeId}/clone | Linode Clone
*LinodeInstancesApi* | [**CreateLinodeInstance**](docs/LinodeInstancesApi.md#createlinodeinstance) | **Post** /linode/instances | Linode Create
*LinodeInstancesApi* | [**CreateSnapshot**](docs/LinodeInstancesApi.md#createsnapshot) | **Post** /linode/instances/{linodeId}/backups | Snapshot Create
*LinodeInstancesApi* | [**DeleteDisk**](docs/LinodeInstancesApi.md#deletedisk) | **Delete** /linode/instances/{linodeId}/disks/{diskId} | Disk Delete
*LinodeInstancesApi* | [**DeleteLinodeConfig**](docs/LinodeInstancesApi.md#deletelinodeconfig) | **Delete** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile Delete
*LinodeInstancesApi* | [**DeleteLinodeInstance**](docs/LinodeInstancesApi.md#deletelinodeinstance) | **Delete** /linode/instances/{linodeId} | Linode Delete
*LinodeInstancesApi* | [**EnableBackups**](docs/LinodeInstancesApi.md#enablebackups) | **Post** /linode/instances/{linodeId}/backups/enable | Backups Enable
*LinodeInstancesApi* | [**GetBackup**](docs/LinodeInstancesApi.md#getbackup) | **Get** /linode/instances/{linodeId}/backups/{backupId} | Backup View
*LinodeInstancesApi* | [**GetBackups**](docs/LinodeInstancesApi.md#getbackups) | **Get** /linode/instances/{linodeId}/backups | Backups List
*LinodeInstancesApi* | [**GetKernel**](docs/LinodeInstancesApi.md#getkernel) | **Get** /linode/kernels/{kernelId} | Kernel View
*LinodeInstancesApi* | [**GetKernels**](docs/LinodeInstancesApi.md#getkernels) | **Get** /linode/kernels | Kernels List
*LinodeInstancesApi* | [**GetLinodeConfig**](docs/LinodeInstancesApi.md#getlinodeconfig) | **Get** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile View
*LinodeInstancesApi* | [**GetLinodeConfigs**](docs/LinodeInstancesApi.md#getlinodeconfigs) | **Get** /linode/instances/{linodeId}/configs | Configuration Profiles List
*LinodeInstancesApi* | [**GetLinodeDisk**](docs/LinodeInstancesApi.md#getlinodedisk) | **Get** /linode/instances/{linodeId}/disks/{diskId} | Disk View
*LinodeInstancesApi* | [**GetLinodeDisks**](docs/LinodeInstancesApi.md#getlinodedisks) | **Get** /linode/instances/{linodeId}/disks | Disks List
*LinodeInstancesApi* | [**GetLinodeFirewalls**](docs/LinodeInstancesApi.md#getlinodefirewalls) | **Get** /linode/instances/{linodeId}/firewalls | Firewalls List
*LinodeInstancesApi* | [**GetLinodeIP**](docs/LinodeInstancesApi.md#getlinodeip) | **Get** /linode/instances/{linodeId}/ips/{address} | IP Address View
*LinodeInstancesApi* | [**GetLinodeIPs**](docs/LinodeInstancesApi.md#getlinodeips) | **Get** /linode/instances/{linodeId}/ips | Networking Information List
*LinodeInstancesApi* | [**GetLinodeInstance**](docs/LinodeInstancesApi.md#getlinodeinstance) | **Get** /linode/instances/{linodeId} | Linode View
*LinodeInstancesApi* | [**GetLinodeInstances**](docs/LinodeInstancesApi.md#getlinodeinstances) | **Get** /linode/instances | Linodes List
*LinodeInstancesApi* | [**GetLinodeStats**](docs/LinodeInstancesApi.md#getlinodestats) | **Get** /linode/instances/{linodeId}/stats | Linode Statistics View
*LinodeInstancesApi* | [**GetLinodeStatsByYearMonth**](docs/LinodeInstancesApi.md#getlinodestatsbyyearmonth) | **Get** /linode/instances/{linodeId}/stats/{year}/{month} | Statistics View (year/month)
*LinodeInstancesApi* | [**GetLinodeTransfer**](docs/LinodeInstancesApi.md#getlinodetransfer) | **Get** /linode/instances/{linodeId}/transfer | Network Transfer View
*LinodeInstancesApi* | [**GetLinodeTransferByYearMonth**](docs/LinodeInstancesApi.md#getlinodetransferbyyearmonth) | **Get** /linode/instances/{linodeId}/transfer/{year}/{month} | Network Transfer View (year/month)
*LinodeInstancesApi* | [**GetLinodeVolumes**](docs/LinodeInstancesApi.md#getlinodevolumes) | **Get** /linode/instances/{linodeId}/volumes | Linode&#39;s Volumes List
*LinodeInstancesApi* | [**MigrateLinodeInstance**](docs/LinodeInstancesApi.md#migratelinodeinstance) | **Post** /linode/instances/{linodeId}/migrate | DC Migration/Pending Host Migration Initiate
*LinodeInstancesApi* | [**MutateLinodeInstance**](docs/LinodeInstancesApi.md#mutatelinodeinstance) | **Post** /linode/instances/{linodeId}/mutate | Linode Upgrade
*LinodeInstancesApi* | [**RebootLinodeInstance**](docs/LinodeInstancesApi.md#rebootlinodeinstance) | **Post** /linode/instances/{linodeId}/reboot | Linode Reboot
*LinodeInstancesApi* | [**RebuildLinodeInstance**](docs/LinodeInstancesApi.md#rebuildlinodeinstance) | **Post** /linode/instances/{linodeId}/rebuild | Linode Rebuild
*LinodeInstancesApi* | [**RemoveLinodeIP**](docs/LinodeInstancesApi.md#removelinodeip) | **Delete** /linode/instances/{linodeId}/ips/{address} | IPv4 Address Delete
*LinodeInstancesApi* | [**RescueLinodeInstance**](docs/LinodeInstancesApi.md#rescuelinodeinstance) | **Post** /linode/instances/{linodeId}/rescue | Linode Boot into Rescue Mode
*LinodeInstancesApi* | [**ResetDiskPassword**](docs/LinodeInstancesApi.md#resetdiskpassword) | **Post** /linode/instances/{linodeId}/disks/{diskId}/password | Disk Root Password Reset
*LinodeInstancesApi* | [**ResetLinodePassword**](docs/LinodeInstancesApi.md#resetlinodepassword) | **Post** /linode/instances/{linodeId}/password | Linode Root Password Reset
*LinodeInstancesApi* | [**ResizeDisk**](docs/LinodeInstancesApi.md#resizedisk) | **Post** /linode/instances/{linodeId}/disks/{diskId}/resize | Disk Resize
*LinodeInstancesApi* | [**ResizeLinodeInstance**](docs/LinodeInstancesApi.md#resizelinodeinstance) | **Post** /linode/instances/{linodeId}/resize | Linode Resize
*LinodeInstancesApi* | [**RestoreBackup**](docs/LinodeInstancesApi.md#restorebackup) | **Post** /linode/instances/{linodeId}/backups/{backupId}/restore | Backup Restore
*LinodeInstancesApi* | [**ShutdownLinodeInstance**](docs/LinodeInstancesApi.md#shutdownlinodeinstance) | **Post** /linode/instances/{linodeId}/shutdown | Linode Shut Down
*LinodeInstancesApi* | [**UpdateDisk**](docs/LinodeInstancesApi.md#updatedisk) | **Put** /linode/instances/{linodeId}/disks/{diskId} | Disk Update
*LinodeInstancesApi* | [**UpdateLinodeConfig**](docs/LinodeInstancesApi.md#updatelinodeconfig) | **Put** /linode/instances/{linodeId}/configs/{configId} | Configuration Profile Update
*LinodeInstancesApi* | [**UpdateLinodeIP**](docs/LinodeInstancesApi.md#updatelinodeip) | **Put** /linode/instances/{linodeId}/ips/{address} | IP Address Update
*LinodeInstancesApi* | [**UpdateLinodeInstance**](docs/LinodeInstancesApi.md#updatelinodeinstance) | **Put** /linode/instances/{linodeId} | Linode Update
*LinodeKubernetesEngineLKEApi* | [**CreateLKECluster**](docs/LinodeKubernetesEngineLKEApi.md#createlkecluster) | **Post** /lke/clusters | Kubernetes Cluster Create
*LinodeKubernetesEngineLKEApi* | [**DeleteLKECluster**](docs/LinodeKubernetesEngineLKEApi.md#deletelkecluster) | **Delete** /lke/clusters/{clusterId} | Kubernetes Cluster Delete
*LinodeKubernetesEngineLKEApi* | [**DeleteLKENodePool**](docs/LinodeKubernetesEngineLKEApi.md#deletelkenodepool) | **Delete** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool Delete
*LinodeKubernetesEngineLKEApi* | [**GetLKECluster**](docs/LinodeKubernetesEngineLKEApi.md#getlkecluster) | **Get** /lke/clusters/{clusterId} | Kubernetes Cluster View
*LinodeKubernetesEngineLKEApi* | [**GetLKEClusterAPIEndpoints**](docs/LinodeKubernetesEngineLKEApi.md#getlkeclusterapiendpoints) | **Get** /lke/clusters/{clusterId}/api-endpoints | Kubernetes API Endpoints List
*LinodeKubernetesEngineLKEApi* | [**GetLKEClusterKubeconfig**](docs/LinodeKubernetesEngineLKEApi.md#getlkeclusterkubeconfig) | **Get** /lke/clusters/{clusterId}/kubeconfig | Kubeconfig View
*LinodeKubernetesEngineLKEApi* | [**GetLKEClusterNode**](docs/LinodeKubernetesEngineLKEApi.md#getlkeclusternode) | **Get** /lke/clusters/{clusterId}/nodes/{nodeId} | Node View
*LinodeKubernetesEngineLKEApi* | [**GetLKEClusterPools**](docs/LinodeKubernetesEngineLKEApi.md#getlkeclusterpools) | **Get** /lke/clusters/{clusterId}/pools | Node Pools List
*LinodeKubernetesEngineLKEApi* | [**GetLKEClusters**](docs/LinodeKubernetesEngineLKEApi.md#getlkeclusters) | **Get** /lke/clusters | Kubernetes Clusters List
*LinodeKubernetesEngineLKEApi* | [**GetLKENodePool**](docs/LinodeKubernetesEngineLKEApi.md#getlkenodepool) | **Get** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool View
*LinodeKubernetesEngineLKEApi* | [**GetLKEVersion**](docs/LinodeKubernetesEngineLKEApi.md#getlkeversion) | **Get** /lke/versions/{version} | Kubernetes Version View
*LinodeKubernetesEngineLKEApi* | [**GetLKEVersions**](docs/LinodeKubernetesEngineLKEApi.md#getlkeversions) | **Get** /lke/versions | Kubernetes Versions List
*LinodeKubernetesEngineLKEApi* | [**PostLKEClusterNodeRecycle**](docs/LinodeKubernetesEngineLKEApi.md#postlkeclusternoderecycle) | **Post** /lke/clusters/{clusterId}/nodes/{nodeId}/recycle | Node Recycle
*LinodeKubernetesEngineLKEApi* | [**PostLKEClusterPoolRecycle**](docs/LinodeKubernetesEngineLKEApi.md#postlkeclusterpoolrecycle) | **Post** /lke/clusters/{clusterId}/pools/{poolId}/recycle | Node Pool Recycle
*LinodeKubernetesEngineLKEApi* | [**PostLKEClusterPools**](docs/LinodeKubernetesEngineLKEApi.md#postlkeclusterpools) | **Post** /lke/clusters/{clusterId}/pools | Node Pool Create
*LinodeKubernetesEngineLKEApi* | [**PostLKEClusterRecycle**](docs/LinodeKubernetesEngineLKEApi.md#postlkeclusterrecycle) | **Post** /lke/clusters/{clusterId}/recycle | Cluster Nodes Recycle
*LinodeKubernetesEngineLKEApi* | [**PutLKECluster**](docs/LinodeKubernetesEngineLKEApi.md#putlkecluster) | **Put** /lke/clusters/{clusterId} | Kubernetes Cluster Update
*LinodeKubernetesEngineLKEApi* | [**PutLKENodePool**](docs/LinodeKubernetesEngineLKEApi.md#putlkenodepool) | **Put** /lke/clusters/{clusterId}/pools/{poolId} | Node Pool Update
*LinodeTypesApi* | [**GetLinodeType**](docs/LinodeTypesApi.md#getlinodetype) | **Get** /linode/types/{typeId} | Type View
*LinodeTypesApi* | [**GetLinodeTypes**](docs/LinodeTypesApi.md#getlinodetypes) | **Get** /linode/types | Types List
*LongviewApi* | [**CreateLongviewClient**](docs/LongviewApi.md#createlongviewclient) | **Post** /longview/clients | Longview Client Create
*LongviewApi* | [**DeleteLongviewClient**](docs/LongviewApi.md#deletelongviewclient) | **Delete** /longview/clients/{clientId} | Longview Client Delete
*LongviewApi* | [**GetLongviewClient**](docs/LongviewApi.md#getlongviewclient) | **Get** /longview/clients/{clientId} | Longview Client View
*LongviewApi* | [**GetLongviewClients**](docs/LongviewApi.md#getlongviewclients) | **Get** /longview/clients | Longview Clients List
*LongviewApi* | [**GetLongviewPlan**](docs/LongviewApi.md#getlongviewplan) | **Get** /longview/plan | Longview Plan View
*LongviewApi* | [**GetLongviewSubscription**](docs/LongviewApi.md#getlongviewsubscription) | **Get** /longview/subscriptions/{subscriptionId} | Longview Subscription View
*LongviewApi* | [**GetLongviewSubscriptions**](docs/LongviewApi.md#getlongviewsubscriptions) | **Get** /longview/subscriptions | Longview Subscriptions List
*LongviewApi* | [**UpdateLongviewClient**](docs/LongviewApi.md#updatelongviewclient) | **Put** /longview/clients/{clientId} | Longview Client Update
*LongviewApi* | [**UpdateLongviewPlan**](docs/LongviewApi.md#updatelongviewplan) | **Put** /longview/plan | Longview Plan Update
*ManagedApi* | [**CreateManagedContact**](docs/ManagedApi.md#createmanagedcontact) | **Post** /managed/contacts | Managed Contact Create
*ManagedApi* | [**CreateManagedCredential**](docs/ManagedApi.md#createmanagedcredential) | **Post** /managed/credentials | Managed Credential Create
*ManagedApi* | [**CreateManagedService**](docs/ManagedApi.md#createmanagedservice) | **Post** /managed/services | Managed Service Create
*ManagedApi* | [**DeleteManagedContact**](docs/ManagedApi.md#deletemanagedcontact) | **Delete** /managed/contacts/{contactId} | Managed Contact Delete
*ManagedApi* | [**DeleteManagedCredential**](docs/ManagedApi.md#deletemanagedcredential) | **Post** /managed/credentials/{credentialId}/revoke | Managed Credential Delete
*ManagedApi* | [**DeleteManagedService**](docs/ManagedApi.md#deletemanagedservice) | **Delete** /managed/services/{serviceId} | Managed Service Delete
*ManagedApi* | [**DisableManagedService**](docs/ManagedApi.md#disablemanagedservice) | **Post** /managed/services/{serviceId}/disable | Managed Service Disable
*ManagedApi* | [**EnableManagedService**](docs/ManagedApi.md#enablemanagedservice) | **Post** /managed/services/{serviceId}/enable | Managed Service Enable
*ManagedApi* | [**GetManagedContact**](docs/ManagedApi.md#getmanagedcontact) | **Get** /managed/contacts/{contactId} | Managed Contact View
*ManagedApi* | [**GetManagedContacts**](docs/ManagedApi.md#getmanagedcontacts) | **Get** /managed/contacts | Managed Contacts List
*ManagedApi* | [**GetManagedCredential**](docs/ManagedApi.md#getmanagedcredential) | **Get** /managed/credentials/{credentialId} | Managed Credential View
*ManagedApi* | [**GetManagedCredentials**](docs/ManagedApi.md#getmanagedcredentials) | **Get** /managed/credentials | Managed Credentials List
*ManagedApi* | [**GetManagedIssue**](docs/ManagedApi.md#getmanagedissue) | **Get** /managed/issues/{issueId} | Managed Issue View
*ManagedApi* | [**GetManagedIssues**](docs/ManagedApi.md#getmanagedissues) | **Get** /managed/issues | Managed Issues List
*ManagedApi* | [**GetManagedLinodeSetting**](docs/ManagedApi.md#getmanagedlinodesetting) | **Get** /managed/linode-settings/{linodeId} | Linode&#39;s Managed Settings View
*ManagedApi* | [**GetManagedLinodeSettings**](docs/ManagedApi.md#getmanagedlinodesettings) | **Get** /managed/linode-settings | Managed Linode Settings List
*ManagedApi* | [**GetManagedService**](docs/ManagedApi.md#getmanagedservice) | **Get** /managed/services/{serviceId} | Managed Service View
*ManagedApi* | [**GetManagedServices**](docs/ManagedApi.md#getmanagedservices) | **Get** /managed/services | Managed Services List
*ManagedApi* | [**GetManagedStats**](docs/ManagedApi.md#getmanagedstats) | **Get** /managed/stats | Managed Stats List
*ManagedApi* | [**UpdateManagedContact**](docs/ManagedApi.md#updatemanagedcontact) | **Put** /managed/contacts/{contactId} | Managed Contact Update
*ManagedApi* | [**UpdateManagedCredential**](docs/ManagedApi.md#updatemanagedcredential) | **Put** /managed/credentials/{credentialId} | Managed Credential Update
*ManagedApi* | [**UpdateManagedCredentialUsernamePassword**](docs/ManagedApi.md#updatemanagedcredentialusernamepassword) | **Post** /managed/credentials/{credentialId}/update | Managed Credential Username and Password Update
*ManagedApi* | [**UpdateManagedLinodeSetting**](docs/ManagedApi.md#updatemanagedlinodesetting) | **Put** /managed/linode-settings/{linodeId} | Linode&#39;s Managed Settings Update
*ManagedApi* | [**UpdateManagedService**](docs/ManagedApi.md#updatemanagedservice) | **Put** /managed/services/{serviceId} | Managed Service Update
*ManagedApi* | [**ViewManagedSSHKey**](docs/ManagedApi.md#viewmanagedsshkey) | **Get** /managed/credentials/sshkey | Managed SSH Key View
*NetworkingApi* | [**AllocateIP**](docs/NetworkingApi.md#allocateip) | **Post** /networking/ips | IP Address Allocate
*NetworkingApi* | [**AssignIPs**](docs/NetworkingApi.md#assignips) | **Post** /networking/ipv4/assign | Linodes Assign IPs
*NetworkingApi* | [**CreateFirewallDevice**](docs/NetworkingApi.md#createfirewalldevice) | **Post** /networking/firewalls/{firewallId}/devices | Firewall Device Create
*NetworkingApi* | [**CreateFirewalls**](docs/NetworkingApi.md#createfirewalls) | **Post** /networking/firewalls | Firewall Create
*NetworkingApi* | [**DeleteFirewall**](docs/NetworkingApi.md#deletefirewall) | **Delete** /networking/firewalls/{firewallId} | Firewall Delete
*NetworkingApi* | [**DeleteFirewallDevice**](docs/NetworkingApi.md#deletefirewalldevice) | **Delete** /networking/firewalls/{firewallId}/devices/{deviceId} | Firewall Device Delete
*NetworkingApi* | [**GetFirewall**](docs/NetworkingApi.md#getfirewall) | **Get** /networking/firewalls/{firewallId} | Firewall View
*NetworkingApi* | [**GetFirewallDevice**](docs/NetworkingApi.md#getfirewalldevice) | **Get** /networking/firewalls/{firewallId}/devices/{deviceId} | Firewall Device View
*NetworkingApi* | [**GetFirewallDevices**](docs/NetworkingApi.md#getfirewalldevices) | **Get** /networking/firewalls/{firewallId}/devices | Firewall Devices List
*NetworkingApi* | [**GetFirewallRules**](docs/NetworkingApi.md#getfirewallrules) | **Get** /networking/firewalls/{firewallId}/rules | Firewall Rules List
*NetworkingApi* | [**GetFirewalls**](docs/NetworkingApi.md#getfirewalls) | **Get** /networking/firewalls | Firewalls List
*NetworkingApi* | [**GetIP**](docs/NetworkingApi.md#getip) | **Get** /networking/ips/{address} | IP Address View
*NetworkingApi* | [**GetIPs**](docs/NetworkingApi.md#getips) | **Get** /networking/ips | IP Addresses List
*NetworkingApi* | [**GetIPv6Pools**](docs/NetworkingApi.md#getipv6pools) | **Get** /networking/ipv6/pools | IPv6 Pools List
*NetworkingApi* | [**GetIPv6Ranges**](docs/NetworkingApi.md#getipv6ranges) | **Get** /networking/ipv6/ranges | IPv6 Ranges List
*NetworkingApi* | [**ShareIPs**](docs/NetworkingApi.md#shareips) | **Post** /networking/ipv4/share | IP Sharing Configure
*NetworkingApi* | [**UpdateFirewall**](docs/NetworkingApi.md#updatefirewall) | **Put** /networking/firewalls/{firewallId} | Firewall Update
*NetworkingApi* | [**UpdateFirewallRules**](docs/NetworkingApi.md#updatefirewallrules) | **Put** /networking/firewalls/{firewallId}/rules | Firewall Rules Update
*NetworkingApi* | [**UpdateIP**](docs/NetworkingApi.md#updateip) | **Put** /networking/ips/{address} | IP Address RDNS Update
*NodeBalancersApi* | [**CreateNodeBalancer**](docs/NodeBalancersApi.md#createnodebalancer) | **Post** /nodebalancers | NodeBalancer Create
*NodeBalancersApi* | [**CreateNodeBalancerConfig**](docs/NodeBalancersApi.md#createnodebalancerconfig) | **Post** /nodebalancers/{nodeBalancerId}/configs | Config Create
*NodeBalancersApi* | [**CreateNodeBalancerNode**](docs/NodeBalancersApi.md#createnodebalancernode) | **Post** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | Node Create
*NodeBalancersApi* | [**DeleteNodeBalancer**](docs/NodeBalancersApi.md#deletenodebalancer) | **Delete** /nodebalancers/{nodeBalancerId} | NodeBalancer Delete
*NodeBalancersApi* | [**DeleteNodeBalancerConfig**](docs/NodeBalancersApi.md#deletenodebalancerconfig) | **Delete** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config Delete
*NodeBalancersApi* | [**DeleteNodeBalancerConfigNode**](docs/NodeBalancersApi.md#deletenodebalancerconfignode) | **Delete** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node Delete
*NodeBalancersApi* | [**GetNodeBalancer**](docs/NodeBalancersApi.md#getnodebalancer) | **Get** /nodebalancers/{nodeBalancerId} | NodeBalancer View
*NodeBalancersApi* | [**GetNodeBalancerConfig**](docs/NodeBalancersApi.md#getnodebalancerconfig) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config View
*NodeBalancersApi* | [**GetNodeBalancerConfigNodes**](docs/NodeBalancersApi.md#getnodebalancerconfignodes) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | Nodes List
*NodeBalancersApi* | [**GetNodeBalancerConfigs**](docs/NodeBalancersApi.md#getnodebalancerconfigs) | **Get** /nodebalancers/{nodeBalancerId}/configs | Configs List
*NodeBalancersApi* | [**GetNodeBalancerNode**](docs/NodeBalancersApi.md#getnodebalancernode) | **Get** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node View
*NodeBalancersApi* | [**GetNodeBalancers**](docs/NodeBalancersApi.md#getnodebalancers) | **Get** /nodebalancers | NodeBalancers List
*NodeBalancersApi* | [**NodebalancersNodeBalancerIdStatsGet**](docs/NodeBalancersApi.md#nodebalancersnodebalanceridstatsget) | **Get** /nodebalancers/{nodeBalancerId}/stats | NodeBalancer Statistics View
*NodeBalancersApi* | [**RebuildNodeBalancerConfig**](docs/NodeBalancersApi.md#rebuildnodebalancerconfig) | **Post** /nodebalancers/{nodeBalancerId}/configs/{configId}/rebuild | Config Rebuild
*NodeBalancersApi* | [**UpdateNodeBalancer**](docs/NodeBalancersApi.md#updatenodebalancer) | **Put** /nodebalancers/{nodeBalancerId} | NodeBalancer Update
*NodeBalancersApi* | [**UpdateNodeBalancerConfig**](docs/NodeBalancersApi.md#updatenodebalancerconfig) | **Put** /nodebalancers/{nodeBalancerId}/configs/{configId} | Config Update
*NodeBalancersApi* | [**UpdateNodeBalancerNode**](docs/NodeBalancersApi.md#updatenodebalancernode) | **Put** /nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Node Update
*ObjectStorageApi* | [**CancelObjectStorage**](docs/ObjectStorageApi.md#cancelobjectstorage) | **Post** /object-storage/cancel | Object Storage Cancel
*ObjectStorageApi* | [**CreateObjectStorageBucket**](docs/ObjectStorageApi.md#createobjectstoragebucket) | **Post** /object-storage/buckets | Object Storage Bucket Create
*ObjectStorageApi* | [**CreateObjectStorageKeys**](docs/ObjectStorageApi.md#createobjectstoragekeys) | **Post** /object-storage/keys | Object Storage Key Create
*ObjectStorageApi* | [**CreateObjectStorageObjectURL**](docs/ObjectStorageApi.md#createobjectstorageobjecturl) | **Post** /object-storage/buckets/{clusterId}/{bucket}/object-url | Object Storage Object URL Create
*ObjectStorageApi* | [**CreateObjectStorageSSL**](docs/ObjectStorageApi.md#createobjectstoragessl) | **Post** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert Upload
*ObjectStorageApi* | [**DeleteObjectStorageBucket**](docs/ObjectStorageApi.md#deleteobjectstoragebucket) | **Delete** /object-storage/buckets/{clusterId}/{bucket} | Object Storage Bucket Remove
*ObjectStorageApi* | [**DeleteObjectStorageKey**](docs/ObjectStorageApi.md#deleteobjectstoragekey) | **Delete** /object-storage/keys/{keyId} | Object Storage Key Revoke
*ObjectStorageApi* | [**DeleteObjectStorageSSL**](docs/ObjectStorageApi.md#deleteobjectstoragessl) | **Delete** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert Delete
*ObjectStorageApi* | [**GetObjectStorageBucket**](docs/ObjectStorageApi.md#getobjectstoragebucket) | **Get** /object-storage/buckets/{clusterId}/{bucket} | Object Storage Bucket View
*ObjectStorageApi* | [**GetObjectStorageBucketContent**](docs/ObjectStorageApi.md#getobjectstoragebucketcontent) | **Get** /object-storage/buckets/{clusterId}/{bucket}/object-list | Object Storage Bucket Contents List
*ObjectStorageApi* | [**GetObjectStorageBucketinCluster**](docs/ObjectStorageApi.md#getobjectstoragebucketincluster) | **Get** /object-storage/buckets/{clusterId} | Object Storage Buckets in Cluster List
*ObjectStorageApi* | [**GetObjectStorageBuckets**](docs/ObjectStorageApi.md#getobjectstoragebuckets) | **Get** /object-storage/buckets | Object Storage Buckets List
*ObjectStorageApi* | [**GetObjectStorageCluster**](docs/ObjectStorageApi.md#getobjectstoragecluster) | **Get** /object-storage/clusters/{clusterId} | Cluster View
*ObjectStorageApi* | [**GetObjectStorageClusters**](docs/ObjectStorageApi.md#getobjectstorageclusters) | **Get** /object-storage/clusters | Clusters List
*ObjectStorageApi* | [**GetObjectStorageKey**](docs/ObjectStorageApi.md#getobjectstoragekey) | **Get** /object-storage/keys/{keyId} | Object Storage Key View
*ObjectStorageApi* | [**GetObjectStorageKeys**](docs/ObjectStorageApi.md#getobjectstoragekeys) | **Get** /object-storage/keys | Object Storage Keys List
*ObjectStorageApi* | [**GetObjectStorageSSL**](docs/ObjectStorageApi.md#getobjectstoragessl) | **Get** /object-storage/buckets/{clusterId}/{bucket}/ssl | Object Storage TLS/SSL Cert View
*ObjectStorageApi* | [**GetObjectStorageTransfer**](docs/ObjectStorageApi.md#getobjectstoragetransfer) | **Get** /object-storage/transfer | Object Storage Transfer View
*ObjectStorageApi* | [**ModifyObjectStorageBucketAccess**](docs/ObjectStorageApi.md#modifyobjectstoragebucketaccess) | **Post** /object-storage/buckets/{clusterId}/{bucket}/access | Object Storage Bucket Access Modify
*ObjectStorageApi* | [**UpdateObjectStorageBucketACL**](docs/ObjectStorageApi.md#updateobjectstoragebucketacl) | **Put** /object-storage/buckets/{clusterId}/{bucket}/object-acl | Object Storage Object ACL Config Update
*ObjectStorageApi* | [**UpdateObjectStorageBucketAccess**](docs/ObjectStorageApi.md#updateobjectstoragebucketaccess) | **Put** /object-storage/buckets/{clusterId}/{bucket}/access | Object Storage Bucket Access Update
*ObjectStorageApi* | [**UpdateObjectStorageKey**](docs/ObjectStorageApi.md#updateobjectstoragekey) | **Put** /object-storage/keys/{keyId} | Object Storage Key Update
*ObjectStorageApi* | [**ViewObjectStorageBucketACL**](docs/ObjectStorageApi.md#viewobjectstoragebucketacl) | **Get** /object-storage/buckets/{clusterId}/{bucket}/object-acl | Object Storage Object ACL Config View
*ProfileApi* | [**AddSSHKey**](docs/ProfileApi.md#addsshkey) | **Post** /profile/sshkeys | SSH Key Add
*ProfileApi* | [**CreatePersonalAccessToken**](docs/ProfileApi.md#createpersonalaccesstoken) | **Post** /profile/tokens | Personal Access Token Create
*ProfileApi* | [**DeletePersonalAccessToken**](docs/ProfileApi.md#deletepersonalaccesstoken) | **Delete** /profile/tokens/{tokenId} | Personal Access Token Revoke
*ProfileApi* | [**DeleteProfileApp**](docs/ProfileApi.md#deleteprofileapp) | **Delete** /profile/apps/{appId} | App Access Revoke
*ProfileApi* | [**DeleteSSHKey**](docs/ProfileApi.md#deletesshkey) | **Delete** /profile/sshkeys/{sshKeyId} | SSH Key Delete
*ProfileApi* | [**GetDevices**](docs/ProfileApi.md#getdevices) | **Get** /profile/devices | Trusted Devices List
*ProfileApi* | [**GetPersonalAccessToken**](docs/ProfileApi.md#getpersonalaccesstoken) | **Get** /profile/tokens/{tokenId} | Personal Access Token View
*ProfileApi* | [**GetPersonalAccessTokens**](docs/ProfileApi.md#getpersonalaccesstokens) | **Get** /profile/tokens | Personal Access Tokens List
*ProfileApi* | [**GetProfile**](docs/ProfileApi.md#getprofile) | **Get** /profile | Profile View
*ProfileApi* | [**GetProfileApp**](docs/ProfileApi.md#getprofileapp) | **Get** /profile/apps/{appId} | Authorized App View
*ProfileApi* | [**GetProfileApps**](docs/ProfileApi.md#getprofileapps) | **Get** /profile/apps | Authorized Apps List
*ProfileApi* | [**GetProfileGrants**](docs/ProfileApi.md#getprofilegrants) | **Get** /profile/grants | Grants List
*ProfileApi* | [**GetProfileLogin**](docs/ProfileApi.md#getprofilelogin) | **Get** /profile/logins/{loginId} | Login View
*ProfileApi* | [**GetProfileLogins**](docs/ProfileApi.md#getprofilelogins) | **Get** /profile/logins | Logins List
*ProfileApi* | [**GetSSHKey**](docs/ProfileApi.md#getsshkey) | **Get** /profile/sshkeys/{sshKeyId} | SSH Key View
*ProfileApi* | [**GetSSHKeys**](docs/ProfileApi.md#getsshkeys) | **Get** /profile/sshkeys | SSH Keys List
*ProfileApi* | [**GetTrustedDevice**](docs/ProfileApi.md#gettrusteddevice) | **Get** /profile/devices/{deviceId} | Trusted Device View
*ProfileApi* | [**GetUserPreferences**](docs/ProfileApi.md#getuserpreferences) | **Get** /profile/preferences | User Preferences View
*ProfileApi* | [**RevokeTrustedDevice**](docs/ProfileApi.md#revoketrusteddevice) | **Delete** /profile/devices/{deviceId} | Trusted Device Revoke
*ProfileApi* | [**TfaConfirm**](docs/ProfileApi.md#tfaconfirm) | **Post** /profile/tfa-enable-confirm | Two Factor Authentication Confirm/Enable
*ProfileApi* | [**TfaDisable**](docs/ProfileApi.md#tfadisable) | **Post** /profile/tfa-disable | Two Factor Authentication Disable
*ProfileApi* | [**TfaEnable**](docs/ProfileApi.md#tfaenable) | **Post** /profile/tfa-enable | Two Factor Secret Create
*ProfileApi* | [**UpdatePersonalAccessToken**](docs/ProfileApi.md#updatepersonalaccesstoken) | **Put** /profile/tokens/{tokenId} | Personal Access Token Update
*ProfileApi* | [**UpdateProfile**](docs/ProfileApi.md#updateprofile) | **Put** /profile | Profile Update
*ProfileApi* | [**UpdateSSHKey**](docs/ProfileApi.md#updatesshkey) | **Put** /profile/sshkeys/{sshKeyId} | SSH Key Update
*ProfileApi* | [**UpdateUserPreferences**](docs/ProfileApi.md#updateuserpreferences) | **Put** /profile/preferences | User Preferences Update
*RegionsApi* | [**GetRegion**](docs/RegionsApi.md#getregion) | **Get** /regions/{regionId} | Region View
*RegionsApi* | [**GetRegions**](docs/RegionsApi.md#getregions) | **Get** /regions | Regions List
*StackScriptsApi* | [**AddStackScript**](docs/StackScriptsApi.md#addstackscript) | **Post** /linode/stackscripts | StackScript Create
*StackScriptsApi* | [**DeleteStackScript**](docs/StackScriptsApi.md#deletestackscript) | **Delete** /linode/stackscripts/{stackscriptId} | StackScript Delete
*StackScriptsApi* | [**GetStackScript**](docs/StackScriptsApi.md#getstackscript) | **Get** /linode/stackscripts/{stackscriptId} | StackScript View
*StackScriptsApi* | [**GetStackScripts**](docs/StackScriptsApi.md#getstackscripts) | **Get** /linode/stackscripts | StackScripts List
*StackScriptsApi* | [**UpdateStackScript**](docs/StackScriptsApi.md#updatestackscript) | **Put** /linode/stackscripts/{stackscriptId} | StackScript Update
*SupportApi* | [**CloseTicket**](docs/SupportApi.md#closeticket) | **Post** /support/tickets/{ticketId}/close | Support Ticket Close
*SupportApi* | [**CreateTicket**](docs/SupportApi.md#createticket) | **Post** /support/tickets | Support Ticket Open
*SupportApi* | [**CreateTicketAttachment**](docs/SupportApi.md#createticketattachment) | **Post** /support/tickets/{ticketId}/attachments | Ticket Attachment Create
*SupportApi* | [**CreateTicketReply**](docs/SupportApi.md#createticketreply) | **Post** /support/tickets/{ticketId}/replies | Reply Create
*SupportApi* | [**GetTicket**](docs/SupportApi.md#getticket) | **Get** /support/tickets/{ticketId} | Support Ticket View
*SupportApi* | [**GetTicketReplies**](docs/SupportApi.md#getticketreplies) | **Get** /support/tickets/{ticketId}/replies | Replies List
*SupportApi* | [**GetTickets**](docs/SupportApi.md#gettickets) | **Get** /support/tickets | Support Tickets List
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /tags | New Tag Create
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /tags/{label} | Tag Delete
*TagsApi* | [**GetTaggedObjects**](docs/TagsApi.md#gettaggedobjects) | **Get** /tags/{label} | Tagged Objects List
*TagsApi* | [**GetTags**](docs/TagsApi.md#gettags) | **Get** /tags | Tags List
*VolumesApi* | [**AttachVolume**](docs/VolumesApi.md#attachvolume) | **Post** /volumes/{volumeId}/attach | Volume Attach
*VolumesApi* | [**CloneVolume**](docs/VolumesApi.md#clonevolume) | **Post** /volumes/{volumeId}/clone | Volume Clone
*VolumesApi* | [**CreateVolume**](docs/VolumesApi.md#createvolume) | **Post** /volumes | Volume Create
*VolumesApi* | [**DeleteVolume**](docs/VolumesApi.md#deletevolume) | **Delete** /volumes/{volumeId} | Volume Delete
*VolumesApi* | [**DetachVolume**](docs/VolumesApi.md#detachvolume) | **Post** /volumes/{volumeId}/detach | Volume Detach
*VolumesApi* | [**GetVolume**](docs/VolumesApi.md#getvolume) | **Get** /volumes/{volumeId} | Volume View
*VolumesApi* | [**GetVolumes**](docs/VolumesApi.md#getvolumes) | **Get** /volumes | Volumes List
*VolumesApi* | [**ResizeVolume**](docs/VolumesApi.md#resizevolume) | **Post** /volumes/{volumeId}/resize | Volume Resize
*VolumesApi* | [**UpdateVolume**](docs/VolumesApi.md#updatevolume) | **Put** /volumes/{volumeId} | Volume Update


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountActivePromotions](docs/AccountActivePromotions.md)
 - [AccountCreditCard](docs/AccountCreditCard.md)
 - [AccountSettings](docs/AccountSettings.md)
 - [AuthorizedApp](docs/AuthorizedApp.md)
 - [Backup](docs/Backup.md)
 - [BackupDisks](docs/BackupDisks.md)
 - [CreditCard](docs/CreditCard.md)
 - [Device](docs/Device.md)
 - [Devices](docs/Devices.md)
 - [Disk](docs/Disk.md)
 - [DiskRequest](docs/DiskRequest.md)
 - [Domain](docs/Domain.md)
 - [DomainRecord](docs/DomainRecord.md)
 - [EntityTransfer](docs/EntityTransfer.md)
 - [EntityTransferEntities](docs/EntityTransferEntities.md)
 - [ErrorObject](docs/ErrorObject.md)
 - [Event](docs/Event.md)
 - [EventEntity](docs/EventEntity.md)
 - [EventSecondaryEntity](docs/EventSecondaryEntity.md)
 - [Firewall](docs/Firewall.md)
 - [FirewallDevices](docs/FirewallDevices.md)
 - [FirewallDevicesEntity](docs/FirewallDevicesEntity.md)
 - [FirewallRuleConfig](docs/FirewallRuleConfig.md)
 - [FirewallRuleConfigAddresses](docs/FirewallRuleConfigAddresses.md)
 - [FirewallRules](docs/FirewallRules.md)
 - [Grant](docs/Grant.md)
 - [GrantsResponse](docs/GrantsResponse.md)
 - [GrantsResponseGlobal](docs/GrantsResponseGlobal.md)
 - [IPAddress](docs/IPAddress.md)
 - [IPAddressPrivate](docs/IPAddressPrivate.md)
 - [IPAddressV6LinkLocal](docs/IPAddressV6LinkLocal.md)
 - [IPAddressV6Slaac](docs/IPAddressV6Slaac.md)
 - [IPv6Pool](docs/IPv6Pool.md)
 - [IPv6Range](docs/IPv6Range.md)
 - [ImagePrivate](docs/ImagePrivate.md)
 - [ImagePublic](docs/ImagePublic.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject10](docs/InlineObject10.md)
 - [InlineObject11](docs/InlineObject11.md)
 - [InlineObject12](docs/InlineObject12.md)
 - [InlineObject13](docs/InlineObject13.md)
 - [InlineObject14](docs/InlineObject14.md)
 - [InlineObject15](docs/InlineObject15.md)
 - [InlineObject16](docs/InlineObject16.md)
 - [InlineObject17](docs/InlineObject17.md)
 - [InlineObject18](docs/InlineObject18.md)
 - [InlineObject19](docs/InlineObject19.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20016Snapshot](docs/InlineResponse20016Snapshot.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20027Data](docs/InlineResponse20027Data.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20028Data](docs/InlineResponse20028Data.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse20045](docs/InlineResponse20045.md)
 - [InlineResponse20046](docs/InlineResponse20046.md)
 - [InlineResponse20047](docs/InlineResponse20047.md)
 - [InlineResponse20048](docs/InlineResponse20048.md)
 - [InlineResponse20049](docs/InlineResponse20049.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse20050](docs/InlineResponse20050.md)
 - [InlineResponse20051](docs/InlineResponse20051.md)
 - [InlineResponse20052](docs/InlineResponse20052.md)
 - [InlineResponse20053](docs/InlineResponse20053.md)
 - [InlineResponse20054](docs/InlineResponse20054.md)
 - [InlineResponse20055](docs/InlineResponse20055.md)
 - [InlineResponse20056](docs/InlineResponse20056.md)
 - [InlineResponse20057](docs/InlineResponse20057.md)
 - [InlineResponse20058](docs/InlineResponse20058.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [InlineResponse409Errors](docs/InlineResponse409Errors.md)
 - [InlineResponseDefault](docs/InlineResponseDefault.md)
 - [Invoice](docs/Invoice.md)
 - [InvoiceItem](docs/InvoiceItem.md)
 - [Kernel](docs/Kernel.md)
 - [LKECluster](docs/LKECluster.md)
 - [LKENodePool](docs/LKENodePool.md)
 - [LKENodePoolRequestBody](docs/LKENodePoolRequestBody.md)
 - [LKENodePoolRequestBodyDisks](docs/LKENodePoolRequestBodyDisks.md)
 - [LKENodeStatus](docs/LKENodeStatus.md)
 - [LKEVersion](docs/LKEVersion.md)
 - [Linode](docs/Linode.md)
 - [LinodeBase](docs/LinodeBase.md)
 - [LinodeConfig](docs/LinodeConfig.md)
 - [LinodeConfigHelpers](docs/LinodeConfigHelpers.md)
 - [LinodeRequest](docs/LinodeRequest.md)
 - [LinodeStats](docs/LinodeStats.md)
 - [LinodeStatsIo](docs/LinodeStatsIo.md)
 - [LinodeStatsNetv4](docs/LinodeStatsNetv4.md)
 - [LinodeStatsNetv6](docs/LinodeStatsNetv6.md)
 - [LinodeType](docs/LinodeType.md)
 - [LinodeTypeAddons](docs/LinodeTypeAddons.md)
 - [LinodeTypeAddonsBackups](docs/LinodeTypeAddonsBackups.md)
 - [LinodeTypeAddonsBackupsPrice](docs/LinodeTypeAddonsBackupsPrice.md)
 - [LinodeTypePrice](docs/LinodeTypePrice.md)
 - [Login](docs/Login.md)
 - [LongviewClient](docs/LongviewClient.md)
 - [LongviewClientApps](docs/LongviewClientApps.md)
 - [LongviewPlan](docs/LongviewPlan.md)
 - [LongviewSubscription](docs/LongviewSubscription.md)
 - [LongviewSubscriptionPrice](docs/LongviewSubscriptionPrice.md)
 - [Maintenance](docs/Maintenance.md)
 - [MaintenanceEntity](docs/MaintenanceEntity.md)
 - [ManagedContact](docs/ManagedContact.md)
 - [ManagedContactPhone](docs/ManagedContactPhone.md)
 - [ManagedCredential](docs/ManagedCredential.md)
 - [ManagedIssue](docs/ManagedIssue.md)
 - [ManagedIssueEntity](docs/ManagedIssueEntity.md)
 - [ManagedLinodeSettings](docs/ManagedLinodeSettings.md)
 - [ManagedLinodeSettingsSsh](docs/ManagedLinodeSettingsSsh.md)
 - [ManagedService](docs/ManagedService.md)
 - [NetworkingFirewallsDevices](docs/NetworkingFirewallsDevices.md)
 - [NodeBalancer](docs/NodeBalancer.md)
 - [NodeBalancerConfig](docs/NodeBalancerConfig.md)
 - [NodeBalancerConfigNodesStatus](docs/NodeBalancerConfigNodesStatus.md)
 - [NodeBalancerNode](docs/NodeBalancerNode.md)
 - [NodeBalancerStats](docs/NodeBalancerStats.md)
 - [NodeBalancerStatsData](docs/NodeBalancerStatsData.md)
 - [NodeBalancerStatsDataTraffic](docs/NodeBalancerStatsDataTraffic.md)
 - [NodeBalancerTransfer](docs/NodeBalancerTransfer.md)
 - [Notification](docs/Notification.md)
 - [NotificationEntity](docs/NotificationEntity.md)
 - [OAuthClient](docs/OAuthClient.md)
 - [ObjectStorageBucket](docs/ObjectStorageBucket.md)
 - [ObjectStorageCluster](docs/ObjectStorageCluster.md)
 - [ObjectStorageKey](docs/ObjectStorageKey.md)
 - [ObjectStorageKeyBucketAccess](docs/ObjectStorageKeyBucketAccess.md)
 - [ObjectStorageObject](docs/ObjectStorageObject.md)
 - [ObjectStorageSSL](docs/ObjectStorageSSL.md)
 - [ObjectStorageSSLResponse](docs/ObjectStorageSSLResponse.md)
 - [PaginationEnvelope](docs/PaginationEnvelope.md)
 - [PayPal](docs/PayPal.md)
 - [PayPalExecute](docs/PayPalExecute.md)
 - [Payment](docs/Payment.md)
 - [PaymentRequest](docs/PaymentRequest.md)
 - [PersonalAccessToken](docs/PersonalAccessToken.md)
 - [Profile](docs/Profile.md)
 - [ProfileReferrals](docs/ProfileReferrals.md)
 - [Region](docs/Region.md)
 - [RegionResolvers](docs/RegionResolvers.md)
 - [RescueDevices](docs/RescueDevices.md)
 - [SSHKey](docs/SSHKey.md)
 - [SSHKeyRequest](docs/SSHKeyRequest.md)
 - [StackScript](docs/StackScript.md)
 - [StatsData](docs/StatsData.md)
 - [StatsDataAvailable](docs/StatsDataAvailable.md)
 - [SupportTicket](docs/SupportTicket.md)
 - [SupportTicketEntity](docs/SupportTicketEntity.md)
 - [SupportTicketReply](docs/SupportTicketReply.md)
 - [SupportTicketRequest](docs/SupportTicketRequest.md)
 - [Tag](docs/Tag.md)
 - [Transfer](docs/Transfer.md)
 - [TrustedDevice](docs/TrustedDevice.md)
 - [User](docs/User.md)
 - [UserDefinedField](docs/UserDefinedField.md)
 - [Volume](docs/Volume.md)


## Documentation For Authorization



### oauth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://login.linode.com/oauth/authorize
- **Scopes**: 
 - **account:read_only**: Allows access to GET information about your Account.
 - **account:read_write**: Allows access to all endpoints related to your Account.
 - **domains:read_only**: Allows access to GET Domains on your Account.
 - **domains:read_write**: Allows access to all Domain endpoints.
 - **events:read_only**: Allows access to GET your Events.
 - **events:read_write**: Allows access to all endpoints related to your Events.
 - **firewall:read_only**: Allows access to GET information about your Firewalls.
 - **firewall:read_write**: Allows acces to all Firewall endpoints.
 - **images:read_only**: Allows access to GET your Images.
 - **images:read_write**: Allows access to all endpoints related to your Images.
 - **ips:read_only**: Allows access to GET your ips.
 - **ips:read_write**: Allows access to all endpoints related to your ips.
 - **linodes:read_only**: Allows access to GET Linodes on your Account.
 - **linodes:read_write**: Allow access to all endpoints related to your Linodes.
 - **lke:read_only**: Allows access to GET LKE Clusters on your Account.
 - **lke:read_write**: Allows access to all endpoints related to LKE Clusters on your Account.
 - **longview:read_only**: Allows access to GET your Longview Clients.
 - **longview:read_write**: Allows access to all endpoints related to your Longview Clients.
 - **maintenance:read_only**: Allows access to GET information about Maintenance on your account.
 - **nodebalancers:read_only**: Allows access to GET NodeBalancers on your Account.
 - **nodebalancers:read_write**: Allows access to all NodeBalancer endpoints.
 - **object_storage:read_only**: Allows access to GET information related to your Object Storage.
 - **object_storage:read_write**: Allows access to all Object Storage endpoints.
 - **stackscripts:read_only**: Allows access to GET your StackScripts.
 - **stackscripts:read_write**: Allows access to all endpoints related to your StackScripts.
 - **volumes:read_only**: Allows access to GET your Volumes.
 - **volumes:read_write**: Allows access to all endpoints related to your Volumes.

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### personalAccessToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@linode.com

