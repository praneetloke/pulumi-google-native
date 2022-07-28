// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DocumentAI.V1
{
    /// <summary>
    /// Creates a processor from the type processor that the user chose. The processor will be at "ENABLED" state by default after its creation.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:documentai/v1:Processor")]
    public partial class Processor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time the processor was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The default processor version.
        /// </summary>
        [Output("defaultProcessorVersion")]
        public Output<string> DefaultProcessorVersion { get; private set; } = null!;

        /// <summary>
        /// The display name of the processor.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        /// </summary>
        [Output("kmsKeyName")]
        public Output<string> KmsKeyName { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the processor. Format: `projects/{project}/locations/{location}/processors/{processor}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The http endpoint that can be called to invoke processing.
        /// </summary>
        [Output("processEndpoint")]
        public Output<string> ProcessEndpoint { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The state of the processor.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Processor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Processor(string name, ProcessorArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:documentai/v1:Processor", name, args ?? new ProcessorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Processor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:documentai/v1:Processor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Processor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Processor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Processor(name, id, options);
        }
    }

    public sealed class ProcessorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time the processor was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The default processor version.
        /// </summary>
        [Input("defaultProcessorVersion")]
        public Input<string>? DefaultProcessorVersion { get; set; }

        /// <summary>
        /// The display name of the processor.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProcessorArgs()
        {
        }
        public static new ProcessorArgs Empty => new ProcessorArgs();
    }
}
