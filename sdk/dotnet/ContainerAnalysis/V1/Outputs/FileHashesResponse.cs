// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// Container message for hashes of byte content of files, used in source messages to verify integrity of source input to the build.
    /// </summary>
    [OutputType]
    public sealed class FileHashesResponse
    {
        /// <summary>
        /// Collection of file hashes.
        /// </summary>
        public readonly ImmutableArray<Outputs.HashResponse> FileHash;

        [OutputConstructor]
        private FileHashesResponse(ImmutableArray<Outputs.HashResponse> fileHash)
        {
            FileHash = fileHash;
        }
    }
}
