// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetResourcefile
    {
        /// <summary>
        /// Gets the contents of a resource file. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).
        /// </summary>
        public static Task<GetResourcefileResult> InvokeAsync(GetResourcefileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcefileResult>("google-native:apigee/v1:getResourcefile", args ?? new GetResourcefileArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the contents of a resource file. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).
        /// </summary>
        public static Output<GetResourcefileResult> Invoke(GetResourcefileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourcefileResult>("google-native:apigee/v1:getResourcefile", args ?? new GetResourcefileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcefileArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetResourcefileArgs()
        {
        }
        public static new GetResourcefileArgs Empty => new GetResourcefileArgs();
    }

    public sealed class GetResourcefileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetResourcefileInvokeArgs()
        {
        }
        public static new GetResourcefileInvokeArgs Empty => new GetResourcefileInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourcefileResult
    {
        /// <summary>
        /// The HTTP Content-Type header value specifying the content type of the body.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// The HTTP request/response body as raw binary.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// Application specific response metadata. Must be set in the first response for streaming APIs.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Extensions;

        [OutputConstructor]
        private GetResourcefileResult(
            string contentType,

            string data,

            ImmutableArray<ImmutableDictionary<string, string>> extensions)
        {
            ContentType = contentType;
            Data = data;
            Extensions = extensions;
        }
    }
}
