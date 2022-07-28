// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// An attestation wrapper that uses the Grafeas `Signature` message. This attestation must define the `serialized_payload` that the `signatures` verify and any metadata necessary to interpret that plaintext. The signatures should always be over the `serialized_payload` bytestring.
    /// </summary>
    public sealed class GenericSignedAttestationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
        /// </summary>
        [Input("contentType")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.GenericSignedAttestationContentType>? ContentType { get; set; }

        /// <summary>
        /// The serialized payload that is verified by one or more `signatures`. The encoding and semantic meaning of this payload must match what is set in `content_type`.
        /// </summary>
        [Input("serializedPayload")]
        public Input<string>? SerializedPayload { get; set; }

        [Input("signatures")]
        private InputList<Inputs.SignatureArgs>? _signatures;

        /// <summary>
        /// One or more signatures over `serialized_payload`. Verifier implementations should consider this attestation message verified if at least one `signature` verifies `serialized_payload`. See `Signature` in common.proto for more details on signature structure and verification.
        /// </summary>
        public InputList<Inputs.SignatureArgs> Signatures
        {
            get => _signatures ?? (_signatures = new InputList<Inputs.SignatureArgs>());
            set => _signatures = value;
        }

        public GenericSignedAttestationArgs()
        {
        }
        public static new GenericSignedAttestationArgs Empty => new GenericSignedAttestationArgs();
    }
}
