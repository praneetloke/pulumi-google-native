// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Defines the IP ranges that want to use NAT for a subnetwork.
    /// </summary>
    public sealed class RouterNatSubnetworkToNatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL for the subnetwork resource that will use NAT.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secondaryIpRangeNames")]
        private InputList<string>? _secondaryIpRangeNames;

        /// <summary>
        /// A list of the secondary ranges of the Subnetwork that are allowed to use NAT. This can be populated only if "LIST_OF_SECONDARY_IP_RANGES" is one of the values in source_ip_ranges_to_nat.
        /// </summary>
        public InputList<string> SecondaryIpRangeNames
        {
            get => _secondaryIpRangeNames ?? (_secondaryIpRangeNames = new InputList<string>());
            set => _secondaryIpRangeNames = value;
        }

        [Input("sourceIpRangesToNat")]
        private InputList<Pulumi.GoogleNative.Compute.V1.RouterNatSubnetworkToNatSourceIpRangesToNatItem>? _sourceIpRangesToNat;

        /// <summary>
        /// Specify the options for NAT ranges in the Subnetwork. All options of a single value are valid except NAT_IP_RANGE_OPTION_UNSPECIFIED. The only valid option with multiple values is: ["PRIMARY_IP_RANGE", "LIST_OF_SECONDARY_IP_RANGES"] Default: [ALL_IP_RANGES]
        /// </summary>
        public InputList<Pulumi.GoogleNative.Compute.V1.RouterNatSubnetworkToNatSourceIpRangesToNatItem> SourceIpRangesToNat
        {
            get => _sourceIpRangesToNat ?? (_sourceIpRangesToNat = new InputList<Pulumi.GoogleNative.Compute.V1.RouterNatSubnetworkToNatSourceIpRangesToNatItem>());
            set => _sourceIpRangesToNat = value;
        }

        public RouterNatSubnetworkToNatArgs()
        {
        }
        public static new RouterNatSubnetworkToNatArgs Empty => new RouterNatSubnetworkToNatArgs();
    }
}
