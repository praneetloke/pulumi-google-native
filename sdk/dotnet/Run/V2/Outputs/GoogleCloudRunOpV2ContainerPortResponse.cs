// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Outputs
{

    /// <summary>
    /// ContainerPort represents a network port in a single container.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRunOpV2ContainerPortResponse
    {
        /// <summary>
        /// Port number the container listens on. This must be a valid TCP port number, 0 &lt; container_port &lt; 65536.
        /// </summary>
        public readonly int ContainerPort;
        /// <summary>
        /// If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GoogleCloudRunOpV2ContainerPortResponse(
            int containerPort,

            string name)
        {
            ContainerPort = containerPort;
            Name = name;
        }
    }
}
