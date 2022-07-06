// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// Include to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. The key must grant the Cloud IAM permission `cloudkms.cryptoKeyVersions.useToDecrypt` to the project's Cloud Healthcare Service Agent service account. For more information, see [Creating a wrapped key] (https://cloud.google.com/dlp/docs/create-wrapped-key).
    /// </summary>
    [OutputType]
    public sealed class KmsWrappedCryptoKeyResponse
    {
        /// <summary>
        /// The resource name of the KMS CryptoKey to use for unwrapping. For example, `projects/{project_id}/locations/{location_id}/keyRings/{keyring}/cryptoKeys/{key}`.
        /// </summary>
        public readonly string CryptoKey;
        /// <summary>
        /// The wrapped data crypto key.
        /// </summary>
        public readonly string WrappedKey;

        [OutputConstructor]
        private KmsWrappedCryptoKeyResponse(
            string cryptoKey,

            string wrappedKey)
        {
            CryptoKey = cryptoKey;
            WrappedKey = wrappedKey;
        }
    }
}
