// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Inputs
{

    /// <summary>
    /// Either an InstancePolicy or an instance template.
    /// </summary>
    public sealed class InstancePolicyOrTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set this field true if users want Batch to help fetch drivers from a third party location and install them for GPUs specified in policy.accelerators or instance_template on their behalf. Default is false.
        /// </summary>
        [Input("installGpuDrivers")]
        public Input<bool>? InstallGpuDrivers { get; set; }

        /// <summary>
        /// Name of an instance template used to create VMs. Named the field as 'instance_template' instead of 'template' to avoid c++ keyword conflict.
        /// </summary>
        [Input("instanceTemplate")]
        public Input<string>? InstanceTemplate { get; set; }

        /// <summary>
        /// InstancePolicy.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.InstancePolicyArgs>? Policy { get; set; }

        public InstancePolicyOrTemplateArgs()
        {
        }
        public static new InstancePolicyOrTemplateArgs Empty => new InstancePolicyOrTemplateArgs();
    }
}
