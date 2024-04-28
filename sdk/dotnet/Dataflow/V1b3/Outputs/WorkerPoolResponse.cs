// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Describes one particular pool of Cloud Dataflow workers to be instantiated by the Cloud Dataflow service in order to perform the computations required by a job. Note that a workflow job may use multiple pools, in order to match the various computational requirements of the various stages of the job.
    /// </summary>
    [OutputType]
    public sealed class WorkerPoolResponse
    {
        /// <summary>
        /// Settings for autoscaling of this WorkerPool.
        /// </summary>
        public readonly Outputs.AutoscalingSettingsResponse AutoscalingSettings;
        /// <summary>
        /// Data disks that are used by a VM in this workflow.
        /// </summary>
        public readonly ImmutableArray<Outputs.DiskResponse> DataDisks;
        /// <summary>
        /// The default package set to install. This allows the service to select a default set of packages which are useful to worker harnesses written in a particular language.
        /// </summary>
        public readonly string DefaultPackageSet;
        /// <summary>
        /// Size of root disk for VMs, in GB. If zero or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly int DiskSizeGb;
        /// <summary>
        /// Fully qualified source image for disks.
        /// </summary>
        public readonly string DiskSourceImage;
        /// <summary>
        /// Type of root disk for VMs. If empty or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// Configuration for VM IPs.
        /// </summary>
        public readonly string IpConfiguration;
        /// <summary>
        /// The kind of the worker pool; currently only `harness` and `shuffle` are supported.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Machine type (e.g. "n1-standard-1"). If empty or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Metadata to set on the Google Compute Engine VMs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Network to which VMs will be assigned. If empty or unspecified, the service will use the network "default".
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The number of threads per worker harness. If empty or unspecified, the service will choose a number of threads (according to the number of cores on the selected machine type for batch, or 1 by convention for streaming).
        /// </summary>
        public readonly int NumThreadsPerWorker;
        /// <summary>
        /// Number of Google Compute Engine workers in this pool needed to execute the job. If zero or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly int NumWorkers;
        /// <summary>
        /// The action to take on host maintenance, as defined by the Google Compute Engine API.
        /// </summary>
        public readonly string OnHostMaintenance;
        /// <summary>
        /// Packages to be installed on workers.
        /// </summary>
        public readonly ImmutableArray<Outputs.PackageResponse> Packages;
        /// <summary>
        /// Extra arguments for this worker pool.
        /// </summary>
        public readonly ImmutableDictionary<string, object> PoolArgs;
        /// <summary>
        /// Set of SDK harness containers needed to execute this pipeline. This will only be set in the Fn API path. For non-cross-language pipelines this should have only one entry. Cross-language pipelines will have two or more entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.SdkHarnessContainerImageResponse> SdkHarnessContainerImages;
        /// <summary>
        /// Subnetwork to which VMs will be assigned, if desired. Expected to be of the form "regions/REGION/subnetworks/SUBNETWORK".
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// Settings passed through to Google Compute Engine workers when using the standard Dataflow task runner. Users should ignore this field.
        /// </summary>
        public readonly Outputs.TaskRunnerSettingsResponse TaskrunnerSettings;
        /// <summary>
        /// Sets the policy for determining when to turndown worker pool. Allowed values are: `TEARDOWN_ALWAYS`, `TEARDOWN_ON_SUCCESS`, and `TEARDOWN_NEVER`. `TEARDOWN_ALWAYS` means workers are always torn down regardless of whether the job succeeds. `TEARDOWN_ON_SUCCESS` means workers are torn down if the job succeeds. `TEARDOWN_NEVER` means the workers are never torn down. If the workers are not torn down by the service, they will continue to run and use Google Compute Engine VM resources in the user's project until they are explicitly terminated by the user. Because of this, Google recommends using the `TEARDOWN_ALWAYS` policy except for small, manually supervised test jobs. If unknown or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly string TeardownPolicy;
        /// <summary>
        /// Docker container image that executes the Cloud Dataflow worker harness, residing in Google Container Registry. Deprecated for the Fn API path. Use sdk_harness_container_images instead.
        /// </summary>
        public readonly string WorkerHarnessContainerImage;
        /// <summary>
        /// Zone to run the worker pools in. If empty or unspecified, the service will attempt to choose a reasonable default.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private WorkerPoolResponse(
            Outputs.AutoscalingSettingsResponse autoscalingSettings,

            ImmutableArray<Outputs.DiskResponse> dataDisks,

            string defaultPackageSet,

            int diskSizeGb,

            string diskSourceImage,

            string diskType,

            string ipConfiguration,

            string kind,

            string machineType,

            ImmutableDictionary<string, string> metadata,

            string network,

            int numThreadsPerWorker,

            int numWorkers,

            string onHostMaintenance,

            ImmutableArray<Outputs.PackageResponse> packages,

            ImmutableDictionary<string, object> poolArgs,

            ImmutableArray<Outputs.SdkHarnessContainerImageResponse> sdkHarnessContainerImages,

            string subnetwork,

            Outputs.TaskRunnerSettingsResponse taskrunnerSettings,

            string teardownPolicy,

            string workerHarnessContainerImage,

            string zone)
        {
            AutoscalingSettings = autoscalingSettings;
            DataDisks = dataDisks;
            DefaultPackageSet = defaultPackageSet;
            DiskSizeGb = diskSizeGb;
            DiskSourceImage = diskSourceImage;
            DiskType = diskType;
            IpConfiguration = ipConfiguration;
            Kind = kind;
            MachineType = machineType;
            Metadata = metadata;
            Network = network;
            NumThreadsPerWorker = numThreadsPerWorker;
            NumWorkers = numWorkers;
            OnHostMaintenance = onHostMaintenance;
            Packages = packages;
            PoolArgs = poolArgs;
            SdkHarnessContainerImages = sdkHarnessContainerImages;
            Subnetwork = subnetwork;
            TaskrunnerSettings = taskrunnerSettings;
            TeardownPolicy = teardownPolicy;
            WorkerHarnessContainerImage = workerHarnessContainerImage;
            Zone = zone;
        }
    }
}
