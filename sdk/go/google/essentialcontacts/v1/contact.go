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
type Contact struct {
	pulumi.CustomResourceState

	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email pulumi.StringOutput `pulumi:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag pulumi.StringOutput `pulumi:"languageTag"`
	// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions pulumi.StringArrayOutput `pulumi:"notificationCategorySubscriptions"`
	Project                           pulumi.StringOutput      `pulumi:"project"`
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime pulumi.StringOutput `pulumi:"validateTime"`
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState pulumi.StringOutput `pulumi:"validationState"`
}

// NewContact registers a new resource with the given unique name, arguments, and options.
func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
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
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Contact
	err := ctx.RegisterResource("google-native:essentialcontacts/v1:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContact gets an existing Contact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("google-native:essentialcontacts/v1:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contact resources.
type contactState struct {
}

type ContactState struct {
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email string `pulumi:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag string `pulumi:"languageTag"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []ContactNotificationCategorySubscriptionsItem `pulumi:"notificationCategorySubscriptions"`
	Project                           *string                                        `pulumi:"project"`
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime *string `pulumi:"validateTime"`
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState *ContactValidationState `pulumi:"validationState"`
}

// The set of arguments for constructing a Contact resource.
type ContactArgs struct {
	// The email address to send notifications to. The email address does not need to be a Google Account.
	Email pulumi.StringInput
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag pulumi.StringInput
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions ContactNotificationCategorySubscriptionsItemArrayInput
	Project                           pulumi.StringPtrInput
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime pulumi.StringPtrInput
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState ContactValidationStatePtrInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

// The email address to send notifications to. The email address does not need to be a Google Account.
func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
func (o ContactOutput) LanguageTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.LanguageTag }).(pulumi.StringOutput)
}

// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
func (o ContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The categories of notifications that the contact will receive communications for.
func (o ContactOutput) NotificationCategorySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringArrayOutput { return v.NotificationCategorySubscriptions }).(pulumi.StringArrayOutput)
}

func (o ContactOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
func (o ContactOutput) ValidateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ValidateTime }).(pulumi.StringOutput)
}

// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
func (o ContactOutput) ValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ValidationState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactInput)(nil)).Elem(), &Contact{})
	pulumi.RegisterOutputType(ContactOutput{})
}
