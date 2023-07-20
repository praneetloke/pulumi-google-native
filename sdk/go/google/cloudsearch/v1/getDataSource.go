// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a datasource. **Note:** This API requires an admin account to execute.
func LookupDataSource(ctx *pulumi.Context, args *LookupDataSourceArgs, opts ...pulumi.InvokeOption) (*LookupDataSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataSourceResult
	err := ctx.Invoke("google-native:cloudsearch/v1:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSourceArgs struct {
	DatasourceId                string `pulumi:"datasourceId"`
	DebugOptionsEnableDebugging *bool  `pulumi:"debugOptionsEnableDebugging"`
}

type LookupDataSourceResult struct {
	// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
	DisableModifications bool `pulumi:"disableModifications"`
	// Disable serving any search or assist results.
	DisableServing bool `pulumi:"disableServing"`
	// Display name of the datasource The maximum length is 300 characters.
	DisplayName string `pulumi:"displayName"`
	// List of service accounts that have indexing access.
	IndexingServiceAccounts []string `pulumi:"indexingServiceAccounts"`
	// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
	ItemsVisibility []GSuitePrincipalResponse `pulumi:"itemsVisibility"`
	// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
	Name string `pulumi:"name"`
	// IDs of the Long Running Operations (LROs) currently running for this schema.
	OperationIds []string `pulumi:"operationIds"`
	// Can a user request to get thumbnail URI for Items indexed in this data source.
	ReturnThumbnailUrls bool `pulumi:"returnThumbnailUrls"`
	// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
	ShortName string `pulumi:"shortName"`
}

func LookupDataSourceOutput(ctx *pulumi.Context, args LookupDataSourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataSourceResult, error) {
			args := v.(LookupDataSourceArgs)
			r, err := LookupDataSource(ctx, &args, opts...)
			var s LookupDataSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataSourceResultOutput)
}

type LookupDataSourceOutputArgs struct {
	DatasourceId                pulumi.StringInput  `pulumi:"datasourceId"`
	DebugOptionsEnableDebugging pulumi.BoolPtrInput `pulumi:"debugOptionsEnableDebugging"`
}

func (LookupDataSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceArgs)(nil)).Elem()
}

type LookupDataSourceResultOutput struct{ *pulumi.OutputState }

func (LookupDataSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceResult)(nil)).Elem()
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutput() LookupDataSourceResultOutput {
	return o
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutputWithContext(ctx context.Context) LookupDataSourceResultOutput {
	return o
}

// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
func (o LookupDataSourceResultOutput) DisableModifications() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataSourceResult) bool { return v.DisableModifications }).(pulumi.BoolOutput)
}

// Disable serving any search or assist results.
func (o LookupDataSourceResultOutput) DisableServing() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataSourceResult) bool { return v.DisableServing }).(pulumi.BoolOutput)
}

// Display name of the datasource The maximum length is 300 characters.
func (o LookupDataSourceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// List of service accounts that have indexing access.
func (o LookupDataSourceResultOutput) IndexingServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []string { return v.IndexingServiceAccounts }).(pulumi.StringArrayOutput)
}

// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
func (o LookupDataSourceResultOutput) ItemsVisibility() GSuitePrincipalResponseArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []GSuitePrincipalResponse { return v.ItemsVisibility }).(GSuitePrincipalResponseArrayOutput)
}

// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
func (o LookupDataSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// IDs of the Long Running Operations (LROs) currently running for this schema.
func (o LookupDataSourceResultOutput) OperationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []string { return v.OperationIds }).(pulumi.StringArrayOutput)
}

// Can a user request to get thumbnail URI for Items indexed in this data source.
func (o LookupDataSourceResultOutput) ReturnThumbnailUrls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataSourceResult) bool { return v.ReturnThumbnailUrls }).(pulumi.BoolOutput)
}

// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
func (o LookupDataSourceResultOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSourceResult) string { return v.ShortName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSourceResultOutput{})
}
