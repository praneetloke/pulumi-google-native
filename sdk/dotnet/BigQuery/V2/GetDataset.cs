// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2
{
    public static class GetDataset
    {
        /// <summary>
        /// Returns the dataset specified by datasetID.
        /// </summary>
        public static Task<GetDatasetResult> InvokeAsync(GetDatasetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatasetResult>("google-native:bigquery/v2:getDataset", args ?? new GetDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the dataset specified by datasetID.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("google-native:bigquery/v2:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatasetArgs : global::Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDatasetArgs()
        {
        }
        public static new GetDatasetArgs Empty => new GetDatasetArgs();
    }

    public sealed class GetDatasetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatasetInvokeArgs()
        {
        }
        public static new GetDatasetInvokeArgs Empty => new GetDatasetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatasetResult
    {
        /// <summary>
        /// [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        /// </summary>
        public readonly ImmutableArray<Outputs.DatasetAccessItemResponse> Access;
        /// <summary>
        /// The time when this dataset was created, in milliseconds since the epoch.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// [Required] A reference that identifies the dataset.
        /// </summary>
        public readonly Outputs.DatasetReferenceResponse DatasetReference;
        /// <summary>
        /// The default collation of the dataset.
        /// </summary>
        public readonly string DefaultCollation;
        public readonly Outputs.EncryptionConfigurationResponse DefaultEncryptionConfiguration;
        /// <summary>
        /// [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        /// </summary>
        public readonly string DefaultPartitionExpirationMs;
        /// <summary>
        /// The default rounding mode of the dataset.
        /// </summary>
        public readonly string DefaultRoundingMode;
        /// <summary>
        /// [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        /// </summary>
        public readonly string DefaultTableExpirationMs;
        /// <summary>
        /// [Optional] A user-friendly description of the dataset.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A hash of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// [Optional] A descriptive name for the dataset.
        /// </summary>
        public readonly string FriendlyName;
        /// <summary>
        /// [Optional] Indicates if table names are case insensitive in the dataset.
        /// </summary>
        public readonly bool IsCaseInsensitive;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// [Optional] Number of hours for the max time travel for all tables in the dataset.
        /// </summary>
        public readonly string MaxTimeTravelHours;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Optional] Storage billing model to be used for all tables in the dataset. Can be set to PHYSICAL. Default is LOGICAL.
        /// </summary>
        public readonly string StorageBillingModel;
        /// <summary>
        /// [Optional]The tags associated with this dataset. Tag keys are globally unique.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatasetTagsItemResponse> Tags;

        [OutputConstructor]
        private GetDatasetResult(
            ImmutableArray<Outputs.DatasetAccessItemResponse> access,

            string creationTime,

            Outputs.DatasetReferenceResponse datasetReference,

            string defaultCollation,

            Outputs.EncryptionConfigurationResponse defaultEncryptionConfiguration,

            string defaultPartitionExpirationMs,

            string defaultRoundingMode,

            string defaultTableExpirationMs,

            string description,

            string etag,

            string friendlyName,

            bool isCaseInsensitive,

            string kind,

            ImmutableDictionary<string, string> labels,

            string lastModifiedTime,

            string location,

            string maxTimeTravelHours,

            bool satisfiesPzs,

            string selfLink,

            string storageBillingModel,

            ImmutableArray<Outputs.DatasetTagsItemResponse> tags)
        {
            Access = access;
            CreationTime = creationTime;
            DatasetReference = datasetReference;
            DefaultCollation = defaultCollation;
            DefaultEncryptionConfiguration = defaultEncryptionConfiguration;
            DefaultPartitionExpirationMs = defaultPartitionExpirationMs;
            DefaultRoundingMode = defaultRoundingMode;
            DefaultTableExpirationMs = defaultTableExpirationMs;
            Description = description;
            Etag = etag;
            FriendlyName = friendlyName;
            IsCaseInsensitive = isCaseInsensitive;
            Kind = kind;
            Labels = labels;
            LastModifiedTime = lastModifiedTime;
            Location = location;
            MaxTimeTravelHours = maxTimeTravelHours;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            StorageBillingModel = storageBillingModel;
            Tags = tags;
        }
    }
}
