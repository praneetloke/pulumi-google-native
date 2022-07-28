// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetSecurityReport
    {
        /// <summary>
        /// Get security report status If the query is still in progress, the `state` is set to "running" After the query has completed successfully, `state` is set to "completed"
        /// </summary>
        public static Task<GetSecurityReportResult> InvokeAsync(GetSecurityReportArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityReportResult>("google-native:apigee/v1:getSecurityReport", args ?? new GetSecurityReportArgs(), options.WithDefaults());

        /// <summary>
        /// Get security report status If the query is still in progress, the `state` is set to "running" After the query has completed successfully, `state` is set to "completed"
        /// </summary>
        public static Output<GetSecurityReportResult> Invoke(GetSecurityReportInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecurityReportResult>("google-native:apigee/v1:getSecurityReport", args ?? new GetSecurityReportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityReportArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("securityReportId", required: true)]
        public string SecurityReportId { get; set; } = null!;

        public GetSecurityReportArgs()
        {
        }
        public static new GetSecurityReportArgs Empty => new GetSecurityReportArgs();
    }

    public sealed class GetSecurityReportInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("securityReportId", required: true)]
        public Input<string> SecurityReportId { get; set; } = null!;

        public GetSecurityReportInvokeArgs()
        {
        }
        public static new GetSecurityReportInvokeArgs Empty => new GetSecurityReportInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityReportResult
    {
        /// <summary>
        /// Creation time of the query.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// Display Name specified by the user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Hostname is available only when query is executed at host level.
        /// </summary>
        public readonly string EnvgroupHostname;
        /// <summary>
        /// Error is set when query fails.
        /// </summary>
        public readonly string Error;
        /// <summary>
        /// ExecutionTime is available only after the query is completed.
        /// </summary>
        public readonly string ExecutionTime;
        /// <summary>
        /// Contains information like metrics, dimenstions etc of the Security Report.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1SecurityReportMetadataResponse QueryParams;
        /// <summary>
        /// Report Definition ID.
        /// </summary>
        public readonly string ReportDefinitionId;
        /// <summary>
        /// Result is available only after the query is completed.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1SecurityReportResultMetadataResponse Result;
        /// <summary>
        /// ResultFileSize is available only after the query is completed.
        /// </summary>
        public readonly string ResultFileSize;
        /// <summary>
        /// ResultRows is available only after the query is completed.
        /// </summary>
        public readonly string ResultRows;
        /// <summary>
        /// Self link of the query. Example: `/organizations/myorg/environments/myenv/securityReports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostSecurityReports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
        /// </summary>
        public readonly string Self;
        /// <summary>
        /// Query state could be "enqueued", "running", "completed", "failed".
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Last updated timestamp for the query.
        /// </summary>
        public readonly string Updated;

        [OutputConstructor]
        private GetSecurityReportResult(
            string created,

            string displayName,

            string envgroupHostname,

            string error,

            string executionTime,

            Outputs.GoogleCloudApigeeV1SecurityReportMetadataResponse queryParams,

            string reportDefinitionId,

            Outputs.GoogleCloudApigeeV1SecurityReportResultMetadataResponse result,

            string resultFileSize,

            string resultRows,

            string self,

            string state,

            string updated)
        {
            Created = created;
            DisplayName = displayName;
            EnvgroupHostname = envgroupHostname;
            Error = error;
            ExecutionTime = executionTime;
            QueryParams = queryParams;
            ReportDefinitionId = reportDefinitionId;
            Result = result;
            ResultFileSize = resultFileSize;
            ResultRows = resultRows;
            Self = self;
            State = state;
            Updated = updated;
        }
    }
}
