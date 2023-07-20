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
    'GetWorkloadIdentityPoolResult',
    'AwaitableGetWorkloadIdentityPoolResult',
    'get_workload_identity_pool',
    'get_workload_identity_pool_output',
]

@pulumi.output_type
class GetWorkloadIdentityPoolResult:
    def __init__(__self__, description=None, disabled=None, display_name=None, name=None, state=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
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
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the pool. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        A display name for the pool. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the pool.
        """
        return pulumi.get(self, "state")


class AwaitableGetWorkloadIdentityPoolResult(GetWorkloadIdentityPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkloadIdentityPoolResult(
            description=self.description,
            disabled=self.disabled,
            display_name=self.display_name,
            name=self.name,
            state=self.state)


def get_workload_identity_pool(location: Optional[str] = None,
                               project: Optional[str] = None,
                               workload_identity_pool_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkloadIdentityPoolResult:
    """
    Gets an individual WorkloadIdentityPool.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['workloadIdentityPoolId'] = workload_identity_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:iam/v1:getWorkloadIdentityPool', __args__, opts=opts, typ=GetWorkloadIdentityPoolResult).value

    return AwaitableGetWorkloadIdentityPoolResult(
        description=pulumi.get(__ret__, 'description'),
        disabled=pulumi.get(__ret__, 'disabled'),
        display_name=pulumi.get(__ret__, 'display_name'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_workload_identity_pool)
def get_workload_identity_pool_output(location: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkloadIdentityPoolResult]:
    """
    Gets an individual WorkloadIdentityPool.
    """
    ...
