# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['InstanceGroupArgs', 'InstanceGroup']

@pulumi.input_type
class InstanceGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_ports: Optional[pulumi.Input[Sequence[pulumi.Input['NamedPortArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceGroup resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
        :param pulumi.Input[Sequence[pulumi.Input['NamedPortArgs']]] named_ports:  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group. 
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if named_ports is not None:
            pulumi.set(__self__, "named_ports", named_ports)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namedPorts")
    def named_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NamedPortArgs']]]]:
        """
         Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group. 
        """
        return pulumi.get(self, "named_ports")

    @named_ports.setter
    def named_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NamedPortArgs']]]]):
        pulumi.set(self, "named_ports", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamedPortArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an instance group in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamedPortArgs']]]] named_ports:  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group. 
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InstanceGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an instance group in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param InstanceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamedPortArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceGroupArgs.__new__(InstanceGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["named_ports"] = named_ports
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["zone"] = zone
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["network"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_id"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["subnetwork"] = None
        super(InstanceGroup, __self__).__init__(
            'google-native:compute/alpha:InstanceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InstanceGroup':
        """
        Get an existing InstanceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceGroupArgs.__new__(InstanceGroupArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["named_ports"] = None
        __props__.__dict__["network"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["subnetwork"] = None
        __props__.__dict__["zone"] = None
        return InstanceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        The creation timestamp for this instance group in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint of the named ports. The system uses this fingerprint to detect conflicts when multiple users change the named ports concurrently.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The resource type, which is always compute#instanceGroup for instance groups.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedPorts")
    def named_ports(self) -> pulumi.Output[Sequence['outputs.NamedPortResponse']]:
        """
         Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group. 
        """
        return pulumi.get(self, "named_ports")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The URL of the network to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The URL of the region where the instance group is located (for regional resources).
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URL for this instance group. The server generates this URL.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The total number of instances in the instance group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Output[str]:
        """
        The URL of the subnetwork to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The URL of the zone where the instance group is located (for zonal resources).
        """
        return pulumi.get(self, "zone")

