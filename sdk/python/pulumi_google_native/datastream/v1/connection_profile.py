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

__all__ = ['ConnectionProfileArgs', 'ConnectionProfile']

@pulumi.input_type
class ConnectionProfileArgs:
    def __init__(__self__, *,
                 connection_profile_id: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 force: Optional[pulumi.Input[str]] = None,
                 forward_ssh_connectivity: Optional[pulumi.Input['ForwardSshTunnelConnectivityArgs']] = None,
                 gcs_profile: Optional[pulumi.Input['GcsProfileArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql_profile: Optional[pulumi.Input['MysqlProfileArgs']] = None,
                 oracle_profile: Optional[pulumi.Input['OracleProfileArgs']] = None,
                 private_connectivity: Optional[pulumi.Input['PrivateConnectivityArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 static_service_ip_connectivity: Optional[pulumi.Input['StaticServiceIpConnectivityArgs']] = None,
                 validate_only: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConnectionProfile resource.
        :param pulumi.Input[str] display_name: Display name.
        :param pulumi.Input['ForwardSshTunnelConnectivityArgs'] forward_ssh_connectivity: Forward SSH tunnel connectivity.
        :param pulumi.Input['GcsProfileArgs'] gcs_profile: Cloud Storage ConnectionProfile configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels.
        :param pulumi.Input['MysqlProfileArgs'] mysql_profile: MySQL ConnectionProfile configuration.
        :param pulumi.Input['OracleProfileArgs'] oracle_profile: Oracle ConnectionProfile configuration.
        :param pulumi.Input['PrivateConnectivityArgs'] private_connectivity: Private connectivity.
        :param pulumi.Input['StaticServiceIpConnectivityArgs'] static_service_ip_connectivity: Static Service IP connectivity.
        """
        pulumi.set(__self__, "connection_profile_id", connection_profile_id)
        pulumi.set(__self__, "display_name", display_name)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if forward_ssh_connectivity is not None:
            pulumi.set(__self__, "forward_ssh_connectivity", forward_ssh_connectivity)
        if gcs_profile is not None:
            pulumi.set(__self__, "gcs_profile", gcs_profile)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mysql_profile is not None:
            pulumi.set(__self__, "mysql_profile", mysql_profile)
        if oracle_profile is not None:
            pulumi.set(__self__, "oracle_profile", oracle_profile)
        if private_connectivity is not None:
            pulumi.set(__self__, "private_connectivity", private_connectivity)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if static_service_ip_connectivity is not None:
            pulumi.set(__self__, "static_service_ip_connectivity", static_service_ip_connectivity)
        if validate_only is not None:
            pulumi.set(__self__, "validate_only", validate_only)

    @property
    @pulumi.getter(name="connectionProfileId")
    def connection_profile_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "connection_profile_id")

    @connection_profile_id.setter
    def connection_profile_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_profile_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="forwardSshConnectivity")
    def forward_ssh_connectivity(self) -> Optional[pulumi.Input['ForwardSshTunnelConnectivityArgs']]:
        """
        Forward SSH tunnel connectivity.
        """
        return pulumi.get(self, "forward_ssh_connectivity")

    @forward_ssh_connectivity.setter
    def forward_ssh_connectivity(self, value: Optional[pulumi.Input['ForwardSshTunnelConnectivityArgs']]):
        pulumi.set(self, "forward_ssh_connectivity", value)

    @property
    @pulumi.getter(name="gcsProfile")
    def gcs_profile(self) -> Optional[pulumi.Input['GcsProfileArgs']]:
        """
        Cloud Storage ConnectionProfile configuration.
        """
        return pulumi.get(self, "gcs_profile")

    @gcs_profile.setter
    def gcs_profile(self, value: Optional[pulumi.Input['GcsProfileArgs']]):
        pulumi.set(self, "gcs_profile", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels.
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
    @pulumi.getter(name="mysqlProfile")
    def mysql_profile(self) -> Optional[pulumi.Input['MysqlProfileArgs']]:
        """
        MySQL ConnectionProfile configuration.
        """
        return pulumi.get(self, "mysql_profile")

    @mysql_profile.setter
    def mysql_profile(self, value: Optional[pulumi.Input['MysqlProfileArgs']]):
        pulumi.set(self, "mysql_profile", value)

    @property
    @pulumi.getter(name="oracleProfile")
    def oracle_profile(self) -> Optional[pulumi.Input['OracleProfileArgs']]:
        """
        Oracle ConnectionProfile configuration.
        """
        return pulumi.get(self, "oracle_profile")

    @oracle_profile.setter
    def oracle_profile(self, value: Optional[pulumi.Input['OracleProfileArgs']]):
        pulumi.set(self, "oracle_profile", value)

    @property
    @pulumi.getter(name="privateConnectivity")
    def private_connectivity(self) -> Optional[pulumi.Input['PrivateConnectivityArgs']]:
        """
        Private connectivity.
        """
        return pulumi.get(self, "private_connectivity")

    @private_connectivity.setter
    def private_connectivity(self, value: Optional[pulumi.Input['PrivateConnectivityArgs']]):
        pulumi.set(self, "private_connectivity", value)

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
    @pulumi.getter(name="staticServiceIpConnectivity")
    def static_service_ip_connectivity(self) -> Optional[pulumi.Input['StaticServiceIpConnectivityArgs']]:
        """
        Static Service IP connectivity.
        """
        return pulumi.get(self, "static_service_ip_connectivity")

    @static_service_ip_connectivity.setter
    def static_service_ip_connectivity(self, value: Optional[pulumi.Input['StaticServiceIpConnectivityArgs']]):
        pulumi.set(self, "static_service_ip_connectivity", value)

    @property
    @pulumi.getter(name="validateOnly")
    def validate_only(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "validate_only")

    @validate_only.setter
    def validate_only(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_only", value)


class ConnectionProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_profile_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[str]] = None,
                 forward_ssh_connectivity: Optional[pulumi.Input[pulumi.InputType['ForwardSshTunnelConnectivityArgs']]] = None,
                 gcs_profile: Optional[pulumi.Input[pulumi.InputType['GcsProfileArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql_profile: Optional[pulumi.Input[pulumi.InputType['MysqlProfileArgs']]] = None,
                 oracle_profile: Optional[pulumi.Input[pulumi.InputType['OracleProfileArgs']]] = None,
                 private_connectivity: Optional[pulumi.Input[pulumi.InputType['PrivateConnectivityArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 static_service_ip_connectivity: Optional[pulumi.Input[pulumi.InputType['StaticServiceIpConnectivityArgs']]] = None,
                 validate_only: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this method to create a connection profile in a project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Display name.
        :param pulumi.Input[pulumi.InputType['ForwardSshTunnelConnectivityArgs']] forward_ssh_connectivity: Forward SSH tunnel connectivity.
        :param pulumi.Input[pulumi.InputType['GcsProfileArgs']] gcs_profile: Cloud Storage ConnectionProfile configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels.
        :param pulumi.Input[pulumi.InputType['MysqlProfileArgs']] mysql_profile: MySQL ConnectionProfile configuration.
        :param pulumi.Input[pulumi.InputType['OracleProfileArgs']] oracle_profile: Oracle ConnectionProfile configuration.
        :param pulumi.Input[pulumi.InputType['PrivateConnectivityArgs']] private_connectivity: Private connectivity.
        :param pulumi.Input[pulumi.InputType['StaticServiceIpConnectivityArgs']] static_service_ip_connectivity: Static Service IP connectivity.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this method to create a connection profile in a project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ConnectionProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_profile_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[str]] = None,
                 forward_ssh_connectivity: Optional[pulumi.Input[pulumi.InputType['ForwardSshTunnelConnectivityArgs']]] = None,
                 gcs_profile: Optional[pulumi.Input[pulumi.InputType['GcsProfileArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql_profile: Optional[pulumi.Input[pulumi.InputType['MysqlProfileArgs']]] = None,
                 oracle_profile: Optional[pulumi.Input[pulumi.InputType['OracleProfileArgs']]] = None,
                 private_connectivity: Optional[pulumi.Input[pulumi.InputType['PrivateConnectivityArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 static_service_ip_connectivity: Optional[pulumi.Input[pulumi.InputType['StaticServiceIpConnectivityArgs']]] = None,
                 validate_only: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConnectionProfileArgs.__new__(ConnectionProfileArgs)

            if connection_profile_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_profile_id'")
            __props__.__dict__["connection_profile_id"] = connection_profile_id
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["force"] = force
            __props__.__dict__["forward_ssh_connectivity"] = forward_ssh_connectivity
            __props__.__dict__["gcs_profile"] = gcs_profile
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["mysql_profile"] = mysql_profile
            __props__.__dict__["oracle_profile"] = oracle_profile
            __props__.__dict__["private_connectivity"] = private_connectivity
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["static_service_ip_connectivity"] = static_service_ip_connectivity
            __props__.__dict__["validate_only"] = validate_only
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(ConnectionProfile, __self__).__init__(
            'google-native:datastream/v1:ConnectionProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConnectionProfile':
        """
        Get an existing ConnectionProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectionProfileArgs.__new__(ConnectionProfileArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["forward_ssh_connectivity"] = None
        __props__.__dict__["gcs_profile"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["mysql_profile"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["oracle_profile"] = None
        __props__.__dict__["private_connectivity"] = None
        __props__.__dict__["static_service_ip_connectivity"] = None
        __props__.__dict__["update_time"] = None
        return ConnectionProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The create time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="forwardSshConnectivity")
    def forward_ssh_connectivity(self) -> pulumi.Output['outputs.ForwardSshTunnelConnectivityResponse']:
        """
        Forward SSH tunnel connectivity.
        """
        return pulumi.get(self, "forward_ssh_connectivity")

    @property
    @pulumi.getter(name="gcsProfile")
    def gcs_profile(self) -> pulumi.Output['outputs.GcsProfileResponse']:
        """
        Cloud Storage ConnectionProfile configuration.
        """
        return pulumi.get(self, "gcs_profile")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mysqlProfile")
    def mysql_profile(self) -> pulumi.Output['outputs.MysqlProfileResponse']:
        """
        MySQL ConnectionProfile configuration.
        """
        return pulumi.get(self, "mysql_profile")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oracleProfile")
    def oracle_profile(self) -> pulumi.Output['outputs.OracleProfileResponse']:
        """
        Oracle ConnectionProfile configuration.
        """
        return pulumi.get(self, "oracle_profile")

    @property
    @pulumi.getter(name="privateConnectivity")
    def private_connectivity(self) -> pulumi.Output['outputs.PrivateConnectivityResponse']:
        """
        Private connectivity.
        """
        return pulumi.get(self, "private_connectivity")

    @property
    @pulumi.getter(name="staticServiceIpConnectivity")
    def static_service_ip_connectivity(self) -> pulumi.Output['outputs.StaticServiceIpConnectivityResponse']:
        """
        Static Service IP connectivity.
        """
        return pulumi.get(self, "static_service_ip_connectivity")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of the resource.
        """
        return pulumi.get(self, "update_time")

