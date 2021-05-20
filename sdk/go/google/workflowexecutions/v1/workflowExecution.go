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
type WorkflowExecution struct {
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

// NewWorkflowExecution registers a new resource with the given unique name, arguments, and options.
func NewWorkflowExecution(ctx *pulumi.Context,
	name string, args *WorkflowExecutionArgs, opts ...pulumi.ResourceOption) (*WorkflowExecution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionId == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.WorkflowId == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowId'")
	}
	var resource WorkflowExecution
	err := ctx.RegisterResource("google-native:workflowexecutions/v1:WorkflowExecution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowExecution gets an existing WorkflowExecution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowExecutionState, opts ...pulumi.ResourceOption) (*WorkflowExecution, error) {
	var resource WorkflowExecution
	err := ctx.ReadResource("google-native:workflowexecutions/v1:WorkflowExecution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowExecution resources.
type workflowExecutionState struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument *string `pulumi:"argument"`
	// Marks the end of execution, successful or not.
	EndTime *string `pulumi:"endTime"`
	// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
	Error *ErrorResponse `pulumi:"error"`
	// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	Name *string `pulumi:"name"`
	// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
	Result *string `pulumi:"result"`
	// Marks the beginning of execution.
	StartTime *string `pulumi:"startTime"`
	// Current state of the execution.
	State *string `pulumi:"state"`
	// Revision of the workflow this execution is using.
	WorkflowRevisionId *string `pulumi:"workflowRevisionId"`
}

type WorkflowExecutionState struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument pulumi.StringPtrInput
	// Marks the end of execution, successful or not.
	EndTime pulumi.StringPtrInput
	// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
	Error ErrorResponsePtrInput
	// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	Name pulumi.StringPtrInput
	// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
	Result pulumi.StringPtrInput
	// Marks the beginning of execution.
	StartTime pulumi.StringPtrInput
	// Current state of the execution.
	State pulumi.StringPtrInput
	// Revision of the workflow this execution is using.
	WorkflowRevisionId pulumi.StringPtrInput
}

func (WorkflowExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowExecutionState)(nil)).Elem()
}

type workflowExecutionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument    *string `pulumi:"argument"`
	ExecutionId string  `pulumi:"executionId"`
	Location    string  `pulumi:"location"`
	Project     string  `pulumi:"project"`
	WorkflowId  string  `pulumi:"workflowId"`
}

// The set of arguments for constructing a WorkflowExecution resource.
type WorkflowExecutionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument    pulumi.StringPtrInput
	ExecutionId pulumi.StringInput
	Location    pulumi.StringInput
	Project     pulumi.StringInput
	WorkflowId  pulumi.StringInput
}

func (WorkflowExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowExecutionArgs)(nil)).Elem()
}

type WorkflowExecutionInput interface {
	pulumi.Input

	ToWorkflowExecutionOutput() WorkflowExecutionOutput
	ToWorkflowExecutionOutputWithContext(ctx context.Context) WorkflowExecutionOutput
}

func (*WorkflowExecution) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowExecution)(nil))
}

func (i *WorkflowExecution) ToWorkflowExecutionOutput() WorkflowExecutionOutput {
	return i.ToWorkflowExecutionOutputWithContext(context.Background())
}

func (i *WorkflowExecution) ToWorkflowExecutionOutputWithContext(ctx context.Context) WorkflowExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowExecutionOutput)
}

type WorkflowExecutionOutput struct {
	*pulumi.OutputState
}

func (WorkflowExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowExecution)(nil))
}

func (o WorkflowExecutionOutput) ToWorkflowExecutionOutput() WorkflowExecutionOutput {
	return o
}

func (o WorkflowExecutionOutput) ToWorkflowExecutionOutputWithContext(ctx context.Context) WorkflowExecutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkflowExecutionOutput{})
}
