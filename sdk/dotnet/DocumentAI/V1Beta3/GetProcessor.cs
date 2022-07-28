// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DocumentAI.V1Beta3
{
    public static class GetProcessor
    {
        /// <summary>
        /// Gets a processor detail.
        /// </summary>
        public static Task<GetProcessorResult> InvokeAsync(GetProcessorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProcessorResult>("google-native:documentai/v1beta3:getProcessor", args ?? new GetProcessorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a processor detail.
        /// </summary>
        public static Output<GetProcessorResult> Invoke(GetProcessorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProcessorResult>("google-native:documentai/v1beta3:getProcessor", args ?? new GetProcessorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProcessorArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("processorId", required: true)]
        public string ProcessorId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetProcessorArgs()
        {
        }
        public static new GetProcessorArgs Empty => new GetProcessorArgs();
    }

    public sealed class GetProcessorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("processorId", required: true)]
        public Input<string> ProcessorId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetProcessorInvokeArgs()
        {
        }
        public static new GetProcessorInvokeArgs Empty => new GetProcessorInvokeArgs();
    }


    [OutputType]
    public sealed class GetProcessorResult
    {
        /// <summary>
        /// The time the processor was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The default processor version.
        /// </summary>
        public readonly string DefaultProcessorVersion;
        /// <summary>
        /// The display name of the processor.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// Immutable. The resource name of the processor. Format: `projects/{project}/locations/{location}/processors/{processor}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The http endpoint that can be called to invoke processing.
        /// </summary>
        public readonly string ProcessEndpoint;
        /// <summary>
        /// The state of the processor.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProcessorResult(
            string createTime,

            string defaultProcessorVersion,

            string displayName,

            string kmsKeyName,

            string name,

            string processEndpoint,

            string state,

            string type)
        {
            CreateTime = createTime;
            DefaultProcessorVersion = defaultProcessorVersion;
            DisplayName = displayName;
            KmsKeyName = kmsKeyName;
            Name = name;
            ProcessEndpoint = processEndpoint;
            State = state;
            Type = type;
        }
    }
}
