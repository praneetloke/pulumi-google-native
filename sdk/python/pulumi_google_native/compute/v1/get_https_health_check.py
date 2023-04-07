# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetHttpsHealthCheckResult',
    'AwaitableGetHttpsHealthCheckResult',
    'get_https_health_check',
    'get_https_health_check_output',
]

@pulumi.output_type
class GetHttpsHealthCheckResult:
    def __init__(__self__, check_interval_sec=None, creation_timestamp=None, description=None, healthy_threshold=None, host=None, kind=None, name=None, port=None, request_path=None, self_link=None, timeout_sec=None, unhealthy_threshold=None):
        if check_interval_sec and not isinstance(check_interval_sec, int):
            raise TypeError("Expected argument 'check_interval_sec' to be a int")
        pulumi.set(__self__, "check_interval_sec", check_interval_sec)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if healthy_threshold and not isinstance(healthy_threshold, int):
            raise TypeError("Expected argument 'healthy_threshold' to be a int")
        pulumi.set(__self__, "healthy_threshold", healthy_threshold)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if request_path and not isinstance(request_path, str):
            raise TypeError("Expected argument 'request_path' to be a str")
        pulumi.set(__self__, "request_path", request_path)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if timeout_sec and not isinstance(timeout_sec, int):
            raise TypeError("Expected argument 'timeout_sec' to be a int")
        pulumi.set(__self__, "timeout_sec", timeout_sec)
        if unhealthy_threshold and not isinstance(unhealthy_threshold, int):
            raise TypeError("Expected argument 'unhealthy_threshold' to be a int")
        pulumi.set(__self__, "unhealthy_threshold", unhealthy_threshold)

    @property
    @pulumi.getter(name="checkIntervalSec")
    def check_interval_sec(self) -> int:
        """
        How often (in seconds) to send a health check. The default value is 5 seconds.
        """
        return pulumi.get(self, "check_interval_sec")

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
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> int:
        """
        A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        """
        return pulumi.get(self, "healthy_threshold")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        """
        return pulumi.get(self, "host")

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
    def port(self) -> int:
        """
        The TCP port number for the HTTPS health check request. The default value is 443.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="requestPath")
    def request_path(self) -> str:
        """
        The request path of the HTTPS health check request. The default value is "/". Must comply with RFC3986.
        """
        return pulumi.get(self, "request_path")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> int:
        """
        How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
        """
        return pulumi.get(self, "timeout_sec")

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> int:
        """
        A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        return pulumi.get(self, "unhealthy_threshold")


class AwaitableGetHttpsHealthCheckResult(GetHttpsHealthCheckResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHttpsHealthCheckResult(
            check_interval_sec=self.check_interval_sec,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            healthy_threshold=self.healthy_threshold,
            host=self.host,
            kind=self.kind,
            name=self.name,
            port=self.port,
            request_path=self.request_path,
            self_link=self.self_link,
            timeout_sec=self.timeout_sec,
            unhealthy_threshold=self.unhealthy_threshold)


def get_https_health_check(https_health_check: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHttpsHealthCheckResult:
    """
    Returns the specified HttpsHealthCheck resource.
    """
    __args__ = dict()
    __args__['httpsHealthCheck'] = https_health_check
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getHttpsHealthCheck', __args__, opts=opts, typ=GetHttpsHealthCheckResult).value

    return AwaitableGetHttpsHealthCheckResult(
        check_interval_sec=__ret__.check_interval_sec,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        healthy_threshold=__ret__.healthy_threshold,
        host=__ret__.host,
        kind=__ret__.kind,
        name=__ret__.name,
        port=__ret__.port,
        request_path=__ret__.request_path,
        self_link=__ret__.self_link,
        timeout_sec=__ret__.timeout_sec,
        unhealthy_threshold=__ret__.unhealthy_threshold)


@_utilities.lift_output_func(get_https_health_check)
def get_https_health_check_output(https_health_check: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHttpsHealthCheckResult]:
    """
    Returns the specified HttpsHealthCheck resource.
    """
    ...
