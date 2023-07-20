// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns information about the reservation.
func LookupReservation(ctx *pulumi.Context, args *LookupReservationArgs, opts ...pulumi.InvokeOption) (*LookupReservationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReservationResult
	err := ctx.Invoke("google-native:bigqueryreservation/v1beta1:getReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReservationArgs struct {
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
	ReservationId string  `pulumi:"reservationId"`
}

type LookupReservationResult struct {
	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency string `pulumi:"concurrency"`
	// Creation time of the reservation.
	CreationTime string `pulumi:"creationTime"`
	// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
	IgnoreIdleSlots bool `pulumi:"ignoreIdleSlots"`
	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary bool `pulumi:"multiRegionAuxiliary"`
	// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
	Name string `pulumi:"name"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceeds the project's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the project's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
	SlotCapacity string `pulumi:"slotCapacity"`
	// Last update time of the reservation.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupReservationOutput(ctx *pulumi.Context, args LookupReservationOutputArgs, opts ...pulumi.InvokeOption) LookupReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReservationResult, error) {
			args := v.(LookupReservationArgs)
			r, err := LookupReservation(ctx, &args, opts...)
			var s LookupReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReservationResultOutput)
}

type LookupReservationOutputArgs struct {
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	ReservationId pulumi.StringInput    `pulumi:"reservationId"`
}

func (LookupReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReservationArgs)(nil)).Elem()
}

type LookupReservationResultOutput struct{ *pulumi.OutputState }

func (LookupReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReservationResult)(nil)).Elem()
}

func (o LookupReservationResultOutput) ToLookupReservationResultOutput() LookupReservationResultOutput {
	return o
}

func (o LookupReservationResultOutput) ToLookupReservationResultOutputWithContext(ctx context.Context) LookupReservationResultOutput {
	return o
}

// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
func (o LookupReservationResultOutput) Concurrency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReservationResult) string { return v.Concurrency }).(pulumi.StringOutput)
}

// Creation time of the reservation.
func (o LookupReservationResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReservationResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
func (o LookupReservationResultOutput) IgnoreIdleSlots() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReservationResult) bool { return v.IgnoreIdleSlots }).(pulumi.BoolOutput)
}

// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
func (o LookupReservationResultOutput) MultiRegionAuxiliary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReservationResult) bool { return v.MultiRegionAuxiliary }).(pulumi.BoolOutput)
}

// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
func (o LookupReservationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReservationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceeds the project's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the project's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
func (o LookupReservationResultOutput) SlotCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReservationResult) string { return v.SlotCapacity }).(pulumi.StringOutput)
}

// Last update time of the reservation.
func (o LookupReservationResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReservationResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReservationResultOutput{})
}
