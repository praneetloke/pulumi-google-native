// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta.Inputs
{

    /// <summary>
    /// Label object for Deployments
    /// </summary>
    public sealed class DeploymentLabelEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the label
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Value of the label
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DeploymentLabelEntryArgs()
        {
        }
        public static new DeploymentLabelEntryArgs Empty => new DeploymentLabelEntryArgs();
    }
}
