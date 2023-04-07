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
    'GetSubnetworkResult',
    'AwaitableGetSubnetworkResult',
    'get_subnetwork',
    'get_subnetwork_output',
]

@pulumi.output_type
class GetSubnetworkResult:
    def __init__(__self__, allow_subnet_cidr_routes_overlap=None, creation_timestamp=None, description=None, enable_flow_logs=None, external_ipv6_prefix=None, fingerprint=None, gateway_address=None, internal_ipv6_prefix=None, ip_cidr_range=None, ipv6_access_type=None, ipv6_cidr_range=None, kind=None, log_config=None, name=None, network=None, private_ip_google_access=None, private_ipv6_google_access=None, purpose=None, region=None, reserved_internal_range=None, role=None, secondary_ip_ranges=None, self_link=None, stack_type=None, state=None):
        if allow_subnet_cidr_routes_overlap and not isinstance(allow_subnet_cidr_routes_overlap, bool):
            raise TypeError("Expected argument 'allow_subnet_cidr_routes_overlap' to be a bool")
        pulumi.set(__self__, "allow_subnet_cidr_routes_overlap", allow_subnet_cidr_routes_overlap)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_flow_logs and not isinstance(enable_flow_logs, bool):
            raise TypeError("Expected argument 'enable_flow_logs' to be a bool")
        pulumi.set(__self__, "enable_flow_logs", enable_flow_logs)
        if external_ipv6_prefix and not isinstance(external_ipv6_prefix, str):
            raise TypeError("Expected argument 'external_ipv6_prefix' to be a str")
        pulumi.set(__self__, "external_ipv6_prefix", external_ipv6_prefix)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if gateway_address and not isinstance(gateway_address, str):
            raise TypeError("Expected argument 'gateway_address' to be a str")
        pulumi.set(__self__, "gateway_address", gateway_address)
        if internal_ipv6_prefix and not isinstance(internal_ipv6_prefix, str):
            raise TypeError("Expected argument 'internal_ipv6_prefix' to be a str")
        pulumi.set(__self__, "internal_ipv6_prefix", internal_ipv6_prefix)
        if ip_cidr_range and not isinstance(ip_cidr_range, str):
            raise TypeError("Expected argument 'ip_cidr_range' to be a str")
        pulumi.set(__self__, "ip_cidr_range", ip_cidr_range)
        if ipv6_access_type and not isinstance(ipv6_access_type, str):
            raise TypeError("Expected argument 'ipv6_access_type' to be a str")
        pulumi.set(__self__, "ipv6_access_type", ipv6_access_type)
        if ipv6_cidr_range and not isinstance(ipv6_cidr_range, str):
            raise TypeError("Expected argument 'ipv6_cidr_range' to be a str")
        pulumi.set(__self__, "ipv6_cidr_range", ipv6_cidr_range)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if log_config and not isinstance(log_config, dict):
            raise TypeError("Expected argument 'log_config' to be a dict")
        pulumi.set(__self__, "log_config", log_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if private_ip_google_access and not isinstance(private_ip_google_access, bool):
            raise TypeError("Expected argument 'private_ip_google_access' to be a bool")
        pulumi.set(__self__, "private_ip_google_access", private_ip_google_access)
        if private_ipv6_google_access and not isinstance(private_ipv6_google_access, str):
            raise TypeError("Expected argument 'private_ipv6_google_access' to be a str")
        pulumi.set(__self__, "private_ipv6_google_access", private_ipv6_google_access)
        if purpose and not isinstance(purpose, str):
            raise TypeError("Expected argument 'purpose' to be a str")
        pulumi.set(__self__, "purpose", purpose)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if reserved_internal_range and not isinstance(reserved_internal_range, str):
            raise TypeError("Expected argument 'reserved_internal_range' to be a str")
        pulumi.set(__self__, "reserved_internal_range", reserved_internal_range)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if secondary_ip_ranges and not isinstance(secondary_ip_ranges, list):
            raise TypeError("Expected argument 'secondary_ip_ranges' to be a list")
        pulumi.set(__self__, "secondary_ip_ranges", secondary_ip_ranges)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if stack_type and not isinstance(stack_type, str):
            raise TypeError("Expected argument 'stack_type' to be a str")
        pulumi.set(__self__, "stack_type", stack_type)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="allowSubnetCidrRoutesOverlap")
    def allow_subnet_cidr_routes_overlap(self) -> bool:
        """
        Whether this subnetwork's ranges can conflict with existing static routes. Setting this to true allows this subnetwork's primary and secondary ranges to overlap with (and contain) static routes that have already been configured on the corresponding network. For example if a static route has range 10.1.0.0/16, a subnet range 10.0.0.0/8 could only be created if allow_conflicting_routes=true. Overlapping is only allowed on subnetwork operations; routes whose ranges conflict with this subnetwork's ranges won't be allowed unless route.allow_conflicting_subnetworks is set to true. Typically packets destined to IPs within the subnetwork (which may contain private/sensitive data) are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
        """
        return pulumi.get(self, "allow_subnet_cidr_routes_overlap")

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
        An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableFlowLogs")
    def enable_flow_logs(self) -> bool:
        """
        Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is determined by the org policy, if there is no org policy specified, then it will default to disabled. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        """
        return pulumi.get(self, "enable_flow_logs")

    @property
    @pulumi.getter(name="externalIpv6Prefix")
    def external_ipv6_prefix(self) -> str:
        """
        The external IPv6 address range that is owned by this subnetwork.
        """
        return pulumi.get(self, "external_ipv6_prefix")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="gatewayAddress")
    def gateway_address(self) -> str:
        """
        The gateway address for default routes to reach destination addresses outside this subnetwork.
        """
        return pulumi.get(self, "gateway_address")

    @property
    @pulumi.getter(name="internalIpv6Prefix")
    def internal_ipv6_prefix(self) -> str:
        """
        The internal IPv6 address range that is assigned to this subnetwork.
        """
        return pulumi.get(self, "internal_ipv6_prefix")

    @property
    @pulumi.getter(name="ipCidrRange")
    def ip_cidr_range(self) -> str:
        """
        The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
        """
        return pulumi.get(self, "ip_cidr_range")

    @property
    @pulumi.getter(name="ipv6AccessType")
    def ipv6_access_type(self) -> str:
        """
        The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack.
        """
        return pulumi.get(self, "ipv6_access_type")

    @property
    @pulumi.getter(name="ipv6CidrRange")
    def ipv6_cidr_range(self) -> str:
        """
        This field is for internal use.
        """
        return pulumi.get(self, "ipv6_cidr_range")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#subnetwork for Subnetwork resources.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> 'outputs.SubnetworkLogConfigResponse':
        """
        This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="privateIpGoogleAccess")
    def private_ip_google_access(self) -> bool:
        """
        Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
        """
        return pulumi.get(self, "private_ip_google_access")

    @property
    @pulumi.getter(name="privateIpv6GoogleAccess")
    def private_ipv6_google_access(self) -> str:
        """
        This field is for internal use. This field can be both set at resource creation time and updated using patch.
        """
        return pulumi.get(self, "private_ipv6_google_access")

    @property
    @pulumi.getter
    def purpose(self) -> str:
        """
        The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="reservedInternalRange")
    def reserved_internal_range(self) -> str:
        """
        The URL of the reserved internal range.
        """
        return pulumi.get(self, "reserved_internal_range")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="secondaryIpRanges")
    def secondary_ip_ranges(self) -> Sequence['outputs.SubnetworkSecondaryRangeResponse']:
        """
        An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
        """
        return pulumi.get(self, "secondary_ip_ranges")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="stackType")
    def stack_type(self) -> str:
        """
        The stack type for the subnet. If set to IPV4_ONLY, new VMs in the subnet are assigned IPv4 addresses only. If set to IPV4_IPV6, new VMs in the subnet can be assigned both IPv4 and IPv6 addresses. If not specified, IPV4_ONLY is used. This field can be both set at resource creation time and updated using patch.
        """
        return pulumi.get(self, "stack_type")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
        """
        return pulumi.get(self, "state")


class AwaitableGetSubnetworkResult(GetSubnetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetworkResult(
            allow_subnet_cidr_routes_overlap=self.allow_subnet_cidr_routes_overlap,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            enable_flow_logs=self.enable_flow_logs,
            external_ipv6_prefix=self.external_ipv6_prefix,
            fingerprint=self.fingerprint,
            gateway_address=self.gateway_address,
            internal_ipv6_prefix=self.internal_ipv6_prefix,
            ip_cidr_range=self.ip_cidr_range,
            ipv6_access_type=self.ipv6_access_type,
            ipv6_cidr_range=self.ipv6_cidr_range,
            kind=self.kind,
            log_config=self.log_config,
            name=self.name,
            network=self.network,
            private_ip_google_access=self.private_ip_google_access,
            private_ipv6_google_access=self.private_ipv6_google_access,
            purpose=self.purpose,
            region=self.region,
            reserved_internal_range=self.reserved_internal_range,
            role=self.role,
            secondary_ip_ranges=self.secondary_ip_ranges,
            self_link=self.self_link,
            stack_type=self.stack_type,
            state=self.state)


def get_subnetwork(project: Optional[str] = None,
                   region: Optional[str] = None,
                   subnetwork: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetworkResult:
    """
    Returns the specified subnetwork.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['subnetwork'] = subnetwork
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getSubnetwork', __args__, opts=opts, typ=GetSubnetworkResult).value

    return AwaitableGetSubnetworkResult(
        allow_subnet_cidr_routes_overlap=__ret__.allow_subnet_cidr_routes_overlap,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        enable_flow_logs=__ret__.enable_flow_logs,
        external_ipv6_prefix=__ret__.external_ipv6_prefix,
        fingerprint=__ret__.fingerprint,
        gateway_address=__ret__.gateway_address,
        internal_ipv6_prefix=__ret__.internal_ipv6_prefix,
        ip_cidr_range=__ret__.ip_cidr_range,
        ipv6_access_type=__ret__.ipv6_access_type,
        ipv6_cidr_range=__ret__.ipv6_cidr_range,
        kind=__ret__.kind,
        log_config=__ret__.log_config,
        name=__ret__.name,
        network=__ret__.network,
        private_ip_google_access=__ret__.private_ip_google_access,
        private_ipv6_google_access=__ret__.private_ipv6_google_access,
        purpose=__ret__.purpose,
        region=__ret__.region,
        reserved_internal_range=__ret__.reserved_internal_range,
        role=__ret__.role,
        secondary_ip_ranges=__ret__.secondary_ip_ranges,
        self_link=__ret__.self_link,
        stack_type=__ret__.stack_type,
        state=__ret__.state)


@_utilities.lift_output_func(get_subnetwork)
def get_subnetwork_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                          region: Optional[pulumi.Input[str]] = None,
                          subnetwork: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetworkResult]:
    """
    Returns the specified subnetwork.
    """
    ...
