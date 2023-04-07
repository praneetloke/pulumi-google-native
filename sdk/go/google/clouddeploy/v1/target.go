// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Target in a given project and location.
type Target struct {
	pulumi.CustomResourceState

	// Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Information specifying an Anthos Cluster.
	AnthosCluster AnthosClusterResponseOutput `pulumi:"anthosCluster"`
	// Time at which the `Target` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the `Target`. Max length is 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
	ExecutionConfigs ExecutionConfigResponseArrayOutput `pulumi:"executionConfigs"`
	// Information specifying a GKE Cluster.
	Gke GkeClusterResponseOutput `pulumi:"gke"`
	// Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Information specifying a multiTarget.
	MultiTarget MultiTargetResponseOutput `pulumi:"multiTarget"`
	// Optional. Name of the `Target`. Format is projects/{project}/locations/{location}/targets/a-z{0,62}.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Optional. Whether or not the `Target` requires approval.
	RequireApproval pulumi.BoolOutput `pulumi:"requireApproval"`
	// Information specifying a Cloud Run deployment target.
	Run CloudRunLocationResponseOutput `pulumi:"run"`
	// Required. ID of the `Target`.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// Unique identifier of the `Target`.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Most recent time at which the `Target` was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Optional. If set to true, the request is validated and the user is provided with an expected result, but no actual change is made.
	ValidateOnly pulumi.BoolPtrOutput `pulumi:"validateOnly"`
}

// NewTarget registers a new resource with the given unique name, arguments, and options.
func NewTarget(ctx *pulumi.Context,
	name string, args *TargetArgs, opts ...pulumi.ResourceOption) (*Target, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"targetId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Target
	err := ctx.RegisterResource("google-native:clouddeploy/v1:Target", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTarget gets an existing Target resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetState, opts ...pulumi.ResourceOption) (*Target, error) {
	var resource Target
	err := ctx.ReadResource("google-native:clouddeploy/v1:Target", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Target resources.
type targetState struct {
}

type TargetState struct {
}

func (TargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetState)(nil)).Elem()
}

type targetArgs struct {
	// Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations map[string]string `pulumi:"annotations"`
	// Information specifying an Anthos Cluster.
	AnthosCluster *AnthosCluster `pulumi:"anthosCluster"`
	// Optional. Description of the `Target`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
	ExecutionConfigs []ExecutionConfig `pulumi:"executionConfigs"`
	// Information specifying a GKE Cluster.
	Gke *GkeCluster `pulumi:"gke"`
	// Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Information specifying a multiTarget.
	MultiTarget *MultiTarget `pulumi:"multiTarget"`
	// Optional. Name of the `Target`. Format is projects/{project}/locations/{location}/targets/a-z{0,62}.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. Whether or not the `Target` requires approval.
	RequireApproval *bool `pulumi:"requireApproval"`
	// Information specifying a Cloud Run deployment target.
	Run *CloudRunLocation `pulumi:"run"`
	// Required. ID of the `Target`.
	TargetId string `pulumi:"targetId"`
	// Optional. If set to true, the request is validated and the user is provided with an expected result, but no actual change is made.
	ValidateOnly *bool `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Target resource.
type TargetArgs struct {
	// Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations pulumi.StringMapInput
	// Information specifying an Anthos Cluster.
	AnthosCluster AnthosClusterPtrInput
	// Optional. Description of the `Target`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
	ExecutionConfigs ExecutionConfigArrayInput
	// Information specifying a GKE Cluster.
	Gke GkeClusterPtrInput
	// Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Information specifying a multiTarget.
	MultiTarget MultiTargetPtrInput
	// Optional. Name of the `Target`. Format is projects/{project}/locations/{location}/targets/a-z{0,62}.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. Whether or not the `Target` requires approval.
	RequireApproval pulumi.BoolPtrInput
	// Information specifying a Cloud Run deployment target.
	Run CloudRunLocationPtrInput
	// Required. ID of the `Target`.
	TargetId pulumi.StringInput
	// Optional. If set to true, the request is validated and the user is provided with an expected result, but no actual change is made.
	ValidateOnly pulumi.BoolPtrInput
}

func (TargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetArgs)(nil)).Elem()
}

type TargetInput interface {
	pulumi.Input

	ToTargetOutput() TargetOutput
	ToTargetOutputWithContext(ctx context.Context) TargetOutput
}

func (*Target) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (i *Target) ToTargetOutput() TargetOutput {
	return i.ToTargetOutputWithContext(context.Background())
}

func (i *Target) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOutput)
}

type TargetOutput struct{ *pulumi.OutputState }

func (TargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (o TargetOutput) ToTargetOutput() TargetOutput {
	return o
}

func (o TargetOutput) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return o
}

// Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
func (o TargetOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Target) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Information specifying an Anthos Cluster.
func (o TargetOutput) AnthosCluster() AnthosClusterResponseOutput {
	return o.ApplyT(func(v *Target) AnthosClusterResponseOutput { return v.AnthosCluster }).(AnthosClusterResponseOutput)
}

// Time at which the `Target` was created.
func (o TargetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the `Target`. Max length is 255 characters.
func (o TargetOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o TargetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
func (o TargetOutput) ExecutionConfigs() ExecutionConfigResponseArrayOutput {
	return o.ApplyT(func(v *Target) ExecutionConfigResponseArrayOutput { return v.ExecutionConfigs }).(ExecutionConfigResponseArrayOutput)
}

// Information specifying a GKE Cluster.
func (o TargetOutput) Gke() GkeClusterResponseOutput {
	return o.ApplyT(func(v *Target) GkeClusterResponseOutput { return v.Gke }).(GkeClusterResponseOutput)
}

// Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
func (o TargetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Target) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o TargetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Information specifying a multiTarget.
func (o TargetOutput) MultiTarget() MultiTargetResponseOutput {
	return o.ApplyT(func(v *Target) MultiTargetResponseOutput { return v.MultiTarget }).(MultiTargetResponseOutput)
}

// Optional. Name of the `Target`. Format is projects/{project}/locations/{location}/targets/a-z{0,62}.
func (o TargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TargetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o TargetOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Target) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Optional. Whether or not the `Target` requires approval.
func (o TargetOutput) RequireApproval() pulumi.BoolOutput {
	return o.ApplyT(func(v *Target) pulumi.BoolOutput { return v.RequireApproval }).(pulumi.BoolOutput)
}

// Information specifying a Cloud Run deployment target.
func (o TargetOutput) Run() CloudRunLocationResponseOutput {
	return o.ApplyT(func(v *Target) CloudRunLocationResponseOutput { return v.Run }).(CloudRunLocationResponseOutput)
}

// Required. ID of the `Target`.
func (o TargetOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// Unique identifier of the `Target`.
func (o TargetOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Most recent time at which the `Target` was updated.
func (o TargetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. If set to true, the request is validated and the user is provided with an expected result, but no actual change is made.
func (o TargetOutput) ValidateOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Target) pulumi.BoolPtrOutput { return v.ValidateOnly }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetInput)(nil)).Elem(), &Target{})
	pulumi.RegisterOutputType(TargetOutput{})
}
