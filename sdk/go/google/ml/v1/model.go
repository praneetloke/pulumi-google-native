// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a model which will later contain one or more versions. You must add at least one version before you can request predictions from the model. Add versions by calling projects.models.versions.create.
type Model struct {
	pulumi.CustomResourceState

	// The default version of the model. This version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.models.versions.setDefault.
	DefaultVersion GoogleCloudMlV1__VersionResponseOutput `pulumi:"defaultVersion"`
	// Optional. The description specified for the model when it was created.
	Description pulumi.StringOutput `pulumi:"description"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name specified for the model when it was created. The model name must be unique within the project it is created in.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
	OnlinePredictionConsoleLogging pulumi.BoolOutput `pulumi:"onlinePredictionConsoleLogging"`
	// Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
	OnlinePredictionLogging pulumi.BoolOutput   `pulumi:"onlinePredictionLogging"`
	Project                 pulumi.StringOutput `pulumi:"project"`
	// Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		args = &ModelArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Model
	err := ctx.RegisterResource("google-native:ml/v1:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("google-native:ml/v1:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
}

type ModelState struct {
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	// Optional. The description specified for the model when it was created.
	Description *string `pulumi:"description"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
	Etag *string `pulumi:"etag"`
	// Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
	Labels map[string]string `pulumi:"labels"`
	// The name specified for the model when it was created. The model name must be unique within the project it is created in.
	Name *string `pulumi:"name"`
	// Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
	OnlinePredictionConsoleLogging *bool `pulumi:"onlinePredictionConsoleLogging"`
	// Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
	OnlinePredictionLogging *bool   `pulumi:"onlinePredictionLogging"`
	Project                 *string `pulumi:"project"`
	// Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
	Regions []string `pulumi:"regions"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// Optional. The description specified for the model when it was created.
	Description pulumi.StringPtrInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
	Etag pulumi.StringPtrInput
	// Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
	Labels pulumi.StringMapInput
	// The name specified for the model when it was created. The model name must be unique within the project it is created in.
	Name pulumi.StringPtrInput
	// Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
	OnlinePredictionConsoleLogging pulumi.BoolPtrInput
	// Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
	OnlinePredictionLogging pulumi.BoolPtrInput
	Project                 pulumi.StringPtrInput
	// Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
	Regions pulumi.StringArrayInput
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}

type ModelInput interface {
	pulumi.Input

	ToModelOutput() ModelOutput
	ToModelOutputWithContext(ctx context.Context) ModelOutput
}

func (*Model) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (i *Model) ToModelOutput() ModelOutput {
	return i.ToModelOutputWithContext(context.Background())
}

func (i *Model) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput)
}

type ModelOutput struct{ *pulumi.OutputState }

func (ModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (o ModelOutput) ToModelOutput() ModelOutput {
	return o
}

func (o ModelOutput) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return o
}

// The default version of the model. This version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.models.versions.setDefault.
func (o ModelOutput) DefaultVersion() GoogleCloudMlV1__VersionResponseOutput {
	return o.ApplyT(func(v *Model) GoogleCloudMlV1__VersionResponseOutput { return v.DefaultVersion }).(GoogleCloudMlV1__VersionResponseOutput)
}

// Optional. The description specified for the model when it was created.
func (o ModelOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
func (o ModelOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
func (o ModelOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Model) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name specified for the model when it was created. The model name must be unique within the project it is created in.
func (o ModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
func (o ModelOutput) OnlinePredictionConsoleLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v *Model) pulumi.BoolOutput { return v.OnlinePredictionConsoleLogging }).(pulumi.BoolOutput)
}

// Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
func (o ModelOutput) OnlinePredictionLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v *Model) pulumi.BoolOutput { return v.OnlinePredictionLogging }).(pulumi.BoolOutput)
}

func (o ModelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
func (o ModelOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Model) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelInput)(nil)).Elem(), &Model{})
	pulumi.RegisterOutputType(ModelOutput{})
}
