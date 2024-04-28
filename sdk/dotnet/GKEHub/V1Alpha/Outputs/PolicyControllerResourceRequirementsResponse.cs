// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// ResourceRequirements describes the compute resource requirements.
    /// </summary>
    [OutputType]
    public sealed class PolicyControllerResourceRequirementsResponse
    {
        /// <summary>
        /// Limits describes the maximum amount of compute resources allowed for use by the running container.
        /// </summary>
        public readonly Outputs.PolicyControllerResourceListResponse Limits;
        /// <summary>
        /// Requests describes the amount of compute resources reserved for the container by the kube-scheduler.
        /// </summary>
        public readonly Outputs.PolicyControllerResourceListResponse Requests;

        [OutputConstructor]
        private PolicyControllerResourceRequirementsResponse(
            Outputs.PolicyControllerResourceListResponse limits,

            Outputs.PolicyControllerResourceListResponse requests)
        {
            Limits = limits;
            Requests = requests;
        }
    }
}
