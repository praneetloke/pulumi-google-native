// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1
{
    public static class GetRuntime
    {
        /// <summary>
        /// Gets details of a single Runtime. The location must be a regional endpoint rather than zonal.
        /// </summary>
        public static Task<GetRuntimeResult> InvokeAsync(GetRuntimeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRuntimeResult>("google-native:notebooks/v1:getRuntime", args ?? new GetRuntimeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Runtime. The location must be a regional endpoint rather than zonal.
        /// </summary>
        public static Output<GetRuntimeResult> Invoke(GetRuntimeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRuntimeResult>("google-native:notebooks/v1:getRuntime", args ?? new GetRuntimeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRuntimeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("runtimeId", required: true)]
        public string RuntimeId { get; set; } = null!;

        public GetRuntimeArgs()
        {
        }
        public static new GetRuntimeArgs Empty => new GetRuntimeArgs();
    }

    public sealed class GetRuntimeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("runtimeId", required: true)]
        public Input<string> RuntimeId { get; set; } = null!;

        public GetRuntimeInvokeArgs()
        {
        }
        public static new GetRuntimeInvokeArgs Empty => new GetRuntimeInvokeArgs();
    }


    [OutputType]
    public sealed class GetRuntimeResult
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// </summary>
        public readonly Outputs.RuntimeAccessConfigResponse AccessConfig;
        /// <summary>
        /// Runtime creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Runtime health_state.
        /// </summary>
        public readonly string HealthState;
        /// <summary>
        /// Optional. The labels to associate with this Managed Notebook or Runtime. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Contains Runtime daemon metrics such as Service status and JupyterLab stats.
        /// </summary>
        public readonly Outputs.RuntimeMetricsResponse Metrics;
        /// <summary>
        /// The resource name of the runtime. Format: `projects/{project}/locations/{location}/runtimes/{runtimeId}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The config settings for software inside the runtime.
        /// </summary>
        public readonly Outputs.RuntimeSoftwareConfigResponse SoftwareConfig;
        /// <summary>
        /// Runtime state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Runtime update time.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// </summary>
        public readonly Outputs.VirtualMachineResponse VirtualMachine;

        [OutputConstructor]
        private GetRuntimeResult(
            Outputs.RuntimeAccessConfigResponse accessConfig,

            string createTime,

            string healthState,

            ImmutableDictionary<string, string> labels,

            Outputs.RuntimeMetricsResponse metrics,

            string name,

            Outputs.RuntimeSoftwareConfigResponse softwareConfig,

            string state,

            string updateTime,

            Outputs.VirtualMachineResponse virtualMachine)
        {
            AccessConfig = accessConfig;
            CreateTime = createTime;
            HealthState = healthState;
            Labels = labels;
            Metrics = metrics;
            Name = name;
            SoftwareConfig = softwareConfig;
            State = state;
            UpdateTime = updateTime;
            VirtualMachine = virtualMachine;
        }
    }
}
