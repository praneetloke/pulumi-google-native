// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BinaryAuthorization.V1Beta1
{
    public static class GetAttestor
    {
        /// <summary>
        /// Gets an attestor. Returns NOT_FOUND if the attestor does not exist.
        /// </summary>
        public static Task<GetAttestorResult> InvokeAsync(GetAttestorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAttestorResult>("google-native:binaryauthorization/v1beta1:getAttestor", args ?? new GetAttestorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an attestor. Returns NOT_FOUND if the attestor does not exist.
        /// </summary>
        public static Output<GetAttestorResult> Invoke(GetAttestorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAttestorResult>("google-native:binaryauthorization/v1beta1:getAttestor", args ?? new GetAttestorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAttestorArgs : global::Pulumi.InvokeArgs
    {
        [Input("attestorId", required: true)]
        public string AttestorId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAttestorArgs()
        {
        }
        public static new GetAttestorArgs Empty => new GetAttestorArgs();
    }

    public sealed class GetAttestorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("attestorId", required: true)]
        public Input<string> AttestorId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAttestorInvokeArgs()
        {
        }
        public static new GetAttestorInvokeArgs Empty => new GetAttestorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAttestorResult
    {
        /// <summary>
        /// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. A checksum, returned by the server, that can be sent on update requests to ensure the attestor has an up-to-date value before attempting to update it. See https://google.aip.dev/154.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Time when the attestor was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        /// </summary>
        public readonly Outputs.UserOwnedDrydockNoteResponse UserOwnedDrydockNote;

        [OutputConstructor]
        private GetAttestorResult(
            string description,

            string etag,

            string name,

            string updateTime,

            Outputs.UserOwnedDrydockNoteResponse userOwnedDrydockNote)
        {
            Description = description;
            Etag = etag;
            Name = name;
            UpdateTime = updateTime;
            UserOwnedDrydockNote = userOwnedDrydockNote;
        }
    }
}
