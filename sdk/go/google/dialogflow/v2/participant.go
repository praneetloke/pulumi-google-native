// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new participant in a conversation.
type Participant struct {
	pulumi.CustomResourceState

	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role pulumi.StringOutput `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel pulumi.StringOutput `pulumi:"sipRecordingMediaLabel"`
}

// NewParticipant registers a new resource with the given unique name, arguments, and options.
func NewParticipant(ctx *pulumi.Context,
	name string, args *ParticipantArgs, opts ...pulumi.ResourceOption) (*Participant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConversationId == nil {
		return nil, errors.New("invalid value for required argument 'ConversationId'")
	}
	var resource Participant
	err := ctx.RegisterResource("google-native:dialogflow/v2:Participant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParticipant gets an existing Participant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParticipant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParticipantState, opts ...pulumi.ResourceOption) (*Participant, error) {
	var resource Participant
	err := ctx.ReadResource("google-native:dialogflow/v2:Participant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Participant resources.
type participantState struct {
}

type ParticipantState struct {
}

func (ParticipantState) ElementType() reflect.Type {
	return reflect.TypeOf((*participantState)(nil)).Elem()
}

type participantArgs struct {
	ConversationId string  `pulumi:"conversationId"`
	Location       *string `pulumi:"location"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role *ParticipantRole `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel *string `pulumi:"sipRecordingMediaLabel"`
}

// The set of arguments for constructing a Participant resource.
type ParticipantArgs struct {
	ConversationId pulumi.StringInput
	Location       pulumi.StringPtrInput
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role ParticipantRolePtrInput
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel pulumi.StringPtrInput
}

func (ParticipantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*participantArgs)(nil)).Elem()
}

type ParticipantInput interface {
	pulumi.Input

	ToParticipantOutput() ParticipantOutput
	ToParticipantOutputWithContext(ctx context.Context) ParticipantOutput
}

func (*Participant) ElementType() reflect.Type {
	return reflect.TypeOf((*Participant)(nil))
}

func (i *Participant) ToParticipantOutput() ParticipantOutput {
	return i.ToParticipantOutputWithContext(context.Background())
}

func (i *Participant) ToParticipantOutputWithContext(ctx context.Context) ParticipantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantOutput)
}

type ParticipantOutput struct {
	*pulumi.OutputState
}

func (ParticipantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Participant)(nil))
}

func (o ParticipantOutput) ToParticipantOutput() ParticipantOutput {
	return o
}

func (o ParticipantOutput) ToParticipantOutputWithContext(ctx context.Context) ParticipantOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ParticipantOutput{})
}
