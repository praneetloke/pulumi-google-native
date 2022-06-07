// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// ApplianceCluster contains information specific to GDC Edge Appliance Clusters.
    /// </summary>
    public sealed class ApplianceClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the Appliance Cluster. For example: //transferappliance.googleapis.com/projects/my-project/locations/us-west1-a/appliances/my-appliance
        /// </summary>
        [Input("resourceLink")]
        public Input<string>? ResourceLink { get; set; }

        public ApplianceClusterArgs()
        {
        }
    }
}
