// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Execution related settings, like retry and service_account.
    /// </summary>
    public sealed class GoogleCloudDataplexV1TaskExecutionSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputMap<string>? _args;

        /// <summary>
        /// Optional. The arguments to pass to the task. The args can use placeholders of the format ${placeholder} as part of key/value string. These will be interpolated before passing the args to the driver. Currently supported placeholders: - ${task_id} - ${job_time} To pass positional args, set the key as TASK_ARGS. The value should be a comma-separated string of all the positional arguments. To use a delimiter other than comma, refer to https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of other keys being present in the args, then TASK_ARGS will be passed as the last argument.
        /// </summary>
        public InputMap<string> Args
        {
            get => _args ?? (_args = new InputMap<string>());
            set => _args = value;
        }

        /// <summary>
        /// Optional. The Cloud KMS key to use for encryption, of the form: projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}.
        /// </summary>
        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        /// <summary>
        /// Optional. The maximum duration after which the job execution is expired.
        /// </summary>
        [Input("maxJobExecutionLifetime")]
        public Input<string>? MaxJobExecutionLifetime { get; set; }

        /// <summary>
        /// Optional. The project in which jobs are run. By default, the project containing the Lake is used. If a project is provided, the executionspec.service_account must belong to this same project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Service account to use to execute a task. If not provided, the default Compute service account for the project is used.
        /// </summary>
        [Input("serviceAccount", required: true)]
        public Input<string> ServiceAccount { get; set; } = null!;

        public GoogleCloudDataplexV1TaskExecutionSpecArgs()
        {
        }
        public static new GoogleCloudDataplexV1TaskExecutionSpecArgs Empty => new GoogleCloudDataplexV1TaskExecutionSpecArgs();
    }
}
