// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ControlSearchSolutionUseCaseItem = {
    /**
     * The value when it's unspecified. In this case, server behavior defaults to SEARCH_SOLUTION_USE_CASE_SEARCH.
     */
    SearchSolutionUseCaseUnspecified: "SEARCH_SOLUTION_USE_CASE_UNSPECIFIED",
    /**
     * Search use case. Expects the traffic has a non-empty query.
     */
    SearchSolutionUseCaseSearch: "SEARCH_SOLUTION_USE_CASE_SEARCH",
    /**
     * Browse use case. Expects the traffic has an empty query.
     */
    SearchSolutionUseCaseBrowse: "SEARCH_SOLUTION_USE_CASE_BROWSE",
} as const;

export type ControlSearchSolutionUseCaseItem = (typeof ControlSearchSolutionUseCaseItem)[keyof typeof ControlSearchSolutionUseCaseItem];

export const ControlSolutionTypesItem = {
    /**
     * Default value.
     */
    SolutionTypeUnspecified: "SOLUTION_TYPE_UNSPECIFIED",
    /**
     * Used for Recommendations AI.
     */
    SolutionTypeRecommendation: "SOLUTION_TYPE_RECOMMENDATION",
    /**
     * Used for Retail Search.
     */
    SolutionTypeSearch: "SOLUTION_TYPE_SEARCH",
} as const;

export type ControlSolutionTypesItem = (typeof ControlSolutionTypesItem)[keyof typeof ControlSolutionTypesItem];

export const GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode = {
    /**
     * Default value.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * Disable Dynamic Facet.
     */
    Disabled: "DISABLED",
    /**
     * Automatic mode built by Google Retail Search.
     */
    Enabled: "ENABLED",
} as const;

/**
 * Mode of the DynamicFacet feature. Defaults to Mode.DISABLED if it's unset.
 */
export type GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode = (typeof GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode)[keyof typeof GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode];

export const GoogleCloudRetailV2SearchRequestPersonalizationSpecMode = {
    /**
     * Default value. In this case, server behavior defaults to Mode.AUTO.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * Let CRS decide whether to use personalization based on quality of user event data.
     */
    Auto: "AUTO",
    /**
     * Disable personalization.
     */
    Disabled: "DISABLED",
} as const;

/**
 * Defaults to Mode.AUTO.
 */
export type GoogleCloudRetailV2SearchRequestPersonalizationSpecMode = (typeof GoogleCloudRetailV2SearchRequestPersonalizationSpecMode)[keyof typeof GoogleCloudRetailV2SearchRequestPersonalizationSpecMode];

export const ModelFilteringOption = {
    /**
     * Value used when unset. In this case, server behavior defaults to RECOMMENDATIONS_FILTERING_DISABLED.
     */
    RecommendationsFilteringOptionUnspecified: "RECOMMENDATIONS_FILTERING_OPTION_UNSPECIFIED",
    /**
     * Recommendation filtering is disabled.
     */
    RecommendationsFilteringDisabled: "RECOMMENDATIONS_FILTERING_DISABLED",
    /**
     * Recommendation filtering is enabled.
     */
    RecommendationsFilteringEnabled: "RECOMMENDATIONS_FILTERING_ENABLED",
} as const;

/**
 * Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
 */
export type ModelFilteringOption = (typeof ModelFilteringOption)[keyof typeof ModelFilteringOption];

export const ModelPeriodicTuningState = {
    /**
     * Unspecified default value, should never be explicitly set.
     */
    PeriodicTuningStateUnspecified: "PERIODIC_TUNING_STATE_UNSPECIFIED",
    /**
     * The model has periodic tuning disabled. Tuning can be reenabled by calling the `EnableModelPeriodicTuning` method or by calling the `TuneModel` method.
     */
    PeriodicTuningDisabled: "PERIODIC_TUNING_DISABLED",
    /**
     * The model cannot be tuned with periodic tuning OR the `TuneModel` method. Hide the options in customer UI and reject any requests through the backend self serve API.
     */
    AllTuningDisabled: "ALL_TUNING_DISABLED",
    /**
     * The model has periodic tuning enabled. Tuning can be disabled by calling the `DisableModelPeriodicTuning` method.
     */
    PeriodicTuningEnabled: "PERIODIC_TUNING_ENABLED",
} as const;

/**
 * Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
 */
export type ModelPeriodicTuningState = (typeof ModelPeriodicTuningState)[keyof typeof ModelPeriodicTuningState];

export const ModelTrainingState = {
    /**
     * Unspecified training state.
     */
    TrainingStateUnspecified: "TRAINING_STATE_UNSPECIFIED",
    /**
     * The model training is paused.
     */
    Paused: "PAUSED",
    /**
     * The model is training.
     */
    Training: "TRAINING",
} as const;

/**
 * Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
 */
export type ModelTrainingState = (typeof ModelTrainingState)[keyof typeof ModelTrainingState];

export const ProductAvailability = {
    /**
     * Default product availability. Default to Availability.IN_STOCK if unset.
     */
    AvailabilityUnspecified: "AVAILABILITY_UNSPECIFIED",
    /**
     * Product in stock.
     */
    InStock: "IN_STOCK",
    /**
     * Product out of stock.
     */
    OutOfStock: "OUT_OF_STOCK",
    /**
     * Product that is in pre-order state.
     */
    Preorder: "PREORDER",
    /**
     * Product that is back-ordered (i.e. temporarily out of stock).
     */
    Backorder: "BACKORDER",
} as const;

/**
 * The online availability of the Product. Default to Availability.IN_STOCK. Corresponding properties: Google Merchant Center property [availability](https://support.google.com/merchants/answer/6324448). Schema.org property [Offer.availability](https://schema.org/availability).
 */
export type ProductAvailability = (typeof ProductAvailability)[keyof typeof ProductAvailability];

export const ProductType = {
    /**
     * Default value. Default to Catalog.product_level_config.ingestion_product_type if unset.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * The primary type. As the primary unit for predicting, indexing and search serving, a Type.PRIMARY Product is grouped with multiple Type.VARIANT Products.
     */
    Primary: "PRIMARY",
    /**
     * The variant type. Type.VARIANT Products usually share some common attributes on the same Type.PRIMARY Products, but they have variant attributes like different colors, sizes and prices, etc.
     */
    Variant: "VARIANT",
    /**
     * The collection type. Collection products are bundled Type.PRIMARY Products or Type.VARIANT Products that are sold together, such as a jewelry set with necklaces, earrings and rings, etc.
     */
    Collection: "COLLECTION",
} as const;

/**
 * Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
 */
export type ProductType = (typeof ProductType)[keyof typeof ProductType];

export const ServingConfigDiversityType = {
    /**
     * Default value.
     */
    DiversityTypeUnspecified: "DIVERSITY_TYPE_UNSPECIFIED",
    /**
     * Rule based diversity.
     */
    RuleBasedDiversity: "RULE_BASED_DIVERSITY",
    /**
     * Data driven diversity.
     */
    DataDrivenDiversity: "DATA_DRIVEN_DIVERSITY",
} as const;

/**
 * What kind of diversity to use - data driven or rule based. If unset, the server behavior defaults to RULE_BASED_DIVERSITY.
 */
export type ServingConfigDiversityType = (typeof ServingConfigDiversityType)[keyof typeof ServingConfigDiversityType];

export const ServingConfigSolutionTypesItem = {
    /**
     * Default value.
     */
    SolutionTypeUnspecified: "SOLUTION_TYPE_UNSPECIFIED",
    /**
     * Used for Recommendations AI.
     */
    SolutionTypeRecommendation: "SOLUTION_TYPE_RECOMMENDATION",
    /**
     * Used for Retail Search.
     */
    SolutionTypeSearch: "SOLUTION_TYPE_SEARCH",
} as const;

export type ServingConfigSolutionTypesItem = (typeof ServingConfigSolutionTypesItem)[keyof typeof ServingConfigSolutionTypesItem];
