// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// A concurrency control configuration. Defines a group config that, when attached to an instance, recognizes that instance as part of a group of instances where only up the concurrency_limit of instances in that group can undergo simultaneous maintenance. For more information: go/concurrency-control-design-doc
    /// </summary>
    public sealed class ResourcePolicyVmMaintenancePolicyConcurrencyControlArgs : global::Pulumi.ResourceArgs
    {
        [Input("concurrencyLimit")]
        public Input<int>? ConcurrencyLimit { get; set; }

        public ResourcePolicyVmMaintenancePolicyConcurrencyControlArgs()
        {
        }
        public static new ResourcePolicyVmMaintenancePolicyConcurrencyControlArgs Empty => new ResourcePolicyVmMaintenancePolicyConcurrencyControlArgs();
    }
}
