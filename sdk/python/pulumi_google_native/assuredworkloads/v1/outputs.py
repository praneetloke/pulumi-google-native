# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse',
    'GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse',
    'GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse',
]

@pulumi.output_type
class GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse(dict):
    """
    Settings specific to the Key Management Service.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nextRotationTime":
            suggest = "next_rotation_time"
        elif key == "rotationPeriod":
            suggest = "rotation_period"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 next_rotation_time: str,
                 rotation_period: str):
        """
        Settings specific to the Key Management Service.
        :param str next_rotation_time: Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
        :param str rotation_period: Required. Input only. Immutable. [next_rotation_time] will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
        """
        pulumi.set(__self__, "next_rotation_time", next_rotation_time)
        pulumi.set(__self__, "rotation_period", rotation_period)

    @property
    @pulumi.getter(name="nextRotationTime")
    def next_rotation_time(self) -> str:
        """
        Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
        """
        return pulumi.get(self, "next_rotation_time")

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> str:
        """
        Required. Input only. Immutable. [next_rotation_time] will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
        """
        return pulumi.get(self, "rotation_period")


@pulumi.output_type
class GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse(dict):
    """
    Represent the resources that are children of this Workload.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceId":
            suggest = "resource_id"
        elif key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 resource_id: str,
                 resource_type: str):
        """
        Represent the resources that are children of this Workload.
        :param str resource_id: Resource identifier. For a project this represents project_number.
        :param str resource_type: Indicates the type of resource.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        Resource identifier. For a project this represents project_number.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        Indicates the type of resource.
        """
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse(dict):
    """
    Represent the custom settings for the resources to be created.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceId":
            suggest = "resource_id"
        elif key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 resource_id: str,
                 resource_type: str):
        """
        Represent the custom settings for the resources to be created.
        :param str resource_id: Resource identifier. For a project this represents project_id. If the project is already taken, the workload creation will fail.
        :param str resource_type: Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT)
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        Resource identifier. For a project this represents project_id. If the project is already taken, the workload creation will fail.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT)
        """
        return pulumi.get(self, "resource_type")


