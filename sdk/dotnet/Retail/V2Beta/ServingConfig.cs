// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta
{
    /// <summary>
    /// Creates a ServingConfig. A maximum of 100 ServingConfigs are allowed in a Catalog, otherwise a FAILED_PRECONDITION error is returned.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:retail/v2beta:ServingConfig")]
    public partial class ServingConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("boostControlIds")]
        public Output<ImmutableArray<string>> BoostControlIds { get; private set; } = null!;

        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Output("diversityLevel")]
        public Output<string> DiversityLevel { get; private set; } = null!;

        /// <summary>
        /// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("doNotAssociateControlIds")]
        public Output<ImmutableArray<string>> DoNotAssociateControlIds { get; private set; } = null!;

        /// <summary>
        /// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("dynamicFacetSpec")]
        public Output<Outputs.GoogleCloudRetailV2betaSearchRequestDynamicFacetSpecResponse> DynamicFacetSpec { get; private set; } = null!;

        /// <summary>
        /// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Output("enableCategoryFilterLevel")]
        public Output<string> EnableCategoryFilterLevel { get; private set; } = null!;

        /// <summary>
        /// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("facetControlIds")]
        public Output<ImmutableArray<string>> FacetControlIds { get; private set; } = null!;

        /// <summary>
        /// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("filterControlIds")]
        public Output<ImmutableArray<string>> FilterControlIds { get; private set; } = null!;

        /// <summary>
        /// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("ignoreControlIds")]
        public Output<ImmutableArray<string>> IgnoreControlIds { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Output("modelId")]
        public Output<string> ModelId { get; private set; } = null!;

        /// <summary>
        /// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("onewaySynonymsControlIds")]
        public Output<ImmutableArray<string>> OnewaySynonymsControlIds { get; private set; } = null!;

        /// <summary>
        /// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Output("priceRerankingLevel")]
        public Output<string> PriceRerankingLevel { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("redirectControlIds")]
        public Output<ImmutableArray<string>> RedirectControlIds { get; private set; } = null!;

        /// <summary>
        /// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("replacementControlIds")]
        public Output<ImmutableArray<string>> ReplacementControlIds { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the ServingConfig, which will become the final component of the ServingConfig's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        /// </summary>
        [Output("servingConfigId")]
        public Output<string> ServingConfigId { get; private set; } = null!;

        /// <summary>
        /// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
        /// </summary>
        [Output("solutionTypes")]
        public Output<ImmutableArray<string>> SolutionTypes { get; private set; } = null!;

        /// <summary>
        /// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("twowaySynonymsControlIds")]
        public Output<ImmutableArray<string>> TwowaySynonymsControlIds { get; private set; } = null!;


        /// <summary>
        /// Create a ServingConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServingConfig(string name, ServingConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:retail/v2beta:ServingConfig", name, args ?? new ServingConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServingConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:retail/v2beta:ServingConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "catalogId",
                    "location",
                    "project",
                    "servingConfigId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServingConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServingConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServingConfig(name, id, options);
        }
    }

    public sealed class ServingConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("boostControlIds")]
        private InputList<string>? _boostControlIds;

        /// <summary>
        /// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> BoostControlIds
        {
            get => _boostControlIds ?? (_boostControlIds = new InputList<string>());
            set => _boostControlIds = value;
        }

        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Input("diversityLevel")]
        public Input<string>? DiversityLevel { get; set; }

        [Input("doNotAssociateControlIds")]
        private InputList<string>? _doNotAssociateControlIds;

        /// <summary>
        /// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> DoNotAssociateControlIds
        {
            get => _doNotAssociateControlIds ?? (_doNotAssociateControlIds = new InputList<string>());
            set => _doNotAssociateControlIds = value;
        }

        /// <summary>
        /// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Input("dynamicFacetSpec")]
        public Input<Inputs.GoogleCloudRetailV2betaSearchRequestDynamicFacetSpecArgs>? DynamicFacetSpec { get; set; }

        /// <summary>
        /// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Input("enableCategoryFilterLevel")]
        public Input<string>? EnableCategoryFilterLevel { get; set; }

        [Input("facetControlIds")]
        private InputList<string>? _facetControlIds;

        /// <summary>
        /// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> FacetControlIds
        {
            get => _facetControlIds ?? (_facetControlIds = new InputList<string>());
            set => _facetControlIds = value;
        }

        [Input("filterControlIds")]
        private InputList<string>? _filterControlIds;

        /// <summary>
        /// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> FilterControlIds
        {
            get => _filterControlIds ?? (_filterControlIds = new InputList<string>());
            set => _filterControlIds = value;
        }

        [Input("ignoreControlIds")]
        private InputList<string>? _ignoreControlIds;

        /// <summary>
        /// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> IgnoreControlIds
        {
            get => _ignoreControlIds ?? (_ignoreControlIds = new InputList<string>());
            set => _ignoreControlIds = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Input("modelId")]
        public Input<string>? ModelId { get; set; }

        /// <summary>
        /// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onewaySynonymsControlIds")]
        private InputList<string>? _onewaySynonymsControlIds;

        /// <summary>
        /// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> OnewaySynonymsControlIds
        {
            get => _onewaySynonymsControlIds ?? (_onewaySynonymsControlIds = new InputList<string>());
            set => _onewaySynonymsControlIds = value;
        }

        /// <summary>
        /// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        [Input("priceRerankingLevel")]
        public Input<string>? PriceRerankingLevel { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("redirectControlIds")]
        private InputList<string>? _redirectControlIds;

        /// <summary>
        /// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> RedirectControlIds
        {
            get => _redirectControlIds ?? (_redirectControlIds = new InputList<string>());
            set => _redirectControlIds = value;
        }

        [Input("replacementControlIds")]
        private InputList<string>? _replacementControlIds;

        /// <summary>
        /// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> ReplacementControlIds
        {
            get => _replacementControlIds ?? (_replacementControlIds = new InputList<string>());
            set => _replacementControlIds = value;
        }

        /// <summary>
        /// Required. The ID to use for the ServingConfig, which will become the final component of the ServingConfig's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        /// </summary>
        [Input("servingConfigId", required: true)]
        public Input<string> ServingConfigId { get; set; } = null!;

        [Input("solutionTypes", required: true)]
        private InputList<Pulumi.GoogleNative.Retail.V2Beta.ServingConfigSolutionTypesItem>? _solutionTypes;

        /// <summary>
        /// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Retail.V2Beta.ServingConfigSolutionTypesItem> SolutionTypes
        {
            get => _solutionTypes ?? (_solutionTypes = new InputList<Pulumi.GoogleNative.Retail.V2Beta.ServingConfigSolutionTypesItem>());
            set => _solutionTypes = value;
        }

        [Input("twowaySynonymsControlIds")]
        private InputList<string>? _twowaySynonymsControlIds;

        /// <summary>
        /// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<string> TwowaySynonymsControlIds
        {
            get => _twowaySynonymsControlIds ?? (_twowaySynonymsControlIds = new InputList<string>());
            set => _twowaySynonymsControlIds = value;
        }

        public ServingConfigArgs()
        {
        }
        public static new ServingConfigArgs Empty => new ServingConfigArgs();
    }
}
