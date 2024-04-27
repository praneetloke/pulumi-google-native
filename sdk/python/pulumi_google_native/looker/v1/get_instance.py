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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, admin_settings=None, consumer_network=None, create_time=None, custom_domain=None, deny_maintenance_period=None, egress_public_ip=None, encryption_config=None, ingress_private_ip=None, ingress_public_ip=None, last_deny_maintenance_period=None, looker_uri=None, looker_version=None, maintenance_schedule=None, maintenance_window=None, name=None, oauth_config=None, platform_edition=None, private_ip_enabled=None, public_ip_enabled=None, reserved_range=None, state=None, update_time=None, user_metadata=None):
        if admin_settings and not isinstance(admin_settings, dict):
            raise TypeError("Expected argument 'admin_settings' to be a dict")
        pulumi.set(__self__, "admin_settings", admin_settings)
        if consumer_network and not isinstance(consumer_network, str):
            raise TypeError("Expected argument 'consumer_network' to be a str")
        pulumi.set(__self__, "consumer_network", consumer_network)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if custom_domain and not isinstance(custom_domain, dict):
            raise TypeError("Expected argument 'custom_domain' to be a dict")
        pulumi.set(__self__, "custom_domain", custom_domain)
        if deny_maintenance_period and not isinstance(deny_maintenance_period, dict):
            raise TypeError("Expected argument 'deny_maintenance_period' to be a dict")
        pulumi.set(__self__, "deny_maintenance_period", deny_maintenance_period)
        if egress_public_ip and not isinstance(egress_public_ip, str):
            raise TypeError("Expected argument 'egress_public_ip' to be a str")
        pulumi.set(__self__, "egress_public_ip", egress_public_ip)
        if encryption_config and not isinstance(encryption_config, dict):
            raise TypeError("Expected argument 'encryption_config' to be a dict")
        pulumi.set(__self__, "encryption_config", encryption_config)
        if ingress_private_ip and not isinstance(ingress_private_ip, str):
            raise TypeError("Expected argument 'ingress_private_ip' to be a str")
        pulumi.set(__self__, "ingress_private_ip", ingress_private_ip)
        if ingress_public_ip and not isinstance(ingress_public_ip, str):
            raise TypeError("Expected argument 'ingress_public_ip' to be a str")
        pulumi.set(__self__, "ingress_public_ip", ingress_public_ip)
        if last_deny_maintenance_period and not isinstance(last_deny_maintenance_period, dict):
            raise TypeError("Expected argument 'last_deny_maintenance_period' to be a dict")
        pulumi.set(__self__, "last_deny_maintenance_period", last_deny_maintenance_period)
        if looker_uri and not isinstance(looker_uri, str):
            raise TypeError("Expected argument 'looker_uri' to be a str")
        pulumi.set(__self__, "looker_uri", looker_uri)
        if looker_version and not isinstance(looker_version, str):
            raise TypeError("Expected argument 'looker_version' to be a str")
        pulumi.set(__self__, "looker_version", looker_version)
        if maintenance_schedule and not isinstance(maintenance_schedule, dict):
            raise TypeError("Expected argument 'maintenance_schedule' to be a dict")
        pulumi.set(__self__, "maintenance_schedule", maintenance_schedule)
        if maintenance_window and not isinstance(maintenance_window, dict):
            raise TypeError("Expected argument 'maintenance_window' to be a dict")
        pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if oauth_config and not isinstance(oauth_config, dict):
            raise TypeError("Expected argument 'oauth_config' to be a dict")
        pulumi.set(__self__, "oauth_config", oauth_config)
        if platform_edition and not isinstance(platform_edition, str):
            raise TypeError("Expected argument 'platform_edition' to be a str")
        pulumi.set(__self__, "platform_edition", platform_edition)
        if private_ip_enabled and not isinstance(private_ip_enabled, bool):
            raise TypeError("Expected argument 'private_ip_enabled' to be a bool")
        pulumi.set(__self__, "private_ip_enabled", private_ip_enabled)
        if public_ip_enabled and not isinstance(public_ip_enabled, bool):
            raise TypeError("Expected argument 'public_ip_enabled' to be a bool")
        pulumi.set(__self__, "public_ip_enabled", public_ip_enabled)
        if reserved_range and not isinstance(reserved_range, str):
            raise TypeError("Expected argument 'reserved_range' to be a str")
        pulumi.set(__self__, "reserved_range", reserved_range)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if user_metadata and not isinstance(user_metadata, dict):
            raise TypeError("Expected argument 'user_metadata' to be a dict")
        pulumi.set(__self__, "user_metadata", user_metadata)

    @property
    @pulumi.getter(name="adminSettings")
    def admin_settings(self) -> 'outputs.AdminSettingsResponse':
        """
        Looker Instance Admin settings.
        """
        return pulumi.get(self, "admin_settings")

    @property
    @pulumi.getter(name="consumerNetwork")
    def consumer_network(self) -> str:
        """
        Network name in the consumer project. Format: `projects/{project}/global/networks/{network}`. Note that the consumer network may be in a different GCP project than the consumer project that is hosting the Looker Instance.
        """
        return pulumi.get(self, "consumer_network")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the Looker instance provisioning was first requested.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="customDomain")
    def custom_domain(self) -> 'outputs.CustomDomainResponse':
        return pulumi.get(self, "custom_domain")

    @property
    @pulumi.getter(name="denyMaintenancePeriod")
    def deny_maintenance_period(self) -> 'outputs.DenyMaintenancePeriodResponse':
        """
        Maintenance denial period for this instance.
        """
        return pulumi.get(self, "deny_maintenance_period")

    @property
    @pulumi.getter(name="egressPublicIp")
    def egress_public_ip(self) -> str:
        """
        Public Egress IP (IPv4).
        """
        return pulumi.get(self, "egress_public_ip")

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> 'outputs.EncryptionConfigResponse':
        """
        Encryption configuration (CMEK). Only set if CMEK has been enabled on the instance.
        """
        return pulumi.get(self, "encryption_config")

    @property
    @pulumi.getter(name="ingressPrivateIp")
    def ingress_private_ip(self) -> str:
        """
        Private Ingress IP (IPv4).
        """
        return pulumi.get(self, "ingress_private_ip")

    @property
    @pulumi.getter(name="ingressPublicIp")
    def ingress_public_ip(self) -> str:
        """
        Public Ingress IP (IPv4).
        """
        return pulumi.get(self, "ingress_public_ip")

    @property
    @pulumi.getter(name="lastDenyMaintenancePeriod")
    def last_deny_maintenance_period(self) -> 'outputs.DenyMaintenancePeriodResponse':
        """
        Last computed maintenance denial period for this instance.
        """
        return pulumi.get(self, "last_deny_maintenance_period")

    @property
    @pulumi.getter(name="lookerUri")
    def looker_uri(self) -> str:
        """
        Looker instance URI which can be used to access the Looker Instance UI.
        """
        return pulumi.get(self, "looker_uri")

    @property
    @pulumi.getter(name="lookerVersion")
    def looker_version(self) -> str:
        """
        The Looker version that the instance is using.
        """
        return pulumi.get(self, "looker_version")

    @property
    @pulumi.getter(name="maintenanceSchedule")
    def maintenance_schedule(self) -> 'outputs.MaintenanceScheduleResponse':
        """
        Maintenance schedule for this instance.
        """
        return pulumi.get(self, "maintenance_schedule")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> 'outputs.MaintenanceWindowResponse':
        """
        Maintenance window for this instance.
        """
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Format: `projects/{project}/locations/{location}/instances/{instance}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oauthConfig")
    def oauth_config(self) -> 'outputs.OAuthConfigResponse':
        """
        Looker instance OAuth login settings.
        """
        return pulumi.get(self, "oauth_config")

    @property
    @pulumi.getter(name="platformEdition")
    def platform_edition(self) -> str:
        """
        Platform edition.
        """
        return pulumi.get(self, "platform_edition")

    @property
    @pulumi.getter(name="privateIpEnabled")
    def private_ip_enabled(self) -> bool:
        """
        Whether private IP is enabled on the Looker instance.
        """
        return pulumi.get(self, "private_ip_enabled")

    @property
    @pulumi.getter(name="publicIpEnabled")
    def public_ip_enabled(self) -> bool:
        """
        Whether public IP is enabled on the Looker instance.
        """
        return pulumi.get(self, "public_ip_enabled")

    @property
    @pulumi.getter(name="reservedRange")
    def reserved_range(self) -> str:
        """
        Name of a reserved IP address range within the Instance.consumer_network, to be used for private services access connection. May or may not be specified in a create request.
        """
        return pulumi.get(self, "reserved_range")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the Looker instance was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> 'outputs.UserMetadataResponse':
        """
        User metadata.
        """
        return pulumi.get(self, "user_metadata")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            admin_settings=self.admin_settings,
            consumer_network=self.consumer_network,
            create_time=self.create_time,
            custom_domain=self.custom_domain,
            deny_maintenance_period=self.deny_maintenance_period,
            egress_public_ip=self.egress_public_ip,
            encryption_config=self.encryption_config,
            ingress_private_ip=self.ingress_private_ip,
            ingress_public_ip=self.ingress_public_ip,
            last_deny_maintenance_period=self.last_deny_maintenance_period,
            looker_uri=self.looker_uri,
            looker_version=self.looker_version,
            maintenance_schedule=self.maintenance_schedule,
            maintenance_window=self.maintenance_window,
            name=self.name,
            oauth_config=self.oauth_config,
            platform_edition=self.platform_edition,
            private_ip_enabled=self.private_ip_enabled,
            public_ip_enabled=self.public_ip_enabled,
            reserved_range=self.reserved_range,
            state=self.state,
            update_time=self.update_time,
            user_metadata=self.user_metadata)


def get_instance(instance_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets details of a single Instance.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:looker/v1:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        admin_settings=pulumi.get(__ret__, 'admin_settings'),
        consumer_network=pulumi.get(__ret__, 'consumer_network'),
        create_time=pulumi.get(__ret__, 'create_time'),
        custom_domain=pulumi.get(__ret__, 'custom_domain'),
        deny_maintenance_period=pulumi.get(__ret__, 'deny_maintenance_period'),
        egress_public_ip=pulumi.get(__ret__, 'egress_public_ip'),
        encryption_config=pulumi.get(__ret__, 'encryption_config'),
        ingress_private_ip=pulumi.get(__ret__, 'ingress_private_ip'),
        ingress_public_ip=pulumi.get(__ret__, 'ingress_public_ip'),
        last_deny_maintenance_period=pulumi.get(__ret__, 'last_deny_maintenance_period'),
        looker_uri=pulumi.get(__ret__, 'looker_uri'),
        looker_version=pulumi.get(__ret__, 'looker_version'),
        maintenance_schedule=pulumi.get(__ret__, 'maintenance_schedule'),
        maintenance_window=pulumi.get(__ret__, 'maintenance_window'),
        name=pulumi.get(__ret__, 'name'),
        oauth_config=pulumi.get(__ret__, 'oauth_config'),
        platform_edition=pulumi.get(__ret__, 'platform_edition'),
        private_ip_enabled=pulumi.get(__ret__, 'private_ip_enabled'),
        public_ip_enabled=pulumi.get(__ret__, 'public_ip_enabled'),
        reserved_range=pulumi.get(__ret__, 'reserved_range'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'),
        user_metadata=pulumi.get(__ret__, 'user_metadata'))


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets details of a single Instance.
    """
    ...