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
    'GetKeyRingResult',
    'AwaitableGetKeyRingResult',
    'get_key_ring',
    'get_key_ring_output',
]

@pulumi.output_type
class GetKeyRingResult:
    def __init__(__self__, create_time=None, name=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time at which this KeyRing was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the KeyRing in the format `projects/*/locations/*/keyRings/*`.
        """
        return pulumi.get(self, "name")


class AwaitableGetKeyRingResult(GetKeyRingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyRingResult(
            create_time=self.create_time,
            name=self.name)


def get_key_ring(key_ring_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyRingResult:
    """
    Returns metadata for a given KeyRing.
    """
    __args__ = dict()
    __args__['keyRingId'] = key_ring_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudkms/v1:getKeyRing', __args__, opts=opts, typ=GetKeyRingResult).value

    return AwaitableGetKeyRingResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_key_ring)
def get_key_ring_output(key_ring_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyRingResult]:
    """
    Returns metadata for a given KeyRing.
    """
    ...
