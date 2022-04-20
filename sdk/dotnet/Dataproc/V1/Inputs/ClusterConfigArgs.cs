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
    /// The cluster config.
    /// </summary>
    public sealed class ClusterConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.
        /// </summary>
        [Input("autoscalingConfig")]
        public Input<Inputs.AutoscalingConfigArgs>? AutoscalingConfig { get; set; }

        /// <summary>
        /// Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see Dataproc staging and temp buckets (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). This field requires a Cloud Storage bucket name, not a gs://... URI to a Cloud Storage bucket.
        /// </summary>
        [Input("configBucket")]
        public Input<string>? ConfigBucket { get; set; }

        /// <summary>
        /// Optional. The configuration(s) for a dataproc metric(s).
        /// </summary>
        [Input("dataprocMetricConfig")]
        public Input<Inputs.DataprocMetricConfigArgs>? DataprocMetricConfig { get; set; }

        /// <summary>
        /// Optional. Encryption settings for the cluster.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Optional. Port/endpoint configuration for this cluster
        /// </summary>
        [Input("endpointConfig")]
        public Input<Inputs.EndpointConfigArgs>? EndpointConfig { get; set; }

        /// <summary>
        /// Optional. The shared Compute Engine config settings for all instances in a cluster.
        /// </summary>
        [Input("gceClusterConfig")]
        public Input<Inputs.GceClusterConfigArgs>? GceClusterConfig { get; set; }

        /// <summary>
        /// Optional. BETA. The Kubernetes Engine config for Dataproc clusters deployed to Kubernetes. Setting this is considered mutually exclusive with Compute Engine-based options such as gce_cluster_config, master_config, worker_config, secondary_worker_config, and autoscaling_config.
        /// </summary>
        [Input("gkeClusterConfig")]
        public Input<Inputs.GkeClusterConfigArgs>? GkeClusterConfig { get; set; }

        [Input("initializationActions")]
        private InputList<Inputs.NodeInitializationActionArgs>? _initializationActions;

        /// <summary>
        /// Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's role metadata to run an executable on a master or worker node, as shown below using curl (you can also use wget): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if [[ "${ROLE}" == 'Master' ]]; then ... master specific actions ... else ... worker specific actions ... fi 
        /// </summary>
        public InputList<Inputs.NodeInitializationActionArgs> InitializationActions
        {
            get => _initializationActions ?? (_initializationActions = new InputList<Inputs.NodeInitializationActionArgs>());
            set => _initializationActions = value;
        }

        /// <summary>
        /// Optional. Lifecycle setting for the cluster.
        /// </summary>
        [Input("lifecycleConfig")]
        public Input<Inputs.LifecycleConfigArgs>? LifecycleConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for the cluster's master instance.
        /// </summary>
        [Input("masterConfig")]
        public Input<Inputs.InstanceGroupConfigArgs>? MasterConfig { get; set; }

        /// <summary>
        /// Optional. Metastore configuration.
        /// </summary>
        [Input("metastoreConfig")]
        public Input<Inputs.MetastoreConfigArgs>? MetastoreConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for a cluster's secondary worker instances
        /// </summary>
        [Input("secondaryWorkerConfig")]
        public Input<Inputs.InstanceGroupConfigArgs>? SecondaryWorkerConfig { get; set; }

        /// <summary>
        /// Optional. Security settings for the cluster.
        /// </summary>
        [Input("securityConfig")]
        public Input<Inputs.SecurityConfigArgs>? SecurityConfig { get; set; }

        /// <summary>
        /// Optional. The config settings for cluster software.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.SoftwareConfigArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket (see Dataproc staging and temp buckets (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). This field requires a Cloud Storage bucket name, not a gs://... URI to a Cloud Storage bucket.
        /// </summary>
        [Input("tempBucket")]
        public Input<string>? TempBucket { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for the cluster's worker instances.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.InstanceGroupConfigArgs>? WorkerConfig { get; set; }

        public ClusterConfigArgs()
        {
        }
    }
}
