// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha2.Inputs
{

    /// <summary>
    /// MultiCloudCluster contains information specific to GKE Multi-Cloud clusters.
    /// </summary>
    public sealed class MultiCloudClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the GKE Multi-Cloud cluster. For example: //gkemulticloud.googleapis.com/projects/my-project/locations/us-west1-a/awsClusters/my-cluster //gkemulticloud.googleapis.com/projects/my-project/locations/us-west1-a/azureClusters/my-cluster
        /// </summary>
        [Input("resourceLink")]
        public Input<string>? ResourceLink { get; set; }

        public MultiCloudClusterArgs()
        {
        }
        public static new MultiCloudClusterArgs Empty => new MultiCloudClusterArgs();
    }
}
