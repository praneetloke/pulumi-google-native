// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha
{
    public static class GetFleet
    {
        /// <summary>
        /// Returns the details of a fleet.
        /// </summary>
        public static Task<GetFleetResult> InvokeAsync(GetFleetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFleetResult>("google-native:gkehub/v1alpha:getFleet", args ?? new GetFleetArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the details of a fleet.
        /// </summary>
        public static Output<GetFleetResult> Invoke(GetFleetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFleetResult>("google-native:gkehub/v1alpha:getFleet", args ?? new GetFleetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFleetArgs : global::Pulumi.InvokeArgs
    {
        [Input("fleetId", required: true)]
        public string FleetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFleetArgs()
        {
        }
        public static new GetFleetArgs Empty => new GetFleetArgs();
    }

    public sealed class GetFleetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("fleetId", required: true)]
        public Input<string> FleetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFleetInvokeArgs()
        {
        }
        public static new GetFleetInvokeArgs Empty => new GetFleetInvokeArgs();
    }


    [OutputType]
    public sealed class GetFleetResult
    {
        /// <summary>
        /// When the Fleet was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// When the Fleet was deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The name for the fleet. The name must meet the following constraints: + The name of a fleet should be unique within the organization; + It must consist of lower case alphanumeric characters or `-`; + The length of the name must be less than or equal to 63; + Unicode names must be expressed in Punycode format (rfc3492). Examples: + prod-fleet + xn--wlq33vhyw9jb （Punycode form for "生产环境")
        /// </summary>
        public readonly string FleetName;
        /// <summary>
        /// The full, unique resource name of this fleet in the format of `projects/{project}/locations/{location}/fleets/{fleet}`. Each GCP project can have at most one fleet resource, named "default".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Google-generated UUID for this resource. This is unique across all Fleet resources. If a Fleet resource is deleted and another resource with the same name is created, it gets a different uid.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// When the Fleet was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetFleetResult(
            string createTime,

            string deleteTime,

            string displayName,

            string fleetName,

            string name,

            string uid,

            string updateTime)
        {
            CreateTime = createTime;
            DeleteTime = deleteTime;
            DisplayName = displayName;
            FleetName = fleetName;
            Name = name;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
