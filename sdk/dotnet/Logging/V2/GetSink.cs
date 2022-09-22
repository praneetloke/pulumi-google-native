// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetSink
    {
        /// <summary>
        /// Gets a sink.
        /// </summary>
        public static Task<GetSinkResult> InvokeAsync(GetSinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSinkResult>("google-native:logging/v2:getSink", args ?? new GetSinkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a sink.
        /// </summary>
        public static Output<GetSinkResult> Invoke(GetSinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSinkResult>("google-native:logging/v2:getSink", args ?? new GetSinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSinkArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("sinkId", required: true)]
        public string SinkId { get; set; } = null!;

        public GetSinkArgs()
        {
        }
        public static new GetSinkArgs Empty => new GetSinkArgs();
    }

    public sealed class GetSinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sinkId", required: true)]
        public Input<string> SinkId { get; set; } = null!;

        public GetSinkInvokeArgs()
        {
        }
        public static new GetSinkInvokeArgs Empty => new GetSinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetSinkResult
    {
        /// <summary>
        /// Optional. Options that affect sinks exporting data to BigQuery.
        /// </summary>
        public readonly Outputs.BigQueryOptionsResponse BigqueryOptions;
        /// <summary>
        /// The creation timestamp of the sink.This field may not be present for older sinks.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A description of this sink.The maximum length of the description is 8000 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// Optional. If set to true, then this sink is disabled and it does not export any log entries.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
        /// </summary>
        public readonly ImmutableArray<Outputs.LogExclusionResponse> Exclusions;
        /// <summary>
        /// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity&gt;=ERROR
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
        /// </summary>
        public readonly bool IncludeChildren;
        /// <summary>
        /// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Deprecated. This field is unused.
        /// </summary>
        public readonly string OutputVersionFormat;
        /// <summary>
        /// The last update timestamp of the sink.This field may not be present for older sinks.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is either set by specifying custom_writer_identity or set automatically by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink cannot have a writer_identity and no additional permissions are required.
        /// </summary>
        public readonly string WriterIdentity;

        [OutputConstructor]
        private GetSinkResult(
            Outputs.BigQueryOptionsResponse bigqueryOptions,

            string createTime,

            string description,

            string destination,

            bool disabled,

            ImmutableArray<Outputs.LogExclusionResponse> exclusions,

            string filter,

            bool includeChildren,

            string name,

            string outputVersionFormat,

            string updateTime,

            string writerIdentity)
        {
            BigqueryOptions = bigqueryOptions;
            CreateTime = createTime;
            Description = description;
            Destination = destination;
            Disabled = disabled;
            Exclusions = exclusions;
            Filter = filter;
            IncludeChildren = includeChildren;
            Name = name;
            OutputVersionFormat = outputVersionFormat;
            UpdateTime = updateTime;
            WriterIdentity = writerIdentity;
        }
    }
}
