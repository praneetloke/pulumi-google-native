# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetServingConfigResult',
    'AwaitableGetServingConfigResult',
    'get_serving_config',
    'get_serving_config_output',
]

@pulumi.output_type
class GetServingConfigResult:
    def __init__(__self__, boost_control_ids=None, display_name=None, diversity_level=None, diversity_type=None, do_not_associate_control_ids=None, dynamic_facet_spec=None, enable_category_filter_level=None, facet_control_ids=None, filter_control_ids=None, ignore_control_ids=None, model_id=None, name=None, oneway_synonyms_control_ids=None, personalization_spec=None, price_reranking_level=None, redirect_control_ids=None, replacement_control_ids=None, solution_types=None, twoway_synonyms_control_ids=None):
        if boost_control_ids and not isinstance(boost_control_ids, list):
            raise TypeError("Expected argument 'boost_control_ids' to be a list")
        pulumi.set(__self__, "boost_control_ids", boost_control_ids)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if diversity_level and not isinstance(diversity_level, str):
            raise TypeError("Expected argument 'diversity_level' to be a str")
        pulumi.set(__self__, "diversity_level", diversity_level)
        if diversity_type and not isinstance(diversity_type, str):
            raise TypeError("Expected argument 'diversity_type' to be a str")
        pulumi.set(__self__, "diversity_type", diversity_type)
        if do_not_associate_control_ids and not isinstance(do_not_associate_control_ids, list):
            raise TypeError("Expected argument 'do_not_associate_control_ids' to be a list")
        pulumi.set(__self__, "do_not_associate_control_ids", do_not_associate_control_ids)
        if dynamic_facet_spec and not isinstance(dynamic_facet_spec, dict):
            raise TypeError("Expected argument 'dynamic_facet_spec' to be a dict")
        pulumi.set(__self__, "dynamic_facet_spec", dynamic_facet_spec)
        if enable_category_filter_level and not isinstance(enable_category_filter_level, str):
            raise TypeError("Expected argument 'enable_category_filter_level' to be a str")
        pulumi.set(__self__, "enable_category_filter_level", enable_category_filter_level)
        if facet_control_ids and not isinstance(facet_control_ids, list):
            raise TypeError("Expected argument 'facet_control_ids' to be a list")
        pulumi.set(__self__, "facet_control_ids", facet_control_ids)
        if filter_control_ids and not isinstance(filter_control_ids, list):
            raise TypeError("Expected argument 'filter_control_ids' to be a list")
        pulumi.set(__self__, "filter_control_ids", filter_control_ids)
        if ignore_control_ids and not isinstance(ignore_control_ids, list):
            raise TypeError("Expected argument 'ignore_control_ids' to be a list")
        pulumi.set(__self__, "ignore_control_ids", ignore_control_ids)
        if model_id and not isinstance(model_id, str):
            raise TypeError("Expected argument 'model_id' to be a str")
        pulumi.set(__self__, "model_id", model_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if oneway_synonyms_control_ids and not isinstance(oneway_synonyms_control_ids, list):
            raise TypeError("Expected argument 'oneway_synonyms_control_ids' to be a list")
        pulumi.set(__self__, "oneway_synonyms_control_ids", oneway_synonyms_control_ids)
        if personalization_spec and not isinstance(personalization_spec, dict):
            raise TypeError("Expected argument 'personalization_spec' to be a dict")
        pulumi.set(__self__, "personalization_spec", personalization_spec)
        if price_reranking_level and not isinstance(price_reranking_level, str):
            raise TypeError("Expected argument 'price_reranking_level' to be a str")
        pulumi.set(__self__, "price_reranking_level", price_reranking_level)
        if redirect_control_ids and not isinstance(redirect_control_ids, list):
            raise TypeError("Expected argument 'redirect_control_ids' to be a list")
        pulumi.set(__self__, "redirect_control_ids", redirect_control_ids)
        if replacement_control_ids and not isinstance(replacement_control_ids, list):
            raise TypeError("Expected argument 'replacement_control_ids' to be a list")
        pulumi.set(__self__, "replacement_control_ids", replacement_control_ids)
        if solution_types and not isinstance(solution_types, list):
            raise TypeError("Expected argument 'solution_types' to be a list")
        pulumi.set(__self__, "solution_types", solution_types)
        if twoway_synonyms_control_ids and not isinstance(twoway_synonyms_control_ids, list):
            raise TypeError("Expected argument 'twoway_synonyms_control_ids' to be a list")
        pulumi.set(__self__, "twoway_synonyms_control_ids", twoway_synonyms_control_ids)

    @property
    @pulumi.getter(name="boostControlIds")
    def boost_control_ids(self) -> Sequence[str]:
        """
        Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "boost_control_ids")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="diversityLevel")
    def diversity_level(self) -> str:
        """
        How much diversity to use in recommendation model results e.g. `medium-diversity` or `high-diversity`. Currently supported values: * `no-diversity` * `low-diversity` * `medium-diversity` * `high-diversity` * `auto-diversity` If not specified, we choose default based on recommendation model type. Default value: `no-diversity`. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        """
        return pulumi.get(self, "diversity_level")

    @property
    @pulumi.getter(name="diversityType")
    def diversity_type(self) -> str:
        """
        What kind of diversity to use - data driven or rule based. If unset, the server behavior defaults to RULE_BASED_DIVERSITY.
        """
        return pulumi.get(self, "diversity_type")

    @property
    @pulumi.getter(name="doNotAssociateControlIds")
    def do_not_associate_control_ids(self) -> Sequence[str]:
        """
        Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "do_not_associate_control_ids")

    @property
    @pulumi.getter(name="dynamicFacetSpec")
    def dynamic_facet_spec(self) -> 'outputs.GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponse':
        """
        The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "dynamic_facet_spec")

    @property
    @pulumi.getter(name="enableCategoryFilterLevel")
    def enable_category_filter_level(self) -> str:
        """
        Whether to add additional category filters on the `similar-items` model. If not specified, we enable it by default. Allowed values are: * `no-category-match`: No additional filtering of original results from the model and the customer's filters. * `relaxed-category-match`: Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        """
        return pulumi.get(self, "enable_category_filter_level")

    @property
    @pulumi.getter(name="facetControlIds")
    def facet_control_ids(self) -> Sequence[str]:
        """
        Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "facet_control_ids")

    @property
    @pulumi.getter(name="filterControlIds")
    def filter_control_ids(self) -> Sequence[str]:
        """
        Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "filter_control_ids")

    @property
    @pulumi.getter(name="ignoreControlIds")
    def ignore_control_ids(self) -> Sequence[str]:
        """
        Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "ignore_control_ids")

    @property
    @pulumi.getter(name="modelId")
    def model_id(self) -> str:
        """
        The id of the model in the same Catalog to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
        """
        return pulumi.get(self, "model_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onewaySynonymsControlIds")
    def oneway_synonyms_control_ids(self) -> Sequence[str]:
        """
        Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "oneway_synonyms_control_ids")

    @property
    @pulumi.getter(name="personalizationSpec")
    def personalization_spec(self) -> 'outputs.GoogleCloudRetailV2alphaSearchRequestPersonalizationSpecResponse':
        """
        The specification for personalization spec. Can only be set if solution_types is SOLUTION_TYPE_SEARCH. Notice that if both ServingConfig.personalization_spec and SearchRequest.personalization_spec are set. SearchRequest.personalization_spec will override ServingConfig.personalization_spec.
        """
        return pulumi.get(self, "personalization_spec")

    @property
    @pulumi.getter(name="priceRerankingLevel")
    def price_reranking_level(self) -> str:
        """
        How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * `no-price-reranking` * `low-price-reranking` * `medium-price-reranking` * `high-price-reranking` If not specified, we choose default based on model type. Default value: `no-price-reranking`. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        """
        return pulumi.get(self, "price_reranking_level")

    @property
    @pulumi.getter(name="redirectControlIds")
    def redirect_control_ids(self) -> Sequence[str]:
        """
        Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "redirect_control_ids")

    @property
    @pulumi.getter(name="replacementControlIds")
    def replacement_control_ids(self) -> Sequence[str]:
        """
        Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "replacement_control_ids")

    @property
    @pulumi.getter(name="solutionTypes")
    def solution_types(self) -> Sequence[str]:
        """
        Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
        """
        return pulumi.get(self, "solution_types")

    @property
    @pulumi.getter(name="twowaySynonymsControlIds")
    def twoway_synonyms_control_ids(self) -> Sequence[str]:
        """
        Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "twoway_synonyms_control_ids")


class AwaitableGetServingConfigResult(GetServingConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServingConfigResult(
            boost_control_ids=self.boost_control_ids,
            display_name=self.display_name,
            diversity_level=self.diversity_level,
            diversity_type=self.diversity_type,
            do_not_associate_control_ids=self.do_not_associate_control_ids,
            dynamic_facet_spec=self.dynamic_facet_spec,
            enable_category_filter_level=self.enable_category_filter_level,
            facet_control_ids=self.facet_control_ids,
            filter_control_ids=self.filter_control_ids,
            ignore_control_ids=self.ignore_control_ids,
            model_id=self.model_id,
            name=self.name,
            oneway_synonyms_control_ids=self.oneway_synonyms_control_ids,
            personalization_spec=self.personalization_spec,
            price_reranking_level=self.price_reranking_level,
            redirect_control_ids=self.redirect_control_ids,
            replacement_control_ids=self.replacement_control_ids,
            solution_types=self.solution_types,
            twoway_synonyms_control_ids=self.twoway_synonyms_control_ids)


def get_serving_config(catalog_id: Optional[str] = None,
                       location: Optional[str] = None,
                       project: Optional[str] = None,
                       serving_config_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServingConfigResult:
    """
    Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['servingConfigId'] = serving_config_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:retail/v2alpha:getServingConfig', __args__, opts=opts, typ=GetServingConfigResult).value

    return AwaitableGetServingConfigResult(
        boost_control_ids=pulumi.get(__ret__, 'boost_control_ids'),
        display_name=pulumi.get(__ret__, 'display_name'),
        diversity_level=pulumi.get(__ret__, 'diversity_level'),
        diversity_type=pulumi.get(__ret__, 'diversity_type'),
        do_not_associate_control_ids=pulumi.get(__ret__, 'do_not_associate_control_ids'),
        dynamic_facet_spec=pulumi.get(__ret__, 'dynamic_facet_spec'),
        enable_category_filter_level=pulumi.get(__ret__, 'enable_category_filter_level'),
        facet_control_ids=pulumi.get(__ret__, 'facet_control_ids'),
        filter_control_ids=pulumi.get(__ret__, 'filter_control_ids'),
        ignore_control_ids=pulumi.get(__ret__, 'ignore_control_ids'),
        model_id=pulumi.get(__ret__, 'model_id'),
        name=pulumi.get(__ret__, 'name'),
        oneway_synonyms_control_ids=pulumi.get(__ret__, 'oneway_synonyms_control_ids'),
        personalization_spec=pulumi.get(__ret__, 'personalization_spec'),
        price_reranking_level=pulumi.get(__ret__, 'price_reranking_level'),
        redirect_control_ids=pulumi.get(__ret__, 'redirect_control_ids'),
        replacement_control_ids=pulumi.get(__ret__, 'replacement_control_ids'),
        solution_types=pulumi.get(__ret__, 'solution_types'),
        twoway_synonyms_control_ids=pulumi.get(__ret__, 'twoway_synonyms_control_ids'))


@_utilities.lift_output_func(get_serving_config)
def get_serving_config_output(catalog_id: Optional[pulumi.Input[str]] = None,
                              location: Optional[pulumi.Input[str]] = None,
                              project: Optional[pulumi.Input[Optional[str]]] = None,
                              serving_config_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServingConfigResult]:
    """
    Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
    """
    ...
