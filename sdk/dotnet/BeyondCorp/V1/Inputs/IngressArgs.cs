// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1.Inputs
{

    /// <summary>
    /// Settings of how to connect to the ClientGateway. One of the following options should be set.
    /// </summary>
    public sealed class IngressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The basic ingress config for ClientGateways.
        /// </summary>
        [Input("config")]
        public Input<Inputs.ConfigArgs>? Config { get; set; }

        public IngressArgs()
        {
        }
        public static new IngressArgs Empty => new IngressArgs();
    }
}
