// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new UtilizationReport.
// Auto-naming is currently not supported for this resource.
type UtilizationReport struct {
	pulumi.CustomResourceState

	// The time the report was created (this refers to the time of the request, not the time the report creation completed).
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The report display name, as assigned by the user.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Provides details on the state of the report in case of an error.
	Error StatusResponseOutput `pulumi:"error"`
	// The point in time when the time frame ends. Notice that the time frame is counted backwards. For instance if the "frame_end_time" value is 2021/01/20 and the time frame is WEEK then the report covers the week between 2021/01/20 and 2021/01/14.
	FrameEndTime pulumi.StringOutput `pulumi:"frameEndTime"`
	Location     pulumi.StringOutput `pulumi:"location"`
	// The report unique name.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	SourceId  pulumi.StringOutput    `pulumi:"sourceId"`
	// Current state of the report.
	State pulumi.StringOutput `pulumi:"state"`
	// The time the state was last set.
	StateTime pulumi.StringOutput `pulumi:"stateTime"`
	// Time frame of the report.
	TimeFrame pulumi.StringOutput `pulumi:"timeFrame"`
	// Required. The ID to use for the report, which will become the final component of the reports's resource name. This value maximum length is 63 characters, and valid characters are /a-z-/. It must start with an english letter and must not end with a hyphen.
	UtilizationReportId pulumi.StringOutput `pulumi:"utilizationReportId"`
	// Total number of VMs included in the report.
	VmCount pulumi.IntOutput `pulumi:"vmCount"`
	// List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
	Vms VmUtilizationInfoResponseArrayOutput `pulumi:"vms"`
	// Total number of VMs included in the report.
	VmsCount pulumi.IntOutput `pulumi:"vmsCount"`
}

// NewUtilizationReport registers a new resource with the given unique name, arguments, and options.
func NewUtilizationReport(ctx *pulumi.Context,
	name string, args *UtilizationReportArgs, opts ...pulumi.ResourceOption) (*UtilizationReport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	if args.UtilizationReportId == nil {
		return nil, errors.New("invalid value for required argument 'UtilizationReportId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"sourceId",
		"utilizationReportId",
	})
	opts = append(opts, replaceOnChanges)
	var resource UtilizationReport
	err := ctx.RegisterResource("google-native:vmmigration/v1alpha1:UtilizationReport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUtilizationReport gets an existing UtilizationReport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUtilizationReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UtilizationReportState, opts ...pulumi.ResourceOption) (*UtilizationReport, error) {
	var resource UtilizationReport
	err := ctx.ReadResource("google-native:vmmigration/v1alpha1:UtilizationReport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UtilizationReport resources.
type utilizationReportState struct {
}

type UtilizationReportState struct {
}

func (UtilizationReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*utilizationReportState)(nil)).Elem()
}

type utilizationReportArgs struct {
	// The report display name, as assigned by the user.
	DisplayName *string `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	Project     *string `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	SourceId  string  `pulumi:"sourceId"`
	// Time frame of the report.
	TimeFrame *UtilizationReportTimeFrame `pulumi:"timeFrame"`
	// Required. The ID to use for the report, which will become the final component of the reports's resource name. This value maximum length is 63 characters, and valid characters are /a-z-/. It must start with an english letter and must not end with a hyphen.
	UtilizationReportId string `pulumi:"utilizationReportId"`
	// List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
	Vms []VmUtilizationInfo `pulumi:"vms"`
}

// The set of arguments for constructing a UtilizationReport resource.
type UtilizationReportArgs struct {
	// The report display name, as assigned by the user.
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	Project     pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	SourceId  pulumi.StringInput
	// Time frame of the report.
	TimeFrame UtilizationReportTimeFramePtrInput
	// Required. The ID to use for the report, which will become the final component of the reports's resource name. This value maximum length is 63 characters, and valid characters are /a-z-/. It must start with an english letter and must not end with a hyphen.
	UtilizationReportId pulumi.StringInput
	// List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
	Vms VmUtilizationInfoArrayInput
}

func (UtilizationReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*utilizationReportArgs)(nil)).Elem()
}

type UtilizationReportInput interface {
	pulumi.Input

	ToUtilizationReportOutput() UtilizationReportOutput
	ToUtilizationReportOutputWithContext(ctx context.Context) UtilizationReportOutput
}

func (*UtilizationReport) ElementType() reflect.Type {
	return reflect.TypeOf((**UtilizationReport)(nil)).Elem()
}

func (i *UtilizationReport) ToUtilizationReportOutput() UtilizationReportOutput {
	return i.ToUtilizationReportOutputWithContext(context.Background())
}

func (i *UtilizationReport) ToUtilizationReportOutputWithContext(ctx context.Context) UtilizationReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtilizationReportOutput)
}

type UtilizationReportOutput struct{ *pulumi.OutputState }

func (UtilizationReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UtilizationReport)(nil)).Elem()
}

func (o UtilizationReportOutput) ToUtilizationReportOutput() UtilizationReportOutput {
	return o
}

func (o UtilizationReportOutput) ToUtilizationReportOutputWithContext(ctx context.Context) UtilizationReportOutput {
	return o
}

// The time the report was created (this refers to the time of the request, not the time the report creation completed).
func (o UtilizationReportOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The report display name, as assigned by the user.
func (o UtilizationReportOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Provides details on the state of the report in case of an error.
func (o UtilizationReportOutput) Error() StatusResponseOutput {
	return o.ApplyT(func(v *UtilizationReport) StatusResponseOutput { return v.Error }).(StatusResponseOutput)
}

// The point in time when the time frame ends. Notice that the time frame is counted backwards. For instance if the "frame_end_time" value is 2021/01/20 and the time frame is WEEK then the report covers the week between 2021/01/20 and 2021/01/14.
func (o UtilizationReportOutput) FrameEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.FrameEndTime }).(pulumi.StringOutput)
}

func (o UtilizationReportOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The report unique name.
func (o UtilizationReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UtilizationReportOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o UtilizationReportOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

func (o UtilizationReportOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

// Current state of the report.
func (o UtilizationReportOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The time the state was last set.
func (o UtilizationReportOutput) StateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.StateTime }).(pulumi.StringOutput)
}

// Time frame of the report.
func (o UtilizationReportOutput) TimeFrame() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.TimeFrame }).(pulumi.StringOutput)
}

// Required. The ID to use for the report, which will become the final component of the reports's resource name. This value maximum length is 63 characters, and valid characters are /a-z-/. It must start with an english letter and must not end with a hyphen.
func (o UtilizationReportOutput) UtilizationReportId() pulumi.StringOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.StringOutput { return v.UtilizationReportId }).(pulumi.StringOutput)
}

// Total number of VMs included in the report.
func (o UtilizationReportOutput) VmCount() pulumi.IntOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.IntOutput { return v.VmCount }).(pulumi.IntOutput)
}

// List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
func (o UtilizationReportOutput) Vms() VmUtilizationInfoResponseArrayOutput {
	return o.ApplyT(func(v *UtilizationReport) VmUtilizationInfoResponseArrayOutput { return v.Vms }).(VmUtilizationInfoResponseArrayOutput)
}

// Total number of VMs included in the report.
func (o UtilizationReportOutput) VmsCount() pulumi.IntOutput {
	return o.ApplyT(func(v *UtilizationReport) pulumi.IntOutput { return v.VmsCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UtilizationReportInput)(nil)).Elem(), &UtilizationReport{})
	pulumi.RegisterOutputType(UtilizationReportOutput{})
}
