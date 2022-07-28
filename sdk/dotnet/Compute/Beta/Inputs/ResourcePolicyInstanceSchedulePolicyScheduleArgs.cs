// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Schedule for an instance operation.
    /// </summary>
    public sealed class ResourcePolicyInstanceSchedulePolicyScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the frequency for the operation, using the unix-cron format.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        public ResourcePolicyInstanceSchedulePolicyScheduleArgs()
        {
        }
        public static new ResourcePolicyInstanceSchedulePolicyScheduleArgs Empty => new ResourcePolicyInstanceSchedulePolicyScheduleArgs();
    }
}
