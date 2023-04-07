// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Outputs
{

    /// <summary>
    /// TaskTemplate describes the data a task should have when created from a template.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRunV2TaskTemplateResponse
    {
        /// <summary>
        /// Holds the single container that defines the unit of execution for this task.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRunV2ContainerResponse> Containers;
        /// <summary>
        /// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
        /// </summary>
        public readonly string EncryptionKey;
        /// <summary>
        /// The execution environment being used to host this Task.
        /// </summary>
        public readonly string ExecutionEnvironment;
        /// <summary>
        /// Number of retries allowed per Task, before marking this Task failed. Defaults to 3.
        /// </summary>
        public readonly int MaxRetries;
        /// <summary>
        /// Email address of the IAM service account associated with the Task of a Job. The service account represents the identity of the running task, and determines what permissions the task has. If not provided, the task will use the project's default service account.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout. Defaults to 600 seconds.
        /// </summary>
        public readonly string Timeout;
        /// <summary>
        /// A list of Volumes to make available to containers.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRunV2VolumeResponse> Volumes;
        /// <summary>
        /// VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2VpcAccessResponse VpcAccess;

        [OutputConstructor]
        private GoogleCloudRunV2TaskTemplateResponse(
            ImmutableArray<Outputs.GoogleCloudRunV2ContainerResponse> containers,

            string encryptionKey,

            string executionEnvironment,

            int maxRetries,

            string serviceAccount,

            string timeout,

            ImmutableArray<Outputs.GoogleCloudRunV2VolumeResponse> volumes,

            Outputs.GoogleCloudRunV2VpcAccessResponse vpcAccess)
        {
            Containers = containers;
            EncryptionKey = encryptionKey;
            ExecutionEnvironment = executionEnvironment;
            MaxRetries = maxRetries;
            ServiceAccount = serviceAccount;
            Timeout = timeout;
            Volumes = volumes;
            VpcAccess = vpcAccess;
        }
    }
}
