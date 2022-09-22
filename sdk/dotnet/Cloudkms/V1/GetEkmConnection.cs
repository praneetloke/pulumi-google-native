// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    public static class GetEkmConnection
    {
        /// <summary>
        /// Returns metadata for a given EkmConnection.
        /// </summary>
        public static Task<GetEkmConnectionResult> InvokeAsync(GetEkmConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEkmConnectionResult>("google-native:cloudkms/v1:getEkmConnection", args ?? new GetEkmConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Returns metadata for a given EkmConnection.
        /// </summary>
        public static Output<GetEkmConnectionResult> Invoke(GetEkmConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEkmConnectionResult>("google-native:cloudkms/v1:getEkmConnection", args ?? new GetEkmConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEkmConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("ekmConnectionId", required: true)]
        public string EkmConnectionId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEkmConnectionArgs()
        {
        }
        public static new GetEkmConnectionArgs Empty => new GetEkmConnectionArgs();
    }

    public sealed class GetEkmConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ekmConnectionId", required: true)]
        public Input<string> EkmConnectionId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEkmConnectionInvokeArgs()
        {
        }
        public static new GetEkmConnectionInvokeArgs Empty => new GetEkmConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetEkmConnectionResult
    {
        /// <summary>
        /// The time at which the EkmConnection was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The resource name for the EkmConnection in the format `projects/*/locations/*/ekmConnections/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceResolverResponse> ServiceResolvers;

        [OutputConstructor]
        private GetEkmConnectionResult(
            string createTime,

            string etag,

            string name,

            ImmutableArray<Outputs.ServiceResolverResponse> serviceResolvers)
        {
            CreateTime = createTime;
            Etag = etag;
            Name = name;
            ServiceResolvers = serviceResolvers;
        }
    }
}
