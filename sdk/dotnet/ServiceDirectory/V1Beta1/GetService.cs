// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceDirectory.V1Beta1
{
    public static class GetService
    {
        /// <summary>
        /// Gets a service.
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("google-native:servicedirectory/v1beta1:getService", args ?? new GetServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a service.
        /// </summary>
        public static Output<GetServiceResult> Invoke(GetServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceResult>("google-native:servicedirectory/v1beta1:getService", args ?? new GetServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetServiceArgs()
        {
        }
        public static new GetServiceArgs Empty => new GetServiceArgs();
    }

    public sealed class GetServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetServiceInvokeArgs()
        {
        }
        public static new GetServiceInvokeArgs Empty => new GetServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The timestamp when the service was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Endpoints associated with this service. Returned on LookupService.ResolveService. Control plane clients should use RegistrationService.ListEndpoints.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointResponse> Endpoints;
        /// <summary>
        /// Optional. Metadata for the service. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 2000 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Immutable. The resource name for the service in the format `projects/*/locations/*/namespaces/*/services/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A globally unique identifier (in UUID4 format) for this service.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The timestamp when the service was last updated. Note: endpoints being created/deleted/updated within the service are not considered service updates for the purpose of this timestamp.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetServiceResult(
            string createTime,

            ImmutableArray<Outputs.EndpointResponse> endpoints,

            ImmutableDictionary<string, string> metadata,

            string name,

            string uid,

            string updateTime)
        {
            CreateTime = createTime;
            Endpoints = endpoints;
            Metadata = metadata;
            Name = name;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
