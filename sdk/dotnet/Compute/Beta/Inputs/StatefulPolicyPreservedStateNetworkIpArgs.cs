// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class StatefulPolicyPreservedStateNetworkIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// These stateful IPs will never be released during autohealing, update or VM instance recreate operations. This flag is used to configure if the IP reservation should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted.
        /// </summary>
        [Input("autoDelete")]
        public Input<Pulumi.GoogleNative.Compute.Beta.StatefulPolicyPreservedStateNetworkIpAutoDelete>? AutoDelete { get; set; }

        public StatefulPolicyPreservedStateNetworkIpArgs()
        {
        }
        public static new StatefulPolicyPreservedStateNetworkIpArgs Empty => new StatefulPolicyPreservedStateNetworkIpArgs();
    }
}
