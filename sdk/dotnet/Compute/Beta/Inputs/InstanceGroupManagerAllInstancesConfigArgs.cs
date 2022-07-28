// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class InstanceGroupManagerAllInstancesConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Properties to set on all instances in the group. You can add or modify properties using the instanceGroupManagers.patch or regionInstanceGroupManagers.patch. After setting allInstancesConfig on the group, you must update the group's instances to apply the configuration. To apply the configuration, set the group's updatePolicy.type field to use proactive updates or use the applyUpdatesToInstances method.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.InstancePropertiesPatchArgs>? Properties { get; set; }

        public InstanceGroupManagerAllInstancesConfigArgs()
        {
        }
        public static new InstanceGroupManagerAllInstancesConfigArgs Empty => new InstanceGroupManagerAllInstancesConfigArgs();
    }
}
