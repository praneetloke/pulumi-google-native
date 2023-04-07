// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Configuration detail of a trigger.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaTriggerConfigResponse
    {
        /// <summary>
        /// Optional. An alert threshold configuration for the [trigger + client + integration] tuple. If these values are not specified in the trigger config, default values will be populated by the system. Note that there must be exactly one alert threshold configured per [client + trigger + integration] when published.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaIntegrationAlertConfigResponse> AlertConfig;
        /// <summary>
        /// Optional. Cloud Scheduler Trigger related metadata
        /// </summary>
        public readonly Outputs.GoogleCloudIntegrationsV1alphaCloudSchedulerConfigResponse CloudSchedulerConfig;
        /// <summary>
        /// Optional. User-provided description intended to give additional business context about the task.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Optional Error catcher id of the error catch flow which will be executed when execution error happens in the task
        /// </summary>
        public readonly string ErrorCatcherId;
        /// <summary>
        /// Optional. The user created label for a particular trigger.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Optional. Dictates how next tasks will be executed.
        /// </summary>
        public readonly string NextTasksExecutionPolicy;
        /// <summary>
        /// Optional. Informs the front-end application where to draw this error catcher config on the UI.
        /// </summary>
        public readonly Outputs.GoogleCloudIntegrationsV1alphaCoordinateResponse Position;
        /// <summary>
        /// Optional. Configurable properties of the trigger, not to be confused with integration parameters. E.g. "name" is a property for API triggers and "subscription" is a property for Pub/sub triggers.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Optional. Set of tasks numbers from where the integration execution is started by this trigger. If this is empty, then integration is executed with default start tasks. In the list of start tasks, none of two tasks can have direct ancestor-descendant relationships (i.e. in a same integration execution graph).
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaNextTaskResponse> StartTasks;
        /// <summary>
        /// Optional. The backend trigger ID.
        /// </summary>
        public readonly string TriggerId;
        /// <summary>
        /// A number to uniquely identify each trigger config within the integration on UI.
        /// </summary>
        public readonly string TriggerNumber;
        /// <summary>
        /// Optional. Type of trigger
        /// </summary>
        public readonly string TriggerType;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaTriggerConfigResponse(
            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaIntegrationAlertConfigResponse> alertConfig,

            Outputs.GoogleCloudIntegrationsV1alphaCloudSchedulerConfigResponse cloudSchedulerConfig,

            string description,

            string errorCatcherId,

            string label,

            string nextTasksExecutionPolicy,

            Outputs.GoogleCloudIntegrationsV1alphaCoordinateResponse position,

            ImmutableDictionary<string, string> properties,

            ImmutableArray<Outputs.GoogleCloudIntegrationsV1alphaNextTaskResponse> startTasks,

            string triggerId,

            string triggerNumber,

            string triggerType)
        {
            AlertConfig = alertConfig;
            CloudSchedulerConfig = cloudSchedulerConfig;
            Description = description;
            ErrorCatcherId = errorCatcherId;
            Label = label;
            NextTasksExecutionPolicy = nextTasksExecutionPolicy;
            Position = position;
            Properties = properties;
            StartTasks = startTasks;
            TriggerId = triggerId;
            TriggerNumber = triggerNumber;
            TriggerType = triggerType;
        }
    }
}
