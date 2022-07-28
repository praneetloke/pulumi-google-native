// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1
{
    public static class GetKey
    {
        /// <summary>
        /// Returns the specified key.
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("google-native:recaptchaenterprise/v1:getKey", args ?? new GetKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified key.
        /// </summary>
        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKeyResult>("google-native:recaptchaenterprise/v1:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetKeyArgs()
        {
        }
        public static new GetKeyArgs Empty => new GetKeyArgs();
    }

    public sealed class GetKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetKeyInvokeArgs()
        {
        }
        public static new GetKeyInvokeArgs Empty => new GetKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// Settings for keys that can be used by Android apps.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse AndroidSettings;
        /// <summary>
        /// The timestamp corresponding to the creation of this Key.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human-readable display name of this key. Modifiable by user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Settings for keys that can be used by iOS apps.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse IosSettings;
        /// <summary>
        /// See Creating and managing labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name for the Key in the format "projects/{project}/keys/{key}".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Options for user acceptance testing.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse TestingOptions;
        /// <summary>
        /// Settings for WAF
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1WafSettingsResponse WafSettings;
        /// <summary>
        /// Settings for keys that can be used by websites.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse WebSettings;

        [OutputConstructor]
        private GetKeyResult(
            Outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse androidSettings,

            string createTime,

            string displayName,

            Outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse iosSettings,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse testingOptions,

            Outputs.GoogleCloudRecaptchaenterpriseV1WafSettingsResponse wafSettings,

            Outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse webSettings)
        {
            AndroidSettings = androidSettings;
            CreateTime = createTime;
            DisplayName = displayName;
            IosSettings = iosSettings;
            Labels = labels;
            Name = name;
            TestingOptions = testingOptions;
            WafSettings = wafSettings;
            WebSettings = webSettings;
        }
    }
}
