// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1
{
    public static class GetServerTlsPolicy
    {
        /// <summary>
        /// Gets details of a single ServerTlsPolicy.
        /// </summary>
        public static Task<GetServerTlsPolicyResult> InvokeAsync(GetServerTlsPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerTlsPolicyResult>("google-native:networksecurity/v1beta1:getServerTlsPolicy", args ?? new GetServerTlsPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single ServerTlsPolicy.
        /// </summary>
        public static Output<GetServerTlsPolicyResult> Invoke(GetServerTlsPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServerTlsPolicyResult>("google-native:networksecurity/v1beta1:getServerTlsPolicy", args ?? new GetServerTlsPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerTlsPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serverTlsPolicyId", required: true)]
        public string ServerTlsPolicyId { get; set; } = null!;

        public GetServerTlsPolicyArgs()
        {
        }
        public static new GetServerTlsPolicyArgs Empty => new GetServerTlsPolicyArgs();
    }

    public sealed class GetServerTlsPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serverTlsPolicyId", required: true)]
        public Input<string> ServerTlsPolicyId { get; set; } = null!;

        public GetServerTlsPolicyInvokeArgs()
        {
        }
        public static new GetServerTlsPolicyInvokeArgs Empty => new GetServerTlsPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerTlsPolicyResult
    {
        /// <summary>
        ///  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
        /// </summary>
        public readonly bool AllowOpen;
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Set of label tags associated with the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        ///  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        /// </summary>
        public readonly Outputs.MTLSPolicyResponse MtlsPolicy;
        /// <summary>
        /// Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        ///  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        /// </summary>
        public readonly Outputs.GoogleCloudNetworksecurityV1beta1CertificateProviderResponse ServerCertificate;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetServerTlsPolicyResult(
            bool allowOpen,

            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            Outputs.MTLSPolicyResponse mtlsPolicy,

            string name,

            Outputs.GoogleCloudNetworksecurityV1beta1CertificateProviderResponse serverCertificate,

            string updateTime)
        {
            AllowOpen = allowOpen;
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            MtlsPolicy = mtlsPolicy;
            Name = name;
            ServerCertificate = serverCertificate;
            UpdateTime = updateTime;
        }
    }
}
