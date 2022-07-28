// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Inputs
{

    /// <summary>
    /// Configures a RRSetRoutingPolicy that routes in a weighted round robin fashion.
    /// </summary>
    public sealed class RRSetRoutingPolicyWrrPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("items")]
        private InputList<Inputs.RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs>? _items;
        public InputList<Inputs.RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.RRSetRoutingPolicyWrrPolicyWrrPolicyItemArgs>());
            set => _items = value;
        }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public RRSetRoutingPolicyWrrPolicyArgs()
        {
        }
        public static new RRSetRoutingPolicyWrrPolicyArgs Empty => new RRSetRoutingPolicyWrrPolicyArgs();
    }
}
