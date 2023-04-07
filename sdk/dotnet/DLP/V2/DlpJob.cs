// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    /// <summary>
    /// Creates a new job to inspect storage or calculate risk metrics. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more. When no InfoTypes or CustomInfoTypes are specified in inspect jobs, the system will automatically choose what detectors to run. By default this may be all types, but may change over time as detectors are updated.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dlp/v2:DlpJob")]
    public partial class DlpJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Events that should occur after the job has completed.
        /// </summary>
        [Output("actionDetails")]
        public Output<ImmutableArray<Outputs.GooglePrivacyDlpV2ActionDetailsResponse>> ActionDetails { get; private set; } = null!;

        /// <summary>
        /// Time when the job was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Time when the job finished.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// A stream of errors encountered running the job.
        /// </summary>
        [Output("errors")]
        public Output<ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse>> Errors { get; private set; } = null!;

        /// <summary>
        /// Results from inspecting a data source.
        /// </summary>
        [Output("inspectDetails")]
        public Output<Outputs.GooglePrivacyDlpV2InspectDataSourceDetailsResponse> InspectDetails { get; private set; } = null!;

        /// <summary>
        /// If created by a job trigger, the resource name of the trigger that instantiated the job.
        /// </summary>
        [Output("jobTriggerName")]
        public Output<string> JobTriggerName { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The server-assigned name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Results from analyzing risk of a data source.
        /// </summary>
        [Output("riskDetails")]
        public Output<Outputs.GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse> RiskDetails { get; private set; } = null!;

        /// <summary>
        /// Time when the job started.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// State of a job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of job.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DlpJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DlpJob(string name, DlpJobArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:DlpJob", name, args ?? new DlpJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DlpJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:DlpJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DlpJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DlpJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DlpJob(name, id, options);
        }
    }

    public sealed class DlpJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An inspection job scans a storage repository for InfoTypes.
        /// </summary>
        [Input("inspectJob")]
        public Input<Inputs.GooglePrivacyDlpV2InspectJobConfigArgs>? InspectJob { get; set; }

        /// <summary>
        /// The job id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// Deprecated. This field has no effect.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A risk analysis job calculates re-identification risk metrics for a BigQuery table.
        /// </summary>
        [Input("riskJob")]
        public Input<Inputs.GooglePrivacyDlpV2RiskAnalysisJobConfigArgs>? RiskJob { get; set; }

        public DlpJobArgs()
        {
        }
        public static new DlpJobArgs Empty => new DlpJobArgs();
    }
}
