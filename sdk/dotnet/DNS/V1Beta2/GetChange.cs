// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2
{
    public static class GetChange
    {
        /// <summary>
        /// Fetches the representation of an existing Change.
        /// </summary>
        public static Task<GetChangeResult> InvokeAsync(GetChangeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChangeResult>("google-native:dns/v1beta2:getChange", args ?? new GetChangeArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches the representation of an existing Change.
        /// </summary>
        public static Output<GetChangeResult> Invoke(GetChangeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChangeResult>("google-native:dns/v1beta2:getChange", args ?? new GetChangeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChangeArgs : global::Pulumi.InvokeArgs
    {
        [Input("changeId", required: true)]
        public string ChangeId { get; set; } = null!;

        [Input("clientOperationId")]
        public string? ClientOperationId { get; set; }

        [Input("managedZone", required: true)]
        public string ManagedZone { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetChangeArgs()
        {
        }
        public static new GetChangeArgs Empty => new GetChangeArgs();
    }

    public sealed class GetChangeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("changeId", required: true)]
        public Input<string> ChangeId { get; set; } = null!;

        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetChangeInvokeArgs()
        {
        }
        public static new GetChangeInvokeArgs Empty => new GetChangeInvokeArgs();
    }


    [OutputType]
    public sealed class GetChangeResult
    {
        /// <summary>
        /// Which ResourceRecordSets to add?
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceRecordSetResponse> Additions;
        /// <summary>
        /// Which ResourceRecordSets to remove? Must match existing data exactly.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceRecordSetResponse> Deletions;
        /// <summary>
        /// If the DNS queries for the zone will be served.
        /// </summary>
        public readonly bool IsServing;
        public readonly string Kind;
        /// <summary>
        /// The time that this operation was started by the server (output only). This is in RFC3339 text format.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetChangeResult(
            ImmutableArray<Outputs.ResourceRecordSetResponse> additions,

            ImmutableArray<Outputs.ResourceRecordSetResponse> deletions,

            bool isServing,

            string kind,

            string startTime,

            string status)
        {
            Additions = additions;
            Deletions = deletions;
            IsServing = isServing;
            Kind = kind;
            StartTime = startTime;
            Status = status;
        }
    }
}
