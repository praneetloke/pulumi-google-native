// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// TaskTemplate describes the data a task should have when created from a template.
    /// </summary>
    public sealed class GoogleCloudRunV2TaskTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("containers")]
        private InputList<Inputs.GoogleCloudRunV2ContainerArgs>? _containers;

        /// <summary>
        /// Holds the single container that defines the unit of execution for this task.
        /// </summary>
        public InputList<Inputs.GoogleCloudRunV2ContainerArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.GoogleCloudRunV2ContainerArgs>());
            set => _containers = value;
        }

        /// <summary>
        /// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// The execution environment being used to host this Task.
        /// </summary>
        [Input("executionEnvironment")]
        public Input<Pulumi.GoogleNative.Run.V2.GoogleCloudRunV2TaskTemplateExecutionEnvironment>? ExecutionEnvironment { get; set; }

        /// <summary>
        /// Number of retries allowed per Task, before marking this Task failed.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// Email address of the IAM service account associated with the Task of a Job. The service account represents the identity of the running task, and determines what permissions the task has. If not provided, the task will use the project's default service account.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        [Input("volumes")]
        private InputList<Inputs.GoogleCloudRunV2VolumeArgs>? _volumes;

        /// <summary>
        /// A list of Volumes to make available to containers.
        /// </summary>
        public InputList<Inputs.GoogleCloudRunV2VolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.GoogleCloudRunV2VolumeArgs>());
            set => _volumes = value;
        }

        /// <summary>
        /// VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
        /// </summary>
        [Input("vpcAccess")]
        public Input<Inputs.GoogleCloudRunV2VpcAccessArgs>? VpcAccess { get; set; }

        public GoogleCloudRunV2TaskTemplateArgs()
        {
        }
        public static new GoogleCloudRunV2TaskTemplateArgs Empty => new GoogleCloudRunV2TaskTemplateArgs();
    }
}
