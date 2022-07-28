// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3
{
    /// <summary>
    /// Creates a PerfSampleSeries. May return any of the following error code(s): - ALREADY_EXISTS - PerfMetricSummary already exists for the given Step - NOT_FOUND - The containing Step does not exist
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:toolresults/v1beta3:PerfSampleSeries")]
    public partial class PerfSampleSeries : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Basic series represented by a line chart
        /// </summary>
        [Output("basicPerfSampleSeries")]
        public Output<Outputs.BasicPerfSampleSeriesResponse> BasicPerfSampleSeries { get; private set; } = null!;

        [Output("executionId")]
        public Output<string> ExecutionId { get; private set; } = null!;

        [Output("historyId")]
        public Output<string> HistoryId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A sample series id 
        /// </summary>
        [Output("sampleSeriesId")]
        public Output<string> SampleSeriesId { get; private set; } = null!;

        [Output("stepId")]
        public Output<string> StepId { get; private set; } = null!;


        /// <summary>
        /// Create a PerfSampleSeries resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PerfSampleSeries(string name, PerfSampleSeriesArgs args, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:PerfSampleSeries", name, args ?? new PerfSampleSeriesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PerfSampleSeries(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:PerfSampleSeries", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "executionId",
                    "historyId",
                    "project",
                    "stepId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PerfSampleSeries resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PerfSampleSeries Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PerfSampleSeries(name, id, options);
        }
    }

    public sealed class PerfSampleSeriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic series represented by a line chart
        /// </summary>
        [Input("basicPerfSampleSeries")]
        public Input<Inputs.BasicPerfSampleSeriesArgs>? BasicPerfSampleSeries { get; set; }

        [Input("executionId", required: true)]
        public Input<string> ExecutionId { get; set; } = null!;

        [Input("historyId", required: true)]
        public Input<string> HistoryId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("stepId", required: true)]
        public Input<string> StepId { get; set; } = null!;

        public PerfSampleSeriesArgs()
        {
        }
        public static new PerfSampleSeriesArgs Empty => new PerfSampleSeriesArgs();
    }
}
