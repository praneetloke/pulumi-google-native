# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'TagKeyPurpose',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class TagKeyPurpose(str, Enum):
    """
    Optional. A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. A purpose does not grant a policy engine exclusive rights to the Tag, and it may be referenced by other policy engines. A purpose cannot be changed once set.
    """
    PURPOSE_UNSPECIFIED = "PURPOSE_UNSPECIFIED"
    """
    Unspecified purpose.
    """
    GCE_FIREWALL = "GCE_FIREWALL"
    """
    Purpose for Compute Engine firewalls. A corresponding purpose_data should be set for the network the tag is intended for. The key should be 'network' and the value should be in the format of the network url id string: http://compute.googleapis.com/v1/projects/{project_number}/global/networks/{network_id}
    """
