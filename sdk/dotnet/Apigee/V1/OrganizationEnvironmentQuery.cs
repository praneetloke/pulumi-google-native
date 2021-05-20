// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Submit a query to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationEnvironmentQuery")]
    public partial class OrganizationEnvironmentQuery : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of the query.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Hostname is available only when query is executed at host level.
        /// </summary>
        [Output("envgroupHostname")]
        public Output<string> EnvgroupHostname { get; private set; } = null!;

        /// <summary>
        /// Error is set when query fails.
        /// </summary>
        [Output("error")]
        public Output<string> Error { get; private set; } = null!;

        /// <summary>
        /// ExecutionTime is available only after the query is completed.
        /// </summary>
        [Output("executionTime")]
        public Output<string> ExecutionTime { get; private set; } = null!;

        /// <summary>
        /// Asynchronous Query Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Contains information like metrics, dimenstions etc of the AsyncQuery.
        /// </summary>
        [Output("queryParams")]
        public Output<Outputs.GoogleCloudApigeeV1QueryMetadataResponse> QueryParams { get; private set; } = null!;

        /// <summary>
        /// Asynchronous Report ID.
        /// </summary>
        [Output("reportDefinitionId")]
        public Output<string> ReportDefinitionId { get; private set; } = null!;

        /// <summary>
        /// Result is available only after the query is completed.
        /// </summary>
        [Output("result")]
        public Output<Outputs.GoogleCloudApigeeV1AsyncQueryResultResponse> Result { get; private set; } = null!;

        /// <summary>
        /// ResultFileSize is available only after the query is completed.
        /// </summary>
        [Output("resultFileSize")]
        public Output<string> ResultFileSize { get; private set; } = null!;

        /// <summary>
        /// ResultRows is available only after the query is completed.
        /// </summary>
        [Output("resultRows")]
        public Output<string> ResultRows { get; private set; } = null!;

        /// <summary>
        /// Self link of the query. Example: `/organizations/myorg/environments/myenv/queries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostQueries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
        /// </summary>
        [Output("self")]
        public Output<string> Self { get; private set; } = null!;

        /// <summary>
        /// Query state could be "enqueued", "running", "completed", "failed".
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Last updated timestamp for the query.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationEnvironmentQuery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationEnvironmentQuery(string name, OrganizationEnvironmentQueryArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentQuery", name, args ?? new OrganizationEnvironmentQueryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationEnvironmentQuery(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationEnvironmentQuery", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationEnvironmentQuery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationEnvironmentQuery Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationEnvironmentQuery(name, id, options);
        }
    }

    public sealed class OrganizationEnvironmentQueryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
        /// </summary>
        [Input("csvDelimiter")]
        public Input<string>? CsvDelimiter { get; set; }

        [Input("dimensions")]
        private InputList<string>? _dimensions;

        /// <summary>
        /// A list of dimensions. https://docs.apigee.com/api-platform/analytics/analytics-reference#dimensions
        /// </summary>
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Hostname needs to be specified if query intends to run at host level. This field is only allowed when query is submitted by CreateHostAsyncQuery where analytics data will be grouped by organization and hostname.
        /// </summary>
        [Input("envgroupHostname")]
        public Input<string>? EnvgroupHostname { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Boolean expression that can be used to filter data. Filter expressions can be combined using AND/OR terms and should be fully parenthesized to avoid ambiguity. See Analytics metrics, dimensions, and filters reference https://docs.apigee.com/api-platform/analytics/analytics-reference for more information on the fields available to filter on. For more information on the tokens that you use to build filter expressions, see Filter expression syntax. https://docs.apigee.com/api-platform/analytics/asynch-reports-api#filter-expression-syntax
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Time unit used to group the result set. Valid values include: second, minute, hour, day, week, or month. If a query includes groupByTimeUnit, then the result is an aggregation based on the specified time unit and the resultant timestamp does not include milliseconds precision. If a query omits groupByTimeUnit, then the resultant timestamp includes milliseconds precision.
        /// </summary>
        [Input("groupByTimeUnit")]
        public Input<string>? GroupByTimeUnit { get; set; }

        /// <summary>
        /// Maximum number of rows that can be returned in the result.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        [Input("metrics")]
        private InputList<Inputs.GoogleCloudApigeeV1QueryMetricArgs>? _metrics;

        /// <summary>
        /// A list of Metrics.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1QueryMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GoogleCloudApigeeV1QueryMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// Asynchronous Query Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the csvDelimiter property.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        [Input("queryId", required: true)]
        public Input<string> QueryId { get; set; } = null!;

        /// <summary>
        /// Asynchronous Report ID.
        /// </summary>
        [Input("reportDefinitionId")]
        public Input<string>? ReportDefinitionId { get; set; }

        /// <summary>
        /// Required. Time range for the query. Can use the following predefined strings to specify the time range: `last60minutes` `last24hours` `last7days` Or, specify the timeRange as a structure describing start and end timestamps in the ISO format: yyyy-mm-ddThh:mm:ssZ. Example: "timeRange": { "start": "2018-07-29T00:13:00Z", "end": "2018-08-01T00:18:00Z" }
        /// </summary>
        [Input("timeRange")]
        public Input<object>? TimeRange { get; set; }

        public OrganizationEnvironmentQueryArgs()
        {
        }
    }
}
