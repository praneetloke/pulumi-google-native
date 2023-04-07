// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// The results of an Action.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2ActionDetailsResponse
    {
        /// <summary>
        /// Outcome of a de-identification action.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DeidentifyDataSourceDetailsResponse DeidentifyDetails;

        [OutputConstructor]
        private GooglePrivacyDlpV2ActionDetailsResponse(Outputs.GooglePrivacyDlpV2DeidentifyDataSourceDetailsResponse deidentifyDetails)
        {
            DeidentifyDetails = deidentifyDetails;
        }
    }
}
