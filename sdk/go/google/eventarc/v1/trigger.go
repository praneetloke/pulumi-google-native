// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new trigger in a particular project and location.
type Trigger struct {
	pulumi.CustomResourceState

	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel pulumi.StringOutput `pulumi:"channel"`
	// The reason(s) why a trigger is in FAILED state.
	Conditions pulumi.StringMapOutput `pulumi:"conditions"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Destination specifies where the events should be sent to.
	Destination DestinationResponseOutput `pulumi:"destination"`
	// This checksum is computed by the server based on the value of other fields, and might be sent only on create requests to ensure that the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType pulumi.StringOutput `pulumi:"eventDataContentType"`
	// Unordered list. The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
	EventFilters EventFilterResponseArrayOutput `pulumi:"eventFilters"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The `iam.serviceAccounts.actAs` permission must be granted on the service account to allow a principal to impersonate the service account. For more information, see the [Roles and permissions](/eventarc/docs/all-roles-permissions) page specific to the trigger destination.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Optional. To deliver messages, Eventarc might use other Google Cloud products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TransportResponseOutput `pulumi:"transport"`
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// Server-assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
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
	if args.EventFilters == nil {
		return nil, errors.New("invalid value for required argument 'EventFilters'")
	}
	if args.TriggerId == nil {
		return nil, errors.New("invalid value for required argument 'TriggerId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"triggerId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Trigger
	err := ctx.RegisterResource("google-native:eventarc/v1:Trigger", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:eventarc/v1:Trigger", name, id, state, &resource, opts...)
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
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel *string `pulumi:"channel"`
	// Destination specifies where the events should be sent to.
	Destination Destination `pulumi:"destination"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType *string `pulumi:"eventDataContentType"`
	// Unordered list. The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
	EventFilters []EventFilter `pulumi:"eventFilters"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The `iam.serviceAccounts.actAs` permission must be granted on the service account to allow a principal to impersonate the service account. For more information, see the [Roles and permissions](/eventarc/docs/all-roles-permissions) page specific to the trigger destination.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Optional. To deliver messages, Eventarc might use other Google Cloud products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport *Transport `pulumi:"transport"`
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId string `pulumi:"triggerId"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel pulumi.StringPtrInput
	// Destination specifies where the events should be sent to.
	Destination DestinationInput
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType pulumi.StringPtrInput
	// Unordered list. The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
	EventFilters EventFilterArrayInput
	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The `iam.serviceAccounts.actAs` permission must be granted on the service account to allow a principal to impersonate the service account. For more information, see the [Roles and permissions](/eventarc/docs/all-roles-permissions) page specific to the trigger destination.
	ServiceAccount pulumi.StringPtrInput
	// Optional. To deliver messages, Eventarc might use other Google Cloud products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TransportPtrInput
	// Required. The user-provided ID to be assigned to the trigger.
	TriggerId pulumi.StringInput
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

// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
func (o TriggerOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Channel }).(pulumi.StringOutput)
}

// The reason(s) why a trigger is in FAILED state.
func (o TriggerOutput) Conditions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Conditions }).(pulumi.StringMapOutput)
}

// The creation time.
func (o TriggerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Destination specifies where the events should be sent to.
func (o TriggerOutput) Destination() DestinationResponseOutput {
	return o.ApplyT(func(v *Trigger) DestinationResponseOutput { return v.Destination }).(DestinationResponseOutput)
}

// This checksum is computed by the server based on the value of other fields, and might be sent only on create requests to ensure that the client has an up-to-date value before proceeding.
func (o TriggerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
func (o TriggerOutput) EventDataContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.EventDataContentType }).(pulumi.StringOutput)
}

// Unordered list. The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
func (o TriggerOutput) EventFilters() EventFilterResponseArrayOutput {
	return o.ApplyT(func(v *Trigger) EventFilterResponseArrayOutput { return v.EventFilters }).(EventFilterResponseArrayOutput)
}

// Optional. User labels attached to the triggers that can be used to group resources.
func (o TriggerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o TriggerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TriggerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The `iam.serviceAccounts.actAs` permission must be granted on the service account to allow a principal to impersonate the service account. For more information, see the [Roles and permissions](/eventarc/docs/all-roles-permissions) page specific to the trigger destination.
func (o TriggerOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Optional. To deliver messages, Eventarc might use other Google Cloud products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
func (o TriggerOutput) Transport() TransportResponseOutput {
	return o.ApplyT(func(v *Trigger) TransportResponseOutput { return v.Transport }).(TransportResponseOutput)
}

// Required. The user-provided ID to be assigned to the trigger.
func (o TriggerOutput) TriggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.TriggerId }).(pulumi.StringOutput)
}

// Server-assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o TriggerOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The last-modified time.
func (o TriggerOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterOutputType(TriggerOutput{})
}
