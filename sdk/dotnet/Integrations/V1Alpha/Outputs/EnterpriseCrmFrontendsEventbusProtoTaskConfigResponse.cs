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
    /// The task configuration details. This is not the implementation of Task. There might be multiple TaskConfigs for the same Task. Next available id: 27
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse
    {
        /// <summary>
        /// Alert configurations on error rate, warning rate, number of runs, durations, etc.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmEventbusProtoTaskAlertConfigResponse> AlertConfigs;
        /// <summary>
        /// Auto-generated.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The creator's email address. Auto-generated from the user's email.
        /// </summary>
        public readonly string CreatorEmail;
        /// <summary>
        /// User-provided description intended to give more business context about the task.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If this config contains a TypedTask, allow validation to succeed if an input is read from the output of another TypedTask whose output type is declared as a superclass of the requested input type. For instance, if the previous task declares an output of type Message, any task with this flag enabled will pass validation when attempting to read any proto Message type from the resultant Event parameter.
        /// </summary>
        public readonly bool DisableStrictTypeValidation;
        /// <summary>
        /// Optional. Determines the number of times the task will be retried on failure and with what retry strategy. This is applicable for asynchronous calls to Eventbus alone (Post To Queue, Schedule etc.).
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoFailurePolicyResponse FailurePolicy;
        /// <summary>
        /// The number of edges leading into this TaskConfig.
        /// </summary>
        public readonly int IncomingEdgeCount;
        /// <summary>
        /// If set, overrides the option configured in the Task implementation class.
        /// </summary>
        public readonly string JsonValidationOption;
        /// <summary>
        /// User-provided label that is attached to this TaskConfig in the UI.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Auto-generated.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The set of tasks that are next in line to be executed as per the execution graph defined for the parent event, specified by `event_config_id`. Each of these next tasks are executed only if the condition associated with them evaluates to true.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmEventbusProtoNextTaskResponse> NextTasks;
        /// <summary>
        /// The policy dictating the execution of the next set of tasks for the current task.
        /// </summary>
        public readonly string NextTasksExecutionPolicy;
        /// <summary>
        /// The customized parameters the user can pass to this task.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parameters;
        /// <summary>
        /// Optional. Informs the front-end application where to draw this task config on the UI.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoCoordinateResponse Position;
        /// <summary>
        /// Optional. Standard filter expression evaluated before execution. Independent of other conditions and tasks. Can be used to enable rollout. e.g. "rollout(5)" will only allow 5% of incoming traffic to task.
        /// </summary>
        public readonly string Precondition;
        /// <summary>
        /// Optional. User-provided label that is attached to precondition in the UI.
        /// </summary>
        public readonly string PreconditionLabel;
        /// <summary>
        /// Optional. Contains information about what needs to be done upon failure (either a permanent error or after it has been retried too many times).
        /// </summary>
        public readonly Outputs.EnterpriseCrmFrontendsEventbusProtoRollbackStrategyResponse RollbackStrategy;
        /// <summary>
        /// Determines what action to take upon successful task completion.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoSuccessPolicyResponse SuccessPolicy;
        /// <summary>
        /// Optional. Determines the number of times the task will be retried on failure and with what retry strategy. This is applicable for synchronous calls to Eventbus alone (Post).
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoFailurePolicyResponse SynchronousCallFailurePolicy;
        /// <summary>
        /// Copy of the task entity that this task config is an instance of.
        /// </summary>
        public readonly Outputs.EnterpriseCrmFrontendsEventbusProtoTaskEntityResponse TaskEntity;
        /// <summary>
        /// The policy dictating the execution strategy of this task.
        /// </summary>
        public readonly string TaskExecutionStrategy;
        /// <summary>
        /// The name for the task.
        /// </summary>
        public readonly string TaskName;
        /// <summary>
        /// REQUIRED: the identifier of this task within its parent event config, specified by the client. This should be unique among all the tasks belong to the same event config. We use this field as the identifier to find next tasks (via field `next_tasks.task_number`).
        /// </summary>
        public readonly string TaskNumber;
        /// <summary>
        /// A string template that allows user to configure task parameters (with either literal default values or tokens which will be resolved at execution time) for the task. It will eventually replace the old "parameters" field. Please refer to go/eventbus-task-spec-example for detailed usage example.
        /// </summary>
        public readonly string TaskSpec;
        /// <summary>
        /// Used to define task-template name if task is of type task-template
        /// </summary>
        public readonly string TaskTemplateName;
        /// <summary>
        /// Defines the type of the task
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse(
            ImmutableArray<Outputs.EnterpriseCrmEventbusProtoTaskAlertConfigResponse> alertConfigs,

            string createTime,

            string creatorEmail,

            string description,

            bool disableStrictTypeValidation,

            Outputs.EnterpriseCrmEventbusProtoFailurePolicyResponse failurePolicy,

            int incomingEdgeCount,

            string jsonValidationOption,

            string label,

            string lastModifiedTime,

            ImmutableArray<Outputs.EnterpriseCrmEventbusProtoNextTaskResponse> nextTasks,

            string nextTasksExecutionPolicy,

            ImmutableDictionary<string, string> parameters,

            Outputs.EnterpriseCrmEventbusProtoCoordinateResponse position,

            string precondition,

            string preconditionLabel,

            Outputs.EnterpriseCrmFrontendsEventbusProtoRollbackStrategyResponse rollbackStrategy,

            Outputs.EnterpriseCrmEventbusProtoSuccessPolicyResponse successPolicy,

            Outputs.EnterpriseCrmEventbusProtoFailurePolicyResponse synchronousCallFailurePolicy,

            Outputs.EnterpriseCrmFrontendsEventbusProtoTaskEntityResponse taskEntity,

            string taskExecutionStrategy,

            string taskName,

            string taskNumber,

            string taskSpec,

            string taskTemplateName,

            string taskType)
        {
            AlertConfigs = alertConfigs;
            CreateTime = createTime;
            CreatorEmail = creatorEmail;
            Description = description;
            DisableStrictTypeValidation = disableStrictTypeValidation;
            FailurePolicy = failurePolicy;
            IncomingEdgeCount = incomingEdgeCount;
            JsonValidationOption = jsonValidationOption;
            Label = label;
            LastModifiedTime = lastModifiedTime;
            NextTasks = nextTasks;
            NextTasksExecutionPolicy = nextTasksExecutionPolicy;
            Parameters = parameters;
            Position = position;
            Precondition = precondition;
            PreconditionLabel = preconditionLabel;
            RollbackStrategy = rollbackStrategy;
            SuccessPolicy = successPolicy;
            SynchronousCallFailurePolicy = synchronousCallFailurePolicy;
            TaskEntity = taskEntity;
            TaskExecutionStrategy = taskExecutionStrategy;
            TaskName = taskName;
            TaskNumber = taskNumber;
            TaskSpec = taskSpec;
            TaskTemplateName = taskTemplateName;
            TaskType = taskType;
        }
    }
}
