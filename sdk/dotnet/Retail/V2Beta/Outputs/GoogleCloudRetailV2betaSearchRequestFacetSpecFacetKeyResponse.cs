// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta.Outputs
{

    /// <summary>
    /// Specifies how a facet is computed.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2betaSearchRequestFacetSpecFacetKeyResponse
    {
        /// <summary>
        /// True to make facet keys case insensitive when getting faceting values with prefixes or contains; false otherwise.
        /// </summary>
        public readonly bool CaseInsensitive;
        /// <summary>
        /// Only get facet values that contains the given strings. For example, suppose "categories" has three values "Women &gt; Shoe", "Women &gt; Dress" and "Men &gt; Shoe". If set "contains" to "Shoe", the "categories" facet will give only "Women &gt; Shoe" and "Men &gt; Shoe". Only supported on textual fields. Maximum is 10.
        /// </summary>
        public readonly ImmutableArray<string> Contains;
        /// <summary>
        /// For all numerical facet keys that appear in the list of products from the catalog, the percentiles 0, 10, 30, 50, 70, 90 and 100 are computed from their distribution weekly. If the model assigns a high score to a numerical facet key and its intervals are not specified in the search request, these percentiles will become the bounds for its intervals and will be returned in the response. If the facet key intervals are specified in the request, then the specified intervals will be returned instead.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRetailV2betaIntervalResponse> Intervals;
        /// <summary>
        /// Supported textual and numerical facet keys in Product object, over which the facet values are computed. Facet key is case-sensitive. Allowed facet keys when FacetKey.query is not specified: * textual_field = * "brands" * "categories" * "genders" * "ageGroups" * "availability" * "colorFamilies" * "colors" * "sizes" * "materials" * "patterns" * "conditions" * "attributes.key" * "pickupInStore" * "shipToStore" * "sameDayDelivery" * "nextDayDelivery" * "customFulfillment1" * "customFulfillment2" * "customFulfillment3" * "customFulfillment4" * "customFulfillment5" * "inventory(place_id,attributes.key)" * numerical_field = * "price" * "discount" * "rating" * "ratingCount" * "attributes.key" * "inventory(place_id,price)" * "inventory(place_id,original_price)" * "inventory(place_id,attributes.key)"
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The order in which SearchResponse.Facet.values are returned. Allowed values are: * "count desc", which means order by SearchResponse.Facet.values.count descending. * "value desc", which means order by SearchResponse.Facet.values.value descending. Only applies to textual facets. If not set, textual values are sorted in [natural order](https://en.wikipedia.org/wiki/Natural_sort_order); numerical intervals are sorted in the order given by FacetSpec.FacetKey.intervals; FulfillmentInfo.place_ids are sorted in the order given by FacetSpec.FacetKey.restricted_values.
        /// </summary>
        public readonly string OrderBy;
        /// <summary>
        /// Only get facet values that start with the given string prefix. For example, suppose "categories" has three values "Women &gt; Shoe", "Women &gt; Dress" and "Men &gt; Shoe". If set "prefixes" to "Women", the "categories" facet will give only "Women &gt; Shoe" and "Women &gt; Dress". Only supported on textual fields. Maximum is 10.
        /// </summary>
        public readonly ImmutableArray<string> Prefixes;
        /// <summary>
        /// The query that is used to compute facet for the given facet key. When provided, it will override the default behavior of facet computation. The query syntax is the same as a filter expression. See SearchRequest.filter for detail syntax and limitations. Notice that there is no limitation on FacetKey.key when query is specified. In the response, SearchResponse.Facet.values.value will be always "1" and SearchResponse.Facet.values.count will be the number of results that match the query. For example, you can set a customized facet for "shipToStore", where FacetKey.key is "customizedShipToStore", and FacetKey.query is "availability: ANY(\"IN_STOCK\") AND shipToStore: ANY(\"123\")". Then the facet will count the products that are both in stock and ship to store "123".
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// Only get facet for the given restricted values. For example, when using "pickupInStore" as key and set restricted values to ["store123", "store456"], only facets for "store123" and "store456" are returned. Only supported on predefined textual fields, custom textual attributes and fulfillments. Maximum is 20. Must be set for the fulfillment facet keys: * pickupInStore * shipToStore * sameDayDelivery * nextDayDelivery * customFulfillment1 * customFulfillment2 * customFulfillment3 * customFulfillment4 * customFulfillment5
        /// </summary>
        public readonly ImmutableArray<string> RestrictedValues;
        /// <summary>
        /// Returns the min and max value for each numerical facet intervals. Ignored for textual facets.
        /// </summary>
        public readonly bool ReturnMinMax;

        [OutputConstructor]
        private GoogleCloudRetailV2betaSearchRequestFacetSpecFacetKeyResponse(
            bool caseInsensitive,

            ImmutableArray<string> contains,

            ImmutableArray<Outputs.GoogleCloudRetailV2betaIntervalResponse> intervals,

            string key,

            string orderBy,

            ImmutableArray<string> prefixes,

            string query,

            ImmutableArray<string> restrictedValues,

            bool returnMinMax)
        {
            CaseInsensitive = caseInsensitive;
            Contains = contains;
            Intervals = intervals;
            Key = key;
            OrderBy = orderBy;
            Prefixes = prefixes;
            Query = query;
            RestrictedValues = restrictedValues;
            ReturnMinMax = returnMinMax;
        }
    }
}
