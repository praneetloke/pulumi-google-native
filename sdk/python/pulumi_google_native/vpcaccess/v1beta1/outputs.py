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
    'SubnetResponse',
]

@pulumi.output_type
class SubnetResponse(dict):
    """
    The subnet in which to house the connector
    """
    def __init__(__self__, *,
                 name: str,
                 project: str):
        """
        The subnet in which to house the connector
        :param str name: Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}
        :param str project: Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
        """
        return pulumi.get(self, "project")


