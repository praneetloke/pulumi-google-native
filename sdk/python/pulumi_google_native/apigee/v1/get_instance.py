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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, consumer_accept_list=None, created_at=None, description=None, disk_encryption_key_name=None, display_name=None, host=None, ip_range=None, last_modified_at=None, location=None, name=None, peering_cidr_range=None, port=None, runtime_version=None, service_attachment=None, state=None):
        if consumer_accept_list and not isinstance(consumer_accept_list, list):
            raise TypeError("Expected argument 'consumer_accept_list' to be a list")
        pulumi.set(__self__, "consumer_accept_list", consumer_accept_list)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_encryption_key_name and not isinstance(disk_encryption_key_name, str):
            raise TypeError("Expected argument 'disk_encryption_key_name' to be a str")
        pulumi.set(__self__, "disk_encryption_key_name", disk_encryption_key_name)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if ip_range and not isinstance(ip_range, str):
            raise TypeError("Expected argument 'ip_range' to be a str")
        pulumi.set(__self__, "ip_range", ip_range)
        if last_modified_at and not isinstance(last_modified_at, str):
            raise TypeError("Expected argument 'last_modified_at' to be a str")
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peering_cidr_range and not isinstance(peering_cidr_range, str):
            raise TypeError("Expected argument 'peering_cidr_range' to be a str")
        pulumi.set(__self__, "peering_cidr_range", peering_cidr_range)
        if port and not isinstance(port, str):
            raise TypeError("Expected argument 'port' to be a str")
        pulumi.set(__self__, "port", port)
        if runtime_version and not isinstance(runtime_version, str):
            raise TypeError("Expected argument 'runtime_version' to be a str")
        pulumi.set(__self__, "runtime_version", runtime_version)
        if service_attachment and not isinstance(service_attachment, str):
            raise TypeError("Expected argument 'service_attachment' to be a str")
        pulumi.set(__self__, "service_attachment", service_attachment)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="consumerAcceptList")
    def consumer_accept_list(self) -> Sequence[str]:
        """
        Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
        """
        return pulumi.get(self, "consumer_accept_list")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Time the instance was created in milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> str:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. Display name for the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> str:
        """
        Optional. Comma-separated list of CIDR blocks of length 22 and/or 28 used to create the Apigee instance. Providing CIDR ranges is optional. You can provide just /22 or /28 or both (or neither). Ranges you provide should be freely available as part of a larger named range you have allocated to the Service Networking peering. If this parameter is not provided, Apigee automatically requests an available /22 and /28 CIDR block from Service Networking. Use the /22 CIDR block for configuring your firewall needs to allow traffic from Apigee. Input formats: `a.b.c.d/22` or `e.f.g.h/28` or `a.b.c.d/22,e.f.g.h/28`
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        Time the instance was last modified in milliseconds since epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Compute Engine location where the instance resides.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\\d]$`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> str:
        """
        Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @property
    @pulumi.getter
    def port(self) -> str:
        """
        Port number of the exposed Apigee endpoint.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> str:
        """
        Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> str:
        """
        Resource name of the service attachment created for the instance in the format: `projects/*/regions/*/serviceAttachments/*` Apigee customers can privately forward traffic to this service attachment using the PSC endpoints.
        """
        return pulumi.get(self, "service_attachment")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
        """
        return pulumi.get(self, "state")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            consumer_accept_list=self.consumer_accept_list,
            created_at=self.created_at,
            description=self.description,
            disk_encryption_key_name=self.disk_encryption_key_name,
            display_name=self.display_name,
            host=self.host,
            ip_range=self.ip_range,
            last_modified_at=self.last_modified_at,
            location=self.location,
            name=self.name,
            peering_cidr_range=self.peering_cidr_range,
            port=self.port,
            runtime_version=self.runtime_version,
            service_attachment=self.service_attachment,
            state=self.state)


def get_instance(instance_id: Optional[str] = None,
                 organization_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets the details for an Apigee runtime instance. **Note:** Not supported for Apigee hybrid.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        consumer_accept_list=pulumi.get(__ret__, 'consumer_accept_list'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        disk_encryption_key_name=pulumi.get(__ret__, 'disk_encryption_key_name'),
        display_name=pulumi.get(__ret__, 'display_name'),
        host=pulumi.get(__ret__, 'host'),
        ip_range=pulumi.get(__ret__, 'ip_range'),
        last_modified_at=pulumi.get(__ret__, 'last_modified_at'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        peering_cidr_range=pulumi.get(__ret__, 'peering_cidr_range'),
        port=pulumi.get(__ret__, 'port'),
        runtime_version=pulumi.get(__ret__, 'runtime_version'),
        service_attachment=pulumi.get(__ret__, 'service_attachment'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        organization_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets the details for an Apigee runtime instance. **Note:** Not supported for Apigee hybrid.
    """
    ...
