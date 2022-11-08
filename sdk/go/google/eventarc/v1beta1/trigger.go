// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new trigger in a particular project and location.
type Trigger struct {
	pulumi.CustomResourceState

	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Destination specifies where the events should be sent to.
	Destination DestinationResponseOutput `pulumi:"destination"`
	// This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Unordered list. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
	MatchingCriteria MatchingCriteriaResponseArrayOutput `pulumi:"matchingCriteria"`
	// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TransportResponseOutput `pulumi:"transport"`
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. If set, validate the request and preview the review, but do not actually post it.
	ValidateOnly pulumi.BoolOutput `pulumi:"validateOnly"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.MatchingCriteria == nil {
		return nil, errors.New("invalid value for required argument 'MatchingCriteria'")
	}
	if args.TriggerId == nil {
		return nil, errors.New("invalid value for required argument 'TriggerId'")
	}
	if args.ValidateOnly == nil {
		return nil, errors.New("invalid value for required argument 'ValidateOnly'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"triggerId",
		"validateOnly",
	})
	opts = append(opts, replaceOnChanges)
	var resource Trigger
	err := ctx.RegisterResource("google-native:eventarc/v1beta1:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("google-native:eventarc/v1beta1:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
}

type TriggerState struct {
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Destination specifies where the events should be sent to.
	Destination Destination `pulumi:"destination"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Unordered list. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
	MatchingCriteria []MatchingCriteria `pulumi:"matchingCriteria"`
	// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId string `pulumi:"triggerId"`
	// Required. If set, validate the request and preview the review, but do not actually post it.
	ValidateOnly bool `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Destination specifies where the events should be sent to.
	Destination DestinationInput
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Unordered list. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
	MatchingCriteria MatchingCriteriaArrayInput
	// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
	ServiceAccount pulumi.StringPtrInput
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId pulumi.StringInput
	// Required. If set, validate the request and preview the review, but do not actually post it.
	ValidateOnly pulumi.BoolInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// The creation time.
func (o TriggerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Destination specifies where the events should be sent to.
func (o TriggerOutput) Destination() DestinationResponseOutput {
	return o.ApplyT(func(v *Trigger) DestinationResponseOutput { return v.Destination }).(DestinationResponseOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
func (o TriggerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. User labels attached to the triggers that can be used to group resources.
func (o TriggerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o TriggerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unordered list. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
func (o TriggerOutput) MatchingCriteria() MatchingCriteriaResponseArrayOutput {
	return o.ApplyT(func(v *Trigger) MatchingCriteriaResponseArrayOutput { return v.MatchingCriteria }).(MatchingCriteriaResponseArrayOutput)
}

// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TriggerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
func (o TriggerOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
func (o TriggerOutput) Transport() TransportResponseOutput {
	return o.ApplyT(func(v *Trigger) TransportResponseOutput { return v.Transport }).(TransportResponseOutput)
}

// Required. The user-provided ID to be assigned to the trigger.
func (o TriggerOutput) TriggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.TriggerId }).(pulumi.StringOutput)
}

// The last-modified time.
func (o TriggerOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Required. If set, validate the request and preview the review, but do not actually post it.
func (o TriggerOutput) ValidateOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *Trigger) pulumi.BoolOutput { return v.ValidateOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterOutputType(TriggerOutput{})
}
