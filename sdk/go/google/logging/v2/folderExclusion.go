// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new exclusion in the _Default sink in a specified parent resource. Only log entries belonging to that resource can be excluded. You can have up to 10 exclusions in a resource.
type FolderExclusion struct {
	pulumi.CustomResourceState

	// The creation timestamp of the exclusion.This field may not be present for older exclusions.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A description of this exclusion.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter   pulumi.StringOutput `pulumi:"filter"`
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update timestamp of the exclusion.This field may not be present for older exclusions.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFolderExclusion registers a new resource with the given unique name, arguments, and options.
func NewFolderExclusion(ctx *pulumi.Context,
	name string, args *FolderExclusionArgs, opts ...pulumi.ResourceOption) (*FolderExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"folderId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FolderExclusion
	err := ctx.RegisterResource("google-native:logging/v2:FolderExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderExclusion gets an existing FolderExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderExclusionState, opts ...pulumi.ResourceOption) (*FolderExclusion, error) {
	var resource FolderExclusion
	err := ctx.ReadResource("google-native:logging/v2:FolderExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderExclusion resources.
type folderExclusionState struct {
}

type FolderExclusionState struct {
}

func (FolderExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderExclusionState)(nil)).Elem()
}

type folderExclusionArgs struct {
	// Optional. A description of this exclusion.
	Description *string `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled *bool `pulumi:"disabled"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter   string `pulumi:"filter"`
	FolderId string `pulumi:"folderId"`
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FolderExclusion resource.
type FolderExclusionArgs struct {
	// Optional. A description of this exclusion.
	Description pulumi.StringPtrInput
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolPtrInput
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter   pulumi.StringInput
	FolderId pulumi.StringInput
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
}

func (FolderExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderExclusionArgs)(nil)).Elem()
}

type FolderExclusionInput interface {
	pulumi.Input

	ToFolderExclusionOutput() FolderExclusionOutput
	ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput
}

func (*FolderExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderExclusion)(nil)).Elem()
}

func (i *FolderExclusion) ToFolderExclusionOutput() FolderExclusionOutput {
	return i.ToFolderExclusionOutputWithContext(context.Background())
}

func (i *FolderExclusion) ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderExclusionOutput)
}

type FolderExclusionOutput struct{ *pulumi.OutputState }

func (FolderExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderExclusion)(nil)).Elem()
}

func (o FolderExclusionOutput) ToFolderExclusionOutput() FolderExclusionOutput {
	return o
}

func (o FolderExclusionOutput) ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput {
	return o
}

// The creation timestamp of the exclusion.This field may not be present for older exclusions.
func (o FolderExclusionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A description of this exclusion.
func (o FolderExclusionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
func (o FolderExclusionOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
func (o FolderExclusionOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o FolderExclusionOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
func (o FolderExclusionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The last update timestamp of the exclusion.This field may not be present for older exclusions.
func (o FolderExclusionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderExclusion) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderExclusionInput)(nil)).Elem(), &FolderExclusion{})
	pulumi.RegisterOutputType(FolderExclusionOutput{})
}
