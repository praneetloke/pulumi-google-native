// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Inputs
{

    /// <summary>
    /// The description a notebook execution workload.
    /// </summary>
    public sealed class ExecutionTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration (count and accelerator type) for hardware running notebook execution.
        /// </summary>
        [Input("acceleratorConfig")]
        public Input<Inputs.SchedulerAcceleratorConfigArgs>? AcceleratorConfig { get; set; }

        /// <summary>
        /// Container Image URI to a DLVM Example: 'gcr.io/deeplearning-platform-release/base-cu100' More examples can be found at: https://cloud.google.com/ai-platform/deep-learning-containers/docs/choosing-container
        /// </summary>
        [Input("containerImageUri")]
        public Input<string>? ContainerImageUri { get; set; }

        /// <summary>
        /// Parameters used in Dataproc JobType executions.
        /// </summary>
        [Input("dataprocParameters")]
        public Input<Inputs.DataprocParametersArgs>? DataprocParameters { get; set; }

        /// <summary>
        /// Path to the notebook file to execute. Must be in a Google Cloud Storage bucket. Format: `gs://{bucket_name}/{folder}/{notebook_file_name}` Ex: `gs://notebook_user/scheduled_notebooks/sentiment_notebook.ipynb`
        /// </summary>
        [Input("inputNotebookFile")]
        public Input<string>? InputNotebookFile { get; set; }

        /// <summary>
        /// The type of Job to be used on this execution.
        /// </summary>
        [Input("jobType")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.ExecutionTemplateJobType>? JobType { get; set; }

        /// <summary>
        /// Name of the kernel spec to use. This must be specified if the kernel spec name on the execution target does not match the name in the input notebook file.
        /// </summary>
        [Input("kernelSpec")]
        public Input<string>? KernelSpec { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for execution. If execution is scheduled, a field included will be 'nbs-scheduled'. Otherwise, it is an immediate execution, and an included field will be 'nbs-immediate'. Use fields to efficiently index between various types of executions.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the type of virtual machine to use for your training job's master worker. You must specify this field when `scaleTier` is set to `CUSTOM`. You can use certain Compute Engine machine types directly in this field. The following types are supported: - `n1-standard-4` - `n1-standard-8` - `n1-standard-16` - `n1-standard-32` - `n1-standard-64` - `n1-standard-96` - `n1-highmem-2` - `n1-highmem-4` - `n1-highmem-8` - `n1-highmem-16` - `n1-highmem-32` - `n1-highmem-64` - `n1-highmem-96` - `n1-highcpu-16` - `n1-highcpu-32` - `n1-highcpu-64` - `n1-highcpu-96` Alternatively, you can use the following legacy machine types: - `standard` - `large_model` - `complex_model_s` - `complex_model_m` - `complex_model_l` - `standard_gpu` - `complex_model_m_gpu` - `complex_model_l_gpu` - `standard_p100` - `complex_model_m_p100` - `standard_v100` - `large_model_v100` - `complex_model_m_v100` - `complex_model_l_v100` Finally, if you want to use a TPU for training, specify `cloud_tpu` in this field. Learn more about the [special configuration options for training with TPU](https://cloud.google.com/ai-platform/training/docs/using-tpus#configuring_a_custom_tpu_machine).
        /// </summary>
        [Input("masterType")]
        public Input<string>? MasterType { get; set; }

        /// <summary>
        /// Path to the notebook folder to write to. Must be in a Google Cloud Storage bucket path. Format: `gs://{bucket_name}/{folder}` Ex: `gs://notebook_user/scheduled_notebooks`
        /// </summary>
        [Input("outputNotebookFolder")]
        public Input<string>? OutputNotebookFolder { get; set; }

        /// <summary>
        /// Parameters used within the 'input_notebook_file' notebook.
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// Parameters to be overridden in the notebook during execution. Ref https://papermill.readthedocs.io/en/latest/usage-parameterize.html on how to specifying parameters in the input notebook and pass them here in an YAML file. Ex: `gs://notebook_user/scheduled_notebooks/sentiment_notebook_params.yaml`
        /// </summary>
        [Input("paramsYamlFile")]
        public Input<string>? ParamsYamlFile { get; set; }

        /// <summary>
        /// Scale tier of the hardware used for notebook execution. DEPRECATED Will be discontinued. As right now only CUSTOM is supported.
        /// </summary>
        [Input("scaleTier", required: true)]
        public Input<Pulumi.GoogleNative.Notebooks.V1.ExecutionTemplateScaleTier> ScaleTier { get; set; } = null!;

        /// <summary>
        /// The email address of a service account to use when running the execution. You must have the `iam.serviceAccounts.actAs` permission for the specified service account.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// The name of a Vertex AI [Tensorboard] resource to which this execution will upload Tensorboard logs. Format: `projects/{project}/locations/{location}/tensorboards/{tensorboard}`
        /// </summary>
        [Input("tensorboard")]
        public Input<string>? Tensorboard { get; set; }

        /// <summary>
        /// Parameters used in Vertex AI JobType executions.
        /// </summary>
        [Input("vertexAiParameters")]
        public Input<Inputs.VertexAIParametersArgs>? VertexAiParameters { get; set; }

        public ExecutionTemplateArgs()
        {
        }
        public static new ExecutionTemplateArgs Empty => new ExecutionTemplateArgs();
    }
}
