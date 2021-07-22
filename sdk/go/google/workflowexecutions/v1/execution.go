// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new execution using the latest revision of the given workflow.
// Auto-naming is currently not supported for this resource.
type Execution struct {
	pulumi.CustomResourceState

	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument pulumi.StringOutput `pulumi:"argument"`
	// Marks the end of execution, successful or not.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
	Error ErrorResponseOutput `pulumi:"error"`
	// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	Name pulumi.StringOutput `pulumi:"name"`
	// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
	Result pulumi.StringOutput `pulumi:"result"`
	// Marks the beginning of execution.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Current state of the execution.
	State pulumi.StringOutput `pulumi:"state"`
	// Revision of the workflow this execution is using.
	WorkflowRevisionId pulumi.StringOutput `pulumi:"workflowRevisionId"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkflowId == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowId'")
	}
	var resource Execution
	err := ctx.RegisterResource("google-native:workflowexecutions/v1:Execution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecution gets an existing Execution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionState, opts ...pulumi.ResourceOption) (*Execution, error) {
	var resource Execution
	err := ctx.ReadResource("google-native:workflowexecutions/v1:Execution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Execution resources.
type executionState struct {
}

type ExecutionState struct {
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument   *string `pulumi:"argument"`
	Location   *string `pulumi:"location"`
	Project    *string `pulumi:"project"`
	WorkflowId string  `pulumi:"workflowId"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument   pulumi.StringPtrInput
	Location   pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	WorkflowId pulumi.StringInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}

type ExecutionInput interface {
	pulumi.Input

	ToExecutionOutput() ExecutionOutput
	ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput
}

func (*Execution) ElementType() reflect.Type {
	return reflect.TypeOf((*Execution)(nil))
}

func (i *Execution) ToExecutionOutput() ExecutionOutput {
	return i.ToExecutionOutputWithContext(context.Background())
}

func (i *Execution) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionOutput)
}

type ExecutionOutput struct {
	*pulumi.OutputState
}

func (ExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Execution)(nil))
}

func (o ExecutionOutput) ToExecutionOutput() ExecutionOutput {
	return o
}

func (o ExecutionOutput) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExecutionOutput{})
}
