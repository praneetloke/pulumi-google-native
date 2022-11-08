// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Runtime information about workload execution.
    /// </summary>
    [OutputType]
    public sealed class RuntimeInfoResponse
    {
        /// <summary>
        /// Approximate workload resource usage calculated after workload finishes.
        /// </summary>
        public readonly Outputs.UsageMetricsResponse ApproximateUsage;
        /// <summary>
        /// A URI pointing to the location of the diagnostics tarball.
        /// </summary>
        public readonly string DiagnosticOutputUri;
        /// <summary>
        /// Map of remote access endpoints (such as web interfaces and APIs) to their URIs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Endpoints;
        /// <summary>
        /// A URI pointing to the location of the stdout and stderr of the workload.
        /// </summary>
        public readonly string OutputUri;

        [OutputConstructor]
        private RuntimeInfoResponse(
            Outputs.UsageMetricsResponse approximateUsage,

            string diagnosticOutputUri,

            ImmutableDictionary<string, string> endpoints,

            string outputUri)
        {
            ApproximateUsage = approximateUsage;
            DiagnosticOutputUri = diagnosticOutputUri;
            Endpoints = endpoints;
            OutputUri = outputUri;
        }
    }
}
