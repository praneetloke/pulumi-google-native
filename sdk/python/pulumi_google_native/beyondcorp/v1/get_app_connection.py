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
    'GetAppConnectionResult',
    'AwaitableGetAppConnectionResult',
    'get_app_connection',
    'get_app_connection_output',
]

@pulumi.output_type
class GetAppConnectionResult:
    def __init__(__self__, application_endpoint=None, connectors=None, create_time=None, display_name=None, gateway=None, labels=None, name=None, state=None, type=None, uid=None, update_time=None):
        if application_endpoint and not isinstance(application_endpoint, dict):
            raise TypeError("Expected argument 'application_endpoint' to be a dict")
        pulumi.set(__self__, "application_endpoint", application_endpoint)
        if connectors and not isinstance(connectors, list):
            raise TypeError("Expected argument 'connectors' to be a list")
        pulumi.set(__self__, "connectors", connectors)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if gateway and not isinstance(gateway, dict):
            raise TypeError("Expected argument 'gateway' to be a dict")
        pulumi.set(__self__, "gateway", gateway)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
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

    @property
    @pulumi.getter(name="applicationEndpoint")
    def application_endpoint(self) -> 'outputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionApplicationEndpointResponse':
        """
        Address of the remote application endpoint for the BeyondCorp AppConnection.
        """
        return pulumi.get(self, "application_endpoint")

    @property
    @pulumi.getter
    def connectors(self) -> Sequence[str]:
        """
        Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
        """
        return pulumi.get(self, "connectors")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def gateway(self) -> 'outputs.GoogleCloudBeyondcorpAppconnectionsV1AppConnectionGatewayResponse':
        """
        Optional. Gateway used by the AppConnection.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the AppConnection.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of network connectivity used by the AppConnection.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        A unique identifier for the instance generated by the system.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Timestamp when the resource was last modified.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetAppConnectionResult(GetAppConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppConnectionResult(
            application_endpoint=self.application_endpoint,
            connectors=self.connectors,
            create_time=self.create_time,
            display_name=self.display_name,
            gateway=self.gateway,
            labels=self.labels,
            name=self.name,
            state=self.state,
            type=self.type,
            uid=self.uid,
            update_time=self.update_time)


def get_app_connection(app_connection_id: Optional[str] = None,
                       location: Optional[str] = None,
                       project: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppConnectionResult:
    """
    Gets details of a single AppConnection.
    """
    __args__ = dict()
    __args__['appConnectionId'] = app_connection_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:beyondcorp/v1:getAppConnection', __args__, opts=opts, typ=GetAppConnectionResult).value

    return AwaitableGetAppConnectionResult(
        application_endpoint=pulumi.get(__ret__, 'application_endpoint'),
        connectors=pulumi.get(__ret__, 'connectors'),
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        gateway=pulumi.get(__ret__, 'gateway'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_app_connection)
def get_app_connection_output(app_connection_id: Optional[pulumi.Input[str]] = None,
                              location: Optional[pulumi.Input[str]] = None,
                              project: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppConnectionResult]:
    """
    Gets details of a single AppConnection.
    """
    ...
