// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    public sealed class RRSetRoutingPolicyGeoPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("items")]
        private InputList<Inputs.RRSetRoutingPolicyGeoPolicyGeoPolicyItemArgs>? _items;

        /// <summary>
        /// The primary geo routing configuration. If there are multiple items with the same location, an error is returned instead.
        /// </summary>
        public InputList<Inputs.RRSetRoutingPolicyGeoPolicyGeoPolicyItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.RRSetRoutingPolicyGeoPolicyGeoPolicyItemArgs>());
            set => _items = value;
        }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public RRSetRoutingPolicyGeoPolicyArgs()
        {
        }
    }
}
