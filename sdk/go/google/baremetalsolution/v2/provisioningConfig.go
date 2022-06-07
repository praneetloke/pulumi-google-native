// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create new ProvisioningConfig.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type ProvisioningConfig struct {
	pulumi.CustomResourceState

	// URI to Cloud Console UI view of this provisioning config.
	CloudConsoleUri pulumi.StringOutput `pulumi:"cloudConsoleUri"`
	// Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	//
	// Deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	Email pulumi.StringOutput `pulumi:"email"`
	// A service account to enable customers to access instance credentials upon handover.
	HandoverServiceAccount pulumi.StringOutput `pulumi:"handoverServiceAccount"`
	// Instances to be created.
	Instances InstanceConfigResponseArrayOutput `pulumi:"instances"`
	// Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the provisioning config.
	Name pulumi.StringOutput `pulumi:"name"`
	// Networks to be created.
	Networks NetworkConfigResponseArrayOutput `pulumi:"networks"`
	// State of ProvisioningConfig.
	State pulumi.StringOutput `pulumi:"state"`
	// A generated ticket id to track provisioning request.
	TicketId pulumi.StringOutput `pulumi:"ticketId"`
	// Last update timestamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Volumes to be created.
	Volumes VolumeConfigResponseArrayOutput `pulumi:"volumes"`
}

// NewProvisioningConfig registers a new resource with the given unique name, arguments, and options.
func NewProvisioningConfig(ctx *pulumi.Context,
	name string, args *ProvisioningConfigArgs, opts ...pulumi.ResourceOption) (*ProvisioningConfig, error) {
	if args == nil {
		args = &ProvisioningConfigArgs{}
	}

	var resource ProvisioningConfig
	err := ctx.RegisterResource("google-native:baremetalsolution/v2:ProvisioningConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisioningConfig gets an existing ProvisioningConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisioningConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisioningConfigState, opts ...pulumi.ResourceOption) (*ProvisioningConfig, error) {
	var resource ProvisioningConfig
	err := ctx.ReadResource("google-native:baremetalsolution/v2:ProvisioningConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisioningConfig resources.
type provisioningConfigState struct {
}

type ProvisioningConfigState struct {
}

func (ProvisioningConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisioningConfigState)(nil)).Elem()
}

type provisioningConfigArgs struct {
	// Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	//
	// Deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	Email *string `pulumi:"email"`
	// A service account to enable customers to access instance credentials upon handover.
	HandoverServiceAccount *string `pulumi:"handoverServiceAccount"`
	// Instances to be created.
	Instances []InstanceConfig `pulumi:"instances"`
	// Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
	Location *string `pulumi:"location"`
	// Networks to be created.
	Networks []NetworkConfig `pulumi:"networks"`
	Project  *string         `pulumi:"project"`
	// A generated ticket id to track provisioning request.
	TicketId *string `pulumi:"ticketId"`
	// Volumes to be created.
	Volumes []VolumeConfig `pulumi:"volumes"`
}

// The set of arguments for constructing a ProvisioningConfig resource.
type ProvisioningConfigArgs struct {
	// Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	//
	// Deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
	Email pulumi.StringPtrInput
	// A service account to enable customers to access instance credentials upon handover.
	HandoverServiceAccount pulumi.StringPtrInput
	// Instances to be created.
	Instances InstanceConfigArrayInput
	// Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
	Location pulumi.StringPtrInput
	// Networks to be created.
	Networks NetworkConfigArrayInput
	Project  pulumi.StringPtrInput
	// A generated ticket id to track provisioning request.
	TicketId pulumi.StringPtrInput
	// Volumes to be created.
	Volumes VolumeConfigArrayInput
}

func (ProvisioningConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisioningConfigArgs)(nil)).Elem()
}

type ProvisioningConfigInput interface {
	pulumi.Input

	ToProvisioningConfigOutput() ProvisioningConfigOutput
	ToProvisioningConfigOutputWithContext(ctx context.Context) ProvisioningConfigOutput
}

func (*ProvisioningConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningConfig)(nil)).Elem()
}

func (i *ProvisioningConfig) ToProvisioningConfigOutput() ProvisioningConfigOutput {
	return i.ToProvisioningConfigOutputWithContext(context.Background())
}

func (i *ProvisioningConfig) ToProvisioningConfigOutputWithContext(ctx context.Context) ProvisioningConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningConfigOutput)
}

type ProvisioningConfigOutput struct{ *pulumi.OutputState }

func (ProvisioningConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningConfig)(nil)).Elem()
}

func (o ProvisioningConfigOutput) ToProvisioningConfigOutput() ProvisioningConfigOutput {
	return o
}

func (o ProvisioningConfigOutput) ToProvisioningConfigOutputWithContext(ctx context.Context) ProvisioningConfigOutput {
	return o
}

// URI to Cloud Console UI view of this provisioning config.
func (o ProvisioningConfigOutput) CloudConsoleUri() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.CloudConsoleUri }).(pulumi.StringOutput)
}

// Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
//
// Deprecated: Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
func (o ProvisioningConfigOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// A service account to enable customers to access instance credentials upon handover.
func (o ProvisioningConfigOutput) HandoverServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.HandoverServiceAccount }).(pulumi.StringOutput)
}

// Instances to be created.
func (o ProvisioningConfigOutput) Instances() InstanceConfigResponseArrayOutput {
	return o.ApplyT(func(v *ProvisioningConfig) InstanceConfigResponseArrayOutput { return v.Instances }).(InstanceConfigResponseArrayOutput)
}

// Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
func (o ProvisioningConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the provisioning config.
func (o ProvisioningConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Networks to be created.
func (o ProvisioningConfigOutput) Networks() NetworkConfigResponseArrayOutput {
	return o.ApplyT(func(v *ProvisioningConfig) NetworkConfigResponseArrayOutput { return v.Networks }).(NetworkConfigResponseArrayOutput)
}

// State of ProvisioningConfig.
func (o ProvisioningConfigOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A generated ticket id to track provisioning request.
func (o ProvisioningConfigOutput) TicketId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.TicketId }).(pulumi.StringOutput)
}

// Last update timestamp.
func (o ProvisioningConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisioningConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Volumes to be created.
func (o ProvisioningConfigOutput) Volumes() VolumeConfigResponseArrayOutput {
	return o.ApplyT(func(v *ProvisioningConfig) VolumeConfigResponseArrayOutput { return v.Volumes }).(VolumeConfigResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningConfigInput)(nil)).Elem(), &ProvisioningConfig{})
	pulumi.RegisterOutputType(ProvisioningConfigOutput{})
}
