// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha
{
    /// <summary>
    /// Create a integration with a draft version in the specified project.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:integrations/v1alpha:Version")]
    public partial class Version : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto-generated.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
        /// </summary>
        [Output("databasePersistencePolicy")]
        public Output<string> DatabasePersistencePolicy { get; private set; } = null!;

        /// <summary>
        /// Optional. The integration description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        [Output("integrationParameters")]
        public Output<ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaIntegrationParameterResponse>> IntegrationParameters { get; private set; } = null!;

        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        [Output("integrationParametersInternal")]
        public Output<Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse> IntegrationParametersInternal { get; private set; } = null!;

        /// <summary>
        /// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Output("lastModifierEmail")]
        public Output<string> LastModifierEmail { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Output("lockHolder")]
        public Output<string> LockHolder { get; private set; } = null!;

        /// <summary>
        /// Auto-generated primary key.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set this flag to true, if draft version is to be created for a brand new integration. False, if the request is for an existing integration. For backward compatibility reasons, even if this flag is set to `false` and no existing integration is found, a new draft integration will still be created.
        /// </summary>
        [Output("newIntegration")]
        public Output<bool?> NewIntegration { get; private set; } = null!;

        /// <summary>
        /// Optional. The origin that indicates where this integration is coming from.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// Optional. The id of the template which was used to create this integration_version.
        /// </summary>
        [Output("parentTemplateId")]
        public Output<string> ParentTemplateId { get; private set; } = null!;

        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
        /// </summary>
        [Output("snapshotNumber")]
        public Output<string> SnapshotNumber { get; private set; } = null!;

        /// <summary>
        /// User should not set it as an input.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Generated by eventbus. User should not set it as an input.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        [Output("taskConfigs")]
        public Output<ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTaskConfigResponse>> TaskConfigs { get; private set; } = null!;

        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        [Output("taskConfigsInternal")]
        public Output<ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse>> TaskConfigsInternal { get; private set; } = null!;

        /// <summary>
        /// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
        /// </summary>
        [Output("teardown")]
        public Output<Outputs.EnterpriseCrmEventbusProtoTeardownResponse> Teardown { get; private set; } = null!;

        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        [Output("triggerConfigs")]
        public Output<ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTriggerConfigResponse>> TriggerConfigs { get; private set; } = null!;

        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        [Output("triggerConfigsInternal")]
        public Output<ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse>> TriggerConfigsInternal { get; private set; } = null!;

        /// <summary>
        /// Auto-generated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
        /// </summary>
        [Output("userLabel")]
        public Output<string> UserLabel { get; private set; } = null!;


        /// <summary>
        /// Create a Version resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Version(string name, VersionArgs args, CustomResourceOptions? options = null)
            : base("google-native:integrations/v1alpha:Version", name, args ?? new VersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Version(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:integrations/v1alpha:Version", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "integrationId",
                    "location",
                    "productId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Version resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Version Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Version(name, id, options);
        }
    }

    public sealed class VersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
        /// </summary>
        [Input("databasePersistencePolicy")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.VersionDatabasePersistencePolicy>? DatabasePersistencePolicy { get; set; }

        /// <summary>
        /// Optional. The integration description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        [Input("integrationParameters")]
        private InputList<Inputs.GoogleCloudIntegrationsV1alphaIntegrationParameterArgs>? _integrationParameters;

        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        public InputList<Inputs.GoogleCloudIntegrationsV1alphaIntegrationParameterArgs> IntegrationParameters
        {
            get => _integrationParameters ?? (_integrationParameters = new InputList<Inputs.GoogleCloudIntegrationsV1alphaIntegrationParameterArgs>());
            set => _integrationParameters = value;
        }

        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        [Input("integrationParametersInternal")]
        public Input<Inputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersArgs>? IntegrationParametersInternal { get; set; }

        /// <summary>
        /// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Input("lastModifierEmail")]
        public Input<string>? LastModifierEmail { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Input("lockHolder")]
        public Input<string>? LockHolder { get; set; }

        /// <summary>
        /// Set this flag to true, if draft version is to be created for a brand new integration. False, if the request is for an existing integration. For backward compatibility reasons, even if this flag is set to `false` and no existing integration is found, a new draft integration will still be created.
        /// </summary>
        [Input("newIntegration")]
        public Input<bool>? NewIntegration { get; set; }

        /// <summary>
        /// Optional. The origin that indicates where this integration is coming from.
        /// </summary>
        [Input("origin")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.VersionOrigin>? Origin { get; set; }

        /// <summary>
        /// Optional. The id of the template which was used to create this integration_version.
        /// </summary>
        [Input("parentTemplateId")]
        public Input<string>? ParentTemplateId { get; set; }

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
        /// </summary>
        [Input("snapshotNumber")]
        public Input<string>? SnapshotNumber { get; set; }

        [Input("taskConfigs")]
        private InputList<Inputs.GoogleCloudIntegrationsV1alphaTaskConfigArgs>? _taskConfigs;

        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        public InputList<Inputs.GoogleCloudIntegrationsV1alphaTaskConfigArgs> TaskConfigs
        {
            get => _taskConfigs ?? (_taskConfigs = new InputList<Inputs.GoogleCloudIntegrationsV1alphaTaskConfigArgs>());
            set => _taskConfigs = value;
        }

        [Input("taskConfigsInternal")]
        private InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigArgs>? _taskConfigsInternal;

        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        public InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigArgs> TaskConfigsInternal
        {
            get => _taskConfigsInternal ?? (_taskConfigsInternal = new InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigArgs>());
            set => _taskConfigsInternal = value;
        }

        /// <summary>
        /// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
        /// </summary>
        [Input("teardown")]
        public Input<Inputs.EnterpriseCrmEventbusProtoTeardownArgs>? Teardown { get; set; }

        [Input("triggerConfigs")]
        private InputList<Inputs.GoogleCloudIntegrationsV1alphaTriggerConfigArgs>? _triggerConfigs;

        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        public InputList<Inputs.GoogleCloudIntegrationsV1alphaTriggerConfigArgs> TriggerConfigs
        {
            get => _triggerConfigs ?? (_triggerConfigs = new InputList<Inputs.GoogleCloudIntegrationsV1alphaTriggerConfigArgs>());
            set => _triggerConfigs = value;
        }

        [Input("triggerConfigsInternal")]
        private InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigArgs>? _triggerConfigsInternal;

        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        public InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigArgs> TriggerConfigsInternal
        {
            get => _triggerConfigsInternal ?? (_triggerConfigsInternal = new InputList<Inputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigArgs>());
            set => _triggerConfigsInternal = value;
        }

        /// <summary>
        /// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
        /// </summary>
        [Input("userLabel")]
        public Input<string>? UserLabel { get; set; }

        public VersionArgs()
        {
        }
        public static new VersionArgs Empty => new VersionArgs();
    }
}
