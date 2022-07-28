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
    /// Indicates the location at which a package was found.
    /// </summary>
    public sealed class FileLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For jars that are contained inside .war files, this filepath can indicate the path to war file combined with the path to jar file.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        public FileLocationArgs()
        {
        }
        public static new FileLocationArgs Empty => new FileLocationArgs();
    }
}
