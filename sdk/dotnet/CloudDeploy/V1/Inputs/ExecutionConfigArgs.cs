// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Inputs
{

    /// <summary>
    /// Configuration of the environment to use when calling Skaffold.
    /// </summary>
    public sealed class ExecutionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Cloud Storage location in which to store execution outputs. This can either be a bucket ("gs://my-bucket") or a path within a bucket ("gs://my-bucket/my-dir"). If unspecified, a default bucket located in the same region will be used.
        /// </summary>
        [Input("artifactStorage")]
        public Input<string>? ArtifactStorage { get; set; }

        /// <summary>
        /// Optional. Use default Cloud Build pool.
        /// </summary>
        [Input("defaultPool")]
        public Input<Inputs.DefaultPoolArgs>? DefaultPool { get; set; }

        /// <summary>
        /// Optional. Execution timeout for a Cloud Build Execution. This must be between 10m and 24h in seconds format. If unspecified, a default timeout of 1h is used.
        /// </summary>
        [Input("executionTimeout")]
        public Input<string>? ExecutionTimeout { get; set; }

        /// <summary>
        /// Optional. Use private Cloud Build pool.
        /// </summary>
        [Input("privatePool")]
        public Input<Inputs.PrivatePoolArgs>? PrivatePool { get; set; }

        /// <summary>
        /// Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) is used.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("usages", required: true)]
        private InputList<Pulumi.GoogleNative.CloudDeploy.V1.ExecutionConfigUsagesItem>? _usages;

        /// <summary>
        /// Usages when this configuration should be applied.
        /// </summary>
        public InputList<Pulumi.GoogleNative.CloudDeploy.V1.ExecutionConfigUsagesItem> Usages
        {
            get => _usages ?? (_usages = new InputList<Pulumi.GoogleNative.CloudDeploy.V1.ExecutionConfigUsagesItem>());
            set => _usages = value;
        }

        /// <summary>
        /// Optional. The resource name of the `WorkerPool`, with the format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. If this optional field is unspecified, the default Cloud Build pool will be used.
        /// </summary>
        [Input("workerPool")]
        public Input<string>? WorkerPool { get; set; }

        public ExecutionConfigArgs()
        {
        }
        public static new ExecutionConfigArgs Empty => new ExecutionConfigArgs();
    }
}
