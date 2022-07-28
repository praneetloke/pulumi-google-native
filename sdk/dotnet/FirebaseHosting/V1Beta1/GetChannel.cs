// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1
{
    public static class GetChannel
    {
        /// <summary>
        /// Retrieves information for the specified channel of the specified site.
        /// </summary>
        public static Task<GetChannelResult> InvokeAsync(GetChannelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetChannelResult>("google-native:firebasehosting/v1beta1:getChannel", args ?? new GetChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information for the specified channel of the specified site.
        /// </summary>
        public static Output<GetChannelResult> Invoke(GetChannelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetChannelResult>("google-native:firebasehosting/v1beta1:getChannel", args ?? new GetChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelId", required: true)]
        public string ChannelId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("siteId", required: true)]
        public string SiteId { get; set; } = null!;

        public GetChannelArgs()
        {
        }
        public static new GetChannelArgs Empty => new GetChannelArgs();
    }

    public sealed class GetChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        public GetChannelInvokeArgs()
        {
        }
        public static new GetChannelInvokeArgs Empty => new GetChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetChannelResult
    {
        /// <summary>
        /// The time at which the channel was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Text labels used for extra metadata and/or filtering.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current release for the channel, if any.
        /// </summary>
        public readonly Outputs.ReleaseResponse Release;
        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        public readonly int RetainedReleaseCount;
        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
        /// </summary>
        public readonly string Ttl;
        /// <summary>
        /// The time at which the channel was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetChannelResult(
            string createTime,

            string expireTime,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.ReleaseResponse release,

            int retainedReleaseCount,

            string ttl,

            string updateTime,

            string url)
        {
            CreateTime = createTime;
            ExpireTime = expireTime;
            Labels = labels;
            Name = name;
            Release = release;
            RetainedReleaseCount = retainedReleaseCount;
            Ttl = ttl;
            UpdateTime = updateTime;
            Url = url;
        }
    }
}
