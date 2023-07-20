// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:file/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// The time when the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The description of the instance (2048 characters or less).
	Description string `pulumi:"description"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag string `pulumi:"etag"`
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares []FileShareConfigResponse `pulumi:"fileShares"`
	// KMS key name used for data encryption.
	KmsKeyName string `pulumi:"kmsKeyName"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the instance, in the format `projects/{project}/locations/{location}/instances/{instance}`.
	Name string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks []NetworkConfigResponse `pulumi:"networks"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// The instance state.
	State string `pulumi:"state"`
	// Additional information about the instance state, if available.
	StatusMessage string `pulumi:"statusMessage"`
	// Field indicates all the reasons the instance is in "SUSPENDED" state.
	SuspensionReasons []string `pulumi:"suspensionReasons"`
	// The service tier of the instance.
	Tier string `pulumi:"tier"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// The time when the instance was created.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the instance (2048 characters or less).
func (o LookupInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
func (o LookupInstanceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// File system shares on the instance. For this version, only a single file share is supported.
func (o LookupInstanceResultOutput) FileShares() FileShareConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []FileShareConfigResponse { return v.FileShares }).(FileShareConfigResponseArrayOutput)
}

// KMS key name used for data encryption.
func (o LookupInstanceResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the instance, in the format `projects/{project}/locations/{location}/instances/{instance}`.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// VPC networks to which the instance is connected. For this version, only a single network is supported.
func (o LookupInstanceResultOutput) Networks() NetworkConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []NetworkConfigResponse { return v.Networks }).(NetworkConfigResponseArrayOutput)
}

// Reserved for future use.
func (o LookupInstanceResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// The instance state.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// Additional information about the instance state, if available.
func (o LookupInstanceResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Field indicates all the reasons the instance is in "SUSPENDED" state.
func (o LookupInstanceResultOutput) SuspensionReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.SuspensionReasons }).(pulumi.StringArrayOutput)
}

// The service tier of the instance.
func (o LookupInstanceResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Tier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
