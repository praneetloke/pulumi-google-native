// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Scheduled Notebook in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:notebooks/v1:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    /**
     * Time the schedule was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Cron-tab formatted schedule by which the job will execute Format: minute, hour, day of month, month, day of week e.g. 0 0 * * WED = every Wednesday More examples: https://crontab.guru/examples.html
     */
    public readonly cronSchedule!: pulumi.Output<string>;
    /**
     * A brief description of this environment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name used for UI purposes. Name can only contain alphanumeric characters, hyphens '-', and underscores '_'.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * Notebook Execution Template corresponding to this schedule.
     */
    public readonly executionTemplate!: pulumi.Output<outputs.notebooks.v1.ExecutionTemplateResponse>;
    /**
     * The name of this schedule. Format: `projects/{project_id}/locations/{location}/schedules/{schedule_id}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The most recent execution names triggered from this schedule and their corresponding states.
     */
    public /*out*/ readonly recentExecutions!: pulumi.Output<outputs.notebooks.v1.ExecutionResponse[]>;
    public readonly state!: pulumi.Output<string>;
    /**
     * Timezone on which the cron_schedule. The value of this field must be a time zone name from the tz database. TZ Database: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * Time the schedule was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.scheduleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduleId'");
            }
            inputs["cronSchedule"] = args ? args.cronSchedule : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["executionTemplate"] = args ? args.executionTemplate : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["scheduleId"] = args ? args.scheduleId : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recentExecutions"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["cronSchedule"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["executionTemplate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recentExecutions"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Schedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    /**
     * Cron-tab formatted schedule by which the job will execute Format: minute, hour, day of month, month, day of week e.g. 0 0 * * WED = every Wednesday More examples: https://crontab.guru/examples.html
     */
    cronSchedule?: pulumi.Input<string>;
    /**
     * A brief description of this environment.
     */
    description?: pulumi.Input<string>;
    /**
     * Notebook Execution Template corresponding to this schedule.
     */
    executionTemplate?: pulumi.Input<inputs.notebooks.v1.ExecutionTemplateArgs>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    scheduleId: pulumi.Input<string>;
    state?: pulumi.Input<enums.notebooks.v1.ScheduleState>;
    /**
     * Timezone on which the cron_schedule. The value of this field must be a time zone name from the tz database. TZ Database: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
     */
    timeZone?: pulumi.Input<string>;
}
