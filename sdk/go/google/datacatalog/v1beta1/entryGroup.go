// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A maximum of 10,000 entry groups may be created per organization across all locations. Users should enable the Data Catalog API in the project identified by the `parent` parameter (see [Data Catalog Resource Project] (https://cloud.google.com/data-catalog/docs/concepts/resource-project) for more information).
type EntryGroup struct {
	pulumi.CustomResourceState

	// Timestamps about this EntryGroup. Default value is empty timestamps.
	DataCatalogTimestamps GoogleCloudDatacatalogV1beta1SystemTimestampsResponseOutput `pulumi:"dataCatalogTimestamps"`
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
	Description pulumi.StringOutput `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEntryGroup registers a new resource with the given unique name, arguments, and options.
func NewEntryGroup(ctx *pulumi.Context,
	name string, args *EntryGroupArgs, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource EntryGroup
	err := ctx.RegisterResource("google-native:datacatalog/v1beta1:EntryGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroup gets an existing EntryGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupState, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	var resource EntryGroup
	err := ctx.ReadResource("google-native:datacatalog/v1beta1:EntryGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroup resources.
type entryGroupState struct {
	// Timestamps about this EntryGroup. Default value is empty timestamps.
	DataCatalogTimestamps *GoogleCloudDatacatalogV1beta1SystemTimestampsResponse `pulumi:"dataCatalogTimestamps"`
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
	Name *string `pulumi:"name"`
}

type EntryGroupState struct {
	// Timestamps about this EntryGroup. Default value is empty timestamps.
	DataCatalogTimestamps GoogleCloudDatacatalogV1beta1SystemTimestampsResponsePtrInput
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName pulumi.StringPtrInput
	// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
	Name pulumi.StringPtrInput
}

func (EntryGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupState)(nil)).Elem()
}

type entryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName  *string `pulumi:"displayName"`
	EntryGroupId string  `pulumi:"entryGroupId"`
	Location     string  `pulumi:"location"`
	// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a EntryGroup resource.
type EntryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents. Default value is an empty string.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011". Default value is an empty string.
	DisplayName  pulumi.StringPtrInput
	EntryGroupId pulumi.StringInput
	Location     pulumi.StringInput
	// The resource name of the entry group in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id} Note that this EntryGroup and its child resources may not actually be stored in the location in this name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (EntryGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupArgs)(nil)).Elem()
}

type EntryGroupInput interface {
	pulumi.Input

	ToEntryGroupOutput() EntryGroupOutput
	ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput
}

func (*EntryGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroup)(nil))
}

func (i *EntryGroup) ToEntryGroupOutput() EntryGroupOutput {
	return i.ToEntryGroupOutputWithContext(context.Background())
}

func (i *EntryGroup) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupOutput)
}

type EntryGroupOutput struct {
	*pulumi.OutputState
}

func (EntryGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroup)(nil))
}

func (o EntryGroupOutput) ToEntryGroupOutput() EntryGroupOutput {
	return o
}

func (o EntryGroupOutput) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EntryGroupOutput{})
}
