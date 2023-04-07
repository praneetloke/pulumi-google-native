// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2.Outputs
{

    /// <summary>
    /// HTTP request. The task will be pushed to the worker as an HTTP request. An HTTP request embodies a url, an http method, headers, body and authorization for the http task.
    /// </summary>
    [OutputType]
    public sealed class HttpRequestResponse
    {
        /// <summary>
        /// HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a task with an incompatible HttpMethod.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// HTTP request headers. This map contains the header field names and values. Headers can be set when running the task is created or task is created. These headers represent a subset of the headers that will accompany the task's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is: * Any header that is prefixed with "X-CloudTasks-" will be treated as service header. Service headers define properties of the task and are predefined in CloudTask. * Host: This will be computed by Cloud Tasks and derived from HttpRequest.url. * Content-Length: This will be computed by Cloud Tasks. * User-Agent: This will be set to `"Google-Cloud-Tasks"`. * `X-Google-*`: Google use only. * `X-AppEngine-*`: Google use only. `Content-Type` won't be set by Cloud Tasks. You can explicitly set `Content-Type` to a media type when the task is created. For example, `Content-Type` can be set to `"application/octet-stream"` or `"application/json"`. Headers which can have multiple values (according to RFC2616) can be specified using comma-separated values. The size of the headers must be less than 80KB.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Headers;
        /// <summary>
        /// The HTTP method to use for the request. The default is POST.
        /// </summary>
        public readonly string HttpMethod;
        /// <summary>
        /// If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        /// </summary>
        public readonly Outputs.OAuthTokenResponse OauthToken;
        /// <summary>
        /// If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        /// </summary>
        public readonly Outputs.OidcTokenResponse OidcToken;
        /// <summary>
        /// The full url path that the request will be sent to. This string must begin with either "http://" or "https://". Some examples are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Tasks will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding. The `Location` header response from a redirect response [`300` - `399`] may be followed. The redirect is not counted as a separate attempt.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private HttpRequestResponse(
            string body,

            ImmutableDictionary<string, string> headers,

            string httpMethod,

            Outputs.OAuthTokenResponse oauthToken,

            Outputs.OidcTokenResponse oidcToken,

            string url)
        {
            Body = body;
            Headers = headers;
            HttpMethod = httpMethod;
            OauthToken = oauthToken;
            OidcToken = oidcToken;
            Url = url;
        }
    }
}
