// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetQuery
    {
        /// <summary>
        /// Get query status If the query is still in progress, the `state` is set to "running" After the query has completed successfully, `state` is set to "completed"
        /// </summary>
        public static Task<GetQueryResult> InvokeAsync(GetQueryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueryResult>("google-native:apigee/v1:getQuery", args ?? new GetQueryArgs(), options.WithDefaults());

        /// <summary>
        /// Get query status If the query is still in progress, the `state` is set to "running" After the query has completed successfully, `state` is set to "completed"
        /// </summary>
        public static Output<GetQueryResult> Invoke(GetQueryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetQueryResult>("google-native:apigee/v1:getQuery", args ?? new GetQueryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueryArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("queryId", required: true)]
        public string QueryId { get; set; } = null!;

        public GetQueryArgs()
        {
        }
        public static new GetQueryArgs Empty => new GetQueryArgs();
    }

    public sealed class GetQueryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("queryId", required: true)]
        public Input<string> QueryId { get; set; } = null!;

        public GetQueryInvokeArgs()
        {
        }
        public static new GetQueryInvokeArgs Empty => new GetQueryInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueryResult
    {
        /// <summary>
        /// Creation time of the query.
        /// </summary>
        public readonly string Created;
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
        /// Asynchronous Query Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Contains information like metrics, dimenstions etc of the AsyncQuery.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1QueryMetadataResponse QueryParams;
        /// <summary>
        /// Asynchronous Report ID.
        /// </summary>
        public readonly string ReportDefinitionId;
        /// <summary>
        /// Result is available only after the query is completed.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1AsyncQueryResultResponse Result;
        /// <summary>
        /// ResultFileSize is available only after the query is completed.
        /// </summary>
        public readonly string ResultFileSize;
        /// <summary>
        /// ResultRows is available only after the query is completed.
        /// </summary>
        public readonly string ResultRows;
        /// <summary>
        /// Self link of the query. Example: `/organizations/myorg/environments/myenv/queries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostQueries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
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
        private GetQueryResult(
            string created,

            string envgroupHostname,

            string error,

            string executionTime,

            string name,

            Outputs.GoogleCloudApigeeV1QueryMetadataResponse queryParams,

            string reportDefinitionId,

            Outputs.GoogleCloudApigeeV1AsyncQueryResultResponse result,

            string resultFileSize,

            string resultRows,

            string self,

            string state,

            string updated)
        {
            Created = created;
            EnvgroupHostname = envgroupHostname;
            Error = error;
            ExecutionTime = executionTime;
            Name = name;
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
