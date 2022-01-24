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
    /// A rule is a condition-action pair * A condition defines when a rule is to be triggered. * An action specifies what occurs on that trigger. Currently only boost rules are supported. Currently only supported by the search endpoint.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaRuleResponse
    {
        /// <summary>
        /// A boost action.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleBoostActionResponse BoostAction;
        /// <summary>
        /// The condition that triggers the rule. If the condition is empty, the rule will always apply.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaConditionResponse Condition;
        /// <summary>
        /// Prevents term from being associated with other terms.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleDoNotAssociateActionResponse DoNotAssociateAction;
        /// <summary>
        /// Filters results.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleFilterActionResponse FilterAction;
        /// <summary>
        /// Ignores specific terms from query during search.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleIgnoreActionResponse IgnoreAction;
        /// <summary>
        /// Treats specific term as a synonym with a group of terms. Group of terms will not be treated as synonyms with the specific term.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleOnewaySynonymsActionResponse OnewaySynonymsAction;
        /// <summary>
        /// Redirects a shopper to a specific page.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleRedirectActionResponse RedirectAction;
        /// <summary>
        /// Replaces specific terms in the query.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleReplacementActionResponse ReplacementAction;
        /// <summary>
        /// Treats a set of terms as synonyms of one another.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleTwowaySynonymsActionResponse TwowaySynonymsAction;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaRuleResponse(
            Outputs.GoogleCloudRetailV2alphaRuleBoostActionResponse boostAction,

            Outputs.GoogleCloudRetailV2alphaConditionResponse condition,

            Outputs.GoogleCloudRetailV2alphaRuleDoNotAssociateActionResponse doNotAssociateAction,

            Outputs.GoogleCloudRetailV2alphaRuleFilterActionResponse filterAction,

            Outputs.GoogleCloudRetailV2alphaRuleIgnoreActionResponse ignoreAction,

            Outputs.GoogleCloudRetailV2alphaRuleOnewaySynonymsActionResponse onewaySynonymsAction,

            Outputs.GoogleCloudRetailV2alphaRuleRedirectActionResponse redirectAction,

            Outputs.GoogleCloudRetailV2alphaRuleReplacementActionResponse replacementAction,

            Outputs.GoogleCloudRetailV2alphaRuleTwowaySynonymsActionResponse twowaySynonymsAction)
        {
            BoostAction = boostAction;
            Condition = condition;
            DoNotAssociateAction = doNotAssociateAction;
            FilterAction = filterAction;
            IgnoreAction = ignoreAction;
            OnewaySynonymsAction = onewaySynonymsAction;
            RedirectAction = redirectAction;
            ReplacementAction = replacementAction;
            TwowaySynonymsAction = twowaySynonymsAction;
        }
    }
}
