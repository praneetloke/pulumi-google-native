// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    /// <summary>
    /// Creates a zone resource within a lake.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Aggregated status of the underlying assets of the zone.
        /// </summary>
        [Output("assetStatus")]
        public Output<Outputs.GoogleCloudDataplexV1AssetStatusResponse> AssetStatus { get; private set; } = null!;

        /// <summary>
        /// The time when the zone was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the zone.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Specification of the discovery feature applied to data in this zone.
        /// </summary>
        [Output("discoverySpec")]
        public Output<Outputs.GoogleCloudDataplexV1ZoneDiscoverySpecResponse> DiscoverySpec { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. User defined labels for the zone.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("lakeId")]
        public Output<string> LakeId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of the zone, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Specification of the resources that are referenced by the assets within this zone.
        /// </summary>
        [Output("resourceSpec")]
        public Output<Outputs.GoogleCloudDataplexV1ZoneResourceSpecResponse> ResourceSpec { get; private set; } = null!;

        /// <summary>
        /// Current state of the zone.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Immutable. The type of the zone.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the zone was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Required. Zone identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique across all lakes from all locations in a project. * Must not be one of the reserved IDs (i.e. "default", "global-temp")
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Zone", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "lakeId",
                    "location",
                    "project",
                    "zoneId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the zone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Specification of the discovery feature applied to data in this zone.
        /// </summary>
        [Input("discoverySpec")]
        public Input<Inputs.GoogleCloudDataplexV1ZoneDiscoverySpecArgs>? DiscoverySpec { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User defined labels for the zone.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specification of the resources that are referenced by the assets within this zone.
        /// </summary>
        [Input("resourceSpec", required: true)]
        public Input<Inputs.GoogleCloudDataplexV1ZoneResourceSpecArgs> ResourceSpec { get; set; } = null!;

        /// <summary>
        /// Immutable. The type of the zone.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.ZoneType> Type { get; set; } = null!;

        /// <summary>
        /// Required. Zone identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique across all lakes from all locations in a project. * Must not be one of the reserved IDs (i.e. "default", "global-temp")
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }
}
