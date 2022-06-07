// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ServingConfig. A maximum of 100 ServingConfigs are allowed in a Catalog, otherwise a FAILED_PRECONDITION error is returned.
// Auto-naming is currently not supported for this resource.
type ServingConfig struct {
	pulumi.CustomResourceState

	// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	BoostControlIds pulumi.StringArrayOutput `pulumi:"boostControlIds"`
	// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	DiversityLevel pulumi.StringOutput `pulumi:"diversityLevel"`
	// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DoNotAssociateControlIds pulumi.StringArrayOutput `pulumi:"doNotAssociateControlIds"`
	// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DynamicFacetSpec GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput `pulumi:"dynamicFacetSpec"`
	// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	EnableCategoryFilterLevel pulumi.StringOutput `pulumi:"enableCategoryFilterLevel"`
	// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FacetControlIds pulumi.StringArrayOutput `pulumi:"facetControlIds"`
	// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FilterControlIds pulumi.StringArrayOutput `pulumi:"filterControlIds"`
	// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	IgnoreControlIds pulumi.StringArrayOutput `pulumi:"ignoreControlIds"`
	// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
	ModelId pulumi.StringOutput `pulumi:"modelId"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
	Name pulumi.StringOutput `pulumi:"name"`
	// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	OnewaySynonymsControlIds pulumi.StringArrayOutput `pulumi:"onewaySynonymsControlIds"`
	// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	PriceRerankingLevel pulumi.StringOutput `pulumi:"priceRerankingLevel"`
	// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	RedirectControlIds pulumi.StringArrayOutput `pulumi:"redirectControlIds"`
	// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	ReplacementControlIds pulumi.StringArrayOutput `pulumi:"replacementControlIds"`
	// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
	SolutionTypes pulumi.StringArrayOutput `pulumi:"solutionTypes"`
	// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	TwowaySynonymsControlIds pulumi.StringArrayOutput `pulumi:"twowaySynonymsControlIds"`
}

// NewServingConfig registers a new resource with the given unique name, arguments, and options.
func NewServingConfig(ctx *pulumi.Context,
	name string, args *ServingConfigArgs, opts ...pulumi.ResourceOption) (*ServingConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ServingConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ServingConfigId'")
	}
	if args.SolutionTypes == nil {
		return nil, errors.New("invalid value for required argument 'SolutionTypes'")
	}
	var resource ServingConfig
	err := ctx.RegisterResource("google-native:retail/v2alpha:ServingConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServingConfig gets an existing ServingConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServingConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServingConfigState, opts ...pulumi.ResourceOption) (*ServingConfig, error) {
	var resource ServingConfig
	err := ctx.ReadResource("google-native:retail/v2alpha:ServingConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServingConfig resources.
type servingConfigState struct {
}

type ServingConfigState struct {
}

func (ServingConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*servingConfigState)(nil)).Elem()
}

type servingConfigArgs struct {
	// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	BoostControlIds []string `pulumi:"boostControlIds"`
	CatalogId       string   `pulumi:"catalogId"`
	// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `pulumi:"displayName"`
	// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	DiversityLevel *string `pulumi:"diversityLevel"`
	// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DoNotAssociateControlIds []string `pulumi:"doNotAssociateControlIds"`
	// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DynamicFacetSpec *GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpec `pulumi:"dynamicFacetSpec"`
	// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	EnableCategoryFilterLevel *string `pulumi:"enableCategoryFilterLevel"`
	// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FacetControlIds []string `pulumi:"facetControlIds"`
	// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FilterControlIds []string `pulumi:"filterControlIds"`
	// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	IgnoreControlIds []string `pulumi:"ignoreControlIds"`
	Location         *string  `pulumi:"location"`
	// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
	ModelId *string `pulumi:"modelId"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
	Name *string `pulumi:"name"`
	// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	OnewaySynonymsControlIds []string `pulumi:"onewaySynonymsControlIds"`
	// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	PriceRerankingLevel *string `pulumi:"priceRerankingLevel"`
	Project             *string `pulumi:"project"`
	// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	RedirectControlIds []string `pulumi:"redirectControlIds"`
	// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	ReplacementControlIds []string `pulumi:"replacementControlIds"`
	// Required. The ID to use for the ServingConfig, which will become the final component of the ServingConfig's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
	ServingConfigId string `pulumi:"servingConfigId"`
	// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
	SolutionTypes []ServingConfigSolutionTypesItem `pulumi:"solutionTypes"`
	// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	TwowaySynonymsControlIds []string `pulumi:"twowaySynonymsControlIds"`
}

// The set of arguments for constructing a ServingConfig resource.
type ServingConfigArgs struct {
	// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	BoostControlIds pulumi.StringArrayInput
	CatalogId       pulumi.StringInput
	// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName pulumi.StringInput
	// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	DiversityLevel pulumi.StringPtrInput
	// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DoNotAssociateControlIds pulumi.StringArrayInput
	// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DynamicFacetSpec GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecPtrInput
	// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	EnableCategoryFilterLevel pulumi.StringPtrInput
	// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FacetControlIds pulumi.StringArrayInput
	// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FilterControlIds pulumi.StringArrayInput
	// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	IgnoreControlIds pulumi.StringArrayInput
	Location         pulumi.StringPtrInput
	// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
	ModelId pulumi.StringPtrInput
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
	Name pulumi.StringPtrInput
	// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	OnewaySynonymsControlIds pulumi.StringArrayInput
	// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	PriceRerankingLevel pulumi.StringPtrInput
	Project             pulumi.StringPtrInput
	// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	RedirectControlIds pulumi.StringArrayInput
	// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	ReplacementControlIds pulumi.StringArrayInput
	// Required. The ID to use for the ServingConfig, which will become the final component of the ServingConfig's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
	ServingConfigId pulumi.StringInput
	// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
	SolutionTypes ServingConfigSolutionTypesItemArrayInput
	// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	TwowaySynonymsControlIds pulumi.StringArrayInput
}

func (ServingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servingConfigArgs)(nil)).Elem()
}

type ServingConfigInput interface {
	pulumi.Input

	ToServingConfigOutput() ServingConfigOutput
	ToServingConfigOutputWithContext(ctx context.Context) ServingConfigOutput
}

func (*ServingConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ServingConfig)(nil)).Elem()
}

func (i *ServingConfig) ToServingConfigOutput() ServingConfigOutput {
	return i.ToServingConfigOutputWithContext(context.Background())
}

func (i *ServingConfig) ToServingConfigOutputWithContext(ctx context.Context) ServingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServingConfigOutput)
}

type ServingConfigOutput struct{ *pulumi.OutputState }

func (ServingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServingConfig)(nil)).Elem()
}

func (o ServingConfigOutput) ToServingConfigOutput() ServingConfigOutput {
	return o
}

func (o ServingConfigOutput) ToServingConfigOutputWithContext(ctx context.Context) ServingConfigOutput {
	return o
}

// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) BoostControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.BoostControlIds }).(pulumi.StringArrayOutput)
}

// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
func (o ServingConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o ServingConfigOutput) DiversityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.DiversityLevel }).(pulumi.StringOutput)
}

// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) DoNotAssociateControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.DoNotAssociateControlIds }).(pulumi.StringArrayOutput)
}

// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) DynamicFacetSpec() GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput {
	return o.ApplyT(func(v *ServingConfig) GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput {
		return v.DynamicFacetSpec
	}).(GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput)
}

// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o ServingConfigOutput) EnableCategoryFilterLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.EnableCategoryFilterLevel }).(pulumi.StringOutput)
}

// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) FacetControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.FacetControlIds }).(pulumi.StringArrayOutput)
}

// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) FilterControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.FilterControlIds }).(pulumi.StringArrayOutput)
}

// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) IgnoreControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.IgnoreControlIds }).(pulumi.StringArrayOutput)
}

// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o ServingConfigOutput) ModelId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.ModelId }).(pulumi.StringOutput)
}

// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
func (o ServingConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) OnewaySynonymsControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.OnewaySynonymsControlIds }).(pulumi.StringArrayOutput)
}

// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o ServingConfigOutput) PriceRerankingLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringOutput { return v.PriceRerankingLevel }).(pulumi.StringOutput)
}

// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) RedirectControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.RedirectControlIds }).(pulumi.StringArrayOutput)
}

// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) ReplacementControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.ReplacementControlIds }).(pulumi.StringArrayOutput)
}

// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
func (o ServingConfigOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o ServingConfigOutput) TwowaySynonymsControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServingConfig) pulumi.StringArrayOutput { return v.TwowaySynonymsControlIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServingConfigInput)(nil)).Elem(), &ServingConfig{})
	pulumi.RegisterOutputType(ServingConfigOutput{})
}
