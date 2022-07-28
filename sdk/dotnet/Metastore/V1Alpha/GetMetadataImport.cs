// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha
{
    public static class GetMetadataImport
    {
        /// <summary>
        /// Gets details of a single import.
        /// </summary>
        public static Task<GetMetadataImportResult> InvokeAsync(GetMetadataImportArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetadataImportResult>("google-native:metastore/v1alpha:getMetadataImport", args ?? new GetMetadataImportArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single import.
        /// </summary>
        public static Output<GetMetadataImportResult> Invoke(GetMetadataImportInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMetadataImportResult>("google-native:metastore/v1alpha:getMetadataImport", args ?? new GetMetadataImportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetadataImportArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("metadataImportId", required: true)]
        public string MetadataImportId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetMetadataImportArgs()
        {
        }
        public static new GetMetadataImportArgs Empty => new GetMetadataImportArgs();
    }

    public sealed class GetMetadataImportInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("metadataImportId", required: true)]
        public Input<string> MetadataImportId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetMetadataImportInvokeArgs()
        {
        }
        public static new GetMetadataImportInvokeArgs Empty => new GetMetadataImportInvokeArgs();
    }


    [OutputType]
    public sealed class GetMetadataImportResult
    {
        /// <summary>
        /// The time when the metadata import was started.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Immutable. A database dump from a pre-existing metastore's database.
        /// </summary>
        public readonly Outputs.DatabaseDumpResponse DatabaseDump;
        /// <summary>
        /// The description of the metadata import.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time when the metadata import finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the metadata import.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time when the metadata import was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetMetadataImportResult(
            string createTime,

            Outputs.DatabaseDumpResponse databaseDump,

            string description,

            string endTime,

            string name,

            string state,

            string updateTime)
        {
            CreateTime = createTime;
            DatabaseDump = databaseDump;
            Description = description;
            EndTime = endTime;
            Name = name;
            State = state;
            UpdateTime = updateTime;
        }
    }
}
