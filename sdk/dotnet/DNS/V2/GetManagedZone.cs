// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2
{
    public static class GetManagedZone
    {
        /// <summary>
        /// Fetches the representation of an existing ManagedZone.
        /// </summary>
        public static Task<GetManagedZoneResult> InvokeAsync(GetManagedZoneArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedZoneResult>("google-native:dns/v2:getManagedZone", args ?? new GetManagedZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches the representation of an existing ManagedZone.
        /// </summary>
        public static Output<GetManagedZoneResult> Invoke(GetManagedZoneInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetManagedZoneResult>("google-native:dns/v2:getManagedZone", args ?? new GetManagedZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedZoneArgs : global::Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public string? ClientOperationId { get; set; }

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("managedZone", required: true)]
        public string ManagedZone { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetManagedZoneArgs()
        {
        }
        public static new GetManagedZoneArgs Empty => new GetManagedZoneArgs();
    }

    public sealed class GetManagedZoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetManagedZoneInvokeArgs()
        {
        }
        public static new GetManagedZoneInvokeArgs Empty => new GetManagedZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedZoneResult
    {
        public readonly Outputs.ManagedZoneCloudLoggingConfigResponse CloudLoggingConfig;
        /// <summary>
        /// The time that this resource was created on the server. This is in RFC3339 text format. Output only.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the managed zone's function.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        public readonly string DnsName;
        /// <summary>
        /// DNSSEC configuration.
        /// </summary>
        public readonly Outputs.ManagedZoneDnsSecConfigResponse DnssecConfig;
        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this field contains the set of destinations to forward to.
        /// </summary>
        public readonly Outputs.ManagedZoneForwardingConfigResponse ForwardingConfig;
        public readonly string Kind;
        /// <summary>
        /// User labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// User assigned name for this resource. Must be unique within the project. The name must be 1-63 characters long, must begin with a letter, end with a letter or digit, and only contain lowercase letters, digits or dashes.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optionally specifies the NameServerSet for this ManagedZone. A NameServerSet is a set of DNS name servers that all host the same ManagedZones. Most users leave this field unset. If you need to use this field, contact your account team.
        /// </summary>
        public readonly string NameServerSet;
        /// <summary>
        /// Delegate your managed_zone to these virtual name servers; defined by the server (output only)
        /// </summary>
        public readonly ImmutableArray<string> NameServers;
        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field contains the network to peer with.
        /// </summary>
        public readonly Outputs.ManagedZonePeeringConfigResponse PeeringConfig;
        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        /// </summary>
        public readonly Outputs.ManagedZonePrivateVisibilityConfigResponse PrivateVisibilityConfig;
        /// <summary>
        /// The presence of this field indicates that this is a managed reverse lookup zone and Cloud DNS resolves reverse lookup queries using automatically configured records for VPC resources. This only applies to networks listed under private_visibility_config.
        /// </summary>
        public readonly Outputs.ManagedZoneReverseLookupConfigResponse ReverseLookupConfig;
        /// <summary>
        /// This field links to the associated service directory namespace. Do not set this field for public zones or forwarding zones.
        /// </summary>
        public readonly Outputs.ManagedZoneServiceDirectoryConfigResponse ServiceDirectoryConfig;
        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetManagedZoneResult(
            Outputs.ManagedZoneCloudLoggingConfigResponse cloudLoggingConfig,

            string creationTime,

            string description,

            string dnsName,

            Outputs.ManagedZoneDnsSecConfigResponse dnssecConfig,

            Outputs.ManagedZoneForwardingConfigResponse forwardingConfig,

            string kind,

            ImmutableDictionary<string, string> labels,

            string name,

            string nameServerSet,

            ImmutableArray<string> nameServers,

            Outputs.ManagedZonePeeringConfigResponse peeringConfig,

            Outputs.ManagedZonePrivateVisibilityConfigResponse privateVisibilityConfig,

            Outputs.ManagedZoneReverseLookupConfigResponse reverseLookupConfig,

            Outputs.ManagedZoneServiceDirectoryConfigResponse serviceDirectoryConfig,

            string visibility)
        {
            CloudLoggingConfig = cloudLoggingConfig;
            CreationTime = creationTime;
            Description = description;
            DnsName = dnsName;
            DnssecConfig = dnssecConfig;
            ForwardingConfig = forwardingConfig;
            Kind = kind;
            Labels = labels;
            Name = name;
            NameServerSet = nameServerSet;
            NameServers = nameServers;
            PeeringConfig = peeringConfig;
            PrivateVisibilityConfig = privateVisibilityConfig;
            ReverseLookupConfig = reverseLookupConfig;
            ServiceDirectoryConfig = serviceDirectoryConfig;
            Visibility = visibility;
        }
    }
}
