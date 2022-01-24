// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// TLS configuration information for virtual hosts and TargetServers.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1TlsInfoResponse
    {
        /// <summary>
        /// The SSL/TLS cipher suites to be used. Must be one of the cipher suite names listed in: http://docs.oracle.com/javase/8/docs/technotes/guides/security/StandardNames.html#ciphersuites
        /// </summary>
        public readonly ImmutableArray<string> Ciphers;
        /// <summary>
        /// Optional. Enables two-way TLS.
        /// </summary>
        public readonly bool ClientAuthEnabled;
        /// <summary>
        /// The TLS Common Name of the certificate.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1TlsInfoCommonNameResponse CommonName;
        /// <summary>
        /// Enables TLS. If false, neither one-way nor two-way TLS will be enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// If true, Edge ignores TLS certificate errors. Valid when configuring TLS for target servers and target endpoints, and when configuring virtual hosts that use 2-way TLS. When used with a target endpoint/target server, if the backend system uses SNI and returns a cert with a subject Distinguished Name (DN) that does not match the hostname, there is no way to ignore the error and the connection fails.
        /// </summary>
        public readonly bool IgnoreValidationErrors;
        /// <summary>
        /// Required if `client_auth_enabled` is true. The resource ID for the alias containing the private key and cert.
        /// </summary>
        public readonly string KeyAlias;
        /// <summary>
        /// Required if `client_auth_enabled` is true. The resource ID of the keystore.
        /// </summary>
        public readonly string KeyStore;
        /// <summary>
        /// The TLS versioins to be used.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        /// <summary>
        /// The resource ID of the truststore.
        /// </summary>
        public readonly string TrustStore;

        [OutputConstructor]
        private GoogleCloudApigeeV1TlsInfoResponse(
            ImmutableArray<string> ciphers,

            bool clientAuthEnabled,

            Outputs.GoogleCloudApigeeV1TlsInfoCommonNameResponse commonName,

            bool enabled,

            bool ignoreValidationErrors,

            string keyAlias,

            string keyStore,

            ImmutableArray<string> protocols,

            string trustStore)
        {
            Ciphers = ciphers;
            ClientAuthEnabled = clientAuthEnabled;
            CommonName = commonName;
            Enabled = enabled;
            IgnoreValidationErrors = ignoreValidationErrors;
            KeyAlias = keyAlias;
            KeyStore = keyStore;
            Protocols = protocols;
            TrustStore = trustStore;
        }
    }
}
