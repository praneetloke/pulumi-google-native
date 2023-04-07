# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetRunResult',
    'AwaitableGetRunResult',
    'get_run',
    'get_run_output',
]

@pulumi.output_type
class GetRunResult:
    def __init__(__self__, attributes=None, display_name=None, end_time=None, name=None, start_time=None, state=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, str]:
        """
        Optional. The attributes of the run. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the run). Up to 100 attributes are allowed.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. A human-readable name you can set to display in a user interface. Must be not longer than 1024 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Optional. The timestamp of the end of the run.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The resource name of the run. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}`. Can be specified or auto-assigned. {run} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        The timestamp of the start of the run.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the run.
        """
        return pulumi.get(self, "state")


class AwaitableGetRunResult(GetRunResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRunResult(
            attributes=self.attributes,
            display_name=self.display_name,
            end_time=self.end_time,
            name=self.name,
            start_time=self.start_time,
            state=self.state)


def get_run(location: Optional[str] = None,
            process_id: Optional[str] = None,
            project: Optional[str] = None,
            run_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRunResult:
    """
    Gets the details of the specified run.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['processId'] = process_id
    __args__['project'] = project
    __args__['runId'] = run_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datalineage/v1:getRun', __args__, opts=opts, typ=GetRunResult).value

    return AwaitableGetRunResult(
        attributes=__ret__.attributes,
        display_name=__ret__.display_name,
        end_time=__ret__.end_time,
        name=__ret__.name,
        start_time=__ret__.start_time,
        state=__ret__.state)


@_utilities.lift_output_func(get_run)
def get_run_output(location: Optional[pulumi.Input[str]] = None,
                   process_id: Optional[pulumi.Input[str]] = None,
                   project: Optional[pulumi.Input[Optional[str]]] = None,
                   run_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRunResult]:
    """
    Gets the details of the specified run.
    """
    ...
