// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1.Inputs
{

    /// <summary>
    /// ApplicationEndpoint represents a remote application endpoint.
    /// </summary>
    public sealed class GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname or IP address of the remote application endpoint.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Port of the remote application endpoint.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointArgs()
        {
        }
        public static new GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointArgs Empty => new GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointArgs();
    }
}
