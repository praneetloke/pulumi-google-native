# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetAgentPoolResult',
    'AwaitableGetAgentPoolResult',
    'get_agent_pool',
    'get_agent_pool_output',
]

@pulumi.output_type
class GetAgentPoolResult:
    def __init__(__self__, bandwidth_limit=None, display_name=None, name=None, state=None):
        if bandwidth_limit and not isinstance(bandwidth_limit, dict):
            raise TypeError("Expected argument 'bandwidth_limit' to be a dict")
        pulumi.set(__self__, "bandwidth_limit", bandwidth_limit)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="bandwidthLimit")
    def bandwidth_limit(self) -> 'outputs.BandwidthLimitResponse':
        """
        Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
        """
        return pulumi.get(self, "bandwidth_limit")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Specifies the client-specified AgentPool description.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies a unique string that identifies the agent pool. Format: projects/{project_id}/agentPools/{agent_pool_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Specifies the state of the AgentPool.
        """
        return pulumi.get(self, "state")


class AwaitableGetAgentPoolResult(GetAgentPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentPoolResult(
            bandwidth_limit=self.bandwidth_limit,
            display_name=self.display_name,
            name=self.name,
            state=self.state)


def get_agent_pool(agent_pool_id: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentPoolResult:
    """
    Gets an agent pool.
    """
    __args__ = dict()
    __args__['agentPoolId'] = agent_pool_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:storagetransfer/v1:getAgentPool', __args__, opts=opts, typ=GetAgentPoolResult).value

    return AwaitableGetAgentPoolResult(
        bandwidth_limit=__ret__.bandwidth_limit,
        display_name=__ret__.display_name,
        name=__ret__.name,
        state=__ret__.state)


@_utilities.lift_output_func(get_agent_pool)
def get_agent_pool_output(agent_pool_id: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAgentPoolResult]:
    """
    Gets an agent pool.
    """
    ...