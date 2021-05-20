// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    /// <summary>
    /// Create a new CryptoKeyVersion in a CryptoKey. The server will assign the next sequential id. If unset, state will be set to ENABLED.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion")]
    public partial class KeyRingCryptoKeyCryptoKeyVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.KeyOperationAttestationResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// The time at which this CryptoKeyVersion was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
        /// </summary>
        [Output("destroyEventTime")]
        public Output<string> DestroyEventTime { get; private set; } = null!;

        /// <summary>
        /// The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
        /// </summary>
        [Output("destroyTime")]
        public Output<string> DestroyTime { get; private set; } = null!;

        /// <summary>
        /// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        /// </summary>
        [Output("externalProtectionLevelOptions")]
        public Output<Outputs.ExternalProtectionLevelOptionsResponse> ExternalProtectionLevelOptions { get; private set; } = null!;

        /// <summary>
        /// The time this CryptoKeyVersion's key material was generated.
        /// </summary>
        [Output("generateTime")]
        public Output<string> GenerateTime { get; private set; } = null!;

        /// <summary>
        /// The root cause of an import failure. Only present if state is IMPORT_FAILED.
        /// </summary>
        [Output("importFailureReason")]
        public Output<string> ImportFailureReason { get; private set; } = null!;

        /// <summary>
        /// The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
        /// </summary>
        [Output("importJob")]
        public Output<string> ImportJob { get; private set; } = null!;

        /// <summary>
        /// The time at which this CryptoKeyVersion's key material was imported.
        /// </summary>
        [Output("importTime")]
        public Output<string> ImportTime { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        /// </summary>
        [Output("protectionLevel")]
        public Output<string> ProtectionLevel { get; private set; } = null!;

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a KeyRingCryptoKeyCryptoKeyVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRingCryptoKeyCryptoKeyVersion(string name, KeyRingCryptoKeyCryptoKeyVersionArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion", name, args ?? new KeyRingCryptoKeyCryptoKeyVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyRingCryptoKeyCryptoKeyVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudkms/v1:KeyRingCryptoKeyCryptoKeyVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing KeyRingCryptoKeyCryptoKeyVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRingCryptoKeyCryptoKeyVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KeyRingCryptoKeyCryptoKeyVersion(name, id, options);
        }
    }

    public sealed class KeyRingCryptoKeyCryptoKeyVersionArgs : Pulumi.ResourceArgs
    {
        [Input("cryptoKeyId", required: true)]
        public Input<string> CryptoKeyId { get; set; } = null!;

        [Input("cryptoKeyVersionId", required: true)]
        public Input<string> CryptoKeyVersionId { get; set; } = null!;

        /// <summary>
        /// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        /// </summary>
        [Input("externalProtectionLevelOptions")]
        public Input<Inputs.ExternalProtectionLevelOptionsArgs>? ExternalProtectionLevelOptions { get; set; }

        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public KeyRingCryptoKeyCryptoKeyVersionArgs()
        {
        }
    }
}
