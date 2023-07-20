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
    'GetPrivateConnectionResult',
    'AwaitableGetPrivateConnectionResult',
    'get_private_connection',
    'get_private_connection_output',
]

@pulumi.output_type
class GetPrivateConnectionResult:
    def __init__(__self__, create_time=None, display_name=None, error=None, labels=None, name=None, state=None, update_time=None, vpc_peering_config=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if vpc_peering_config and not isinstance(vpc_peering_config, dict):
            raise TypeError("Expected argument 'vpc_peering_config' to be a dict")
        pulumi.set(__self__, "vpc_peering_config", vpc_peering_config)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The private connection display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.StatusResponse':
        """
        The error details in case of state FAILED.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The resource labels for private connections to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the private connection.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update time of the resource.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcPeeringConfig")
    def vpc_peering_config(self) -> 'outputs.VpcPeeringConfigResponse':
        """
        VPC peering configuration.
        """
        return pulumi.get(self, "vpc_peering_config")


class AwaitableGetPrivateConnectionResult(GetPrivateConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateConnectionResult(
            create_time=self.create_time,
            display_name=self.display_name,
            error=self.error,
            labels=self.labels,
            name=self.name,
            state=self.state,
            update_time=self.update_time,
            vpc_peering_config=self.vpc_peering_config)


def get_private_connection(location: Optional[str] = None,
                           private_connection_id: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateConnectionResult:
    """
    Gets details of a single private connection.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['privateConnectionId'] = private_connection_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datamigration/v1:getPrivateConnection', __args__, opts=opts, typ=GetPrivateConnectionResult).value

    return AwaitableGetPrivateConnectionResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        error=pulumi.get(__ret__, 'error'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'),
        vpc_peering_config=pulumi.get(__ret__, 'vpc_peering_config'))


@_utilities.lift_output_func(get_private_connection)
def get_private_connection_output(location: Optional[pulumi.Input[str]] = None,
                                  private_connection_id: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateConnectionResult]:
    """
    Gets details of a single private connection.
    """
    ...
