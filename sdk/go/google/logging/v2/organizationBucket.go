// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a log bucket that can be used to store log entries. After a bucket has been created, the bucket's location cannot be changed.
// Auto-naming is currently not supported for this resource.
type OrganizationBucket struct {
	pulumi.CustomResourceState

	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled pulumi.BoolOutput `pulumi:"analyticsEnabled"`
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings CmekSettingsResponseOutput `pulumi:"cmekSettings"`
	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of indexed fields and related configuration data.
	IndexConfigs IndexConfigResponseArrayOutput `pulumi:"indexConfigs"`
	// The bucket lifecycle state.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	Location       pulumi.StringOutput `pulumi:"location"`
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields pulumi.StringArrayOutput `pulumi:"restrictedFields"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The last update timestamp of the bucket.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationBucket registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBucket(ctx *pulumi.Context,
	name string, args *OrganizationBucketArgs, opts ...pulumi.ResourceOption) (*OrganizationBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucketId",
		"location",
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationBucket
	err := ctx.RegisterResource("google-native:logging/v2:OrganizationBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBucket gets an existing OrganizationBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBucketState, opts ...pulumi.ResourceOption) (*OrganizationBucket, error) {
	var resource OrganizationBucket
	err := ctx.ReadResource("google-native:logging/v2:OrganizationBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBucket resources.
type organizationBucketState struct {
}

type OrganizationBucketState struct {
}

func (OrganizationBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketState)(nil)).Elem()
}

type organizationBucketArgs struct {
	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled *bool `pulumi:"analyticsEnabled"`
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *CmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// A list of indexed fields and related configuration data.
	IndexConfigs []IndexConfig `pulumi:"indexConfigs"`
	Location     *string       `pulumi:"location"`
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked         *bool  `pulumi:"locked"`
	OrganizationId string `pulumi:"organizationId"`
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields []string `pulumi:"restrictedFields"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a OrganizationBucket resource.
type OrganizationBucketArgs struct {
	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled pulumi.BoolPtrInput
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId pulumi.StringInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings CmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// A list of indexed fields and related configuration data.
	IndexConfigs IndexConfigArrayInput
	Location     pulumi.StringPtrInput
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked         pulumi.BoolPtrInput
	OrganizationId pulumi.StringInput
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields pulumi.StringArrayInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketArgs)(nil)).Elem()
}

type OrganizationBucketInput interface {
	pulumi.Input

	ToOrganizationBucketOutput() OrganizationBucketOutput
	ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput
}

func (*OrganizationBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationBucket)(nil)).Elem()
}

func (i *OrganizationBucket) ToOrganizationBucketOutput() OrganizationBucketOutput {
	return i.ToOrganizationBucketOutputWithContext(context.Background())
}

func (i *OrganizationBucket) ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketOutput)
}

type OrganizationBucketOutput struct{ *pulumi.OutputState }

func (OrganizationBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationBucket)(nil)).Elem()
}

func (o OrganizationBucketOutput) ToOrganizationBucketOutput() OrganizationBucketOutput {
	return o
}

func (o OrganizationBucketOutput) ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput {
	return o
}

// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
func (o OrganizationBucketOutput) AnalyticsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.BoolOutput { return v.AnalyticsEnabled }).(pulumi.BoolOutput)
}

// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
func (o OrganizationBucketOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
func (o OrganizationBucketOutput) CmekSettings() CmekSettingsResponseOutput {
	return o.ApplyT(func(v *OrganizationBucket) CmekSettingsResponseOutput { return v.CmekSettings }).(CmekSettingsResponseOutput)
}

// The creation timestamp of the bucket. This is not set for any of the default buckets.
func (o OrganizationBucketOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this bucket.
func (o OrganizationBucketOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A list of indexed fields and related configuration data.
func (o OrganizationBucketOutput) IndexConfigs() IndexConfigResponseArrayOutput {
	return o.ApplyT(func(v *OrganizationBucket) IndexConfigResponseArrayOutput { return v.IndexConfigs }).(IndexConfigResponseArrayOutput)
}

// The bucket lifecycle state.
func (o OrganizationBucketOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

func (o OrganizationBucketOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
func (o OrganizationBucketOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
func (o OrganizationBucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrganizationBucketOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
func (o OrganizationBucketOutput) RestrictedFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringArrayOutput { return v.RestrictedFields }).(pulumi.StringArrayOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
func (o OrganizationBucketOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

// The last update timestamp of the bucket.
func (o OrganizationBucketOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucket) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationBucketInput)(nil)).Elem(), &OrganizationBucket{})
	pulumi.RegisterOutputType(OrganizationBucketOutput{})
}
