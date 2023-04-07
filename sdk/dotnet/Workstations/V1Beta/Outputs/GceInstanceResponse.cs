// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Outputs
{

    /// <summary>
    /// A runtime using a Compute Engine instance.
    /// </summary>
    [OutputType]
    public sealed class GceInstanceResponse
    {
        /// <summary>
        /// Size of the boot disk in GB. Defaults to 50.
        /// </summary>
        public readonly int BootDiskSizeGb;
        /// <summary>
        /// A set of Compute Engine Confidential VM instance options.
        /// </summary>
        public readonly Outputs.GceConfidentialInstanceConfigResponse ConfidentialInstanceConfig;
        /// <summary>
        /// Whether instances have no public IP address.
        /// </summary>
        public readonly bool DisablePublicIpAddresses;
        /// <summary>
        /// The name of a Compute Engine machine type.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Number of instances to pool for faster workstation starup.
        /// </summary>
        public readonly int PoolSize;
        /// <summary>
        /// Email address of the service account that will be used on VM instances used to support this config. If not set, VMs will run with a Google-managed service account. This service account must have permission to pull the specified container image, otherwise the image must be publicly accessible.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// A set of Compute Engine Shielded instance options.
        /// </summary>
        public readonly Outputs.GceShieldedInstanceConfigResponse ShieldedInstanceConfig;
        /// <summary>
        /// Network tags to add to the Compute Engine machines backing the Workstations.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GceInstanceResponse(
            int bootDiskSizeGb,

            Outputs.GceConfidentialInstanceConfigResponse confidentialInstanceConfig,

            bool disablePublicIpAddresses,

            string machineType,

            int poolSize,

            string serviceAccount,

            Outputs.GceShieldedInstanceConfigResponse shieldedInstanceConfig,

            ImmutableArray<string> tags)
        {
            BootDiskSizeGb = bootDiskSizeGb;
            ConfidentialInstanceConfig = confidentialInstanceConfig;
            DisablePublicIpAddresses = disablePublicIpAddresses;
            MachineType = machineType;
            PoolSize = poolSize;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            Tags = tags;
        }
    }
}
