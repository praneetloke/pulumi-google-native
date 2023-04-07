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
    /// Configuration description of the scanning process. When used with redactContent only info_types and min_likelihood are currently used.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2InspectConfigResponse
    {
        /// <summary>
        /// Deprecated and unused.
        /// </summary>
        public readonly ImmutableArray<string> ContentOptions;
        /// <summary>
        /// CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2CustomInfoTypeResponse> CustomInfoTypes;
        /// <summary>
        /// When true, excludes type information of the findings. This is not used for data profiling.
        /// </summary>
        public readonly bool ExcludeInfoTypes;
        /// <summary>
        /// When true, a contextual quote from the data that triggered a finding is included in the response; see Finding.quote. This is not used for data profiling.
        /// </summary>
        public readonly bool IncludeQuote;
        /// <summary>
        /// Restricts what info_types to look for. The values must correspond to InfoType values returned by ListInfoTypes or listed at https://cloud.google.com/dlp/docs/infotypes-reference. When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run. By default this may be all types, but may change over time as detectors are updated. If you need precise control and predictability as to what detectors are run you should specify specific InfoTypes listed in the reference, otherwise a default list will be used, which may change over time.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeResponse> InfoTypes;
        /// <summary>
        /// Configuration to control the number of findings returned. This is not used for data profiling. When redacting sensitive data from images, finding limits don't apply. They can cause unexpected or inconsistent results, where only some data is redacted. Don't include finding limits in RedactImage requests. Otherwise, Cloud DLP returns an error. When set within `InspectJobConfig`, the specified maximum values aren't hard limits. If an inspection job reaches these limits, the job ends gradually, not abruptly. Therefore, the actual number of findings that Cloud DLP returns can be multiple times higher than these maximum values.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2FindingLimitsResponse Limits;
        /// <summary>
        /// Only returns findings equal or above this threshold. The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood to learn more.
        /// </summary>
        public readonly string MinLikelihood;
        /// <summary>
        /// Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end, other rules are executed in the order they are specified for each info type.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2InspectionRuleSetResponse> RuleSet;

        [OutputConstructor]
        private GooglePrivacyDlpV2InspectConfigResponse(
            ImmutableArray<string> contentOptions,

            ImmutableArray<Outputs.GooglePrivacyDlpV2CustomInfoTypeResponse> customInfoTypes,

            bool excludeInfoTypes,

            bool includeQuote,

            ImmutableArray<Outputs.GooglePrivacyDlpV2InfoTypeResponse> infoTypes,

            Outputs.GooglePrivacyDlpV2FindingLimitsResponse limits,

            string minLikelihood,

            ImmutableArray<Outputs.GooglePrivacyDlpV2InspectionRuleSetResponse> ruleSet)
        {
            ContentOptions = contentOptions;
            CustomInfoTypes = customInfoTypes;
            ExcludeInfoTypes = excludeInfoTypes;
            IncludeQuote = includeQuote;
            InfoTypes = infoTypes;
            Limits = limits;
            MinLikelihood = minLikelihood;
            RuleSet = ruleSet;
        }
    }
}
