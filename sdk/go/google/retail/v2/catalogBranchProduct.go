// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Product.
type CatalogBranchProduct struct {
	pulumi.CustomResourceState

	// Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability pulumi.StringOutput `pulumi:"availability"`
	// The available quantity of the item.
	AvailableQuantity pulumi.IntOutput `pulumi:"availableQuantity"`
	// The timestamp when this Product becomes available for recommendation.
	AvailableTime pulumi.StringOutput `pulumi:"availableTime"`
	// Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
	Categories pulumi.StringArrayOutput `pulumi:"categories"`
	// Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
	Description pulumi.StringOutput `pulumi:"description"`
	// Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Images GoogleCloudRetailV2ImageResponseArrayOutput `pulumi:"images"`
	// Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
	Name pulumi.StringOutput `pulumi:"name"`
	// Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo GoogleCloudRetailV2PriceInfoResponseOutput `pulumi:"priceInfo"`
	// Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId pulumi.StringOutput `pulumi:"primaryProductId"`
	// Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
	Title pulumi.StringOutput `pulumi:"title"`
	// Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
	Type pulumi.StringOutput `pulumi:"type"`
	// Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewCatalogBranchProduct registers a new resource with the given unique name, arguments, and options.
func NewCatalogBranchProduct(ctx *pulumi.Context,
	name string, args *CatalogBranchProductArgs, opts ...pulumi.ResourceOption) (*CatalogBranchProduct, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BranchId == nil {
		return nil, errors.New("invalid value for required argument 'BranchId'")
	}
	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource CatalogBranchProduct
	err := ctx.RegisterResource("google-native:retail/v2:CatalogBranchProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalogBranchProduct gets an existing CatalogBranchProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogBranchProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogBranchProductState, opts ...pulumi.ResourceOption) (*CatalogBranchProduct, error) {
	var resource CatalogBranchProduct
	err := ctx.ReadResource("google-native:retail/v2:CatalogBranchProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CatalogBranchProduct resources.
type catalogBranchProductState struct {
	// Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
	Attributes map[string]string `pulumi:"attributes"`
	// The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability *string `pulumi:"availability"`
	// The available quantity of the item.
	AvailableQuantity *int `pulumi:"availableQuantity"`
	// The timestamp when this Product becomes available for recommendation.
	AvailableTime *string `pulumi:"availableTime"`
	// Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
	Categories []string `pulumi:"categories"`
	// Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
	Description *string `pulumi:"description"`
	// Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Images []GoogleCloudRetailV2ImageResponse `pulumi:"images"`
	// Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
	Name *string `pulumi:"name"`
	// Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo *GoogleCloudRetailV2PriceInfoResponse `pulumi:"priceInfo"`
	// Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId *string `pulumi:"primaryProductId"`
	// Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags []string `pulumi:"tags"`
	// Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
	Title *string `pulumi:"title"`
	// Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
	Type *string `pulumi:"type"`
	// Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
	Uri *string `pulumi:"uri"`
}

type CatalogBranchProductState struct {
	// Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
	Attributes pulumi.StringMapInput
	// The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability pulumi.StringPtrInput
	// The available quantity of the item.
	AvailableQuantity pulumi.IntPtrInput
	// The timestamp when this Product becomes available for recommendation.
	AvailableTime pulumi.StringPtrInput
	// Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
	Categories pulumi.StringArrayInput
	// Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
	Description pulumi.StringPtrInput
	// Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Images GoogleCloudRetailV2ImageResponseArrayInput
	// Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
	Name pulumi.StringPtrInput
	// Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo GoogleCloudRetailV2PriceInfoResponsePtrInput
	// Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId pulumi.StringPtrInput
	// Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags pulumi.StringArrayInput
	// Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
	Title pulumi.StringPtrInput
	// Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
	Type pulumi.StringPtrInput
	// Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
	Uri pulumi.StringPtrInput
}

func (CatalogBranchProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogBranchProductState)(nil)).Elem()
}

type catalogBranchProductArgs struct {
	// Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
	Attributes map[string]string `pulumi:"attributes"`
	// The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability *string `pulumi:"availability"`
	// The available quantity of the item.
	AvailableQuantity *int `pulumi:"availableQuantity"`
	// The timestamp when this Product becomes available for recommendation.
	AvailableTime *string `pulumi:"availableTime"`
	BranchId      string  `pulumi:"branchId"`
	CatalogId     string  `pulumi:"catalogId"`
	// Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
	Categories []string `pulumi:"categories"`
	// Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
	Description *string `pulumi:"description"`
	// Immutable. Product identifier, which is the final component of name. For example, this field is "id_1", if name is `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [id](https://support.google.com/merchants/answer/6324405). Schema.org Property [Product.sku](https://schema.org/sku).
	Id *string `pulumi:"id"`
	// Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Images   []GoogleCloudRetailV2Image `pulumi:"images"`
	Location string                     `pulumi:"location"`
	// Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
	Name *string `pulumi:"name"`
	// Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo *GoogleCloudRetailV2PriceInfo `pulumi:"priceInfo"`
	// Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId *string `pulumi:"primaryProductId"`
	ProductId        string  `pulumi:"productId"`
	Project          string  `pulumi:"project"`
	// Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags []string `pulumi:"tags"`
	// Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
	Title *string `pulumi:"title"`
	// Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
	Type *string `pulumi:"type"`
	// Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a CatalogBranchProduct resource.
type CatalogBranchProductArgs struct {
	// Highly encouraged. Extra product attributes to be included. For example, for products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the attributes here. Features that can take on one of a limited number of possible values. Two types of features can be set are: Textual features. some examples would be the brand/maker of a product, or country of a customer. Numerical features. Some examples would be the height/weight of a product, or age of a customer. For example: `{ "vendor": {"text": ["vendor123", "vendor456"]}, "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]} }`. This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT error is returned: * Max entries count: 150 by default; 100 for Type.VARIANT. * The key must be a UTF-8 encoded string with a length limit of 128 characters.
	Attributes pulumi.StringMapInput
	// The online availability of the Product. Default to Availability.IN_STOCK. Google Merchant Center Property [availability](https://support.google.com/merchants/answer/6324448). Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability pulumi.StringPtrInput
	// The available quantity of the item.
	AvailableQuantity pulumi.IntPtrInput
	// The timestamp when this Product becomes available for recommendation.
	AvailableTime pulumi.StringPtrInput
	BranchId      pulumi.StringInput
	CatalogId     pulumi.StringInput
	// Product categories. This field is repeated for supporting one product belonging to several parallel categories. Strongly recommended using the full path for better search / recommendation quality. To represent full path of category, use '>' sign to separate different hierarchies. If '>' is part of the category name, please replace it with other character(s). For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categories": [ "Shoes & Accessories > Shoes", "Sports & Fitness > Athletic Clothing > Shoes" ] Must be set for Type.PRIMARY Product otherwise an INVALID_ARGUMENT error is returned. At most 250 values are allowed per Product. Empty values are not allowed. Each value must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property google_product_category. Schema.org property [Product.category] (https://schema.org/category). [mc_google_product_category]: https://support.google.com/merchants/answer/6324436
	Categories pulumi.StringArrayInput
	// Product description. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [description](https://support.google.com/merchants/answer/6324468). schema.org property [Product.description](https://schema.org/description).
	Description pulumi.StringPtrInput
	// Immutable. Product identifier, which is the final component of name. For example, this field is "id_1", if name is `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [id](https://support.google.com/merchants/answer/6324405). Schema.org Property [Product.sku](https://schema.org/sku).
	Id pulumi.StringPtrInput
	// Product images for the product.Highly recommended to put the main image to the first. A maximum of 300 images are allowed. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Images   GoogleCloudRetailV2ImageArrayInput
	Location pulumi.StringInput
	// Immutable. Full resource name of the product, such as `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`. The branch ID must be "default_branch".
	Name pulumi.StringPtrInput
	// Product price and cost information. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo GoogleCloudRetailV2PriceInfoPtrInput
	// Variant group identifier. Must be an id, with the same parent branch with this product. Otherwise, an error is thrown. For Type.PRIMARY Products, this field can only be empty or set to the same value as id. For VARIANT Products, this field cannot be empty. A maximum of 2,000 products are allowed to share the same Type.PRIMARY Product. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center Property [item_group_id](https://support.google.com/merchants/answer/6324507). Schema.org Property [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID). This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId pulumi.StringPtrInput
	ProductId        pulumi.StringInput
	Project          pulumi.StringInput
	// Custom tags associated with the product. At most 250 values are allowed per Product. This value must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. This tag can be used for filtering recommendation results by passing the tag as part of the PredictRequest.filter. Google Merchant Center property [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags pulumi.StringArrayInput
	// Required. Product title. This field must be a UTF-8 encoded string with a length limit of 1,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [title](https://support.google.com/merchants/answer/6324415). Schema.org property [Product.name](https://schema.org/name).
	Title pulumi.StringPtrInput
	// Immutable. The type of the product. This field is output-only. Default to Catalog.product_level_config.ingestion_product_type if unset.
	Type pulumi.StringPtrInput
	// Canonical URL directly linking to the product detail page. It is strongly recommended to provide a valid uri for the product, otherwise the service performance could be significantly degraded. This field must be a UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [link](https://support.google.com/merchants/answer/6324416). Schema.org property [Offer.url](https://schema.org/url).
	Uri pulumi.StringPtrInput
}

func (CatalogBranchProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogBranchProductArgs)(nil)).Elem()
}

type CatalogBranchProductInput interface {
	pulumi.Input

	ToCatalogBranchProductOutput() CatalogBranchProductOutput
	ToCatalogBranchProductOutputWithContext(ctx context.Context) CatalogBranchProductOutput
}

func (*CatalogBranchProduct) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogBranchProduct)(nil))
}

func (i *CatalogBranchProduct) ToCatalogBranchProductOutput() CatalogBranchProductOutput {
	return i.ToCatalogBranchProductOutputWithContext(context.Background())
}

func (i *CatalogBranchProduct) ToCatalogBranchProductOutputWithContext(ctx context.Context) CatalogBranchProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogBranchProductOutput)
}

type CatalogBranchProductOutput struct {
	*pulumi.OutputState
}

func (CatalogBranchProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogBranchProduct)(nil))
}

func (o CatalogBranchProductOutput) ToCatalogBranchProductOutput() CatalogBranchProductOutput {
	return o
}

func (o CatalogBranchProductOutput) ToCatalogBranchProductOutputWithContext(ctx context.Context) CatalogBranchProductOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CatalogBranchProductOutput{})
}
