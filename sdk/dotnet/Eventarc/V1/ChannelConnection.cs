// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1
{
    /// <summary>
    /// Create a new ChannelConnection in a particular project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:eventarc/v1:ChannelConnection")]
    public partial class ChannelConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
        /// </summary>
        [Output("activationToken")]
        public Output<string> ActivationToken { get; private set; } = null!;

        /// <summary>
        /// The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
        /// </summary>
        [Output("channel")]
        public Output<string> Channel { get; private set; } = null!;

        /// <summary>
        /// Required. The user-provided ID to be assigned to the channel connection.
        /// </summary>
        [Output("channelConnectionId")]
        public Output<string> ChannelConnectionId { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Server assigned ID of the resource. The server guarantees uniqueness and immutability until deleted.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ChannelConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChannelConnection(string name, ChannelConnectionArgs args, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1:ChannelConnection", name, args ?? new ChannelConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChannelConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1:ChannelConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "channelConnectionId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ChannelConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChannelConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ChannelConnection(name, id, options);
        }
    }

    public sealed class ChannelConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
        /// </summary>
        [Input("activationToken")]
        public Input<string>? ActivationToken { get; set; }

        /// <summary>
        /// The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
        /// </summary>
        [Input("channel", required: true)]
        public Input<string> Channel { get; set; } = null!;

        /// <summary>
        /// Required. The user-provided ID to be assigned to the channel connection.
        /// </summary>
        [Input("channelConnectionId", required: true)]
        public Input<string> ChannelConnectionId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public ChannelConnectionArgs()
        {
        }
        public static new ChannelConnectionArgs Empty => new ChannelConnectionArgs();
    }
}
