// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2
{
    public static class GetInstance
    {
        /// <summary>
        /// Get details about a single server.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:baremetalsolution/v2:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Get details about a single server.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:baremetalsolution/v2:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// Create a time stamp.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The firmware version for the instance.
        /// </summary>
        public readonly string FirmwareVersion;
        /// <summary>
        /// True if you enable hyperthreading for the server, otherwise false. The default value is false.
        /// </summary>
        public readonly bool HyperthreadingEnabled;
        /// <summary>
        /// True if the interactive serial console feature is enabled for the instance, false otherwise. The default value is false.
        /// </summary>
        public readonly bool InteractiveSerialConsoleEnabled;
        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// List of logical interfaces for the instance. The number of logical interfaces will be the same as number of hardware bond/nic on the chosen network template. For the non-multivlan configurations (for eg, existing servers) that use existing default network template (bondaa-bondaa), both the Instance.networks field and the Instance.logical_interfaces fields will be filled to ensure backward compatibility. For the others, only Instance.logical_interfaces will be filled.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudBaremetalsolutionV2LogicalInterfaceResponse> LogicalInterfaces;
        /// <summary>
        /// Text field about info for logging in.
        /// </summary>
        public readonly string LoginInfo;
        /// <summary>
        /// Immutable. List of LUNs associated with this server.
        /// </summary>
        public readonly ImmutableArray<Outputs.LunResponse> Luns;
        /// <summary>
        /// Immutable. The server type. [Available server types](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Immutable. The resource name of this `Instance`. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. Format: `projects/{project}/locations/{location}/instances/{instance}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Instance network template name. For eg, bondaa-bondaa, bondab-nic, etc. Generally, the template name follows the syntax of "bond" or "nic".
        /// </summary>
        public readonly string NetworkTemplate;
        /// <summary>
        /// List of networks associated with this server.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkResponse> Networks;
        /// <summary>
        /// The OS image currently installed on the server.
        /// </summary>
        public readonly string OsImage;
        /// <summary>
        /// Immutable. Pod name. Pod is an independent part of infrastructure. Instance can be connected to the assets (networks, volumes) allocated in the same pod only.
        /// </summary>
        public readonly string Pod;
        /// <summary>
        /// The state of the server.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Update a time stamp.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Input only. List of Volumes to attach to this Instance on creation. This field won't be populated in Get/List responses.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponse> Volumes;
        /// <summary>
        /// The workload profile for the instance.
        /// </summary>
        public readonly string WorkloadProfile;

        [OutputConstructor]
        private GetInstanceResult(
            string createTime,

            string firmwareVersion,

            bool hyperthreadingEnabled,

            bool interactiveSerialConsoleEnabled,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GoogleCloudBaremetalsolutionV2LogicalInterfaceResponse> logicalInterfaces,

            string loginInfo,

            ImmutableArray<Outputs.LunResponse> luns,

            string machineType,

            string name,

            string networkTemplate,

            ImmutableArray<Outputs.NetworkResponse> networks,

            string osImage,

            string pod,

            string state,

            string updateTime,

            ImmutableArray<Outputs.VolumeResponse> volumes,

            string workloadProfile)
        {
            CreateTime = createTime;
            FirmwareVersion = firmwareVersion;
            HyperthreadingEnabled = hyperthreadingEnabled;
            InteractiveSerialConsoleEnabled = interactiveSerialConsoleEnabled;
            Labels = labels;
            LogicalInterfaces = logicalInterfaces;
            LoginInfo = loginInfo;
            Luns = luns;
            MachineType = machineType;
            Name = name;
            NetworkTemplate = networkTemplate;
            Networks = networks;
            OsImage = osImage;
            Pod = pod;
            State = state;
            UpdateTime = updateTime;
            Volumes = volumes;
            WorkloadProfile = workloadProfile;
        }
    }
}
