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
    'GetSfdcChannelResult',
    'AwaitableGetSfdcChannelResult',
    'get_sfdc_channel',
    'get_sfdc_channel_output',
]

@pulumi.output_type
class GetSfdcChannelResult:
    def __init__(__self__, channel_topic=None, create_time=None, delete_time=None, description=None, display_name=None, is_active=None, last_replay_id=None, name=None, update_time=None):
        if channel_topic and not isinstance(channel_topic, str):
            raise TypeError("Expected argument 'channel_topic' to be a str")
        pulumi.set(__self__, "channel_topic", channel_topic)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if last_replay_id and not isinstance(last_replay_id, str):
            raise TypeError("Expected argument 'last_replay_id' to be a str")
        pulumi.set(__self__, "last_replay_id", last_replay_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="channelTopic")
    def channel_topic(self) -> str:
        """
        The Channel topic defined by salesforce once an channel is opened
        """
        return pulumi.get(self, "channel_topic")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the channel is created
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        Time when the channel was deleted. Empty if not deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description for this channel
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Client level unique name/alias to easily reference a channel.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        """
        Indicated if a channel has any active integrations referencing it. Set to false when the channel is created, and set to true if there is any integration published with the channel configured in it.
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="lastReplayId")
    def last_replay_id(self) -> str:
        """
        Last sfdc messsage replay id for channel
        """
        return pulumi.get(self, "last_replay_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the SFDC channel projects/{project}/locations/{location}/sfdcInstances/{sfdc_instance}/sfdcChannels/{sfdc_channel}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time when the channel was last updated
        """
        return pulumi.get(self, "update_time")


class AwaitableGetSfdcChannelResult(GetSfdcChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSfdcChannelResult(
            channel_topic=self.channel_topic,
            create_time=self.create_time,
            delete_time=self.delete_time,
            description=self.description,
            display_name=self.display_name,
            is_active=self.is_active,
            last_replay_id=self.last_replay_id,
            name=self.name,
            update_time=self.update_time)


def get_sfdc_channel(location: Optional[str] = None,
                     product_id: Optional[str] = None,
                     project: Optional[str] = None,
                     sfdc_channel_id: Optional[str] = None,
                     sfdc_instance_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSfdcChannelResult:
    """
    Gets an sfdc channel. If the channel doesn't exist, Code.NOT_FOUND exception will be thrown.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['productId'] = product_id
    __args__['project'] = project
    __args__['sfdcChannelId'] = sfdc_channel_id
    __args__['sfdcInstanceId'] = sfdc_instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:integrations/v1alpha:getSfdcChannel', __args__, opts=opts, typ=GetSfdcChannelResult).value

    return AwaitableGetSfdcChannelResult(
        channel_topic=__ret__.channel_topic,
        create_time=__ret__.create_time,
        delete_time=__ret__.delete_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        is_active=__ret__.is_active,
        last_replay_id=__ret__.last_replay_id,
        name=__ret__.name,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_sfdc_channel)
def get_sfdc_channel_output(location: Optional[pulumi.Input[str]] = None,
                            product_id: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            sfdc_channel_id: Optional[pulumi.Input[str]] = None,
                            sfdc_instance_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSfdcChannelResult]:
    """
    Gets an sfdc channel. If the channel doesn't exist, Code.NOT_FOUND exception will be thrown.
    """
    ...
