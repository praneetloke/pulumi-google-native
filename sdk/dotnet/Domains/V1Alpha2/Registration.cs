// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Alpha2
{
    /// <summary>
    /// Registers a new domain name and creates a corresponding `Registration` resource. Call `RetrieveRegisterParameters` first to check availability of the domain name and determine parameters like price that are needed to build a call to this method. A successful call creates a `Registration` resource in state `REGISTRATION_PENDING`, which resolves to `ACTIVE` within 1-2 minutes, indicating that the domain was successfully registered. If the resource ends up in state `REGISTRATION_FAILED`, it indicates that the domain was not registered successfully, and you can safely delete the resource and retry registration.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:domains/v1alpha2:Registration")]
    public partial class Registration : Pulumi.CustomResource
    {
        /// <summary>
        /// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
        /// </summary>
        [Output("contactSettings")]
        public Output<Outputs.ContactSettingsResponse> ContactSettings { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the `Registration` resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
        /// </summary>
        [Output("dnsSettings")]
        public Output<Outputs.DnsSettingsResponse> DnsSettings { get; private set; } = null!;

        /// <summary>
        /// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The expiration timestamp of the `Registration`.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The set of issues with the `Registration` that require attention.
        /// </summary>
        [Output("issues")]
        public Output<ImmutableArray<string>> Issues { get; private set; } = null!;

        /// <summary>
        /// Set of labels associated with the `Registration`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
        /// </summary>
        [Output("managementSettings")]
        public Output<Outputs.ManagementSettingsResponse> ManagementSettings { get; private set; } = null!;

        /// <summary>
        /// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not yet been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
        /// </summary>
        [Output("pendingContactSettings")]
        public Output<Outputs.ContactSettingsResponse> PendingContactSettings { get; private set; } = null!;

        /// <summary>
        /// The state of the `Registration`
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
        /// </summary>
        [Output("supportedPrivacy")]
        public Output<ImmutableArray<string>> SupportedPrivacy { get; private set; } = null!;


        /// <summary>
        /// Create a Registration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Registration(string name, RegistrationArgs args, CustomResourceOptions? options = null)
            : base("google-native:domains/v1alpha2:Registration", name, args ?? new RegistrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Registration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:domains/v1alpha2:Registration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Registration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Registration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Registration(name, id, options);
        }
    }

    public sealed class RegistrationArgs : Pulumi.ResourceArgs
    {
        [Input("contactNotices")]
        private InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationContactNoticesItem>? _contactNotices;

        /// <summary>
        /// The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationContactNoticesItem> ContactNotices
        {
            get => _contactNotices ?? (_contactNotices = new InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationContactNoticesItem>());
            set => _contactNotices = value;
        }

        /// <summary>
        /// Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
        /// </summary>
        [Input("contactSettings", required: true)]
        public Input<Inputs.ContactSettingsArgs> ContactSettings { get; set; } = null!;

        /// <summary>
        /// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.DnsSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("domainNotices")]
        private InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationDomainNoticesItem>? _domainNotices;

        /// <summary>
        /// The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationDomainNoticesItem> DomainNotices
        {
            get => _domainNotices ?? (_domainNotices = new InputList<Pulumi.GoogleNative.Domains.V1Alpha2.RegistrationDomainNoticesItem>());
            set => _domainNotices = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of labels associated with the `Registration`.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
        /// </summary>
        [Input("managementSettings")]
        public Input<Inputs.ManagementSettingsArgs>? ManagementSettings { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// When true, only validation will be performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
        /// </summary>
        [Input("validateOnly")]
        public Input<bool>? ValidateOnly { get; set; }

        /// <summary>
        /// Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
        /// </summary>
        [Input("yearlyPrice", required: true)]
        public Input<Inputs.MoneyArgs> YearlyPrice { get; set; } = null!;

        public RegistrationArgs()
        {
        }
    }
}
