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
    'GetNetworkAttachmentResult',
    'AwaitableGetNetworkAttachmentResult',
    'get_network_attachment',
    'get_network_attachment_output',
]

@pulumi.output_type
class GetNetworkAttachmentResult:
    def __init__(__self__, connection_endpoints=None, connection_preference=None, creation_timestamp=None, description=None, fingerprint=None, kind=None, name=None, network=None, producer_accept_lists=None, producer_reject_lists=None, region=None, self_link=None, self_link_with_id=None, subnetworks=None):
        if connection_endpoints and not isinstance(connection_endpoints, list):
            raise TypeError("Expected argument 'connection_endpoints' to be a list")
        pulumi.set(__self__, "connection_endpoints", connection_endpoints)
        if connection_preference and not isinstance(connection_preference, str):
            raise TypeError("Expected argument 'connection_preference' to be a str")
        pulumi.set(__self__, "connection_preference", connection_preference)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if producer_accept_lists and not isinstance(producer_accept_lists, list):
            raise TypeError("Expected argument 'producer_accept_lists' to be a list")
        pulumi.set(__self__, "producer_accept_lists", producer_accept_lists)
        if producer_reject_lists and not isinstance(producer_reject_lists, list):
            raise TypeError("Expected argument 'producer_reject_lists' to be a list")
        pulumi.set(__self__, "producer_reject_lists", producer_reject_lists)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if subnetworks and not isinstance(subnetworks, list):
            raise TypeError("Expected argument 'subnetworks' to be a list")
        pulumi.set(__self__, "subnetworks", subnetworks)

    @property
    @pulumi.getter(name="connectionEndpoints")
    def connection_endpoints(self) -> Sequence['outputs.NetworkAttachmentConnectedEndpointResponse']:
        """
        An array of connections for all the producers connected to this network attachment.
        """
        return pulumi.get(self, "connection_endpoints")

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> str:
        return pulumi.get(self, "connection_preference")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="producerAcceptLists")
    def producer_accept_lists(self) -> Sequence[str]:
        """
        Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
        """
        return pulumi.get(self, "producer_accept_lists")

    @property
    @pulumi.getter(name="producerRejectLists")
    def producer_reject_lists(self) -> Sequence[str]:
        """
        Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
        """
        return pulumi.get(self, "producer_reject_lists")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource's resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def subnetworks(self) -> Sequence[str]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
        """
        return pulumi.get(self, "subnetworks")


class AwaitableGetNetworkAttachmentResult(GetNetworkAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkAttachmentResult(
            connection_endpoints=self.connection_endpoints,
            connection_preference=self.connection_preference,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            fingerprint=self.fingerprint,
            kind=self.kind,
            name=self.name,
            network=self.network,
            producer_accept_lists=self.producer_accept_lists,
            producer_reject_lists=self.producer_reject_lists,
            region=self.region,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            subnetworks=self.subnetworks)


def get_network_attachment(network_attachment: Optional[str] = None,
                           project: Optional[str] = None,
                           region: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkAttachmentResult:
    """
    Returns the specified NetworkAttachment resource in the given scope.
    """
    __args__ = dict()
    __args__['networkAttachment'] = network_attachment
    __args__['project'] = project
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getNetworkAttachment', __args__, opts=opts, typ=GetNetworkAttachmentResult).value

    return AwaitableGetNetworkAttachmentResult(
        connection_endpoints=__ret__.connection_endpoints,
        connection_preference=__ret__.connection_preference,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        fingerprint=__ret__.fingerprint,
        kind=__ret__.kind,
        name=__ret__.name,
        network=__ret__.network,
        producer_accept_lists=__ret__.producer_accept_lists,
        producer_reject_lists=__ret__.producer_reject_lists,
        region=__ret__.region,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        subnetworks=__ret__.subnetworks)


@_utilities.lift_output_func(get_network_attachment)
def get_network_attachment_output(network_attachment: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkAttachmentResult]:
    """
    Returns the specified NetworkAttachment resource in the given scope.
    """
    ...
