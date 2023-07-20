// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an InboundSamlSsoProfile for a customer.
// Auto-naming is currently not supported for this resource.
type InboundSamlSsoProfile struct {
	pulumi.CustomResourceState

	// Immutable. The customer. For example: `customers/C0123abc`.
	Customer pulumi.StringOutput `pulumi:"customer"`
	// Human-readable name of the SAML SSO profile.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// SAML identity provider configuration.
	IdpConfig SamlIdpConfigResponseOutput `pulumi:"idpConfig"`
	// [Resource name](https://cloud.google.com/apis/design/resource_names) of the SAML SSO profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
	SpConfig SamlSpConfigResponseOutput `pulumi:"spConfig"`
}

// NewInboundSamlSsoProfile registers a new resource with the given unique name, arguments, and options.
func NewInboundSamlSsoProfile(ctx *pulumi.Context,
	name string, args *InboundSamlSsoProfileArgs, opts ...pulumi.ResourceOption) (*InboundSamlSsoProfile, error) {
	if args == nil {
		args = &InboundSamlSsoProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InboundSamlSsoProfile
	err := ctx.RegisterResource("google-native:cloudidentity/v1:InboundSamlSsoProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundSamlSsoProfile gets an existing InboundSamlSsoProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundSamlSsoProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundSamlSsoProfileState, opts ...pulumi.ResourceOption) (*InboundSamlSsoProfile, error) {
	var resource InboundSamlSsoProfile
	err := ctx.ReadResource("google-native:cloudidentity/v1:InboundSamlSsoProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundSamlSsoProfile resources.
type inboundSamlSsoProfileState struct {
}

type InboundSamlSsoProfileState struct {
}

func (InboundSamlSsoProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlSsoProfileState)(nil)).Elem()
}

type inboundSamlSsoProfileArgs struct {
	// Immutable. The customer. For example: `customers/C0123abc`.
	Customer *string `pulumi:"customer"`
	// Human-readable name of the SAML SSO profile.
	DisplayName *string `pulumi:"displayName"`
	// SAML identity provider configuration.
	IdpConfig *SamlIdpConfig `pulumi:"idpConfig"`
	// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
	SpConfig *SamlSpConfig `pulumi:"spConfig"`
}

// The set of arguments for constructing a InboundSamlSsoProfile resource.
type InboundSamlSsoProfileArgs struct {
	// Immutable. The customer. For example: `customers/C0123abc`.
	Customer pulumi.StringPtrInput
	// Human-readable name of the SAML SSO profile.
	DisplayName pulumi.StringPtrInput
	// SAML identity provider configuration.
	IdpConfig SamlIdpConfigPtrInput
	// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
	SpConfig SamlSpConfigPtrInput
}

func (InboundSamlSsoProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundSamlSsoProfileArgs)(nil)).Elem()
}

type InboundSamlSsoProfileInput interface {
	pulumi.Input

	ToInboundSamlSsoProfileOutput() InboundSamlSsoProfileOutput
	ToInboundSamlSsoProfileOutputWithContext(ctx context.Context) InboundSamlSsoProfileOutput
}

func (*InboundSamlSsoProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlSsoProfile)(nil)).Elem()
}

func (i *InboundSamlSsoProfile) ToInboundSamlSsoProfileOutput() InboundSamlSsoProfileOutput {
	return i.ToInboundSamlSsoProfileOutputWithContext(context.Background())
}

func (i *InboundSamlSsoProfile) ToInboundSamlSsoProfileOutputWithContext(ctx context.Context) InboundSamlSsoProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundSamlSsoProfileOutput)
}

type InboundSamlSsoProfileOutput struct{ *pulumi.OutputState }

func (InboundSamlSsoProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundSamlSsoProfile)(nil)).Elem()
}

func (o InboundSamlSsoProfileOutput) ToInboundSamlSsoProfileOutput() InboundSamlSsoProfileOutput {
	return o
}

func (o InboundSamlSsoProfileOutput) ToInboundSamlSsoProfileOutputWithContext(ctx context.Context) InboundSamlSsoProfileOutput {
	return o
}

// Immutable. The customer. For example: `customers/C0123abc`.
func (o InboundSamlSsoProfileOutput) Customer() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundSamlSsoProfile) pulumi.StringOutput { return v.Customer }).(pulumi.StringOutput)
}

// Human-readable name of the SAML SSO profile.
func (o InboundSamlSsoProfileOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundSamlSsoProfile) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// SAML identity provider configuration.
func (o InboundSamlSsoProfileOutput) IdpConfig() SamlIdpConfigResponseOutput {
	return o.ApplyT(func(v *InboundSamlSsoProfile) SamlIdpConfigResponseOutput { return v.IdpConfig }).(SamlIdpConfigResponseOutput)
}

// [Resource name](https://cloud.google.com/apis/design/resource_names) of the SAML SSO profile.
func (o InboundSamlSsoProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundSamlSsoProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SAML service provider configuration for this SAML SSO profile. These are the service provider details provided by Google that should be configured on the corresponding identity provider.
func (o InboundSamlSsoProfileOutput) SpConfig() SamlSpConfigResponseOutput {
	return o.ApplyT(func(v *InboundSamlSsoProfile) SamlSpConfigResponseOutput { return v.SpConfig }).(SamlSpConfigResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InboundSamlSsoProfileInput)(nil)).Elem(), &InboundSamlSsoProfile{})
	pulumi.RegisterOutputType(InboundSamlSsoProfileOutput{})
}
