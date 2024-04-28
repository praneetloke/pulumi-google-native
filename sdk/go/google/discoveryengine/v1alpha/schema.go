// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Schema.
type Schema struct {
	pulumi.CustomResourceState

	CollectionId pulumi.StringOutput `pulumi:"collectionId"`
	DataStoreId  pulumi.StringOutput `pulumi:"dataStoreId"`
	// Configurations for fields of the schema.
	FieldConfigs GoogleCloudDiscoveryengineV1alphaFieldConfigResponseArrayOutput `pulumi:"fieldConfigs"`
	// The JSON representation of the schema.
	JsonSchema pulumi.StringOutput `pulumi:"jsonSchema"`
	Location   pulumi.StringOutput `pulumi:"location"`
	// Immutable. The full resource name of the schema, in the format of `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. The ID to use for the Schema, which will become the final component of the Schema.name. This field should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
	SchemaId pulumi.StringOutput `pulumi:"schemaId"`
	// The structured representation of the schema.
	StructSchema pulumi.MapOutput `pulumi:"structSchema"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionId == nil {
		return nil, errors.New("invalid value for required argument 'CollectionId'")
	}
	if args.DataStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreId'")
	}
	if args.SchemaId == nil {
		return nil, errors.New("invalid value for required argument 'SchemaId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"collectionId",
		"dataStoreId",
		"location",
		"project",
		"schemaId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Schema
	err := ctx.RegisterResource("google-native:discoveryengine/v1alpha:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("google-native:discoveryengine/v1alpha:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	CollectionId string `pulumi:"collectionId"`
	DataStoreId  string `pulumi:"dataStoreId"`
	// The JSON representation of the schema.
	JsonSchema *string `pulumi:"jsonSchema"`
	Location   *string `pulumi:"location"`
	// Immutable. The full resource name of the schema, in the format of `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The ID to use for the Schema, which will become the final component of the Schema.name. This field should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
	SchemaId string `pulumi:"schemaId"`
	// The structured representation of the schema.
	StructSchema map[string]interface{} `pulumi:"structSchema"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	CollectionId pulumi.StringInput
	DataStoreId  pulumi.StringInput
	// The JSON representation of the schema.
	JsonSchema pulumi.StringPtrInput
	Location   pulumi.StringPtrInput
	// Immutable. The full resource name of the schema, in the format of `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The ID to use for the Schema, which will become the final component of the Schema.name. This field should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
	SchemaId pulumi.StringInput
	// The structured representation of the schema.
	StructSchema pulumi.MapInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) CollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.CollectionId }).(pulumi.StringOutput)
}

func (o SchemaOutput) DataStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.DataStoreId }).(pulumi.StringOutput)
}

// Configurations for fields of the schema.
func (o SchemaOutput) FieldConfigs() GoogleCloudDiscoveryengineV1alphaFieldConfigResponseArrayOutput {
	return o.ApplyT(func(v *Schema) GoogleCloudDiscoveryengineV1alphaFieldConfigResponseArrayOutput { return v.FieldConfigs }).(GoogleCloudDiscoveryengineV1alphaFieldConfigResponseArrayOutput)
}

// The JSON representation of the schema.
func (o SchemaOutput) JsonSchema() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.JsonSchema }).(pulumi.StringOutput)
}

func (o SchemaOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The full resource name of the schema, in the format of `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchemaOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. The ID to use for the Schema, which will become the final component of the Schema.name. This field should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
func (o SchemaOutput) SchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaId }).(pulumi.StringOutput)
}

// The structured representation of the schema.
func (o SchemaOutput) StructSchema() pulumi.MapOutput {
	return o.ApplyT(func(v *Schema) pulumi.MapOutput { return v.StructSchema }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterOutputType(SchemaOutput{})
}
