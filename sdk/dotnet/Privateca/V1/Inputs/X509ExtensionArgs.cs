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
    /// An X509Extension specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.
    /// </summary>
    public sealed class X509ExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
        /// </summary>
        [Input("critical")]
        public Input<bool>? Critical { get; set; }

        /// <summary>
        /// The OID for this X.509 extension.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<Inputs.ObjectIdArgs> ObjectId { get; set; } = null!;

        /// <summary>
        /// The value of this X.509 extension.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public X509ExtensionArgs()
        {
        }
        public static new X509ExtensionArgs Empty => new X509ExtensionArgs();
    }
}
