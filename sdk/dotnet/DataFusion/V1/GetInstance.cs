// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataFusion.V1
{
    public static class GetInstance
    {
        /// <summary>
        /// Gets details of a single Data Fusion instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:datafusion/v1:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Data Fusion instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:datafusion/v1:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// List of accelerators enabled for this CDF instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorResponse> Accelerators;
        /// <summary>
        /// Endpoint on which the REST APIs is accessible.
        /// </summary>
        public readonly string ApiEndpoint;
        /// <summary>
        /// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        /// </summary>
        public readonly ImmutableArray<Outputs.VersionResponse> AvailableVersion;
        /// <summary>
        /// The time the instance was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
        /// </summary>
        public readonly Outputs.CryptoKeyConfigResponse CryptoKeyConfig;
        /// <summary>
        /// Optional. Reserved for future use.
        /// </summary>
        public readonly bool DataplexDataLineageIntegrationEnabled;
        /// <summary>
        /// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        /// </summary>
        public readonly string DataprocServiceAccount;
        /// <summary>
        /// A description of this instance.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If the instance state is DISABLED, the reason for disabling the instance.
        /// </summary>
        public readonly ImmutableArray<string> DisabledReason;
        /// <summary>
        /// Display name for an instance.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Option to enable granular role-based access control.
        /// </summary>
        public readonly bool EnableRbac;
        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        public readonly bool EnableStackdriverLogging;
        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        public readonly bool EnableStackdriverMonitoring;
        /// <summary>
        /// Option to enable granular zone separation.
        /// </summary>
        public readonly bool EnableZoneSeparation;
        /// <summary>
        /// Option to enable and pass metadata for event publishing.
        /// </summary>
        public readonly Outputs.EventPublishConfigResponse EventPublishConfig;
        /// <summary>
        /// Cloud Storage bucket generated by Data Fusion in the customer project.
        /// </summary>
        public readonly string GcsBucket;
        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        public readonly Outputs.NetworkConfigResponse NetworkConfig;
        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Options;
        /// <summary>
        /// P4 service account for the customer project.
        /// </summary>
        public readonly string P4ServiceAccount;
        /// <summary>
        /// Optional. Current patch revision of the Data Fusion.
        /// </summary>
        public readonly string PatchRevision;
        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        /// </summary>
        public readonly bool PrivateInstance;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Endpoint on which the Data Fusion UI is accessible.
        /// </summary>
        public readonly string ServiceEndpoint;
        /// <summary>
        /// The current state of this Data Fusion instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the current state of this Data Fusion instance if available.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The name of the tenant project.
        /// </summary>
        public readonly string TenantProjectId;
        /// <summary>
        /// Instance type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The time the instance was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Current version of the Data Fusion. Only specifiable in Update.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Endpoint on which the Data Fusion UI is accessible to third-party users
        /// </summary>
        public readonly string WorkforceIdentityServiceEndpoint;
        /// <summary>
        /// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.AcceleratorResponse> accelerators,

            string apiEndpoint,

            ImmutableArray<Outputs.VersionResponse> availableVersion,

            string createTime,

            Outputs.CryptoKeyConfigResponse cryptoKeyConfig,

            bool dataplexDataLineageIntegrationEnabled,

            string dataprocServiceAccount,

            string description,

            ImmutableArray<string> disabledReason,

            string displayName,

            bool enableRbac,

            bool enableStackdriverLogging,

            bool enableStackdriverMonitoring,

            bool enableZoneSeparation,

            Outputs.EventPublishConfigResponse eventPublishConfig,

            string gcsBucket,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.NetworkConfigResponse networkConfig,

            ImmutableDictionary<string, string> options,

            string p4ServiceAccount,

            string patchRevision,

            bool privateInstance,

            bool satisfiesPzs,

            string serviceAccount,

            string serviceEndpoint,

            string state,

            string stateMessage,

            string tenantProjectId,

            string type,

            string updateTime,

            string version,

            string workforceIdentityServiceEndpoint,

            string zone)
        {
            Accelerators = accelerators;
            ApiEndpoint = apiEndpoint;
            AvailableVersion = availableVersion;
            CreateTime = createTime;
            CryptoKeyConfig = cryptoKeyConfig;
            DataplexDataLineageIntegrationEnabled = dataplexDataLineageIntegrationEnabled;
            DataprocServiceAccount = dataprocServiceAccount;
            Description = description;
            DisabledReason = disabledReason;
            DisplayName = displayName;
            EnableRbac = enableRbac;
            EnableStackdriverLogging = enableStackdriverLogging;
            EnableStackdriverMonitoring = enableStackdriverMonitoring;
            EnableZoneSeparation = enableZoneSeparation;
            EventPublishConfig = eventPublishConfig;
            GcsBucket = gcsBucket;
            Labels = labels;
            Name = name;
            NetworkConfig = networkConfig;
            Options = options;
            P4ServiceAccount = p4ServiceAccount;
            PatchRevision = patchRevision;
            PrivateInstance = privateInstance;
            SatisfiesPzs = satisfiesPzs;
            ServiceAccount = serviceAccount;
            ServiceEndpoint = serviceEndpoint;
            State = state;
            StateMessage = stateMessage;
            TenantProjectId = tenantProjectId;
            Type = type;
            UpdateTime = updateTime;
            Version = version;
            WorkforceIdentityServiceEndpoint = workforceIdentityServiceEndpoint;
            Zone = zone;
        }
    }
}
