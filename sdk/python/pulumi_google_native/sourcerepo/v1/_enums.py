# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'PubsubConfigMessageFormat',
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


class PubsubConfigMessageFormat(str, Enum):
    """
    The format of the Cloud Pub/Sub messages.
    """
    MESSAGE_FORMAT_UNSPECIFIED = "MESSAGE_FORMAT_UNSPECIFIED"
    """
    Unspecified.
    """
    PROTOBUF = "PROTOBUF"
    """
    The message payload is a serialized protocol buffer of SourceRepoEvent.
    """
    JSON = "JSON"
    """
    The message payload is a JSON string of SourceRepoEvent.
    """
