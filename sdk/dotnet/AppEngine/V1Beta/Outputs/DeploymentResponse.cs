// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// Code and application artifacts used to deploy a version to App Engine.
    /// </summary>
    [OutputType]
    public sealed class DeploymentResponse
    {
        /// <summary>
        /// Google Cloud Build build information. Only applicable for instances running in the App Engine flexible environment.
        /// </summary>
        public readonly Outputs.BuildInfoResponse Build;
        /// <summary>
        /// Options for any Google Cloud Build builds created as a part of this deployment.These options will only be used if a new build is created, such as when deploying to the App Engine flexible environment using files or zip.
        /// </summary>
        public readonly Outputs.CloudBuildOptionsResponse CloudBuildOptions;
        /// <summary>
        /// The Docker image for the container that runs the version. Only applicable for instances running in the App Engine flexible environment.
        /// </summary>
        public readonly Outputs.ContainerInfoResponse Container;
        /// <summary>
        /// Manifest of the files stored in Google Cloud Storage that are included as part of this version. All files must be readable using the credentials supplied with this call.
        /// </summary>
        public readonly Outputs.FileInfoResponse Files;
        /// <summary>
        /// The zip file for this deployment, if this is a zip deployment.
        /// </summary>
        public readonly Outputs.ZipInfoResponse Zip;

        [OutputConstructor]
        private DeploymentResponse(
            Outputs.BuildInfoResponse build,

            Outputs.CloudBuildOptionsResponse cloudBuildOptions,

            Outputs.ContainerInfoResponse container,

            Outputs.FileInfoResponse files,

            Outputs.ZipInfoResponse zip)
        {
            Build = build;
            CloudBuildOptions = cloudBuildOptions;
            Container = container;
            Files = files;
            Zip = zip;
        }
    }
}
