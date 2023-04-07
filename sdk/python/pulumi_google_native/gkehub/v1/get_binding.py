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
    'GetBindingResult',
    'AwaitableGetBindingResult',
    'get_binding',
    'get_binding_output',
]

@pulumi.output_type
class GetBindingResult:
    def __init__(__self__, create_time=None, delete_time=None, fleet=None, name=None, scope=None, state=None, uid=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if fleet and not isinstance(fleet, bool):
            raise TypeError("Expected argument 'fleet' to be a bool")
        pulumi.set(__self__, "fleet", fleet)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
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
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        When the membership binding was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        When the membership binding was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def fleet(self) -> bool:
        """
        Whether the membershipbinding is Fleet-wide; true means that this Membership should be bound to all Namespaces in this entire Fleet.
        """
        return pulumi.get(self, "fleet")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the membershipbinding itself `projects/{project}/locations/{location}/memberships/{membership}/bindings/{membershipbinding}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scope(self) -> str:
        """
        A Workspace resource name in the format `projects/*/locations/*/scopes/*`.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def state(self) -> 'outputs.MembershipBindingLifecycleStateResponse':
        """
        State of the membership binding resource.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Google-generated UUID for this resource. This is unique across all membershipbinding resources. If a membershipbinding resource is deleted and another resource with the same name is created, it gets a different uid.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        When the membership binding was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetBindingResult(GetBindingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBindingResult(
            create_time=self.create_time,
            delete_time=self.delete_time,
            fleet=self.fleet,
            name=self.name,
            scope=self.scope,
            state=self.state,
            uid=self.uid,
            update_time=self.update_time)


def get_binding(binding_id: Optional[str] = None,
                location: Optional[str] = None,
                membership_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBindingResult:
    """
    Returns the details of a MembershipBinding.
    """
    __args__ = dict()
    __args__['bindingId'] = binding_id
    __args__['location'] = location
    __args__['membershipId'] = membership_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:gkehub/v1:getBinding', __args__, opts=opts, typ=GetBindingResult).value

    return AwaitableGetBindingResult(
        create_time=__ret__.create_time,
        delete_time=__ret__.delete_time,
        fleet=__ret__.fleet,
        name=__ret__.name,
        scope=__ret__.scope,
        state=__ret__.state,
        uid=__ret__.uid,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_binding)
def get_binding_output(binding_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       membership_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBindingResult]:
    """
    Returns the details of a MembershipBinding.
    """
    ...
