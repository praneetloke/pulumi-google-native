// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ClientConnectorService in a given project and location.
type ClientConnectorService struct {
	pulumi.CustomResourceState

	// Optional. User-settable client connector service resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter. A random system generated name will be assigned if not specified by the user.
	ClientConnectorServiceId pulumi.StringPtrOutput `pulumi:"clientConnectorServiceId"`
	// [Output only] Create time stamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The details of the egress settings.
	Egress EgressResponseOutput `pulumi:"egress"`
	// The details of the ingress settings.
	Ingress  IngressResponseOutput `pulumi:"ingress"`
	Location pulumi.StringOutput   `pulumi:"location"`
	// Name of resource. The name is ignored during creation.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The operational state of the ClientConnectorService.
	State pulumi.StringOutput `pulumi:"state"`
	// [Output only] Update time stamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
	ValidateOnly pulumi.StringPtrOutput `pulumi:"validateOnly"`
}

// NewClientConnectorService registers a new resource with the given unique name, arguments, and options.
func NewClientConnectorService(ctx *pulumi.Context,
	name string, args *ClientConnectorServiceArgs, opts ...pulumi.ResourceOption) (*ClientConnectorService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Egress == nil {
		return nil, errors.New("invalid value for required argument 'Egress'")
	}
	if args.Ingress == nil {
		return nil, errors.New("invalid value for required argument 'Ingress'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource ClientConnectorService
	err := ctx.RegisterResource("google-native:beyondcorp/v1:ClientConnectorService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientConnectorService gets an existing ClientConnectorService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientConnectorService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientConnectorServiceState, opts ...pulumi.ResourceOption) (*ClientConnectorService, error) {
	var resource ClientConnectorService
	err := ctx.ReadResource("google-native:beyondcorp/v1:ClientConnectorService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientConnectorService resources.
type clientConnectorServiceState struct {
}

type ClientConnectorServiceState struct {
}

func (ClientConnectorServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientConnectorServiceState)(nil)).Elem()
}

type clientConnectorServiceArgs struct {
	// Optional. User-settable client connector service resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter. A random system generated name will be assigned if not specified by the user.
	ClientConnectorServiceId *string `pulumi:"clientConnectorServiceId"`
	// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
	DisplayName *string `pulumi:"displayName"`
	// The details of the egress settings.
	Egress Egress `pulumi:"egress"`
	// The details of the ingress settings.
	Ingress  Ingress `pulumi:"ingress"`
	Location *string `pulumi:"location"`
	// Name of resource. The name is ignored during creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
	ValidateOnly *string `pulumi:"validateOnly"`
}

// The set of arguments for constructing a ClientConnectorService resource.
type ClientConnectorServiceArgs struct {
	// Optional. User-settable client connector service resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter. A random system generated name will be assigned if not specified by the user.
	ClientConnectorServiceId pulumi.StringPtrInput
	// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
	DisplayName pulumi.StringPtrInput
	// The details of the egress settings.
	Egress EgressInput
	// The details of the ingress settings.
	Ingress  IngressInput
	Location pulumi.StringPtrInput
	// Name of resource. The name is ignored during creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
	ValidateOnly pulumi.StringPtrInput
}

func (ClientConnectorServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientConnectorServiceArgs)(nil)).Elem()
}

type ClientConnectorServiceInput interface {
	pulumi.Input

	ToClientConnectorServiceOutput() ClientConnectorServiceOutput
	ToClientConnectorServiceOutputWithContext(ctx context.Context) ClientConnectorServiceOutput
}

func (*ClientConnectorService) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientConnectorService)(nil)).Elem()
}

func (i *ClientConnectorService) ToClientConnectorServiceOutput() ClientConnectorServiceOutput {
	return i.ToClientConnectorServiceOutputWithContext(context.Background())
}

func (i *ClientConnectorService) ToClientConnectorServiceOutputWithContext(ctx context.Context) ClientConnectorServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientConnectorServiceOutput)
}

type ClientConnectorServiceOutput struct{ *pulumi.OutputState }

func (ClientConnectorServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientConnectorService)(nil)).Elem()
}

func (o ClientConnectorServiceOutput) ToClientConnectorServiceOutput() ClientConnectorServiceOutput {
	return o
}

func (o ClientConnectorServiceOutput) ToClientConnectorServiceOutputWithContext(ctx context.Context) ClientConnectorServiceOutput {
	return o
}

// Optional. User-settable client connector service resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter. A random system generated name will be assigned if not specified by the user.
func (o ClientConnectorServiceOutput) ClientConnectorServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringPtrOutput { return v.ClientConnectorServiceId }).(pulumi.StringPtrOutput)
}

// [Output only] Create time stamp.
func (o ClientConnectorServiceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
func (o ClientConnectorServiceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The details of the egress settings.
func (o ClientConnectorServiceOutput) Egress() EgressResponseOutput {
	return o.ApplyT(func(v *ClientConnectorService) EgressResponseOutput { return v.Egress }).(EgressResponseOutput)
}

// The details of the ingress settings.
func (o ClientConnectorServiceOutput) Ingress() IngressResponseOutput {
	return o.ApplyT(func(v *ClientConnectorService) IngressResponseOutput { return v.Ingress }).(IngressResponseOutput)
}

func (o ClientConnectorServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of resource. The name is ignored during creation.
func (o ClientConnectorServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientConnectorServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o ClientConnectorServiceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The operational state of the ClientConnectorService.
func (o ClientConnectorServiceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// [Output only] Update time stamp.
func (o ClientConnectorServiceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
func (o ClientConnectorServiceOutput) ValidateOnly() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientConnectorService) pulumi.StringPtrOutput { return v.ValidateOnly }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientConnectorServiceInput)(nil)).Elem(), &ClientConnectorService{})
	pulumi.RegisterOutputType(ClientConnectorServiceOutput{})
}
