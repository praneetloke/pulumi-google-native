// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Sets the scheduling options for an Instance. NextID: 21
    /// </summary>
    [OutputType]
    public sealed class SchedulingResponse
    {
        /// <summary>
        /// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). You can only set the automatic restart option for standard instances. Preemptible instances cannot be automatically restarted. By default, this is set to true so an instance is automatically restarted if it is terminated by Compute Engine.
        /// </summary>
        public readonly bool AutomaticRestart;
        /// <summary>
        /// Specifies the termination action for the instance.
        /// </summary>
        public readonly string InstanceTerminationAction;
        /// <summary>
        /// An opaque location hint used to place the instance close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        public readonly string LocationHint;
        /// <summary>
        /// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
        /// </summary>
        public readonly int MinNodeCpus;
        /// <summary>
        /// A set of node affinity and anti-affinity configurations. Refer to Configuring node affinity for more information. Overrides reservationAffinity.
        /// </summary>
        public readonly ImmutableArray<Outputs.SchedulingNodeAffinityResponse> NodeAffinities;
        /// <summary>
        /// Defines the maintenance behavior for this instance. For standard instances, the default behavior is MIGRATE. For preemptible instances, the default and only possible behavior is TERMINATE. For more information, see Set VM host maintenance policy.
        /// </summary>
        public readonly string OnHostMaintenance;
        /// <summary>
        /// Defines whether the instance is preemptible. This can only be set during instance creation or while the instance is stopped and therefore, in a `TERMINATED` state. See Instance Life Cycle for more information on the possible instance states.
        /// </summary>
        public readonly bool Preemptible;
        /// <summary>
        /// Specifies the provisioning model of the instance.
        /// </summary>
        public readonly string ProvisioningModel;

        [OutputConstructor]
        private SchedulingResponse(
            bool automaticRestart,

            string instanceTerminationAction,

            string locationHint,

            int minNodeCpus,

            ImmutableArray<Outputs.SchedulingNodeAffinityResponse> nodeAffinities,

            string onHostMaintenance,

            bool preemptible,

            string provisioningModel)
        {
            AutomaticRestart = automaticRestart;
            InstanceTerminationAction = instanceTerminationAction;
            LocationHint = locationHint;
            MinNodeCpus = minNodeCpus;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            Preemptible = preemptible;
            ProvisioningModel = provisioningModel;
        }
    }
}
