// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class InstanceGroupManagerUpdatePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance redistribution policy for regional managed instance groups. Valid values are: - PROACTIVE (default): The group attempts to maintain an even distribution of VM instances across zones in the region. - NONE: For non-autoscaled groups, proactive redistribution is disabled. 
        /// </summary>
        [Input("instanceRedistributionType")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionType>? InstanceRedistributionType { get; set; }

        /// <summary>
        /// The maximum number of instances that can be created above the specified targetSize during the update process. This value can be either a fixed number or, if the group has 10 or more instances, a percentage. If you set a percentage, the number of instances is rounded if necessary. The default value for maxSurge is a fixed value equal to the number of zones in which the managed instance group operates. At least one of either maxSurge or maxUnavailable must be greater than 0. Learn more about maxSurge.
        /// </summary>
        [Input("maxSurge")]
        public Input<Inputs.FixedOrPercentArgs>? MaxSurge { get; set; }

        /// <summary>
        /// The maximum number of instances that can be unavailable during the update process. An instance is considered available if all of the following conditions are satisfied: - The instance's status is RUNNING. - If there is a health check on the instance group, the instance's health check status must be HEALTHY at least once. If there is no health check on the group, then the instance only needs to have a status of RUNNING to be considered available. This value can be either a fixed number or, if the group has 10 or more instances, a percentage. If you set a percentage, the number of instances is rounded if necessary. The default value for maxUnavailable is a fixed value equal to the number of zones in which the managed instance group operates. At least one of either maxSurge or maxUnavailable must be greater than 0. Learn more about maxUnavailable.
        /// </summary>
        [Input("maxUnavailable")]
        public Input<Inputs.FixedOrPercentArgs>? MaxUnavailable { get; set; }

        /// <summary>
        /// Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600].
        /// </summary>
        [Input("minReadySec")]
        public Input<int>? MinReadySec { get; set; }

        /// <summary>
        /// Minimal action to be taken on an instance. Use this option to minimize disruption as much as possible or to apply a more disruptive action than is necessary. - To limit disruption as much as possible, set the minimal action to REFRESH. If your update requires a more disruptive action, Compute Engine performs the necessary action to execute the update. - To apply a more disruptive action than is strictly necessary, set the minimal action to RESTART or REPLACE. For example, Compute Engine does not need to restart a VM to change its metadata. But if your application reads instance metadata only when a VM is restarted, you can set the minimal action to RESTART in order to pick up metadata changes. 
        /// </summary>
        [Input("minimalAction")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerUpdatePolicyMinimalAction>? MinimalAction { get; set; }

        /// <summary>
        /// Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
        /// </summary>
        [Input("mostDisruptiveAllowedAction")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedAction>? MostDisruptiveAllowedAction { get; set; }

        /// <summary>
        /// What action should be used to replace instances. See minimal_action.REPLACE
        /// </summary>
        [Input("replacementMethod")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerUpdatePolicyReplacementMethod>? ReplacementMethod { get; set; }

        /// <summary>
        /// The type of update process. You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerUpdatePolicyType>? Type { get; set; }

        public InstanceGroupManagerUpdatePolicyArgs()
        {
        }
        public static new InstanceGroupManagerUpdatePolicyArgs Empty => new InstanceGroupManagerUpdatePolicyArgs();
    }
}
