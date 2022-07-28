// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Parameters that can be configured on Linux nodes.
    /// </summary>
    public sealed class LinuxNodeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// cgroup_mode specifies the cgroup mode to be used on the node.
        /// </summary>
        [Input("cgroupMode")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.LinuxNodeConfigCgroupMode>? CgroupMode { get; set; }

        [Input("sysctls")]
        private InputMap<string>? _sysctls;

        /// <summary>
        /// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. The following parameters are supported. net.core.busy_poll net.core.busy_read net.core.netdev_max_backlog net.core.rmem_max net.core.wmem_default net.core.wmem_max net.core.optmem_max net.core.somaxconn net.ipv4.tcp_rmem net.ipv4.tcp_wmem net.ipv4.tcp_tw_reuse
        /// </summary>
        public InputMap<string> Sysctls
        {
            get => _sysctls ?? (_sysctls = new InputMap<string>());
            set => _sysctls = value;
        }

        public LinuxNodeConfigArgs()
        {
        }
        public static new LinuxNodeConfigArgs Empty => new LinuxNodeConfigArgs();
    }
}
