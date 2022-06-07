// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha2.Outputs
{

    /// <summary>
    /// ApplianceCluster contains information specific to GDC Edge Appliance Clusters.
    /// </summary>
    [OutputType]
    public sealed class ApplianceClusterResponse
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the Appliance Cluster. For example: //transferappliance.googleapis.com/projects/my-project/locations/us-west1-a/appliances/my-appliance
        /// </summary>
        public readonly string ResourceLink;

        [OutputConstructor]
        private ApplianceClusterResponse(string resourceLink)
        {
            ResourceLink = resourceLink;
        }
    }
}
