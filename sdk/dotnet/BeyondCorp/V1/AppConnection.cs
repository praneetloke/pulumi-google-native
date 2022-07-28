// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1
{
    /// <summary>
    /// Creates a new AppConnection in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:beyondcorp/v1:AppConnection")]
    public partial class AppConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Output("appConnectionId")]
        public Output<string?> AppConnectionId { get; private set; } = null!;

        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp AppConnection.
        /// </summary>
        [Output("applicationEndpoint")]
        public Output<Outputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointResponse> ApplicationEndpoint { get; private set; } = null!;

        /// <summary>
        /// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
        /// </summary>
        [Output("connectors")]
        public Output<ImmutableArray<string>> Connectors { get; private set; } = null!;

        /// <summary>
        /// Timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Gateway used by the AppConnection.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionGatewayResponse> Gateway { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The current state of the AppConnection.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of network connectivity used by the AppConnection.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the instance generated by the system.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Timestamp when the resource was last modified.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
        /// </summary>
        [Output("validateOnly")]
        public Output<string?> ValidateOnly { get; private set; } = null!;


        /// <summary>
        /// Create a AppConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppConnection(string name, AppConnectionArgs args, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1:AppConnection", name, args ?? new AppConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1:AppConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
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
        /// Get an existing AppConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppConnection(name, id, options);
        }
    }

    public sealed class AppConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Input("appConnectionId")]
        public Input<string>? AppConnectionId { get; set; }

        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp AppConnection.
        /// </summary>
        [Input("applicationEndpoint", required: true)]
        public Input<Inputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointArgs> ApplicationEndpoint { get; set; } = null!;

        [Input("connectors")]
        private InputList<string>? _connectors;

        /// <summary>
        /// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
        /// </summary>
        public InputList<string> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<string>());
            set => _connectors = value;
        }

        /// <summary>
        /// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Gateway used by the AppConnection.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionGatewayArgs>? Gateway { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The type of network connectivity used by the AppConnection.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.BeyondCorp.V1.AppConnectionType> Type { get; set; } = null!;

        /// <summary>
        /// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public AppConnectionArgs()
        {
        }
        public static new AppConnectionArgs Empty => new AppConnectionArgs();
    }
}
