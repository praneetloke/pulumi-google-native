// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta
{
    /// <summary>
    /// Create an OS Config patch deployment.
    /// </summary>
    [GoogleNativeResourceType("google-native:osconfig/v1beta:PatchDeployment")]
    public partial class PatchDeployment : Pulumi.CustomResource
    {
        /// <summary>
        /// Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Duration of the patch. After the duration ends, the patch times out.
        /// </summary>
        [Output("duration")]
        public Output<string> Duration { get; private set; } = null!;

        /// <summary>
        /// Required. VM instances to patch.
        /// </summary>
        [Output("instanceFilter")]
        public Output<Outputs.PatchInstanceFilterResponse> InstanceFilter { get; private set; } = null!;

        /// <summary>
        /// The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Output("lastExecuteTime")]
        public Output<string> LastExecuteTime { get; private set; } = null!;

        /// <summary>
        /// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. Schedule a one-time execution.
        /// </summary>
        [Output("oneTimeSchedule")]
        public Output<Outputs.OneTimeScheduleResponse> OneTimeSchedule { get; private set; } = null!;

        /// <summary>
        /// Optional. Patch configuration that is applied.
        /// </summary>
        [Output("patchConfig")]
        public Output<Outputs.PatchConfigResponse> PatchConfig { get; private set; } = null!;

        /// <summary>
        /// Required. Schedule recurring executions.
        /// </summary>
        [Output("recurringSchedule")]
        public Output<Outputs.RecurringScheduleResponse> RecurringSchedule { get; private set; } = null!;

        /// <summary>
        /// Optional. Rollout strategy of the patch job.
        /// </summary>
        [Output("rollout")]
        public Output<Outputs.PatchRolloutResponse> Rollout { get; private set; } = null!;

        /// <summary>
        /// Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a PatchDeployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PatchDeployment(string name, PatchDeploymentArgs args, CustomResourceOptions? options = null)
            : base("google-native:osconfig/v1beta:PatchDeployment", name, args ?? new PatchDeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PatchDeployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:osconfig/v1beta:PatchDeployment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PatchDeployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PatchDeployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PatchDeployment(name, id, options);
        }
    }

    public sealed class PatchDeploymentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Duration of the patch. After the duration ends, the patch times out.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Required. VM instances to patch.
        /// </summary>
        [Input("instanceFilter")]
        public Input<Inputs.PatchInstanceFilterArgs>? InstanceFilter { get; set; }

        /// <summary>
        /// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required. Schedule a one-time execution.
        /// </summary>
        [Input("oneTimeSchedule")]
        public Input<Inputs.OneTimeScheduleArgs>? OneTimeSchedule { get; set; }

        /// <summary>
        /// Optional. Patch configuration that is applied.
        /// </summary>
        [Input("patchConfig")]
        public Input<Inputs.PatchConfigArgs>? PatchConfig { get; set; }

        [Input("patchDeploymentId", required: true)]
        public Input<string> PatchDeploymentId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Required. Schedule recurring executions.
        /// </summary>
        [Input("recurringSchedule")]
        public Input<Inputs.RecurringScheduleArgs>? RecurringSchedule { get; set; }

        /// <summary>
        /// Optional. Rollout strategy of the patch job.
        /// </summary>
        [Input("rollout")]
        public Input<Inputs.PatchRolloutArgs>? Rollout { get; set; }

        public PatchDeploymentArgs()
        {
        }
    }
}
