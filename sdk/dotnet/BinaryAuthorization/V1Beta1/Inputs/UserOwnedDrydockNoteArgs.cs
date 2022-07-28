// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BinaryAuthorization.V1Beta1.Inputs
{

    /// <summary>
    /// An user owned drydock note references a Drydock ATTESTATION_AUTHORITY Note created by the user.
    /// </summary>
    public sealed class UserOwnedDrydockNoteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Drydock resource name of a ATTESTATION_AUTHORITY Note, created by the user, in the format: `projects/*/notes/*` (or the legacy `providers/*/notes/*`). This field may not be updated. An attestation by this attestor is stored as a Drydock ATTESTATION_AUTHORITY Occurrence that names a container image and that links to this Note. Drydock is an external dependency.
        /// </summary>
        [Input("noteReference", required: true)]
        public Input<string> NoteReference { get; set; } = null!;

        [Input("publicKeys")]
        private InputList<Inputs.AttestorPublicKeyArgs>? _publicKeys;

        /// <summary>
        /// Optional. Public keys that verify attestations signed by this attestor. This field may be updated. If this field is non-empty, one of the specified public keys must verify that an attestation was signed by this attestor for the image specified in the admission request. If this field is empty, this attestor always returns that no valid attestations exist.
        /// </summary>
        public InputList<Inputs.AttestorPublicKeyArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<Inputs.AttestorPublicKeyArgs>());
            set => _publicKeys = value;
        }

        public UserOwnedDrydockNoteArgs()
        {
        }
        public static new UserOwnedDrydockNoteArgs Empty => new UserOwnedDrydockNoteArgs();
    }
}
