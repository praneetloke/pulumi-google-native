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
    /// An X509Parameters is used to describe certain fields of an X.509 certificate, such as the key usage fields, fields specific to CA certificates, certificate policy extensions and custom extensions.
    /// </summary>
    public sealed class X509ParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalExtensions")]
        private InputList<Inputs.X509ExtensionArgs>? _additionalExtensions;

        /// <summary>
        /// Optional. Describes custom X.509 extensions.
        /// </summary>
        public InputList<Inputs.X509ExtensionArgs> AdditionalExtensions
        {
            get => _additionalExtensions ?? (_additionalExtensions = new InputList<Inputs.X509ExtensionArgs>());
            set => _additionalExtensions = value;
        }

        [Input("aiaOcspServers")]
        private InputList<string>? _aiaOcspServers;

        /// <summary>
        /// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
        /// </summary>
        public InputList<string> AiaOcspServers
        {
            get => _aiaOcspServers ?? (_aiaOcspServers = new InputList<string>());
            set => _aiaOcspServers = value;
        }

        /// <summary>
        /// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
        /// </summary>
        [Input("caOptions")]
        public Input<Inputs.CaOptionsArgs>? CaOptions { get; set; }

        /// <summary>
        /// Optional. Indicates the intended use for keys that correspond to a certificate.
        /// </summary>
        [Input("keyUsage")]
        public Input<Inputs.KeyUsageArgs>? KeyUsage { get; set; }

        /// <summary>
        /// Optional. Describes the X.509 name constraints extension.
        /// </summary>
        [Input("nameConstraints")]
        public Input<Inputs.NameConstraintsArgs>? NameConstraints { get; set; }

        [Input("policyIds")]
        private InputList<Inputs.ObjectIdArgs>? _policyIds;

        /// <summary>
        /// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
        /// </summary>
        public InputList<Inputs.ObjectIdArgs> PolicyIds
        {
            get => _policyIds ?? (_policyIds = new InputList<Inputs.ObjectIdArgs>());
            set => _policyIds = value;
        }

        public X509ParametersArgs()
        {
        }
        public static new X509ParametersArgs Empty => new X509ParametersArgs();
    }
}
