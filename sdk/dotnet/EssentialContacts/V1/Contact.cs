// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.EssentialContacts.V1
{
    /// <summary>
    /// Adds a new contact for a resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:essentialcontacts/v1:Contact")]
    public partial class Contact : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The email address to send notifications to. This does not need to be a Google account.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        /// </summary>
        [Output("languageTag")]
        public Output<string> LanguageTag { get; private set; } = null!;

        /// <summary>
        /// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The categories of notifications that the contact will receive communications for.
        /// </summary>
        [Output("notificationCategorySubscriptions")]
        public Output<ImmutableArray<string>> NotificationCategorySubscriptions { get; private set; } = null!;

        /// <summary>
        /// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        /// </summary>
        [Output("validateTime")]
        public Output<string> ValidateTime { get; private set; } = null!;

        /// <summary>
        /// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        /// </summary>
        [Output("validationState")]
        public Output<string> ValidationState { get; private set; } = null!;


        /// <summary>
        /// Create a Contact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Contact(string name, ContactArgs args, CustomResourceOptions? options = null)
            : base("google-native:essentialcontacts/v1:Contact", name, args ?? new ContactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Contact(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:essentialcontacts/v1:Contact", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Contact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Contact Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Contact(name, id, options);
        }
    }

    public sealed class ContactArgs : Pulumi.ResourceArgs
    {
        [Input("contactId", required: true)]
        public Input<string> ContactId { get; set; } = null!;

        /// <summary>
        /// Required. The email address to send notifications to. This does not need to be a Google account.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        /// </summary>
        [Input("languageTag")]
        public Input<string>? LanguageTag { get; set; }

        /// <summary>
        /// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationCategorySubscriptions")]
        private InputList<string>? _notificationCategorySubscriptions;

        /// <summary>
        /// The categories of notifications that the contact will receive communications for.
        /// </summary>
        public InputList<string> NotificationCategorySubscriptions
        {
            get => _notificationCategorySubscriptions ?? (_notificationCategorySubscriptions = new InputList<string>());
            set => _notificationCategorySubscriptions = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        /// </summary>
        [Input("validateTime")]
        public Input<string>? ValidateTime { get; set; }

        /// <summary>
        /// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        /// </summary>
        [Input("validationState")]
        public Input<string>? ValidationState { get; set; }

        public ContactArgs()
        {
        }
    }
}
