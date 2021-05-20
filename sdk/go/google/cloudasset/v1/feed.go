// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a feed in a parent project/folder/organization to listen to its asset updates.
type Feed struct {
	pulumi.CustomResourceState

	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more info.
	AssetNames pulumi.StringArrayOutput `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` See [this topic](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for a list of all supported asset types.
	AssetTypes pulumi.StringArrayOutput `pulumi:"assetTypes"`
	// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes#feed_with_condition) for detailed instructions.
	Condition ExprResponseOutput `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig FeedOutputConfigResponseOutput `pulumi:"feedOutputConfig"`
	// Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFeed registers a new resource with the given unique name, arguments, and options.
func NewFeed(ctx *pulumi.Context,
	name string, args *FeedArgs, opts ...pulumi.ResourceOption) (*Feed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeedId == nil {
		return nil, errors.New("invalid value for required argument 'FeedId'")
	}
	if args.V1Id == nil {
		return nil, errors.New("invalid value for required argument 'V1Id'")
	}
	if args.V1Id1 == nil {
		return nil, errors.New("invalid value for required argument 'V1Id1'")
	}
	var resource Feed
	err := ctx.RegisterResource("google-native:cloudasset/v1:Feed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeed gets an existing Feed resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeed(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeedState, opts ...pulumi.ResourceOption) (*Feed, error) {
	var resource Feed
	err := ctx.ReadResource("google-native:cloudasset/v1:Feed", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Feed resources.
type feedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` See [this topic](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for a list of all supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes#feed_with_condition) for detailed instructions.
	Condition *ExprResponse `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType *string `pulumi:"contentType"`
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig *FeedOutputConfigResponse `pulumi:"feedOutputConfig"`
	// Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Name *string `pulumi:"name"`
}

type FeedState struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` See [this topic](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for a list of all supported asset types.
	AssetTypes pulumi.StringArrayInput
	// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes#feed_with_condition) for detailed instructions.
	Condition ExprResponsePtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType pulumi.StringPtrInput
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig FeedOutputConfigResponsePtrInput
	// Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Name pulumi.StringPtrInput
}

func (FeedState) ElementType() reflect.Type {
	return reflect.TypeOf((*feedState)(nil)).Elem()
}

type feedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more info.
	AssetNames []string `pulumi:"assetNames"`
	// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` See [this topic](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for a list of all supported asset types.
	AssetTypes []string `pulumi:"assetTypes"`
	// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes#feed_with_condition) for detailed instructions.
	Condition *Expr `pulumi:"condition"`
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType *string `pulumi:"contentType"`
	// Required. This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
	FeedId string `pulumi:"feedId"`
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig *FeedOutputConfig `pulumi:"feedOutputConfig"`
	// Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Name  *string `pulumi:"name"`
	V1Id  string  `pulumi:"v1Id"`
	V1Id1 string  `pulumi:"v1Id1"`
}

// The set of arguments for constructing a Feed resource.
type FeedArgs struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more info.
	AssetNames pulumi.StringArrayInput
	// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` See [this topic](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for a list of all supported asset types.
	AssetTypes pulumi.StringArrayInput
	// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes#feed_with_condition) for detailed instructions.
	Condition ExprPtrInput
	// Asset content type. If not specified, no content but the asset name and type will be returned.
	ContentType pulumi.StringPtrInput
	// Required. This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
	FeedId pulumi.StringInput
	// Required. Feed output configuration defining where the asset updates are published to.
	FeedOutputConfig FeedOutputConfigPtrInput
	// Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Name  pulumi.StringPtrInput
	V1Id  pulumi.StringInput
	V1Id1 pulumi.StringInput
}

func (FeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*feedArgs)(nil)).Elem()
}

type FeedInput interface {
	pulumi.Input

	ToFeedOutput() FeedOutput
	ToFeedOutputWithContext(ctx context.Context) FeedOutput
}

func (*Feed) ElementType() reflect.Type {
	return reflect.TypeOf((*Feed)(nil))
}

func (i *Feed) ToFeedOutput() FeedOutput {
	return i.ToFeedOutputWithContext(context.Background())
}

func (i *Feed) ToFeedOutputWithContext(ctx context.Context) FeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedOutput)
}

type FeedOutput struct {
	*pulumi.OutputState
}

func (FeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Feed)(nil))
}

func (o FeedOutput) ToFeedOutput() FeedOutput {
	return o
}

func (o FeedOutput) ToFeedOutputWithContext(ctx context.Context) FeedOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FeedOutput{})
}
