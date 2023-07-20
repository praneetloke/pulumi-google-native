// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a resident SecurityHealthAnalyticsCustomModule at the scope of the given CRM parent, and also creates inherited SecurityHealthAnalyticsCustomModules for all CRM descendants of the given parent. These modules are enabled by default.
// Auto-naming is currently not supported for this resource.
type CustomModule struct {
	pulumi.CustomResourceState

	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
	AncestorModule pulumi.StringOutput `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig GoogleCloudSecuritycenterV1CustomConfigResponseOutput `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The enablement state of the custom module.
	EnablementState pulumi.StringOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The time at which the custom module was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCustomModule registers a new resource with the given unique name, arguments, and options.
func NewCustomModule(ctx *pulumi.Context,
	name string, args *CustomModuleArgs, opts ...pulumi.ResourceOption) (*CustomModule, error) {
	if args == nil {
		args = &CustomModuleArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomModule
	err := ctx.RegisterResource("google-native:securitycenter/v1:CustomModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomModule gets an existing CustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomModuleState, opts ...pulumi.ResourceOption) (*CustomModule, error) {
	var resource CustomModule
	err := ctx.ReadResource("google-native:securitycenter/v1:CustomModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomModule resources.
type customModuleState struct {
}

type CustomModuleState struct {
}

func (CustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*customModuleState)(nil)).Elem()
}

type customModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig *GoogleCloudSecuritycenterV1CustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName *string `pulumi:"displayName"`
	// The enablement state of the custom module.
	EnablementState *CustomModuleEnablementState `pulumi:"enablementState"`
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a CustomModule resource.
type CustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig GoogleCloudSecuritycenterV1CustomConfigPtrInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrInput
	// The enablement state of the custom module.
	EnablementState CustomModuleEnablementStatePtrInput
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (CustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customModuleArgs)(nil)).Elem()
}

type CustomModuleInput interface {
	pulumi.Input

	ToCustomModuleOutput() CustomModuleOutput
	ToCustomModuleOutputWithContext(ctx context.Context) CustomModuleOutput
}

func (*CustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomModule)(nil)).Elem()
}

func (i *CustomModule) ToCustomModuleOutput() CustomModuleOutput {
	return i.ToCustomModuleOutputWithContext(context.Background())
}

func (i *CustomModule) ToCustomModuleOutputWithContext(ctx context.Context) CustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomModuleOutput)
}

type CustomModuleOutput struct{ *pulumi.OutputState }

func (CustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomModule)(nil)).Elem()
}

func (o CustomModuleOutput) ToCustomModuleOutput() CustomModuleOutput {
	return o
}

func (o CustomModuleOutput) ToCustomModuleOutputWithContext(ctx context.Context) CustomModuleOutput {
	return o
}

// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
func (o CustomModuleOutput) AncestorModule() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.AncestorModule }).(pulumi.StringOutput)
}

// The user specified custom configuration for the module.
func (o CustomModuleOutput) CustomConfig() GoogleCloudSecuritycenterV1CustomConfigResponseOutput {
	return o.ApplyT(func(v *CustomModule) GoogleCloudSecuritycenterV1CustomConfigResponseOutput { return v.CustomConfig }).(GoogleCloudSecuritycenterV1CustomConfigResponseOutput)
}

// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
func (o CustomModuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The enablement state of the custom module.
func (o CustomModuleOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.EnablementState }).(pulumi.StringOutput)
}

// The editor that last updated the custom module.
func (o CustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.LastEditor }).(pulumi.StringOutput)
}

// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
func (o CustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomModuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The time at which the custom module was last updated.
func (o CustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomModule) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomModuleInput)(nil)).Elem(), &CustomModule{})
	pulumi.RegisterOutputType(CustomModuleOutput{})
}
