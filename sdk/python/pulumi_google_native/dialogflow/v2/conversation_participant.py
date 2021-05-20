# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ConversationParticipantArgs', 'ConversationParticipant']

@pulumi.input_type
class ConversationParticipantArgs:
    def __init__(__self__, *,
                 conversation_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 participant_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sip_recording_media_label: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConversationParticipant resource.
        :param pulumi.Input[str] name: Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        :param pulumi.Input[str] role: Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        :param pulumi.Input[str] sip_recording_media_label: Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
        """
        pulumi.set(__self__, "conversation_id", conversation_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "participant_id", participant_id)
        pulumi.set(__self__, "project", project)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if sip_recording_media_label is not None:
            pulumi.set(__self__, "sip_recording_media_label", sip_recording_media_label)

    @property
    @pulumi.getter(name="conversationId")
    def conversation_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "conversation_id")

    @conversation_id.setter
    def conversation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "conversation_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="participantId")
    def participant_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "participant_id")

    @participant_id.setter
    def participant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "participant_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="sipRecordingMediaLabel")
    def sip_recording_media_label(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
        """
        return pulumi.get(self, "sip_recording_media_label")

    @sip_recording_media_label.setter
    def sip_recording_media_label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip_recording_media_label", value)


class ConversationParticipant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conversation_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 participant_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sip_recording_media_label: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new participant in a conversation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        :param pulumi.Input[str] role: Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        :param pulumi.Input[str] sip_recording_media_label: Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConversationParticipantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new participant in a conversation.

        :param str resource_name: The name of the resource.
        :param ConversationParticipantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConversationParticipantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conversation_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 participant_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sip_recording_media_label: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConversationParticipantArgs.__new__(ConversationParticipantArgs)

            if conversation_id is None and not opts.urn:
                raise TypeError("Missing required property 'conversation_id'")
            __props__.__dict__["conversation_id"] = conversation_id
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if participant_id is None and not opts.urn:
                raise TypeError("Missing required property 'participant_id'")
            __props__.__dict__["participant_id"] = participant_id
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["role"] = role
            __props__.__dict__["sip_recording_media_label"] = sip_recording_media_label
        super(ConversationParticipant, __self__).__init__(
            'google-native:dialogflow/v2:ConversationParticipant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConversationParticipant':
        """
        Get an existing ConversationParticipant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConversationParticipantArgs.__new__(ConversationParticipantArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["sip_recording_media_label"] = None
        return ConversationParticipant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="sipRecordingMediaLabel")
    def sip_recording_media_label(self) -> pulumi.Output[str]:
        """
        Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
        """
        return pulumi.get(self, "sip_recording_media_label")

