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
    'FileResponse',
    'MetadataResponse',
    'SourceResponse',
]

@pulumi.output_type
class FileResponse(dict):
    """
    `File` containing source content.
    """
    def __init__(__self__, *,
                 content: str,
                 fingerprint: str,
                 name: str):
        """
        `File` containing source content.
        :param str content: Textual Content.
        :param str fingerprint: Fingerprint (e.g. github sha) associated with the `File`.
        :param str name: File name.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "fingerprint", fingerprint)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        Textual Content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint (e.g. github sha) associated with the `File`.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        File name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class MetadataResponse(dict):
    """
    Metadata for a Ruleset.
    """
    def __init__(__self__, *,
                 services: Sequence[str]):
        """
        Metadata for a Ruleset.
        :param Sequence[str] services: Services that this ruleset has declarations for (e.g., "cloud.firestore"). There may be 0+ of these.
        """
        pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter
    def services(self) -> Sequence[str]:
        """
        Services that this ruleset has declarations for (e.g., "cloud.firestore"). There may be 0+ of these.
        """
        return pulumi.get(self, "services")


@pulumi.output_type
class SourceResponse(dict):
    """
    `Source` is one or more `File` messages comprising a logical set of rules.
    """
    def __init__(__self__, *,
                 files: Sequence['outputs.FileResponse']):
        """
        `Source` is one or more `File` messages comprising a logical set of rules.
        :param Sequence['FileResponse'] files: `File` set constituting the `Source` bundle.
        """
        pulumi.set(__self__, "files", files)

    @property
    @pulumi.getter
    def files(self) -> Sequence['outputs.FileResponse']:
        """
        `File` set constituting the `Source` bundle.
        """
        return pulumi.get(self, "files")


