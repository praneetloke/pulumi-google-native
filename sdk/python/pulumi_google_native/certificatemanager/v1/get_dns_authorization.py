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
    'GetDnsAuthorizationResult',
    'AwaitableGetDnsAuthorizationResult',
    'get_dns_authorization',
    'get_dns_authorization_output',
]

@pulumi.output_type
class GetDnsAuthorizationResult:
    def __init__(__self__, create_time=None, description=None, dns_resource_record=None, domain=None, labels=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dns_resource_record and not isinstance(dns_resource_record, dict):
            raise TypeError("Expected argument 'dns_resource_record' to be a dict")
        pulumi.set(__self__, "dns_resource_record", dns_resource_record)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
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
        The creation timestamp of a DnsAuthorization.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        One or more paragraphs of text description of a DnsAuthorization.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsResourceRecord")
    def dns_resource_record(self) -> 'outputs.DnsResourceRecordResponse':
        """
        DNS Resource Record that needs to be added to DNS configuration.
        """
        return pulumi.get(self, "dns_resource_record")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Immutable. A domain that is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Set of labels associated with a DnsAuthorization.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*/locations/*/dnsAuthorizations/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of a DnsAuthorization.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDnsAuthorizationResult(GetDnsAuthorizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDnsAuthorizationResult(
            create_time=self.create_time,
            description=self.description,
            dns_resource_record=self.dns_resource_record,
            domain=self.domain,
            labels=self.labels,
            name=self.name,
            update_time=self.update_time)


def get_dns_authorization(dns_authorization_id: Optional[str] = None,
                          location: Optional[str] = None,
                          project: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDnsAuthorizationResult:
    """
    Gets details of a single DnsAuthorization.
    """
    __args__ = dict()
    __args__['dnsAuthorizationId'] = dns_authorization_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:certificatemanager/v1:getDnsAuthorization', __args__, opts=opts, typ=GetDnsAuthorizationResult).value

    return AwaitableGetDnsAuthorizationResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        dns_resource_record=pulumi.get(__ret__, 'dns_resource_record'),
        domain=pulumi.get(__ret__, 'domain'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_dns_authorization)
def get_dns_authorization_output(dns_authorization_id: Optional[pulumi.Input[str]] = None,
                                 location: Optional[pulumi.Input[str]] = None,
                                 project: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDnsAuthorizationResult]:
    """
    Gets details of a single DnsAuthorization.
    """
    ...
