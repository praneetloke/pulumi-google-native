// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a view over logs in a bucket. A bucket may contain a maximum of 50 views.
type BucketView struct {
	pulumi.CustomResourceState

	// The creation timestamp of the view.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this view.
	Description pulumi.StringOutput `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update timestamp of the view.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBucketView registers a new resource with the given unique name, arguments, and options.
func NewBucketView(ctx *pulumi.Context,
	name string, args *BucketViewArgs, opts ...pulumi.ResourceOption) (*BucketView, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	var resource BucketView
	err := ctx.RegisterResource("google-native:logging/v2:BucketView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketView gets an existing BucketView resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketViewState, opts ...pulumi.ResourceOption) (*BucketView, error) {
	var resource BucketView
	err := ctx.ReadResource("google-native:logging/v2:BucketView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketView resources.
type bucketViewState struct {
	// The creation timestamp of the view.
	CreateTime *string `pulumi:"createTime"`
	// Describes this view.
	Description *string `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter *string `pulumi:"filter"`
	// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
	Name *string `pulumi:"name"`
	// The last update timestamp of the view.
	UpdateTime *string `pulumi:"updateTime"`
}

type BucketViewState struct {
	// The creation timestamp of the view.
	CreateTime pulumi.StringPtrInput
	// Describes this view.
	Description pulumi.StringPtrInput
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter pulumi.StringPtrInput
	// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
	Name pulumi.StringPtrInput
	// The last update timestamp of the view.
	UpdateTime pulumi.StringPtrInput
}

func (BucketViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketViewState)(nil)).Elem()
}

type bucketViewArgs struct {
	BucketId string `pulumi:"bucketId"`
	// Describes this view.
	Description *string `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter   *string `pulumi:"filter"`
	Location string  `pulumi:"location"`
	// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	ViewId  string  `pulumi:"viewId"`
}

// The set of arguments for constructing a BucketView resource.
type BucketViewArgs struct {
	BucketId pulumi.StringInput
	// Describes this view.
	Description pulumi.StringPtrInput
	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter   pulumi.StringPtrInput
	Location pulumi.StringInput
	// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	ViewId  pulumi.StringInput
}

func (BucketViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketViewArgs)(nil)).Elem()
}

type BucketViewInput interface {
	pulumi.Input

	ToBucketViewOutput() BucketViewOutput
	ToBucketViewOutputWithContext(ctx context.Context) BucketViewOutput
}

func (*BucketView) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketView)(nil))
}

func (i *BucketView) ToBucketViewOutput() BucketViewOutput {
	return i.ToBucketViewOutputWithContext(context.Background())
}

func (i *BucketView) ToBucketViewOutputWithContext(ctx context.Context) BucketViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketViewOutput)
}

type BucketViewOutput struct {
	*pulumi.OutputState
}

func (BucketViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketView)(nil))
}

func (o BucketViewOutput) ToBucketViewOutput() BucketViewOutput {
	return o
}

func (o BucketViewOutput) ToBucketViewOutputWithContext(ctx context.Context) BucketViewOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketViewOutput{})
}
