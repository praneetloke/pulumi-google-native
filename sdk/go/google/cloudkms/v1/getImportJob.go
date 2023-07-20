// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns metadata for a given ImportJob.
func LookupImportJob(ctx *pulumi.Context, args *LookupImportJobArgs, opts ...pulumi.InvokeOption) (*LookupImportJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImportJobResult
	err := ctx.Invoke("google-native:cloudkms/v1:getImportJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImportJobArgs struct {
	ImportJobId string  `pulumi:"importJobId"`
	KeyRingId   string  `pulumi:"keyRingId"`
	Location    string  `pulumi:"location"`
	Project     *string `pulumi:"project"`
}

type LookupImportJobResult struct {
	// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
	Attestation KeyOperationAttestationResponse `pulumi:"attestation"`
	// The time at which this ImportJob was created.
	CreateTime string `pulumi:"createTime"`
	// The time this ImportJob expired. Only present if state is EXPIRED.
	ExpireEventTime string `pulumi:"expireEventTime"`
	// The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
	ExpireTime string `pulumi:"expireTime"`
	// The time this ImportJob's key material was generated.
	GenerateTime string `pulumi:"generateTime"`
	// Immutable. The wrapping method to be used for incoming key material.
	ImportMethod string `pulumi:"importMethod"`
	// The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
	Name string `pulumi:"name"`
	// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
	ProtectionLevel string `pulumi:"protectionLevel"`
	// The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
	PublicKey WrappingPublicKeyResponse `pulumi:"publicKey"`
	// The current state of the ImportJob, indicating if it can be used.
	State string `pulumi:"state"`
}

func LookupImportJobOutput(ctx *pulumi.Context, args LookupImportJobOutputArgs, opts ...pulumi.InvokeOption) LookupImportJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportJobResult, error) {
			args := v.(LookupImportJobArgs)
			r, err := LookupImportJob(ctx, &args, opts...)
			var s LookupImportJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportJobResultOutput)
}

type LookupImportJobOutputArgs struct {
	ImportJobId pulumi.StringInput    `pulumi:"importJobId"`
	KeyRingId   pulumi.StringInput    `pulumi:"keyRingId"`
	Location    pulumi.StringInput    `pulumi:"location"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupImportJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportJobArgs)(nil)).Elem()
}

type LookupImportJobResultOutput struct{ *pulumi.OutputState }

func (LookupImportJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportJobResult)(nil)).Elem()
}

func (o LookupImportJobResultOutput) ToLookupImportJobResultOutput() LookupImportJobResultOutput {
	return o
}

func (o LookupImportJobResultOutput) ToLookupImportJobResultOutputWithContext(ctx context.Context) LookupImportJobResultOutput {
	return o
}

// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
func (o LookupImportJobResultOutput) Attestation() KeyOperationAttestationResponseOutput {
	return o.ApplyT(func(v LookupImportJobResult) KeyOperationAttestationResponse { return v.Attestation }).(KeyOperationAttestationResponseOutput)
}

// The time at which this ImportJob was created.
func (o LookupImportJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time this ImportJob expired. Only present if state is EXPIRED.
func (o LookupImportJobResultOutput) ExpireEventTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.ExpireEventTime }).(pulumi.StringOutput)
}

// The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
func (o LookupImportJobResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// The time this ImportJob's key material was generated.
func (o LookupImportJobResultOutput) GenerateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.GenerateTime }).(pulumi.StringOutput)
}

// Immutable. The wrapping method to be used for incoming key material.
func (o LookupImportJobResultOutput) ImportMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.ImportMethod }).(pulumi.StringOutput)
}

// The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
func (o LookupImportJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
func (o LookupImportJobResultOutput) ProtectionLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.ProtectionLevel }).(pulumi.StringOutput)
}

// The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
func (o LookupImportJobResultOutput) PublicKey() WrappingPublicKeyResponseOutput {
	return o.ApplyT(func(v LookupImportJobResult) WrappingPublicKeyResponse { return v.PublicKey }).(WrappingPublicKeyResponseOutput)
}

// The current state of the ImportJob, indicating if it can be used.
func (o LookupImportJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportJobResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportJobResultOutput{})
}
