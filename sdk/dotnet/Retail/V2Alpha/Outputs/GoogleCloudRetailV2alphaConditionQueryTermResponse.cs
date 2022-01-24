// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// Query terms that we want to match on.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaConditionQueryTermResponse
    {
        /// <summary>
        /// Whether this is supposed to be a full or partial match.
        /// </summary>
        public readonly bool FullMatch;
        /// <summary>
        /// The value of the term to match on. Value cannot be empty. Value can have at most 3 terms if specified as a partial match. Each space separated string is considered as one term. Example) "a b c" is 3 terms and allowed, " a b c d" is 4 terms and not allowed for partial match.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaConditionQueryTermResponse(
            bool fullMatch,

            string value)
        {
            FullMatch = fullMatch;
            Value = value;
        }
    }
}
