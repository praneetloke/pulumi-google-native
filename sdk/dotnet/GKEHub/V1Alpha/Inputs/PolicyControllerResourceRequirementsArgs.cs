// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Inputs
{

    /// <summary>
    /// ResourceRequirements describes the compute resource requirements.
    /// </summary>
    public sealed class PolicyControllerResourceRequirementsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limits describes the maximum amount of compute resources allowed for use by the running container.
        /// </summary>
        [Input("limits")]
        public Input<Inputs.PolicyControllerResourceListArgs>? Limits { get; set; }

        /// <summary>
        /// Requests describes the amount of compute resources reserved for the container by the kube-scheduler.
        /// </summary>
        [Input("requests")]
        public Input<Inputs.PolicyControllerResourceListArgs>? Requests { get; set; }

        public PolicyControllerResourceRequirementsArgs()
        {
        }
        public static new PolicyControllerResourceRequirementsArgs Empty => new PolicyControllerResourceRequirementsArgs();
    }
}
