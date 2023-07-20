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
    'GetDomainMappingResult',
    'AwaitableGetDomainMappingResult',
    'get_domain_mapping',
    'get_domain_mapping_output',
]

@pulumi.output_type
class GetDomainMappingResult:
    def __init__(__self__, name=None, resource_records=None, ssl_settings=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_records and not isinstance(resource_records, list):
            raise TypeError("Expected argument 'resource_records' to be a list")
        pulumi.set(__self__, "resource_records", resource_records)
        if ssl_settings and not isinstance(ssl_settings, dict):
            raise TypeError("Expected argument 'ssl_settings' to be a dict")
        pulumi.set(__self__, "ssl_settings", ssl_settings)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceRecords")
    def resource_records(self) -> Sequence['outputs.ResourceRecordResponse']:
        """
        The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.
        """
        return pulumi.get(self, "resource_records")

    @property
    @pulumi.getter(name="sslSettings")
    def ssl_settings(self) -> 'outputs.SslSettingsResponse':
        """
        SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
        """
        return pulumi.get(self, "ssl_settings")


class AwaitableGetDomainMappingResult(GetDomainMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainMappingResult(
            name=self.name,
            resource_records=self.resource_records,
            ssl_settings=self.ssl_settings)


def get_domain_mapping(app_id: Optional[str] = None,
                       domain_mapping_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainMappingResult:
    """
    Gets the specified domain mapping.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['domainMappingId'] = domain_mapping_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:appengine/v1alpha:getDomainMapping', __args__, opts=opts, typ=GetDomainMappingResult).value

    return AwaitableGetDomainMappingResult(
        name=pulumi.get(__ret__, 'name'),
        resource_records=pulumi.get(__ret__, 'resource_records'),
        ssl_settings=pulumi.get(__ret__, 'ssl_settings'))


@_utilities.lift_output_func(get_domain_mapping)
def get_domain_mapping_output(app_id: Optional[pulumi.Input[str]] = None,
                              domain_mapping_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainMappingResult]:
    """
    Gets the specified domain mapping.
    """
    ...
