# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetQueueResult',
    'AwaitableGetQueueResult',
    'get_queue',
    'get_queue_output',
]

@pulumi.output_type
class GetQueueResult:
    def __init__(__self__, app_engine_routing_override=None, name=None, purge_time=None, rate_limits=None, retry_config=None, stackdriver_logging_config=None, state=None):
        if app_engine_routing_override and not isinstance(app_engine_routing_override, dict):
            raise TypeError("Expected argument 'app_engine_routing_override' to be a dict")
        pulumi.set(__self__, "app_engine_routing_override", app_engine_routing_override)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if purge_time and not isinstance(purge_time, str):
            raise TypeError("Expected argument 'purge_time' to be a str")
        pulumi.set(__self__, "purge_time", purge_time)
        if rate_limits and not isinstance(rate_limits, dict):
            raise TypeError("Expected argument 'rate_limits' to be a dict")
        pulumi.set(__self__, "rate_limits", rate_limits)
        if retry_config and not isinstance(retry_config, dict):
            raise TypeError("Expected argument 'retry_config' to be a dict")
        pulumi.set(__self__, "retry_config", retry_config)
        if stackdriver_logging_config and not isinstance(stackdriver_logging_config, dict):
            raise TypeError("Expected argument 'stackdriver_logging_config' to be a dict")
        pulumi.set(__self__, "stackdriver_logging_config", stackdriver_logging_config)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="appEngineRoutingOverride")
    def app_engine_routing_override(self) -> 'outputs.AppEngineRoutingResponse':
        """
        Overrides for task-level app_engine_routing. These settings apply only to App Engine tasks in this queue. Http tasks are not affected. If set, `app_engine_routing_override` is used for all App Engine tasks in the queue, no matter what the setting is for the task-level app_engine_routing.
        """
        return pulumi.get(self, "app_engine_routing_override")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Caller-specified and required in CreateQueue, after which it becomes output only. The queue name. The queue name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the queue's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="purgeTime")
    def purge_time(self) -> str:
        """
        The last time this queue was purged. All tasks that were created before this time were purged. A queue can be purged using PurgeQueue, the [App Engine Task Queue SDK, or the Cloud Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue). Purge time will be truncated to the nearest microsecond. Purge time will be unset if the queue has never been purged.
        """
        return pulumi.get(self, "purge_time")

    @property
    @pulumi.getter(name="rateLimits")
    def rate_limits(self) -> 'outputs.RateLimitsResponse':
        """
        Rate limits for task dispatches. rate_limits and retry_config are related because they both control task attempts. However they control task attempts in different ways: * rate_limits controls the total rate of dispatches from a queue (i.e. all traffic dispatched from the queue, regardless of whether the dispatch is from a first attempt or a retry). * retry_config controls what happens to particular a task after its first attempt fails. That is, retry_config controls task retries (the second attempt, third attempt, etc). The queue's actual dispatch rate is the result of: * Number of tasks in the queue * User-specified throttling: rate_limits, retry_config, and the queue's state. * System throttling due to `429` (Too Many Requests) or `503` (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic spikes.
        """
        return pulumi.get(self, "rate_limits")

    @property
    @pulumi.getter(name="retryConfig")
    def retry_config(self) -> 'outputs.RetryConfigResponse':
        """
        Settings that determine the retry behavior. * For tasks created using Cloud Tasks: the queue-level retry settings apply to all tasks in the queue that were created using Cloud Tasks. Retry settings cannot be set on individual tasks. * For tasks created using the App Engine SDK: the queue-level retry settings apply to all tasks in the queue which do not have retry settings explicitly set on the task and were created by the App Engine SDK. See [App Engine documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
        """
        return pulumi.get(self, "retry_config")

    @property
    @pulumi.getter(name="stackdriverLoggingConfig")
    def stackdriver_logging_config(self) -> 'outputs.StackdriverLoggingConfigResponse':
        """
        Configuration options for writing logs to [Stackdriver Logging](https://cloud.google.com/logging/docs/). If this field is unset, then no logs are written.
        """
        return pulumi.get(self, "stackdriver_logging_config")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the queue. `state` can only be changed by calling PauseQueue, ResumeQueue, or uploading [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref). UpdateQueue cannot be used to change `state`.
        """
        return pulumi.get(self, "state")


class AwaitableGetQueueResult(GetQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueueResult(
            app_engine_routing_override=self.app_engine_routing_override,
            name=self.name,
            purge_time=self.purge_time,
            rate_limits=self.rate_limits,
            retry_config=self.retry_config,
            stackdriver_logging_config=self.stackdriver_logging_config,
            state=self.state)


def get_queue(location: Optional[str] = None,
              project: Optional[str] = None,
              queue_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQueueResult:
    """
    Gets a queue.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['queueId'] = queue_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudtasks/v2:getQueue', __args__, opts=opts, typ=GetQueueResult).value

    return AwaitableGetQueueResult(
        app_engine_routing_override=pulumi.get(__ret__, 'app_engine_routing_override'),
        name=pulumi.get(__ret__, 'name'),
        purge_time=pulumi.get(__ret__, 'purge_time'),
        rate_limits=pulumi.get(__ret__, 'rate_limits'),
        retry_config=pulumi.get(__ret__, 'retry_config'),
        stackdriver_logging_config=pulumi.get(__ret__, 'stackdriver_logging_config'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_queue)
def get_queue_output(location: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     queue_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQueueResult]:
    """
    Gets a queue.
    """
    ...
