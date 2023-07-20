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
    'GetClientConnectorServiceResult',
    'AwaitableGetClientConnectorServiceResult',
    'get_client_connector_service',
    'get_client_connector_service_output',
]

@pulumi.output_type
class GetClientConnectorServiceResult:
    def __init__(__self__, create_time=None, display_name=None, egress=None, ingress=None, name=None, state=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if egress and not isinstance(egress, dict):
            raise TypeError("Expected argument 'egress' to be a dict")
        pulumi.set(__self__, "egress", egress)
        if ingress and not isinstance(ingress, dict):
            raise TypeError("Expected argument 'ingress' to be a dict")
        pulumi.set(__self__, "ingress", ingress)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        [Output only] Create time stamp.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def egress(self) -> 'outputs.EgressResponse':
        """
        The details of the egress settings.
        """
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter
    def ingress(self) -> 'outputs.IngressResponse':
        """
        The details of the ingress settings.
        """
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of resource. The name is ignored during creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The operational state of the ClientConnectorService.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        [Output only] Update time stamp.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetClientConnectorServiceResult(GetClientConnectorServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientConnectorServiceResult(
            create_time=self.create_time,
            display_name=self.display_name,
            egress=self.egress,
            ingress=self.ingress,
            name=self.name,
            state=self.state,
            update_time=self.update_time)


def get_client_connector_service(client_connector_service_id: Optional[str] = None,
                                 location: Optional[str] = None,
                                 project: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientConnectorServiceResult:
    """
    Gets details of a single ClientConnectorService.
    """
    __args__ = dict()
    __args__['clientConnectorServiceId'] = client_connector_service_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:beyondcorp/v1alpha:getClientConnectorService', __args__, opts=opts, typ=GetClientConnectorServiceResult).value

    return AwaitableGetClientConnectorServiceResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        egress=pulumi.get(__ret__, 'egress'),
        ingress=pulumi.get(__ret__, 'ingress'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_client_connector_service)
def get_client_connector_service_output(client_connector_service_id: Optional[pulumi.Input[str]] = None,
                                        location: Optional[pulumi.Input[str]] = None,
                                        project: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientConnectorServiceResult]:
    """
    Gets details of a single ClientConnectorService.
    """
    ...
