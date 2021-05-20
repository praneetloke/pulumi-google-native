// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a task and adds it to a queue. Tasks cannot be updated after creation; there is no UpdateTask command. * The maximum task size is 100KB.
 */
export class QueueTask extends pulumi.CustomResource {
    /**
     * Get an existing QueueTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): QueueTask {
        return new QueueTask(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudtasks/v2:QueueTask';

    /**
     * Returns true if the given object is an instance of QueueTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueueTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueueTask.__pulumiType;
    }

    /**
     * HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
     */
    public readonly appEngineHttpRequest!: pulumi.Output<outputs.cloudtasks.v2.AppEngineHttpRequestResponse>;
    /**
     * The time that the task was created. `create_time` will be truncated to the nearest second.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
     */
    public readonly dispatchCount!: pulumi.Output<number>;
    /**
     * The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
     */
    public readonly dispatchDeadline!: pulumi.Output<string>;
    /**
     * The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
     */
    public readonly firstAttempt!: pulumi.Output<outputs.cloudtasks.v2.AttemptResponse>;
    /**
     * HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
     */
    public readonly httpRequest!: pulumi.Output<outputs.cloudtasks.v2.HttpRequestResponse>;
    /**
     * The status of the task's last attempt.
     */
    public readonly lastAttempt!: pulumi.Output<outputs.cloudtasks.v2.AttemptResponse>;
    /**
     * Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of attempts which have received a response.
     */
    public readonly responseCount!: pulumi.Output<number>;
    /**
     * The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
     */
    public readonly scheduleTime!: pulumi.Output<string>;
    /**
     * The view specifies which subset of the Task has been returned.
     */
    public readonly view!: pulumi.Output<string>;

    /**
     * Create a QueueTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueTaskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.queueId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queueId'");
            }
            if ((!args || args.taskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskId'");
            }
            inputs["appEngineHttpRequest"] = args ? args.appEngineHttpRequest : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["dispatchCount"] = args ? args.dispatchCount : undefined;
            inputs["dispatchDeadline"] = args ? args.dispatchDeadline : undefined;
            inputs["firstAttempt"] = args ? args.firstAttempt : undefined;
            inputs["httpRequest"] = args ? args.httpRequest : undefined;
            inputs["lastAttempt"] = args ? args.lastAttempt : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["queueId"] = args ? args.queueId : undefined;
            inputs["responseCount"] = args ? args.responseCount : undefined;
            inputs["responseView"] = args ? args.responseView : undefined;
            inputs["scheduleTime"] = args ? args.scheduleTime : undefined;
            inputs["taskId"] = args ? args.taskId : undefined;
            inputs["view"] = args ? args.view : undefined;
        } else {
            inputs["appEngineHttpRequest"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["dispatchCount"] = undefined /*out*/;
            inputs["dispatchDeadline"] = undefined /*out*/;
            inputs["firstAttempt"] = undefined /*out*/;
            inputs["httpRequest"] = undefined /*out*/;
            inputs["lastAttempt"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["responseCount"] = undefined /*out*/;
            inputs["scheduleTime"] = undefined /*out*/;
            inputs["view"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(QueueTask.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a QueueTask resource.
 */
export interface QueueTaskArgs {
    /**
     * HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
     */
    readonly appEngineHttpRequest?: pulumi.Input<inputs.cloudtasks.v2.AppEngineHttpRequestArgs>;
    /**
     * The time that the task was created. `create_time` will be truncated to the nearest second.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
     */
    readonly dispatchCount?: pulumi.Input<number>;
    /**
     * The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
     */
    readonly dispatchDeadline?: pulumi.Input<string>;
    /**
     * The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
     */
    readonly firstAttempt?: pulumi.Input<inputs.cloudtasks.v2.AttemptArgs>;
    /**
     * HTTP request that is sent to the worker. An HTTP task is a task that has HttpRequest set.
     */
    readonly httpRequest?: pulumi.Input<inputs.cloudtasks.v2.HttpRequestArgs>;
    /**
     * The status of the task's last attempt.
     */
    readonly lastAttempt?: pulumi.Input<inputs.cloudtasks.v2.AttemptArgs>;
    readonly location: pulumi.Input<string>;
    /**
     * Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly queueId: pulumi.Input<string>;
    /**
     * The number of attempts which have received a response.
     */
    readonly responseCount?: pulumi.Input<number>;
    /**
     * The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
     */
    readonly responseView?: pulumi.Input<string>;
    /**
     * The time when the task is scheduled to be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
     */
    readonly scheduleTime?: pulumi.Input<string>;
    readonly taskId: pulumi.Input<string>;
    /**
     * The view specifies which subset of the Task has been returned.
     */
    readonly view?: pulumi.Input<string>;
}
