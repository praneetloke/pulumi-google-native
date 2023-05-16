// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.Outputs
{

    /// <summary>
    /// Settings specific to keys that can be used by iOS apps.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse
    {
        /// <summary>
        /// If set to true, allowed_bundle_ids are not enforced.
        /// </summary>
        public readonly bool AllowAllBundleIds;
        /// <summary>
        /// iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'
        /// </summary>
        public readonly ImmutableArray<string> AllowedBundleIds;
        /// <summary>
        /// Apple Developer account details for the app the reCAPTCHA key will protect. reCAPTCHA Enterprise leverages platform specific checks like Apple AppAttest and Apple DeviceCheck to protect your app from abuse. Providing these fields allows reCAPTCHA Enterprise to get a better assessment of the integrity of your app.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1AppleDeveloperIdResponse AppleDeveloperId;

        [OutputConstructor]
        private GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse(
            bool allowAllBundleIds,

            ImmutableArray<string> allowedBundleIds,

            Outputs.GoogleCloudRecaptchaenterpriseV1AppleDeveloperIdResponse appleDeveloperId)
        {
            AllowAllBundleIds = allowAllBundleIds;
            AllowedBundleIds = allowedBundleIds;
            AppleDeveloperId = appleDeveloperId;
        }
    }
}
