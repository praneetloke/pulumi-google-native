// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.GoogleNative
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("google-native");
        /// <summary>
        /// Additional user-agent string to append to the default one (&lt;prod_name&gt;/&lt;ver&gt;).
        /// </summary>
        public static string? AppendUserAgent { get; set; } = __config.Get("appendUserAgent");

        /// <summary>
        /// This will disable the Pulumi Partner Name which is used if a custom `partnerName` isn't specified.
        /// </summary>
        public static bool? DisablePartnerName { get; set; } = __config.GetBoolean("disablePartnerName");

        /// <summary>
        /// A Google Partner Name to facilitate partner resource usage attribution.
        /// </summary>
        public static string? PartnerName { get; set; } = __config.Get("partnerName");

        /// <summary>
        /// The default project to manage resources in. If another project is specified on a resource, it will take precedence.
        /// </summary>
        public static string? Project { get; set; } = __config.Get("project");

        /// <summary>
        /// The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.
        /// </summary>
        public static string? Region { get; set; } = __config.Get("region");

        /// <summary>
        /// The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.
        /// </summary>
        public static string? Zone { get; set; } = __config.Get("zone");

    }
}
