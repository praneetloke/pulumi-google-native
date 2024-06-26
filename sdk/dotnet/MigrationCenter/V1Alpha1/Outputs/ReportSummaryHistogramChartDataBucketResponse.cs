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
    /// A histogram bucket with a lower and upper bound, and a count of items with a field value between those bounds. The lower bound is inclusive and the upper bound is exclusive. Lower bound may be -infinity and upper bound may be infinity.
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryHistogramChartDataBucketResponse
    {
        /// <summary>
        /// Count of items in the bucket.
        /// </summary>
        public readonly string Count;
        /// <summary>
        /// Lower bound - inclusive.
        /// </summary>
        public readonly string LowerBound;
        /// <summary>
        /// Upper bound - exclusive.
        /// </summary>
        public readonly string UpperBound;

        [OutputConstructor]
        private ReportSummaryHistogramChartDataBucketResponse(
            string count,

            string lowerBound,

            string upperBound)
        {
            Count = count;
            LowerBound = lowerBound;
            UpperBound = upperBound;
        }
    }
}
