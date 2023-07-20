// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Experiment in the specified Environment.
type Experiment struct {
	pulumi.CustomResourceState

	AgentId pulumi.StringOutput `pulumi:"agentId"`
	// Creation time of this experiment.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The definition of the experiment.
	Definition GoogleCloudDialogflowCxV3ExperimentDefinitionResponseOutput `pulumi:"definition"`
	// The human-readable description of the experiment.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// End time of this experiment.
	EndTime       pulumi.StringOutput `pulumi:"endTime"`
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength pulumi.StringOutput `pulumi:"experimentLength"`
	// Last update time of this experiment.
	LastUpdateTime pulumi.StringOutput `pulumi:"lastUpdateTime"`
	Location       pulumi.StringOutput `pulumi:"location"`
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Inference result of the experiment.
	Result GoogleCloudDialogflowCxV3ExperimentResultResponseOutput `pulumi:"result"`
	// The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
	RolloutConfig GoogleCloudDialogflowCxV3RolloutConfigResponseOutput `pulumi:"rolloutConfig"`
	// The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
	RolloutFailureReason pulumi.StringOutput `pulumi:"rolloutFailureReason"`
	// State of the auto rollout process.
	RolloutState GoogleCloudDialogflowCxV3RolloutStateResponseOutput `pulumi:"rolloutState"`
	// Start time of this experiment.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
	State pulumi.StringOutput `pulumi:"state"`
	// The history of updates to the experiment variants.
	VariantsHistory GoogleCloudDialogflowCxV3VariantsHistoryResponseArrayOutput `pulumi:"variantsHistory"`
}

// NewExperiment registers a new resource with the given unique name, arguments, and options.
func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"agentId",
		"environmentId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Experiment
	err := ctx.RegisterResource("google-native:dialogflow/v3:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperiment gets an existing Experiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("google-native:dialogflow/v3:Experiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Experiment resources.
type experimentState struct {
}

type ExperimentState struct {
}

func (ExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentState)(nil)).Elem()
}

type experimentArgs struct {
	AgentId string `pulumi:"agentId"`
	// Creation time of this experiment.
	CreateTime *string `pulumi:"createTime"`
	// The definition of the experiment.
	Definition *GoogleCloudDialogflowCxV3ExperimentDefinition `pulumi:"definition"`
	// The human-readable description of the experiment.
	Description *string `pulumi:"description"`
	// The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName string `pulumi:"displayName"`
	// End time of this experiment.
	EndTime       *string `pulumi:"endTime"`
	EnvironmentId string  `pulumi:"environmentId"`
	// Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength *string `pulumi:"experimentLength"`
	// Last update time of this experiment.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`
	Location       *string `pulumi:"location"`
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Inference result of the experiment.
	Result *GoogleCloudDialogflowCxV3ExperimentResult `pulumi:"result"`
	// The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
	RolloutConfig *GoogleCloudDialogflowCxV3RolloutConfig `pulumi:"rolloutConfig"`
	// The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
	RolloutFailureReason *string `pulumi:"rolloutFailureReason"`
	// State of the auto rollout process.
	RolloutState *GoogleCloudDialogflowCxV3RolloutState `pulumi:"rolloutState"`
	// Start time of this experiment.
	StartTime *string `pulumi:"startTime"`
	// The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
	State *ExperimentStateEnum `pulumi:"state"`
	// The history of updates to the experiment variants.
	VariantsHistory []GoogleCloudDialogflowCxV3VariantsHistory `pulumi:"variantsHistory"`
}

// The set of arguments for constructing a Experiment resource.
type ExperimentArgs struct {
	AgentId pulumi.StringInput
	// Creation time of this experiment.
	CreateTime pulumi.StringPtrInput
	// The definition of the experiment.
	Definition GoogleCloudDialogflowCxV3ExperimentDefinitionPtrInput
	// The human-readable description of the experiment.
	Description pulumi.StringPtrInput
	// The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName pulumi.StringInput
	// End time of this experiment.
	EndTime       pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength pulumi.StringPtrInput
	// Last update time of this experiment.
	LastUpdateTime pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Inference result of the experiment.
	Result GoogleCloudDialogflowCxV3ExperimentResultPtrInput
	// The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
	RolloutConfig GoogleCloudDialogflowCxV3RolloutConfigPtrInput
	// The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
	RolloutFailureReason pulumi.StringPtrInput
	// State of the auto rollout process.
	RolloutState GoogleCloudDialogflowCxV3RolloutStatePtrInput
	// Start time of this experiment.
	StartTime pulumi.StringPtrInput
	// The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
	State ExperimentStateEnumPtrInput
	// The history of updates to the experiment variants.
	VariantsHistory GoogleCloudDialogflowCxV3VariantsHistoryArrayInput
}

func (ExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentArgs)(nil)).Elem()
}

type ExperimentInput interface {
	pulumi.Input

	ToExperimentOutput() ExperimentOutput
	ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput
}

func (*Experiment) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (i *Experiment) ToExperimentOutput() ExperimentOutput {
	return i.ToExperimentOutputWithContext(context.Background())
}

func (i *Experiment) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentOutput)
}

type ExperimentOutput struct{ *pulumi.OutputState }

func (ExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (o ExperimentOutput) ToExperimentOutput() ExperimentOutput {
	return o
}

func (o ExperimentOutput) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return o
}

func (o ExperimentOutput) AgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.AgentId }).(pulumi.StringOutput)
}

// Creation time of this experiment.
func (o ExperimentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The definition of the experiment.
func (o ExperimentOutput) Definition() GoogleCloudDialogflowCxV3ExperimentDefinitionResponseOutput {
	return o.ApplyT(func(v *Experiment) GoogleCloudDialogflowCxV3ExperimentDefinitionResponseOutput { return v.Definition }).(GoogleCloudDialogflowCxV3ExperimentDefinitionResponseOutput)
}

// The human-readable description of the experiment.
func (o ExperimentOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
func (o ExperimentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// End time of this experiment.
func (o ExperimentOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

func (o ExperimentOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
func (o ExperimentOutput) ExperimentLength() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.ExperimentLength }).(pulumi.StringOutput)
}

// Last update time of this experiment.
func (o ExperimentOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.LastUpdateTime }).(pulumi.StringOutput)
}

func (o ExperimentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
func (o ExperimentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExperimentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Inference result of the experiment.
func (o ExperimentOutput) Result() GoogleCloudDialogflowCxV3ExperimentResultResponseOutput {
	return o.ApplyT(func(v *Experiment) GoogleCloudDialogflowCxV3ExperimentResultResponseOutput { return v.Result }).(GoogleCloudDialogflowCxV3ExperimentResultResponseOutput)
}

// The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
func (o ExperimentOutput) RolloutConfig() GoogleCloudDialogflowCxV3RolloutConfigResponseOutput {
	return o.ApplyT(func(v *Experiment) GoogleCloudDialogflowCxV3RolloutConfigResponseOutput { return v.RolloutConfig }).(GoogleCloudDialogflowCxV3RolloutConfigResponseOutput)
}

// The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
func (o ExperimentOutput) RolloutFailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.RolloutFailureReason }).(pulumi.StringOutput)
}

// State of the auto rollout process.
func (o ExperimentOutput) RolloutState() GoogleCloudDialogflowCxV3RolloutStateResponseOutput {
	return o.ApplyT(func(v *Experiment) GoogleCloudDialogflowCxV3RolloutStateResponseOutput { return v.RolloutState }).(GoogleCloudDialogflowCxV3RolloutStateResponseOutput)
}

// Start time of this experiment.
func (o ExperimentOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
func (o ExperimentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The history of updates to the experiment variants.
func (o ExperimentOutput) VariantsHistory() GoogleCloudDialogflowCxV3VariantsHistoryResponseArrayOutput {
	return o.ApplyT(func(v *Experiment) GoogleCloudDialogflowCxV3VariantsHistoryResponseArrayOutput {
		return v.VariantsHistory
	}).(GoogleCloudDialogflowCxV3VariantsHistoryResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentInput)(nil)).Elem(), &Experiment{})
	pulumi.RegisterOutputType(ExperimentOutput{})
}
