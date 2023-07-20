// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a conversation participant.
func LookupParticipant(ctx *pulumi.Context, args *LookupParticipantArgs, opts ...pulumi.InvokeOption) (*LookupParticipantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupParticipantResult
	err := ctx.Invoke("google-native:dialogflow/v2beta1:getParticipant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParticipantArgs struct {
	ConversationId string  `pulumi:"conversationId"`
	Location       string  `pulumi:"location"`
	ParticipantId  string  `pulumi:"participantId"`
	Project        *string `pulumi:"project"`
}

type LookupParticipantResult struct {
	// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
	DocumentsMetadataFilters map[string]string `pulumi:"documentsMetadataFilters"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name string `pulumi:"name"`
	// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
	ObfuscatedExternalUserId string `pulumi:"obfuscatedExternalUserId"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role string `pulumi:"role"`
}

func LookupParticipantOutput(ctx *pulumi.Context, args LookupParticipantOutputArgs, opts ...pulumi.InvokeOption) LookupParticipantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupParticipantResult, error) {
			args := v.(LookupParticipantArgs)
			r, err := LookupParticipant(ctx, &args, opts...)
			var s LookupParticipantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupParticipantResultOutput)
}

type LookupParticipantOutputArgs struct {
	ConversationId pulumi.StringInput    `pulumi:"conversationId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	ParticipantId  pulumi.StringInput    `pulumi:"participantId"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupParticipantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParticipantArgs)(nil)).Elem()
}

type LookupParticipantResultOutput struct{ *pulumi.OutputState }

func (LookupParticipantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParticipantResult)(nil)).Elem()
}

func (o LookupParticipantResultOutput) ToLookupParticipantResultOutput() LookupParticipantResultOutput {
	return o
}

func (o LookupParticipantResultOutput) ToLookupParticipantResultOutputWithContext(ctx context.Context) LookupParticipantResultOutput {
	return o
}

// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
func (o LookupParticipantResultOutput) DocumentsMetadataFilters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupParticipantResult) map[string]string { return v.DocumentsMetadataFilters }).(pulumi.StringMapOutput)
}

// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
func (o LookupParticipantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
func (o LookupParticipantResultOutput) ObfuscatedExternalUserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.ObfuscatedExternalUserId }).(pulumi.StringOutput)
}

// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
func (o LookupParticipantResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParticipantResultOutput{})
}
