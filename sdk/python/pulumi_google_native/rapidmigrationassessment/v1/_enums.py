# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AnnotationType',
]


class AnnotationType(str, Enum):
    """
    Type of an annotation.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Unknown type
    """
    TYPE_LEGACY_EXPORT_CONSENT = "TYPE_LEGACY_EXPORT_CONSENT"
    """
    Indicates that this project has opted into StratoZone export.
    """
    TYPE_QWIKLAB = "TYPE_QWIKLAB"
    """
    Indicates that this project is created by Qwiklab.
    """