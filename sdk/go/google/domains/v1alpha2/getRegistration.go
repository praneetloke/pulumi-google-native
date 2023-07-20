// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a `Registration` resource.
func LookupRegistration(ctx *pulumi.Context, args *LookupRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistrationResult
	err := ctx.Invoke("google-native:domains/v1alpha2:getRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationArgs struct {
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
	RegistrationId string  `pulumi:"registrationId"`
}

type LookupRegistrationResult struct {
	// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsResponse `pulumi:"contactSettings"`
	// The creation timestamp of the `Registration` resource.
	CreateTime string `pulumi:"createTime"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsResponse `pulumi:"dnsSettings"`
	// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName string `pulumi:"domainName"`
	// The expiration timestamp of the `Registration`.
	ExpireTime string `pulumi:"expireTime"`
	// The set of issues with the `Registration` that require attention.
	Issues []string `pulumi:"issues"`
	// Set of labels associated with the `Registration`.
	Labels map[string]string `pulumi:"labels"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsResponse `pulumi:"managementSettings"`
	// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
	Name string `pulumi:"name"`
	// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
	PendingContactSettings ContactSettingsResponse `pulumi:"pendingContactSettings"`
	// The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
	RegisterFailureReason string `pulumi:"registerFailureReason"`
	// The state of the `Registration`
	State string `pulumi:"state"`
	// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
	SupportedPrivacy []string `pulumi:"supportedPrivacy"`
	// The reason the domain transfer failed. Only set for domains in TRANSFER_FAILED state.
	TransferFailureReason string `pulumi:"transferFailureReason"`
}

func LookupRegistrationOutput(ctx *pulumi.Context, args LookupRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationResult, error) {
			args := v.(LookupRegistrationArgs)
			r, err := LookupRegistration(ctx, &args, opts...)
			var s LookupRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationResultOutput)
}

type LookupRegistrationOutputArgs struct {
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
	RegistrationId pulumi.StringInput    `pulumi:"registrationId"`
}

func (LookupRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationArgs)(nil)).Elem()
}

type LookupRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationResult)(nil)).Elem()
}

func (o LookupRegistrationResultOutput) ToLookupRegistrationResultOutput() LookupRegistrationResultOutput {
	return o
}

func (o LookupRegistrationResultOutput) ToLookupRegistrationResultOutputWithContext(ctx context.Context) LookupRegistrationResultOutput {
	return o
}

// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
func (o LookupRegistrationResultOutput) ContactSettings() ContactSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegistrationResult) ContactSettingsResponse { return v.ContactSettings }).(ContactSettingsResponseOutput)
}

// The creation timestamp of the `Registration` resource.
func (o LookupRegistrationResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
func (o LookupRegistrationResultOutput) DnsSettings() DnsSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegistrationResult) DnsSettingsResponse { return v.DnsSettings }).(DnsSettingsResponseOutput)
}

// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
func (o LookupRegistrationResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The expiration timestamp of the `Registration`.
func (o LookupRegistrationResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// The set of issues with the `Registration` that require attention.
func (o LookupRegistrationResultOutput) Issues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegistrationResult) []string { return v.Issues }).(pulumi.StringArrayOutput)
}

// Set of labels associated with the `Registration`.
func (o LookupRegistrationResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistrationResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
func (o LookupRegistrationResultOutput) ManagementSettings() ManagementSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegistrationResult) ManagementSettingsResponse { return v.ManagementSettings }).(ManagementSettingsResponseOutput)
}

// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
func (o LookupRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
func (o LookupRegistrationResultOutput) PendingContactSettings() ContactSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegistrationResult) ContactSettingsResponse { return v.PendingContactSettings }).(ContactSettingsResponseOutput)
}

// The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
func (o LookupRegistrationResultOutput) RegisterFailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.RegisterFailureReason }).(pulumi.StringOutput)
}

// The state of the `Registration`
func (o LookupRegistrationResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.State }).(pulumi.StringOutput)
}

// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
func (o LookupRegistrationResultOutput) SupportedPrivacy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegistrationResult) []string { return v.SupportedPrivacy }).(pulumi.StringArrayOutput)
}

// The reason the domain transfer failed. Only set for domains in TRANSFER_FAILED state.
func (o LookupRegistrationResultOutput) TransferFailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.TransferFailureReason }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationResultOutput{})
}
