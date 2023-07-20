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
    'GetOrganizationExclusionResult',
    'AwaitableGetOrganizationExclusionResult',
    'get_organization_exclusion',
    'get_organization_exclusion_output',
]

@pulumi.output_type
class GetOrganizationExclusionResult:
    def __init__(__self__, create_time=None, description=None, disabled=None, filter=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of the exclusion.This field may not be present for older exclusions.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A description of this exclusion.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of the exclusion.This field may not be present for older exclusions.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetOrganizationExclusionResult(GetOrganizationExclusionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationExclusionResult(
            create_time=self.create_time,
            description=self.description,
            disabled=self.disabled,
            filter=self.filter,
            name=self.name,
            update_time=self.update_time)


def get_organization_exclusion(exclusion_id: Optional[str] = None,
                               organization_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationExclusionResult:
    """
    Gets the description of an exclusion in the _Default sink.
    """
    __args__ = dict()
    __args__['exclusionId'] = exclusion_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:logging/v2:getOrganizationExclusion', __args__, opts=opts, typ=GetOrganizationExclusionResult).value

    return AwaitableGetOrganizationExclusionResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        disabled=pulumi.get(__ret__, 'disabled'),
        filter=pulumi.get(__ret__, 'filter'),
        name=pulumi.get(__ret__, 'name'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_organization_exclusion)
def get_organization_exclusion_output(exclusion_id: Optional[pulumi.Input[str]] = None,
                                      organization_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationExclusionResult]:
    """
    Gets the description of an exclusion in the _Default sink.
    """
    ...
