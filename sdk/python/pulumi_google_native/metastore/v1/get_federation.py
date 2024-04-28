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
    'GetFederationResult',
    'AwaitableGetFederationResult',
    'get_federation',
    'get_federation_output',
]

@pulumi.output_type
class GetFederationResult:
    def __init__(__self__, backend_metastores=None, create_time=None, endpoint_uri=None, labels=None, name=None, state=None, state_message=None, uid=None, update_time=None, version=None):
        if backend_metastores and not isinstance(backend_metastores, dict):
            raise TypeError("Expected argument 'backend_metastores' to be a dict")
        pulumi.set(__self__, "backend_metastores", backend_metastores)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if endpoint_uri and not isinstance(endpoint_uri, str):
            raise TypeError("Expected argument 'endpoint_uri' to be a str")
        pulumi.set(__self__, "endpoint_uri", endpoint_uri)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_message and not isinstance(state_message, str):
            raise TypeError("Expected argument 'state_message' to be a str")
        pulumi.set(__self__, "state_message", state_message)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="backendMetastores")
    def backend_metastores(self) -> 'outputs.BackendMetastoreResponse':
        """
        A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
        """
        return pulumi.get(self, "backend_metastores")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the metastore federation was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="endpointUri")
    def endpoint_uri(self) -> str:
        """
        The federation endpoint.
        """
        return pulumi.get(self, "endpoint_uri")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        User-defined labels for the metastore federation.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The relative resource name of the federation, of the form: projects/{project_number}/locations/{location_id}/federations/{federation_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the federation.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateMessage")
    def state_message(self) -> str:
        """
        Additional information about the current state of the metastore federation, if available.
        """
        return pulumi.get(self, "state_message")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        The globally unique resource identifier of the metastore federation.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the metastore federation was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Immutable. The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
        """
        return pulumi.get(self, "version")


class AwaitableGetFederationResult(GetFederationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFederationResult(
            backend_metastores=self.backend_metastores,
            create_time=self.create_time,
            endpoint_uri=self.endpoint_uri,
            labels=self.labels,
            name=self.name,
            state=self.state,
            state_message=self.state_message,
            uid=self.uid,
            update_time=self.update_time,
            version=self.version)


def get_federation(federation_id: Optional[str] = None,
                   location: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFederationResult:
    """
    Gets the details of a single federation.
    """
    __args__ = dict()
    __args__['federationId'] = federation_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:metastore/v1:getFederation', __args__, opts=opts, typ=GetFederationResult).value

    return AwaitableGetFederationResult(
        backend_metastores=pulumi.get(__ret__, 'backend_metastores'),
        create_time=pulumi.get(__ret__, 'create_time'),
        endpoint_uri=pulumi.get(__ret__, 'endpoint_uri'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        state_message=pulumi.get(__ret__, 'state_message'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_federation)
def get_federation_output(federation_id: Optional[pulumi.Input[str]] = None,
                          location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFederationResult]:
    """
    Gets the details of a single federation.
    """
    ...
