// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AnalyticsHub.V1Beta1.Outputs
{

    /// <summary>
    /// Restricted export config, used to configure restricted export on linked dataset.
    /// </summary>
    [OutputType]
    public sealed class RestrictedExportConfigResponse
    {
        /// <summary>
        /// If true, enable restricted export.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// If true, restrict export of query result derived from restricted linked dataset table.
        /// </summary>
        public readonly bool RestrictQueryResult;

        [OutputConstructor]
        private RestrictedExportConfigResponse(
            bool enabled,

            bool restrictQueryResult)
        {
            Enabled = enabled;
            RestrictQueryResult = restrictQueryResult;
        }
    }
}
