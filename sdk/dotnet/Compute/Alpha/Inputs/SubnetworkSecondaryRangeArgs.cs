// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Represents a secondary IP range of a subnetwork.
    /// </summary>
    public sealed class SubnetworkSecondaryRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of IP addresses belonging to this subnetwork secondary range. Provide this property when you create the subnetwork. Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network. Only IPv4 is supported. The range can be any range listed in the Valid ranges list.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance. The name must be 1-63 characters long, and comply with RFC1035. The name must be unique within the subnetwork.
        /// </summary>
        [Input("rangeName")]
        public Input<string>? RangeName { get; set; }

        /// <summary>
        /// The URL of the reserved internal range.
        /// </summary>
        [Input("reservedInternalRange")]
        public Input<string>? ReservedInternalRange { get; set; }

        public SubnetworkSecondaryRangeArgs()
        {
        }
        public static new SubnetworkSecondaryRangeArgs Empty => new SubnetworkSecondaryRangeArgs();
    }
}
