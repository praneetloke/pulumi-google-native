// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2
{
    public static class GetInstance
    {
        /// <summary>
        /// Gets information about an instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:bigtableadmin/v2:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:bigtableadmin/v2:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
    }

    public sealed class GetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// A server-assigned timestamp representing when this Instance was created. For instances created before this field was added (August 2021), this value is `seconds: 0, nanos: 1`.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The descriptive name for this instance as it appears in UIs. Can be changed at any time, but should be kept globally unique to avoid confusion.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. They can be used to filter resources and aggregate metrics. * Label keys must be between 1 and 63 characters long and must conform to the regular expression: `\p{Ll}\p{Lo}{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`. * No more than 64 labels can be associated with a given resource. * Keys and values must both be under 128 bytes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The unique name of the instance. Values are of the form `projects/{project}/instances/a-z+[a-z0-9]`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// The current state of the instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of the instance. Defaults to `PRODUCTION`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceResult(
            string createTime,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            bool satisfiesPzs,

            string state,

            string type)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            SatisfiesPzs = satisfiesPzs;
            State = state;
            Type = type;
        }
    }
}
