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
    'GetProductResult',
    'AwaitableGetProductResult',
    'get_product',
    'get_product_output',
]

@pulumi.output_type
class GetProductResult:
    def __init__(__self__, attributes=None, audience=None, availability=None, available_quantity=None, available_time=None, brands=None, categories=None, collection_member_ids=None, color_info=None, conditions=None, description=None, expire_time=None, fulfillment_info=None, gtin=None, images=None, language_code=None, local_inventories=None, materials=None, name=None, patterns=None, price_info=None, primary_product_id=None, promotions=None, publish_time=None, rating=None, retrievable_fields=None, sizes=None, tags=None, title=None, ttl=None, type=None, uri=None, variants=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if audience and not isinstance(audience, dict):
            raise TypeError("Expected argument 'audience' to be a dict")
        pulumi.set(__self__, "audience", audience)
        if availability and not isinstance(availability, str):
            raise TypeError("Expected argument 'availability' to be a str")
        pulumi.set(__self__, "availability", availability)
        if available_quantity and not isinstance(available_quantity, int):
            raise TypeError("Expected argument 'available_quantity' to be a int")
        pulumi.set(__self__, "available_quantity", available_quantity)
        if available_time and not isinstance(available_time, str):
            raise TypeError("Expected argument 'available_time' to be a str")
        pulumi.set(__self__, "available_time", available_time)
        if brands and not isinstance(brands, list):
            raise TypeError("Expected argument 'brands' to be a list")
        pulumi.set(__self__, "brands", brands)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if collection_member_ids and not isinstance(collection_member_ids, list):
            raise TypeError("Expected argument 'collection_member_ids' to be a list")
        pulumi.set(__self__, "collection_member_ids", collection_member_ids)
        if color_info and not isinstance(color_info, dict):
            raise TypeError("Expected argument 'color_info' to be a dict")
        pulumi.set(__self__, "color_info", color_info)
        if conditions and not isinstance(conditions, list):
            raise TypeError("Expected argument 'conditions' to be a list")
        pulumi.set(__self__, "conditions", conditions)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if fulfillment_info and not isinstance(fulfillment_info, list):
            raise TypeError("Expected argument 'fulfillment_info' to be a list")
        pulumi.set(__self__, "fulfillment_info", fulfillment_info)
        if gtin and not isinstance(gtin, str):
            raise TypeError("Expected argument 'gtin' to be a str")
        pulumi.set(__self__, "gtin", gtin)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if language_code and not isinstance(language_code, str):
            raise TypeError("Expected argument 'language_code' to be a str")
        pulumi.set(__self__, "language_code", language_code)
        if local_inventories and not isinstance(local_inventories, list):
            raise TypeError("Expected argument 'local_inventories' to be a list")
        pulumi.set(__self__, "local_inventories", local_inventories)
        if materials and not isinstance(materials, list):
            raise TypeError("Expected argument 'materials' to be a list")
        pulumi.set(__self__, "materials", materials)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if patterns and not isinstance(patterns, list):
            raise TypeError("Expected argument 'patterns' to be a list")
        pulumi.set(__self__, "patterns", patterns)
        if price_info and not isinstance(price_info, dict):
            raise TypeError("Expected argument 'price_info' to be a dict")
        pulumi.set(__self__, "price_info", price_info)
        if primary_product_id and not isinstance(primary_product_id, str):
            raise TypeError("Expected argument 'primary_product_id' to be a str")
        pulumi.set(__self__, "primary_product_id", primary_product_id)
        if promotions and not isinstance(promotions, list):
            raise TypeError("Expected argument 'promotions' to be a list")
        pulumi.set(__self__, "promotions", promotions)
        if publish_time and not isinstance(publish_time, str):
            raise TypeError("Expected argument 'publish_time' to be a str")
        pulumi.set(__self__, "publish_time", publish_time)
        if rating and not isinstance(rating, dict):
            raise TypeError("Expected argument 'rating' to be a dict")
        pulumi.set(__self__, "rating", rating)
        if retrievable_fields and not isinstance(retrievable_fields, str):
            raise TypeError("Expected argument 'retrievable_fields' to be a str")
        if retrievable_fields is not None:
            warnings.warn("""Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form \"attributes.key\" where \"key\" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info The maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse can increase response payload size and serving latency. This field is deprecated. Use the retrievable site-wide control instead.""", DeprecationWarning)
            pulumi.log.warn("""retrievable_fields is deprecated: Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form \"attributes.key\" where \"key\" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info The maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse can increase response payload size and serving latency. This field is deprecated. Use the retrievable site-wide control instead.""")

        pulumi.set(__self__, "retrievable_fields", retrievable_fields)
        if sizes and not isinstance(sizes, list):
            raise TypeError("Expected argument 'sizes' to be a list")
        pulumi.set(__self__, "sizes", sizes)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if ttl and not isinstance(ttl, str):
            raise TypeError("Expected argument 'ttl' to be a str")
        pulumi.set(__self__, "ttl", ttl)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)
        if variants and not isinstance(variants, list):
            raise TypeError("Expected argument 'variants' to be a list")
        pulumi.set(__self__, "variants", variants)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, str]:
        """
        Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 200. * The key must be a UTF-8 encoded string with a length limit of 128 characters. * For indexable attribute, the key must match the pattern: `a-zA-Z0-9*`. For example, `key0LikeThis` or `KEY_1_LIKE_THIS`. * For text attributes, at most 400 values are allowed. Empty values are not allowed. Each value must be a non-empty UTF-8 encoded string with a length limit of 256 characters. * For number attributes, at most 400 values are allowed.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def audience(self) -> 'outputs.GoogleCloudRetailV2AudienceResponse':
        """
        The target group associated with a given audience (e.g. male, veterans, car owners, musicians, etc.) of the product.
        """
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter
    def availability(self) -> str:
        """
        The online availability of the Product. Default to Availability.IN_STOCK. Corresponding properties: Google Merchant Center property [availability](https://support.google.com/merchants/answer/6324448). Schema.org property [Offer.availability](https://schema.org/availability).
        """
        return pulumi.get(self, "availability")

    @property
    @pulumi.getter(name="availableQuantity")
    def available_quantity(self) -> int:
        """
        The available quantity of the item.
        """
        return pulumi.get(self, "available_quantity")

    @property
    @pulumi.getter(name="availableTime")
    def available_time(self) -> str:
        """
        The timestamp when this Product becomes available for SearchService.Search. Note that this is only applicable to Type.PRIMARY and Type.COLLECTION, and ignored for Type.VARIANT.
        """
        return pulumi.get(self, "available_time")

    @property
    @pulumi.getter
    def brands(self) -> Sequence[str]:
        """
        The brands of the product. A maximum of 30 brands are allowed. Each brand must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [brand](https://support.google.com/merchants/answer/6324351). Schema.org property [Product.brand](https://schema.org/brand).
        """
        return pulumi.get(self, "brands")

    @property
    @pulumi.getter
    def categories(self) -> Sequence[str]:
        """
        Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter(name="collectionMemberIds")
    def collection_member_ids(self) -> Sequence[str]:
        """
        The id of the collection members when type is Type.COLLECTION. Non-existent product ids are allowed. The type of the members must be either Type.PRIMARY or Type.VARIANT otherwise an INVALID_ARGUMENT error is thrown. Should not set it for other types. A maximum of 1000 values are allowed. Otherwise, an INVALID_ARGUMENT error is return.
        """
        return pulumi.get(self, "collection_member_ids")

    @property
    @pulumi.getter(name="colorInfo")
    def color_info(self) -> 'outputs.GoogleCloudRetailV2ColorInfoResponse':
        """
        The color of the product. Corresponding properties: Google Merchant Center property [color](https://support.google.com/merchants/answer/6324487). Schema.org property [Product.color](https://schema.org/color).
        """
        return pulumi.get(self, "color_info")

    @property
    @pulumi.getter
    def conditions(self) -> Sequence[str]:
        """
        The condition of the product. Strongly encouraged to use the standard values: "new", "refurbished", "used". A maximum of 1 value is allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [condition](https://support.google.com/merchants/answer/6324469). Schema.org property [Offer.itemCondition](https://schema.org/itemCondition).
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). Schema.org property [Product.description](https://schema.org/description).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The timestamp when this product becomes unavailable for SearchService.Search. Note that this is only applicable to Type.PRIMARY and Type.COLLECTION, and ignored for Type.VARIANT. In general, we suggest the users to delete the stale products explicitly, instead of using this field to determine staleness. If it is set, the Product is not available for SearchService.Search after expire_time. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts. expire_time must be later than available_time and publish_time, otherwise an INVALID_ARGUMENT error is thrown. Corresponding properties: Google Merchant Center property [expiration_date](https://support.google.com/merchants/answer/6324499).
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="fulfillmentInfo")
    def fulfillment_info(self) -> Sequence['outputs.GoogleCloudRetailV2FulfillmentInfoResponse']:
        """
        Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods. All the elements must have distinct FulfillmentInfo.type. Otherwise, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "fulfillment_info")

    @property
    @pulumi.getter
    def gtin(self) -> str:
        """
        The Global Trade Item Number (GTIN) of the product. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. This field must be a Unigram. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [gtin](https://support.google.com/merchants/answer/6324461). Schema.org property [Product.isbn](https://schema.org/isbn), [Product.gtin8](https://schema.org/gtin8), [Product.gtin12](https://schema.org/gtin12), [Product.gtin13](https://schema.org/gtin13), or [Product.gtin14](https://schema.org/gtin14). If the value is not a valid GTIN, an INVALID_ARGUMENT error is returned.
        """
        return pulumi.get(self, "gtin")

    @property
    @pulumi.getter
    def images(self) -> Sequence['outputs.GoogleCloudRetailV2ImageResponse']:
        """
        Product images for the product. We highly recommend putting the main image first. A maximum of 300 images are allowed. Corresponding properties: Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> str:
        """
        Language of the title/description and other string attributes. Use language tags defined by [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). For product prediction, this field is ignored and the model automatically detects the text language. The Product can include text in different languages, but duplicating Products to provide text in multiple languages can result in degraded model performance. For product search this field is in use. It defaults to "en-US" if unset.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter(name="localInventories")
    def local_inventories(self) -> Sequence['outputs.GoogleCloudRetailV2LocalInventoryResponse']:
        """
        A list of local inventories specific to different places. This field can be managed by ProductService.AddLocalInventories and ProductService.RemoveLocalInventories APIs if fine-grained, high-volume updates are necessary.
        """
        return pulumi.get(self, "local_inventories")

    @property
    @pulumi.getter
    def materials(self) -> Sequence[str]:
        """
        The material of the product. For example, "leather", "wooden". A maximum of 20 values are allowed. Each value must be a UTF-8 encoded string with a length limit of 200 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [material](https://support.google.com/merchants/answer/6324410). Schema.org property [Product.material](https://schema.org/material).
        """
        return pulumi.get(self, "materials")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def patterns(self) -> Sequence[str]:
        """
        The pattern or graphic print of the product. For example, "striped", "polka dot", "paisley". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [pattern](https://support.google.com/merchants/answer/6324483). Schema.org property [Product.pattern](https://schema.org/pattern).
        """
        return pulumi.get(self, "patterns")

    @property
    @pulumi.getter(name="priceInfo")
    def price_info(self) -> 'outputs.GoogleCloudRetailV2PriceInfoResponse':
        """
        Product price and cost information. Corresponding properties: Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
        """
        return pulumi.get(self, "price_info")

    @property
    @pulumi.getter(name="primaryProductId")
    def primary_product_id(self) -> str:
        """
        Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID).
        """
        return pulumi.get(self, "primary_product_id")

    @property
    @pulumi.getter
    def promotions(self) -> Sequence['outputs.GoogleCloudRetailV2PromotionResponse']:
        """
        The promotions applied to the product. A maximum of 10 values are allowed per Product. Only Promotion.promotion_id will be used, other fields will be ignored if set.
        """
        return pulumi.get(self, "promotions")

    @property
    @pulumi.getter(name="publishTime")
    def publish_time(self) -> str:
        """
        The timestamp when the product is published by the retailer for the first time, which indicates the freshness of the products. Note that this field is different from available_time, given it purely describes product freshness regardless of when it is available on search and recommendation.
        """
        return pulumi.get(self, "publish_time")

    @property
    @pulumi.getter
    def rating(self) -> 'outputs.GoogleCloudRetailV2RatingResponse':
        """
        The rating of this product.
        """
        return pulumi.get(self, "rating")

    @property
    @pulumi.getter(name="retrievableFields")
    def retrievable_fields(self) -> str:
        """
        Indicates which fields in the Products are returned in SearchResponse. Supported fields for all types: * audience * availability * brands * color_info * conditions * gtin * materials * name * patterns * price_info * rating * sizes * title * uri Supported fields only for Type.PRIMARY and Type.COLLECTION: * categories * description * images Supported fields only for Type.VARIANT: * Only the first image in images To mark attributes as retrievable, include paths of the form "attributes.key" where "key" is the key of a custom attribute, as specified in attributes. For Type.PRIMARY and Type.COLLECTION, the following fields are always returned in SearchResponse by default: * name For Type.VARIANT, the following fields are always returned in by default: * name * color_info The maximum number of paths is 30. Otherwise, an INVALID_ARGUMENT error is returned. Note: Returning more fields in SearchResponse can increase response payload size and serving latency. This field is deprecated. Use the retrievable site-wide control instead.
        """
        return pulumi.get(self, "retrievable_fields")

    @property
    @pulumi.getter
    def sizes(self) -> Sequence[str]:
        """
        The size of the product. To represent different size systems or size types, consider using this format: [[[size_system:]size_type:]size_value]. For example, in "US:MENS:M", "US" represents size system; "MENS" represents size type; "M" represents size value. In "GIRLS:27", size system is empty; "GIRLS" represents size type; "27" represents size value. In "32 inches", both size system and size type are empty, while size value is "32 inches". A maximum of 20 values are allowed per Product. Each value must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [size](https://support.google.com/merchants/answer/6324492), [size_type](https://support.google.com/merchants/answer/6324497), and [size_system](https://support.google.com/merchants/answer/6324502). Schema.org property [Product.size](https://schema.org/size).
        """
        return pulumi.get(self, "sizes")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Corresponding properties: Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def ttl(self) -> str:
        """
        Input only. The TTL (time to live) of the product. Note that this is only applicable to Type.PRIMARY and Type.COLLECTION, and ignored for Type.VARIANT. In general, we suggest the users to delete the stale products explicitly, instead of using this field to determine staleness. If it is set, it must be a non-negative value, and expire_time is set as current timestamp plus ttl. The derived expire_time is returned in the output and ttl is left blank when retrieving the Product. If it is set, the product is not available for SearchService.Search after current timestamp plus ttl. However, the product can still be retrieved by ProductService.GetProduct and ProductService.ListProducts.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Immutable. The type of the product. Default to Catalog.product_level_config.ingestion_product_type if unset.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Corresponding properties: Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def variants(self) -> Sequence['outputs.GoogleCloudRetailV2ProductResponse']:
        """
        Product variants grouped together on primary product which share similar product attributes. It's automatically grouped by primary_product_id for all the product variants. Only populated for Type.PRIMARY Products. Note: This field is OUTPUT_ONLY for ProductService.GetProduct. Do not set this field in API requests.
        """
        return pulumi.get(self, "variants")


class AwaitableGetProductResult(GetProductResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductResult(
            attributes=self.attributes,
            audience=self.audience,
            availability=self.availability,
            available_quantity=self.available_quantity,
            available_time=self.available_time,
            brands=self.brands,
            categories=self.categories,
            collection_member_ids=self.collection_member_ids,
            color_info=self.color_info,
            conditions=self.conditions,
            description=self.description,
            expire_time=self.expire_time,
            fulfillment_info=self.fulfillment_info,
            gtin=self.gtin,
            images=self.images,
            language_code=self.language_code,
            local_inventories=self.local_inventories,
            materials=self.materials,
            name=self.name,
            patterns=self.patterns,
            price_info=self.price_info,
            primary_product_id=self.primary_product_id,
            promotions=self.promotions,
            publish_time=self.publish_time,
            rating=self.rating,
            retrievable_fields=self.retrievable_fields,
            sizes=self.sizes,
            tags=self.tags,
            title=self.title,
            ttl=self.ttl,
            type=self.type,
            uri=self.uri,
            variants=self.variants)


def get_product(branch_id: Optional[str] = None,
                catalog_id: Optional[str] = None,
                location: Optional[str] = None,
                product_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductResult:
    """
    Gets a Product.
    """
    __args__ = dict()
    __args__['branchId'] = branch_id
    __args__['catalogId'] = catalog_id
    __args__['location'] = location
    __args__['productId'] = product_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:retail/v2:getProduct', __args__, opts=opts, typ=GetProductResult).value

    return AwaitableGetProductResult(
        attributes=__ret__.attributes,
        audience=__ret__.audience,
        availability=__ret__.availability,
        available_quantity=__ret__.available_quantity,
        available_time=__ret__.available_time,
        brands=__ret__.brands,
        categories=__ret__.categories,
        collection_member_ids=__ret__.collection_member_ids,
        color_info=__ret__.color_info,
        conditions=__ret__.conditions,
        description=__ret__.description,
        expire_time=__ret__.expire_time,
        fulfillment_info=__ret__.fulfillment_info,
        gtin=__ret__.gtin,
        images=__ret__.images,
        language_code=__ret__.language_code,
        local_inventories=__ret__.local_inventories,
        materials=__ret__.materials,
        name=__ret__.name,
        patterns=__ret__.patterns,
        price_info=__ret__.price_info,
        primary_product_id=__ret__.primary_product_id,
        promotions=__ret__.promotions,
        publish_time=__ret__.publish_time,
        rating=__ret__.rating,
        retrievable_fields=__ret__.retrievable_fields,
        sizes=__ret__.sizes,
        tags=__ret__.tags,
        title=__ret__.title,
        ttl=__ret__.ttl,
        type=__ret__.type,
        uri=__ret__.uri,
        variants=__ret__.variants)


@_utilities.lift_output_func(get_product)
def get_product_output(branch_id: Optional[pulumi.Input[str]] = None,
                       catalog_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       product_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductResult]:
    """
    Gets a Product.
    """
    ...
