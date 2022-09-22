// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3
{
    public static class GetTemplate
    {
        /// <summary>
        /// Get the template associated with a template.
        /// </summary>
        public static Task<GetTemplateResult> InvokeAsync(GetTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateResult>("google-native:dataflow/v1b3:getTemplate", args ?? new GetTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Get the template associated with a template.
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("google-native:dataflow/v1b3:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("gcsPath", required: true)]
        public string GcsPath { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetTemplateArgs()
        {
        }
        public static new GetTemplateArgs Empty => new GetTemplateArgs();
    }

    public sealed class GetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("gcsPath", required: true)]
        public Input<string> GcsPath { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetTemplateInvokeArgs()
        {
        }
        public static new GetTemplateInvokeArgs Empty => new GetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateResult
    {
        /// <summary>
        /// The template metadata describing the template name, available parameters, etc.
        /// </summary>
        public readonly Outputs.TemplateMetadataResponse Metadata;
        /// <summary>
        /// Describes the runtime metadata with SDKInfo and available parameters.
        /// </summary>
        public readonly Outputs.RuntimeMetadataResponse RuntimeMetadata;
        /// <summary>
        /// The status of the get template request. Any problems with the request will be indicated in the error_details.
        /// </summary>
        public readonly Outputs.StatusResponse Status;
        /// <summary>
        /// Template Type.
        /// </summary>
        public readonly string TemplateType;

        [OutputConstructor]
        private GetTemplateResult(
            Outputs.TemplateMetadataResponse metadata,

            Outputs.RuntimeMetadataResponse runtimeMetadata,

            Outputs.StatusResponse status,

            string templateType)
        {
            Metadata = metadata;
            RuntimeMetadata = runtimeMetadata;
            Status = status;
            TemplateType = templateType;
        }
    }
}
