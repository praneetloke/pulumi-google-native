// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration information for an environment.
    /// </summary>
    public sealed class EnvironmentConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The configuration settings for Cloud SQL instance used internally by Apache Airflow software.
        /// </summary>
        [Input("databaseConfig")]
        public Input<Inputs.DatabaseConfigArgs>? DatabaseConfig { get; set; }

        /// <summary>
        /// Optional. The encryption options for the Cloud Composer environment and its dependencies. Cannot be updated.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Optional. The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
        /// </summary>
        [Input("environmentSize")]
        public Input<Pulumi.GoogleNative.Composer.V1Beta1.EnvironmentConfigEnvironmentSize>? EnvironmentSize { get; set; }

        /// <summary>
        /// Optional. The maintenance window is the period when Cloud Composer components may undergo maintenance. It is defined so that maintenance is not executed during peak hours or critical time periods. The system will not be under maintenance for every occurrence of this window, but when maintenance is planned, it will be scheduled during the window. The maintenance window period must encompass at least 12 hours per week. This may be split into multiple chunks, each with a size of at least 4 hours. If this value is omitted, Cloud Composer components may be subject to maintenance at any time.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Optional. The configuration options for GKE cluster master authorized networks. By default master authorized networks feature is: - in case of private environment: enabled with no external networks allowlisted. - in case of public environment: disabled.
        /// </summary>
        [Input("masterAuthorizedNetworksConfig")]
        public Input<Inputs.MasterAuthorizedNetworksConfigArgs>? MasterAuthorizedNetworksConfig { get; set; }

        /// <summary>
        /// The configuration used for the Kubernetes Engine cluster.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodeConfigArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The configuration used for the Private IP Cloud Composer environment.
        /// </summary>
        [Input("privateEnvironmentConfig")]
        public Input<Inputs.PrivateEnvironmentConfigArgs>? PrivateEnvironmentConfig { get; set; }

        /// <summary>
        /// The configuration settings for software inside the environment.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.SoftwareConfigArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// Optional. The configuration settings for the Airflow web server App Engine instance. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
        /// </summary>
        [Input("webServerConfig")]
        public Input<Inputs.WebServerConfigArgs>? WebServerConfig { get; set; }

        /// <summary>
        /// Optional. The network-level access control policy for the Airflow web server. If unspecified, no network-level access restrictions will be applied.
        /// </summary>
        [Input("webServerNetworkAccessControl")]
        public Input<Inputs.WebServerNetworkAccessControlArgs>? WebServerNetworkAccessControl { get; set; }

        /// <summary>
        /// Optional. The workloads configuration settings for the GKE cluster associated with the Cloud Composer environment. The GKE cluster runs Airflow scheduler, web server and workers workloads. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
        /// </summary>
        [Input("workloadsConfig")]
        public Input<Inputs.WorkloadsConfigArgs>? WorkloadsConfig { get; set; }

        public EnvironmentConfigArgs()
        {
        }
        public static new EnvironmentConfigArgs Empty => new EnvironmentConfigArgs();
    }
}
