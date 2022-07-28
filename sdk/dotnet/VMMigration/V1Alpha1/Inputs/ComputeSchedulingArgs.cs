// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Inputs
{

    /// <summary>
    /// Scheduling information for VM on maintenance/restart behaviour and node allocation in sole tenant nodes.
    /// </summary>
    public sealed class ComputeSchedulingArgs : global::Pulumi.ResourceArgs
    {
        [Input("automaticRestart")]
        public Input<bool>? AutomaticRestart { get; set; }

        /// <summary>
        /// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node. Ignored if no node_affinites are configured.
        /// </summary>
        [Input("minNodeCpus")]
        public Input<int>? MinNodeCpus { get; set; }

        [Input("nodeAffinities")]
        private InputList<Inputs.SchedulingNodeAffinityArgs>? _nodeAffinities;

        /// <summary>
        /// A set of node affinity and anti-affinity configurations for sole tenant nodes.
        /// </summary>
        public InputList<Inputs.SchedulingNodeAffinityArgs> NodeAffinities
        {
            get => _nodeAffinities ?? (_nodeAffinities = new InputList<Inputs.SchedulingNodeAffinityArgs>());
            set => _nodeAffinities = value;
        }

        /// <summary>
        /// How the instance should behave when the host machine undergoes maintenance that may temporarily impact instance performance.
        /// </summary>
        [Input("onHostMaintenance")]
        public Input<Pulumi.GoogleNative.VMMigration.V1Alpha1.ComputeSchedulingOnHostMaintenance>? OnHostMaintenance { get; set; }

        /// <summary>
        /// Whether the Instance should be automatically restarted whenever it is terminated by Compute Engine (not terminated by user). This configuration is identical to `automaticRestart` field in Compute Engine create instance under scheduling. It was changed to an enum (instead of a boolean) to match the default value in Compute Engine which is automatic restart.
        /// </summary>
        [Input("restartType")]
        public Input<Pulumi.GoogleNative.VMMigration.V1Alpha1.ComputeSchedulingRestartType>? RestartType { get; set; }

        public ComputeSchedulingArgs()
        {
        }
        public static new ComputeSchedulingArgs Empty => new ComputeSchedulingArgs();
    }
}
