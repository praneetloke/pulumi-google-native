// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// The actual payload that contains the SBOM Reference data. The payload follows the intoto statement specification. See https://github.com/in-toto/attestation/blob/main/spec/v1.0/statement.md for more details.
    /// </summary>
    public sealed class SbomReferenceIntotoPayloadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Additional parameters of the Predicate. Includes the actual data about the SBOM.
        /// </summary>
        [Input("predicate")]
        public Input<Inputs.SbomReferenceIntotoPredicateArgs>? Predicate { get; set; }

        /// <summary>
        /// URI identifying the type of the Predicate.
        /// </summary>
        [Input("predicateType")]
        public Input<string>? PredicateType { get; set; }

        [Input("subject")]
        private InputList<Inputs.SubjectArgs>? _subject;

        /// <summary>
        /// Set of software artifacts that the attestation applies to. Each element represents a single software artifact.
        /// </summary>
        public InputList<Inputs.SubjectArgs> Subject
        {
            get => _subject ?? (_subject = new InputList<Inputs.SubjectArgs>());
            set => _subject = value;
        }

        /// <summary>
        /// Identifier for the schema of the Statement.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SbomReferenceIntotoPayloadArgs()
        {
        }
        public static new SbomReferenceIntotoPayloadArgs Empty => new SbomReferenceIntotoPayloadArgs();
    }
}
