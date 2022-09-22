// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetTargetGrpcProxy
    {
        /// <summary>
        /// Returns the specified TargetGrpcProxy resource in the given scope.
        /// </summary>
        public static Task<GetTargetGrpcProxyResult> InvokeAsync(GetTargetGrpcProxyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTargetGrpcProxyResult>("google-native:compute/alpha:getTargetGrpcProxy", args ?? new GetTargetGrpcProxyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified TargetGrpcProxy resource in the given scope.
        /// </summary>
        public static Output<GetTargetGrpcProxyResult> Invoke(GetTargetGrpcProxyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetGrpcProxyResult>("google-native:compute/alpha:getTargetGrpcProxy", args ?? new GetTargetGrpcProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetGrpcProxyArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("targetGrpcProxy", required: true)]
        public string TargetGrpcProxy { get; set; } = null!;

        public GetTargetGrpcProxyArgs()
        {
        }
        public static new GetTargetGrpcProxyArgs Empty => new GetTargetGrpcProxyArgs();
    }

    public sealed class GetTargetGrpcProxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetGrpcProxy", required: true)]
        public Input<string> TargetGrpcProxy { get; set; } = null!;

        public GetTargetGrpcProxyInvokeArgs()
        {
        }
        public static new GetTargetGrpcProxyInvokeArgs Empty => new GetTargetGrpcProxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetTargetGrpcProxyResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetGrpcProxy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Type of the resource. Always compute#targetGrpcProxy for target grpc proxies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL with id for the resource.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
        /// </summary>
        public readonly string UrlMap;
        /// <summary>
        /// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
        /// </summary>
        public readonly bool ValidateForProxyless;

        [OutputConstructor]
        private GetTargetGrpcProxyResult(
            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string name,

            string selfLink,

            string selfLinkWithId,

            string urlMap,

            bool validateForProxyless)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            UrlMap = urlMap;
            ValidateForProxyless = validateForProxyless;
        }
    }
}
