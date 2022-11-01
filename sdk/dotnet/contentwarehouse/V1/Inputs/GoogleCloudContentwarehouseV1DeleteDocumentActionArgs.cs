// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Inputs
{

    /// <summary>
    /// Represents the action responsible for deleting the document.
    /// </summary>
    public sealed class GoogleCloudContentwarehouseV1DeleteDocumentActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean field to select between hard vs soft delete options. Set 'true' for 'hard delete' and 'false' for 'soft delete'.
        /// </summary>
        [Input("enableHardDelete")]
        public Input<bool>? EnableHardDelete { get; set; }

        public GoogleCloudContentwarehouseV1DeleteDocumentActionArgs()
        {
        }
        public static new GoogleCloudContentwarehouseV1DeleteDocumentActionArgs Empty => new GoogleCloudContentwarehouseV1DeleteDocumentActionArgs();
    }
}