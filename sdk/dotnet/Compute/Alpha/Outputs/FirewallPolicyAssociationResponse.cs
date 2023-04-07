// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class FirewallPolicyAssociationResponse
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        public readonly string AttachmentTarget;
        /// <summary>
        /// Deprecated, please use short name instead. The display name of the firewall policy of the association.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The firewall policy ID of the association.
        /// </summary>
        public readonly string FirewallPolicyId;
        /// <summary>
        /// The name for an association.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An integer indicating the priority of an association. The priority must be a positive value between 1 and 2147483647. Firewall Policies are evaluated from highest to lowest priority where 1 is the highest priority and 2147483647 is the lowest priority. The default value is `1000`. If two associations have the same priority then lexicographical order on association names is applied.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The short name of the firewall policy of the association.
        /// </summary>
        public readonly string ShortName;

        [OutputConstructor]
        private FirewallPolicyAssociationResponse(
            string attachmentTarget,

            string displayName,

            string firewallPolicyId,

            string name,

            int priority,

            string shortName)
        {
            AttachmentTarget = attachmentTarget;
            DisplayName = displayName;
            FirewallPolicyId = firewallPolicyId;
            Name = name;
            Priority = priority;
            ShortName = shortName;
        }
    }
}
