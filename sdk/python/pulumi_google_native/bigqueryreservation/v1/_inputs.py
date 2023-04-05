# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'AutoscaleArgs',
]

@pulumi.input_type
class AutoscaleArgs:
    def __init__(__self__, *,
                 max_slots: Optional[pulumi.Input[str]] = None):
        """
        Auto scaling settings.
        :param pulumi.Input[str] max_slots: Number of slots to be scaled when needed.
        """
        if max_slots is not None:
            pulumi.set(__self__, "max_slots", max_slots)

    @property
    @pulumi.getter(name="maxSlots")
    def max_slots(self) -> Optional[pulumi.Input[str]]:
        """
        Number of slots to be scaled when needed.
        """
        return pulumi.get(self, "max_slots")

    @max_slots.setter
    def max_slots(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_slots", value)

