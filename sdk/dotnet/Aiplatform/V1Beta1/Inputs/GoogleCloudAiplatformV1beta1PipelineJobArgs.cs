// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Inputs
{

    /// <summary>
    /// An instance of a machine learning PipelineJob.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1beta1PipelineJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the Pipeline. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Customer-managed encryption key spec for a pipelineJob. If set, this PipelineJob and all of its sub-resources will be secured by this key.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1EncryptionSpecArgs>? EncryptionSpec { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize PipelineJob. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. Note there is some reserved label key for Vertex AI Pipelines. - `vertex-ai-pipelines-run-billing-id`, user set value will get overrided.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The full name of the Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the Pipeline Job's workload should be peered. For example, `projects/12345/global/networks/myVPC`. [Format](/compute/docs/reference/rest/v1/networks/insert) is of the form `projects/{project}/global/networks/{network}`. Where {project} is a project number, as in `12345`, and {network} is a network name. Private services access must already be configured for the network. Pipeline job will apply the network configuration to the Google Cloud resources being launched, if applied, such as Vertex AI Training or Dataflow job. If left unspecified, the workload is not peered with any network.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("pipelineSpec")]
        private InputMap<object>? _pipelineSpec;

        /// <summary>
        /// The spec of the pipeline.
        /// </summary>
        public InputMap<object> PipelineSpec
        {
            get => _pipelineSpec ?? (_pipelineSpec = new InputMap<object>());
            set => _pipelineSpec = value;
        }

        [Input("reservedIpRanges")]
        private InputList<string>? _reservedIpRanges;

        /// <summary>
        /// A list of names for the reserved ip ranges under the VPC network that can be used for this Pipeline Job's workload. If set, we will deploy the Pipeline Job's workload within the provided ip ranges. Otherwise, the job will be deployed to any ip ranges under the provided VPC network. Example: ['vertex-ai-ip-range'].
        /// </summary>
        public InputList<string> ReservedIpRanges
        {
            get => _reservedIpRanges ?? (_reservedIpRanges = new InputList<string>());
            set => _reservedIpRanges = value;
        }

        /// <summary>
        /// Runtime config of the pipeline.
        /// </summary>
        [Input("runtimeConfig")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigArgs>? RuntimeConfig { get; set; }

        /// <summary>
        /// The service account that the pipeline workload runs as. If not specified, the Compute Engine default service account in the project will be used. See https://cloud.google.com/compute/docs/access/service-accounts#default_service_account Users starting the pipeline must have the `iam.serviceAccounts.actAs` permission on this service account.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// A template uri from where the PipelineJob.pipeline_spec, if empty, will be downloaded. Currently, only uri from Vertex Template Registry &amp; Gallery is supported. Reference to https://cloud.google.com/vertex-ai/docs/pipelines/create-pipeline-template.
        /// </summary>
        [Input("templateUri")]
        public Input<string>? TemplateUri { get; set; }

        public GoogleCloudAiplatformV1beta1PipelineJobArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1PipelineJobArgs Empty => new GoogleCloudAiplatformV1beta1PipelineJobArgs();
    }
}
