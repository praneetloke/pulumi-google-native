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
    'GetAdaptiveMtDatasetResult',
    'AwaitableGetAdaptiveMtDatasetResult',
    'get_adaptive_mt_dataset',
    'get_adaptive_mt_dataset_output',
]

@pulumi.output_type
class GetAdaptiveMtDatasetResult:
    def __init__(__self__, create_time=None, display_name=None, example_count=None, name=None, source_language_code=None, target_language_code=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if example_count and not isinstance(example_count, int):
            raise TypeError("Expected argument 'example_count' to be a int")
        pulumi.set(__self__, "example_count", example_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if source_language_code and not isinstance(source_language_code, str):
            raise TypeError("Expected argument 'source_language_code' to be a str")
        pulumi.set(__self__, "source_language_code", source_language_code)
        if target_language_code and not isinstance(target_language_code, str):
            raise TypeError("Expected argument 'target_language_code' to be a str")
        pulumi.set(__self__, "target_language_code", target_language_code)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when this dataset was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of the dataset to show in the interface. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores (_), and ASCII digits 0-9.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="exampleCount")
    def example_count(self) -> int:
        """
        The number of examples in the dataset.
        """
        return pulumi.get(self, "example_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the dataset, in form of `projects/{project-number-or-id}/locations/{location_id}/adaptiveMtDatasets/{dataset_id}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sourceLanguageCode")
    def source_language_code(self) -> str:
        """
        The BCP-47 language code of the source language.
        """
        return pulumi.get(self, "source_language_code")

    @property
    @pulumi.getter(name="targetLanguageCode")
    def target_language_code(self) -> str:
        """
        The BCP-47 language code of the target language.
        """
        return pulumi.get(self, "target_language_code")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Timestamp when this dataset was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetAdaptiveMtDatasetResult(GetAdaptiveMtDatasetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAdaptiveMtDatasetResult(
            create_time=self.create_time,
            display_name=self.display_name,
            example_count=self.example_count,
            name=self.name,
            source_language_code=self.source_language_code,
            target_language_code=self.target_language_code,
            update_time=self.update_time)


def get_adaptive_mt_dataset(adaptive_mt_dataset_id: Optional[str] = None,
                            location: Optional[str] = None,
                            project: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAdaptiveMtDatasetResult:
    """
    Gets the Adaptive MT dataset.
    """
    __args__ = dict()
    __args__['adaptiveMtDatasetId'] = adaptive_mt_dataset_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:translate/v3:getAdaptiveMtDataset', __args__, opts=opts, typ=GetAdaptiveMtDatasetResult).value

    return AwaitableGetAdaptiveMtDatasetResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        example_count=pulumi.get(__ret__, 'example_count'),
        name=pulumi.get(__ret__, 'name'),
        source_language_code=pulumi.get(__ret__, 'source_language_code'),
        target_language_code=pulumi.get(__ret__, 'target_language_code'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_adaptive_mt_dataset)
def get_adaptive_mt_dataset_output(adaptive_mt_dataset_id: Optional[pulumi.Input[str]] = None,
                                   location: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAdaptiveMtDatasetResult]:
    """
    Gets the Adaptive MT dataset.
    """
    ...
