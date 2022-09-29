// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// TCPSocketAction describes an action based on opening a socket
    /// </summary>
    public sealed class GoogleCloudRunV2TCPSocketActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host name to connect to, defaults to the pod IP.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. This field is currently limited to integer types only because of proto's inability to properly support the IntOrString golang type.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public GoogleCloudRunV2TCPSocketActionArgs()
        {
        }
        public static new GoogleCloudRunV2TCPSocketActionArgs Empty => new GoogleCloudRunV2TCPSocketActionArgs();
    }
}
