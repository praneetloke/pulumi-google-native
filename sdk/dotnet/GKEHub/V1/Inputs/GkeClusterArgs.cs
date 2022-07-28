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
    /// GkeCluster contains information specific to GKE clusters.
    /// </summary>
    public sealed class GkeClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Self-link of the GCP resource for the GKE cluster. For example: //container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster Zonal clusters are also supported.
        /// </summary>
        [Input("resourceLink")]
        public Input<string>? ResourceLink { get; set; }

        public GkeClusterArgs()
        {
        }
        public static new GkeClusterArgs Empty => new GkeClusterArgs();
    }
}
