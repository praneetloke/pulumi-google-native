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
    /// Describes the Summary view of a Report, which contains aggregated values for all the groups and preference sets included in this Report.
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryResponse
    {
        /// <summary>
        /// Aggregate statistics for all the assets across all the groups.
        /// </summary>
        public readonly Outputs.ReportSummaryAssetAggregateStatsResponse AllAssetsStats;
        /// <summary>
        /// Findings for each Group included in this report.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportSummaryGroupFindingResponse> GroupFindings;

        [OutputConstructor]
        private ReportSummaryResponse(
            Outputs.ReportSummaryAssetAggregateStatsResponse allAssetsStats,

            ImmutableArray<Outputs.ReportSummaryGroupFindingResponse> groupFindings)
        {
            AllAssetsStats = allAssetsStats;
            GroupFindings = groupFindings;
        }
    }
}
