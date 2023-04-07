// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha
{
    public static class GetVersion
    {
        /// <summary>
        /// Get a integration in the specified project.
        /// </summary>
        public static Task<GetVersionResult> InvokeAsync(GetVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVersionResult>("google-native:integrations/v1alpha:getVersion", args ?? new GetVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Get a integration in the specified project.
        /// </summary>
        public static Output<GetVersionResult> Invoke(GetVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVersionResult>("google-native:integrations/v1alpha:getVersion", args ?? new GetVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public string IntegrationId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetVersionArgs()
        {
        }
        public static new GetVersionArgs Empty => new GetVersionArgs();
    }

    public sealed class GetVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetVersionInvokeArgs()
        {
        }
        public static new GetVersionInvokeArgs Empty => new GetVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVersionResult
    {
        /// <summary>
        /// Auto-generated.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
        /// </summary>
        public readonly string DatabasePersistencePolicy;
        /// <summary>
        /// Optional. The integration description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Error Catch Task configuration for the integration. It's optional.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse> ErrorCatcherConfigs;
        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaIntegrationParameterResponse> IntegrationParameters;
        /// <summary>
        /// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
        /// </summary>
        public readonly Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse IntegrationParametersInternal;
        /// <summary>
        /// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        public readonly string LastModifierEmail;
        /// <summary>
        /// Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        public readonly string LockHolder;
        /// <summary>
        /// Auto-generated primary key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. The origin that indicates where this integration is coming from.
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// Optional. The id of the template which was used to create this integration_version.
        /// </summary>
        public readonly string ParentTemplateId;
        /// <summary>
        /// Optional. The run-as service account email, if set and auth config is not configured, that will be used to generate auth token to be used in Connector task, Rest caller task and Cloud function task.
        /// </summary>
        public readonly string RunAsServiceAccount;
        /// <summary>
        /// Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
        /// </summary>
        public readonly string SnapshotNumber;
        /// <summary>
        /// User should not set it as an input.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Generated by eventbus. User should not set it as an input.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTaskConfigResponse> TaskConfigs;
        /// <summary>
        /// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse> TaskConfigsInternal;
        /// <summary>
        /// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoTeardownResponse Teardown;
        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTriggerConfigResponse> TriggerConfigs;
        /// <summary>
        /// Optional. Trigger configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse> TriggerConfigsInternal;
        /// <summary>
        /// Auto-generated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
        /// </summary>
        public readonly string UserLabel;

        [OutputConstructor]
        private GetVersionResult(
            string createTime,

            string databasePersistencePolicy,

            string description,

            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse> errorCatcherConfigs,

            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaIntegrationParameterResponse> integrationParameters,

            Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse integrationParametersInternal,

            string lastModifierEmail,

            string lockHolder,

            string name,

            string origin,

            string parentTemplateId,

            string runAsServiceAccount,

            string snapshotNumber,

            string state,

            string status,

            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTaskConfigResponse> taskConfigs,

            ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse> taskConfigsInternal,

            Outputs.EnterpriseCrmEventbusProtoTeardownResponse teardown,

            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaTriggerConfigResponse> triggerConfigs,

            ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse> triggerConfigsInternal,

            string updateTime,

            string userLabel)
        {
            CreateTime = createTime;
            DatabasePersistencePolicy = databasePersistencePolicy;
            Description = description;
            ErrorCatcherConfigs = errorCatcherConfigs;
            IntegrationParameters = integrationParameters;
            IntegrationParametersInternal = integrationParametersInternal;
            LastModifierEmail = lastModifierEmail;
            LockHolder = lockHolder;
            Name = name;
            Origin = origin;
            ParentTemplateId = parentTemplateId;
            RunAsServiceAccount = runAsServiceAccount;
            SnapshotNumber = snapshotNumber;
            State = state;
            Status = status;
            TaskConfigs = taskConfigs;
            TaskConfigsInternal = taskConfigsInternal;
            Teardown = teardown;
            TriggerConfigs = triggerConfigs;
            TriggerConfigsInternal = triggerConfigsInternal;
            UpdateTime = updateTime;
            UserLabel = userLabel;
        }
    }
}
