// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetNetworkAttachment
    {
        /// <summary>
        /// Returns the specified NetworkAttachment resource in the given scope.
        /// </summary>
        public static Task<GetNetworkAttachmentResult> InvokeAsync(GetNetworkAttachmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkAttachmentResult>("google-native:compute/beta:getNetworkAttachment", args ?? new GetNetworkAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified NetworkAttachment resource in the given scope.
        /// </summary>
        public static Output<GetNetworkAttachmentResult> Invoke(GetNetworkAttachmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkAttachmentResult>("google-native:compute/beta:getNetworkAttachment", args ?? new GetNetworkAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkAttachmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkAttachment", required: true)]
        public string NetworkAttachment { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetNetworkAttachmentArgs()
        {
        }
        public static new GetNetworkAttachmentArgs Empty => new GetNetworkAttachmentArgs();
    }

    public sealed class GetNetworkAttachmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("networkAttachment", required: true)]
        public Input<string> NetworkAttachment { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetNetworkAttachmentInvokeArgs()
        {
        }
        public static new GetNetworkAttachmentInvokeArgs Empty => new GetNetworkAttachmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkAttachmentResult
    {
        /// <summary>
        /// An array of connections for all the producers connected to this network attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkAttachmentConnectedEndpointResponse> ConnectionEndpoints;
        public readonly string ConnectionPreference;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
        /// </summary>
        public readonly ImmutableArray<string> ProducerAcceptLists;
        /// <summary>
        /// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
        /// </summary>
        public readonly ImmutableArray<string> ProducerRejectLists;
        /// <summary>
        /// URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource's resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
        /// </summary>
        public readonly ImmutableArray<string> Subnetworks;

        [OutputConstructor]
        private GetNetworkAttachmentResult(
            ImmutableArray<Outputs.NetworkAttachmentConnectedEndpointResponse> connectionEndpoints,

            string connectionPreference,

            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string name,

            string network,

            ImmutableArray<string> producerAcceptLists,

            ImmutableArray<string> producerRejectLists,

            string region,

            string selfLink,

            string selfLinkWithId,

            ImmutableArray<string> subnetworks)
        {
            ConnectionEndpoints = connectionEndpoints;
            ConnectionPreference = connectionPreference;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            Network = network;
            ProducerAcceptLists = producerAcceptLists;
            ProducerRejectLists = producerRejectLists;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            Subnetworks = subnetworks;
        }
    }
}
