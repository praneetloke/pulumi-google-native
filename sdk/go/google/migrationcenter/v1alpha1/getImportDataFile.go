// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an import data file.
func LookupImportDataFile(ctx *pulumi.Context, args *LookupImportDataFileArgs, opts ...pulumi.InvokeOption) (*LookupImportDataFileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImportDataFileResult
	err := ctx.Invoke("google-native:migrationcenter/v1alpha1:getImportDataFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImportDataFileArgs struct {
	ImportDataFileId string  `pulumi:"importDataFileId"`
	ImportJobId      string  `pulumi:"importJobId"`
	Location         string  `pulumi:"location"`
	Project          *string `pulumi:"project"`
}

type LookupImportDataFileResult struct {
	// The timestamp when the file was created.
	CreateTime string `pulumi:"createTime"`
	// User-friendly display name. Maximum length is 63 characters.
	DisplayName string `pulumi:"displayName"`
	// The payload format.
	Format string `pulumi:"format"`
	// The name of the file.
	Name string `pulumi:"name"`
	// The state of the import data file.
	State string `pulumi:"state"`
	// Information about a file that is uploaded to a storage service.
	UploadFileInfo UploadFileInfoResponse `pulumi:"uploadFileInfo"`
}

func LookupImportDataFileOutput(ctx *pulumi.Context, args LookupImportDataFileOutputArgs, opts ...pulumi.InvokeOption) LookupImportDataFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportDataFileResult, error) {
			args := v.(LookupImportDataFileArgs)
			r, err := LookupImportDataFile(ctx, &args, opts...)
			var s LookupImportDataFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportDataFileResultOutput)
}

type LookupImportDataFileOutputArgs struct {
	ImportDataFileId pulumi.StringInput    `pulumi:"importDataFileId"`
	ImportJobId      pulumi.StringInput    `pulumi:"importJobId"`
	Location         pulumi.StringInput    `pulumi:"location"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupImportDataFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportDataFileArgs)(nil)).Elem()
}

type LookupImportDataFileResultOutput struct{ *pulumi.OutputState }

func (LookupImportDataFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportDataFileResult)(nil)).Elem()
}

func (o LookupImportDataFileResultOutput) ToLookupImportDataFileResultOutput() LookupImportDataFileResultOutput {
	return o
}

func (o LookupImportDataFileResultOutput) ToLookupImportDataFileResultOutputWithContext(ctx context.Context) LookupImportDataFileResultOutput {
	return o
}

// The timestamp when the file was created.
func (o LookupImportDataFileResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// User-friendly display name. Maximum length is 63 characters.
func (o LookupImportDataFileResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The payload format.
func (o LookupImportDataFileResultOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) string { return v.Format }).(pulumi.StringOutput)
}

// The name of the file.
func (o LookupImportDataFileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the import data file.
func (o LookupImportDataFileResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) string { return v.State }).(pulumi.StringOutput)
}

// Information about a file that is uploaded to a storage service.
func (o LookupImportDataFileResultOutput) UploadFileInfo() UploadFileInfoResponseOutput {
	return o.ApplyT(func(v LookupImportDataFileResult) UploadFileInfoResponse { return v.UploadFileInfo }).(UploadFileInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportDataFileResultOutput{})
}
