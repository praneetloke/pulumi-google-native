// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// A configuration for running an Apache Spark (http://spark.apache.org/) batch workload.
    /// </summary>
    public sealed class SparkBatchArgs : Pulumi.ResourceArgs
    {
        [Input("archiveUris")]
        private InputList<string>? _archiveUris;

        /// <summary>
        /// Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public InputList<string> ArchiveUris
        {
            get => _archiveUris ?? (_archiveUris = new InputList<string>());
            set => _archiveUris = value;
        }

        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// Optional. The arguments to pass to the driver. Do not include arguments that can be set as batch properties, such as --conf, since a collision can occur that causes an incorrect batch submission.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("fileUris")]
        private InputList<string>? _fileUris;

        /// <summary>
        /// Optional. HCFS URIs of files to be placed in the working directory of each executor.
        /// </summary>
        public InputList<string> FileUris
        {
            get => _fileUris ?? (_fileUris = new InputList<string>());
            set => _fileUris = value;
        }

        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// Optional. HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        /// <summary>
        /// Optional. The name of the driver main class. The jar file that contains the class must be in the classpath or specified in jar_file_uris.
        /// </summary>
        [Input("mainClass")]
        public Input<string>? MainClass { get; set; }

        /// <summary>
        /// Optional. The HCFS URI of the jar file that contains the main class.
        /// </summary>
        [Input("mainJarFileUri")]
        public Input<string>? MainJarFileUri { get; set; }

        public SparkBatchArgs()
        {
        }
    }
}
