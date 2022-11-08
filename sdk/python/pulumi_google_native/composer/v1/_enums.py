# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'EnvironmentConfigEnvironmentSize',
    'EnvironmentState',
    'NetworkingConfigConnectionType',
]


class EnvironmentConfigEnvironmentSize(str, Enum):
    """
    Optional. The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
    """
    ENVIRONMENT_SIZE_UNSPECIFIED = "ENVIRONMENT_SIZE_UNSPECIFIED"
    """
    The size of the environment is unspecified.
    """
    ENVIRONMENT_SIZE_SMALL = "ENVIRONMENT_SIZE_SMALL"
    """
    The environment size is small.
    """
    ENVIRONMENT_SIZE_MEDIUM = "ENVIRONMENT_SIZE_MEDIUM"
    """
    The environment size is medium.
    """
    ENVIRONMENT_SIZE_LARGE = "ENVIRONMENT_SIZE_LARGE"
    """
    The environment size is large.
    """


class EnvironmentState(str, Enum):
    """
    The current state of the environment.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    The state of the environment is unknown.
    """
    CREATING = "CREATING"
    """
    The environment is in the process of being created.
    """
    RUNNING = "RUNNING"
    """
    The environment is currently running and healthy. It is ready for use.
    """
    UPDATING = "UPDATING"
    """
    The environment is being updated. It remains usable but cannot receive additional update requests or be deleted at this time.
    """
    DELETING = "DELETING"
    """
    The environment is undergoing deletion. It cannot be used.
    """
    ERROR = "ERROR"
    """
    The environment has encountered an error and cannot be used.
    """


class NetworkingConfigConnectionType(str, Enum):
    """
    Optional. Indicates the user requested specifc connection type between Tenant and Customer projects. You cannot set networking connection type in public IP environment.
    """
    CONNECTION_TYPE_UNSPECIFIED = "CONNECTION_TYPE_UNSPECIFIED"
    """
    No specific connection type was requested, so the environment uses the default value corresponding to the rest of its configuration.
    """
    VPC_PEERING = "VPC_PEERING"
    """
    Requests the use of VPC peerings for connecting the Customer and Tenant projects.
    """
    PRIVATE_SERVICE_CONNECT = "PRIVATE_SERVICE_CONNECT"
    """
    Requests the use of Private Service Connect for connecting the Customer and Tenant projects.
    """
