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
    /// TransferOptions define the actions to be performed on objects in a transfer.
    /// </summary>
    public sealed class TransferOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether objects should be deleted from the source after they are transferred to the sink. **Note:** This option and delete_objects_unique_in_sink are mutually exclusive.
        /// </summary>
        [Input("deleteObjectsFromSourceAfterTransfer")]
        public Input<bool>? DeleteObjectsFromSourceAfterTransfer { get; set; }

        /// <summary>
        /// Whether objects that exist only in the sink should be deleted. **Note:** This option and delete_objects_from_source_after_transfer are mutually exclusive.
        /// </summary>
        [Input("deleteObjectsUniqueInSink")]
        public Input<bool>? DeleteObjectsUniqueInSink { get; set; }

        /// <summary>
        /// Represents the selected metadata options for a transfer job. This feature is in Preview.
        /// </summary>
        [Input("metadataOptions")]
        public Input<Inputs.MetadataOptionsArgs>? MetadataOptions { get; set; }

        /// <summary>
        /// When to overwrite objects that already exist in the sink. The default is that only objects that are different from the source are ovewritten. If true, all objects in the sink whose name matches an object in the source are overwritten with the source object.
        /// </summary>
        [Input("overwriteObjectsAlreadyExistingInSink")]
        public Input<bool>? OverwriteObjectsAlreadyExistingInSink { get; set; }

        /// <summary>
        /// When to overwrite objects that already exist in the sink. If not set overwrite behavior is determined by overwrite_objects_already_existing_in_sink.
        /// </summary>
        [Input("overwriteWhen")]
        public Input<Pulumi.GoogleNative.StorageTransfer.V1.TransferOptionsOverwriteWhen>? OverwriteWhen { get; set; }

        public TransferOptionsArgs()
        {
        }
    }
}
