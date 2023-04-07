// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1Alpha1.Outputs
{

    /// <summary>
    /// Aggregate statistics for a collection of assets.
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryAssetAggregateStatsResponse
    {
        /// <summary>
        /// Count of assets grouped by age.
        /// </summary>
        public readonly Outputs.ReportSummaryChartDataResponse AssetAge;
        /// <summary>
        /// Histogram showing a distribution of CPU core counts.
        /// </summary>
        public readonly Outputs.ReportSummaryHistogramChartDataResponse CoreCountHistogram;
        /// <summary>
        /// Histogram showing a distribution of memory sizes.
        /// </summary>
        public readonly Outputs.ReportSummaryHistogramChartDataResponse MemoryBytesHistogram;
        /// <summary>
        /// Total memory split into Used/Free buckets.
        /// </summary>
        public readonly Outputs.ReportSummaryChartDataResponse MemoryUtilization;
        /// <summary>
        /// Total memory split into Used/Free buckets.
        /// </summary>
        public readonly Outputs.ReportSummaryUtilizationChartDataResponse MemoryUtilizationChart;
        /// <summary>
        /// Count of assets grouped by Operating System families.
        /// </summary>
        public readonly Outputs.ReportSummaryChartDataResponse OperatingSystem;
        /// <summary>
        /// Histogram showing a distribution of memory sizes.
        /// </summary>
        public readonly Outputs.ReportSummaryHistogramChartDataResponse StorageBytesHistogram;
        /// <summary>
        /// Total storage split into Used/Free buckets.
        /// </summary>
        public readonly Outputs.ReportSummaryChartDataResponse StorageUtilization;
        /// <summary>
        /// Total memory split into Used/Free buckets.
        /// </summary>
        public readonly Outputs.ReportSummaryUtilizationChartDataResponse StorageUtilizationChart;
        /// <summary>
        /// Count of the number of unique assets in this collection.
        /// </summary>
        public readonly string TotalAssets;
        /// <summary>
        /// Sum of the CPU core count of all the assets in this collection.
        /// </summary>
        public readonly string TotalCores;
        /// <summary>
        /// Sum of the memory in bytes of all the assets in this collection.
        /// </summary>
        public readonly string TotalMemoryBytes;
        /// <summary>
        /// Sum of persistent storage in bytes of all the assets in this collection.
        /// </summary>
        public readonly string TotalStorageBytes;

        [OutputConstructor]
        private ReportSummaryAssetAggregateStatsResponse(
            Outputs.ReportSummaryChartDataResponse assetAge,

            Outputs.ReportSummaryHistogramChartDataResponse coreCountHistogram,

            Outputs.ReportSummaryHistogramChartDataResponse memoryBytesHistogram,

            Outputs.ReportSummaryChartDataResponse memoryUtilization,

            Outputs.ReportSummaryUtilizationChartDataResponse memoryUtilizationChart,

            Outputs.ReportSummaryChartDataResponse operatingSystem,

            Outputs.ReportSummaryHistogramChartDataResponse storageBytesHistogram,

            Outputs.ReportSummaryChartDataResponse storageUtilization,

            Outputs.ReportSummaryUtilizationChartDataResponse storageUtilizationChart,

            string totalAssets,

            string totalCores,

            string totalMemoryBytes,

            string totalStorageBytes)
        {
            AssetAge = assetAge;
            CoreCountHistogram = coreCountHistogram;
            MemoryBytesHistogram = memoryBytesHistogram;
            MemoryUtilization = memoryUtilization;
            MemoryUtilizationChart = memoryUtilizationChart;
            OperatingSystem = operatingSystem;
            StorageBytesHistogram = storageBytesHistogram;
            StorageUtilization = storageUtilization;
            StorageUtilizationChart = storageUtilizationChart;
            TotalAssets = totalAssets;
            TotalCores = totalCores;
            TotalMemoryBytes = totalMemoryBytes;
            TotalStorageBytes = totalStorageBytes;
        }
    }
}
