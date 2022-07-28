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
    /// Defines controls over all certificate issuance within a CaPool.
    /// </summary>
    public sealed class IssuancePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. If specified, then only methods allowed in the IssuanceModes may be used to issue Certificates.
        /// </summary>
        [Input("allowedIssuanceModes")]
        public Input<Inputs.IssuanceModesArgs>? AllowedIssuanceModes { get; set; }

        [Input("allowedKeyTypes")]
        private InputList<Inputs.AllowedKeyTypeArgs>? _allowedKeyTypes;

        /// <summary>
        /// Optional. If any AllowedKeyType is specified, then the certificate request's public key must match one of the key types listed here. Otherwise, any key may be used.
        /// </summary>
        public InputList<Inputs.AllowedKeyTypeArgs> AllowedKeyTypes
        {
            get => _allowedKeyTypes ?? (_allowedKeyTypes = new InputList<Inputs.AllowedKeyTypeArgs>());
            set => _allowedKeyTypes = value;
        }

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all certificates issued through this CaPool. If a certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If a certificate request uses a CertificateTemplate that defines conflicting predefined_values for the same properties, the certificate issuance request will fail.
        /// </summary>
        [Input("baselineValues")]
        public Input<Inputs.X509ParametersArgs>? BaselineValues { get; set; }

        /// <summary>
        /// Optional. Describes constraints on identities that may appear in Certificates issued through this CaPool. If this is omitted, then this CaPool will not add restrictions on a certificate's identity.
        /// </summary>
        [Input("identityConstraints")]
        public Input<Inputs.CertificateIdentityConstraintsArgs>? IdentityConstraints { get; set; }

        /// <summary>
        /// Optional. The maximum lifetime allowed for issued Certificates. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.
        /// </summary>
        [Input("maximumLifetime")]
        public Input<string>? MaximumLifetime { get; set; }

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued through this CaPool. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If a certificate request uses a CertificateTemplate with predefined_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this CaPool will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CaPool's baseline_values.
        /// </summary>
        [Input("passthroughExtensions")]
        public Input<Inputs.CertificateExtensionConstraintsArgs>? PassthroughExtensions { get; set; }

        public IssuancePolicyArgs()
        {
        }
        public static new IssuancePolicyArgs Empty => new IssuancePolicyArgs();
    }
}
