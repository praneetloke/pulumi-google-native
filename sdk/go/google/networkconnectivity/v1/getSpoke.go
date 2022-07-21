// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details about a Network Connectivity Center spoke.
func LookupSpoke(ctx *pulumi.Context, args *LookupSpokeArgs, opts ...pulumi.InvokeOption) (*LookupSpokeResult, error) {
	var rv LookupSpokeResult
	err := ctx.Invoke("google-native:networkconnectivity/v1:getSpoke", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpokeArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	SpokeId  string  `pulumi:"spokeId"`
}

type LookupSpokeResult struct {
	// The time the spoke was created.
	CreateTime string `pulumi:"createTime"`
	// An optional description of the spoke.
	Description string `pulumi:"description"`
	// Immutable. The name of the hub that this spoke is attached to.
	Hub string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `pulumi:"labels"`
	// VLAN attachments that are associated with the spoke.
	LinkedInterconnectAttachments LinkedInterconnectAttachmentsResponse `pulumi:"linkedInterconnectAttachments"`
	// Router appliance instances that are associated with the spoke.
	LinkedRouterApplianceInstances LinkedRouterApplianceInstancesResponse `pulumi:"linkedRouterApplianceInstances"`
	// VPN tunnels that are associated with the spoke.
	LinkedVpnTunnels LinkedVpnTunnelsResponse `pulumi:"linkedVpnTunnels"`
	// Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
	Name string `pulumi:"name"`
	// The current lifecycle state of this spoke.
	State string `pulumi:"state"`
	// The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId string `pulumi:"uniqueId"`
	// The time the spoke was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupSpokeOutput(ctx *pulumi.Context, args LookupSpokeOutputArgs, opts ...pulumi.InvokeOption) LookupSpokeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpokeResult, error) {
			args := v.(LookupSpokeArgs)
			r, err := LookupSpoke(ctx, &args, opts...)
			var s LookupSpokeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpokeResultOutput)
}

type LookupSpokeOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	SpokeId  pulumi.StringInput    `pulumi:"spokeId"`
}

func (LookupSpokeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpokeArgs)(nil)).Elem()
}

type LookupSpokeResultOutput struct{ *pulumi.OutputState }

func (LookupSpokeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpokeResult)(nil)).Elem()
}

func (o LookupSpokeResultOutput) ToLookupSpokeResultOutput() LookupSpokeResultOutput {
	return o
}

func (o LookupSpokeResultOutput) ToLookupSpokeResultOutputWithContext(ctx context.Context) LookupSpokeResultOutput {
	return o
}

// The time the spoke was created.
func (o LookupSpokeResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// An optional description of the spoke.
func (o LookupSpokeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.Description }).(pulumi.StringOutput)
}

// Immutable. The name of the hub that this spoke is attached to.
func (o LookupSpokeResultOutput) Hub() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.Hub }).(pulumi.StringOutput)
}

// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
func (o LookupSpokeResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSpokeResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// VLAN attachments that are associated with the spoke.
func (o LookupSpokeResultOutput) LinkedInterconnectAttachments() LinkedInterconnectAttachmentsResponseOutput {
	return o.ApplyT(func(v LookupSpokeResult) LinkedInterconnectAttachmentsResponse {
		return v.LinkedInterconnectAttachments
	}).(LinkedInterconnectAttachmentsResponseOutput)
}

// Router appliance instances that are associated with the spoke.
func (o LookupSpokeResultOutput) LinkedRouterApplianceInstances() LinkedRouterApplianceInstancesResponseOutput {
	return o.ApplyT(func(v LookupSpokeResult) LinkedRouterApplianceInstancesResponse {
		return v.LinkedRouterApplianceInstances
	}).(LinkedRouterApplianceInstancesResponseOutput)
}

// VPN tunnels that are associated with the spoke.
func (o LookupSpokeResultOutput) LinkedVpnTunnels() LinkedVpnTunnelsResponseOutput {
	return o.ApplyT(func(v LookupSpokeResult) LinkedVpnTunnelsResponse { return v.LinkedVpnTunnels }).(LinkedVpnTunnelsResponseOutput)
}

// Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
func (o LookupSpokeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current lifecycle state of this spoke.
func (o LookupSpokeResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.State }).(pulumi.StringOutput)
}

// The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
func (o LookupSpokeResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

// The time the spoke was last updated.
func (o LookupSpokeResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpokeResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpokeResultOutput{})
}
