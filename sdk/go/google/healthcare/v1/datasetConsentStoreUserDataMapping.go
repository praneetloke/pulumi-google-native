// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new User data mapping in the parent consent store.
type DatasetConsentStoreUserDataMapping struct {
	pulumi.CustomResourceState

	// Indicates the time when this mapping was archived.
	ArchiveTime pulumi.StringOutput `pulumi:"archiveTime"`
	// Indicates whether this mapping is archived.
	Archived pulumi.BoolOutput `pulumi:"archived"`
	// Required. A unique identifier for the mapped resource.
	DataId pulumi.StringOutput `pulumi:"dataId"`
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes AttributeResponseArrayOutput `pulumi:"resourceAttributes"`
	// Required. User's UUID provided by the client.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewDatasetConsentStoreUserDataMapping registers a new resource with the given unique name, arguments, and options.
func NewDatasetConsentStoreUserDataMapping(ctx *pulumi.Context,
	name string, args *DatasetConsentStoreUserDataMappingArgs, opts ...pulumi.ResourceOption) (*DatasetConsentStoreUserDataMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.UserDataMappingId == nil {
		return nil, errors.New("invalid value for required argument 'UserDataMappingId'")
	}
	var resource DatasetConsentStoreUserDataMapping
	err := ctx.RegisterResource("google-native:healthcare/v1:DatasetConsentStoreUserDataMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetConsentStoreUserDataMapping gets an existing DatasetConsentStoreUserDataMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetConsentStoreUserDataMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetConsentStoreUserDataMappingState, opts ...pulumi.ResourceOption) (*DatasetConsentStoreUserDataMapping, error) {
	var resource DatasetConsentStoreUserDataMapping
	err := ctx.ReadResource("google-native:healthcare/v1:DatasetConsentStoreUserDataMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetConsentStoreUserDataMapping resources.
type datasetConsentStoreUserDataMappingState struct {
	// Indicates the time when this mapping was archived.
	ArchiveTime *string `pulumi:"archiveTime"`
	// Indicates whether this mapping is archived.
	Archived *bool `pulumi:"archived"`
	// Required. A unique identifier for the mapped resource.
	DataId *string `pulumi:"dataId"`
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name *string `pulumi:"name"`
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes []AttributeResponse `pulumi:"resourceAttributes"`
	// Required. User's UUID provided by the client.
	UserId *string `pulumi:"userId"`
}

type DatasetConsentStoreUserDataMappingState struct {
	// Indicates the time when this mapping was archived.
	ArchiveTime pulumi.StringPtrInput
	// Indicates whether this mapping is archived.
	Archived pulumi.BoolPtrInput
	// Required. A unique identifier for the mapped resource.
	DataId pulumi.StringPtrInput
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name pulumi.StringPtrInput
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes AttributeResponseArrayInput
	// Required. User's UUID provided by the client.
	UserId pulumi.StringPtrInput
}

func (DatasetConsentStoreUserDataMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetConsentStoreUserDataMappingState)(nil)).Elem()
}

type datasetConsentStoreUserDataMappingArgs struct {
	ConsentStoreId string `pulumi:"consentStoreId"`
	// Required. A unique identifier for the mapped resource.
	DataId    *string `pulumi:"dataId"`
	DatasetId string  `pulumi:"datasetId"`
	Location  string  `pulumi:"location"`
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes []Attribute `pulumi:"resourceAttributes"`
	UserDataMappingId  string      `pulumi:"userDataMappingId"`
	// Required. User's UUID provided by the client.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a DatasetConsentStoreUserDataMapping resource.
type DatasetConsentStoreUserDataMappingArgs struct {
	ConsentStoreId pulumi.StringInput
	// Required. A unique identifier for the mapped resource.
	DataId    pulumi.StringPtrInput
	DatasetId pulumi.StringInput
	Location  pulumi.StringInput
	// Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
	ResourceAttributes AttributeArrayInput
	UserDataMappingId  pulumi.StringInput
	// Required. User's UUID provided by the client.
	UserId pulumi.StringPtrInput
}

func (DatasetConsentStoreUserDataMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetConsentStoreUserDataMappingArgs)(nil)).Elem()
}

type DatasetConsentStoreUserDataMappingInput interface {
	pulumi.Input

	ToDatasetConsentStoreUserDataMappingOutput() DatasetConsentStoreUserDataMappingOutput
	ToDatasetConsentStoreUserDataMappingOutputWithContext(ctx context.Context) DatasetConsentStoreUserDataMappingOutput
}

func (*DatasetConsentStoreUserDataMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetConsentStoreUserDataMapping)(nil))
}

func (i *DatasetConsentStoreUserDataMapping) ToDatasetConsentStoreUserDataMappingOutput() DatasetConsentStoreUserDataMappingOutput {
	return i.ToDatasetConsentStoreUserDataMappingOutputWithContext(context.Background())
}

func (i *DatasetConsentStoreUserDataMapping) ToDatasetConsentStoreUserDataMappingOutputWithContext(ctx context.Context) DatasetConsentStoreUserDataMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetConsentStoreUserDataMappingOutput)
}

type DatasetConsentStoreUserDataMappingOutput struct {
	*pulumi.OutputState
}

func (DatasetConsentStoreUserDataMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetConsentStoreUserDataMapping)(nil))
}

func (o DatasetConsentStoreUserDataMappingOutput) ToDatasetConsentStoreUserDataMappingOutput() DatasetConsentStoreUserDataMappingOutput {
	return o
}

func (o DatasetConsentStoreUserDataMappingOutput) ToDatasetConsentStoreUserDataMappingOutputWithContext(ctx context.Context) DatasetConsentStoreUserDataMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetConsentStoreUserDataMappingOutput{})
}
