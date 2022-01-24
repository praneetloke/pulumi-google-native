// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Inputs
{

    /// <summary>
    /// A boost action to apply to results matching condition specified above.
    /// </summary>
    public sealed class GoogleCloudRetailV2alphaRuleBoostActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Strength of the condition boost, which must be in [-1, 1]. Negative boost means demotion. Default is 0.0. Setting to 1.0 gives the item a big promotion. However, it does not necessarily mean that the boosted item will be the top result at all times, nor that other items will be excluded. Results could still be shown even when none of them matches the condition. And results that are significantly more relevant to the search query can still trump your heavily favored but irrelevant items. Setting to -1.0 gives the item a big demotion. However, results that are deeply relevant might still be shown. The item will have an upstream battle to get a fairly high ranking, but it is not blocked out completely. Setting to 0.0 means no boost applied. The boosting condition is ignored.
        /// </summary>
        [Input("boost")]
        public Input<double>? Boost { get; set; }

        /// <summary>
        /// The filter can have a max size of 5000 characters. An expression which specifies which products to apply an action to. The syntax and supported fields are the same as a filter expression. See SearchRequest.filter for detail syntax and limitations. Examples: * To boost products with product ID "product_1" or "product_2", and color "Red" or "Blue": *(id: ANY("product_1", "product_2")) * *AND * *(colorFamily: ANY("Red", "Blue")) *
        /// </summary>
        [Input("productsFilter")]
        public Input<string>? ProductsFilter { get; set; }

        public GoogleCloudRetailV2alphaRuleBoostActionArgs()
        {
        }
    }
}
