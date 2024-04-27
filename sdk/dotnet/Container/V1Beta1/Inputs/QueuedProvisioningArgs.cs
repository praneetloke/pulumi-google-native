// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// QueuedProvisioning defines the queued provisioning used by the node pool.
    /// </summary>
    public sealed class QueuedProvisioningArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Denotes that this nodepool is QRM specific, meaning nodes can be only obtained through queuing via the Cluster Autoscaler ProvisioningRequest API.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public QueuedProvisioningArgs()
        {
        }
        public static new QueuedProvisioningArgs Empty => new QueuedProvisioningArgs();
    }
}