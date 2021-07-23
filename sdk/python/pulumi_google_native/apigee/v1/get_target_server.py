# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetTargetServerResult',
    'AwaitableGetTargetServerResult',
    'get_target_server',
]

@pulumi.output_type
class GetTargetServerResult:
    def __init__(__self__, description=None, host=None, is_enabled=None, name=None, port=None, protocol=None, s_sl_info=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if is_enabled and not isinstance(is_enabled, bool):
            raise TypeError("Expected argument 'is_enabled' to be a bool")
        pulumi.set(__self__, "is_enabled", is_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if s_sl_info and not isinstance(s_sl_info, dict):
            raise TypeError("Expected argument 's_sl_info' to be a dict")
        pulumi.set(__self__, "s_sl_info", s_sl_info)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A human-readable description of this TargetServer.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> bool:
        """
        Optional. Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource id of this target server. Values must match the regular expression 
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Immutable. The protocol used by this TargetServer.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sSLInfo")
    def s_sl_info(self) -> 'outputs.GoogleCloudApigeeV1TlsInfoResponse':
        """
        Optional. Specifies TLS configuration info for this TargetServer. The JSON name is `sSLInfo` for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        """
        return pulumi.get(self, "s_sl_info")


class AwaitableGetTargetServerResult(GetTargetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetServerResult(
            description=self.description,
            host=self.host,
            is_enabled=self.is_enabled,
            name=self.name,
            port=self.port,
            protocol=self.protocol,
            s_sl_info=self.s_sl_info)


def get_target_server(environment_id: Optional[str] = None,
                      organization_id: Optional[str] = None,
                      targetserver_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetServerResult:
    """
    Gets a TargetServer resource.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    __args__['organizationId'] = organization_id
    __args__['targetserverId'] = targetserver_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getTargetServer', __args__, opts=opts, typ=GetTargetServerResult).value

    return AwaitableGetTargetServerResult(
        description=__ret__.description,
        host=__ret__.host,
        is_enabled=__ret__.is_enabled,
        name=__ret__.name,
        port=__ret__.port,
        protocol=__ret__.protocol,
        s_sl_info=__ret__.s_sl_info)