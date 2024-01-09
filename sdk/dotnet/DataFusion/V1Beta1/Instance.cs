// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataFusion.V1Beta1
{
    /// <summary>
    /// Creates a new Data Fusion instance in the specified project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:datafusion/v1beta1:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of accelerators enabled for this CDF instance.
        /// </summary>
        [Output("accelerators")]
        public Output<ImmutableArray<Outputs.AcceleratorResponse>> Accelerators { get; private set; } = null!;

        /// <summary>
        /// Endpoint on which the REST APIs is accessible.
        /// </summary>
        [Output("apiEndpoint")]
        public Output<string> ApiEndpoint { get; private set; } = null!;

        /// <summary>
        /// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        /// </summary>
        [Output("availableVersion")]
        public Output<ImmutableArray<Outputs.VersionResponse>> AvailableVersion { get; private set; } = null!;

        /// <summary>
        /// The time the instance was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
        /// </summary>
        [Output("cryptoKeyConfig")]
        public Output<Outputs.CryptoKeyConfigResponse> CryptoKeyConfig { get; private set; } = null!;

        /// <summary>
        /// Optional. Reserved for future use.
        /// </summary>
        [Output("dataplexDataLineageIntegrationEnabled")]
        public Output<bool> DataplexDataLineageIntegrationEnabled { get; private set; } = null!;

        /// <summary>
        /// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        /// </summary>
        [Output("dataprocServiceAccount")]
        public Output<string> DataprocServiceAccount { get; private set; } = null!;

        /// <summary>
        /// A description of this instance.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If the instance state is DISABLED, the reason for disabling the instance.
        /// </summary>
        [Output("disabledReason")]
        public Output<ImmutableArray<string>> DisabledReason { get; private set; } = null!;

        /// <summary>
        /// Display name for an instance.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Option to enable granular role-based access control.
        /// </summary>
        [Output("enableRbac")]
        public Output<bool> EnableRbac { get; private set; } = null!;

        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        [Output("enableStackdriverLogging")]
        public Output<bool> EnableStackdriverLogging { get; private set; } = null!;

        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        [Output("enableStackdriverMonitoring")]
        public Output<bool> EnableStackdriverMonitoring { get; private set; } = null!;

        /// <summary>
        /// Option to enable zone separation.
        /// </summary>
        [Output("enableZoneSeparation")]
        public Output<bool> EnableZoneSeparation { get; private set; } = null!;

        /// <summary>
        /// Option to enable and pass metadata for event publishing.
        /// </summary>
        [Output("eventPublishConfig")]
        public Output<Outputs.EventPublishConfigResponse> EventPublishConfig { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage bucket generated by Data Fusion in the customer project.
        /// </summary>
        [Output("gcsBucket")]
        public Output<string> GcsBucket { get; private set; } = null!;

        /// <summary>
        /// Required. The name of the instance to create.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, string>> Options { get; private set; } = null!;

        /// <summary>
        /// P4 service account for the customer project.
        /// </summary>
        [Output("p4ServiceAccount")]
        public Output<string> P4ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Optional. Current patch revision of the Data Fusion.
        /// </summary>
        [Output("patchRevision")]
        public Output<string> PatchRevision { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        /// </summary>
        [Output("privateInstance")]
        public Output<bool> PrivateInstance { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Endpoint on which the Data Fusion UI is accessible.
        /// </summary>
        [Output("serviceEndpoint")]
        public Output<string> ServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// The current state of this Data Fusion instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Additional information about the current state of this Data Fusion instance if available.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// The name of the tenant project.
        /// </summary>
        [Output("tenantProjectId")]
        public Output<string> TenantProjectId { get; private set; } = null!;

        /// <summary>
        /// Instance type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time the instance was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Current version of Data Fusion.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Endpoint on which the Data Fusion UI is accessible to third-party users.
        /// </summary>
        [Output("workforceIdentityServiceEndpoint")]
        public Output<string> WorkforceIdentityServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("google-native:datafusion/v1beta1:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datafusion/v1beta1:Instance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "instanceId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
        /// </summary>
        [Input("cryptoKeyConfig")]
        public Input<Inputs.CryptoKeyConfigArgs>? CryptoKeyConfig { get; set; }

        /// <summary>
        /// Optional. Reserved for future use.
        /// </summary>
        [Input("dataplexDataLineageIntegrationEnabled")]
        public Input<bool>? DataplexDataLineageIntegrationEnabled { get; set; }

        /// <summary>
        /// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        /// </summary>
        [Input("dataprocServiceAccount")]
        public Input<string>? DataprocServiceAccount { get; set; }

        /// <summary>
        /// A description of this instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name for an instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Option to enable granular role-based access control.
        /// </summary>
        [Input("enableRbac")]
        public Input<bool>? EnableRbac { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Logging.
        /// </summary>
        [Input("enableStackdriverLogging")]
        public Input<bool>? EnableStackdriverLogging { get; set; }

        /// <summary>
        /// Option to enable Stackdriver Monitoring.
        /// </summary>
        [Input("enableStackdriverMonitoring")]
        public Input<bool>? EnableStackdriverMonitoring { get; set; }

        /// <summary>
        /// Option to enable zone separation.
        /// </summary>
        [Input("enableZoneSeparation")]
        public Input<bool>? EnableZoneSeparation { get; set; }

        /// <summary>
        /// Option to enable and pass metadata for event publishing.
        /// </summary>
        [Input("eventPublishConfig")]
        public Input<Inputs.EventPublishConfigArgs>? EventPublishConfig { get; set; }

        /// <summary>
        /// Required. The name of the instance to create.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Network configuration options. These are required when a private Data Fusion instance is to be created.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        [Input("options")]
        private InputMap<string>? _options;

        /// <summary>
        /// Map of additional options used to configure the behavior of Data Fusion instance.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// Optional. Current patch revision of the Data Fusion.
        /// </summary>
        [Input("patchRevision")]
        public Input<string>? PatchRevision { get; set; }

        /// <summary>
        /// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        /// </summary>
        [Input("privateInstance")]
        public Input<bool>? PrivateInstance { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Instance type.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.DataFusion.V1Beta1.InstanceType> Type { get; set; } = null!;

        /// <summary>
        /// Current version of Data Fusion.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }
}
