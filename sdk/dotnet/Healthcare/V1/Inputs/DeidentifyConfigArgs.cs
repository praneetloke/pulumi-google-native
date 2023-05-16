// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// Configures de-id options specific to different types of content. Each submessage customizes the handling of an https://tools.ietf.org/html/rfc6838 media type or subtype. Configs are applied in a nested manner at runtime.
    /// </summary>
    public sealed class DeidentifyConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures de-id of application/DICOM content.
        /// </summary>
        [Input("dicom")]
        public Input<Inputs.DicomConfigArgs>? Dicom { get; set; }

        /// <summary>
        /// Configures de-id of application/FHIR content.
        /// </summary>
        [Input("fhir")]
        public Input<Inputs.FhirConfigArgs>? Fhir { get; set; }

        /// <summary>
        /// Configures de-identification of image pixels wherever they are found in the source_dataset.
        /// </summary>
        [Input("image")]
        public Input<Inputs.ImageConfigArgs>? Image { get; set; }

        /// <summary>
        /// Configures de-identification of text wherever it is found in the source_dataset.
        /// </summary>
        [Input("text")]
        public Input<Inputs.TextConfigArgs>? Text { get; set; }

        /// <summary>
        /// Ensures in-flight data remains in the region of origin during de-identification. Using this option results in a significant reduction of throughput, and is not compatible with `LOCATION` or `ORGANIZATION_NAME` infoTypes. `LOCATION` must be excluded within `TextConfig`, and must also be excluded within `ImageConfig` if image redaction is required.
        /// </summary>
        [Input("useRegionalDataProcessing")]
        public Input<bool>? UseRegionalDataProcessing { get; set; }

        public DeidentifyConfigArgs()
        {
        }
        public static new DeidentifyConfigArgs Empty => new DeidentifyConfigArgs();
    }
}
