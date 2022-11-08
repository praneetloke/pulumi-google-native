// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    public static class GetInternalRange
    {
        /// <summary>
        /// Gets details of a single InternalRange.
        /// </summary>
        public static Task<GetInternalRangeResult> InvokeAsync(GetInternalRangeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInternalRangeResult>("google-native:networkconnectivity/v1:getInternalRange", args ?? new GetInternalRangeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single InternalRange.
        /// </summary>
        public static Output<GetInternalRangeResult> Invoke(GetInternalRangeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInternalRangeResult>("google-native:networkconnectivity/v1:getInternalRange", args ?? new GetInternalRangeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInternalRangeArgs : global::Pulumi.InvokeArgs
    {
        [Input("internalRangeId", required: true)]
        public string InternalRangeId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInternalRangeArgs()
        {
        }
        public static new GetInternalRangeArgs Empty => new GetInternalRangeArgs();
    }

    public sealed class GetInternalRangeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("internalRangeId", required: true)]
        public Input<string> InternalRangeId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInternalRangeInvokeArgs()
        {
        }
        public static new GetInternalRangeInvokeArgs Empty => new GetInternalRangeInvokeArgs();
    }


    [OutputType]
    public sealed class GetInternalRangeResult
    {
        /// <summary>
        /// Time when the InternalRange was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// A description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// IP range that this InternalRange defines.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// User-defined labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of a InternalRange. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL or resource ID of the network in which to reserve the Internal Range. The network cannot be deleted if there are any reserved Internal Ranges referring to it. Legacy network is not supported. This can only be specified for a global internal address. Example: - URL: /compute/v1/projects/{project}/global/networks/{resourceId} - ID: network123
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Optional. Types of resources that are allowed to overlap with the current InternalRange.
        /// </summary>
        public readonly ImmutableArray<string> Overlaps;
        /// <summary>
        /// The type of peering set for this InternalRange.
        /// </summary>
        public readonly string Peering;
        /// <summary>
        /// An alternate to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, it's an error if the range sizes don't match. Can also be used during updates to change the range size.
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
        /// </summary>
        public readonly ImmutableArray<string> TargetCidrRange;
        /// <summary>
        /// Time when the InternalRange was updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The type of usage set for this InternalRange.
        /// </summary>
        public readonly string Usage;
        /// <summary>
        /// The list of resources that refer to this internal range. Resources that use the InternalRange for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this InternalRange. Having a user, based on this reference, prevents deletion of the InternalRange referred to. Can be empty.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetInternalRangeResult(
            string createTime,

            string description,

            string ipCidrRange,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            ImmutableArray<string> overlaps,

            string peering,

            int prefixLength,

            ImmutableArray<string> targetCidrRange,

            string updateTime,

            string usage,

            ImmutableArray<string> users)
        {
            CreateTime = createTime;
            Description = description;
            IpCidrRange = ipCidrRange;
            Labels = labels;
            Name = name;
            Network = network;
            Overlaps = overlaps;
            Peering = peering;
            PrefixLength = prefixLength;
            TargetCidrRange = targetCidrRange;
            UpdateTime = updateTime;
            Usage = usage;
            Users = users;
        }
    }
}
