// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve an inbound SAML configuration for an Identity Toolkit project.
func LookupInboundSamlConfig(ctx *pulumi.Context, args *LookupInboundSamlConfigArgs, opts ...pulumi.InvokeOption) (*LookupInboundSamlConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInboundSamlConfigResult
	err := ctx.Invoke("google-native:identitytoolkit/v2:getInboundSamlConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundSamlConfigArgs struct {
	InboundSamlConfigId string  `pulumi:"inboundSamlConfigId"`
	Project             *string `pulumi:"project"`
	TenantId            string  `pulumi:"tenantId"`
}

type LookupInboundSamlConfigResult struct {
	// The config's display name set by developers.
	DisplayName string `pulumi:"displayName"`
	// True if allows the user to sign in with the provider.
	Enabled bool `pulumi:"enabled"`
	// The SAML IdP (Identity Provider) configuration when the project acts as the relying party.
	IdpConfig GoogleCloudIdentitytoolkitAdminV2IdpConfigResponse `pulumi:"idpConfig"`
	// The name of the InboundSamlConfig resource, for example: 'projects/my-awesome-project/inboundSamlConfigs/my-config-id'. Ignored during create requests.
	Name string `pulumi:"name"`
	// The SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an authentication assertion issued by a SAML identity provider.
	SpConfig GoogleCloudIdentitytoolkitAdminV2SpConfigResponse `pulumi:"spConfig"`
}

func LookupInboundSamlConfigOutput(ctx *pulumi.Context, args LookupInboundSamlConfigOutputArgs, opts ...pulumi.InvokeOption) LookupInboundSamlConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInboundSamlConfigResult, error) {
			args := v.(LookupInboundSamlConfigArgs)
			r, err := LookupInboundSamlConfig(ctx, &args, opts...)
			var s LookupInboundSamlConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInboundSamlConfigResultOutput)
}

type LookupInboundSamlConfigOutputArgs struct {
	InboundSamlConfigId pulumi.StringInput    `pulumi:"inboundSamlConfigId"`
	Project             pulumi.StringPtrInput `pulumi:"project"`
	TenantId            pulumi.StringInput    `pulumi:"tenantId"`
}

func (LookupInboundSamlConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundSamlConfigArgs)(nil)).Elem()
}

type LookupInboundSamlConfigResultOutput struct{ *pulumi.OutputState }

func (LookupInboundSamlConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundSamlConfigResult)(nil)).Elem()
}

func (o LookupInboundSamlConfigResultOutput) ToLookupInboundSamlConfigResultOutput() LookupInboundSamlConfigResultOutput {
	return o
}

func (o LookupInboundSamlConfigResultOutput) ToLookupInboundSamlConfigResultOutputWithContext(ctx context.Context) LookupInboundSamlConfigResultOutput {
	return o
}

// The config's display name set by developers.
func (o LookupInboundSamlConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundSamlConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// True if allows the user to sign in with the provider.
func (o LookupInboundSamlConfigResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInboundSamlConfigResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The SAML IdP (Identity Provider) configuration when the project acts as the relying party.
func (o LookupInboundSamlConfigResultOutput) IdpConfig() GoogleCloudIdentitytoolkitAdminV2IdpConfigResponseOutput {
	return o.ApplyT(func(v LookupInboundSamlConfigResult) GoogleCloudIdentitytoolkitAdminV2IdpConfigResponse {
		return v.IdpConfig
	}).(GoogleCloudIdentitytoolkitAdminV2IdpConfigResponseOutput)
}

// The name of the InboundSamlConfig resource, for example: 'projects/my-awesome-project/inboundSamlConfigs/my-config-id'. Ignored during create requests.
func (o LookupInboundSamlConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundSamlConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an authentication assertion issued by a SAML identity provider.
func (o LookupInboundSamlConfigResultOutput) SpConfig() GoogleCloudIdentitytoolkitAdminV2SpConfigResponseOutput {
	return o.ApplyT(func(v LookupInboundSamlConfigResult) GoogleCloudIdentitytoolkitAdminV2SpConfigResponse {
		return v.SpConfig
	}).(GoogleCloudIdentitytoolkitAdminV2SpConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundSamlConfigResultOutput{})
}
