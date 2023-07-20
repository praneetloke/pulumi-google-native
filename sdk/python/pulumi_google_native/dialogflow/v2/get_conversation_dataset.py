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
    'GetConversationDatasetResult',
    'AwaitableGetConversationDatasetResult',
    'get_conversation_dataset',
    'get_conversation_dataset_output',
]

@pulumi.output_type
class GetConversationDatasetResult:
    def __init__(__self__, conversation_count=None, conversation_info=None, create_time=None, description=None, display_name=None, input_config=None, name=None):
        if conversation_count and not isinstance(conversation_count, str):
            raise TypeError("Expected argument 'conversation_count' to be a str")
        pulumi.set(__self__, "conversation_count", conversation_count)
        if conversation_info and not isinstance(conversation_info, dict):
            raise TypeError("Expected argument 'conversation_info' to be a dict")
        pulumi.set(__self__, "conversation_info", conversation_info)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if input_config and not isinstance(input_config, dict):
            raise TypeError("Expected argument 'input_config' to be a dict")
        pulumi.set(__self__, "input_config", input_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="conversationCount")
    def conversation_count(self) -> str:
        """
        The number of conversations this conversation dataset contains.
        """
        return pulumi.get(self, "conversation_count")

    @property
    @pulumi.getter(name="conversationInfo")
    def conversation_info(self) -> 'outputs.GoogleCloudDialogflowV2ConversationInfoResponse':
        """
        Metadata set during conversation data import.
        """
        return pulumi.get(self, "conversation_info")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of this dataset.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The description of the dataset. Maximum of 10000 bytes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of the dataset. Maximum of 64 bytes.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="inputConfig")
    def input_config(self) -> 'outputs.GoogleCloudDialogflowV2InputConfigResponse':
        """
        Input configurations set during conversation data import.
        """
        return pulumi.get(self, "input_config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        ConversationDataset resource name. Format: `projects//locations//conversationDatasets/`
        """
        return pulumi.get(self, "name")


class AwaitableGetConversationDatasetResult(GetConversationDatasetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConversationDatasetResult(
            conversation_count=self.conversation_count,
            conversation_info=self.conversation_info,
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            input_config=self.input_config,
            name=self.name)


def get_conversation_dataset(conversation_dataset_id: Optional[str] = None,
                             location: Optional[str] = None,
                             project: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConversationDatasetResult:
    """
    Retrieves the specified conversation dataset.
    """
    __args__ = dict()
    __args__['conversationDatasetId'] = conversation_dataset_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v2:getConversationDataset', __args__, opts=opts, typ=GetConversationDatasetResult).value

    return AwaitableGetConversationDatasetResult(
        conversation_count=pulumi.get(__ret__, 'conversation_count'),
        conversation_info=pulumi.get(__ret__, 'conversation_info'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        input_config=pulumi.get(__ret__, 'input_config'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_conversation_dataset)
def get_conversation_dataset_output(conversation_dataset_id: Optional[pulumi.Input[str]] = None,
                                    location: Optional[pulumi.Input[str]] = None,
                                    project: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConversationDatasetResult]:
    """
    Retrieves the specified conversation dataset.
    """
    ...
