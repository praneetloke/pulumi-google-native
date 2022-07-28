// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1
{
    public static class GetCertificateMap
    {
        /// <summary>
        /// Gets details of a single CertificateMap.
        /// </summary>
        public static Task<GetCertificateMapResult> InvokeAsync(GetCertificateMapArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateMapResult>("google-native:certificatemanager/v1:getCertificateMap", args ?? new GetCertificateMapArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single CertificateMap.
        /// </summary>
        public static Output<GetCertificateMapResult> Invoke(GetCertificateMapInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCertificateMapResult>("google-native:certificatemanager/v1:getCertificateMap", args ?? new GetCertificateMapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateMapArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateMapId", required: true)]
        public string CertificateMapId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateMapArgs()
        {
        }
        public static new GetCertificateMapArgs Empty => new GetCertificateMapArgs();
    }

    public sealed class GetCertificateMapInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateMapId", required: true)]
        public Input<string> CertificateMapId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateMapInvokeArgs()
        {
        }
        public static new GetCertificateMapInvokeArgs Empty => new GetCertificateMapInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateMapResult
    {
        /// <summary>
        /// The creation timestamp of a Certificate Map.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// One or more paragraphs of text description of a certificate map.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A list of GCLB targets which use this Certificate Map.
        /// </summary>
        public readonly ImmutableArray<Outputs.GclbTargetResponse> GclbTargets;
        /// <summary>
        /// Set of labels associated with a Certificate Map.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*/locations/*/certificateMaps/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The update timestamp of a Certificate Map.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetCertificateMapResult(
            string createTime,

            string description,

            ImmutableArray<Outputs.GclbTargetResponse> gclbTargets,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            GclbTargets = gclbTargets;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
