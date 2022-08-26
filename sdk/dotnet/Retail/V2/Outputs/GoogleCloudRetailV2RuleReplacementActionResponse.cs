// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2.Outputs
{

    /// <summary>
    /// Replaces a term in the query. Multiple replacement candidates can be specified. All `query_terms` will be replaced with the replacement term. Example: Replace "gShoe" with "google shoe".
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2RuleReplacementActionResponse
    {
        /// <summary>
        /// Terms from the search query. Will be replaced by replacement term. Can specify up to 100 terms.
        /// </summary>
        public readonly ImmutableArray<string> QueryTerms;
        /// <summary>
        /// Term that will be used for replacement.
        /// </summary>
        public readonly string ReplacementTerm;
        /// <summary>
        /// Will be [deprecated = true] post migration;
        /// </summary>
        public readonly string Term;

        [OutputConstructor]
        private GoogleCloudRetailV2RuleReplacementActionResponse(
            ImmutableArray<string> queryTerms,

            string replacementTerm,

            string term)
        {
            QueryTerms = queryTerms;
            ReplacementTerm = replacementTerm;
            Term = term;
        }
    }
}