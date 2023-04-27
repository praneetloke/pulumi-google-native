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
    /// Create a metadata entity.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Entity")]
    public partial class Entity : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Identifies the access mechanism to the entity. Not user settable.
        /// </summary>
        [Output("access")]
        public Output<Outputs.GoogleCloudDataplexV1StorageAccessResponse> Access { get; private set; } = null!;

        /// <summary>
        /// Immutable. The ID of the asset associated with the storage location containing the entity data. The entity must be with in the same zone with the asset.
        /// </summary>
        [Output("asset")]
        public Output<string> Asset { get; private set; } = null!;

        /// <summary>
        /// The name of the associated Data Catalog entry.
        /// </summary>
        [Output("catalogEntry")]
        public Output<string> CatalogEntry { get; private set; } = null!;

        /// <summary>
        /// Metadata stores that the entity is compatible with.
        /// </summary>
        [Output("compatibility")]
        public Output<Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusResponse> Compatibility { get; private set; } = null!;

        /// <summary>
        /// The time when the entity was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The storage path of the entity data. For Cloud Storage data, this is the fully-qualified path to the entity, such as gs://bucket/path/to/data. For BigQuery data, this is the name of the table resource, such as projects/project_id/datasets/dataset_id/tables/table_id.
        /// </summary>
        [Output("dataPath")]
        public Output<string> DataPath { get; private set; } = null!;

        /// <summary>
        /// Optional. The set of items within the data path constituting the data in the entity, represented as a glob path. Example: gs://bucket/path/to/data/**/*.csv.
        /// </summary>
        [Output("dataPathPattern")]
        public Output<string> DataPathPattern { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly longer description text. Must be shorter than or equal to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Display name must be shorter than or equal to 256 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. The etag associated with the entity, which can be retrieved with a GetEntity request. Required for update and delete requests.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Identifies the storage format of the entity data. It does not apply to entities with data stored in BigQuery.
        /// </summary>
        [Output("format")]
        public Output<Outputs.GoogleCloudDataplexV1StorageFormatResponse> Format { get; private set; } = null!;

        [Output("lakeId")]
        public Output<string> LakeId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the entity, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The description of the data structure and layout. The schema is not included in list responses. It is only included in SCHEMA and FULL entity views of a GetEntity response.
        /// </summary>
        [Output("schema")]
        public Output<Outputs.GoogleCloudDataplexV1SchemaResponse> Schema { get; private set; } = null!;

        /// <summary>
        /// Immutable. Identifies the storage system of the entity data.
        /// </summary>
        [Output("system")]
        public Output<string> System { get; private set; } = null!;

        /// <summary>
        /// Immutable. The type of entity.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// System generated unique ID for the Entity. This ID will be different if the Entity is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the entity was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Entity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Entity(string name, EntityArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Entity", name, args ?? new EntityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Entity(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Entity", name, null, MakeResourceOptions(options, id))
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
                    "zone",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Entity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Entity Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Entity(name, id, options);
        }
    }

    public sealed class EntityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The ID of the asset associated with the storage location containing the entity data. The entity must be with in the same zone with the asset.
        /// </summary>
        [Input("asset", required: true)]
        public Input<string> Asset { get; set; } = null!;

        /// <summary>
        /// Immutable. The storage path of the entity data. For Cloud Storage data, this is the fully-qualified path to the entity, such as gs://bucket/path/to/data. For BigQuery data, this is the name of the table resource, such as projects/project_id/datasets/dataset_id/tables/table_id.
        /// </summary>
        [Input("dataPath", required: true)]
        public Input<string> DataPath { get; set; } = null!;

        /// <summary>
        /// Optional. The set of items within the data path constituting the data in the entity, represented as a glob path. Example: gs://bucket/path/to/data/**/*.csv.
        /// </summary>
        [Input("dataPathPattern")]
        public Input<string>? DataPathPattern { get; set; }

        /// <summary>
        /// Optional. User friendly longer description text. Must be shorter than or equal to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Display name must be shorter than or equal to 256 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. The etag associated with the entity, which can be retrieved with a GetEntity request. Required for update and delete requests.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Identifies the storage format of the entity data. It does not apply to entities with data stored in BigQuery.
        /// </summary>
        [Input("format", required: true)]
        public Input<Inputs.GoogleCloudDataplexV1StorageFormatArgs> Format { get; set; } = null!;

        /// <summary>
        /// A user-provided entity ID. It is mutable, and will be used as the published table name. Specifying a new ID in an update entity request will override the existing value. The ID must contain only letters (a-z, A-Z), numbers (0-9), and underscores, and consist of 256 or fewer characters.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The description of the data structure and layout. The schema is not included in list responses. It is only included in SCHEMA and FULL entity views of a GetEntity response.
        /// </summary>
        [Input("schema", required: true)]
        public Input<Inputs.GoogleCloudDataplexV1SchemaArgs> Schema { get; set; } = null!;

        /// <summary>
        /// Immutable. Identifies the storage system of the entity data.
        /// </summary>
        [Input("system", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.EntitySystem> System { get; set; } = null!;

        /// <summary>
        /// Immutable. The type of entity.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.EntityType> Type { get; set; } = null!;

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public EntityArgs()
        {
        }
        public static new EntityArgs Empty => new EntityArgs();
    }
}
