// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// Message contains the transport layer information to verify the proxy server.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponse
    {
        /// <summary>
        /// PEM encoded CA certificate associated with the proxy server certificate.
        /// </summary>
        public readonly string ServerCaCertPem;
        /// <summary>
        /// Optional. PEM encoded CA certificate associated with the certificate used by proxy server for SSL decryption.
        /// </summary>
        public readonly string SslDecryptCaCertPem;

        [OutputConstructor]
        private GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponse(
            string serverCaCertPem,

            string sslDecryptCaCertPem)
        {
            ServerCaCertPem = serverCaCertPem;
            SslDecryptCaCertPem = sslDecryptCaCertPem;
        }
    }
}
