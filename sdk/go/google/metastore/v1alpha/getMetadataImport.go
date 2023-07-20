// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single import.
func LookupMetadataImport(ctx *pulumi.Context, args *LookupMetadataImportArgs, opts ...pulumi.InvokeOption) (*LookupMetadataImportResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMetadataImportResult
	err := ctx.Invoke("google-native:metastore/v1alpha:getMetadataImport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetadataImportArgs struct {
	Location         string  `pulumi:"location"`
	MetadataImportId string  `pulumi:"metadataImportId"`
	Project          *string `pulumi:"project"`
	ServiceId        string  `pulumi:"serviceId"`
}

type LookupMetadataImportResult struct {
	// The time when the metadata import was started.
	CreateTime string `pulumi:"createTime"`
	// Immutable. A database dump from a pre-existing metastore's database.
	DatabaseDump DatabaseDumpResponse `pulumi:"databaseDump"`
	// The description of the metadata import.
	Description string `pulumi:"description"`
	// The time when the metadata import finished.
	EndTime string `pulumi:"endTime"`
	// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
	Name string `pulumi:"name"`
	// The current state of the metadata import.
	State string `pulumi:"state"`
	// The time when the metadata import was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupMetadataImportOutput(ctx *pulumi.Context, args LookupMetadataImportOutputArgs, opts ...pulumi.InvokeOption) LookupMetadataImportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetadataImportResult, error) {
			args := v.(LookupMetadataImportArgs)
			r, err := LookupMetadataImport(ctx, &args, opts...)
			var s LookupMetadataImportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetadataImportResultOutput)
}

type LookupMetadataImportOutputArgs struct {
	Location         pulumi.StringInput    `pulumi:"location"`
	MetadataImportId pulumi.StringInput    `pulumi:"metadataImportId"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
	ServiceId        pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupMetadataImportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataImportArgs)(nil)).Elem()
}

type LookupMetadataImportResultOutput struct{ *pulumi.OutputState }

func (LookupMetadataImportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataImportResult)(nil)).Elem()
}

func (o LookupMetadataImportResultOutput) ToLookupMetadataImportResultOutput() LookupMetadataImportResultOutput {
	return o
}

func (o LookupMetadataImportResultOutput) ToLookupMetadataImportResultOutputWithContext(ctx context.Context) LookupMetadataImportResultOutput {
	return o
}

// The time when the metadata import was started.
func (o LookupMetadataImportResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Immutable. A database dump from a pre-existing metastore's database.
func (o LookupMetadataImportResultOutput) DatabaseDump() DatabaseDumpResponseOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) DatabaseDumpResponse { return v.DatabaseDump }).(DatabaseDumpResponseOutput)
}

// The description of the metadata import.
func (o LookupMetadataImportResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.Description }).(pulumi.StringOutput)
}

// The time when the metadata import finished.
func (o LookupMetadataImportResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
func (o LookupMetadataImportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of the metadata import.
func (o LookupMetadataImportResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.State }).(pulumi.StringOutput)
}

// The time when the metadata import was last updated.
func (o LookupMetadataImportResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataImportResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetadataImportResultOutput{})
}
