// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    /// <summary>
    /// A routing block which contains the routing information for one WRR item.
    /// </summary>
    public sealed class RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("rrdatas")]
        private InputList<string>? _rrdatas;
        public InputList<string> Rrdatas
        {
            get => _rrdatas ?? (_rrdatas = new InputList<string>());
            set => _rrdatas = value;
        }

        [Input("signatureRrdatas")]
        private InputList<string>? _signatureRrdatas;

        /// <summary>
        /// DNSSEC generated signatures for all the rrdata within this item. Note that if health checked targets are provided for DNSSEC enabled zones, there's a restriction of 1 ip per item. .
        /// </summary>
        public InputList<string> SignatureRrdatas
        {
            get => _signatureRrdatas ?? (_signatureRrdatas = new InputList<string>());
            set => _signatureRrdatas = value;
        }

        /// <summary>
        /// The weight corresponding to this subset of rrdata. When multiple WeightedRoundRobinPolicyItems are configured, the probability of returning an rrset is proportional to its weight relative to the sum of weights configured for all items. This weight should be non-negative.
        /// </summary>
        [Input("weight")]
        public Input<double>? Weight { get; set; }

        public RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs()
        {
        }
        public static new RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs Empty => new RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs();
    }
}
