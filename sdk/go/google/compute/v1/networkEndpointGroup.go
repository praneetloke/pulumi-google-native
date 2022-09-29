// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a network endpoint group in the specified project using the parameters that are included in the request.
type NetworkEndpointGroup struct {
	pulumi.CustomResourceState

	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEngineResponseOutput `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionResponseOutput `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunResponseOutput `pulumi:"cloudRun"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntOutput `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network pulumi.StringOutput `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType pulumi.StringOutput                       `pulumi:"networkEndpointType"`
	Project             pulumi.StringOutput                       `pulumi:"project"`
	PscData             NetworkEndpointGroupPscDataResponseOutput `pulumi:"pscData"`
	// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService pulumi.StringOutput `pulumi:"pscTargetService"`
	// The URL of the region where the network endpoint group is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size pulumi.IntOutput `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	Zone       pulumi.StringOutput `pulumi:"zone"`
}

// NewNetworkEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkEndpointGroup(ctx *pulumi.Context,
	name string, args *NetworkEndpointGroupArgs, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	if args == nil {
		args = &NetworkEndpointGroupArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"zone",
	})
	opts = append(opts, replaceOnChanges)
	var resource NetworkEndpointGroup
	err := ctx.RegisterResource("google-native:compute/v1:NetworkEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkEndpointGroup gets an existing NetworkEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkEndpointGroupState, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	var resource NetworkEndpointGroup
	err := ctx.ReadResource("google-native:compute/v1:NetworkEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkEndpointGroup resources.
type networkEndpointGroupState struct {
}

type NetworkEndpointGroupState struct {
}

func (NetworkEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupState)(nil)).Elem()
}

type networkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine *NetworkEndpointGroupAppEngine `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction *NetworkEndpointGroupCloudFunction `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun *NetworkEndpointGroupCloudRun `pulumi:"cloudRun"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network *string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType *NetworkEndpointGroupNetworkEndpointType `pulumi:"networkEndpointType"`
	Project             *string                                  `pulumi:"project"`
	PscData             *NetworkEndpointGroupPscData             `pulumi:"pscData"`
	// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService *string `pulumi:"pscTargetService"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
	Zone       *string `pulumi:"zone"`
}

// The set of arguments for constructing a NetworkEndpointGroup resource.
type NetworkEndpointGroupArgs struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations pulumi.StringMapInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEnginePtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionPtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunPtrInput
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType NetworkEndpointGroupNetworkEndpointTypePtrInput
	Project             pulumi.StringPtrInput
	PscData             NetworkEndpointGroupPscDataPtrInput
	// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
	Zone       pulumi.StringPtrInput
}

func (NetworkEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupArgs)(nil)).Elem()
}

type NetworkEndpointGroupInput interface {
	pulumi.Input

	ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput
	ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput
}

func (*NetworkEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEndpointGroup)(nil)).Elem()
}

func (i *NetworkEndpointGroup) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return i.ToNetworkEndpointGroupOutputWithContext(context.Background())
}

func (i *NetworkEndpointGroup) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointGroupOutput)
}

type NetworkEndpointGroupOutput struct{ *pulumi.OutputState }

func (NetworkEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEndpointGroup)(nil)).Elem()
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return o
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return o
}

// Metadata defined as annotations on the network endpoint group.
func (o NetworkEndpointGroupOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o NetworkEndpointGroupOutput) AppEngine() NetworkEndpointGroupAppEngineResponseOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) NetworkEndpointGroupAppEngineResponseOutput { return v.AppEngine }).(NetworkEndpointGroupAppEngineResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o NetworkEndpointGroupOutput) CloudFunction() NetworkEndpointGroupCloudFunctionResponseOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) NetworkEndpointGroupCloudFunctionResponseOutput { return v.CloudFunction }).(NetworkEndpointGroupCloudFunctionResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o NetworkEndpointGroupOutput) CloudRun() NetworkEndpointGroupCloudRunResponseOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) NetworkEndpointGroupCloudRunResponseOutput { return v.CloudRun }).(NetworkEndpointGroupCloudRunResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o NetworkEndpointGroupOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The default port used if the port number is not specified in the network endpoint.
func (o NetworkEndpointGroupOutput) DefaultPort() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.IntOutput { return v.DefaultPort }).(pulumi.IntOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o NetworkEndpointGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
func (o NetworkEndpointGroupOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o NetworkEndpointGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
func (o NetworkEndpointGroupOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
func (o NetworkEndpointGroupOutput) NetworkEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.NetworkEndpointType }).(pulumi.StringOutput)
}

func (o NetworkEndpointGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NetworkEndpointGroupOutput) PscData() NetworkEndpointGroupPscDataResponseOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) NetworkEndpointGroupPscDataResponseOutput { return v.PscData }).(NetworkEndpointGroupPscDataResponseOutput)
}

// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
func (o NetworkEndpointGroupOutput) PscTargetService() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.PscTargetService }).(pulumi.StringOutput)
}

// The URL of the region where the network endpoint group is located.
func (o NetworkEndpointGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o NetworkEndpointGroupOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o NetworkEndpointGroupOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// [Output only] Number of network endpoints in the network endpoint group.
func (o NetworkEndpointGroupOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
func (o NetworkEndpointGroupOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Subnetwork }).(pulumi.StringOutput)
}

func (o NetworkEndpointGroupOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEndpointGroup) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkEndpointGroupInput)(nil)).Elem(), &NetworkEndpointGroup{})
	pulumi.RegisterOutputType(NetworkEndpointGroupOutput{})
}
