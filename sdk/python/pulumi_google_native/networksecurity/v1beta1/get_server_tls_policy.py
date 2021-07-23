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
    'GetServerTlsPolicyResult',
    'AwaitableGetServerTlsPolicyResult',
    'get_server_tls_policy',
]

@pulumi.output_type
class GetServerTlsPolicyResult:
    def __init__(__self__, allow_open=None, create_time=None, description=None, labels=None, mtls_policy=None, name=None, server_certificate=None, update_time=None):
        if allow_open and not isinstance(allow_open, bool):
            raise TypeError("Expected argument 'allow_open' to be a bool")
        pulumi.set(__self__, "allow_open", allow_open)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if mtls_policy and not isinstance(mtls_policy, dict):
            raise TypeError("Expected argument 'mtls_policy' to be a dict")
        pulumi.set(__self__, "mtls_policy", mtls_policy)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if server_certificate and not isinstance(server_certificate, dict):
            raise TypeError("Expected argument 'server_certificate' to be a dict")
        pulumi.set(__self__, "server_certificate", server_certificate)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="allowOpen")
    def allow_open(self) -> bool:
        """
         Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        """
        return pulumi.get(self, "allow_open")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Set of label tags associated with the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mtlsPolicy")
    def mtls_policy(self) -> 'outputs.MTLSPolicyResponse':
        """
         Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        """
        return pulumi.get(self, "mtls_policy")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> 'outputs.GoogleCloudNetworksecurityV1beta1CertificateProviderResponse':
        """
         Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp when the resource was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetServerTlsPolicyResult(GetServerTlsPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerTlsPolicyResult(
            allow_open=self.allow_open,
            create_time=self.create_time,
            description=self.description,
            labels=self.labels,
            mtls_policy=self.mtls_policy,
            name=self.name,
            server_certificate=self.server_certificate,
            update_time=self.update_time)


def get_server_tls_policy(location: Optional[str] = None,
                          project: Optional[str] = None,
                          server_tls_policy_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerTlsPolicyResult:
    """
    Gets details of a single ServerTlsPolicy.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['serverTlsPolicyId'] = server_tls_policy_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:networksecurity/v1beta1:getServerTlsPolicy', __args__, opts=opts, typ=GetServerTlsPolicyResult).value

    return AwaitableGetServerTlsPolicyResult(
        allow_open=__ret__.allow_open,
        create_time=__ret__.create_time,
        description=__ret__.description,
        labels=__ret__.labels,
        mtls_policy=__ret__.mtls_policy,
        name=__ret__.name,
        server_certificate=__ret__.server_certificate,
        update_time=__ret__.update_time)