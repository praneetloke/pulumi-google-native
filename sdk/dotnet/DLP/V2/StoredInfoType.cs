// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    /// <summary>
    /// Creates a pre-built stored infoType to be used for inspection. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dlp/v2:StoredInfoType")]
    public partial class StoredInfoType : Pulumi.CustomResource
    {
        /// <summary>
        /// Current version of the stored info type.
        /// </summary>
        [Output("currentVersion")]
        public Output<Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse> CurrentVersion { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Pending versions of the stored info type. Empty if no versions are pending.
        /// </summary>
        [Output("pendingVersions")]
        public Output<ImmutableArray<Outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse>> PendingVersions { get; private set; } = null!;


        /// <summary>
        /// Create a StoredInfoType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StoredInfoType(string name, StoredInfoTypeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:StoredInfoType", name, args ?? new StoredInfoTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StoredInfoType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dlp/v2:StoredInfoType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing StoredInfoType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StoredInfoType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StoredInfoType(name, id, options);
        }
    }

    public sealed class StoredInfoTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the storedInfoType to create.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.GooglePrivacyDlpV2StoredInfoTypeConfigArgs> Config { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        /// </summary>
        [Input("storedInfoTypeId")]
        public Input<string>? StoredInfoTypeId { get; set; }

        public StoredInfoTypeArgs()
        {
        }
    }
}
