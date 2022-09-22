// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1
{
    public static class GetDataset
    {
        /// <summary>
        /// Gets dataset by resource name.
        /// </summary>
        public static Task<GetDatasetResult> InvokeAsync(GetDatasetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatasetResult>("google-native:datalabeling/v1beta1:getDataset", args ?? new GetDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets dataset by resource name.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("google-native:datalabeling/v1beta1:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());
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
        /// The names of any related resources that are blocking changes to the dataset.
        /// </summary>
        public readonly ImmutableArray<string> BlockingResources;
        /// <summary>
        /// Time the dataset is created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The number of data items in the dataset.
        /// </summary>
        public readonly string DataItemCount;
        /// <summary>
        /// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the dataset. Maximum of 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// This is populated with the original input configs where ImportData is called. It is available only after the clients import data to this dataset.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1InputConfigResponse> InputConfigs;
        /// <summary>
        /// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
        /// </summary>
        public readonly string LastMigrateTime;
        /// <summary>
        /// Dataset resource name, format is: projects/{project_id}/datasets/{dataset_id}
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDatasetResult(
            ImmutableArray<string> blockingResources,

            string createTime,

            string dataItemCount,

            string description,

            string displayName,

            ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1InputConfigResponse> inputConfigs,

            string lastMigrateTime,

            string name)
        {
            BlockingResources = blockingResources;
            CreateTime = createTime;
            DataItemCount = dataItemCount;
            Description = description;
            DisplayName = displayName;
            InputConfigs = inputConfigs;
            LastMigrateTime = lastMigrateTime;
            Name = name;
        }
    }
}
