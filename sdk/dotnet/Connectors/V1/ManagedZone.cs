// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1
{
    /// <summary>
    /// Creates a new ManagedZone in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:connectors/v1:ManagedZone")]
    public partial class ManagedZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Created time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// DNS Name of the resource
        /// </summary>
        [Output("dns")]
        public Output<string> Dns { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        /// </summary>
        [Output("managedZoneId")]
        public Output<string> ManagedZoneId { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Managed Zone. Format: projects/{project}/locations/global/managedZones/{managed_zone}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the Target Project
        /// </summary>
        [Output("targetProject")]
        public Output<string> TargetProject { get; private set; } = null!;

        /// <summary>
        /// The name of the Target Project VPC Network
        /// </summary>
        [Output("targetVpc")]
        public Output<string> TargetVpc { get; private set; } = null!;

        /// <summary>
        /// Updated time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedZone(string name, ManagedZoneArgs args, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:ManagedZone", name, args ?? new ManagedZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedZone(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:ManagedZone", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "managedZoneId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedZone Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedZone(name, id, options);
        }
    }

    public sealed class ManagedZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// DNS Name of the resource
        /// </summary>
        [Input("dns", required: true)]
        public Input<string> Dns { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        /// </summary>
        [Input("managedZoneId", required: true)]
        public Input<string> ManagedZoneId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the Target Project
        /// </summary>
        [Input("targetProject", required: true)]
        public Input<string> TargetProject { get; set; } = null!;

        /// <summary>
        /// The name of the Target Project VPC Network
        /// </summary>
        [Input("targetVpc", required: true)]
        public Input<string> TargetVpc { get; set; } = null!;

        public ManagedZoneArgs()
        {
        }
        public static new ManagedZoneArgs Empty => new ManagedZoneArgs();
    }
}