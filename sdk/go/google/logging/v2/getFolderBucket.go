// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a log bucket.
func LookupFolderBucket(ctx *pulumi.Context, args *LookupFolderBucketArgs, opts ...pulumi.InvokeOption) (*LookupFolderBucketResult, error) {
	var rv LookupFolderBucketResult
	err := ctx.Invoke("google-native:logging/v2:getFolderBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderBucketArgs struct {
	BucketId string `pulumi:"bucketId"`
	FolderId string `pulumi:"folderId"`
	Location string `pulumi:"location"`
}

type LookupFolderBucketResult struct {
	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled bool `pulumi:"analyticsEnabled"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings CmekSettingsResponse `pulumi:"cmekSettings"`
	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime string `pulumi:"createTime"`
	// Describes this bucket.
	Description string `pulumi:"description"`
	// A list of indexed fields and related configuration data.
	IndexConfigs []IndexConfigResponse `pulumi:"indexConfigs"`
	// The bucket lifecycle state.
	LifecycleState string `pulumi:"lifecycleState"`
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked bool `pulumi:"locked"`
	// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
	Name string `pulumi:"name"`
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields []string `pulumi:"restrictedFields"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays int `pulumi:"retentionDays"`
	// The last update timestamp of the bucket.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFolderBucketOutput(ctx *pulumi.Context, args LookupFolderBucketOutputArgs, opts ...pulumi.InvokeOption) LookupFolderBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderBucketResult, error) {
			args := v.(LookupFolderBucketArgs)
			r, err := LookupFolderBucket(ctx, &args, opts...)
			var s LookupFolderBucketResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderBucketResultOutput)
}

type LookupFolderBucketOutputArgs struct {
	BucketId pulumi.StringInput `pulumi:"bucketId"`
	FolderId pulumi.StringInput `pulumi:"folderId"`
	Location pulumi.StringInput `pulumi:"location"`
}

func (LookupFolderBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderBucketArgs)(nil)).Elem()
}

type LookupFolderBucketResultOutput struct{ *pulumi.OutputState }

func (LookupFolderBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderBucketResult)(nil)).Elem()
}

func (o LookupFolderBucketResultOutput) ToLookupFolderBucketResultOutput() LookupFolderBucketResultOutput {
	return o
}

func (o LookupFolderBucketResultOutput) ToLookupFolderBucketResultOutputWithContext(ctx context.Context) LookupFolderBucketResultOutput {
	return o
}

// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
func (o LookupFolderBucketResultOutput) AnalyticsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) bool { return v.AnalyticsEnabled }).(pulumi.BoolOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
func (o LookupFolderBucketResultOutput) CmekSettings() CmekSettingsResponseOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) CmekSettingsResponse { return v.CmekSettings }).(CmekSettingsResponseOutput)
}

// The creation timestamp of the bucket. This is not set for any of the default buckets.
func (o LookupFolderBucketResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this bucket.
func (o LookupFolderBucketResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) string { return v.Description }).(pulumi.StringOutput)
}

// A list of indexed fields and related configuration data.
func (o LookupFolderBucketResultOutput) IndexConfigs() IndexConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) []IndexConfigResponse { return v.IndexConfigs }).(IndexConfigResponseArrayOutput)
}

// The bucket lifecycle state.
func (o LookupFolderBucketResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
func (o LookupFolderBucketResultOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) bool { return v.Locked }).(pulumi.BoolOutput)
}

// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
func (o LookupFolderBucketResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) string { return v.Name }).(pulumi.StringOutput)
}

// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
func (o LookupFolderBucketResultOutput) RestrictedFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) []string { return v.RestrictedFields }).(pulumi.StringArrayOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
func (o LookupFolderBucketResultOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) int { return v.RetentionDays }).(pulumi.IntOutput)
}

// The last update timestamp of the bucket.
func (o LookupFolderBucketResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBucketResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderBucketResultOutput{})
}
