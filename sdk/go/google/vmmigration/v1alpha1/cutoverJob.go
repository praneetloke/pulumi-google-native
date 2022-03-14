// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Initiates a Cutover of a specific migrating VM. The returned LRO is completed when the cutover job resource is created and the job is initiated.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type CutoverJob struct {
	pulumi.CustomResourceState

	// Details of the target VM in Compute Engine.
	ComputeEngineTargetDetails ComputeEngineTargetDetailsResponseOutput `pulumi:"computeEngineTargetDetails"`
	// The time the cutover job was created (as an API call, not when it was actually created in the target).
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Provides details for the errors that led to the Cutover Job's state.
	Error StatusResponseOutput `pulumi:"error"`
	// The name of the cutover job.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current progress in percentage of the cutover job.
	Progress pulumi.IntOutput `pulumi:"progress"`
	// The current progress in percentage of the cutover job.
	ProgressPercent pulumi.IntOutput `pulumi:"progressPercent"`
	// State of the cutover job.
	State pulumi.StringOutput `pulumi:"state"`
	// A message providing possible extra details about the current state.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The time the state was last updated.
	StateTime pulumi.StringOutput `pulumi:"stateTime"`
}

// NewCutoverJob registers a new resource with the given unique name, arguments, and options.
func NewCutoverJob(ctx *pulumi.Context,
	name string, args *CutoverJobArgs, opts ...pulumi.ResourceOption) (*CutoverJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CutoverJobId == nil {
		return nil, errors.New("invalid value for required argument 'CutoverJobId'")
	}
	if args.MigratingVmId == nil {
		return nil, errors.New("invalid value for required argument 'MigratingVmId'")
	}
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	var resource CutoverJob
	err := ctx.RegisterResource("google-native:vmmigration/v1alpha1:CutoverJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCutoverJob gets an existing CutoverJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCutoverJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CutoverJobState, opts ...pulumi.ResourceOption) (*CutoverJob, error) {
	var resource CutoverJob
	err := ctx.ReadResource("google-native:vmmigration/v1alpha1:CutoverJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CutoverJob resources.
type cutoverJobState struct {
}

type CutoverJobState struct {
}

func (CutoverJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cutoverJobState)(nil)).Elem()
}

type cutoverJobArgs struct {
	// Required. The cutover job identifier.
	CutoverJobId  string  `pulumi:"cutoverJobId"`
	Location      *string `pulumi:"location"`
	MigratingVmId string  `pulumi:"migratingVmId"`
	Project       *string `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	SourceId  string  `pulumi:"sourceId"`
}

// The set of arguments for constructing a CutoverJob resource.
type CutoverJobArgs struct {
	// Required. The cutover job identifier.
	CutoverJobId  pulumi.StringInput
	Location      pulumi.StringPtrInput
	MigratingVmId pulumi.StringInput
	Project       pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	SourceId  pulumi.StringInput
}

func (CutoverJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cutoverJobArgs)(nil)).Elem()
}

type CutoverJobInput interface {
	pulumi.Input

	ToCutoverJobOutput() CutoverJobOutput
	ToCutoverJobOutputWithContext(ctx context.Context) CutoverJobOutput
}

func (*CutoverJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CutoverJob)(nil)).Elem()
}

func (i *CutoverJob) ToCutoverJobOutput() CutoverJobOutput {
	return i.ToCutoverJobOutputWithContext(context.Background())
}

func (i *CutoverJob) ToCutoverJobOutputWithContext(ctx context.Context) CutoverJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CutoverJobOutput)
}

type CutoverJobOutput struct{ *pulumi.OutputState }

func (CutoverJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CutoverJob)(nil)).Elem()
}

func (o CutoverJobOutput) ToCutoverJobOutput() CutoverJobOutput {
	return o
}

func (o CutoverJobOutput) ToCutoverJobOutputWithContext(ctx context.Context) CutoverJobOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CutoverJobInput)(nil)).Elem(), &CutoverJob{})
	pulumi.RegisterOutputType(CutoverJobOutput{})
}
