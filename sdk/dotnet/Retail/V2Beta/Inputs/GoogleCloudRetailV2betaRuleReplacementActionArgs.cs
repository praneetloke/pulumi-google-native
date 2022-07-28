// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta.Inputs
{

    /// <summary>
    /// Replaces a term in the query. Multiple replacement candidates can be specified. All `query_terms` will be replaced with the replacement term. Example: Replace "gShoe" with "google shoe".
    /// </summary>
    public sealed class GoogleCloudRetailV2betaRuleReplacementActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("queryTerms")]
        private InputList<string>? _queryTerms;

        /// <summary>
        /// Terms from the search query. Will be replaced by replacement term. Can specify up to 100 terms.
        /// </summary>
        public InputList<string> QueryTerms
        {
            get => _queryTerms ?? (_queryTerms = new InputList<string>());
            set => _queryTerms = value;
        }

        /// <summary>
        /// Term that will be used for replacement.
        /// </summary>
        [Input("replacementTerm")]
        public Input<string>? ReplacementTerm { get; set; }

        /// <summary>
        /// Will be [deprecated = true] post migration;
        /// </summary>
        [Input("term")]
        public Input<string>? Term { get; set; }

        public GoogleCloudRetailV2betaRuleReplacementActionArgs()
        {
        }
        public static new GoogleCloudRetailV2betaRuleReplacementActionArgs Empty => new GoogleCloudRetailV2betaRuleReplacementActionArgs();
    }
}
