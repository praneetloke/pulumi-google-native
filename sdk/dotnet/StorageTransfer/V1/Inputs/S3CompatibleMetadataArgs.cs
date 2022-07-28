// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// S3CompatibleMetadata contains the metadata fields that apply to the basic types of S3-compatible data providers.
    /// </summary>
    public sealed class S3CompatibleMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the authentication and authorization method used by the storage service. When not specified, Transfer Service will attempt to determine right auth method to use.
        /// </summary>
        [Input("authMethod")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.S3CompatibleMetadataAuthMethod>? AuthMethod { get; set; }

        /// <summary>
        /// The Listing API to use for discovering objects. When not specified, Transfer Service will attempt to determine the right API to use.
        /// </summary>
        [Input("listApi")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.S3CompatibleMetadataListApi>? ListApi { get; set; }

        /// <summary>
        /// Specifies the network protocol of the agent. When not specified, the default value of NetworkProtocol NETWORK_PROTOCOL_HTTPS is used.
        /// </summary>
        [Input("protocol")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.S3CompatibleMetadataProtocol>? Protocol { get; set; }

        /// <summary>
        /// Specifies the API request model used to call the storage service. When not specified, the default value of RequestModel REQUEST_MODEL_VIRTUAL_HOSTED_STYLE is used.
        /// </summary>
        [Input("requestModel")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.S3CompatibleMetadataRequestModel>? RequestModel { get; set; }

        public S3CompatibleMetadataArgs()
        {
        }
        public static new S3CompatibleMetadataArgs Empty => new S3CompatibleMetadataArgs();
    }
}
