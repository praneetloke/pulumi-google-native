// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1.Inputs
{

    /// <summary>
    /// Represents an OpenId Connect 1.0 identity provider.
    /// </summary>
    public sealed class GoogleIamAdminV1WorkforcePoolProviderOidcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID. Must match the audience claim of the JWT issued by the identity provider.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
        /// </summary>
        [Input("issuerUri", required: true)]
        public Input<string> IssuerUri { get; set; } = null!;

        public GoogleIamAdminV1WorkforcePoolProviderOidcArgs()
        {
        }
        public static new GoogleIamAdminV1WorkforcePoolProviderOidcArgs Empty => new GoogleIamAdminV1WorkforcePoolProviderOidcArgs();
    }
}
