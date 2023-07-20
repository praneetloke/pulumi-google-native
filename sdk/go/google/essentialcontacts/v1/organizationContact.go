// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds a new contact for a resource.
// Auto-naming is currently not supported for this resource.
type OrganizationContact struct {
	pulumi.CustomResourceState

	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email pulumi.StringOutput `pulumi:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag pulumi.StringOutput `pulumi:"languageTag"`
	// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions pulumi.StringArrayOutput `pulumi:"notificationCategorySubscriptions"`
	OrganizationId                    pulumi.StringOutput      `pulumi:"organizationId"`
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime pulumi.StringOutput `pulumi:"validateTime"`
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState pulumi.StringOutput `pulumi:"validationState"`
}

// NewOrganizationContact registers a new resource with the given unique name, arguments, and options.
func NewOrganizationContact(ctx *pulumi.Context,
	name string, args *OrganizationContactArgs, opts ...pulumi.ResourceOption) (*OrganizationContact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.LanguageTag == nil {
		return nil, errors.New("invalid value for required argument 'LanguageTag'")
	}
	if args.NotificationCategorySubscriptions == nil {
		return nil, errors.New("invalid value for required argument 'NotificationCategorySubscriptions'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationContact
	err := ctx.RegisterResource("google-native:essentialcontacts/v1:OrganizationContact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationContact gets an existing OrganizationContact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationContactState, opts ...pulumi.ResourceOption) (*OrganizationContact, error) {
	var resource OrganizationContact
	err := ctx.ReadResource("google-native:essentialcontacts/v1:OrganizationContact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationContact resources.
type organizationContactState struct {
}

type OrganizationContactState struct {
}

func (OrganizationContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationContactState)(nil)).Elem()
}

type organizationContactArgs struct {
	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email string `pulumi:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag string `pulumi:"languageTag"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []OrganizationContactNotificationCategorySubscriptionsItem `pulumi:"notificationCategorySubscriptions"`
	OrganizationId                    string                                                     `pulumi:"organizationId"`
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime *string `pulumi:"validateTime"`
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState *OrganizationContactValidationState `pulumi:"validationState"`
}

// The set of arguments for constructing a OrganizationContact resource.
type OrganizationContactArgs struct {
	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email pulumi.StringInput
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag pulumi.StringInput
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions OrganizationContactNotificationCategorySubscriptionsItemArrayInput
	OrganizationId                    pulumi.StringInput
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime pulumi.StringPtrInput
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState OrganizationContactValidationStatePtrInput
}

func (OrganizationContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationContactArgs)(nil)).Elem()
}

type OrganizationContactInput interface {
	pulumi.Input

	ToOrganizationContactOutput() OrganizationContactOutput
	ToOrganizationContactOutputWithContext(ctx context.Context) OrganizationContactOutput
}

func (*OrganizationContact) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationContact)(nil)).Elem()
}

func (i *OrganizationContact) ToOrganizationContactOutput() OrganizationContactOutput {
	return i.ToOrganizationContactOutputWithContext(context.Background())
}

func (i *OrganizationContact) ToOrganizationContactOutputWithContext(ctx context.Context) OrganizationContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationContactOutput)
}

type OrganizationContactOutput struct{ *pulumi.OutputState }

func (OrganizationContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationContact)(nil)).Elem()
}

func (o OrganizationContactOutput) ToOrganizationContactOutput() OrganizationContactOutput {
	return o
}

func (o OrganizationContactOutput) ToOrganizationContactOutputWithContext(ctx context.Context) OrganizationContactOutput {
	return o
}

// The email address to send notifications to. The email address does not need to be a Google Account.
func (o OrganizationContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
func (o OrganizationContactOutput) LanguageTag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.LanguageTag }).(pulumi.StringOutput)
}

// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
func (o OrganizationContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The categories of notifications that the contact will receive communications for.
func (o OrganizationContactOutput) NotificationCategorySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringArrayOutput { return v.NotificationCategorySubscriptions }).(pulumi.StringArrayOutput)
}

func (o OrganizationContactOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
func (o OrganizationContactOutput) ValidateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.ValidateTime }).(pulumi.StringOutput)
}

// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
func (o OrganizationContactOutput) ValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationContact) pulumi.StringOutput { return v.ValidationState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationContactInput)(nil)).Elem(), &OrganizationContact{})
	pulumi.RegisterOutputType(OrganizationContactOutput{})
}
