// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new FHIR store within the parent dataset.
// Auto-naming is currently not supported for this resource.
type FhirStore struct {
	pulumi.CustomResourceState

	// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
	ComplexDataTypeReferenceParsing pulumi.StringOutput `pulumi:"complexDataTypeReferenceParsing"`
	DatasetId                       pulumi.StringOutput `pulumi:"datasetId"`
	// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
	DefaultSearchHandlingStrict pulumi.BoolOutput `pulumi:"defaultSearchHandlingStrict"`
	// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
	DisableReferentialIntegrity pulumi.BoolOutput `pulumi:"disableReferentialIntegrity"`
	// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
	DisableResourceVersioning pulumi.BoolOutput `pulumi:"disableResourceVersioning"`
	// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
	EnableUpdateCreate pulumi.BoolOutput `pulumi:"enableUpdateCreate"`
	// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	FhirStoreId pulumi.StringPtrOutput `pulumi:"fhirStoreId"`
	// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	//
	// Deprecated: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	NotificationConfig NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	// Specifies where and whether to send notifications upon changes to a Fhir store.
	NotificationConfigs FhirNotificationConfigResponseArrayOutput `pulumi:"notificationConfigs"`
	Project             pulumi.StringOutput                       `pulumi:"project"`
	// Configuration for how FHIR resources can be searched.
	SearchConfig SearchConfigResponseOutput `pulumi:"searchConfig"`
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
	StreamConfigs StreamConfigResponseArrayOutput `pulumi:"streamConfigs"`
	// Configuration for how to validate incoming FHIR resources against configured profiles.
	ValidationConfig ValidationConfigResponseOutput `pulumi:"validationConfig"`
	// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewFhirStore registers a new resource with the given unique name, arguments, and options.
func NewFhirStore(ctx *pulumi.Context,
	name string, args *FhirStoreArgs, opts ...pulumi.ResourceOption) (*FhirStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"datasetId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource FhirStore
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:FhirStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStore gets an existing FhirStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreState, opts ...pulumi.ResourceOption) (*FhirStore, error) {
	var resource FhirStore
	err := ctx.ReadResource("google-native:healthcare/v1beta1:FhirStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStore resources.
type fhirStoreState struct {
}

type FhirStoreState struct {
}

func (FhirStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreState)(nil)).Elem()
}

type fhirStoreArgs struct {
	// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
	ComplexDataTypeReferenceParsing *FhirStoreComplexDataTypeReferenceParsing `pulumi:"complexDataTypeReferenceParsing"`
	DatasetId                       string                                    `pulumi:"datasetId"`
	// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
	DefaultSearchHandlingStrict *bool `pulumi:"defaultSearchHandlingStrict"`
	// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
	DisableReferentialIntegrity *bool `pulumi:"disableReferentialIntegrity"`
	// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
	DisableResourceVersioning *bool `pulumi:"disableResourceVersioning"`
	// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
	EnableUpdateCreate *bool `pulumi:"enableUpdateCreate"`
	// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	FhirStoreId *string `pulumi:"fhirStoreId"`
	// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	//
	// Deprecated: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	NotificationConfig *NotificationConfig `pulumi:"notificationConfig"`
	// Specifies where and whether to send notifications upon changes to a Fhir store.
	NotificationConfigs []FhirNotificationConfig `pulumi:"notificationConfigs"`
	Project             *string                  `pulumi:"project"`
	// Configuration for how FHIR resources can be searched.
	SearchConfig *SearchConfig `pulumi:"searchConfig"`
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
	StreamConfigs []StreamConfig `pulumi:"streamConfigs"`
	// Configuration for how to validate incoming FHIR resources against configured profiles.
	ValidationConfig *ValidationConfig `pulumi:"validationConfig"`
	// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
	Version *FhirStoreVersion `pulumi:"version"`
}

// The set of arguments for constructing a FhirStore resource.
type FhirStoreArgs struct {
	// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
	ComplexDataTypeReferenceParsing FhirStoreComplexDataTypeReferenceParsingPtrInput
	DatasetId                       pulumi.StringInput
	// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
	DefaultSearchHandlingStrict pulumi.BoolPtrInput
	// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
	DisableReferentialIntegrity pulumi.BoolPtrInput
	// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
	DisableResourceVersioning pulumi.BoolPtrInput
	// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
	EnableUpdateCreate pulumi.BoolPtrInput
	// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	FhirStoreId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	//
	// Deprecated: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
	NotificationConfig NotificationConfigPtrInput
	// Specifies where and whether to send notifications upon changes to a Fhir store.
	NotificationConfigs FhirNotificationConfigArrayInput
	Project             pulumi.StringPtrInput
	// Configuration for how FHIR resources can be searched.
	SearchConfig SearchConfigPtrInput
	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
	StreamConfigs StreamConfigArrayInput
	// Configuration for how to validate incoming FHIR resources against configured profiles.
	ValidationConfig ValidationConfigPtrInput
	// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
	Version FhirStoreVersionPtrInput
}

func (FhirStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreArgs)(nil)).Elem()
}

type FhirStoreInput interface {
	pulumi.Input

	ToFhirStoreOutput() FhirStoreOutput
	ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput
}

func (*FhirStore) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStore)(nil)).Elem()
}

func (i *FhirStore) ToFhirStoreOutput() FhirStoreOutput {
	return i.ToFhirStoreOutputWithContext(context.Background())
}

func (i *FhirStore) ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreOutput)
}

type FhirStoreOutput struct{ *pulumi.OutputState }

func (FhirStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStore)(nil)).Elem()
}

func (o FhirStoreOutput) ToFhirStoreOutput() FhirStoreOutput {
	return o
}

func (o FhirStoreOutput) ToFhirStoreOutputWithContext(ctx context.Context) FhirStoreOutput {
	return o
}

// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
func (o FhirStoreOutput) ComplexDataTypeReferenceParsing() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.ComplexDataTypeReferenceParsing }).(pulumi.StringOutput)
}

func (o FhirStoreOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
func (o FhirStoreOutput) DefaultSearchHandlingStrict() pulumi.BoolOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.BoolOutput { return v.DefaultSearchHandlingStrict }).(pulumi.BoolOutput)
}

// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
func (o FhirStoreOutput) DisableReferentialIntegrity() pulumi.BoolOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.BoolOutput { return v.DisableReferentialIntegrity }).(pulumi.BoolOutput)
}

// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
func (o FhirStoreOutput) DisableResourceVersioning() pulumi.BoolOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.BoolOutput { return v.DisableResourceVersioning }).(pulumi.BoolOutput)
}

// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
func (o FhirStoreOutput) EnableUpdateCreate() pulumi.BoolOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.BoolOutput { return v.EnableUpdateCreate }).(pulumi.BoolOutput)
}

// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
func (o FhirStoreOutput) FhirStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringPtrOutput { return v.FhirStoreId }).(pulumi.StringPtrOutput)
}

// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
func (o FhirStoreOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o FhirStoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
func (o FhirStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
//
// Deprecated: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource". Deprecated. Use `notification_configs` instead.
func (o FhirStoreOutput) NotificationConfig() NotificationConfigResponseOutput {
	return o.ApplyT(func(v *FhirStore) NotificationConfigResponseOutput { return v.NotificationConfig }).(NotificationConfigResponseOutput)
}

// Specifies where and whether to send notifications upon changes to a Fhir store.
func (o FhirStoreOutput) NotificationConfigs() FhirNotificationConfigResponseArrayOutput {
	return o.ApplyT(func(v *FhirStore) FhirNotificationConfigResponseArrayOutput { return v.NotificationConfigs }).(FhirNotificationConfigResponseArrayOutput)
}

func (o FhirStoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Configuration for how FHIR resources can be searched.
func (o FhirStoreOutput) SearchConfig() SearchConfigResponseOutput {
	return o.ApplyT(func(v *FhirStore) SearchConfigResponseOutput { return v.SearchConfig }).(SearchConfigResponseOutput)
}

// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
func (o FhirStoreOutput) StreamConfigs() StreamConfigResponseArrayOutput {
	return o.ApplyT(func(v *FhirStore) StreamConfigResponseArrayOutput { return v.StreamConfigs }).(StreamConfigResponseArrayOutput)
}

// Configuration for how to validate incoming FHIR resources against configured profiles.
func (o FhirStoreOutput) ValidationConfig() ValidationConfigResponseOutput {
	return o.ApplyT(func(v *FhirStore) ValidationConfigResponseOutput { return v.ValidationConfig }).(ValidationConfigResponseOutput)
}

// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
func (o FhirStoreOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirStore) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreInput)(nil)).Elem(), &FhirStore{})
	pulumi.RegisterOutputType(FhirStoreOutput{})
}
