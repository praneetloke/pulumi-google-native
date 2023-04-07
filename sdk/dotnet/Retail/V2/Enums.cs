// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Retail.V2
{
    [EnumType]
    public readonly struct ControlSearchSolutionUseCaseItem : IEquatable<ControlSearchSolutionUseCaseItem>
    {
        private readonly string _value;

        private ControlSearchSolutionUseCaseItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The value when it's unspecified. In this case, server behavior defaults to SEARCH_SOLUTION_USE_CASE_SEARCH.
        /// </summary>
        public static ControlSearchSolutionUseCaseItem SearchSolutionUseCaseUnspecified { get; } = new ControlSearchSolutionUseCaseItem("SEARCH_SOLUTION_USE_CASE_UNSPECIFIED");
        /// <summary>
        /// Search use case. Expects the traffic has a non-empty query.
        /// </summary>
        public static ControlSearchSolutionUseCaseItem SearchSolutionUseCaseSearch { get; } = new ControlSearchSolutionUseCaseItem("SEARCH_SOLUTION_USE_CASE_SEARCH");
        /// <summary>
        /// Browse use case. Expects the traffic has an empty query.
        /// </summary>
        public static ControlSearchSolutionUseCaseItem SearchSolutionUseCaseBrowse { get; } = new ControlSearchSolutionUseCaseItem("SEARCH_SOLUTION_USE_CASE_BROWSE");

        public static bool operator ==(ControlSearchSolutionUseCaseItem left, ControlSearchSolutionUseCaseItem right) => left.Equals(right);
        public static bool operator !=(ControlSearchSolutionUseCaseItem left, ControlSearchSolutionUseCaseItem right) => !left.Equals(right);

        public static explicit operator string(ControlSearchSolutionUseCaseItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ControlSearchSolutionUseCaseItem other && Equals(other);
        public bool Equals(ControlSearchSolutionUseCaseItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ControlSolutionTypesItem : IEquatable<ControlSolutionTypesItem>
    {
        private readonly string _value;

        private ControlSolutionTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static ControlSolutionTypesItem SolutionTypeUnspecified { get; } = new ControlSolutionTypesItem("SOLUTION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Used for Recommendations AI.
        /// </summary>
        public static ControlSolutionTypesItem SolutionTypeRecommendation { get; } = new ControlSolutionTypesItem("SOLUTION_TYPE_RECOMMENDATION");
        /// <summary>
        /// Used for Retail Search.
        /// </summary>
        public static ControlSolutionTypesItem SolutionTypeSearch { get; } = new ControlSolutionTypesItem("SOLUTION_TYPE_SEARCH");

        public static bool operator ==(ControlSolutionTypesItem left, ControlSolutionTypesItem right) => left.Equals(right);
        public static bool operator !=(ControlSolutionTypesItem left, ControlSolutionTypesItem right) => !left.Equals(right);

        public static explicit operator string(ControlSolutionTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ControlSolutionTypesItem other && Equals(other);
        public bool Equals(ControlSolutionTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Mode of the DynamicFacet feature. Defaults to Mode.DISABLED if it's unset.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode : IEquatable<GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode>
    {
        private readonly string _value;

        private GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode ModeUnspecified { get; } = new GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode("MODE_UNSPECIFIED");
        /// <summary>
        /// Disable Dynamic Facet.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode Disabled { get; } = new GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode("DISABLED");
        /// <summary>
        /// Automatic mode built by Google Retail Search.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode Enabled { get; } = new GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode("ENABLED");

        public static bool operator ==(GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode left, GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode left, GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode other && Equals(other);
        public bool Equals(GoogleCloudRetailV2SearchRequestDynamicFacetSpecMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defaults to Mode.AUTO.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRetailV2SearchRequestPersonalizationSpecMode : IEquatable<GoogleCloudRetailV2SearchRequestPersonalizationSpecMode>
    {
        private readonly string _value;

        private GoogleCloudRetailV2SearchRequestPersonalizationSpecMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value. In this case, server behavior defaults to Mode.AUTO.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestPersonalizationSpecMode ModeUnspecified { get; } = new GoogleCloudRetailV2SearchRequestPersonalizationSpecMode("MODE_UNSPECIFIED");
        /// <summary>
        /// Let CRS decide whether to use personalization based on quality of user event data.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestPersonalizationSpecMode Auto { get; } = new GoogleCloudRetailV2SearchRequestPersonalizationSpecMode("AUTO");
        /// <summary>
        /// Disable personalization.
        /// </summary>
        public static GoogleCloudRetailV2SearchRequestPersonalizationSpecMode Disabled { get; } = new GoogleCloudRetailV2SearchRequestPersonalizationSpecMode("DISABLED");

        public static bool operator ==(GoogleCloudRetailV2SearchRequestPersonalizationSpecMode left, GoogleCloudRetailV2SearchRequestPersonalizationSpecMode right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRetailV2SearchRequestPersonalizationSpecMode left, GoogleCloudRetailV2SearchRequestPersonalizationSpecMode right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRetailV2SearchRequestPersonalizationSpecMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRetailV2SearchRequestPersonalizationSpecMode other && Equals(other);
        public bool Equals(GoogleCloudRetailV2SearchRequestPersonalizationSpecMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
    /// </summary>
    [EnumType]
    public readonly struct ModelFilteringOption : IEquatable<ModelFilteringOption>
    {
        private readonly string _value;

        private ModelFilteringOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Value used when unset. In this case, server behavior defaults to RECOMMENDATIONS_FILTERING_DISABLED.
        /// </summary>
        public static ModelFilteringOption RecommendationsFilteringOptionUnspecified { get; } = new ModelFilteringOption("RECOMMENDATIONS_FILTERING_OPTION_UNSPECIFIED");
        /// <summary>
        /// Recommendation filtering is disabled.
        /// </summary>
        public static ModelFilteringOption RecommendationsFilteringDisabled { get; } = new ModelFilteringOption("RECOMMENDATIONS_FILTERING_DISABLED");
        /// <summary>
        /// Recommendation filtering is enabled.
        /// </summary>
        public static ModelFilteringOption RecommendationsFilteringEnabled { get; } = new ModelFilteringOption("RECOMMENDATIONS_FILTERING_ENABLED");

        public static bool operator ==(ModelFilteringOption left, ModelFilteringOption right) => left.Equals(right);
        public static bool operator !=(ModelFilteringOption left, ModelFilteringOption right) => !left.Equals(right);

        public static explicit operator string(ModelFilteringOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ModelFilteringOption other && Equals(other);
        public bool Equals(ModelFilteringOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
    /// </summary>
    [EnumType]
    public readonly struct ModelPeriodicTuningState : IEquatable<ModelPeriodicTuningState>
    {
        private readonly string _value;

        private ModelPeriodicTuningState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified default value, should never be explicitly set.
        /// </summary>
        public static ModelPeriodicTuningState PeriodicTuningStateUnspecified { get; } = new ModelPeriodicTuningState("PERIODIC_TUNING_STATE_UNSPECIFIED");
        /// <summary>
        /// The model has periodic tuning disabled. Tuning can be reenabled by calling the `EnableModelPeriodicTuning` method or by calling the `TuneModel` method.
        /// </summary>
        public static ModelPeriodicTuningState PeriodicTuningDisabled { get; } = new ModelPeriodicTuningState("PERIODIC_TUNING_DISABLED");
        /// <summary>
        /// The model cannot be tuned with periodic tuning OR the `TuneModel` method. Hide the options in customer UI and reject any requests through the backend self serve API.
        /// </summary>
        public static ModelPeriodicTuningState AllTuningDisabled { get; } = new ModelPeriodicTuningState("ALL_TUNING_DISABLED");
        /// <summary>
        /// The model has periodic tuning enabled. Tuning can be disabled by calling the `DisableModelPeriodicTuning` method.
        /// </summary>
        public static ModelPeriodicTuningState PeriodicTuningEnabled { get; } = new ModelPeriodicTuningState("PERIODIC_TUNING_ENABLED");

        public static bool operator ==(ModelPeriodicTuningState left, ModelPeriodicTuningState right) => left.Equals(right);
        public static bool operator !=(ModelPeriodicTuningState left, ModelPeriodicTuningState right) => !left.Equals(right);

        public static explicit operator string(ModelPeriodicTuningState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ModelPeriodicTuningState other && Equals(other);
        public bool Equals(ModelPeriodicTuningState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
    /// </summary>
    [EnumType]
    public readonly struct ModelTrainingState : IEquatable<ModelTrainingState>
    {
        private readonly string _value;

        private ModelTrainingState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified training state.
        /// </summary>
        public static ModelTrainingState TrainingStateUnspecified { get; } = new ModelTrainingState("TRAINING_STATE_UNSPECIFIED");
        /// <summary>
        /// The model training is paused.
        /// </summary>
        public static ModelTrainingState Paused { get; } = new ModelTrainingState("PAUSED");
        /// <summary>
        /// The model is training.
        /// </summary>
        public static ModelTrainingState Training { get; } = new ModelTrainingState("TRAINING");

        public static bool operator ==(ModelTrainingState left, ModelTrainingState right) => left.Equals(right);
        public static bool operator !=(ModelTrainingState left, ModelTrainingState right) => !left.Equals(right);

        public static explicit operator string(ModelTrainingState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ModelTrainingState other && Equals(other);
        public bool Equals(ModelTrainingState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The online availability of the Product. Default to Availability.IN_STOCK. Corresponding properties: Google Merchant Center property [availability](https://support.google.com/merchants/answer/6324448). Schema.org property [Offer.availability](https://schema.org/availability).
    /// </summary>
    [EnumType]
    public readonly struct ProductAvailability : IEquatable<ProductAvailability>
    {
        private readonly string _value;

        private ProductAvailability(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default product availability. Default to Availability.IN_STOCK if unset.
        /// </summary>
        public static ProductAvailability AvailabilityUnspecified { get; } = new ProductAvailability("AVAILABILITY_UNSPECIFIED");
        /// <summary>
        /// Product in stock.
        /// </summary>
        public static ProductAvailability InStock { get; } = new ProductAvailability("IN_STOCK");
        /// <summary>
        /// Product out of stock.
        /// </summary>
        public static ProductAvailability OutOfStock { get; } = new ProductAvailability("OUT_OF_STOCK");
        /// <summary>
        /// Product that is in pre-order state.
        /// </summary>
        public static ProductAvailability Preorder { get; } = new ProductAvailability("PREORDER");
        /// <summary>
        /// Product that is back-ordered (i.e. temporarily out of stock).
        /// </summary>
        public static ProductAvailability Backorder { get; } = new ProductAvailability("BACKORDER");

        public static bool operator ==(ProductAvailability left, ProductAvailability right) => left.Equals(right);
        public static bool operator !=(ProductAvailability left, ProductAvailability right) => !left.Equals(right);

        public static explicit operator string(ProductAvailability value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProductAvailability other && Equals(other);
        public bool Equals(ProductAvailability other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
    /// </summary>
    [EnumType]
    public readonly struct ProductType : IEquatable<ProductType>
    {
        private readonly string _value;

        private ProductType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value. Default to Catalog.product_level_config.ingestion_product_type if unset.
        /// </summary>
        public static ProductType TypeUnspecified { get; } = new ProductType("TYPE_UNSPECIFIED");
        /// <summary>
        /// The primary type. As the primary unit for predicting, indexing and search serving, a Type.PRIMARY Product is grouped with multiple Type.VARIANT Products.
        /// </summary>
        public static ProductType Primary { get; } = new ProductType("PRIMARY");
        /// <summary>
        /// The variant type. Type.VARIANT Products usually share some common attributes on the same Type.PRIMARY Products, but they have variant attributes like different colors, sizes and prices, etc.
        /// </summary>
        public static ProductType Variant { get; } = new ProductType("VARIANT");
        /// <summary>
        /// The collection type. Collection products are bundled Type.PRIMARY Products or Type.VARIANT Products that are sold together, such as a jewelry set with necklaces, earrings and rings, etc.
        /// </summary>
        public static ProductType Collection { get; } = new ProductType("COLLECTION");

        public static bool operator ==(ProductType left, ProductType right) => left.Equals(right);
        public static bool operator !=(ProductType left, ProductType right) => !left.Equals(right);

        public static explicit operator string(ProductType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProductType other && Equals(other);
        public bool Equals(ProductType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// What kind of diversity to use - data driven or rule based. If unset, the server behavior defaults to RULE_BASED_DIVERSITY.
    /// </summary>
    [EnumType]
    public readonly struct ServingConfigDiversityType : IEquatable<ServingConfigDiversityType>
    {
        private readonly string _value;

        private ServingConfigDiversityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static ServingConfigDiversityType DiversityTypeUnspecified { get; } = new ServingConfigDiversityType("DIVERSITY_TYPE_UNSPECIFIED");
        /// <summary>
        /// Rule based diversity.
        /// </summary>
        public static ServingConfigDiversityType RuleBasedDiversity { get; } = new ServingConfigDiversityType("RULE_BASED_DIVERSITY");
        /// <summary>
        /// Data driven diversity.
        /// </summary>
        public static ServingConfigDiversityType DataDrivenDiversity { get; } = new ServingConfigDiversityType("DATA_DRIVEN_DIVERSITY");

        public static bool operator ==(ServingConfigDiversityType left, ServingConfigDiversityType right) => left.Equals(right);
        public static bool operator !=(ServingConfigDiversityType left, ServingConfigDiversityType right) => !left.Equals(right);

        public static explicit operator string(ServingConfigDiversityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServingConfigDiversityType other && Equals(other);
        public bool Equals(ServingConfigDiversityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServingConfigSolutionTypesItem : IEquatable<ServingConfigSolutionTypesItem>
    {
        private readonly string _value;

        private ServingConfigSolutionTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static ServingConfigSolutionTypesItem SolutionTypeUnspecified { get; } = new ServingConfigSolutionTypesItem("SOLUTION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Used for Recommendations AI.
        /// </summary>
        public static ServingConfigSolutionTypesItem SolutionTypeRecommendation { get; } = new ServingConfigSolutionTypesItem("SOLUTION_TYPE_RECOMMENDATION");
        /// <summary>
        /// Used for Retail Search.
        /// </summary>
        public static ServingConfigSolutionTypesItem SolutionTypeSearch { get; } = new ServingConfigSolutionTypesItem("SOLUTION_TYPE_SEARCH");

        public static bool operator ==(ServingConfigSolutionTypesItem left, ServingConfigSolutionTypesItem right) => left.Equals(right);
        public static bool operator !=(ServingConfigSolutionTypesItem left, ServingConfigSolutionTypesItem right) => !left.Equals(right);

        public static explicit operator string(ServingConfigSolutionTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServingConfigSolutionTypesItem other && Equals(other);
        public bool Equals(ServingConfigSolutionTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
