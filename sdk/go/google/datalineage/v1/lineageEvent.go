// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new lineage event.
type LineageEvent struct {
	pulumi.CustomResourceState

	// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Optional. List of source-target pairs. Can't contain more than 100 tuples.
	Links    GoogleCloudDatacatalogLineageV1EventLinkResponseArrayOutput `pulumi:"links"`
	Location pulumi.StringOutput                                         `pulumi:"location"`
	// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name      pulumi.StringOutput `pulumi:"name"`
	ProcessId pulumi.StringOutput `pulumi:"processId"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	RunId     pulumi.StringOutput    `pulumi:"runId"`
	// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewLineageEvent registers a new resource with the given unique name, arguments, and options.
func NewLineageEvent(ctx *pulumi.Context,
	name string, args *LineageEventArgs, opts ...pulumi.ResourceOption) (*LineageEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProcessId == nil {
		return nil, errors.New("invalid value for required argument 'ProcessId'")
	}
	if args.RunId == nil {
		return nil, errors.New("invalid value for required argument 'RunId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"processId",
		"project",
		"runId",
	})
	opts = append(opts, replaceOnChanges)
	var resource LineageEvent
	err := ctx.RegisterResource("google-native:datalineage/v1:LineageEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLineageEvent gets an existing LineageEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLineageEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LineageEventState, opts ...pulumi.ResourceOption) (*LineageEvent, error) {
	var resource LineageEvent
	err := ctx.ReadResource("google-native:datalineage/v1:LineageEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LineageEvent resources.
type lineageEventState struct {
}

type LineageEventState struct {
}

func (LineageEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*lineageEventState)(nil)).Elem()
}

type lineageEventArgs struct {
	// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
	EndTime *string `pulumi:"endTime"`
	// Optional. List of source-target pairs. Can't contain more than 100 tuples.
	Links    []GoogleCloudDatacatalogLineageV1EventLink `pulumi:"links"`
	Location *string                                    `pulumi:"location"`
	// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name      *string `pulumi:"name"`
	ProcessId string  `pulumi:"processId"`
	Project   *string `pulumi:"project"`
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId *string `pulumi:"requestId"`
	RunId     string  `pulumi:"runId"`
	// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a LineageEvent resource.
type LineageEventArgs struct {
	// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
	EndTime pulumi.StringPtrInput
	// Optional. List of source-target pairs. Can't contain more than 100 tuples.
	Links    GoogleCloudDatacatalogLineageV1EventLinkArrayInput
	Location pulumi.StringPtrInput
	// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name      pulumi.StringPtrInput
	ProcessId pulumi.StringInput
	Project   pulumi.StringPtrInput
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId pulumi.StringPtrInput
	RunId     pulumi.StringInput
	// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
	StartTime pulumi.StringPtrInput
}

func (LineageEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lineageEventArgs)(nil)).Elem()
}

type LineageEventInput interface {
	pulumi.Input

	ToLineageEventOutput() LineageEventOutput
	ToLineageEventOutputWithContext(ctx context.Context) LineageEventOutput
}

func (*LineageEvent) ElementType() reflect.Type {
	return reflect.TypeOf((**LineageEvent)(nil)).Elem()
}

func (i *LineageEvent) ToLineageEventOutput() LineageEventOutput {
	return i.ToLineageEventOutputWithContext(context.Background())
}

func (i *LineageEvent) ToLineageEventOutputWithContext(ctx context.Context) LineageEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineageEventOutput)
}

type LineageEventOutput struct{ *pulumi.OutputState }

func (LineageEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LineageEvent)(nil)).Elem()
}

func (o LineageEventOutput) ToLineageEventOutput() LineageEventOutput {
	return o
}

func (o LineageEventOutput) ToLineageEventOutputWithContext(ctx context.Context) LineageEventOutput {
	return o
}

// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
func (o LineageEventOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Optional. List of source-target pairs. Can't contain more than 100 tuples.
func (o LineageEventOutput) Links() GoogleCloudDatacatalogLineageV1EventLinkResponseArrayOutput {
	return o.ApplyT(func(v *LineageEvent) GoogleCloudDatacatalogLineageV1EventLinkResponseArrayOutput { return v.Links }).(GoogleCloudDatacatalogLineageV1EventLinkResponseArrayOutput)
}

func (o LineageEventOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
func (o LineageEventOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LineageEventOutput) ProcessId() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.ProcessId }).(pulumi.StringOutput)
}

func (o LineageEventOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
func (o LineageEventOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

func (o LineageEventOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.RunId }).(pulumi.StringOutput)
}

// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
func (o LineageEventOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LineageEvent) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LineageEventInput)(nil)).Elem(), &LineageEvent{})
	pulumi.RegisterOutputType(LineageEventOutput{})
}
