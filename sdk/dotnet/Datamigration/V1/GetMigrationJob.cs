// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1
{
    public static class GetMigrationJob
    {
        /// <summary>
        /// Gets details of a single migration job.
        /// </summary>
        public static Task<GetMigrationJobResult> InvokeAsync(GetMigrationJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMigrationJobResult>("google-native:datamigration/v1:getMigrationJob", args ?? new GetMigrationJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single migration job.
        /// </summary>
        public static Output<GetMigrationJobResult> Invoke(GetMigrationJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMigrationJobResult>("google-native:datamigration/v1:getMigrationJob", args ?? new GetMigrationJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMigrationJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("migrationJobId", required: true)]
        public string MigrationJobId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetMigrationJobArgs()
        {
        }
        public static new GetMigrationJobArgs Empty => new GetMigrationJobArgs();
    }

    public sealed class GetMigrationJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("migrationJobId", required: true)]
        public Input<string> MigrationJobId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetMigrationJobInvokeArgs()
        {
        }
        public static new GetMigrationJobInvokeArgs Empty => new GetMigrationJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetMigrationJobResult
    {
        /// <summary>
        /// The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The resource name (URI) of the destination connection profile.
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The database engine type and provider of the destination.
        /// </summary>
        public readonly Outputs.DatabaseTypeResponse DestinationDatabase;
        /// <summary>
        /// The migration job display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The initial dump flags. This field and the "dump_path" field are mutually exclusive.
        /// </summary>
        public readonly Outputs.DumpFlagsResponse DumpFlags;
        /// <summary>
        /// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]). This field and the "dump_flags" field are mutually exclusive.
        /// </summary>
        public readonly string DumpPath;
        /// <summary>
        /// The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// If the migration job is completed, the time when it was completed.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The error details in case of state FAILED.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current migration job phase.
        /// </summary>
        public readonly string Phase;
        /// <summary>
        /// The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        /// </summary>
        public readonly Outputs.ReverseSshConnectivityResponse ReverseSshConnectivity;
        /// <summary>
        /// The resource name (URI) of the source connection profile.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The database engine type and provider of the source.
        /// </summary>
        public readonly Outputs.DatabaseTypeResponse SourceDatabase;
        /// <summary>
        /// The current migration job state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// static ip connectivity data (default, no additional details needed).
        /// </summary>
        public readonly Outputs.StaticIpConnectivityResponse StaticIpConnectivity;
        /// <summary>
        /// The migration job type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The details of the VPC network that the source database is located in.
        /// </summary>
        public readonly Outputs.VpcPeeringConnectivityResponse VpcPeeringConnectivity;

        [OutputConstructor]
        private GetMigrationJobResult(
            string createTime,

            string destination,

            Outputs.DatabaseTypeResponse destinationDatabase,

            string displayName,

            Outputs.DumpFlagsResponse dumpFlags,

            string dumpPath,

            string duration,

            string endTime,

            Outputs.StatusResponse error,

            ImmutableDictionary<string, string> labels,

            string name,

            string phase,

            Outputs.ReverseSshConnectivityResponse reverseSshConnectivity,

            string source,

            Outputs.DatabaseTypeResponse sourceDatabase,

            string state,

            Outputs.StaticIpConnectivityResponse staticIpConnectivity,

            string type,

            string updateTime,

            Outputs.VpcPeeringConnectivityResponse vpcPeeringConnectivity)
        {
            CreateTime = createTime;
            Destination = destination;
            DestinationDatabase = destinationDatabase;
            DisplayName = displayName;
            DumpFlags = dumpFlags;
            DumpPath = dumpPath;
            Duration = duration;
            EndTime = endTime;
            Error = error;
            Labels = labels;
            Name = name;
            Phase = phase;
            ReverseSshConnectivity = reverseSshConnectivity;
            Source = source;
            SourceDatabase = sourceDatabase;
            State = state;
            StaticIpConnectivity = staticIpConnectivity;
            Type = type;
            UpdateTime = updateTime;
            VpcPeeringConnectivity = vpcPeeringConnectivity;
        }
    }
}
