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
from ._enums import *

__all__ = ['AppGatewayArgs', 'AppGateway']

@pulumi.input_type
class AppGatewayArgs:
    def __init__(__self__, *,
                 host_type: pulumi.Input['AppGatewayHostType'],
                 type: pulumi.Input['AppGatewayType'],
                 app_gateway_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppGateway resource.
        :param pulumi.Input['AppGatewayHostType'] host_type: The type of hosting used by the AppGateway.
        :param pulumi.Input['AppGatewayType'] type: The type of network connectivity used by the AppGateway.
        :param pulumi.Input[str] app_gateway_id: Optional. User-settable AppGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        :param pulumi.Input[str] display_name: Optional. An arbitrary user-provided name for the AppGateway. Cannot exceed 64 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: Unique resource name of the AppGateway. The name is ignored when creating an AppGateway.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        pulumi.set(__self__, "host_type", host_type)
        pulumi.set(__self__, "type", type)
        if app_gateway_id is not None:
            pulumi.set(__self__, "app_gateway_id", app_gateway_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> pulumi.Input['AppGatewayHostType']:
        """
        The type of hosting used by the AppGateway.
        """
        return pulumi.get(self, "host_type")

    @host_type.setter
    def host_type(self, value: pulumi.Input['AppGatewayHostType']):
        pulumi.set(self, "host_type", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['AppGatewayType']:
        """
        The type of network connectivity used by the AppGateway.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['AppGatewayType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="appGatewayId")
    def app_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User-settable AppGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        """
        return pulumi.get(self, "app_gateway_id")

    @app_gateway_id.setter
    def app_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_gateway_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. An arbitrary user-provided name for the AppGateway. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique resource name of the AppGateway. The name is ignored when creating an AppGateway.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class AppGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_gateway_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input['AppGatewayHostType']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['AppGatewayType']] = None,
                 __props__=None):
        """
        Creates a new AppGateway in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_gateway_id: Optional. User-settable AppGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        :param pulumi.Input[str] display_name: Optional. An arbitrary user-provided name for the AppGateway. Cannot exceed 64 characters.
        :param pulumi.Input['AppGatewayHostType'] host_type: The type of hosting used by the AppGateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: Unique resource name of the AppGateway. The name is ignored when creating an AppGateway.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['AppGatewayType'] type: The type of network connectivity used by the AppGateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new AppGateway in a given project and location.

        :param str resource_name: The name of the resource.
        :param AppGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_gateway_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input['AppGatewayHostType']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['AppGatewayType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppGatewayArgs.__new__(AppGatewayArgs)

            __props__.__dict__["app_gateway_id"] = app_gateway_id
            __props__.__dict__["display_name"] = display_name
            if host_type is None and not opts.urn:
                raise TypeError("Missing required property 'host_type'")
            __props__.__dict__["host_type"] = host_type
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["allocated_connections"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["uri"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(AppGateway, __self__).__init__(
            'google-native:beyondcorp/v1alpha:AppGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AppGateway':
        """
        Get an existing AppGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AppGatewayArgs.__new__(AppGatewayArgs)

        __props__.__dict__["allocated_connections"] = None
        __props__.__dict__["app_gateway_id"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["host_type"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["uri"] = None
        return AppGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocatedConnections")
    def allocated_connections(self) -> pulumi.Output[Sequence['outputs.AllocatedConnectionResponse']]:
        """
        A list of connections allocated for the Gateway
        """
        return pulumi.get(self, "allocated_connections")

    @property
    @pulumi.getter(name="appGatewayId")
    def app_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. User-settable AppGateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        """
        return pulumi.get(self, "app_gateway_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. An arbitrary user-provided name for the AppGateway. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> pulumi.Output[str]:
        """
        The type of hosting used by the AppGateway.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique resource name of the AppGateway. The name is ignored when creating an AppGateway.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the AppGateway.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of network connectivity used by the AppGateway.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        A unique identifier for the instance generated by the system.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the resource was last modified.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        Server-defined URI for this resource.
        """
        return pulumi.get(self, "uri")

