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
    'GetTlsInspectionPolicyResult',
    'AwaitableGetTlsInspectionPolicyResult',
    'get_tls_inspection_policy',
    'get_tls_inspection_policy_output',
]

@pulumi.output_type
class GetTlsInspectionPolicyResult:
    def __init__(__self__, ca_pool=None, create_time=None, description=None, name=None, update_time=None):
        if ca_pool and not isinstance(ca_pool, str):
            raise TypeError("Expected argument 'ca_pool' to be a str")
        pulumi.set(__self__, "ca_pool", ca_pool)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="caPool")
    def ca_pool(self) -> str:
        """
        A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
        """
        return pulumi.get(self, "ca_pool")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp when the resource was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetTlsInspectionPolicyResult(GetTlsInspectionPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTlsInspectionPolicyResult(
            ca_pool=self.ca_pool,
            create_time=self.create_time,
            description=self.description,
            name=self.name,
            update_time=self.update_time)


def get_tls_inspection_policy(location: Optional[str] = None,
                              project: Optional[str] = None,
                              tls_inspection_policy_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTlsInspectionPolicyResult:
    """
    Gets details of a single TlsInspectionPolicy.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['tlsInspectionPolicyId'] = tls_inspection_policy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:networksecurity/v1:getTlsInspectionPolicy', __args__, opts=opts, typ=GetTlsInspectionPolicyResult).value

    return AwaitableGetTlsInspectionPolicyResult(
        ca_pool=pulumi.get(__ret__, 'ca_pool'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        name=pulumi.get(__ret__, 'name'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_tls_inspection_policy)
def get_tls_inspection_policy_output(location: Optional[pulumi.Input[str]] = None,
                                     project: Optional[pulumi.Input[Optional[str]]] = None,
                                     tls_inspection_policy_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTlsInspectionPolicyResult]:
    """
    Gets details of a single TlsInspectionPolicy.
    """
    ...
