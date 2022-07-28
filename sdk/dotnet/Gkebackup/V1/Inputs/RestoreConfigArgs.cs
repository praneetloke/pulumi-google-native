// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1.Inputs
{

    /// <summary>
    /// Configuration of a restore. Next id: 9
    /// </summary>
    public sealed class RestoreConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Restore all namespaced resources in the Backup if set to "True". Specifying this field to "False" is an error.
        /// </summary>
        [Input("allNamespaces")]
        public Input<bool>? AllNamespaces { get; set; }

        /// <summary>
        /// Defines the behavior for handling the situation where cluster-scoped resources being restored already exist in the target cluster. This MUST be set to a value other than CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED if cluster_resource_restore_scope is not empty.
        /// </summary>
        [Input("clusterResourceConflictPolicy")]
        public Input<Pulumi.GoogleNative.Gkebackup.V1.RestoreConfigClusterResourceConflictPolicy>? ClusterResourceConflictPolicy { get; set; }

        /// <summary>
        /// Identifies the cluster-scoped resources to restore from the Backup. Not specifying it means NO cluster resource will be restored.
        /// </summary>
        [Input("clusterResourceRestoreScope")]
        public Input<Inputs.ClusterResourceRestoreScopeArgs>? ClusterResourceRestoreScope { get; set; }

        /// <summary>
        /// Defines the behavior for handling the situation where sets of namespaced resources being restored already exist in the target cluster. This MUST be set to a value other than NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED.
        /// </summary>
        [Input("namespacedResourceRestoreMode")]
        public Input<Pulumi.GoogleNative.Gkebackup.V1.RestoreConfigNamespacedResourceRestoreMode>? NamespacedResourceRestoreMode { get; set; }

        /// <summary>
        /// A list of selected ProtectedApplications to restore. The listed ProtectedApplications and all the resources to which they refer will be restored.
        /// </summary>
        [Input("selectedApplications")]
        public Input<Inputs.NamespacedNamesArgs>? SelectedApplications { get; set; }

        /// <summary>
        /// A list of selected Namespaces to restore from the Backup. The listed Namespaces and all resources contained in them will be restored.
        /// </summary>
        [Input("selectedNamespaces")]
        public Input<Inputs.NamespacesArgs>? SelectedNamespaces { get; set; }

        [Input("substitutionRules")]
        private InputList<Inputs.SubstitutionRuleArgs>? _substitutionRules;

        /// <summary>
        /// A list of transformation rules to be applied against Kubernetes resources as they are selected for restoration from a Backup. Rules are executed in order defined - this order matters, as changes made by a rule may impact the filtering logic of subsequent rules. An empty list means no substitution will occur.
        /// </summary>
        public InputList<Inputs.SubstitutionRuleArgs> SubstitutionRules
        {
            get => _substitutionRules ?? (_substitutionRules = new InputList<Inputs.SubstitutionRuleArgs>());
            set => _substitutionRules = value;
        }

        /// <summary>
        /// Specifies the mechanism to be used to restore volume data. Default: VOLUME_DATA_RESTORE_POLICY_UNSPECIFIED (will be treated as NO_VOLUME_DATA_RESTORATION).
        /// </summary>
        [Input("volumeDataRestorePolicy")]
        public Input<Pulumi.GoogleNative.Gkebackup.V1.RestoreConfigVolumeDataRestorePolicy>? VolumeDataRestorePolicy { get; set; }

        public RestoreConfigArgs()
        {
        }
        public static new RestoreConfigArgs Empty => new RestoreConfigArgs();
    }
}
