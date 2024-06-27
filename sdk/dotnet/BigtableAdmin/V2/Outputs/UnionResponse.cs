// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// A GcRule which deletes cells matching any of the given rules.
    /// </summary>
    [OutputType]
    public sealed class UnionResponse
    {
        /// <summary>
        /// Delete cells which would be deleted by any element of `rules`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GcRuleResponse> Rules;

        [OutputConstructor]
        private UnionResponse(ImmutableArray<Outputs.GcRuleResponse> rules)
        {
            Rules = rules;
        }
    }
}