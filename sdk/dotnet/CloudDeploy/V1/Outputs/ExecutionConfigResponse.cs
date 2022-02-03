// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Configuration of the environment to use when calling Skaffold.
    /// </summary>
    [OutputType]
    public sealed class ExecutionConfigResponse
    {
        /// <summary>
        /// Optional. Cloud Storage location where execution outputs should be stored. This can either be a bucket ("gs://my-bucket") or a path within a bucket ("gs://my-bucket/my-dir"). If unspecified, a default bucket located in the same region will be used.
        /// </summary>
        public readonly string ArtifactStorage;
        /// <summary>
        /// Optional. Use default Cloud Build pool.
        /// </summary>
        public readonly Outputs.DefaultPoolResponse DefaultPool;
        /// <summary>
        /// Optional. Use private Cloud Build pool.
        /// </summary>
        public readonly Outputs.PrivatePoolResponse PrivatePool;
        /// <summary>
        /// Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) will be used.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Usages when this configuration should be applied.
        /// </summary>
        public readonly ImmutableArray<string> Usages;
        /// <summary>
        /// Optional. The resource name of the `WorkerPool`, with the format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. If this optional field is unspecified, the default Cloud Build pool will be used.
        /// </summary>
        public readonly string WorkerPool;

        [OutputConstructor]
        private ExecutionConfigResponse(
            string artifactStorage,

            Outputs.DefaultPoolResponse defaultPool,

            Outputs.PrivatePoolResponse privatePool,

            string serviceAccount,

            ImmutableArray<string> usages,

            string workerPool)
        {
            ArtifactStorage = artifactStorage;
            DefaultPool = defaultPool;
            PrivatePool = privatePool;
            ServiceAccount = serviceAccount;
            Usages = usages;
            WorkerPool = workerPool;
        }
    }
}
