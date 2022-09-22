// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1
{
    public static class GetHub
    {
        /// <summary>
        /// Gets details about a Network Connectivity Center hub.
        /// </summary>
        public static Task<GetHubResult> InvokeAsync(GetHubArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHubResult>("google-native:networkconnectivity/v1alpha1:getHub", args ?? new GetHubArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details about a Network Connectivity Center hub.
        /// </summary>
        public static Output<GetHubResult> Invoke(GetHubInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHubResult>("google-native:networkconnectivity/v1alpha1:getHub", args ?? new GetHubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHubArgs : global::Pulumi.InvokeArgs
    {
        [Input("hubId", required: true)]
        public string HubId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetHubArgs()
        {
        }
        public static new GetHubArgs Empty => new GetHubArgs();
    }

    public sealed class GetHubInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("hubId", required: true)]
        public Input<string> HubId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetHubInvokeArgs()
        {
        }
        public static new GetHubInvokeArgs Empty => new GetHubInvokeArgs();
    }


    [OutputType]
    public sealed class GetHubResult
    {
        /// <summary>
        /// Time when the Hub was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Short description of the hub resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-defined labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of a Hub resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of the URIs of all attached spokes. This field is deprecated and will not be included in future API versions. Call ListSpokes on each region instead.
        /// </summary>
        public readonly ImmutableArray<string> Spokes;
        /// <summary>
        /// The current lifecycle state of this Hub.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Google-generated UUID for this resource. This is unique across all Hub resources. If a Hub resource is deleted and another with the same name is created, it gets a different unique_id.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// Time when the Hub was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetHubResult(
            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<string> spokes,

            string state,

            string uniqueId,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            Spokes = spokes;
            State = state;
            UniqueId = uniqueId;
            UpdateTime = updateTime;
        }
    }
}
