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
    'GetNodeGroupResult',
    'AwaitableGetNodeGroupResult',
    'get_node_group',
    'get_node_group_output',
]

@pulumi.output_type
class GetNodeGroupResult:
    def __init__(__self__, autoscaling_policy=None, creation_timestamp=None, description=None, fingerprint=None, kind=None, location_hint=None, maintenance_policy=None, maintenance_window=None, name=None, node_template=None, self_link=None, share_settings=None, size=None, status=None, zone=None):
        if autoscaling_policy and not isinstance(autoscaling_policy, dict):
            raise TypeError("Expected argument 'autoscaling_policy' to be a dict")
        pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
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
        if location_hint and not isinstance(location_hint, str):
            raise TypeError("Expected argument 'location_hint' to be a str")
        pulumi.set(__self__, "location_hint", location_hint)
        if maintenance_policy and not isinstance(maintenance_policy, str):
            raise TypeError("Expected argument 'maintenance_policy' to be a str")
        pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if maintenance_window and not isinstance(maintenance_window, dict):
            raise TypeError("Expected argument 'maintenance_window' to be a dict")
        pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_template and not isinstance(node_template, str):
            raise TypeError("Expected argument 'node_template' to be a str")
        pulumi.set(__self__, "node_template", node_template)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if share_settings and not isinstance(share_settings, dict):
            raise TypeError("Expected argument 'share_settings' to be a dict")
        pulumi.set(__self__, "share_settings", share_settings)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> 'outputs.NodeGroupAutoscalingPolicyResponse':
        """
        Specifies how autoscaling should behave.
        """
        return pulumi.get(self, "autoscaling_policy")

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
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The type of the resource. Always compute#nodeGroup for node group.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="locationHint")
    def location_hint(self) -> str:
        """
        An opaque location hint used to place the Node close to other resources. This field is for use by internal tools that use the public API. The location hint here on the NodeGroup overrides any location_hint present in the NodeTemplate.
        """
        return pulumi.get(self, "location_hint")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> str:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. For more information, see Maintenance policies.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> 'outputs.NodeGroupMaintenanceWindowResponse':
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> str:
        """
        URL of the node template to create the node group from.
        """
        return pulumi.get(self, "node_template")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="shareSettings")
    def share_settings(self) -> 'outputs.ShareSettingsResponse':
        """
        Share-settings for the node group
        """
        return pulumi.get(self, "share_settings")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The total number of nodes in the node group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The name of the zone where the node group resides, such as us-central1-a.
        """
        return pulumi.get(self, "zone")


class AwaitableGetNodeGroupResult(GetNodeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodeGroupResult(
            autoscaling_policy=self.autoscaling_policy,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            fingerprint=self.fingerprint,
            kind=self.kind,
            location_hint=self.location_hint,
            maintenance_policy=self.maintenance_policy,
            maintenance_window=self.maintenance_window,
            name=self.name,
            node_template=self.node_template,
            self_link=self.self_link,
            share_settings=self.share_settings,
            size=self.size,
            status=self.status,
            zone=self.zone)


def get_node_group(node_group: Optional[str] = None,
                   project: Optional[str] = None,
                   zone: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodeGroupResult:
    """
    Returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the "nodes" field should not be used. Use nodeGroups.listNodes instead.
    """
    __args__ = dict()
    __args__['nodeGroup'] = node_group
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getNodeGroup', __args__, opts=opts, typ=GetNodeGroupResult).value

    return AwaitableGetNodeGroupResult(
        autoscaling_policy=pulumi.get(__ret__, 'autoscaling_policy'),
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        description=pulumi.get(__ret__, 'description'),
        fingerprint=pulumi.get(__ret__, 'fingerprint'),
        kind=pulumi.get(__ret__, 'kind'),
        location_hint=pulumi.get(__ret__, 'location_hint'),
        maintenance_policy=pulumi.get(__ret__, 'maintenance_policy'),
        maintenance_window=pulumi.get(__ret__, 'maintenance_window'),
        name=pulumi.get(__ret__, 'name'),
        node_template=pulumi.get(__ret__, 'node_template'),
        self_link=pulumi.get(__ret__, 'self_link'),
        share_settings=pulumi.get(__ret__, 'share_settings'),
        size=pulumi.get(__ret__, 'size'),
        status=pulumi.get(__ret__, 'status'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_node_group)
def get_node_group_output(node_group: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          zone: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNodeGroupResult]:
    """
    Returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the "nodes" field should not be used. Use nodeGroups.listNodes instead.
    """
    ...
