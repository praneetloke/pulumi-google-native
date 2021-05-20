// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new exclusion in a specified parent resource. Only log entries belonging to that resource can be excluded. You can have up to 10 exclusions in a resource.
type Exclusion struct {
	pulumi.CustomResourceState

	// The creation timestamp of the exclusion.This field may not be present for older exclusions.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A description of this exclusion.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update timestamp of the exclusion.This field may not be present for older exclusions.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewExclusion registers a new resource with the given unique name, arguments, and options.
func NewExclusion(ctx *pulumi.Context,
	name string, args *ExclusionArgs, opts ...pulumi.ResourceOption) (*Exclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExclusionId == nil {
		return nil, errors.New("invalid value for required argument 'ExclusionId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Exclusion
	err := ctx.RegisterResource("google-native:logging/v2:Exclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExclusion gets an existing Exclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExclusionState, opts ...pulumi.ResourceOption) (*Exclusion, error) {
	var resource Exclusion
	err := ctx.ReadResource("google-native:logging/v2:Exclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Exclusion resources.
type exclusionState struct {
	// The creation timestamp of the exclusion.This field may not be present for older exclusions.
	CreateTime *string `pulumi:"createTime"`
	// Optional. A description of this exclusion.
	Description *string `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled *bool `pulumi:"disabled"`
	// Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
	Filter *string `pulumi:"filter"`
	// Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The last update timestamp of the exclusion.This field may not be present for older exclusions.
	UpdateTime *string `pulumi:"updateTime"`
}

type ExclusionState struct {
	// The creation timestamp of the exclusion.This field may not be present for older exclusions.
	CreateTime pulumi.StringPtrInput
	// Optional. A description of this exclusion.
	Description pulumi.StringPtrInput
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolPtrInput
	// Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
	Filter pulumi.StringPtrInput
	// Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The last update timestamp of the exclusion.This field may not be present for older exclusions.
	UpdateTime pulumi.StringPtrInput
}

func (ExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*exclusionState)(nil)).Elem()
}

type exclusionArgs struct {
	// Optional. A description of this exclusion.
	Description *string `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled    *bool  `pulumi:"disabled"`
	ExclusionId string `pulumi:"exclusionId"`
	// Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
	Filter *string `pulumi:"filter"`
	// Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a Exclusion resource.
type ExclusionArgs struct {
	// Optional. A description of this exclusion.
	Description pulumi.StringPtrInput
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled    pulumi.BoolPtrInput
	ExclusionId pulumi.StringInput
	// Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
	Filter pulumi.StringPtrInput
	// Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (ExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exclusionArgs)(nil)).Elem()
}

type ExclusionInput interface {
	pulumi.Input

	ToExclusionOutput() ExclusionOutput
	ToExclusionOutputWithContext(ctx context.Context) ExclusionOutput
}

func (*Exclusion) ElementType() reflect.Type {
	return reflect.TypeOf((*Exclusion)(nil))
}

func (i *Exclusion) ToExclusionOutput() ExclusionOutput {
	return i.ToExclusionOutputWithContext(context.Background())
}

func (i *Exclusion) ToExclusionOutputWithContext(ctx context.Context) ExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExclusionOutput)
}

type ExclusionOutput struct {
	*pulumi.OutputState
}

func (ExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Exclusion)(nil))
}

func (o ExclusionOutput) ToExclusionOutput() ExclusionOutput {
	return o
}

func (o ExclusionOutput) ToExclusionOutputWithContext(ctx context.Context) ExclusionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExclusionOutput{})
}
