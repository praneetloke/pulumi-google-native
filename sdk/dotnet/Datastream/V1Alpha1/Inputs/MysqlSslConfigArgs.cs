// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Inputs
{

    /// <summary>
    /// MySQL SSL configuration information.
    /// </summary>
    public sealed class MysqlSslConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input only. PEM-encoded certificate of the CA that signed the source database server's certificate.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        /// <summary>
        /// Input only. PEM-encoded certificate that will be used by the replica to authenticate against the source database server. If this field is used then the 'client_key' and the 'ca_certificate' fields are mandatory.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// Input only. PEM-encoded private key associated with the Client Certificate. If this field is used then the 'client_certificate' and the 'ca_certificate' fields are mandatory.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        public MysqlSslConfigArgs()
        {
        }
        public static new MysqlSslConfigArgs Empty => new MysqlSslConfigArgs();
    }
}
