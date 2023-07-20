// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a BigQuery export.
func LookupOrganizationBigQueryExport(ctx *pulumi.Context, args *LookupOrganizationBigQueryExportArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationBigQueryExportResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationBigQueryExportResult
	err := ctx.Invoke("google-native:securitycenter/v1:getOrganizationBigQueryExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationBigQueryExportArgs struct {
	BigQueryExportId string `pulumi:"bigQueryExportId"`
	OrganizationId   string `pulumi:"organizationId"`
}

type LookupOrganizationBigQueryExportResult struct {
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on export on creation.
	CreateTime string `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset string `pulumi:"dataset"`
	// The description of the export (max of 1024 characters).
	Description string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes.
	Filter string `pulumi:"filter"`
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if provided on export creation or update.
	MostRecentEditor string `pulumi:"mostRecentEditor"`
	// The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
	Name string `pulumi:"name"`
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal string `pulumi:"principal"`
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if provided on export creation or update.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupOrganizationBigQueryExportOutput(ctx *pulumi.Context, args LookupOrganizationBigQueryExportOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationBigQueryExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationBigQueryExportResult, error) {
			args := v.(LookupOrganizationBigQueryExportArgs)
			r, err := LookupOrganizationBigQueryExport(ctx, &args, opts...)
			var s LookupOrganizationBigQueryExportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationBigQueryExportResultOutput)
}

type LookupOrganizationBigQueryExportOutputArgs struct {
	BigQueryExportId pulumi.StringInput `pulumi:"bigQueryExportId"`
	OrganizationId   pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupOrganizationBigQueryExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationBigQueryExportArgs)(nil)).Elem()
}

type LookupOrganizationBigQueryExportResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationBigQueryExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationBigQueryExportResult)(nil)).Elem()
}

func (o LookupOrganizationBigQueryExportResultOutput) ToLookupOrganizationBigQueryExportResultOutput() LookupOrganizationBigQueryExportResultOutput {
	return o
}

func (o LookupOrganizationBigQueryExportResultOutput) ToLookupOrganizationBigQueryExportResultOutputWithContext(ctx context.Context) LookupOrganizationBigQueryExportResultOutput {
	return o
}

// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on export on creation.
func (o LookupOrganizationBigQueryExportResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
func (o LookupOrganizationBigQueryExportResultOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.Dataset }).(pulumi.StringOutput)
}

// The description of the export (max of 1024 characters).
func (o LookupOrganizationBigQueryExportResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.Description }).(pulumi.StringOutput)
}

// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes.
func (o LookupOrganizationBigQueryExportResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.Filter }).(pulumi.StringOutput)
}

// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if provided on export creation or update.
func (o LookupOrganizationBigQueryExportResultOutput) MostRecentEditor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.MostRecentEditor }).(pulumi.StringOutput)
}

// The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
func (o LookupOrganizationBigQueryExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service account that needs permission to create table and upload data to the BigQuery dataset.
func (o LookupOrganizationBigQueryExportResultOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.Principal }).(pulumi.StringOutput)
}

// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if provided on export creation or update.
func (o LookupOrganizationBigQueryExportResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBigQueryExportResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationBigQueryExportResultOutput{})
}
