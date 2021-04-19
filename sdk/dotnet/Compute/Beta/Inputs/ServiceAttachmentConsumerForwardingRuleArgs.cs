// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// [Output Only] A consumer forwarding rule connected to this service attachment. [Deprecated] Do not use.
    /// </summary>
    public sealed class ServiceAttachmentConsumerForwardingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The url of a consumer forwarding rule.
        /// </summary>
        [Input("forwardingRule")]
        public Input<string>? ForwardingRule { get; set; }

        /// <summary>
        /// The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        [Input("pscConnectionId")]
        public Input<string>? PscConnectionId { get; set; }

        /// <summary>
        /// The status of the forwarding rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ServiceAttachmentConsumerForwardingRuleArgs()
        {
        }
    }
}
