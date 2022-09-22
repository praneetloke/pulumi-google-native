// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    public static class GetCryptoKey
    {
        /// <summary>
        /// Returns metadata for a given CryptoKey, as well as its primary CryptoKeyVersion.
        /// </summary>
        public static Task<GetCryptoKeyResult> InvokeAsync(GetCryptoKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCryptoKeyResult>("google-native:cloudkms/v1:getCryptoKey", args ?? new GetCryptoKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns metadata for a given CryptoKey, as well as its primary CryptoKeyVersion.
        /// </summary>
        public static Output<GetCryptoKeyResult> Invoke(GetCryptoKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCryptoKeyResult>("google-native:cloudkms/v1:getCryptoKey", args ?? new GetCryptoKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCryptoKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("cryptoKeyId", required: true)]
        public string CryptoKeyId { get; set; } = null!;

        [Input("keyRingId", required: true)]
        public string KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCryptoKeyArgs()
        {
        }
        public static new GetCryptoKeyArgs Empty => new GetCryptoKeyArgs();
    }

    public sealed class GetCryptoKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cryptoKeyId", required: true)]
        public Input<string> CryptoKeyId { get; set; } = null!;

        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCryptoKeyInvokeArgs()
        {
        }
        public static new GetCryptoKeyInvokeArgs Empty => new GetCryptoKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetCryptoKeyResult
    {
        /// <summary>
        /// The time at which this CryptoKey was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Immutable. The resource name of the backend environment where the key material for all CryptoKeyVersions associated with this CryptoKey reside and where all related cryptographic operations are performed. Only applicable if CryptoKeyVersions have a ProtectionLevel of EXTERNAL_VPC, with the resource name in the format `projects/*/locations/*/ekmConnections/*`. Note, this list is non-exhaustive and may apply to additional ProtectionLevels in the future.
        /// </summary>
        public readonly string CryptoKeyBackend;
        /// <summary>
        /// Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED. If not specified at creation time, the default duration is 24 hours.
        /// </summary>
        public readonly string DestroyScheduledDuration;
        /// <summary>
        /// Immutable. Whether this key may contain imported versions only.
        /// </summary>
        public readonly bool ImportOnly;
        /// <summary>
        /// Labels with user-defined metadata. For more information, see [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// At next_rotation_time, the Key Management Service will automatically: 1. Create a new version of this CryptoKey. 2. Mark the new version as primary. Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
        /// </summary>
        public readonly string NextRotationTime;
        /// <summary>
        /// A copy of the "primary" CryptoKeyVersion that will be used by Encrypt when this CryptoKey is given in EncryptRequest.name. The CryptoKey's primary version can be updated via UpdateCryptoKeyPrimaryVersion. Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be omitted.
        /// </summary>
        public readonly Outputs.CryptoKeyVersionResponse Primary;
        /// <summary>
        /// Immutable. The immutable purpose of this CryptoKey.
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// next_rotation_time will be advanced by this period when the service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours. If rotation_period is set, next_rotation_time must also be set. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
        /// </summary>
        public readonly string RotationPeriod;
        /// <summary>
        /// A template describing settings for new CryptoKeyVersion instances. The properties of new CryptoKeyVersion instances created by either CreateCryptoKeyVersion or auto-rotation are controlled by this template.
        /// </summary>
        public readonly Outputs.CryptoKeyVersionTemplateResponse VersionTemplate;

        [OutputConstructor]
        private GetCryptoKeyResult(
            string createTime,

            string cryptoKeyBackend,

            string destroyScheduledDuration,

            bool importOnly,

            ImmutableDictionary<string, string> labels,

            string name,

            string nextRotationTime,

            Outputs.CryptoKeyVersionResponse primary,

            string purpose,

            string rotationPeriod,

            Outputs.CryptoKeyVersionTemplateResponse versionTemplate)
        {
            CreateTime = createTime;
            CryptoKeyBackend = cryptoKeyBackend;
            DestroyScheduledDuration = destroyScheduledDuration;
            ImportOnly = importOnly;
            Labels = labels;
            Name = name;
            NextRotationTime = nextRotationTime;
            Primary = primary;
            Purpose = purpose;
            RotationPeriod = rotationPeriod;
            VersionTemplate = versionTemplate;
        }
    }
}
