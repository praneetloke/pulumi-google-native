// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// OCI repo configuration for a single cluster
    /// </summary>
    [OutputType]
    public sealed class ConfigManagementOciConfigResponse
    {
        /// <summary>
        /// The Google Cloud Service Account Email used for auth when secret_type is gcpServiceAccount.
        /// </summary>
        public readonly string GcpServiceAccountEmail;
        /// <summary>
        /// The absolute path of the directory that contains the local resources. Default: the root directory of the image.
        /// </summary>
        public readonly string PolicyDir;
        /// <summary>
        /// Type of secret configured for access to the Git repo.
        /// </summary>
        public readonly string SecretType;
        /// <summary>
        /// The OCI image repository URL for the package to sync from. e.g. `LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME`.
        /// </summary>
        public readonly string SyncRepo;
        /// <summary>
        /// Period in seconds between consecutive syncs. Default: 15.
        /// </summary>
        public readonly string SyncWaitSecs;

        [OutputConstructor]
        private ConfigManagementOciConfigResponse(
            string gcpServiceAccountEmail,

            string policyDir,

            string secretType,

            string syncRepo,

            string syncWaitSecs)
        {
            GcpServiceAccountEmail = gcpServiceAccountEmail;
            PolicyDir = policyDir;
            SecretType = secretType;
            SyncRepo = syncRepo;
            SyncWaitSecs = syncWaitSecs;
        }
    }
}