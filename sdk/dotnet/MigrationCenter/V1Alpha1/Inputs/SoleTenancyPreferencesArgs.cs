// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1Alpha1.Inputs
{

    /// <summary>
    /// Preferences concerning Sole Tenancy nodes and VMs.
    /// </summary>
    public sealed class SoleTenancyPreferencesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with.
        /// </summary>
        [Input("commitmentPlan")]
        public Input<Pulumi.GoogleNative.MigrationCenter.V1Alpha1.SoleTenancyPreferencesCommitmentPlan>? CommitmentPlan { get; set; }

        /// <summary>
        /// CPU overcommit ratio. Acceptable values are between 1.0 and 2.0 inclusive.
        /// </summary>
        [Input("cpuOvercommitRatio")]
        public Input<double>? CpuOvercommitRatio { get; set; }

        /// <summary>
        /// Sole Tenancy nodes maintenance policy.
        /// </summary>
        [Input("hostMaintenancePolicy")]
        public Input<Pulumi.GoogleNative.MigrationCenter.V1Alpha1.SoleTenancyPreferencesHostMaintenancePolicy>? HostMaintenancePolicy { get; set; }

        [Input("nodeTypes")]
        private InputList<Inputs.SoleTenantNodeTypeArgs>? _nodeTypes;

        /// <summary>
        /// A list of sole tenant node types. An empty list means that all possible node types will be considered.
        /// </summary>
        public InputList<Inputs.SoleTenantNodeTypeArgs> NodeTypes
        {
            get => _nodeTypes ?? (_nodeTypes = new InputList<Inputs.SoleTenantNodeTypeArgs>());
            set => _nodeTypes = value;
        }

        public SoleTenancyPreferencesArgs()
        {
        }
        public static new SoleTenancyPreferencesArgs Empty => new SoleTenancyPreferencesArgs();
    }
}