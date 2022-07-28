// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSearch.V1.Inputs
{

    /// <summary>
    /// Scoring configurations for a source while processing a Search or Suggest request.
    /// </summary>
    public sealed class ScoringConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use freshness as a ranking signal. By default, freshness is used as a ranking signal. Note that this setting is not available in the Admin UI.
        /// </summary>
        [Input("disableFreshness")]
        public Input<bool>? DisableFreshness { get; set; }

        /// <summary>
        /// Whether to personalize the results. By default, personal signals will be used to boost results.
        /// </summary>
        [Input("disablePersonalization")]
        public Input<bool>? DisablePersonalization { get; set; }

        public ScoringConfigArgs()
        {
        }
        public static new ScoringConfigArgs Empty => new ScoringConfigArgs();
    }
}
