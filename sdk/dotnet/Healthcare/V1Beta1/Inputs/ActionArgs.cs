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
    /// Specifies a selection of tags and an `Action` to apply to each one.
    /// </summary>
    public sealed class ActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inspect image and transform sensitive burnt-in text. Doesn't apply to elements nested in a sequence, which revert to `Keep`. Supported [tags](http://dicom.nema.org/medical/dicom/2018e/output/chtml/part06/chapter_6.html): PixelData
        /// </summary>
        [Input("cleanImageTag")]
        public Input<Inputs.ImageConfigArgs>? CleanImageTag { get; set; }

        /// <summary>
        /// Inspect text and transform sensitive text. Configurable via TextConfig. Supported Value Representations: AE, LO, LT, PN, SH, ST, UC, UT, DA, DT, AS
        /// </summary>
        [Input("cleanTextTag")]
        public Input<Inputs.CleanTextTagArgs>? CleanTextTag { get; set; }

        /// <summary>
        /// Delete tag.
        /// </summary>
        [Input("deleteTag")]
        public Input<Inputs.DeleteTagArgs>? DeleteTag { get; set; }

        /// <summary>
        /// Keep tag unchanged.
        /// </summary>
        [Input("keepTag")]
        public Input<Inputs.KeepTagArgs>? KeepTag { get; set; }

        [Input("queries")]
        private InputList<string>? _queries;

        /// <summary>
        /// Select all tags with the listed tag IDs, names, or Value Representations (VRs). Examples: ID: "00100010" Keyword: "PatientName" VR: "PN"
        /// </summary>
        public InputList<string> Queries
        {
            get => _queries ?? (_queries = new InputList<string>());
            set => _queries = value;
        }

        /// <summary>
        /// Recursively apply DICOM de-id to tags nested in a sequence. Supported [Value Representation] (http://dicom.nema.org/medical/dicom/2018e/output/chtml/part05/sect_6.2.html#table_6.2-1): SQ
        /// </summary>
        [Input("recurseTag")]
        public Input<Inputs.RecurseTagArgs>? RecurseTag { get; set; }

        /// <summary>
        /// Replace UID with a new generated UID. Supported [Value Representation] (http://dicom.nema.org/medical/dicom/2018e/output/chtml/part05/sect_6.2.html#table_6.2-1): UI
        /// </summary>
        [Input("regenUidTag")]
        public Input<Inputs.RegenUidTagArgs>? RegenUidTag { get; set; }

        /// <summary>
        /// Replace with empty tag.
        /// </summary>
        [Input("removeTag")]
        public Input<Inputs.RemoveTagArgs>? RemoveTag { get; set; }

        /// <summary>
        /// Reset tag to a placeholder value.
        /// </summary>
        [Input("resetTag")]
        public Input<Inputs.ResetTagArgs>? ResetTag { get; set; }

        public ActionArgs()
        {
        }
        public static new ActionArgs Empty => new ActionArgs();
    }
}
