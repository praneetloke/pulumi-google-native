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
    'GetRegionUrlMapResult',
    'AwaitableGetRegionUrlMapResult',
    'get_region_url_map',
    'get_region_url_map_output',
]

@pulumi.output_type
class GetRegionUrlMapResult:
    def __init__(__self__, creation_timestamp=None, default_custom_error_response_policy=None, default_route_action=None, default_service=None, default_url_redirect=None, description=None, fingerprint=None, header_action=None, host_rules=None, kind=None, name=None, path_matchers=None, region=None, self_link=None, tests=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if default_custom_error_response_policy and not isinstance(default_custom_error_response_policy, dict):
            raise TypeError("Expected argument 'default_custom_error_response_policy' to be a dict")
        pulumi.set(__self__, "default_custom_error_response_policy", default_custom_error_response_policy)
        if default_route_action and not isinstance(default_route_action, dict):
            raise TypeError("Expected argument 'default_route_action' to be a dict")
        pulumi.set(__self__, "default_route_action", default_route_action)
        if default_service and not isinstance(default_service, str):
            raise TypeError("Expected argument 'default_service' to be a str")
        pulumi.set(__self__, "default_service", default_service)
        if default_url_redirect and not isinstance(default_url_redirect, dict):
            raise TypeError("Expected argument 'default_url_redirect' to be a dict")
        pulumi.set(__self__, "default_url_redirect", default_url_redirect)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if header_action and not isinstance(header_action, dict):
            raise TypeError("Expected argument 'header_action' to be a dict")
        pulumi.set(__self__, "header_action", header_action)
        if host_rules and not isinstance(host_rules, list):
            raise TypeError("Expected argument 'host_rules' to be a list")
        pulumi.set(__self__, "host_rules", host_rules)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path_matchers and not isinstance(path_matchers, list):
            raise TypeError("Expected argument 'path_matchers' to be a list")
        pulumi.set(__self__, "path_matchers", path_matchers)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if tests and not isinstance(tests, list):
            raise TypeError("Expected argument 'tests' to be a list")
        pulumi.set(__self__, "tests", tests)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="defaultCustomErrorResponsePolicy")
    def default_custom_error_response_policy(self) -> 'outputs.CustomErrorResponsePolicyResponse':
        """
        defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
        """
        return pulumi.get(self, "default_custom_error_response_policy")

    @property
    @pulumi.getter(name="defaultRouteAction")
    def default_route_action(self) -> 'outputs.HttpRouteActionResponse':
        """
        defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
        """
        return pulumi.get(self, "default_route_action")

    @property
    @pulumi.getter(name="defaultService")
    def default_service(self) -> str:
        """
        The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
        """
        return pulumi.get(self, "default_service")

    @property
    @pulumi.getter(name="defaultUrlRedirect")
    def default_url_redirect(self) -> 'outputs.HttpRedirectActionResponse':
        """
        When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
        """
        return pulumi.get(self, "default_url_redirect")

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
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="headerAction")
    def header_action(self) -> 'outputs.HttpHeaderActionResponse':
        """
        Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        """
        return pulumi.get(self, "header_action")

    @property
    @pulumi.getter(name="hostRules")
    def host_rules(self) -> Sequence['outputs.HostRuleResponse']:
        """
        The list of host rules to use against the URL.
        """
        return pulumi.get(self, "host_rules")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#urlMaps for url maps.
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
    @pulumi.getter(name="pathMatchers")
    def path_matchers(self) -> Sequence['outputs.PathMatcherResponse']:
        """
        The list of named PathMatchers to use against the URL.
        """
        return pulumi.get(self, "path_matchers")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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
    def tests(self) -> Sequence['outputs.UrlMapTestResponse']:
        """
        The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        """
        return pulumi.get(self, "tests")


class AwaitableGetRegionUrlMapResult(GetRegionUrlMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionUrlMapResult(
            creation_timestamp=self.creation_timestamp,
            default_custom_error_response_policy=self.default_custom_error_response_policy,
            default_route_action=self.default_route_action,
            default_service=self.default_service,
            default_url_redirect=self.default_url_redirect,
            description=self.description,
            fingerprint=self.fingerprint,
            header_action=self.header_action,
            host_rules=self.host_rules,
            kind=self.kind,
            name=self.name,
            path_matchers=self.path_matchers,
            region=self.region,
            self_link=self.self_link,
            tests=self.tests)


def get_region_url_map(project: Optional[str] = None,
                       region: Optional[str] = None,
                       url_map: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionUrlMapResult:
    """
    Returns the specified UrlMap resource.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['urlMap'] = url_map
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getRegionUrlMap', __args__, opts=opts, typ=GetRegionUrlMapResult).value

    return AwaitableGetRegionUrlMapResult(
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        default_custom_error_response_policy=pulumi.get(__ret__, 'default_custom_error_response_policy'),
        default_route_action=pulumi.get(__ret__, 'default_route_action'),
        default_service=pulumi.get(__ret__, 'default_service'),
        default_url_redirect=pulumi.get(__ret__, 'default_url_redirect'),
        description=pulumi.get(__ret__, 'description'),
        fingerprint=pulumi.get(__ret__, 'fingerprint'),
        header_action=pulumi.get(__ret__, 'header_action'),
        host_rules=pulumi.get(__ret__, 'host_rules'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        path_matchers=pulumi.get(__ret__, 'path_matchers'),
        region=pulumi.get(__ret__, 'region'),
        self_link=pulumi.get(__ret__, 'self_link'),
        tests=pulumi.get(__ret__, 'tests'))


@_utilities.lift_output_func(get_region_url_map)
def get_region_url_map_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                              region: Optional[pulumi.Input[str]] = None,
                              url_map: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionUrlMapResult]:
    """
    Returns the specified UrlMap resource.
    """
    ...
