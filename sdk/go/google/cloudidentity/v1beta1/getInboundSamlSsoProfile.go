// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an InboundSamlSsoProfile.
func LookupInboundSamlSsoProfile(ctx *pulumi.Context, args *LookupInboundSamlSsoProfileArgs, opts ...pulumi.InvokeOption) (*LookupInboundSamlSsoProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInboundSamlSsoProfileResult
	err := ctx.Invoke("google-native:cloudidentity/v1beta1:getInboundSamlSsoProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundSamlSsoProfileArgs struct {
	InboundSamlSsoProfileId string `pulumi:"inboundSamlSsoProfileId"`
}

type LookupInboundSamlSsoProfileResult struct {
	// Immutable. The customer. For example: `customers/C0123abc`.
	Customer string `pulumi:"customer"`
	// Human-readable name of the SAML SSO profile.
	DisplayName string `pulumi:"displayName"`
	// SAML identity provider configuration.
	IdpConfig SamlIdpConfigResponse `pulumi:"idpConfig"`
	// [Resource name](https://cloud.google.com/apis/design/resource_names) of the SAML SSO profile.
	Name string `pulumi:"name"`
	// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
	SpConfig SamlSpConfigResponse `pulumi:"spConfig"`
}

func LookupInboundSamlSsoProfileOutput(ctx *pulumi.Context, args LookupInboundSamlSsoProfileOutputArgs, opts ...pulumi.InvokeOption) LookupInboundSamlSsoProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInboundSamlSsoProfileResult, error) {
			args := v.(LookupInboundSamlSsoProfileArgs)
			r, err := LookupInboundSamlSsoProfile(ctx, &args, opts...)
			var s LookupInboundSamlSsoProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInboundSamlSsoProfileResultOutput)
}

type LookupInboundSamlSsoProfileOutputArgs struct {
	InboundSamlSsoProfileId pulumi.StringInput `pulumi:"inboundSamlSsoProfileId"`
}

func (LookupInboundSamlSsoProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundSamlSsoProfileArgs)(nil)).Elem()
}

type LookupInboundSamlSsoProfileResultOutput struct{ *pulumi.OutputState }

func (LookupInboundSamlSsoProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundSamlSsoProfileResult)(nil)).Elem()
}

func (o LookupInboundSamlSsoProfileResultOutput) ToLookupInboundSamlSsoProfileResultOutput() LookupInboundSamlSsoProfileResultOutput {
	return o
}

func (o LookupInboundSamlSsoProfileResultOutput) ToLookupInboundSamlSsoProfileResultOutputWithContext(ctx context.Context) LookupInboundSamlSsoProfileResultOutput {
	return o
}

// Immutable. The customer. For example: `customers/C0123abc`.
func (o LookupInboundSamlSsoProfileResultOutput) Customer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundSamlSsoProfileResult) string { return v.Customer }).(pulumi.StringOutput)
}

// Human-readable name of the SAML SSO profile.
func (o LookupInboundSamlSsoProfileResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundSamlSsoProfileResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// SAML identity provider configuration.
func (o LookupInboundSamlSsoProfileResultOutput) IdpConfig() SamlIdpConfigResponseOutput {
	return o.ApplyT(func(v LookupInboundSamlSsoProfileResult) SamlIdpConfigResponse { return v.IdpConfig }).(SamlIdpConfigResponseOutput)
}

// [Resource name](https://cloud.google.com/apis/design/resource_names) of the SAML SSO profile.
func (o LookupInboundSamlSsoProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundSamlSsoProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
func (o LookupInboundSamlSsoProfileResultOutput) SpConfig() SamlSpConfigResponseOutput {
	return o.ApplyT(func(v LookupInboundSamlSsoProfileResult) SamlSpConfigResponse { return v.SpConfig }).(SamlSpConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundSamlSsoProfileResultOutput{})
}
