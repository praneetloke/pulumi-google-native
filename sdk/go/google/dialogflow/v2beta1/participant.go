// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new participant in a conversation.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Participant struct {
	pulumi.CustomResourceState

	ConversationId pulumi.StringOutput `pulumi:"conversationId"`
	// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
	DocumentsMetadataFilters pulumi.StringMapOutput `pulumi:"documentsMetadataFilters"`
	Location                 pulumi.StringOutput    `pulumi:"location"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
	ObfuscatedExternalUserId pulumi.StringOutput `pulumi:"obfuscatedExternalUserId"`
	Project                  pulumi.StringOutput `pulumi:"project"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role pulumi.StringOutput `pulumi:"role"`
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
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"conversationId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Participant
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:Participant", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:Participant", name, id, state, &resource, opts...)
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
	ConversationId string `pulumi:"conversationId"`
	// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
	DocumentsMetadataFilters map[string]string `pulumi:"documentsMetadataFilters"`
	Location                 *string           `pulumi:"location"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name *string `pulumi:"name"`
	// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
	ObfuscatedExternalUserId *string `pulumi:"obfuscatedExternalUserId"`
	Project                  *string `pulumi:"project"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role *ParticipantRole `pulumi:"role"`
}

// The set of arguments for constructing a Participant resource.
type ParticipantArgs struct {
	ConversationId pulumi.StringInput
	// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
	DocumentsMetadataFilters pulumi.StringMapInput
	Location                 pulumi.StringPtrInput
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name pulumi.StringPtrInput
	// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
	ObfuscatedExternalUserId pulumi.StringPtrInput
	Project                  pulumi.StringPtrInput
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role ParticipantRolePtrInput
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
	return reflect.TypeOf((**Participant)(nil)).Elem()
}

func (i *Participant) ToParticipantOutput() ParticipantOutput {
	return i.ToParticipantOutputWithContext(context.Background())
}

func (i *Participant) ToParticipantOutputWithContext(ctx context.Context) ParticipantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParticipantOutput)
}

type ParticipantOutput struct{ *pulumi.OutputState }

func (ParticipantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Participant)(nil)).Elem()
}

func (o ParticipantOutput) ToParticipantOutput() ParticipantOutput {
	return o
}

func (o ParticipantOutput) ToParticipantOutputWithContext(ctx context.Context) ParticipantOutput {
	return o
}

func (o ParticipantOutput) ConversationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.ConversationId }).(pulumi.StringOutput)
}

// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
func (o ParticipantOutput) DocumentsMetadataFilters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringMapOutput { return v.DocumentsMetadataFilters }).(pulumi.StringMapOutput)
}

func (o ParticipantOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
func (o ParticipantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
func (o ParticipantOutput) ObfuscatedExternalUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.ObfuscatedExternalUserId }).(pulumi.StringOutput)
}

func (o ParticipantOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
func (o ParticipantOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Participant) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParticipantInput)(nil)).Elem(), &Participant{})
	pulumi.RegisterOutputType(ParticipantOutput{})
}
