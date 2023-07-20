// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new reservation. For more information, read Reserving zonal resources.
type Reservation struct {
	pulumi.CustomResourceState

	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringOutput `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#reservations for reservations.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
	ResourcePolicies pulumi.StringMapOutput `pulumi:"resourcePolicies"`
	// Status information for Reservation resource.
	ResourceStatus AllocationResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
	ShareSettings ShareSettingsResponseOutput `pulumi:"shareSettings"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation AllocationSpecificSKUReservationResponseOutput `pulumi:"specificReservation"`
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired pulumi.BoolOutput `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status pulumi.StringOutput `pulumi:"status"`
	Zone   pulumi.StringOutput `pulumi:"zone"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil {
		args = &ReservationArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"zone",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Reservation
	err := ctx.RegisterResource("google-native:compute/v1:Reservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservation gets an existing Reservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservationState, opts ...pulumi.ResourceOption) (*Reservation, error) {
	var resource Reservation
	err := ctx.ReadResource("google-native:compute/v1:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
}

type ReservationState struct {
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
	ResourcePolicies map[string]string `pulumi:"resourcePolicies"`
	// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
	ShareSettings *ShareSettings `pulumi:"shareSettings"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation *AllocationSpecificSKUReservation `pulumi:"specificReservation"`
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
	ResourcePolicies pulumi.StringMapInput
	// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
	ShareSettings ShareSettingsPtrInput
	// Reservation for instances with specific machine shapes.
	SpecificReservation AllocationSpecificSKUReservationPtrInput
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired pulumi.BoolPtrInput
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone pulumi.StringPtrInput
}

func (ReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationArgs)(nil)).Elem()
}

type ReservationInput interface {
	pulumi.Input

	ToReservationOutput() ReservationOutput
	ToReservationOutputWithContext(ctx context.Context) ReservationOutput
}

func (*Reservation) ElementType() reflect.Type {
	return reflect.TypeOf((**Reservation)(nil)).Elem()
}

func (i *Reservation) ToReservationOutput() ReservationOutput {
	return i.ToReservationOutputWithContext(context.Background())
}

func (i *Reservation) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationOutput)
}

type ReservationOutput struct{ *pulumi.OutputState }

func (ReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Reservation)(nil)).Elem()
}

func (o ReservationOutput) ToReservationOutput() ReservationOutput {
	return o
}

func (o ReservationOutput) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return o
}

// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
func (o ReservationOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Commitment }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ReservationOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o ReservationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#reservations for reservations.
func (o ReservationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o ReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReservationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o ReservationOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
func (o ReservationOutput) ResourcePolicies() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringMapOutput { return v.ResourcePolicies }).(pulumi.StringMapOutput)
}

// Status information for Reservation resource.
func (o ReservationOutput) ResourceStatus() AllocationResourceStatusResponseOutput {
	return o.ApplyT(func(v *Reservation) AllocationResourceStatusResponseOutput { return v.ResourceStatus }).(AllocationResourceStatusResponseOutput)
}

// Reserved for future use.
func (o ReservationOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v *Reservation) pulumi.BoolOutput { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o ReservationOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
func (o ReservationOutput) ShareSettings() ShareSettingsResponseOutput {
	return o.ApplyT(func(v *Reservation) ShareSettingsResponseOutput { return v.ShareSettings }).(ShareSettingsResponseOutput)
}

// Reservation for instances with specific machine shapes.
func (o ReservationOutput) SpecificReservation() AllocationSpecificSKUReservationResponseOutput {
	return o.ApplyT(func(v *Reservation) AllocationSpecificSKUReservationResponseOutput { return v.SpecificReservation }).(AllocationSpecificSKUReservationResponseOutput)
}

// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
func (o ReservationOutput) SpecificReservationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *Reservation) pulumi.BoolOutput { return v.SpecificReservationRequired }).(pulumi.BoolOutput)
}

// The status of the reservation.
func (o ReservationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ReservationOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Reservation) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReservationInput)(nil)).Elem(), &Reservation{})
	pulumi.RegisterOutputType(ReservationOutput{})
}
