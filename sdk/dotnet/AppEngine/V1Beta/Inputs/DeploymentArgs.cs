// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Code and application artifacts used to deploy a version to App Engine.
    /// </summary>
    public sealed class DeploymentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google Cloud Build build information. Only applicable for instances running in the App Engine flexible environment.
        /// </summary>
        [Input("build")]
        public Input<Inputs.BuildInfoArgs>? Build { get; set; }

        /// <summary>
        /// Options for any Google Cloud Build builds created as a part of this deployment.These options will only be used if a new build is created, such as when deploying to the App Engine flexible environment using files or zip.
        /// </summary>
        [Input("cloudBuildOptions")]
        public Input<Inputs.CloudBuildOptionsArgs>? CloudBuildOptions { get; set; }

        /// <summary>
        /// The Docker image for the container that runs the version. Only applicable for instances running in the App Engine flexible environment.
        /// </summary>
        [Input("container")]
        public Input<Inputs.ContainerInfoArgs>? Container { get; set; }

        /// <summary>
        /// Manifest of the files stored in Google Cloud Storage that are included as part of this version. All files must be readable using the credentials supplied with this call.
        /// </summary>
        [Input("files")]
        public Input<Inputs.FileInfoArgs>? Files { get; set; }

        /// <summary>
        /// The zip file for this deployment, if this is a zip deployment.
        /// </summary>
        [Input("zip")]
        public Input<Inputs.ZipInfoArgs>? Zip { get; set; }

        public DeploymentArgs()
        {
        }
        public static new DeploymentArgs Empty => new DeploymentArgs();
    }
}
