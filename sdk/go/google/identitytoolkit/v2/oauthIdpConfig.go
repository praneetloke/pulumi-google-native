// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an Oidc Idp configuration for an Identity Toolkit project.
type OauthIdpConfig struct {
	pulumi.CustomResourceState

	// The client id of an OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The config's display name set by developers.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// True if allows the user to sign in with the provider.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id to use for this config.
	OauthIdpConfigId pulumi.StringPtrOutput `pulumi:"oauthIdpConfigId"`
	Project          pulumi.StringOutput    `pulumi:"project"`
	// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
	ResponseType GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponseOutput `pulumi:"responseType"`
	TenantId     pulumi.StringOutput                                              `pulumi:"tenantId"`
}

// NewOauthIdpConfig registers a new resource with the given unique name, arguments, and options.
func NewOauthIdpConfig(ctx *pulumi.Context,
	name string, args *OauthIdpConfigArgs, opts ...pulumi.ResourceOption) (*OauthIdpConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"tenantId",
	})
	opts = append(opts, replaceOnChanges)
	var resource OauthIdpConfig
	err := ctx.RegisterResource("google-native:identitytoolkit/v2:OauthIdpConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOauthIdpConfig gets an existing OauthIdpConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOauthIdpConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OauthIdpConfigState, opts ...pulumi.ResourceOption) (*OauthIdpConfig, error) {
	var resource OauthIdpConfig
	err := ctx.ReadResource("google-native:identitytoolkit/v2:OauthIdpConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OauthIdpConfig resources.
type oauthIdpConfigState struct {
}

type OauthIdpConfigState struct {
}

func (OauthIdpConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthIdpConfigState)(nil)).Elem()
}

type oauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId *string `pulumi:"clientId"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret *string `pulumi:"clientSecret"`
	// The config's display name set by developers.
	DisplayName *string `pulumi:"displayName"`
	// True if allows the user to sign in with the provider.
	Enabled *bool `pulumi:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer *string `pulumi:"issuer"`
	// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
	Name *string `pulumi:"name"`
	// The id to use for this config.
	OauthIdpConfigId *string `pulumi:"oauthIdpConfigId"`
	Project          *string `pulumi:"project"`
	// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
	ResponseType *GoogleCloudIdentitytoolkitAdminV2OAuthResponseType `pulumi:"responseType"`
	TenantId     string                                              `pulumi:"tenantId"`
}

// The set of arguments for constructing a OauthIdpConfig resource.
type OauthIdpConfigArgs struct {
	// The client id of an OAuth client.
	ClientId pulumi.StringPtrInput
	// The client secret of the OAuth client, to enable OIDC code flow.
	ClientSecret pulumi.StringPtrInput
	// The config's display name set by developers.
	DisplayName pulumi.StringPtrInput
	// True if allows the user to sign in with the provider.
	Enabled pulumi.BoolPtrInput
	// For OIDC Idps, the issuer identifier.
	Issuer pulumi.StringPtrInput
	// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
	Name pulumi.StringPtrInput
	// The id to use for this config.
	OauthIdpConfigId pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
	ResponseType GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypePtrInput
	TenantId     pulumi.StringInput
}

func (OauthIdpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthIdpConfigArgs)(nil)).Elem()
}

type OauthIdpConfigInput interface {
	pulumi.Input

	ToOauthIdpConfigOutput() OauthIdpConfigOutput
	ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput
}

func (*OauthIdpConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthIdpConfig)(nil)).Elem()
}

func (i *OauthIdpConfig) ToOauthIdpConfigOutput() OauthIdpConfigOutput {
	return i.ToOauthIdpConfigOutputWithContext(context.Background())
}

func (i *OauthIdpConfig) ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthIdpConfigOutput)
}

type OauthIdpConfigOutput struct{ *pulumi.OutputState }

func (OauthIdpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthIdpConfig)(nil)).Elem()
}

func (o OauthIdpConfigOutput) ToOauthIdpConfigOutput() OauthIdpConfigOutput {
	return o
}

func (o OauthIdpConfigOutput) ToOauthIdpConfigOutputWithContext(ctx context.Context) OauthIdpConfigOutput {
	return o
}

// The client id of an OAuth client.
func (o OauthIdpConfigOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the OAuth client, to enable OIDC code flow.
func (o OauthIdpConfigOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// The config's display name set by developers.
func (o OauthIdpConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// True if allows the user to sign in with the provider.
func (o OauthIdpConfigOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// For OIDC Idps, the issuer identifier.
func (o OauthIdpConfigOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
func (o OauthIdpConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id to use for this config.
func (o OauthIdpConfigOutput) OauthIdpConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringPtrOutput { return v.OauthIdpConfigId }).(pulumi.StringPtrOutput)
}

func (o OauthIdpConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
func (o OauthIdpConfigOutput) ResponseType() GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponseOutput {
	return o.ApplyT(func(v *OauthIdpConfig) GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponseOutput {
		return v.ResponseType
	}).(GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponseOutput)
}

func (o OauthIdpConfigOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OauthIdpConfig) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OauthIdpConfigInput)(nil)).Elem(), &OauthIdpConfig{})
	pulumi.RegisterOutputType(OauthIdpConfigOutput{})
}