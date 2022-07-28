// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetRegionSecurityPolicy
    {
        /// <summary>
        /// List all of the ordered rules present in a single specified policy.
        /// </summary>
        public static Task<GetRegionSecurityPolicyResult> InvokeAsync(GetRegionSecurityPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionSecurityPolicyResult>("google-native:compute/v1:getRegionSecurityPolicy", args ?? new GetRegionSecurityPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// List all of the ordered rules present in a single specified policy.
        /// </summary>
        public static Output<GetRegionSecurityPolicyResult> Invoke(GetRegionSecurityPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionSecurityPolicyResult>("google-native:compute/v1:getRegionSecurityPolicy", args ?? new GetRegionSecurityPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionSecurityPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("securityPolicy", required: true)]
        public string SecurityPolicy { get; set; } = null!;

        public GetRegionSecurityPolicyArgs()
        {
        }
        public static new GetRegionSecurityPolicyArgs Empty => new GetRegionSecurityPolicyArgs();
    }

    public sealed class GetRegionSecurityPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("securityPolicy", required: true)]
        public Input<string> SecurityPolicy { get; set; } = null!;

        public GetRegionSecurityPolicyInvokeArgs()
        {
        }
        public static new GetRegionSecurityPolicyInvokeArgs Empty => new GetRegionSecurityPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionSecurityPolicyResult
    {
        public readonly Outputs.SecurityPolicyAdaptiveProtectionConfigResponse AdaptiveProtectionConfig;
        public readonly Outputs.SecurityPolicyAdvancedOptionsConfigResponse AdvancedOptionsConfig;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        public readonly Outputs.SecurityPolicyDdosProtectionConfigResponse DdosProtectionConfig;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.SecurityPolicyRecaptchaOptionsConfigResponse RecaptchaOptionsConfig;
        /// <summary>
        /// URL of the region where the regional security policy resides. This field is not applicable to global security policies.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityPolicyRuleResponse> Rules;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegionSecurityPolicyResult(
            Outputs.SecurityPolicyAdaptiveProtectionConfigResponse adaptiveProtectionConfig,

            Outputs.SecurityPolicyAdvancedOptionsConfigResponse advancedOptionsConfig,

            string creationTimestamp,

            Outputs.SecurityPolicyDdosProtectionConfigResponse ddosProtectionConfig,

            string description,

            string fingerprint,

            string kind,

            string name,

            Outputs.SecurityPolicyRecaptchaOptionsConfigResponse recaptchaOptionsConfig,

            string region,

            ImmutableArray<Outputs.SecurityPolicyRuleResponse> rules,

            string selfLink,

            string type)
        {
            AdaptiveProtectionConfig = adaptiveProtectionConfig;
            AdvancedOptionsConfig = advancedOptionsConfig;
            CreationTimestamp = creationTimestamp;
            DdosProtectionConfig = ddosProtectionConfig;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            RecaptchaOptionsConfig = recaptchaOptionsConfig;
            Region = region;
            Rules = rules;
            SelfLink = selfLink;
            Type = type;
        }
    }
}
