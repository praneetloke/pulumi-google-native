// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2
{
    public static class GetParticipant
    {
        /// <summary>
        /// Retrieves a conversation participant.
        /// </summary>
        public static Task<GetParticipantResult> InvokeAsync(GetParticipantArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetParticipantResult>("google-native:dialogflow/v2:getParticipant", args ?? new GetParticipantArgs(), options.WithVersion());
    }


    public sealed class GetParticipantArgs : Pulumi.InvokeArgs
    {
        [Input("conversationId", required: true)]
        public string ConversationId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("participantId", required: true)]
        public string ParticipantId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetParticipantArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetParticipantResult
    {
        /// <summary>
        /// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
        /// </summary>
        public readonly string SipRecordingMediaLabel;

        [OutputConstructor]
        private GetParticipantResult(
            string name,

            string role,

            string sipRecordingMediaLabel)
        {
            Name = name;
            Role = role;
            SipRecordingMediaLabel = sipRecordingMediaLabel;
        }
    }
}