// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TcpRoute in a given project and location.
type TcpRoute struct {
	pulumi.CustomResourceState

	// The timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
	Gateways pulumi.StringArrayOutput `pulumi:"gateways"`
	// Optional. Set of label tags associated with the TcpRoute resource.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayOutput `pulumi:"meshes"`
	// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
	Rules TcpRouteRouteRuleResponseArrayOutput `pulumi:"rules"`
	// Server-defined URL of this resource
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Required. Short name of the TcpRoute resource to be created. E.g. TODO(Add an example).
	TcpRouteId pulumi.StringOutput `pulumi:"tcpRouteId"`
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTcpRoute registers a new resource with the given unique name, arguments, and options.
func NewTcpRoute(ctx *pulumi.Context,
	name string, args *TcpRouteArgs, opts ...pulumi.ResourceOption) (*TcpRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.TcpRouteId == nil {
		return nil, errors.New("invalid value for required argument 'TcpRouteId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"tcpRouteId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TcpRoute
	err := ctx.RegisterResource("google-native:networkservices/v1beta1:TcpRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTcpRoute gets an existing TcpRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTcpRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TcpRouteState, opts ...pulumi.ResourceOption) (*TcpRoute, error) {
	var resource TcpRoute
	err := ctx.ReadResource("google-native:networkservices/v1beta1:TcpRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TcpRoute resources.
type tcpRouteState struct {
}

type TcpRouteState struct {
}

func (TcpRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteState)(nil)).Elem()
}

type tcpRouteArgs struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
	Gateways []string `pulumi:"gateways"`
	// Optional. Set of label tags associated with the TcpRoute resource.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
	Rules []TcpRouteRouteRule `pulumi:"rules"`
	// Required. Short name of the TcpRoute resource to be created. E.g. TODO(Add an example).
	TcpRouteId string `pulumi:"tcpRouteId"`
}

// The set of arguments for constructing a TcpRoute resource.
type TcpRouteArgs struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
	Gateways pulumi.StringArrayInput
	// Optional. Set of label tags associated with the TcpRoute resource.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayInput
	// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
	Rules TcpRouteRouteRuleArrayInput
	// Required. Short name of the TcpRoute resource to be created. E.g. TODO(Add an example).
	TcpRouteId pulumi.StringInput
}

func (TcpRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteArgs)(nil)).Elem()
}

type TcpRouteInput interface {
	pulumi.Input

	ToTcpRouteOutput() TcpRouteOutput
	ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput
}

func (*TcpRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRoute)(nil)).Elem()
}

func (i *TcpRoute) ToTcpRouteOutput() TcpRouteOutput {
	return i.ToTcpRouteOutputWithContext(context.Background())
}

func (i *TcpRoute) ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteOutput)
}

type TcpRouteOutput struct{ *pulumi.OutputState }

func (TcpRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRoute)(nil)).Elem()
}

func (o TcpRouteOutput) ToTcpRouteOutput() TcpRouteOutput {
	return o
}

func (o TcpRouteOutput) ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput {
	return o
}

// The timestamp when the resource was created.
func (o TcpRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A free-text description of the resource. Max length 1024 characters.
func (o TcpRouteOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
func (o TcpRouteOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringArrayOutput { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Optional. Set of label tags associated with the TcpRoute resource.
func (o TcpRouteOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o TcpRouteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
func (o TcpRouteOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringArrayOutput { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
func (o TcpRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TcpRouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
func (o TcpRouteOutput) Rules() TcpRouteRouteRuleResponseArrayOutput {
	return o.ApplyT(func(v *TcpRoute) TcpRouteRouteRuleResponseArrayOutput { return v.Rules }).(TcpRouteRouteRuleResponseArrayOutput)
}

// Server-defined URL of this resource
func (o TcpRouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Required. Short name of the TcpRoute resource to be created. E.g. TODO(Add an example).
func (o TcpRouteOutput) TcpRouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.TcpRouteId }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o TcpRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteInput)(nil)).Elem(), &TcpRoute{})
	pulumi.RegisterOutputType(TcpRouteOutput{})
}
