// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    /// <summary>
    /// A segment of a full transcript.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1ConversationTranscriptTranscriptSegmentResponse
    {
        /// <summary>
        /// For conversations derived from multi-channel audio, this is the channel number corresponding to the audio from that channel. For audioChannelCount = N, its output values can range from '1' to 'N'. A channel tag of 0 indicates that the audio is mono.
        /// </summary>
        public readonly int ChannelTag;
        /// <summary>
        /// A confidence estimate between 0.0 and 1.0 of the fidelity of this segment. A default value of 0.0 indicates that the value is unset.
        /// </summary>
        public readonly double Confidence;
        /// <summary>
        /// The language code of this segment as a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// The participant of this segment.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ConversationParticipantResponse SegmentParticipant;
        /// <summary>
        /// The text of this segment.
        /// </summary>
        public readonly string Text;
        /// <summary>
        /// A list of the word-specific information for each word in the segment.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudContactcenterinsightsV1ConversationTranscriptTranscriptSegmentWordInfoResponse> Words;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1ConversationTranscriptTranscriptSegmentResponse(
            int channelTag,

            double confidence,

            string languageCode,

            Outputs.GoogleCloudContactcenterinsightsV1ConversationParticipantResponse segmentParticipant,

            string text,

            ImmutableArray<Outputs.GoogleCloudContactcenterinsightsV1ConversationTranscriptTranscriptSegmentWordInfoResponse> words)
        {
            ChannelTag = channelTag;
            Confidence = confidence;
            LanguageCode = languageCode;
            SegmentParticipant = segmentParticipant;
            Text = text;
            Words = words;
        }
    }
}