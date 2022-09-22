// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1
{
    public static class GetDomainMapping
    {
        /// <summary>
        /// Get information about a domain mapping.
        /// </summary>
        public static Task<GetDomainMappingResult> InvokeAsync(GetDomainMappingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainMappingResult>("google-native:run/v1:getDomainMapping", args ?? new GetDomainMappingArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a domain mapping.
        /// </summary>
        public static Output<GetDomainMappingResult> Invoke(GetDomainMappingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainMappingResult>("google-native:run/v1:getDomainMapping", args ?? new GetDomainMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainMappingArgs : global::Pulumi.InvokeArgs
    {
        [Input("domainmappingId", required: true)]
        public string DomainmappingId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDomainMappingArgs()
        {
        }
        public static new GetDomainMappingArgs Empty => new GetDomainMappingArgs();
    }

    public sealed class GetDomainMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("domainmappingId", required: true)]
        public Input<string> DomainmappingId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDomainMappingInvokeArgs()
        {
        }
        public static new GetDomainMappingInvokeArgs Empty => new GetDomainMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainMappingResult
    {
        /// <summary>
        /// The API version for this call such as "domains.cloudrun.com/v1".
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// The kind of resource, in this case "DomainMapping".
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Metadata associated with this BuildTemplate.
        /// </summary>
        public readonly Outputs.ObjectMetaResponse Metadata;
        /// <summary>
        /// The spec for this DomainMapping.
        /// </summary>
        public readonly Outputs.DomainMappingSpecResponse Spec;
        /// <summary>
        /// The current status of the DomainMapping.
        /// </summary>
        public readonly Outputs.DomainMappingStatusResponse Status;

        [OutputConstructor]
        private GetDomainMappingResult(
            string apiVersion,

            string kind,

            Outputs.ObjectMetaResponse metadata,

            Outputs.DomainMappingSpecResponse spec,

            Outputs.DomainMappingStatusResponse status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
