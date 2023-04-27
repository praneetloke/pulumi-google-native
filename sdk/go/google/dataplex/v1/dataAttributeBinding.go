// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a DataAttributeBinding resource.
// Auto-naming is currently not supported for this resource.
type DataAttributeBinding struct {
	pulumi.CustomResourceState

	// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
	Attributes pulumi.StringArrayOutput `pulumi:"attributes"`
	// The time when the DataAttributeBinding was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
	DataAttributeBindingId pulumi.StringOutput `pulumi:"dataAttributeBindingId"`
	// Optional. Description of the DataAttributeBinding.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. User friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. User-defined labels for the DataAttributeBinding.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The relative resource name of the Data Attribute Binding, of the form: projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
	Paths   GoogleCloudDataplexV1DataAttributeBindingPathResponseArrayOutput `pulumi:"paths"`
	Project pulumi.StringOutput                                              `pulumi:"project"`
	// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// System generated globally unique ID for the DataAttributeBinding. This ID will be different if the DataAttributeBinding is deleted and re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the DataAttributeBinding was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataAttributeBinding registers a new resource with the given unique name, arguments, and options.
func NewDataAttributeBinding(ctx *pulumi.Context,
	name string, args *DataAttributeBindingArgs, opts ...pulumi.ResourceOption) (*DataAttributeBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataAttributeBindingId == nil {
		return nil, errors.New("invalid value for required argument 'DataAttributeBindingId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dataAttributeBindingId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource DataAttributeBinding
	err := ctx.RegisterResource("google-native:dataplex/v1:DataAttributeBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataAttributeBinding gets an existing DataAttributeBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataAttributeBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataAttributeBindingState, opts ...pulumi.ResourceOption) (*DataAttributeBinding, error) {
	var resource DataAttributeBinding
	err := ctx.ReadResource("google-native:dataplex/v1:DataAttributeBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataAttributeBinding resources.
type dataAttributeBindingState struct {
}

type DataAttributeBindingState struct {
}

func (DataAttributeBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAttributeBindingState)(nil)).Elem()
}

type dataAttributeBindingArgs struct {
	// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
	Attributes []string `pulumi:"attributes"`
	// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
	DataAttributeBindingId string `pulumi:"dataAttributeBindingId"`
	// Optional. Description of the DataAttributeBinding.
	Description *string `pulumi:"description"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
	Etag *string `pulumi:"etag"`
	// Optional. User-defined labels for the DataAttributeBinding.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
	Paths   []GoogleCloudDataplexV1DataAttributeBindingPath `pulumi:"paths"`
	Project *string                                         `pulumi:"project"`
	// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
	Resource *string `pulumi:"resource"`
}

// The set of arguments for constructing a DataAttributeBinding resource.
type DataAttributeBindingArgs struct {
	// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
	Attributes pulumi.StringArrayInput
	// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
	DataAttributeBindingId pulumi.StringInput
	// Optional. Description of the DataAttributeBinding.
	Description pulumi.StringPtrInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
	Etag pulumi.StringPtrInput
	// Optional. User-defined labels for the DataAttributeBinding.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
	Paths   GoogleCloudDataplexV1DataAttributeBindingPathArrayInput
	Project pulumi.StringPtrInput
	// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
	Resource pulumi.StringPtrInput
}

func (DataAttributeBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAttributeBindingArgs)(nil)).Elem()
}

type DataAttributeBindingInput interface {
	pulumi.Input

	ToDataAttributeBindingOutput() DataAttributeBindingOutput
	ToDataAttributeBindingOutputWithContext(ctx context.Context) DataAttributeBindingOutput
}

func (*DataAttributeBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAttributeBinding)(nil)).Elem()
}

func (i *DataAttributeBinding) ToDataAttributeBindingOutput() DataAttributeBindingOutput {
	return i.ToDataAttributeBindingOutputWithContext(context.Background())
}

func (i *DataAttributeBinding) ToDataAttributeBindingOutputWithContext(ctx context.Context) DataAttributeBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAttributeBindingOutput)
}

type DataAttributeBindingOutput struct{ *pulumi.OutputState }

func (DataAttributeBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAttributeBinding)(nil)).Elem()
}

func (o DataAttributeBindingOutput) ToDataAttributeBindingOutput() DataAttributeBindingOutput {
	return o
}

func (o DataAttributeBindingOutput) ToDataAttributeBindingOutputWithContext(ctx context.Context) DataAttributeBindingOutput {
	return o
}

// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
func (o DataAttributeBindingOutput) Attributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringArrayOutput { return v.Attributes }).(pulumi.StringArrayOutput)
}

// The time when the DataAttributeBinding was created.
func (o DataAttributeBindingOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
func (o DataAttributeBindingOutput) DataAttributeBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.DataAttributeBindingId }).(pulumi.StringOutput)
}

// Optional. Description of the DataAttributeBinding.
func (o DataAttributeBindingOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. User friendly display name.
func (o DataAttributeBindingOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
func (o DataAttributeBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. User-defined labels for the DataAttributeBinding.
func (o DataAttributeBindingOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o DataAttributeBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The relative resource name of the Data Attribute Binding, of the form: projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
func (o DataAttributeBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
func (o DataAttributeBindingOutput) Paths() GoogleCloudDataplexV1DataAttributeBindingPathResponseArrayOutput {
	return o.ApplyT(func(v *DataAttributeBinding) GoogleCloudDataplexV1DataAttributeBindingPathResponseArrayOutput {
		return v.Paths
	}).(GoogleCloudDataplexV1DataAttributeBindingPathResponseArrayOutput)
}

func (o DataAttributeBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
func (o DataAttributeBindingOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

// System generated globally unique ID for the DataAttributeBinding. This ID will be different if the DataAttributeBinding is deleted and re-created with the same name.
func (o DataAttributeBindingOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the DataAttributeBinding was last updated.
func (o DataAttributeBindingOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAttributeBinding) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataAttributeBindingInput)(nil)).Elem(), &DataAttributeBinding{})
	pulumi.RegisterOutputType(DataAttributeBindingOutput{})
}
