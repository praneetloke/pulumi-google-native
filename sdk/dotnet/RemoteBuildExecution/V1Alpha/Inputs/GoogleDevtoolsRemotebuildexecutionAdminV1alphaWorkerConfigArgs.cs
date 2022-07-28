// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RemoteBuildExecution.V1Alpha.Inputs
{

    /// <summary>
    /// Defines the configuration to be used for creating workers in the worker pool.
    /// </summary>
    public sealed class GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accelerator card attached to each VM.
        /// </summary>
        [Input("accelerator")]
        public Input<Inputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigArgs>? Accelerator { get; set; }

        /// <summary>
        /// Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/
        /// </summary>
        [Input("diskSizeGb", required: true)]
        public Input<string> DiskSizeGb { get; set; } = null!;

        /// <summary>
        /// Disk Type to use for the worker. See [Storage options](https://cloud.google.com/compute/docs/disks/#introduction). Currently only `pd-standard` and `pd-ssd` are supported.
        /// </summary>
        [Input("diskType", required: true)]
        public Input<string> DiskType { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels associated with the workers. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International letters are permitted. Label keys must start with a letter. Label values are optional. There can not be more than 64 labels per resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Machine type of the worker, such as `e2-standard-2`. See https://cloud.google.com/compute/docs/machine-types for a list of supported machine types. Note that `f1-micro` and `g1-small` are not yet supported.
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        /// <summary>
        /// The maximum number of actions a worker can execute concurrently.
        /// </summary>
        [Input("maxConcurrentActions")]
        public Input<string>? MaxConcurrentActions { get; set; }

        /// <summary>
        /// Minimum CPU platform to use when creating the worker. See [CPU Platforms](https://cloud.google.com/compute/docs/cpu-platforms).
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// Determines the type of network access granted to workers. Possible values: - "public": Workers can connect to the public internet. - "private": Workers can only connect to Google APIs and services. - "restricted-private": Workers can only connect to Google APIs that are reachable through `restricted.googleapis.com` (`199.36.153.4/30`).
        /// </summary>
        [Input("networkAccess")]
        public Input<string>? NetworkAccess { get; set; }

        /// <summary>
        /// Determines whether the worker is reserved (equivalent to a Compute Engine on-demand VM and therefore won't be preempted). See [Preemptible VMs](https://cloud.google.com/preemptible-vms/) for more details.
        /// </summary>
        [Input("reserved")]
        public Input<bool>? Reserved { get; set; }

        /// <summary>
        /// The node type name to be used for sole-tenant nodes.
        /// </summary>
        [Input("soleTenantNodeType")]
        public Input<string>? SoleTenantNodeType { get; set; }

        /// <summary>
        /// The name of the image used by each VM.
        /// </summary>
        [Input("vmImage")]
        public Input<string>? VmImage { get; set; }

        public GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs()
        {
        }
        public static new GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs Empty => new GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigArgs();
    }
}
