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
    'GetRegionTargetHttpsProxyResult',
    'AwaitableGetRegionTargetHttpsProxyResult',
    'get_region_target_https_proxy',
    'get_region_target_https_proxy_output',
]

@pulumi.output_type
class GetRegionTargetHttpsProxyResult:
    def __init__(__self__, authorization_policy=None, certificate_map=None, creation_timestamp=None, description=None, fingerprint=None, kind=None, name=None, proxy_bind=None, quic_override=None, region=None, self_link=None, server_tls_policy=None, ssl_certificates=None, ssl_policy=None, url_map=None):
        if authorization_policy and not isinstance(authorization_policy, str):
            raise TypeError("Expected argument 'authorization_policy' to be a str")
        pulumi.set(__self__, "authorization_policy", authorization_policy)
        if certificate_map and not isinstance(certificate_map, str):
            raise TypeError("Expected argument 'certificate_map' to be a str")
        pulumi.set(__self__, "certificate_map", certificate_map)
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
        if proxy_bind and not isinstance(proxy_bind, bool):
            raise TypeError("Expected argument 'proxy_bind' to be a bool")
        pulumi.set(__self__, "proxy_bind", proxy_bind)
        if quic_override and not isinstance(quic_override, str):
            raise TypeError("Expected argument 'quic_override' to be a str")
        pulumi.set(__self__, "quic_override", quic_override)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if server_tls_policy and not isinstance(server_tls_policy, str):
            raise TypeError("Expected argument 'server_tls_policy' to be a str")
        pulumi.set(__self__, "server_tls_policy", server_tls_policy)
        if ssl_certificates and not isinstance(ssl_certificates, list):
            raise TypeError("Expected argument 'ssl_certificates' to be a list")
        pulumi.set(__self__, "ssl_certificates", ssl_certificates)
        if ssl_policy and not isinstance(ssl_policy, str):
            raise TypeError("Expected argument 'ssl_policy' to be a str")
        pulumi.set(__self__, "ssl_policy", ssl_policy)
        if url_map and not isinstance(url_map, str):
            raise TypeError("Expected argument 'url_map' to be a str")
        pulumi.set(__self__, "url_map", url_map)

    @property
    @pulumi.getter(name="authorizationPolicy")
    def authorization_policy(self) -> str:
        """
        Optional. A URL referring to a networksecurity.AuthorizationPolicy resource that describes how the proxy should authorize inbound traffic. If left blank, access will not be restricted by an authorization policy. Refer to the AuthorizationPolicy resource for additional details. authorizationPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED. Note: This field currently has no impact.
        """
        return pulumi.get(self, "authorization_policy")

    @property
    @pulumi.getter(name="certificateMap")
    def certificate_map(self) -> str:
        """
        URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
        """
        return pulumi.get(self, "certificate_map")

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
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpsProxy. An up-to-date fingerprint must be provided in order to patch the TargetHttpsProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpsProxy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of resource. Always compute#targetHttpsProxy for target HTTPS proxies.
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
    @pulumi.getter(name="proxyBind")
    def proxy_bind(self) -> bool:
        """
        This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
        """
        return pulumi.get(self, "proxy_bind")

    @property
    @pulumi.getter(name="quicOverride")
    def quic_override(self) -> str:
        """
        Specifies the QUIC override policy for this TargetHttpsProxy resource. This setting determines whether the load balancer attempts to negotiate QUIC with clients. You can specify NONE, ENABLE, or DISABLE. - When quic-override is set to NONE, Google manages whether QUIC is used. - When quic-override is set to ENABLE, the load balancer uses QUIC when possible. - When quic-override is set to DISABLE, the load balancer doesn't use QUIC. - If the quic-override flag is not specified, NONE is implied. 
        """
        return pulumi.get(self, "quic_override")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional TargetHttpsProxy resides. This field is not applicable to global TargetHttpsProxies.
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
    @pulumi.getter(name="serverTlsPolicy")
    def server_tls_policy(self) -> str:
        """
        Optional. A URL referring to a networksecurity.ServerTlsPolicy resource that describes how the proxy should authenticate inbound traffic. serverTlsPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED. For details which ServerTlsPolicy resources are accepted with INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED loadBalancingScheme consult ServerTlsPolicy documentation. If left blank, communications are not encrypted.
        """
        return pulumi.get(self, "server_tls_policy")

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> Sequence[str]:
        """
        URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "ssl_certificates")

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> str:
        """
        URL of SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the TargetHttpsProxy resource has no SSL policy configured.
        """
        return pulumi.get(self, "ssl_policy")

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> str:
        """
        A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService. For example, the following are all valid URLs for specifying a URL map: - https://www.googleapis.compute/v1/projects/project/global/urlMaps/ url-map - projects/project/global/urlMaps/url-map - global/urlMaps/url-map 
        """
        return pulumi.get(self, "url_map")


class AwaitableGetRegionTargetHttpsProxyResult(GetRegionTargetHttpsProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionTargetHttpsProxyResult(
            authorization_policy=self.authorization_policy,
            certificate_map=self.certificate_map,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            fingerprint=self.fingerprint,
            kind=self.kind,
            name=self.name,
            proxy_bind=self.proxy_bind,
            quic_override=self.quic_override,
            region=self.region,
            self_link=self.self_link,
            server_tls_policy=self.server_tls_policy,
            ssl_certificates=self.ssl_certificates,
            ssl_policy=self.ssl_policy,
            url_map=self.url_map)


def get_region_target_https_proxy(project: Optional[str] = None,
                                  region: Optional[str] = None,
                                  target_https_proxy: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionTargetHttpsProxyResult:
    """
    Returns the specified TargetHttpsProxy resource in the specified region.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['targetHttpsProxy'] = target_https_proxy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getRegionTargetHttpsProxy', __args__, opts=opts, typ=GetRegionTargetHttpsProxyResult).value

    return AwaitableGetRegionTargetHttpsProxyResult(
        authorization_policy=__ret__.authorization_policy,
        certificate_map=__ret__.certificate_map,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        fingerprint=__ret__.fingerprint,
        kind=__ret__.kind,
        name=__ret__.name,
        proxy_bind=__ret__.proxy_bind,
        quic_override=__ret__.quic_override,
        region=__ret__.region,
        self_link=__ret__.self_link,
        server_tls_policy=__ret__.server_tls_policy,
        ssl_certificates=__ret__.ssl_certificates,
        ssl_policy=__ret__.ssl_policy,
        url_map=__ret__.url_map)


@_utilities.lift_output_func(get_region_target_https_proxy)
def get_region_target_https_proxy_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                         region: Optional[pulumi.Input[str]] = None,
                                         target_https_proxy: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionTargetHttpsProxyResult]:
    """
    Returns the specified TargetHttpsProxy resource in the specified region.
    """
    ...
