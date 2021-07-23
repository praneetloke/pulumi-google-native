# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetRegionHealthCheckServiceResult',
    'AwaitableGetRegionHealthCheckServiceResult',
    'get_region_health_check_service',
]

@pulumi.output_type
class GetRegionHealthCheckServiceResult:
    def __init__(__self__, creation_timestamp=None, description=None, fingerprint=None, health_checks=None, health_status_aggregation_policy=None, kind=None, name=None, network_endpoint_groups=None, notification_endpoints=None, region=None, self_link=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if health_checks and not isinstance(health_checks, list):
            raise TypeError("Expected argument 'health_checks' to be a list")
        pulumi.set(__self__, "health_checks", health_checks)
        if health_status_aggregation_policy and not isinstance(health_status_aggregation_policy, str):
            raise TypeError("Expected argument 'health_status_aggregation_policy' to be a str")
        pulumi.set(__self__, "health_status_aggregation_policy", health_status_aggregation_policy)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_endpoint_groups and not isinstance(network_endpoint_groups, list):
            raise TypeError("Expected argument 'network_endpoint_groups' to be a list")
        pulumi.set(__self__, "network_endpoint_groups", network_endpoint_groups)
        if notification_endpoints and not isinstance(notification_endpoints, list):
            raise TypeError("Expected argument 'notification_endpoints' to be a list")
        pulumi.set(__self__, "notification_endpoints", notification_endpoints)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

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
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="healthChecks")
    def health_checks(self) -> Sequence[str]:
        """
        A list of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT or portSpecification=USE_FIXED_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks must belong to the same region as zones of NEGs.
        """
        return pulumi.get(self, "health_checks")

    @property
    @pulumi.getter(name="healthStatusAggregationPolicy")
    def health_status_aggregation_policy(self) -> str:
        """
        Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified. - NO_AGGREGATION. An EndpointHealth message is returned for each pair in the health check service. - AND. If any health check of an endpoint reports UNHEALTHY, then UNHEALTHY is the HealthState of the endpoint. If all health checks report HEALTHY, the HealthState of the endpoint is HEALTHY. .
        """
        return pulumi.get(self, "health_status_aggregation_policy")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkEndpointGroups")
    def network_endpoint_groups(self) -> Sequence[str]:
        """
        A list of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
        """
        return pulumi.get(self, "network_endpoint_groups")

    @property
    @pulumi.getter(name="notificationEndpoints")
    def notification_endpoints(self) -> Sequence[str]:
        """
        A list of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
        """
        return pulumi.get(self, "notification_endpoints")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")


class AwaitableGetRegionHealthCheckServiceResult(GetRegionHealthCheckServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionHealthCheckServiceResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            fingerprint=self.fingerprint,
            health_checks=self.health_checks,
            health_status_aggregation_policy=self.health_status_aggregation_policy,
            kind=self.kind,
            name=self.name,
            network_endpoint_groups=self.network_endpoint_groups,
            notification_endpoints=self.notification_endpoints,
            region=self.region,
            self_link=self.self_link)


def get_region_health_check_service(health_check_service: Optional[str] = None,
                                    project: Optional[str] = None,
                                    region: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionHealthCheckServiceResult:
    """
    Returns the specified regional HealthCheckService resource.
    """
    __args__ = dict()
    __args__['healthCheckService'] = health_check_service
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getRegionHealthCheckService', __args__, opts=opts, typ=GetRegionHealthCheckServiceResult).value

    return AwaitableGetRegionHealthCheckServiceResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        fingerprint=__ret__.fingerprint,
        health_checks=__ret__.health_checks,
        health_status_aggregation_policy=__ret__.health_status_aggregation_policy,
        kind=__ret__.kind,
        name=__ret__.name,
        network_endpoint_groups=__ret__.network_endpoint_groups,
        notification_endpoints=__ret__.notification_endpoints,
        region=__ret__.region,
        self_link=__ret__.self_link)