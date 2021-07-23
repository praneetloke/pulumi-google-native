# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DeliveryConfigDeliveryRequirement',
]


class DeliveryConfigDeliveryRequirement(str, Enum):
    """
    The DeliveryRequirement for this subscription.
    """
    DELIVERY_REQUIREMENT_UNSPECIFIED = "DELIVERY_REQUIREMENT_UNSPECIFIED"
    """Default value. This value is unused."""
    DELIVER_IMMEDIATELY = "DELIVER_IMMEDIATELY"
    """The server does not wait for a published message to be successfully written to storage before delivering it to subscribers."""
    DELIVER_AFTER_STORED = "DELIVER_AFTER_STORED"
    """The server will not deliver a published message to subscribers until the message has been successfully written to storage. This will result in higher end-to-end latency, but consistent delivery."""