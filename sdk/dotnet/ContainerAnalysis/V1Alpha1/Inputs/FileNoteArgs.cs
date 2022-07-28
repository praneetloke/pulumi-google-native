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
    /// FileNote represents an SPDX File Information section: https://spdx.github.io/spdx-spec/4-file-information/
    /// </summary>
    public sealed class FileNoteArgs : global::Pulumi.ResourceArgs
    {
        [Input("checksum")]
        private InputList<string>? _checksum;

        /// <summary>
        /// Provide a unique identifier to match analysis information on each specific file in a package
        /// </summary>
        public InputList<string> Checksum
        {
            get => _checksum ?? (_checksum = new InputList<string>());
            set => _checksum = value;
        }

        /// <summary>
        /// This field provides information about the type of file identified
        /// </summary>
        [Input("fileType")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.FileNoteFileType>? FileType { get; set; }

        /// <summary>
        /// Identify the full path and filename that corresponds to the file information in this section
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public FileNoteArgs()
        {
        }
        public static new FileNoteArgs Empty => new FileNoteArgs();
    }
}
