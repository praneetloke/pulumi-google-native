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
    'GetScopeResult',
    'AwaitableGetScopeResult',
    'get_scope',
    'get_scope_output',
]

@pulumi.output_type
class GetScopeResult:
    def __init__(__self__, all_memberships=None, create_time=None, delete_time=None, name=None, state=None, uid=None, update_time=None):
        if all_memberships and not isinstance(all_memberships, bool):
            raise TypeError("Expected argument 'all_memberships' to be a bool")
        pulumi.set(__self__, "all_memberships", all_memberships)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, dict):
            raise TypeError("Expected argument 'state' to be a dict")
        pulumi.set(__self__, "state", state)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="allMemberships")
    def all_memberships(self) -> bool:
        """
        If true, all Memberships in the Fleet bind to this Scope.
        """
        return pulumi.get(self, "all_memberships")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        When the scope was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        When the scope was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the scope `projects/{project}/locations/{location}/scopes/{scope}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> 'outputs.ScopeLifecycleStateResponse':
        """
        State of the scope resource.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Google-generated UUID for this resource. This is unique across all scope resources. If a scope resource is deleted and another resource with the same name is created, it gets a different uid.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        When the scope was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetScopeResult(GetScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScopeResult(
            all_memberships=self.all_memberships,
            create_time=self.create_time,
            delete_time=self.delete_time,
            name=self.name,
            state=self.state,
            uid=self.uid,
            update_time=self.update_time)


def get_scope(location: Optional[str] = None,
              project: Optional[str] = None,
              scope_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScopeResult:
    """
    Returns the details of a Scope.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['scopeId'] = scope_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:gkehub/v1alpha:getScope', __args__, opts=opts, typ=GetScopeResult).value

    return AwaitableGetScopeResult(
        all_memberships=pulumi.get(__ret__, 'all_memberships'),
        create_time=pulumi.get(__ret__, 'create_time'),
        delete_time=pulumi.get(__ret__, 'delete_time'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_scope)
def get_scope_output(location: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     scope_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScopeResult]:
    """
    Returns the details of a Scope.
    """
    ...
