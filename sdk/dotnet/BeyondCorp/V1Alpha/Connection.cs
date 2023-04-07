// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha
{
    /// <summary>
    /// Creates a new Connection in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:beyondcorp/v1alpha:Connection")]
    public partial class Connection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp Connection.
        /// </summary>
        [Output("applicationEndpoint")]
        public Output<Outputs.ApplicationEndpointResponse> ApplicationEndpoint { get; private set; } = null!;

        /// <summary>
        /// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Output("connectionId")]
        public Output<string?> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
        /// </summary>
        [Output("connectors")]
        public Output<ImmutableArray<string>> Connectors { get; private set; } = null!;

        /// <summary>
        /// Timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Gateway used by the connection.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.GatewayResponse> Gateway { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Unique resource name of the connection. The name is ignored when creating a connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The current state of the connection.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of network connectivity used by the connection.
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
        public Output<bool?> ValidateOnly { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1alpha:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1alpha:Connection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, options);
        }
    }

    public sealed class ConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp Connection.
        /// </summary>
        [Input("applicationEndpoint", required: true)]
        public Input<Inputs.ApplicationEndpointArgs> ApplicationEndpoint { get; set; } = null!;

        /// <summary>
        /// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        [Input("connectors")]
        private InputList<string>? _connectors;

        /// <summary>
        /// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
        /// </summary>
        public InputList<string> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<string>());
            set => _connectors = value;
        }

        /// <summary>
        /// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Gateway used by the connection.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.GatewayArgs>? Gateway { get; set; }

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
        /// Unique resource name of the connection. The name is ignored when creating a connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The type of network connectivity used by the connection.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.BeyondCorp.V1Alpha.ConnectionType> Type { get; set; } = null!;

        /// <summary>
        /// Optional. If set, validates request by executing a dry-run which would not alter the resource in any way.
        /// </summary>
        [Input("validateOnly")]
        public Input<bool>? ValidateOnly { get; set; }

        public ConnectionArgs()
        {
        }
        public static new ConnectionArgs Empty => new ConnectionArgs();
    }
}
