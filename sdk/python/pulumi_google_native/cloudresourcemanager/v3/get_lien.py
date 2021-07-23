# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetLienResult',
    'AwaitableGetLienResult',
    'get_lien',
]

@pulumi.output_type
class GetLienResult:
    def __init__(__self__, create_time=None, name=None, origin=None, parent=None, reason=None, restrictions=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        pulumi.set(__self__, "parent", parent)
        if reason and not isinstance(reason, str):
            raise TypeError("Expected argument 'reason' to be a str")
        pulumi.set(__self__, "reason", reason)
        if restrictions and not isinstance(restrictions, list):
            raise TypeError("Expected argument 'restrictions' to be a list")
        pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of this Lien.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def origin(self) -> str:
        """
        A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter
    def parent(self) -> str:
        """
        A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def reason(self) -> str:
        """
        Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
        """
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter
    def restrictions(self) -> Sequence[str]:
        """
        The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
        """
        return pulumi.get(self, "restrictions")


class AwaitableGetLienResult(GetLienResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLienResult(
            create_time=self.create_time,
            name=self.name,
            origin=self.origin,
            parent=self.parent,
            reason=self.reason,
            restrictions=self.restrictions)


def get_lien(lien_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLienResult:
    """
    Retrieve a Lien by `name`. Callers of this method will require permission on the `parent` resource. For example, a Lien with a `parent` of `projects/1234` requires permission `resourcemanager.projects.get`
    """
    __args__ = dict()
    __args__['lienId'] = lien_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:cloudresourcemanager/v3:getLien', __args__, opts=opts, typ=GetLienResult).value

    return AwaitableGetLienResult(
        create_time=__ret__.create_time,
        name=__ret__.name,
        origin=__ret__.origin,
        parent=__ret__.parent,
        reason=__ret__.reason,
        restrictions=__ret__.restrictions)