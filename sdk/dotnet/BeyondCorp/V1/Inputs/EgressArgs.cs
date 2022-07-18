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
    /// The details of the egress info. One of the following options should be set.
    /// </summary>
    public sealed class EgressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A VPC from the consumer project.
        /// </summary>
        [Input("peeredVpc")]
        public Input<Inputs.PeeredVpcArgs>? PeeredVpc { get; set; }

        public EgressArgs()
        {
        }
    }
}
