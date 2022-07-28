// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// VPC access connector specification.
    /// </summary>
    public sealed class VpcAccessConnectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The egress setting for the connector, controlling what traffic is diverted through it.
        /// </summary>
        [Input("egressSetting")]
        public Input<Pulumi.GoogleNative.AppEngine.V1Beta.VpcAccessConnectorEgressSetting>? EgressSetting { get; set; }

        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VpcAccessConnectorArgs()
        {
        }
        public static new VpcAccessConnectorArgs Empty => new VpcAccessConnectorArgs();
    }
}
