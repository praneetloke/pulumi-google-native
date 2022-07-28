// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// A file or directory to install on the device before the test starts.
    /// </summary>
    public sealed class RegularFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source file.
        /// </summary>
        [Input("content", required: true)]
        public Input<Inputs.FileReferenceArgs> Content { get; set; } = null!;

        /// <summary>
        /// Where to put the content on the device. Must be an absolute, allowlisted path. If the file exists, it will be replaced. The following device-side directories and any of their subdirectories are allowlisted: ${EXTERNAL_STORAGE}, /sdcard, or /storage ${ANDROID_DATA}/local/tmp, or /data/local/tmp Specifying a path outside of these directory trees is invalid. The paths /sdcard and /data will be made available and treated as implicit path substitutions. E.g. if /sdcard on a particular device does not map to external storage, the system will replace it with the external storage path prefix for that device and copy the file there. It is strongly advised to use the Environment API in app and test code to access files on the device in a portable way.
        /// </summary>
        [Input("devicePath", required: true)]
        public Input<string> DevicePath { get; set; } = null!;

        public RegularFileArgs()
        {
        }
        public static new RegularFileArgs Empty => new RegularFileArgs();
    }
}
