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
    'GetPatchDeploymentResult',
    'AwaitableGetPatchDeploymentResult',
    'get_patch_deployment',
    'get_patch_deployment_output',
]

@pulumi.output_type
class GetPatchDeploymentResult:
    def __init__(__self__, create_time=None, description=None, duration=None, instance_filter=None, last_execute_time=None, name=None, one_time_schedule=None, patch_config=None, recurring_schedule=None, rollout=None, state=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if duration and not isinstance(duration, str):
            raise TypeError("Expected argument 'duration' to be a str")
        pulumi.set(__self__, "duration", duration)
        if instance_filter and not isinstance(instance_filter, dict):
            raise TypeError("Expected argument 'instance_filter' to be a dict")
        pulumi.set(__self__, "instance_filter", instance_filter)
        if last_execute_time and not isinstance(last_execute_time, str):
            raise TypeError("Expected argument 'last_execute_time' to be a str")
        pulumi.set(__self__, "last_execute_time", last_execute_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if one_time_schedule and not isinstance(one_time_schedule, dict):
            raise TypeError("Expected argument 'one_time_schedule' to be a dict")
        pulumi.set(__self__, "one_time_schedule", one_time_schedule)
        if patch_config and not isinstance(patch_config, dict):
            raise TypeError("Expected argument 'patch_config' to be a dict")
        pulumi.set(__self__, "patch_config", patch_config)
        if recurring_schedule and not isinstance(recurring_schedule, dict):
            raise TypeError("Expected argument 'recurring_schedule' to be a dict")
        pulumi.set(__self__, "recurring_schedule", recurring_schedule)
        if rollout and not isinstance(rollout, dict):
            raise TypeError("Expected argument 'rollout' to be a dict")
        pulumi.set(__self__, "rollout", rollout)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        Optional. Duration of the patch. After the duration ends, the patch times out.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="instanceFilter")
    def instance_filter(self) -> 'outputs.PatchInstanceFilterResponse':
        """
        VM instances to patch.
        """
        return pulumi.get(self, "instance_filter")

    @property
    @pulumi.getter(name="lastExecuteTime")
    def last_execute_time(self) -> str:
        """
        The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "last_execute_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oneTimeSchedule")
    def one_time_schedule(self) -> 'outputs.OneTimeScheduleResponse':
        """
        Schedule a one-time execution.
        """
        return pulumi.get(self, "one_time_schedule")

    @property
    @pulumi.getter(name="patchConfig")
    def patch_config(self) -> 'outputs.PatchConfigResponse':
        """
        Optional. Patch configuration that is applied.
        """
        return pulumi.get(self, "patch_config")

    @property
    @pulumi.getter(name="recurringSchedule")
    def recurring_schedule(self) -> 'outputs.RecurringScheduleResponse':
        """
        Schedule recurring executions.
        """
        return pulumi.get(self, "recurring_schedule")

    @property
    @pulumi.getter
    def rollout(self) -> 'outputs.PatchRolloutResponse':
        """
        Optional. Rollout strategy of the patch job.
        """
        return pulumi.get(self, "rollout")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the patch deployment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetPatchDeploymentResult(GetPatchDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPatchDeploymentResult(
            create_time=self.create_time,
            description=self.description,
            duration=self.duration,
            instance_filter=self.instance_filter,
            last_execute_time=self.last_execute_time,
            name=self.name,
            one_time_schedule=self.one_time_schedule,
            patch_config=self.patch_config,
            recurring_schedule=self.recurring_schedule,
            rollout=self.rollout,
            state=self.state,
            update_time=self.update_time)


def get_patch_deployment(patch_deployment_id: Optional[str] = None,
                         project: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPatchDeploymentResult:
    """
    Get an OS Config patch deployment.
    """
    __args__ = dict()
    __args__['patchDeploymentId'] = patch_deployment_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:osconfig/v1beta:getPatchDeployment', __args__, opts=opts, typ=GetPatchDeploymentResult).value

    return AwaitableGetPatchDeploymentResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        duration=pulumi.get(__ret__, 'duration'),
        instance_filter=pulumi.get(__ret__, 'instance_filter'),
        last_execute_time=pulumi.get(__ret__, 'last_execute_time'),
        name=pulumi.get(__ret__, 'name'),
        one_time_schedule=pulumi.get(__ret__, 'one_time_schedule'),
        patch_config=pulumi.get(__ret__, 'patch_config'),
        recurring_schedule=pulumi.get(__ret__, 'recurring_schedule'),
        rollout=pulumi.get(__ret__, 'rollout'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_patch_deployment)
def get_patch_deployment_output(patch_deployment_id: Optional[pulumi.Input[str]] = None,
                                project: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPatchDeploymentResult]:
    """
    Get an OS Config patch deployment.
    """
    ...
