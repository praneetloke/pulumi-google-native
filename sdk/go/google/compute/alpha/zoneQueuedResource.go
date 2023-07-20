// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a QueuedResource.
type ZoneQueuedResource struct {
	pulumi.CustomResourceState

	// Specification of VM instances to create.
	BulkInsertInstanceResource BulkInsertInstanceResourceResponseOutput `pulumi:"bulkInsertInstanceResource"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#queuedResource for QueuedResources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Queuing parameters for the requested capacity.
	QueuingPolicy QueuingPolicyResponseOutput `pulumi:"queuingPolicy"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// [Output only] High-level status of the request.
	State pulumi.StringOutput `pulumi:"state"`
	// [Output only] Result of queuing and provisioning based on deferred capacity.
	Status QueuedResourceStatusResponseOutput `pulumi:"status"`
	Zone   pulumi.StringOutput                `pulumi:"zone"`
}

// NewZoneQueuedResource registers a new resource with the given unique name, arguments, and options.
func NewZoneQueuedResource(ctx *pulumi.Context,
	name string, args *ZoneQueuedResourceArgs, opts ...pulumi.ResourceOption) (*ZoneQueuedResource, error) {
	if args == nil {
		args = &ZoneQueuedResourceArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"zone",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ZoneQueuedResource
	err := ctx.RegisterResource("google-native:compute/alpha:ZoneQueuedResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneQueuedResource gets an existing ZoneQueuedResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneQueuedResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneQueuedResourceState, opts ...pulumi.ResourceOption) (*ZoneQueuedResource, error) {
	var resource ZoneQueuedResource
	err := ctx.ReadResource("google-native:compute/alpha:ZoneQueuedResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneQueuedResource resources.
type zoneQueuedResourceState struct {
}

type ZoneQueuedResourceState struct {
}

func (ZoneQueuedResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneQueuedResourceState)(nil)).Elem()
}

type zoneQueuedResourceArgs struct {
	// Specification of VM instances to create.
	BulkInsertInstanceResource *BulkInsertInstanceResource `pulumi:"bulkInsertInstanceResource"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Queuing parameters for the requested capacity.
	QueuingPolicy *QueuingPolicy `pulumi:"queuingPolicy"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	Zone      *string `pulumi:"zone"`
}

// The set of arguments for constructing a ZoneQueuedResource resource.
type ZoneQueuedResourceArgs struct {
	// Specification of VM instances to create.
	BulkInsertInstanceResource BulkInsertInstanceResourcePtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Queuing parameters for the requested capacity.
	QueuingPolicy QueuingPolicyPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	Zone      pulumi.StringPtrInput
}

func (ZoneQueuedResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneQueuedResourceArgs)(nil)).Elem()
}

type ZoneQueuedResourceInput interface {
	pulumi.Input

	ToZoneQueuedResourceOutput() ZoneQueuedResourceOutput
	ToZoneQueuedResourceOutputWithContext(ctx context.Context) ZoneQueuedResourceOutput
}

func (*ZoneQueuedResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneQueuedResource)(nil)).Elem()
}

func (i *ZoneQueuedResource) ToZoneQueuedResourceOutput() ZoneQueuedResourceOutput {
	return i.ToZoneQueuedResourceOutputWithContext(context.Background())
}

func (i *ZoneQueuedResource) ToZoneQueuedResourceOutputWithContext(ctx context.Context) ZoneQueuedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneQueuedResourceOutput)
}

type ZoneQueuedResourceOutput struct{ *pulumi.OutputState }

func (ZoneQueuedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneQueuedResource)(nil)).Elem()
}

func (o ZoneQueuedResourceOutput) ToZoneQueuedResourceOutput() ZoneQueuedResourceOutput {
	return o
}

func (o ZoneQueuedResourceOutput) ToZoneQueuedResourceOutputWithContext(ctx context.Context) ZoneQueuedResourceOutput {
	return o
}

// Specification of VM instances to create.
func (o ZoneQueuedResourceOutput) BulkInsertInstanceResource() BulkInsertInstanceResourceResponseOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) BulkInsertInstanceResourceResponseOutput {
		return v.BulkInsertInstanceResource
	}).(BulkInsertInstanceResourceResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ZoneQueuedResourceOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o ZoneQueuedResourceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#queuedResource for QueuedResources.
func (o ZoneQueuedResourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o ZoneQueuedResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ZoneQueuedResourceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Queuing parameters for the requested capacity.
func (o ZoneQueuedResourceOutput) QueuingPolicy() QueuingPolicyResponseOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) QueuingPolicyResponseOutput { return v.QueuingPolicy }).(QueuingPolicyResponseOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o ZoneQueuedResourceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// [Output only] Server-defined URL for the resource.
func (o ZoneQueuedResourceOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o ZoneQueuedResourceOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// [Output only] High-level status of the request.
func (o ZoneQueuedResourceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// [Output only] Result of queuing and provisioning based on deferred capacity.
func (o ZoneQueuedResourceOutput) Status() QueuedResourceStatusResponseOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) QueuedResourceStatusResponseOutput { return v.Status }).(QueuedResourceStatusResponseOutput)
}

func (o ZoneQueuedResourceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneQueuedResource) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneQueuedResourceInput)(nil)).Elem(), &ZoneQueuedResource{})
	pulumi.RegisterOutputType(ZoneQueuedResourceOutput{})
}
