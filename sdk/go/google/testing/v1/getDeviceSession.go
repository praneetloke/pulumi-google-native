// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GET /v1/projects/{project_id}/deviceSessions/{device_session_id} Return a DeviceSession, which documents the allocation status and whether the device is allocated. Clients making requests from this API must poll GetDeviceSession.
func LookupDeviceSession(ctx *pulumi.Context, args *LookupDeviceSessionArgs, opts ...pulumi.InvokeOption) (*LookupDeviceSessionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeviceSessionResult
	err := ctx.Invoke("google-native:testing/v1:getDeviceSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceSessionArgs struct {
	DeviceSessionId string  `pulumi:"deviceSessionId"`
	Project         *string `pulumi:"project"`
}

type LookupDeviceSessionResult struct {
	// The timestamp that the session first became ACTIVE.
	ActiveStartTime string `pulumi:"activeStartTime"`
	// The requested device
	AndroidDevice AndroidDeviceResponse `pulumi:"androidDevice"`
	// The time that the Session was created.
	CreateTime string `pulumi:"createTime"`
	// The title of the DeviceSession to be presented in the UI.
	DisplayName string `pulumi:"displayName"`
	// Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
	ExpireTime string `pulumi:"expireTime"`
	// The interval of time that this device must be interacted with before it transitions from ACTIVE to TIMEOUT_INACTIVITY.
	InactivityTimeout string `pulumi:"inactivityTimeout"`
	// Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
	Name string `pulumi:"name"`
	// Current state of the DeviceSession.
	State string `pulumi:"state"`
	// The historical state transitions of the session_state message including the current session state.
	StateHistories []SessionStateEventResponse `pulumi:"stateHistories"`
	// Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
	Ttl string `pulumi:"ttl"`
}

func LookupDeviceSessionOutput(ctx *pulumi.Context, args LookupDeviceSessionOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceSessionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceSessionResult, error) {
			args := v.(LookupDeviceSessionArgs)
			r, err := LookupDeviceSession(ctx, &args, opts...)
			var s LookupDeviceSessionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceSessionResultOutput)
}

type LookupDeviceSessionOutputArgs struct {
	DeviceSessionId pulumi.StringInput    `pulumi:"deviceSessionId"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDeviceSessionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceSessionArgs)(nil)).Elem()
}

type LookupDeviceSessionResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceSessionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceSessionResult)(nil)).Elem()
}

func (o LookupDeviceSessionResultOutput) ToLookupDeviceSessionResultOutput() LookupDeviceSessionResultOutput {
	return o
}

func (o LookupDeviceSessionResultOutput) ToLookupDeviceSessionResultOutputWithContext(ctx context.Context) LookupDeviceSessionResultOutput {
	return o
}

// The timestamp that the session first became ACTIVE.
func (o LookupDeviceSessionResultOutput) ActiveStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.ActiveStartTime }).(pulumi.StringOutput)
}

// The requested device
func (o LookupDeviceSessionResultOutput) AndroidDevice() AndroidDeviceResponseOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) AndroidDeviceResponse { return v.AndroidDevice }).(AndroidDeviceResponseOutput)
}

// The time that the Session was created.
func (o LookupDeviceSessionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The title of the DeviceSession to be presented in the UI.
func (o LookupDeviceSessionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
func (o LookupDeviceSessionResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// The interval of time that this device must be interacted with before it transitions from ACTIVE to TIMEOUT_INACTIVITY.
func (o LookupDeviceSessionResultOutput) InactivityTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.InactivityTimeout }).(pulumi.StringOutput)
}

// Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
func (o LookupDeviceSessionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current state of the DeviceSession.
func (o LookupDeviceSessionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.State }).(pulumi.StringOutput)
}

// The historical state transitions of the session_state message including the current session state.
func (o LookupDeviceSessionResultOutput) StateHistories() SessionStateEventResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) []SessionStateEventResponse { return v.StateHistories }).(SessionStateEventResponseArrayOutput)
}

// Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
func (o LookupDeviceSessionResultOutput) Ttl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSessionResult) string { return v.Ttl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceSessionResultOutput{})
}