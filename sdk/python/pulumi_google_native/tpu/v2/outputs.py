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

__all__ = [
    'AccessConfigResponse',
    'AttachedDiskResponse',
    'NetworkConfigResponse',
    'NetworkEndpointResponse',
    'SchedulingConfigResponse',
    'ServiceAccountResponse',
    'ShieldedInstanceConfigResponse',
    'SymptomResponse',
]

@pulumi.output_type
class AccessConfigResponse(dict):
    """
    An access config attached to the TPU worker.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "externalIp":
            suggest = "external_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 external_ip: str):
        """
        An access config attached to the TPU worker.
        :param str external_ip: An external IP address associated with the TPU worker.
        """
        pulumi.set(__self__, "external_ip", external_ip)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> str:
        """
        An external IP address associated with the TPU worker.
        """
        return pulumi.get(self, "external_ip")


@pulumi.output_type
class AttachedDiskResponse(dict):
    """
    A node-attached disk resource. Next ID: 8;
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceDisk":
            suggest = "source_disk"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AttachedDiskResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AttachedDiskResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AttachedDiskResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mode: str,
                 source_disk: str):
        """
        A node-attached disk resource. Next ID: 8;
        :param str mode: The mode in which to attach this disk. If not specified, the default is READ_WRITE mode. Only applicable to data_disks.
        :param str source_disk: Specifies the full path to an existing disk. For example: "projects/my-project/zones/us-central1-c/disks/my-disk".
        """
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "source_disk", source_disk)

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        The mode in which to attach this disk. If not specified, the default is READ_WRITE mode. Only applicable to data_disks.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="sourceDisk")
    def source_disk(self) -> str:
        """
        Specifies the full path to an existing disk. For example: "projects/my-project/zones/us-central1-c/disks/my-disk".
        """
        return pulumi.get(self, "source_disk")


@pulumi.output_type
class NetworkConfigResponse(dict):
    """
    Network related configurations.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "canIpForward":
            suggest = "can_ip_forward"
        elif key == "enableExternalIps":
            suggest = "enable_external_ips"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 can_ip_forward: bool,
                 enable_external_ips: bool,
                 network: str,
                 subnetwork: str):
        """
        Network related configurations.
        :param bool can_ip_forward: Allows the TPU node to send and receive packets with non-matching destination or source IPs. This is required if you plan to use the TPU workers to forward routes.
        :param bool enable_external_ips: Indicates that external IP addresses would be associated with the TPU workers. If set to false, the specified subnetwork or network should have Private Google Access enabled.
        :param str network: The name of the network for the TPU node. It must be a preexisting Google Compute Engine network. If none is provided, "default" will be used.
        :param str subnetwork: The name of the subnetwork for the TPU node. It must be a preexisting Google Compute Engine subnetwork. If none is provided, "default" will be used.
        """
        pulumi.set(__self__, "can_ip_forward", can_ip_forward)
        pulumi.set(__self__, "enable_external_ips", enable_external_ips)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "subnetwork", subnetwork)

    @property
    @pulumi.getter(name="canIpForward")
    def can_ip_forward(self) -> bool:
        """
        Allows the TPU node to send and receive packets with non-matching destination or source IPs. This is required if you plan to use the TPU workers to forward routes.
        """
        return pulumi.get(self, "can_ip_forward")

    @property
    @pulumi.getter(name="enableExternalIps")
    def enable_external_ips(self) -> bool:
        """
        Indicates that external IP addresses would be associated with the TPU workers. If set to false, the specified subnetwork or network should have Private Google Access enabled.
        """
        return pulumi.get(self, "enable_external_ips")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The name of the network for the TPU node. It must be a preexisting Google Compute Engine network. If none is provided, "default" will be used.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def subnetwork(self) -> str:
        """
        The name of the subnetwork for the TPU node. It must be a preexisting Google Compute Engine subnetwork. If none is provided, "default" will be used.
        """
        return pulumi.get(self, "subnetwork")


@pulumi.output_type
class NetworkEndpointResponse(dict):
    """
    A network endpoint over which a TPU worker can be reached.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessConfig":
            suggest = "access_config"
        elif key == "ipAddress":
            suggest = "ip_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkEndpointResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkEndpointResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkEndpointResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_config: 'outputs.AccessConfigResponse',
                 ip_address: str,
                 port: int):
        """
        A network endpoint over which a TPU worker can be reached.
        :param 'AccessConfigResponse' access_config: The access config for the TPU worker.
        :param str ip_address: The internal IP address of this network endpoint.
        :param int port: The port of this network endpoint.
        """
        pulumi.set(__self__, "access_config", access_config)
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="accessConfig")
    def access_config(self) -> 'outputs.AccessConfigResponse':
        """
        The access config for the TPU worker.
        """
        return pulumi.get(self, "access_config")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The internal IP address of this network endpoint.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port of this network endpoint.
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class SchedulingConfigResponse(dict):
    """
    Sets the scheduling options for this node.
    """
    def __init__(__self__, *,
                 preemptible: bool,
                 reserved: bool):
        """
        Sets the scheduling options for this node.
        :param bool preemptible: Defines whether the node is preemptible.
        :param bool reserved: Whether the node is created under a reservation.
        """
        pulumi.set(__self__, "preemptible", preemptible)
        pulumi.set(__self__, "reserved", reserved)

    @property
    @pulumi.getter
    def preemptible(self) -> bool:
        """
        Defines whether the node is preemptible.
        """
        return pulumi.get(self, "preemptible")

    @property
    @pulumi.getter
    def reserved(self) -> bool:
        """
        Whether the node is created under a reservation.
        """
        return pulumi.get(self, "reserved")


@pulumi.output_type
class ServiceAccountResponse(dict):
    """
    A service account.
    """
    def __init__(__self__, *,
                 email: str,
                 scope: Sequence[str]):
        """
        A service account.
        :param str email: Email address of the service account. If empty, default Compute service account will be used.
        :param Sequence[str] scope: The list of scopes to be made available for this service account. If empty, access to all Cloud APIs will be allowed.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Email address of the service account. If empty, default Compute service account will be used.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def scope(self) -> Sequence[str]:
        """
        The list of scopes to be made available for this service account. If empty, access to all Cloud APIs will be allowed.
        """
        return pulumi.get(self, "scope")


@pulumi.output_type
class ShieldedInstanceConfigResponse(dict):
    """
    A set of Shielded Instance options.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enableSecureBoot":
            suggest = "enable_secure_boot"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ShieldedInstanceConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ShieldedInstanceConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ShieldedInstanceConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enable_secure_boot: bool):
        """
        A set of Shielded Instance options.
        :param bool enable_secure_boot: Defines whether the instance has Secure Boot enabled.
        """
        pulumi.set(__self__, "enable_secure_boot", enable_secure_boot)

    @property
    @pulumi.getter(name="enableSecureBoot")
    def enable_secure_boot(self) -> bool:
        """
        Defines whether the instance has Secure Boot enabled.
        """
        return pulumi.get(self, "enable_secure_boot")


@pulumi.output_type
class SymptomResponse(dict):
    """
    A Symptom instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "symptomType":
            suggest = "symptom_type"
        elif key == "workerId":
            suggest = "worker_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SymptomResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SymptomResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SymptomResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 create_time: str,
                 details: str,
                 symptom_type: str,
                 worker_id: str):
        """
        A Symptom instance.
        :param str create_time: Timestamp when the Symptom is created.
        :param str details: Detailed information of the current Symptom.
        :param str symptom_type: Type of the Symptom.
        :param str worker_id: A string used to uniquely distinguish a worker within a TPU node.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "symptom_type", symptom_type)
        pulumi.set(__self__, "worker_id", worker_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when the Symptom is created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def details(self) -> str:
        """
        Detailed information of the current Symptom.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="symptomType")
    def symptom_type(self) -> str:
        """
        Type of the Symptom.
        """
        return pulumi.get(self, "symptom_type")

    @property
    @pulumi.getter(name="workerId")
    def worker_id(self) -> str:
        """
        A string used to uniquely distinguish a worker within a TPU node.
        """
        return pulumi.get(self, "worker_id")


