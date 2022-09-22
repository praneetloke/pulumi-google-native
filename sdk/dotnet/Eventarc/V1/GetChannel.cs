// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1
{
    public static class GetChannel
    {
        /// <summary>
        /// Get a single Channel.
        /// </summary>
        public static Task<GetChannelResult> InvokeAsync(GetChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChannelResult>("google-native:eventarc/v1:getChannel", args ?? new GetChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Get a single Channel.
        /// </summary>
        public static Output<GetChannelResult> Invoke(GetChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChannelResult>("google-native:eventarc/v1:getChannel", args ?? new GetChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelId", required: true)]
        public string ChannelId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetChannelArgs()
        {
        }
        public static new GetChannelArgs Empty => new GetChannelArgs();
    }

    public sealed class GetChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetChannelInvokeArgs()
        {
        }
        public static new GetChannelInvokeArgs Empty => new GetChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetChannelResult
    {
        /// <summary>
        /// The activation token for the channel. The token must be used by the provider to register the channel for publishing.
        /// </summary>
        public readonly string ActivationToken;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
        /// </summary>
        public readonly string CryptoKeyName;
        /// <summary>
        /// The resource name of the channel. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/channels/{channel_id}` format.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
        /// </summary>
        public readonly string Provider;
        /// <summary>
        /// The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
        /// </summary>
        public readonly string PubsubTopic;
        /// <summary>
        /// The state of a Channel.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The last-modified time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetChannelResult(
            string activationToken,

            string createTime,

            string cryptoKeyName,

            string name,

            string provider,

            string pubsubTopic,

            string state,

            string uid,

            string updateTime)
        {
            ActivationToken = activationToken;
            CreateTime = createTime;
            CryptoKeyName = cryptoKeyName;
            Name = name;
            Provider = provider;
            PubsubTopic = pubsubTopic;
            State = state;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
