// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    public static class GetDlpJob
    {
        /// <summary>
        /// Gets the latest state of a long-running DlpJob. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more.
        /// </summary>
        public static Task<GetDlpJobResult> InvokeAsync(GetDlpJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDlpJobResult>("google-native:dlp/v2:getDlpJob", args ?? new GetDlpJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the latest state of a long-running DlpJob. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more.
        /// </summary>
        public static Output<GetDlpJobResult> Invoke(GetDlpJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDlpJobResult>("google-native:dlp/v2:getDlpJob", args ?? new GetDlpJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDlpJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("dlpJobId", required: true)]
        public string DlpJobId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDlpJobArgs()
        {
        }
        public static new GetDlpJobArgs Empty => new GetDlpJobArgs();
    }

    public sealed class GetDlpJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dlpJobId", required: true)]
        public Input<string> DlpJobId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDlpJobInvokeArgs()
        {
        }
        public static new GetDlpJobInvokeArgs Empty => new GetDlpJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetDlpJobResult
    {
        /// <summary>
        /// Time when the job was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Time when the job finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// A stream of errors encountered running the job.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse> Errors;
        /// <summary>
        /// Results from inspecting a data source.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2InspectDataSourceDetailsResponse InspectDetails;
        /// <summary>
        /// If created by a job trigger, the resource name of the trigger that instantiated the job.
        /// </summary>
        public readonly string JobTriggerName;
        /// <summary>
        /// The server-assigned name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Results from analyzing risk of a data source.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse RiskDetails;
        /// <summary>
        /// Time when the job started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// State of a job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of job.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDlpJobResult(
            string createTime,

            string endTime,

            ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse> errors,

            Outputs.GooglePrivacyDlpV2InspectDataSourceDetailsResponse inspectDetails,

            string jobTriggerName,

            string name,

            Outputs.GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse riskDetails,

            string startTime,

            string state,

            string type)
        {
            CreateTime = createTime;
            EndTime = endTime;
            Errors = errors;
            InspectDetails = inspectDetails;
            JobTriggerName = jobTriggerName;
            Name = name;
            RiskDetails = riskDetails;
            StartTime = startTime;
            State = state;
            Type = type;
        }
    }
}
