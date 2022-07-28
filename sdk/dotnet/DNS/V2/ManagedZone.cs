// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2
{
    /// <summary>
    /// Creates a new ManagedZone.
    /// </summary>
    [GoogleNativeResourceType("google-native:dns/v2:ManagedZone")]
    public partial class ManagedZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        /// </summary>
        [Output("clientOperationId")]
        public Output<string?> ClientOperationId { get; private set; } = null!;

        [Output("cloudLoggingConfig")]
        public Output<Outputs.ManagedZoneCloudLoggingConfigResponse> CloudLoggingConfig { get; private set; } = null!;

        /// <summary>
        /// The time that this resource was created on the server. This is in RFC3339 text format. Output only.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the managed zone's function.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// DNSSEC configuration.
        /// </summary>
        [Output("dnssecConfig")]
        public Output<Outputs.ManagedZoneDnsSecConfigResponse> DnssecConfig { get; private set; } = null!;

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this field contains the set of destinations to forward to.
        /// </summary>
        [Output("forwardingConfig")]
        public Output<Outputs.ManagedZoneForwardingConfigResponse> ForwardingConfig { get; private set; } = null!;

        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// User labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// User assigned name for this resource. Must be unique within the project. The name must be 1-63 characters long, must begin with a letter, end with a letter or digit, and only contain lowercase letters, digits or dashes.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optionally specifies the NameServerSet for this ManagedZone. A NameServerSet is a set of DNS name servers that all host the same ManagedZones. Most users leave this field unset. If you need to use this field, contact your account team.
        /// </summary>
        [Output("nameServerSet")]
        public Output<string> NameServerSet { get; private set; } = null!;

        /// <summary>
        /// Delegate your managed_zone to these virtual name servers; defined by the server (output only)
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field contains the network to peer with.
        /// </summary>
        [Output("peeringConfig")]
        public Output<Outputs.ManagedZonePeeringConfigResponse> PeeringConfig { get; private set; } = null!;

        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        /// </summary>
        [Output("privateVisibilityConfig")]
        public Output<Outputs.ManagedZonePrivateVisibilityConfigResponse> PrivateVisibilityConfig { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The presence of this field indicates that this is a managed reverse lookup zone and Cloud DNS resolves reverse lookup queries using automatically configured records for VPC resources. This only applies to networks listed under private_visibility_config.
        /// </summary>
        [Output("reverseLookupConfig")]
        public Output<Outputs.ManagedZoneReverseLookupConfigResponse> ReverseLookupConfig { get; private set; } = null!;

        /// <summary>
        /// This field links to the associated service directory namespace. Do not set this field for public zones or forwarding zones.
        /// </summary>
        [Output("serviceDirectoryConfig")]
        public Output<Outputs.ManagedZoneServiceDirectoryConfigResponse> ServiceDirectoryConfig { get; private set; } = null!;

        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedZone(string name, ManagedZoneArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:dns/v2:ManagedZone", name, args ?? new ManagedZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedZone(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dns/v2:ManagedZone", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedZone Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedZone(name, id, options);
        }
    }

    public sealed class ManagedZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        /// </summary>
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("cloudLoggingConfig")]
        public Input<Inputs.ManagedZoneCloudLoggingConfigArgs>? CloudLoggingConfig { get; set; }

        /// <summary>
        /// The time that this resource was created on the server. This is in RFC3339 text format. Output only.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the managed zone's function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The DNS name of this managed zone, for instance "example.com.".
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// DNSSEC configuration.
        /// </summary>
        [Input("dnssecConfig")]
        public Input<Inputs.ManagedZoneDnsSecConfigArgs>? DnssecConfig { get; set; }

        /// <summary>
        /// The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this field contains the set of destinations to forward to.
        /// </summary>
        [Input("forwardingConfig")]
        public Input<Inputs.ManagedZoneForwardingConfigArgs>? ForwardingConfig { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// User assigned name for this resource. Must be unique within the project. The name must be 1-63 characters long, must begin with a letter, end with a letter or digit, and only contain lowercase letters, digits or dashes.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optionally specifies the NameServerSet for this ManagedZone. A NameServerSet is a set of DNS name servers that all host the same ManagedZones. Most users leave this field unset. If you need to use this field, contact your account team.
        /// </summary>
        [Input("nameServerSet")]
        public Input<string>? NameServerSet { get; set; }

        /// <summary>
        /// The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field contains the network to peer with.
        /// </summary>
        [Input("peeringConfig")]
        public Input<Inputs.ManagedZonePeeringConfigArgs>? PeeringConfig { get; set; }

        /// <summary>
        /// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        /// </summary>
        [Input("privateVisibilityConfig")]
        public Input<Inputs.ManagedZonePrivateVisibilityConfigArgs>? PrivateVisibilityConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The presence of this field indicates that this is a managed reverse lookup zone and Cloud DNS resolves reverse lookup queries using automatically configured records for VPC resources. This only applies to networks listed under private_visibility_config.
        /// </summary>
        [Input("reverseLookupConfig")]
        public Input<Inputs.ManagedZoneReverseLookupConfigArgs>? ReverseLookupConfig { get; set; }

        /// <summary>
        /// This field links to the associated service directory namespace. Do not set this field for public zones or forwarding zones.
        /// </summary>
        [Input("serviceDirectoryConfig")]
        public Input<Inputs.ManagedZoneServiceDirectoryConfigArgs>? ServiceDirectoryConfig { get; set; }

        /// <summary>
        /// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.GoogleNative.DNS.V2.ManagedZoneVisibility>? Visibility { get; set; }

        public ManagedZoneArgs()
        {
        }
        public static new ManagedZoneArgs Empty => new ManagedZoneArgs();
    }
}
