// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a tenant. Requires read permission on the Tenant resource.
func LookupTenant(ctx *pulumi.Context, args *LookupTenantArgs, opts ...pulumi.InvokeOption) (*LookupTenantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTenantResult
	err := ctx.Invoke("google-native:identitytoolkit/v2:getTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTenantArgs struct {
	Project  *string `pulumi:"project"`
	TenantId string  `pulumi:"tenantId"`
}

type LookupTenantResult struct {
	// Whether to allow email/password user authentication.
	AllowPasswordSignup bool `pulumi:"allowPasswordSignup"`
	// Whether anonymous users will be auto-deleted after a period of 30 days.
	AutodeleteAnonymousUsers bool `pulumi:"autodeleteAnonymousUsers"`
	// Options related to how clients making requests on behalf of a project should be configured.
	Client GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponse `pulumi:"client"`
	// Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
	DisableAuth bool `pulumi:"disableAuth"`
	// Display name of the tenant.
	DisplayName string `pulumi:"displayName"`
	// Configuration for settings related to email privacy and public visibility.
	EmailPrivacyConfig GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponse `pulumi:"emailPrivacyConfig"`
	// Whether to enable anonymous user authentication.
	EnableAnonymousUser bool `pulumi:"enableAnonymousUser"`
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin bool `pulumi:"enableEmailLinkSignin"`
	// Hash config information of a tenant for display on Pantheon. This can only be displayed on Pantheon to avoid the sensitive information to get accidentally leaked. Only returned in GetTenant response to restrict reading of this information. Requires firebaseauth.configs.getHashConfig permission on the agent project for returning this field.
	HashConfig GoogleCloudIdentitytoolkitAdminV2HashConfigResponse `pulumi:"hashConfig"`
	// Specify the settings that the tenant could inherit.
	Inheritance GoogleCloudIdentitytoolkitAdminV2InheritanceResponse `pulumi:"inheritance"`
	// The tenant-level configuration of MFA options.
	MfaConfig GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponse `pulumi:"mfaConfig"`
	// Configuration related to monitoring project activity.
	Monitoring GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponse `pulumi:"monitoring"`
	// Resource name of a tenant. For example: "projects/{project-id}/tenants/{tenant-id}"
	Name string `pulumi:"name"`
	// The tenant-level password policy config
	PasswordPolicyConfig GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponse `pulumi:"passwordPolicyConfig"`
	// The tenant-level reCAPTCHA config.
	RecaptchaConfig GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponse `pulumi:"recaptchaConfig"`
	// Configures which regions are enabled for SMS verification code sending.
	SmsRegionConfig GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponse `pulumi:"smsRegionConfig"`
	// A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
	TestPhoneNumbers map[string]string `pulumi:"testPhoneNumbers"`
}

func LookupTenantOutput(ctx *pulumi.Context, args LookupTenantOutputArgs, opts ...pulumi.InvokeOption) LookupTenantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTenantResult, error) {
			args := v.(LookupTenantArgs)
			r, err := LookupTenant(ctx, &args, opts...)
			var s LookupTenantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTenantResultOutput)
}

type LookupTenantOutputArgs struct {
	Project  pulumi.StringPtrInput `pulumi:"project"`
	TenantId pulumi.StringInput    `pulumi:"tenantId"`
}

func (LookupTenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantArgs)(nil)).Elem()
}

type LookupTenantResultOutput struct{ *pulumi.OutputState }

func (LookupTenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantResult)(nil)).Elem()
}

func (o LookupTenantResultOutput) ToLookupTenantResultOutput() LookupTenantResultOutput {
	return o
}

func (o LookupTenantResultOutput) ToLookupTenantResultOutputWithContext(ctx context.Context) LookupTenantResultOutput {
	return o
}

// Whether to allow email/password user authentication.
func (o LookupTenantResultOutput) AllowPasswordSignup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantResult) bool { return v.AllowPasswordSignup }).(pulumi.BoolOutput)
}

// Whether anonymous users will be auto-deleted after a period of 30 days.
func (o LookupTenantResultOutput) AutodeleteAnonymousUsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantResult) bool { return v.AutodeleteAnonymousUsers }).(pulumi.BoolOutput)
}

// Options related to how clients making requests on behalf of a project should be configured.
func (o LookupTenantResultOutput) Client() GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponse {
		return v.Client
	}).(GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponseOutput)
}

// Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
func (o LookupTenantResultOutput) DisableAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantResult) bool { return v.DisableAuth }).(pulumi.BoolOutput)
}

// Display name of the tenant.
func (o LookupTenantResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Configuration for settings related to email privacy and public visibility.
func (o LookupTenantResultOutput) EmailPrivacyConfig() GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponse {
		return v.EmailPrivacyConfig
	}).(GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponseOutput)
}

// Whether to enable anonymous user authentication.
func (o LookupTenantResultOutput) EnableAnonymousUser() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantResult) bool { return v.EnableAnonymousUser }).(pulumi.BoolOutput)
}

// Whether to enable email link user authentication.
func (o LookupTenantResultOutput) EnableEmailLinkSignin() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantResult) bool { return v.EnableEmailLinkSignin }).(pulumi.BoolOutput)
}

// Hash config information of a tenant for display on Pantheon. This can only be displayed on Pantheon to avoid the sensitive information to get accidentally leaked. Only returned in GetTenant response to restrict reading of this information. Requires firebaseauth.configs.getHashConfig permission on the agent project for returning this field.
func (o LookupTenantResultOutput) HashConfig() GoogleCloudIdentitytoolkitAdminV2HashConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2HashConfigResponse { return v.HashConfig }).(GoogleCloudIdentitytoolkitAdminV2HashConfigResponseOutput)
}

// Specify the settings that the tenant could inherit.
func (o LookupTenantResultOutput) Inheritance() GoogleCloudIdentitytoolkitAdminV2InheritanceResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2InheritanceResponse { return v.Inheritance }).(GoogleCloudIdentitytoolkitAdminV2InheritanceResponseOutput)
}

// The tenant-level configuration of MFA options.
func (o LookupTenantResultOutput) MfaConfig() GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponse {
		return v.MfaConfig
	}).(GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponseOutput)
}

// Configuration related to monitoring project activity.
func (o LookupTenantResultOutput) Monitoring() GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponse {
		return v.Monitoring
	}).(GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponseOutput)
}

// Resource name of a tenant. For example: "projects/{project-id}/tenants/{tenant-id}"
func (o LookupTenantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantResult) string { return v.Name }).(pulumi.StringOutput)
}

// The tenant-level password policy config
func (o LookupTenantResultOutput) PasswordPolicyConfig() GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponse {
		return v.PasswordPolicyConfig
	}).(GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponseOutput)
}

// The tenant-level reCAPTCHA config.
func (o LookupTenantResultOutput) RecaptchaConfig() GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponse {
		return v.RecaptchaConfig
	}).(GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponseOutput)
}

// Configures which regions are enabled for SMS verification code sending.
func (o LookupTenantResultOutput) SmsRegionConfig() GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponseOutput {
	return o.ApplyT(func(v LookupTenantResult) GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponse {
		return v.SmsRegionConfig
	}).(GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponseOutput)
}

// A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
func (o LookupTenantResultOutput) TestPhoneNumbers() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTenantResult) map[string]string { return v.TestPhoneNumbers }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantResultOutput{})
}
