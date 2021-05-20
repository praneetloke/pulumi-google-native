// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a ServiceLevelObjective for the given Service.
type ServiceServiceLevelObjective struct {
	pulumi.CustomResourceState

	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod pulumi.StringOutput `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal pulumi.Float64Output `pulumi:"goal"`
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringOutput `pulumi:"name"`
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod pulumi.StringOutput `pulumi:"rollingPeriod"`
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator ServiceLevelIndicatorResponseOutput `pulumi:"serviceLevelIndicator"`
}

// NewServiceServiceLevelObjective registers a new resource with the given unique name, arguments, and options.
func NewServiceServiceLevelObjective(ctx *pulumi.Context,
	name string, args *ServiceServiceLevelObjectiveArgs, opts ...pulumi.ResourceOption) (*ServiceServiceLevelObjective, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.V3Id == nil {
		return nil, errors.New("invalid value for required argument 'V3Id'")
	}
	if args.V3Id1 == nil {
		return nil, errors.New("invalid value for required argument 'V3Id1'")
	}
	var resource ServiceServiceLevelObjective
	err := ctx.RegisterResource("google-native:monitoring/v3:ServiceServiceLevelObjective", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceServiceLevelObjective gets an existing ServiceServiceLevelObjective resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceServiceLevelObjective(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceServiceLevelObjectiveState, opts ...pulumi.ResourceOption) (*ServiceServiceLevelObjective, error) {
	var resource ServiceServiceLevelObjective
	err := ctx.ReadResource("google-native:monitoring/v3:ServiceServiceLevelObjective", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceServiceLevelObjective resources.
type serviceServiceLevelObjectiveState struct {
	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal *float64 `pulumi:"goal"`
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name *string `pulumi:"name"`
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod *string `pulumi:"rollingPeriod"`
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator *ServiceLevelIndicatorResponse `pulumi:"serviceLevelIndicator"`
}

type ServiceServiceLevelObjectiveState struct {
	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal pulumi.Float64PtrInput
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringPtrInput
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod pulumi.StringPtrInput
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator ServiceLevelIndicatorResponsePtrInput
}

func (ServiceServiceLevelObjectiveState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceServiceLevelObjectiveState)(nil)).Elem()
}

type serviceServiceLevelObjectiveArgs struct {
	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal *float64 `pulumi:"goal"`
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name *string `pulumi:"name"`
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod *string `pulumi:"rollingPeriod"`
	ServiceId     string  `pulumi:"serviceId"`
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator   *ServiceLevelIndicator `pulumi:"serviceLevelIndicator"`
	ServiceLevelObjectiveId *string                `pulumi:"serviceLevelObjectiveId"`
	V3Id                    string                 `pulumi:"v3Id"`
	V3Id1                   string                 `pulumi:"v3Id1"`
}

// The set of arguments for constructing a ServiceServiceLevelObjective resource.
type ServiceServiceLevelObjectiveArgs struct {
	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal pulumi.Float64PtrInput
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringPtrInput
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod pulumi.StringPtrInput
	ServiceId     pulumi.StringInput
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator   ServiceLevelIndicatorPtrInput
	ServiceLevelObjectiveId pulumi.StringPtrInput
	V3Id                    pulumi.StringInput
	V3Id1                   pulumi.StringInput
}

func (ServiceServiceLevelObjectiveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceServiceLevelObjectiveArgs)(nil)).Elem()
}

type ServiceServiceLevelObjectiveInput interface {
	pulumi.Input

	ToServiceServiceLevelObjectiveOutput() ServiceServiceLevelObjectiveOutput
	ToServiceServiceLevelObjectiveOutputWithContext(ctx context.Context) ServiceServiceLevelObjectiveOutput
}

func (*ServiceServiceLevelObjective) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceServiceLevelObjective)(nil))
}

func (i *ServiceServiceLevelObjective) ToServiceServiceLevelObjectiveOutput() ServiceServiceLevelObjectiveOutput {
	return i.ToServiceServiceLevelObjectiveOutputWithContext(context.Background())
}

func (i *ServiceServiceLevelObjective) ToServiceServiceLevelObjectiveOutputWithContext(ctx context.Context) ServiceServiceLevelObjectiveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceServiceLevelObjectiveOutput)
}

type ServiceServiceLevelObjectiveOutput struct {
	*pulumi.OutputState
}

func (ServiceServiceLevelObjectiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceServiceLevelObjective)(nil))
}

func (o ServiceServiceLevelObjectiveOutput) ToServiceServiceLevelObjectiveOutput() ServiceServiceLevelObjectiveOutput {
	return o
}

func (o ServiceServiceLevelObjectiveOutput) ToServiceServiceLevelObjectiveOutputWithContext(ctx context.Context) ServiceServiceLevelObjectiveOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceServiceLevelObjectiveOutput{})
}
