// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1
{
    /// <summary>
    /// Creates a new private connection in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:datamigration/v1:PrivateConnection")]
    public partial class PrivateConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The create time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The private connection display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The error details in case of state FAILED.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        /// <summary>
        /// The resource labels for private connections to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. The private connection identifier.
        /// </summary>
        [Output("privateConnectionId")]
        public Output<string> PrivateConnectionId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Optional. If set to true, will skip validations.
        /// </summary>
        [Output("skipValidation")]
        public Output<bool?> SkipValidation { get; private set; } = null!;

        /// <summary>
        /// The state of the private connection.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The last update time of the resource.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// VPC peering configuration.
        /// </summary>
        [Output("vpcPeeringConfig")]
        public Output<Outputs.VpcPeeringConfigResponse> VpcPeeringConfig { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateConnection(string name, PrivateConnectionArgs args, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:PrivateConnection", name, args ?? new PrivateConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:PrivateConnection", name, null, MakeResourceOptions(options, id))
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
                    "privateConnectionId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateConnection(name, id, options);
        }
    }

    public sealed class PrivateConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The private connection display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for private connections to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required. The private connection identifier.
        /// </summary>
        [Input("privateConnectionId", required: true)]
        public Input<string> PrivateConnectionId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Optional. If set to true, will skip validations.
        /// </summary>
        [Input("skipValidation")]
        public Input<bool>? SkipValidation { get; set; }

        /// <summary>
        /// VPC peering configuration.
        /// </summary>
        [Input("vpcPeeringConfig")]
        public Input<Inputs.VpcPeeringConfigArgs>? VpcPeeringConfig { get; set; }

        public PrivateConnectionArgs()
        {
        }
        public static new PrivateConnectionArgs Empty => new PrivateConnectionArgs();
    }
}
