// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// Message describing compute engine instance filter
type GceInstanceFilter struct {
	// Service account of compute engine
	ServiceAccounts []string `pulumi:"serviceAccounts"`
}

// GceInstanceFilterInput is an input type that accepts GceInstanceFilterArgs and GceInstanceFilterOutput values.
// You can construct a concrete instance of `GceInstanceFilterInput` via:
//
//	GceInstanceFilterArgs{...}
type GceInstanceFilterInput interface {
	pulumi.Input

	ToGceInstanceFilterOutput() GceInstanceFilterOutput
	ToGceInstanceFilterOutputWithContext(context.Context) GceInstanceFilterOutput
}

// Message describing compute engine instance filter
type GceInstanceFilterArgs struct {
	// Service account of compute engine
	ServiceAccounts pulumi.StringArrayInput `pulumi:"serviceAccounts"`
}

func (GceInstanceFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GceInstanceFilter)(nil)).Elem()
}

func (i GceInstanceFilterArgs) ToGceInstanceFilterOutput() GceInstanceFilterOutput {
	return i.ToGceInstanceFilterOutputWithContext(context.Background())
}

func (i GceInstanceFilterArgs) ToGceInstanceFilterOutputWithContext(ctx context.Context) GceInstanceFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GceInstanceFilterOutput)
}

func (i GceInstanceFilterArgs) ToGceInstanceFilterPtrOutput() GceInstanceFilterPtrOutput {
	return i.ToGceInstanceFilterPtrOutputWithContext(context.Background())
}

func (i GceInstanceFilterArgs) ToGceInstanceFilterPtrOutputWithContext(ctx context.Context) GceInstanceFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GceInstanceFilterOutput).ToGceInstanceFilterPtrOutputWithContext(ctx)
}

// GceInstanceFilterPtrInput is an input type that accepts GceInstanceFilterArgs, GceInstanceFilterPtr and GceInstanceFilterPtrOutput values.
// You can construct a concrete instance of `GceInstanceFilterPtrInput` via:
//
//	        GceInstanceFilterArgs{...}
//
//	or:
//
//	        nil
type GceInstanceFilterPtrInput interface {
	pulumi.Input

	ToGceInstanceFilterPtrOutput() GceInstanceFilterPtrOutput
	ToGceInstanceFilterPtrOutputWithContext(context.Context) GceInstanceFilterPtrOutput
}

type gceInstanceFilterPtrType GceInstanceFilterArgs

func GceInstanceFilterPtr(v *GceInstanceFilterArgs) GceInstanceFilterPtrInput {
	return (*gceInstanceFilterPtrType)(v)
}

func (*gceInstanceFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GceInstanceFilter)(nil)).Elem()
}

func (i *gceInstanceFilterPtrType) ToGceInstanceFilterPtrOutput() GceInstanceFilterPtrOutput {
	return i.ToGceInstanceFilterPtrOutputWithContext(context.Background())
}

func (i *gceInstanceFilterPtrType) ToGceInstanceFilterPtrOutputWithContext(ctx context.Context) GceInstanceFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GceInstanceFilterPtrOutput)
}

// Message describing compute engine instance filter
type GceInstanceFilterOutput struct{ *pulumi.OutputState }

func (GceInstanceFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GceInstanceFilter)(nil)).Elem()
}

func (o GceInstanceFilterOutput) ToGceInstanceFilterOutput() GceInstanceFilterOutput {
	return o
}

func (o GceInstanceFilterOutput) ToGceInstanceFilterOutputWithContext(ctx context.Context) GceInstanceFilterOutput {
	return o
}

func (o GceInstanceFilterOutput) ToGceInstanceFilterPtrOutput() GceInstanceFilterPtrOutput {
	return o.ToGceInstanceFilterPtrOutputWithContext(context.Background())
}

func (o GceInstanceFilterOutput) ToGceInstanceFilterPtrOutputWithContext(ctx context.Context) GceInstanceFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GceInstanceFilter) *GceInstanceFilter {
		return &v
	}).(GceInstanceFilterPtrOutput)
}

// Service account of compute engine
func (o GceInstanceFilterOutput) ServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GceInstanceFilter) []string { return v.ServiceAccounts }).(pulumi.StringArrayOutput)
}

type GceInstanceFilterPtrOutput struct{ *pulumi.OutputState }

func (GceInstanceFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GceInstanceFilter)(nil)).Elem()
}

func (o GceInstanceFilterPtrOutput) ToGceInstanceFilterPtrOutput() GceInstanceFilterPtrOutput {
	return o
}

func (o GceInstanceFilterPtrOutput) ToGceInstanceFilterPtrOutputWithContext(ctx context.Context) GceInstanceFilterPtrOutput {
	return o
}

func (o GceInstanceFilterPtrOutput) Elem() GceInstanceFilterOutput {
	return o.ApplyT(func(v *GceInstanceFilter) GceInstanceFilter {
		if v != nil {
			return *v
		}
		var ret GceInstanceFilter
		return ret
	}).(GceInstanceFilterOutput)
}

// Service account of compute engine
func (o GceInstanceFilterPtrOutput) ServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GceInstanceFilter) []string {
		if v == nil {
			return nil
		}
		return v.ServiceAccounts
	}).(pulumi.StringArrayOutput)
}

// Message describing compute engine instance filter
type GceInstanceFilterResponse struct {
	// Service account of compute engine
	ServiceAccounts []string `pulumi:"serviceAccounts"`
}

// Message describing compute engine instance filter
type GceInstanceFilterResponseOutput struct{ *pulumi.OutputState }

func (GceInstanceFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GceInstanceFilterResponse)(nil)).Elem()
}

func (o GceInstanceFilterResponseOutput) ToGceInstanceFilterResponseOutput() GceInstanceFilterResponseOutput {
	return o
}

func (o GceInstanceFilterResponseOutput) ToGceInstanceFilterResponseOutputWithContext(ctx context.Context) GceInstanceFilterResponseOutput {
	return o
}

// Service account of compute engine
func (o GceInstanceFilterResponseOutput) ServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GceInstanceFilterResponse) []string { return v.ServiceAccounts }).(pulumi.StringArrayOutput)
}

// Message describing resource filters
type ResourceFilter struct {
	// Filter compute engine resource
	GceInstanceFilter *GceInstanceFilter `pulumi:"gceInstanceFilter"`
	// The label used for filter resource
	InclusionLabels map[string]string `pulumi:"inclusionLabels"`
	// The id pattern for filter resource
	ResourceIdPatterns []string `pulumi:"resourceIdPatterns"`
	// The scopes of evaluation resource
	Scopes []string `pulumi:"scopes"`
}

// ResourceFilterInput is an input type that accepts ResourceFilterArgs and ResourceFilterOutput values.
// You can construct a concrete instance of `ResourceFilterInput` via:
//
//	ResourceFilterArgs{...}
type ResourceFilterInput interface {
	pulumi.Input

	ToResourceFilterOutput() ResourceFilterOutput
	ToResourceFilterOutputWithContext(context.Context) ResourceFilterOutput
}

// Message describing resource filters
type ResourceFilterArgs struct {
	// Filter compute engine resource
	GceInstanceFilter GceInstanceFilterPtrInput `pulumi:"gceInstanceFilter"`
	// The label used for filter resource
	InclusionLabels pulumi.StringMapInput `pulumi:"inclusionLabels"`
	// The id pattern for filter resource
	ResourceIdPatterns pulumi.StringArrayInput `pulumi:"resourceIdPatterns"`
	// The scopes of evaluation resource
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
}

func (ResourceFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFilter)(nil)).Elem()
}

func (i ResourceFilterArgs) ToResourceFilterOutput() ResourceFilterOutput {
	return i.ToResourceFilterOutputWithContext(context.Background())
}

func (i ResourceFilterArgs) ToResourceFilterOutputWithContext(ctx context.Context) ResourceFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceFilterOutput)
}

func (i ResourceFilterArgs) ToResourceFilterPtrOutput() ResourceFilterPtrOutput {
	return i.ToResourceFilterPtrOutputWithContext(context.Background())
}

func (i ResourceFilterArgs) ToResourceFilterPtrOutputWithContext(ctx context.Context) ResourceFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceFilterOutput).ToResourceFilterPtrOutputWithContext(ctx)
}

// ResourceFilterPtrInput is an input type that accepts ResourceFilterArgs, ResourceFilterPtr and ResourceFilterPtrOutput values.
// You can construct a concrete instance of `ResourceFilterPtrInput` via:
//
//	        ResourceFilterArgs{...}
//
//	or:
//
//	        nil
type ResourceFilterPtrInput interface {
	pulumi.Input

	ToResourceFilterPtrOutput() ResourceFilterPtrOutput
	ToResourceFilterPtrOutputWithContext(context.Context) ResourceFilterPtrOutput
}

type resourceFilterPtrType ResourceFilterArgs

func ResourceFilterPtr(v *ResourceFilterArgs) ResourceFilterPtrInput {
	return (*resourceFilterPtrType)(v)
}

func (*resourceFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceFilter)(nil)).Elem()
}

func (i *resourceFilterPtrType) ToResourceFilterPtrOutput() ResourceFilterPtrOutput {
	return i.ToResourceFilterPtrOutputWithContext(context.Background())
}

func (i *resourceFilterPtrType) ToResourceFilterPtrOutputWithContext(ctx context.Context) ResourceFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceFilterPtrOutput)
}

// Message describing resource filters
type ResourceFilterOutput struct{ *pulumi.OutputState }

func (ResourceFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFilter)(nil)).Elem()
}

func (o ResourceFilterOutput) ToResourceFilterOutput() ResourceFilterOutput {
	return o
}

func (o ResourceFilterOutput) ToResourceFilterOutputWithContext(ctx context.Context) ResourceFilterOutput {
	return o
}

func (o ResourceFilterOutput) ToResourceFilterPtrOutput() ResourceFilterPtrOutput {
	return o.ToResourceFilterPtrOutputWithContext(context.Background())
}

func (o ResourceFilterOutput) ToResourceFilterPtrOutputWithContext(ctx context.Context) ResourceFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceFilter) *ResourceFilter {
		return &v
	}).(ResourceFilterPtrOutput)
}

// Filter compute engine resource
func (o ResourceFilterOutput) GceInstanceFilter() GceInstanceFilterPtrOutput {
	return o.ApplyT(func(v ResourceFilter) *GceInstanceFilter { return v.GceInstanceFilter }).(GceInstanceFilterPtrOutput)
}

// The label used for filter resource
func (o ResourceFilterOutput) InclusionLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceFilter) map[string]string { return v.InclusionLabels }).(pulumi.StringMapOutput)
}

// The id pattern for filter resource
func (o ResourceFilterOutput) ResourceIdPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceFilter) []string { return v.ResourceIdPatterns }).(pulumi.StringArrayOutput)
}

// The scopes of evaluation resource
func (o ResourceFilterOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceFilter) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ResourceFilterPtrOutput struct{ *pulumi.OutputState }

func (ResourceFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceFilter)(nil)).Elem()
}

func (o ResourceFilterPtrOutput) ToResourceFilterPtrOutput() ResourceFilterPtrOutput {
	return o
}

func (o ResourceFilterPtrOutput) ToResourceFilterPtrOutputWithContext(ctx context.Context) ResourceFilterPtrOutput {
	return o
}

func (o ResourceFilterPtrOutput) Elem() ResourceFilterOutput {
	return o.ApplyT(func(v *ResourceFilter) ResourceFilter {
		if v != nil {
			return *v
		}
		var ret ResourceFilter
		return ret
	}).(ResourceFilterOutput)
}

// Filter compute engine resource
func (o ResourceFilterPtrOutput) GceInstanceFilter() GceInstanceFilterPtrOutput {
	return o.ApplyT(func(v *ResourceFilter) *GceInstanceFilter {
		if v == nil {
			return nil
		}
		return v.GceInstanceFilter
	}).(GceInstanceFilterPtrOutput)
}

// The label used for filter resource
func (o ResourceFilterPtrOutput) InclusionLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceFilter) map[string]string {
		if v == nil {
			return nil
		}
		return v.InclusionLabels
	}).(pulumi.StringMapOutput)
}

// The id pattern for filter resource
func (o ResourceFilterPtrOutput) ResourceIdPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceFilter) []string {
		if v == nil {
			return nil
		}
		return v.ResourceIdPatterns
	}).(pulumi.StringArrayOutput)
}

// The scopes of evaluation resource
func (o ResourceFilterPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceFilter) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

// Message describing resource filters
type ResourceFilterResponse struct {
	// Filter compute engine resource
	GceInstanceFilter GceInstanceFilterResponse `pulumi:"gceInstanceFilter"`
	// The label used for filter resource
	InclusionLabels map[string]string `pulumi:"inclusionLabels"`
	// The id pattern for filter resource
	ResourceIdPatterns []string `pulumi:"resourceIdPatterns"`
	// The scopes of evaluation resource
	Scopes []string `pulumi:"scopes"`
}

// Message describing resource filters
type ResourceFilterResponseOutput struct{ *pulumi.OutputState }

func (ResourceFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceFilterResponse)(nil)).Elem()
}

func (o ResourceFilterResponseOutput) ToResourceFilterResponseOutput() ResourceFilterResponseOutput {
	return o
}

func (o ResourceFilterResponseOutput) ToResourceFilterResponseOutputWithContext(ctx context.Context) ResourceFilterResponseOutput {
	return o
}

// Filter compute engine resource
func (o ResourceFilterResponseOutput) GceInstanceFilter() GceInstanceFilterResponseOutput {
	return o.ApplyT(func(v ResourceFilterResponse) GceInstanceFilterResponse { return v.GceInstanceFilter }).(GceInstanceFilterResponseOutput)
}

// The label used for filter resource
func (o ResourceFilterResponseOutput) InclusionLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceFilterResponse) map[string]string { return v.InclusionLabels }).(pulumi.StringMapOutput)
}

// The id pattern for filter resource
func (o ResourceFilterResponseOutput) ResourceIdPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceFilterResponse) []string { return v.ResourceIdPatterns }).(pulumi.StringArrayOutput)
}

// The scopes of evaluation resource
func (o ResourceFilterResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceFilterResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Message describing resource status
type ResourceStatusResponse struct {
	// the new version of rule id if exists
	RulesNewerVersions []string `pulumi:"rulesNewerVersions"`
	// State of the resource
	State string `pulumi:"state"`
}

// Message describing resource status
type ResourceStatusResponseOutput struct{ *pulumi.OutputState }

func (ResourceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceStatusResponse)(nil)).Elem()
}

func (o ResourceStatusResponseOutput) ToResourceStatusResponseOutput() ResourceStatusResponseOutput {
	return o
}

func (o ResourceStatusResponseOutput) ToResourceStatusResponseOutputWithContext(ctx context.Context) ResourceStatusResponseOutput {
	return o
}

// the new version of rule id if exists
func (o ResourceStatusResponseOutput) RulesNewerVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceStatusResponse) []string { return v.RulesNewerVersions }).(pulumi.StringArrayOutput)
}

// State of the resource
func (o ResourceStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GceInstanceFilterInput)(nil)).Elem(), GceInstanceFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GceInstanceFilterPtrInput)(nil)).Elem(), GceInstanceFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceFilterInput)(nil)).Elem(), ResourceFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceFilterPtrInput)(nil)).Elem(), ResourceFilterArgs{})
	pulumi.RegisterOutputType(GceInstanceFilterOutput{})
	pulumi.RegisterOutputType(GceInstanceFilterPtrOutput{})
	pulumi.RegisterOutputType(GceInstanceFilterResponseOutput{})
	pulumi.RegisterOutputType(ResourceFilterOutput{})
	pulumi.RegisterOutputType(ResourceFilterPtrOutput{})
	pulumi.RegisterOutputType(ResourceFilterResponseOutput{})
	pulumi.RegisterOutputType(ResourceStatusResponseOutput{})
}
