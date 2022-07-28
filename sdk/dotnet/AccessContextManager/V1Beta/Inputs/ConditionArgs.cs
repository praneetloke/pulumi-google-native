// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta.Inputs
{

    /// <summary>
    /// A condition necessary for an `AccessLevel` to be granted. The Condition is an AND over its fields. So a Condition is true if: 1) the request IP is from one of the listed subnetworks AND 2) the originating device complies with the listed device policy AND 3) all listed access levels are granted AND 4) the request was sent at a time allowed by the DateTimeRestriction.
    /// </summary>
    public sealed class ConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device specific restrictions, all restrictions must hold for the Condition to be true. If not specified, all devices are allowed.
        /// </summary>
        [Input("devicePolicy")]
        public Input<Inputs.DevicePolicyArgs>? DevicePolicy { get; set; }

        [Input("ipSubnetworks")]
        private InputList<string>? _ipSubnetworks;

        /// <summary>
        /// CIDR block IP subnetwork specification. May be IPv4 or IPv6. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly, for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32" is not. The originating IP of a request must be in one of the listed subnets in order for this Condition to be true. If empty, all IP addresses are allowed.
        /// </summary>
        public InputList<string> IpSubnetworks
        {
            get => _ipSubnetworks ?? (_ipSubnetworks = new InputList<string>());
            set => _ipSubnetworks = value;
        }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// The request must be made by one of the provided user or service accounts. Groups are not supported. Syntax: `user:{emailid}` `serviceAccount:{emailid}` If not specified, a request may come from any user.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Whether to negate the Condition. If true, the Condition becomes a NAND over its non-empty fields, each field must be false for the Condition overall to be satisfied. Defaults to false.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// The request must originate from one of the provided countries/regions. Must be valid ISO 3166-1 alpha-2 codes.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("requiredAccessLevels")]
        private InputList<string>? _requiredAccessLevels;

        /// <summary>
        /// A list of other access levels defined in the same `Policy`, referenced by resource name. Referencing an `AccessLevel` which does not exist is an error. All access levels listed must be granted for the Condition to be true. Example: "`accessPolicies/MY_POLICY/accessLevels/LEVEL_NAME"`
        /// </summary>
        public InputList<string> RequiredAccessLevels
        {
            get => _requiredAccessLevels ?? (_requiredAccessLevels = new InputList<string>());
            set => _requiredAccessLevels = value;
        }

        public ConditionArgs()
        {
        }
        public static new ConditionArgs Empty => new ConditionArgs();
    }
}
