// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns information about the capacity commitment.
func LookupCapacityCommitment(ctx *pulumi.Context, args *LookupCapacityCommitmentArgs, opts ...pulumi.InvokeOption) (*LookupCapacityCommitmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityCommitmentResult
	err := ctx.Invoke("google-native:bigqueryreservation/v1:getCapacityCommitment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityCommitmentArgs struct {
	CapacityCommitmentId string  `pulumi:"capacityCommitmentId"`
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
}

type LookupCapacityCommitmentResult struct {
	// The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	CommitmentEndTime string `pulumi:"commitmentEndTime"`
	// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	CommitmentStartTime string `pulumi:"commitmentStartTime"`
	// Edition of the capacity commitment.
	Edition string `pulumi:"edition"`
	// For FAILED commitment plan, provides the reason of failure.
	FailureStatus StatusResponse `pulumi:"failureStatus"`
	// Applicable only for commitments located within one of the BigQuery multi-regions (US or EU). If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
	MultiRegionAuxiliary bool `pulumi:"multiRegionAuxiliary"`
	// The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123` The commitment_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	Name string `pulumi:"name"`
	// Capacity commitment commitment plan.
	Plan string `pulumi:"plan"`
	// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
	RenewalPlan string `pulumi:"renewalPlan"`
	// Number of slots in this commitment.
	SlotCount string `pulumi:"slotCount"`
	// State of the commitment.
	State string `pulumi:"state"`
}

func LookupCapacityCommitmentOutput(ctx *pulumi.Context, args LookupCapacityCommitmentOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityCommitmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityCommitmentResult, error) {
			args := v.(LookupCapacityCommitmentArgs)
			r, err := LookupCapacityCommitment(ctx, &args, opts...)
			var s LookupCapacityCommitmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityCommitmentResultOutput)
}

type LookupCapacityCommitmentOutputArgs struct {
	CapacityCommitmentId pulumi.StringInput    `pulumi:"capacityCommitmentId"`
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCapacityCommitmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityCommitmentArgs)(nil)).Elem()
}

type LookupCapacityCommitmentResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityCommitmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityCommitmentResult)(nil)).Elem()
}

func (o LookupCapacityCommitmentResultOutput) ToLookupCapacityCommitmentResultOutput() LookupCapacityCommitmentResultOutput {
	return o
}

func (o LookupCapacityCommitmentResultOutput) ToLookupCapacityCommitmentResultOutputWithContext(ctx context.Context) LookupCapacityCommitmentResultOutput {
	return o
}

// The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
func (o LookupCapacityCommitmentResultOutput) CommitmentEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.CommitmentEndTime }).(pulumi.StringOutput)
}

// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
func (o LookupCapacityCommitmentResultOutput) CommitmentStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.CommitmentStartTime }).(pulumi.StringOutput)
}

// Edition of the capacity commitment.
func (o LookupCapacityCommitmentResultOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.Edition }).(pulumi.StringOutput)
}

// For FAILED commitment plan, provides the reason of failure.
func (o LookupCapacityCommitmentResultOutput) FailureStatus() StatusResponseOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) StatusResponse { return v.FailureStatus }).(StatusResponseOutput)
}

// Applicable only for commitments located within one of the BigQuery multi-regions (US or EU). If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
func (o LookupCapacityCommitmentResultOutput) MultiRegionAuxiliary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) bool { return v.MultiRegionAuxiliary }).(pulumi.BoolOutput)
}

// The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123` The commitment_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
func (o LookupCapacityCommitmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Capacity commitment commitment plan.
func (o LookupCapacityCommitmentResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.Plan }).(pulumi.StringOutput)
}

// The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
func (o LookupCapacityCommitmentResultOutput) RenewalPlan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.RenewalPlan }).(pulumi.StringOutput)
}

// Number of slots in this commitment.
func (o LookupCapacityCommitmentResultOutput) SlotCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.SlotCount }).(pulumi.StringOutput)
}

// State of the commitment.
func (o LookupCapacityCommitmentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityCommitmentResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityCommitmentResultOutput{})
}
