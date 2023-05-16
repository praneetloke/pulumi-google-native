// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// Options relating to the publication of each CertificateAuthority's CA certificate and CRLs and their inclusion as extensions in issued Certificates. The options set here apply to certificates issued by any CertificateAuthority in the CaPool.
    /// </summary>
    public sealed class PublishingOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Specifies the encoding format of each CertificateAuthority's CA certificate and CRLs. If this is omitted, CA certificates and CRLs will be published in PEM.
        /// </summary>
        [Input("encodingFormat")]
        public Input<Pulumi.GoogleNative.Privateca.V1.PublishingOptionsEncodingFormat>? EncodingFormat { get; set; }

        /// <summary>
        /// Optional. When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access" X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding X.509 extension will not be written in issued certificates.
        /// </summary>
        [Input("publishCaCert")]
        public Input<bool>? PublishCaCert { get; set; }

        /// <summary>
        /// Optional. When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are also rebuilt shortly after a certificate is revoked.
        /// </summary>
        [Input("publishCrl")]
        public Input<bool>? PublishCrl { get; set; }

        public PublishingOptionsArgs()
        {
        }
        public static new PublishingOptionsArgs Empty => new PublishingOptionsArgs();
    }
}
