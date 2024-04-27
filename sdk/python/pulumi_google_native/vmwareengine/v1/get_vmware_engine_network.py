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
    'GetVmwareEngineNetworkResult',
    'AwaitableGetVmwareEngineNetworkResult',
    'get_vmware_engine_network',
    'get_vmware_engine_network_output',
]

@pulumi.output_type
class GetVmwareEngineNetworkResult:
    def __init__(__self__, create_time=None, description=None, etag=None, name=None, state=None, type=None, uid=None, update_time=None, vpc_networks=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if vpc_networks and not isinstance(vpc_networks, list):
            raise TypeError("Expected argument 'vpc_networks' to be a list")
        pulumi.set(__self__, "vpc_networks", vpc_networks)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of this resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        User-provided description for this VMware Engine network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Checksum that may be sent on update and delete requests to ensure that the user-provided value is up to date before the server processes a request. The server computes checksums based on the value of other fields in the request.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the VMware Engine network. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/global/vmwareEngineNetworks/my-network`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the VMware Engine network.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        VMware Engine network type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System-generated unique identifier for the resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Last update time of this resource.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcNetworks")
    def vpc_networks(self) -> Sequence['outputs.VpcNetworkResponse']:
        """
        VMware Engine service VPC networks that provide connectivity from a private cloud to customer projects, the internet, and other Google Cloud services.
        """
        return pulumi.get(self, "vpc_networks")


class AwaitableGetVmwareEngineNetworkResult(GetVmwareEngineNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVmwareEngineNetworkResult(
            create_time=self.create_time,
            description=self.description,
            etag=self.etag,
            name=self.name,
            state=self.state,
            type=self.type,
            uid=self.uid,
            update_time=self.update_time,
            vpc_networks=self.vpc_networks)


def get_vmware_engine_network(location: Optional[str] = None,
                              project: Optional[str] = None,
                              vmware_engine_network_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVmwareEngineNetworkResult:
    """
    Retrieves a `VmwareEngineNetwork` resource by its resource name. The resource contains details of the VMware Engine network, such as its VMware Engine network type, peered networks in a service project, and state (for example, `CREATING`, `ACTIVE`, `DELETING`).
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['vmwareEngineNetworkId'] = vmware_engine_network_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:vmwareengine/v1:getVmwareEngineNetwork', __args__, opts=opts, typ=GetVmwareEngineNetworkResult).value

    return AwaitableGetVmwareEngineNetworkResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'),
        vpc_networks=pulumi.get(__ret__, 'vpc_networks'))


@_utilities.lift_output_func(get_vmware_engine_network)
def get_vmware_engine_network_output(location: Optional[pulumi.Input[str]] = None,
                                     project: Optional[pulumi.Input[Optional[str]]] = None,
                                     vmware_engine_network_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVmwareEngineNetworkResult]:
    """
    Retrieves a `VmwareEngineNetwork` resource by its resource name. The resource contains details of the VMware Engine network, such as its VMware Engine network type, peered networks in a service project, and state (for example, `CREATING`, `ACTIVE`, `DELETING`).
    """
    ...