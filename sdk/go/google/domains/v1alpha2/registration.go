// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Registers a new domain name and creates a corresponding `Registration` resource. Call `RetrieveRegisterParameters` first to check availability of the domain name and determine parameters like price that are needed to build a call to this method. A successful call creates a `Registration` resource in state `REGISTRATION_PENDING`, which resolves to `ACTIVE` within 1-2 minutes, indicating that the domain was successfully registered. If the resource ends up in state `REGISTRATION_FAILED`, it indicates that the domain was not registered successfully, and you can safely delete the resource and retry registration.
// Auto-naming is currently not supported for this resource.
type Registration struct {
	pulumi.CustomResourceState

	// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsResponseOutput `pulumi:"contactSettings"`
	// The creation timestamp of the `Registration` resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsResponseOutput `pulumi:"dnsSettings"`
	// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The expiration timestamp of the `Registration`.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The set of issues with the `Registration` that require attention.
	Issues pulumi.StringArrayOutput `pulumi:"issues"`
	// Set of labels associated with the `Registration`.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsResponseOutput `pulumi:"managementSettings"`
	// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
	PendingContactSettings ContactSettingsResponseOutput `pulumi:"pendingContactSettings"`
	Project                pulumi.StringOutput           `pulumi:"project"`
	// The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
	RegisterFailureReason pulumi.StringOutput `pulumi:"registerFailureReason"`
	// The state of the `Registration`
	State pulumi.StringOutput `pulumi:"state"`
	// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
	SupportedPrivacy pulumi.StringArrayOutput `pulumi:"supportedPrivacy"`
	// The reason the domain transfer failed. Only set for domains in TRANSFER_FAILED state.
	TransferFailureReason pulumi.StringOutput `pulumi:"transferFailureReason"`
}

// NewRegistration registers a new resource with the given unique name, arguments, and options.
func NewRegistration(ctx *pulumi.Context,
	name string, args *RegistrationArgs, opts ...pulumi.ResourceOption) (*Registration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactSettings == nil {
		return nil, errors.New("invalid value for required argument 'ContactSettings'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.YearlyPrice == nil {
		return nil, errors.New("invalid value for required argument 'YearlyPrice'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Registration
	err := ctx.RegisterResource("google-native:domains/v1alpha2:Registration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistration gets an existing Registration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationState, opts ...pulumi.ResourceOption) (*Registration, error) {
	var resource Registration
	err := ctx.ReadResource("google-native:domains/v1alpha2:Registration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registration resources.
type registrationState struct {
}

type RegistrationState struct {
}

func (RegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationState)(nil)).Elem()
}

type registrationArgs struct {
	// The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
	ContactNotices []RegistrationContactNoticesItem `pulumi:"contactNotices"`
	// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettings `pulumi:"contactSettings"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings *DnsSettings `pulumi:"dnsSettings"`
	// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName string `pulumi:"domainName"`
	// The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
	DomainNotices []RegistrationDomainNoticesItem `pulumi:"domainNotices"`
	// Set of labels associated with the `Registration`.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings *ManagementSettings `pulumi:"managementSettings"`
	Project            *string             `pulumi:"project"`
	// When true, only validation is performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
	ValidateOnly *bool `pulumi:"validateOnly"`
	// Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
	YearlyPrice Money `pulumi:"yearlyPrice"`
}

// The set of arguments for constructing a Registration resource.
type RegistrationArgs struct {
	// The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
	ContactNotices RegistrationContactNoticesItemArrayInput
	// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsInput
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsPtrInput
	// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName pulumi.StringInput
	// The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
	DomainNotices RegistrationDomainNoticesItemArrayInput
	// Set of labels associated with the `Registration`.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsPtrInput
	Project            pulumi.StringPtrInput
	// When true, only validation is performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
	ValidateOnly pulumi.BoolPtrInput
	// Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
	YearlyPrice MoneyInput
}

func (RegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationArgs)(nil)).Elem()
}

type RegistrationInput interface {
	pulumi.Input

	ToRegistrationOutput() RegistrationOutput
	ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput
}

func (*Registration) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (i *Registration) ToRegistrationOutput() RegistrationOutput {
	return i.ToRegistrationOutputWithContext(context.Background())
}

func (i *Registration) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationOutput)
}

type RegistrationOutput struct{ *pulumi.OutputState }

func (RegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (o RegistrationOutput) ToRegistrationOutput() RegistrationOutput {
	return o
}

func (o RegistrationOutput) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return o
}

// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
func (o RegistrationOutput) ContactSettings() ContactSettingsResponseOutput {
	return o.ApplyT(func(v *Registration) ContactSettingsResponseOutput { return v.ContactSettings }).(ContactSettingsResponseOutput)
}

// The creation timestamp of the `Registration` resource.
func (o RegistrationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
func (o RegistrationOutput) DnsSettings() DnsSettingsResponseOutput {
	return o.ApplyT(func(v *Registration) DnsSettingsResponseOutput { return v.DnsSettings }).(DnsSettingsResponseOutput)
}

// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
func (o RegistrationOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The expiration timestamp of the `Registration`.
func (o RegistrationOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// The set of issues with the `Registration` that require attention.
func (o RegistrationOutput) Issues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringArrayOutput { return v.Issues }).(pulumi.StringArrayOutput)
}

// Set of labels associated with the `Registration`.
func (o RegistrationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o RegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
func (o RegistrationOutput) ManagementSettings() ManagementSettingsResponseOutput {
	return o.ApplyT(func(v *Registration) ManagementSettingsResponseOutput { return v.ManagementSettings }).(ManagementSettingsResponseOutput)
}

// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
func (o RegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
func (o RegistrationOutput) PendingContactSettings() ContactSettingsResponseOutput {
	return o.ApplyT(func(v *Registration) ContactSettingsResponseOutput { return v.PendingContactSettings }).(ContactSettingsResponseOutput)
}

func (o RegistrationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
func (o RegistrationOutput) RegisterFailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.RegisterFailureReason }).(pulumi.StringOutput)
}

// The state of the `Registration`
func (o RegistrationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
func (o RegistrationOutput) SupportedPrivacy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringArrayOutput { return v.SupportedPrivacy }).(pulumi.StringArrayOutput)
}

// The reason the domain transfer failed. Only set for domains in TRANSFER_FAILED state.
func (o RegistrationOutput) TransferFailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.TransferFailureReason }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationInput)(nil)).Elem(), &Registration{})
	pulumi.RegisterOutputType(RegistrationOutput{})
}
