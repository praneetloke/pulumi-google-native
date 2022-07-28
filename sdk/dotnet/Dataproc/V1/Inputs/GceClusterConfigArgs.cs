// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Common config settings for resources of Compute Engine cluster instances, applicable to all instances in the cluster.
    /// </summary>
    public sealed class GceClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Confidential Instance Config for clusters using Confidential VMs (https://cloud.google.com/compute/confidential-vm/docs).
        /// </summary>
        [Input("confidentialInstanceConfig")]
        public Input<Inputs.ConfidentialInstanceConfigArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// Optional. If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This internal_ip_only restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses.
        /// </summary>
        [Input("internalIpOnly")]
        public Input<bool>? InternalIpOnly { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The Compute Engine metadata entries to add to all instances (see Project and instance metadata (https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither network_uri nor subnetwork_uri is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see Using Subnetworks (https://cloud.google.com/compute/docs/subnetworks) for more information).A full URL, partial URI, or short name are valid. Examples: https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default projects/[project_id]/regions/global/default default
        /// </summary>
        [Input("networkUri")]
        public Input<string>? NetworkUri { get; set; }

        /// <summary>
        /// Optional. Node Group Affinity for sole-tenant clusters.
        /// </summary>
        [Input("nodeGroupAffinity")]
        public Input<Inputs.NodeGroupAffinityArgs>? NodeGroupAffinity { get; set; }

        /// <summary>
        /// Optional. The type of IPv6 access for a cluster.
        /// </summary>
        [Input("privateIpv6GoogleAccess")]
        public Input<Pulumi.GoogleNative.Dataproc.V1.GceClusterConfigPrivateIpv6GoogleAccess>? PrivateIpv6GoogleAccess { get; set; }

        /// <summary>
        /// Optional. Reservation Affinity for consuming Zonal reservation.
        /// </summary>
        [Input("reservationAffinity")]
        public Input<Inputs.ReservationAffinityArgs>? ReservationAffinity { get; set; }

        /// <summary>
        /// Optional. The Dataproc service account (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see VM Data Plane identity (https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services.If not specified, the Compute Engine default service account (https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("serviceAccountScopes")]
        private InputList<string>? _serviceAccountScopes;

        /// <summary>
        /// Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: https://www.googleapis.com/auth/cloud.useraccounts.readonly https://www.googleapis.com/auth/devstorage.read_write https://www.googleapis.com/auth/logging.writeIf no scopes are specified, the following defaults are also provided: https://www.googleapis.com/auth/bigquery https://www.googleapis.com/auth/bigtable.admin.table https://www.googleapis.com/auth/bigtable.data https://www.googleapis.com/auth/devstorage.full_control
        /// </summary>
        public InputList<string> ServiceAccountScopes
        {
            get => _serviceAccountScopes ?? (_serviceAccountScopes = new InputList<string>());
            set => _serviceAccountScopes = value;
        }

        /// <summary>
        /// Optional. Shielded Instance Config for clusters using Compute Engine Shielded VMs (https://cloud.google.com/security/shielded-cloud/shielded-vm).
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.ShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri.A full URL, partial URI, or short name are valid. Examples: https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0 projects/[project_id]/regions/us-east1/subnetworks/sub0 sub0
        /// </summary>
        [Input("subnetworkUri")]
        public Input<string>? SubnetworkUri { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The Compute Engine tags to add to all instances (see Tagging instances (https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Optional. The zone where the Compute Engine cluster will be located. On a create request, it is required in the "global" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present.A full URL, partial URI, or short name are valid. Examples: https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone] projects/[project_id]/zones/[zone] us-central1-f
        /// </summary>
        [Input("zoneUri")]
        public Input<string>? ZoneUri { get; set; }

        public GceClusterConfigArgs()
        {
        }
        public static new GceClusterConfigArgs Empty => new GceClusterConfigArgs();
    }
}
