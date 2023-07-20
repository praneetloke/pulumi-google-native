// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a Folder identified by the supplied resource name. Valid Folder resource names have the format `folders/{folder_id}` (for example, `folders/1234`). The caller must have `resourcemanager.folders.get` permission on the identified folder.
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFolderResult
	err := ctx.Invoke("google-native:cloudresourcemanager/v2beta1:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderArgs struct {
	FolderId string `pulumi:"folderId"`
}

type LookupFolderResult struct {
	// Timestamp when the Folder was created. Assigned by the server.
	CreateTime string `pulumi:"createTime"`
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
	DisplayName string `pulumi:"displayName"`
	// The lifecycle state of the folder. Updates to the lifecycle_state must be performed via DeleteFolder and UndeleteFolder.
	LifecycleState string `pulumi:"lifecycleState"`
	// The resource name of the Folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
	Name string `pulumi:"name"`
	// The Folder's parent's resource name. Updates to the folder's parent must be performed via MoveFolder.
	Parent string `pulumi:"parent"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderResult, error) {
			args := v.(LookupFolderArgs)
			r, err := LookupFolder(ctx, &args, opts...)
			var s LookupFolderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderResultOutput)
}

type LookupFolderOutputArgs struct {
	FolderId pulumi.StringInput `pulumi:"folderId"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// Timestamp when the Folder was created. Assigned by the server.
func (o LookupFolderResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
func (o LookupFolderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The lifecycle state of the folder. Updates to the lifecycle_state must be performed via DeleteFolder and UndeleteFolder.
func (o LookupFolderResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

// The resource name of the Folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
func (o LookupFolderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Folder's parent's resource name. Updates to the folder's parent must be performed via MoveFolder.
func (o LookupFolderResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Parent }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
