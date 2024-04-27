# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GoogleCloudDocumentaiV1ProcessorVersionAliasResponse',
]

@pulumi.output_type
class GoogleCloudDocumentaiV1ProcessorVersionAliasResponse(dict):
    """
    Contains the alias and the aliased resource name of processor version.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "processorVersion":
            suggest = "processor_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudDocumentaiV1ProcessorVersionAliasResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudDocumentaiV1ProcessorVersionAliasResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudDocumentaiV1ProcessorVersionAliasResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alias: str,
                 processor_version: str):
        """
        Contains the alias and the aliased resource name of processor version.
        :param str alias: The alias in the form of `processor_version` resource name.
        :param str processor_version: The resource name of aliased processor version.
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "processor_version", processor_version)

    @property
    @pulumi.getter
    def alias(self) -> str:
        """
        The alias in the form of `processor_version` resource name.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="processorVersion")
    def processor_version(self) -> str:
        """
        The resource name of aliased processor version.
        """
        return pulumi.get(self, "processor_version")

