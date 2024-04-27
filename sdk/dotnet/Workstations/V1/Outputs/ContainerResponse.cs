// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1.Outputs
{

    /// <summary>
    /// A Docker container.
    /// </summary>
    [OutputType]
    public sealed class ContainerResponse
    {
        /// <summary>
        /// Optional. Arguments passed to the entrypoint.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Optional. If set, overrides the default ENTRYPOINT specified by the image.
        /// </summary>
        public readonly ImmutableArray<string> Command;
        /// <summary>
        /// Optional. Environment variables passed to the container's entrypoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Env;
        /// <summary>
        /// Optional. A Docker container image that defines a custom environment. Cloud Workstations provides a number of [preconfigured images](https://cloud.google.com/workstations/docs/preconfigured-base-images), but you can create your own [custom container images](https://cloud.google.com/workstations/docs/custom-container-images). If using a private image, the `host.gceInstance.serviceAccount` field must be specified in the workstation configuration. If using a custom container image, the service account must have [Artifact Registry Reader](https://cloud.google.com/artifact-registry/docs/access-control#roles) permission to pull the specified image. Otherwise, the image must be publicly accessible.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// Optional. If set, overrides the USER specified in the image with the given uid.
        /// </summary>
        public readonly int RunAsUser;
        /// <summary>
        /// Optional. If set, overrides the default DIR specified by the image.
        /// </summary>
        public readonly string WorkingDir;

        [OutputConstructor]
        private ContainerResponse(
            ImmutableArray<string> args,

            ImmutableArray<string> command,

            ImmutableDictionary<string, string> env,

            string image,

            int runAsUser,

            string workingDir)
        {
            Args = args;
            Command = command;
            Env = env;
            Image = image;
            RunAsUser = runAsUser;
            WorkingDir = workingDir;
        }
    }
}