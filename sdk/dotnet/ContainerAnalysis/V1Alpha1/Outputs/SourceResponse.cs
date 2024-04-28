// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Source describes the location of the source used for the build.
    /// </summary>
    [OutputType]
    public sealed class SourceResponse
    {
        /// <summary>
        /// If provided, some of the source code used for the build may be found in these locations, in the case where the source repository had multiple remotes or submodules. This list will not include the context specified in the context field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleDevtoolsContaineranalysisV1alpha1SourceContextResponse> AdditionalContexts;
        /// <summary>
        /// If provided, the input binary artifacts for the build came from this location.
        /// </summary>
        public readonly Outputs.StorageSourceResponse ArtifactStorageSource;
        /// <summary>
        /// If provided, the source code used for the build came from this location.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1SourceContextResponse Context;
        /// <summary>
        /// Hash(es) of the build source, which can be used to verify that the original source integrity was maintained in the build. The keys to this map are file paths used as build source and the values contain the hash values for those files. If the build source came in a single package such as a gzipped tarfile (.tar.gz), the FileHash will be for the single path to that file.
        /// </summary>
        public readonly Outputs.FileHashesResponse FileHashes;
        /// <summary>
        /// If provided, get source from this location in a Cloud Repo.
        /// </summary>
        public readonly Outputs.RepoSourceResponse RepoSource;
        /// <summary>
        /// If provided, get the source from this location in Google Cloud Storage.
        /// </summary>
        public readonly Outputs.StorageSourceResponse StorageSource;

        [OutputConstructor]
        private SourceResponse(
            ImmutableArray<Outputs.GoogleDevtoolsContaineranalysisV1alpha1SourceContextResponse> additionalContexts,

            Outputs.StorageSourceResponse artifactStorageSource,

            Outputs.GoogleDevtoolsContaineranalysisV1alpha1SourceContextResponse context,

            Outputs.FileHashesResponse fileHashes,

            Outputs.RepoSourceResponse repoSource,

            Outputs.StorageSourceResponse storageSource)
        {
            AdditionalContexts = additionalContexts;
            ArtifactStorageSource = artifactStorageSource;
            Context = context;
            FileHashes = fileHashes;
            RepoSource = repoSource;
            StorageSource = storageSource;
        }
    }
}
