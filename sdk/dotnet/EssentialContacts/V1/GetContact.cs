// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.EssentialContacts.V1
{
    public static class GetContact
    {
        /// <summary>
        /// Gets a single contact.
        /// </summary>
        public static Task<GetContactResult> InvokeAsync(GetContactArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactResult>("google-native:essentialcontacts/v1:getContact", args ?? new GetContactArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a single contact.
        /// </summary>
        public static Output<GetContactResult> Invoke(GetContactInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactResult>("google-native:essentialcontacts/v1:getContact", args ?? new GetContactInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactArgs : global::Pulumi.InvokeArgs
    {
        [Input("contactId", required: true)]
        public string ContactId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetContactArgs()
        {
        }
        public static new GetContactArgs Empty => new GetContactArgs();
    }

    public sealed class GetContactInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("contactId", required: true)]
        public Input<string> ContactId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetContactInvokeArgs()
        {
        }
        public static new GetContactInvokeArgs Empty => new GetContactInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactResult
    {
        /// <summary>
        /// The email address to send notifications to. The email address does not need to be a Google Account.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        /// </summary>
        public readonly string LanguageTag;
        /// <summary>
        /// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The categories of notifications that the contact will receive communications for.
        /// </summary>
        public readonly ImmutableArray<string> NotificationCategorySubscriptions;
        /// <summary>
        /// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        /// </summary>
        public readonly string ValidateTime;
        /// <summary>
        /// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        /// </summary>
        public readonly string ValidationState;

        [OutputConstructor]
        private GetContactResult(
            string email,

            string languageTag,

            string name,

            ImmutableArray<string> notificationCategorySubscriptions,

            string validateTime,

            string validationState)
        {
            Email = email;
            LanguageTag = languageTag;
            Name = name;
            NotificationCategorySubscriptions = notificationCategorySubscriptions;
            ValidateTime = validateTime;
            ValidationState = validationState;
        }
    }
}
