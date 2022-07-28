// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Binds the resources in an API proxy or remote service with the allowed REST methods and associated quota enforcement.
    /// </summary>
    public sealed class GoogleCloudApigeeV1OperationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the API proxy or remote service with which the resources, methods, and quota are associated.
        /// </summary>
        [Input("apiSource", required: true)]
        public Input<string> ApiSource { get; set; } = null!;

        [Input("attributes")]
        private InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>? _attributes;

        /// <summary>
        /// Custom attributes associated with the operation.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1AttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>());
            set => _attributes = value;
        }

        [Input("operations")]
        private InputList<Inputs.GoogleCloudApigeeV1OperationArgs>? _operations;

        /// <summary>
        /// List of resource/method pairs for the API proxy or remote service to which quota will applied. **Note**: Currently, you can specify only a single resource/method pair. The call will fail if more than one resource/method pair is provided.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1OperationArgs> Operations
        {
            get => _operations ?? (_operations = new InputList<Inputs.GoogleCloudApigeeV1OperationArgs>());
            set => _operations = value;
        }

        /// <summary>
        /// Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.
        /// </summary>
        [Input("quota")]
        public Input<Inputs.GoogleCloudApigeeV1QuotaArgs>? Quota { get; set; }

        public GoogleCloudApigeeV1OperationConfigArgs()
        {
        }
        public static new GoogleCloudApigeeV1OperationConfigArgs Empty => new GoogleCloudApigeeV1OperationConfigArgs();
    }
}
