// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// Provides information about the analysis status of a discovered resource.
    /// </summary>
    [OutputType]
    public sealed class DiscoveryOccurrenceResponse
    {
        public readonly Outputs.AnalysisCompletedResponse AnalysisCompleted;
        /// <summary>
        /// Indicates any errors encountered during analysis of a resource. There could be 0 or more of these errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusResponse> AnalysisError;
        /// <summary>
        /// The status of discovery for the resource.
        /// </summary>
        public readonly string AnalysisStatus;
        /// <summary>
        /// When an error is encountered this will contain a LocalizedMessage under details to show to the user. The LocalizedMessage is output only and populated by the API.
        /// </summary>
        public readonly Outputs.StatusResponse AnalysisStatusError;
        /// <summary>
        /// The time occurrences related to this discovery occurrence were archived.
        /// </summary>
        public readonly string ArchiveTime;
        /// <summary>
        /// Whether the resource is continuously analyzed.
        /// </summary>
        public readonly string ContinuousAnalysis;
        /// <summary>
        /// The CPE of the resource being scanned.
        /// </summary>
        public readonly string Cpe;
        /// <summary>
        /// The last time this resource was scanned.
        /// </summary>
        public readonly string LastScanTime;

        [OutputConstructor]
        private DiscoveryOccurrenceResponse(
            Outputs.AnalysisCompletedResponse analysisCompleted,

            ImmutableArray<Outputs.StatusResponse> analysisError,

            string analysisStatus,

            Outputs.StatusResponse analysisStatusError,

            string archiveTime,

            string continuousAnalysis,

            string cpe,

            string lastScanTime)
        {
            AnalysisCompleted = analysisCompleted;
            AnalysisError = analysisError;
            AnalysisStatus = analysisStatus;
            AnalysisStatusError = analysisStatusError;
            ArchiveTime = archiveTime;
            ContinuousAnalysis = continuousAnalysis;
            Cpe = cpe;
            LastScanTime = lastScanTime;
        }
    }
}
