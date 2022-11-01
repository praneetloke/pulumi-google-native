// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Inputs
{

    /// <summary>
    /// A table representation similar to HTML table structure.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentPageTableArgs : global::Pulumi.ResourceArgs
    {
        [Input("bodyRows")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs>? _bodyRows;

        /// <summary>
        /// Body rows of the table.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs> BodyRows
        {
            get => _bodyRows ?? (_bodyRows = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs>());
            set => _bodyRows = value;
        }

        [Input("detectedLanguages")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs>? _detectedLanguages;

        /// <summary>
        /// A list of detected languages together with confidence.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs> DetectedLanguages
        {
            get => _detectedLanguages ?? (_detectedLanguages = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs>());
            set => _detectedLanguages = value;
        }

        [Input("headerRows")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs>? _headerRows;

        /// <summary>
        /// Header rows of the table.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs> HeaderRows
        {
            get => _headerRows ?? (_headerRows = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableTableRowArgs>());
            set => _headerRows = value;
        }

        /// <summary>
        /// Layout for Table.
        /// </summary>
        [Input("layout")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageLayoutArgs>? Layout { get; set; }

        /// <summary>
        /// The history of this table.
        /// </summary>
        [Input("provenance")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentProvenanceArgs>? Provenance { get; set; }

        public GoogleCloudDocumentaiV1DocumentPageTableArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentPageTableArgs Empty => new GoogleCloudDocumentaiV1DocumentPageTableArgs();
    }
}