// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    /// <summary>
    /// Creates new workflow template.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataproc/v1:WorkflowTemplate")]
    public partial class WorkflowTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time template was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        /// </summary>
        [Output("dagTimeout")]
        public Output<string> DagTimeout { get; private set; } = null!;

        /// <summary>
        /// The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        [Output("jobs")]
        public Output<ImmutableArray<Outputs.OrderedJobResponse>> Jobs { get; private set; } = null!;

        /// <summary>
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.TemplateParameterResponse>> Parameters { get; private set; } = null!;

        /// <summary>
        /// WorkflowTemplate scheduling information.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.WorkflowTemplatePlacementResponse> Placement { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The time template was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowTemplate(string name, WorkflowTemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1:WorkflowTemplate", name, args ?? new WorkflowTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataproc/v1:WorkflowTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkflowTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkflowTemplate(name, id, options);
        }
    }

    public sealed class WorkflowTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        /// </summary>
        [Input("dagTimeout")]
        public Input<string>? DagTimeout { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("jobs", required: true)]
        private InputList<Inputs.OrderedJobArgs>? _jobs;

        /// <summary>
        /// The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        public InputList<Inputs.OrderedJobArgs> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<Inputs.OrderedJobArgs>());
            set => _jobs = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("parameters")]
        private InputList<Inputs.TemplateParameterArgs>? _parameters;

        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        public InputList<Inputs.TemplateParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.TemplateParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// WorkflowTemplate scheduling information.
        /// </summary>
        [Input("placement", required: true)]
        public Input<Inputs.WorkflowTemplatePlacementArgs> Placement { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public WorkflowTemplateArgs()
        {
        }
        public static new WorkflowTemplateArgs Empty => new WorkflowTemplateArgs();
    }
}
