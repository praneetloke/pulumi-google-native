// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Notebooks.V1.Outputs
{

    [OutputType]
    public sealed class VirtualMachineResponse
    {
        /// <summary>
        /// The unique identifier of the Managed Compute Engine instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The user-friendly name of the Managed Compute Engine instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// Virtual Machine configuration settings.
        /// </summary>
        public readonly Outputs.VirtualMachineConfigResponse VirtualMachineConfig;

        [OutputConstructor]
        private VirtualMachineResponse(
            string instanceId,

            string instanceName,

            Outputs.VirtualMachineConfigResponse virtualMachineConfig)
        {
            InstanceId = instanceId;
            InstanceName = instanceName;
            VirtualMachineConfig = virtualMachineConfig;
        }
    }
}
