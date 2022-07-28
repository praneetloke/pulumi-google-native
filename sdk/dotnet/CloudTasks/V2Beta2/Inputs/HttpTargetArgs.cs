// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2.Inputs
{

    /// <summary>
    /// HTTP target. When specified as a Queue, all the tasks with [HttpRequest] will be overridden according to the target.
    /// </summary>
    public sealed class HttpTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerOverrides")]
        private InputList<Inputs.HeaderOverrideArgs>? _headerOverrides;

        /// <summary>
        /// HTTP target headers. This map contains the header field names and values. Headers will be set when running the task is created and/or task is created. These headers represent a subset of the headers that will accompany the task's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is: * Any header that is prefixed with "X-Google-Cloud-Tasks-" will be treated as service header. Service headers define properties of the task and are predefined in CloudTask. * Host: This will be computed by Cloud Tasks and derived from HttpRequest.url. * Content-Length: This will be computed by Cloud Tasks. * User-Agent: This will be set to `"Google-Cloud-Tasks"`. * `X-Google-*`: Google use only. * `X-AppEngine-*`: Google use only. `Content-Type` won't be set by Cloud Tasks. You can explicitly set `Content-Type` to a media type when the task is created. For example, `Content-Type` can be set to `"application/octet-stream"` or `"application/json"`. Headers which can have multiple values (according to RFC2616) can be specified using comma-separated values. The size of the headers must be less than 80KB. Queue-level headers to override headers of all the tasks in the queue.
        /// </summary>
        public InputList<Inputs.HeaderOverrideArgs> HeaderOverrides
        {
            get => _headerOverrides ?? (_headerOverrides = new InputList<Inputs.HeaderOverrideArgs>());
            set => _headerOverrides = value;
        }

        /// <summary>
        /// The HTTP method to use for the request. When specified, it will override HttpRequest for the task. Note that if the value is set to HttpMethod the HttpRequest of the task will be ignored at execution time.
        /// </summary>
        [Input("httpMethod")]
        public Input<Pulumi.GoogleNative.CloudTasks.V2Beta2.HttpTargetHttpMethod>? HttpMethod { get; set; }

        /// <summary>
        /// If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        /// </summary>
        [Input("oauthToken")]
        public Input<Inputs.OAuthTokenArgs>? OauthToken { get; set; }

        /// <summary>
        /// If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        /// </summary>
        [Input("oidcToken")]
        public Input<Inputs.OidcTokenArgs>? OidcToken { get; set; }

        /// <summary>
        /// Uri override. When specified modifies the execution Uri for all the tasks in the queue.
        /// </summary>
        [Input("uriOverride")]
        public Input<Inputs.UriOverrideArgs>? UriOverride { get; set; }

        public HttpTargetArgs()
        {
        }
        public static new HttpTargetArgs Empty => new HttpTargetArgs();
    }
}
