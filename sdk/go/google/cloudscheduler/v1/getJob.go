// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("google-native:cloudscheduler/v1:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId    string  `pulumi:"jobId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupJobResult struct {
	// App Engine HTTP target.
	AppEngineHttpTarget AppEngineHttpTargetResponse `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The default and the allowed values depend on the type of target: * For HTTP targets, the default is 3 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine HTTP targets, 0 indicates that the request has the default deadline. The default deadline depends on the scaling type of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. * For Pub/Sub targets, this field is ignored.
	AttemptDeadline string `pulumi:"attemptDeadline"`
	// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
	Description string `pulumi:"description"`
	// HTTP target.
	HttpTarget HttpTargetResponse `pulumi:"httpTarget"`
	// The time the last job attempt started.
	LastAttemptTime string `pulumi:"lastAttemptTime"`
	// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `pulumi:"name"`
	// Pub/Sub target.
	PubsubTarget PubsubTargetResponse `pulumi:"pubsubTarget"`
	// Settings that determine the retry behavior.
	RetryConfig RetryConfigResponse `pulumi:"retryConfig"`
	// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
	Schedule string `pulumi:"schedule"`
	// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
	ScheduleTime string `pulumi:"scheduleTime"`
	// State of the job.
	State string `pulumi:"state"`
	// The response from the target for the last attempted execution.
	Status StatusResponse `pulumi:"status"`
	// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
	TimeZone string `pulumi:"timeZone"`
	// The creation time of the job.
	UserUpdateTime string `pulumi:"userUpdateTime"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobId    pulumi.StringInput    `pulumi:"jobId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// App Engine HTTP target.
func (o LookupJobResultOutput) AppEngineHttpTarget() AppEngineHttpTargetResponseOutput {
	return o.ApplyT(func(v LookupJobResult) AppEngineHttpTargetResponse { return v.AppEngineHttpTarget }).(AppEngineHttpTargetResponseOutput)
}

// The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The default and the allowed values depend on the type of target: * For HTTP targets, the default is 3 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine HTTP targets, 0 indicates that the request has the default deadline. The default deadline depends on the scaling type of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. * For Pub/Sub targets, this field is ignored.
func (o LookupJobResultOutput) AttemptDeadline() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.AttemptDeadline }).(pulumi.StringOutput)
}

// Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
func (o LookupJobResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Description }).(pulumi.StringOutput)
}

// HTTP target.
func (o LookupJobResultOutput) HttpTarget() HttpTargetResponseOutput {
	return o.ApplyT(func(v LookupJobResult) HttpTargetResponse { return v.HttpTarget }).(HttpTargetResponseOutput)
}

// The time the last job attempt started.
func (o LookupJobResultOutput) LastAttemptTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.LastAttemptTime }).(pulumi.StringOutput)
}

// Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Pub/Sub target.
func (o LookupJobResultOutput) PubsubTarget() PubsubTargetResponseOutput {
	return o.ApplyT(func(v LookupJobResult) PubsubTargetResponse { return v.PubsubTarget }).(PubsubTargetResponseOutput)
}

// Settings that determine the retry behavior.
func (o LookupJobResultOutput) RetryConfig() RetryConfigResponseOutput {
	return o.ApplyT(func(v LookupJobResult) RetryConfigResponse { return v.RetryConfig }).(RetryConfigResponseOutput)
}

// Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
func (o LookupJobResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Schedule }).(pulumi.StringOutput)
}

// The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
func (o LookupJobResultOutput) ScheduleTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ScheduleTime }).(pulumi.StringOutput)
}

// State of the job.
func (o LookupJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.State }).(pulumi.StringOutput)
}

// The response from the target for the last attempted execution.
func (o LookupJobResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupJobResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
func (o LookupJobResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// The creation time of the job.
func (o LookupJobResultOutput) UserUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.UserUpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
