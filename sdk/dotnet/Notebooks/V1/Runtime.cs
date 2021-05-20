// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1
{
    /// <summary>
    /// Creates a new Runtime in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:notebooks/v1:Runtime")]
    public partial class Runtime : Pulumi.CustomResource
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// </summary>
        [Output("accessConfig")]
        public Output<Outputs.RuntimeAccessConfigResponse> AccessConfig { get; private set; } = null!;

        /// <summary>
        /// Runtime creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Runtime health_state.
        /// </summary>
        [Output("healthState")]
        public Output<string> HealthState { get; private set; } = null!;

        /// <summary>
        /// Contains Runtime daemon metrics such as Service status and JupyterLab stats.
        /// </summary>
        [Output("metrics")]
        public Output<Outputs.RuntimeMetricsResponse> Metrics { get; private set; } = null!;

        /// <summary>
        /// The resource name of the runtime. Format: `projects/{project}/locations/{location}/runtimes/{runtime}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The config settings for software inside the runtime.
        /// </summary>
        [Output("softwareConfig")]
        public Output<Outputs.RuntimeSoftwareConfigResponse> SoftwareConfig { get; private set; } = null!;

        /// <summary>
        /// Runtime state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Runtime update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// </summary>
        [Output("virtualMachine")]
        public Output<Outputs.VirtualMachineResponse> VirtualMachine { get; private set; } = null!;


        /// <summary>
        /// Create a Runtime resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Runtime(string name, RuntimeArgs args, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Runtime", name, args ?? new RuntimeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Runtime(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:notebooks/v1:Runtime", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Runtime resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Runtime Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Runtime(name, id, options);
        }
    }

    public sealed class RuntimeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// </summary>
        [Input("accessConfig")]
        public Input<Inputs.RuntimeAccessConfigArgs>? AccessConfig { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("runtimeId", required: true)]
        public Input<string> RuntimeId { get; set; } = null!;

        /// <summary>
        /// The config settings for software inside the runtime.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.RuntimeSoftwareConfigArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// </summary>
        [Input("virtualMachine")]
        public Input<Inputs.VirtualMachineArgs>? VirtualMachine { get; set; }

        public RuntimeArgs()
        {
        }
    }
}
