// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// A remote or local file.
    /// </summary>
    public sealed class OSPolicyResourceFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to false. When false, files are subject to validations based on the file type: Remote: A checksum must be specified. Cloud Storage: An object generation number must be specified.
        /// </summary>
        [Input("allowInsecure")]
        public Input<bool>? AllowInsecure { get; set; }

        /// <summary>
        /// A Cloud Storage object.
        /// </summary>
        [Input("gcs")]
        public Input<Inputs.OSPolicyResourceFileGcsArgs>? Gcs { get; set; }

        /// <summary>
        /// A local path within the VM to use.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        /// <summary>
        /// A generic remote file.
        /// </summary>
        [Input("remote")]
        public Input<Inputs.OSPolicyResourceFileRemoteArgs>? Remote { get; set; }

        public OSPolicyResourceFileArgs()
        {
        }
        public static new OSPolicyResourceFileArgs Empty => new OSPolicyResourceFileArgs();
    }
}
