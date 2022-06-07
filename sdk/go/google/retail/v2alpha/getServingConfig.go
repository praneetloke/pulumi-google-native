// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
func LookupServingConfig(ctx *pulumi.Context, args *LookupServingConfigArgs, opts ...pulumi.InvokeOption) (*LookupServingConfigResult, error) {
	var rv LookupServingConfigResult
	err := ctx.Invoke("google-native:retail/v2alpha:getServingConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServingConfigArgs struct {
	CatalogId       string  `pulumi:"catalogId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
	ServingConfigId string  `pulumi:"servingConfigId"`
}

type LookupServingConfigResult struct {
	// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	BoostControlIds []string `pulumi:"boostControlIds"`
	// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `pulumi:"displayName"`
	// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	DiversityLevel string `pulumi:"diversityLevel"`
	// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DoNotAssociateControlIds []string `pulumi:"doNotAssociateControlIds"`
	// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	DynamicFacetSpec GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponse `pulumi:"dynamicFacetSpec"`
	// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	EnableCategoryFilterLevel string `pulumi:"enableCategoryFilterLevel"`
	// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FacetControlIds []string `pulumi:"facetControlIds"`
	// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	FilterControlIds []string `pulumi:"filterControlIds"`
	// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	IgnoreControlIds []string `pulumi:"ignoreControlIds"`
	// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
	ModelId string `pulumi:"modelId"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
	Name string `pulumi:"name"`
	// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	OnewaySynonymsControlIds []string `pulumi:"onewaySynonymsControlIds"`
	// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
	PriceRerankingLevel string `pulumi:"priceRerankingLevel"`
	// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	RedirectControlIds []string `pulumi:"redirectControlIds"`
	// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	ReplacementControlIds []string `pulumi:"replacementControlIds"`
	// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
	SolutionTypes []string `pulumi:"solutionTypes"`
	// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
	TwowaySynonymsControlIds []string `pulumi:"twowaySynonymsControlIds"`
}

func LookupServingConfigOutput(ctx *pulumi.Context, args LookupServingConfigOutputArgs, opts ...pulumi.InvokeOption) LookupServingConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServingConfigResult, error) {
			args := v.(LookupServingConfigArgs)
			r, err := LookupServingConfig(ctx, &args, opts...)
			var s LookupServingConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServingConfigResultOutput)
}

type LookupServingConfigOutputArgs struct {
	CatalogId       pulumi.StringInput    `pulumi:"catalogId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
	ServingConfigId pulumi.StringInput    `pulumi:"servingConfigId"`
}

func (LookupServingConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServingConfigArgs)(nil)).Elem()
}

type LookupServingConfigResultOutput struct{ *pulumi.OutputState }

func (LookupServingConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServingConfigResult)(nil)).Elem()
}

func (o LookupServingConfigResultOutput) ToLookupServingConfigResultOutput() LookupServingConfigResultOutput {
	return o
}

func (o LookupServingConfigResultOutput) ToLookupServingConfigResultOutputWithContext(ctx context.Context) LookupServingConfigResultOutput {
	return o
}

// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) BoostControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.BoostControlIds }).(pulumi.StringArrayOutput)
}

// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
func (o LookupServingConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o LookupServingConfigResultOutput) DiversityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.DiversityLevel }).(pulumi.StringOutput)
}

// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) DoNotAssociateControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.DoNotAssociateControlIds }).(pulumi.StringArrayOutput)
}

// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) DynamicFacetSpec() GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput {
	return o.ApplyT(func(v LookupServingConfigResult) GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponse {
		return v.DynamicFacetSpec
	}).(GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponseOutput)
}

// Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o LookupServingConfigResultOutput) EnableCategoryFilterLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.EnableCategoryFilterLevel }).(pulumi.StringOutput)
}

// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) FacetControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.FacetControlIds }).(pulumi.StringArrayOutput)
}

// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) FilterControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.FilterControlIds }).(pulumi.StringArrayOutput)
}

// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) IgnoreControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.IgnoreControlIds }).(pulumi.StringArrayOutput)
}

// The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o LookupServingConfigResultOutput) ModelId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.ModelId }).(pulumi.StringOutput)
}

// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
func (o LookupServingConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) OnewaySynonymsControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.OnewaySynonymsControlIds }).(pulumi.StringArrayOutput)
}

// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
func (o LookupServingConfigResultOutput) PriceRerankingLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServingConfigResult) string { return v.PriceRerankingLevel }).(pulumi.StringOutput)
}

// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) RedirectControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.RedirectControlIds }).(pulumi.StringArrayOutput)
}

// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) ReplacementControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.ReplacementControlIds }).(pulumi.StringArrayOutput)
}

// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
func (o LookupServingConfigResultOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
func (o LookupServingConfigResultOutput) TwowaySynonymsControlIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServingConfigResult) []string { return v.TwowaySynonymsControlIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServingConfigResultOutput{})
}
