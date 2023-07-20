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
    'GetServiceLevelObjectiveResult',
    'AwaitableGetServiceLevelObjectiveResult',
    'get_service_level_objective',
    'get_service_level_objective_output',
]

@pulumi.output_type
class GetServiceLevelObjectiveResult:
    def __init__(__self__, calendar_period=None, display_name=None, goal=None, name=None, rolling_period=None, service_level_indicator=None, user_labels=None):
        if calendar_period and not isinstance(calendar_period, str):
            raise TypeError("Expected argument 'calendar_period' to be a str")
        pulumi.set(__self__, "calendar_period", calendar_period)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if goal and not isinstance(goal, float):
            raise TypeError("Expected argument 'goal' to be a float")
        pulumi.set(__self__, "goal", goal)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rolling_period and not isinstance(rolling_period, str):
            raise TypeError("Expected argument 'rolling_period' to be a str")
        pulumi.set(__self__, "rolling_period", rolling_period)
        if service_level_indicator and not isinstance(service_level_indicator, dict):
            raise TypeError("Expected argument 'service_level_indicator' to be a dict")
        pulumi.set(__self__, "service_level_indicator", service_level_indicator)
        if user_labels and not isinstance(user_labels, dict):
            raise TypeError("Expected argument 'user_labels' to be a dict")
        pulumi.set(__self__, "user_labels", user_labels)

    @property
    @pulumi.getter(name="calendarPeriod")
    def calendar_period(self) -> str:
        """
        A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
        """
        return pulumi.get(self, "calendar_period")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Name used for UI elements listing this SLO.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def goal(self) -> float:
        """
        The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
        """
        return pulumi.get(self, "goal")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME] 
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rollingPeriod")
    def rolling_period(self) -> str:
        """
        A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
        """
        return pulumi.get(self, "rolling_period")

    @property
    @pulumi.getter(name="serviceLevelIndicator")
    def service_level_indicator(self) -> 'outputs.ServiceLevelIndicatorResponse':
        """
        The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
        """
        return pulumi.get(self, "service_level_indicator")

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> Mapping[str, str]:
        """
        Labels which have been used to annotate the service-level objective. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
        """
        return pulumi.get(self, "user_labels")


class AwaitableGetServiceLevelObjectiveResult(GetServiceLevelObjectiveResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceLevelObjectiveResult(
            calendar_period=self.calendar_period,
            display_name=self.display_name,
            goal=self.goal,
            name=self.name,
            rolling_period=self.rolling_period,
            service_level_indicator=self.service_level_indicator,
            user_labels=self.user_labels)


def get_service_level_objective(service_id: Optional[str] = None,
                                service_level_objective_id: Optional[str] = None,
                                v3_id: Optional[str] = None,
                                v3_id1: Optional[str] = None,
                                view: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceLevelObjectiveResult:
    """
    Get a ServiceLevelObjective by name.
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    __args__['serviceLevelObjectiveId'] = service_level_objective_id
    __args__['v3Id'] = v3_id
    __args__['v3Id1'] = v3_id1
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:monitoring/v3:getServiceLevelObjective', __args__, opts=opts, typ=GetServiceLevelObjectiveResult).value

    return AwaitableGetServiceLevelObjectiveResult(
        calendar_period=pulumi.get(__ret__, 'calendar_period'),
        display_name=pulumi.get(__ret__, 'display_name'),
        goal=pulumi.get(__ret__, 'goal'),
        name=pulumi.get(__ret__, 'name'),
        rolling_period=pulumi.get(__ret__, 'rolling_period'),
        service_level_indicator=pulumi.get(__ret__, 'service_level_indicator'),
        user_labels=pulumi.get(__ret__, 'user_labels'))


@_utilities.lift_output_func(get_service_level_objective)
def get_service_level_objective_output(service_id: Optional[pulumi.Input[str]] = None,
                                       service_level_objective_id: Optional[pulumi.Input[str]] = None,
                                       v3_id: Optional[pulumi.Input[str]] = None,
                                       v3_id1: Optional[pulumi.Input[str]] = None,
                                       view: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceLevelObjectiveResult]:
    """
    Get a ServiceLevelObjective by name.
    """
    ...
