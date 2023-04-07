// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The task configuration details. This is not the implementation of Task. There might be multiple TaskConfigs for the same Task.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaTaskConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User-provided description intended to give additional business context about the task.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. User-provided label that is attached to this TaskConfig in the UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Optional Error catcher id of the error catch flow which will be executed when execution error happens in the task
        /// </summary>
        [Input("errorCatcherId")]
        public Input<string>? ErrorCatcherId { get; set; }

        /// <summary>
        /// Optional. External task type of the task
        /// </summary>
        [Input("externalTaskType")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaTaskConfigExternalTaskType>? ExternalTaskType { get; set; }

        /// <summary>
        /// Optional. Determines the number of times the task will be retried on failure and with what retry strategy. This is applicable for asynchronous calls to Eventbus alone (Post To Queue, Schedule etc.).
        /// </summary>
        [Input("failurePolicy")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaFailurePolicyArgs>? FailurePolicy { get; set; }

        /// <summary>
        /// Optional. If set, overrides the option configured in the Task implementation class.
        /// </summary>
        [Input("jsonValidationOption")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaTaskConfigJsonValidationOption>? JsonValidationOption { get; set; }

        [Input("nextTasks")]
        private InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs>? _nextTasks;

        /// <summary>
        /// Optional. The set of tasks that are next in line to be executed as per the execution graph defined for the parent event, specified by `event_config_id`. Each of these next tasks are executed only if the condition associated with them evaluates to true.
        /// </summary>
        public InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs> NextTasks
        {
            get => _nextTasks ?? (_nextTasks = new InputList<Inputs.GoogleCloudIntegrationsV1alphaNextTaskArgs>());
            set => _nextTasks = value;
        }

        /// <summary>
        /// Optional. The policy dictating the execution of the next set of tasks for the current task.
        /// </summary>
        [Input("nextTasksExecutionPolicy")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaTaskConfigNextTasksExecutionPolicy>? NextTasksExecutionPolicy { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Optional. The customized parameters the user can pass to this task.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Optional. Informs the front-end application where to draw this error catcher config on the UI.
        /// </summary>
        [Input("position")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaCoordinateArgs>? Position { get; set; }

        /// <summary>
        /// Optional. Determines what action to take upon successful task completion.
        /// </summary>
        [Input("successPolicy")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaSuccessPolicyArgs>? SuccessPolicy { get; set; }

        /// <summary>
        /// Optional. Determines the number of times the task will be retried on failure and with what retry strategy. This is applicable for synchronous calls to Eventbus alone (Post).
        /// </summary>
        [Input("synchronousCallFailurePolicy")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaFailurePolicyArgs>? SynchronousCallFailurePolicy { get; set; }

        /// <summary>
        /// Optional. The name for the task.
        /// </summary>
        [Input("task")]
        public Input<string>? Task { get; set; }

        /// <summary>
        /// Optional. The policy dictating the execution strategy of this task.
        /// </summary>
        [Input("taskExecutionStrategy")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaTaskConfigTaskExecutionStrategy>? TaskExecutionStrategy { get; set; }

        /// <summary>
        /// The identifier of this task within its parent event config, specified by the client. This should be unique among all the tasks belong to the same event config. We use this field as the identifier to find next tasks (via field `next_tasks.task_id`).
        /// </summary>
        [Input("taskId", required: true)]
        public Input<string> TaskId { get; set; } = null!;

        /// <summary>
        /// Optional. Used to define task-template name if task is of type task-template
        /// </summary>
        [Input("taskTemplate")]
        public Input<string>? TaskTemplate { get; set; }

        public GoogleCloudIntegrationsV1alphaTaskConfigArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaTaskConfigArgs Empty => new GoogleCloudIntegrationsV1alphaTaskConfigArgs();
    }
}
