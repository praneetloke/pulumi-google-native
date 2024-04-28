// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// Configuration for Policy Controller
    /// </summary>
    [OutputType]
    public sealed class PolicyControllerHubConfigResponse
    {
        /// <summary>
        /// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
        /// </summary>
        public readonly string AuditIntervalSeconds;
        /// <summary>
        /// The maximum number of audit violations to be stored in a constraint. If not set, the internal default (currently 20) will be used.
        /// </summary>
        public readonly string ConstraintViolationLimit;
        /// <summary>
        /// Map of deployment configs to deployments ("admission", "audit", "mutation').
        /// </summary>
        public readonly Outputs.PolicyControllerPolicyControllerDeploymentConfigResponse DeploymentConfigs;
        /// <summary>
        /// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
        /// </summary>
        public readonly ImmutableArray<string> ExemptableNamespaces;
        /// <summary>
        /// The install_spec represents the intended state specified by the latest request that mutated install_spec in the feature spec, not the lifecycle state of the feature observed by the Hub feature controller that is reported in the feature state.
        /// </summary>
        public readonly string InstallSpec;
        /// <summary>
        /// Logs all denies and dry run failures.
        /// </summary>
        public readonly bool LogDeniesEnabled;
        /// <summary>
        /// Monitoring specifies the configuration of monitoring.
        /// </summary>
        public readonly Outputs.PolicyControllerMonitoringConfigResponse Monitoring;
        /// <summary>
        /// Enables the ability to mutate resources using Policy Controller.
        /// </summary>
        public readonly bool MutationEnabled;
        /// <summary>
        /// Specifies the desired policy content on the cluster
        /// </summary>
        public readonly Outputs.PolicyControllerPolicyContentSpecResponse PolicyContent;
        /// <summary>
        /// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
        /// </summary>
        public readonly bool ReferentialRulesEnabled;

        [OutputConstructor]
        private PolicyControllerHubConfigResponse(
            string auditIntervalSeconds,

            string constraintViolationLimit,

            Outputs.PolicyControllerPolicyControllerDeploymentConfigResponse deploymentConfigs,

            ImmutableArray<string> exemptableNamespaces,

            string installSpec,

            bool logDeniesEnabled,

            Outputs.PolicyControllerMonitoringConfigResponse monitoring,

            bool mutationEnabled,

            Outputs.PolicyControllerPolicyContentSpecResponse policyContent,

            bool referentialRulesEnabled)
        {
            AuditIntervalSeconds = auditIntervalSeconds;
            ConstraintViolationLimit = constraintViolationLimit;
            DeploymentConfigs = deploymentConfigs;
            ExemptableNamespaces = exemptableNamespaces;
            InstallSpec = installSpec;
            LogDeniesEnabled = logDeniesEnabled;
            Monitoring = monitoring;
            MutationEnabled = mutationEnabled;
            PolicyContent = policyContent;
            ReferentialRulesEnabled = referentialRulesEnabled;
        }
    }
}
