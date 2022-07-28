// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// The Dataproc cluster config for a cluster that does not directly control the underlying compute resources, such as a Dataproc-on-GKE cluster (https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke).
    /// </summary>
    public sealed class VirtualClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Configuration of auxiliary services used by this cluster.
        /// </summary>
        [Input("auxiliaryServicesConfig")]
        public Input<Inputs.AuxiliaryServicesConfigArgs>? AuxiliaryServicesConfig { get; set; }

        /// <summary>
        /// The configuration for running the Dataproc cluster on Kubernetes.
        /// </summary>
        [Input("kubernetesClusterConfig", required: true)]
        public Input<Inputs.KubernetesClusterConfigArgs> KubernetesClusterConfig { get; set; } = null!;

        /// <summary>
        /// Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see Dataproc staging and temp buckets (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). This field requires a Cloud Storage bucket name, not a gs://... URI to a Cloud Storage bucket.
        /// </summary>
        [Input("stagingBucket")]
        public Input<string>? StagingBucket { get; set; }

        public VirtualClusterConfigArgs()
        {
        }
        public static new VirtualClusterConfigArgs Empty => new VirtualClusterConfigArgs();
    }
}
