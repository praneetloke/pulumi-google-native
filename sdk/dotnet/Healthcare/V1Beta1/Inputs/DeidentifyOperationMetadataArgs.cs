// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Details about the work the de-identify operation performed.
    /// </summary>
    public sealed class DeidentifyOperationMetadataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details about the FHIR store to write the output to.
        /// </summary>
        [Input("fhirOutput")]
        public Input<Inputs.FhirOutputArgs>? FhirOutput { get; set; }

        public DeidentifyOperationMetadataArgs()
        {
        }
    }
}
