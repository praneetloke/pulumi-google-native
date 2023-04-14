// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Asynchronously creates a linked dataset in BigQuery which makes it possible to use BigQuery to read the logs stored in the log bucket. A log bucket may currently only contain one link.
// Auto-naming is currently not supported for this resource.
type BillingAccountBucketLink struct {
	pulumi.CustomResourceState

	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset  BigQueryDatasetResponseOutput `pulumi:"bigqueryDataset"`
	BillingAccountId pulumi.StringOutput           `pulumi:"billingAccountId"`
	BucketId         pulumi.StringOutput           `pulumi:"bucketId"`
	// The creation timestamp of the link.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this link.The maximum length of the description is 8000 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource lifecycle state.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   pulumi.StringOutput `pulumi:"linkId"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewBillingAccountBucketLink registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountBucketLink(ctx *pulumi.Context,
	name string, args *BillingAccountBucketLinkArgs, opts ...pulumi.ResourceOption) (*BillingAccountBucketLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.LinkId == nil {
		return nil, errors.New("invalid value for required argument 'LinkId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"billingAccountId",
		"bucketId",
		"linkId",
		"location",
	})
	opts = append(opts, replaceOnChanges)
	var resource BillingAccountBucketLink
	err := ctx.RegisterResource("google-native:logging/v2:BillingAccountBucketLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountBucketLink gets an existing BillingAccountBucketLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountBucketLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountBucketLinkState, opts ...pulumi.ResourceOption) (*BillingAccountBucketLink, error) {
	var resource BillingAccountBucketLink
	err := ctx.ReadResource("google-native:logging/v2:BillingAccountBucketLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountBucketLink resources.
type billingAccountBucketLinkState struct {
}

type BillingAccountBucketLinkState struct {
}

func (BillingAccountBucketLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketLinkState)(nil)).Elem()
}

type billingAccountBucketLinkArgs struct {
	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset  *BigQueryDataset `pulumi:"bigqueryDataset"`
	BillingAccountId string           `pulumi:"billingAccountId"`
	BucketId         string           `pulumi:"bucketId"`
	// Describes this link.The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   string  `pulumi:"linkId"`
	Location *string `pulumi:"location"`
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a BillingAccountBucketLink resource.
type BillingAccountBucketLinkArgs struct {
	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset  BigQueryDatasetPtrInput
	BillingAccountId pulumi.StringInput
	BucketId         pulumi.StringInput
	// Describes this link.The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   pulumi.StringInput
	Location pulumi.StringPtrInput
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name pulumi.StringPtrInput
}

func (BillingAccountBucketLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketLinkArgs)(nil)).Elem()
}

type BillingAccountBucketLinkInput interface {
	pulumi.Input

	ToBillingAccountBucketLinkOutput() BillingAccountBucketLinkOutput
	ToBillingAccountBucketLinkOutputWithContext(ctx context.Context) BillingAccountBucketLinkOutput
}

func (*BillingAccountBucketLink) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketLink)(nil)).Elem()
}

func (i *BillingAccountBucketLink) ToBillingAccountBucketLinkOutput() BillingAccountBucketLinkOutput {
	return i.ToBillingAccountBucketLinkOutputWithContext(context.Background())
}

func (i *BillingAccountBucketLink) ToBillingAccountBucketLinkOutputWithContext(ctx context.Context) BillingAccountBucketLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBucketLinkOutput)
}

type BillingAccountBucketLinkOutput struct{ *pulumi.OutputState }

func (BillingAccountBucketLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketLink)(nil)).Elem()
}

func (o BillingAccountBucketLinkOutput) ToBillingAccountBucketLinkOutput() BillingAccountBucketLinkOutput {
	return o
}

func (o BillingAccountBucketLinkOutput) ToBillingAccountBucketLinkOutputWithContext(ctx context.Context) BillingAccountBucketLinkOutput {
	return o
}

// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
func (o BillingAccountBucketLinkOutput) BigqueryDataset() BigQueryDatasetResponseOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) BigQueryDatasetResponseOutput { return v.BigqueryDataset }).(BigQueryDatasetResponseOutput)
}

func (o BillingAccountBucketLinkOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

func (o BillingAccountBucketLinkOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The creation timestamp of the link.
func (o BillingAccountBucketLinkOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this link.The maximum length of the description is 8000 characters.
func (o BillingAccountBucketLinkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The resource lifecycle state.
func (o BillingAccountBucketLinkOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
func (o BillingAccountBucketLinkOutput) LinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.LinkId }).(pulumi.StringOutput)
}

func (o BillingAccountBucketLinkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
func (o BillingAccountBucketLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountBucketLinkInput)(nil)).Elem(), &BillingAccountBucketLink{})
	pulumi.RegisterOutputType(BillingAccountBucketLinkOutput{})
}
