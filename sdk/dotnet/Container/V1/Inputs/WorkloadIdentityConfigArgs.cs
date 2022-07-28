// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
    /// </summary>
    public sealed class WorkloadIdentityConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The workload pool to attach all Kubernetes service accounts to.
        /// </summary>
        [Input("workloadPool")]
        public Input<string>? WorkloadPool { get; set; }

        public WorkloadIdentityConfigArgs()
        {
        }
        public static new WorkloadIdentityConfigArgs Empty => new WorkloadIdentityConfigArgs();
    }
}
