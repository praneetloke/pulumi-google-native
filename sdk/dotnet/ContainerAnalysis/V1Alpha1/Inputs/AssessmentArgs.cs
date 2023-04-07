// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Assessment provides all information that is related to a single vulnerability for this product.
    /// </summary>
    public sealed class AssessmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Holds the MITRE standard Common Vulnerabilities and Exposures (CVE) tracking number for the vulnerability.
        /// </summary>
        [Input("cve")]
        public Input<string>? Cve { get; set; }

        [Input("impacts")]
        private InputList<string>? _impacts;

        /// <summary>
        /// Contains information about the impact of this vulnerability, this will change with time.
        /// </summary>
        public InputList<string> Impacts
        {
            get => _impacts ?? (_impacts = new InputList<string>());
            set => _impacts = value;
        }

        /// <summary>
        /// Justification provides the justification when the state of the assessment if NOT_AFFECTED.
        /// </summary>
        [Input("justification")]
        public Input<Inputs.JustificationArgs>? Justification { get; set; }

        /// <summary>
        /// A detailed description of this Vex.
        /// </summary>
        [Input("longDescription")]
        public Input<string>? LongDescription { get; set; }

        [Input("relatedUris")]
        private InputList<Inputs.URIArgs>? _relatedUris;

        /// <summary>
        /// Holds a list of references associated with this vulnerability item and assessment. These uris have additional information about the vulnerability and the assessment itself. E.g. Link to a document which details how this assessment concluded the state of this vulnerability.
        /// </summary>
        public InputList<Inputs.URIArgs> RelatedUris
        {
            get => _relatedUris ?? (_relatedUris = new InputList<Inputs.URIArgs>());
            set => _relatedUris = value;
        }

        [Input("remediations")]
        private InputList<Inputs.RemediationArgs>? _remediations;

        /// <summary>
        /// Specifies details on how to handle (and presumably, fix) a vulnerability.
        /// </summary>
        public InputList<Inputs.RemediationArgs> Remediations
        {
            get => _remediations ?? (_remediations = new InputList<Inputs.RemediationArgs>());
            set => _remediations = value;
        }

        /// <summary>
        /// A one sentence description of this Vex.
        /// </summary>
        [Input("shortDescription")]
        public Input<string>? ShortDescription { get; set; }

        /// <summary>
        /// Provides the state of this Vulnerability assessment.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.AssessmentState>? State { get; set; }

        public AssessmentArgs()
        {
        }
        public static new AssessmentArgs Empty => new AssessmentArgs();
    }
}
