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
    /// Default options to interpret user query.
    /// </summary>
    public sealed class QueryInterpretationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set this flag to disable supplemental results retrieval, setting a flag here will not retrieve supplemental results for queries associated with a given search application. If this flag is set to True, it will take precedence over the option set at Query level. For the default value of False, query level flag will set the correct interpretation for supplemental results.
        /// </summary>
        [Input("forceDisableSupplementalResults")]
        public Input<bool>? ForceDisableSupplementalResults { get; set; }

        /// <summary>
        /// Enable this flag to turn off all internal optimizations like natural language (NL) interpretation of queries, supplemental results retrieval, and usage of synonyms including custom ones. If this flag is set to True, it will take precedence over the option set at Query level. For the default value of False, query level flag will set the correct interpretation for verbatim mode.
        /// </summary>
        [Input("forceVerbatimMode")]
        public Input<bool>? ForceVerbatimMode { get; set; }

        public QueryInterpretationConfigArgs()
        {
        }
        public static new QueryInterpretationConfigArgs Empty => new QueryInterpretationConfigArgs();
    }
}
