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
    'GetGlobalNetworkEndpointGroupResult',
    'AwaitableGetGlobalNetworkEndpointGroupResult',
    'get_global_network_endpoint_group',
    'get_global_network_endpoint_group_output',
]

@pulumi.output_type
class GetGlobalNetworkEndpointGroupResult:
    def __init__(__self__, annotations=None, app_engine=None, cloud_function=None, cloud_run=None, creation_timestamp=None, default_port=None, description=None, kind=None, name=None, network=None, network_endpoint_type=None, psc_data=None, psc_target_service=None, region=None, self_link=None, size=None, subnetwork=None, zone=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        pulumi.set(__self__, "annotations", annotations)
        if app_engine and not isinstance(app_engine, dict):
            raise TypeError("Expected argument 'app_engine' to be a dict")
        pulumi.set(__self__, "app_engine", app_engine)
        if cloud_function and not isinstance(cloud_function, dict):
            raise TypeError("Expected argument 'cloud_function' to be a dict")
        pulumi.set(__self__, "cloud_function", cloud_function)
        if cloud_run and not isinstance(cloud_run, dict):
            raise TypeError("Expected argument 'cloud_run' to be a dict")
        pulumi.set(__self__, "cloud_run", cloud_run)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if default_port and not isinstance(default_port, int):
            raise TypeError("Expected argument 'default_port' to be a int")
        pulumi.set(__self__, "default_port", default_port)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if network_endpoint_type and not isinstance(network_endpoint_type, str):
            raise TypeError("Expected argument 'network_endpoint_type' to be a str")
        pulumi.set(__self__, "network_endpoint_type", network_endpoint_type)
        if psc_data and not isinstance(psc_data, dict):
            raise TypeError("Expected argument 'psc_data' to be a dict")
        pulumi.set(__self__, "psc_data", psc_data)
        if psc_target_service and not isinstance(psc_target_service, str):
            raise TypeError("Expected argument 'psc_target_service' to be a str")
        pulumi.set(__self__, "psc_target_service", psc_target_service)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if subnetwork and not isinstance(subnetwork, str):
            raise TypeError("Expected argument 'subnetwork' to be a str")
        pulumi.set(__self__, "subnetwork", subnetwork)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def annotations(self) -> Mapping[str, str]:
        """
        Metadata defined as annotations on the network endpoint group.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="appEngine")
    def app_engine(self) -> 'outputs.NetworkEndpointGroupAppEngineResponse':
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "app_engine")

    @property
    @pulumi.getter(name="cloudFunction")
    def cloud_function(self) -> 'outputs.NetworkEndpointGroupCloudFunctionResponse':
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_function")

    @property
    @pulumi.getter(name="cloudRun")
    def cloud_run(self) -> 'outputs.NetworkEndpointGroupCloudRunResponse':
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_run")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> int:
        """
        The default port used if the port number is not specified in the network endpoint.
        """
        return pulumi.get(self, "default_port")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> str:
        """
        Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        """
        return pulumi.get(self, "network_endpoint_type")

    @property
    @pulumi.getter(name="pscData")
    def psc_data(self) -> 'outputs.NetworkEndpointGroupPscDataResponse':
        return pulumi.get(self, "psc_data")

    @property
    @pulumi.getter(name="pscTargetService")
    def psc_target_service(self) -> str:
        """
        The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        """
        return pulumi.get(self, "psc_target_service")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The URL of the region where the network endpoint group is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        [Output only] Number of network endpoints in the network endpoint group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def subnetwork(self) -> str:
        """
        Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The URL of the zone where the network endpoint group is located.
        """
        return pulumi.get(self, "zone")


class AwaitableGetGlobalNetworkEndpointGroupResult(GetGlobalNetworkEndpointGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalNetworkEndpointGroupResult(
            annotations=self.annotations,
            app_engine=self.app_engine,
            cloud_function=self.cloud_function,
            cloud_run=self.cloud_run,
            creation_timestamp=self.creation_timestamp,
            default_port=self.default_port,
            description=self.description,
            kind=self.kind,
            name=self.name,
            network=self.network,
            network_endpoint_type=self.network_endpoint_type,
            psc_data=self.psc_data,
            psc_target_service=self.psc_target_service,
            region=self.region,
            self_link=self.self_link,
            size=self.size,
            subnetwork=self.subnetwork,
            zone=self.zone)


def get_global_network_endpoint_group(network_endpoint_group: Optional[str] = None,
                                      project: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlobalNetworkEndpointGroupResult:
    """
    Returns the specified network endpoint group.
    """
    __args__ = dict()
    __args__['networkEndpointGroup'] = network_endpoint_group
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getGlobalNetworkEndpointGroup', __args__, opts=opts, typ=GetGlobalNetworkEndpointGroupResult).value

    return AwaitableGetGlobalNetworkEndpointGroupResult(
        annotations=__ret__.annotations,
        app_engine=__ret__.app_engine,
        cloud_function=__ret__.cloud_function,
        cloud_run=__ret__.cloud_run,
        creation_timestamp=__ret__.creation_timestamp,
        default_port=__ret__.default_port,
        description=__ret__.description,
        kind=__ret__.kind,
        name=__ret__.name,
        network=__ret__.network,
        network_endpoint_type=__ret__.network_endpoint_type,
        psc_data=__ret__.psc_data,
        psc_target_service=__ret__.psc_target_service,
        region=__ret__.region,
        self_link=__ret__.self_link,
        size=__ret__.size,
        subnetwork=__ret__.subnetwork,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_global_network_endpoint_group)
def get_global_network_endpoint_group_output(network_endpoint_group: Optional[pulumi.Input[str]] = None,
                                             project: Optional[pulumi.Input[Optional[str]]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGlobalNetworkEndpointGroupResult]:
    """
    Returns the specified network endpoint group.
    """
    ...
