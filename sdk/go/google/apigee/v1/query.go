// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Submit a query to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Query struct {
	pulumi.CustomResourceState

	// Creation time of the query.
	Created pulumi.StringOutput `pulumi:"created"`
	// Hostname is available only when query is executed at host level.
	EnvgroupHostname pulumi.StringOutput `pulumi:"envgroupHostname"`
	EnvironmentId    pulumi.StringOutput `pulumi:"environmentId"`
	// Error is set when query fails.
	Error pulumi.StringOutput `pulumi:"error"`
	// ExecutionTime is available only after the query is completed.
	ExecutionTime pulumi.StringOutput `pulumi:"executionTime"`
	// Asynchronous Query Name.
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Contains information like metrics, dimenstions etc of the AsyncQuery.
	QueryParams GoogleCloudApigeeV1QueryMetadataResponseOutput `pulumi:"queryParams"`
	// Asynchronous Report ID.
	ReportDefinitionId pulumi.StringOutput `pulumi:"reportDefinitionId"`
	// Result is available only after the query is completed.
	Result GoogleCloudApigeeV1AsyncQueryResultResponseOutput `pulumi:"result"`
	// ResultFileSize is available only after the query is completed.
	ResultFileSize pulumi.StringOutput `pulumi:"resultFileSize"`
	// ResultRows is available only after the query is completed.
	ResultRows pulumi.StringOutput `pulumi:"resultRows"`
	// Self link of the query. Example: `/organizations/myorg/environments/myenv/queries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostQueries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
	Self pulumi.StringOutput `pulumi:"self"`
	// Query state could be "enqueued", "running", "completed", "failed".
	State pulumi.StringOutput `pulumi:"state"`
	// Last updated timestamp for the query.
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewQuery registers a new resource with the given unique name, arguments, and options.
func NewQuery(ctx *pulumi.Context,
	name string, args *QueryArgs, opts ...pulumi.ResourceOption) (*Query, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.TimeRange == nil {
		return nil, errors.New("invalid value for required argument 'TimeRange'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"environmentId",
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Query
	err := ctx.RegisterResource("google-native:apigee/v1:Query", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuery gets an existing Query resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryState, opts ...pulumi.ResourceOption) (*Query, error) {
	var resource Query
	err := ctx.ReadResource("google-native:apigee/v1:Query", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Query resources.
type queryState struct {
}

type QueryState struct {
}

func (QueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryState)(nil)).Elem()
}

type queryArgs struct {
	// Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
	CsvDelimiter *string `pulumi:"csvDelimiter"`
	// A list of dimensions. https://docs.apigee.com/api-platform/analytics/analytics-reference#dimensions
	Dimensions []string `pulumi:"dimensions"`
	// Hostname needs to be specified if query intends to run at host level. This field is only allowed when query is submitted by CreateHostAsyncQuery where analytics data will be grouped by organization and hostname.
	EnvgroupHostname *string `pulumi:"envgroupHostname"`
	EnvironmentId    string  `pulumi:"environmentId"`
	// Boolean expression that can be used to filter data. Filter expressions can be combined using AND/OR terms and should be fully parenthesized to avoid ambiguity. See Analytics metrics, dimensions, and filters reference https://docs.apigee.com/api-platform/analytics/analytics-reference for more information on the fields available to filter on. For more information on the tokens that you use to build filter expressions, see Filter expression syntax. https://docs.apigee.com/api-platform/analytics/asynch-reports-api#filter-expression-syntax
	Filter *string `pulumi:"filter"`
	// Time unit used to group the result set. Valid values include: second, minute, hour, day, week, or month. If a query includes groupByTimeUnit, then the result is an aggregation based on the specified time unit and the resultant timestamp does not include milliseconds precision. If a query omits groupByTimeUnit, then the resultant timestamp includes milliseconds precision.
	GroupByTimeUnit *string `pulumi:"groupByTimeUnit"`
	// Maximum number of rows that can be returned in the result.
	Limit *int `pulumi:"limit"`
	// A list of Metrics.
	Metrics []GoogleCloudApigeeV1QueryMetric `pulumi:"metrics"`
	// Asynchronous Query Name.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the csvDelimiter property.
	OutputFormat *string `pulumi:"outputFormat"`
	// Asynchronous Report ID.
	ReportDefinitionId *string `pulumi:"reportDefinitionId"`
	// Time range for the query. Can use the following predefined strings to specify the time range: `last60minutes` `last24hours` `last7days` Or, specify the timeRange as a structure describing start and end timestamps in the ISO format: yyyy-mm-ddThh:mm:ssZ. Example: "timeRange": { "start": "2018-07-29T00:13:00Z", "end": "2018-08-01T00:18:00Z" }
	TimeRange interface{} `pulumi:"timeRange"`
}

// The set of arguments for constructing a Query resource.
type QueryArgs struct {
	// Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
	CsvDelimiter pulumi.StringPtrInput
	// A list of dimensions. https://docs.apigee.com/api-platform/analytics/analytics-reference#dimensions
	Dimensions pulumi.StringArrayInput
	// Hostname needs to be specified if query intends to run at host level. This field is only allowed when query is submitted by CreateHostAsyncQuery where analytics data will be grouped by organization and hostname.
	EnvgroupHostname pulumi.StringPtrInput
	EnvironmentId    pulumi.StringInput
	// Boolean expression that can be used to filter data. Filter expressions can be combined using AND/OR terms and should be fully parenthesized to avoid ambiguity. See Analytics metrics, dimensions, and filters reference https://docs.apigee.com/api-platform/analytics/analytics-reference for more information on the fields available to filter on. For more information on the tokens that you use to build filter expressions, see Filter expression syntax. https://docs.apigee.com/api-platform/analytics/asynch-reports-api#filter-expression-syntax
	Filter pulumi.StringPtrInput
	// Time unit used to group the result set. Valid values include: second, minute, hour, day, week, or month. If a query includes groupByTimeUnit, then the result is an aggregation based on the specified time unit and the resultant timestamp does not include milliseconds precision. If a query omits groupByTimeUnit, then the resultant timestamp includes milliseconds precision.
	GroupByTimeUnit pulumi.StringPtrInput
	// Maximum number of rows that can be returned in the result.
	Limit pulumi.IntPtrInput
	// A list of Metrics.
	Metrics GoogleCloudApigeeV1QueryMetricArrayInput
	// Asynchronous Query Name.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the csvDelimiter property.
	OutputFormat pulumi.StringPtrInput
	// Asynchronous Report ID.
	ReportDefinitionId pulumi.StringPtrInput
	// Time range for the query. Can use the following predefined strings to specify the time range: `last60minutes` `last24hours` `last7days` Or, specify the timeRange as a structure describing start and end timestamps in the ISO format: yyyy-mm-ddThh:mm:ssZ. Example: "timeRange": { "start": "2018-07-29T00:13:00Z", "end": "2018-08-01T00:18:00Z" }
	TimeRange pulumi.Input
}

func (QueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryArgs)(nil)).Elem()
}

type QueryInput interface {
	pulumi.Input

	ToQueryOutput() QueryOutput
	ToQueryOutputWithContext(ctx context.Context) QueryOutput
}

func (*Query) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (i *Query) ToQueryOutput() QueryOutput {
	return i.ToQueryOutputWithContext(context.Background())
}

func (i *Query) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryOutput)
}

type QueryOutput struct{ *pulumi.OutputState }

func (QueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (o QueryOutput) ToQueryOutput() QueryOutput {
	return o
}

func (o QueryOutput) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return o
}

// Creation time of the query.
func (o QueryOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Hostname is available only when query is executed at host level.
func (o QueryOutput) EnvgroupHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.EnvgroupHostname }).(pulumi.StringOutput)
}

func (o QueryOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// Error is set when query fails.
func (o QueryOutput) Error() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Error }).(pulumi.StringOutput)
}

// ExecutionTime is available only after the query is completed.
func (o QueryOutput) ExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.ExecutionTime }).(pulumi.StringOutput)
}

// Asynchronous Query Name.
func (o QueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueryOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Contains information like metrics, dimenstions etc of the AsyncQuery.
func (o QueryOutput) QueryParams() GoogleCloudApigeeV1QueryMetadataResponseOutput {
	return o.ApplyT(func(v *Query) GoogleCloudApigeeV1QueryMetadataResponseOutput { return v.QueryParams }).(GoogleCloudApigeeV1QueryMetadataResponseOutput)
}

// Asynchronous Report ID.
func (o QueryOutput) ReportDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.ReportDefinitionId }).(pulumi.StringOutput)
}

// Result is available only after the query is completed.
func (o QueryOutput) Result() GoogleCloudApigeeV1AsyncQueryResultResponseOutput {
	return o.ApplyT(func(v *Query) GoogleCloudApigeeV1AsyncQueryResultResponseOutput { return v.Result }).(GoogleCloudApigeeV1AsyncQueryResultResponseOutput)
}

// ResultFileSize is available only after the query is completed.
func (o QueryOutput) ResultFileSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.ResultFileSize }).(pulumi.StringOutput)
}

// ResultRows is available only after the query is completed.
func (o QueryOutput) ResultRows() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.ResultRows }).(pulumi.StringOutput)
}

// Self link of the query. Example: `/organizations/myorg/environments/myenv/queries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostQueries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
func (o QueryOutput) Self() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Self }).(pulumi.StringOutput)
}

// Query state could be "enqueued", "running", "completed", "failed".
func (o QueryOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Last updated timestamp for the query.
func (o QueryOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueryInput)(nil)).Elem(), &Query{})
	pulumi.RegisterOutputType(QueryOutput{})
}
