// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Attribute definition.
func LookupAttributeDefinition(ctx *pulumi.Context, args *LookupAttributeDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupAttributeDefinitionResult, error) {
	var rv LookupAttributeDefinitionResult
	err := ctx.Invoke("google-native:healthcare/v1:getAttributeDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttributeDefinitionArgs struct {
	AttributeDefinitionId string  `pulumi:"attributeDefinitionId"`
	ConsentStoreId        string  `pulumi:"consentStoreId"`
	DatasetId             string  `pulumi:"datasetId"`
	Location              string  `pulumi:"location"`
	Project               *string `pulumi:"project"`
}

type LookupAttributeDefinitionResult struct {
	// Possible values for the attribute. The number of allowed values must not exceed 100. An empty list is invalid. The list can only be expanded after creation.
	AllowedValues []string `pulumi:"allowedValues"`
	// The category of the attribute. The value of this field cannot be changed after creation.
	Category string `pulumi:"category"`
	// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
	ConsentDefaultValues []string `pulumi:"consentDefaultValues"`
	// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
	DataMappingDefaultValue string `pulumi:"dataMappingDefaultValue"`
	// Optional. A description of the attribute.
	Description string `pulumi:"description"`
	// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
	Name string `pulumi:"name"`
}