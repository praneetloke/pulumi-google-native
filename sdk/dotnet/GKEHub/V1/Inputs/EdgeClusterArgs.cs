// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Inputs
{

    /// <summary>
    /// EdgeCluster contains information specific to Google Edge Clusters.
    /// </summary>
    public sealed class EdgeClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the Edge Cluster. For example: //edgecontainer.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster
        /// </summary>
        [Input("resourceLink")]
        public Input<string>? ResourceLink { get; set; }

        public EdgeClusterArgs()
        {
        }
        public static new EdgeClusterArgs Empty => new EdgeClusterArgs();
    }
}
