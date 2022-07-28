// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// A reference to a file.
    /// </summary>
    public sealed class FileReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URI of a file stored in Google Cloud Storage. For example: http://storage.googleapis.com/mybucket/path/to/test.xml or in gsutil format: gs://mybucket/path/to/test.xml with version-specific info, gs://mybucket/path/to/test.xml#1360383693690000 An INVALID_ARGUMENT error will be returned if the URI format is not supported. - In response: always set - In create/update request: always set
        /// </summary>
        [Input("fileUri")]
        public Input<string>? FileUri { get; set; }

        public FileReferenceArgs()
        {
        }
        public static new FileReferenceArgs Empty => new FileReferenceArgs();
    }
}
