// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Network Connectivity Center spoke.
type Spoke struct {
	pulumi.CustomResourceState

	// The time when the Spoke was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Short description of the spoke resource
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource URL of the hub resource that the spoke is attached to
	Hub pulumi.StringOutput `pulumi:"hub"`
	// User-defined labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments pulumi.StringArrayOutput `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances RouterApplianceInstanceResponseArrayOutput `pulumi:"linkedRouterApplianceInstances"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels pulumi.StringArrayOutput `pulumi:"linkedVpnTunnels"`
	Location         pulumi.StringOutput      `pulumi:"location"`
	// Immutable. The name of a Spoke resource.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Optional. Unique id for the Spoke to create.
	SpokeId pulumi.StringPtrOutput `pulumi:"spokeId"`
	// The current lifecycle state of this Hub.
	State pulumi.StringOutput `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// The time when the Spoke was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSpoke registers a new resource with the given unique name, arguments, and options.
func NewSpoke(ctx *pulumi.Context,
	name string, args *SpokeArgs, opts ...pulumi.ResourceOption) (*Spoke, error) {
	if args == nil {
		args = &SpokeArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Spoke
	err := ctx.RegisterResource("google-native:networkconnectivity/v1alpha1:Spoke", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpoke gets an existing Spoke resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpoke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpokeState, opts ...pulumi.ResourceOption) (*Spoke, error) {
	var resource Spoke
	err := ctx.ReadResource("google-native:networkconnectivity/v1alpha1:Spoke", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spoke resources.
type spokeState struct {
}

type SpokeState struct {
}

func (SpokeState) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeState)(nil)).Elem()
}

type spokeArgs struct {
	// The time when the Spoke was created.
	CreateTime *string `pulumi:"createTime"`
	// Short description of the spoke resource
	Description *string `pulumi:"description"`
	// The resource URL of the hub resource that the spoke is attached to
	Hub *string `pulumi:"hub"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments []string `pulumi:"linkedInterconnectAttachments"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances []RouterApplianceInstance `pulumi:"linkedRouterApplianceInstances"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels []string `pulumi:"linkedVpnTunnels"`
	Location         *string  `pulumi:"location"`
	// Immutable. The name of a Spoke resource.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. Unique id for the Spoke to create.
	SpokeId *string `pulumi:"spokeId"`
	// The time when the Spoke was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a Spoke resource.
type SpokeArgs struct {
	// The time when the Spoke was created.
	CreateTime pulumi.StringPtrInput
	// Short description of the spoke resource
	Description pulumi.StringPtrInput
	// The resource URL of the hub resource that the spoke is attached to
	Hub pulumi.StringPtrInput
	// User-defined labels.
	Labels pulumi.StringMapInput
	// The URIs of linked interconnect attachment resources
	LinkedInterconnectAttachments pulumi.StringArrayInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances RouterApplianceInstanceArrayInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels pulumi.StringArrayInput
	Location         pulumi.StringPtrInput
	// Immutable. The name of a Spoke resource.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. Unique id for the Spoke to create.
	SpokeId pulumi.StringPtrInput
	// The time when the Spoke was updated.
	UpdateTime pulumi.StringPtrInput
}

func (SpokeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spokeArgs)(nil)).Elem()
}

type SpokeInput interface {
	pulumi.Input

	ToSpokeOutput() SpokeOutput
	ToSpokeOutputWithContext(ctx context.Context) SpokeOutput
}

func (*Spoke) ElementType() reflect.Type {
	return reflect.TypeOf((**Spoke)(nil)).Elem()
}

func (i *Spoke) ToSpokeOutput() SpokeOutput {
	return i.ToSpokeOutputWithContext(context.Background())
}

func (i *Spoke) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpokeOutput)
}

type SpokeOutput struct{ *pulumi.OutputState }

func (SpokeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spoke)(nil)).Elem()
}

func (o SpokeOutput) ToSpokeOutput() SpokeOutput {
	return o
}

func (o SpokeOutput) ToSpokeOutputWithContext(ctx context.Context) SpokeOutput {
	return o
}

// The time when the Spoke was created.
func (o SpokeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Short description of the spoke resource
func (o SpokeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The resource URL of the hub resource that the spoke is attached to
func (o SpokeOutput) Hub() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Hub }).(pulumi.StringOutput)
}

// User-defined labels.
func (o SpokeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The URIs of linked interconnect attachment resources
func (o SpokeOutput) LinkedInterconnectAttachments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringArrayOutput { return v.LinkedInterconnectAttachments }).(pulumi.StringArrayOutput)
}

// The URIs of linked Router appliance resources
func (o SpokeOutput) LinkedRouterApplianceInstances() RouterApplianceInstanceResponseArrayOutput {
	return o.ApplyT(func(v *Spoke) RouterApplianceInstanceResponseArrayOutput { return v.LinkedRouterApplianceInstances }).(RouterApplianceInstanceResponseArrayOutput)
}

// The URIs of linked VPN tunnel resources
func (o SpokeOutput) LinkedVpnTunnels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringArrayOutput { return v.LinkedVpnTunnels }).(pulumi.StringArrayOutput)
}

func (o SpokeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The name of a Spoke resource.
func (o SpokeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpokeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o SpokeOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Optional. Unique id for the Spoke to create.
func (o SpokeOutput) SpokeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringPtrOutput { return v.SpokeId }).(pulumi.StringPtrOutput)
}

// The current lifecycle state of this Hub.
func (o SpokeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
func (o SpokeOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

// The time when the Spoke was updated.
func (o SpokeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Spoke) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpokeInput)(nil)).Elem(), &Spoke{})
	pulumi.RegisterOutputType(SpokeOutput{})
}
