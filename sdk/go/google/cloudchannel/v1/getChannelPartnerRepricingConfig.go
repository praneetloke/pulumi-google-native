// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about how a Distributor modifies their bill before sending it to a ChannelPartner. Possible Error Codes: * PERMISSION_DENIED: If the account making the request and the account being queried are different. * NOT_FOUND: The ChannelPartnerRepricingConfig was not found. * INTERNAL: Any non-user error related to technical issues in the backend. In this case, contact Cloud Channel support. Return Value: If successful, the ChannelPartnerRepricingConfig resource, otherwise returns an error.
func LookupChannelPartnerRepricingConfig(ctx *pulumi.Context, args *LookupChannelPartnerRepricingConfigArgs, opts ...pulumi.InvokeOption) (*LookupChannelPartnerRepricingConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChannelPartnerRepricingConfigResult
	err := ctx.Invoke("google-native:cloudchannel/v1:getChannelPartnerRepricingConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelPartnerRepricingConfigArgs struct {
	AccountId                       string `pulumi:"accountId"`
	ChannelPartnerLinkId            string `pulumi:"channelPartnerLinkId"`
	ChannelPartnerRepricingConfigId string `pulumi:"channelPartnerRepricingConfigId"`
}

type LookupChannelPartnerRepricingConfigResult struct {
	// Resource name of the ChannelPartnerRepricingConfig. Format: accounts/{account_id}/channelPartnerLinks/{channel_partner_id}/channelPartnerRepricingConfigs/{id}.
	Name string `pulumi:"name"`
	// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
	RepricingConfig GoogleCloudChannelV1RepricingConfigResponse `pulumi:"repricingConfig"`
	// Timestamp of an update to the repricing rule. If `update_time` is after RepricingConfig.effective_invoice_month then it indicates this was set mid-month.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupChannelPartnerRepricingConfigOutput(ctx *pulumi.Context, args LookupChannelPartnerRepricingConfigOutputArgs, opts ...pulumi.InvokeOption) LookupChannelPartnerRepricingConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelPartnerRepricingConfigResult, error) {
			args := v.(LookupChannelPartnerRepricingConfigArgs)
			r, err := LookupChannelPartnerRepricingConfig(ctx, &args, opts...)
			var s LookupChannelPartnerRepricingConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelPartnerRepricingConfigResultOutput)
}

type LookupChannelPartnerRepricingConfigOutputArgs struct {
	AccountId                       pulumi.StringInput `pulumi:"accountId"`
	ChannelPartnerLinkId            pulumi.StringInput `pulumi:"channelPartnerLinkId"`
	ChannelPartnerRepricingConfigId pulumi.StringInput `pulumi:"channelPartnerRepricingConfigId"`
}

func (LookupChannelPartnerRepricingConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelPartnerRepricingConfigArgs)(nil)).Elem()
}

type LookupChannelPartnerRepricingConfigResultOutput struct{ *pulumi.OutputState }

func (LookupChannelPartnerRepricingConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelPartnerRepricingConfigResult)(nil)).Elem()
}

func (o LookupChannelPartnerRepricingConfigResultOutput) ToLookupChannelPartnerRepricingConfigResultOutput() LookupChannelPartnerRepricingConfigResultOutput {
	return o
}

func (o LookupChannelPartnerRepricingConfigResultOutput) ToLookupChannelPartnerRepricingConfigResultOutputWithContext(ctx context.Context) LookupChannelPartnerRepricingConfigResultOutput {
	return o
}

// Resource name of the ChannelPartnerRepricingConfig. Format: accounts/{account_id}/channelPartnerLinks/{channel_partner_id}/channelPartnerRepricingConfigs/{id}.
func (o LookupChannelPartnerRepricingConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerRepricingConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
func (o LookupChannelPartnerRepricingConfigResultOutput) RepricingConfig() GoogleCloudChannelV1RepricingConfigResponseOutput {
	return o.ApplyT(func(v LookupChannelPartnerRepricingConfigResult) GoogleCloudChannelV1RepricingConfigResponse {
		return v.RepricingConfig
	}).(GoogleCloudChannelV1RepricingConfigResponseOutput)
}

// Timestamp of an update to the repricing rule. If `update_time` is after RepricingConfig.effective_invoice_month then it indicates this was set mid-month.
func (o LookupChannelPartnerRepricingConfigResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerRepricingConfigResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelPartnerRepricingConfigResultOutput{})
}
