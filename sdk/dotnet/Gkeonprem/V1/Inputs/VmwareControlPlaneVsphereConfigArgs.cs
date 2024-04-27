// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Inputs
{

    /// <summary>
    /// Specifies control plane node config.
    /// </summary>
    public sealed class VmwareControlPlaneVsphereConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Vsphere datastore used by the control plane Node.
        /// </summary>
        [Input("datastore")]
        public Input<string>? Datastore { get; set; }

        /// <summary>
        /// The Vsphere storage policy used by the control plane Node.
        /// </summary>
        [Input("storagePolicyName")]
        public Input<string>? StoragePolicyName { get; set; }

        public VmwareControlPlaneVsphereConfigArgs()
        {
        }
        public static new VmwareControlPlaneVsphereConfigArgs Empty => new VmwareControlPlaneVsphereConfigArgs();
    }
}