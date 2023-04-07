// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Inputs
{

    /// <summary>
    /// Container runnable.
    /// </summary>
    public sealed class ContainerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, external network access to and from container will be blocked, containers that are with block_external_network as true can still communicate with each other, network cannot be specified in the `container.options` field.
        /// </summary>
        [Input("blockExternalNetwork")]
        public Input<bool>? BlockExternalNetwork { get; set; }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Overrides the `CMD` specified in the container. If there is an ENTRYPOINT (either in the container image or with the entrypoint field below) then commands are appended as arguments to the ENTRYPOINT.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// Overrides the `ENTRYPOINT` specified in the container.
        /// </summary>
        [Input("entrypoint")]
        public Input<string>? Entrypoint { get; set; }

        /// <summary>
        /// The URI to pull the container image from.
        /// </summary>
        [Input("imageUri")]
        public Input<string>? ImageUri { get; set; }

        /// <summary>
        /// Arbitrary additional options to include in the "docker run" command when running this container, e.g. "--network host".
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Optional password for logging in to a docker registry. If password matches `projects/*/secrets/*/versions/*` then Batch will read the password from the Secret Manager;
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Optional username for logging in to a docker registry. If username matches `projects/*/secrets/*/versions/*` then Batch will read the username from the Secret Manager.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        [Input("volumes")]
        private InputList<string>? _volumes;

        /// <summary>
        /// Volumes to mount (bind mount) from the host machine files or directories into the container, formatted to match docker run's --volume option, e.g. /foo:/bar, or /foo:/bar:ro
        /// </summary>
        public InputList<string> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<string>());
            set => _volumes = value;
        }

        public ContainerArgs()
        {
        }
        public static new ContainerArgs Empty => new ContainerArgs();
    }
}
