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
    'GetDlpJobResult',
    'AwaitableGetDlpJobResult',
    'get_dlp_job',
    'get_dlp_job_output',
]

@pulumi.output_type
class GetDlpJobResult:
    def __init__(__self__, action_details=None, create_time=None, end_time=None, errors=None, inspect_details=None, job_trigger_name=None, name=None, risk_details=None, start_time=None, state=None, type=None):
        if action_details and not isinstance(action_details, list):
            raise TypeError("Expected argument 'action_details' to be a list")
        pulumi.set(__self__, "action_details", action_details)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if errors and not isinstance(errors, list):
            raise TypeError("Expected argument 'errors' to be a list")
        pulumi.set(__self__, "errors", errors)
        if inspect_details and not isinstance(inspect_details, dict):
            raise TypeError("Expected argument 'inspect_details' to be a dict")
        pulumi.set(__self__, "inspect_details", inspect_details)
        if job_trigger_name and not isinstance(job_trigger_name, str):
            raise TypeError("Expected argument 'job_trigger_name' to be a str")
        pulumi.set(__self__, "job_trigger_name", job_trigger_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if risk_details and not isinstance(risk_details, dict):
            raise TypeError("Expected argument 'risk_details' to be a dict")
        pulumi.set(__self__, "risk_details", risk_details)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="actionDetails")
    def action_details(self) -> Sequence['outputs.GooglePrivacyDlpV2ActionDetailsResponse']:
        """
        Events that should occur after the job has completed.
        """
        return pulumi.get(self, "action_details")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the job was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Time when the job finished.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def errors(self) -> Sequence['outputs.GooglePrivacyDlpV2ErrorResponse']:
        """
        A stream of errors encountered running the job.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="inspectDetails")
    def inspect_details(self) -> 'outputs.GooglePrivacyDlpV2InspectDataSourceDetailsResponse':
        """
        Results from inspecting a data source.
        """
        return pulumi.get(self, "inspect_details")

    @property
    @pulumi.getter(name="jobTriggerName")
    def job_trigger_name(self) -> str:
        """
        If created by a job trigger, the resource name of the trigger that instantiated the job.
        """
        return pulumi.get(self, "job_trigger_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The server-assigned name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="riskDetails")
    def risk_details(self) -> 'outputs.GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse':
        """
        Results from analyzing risk of a data source.
        """
        return pulumi.get(self, "risk_details")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Time when the job started.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of a job.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of job.
        """
        return pulumi.get(self, "type")


class AwaitableGetDlpJobResult(GetDlpJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDlpJobResult(
            action_details=self.action_details,
            create_time=self.create_time,
            end_time=self.end_time,
            errors=self.errors,
            inspect_details=self.inspect_details,
            job_trigger_name=self.job_trigger_name,
            name=self.name,
            risk_details=self.risk_details,
            start_time=self.start_time,
            state=self.state,
            type=self.type)


def get_dlp_job(dlp_job_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDlpJobResult:
    """
    Gets the latest state of a long-running DlpJob. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more.
    """
    __args__ = dict()
    __args__['dlpJobId'] = dlp_job_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dlp/v2:getDlpJob', __args__, opts=opts, typ=GetDlpJobResult).value

    return AwaitableGetDlpJobResult(
        action_details=pulumi.get(__ret__, 'action_details'),
        create_time=pulumi.get(__ret__, 'create_time'),
        end_time=pulumi.get(__ret__, 'end_time'),
        errors=pulumi.get(__ret__, 'errors'),
        inspect_details=pulumi.get(__ret__, 'inspect_details'),
        job_trigger_name=pulumi.get(__ret__, 'job_trigger_name'),
        name=pulumi.get(__ret__, 'name'),
        risk_details=pulumi.get(__ret__, 'risk_details'),
        start_time=pulumi.get(__ret__, 'start_time'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_dlp_job)
def get_dlp_job_output(dlp_job_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDlpJobResult]:
    """
    Gets the latest state of a long-running DlpJob. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more.
    """
    ...
