// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RapidMigrationAssessment.V1
{
    public static class GetCollector
    {
        /// <summary>
        /// Gets details of a single Collector.
        /// </summary>
        public static Task<GetCollectorResult> InvokeAsync(GetCollectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCollectorResult>("google-native:rapidmigrationassessment/v1:getCollector", args ?? new GetCollectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Collector.
        /// </summary>
        public static Output<GetCollectorResult> Invoke(GetCollectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCollectorResult>("google-native:rapidmigrationassessment/v1:getCollector", args ?? new GetCollectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCollectorArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectorId", required: true)]
        public string CollectorId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCollectorArgs()
        {
        }
        public static new GetCollectorArgs Empty => new GetCollectorArgs();
    }

    public sealed class GetCollectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectorId", required: true)]
        public Input<string> CollectorId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCollectorInvokeArgs()
        {
        }
        public static new GetCollectorInvokeArgs Empty => new GetCollectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetCollectorResult
    {
        /// <summary>
        /// Store cloud storage bucket name (which is a guid) created with this Collector.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Client version.
        /// </summary>
        public readonly string ClientVersion;
        /// <summary>
        /// How many days to collect data.
        /// </summary>
        public readonly int CollectionDays;
        /// <summary>
        /// Create time stamp.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// User specified description of the Collector.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User specified name of the Collector.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Uri for EULA (End User License Agreement) from customer.
        /// </summary>
        public readonly string EulaUri;
        /// <summary>
        /// User specified expected asset count.
        /// </summary>
        public readonly string ExpectedAssetCount;
        /// <summary>
        /// Reference to MC Source Guest Os Scan.
        /// </summary>
        public readonly Outputs.GuestOsScanResponse GuestOsScan;
        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// name of resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Service Account email used to ingest data to this Collector.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// State of the Collector.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Update time stamp.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Reference to MC Source vsphere_scan.
        /// </summary>
        public readonly Outputs.VSphereScanResponse VsphereScan;

        [OutputConstructor]
        private GetCollectorResult(
            string bucket,

            string clientVersion,

            int collectionDays,

            string createTime,

            string description,

            string displayName,

            string eulaUri,

            string expectedAssetCount,

            Outputs.GuestOsScanResponse guestOsScan,

            ImmutableDictionary<string, string> labels,

            string name,

            string serviceAccount,

            string state,

            string updateTime,

            Outputs.VSphereScanResponse vsphereScan)
        {
            Bucket = bucket;
            ClientVersion = clientVersion;
            CollectionDays = collectionDays;
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            EulaUri = eulaUri;
            ExpectedAssetCount = expectedAssetCount;
            GuestOsScan = guestOsScan;
            Labels = labels;
            Name = name;
            ServiceAccount = serviceAccount;
            State = state;
            UpdateTime = updateTime;
            VsphereScan = vsphereScan;
        }
    }
}