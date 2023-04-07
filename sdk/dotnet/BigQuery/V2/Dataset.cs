// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2
{
    /// <summary>
    /// Creates a new empty dataset.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:bigquery/v2:Dataset")]
    public partial class Dataset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        /// </summary>
        [Output("access")]
        public Output<ImmutableArray<Outputs.DatasetAccessItemResponse>> Access { get; private set; } = null!;

        /// <summary>
        /// The time when this dataset was created, in milliseconds since the epoch.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// [Required] A reference that identifies the dataset.
        /// </summary>
        [Output("datasetReference")]
        public Output<Outputs.DatasetReferenceResponse> DatasetReference { get; private set; } = null!;

        /// <summary>
        /// The default collation of the dataset.
        /// </summary>
        [Output("defaultCollation")]
        public Output<string> DefaultCollation { get; private set; } = null!;

        [Output("defaultEncryptionConfiguration")]
        public Output<Outputs.EncryptionConfigurationResponse> DefaultEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        /// </summary>
        [Output("defaultPartitionExpirationMs")]
        public Output<string> DefaultPartitionExpirationMs { get; private set; } = null!;

        /// <summary>
        /// The default rounding mode of the dataset.
        /// </summary>
        [Output("defaultRoundingMode")]
        public Output<string> DefaultRoundingMode { get; private set; } = null!;

        /// <summary>
        /// [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        /// </summary>
        [Output("defaultTableExpirationMs")]
        public Output<string> DefaultTableExpirationMs { get; private set; } = null!;

        /// <summary>
        /// [Optional] A user-friendly description of the dataset.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A hash of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// [Optional] A descriptive name for the dataset.
        /// </summary>
        [Output("friendlyName")]
        public Output<string> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// [Optional] Indicates if table names are case insensitive in the dataset.
        /// </summary>
        [Output("isCaseInsensitive")]
        public Output<bool> IsCaseInsensitive { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// [Optional] Number of hours for the max time travel for all tables in the dataset.
        /// </summary>
        [Output("maxTimeTravelHours")]
        public Output<string> MaxTimeTravelHours { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Optional] Storage billing model to be used for all tables in the dataset. Can be set to PHYSICAL. Default is LOGICAL.
        /// </summary>
        [Output("storageBillingModel")]
        public Output<string> StorageBillingModel { get; private set; } = null!;

        /// <summary>
        /// [Optional]The tags associated with this dataset. Tag keys are globally unique.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DatasetTagsItemResponse>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:bigquery/v2:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:bigquery/v2:Dataset", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, options);
        }
    }

    public sealed class DatasetArgs : global::Pulumi.ResourceArgs
    {
        [Input("access")]
        private InputList<Inputs.DatasetAccessItemArgs>? _access;

        /// <summary>
        /// [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        /// </summary>
        public InputList<Inputs.DatasetAccessItemArgs> Access
        {
            get => _access ?? (_access = new InputList<Inputs.DatasetAccessItemArgs>());
            set => _access = value;
        }

        /// <summary>
        /// [Required] A reference that identifies the dataset.
        /// </summary>
        [Input("datasetReference")]
        public Input<Inputs.DatasetReferenceArgs>? DatasetReference { get; set; }

        [Input("defaultEncryptionConfiguration")]
        public Input<Inputs.EncryptionConfigurationArgs>? DefaultEncryptionConfiguration { get; set; }

        /// <summary>
        /// [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        /// </summary>
        [Input("defaultPartitionExpirationMs")]
        public Input<string>? DefaultPartitionExpirationMs { get; set; }

        /// <summary>
        /// [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        /// </summary>
        [Input("defaultTableExpirationMs")]
        public Input<string>? DefaultTableExpirationMs { get; set; }

        /// <summary>
        /// [Optional] A user-friendly description of the dataset.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Optional] A descriptive name for the dataset.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// [Optional] Indicates if table names are case insensitive in the dataset.
        /// </summary>
        [Input("isCaseInsensitive")]
        public Input<bool>? IsCaseInsensitive { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// [Optional] Number of hours for the max time travel for all tables in the dataset.
        /// </summary>
        [Input("maxTimeTravelHours")]
        public Input<string>? MaxTimeTravelHours { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// [Optional] Storage billing model to be used for all tables in the dataset. Can be set to PHYSICAL. Default is LOGICAL.
        /// </summary>
        [Input("storageBillingModel")]
        public Input<string>? StorageBillingModel { get; set; }

        [Input("tags")]
        private InputList<Inputs.DatasetTagsItemArgs>? _tags;

        /// <summary>
        /// [Optional]The tags associated with this dataset. Tag keys are globally unique.
        /// </summary>
        public InputList<Inputs.DatasetTagsItemArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DatasetTagsItemArgs>());
            set => _tags = value;
        }

        public DatasetArgs()
        {
        }
        public static new DatasetArgs Empty => new DatasetArgs();
    }
}
