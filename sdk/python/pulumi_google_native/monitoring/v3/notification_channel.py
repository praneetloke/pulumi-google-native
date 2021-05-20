# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NotificationChannelArgs', 'NotificationChannel']

@pulumi.input_type
class NotificationChannelArgs:
    def __init__(__self__, *,
                 notification_channel_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 creation_record: Optional[pulumi.Input['MutationRecordArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mutation_records: Optional[pulumi.Input[Sequence[pulumi.Input['MutationRecordArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 verification_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NotificationChannel resource.
        :param pulumi.Input['MutationRecordArgs'] creation_record: Record of the creation of this channel.
        :param pulumi.Input[str] description: An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        :param pulumi.Input[str] display_name: An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        :param pulumi.Input[bool] enabled: Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
        :param pulumi.Input[Sequence[pulumi.Input['MutationRecordArgs']]] mutation_records: Records of the modification of this channel.
        :param pulumi.Input[str] name: The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
        :param pulumi.Input[str] type: The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        :param pulumi.Input[str] verification_status: Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
        """
        pulumi.set(__self__, "notification_channel_id", notification_channel_id)
        pulumi.set(__self__, "project", project)
        if creation_record is not None:
            pulumi.set(__self__, "creation_record", creation_record)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if mutation_records is not None:
            pulumi.set(__self__, "mutation_records", mutation_records)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_labels is not None:
            pulumi.set(__self__, "user_labels", user_labels)
        if verification_status is not None:
            pulumi.set(__self__, "verification_status", verification_status)

    @property
    @pulumi.getter(name="notificationChannelId")
    def notification_channel_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "notification_channel_id")

    @notification_channel_id.setter
    def notification_channel_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "notification_channel_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="creationRecord")
    def creation_record(self) -> Optional[pulumi.Input['MutationRecordArgs']]:
        """
        Record of the creation of this channel.
        """
        return pulumi.get(self, "creation_record")

    @creation_record.setter
    def creation_record(self, value: Optional[pulumi.Input['MutationRecordArgs']]):
        pulumi.set(self, "creation_record", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="mutationRecords")
    def mutation_records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MutationRecordArgs']]]]:
        """
        Records of the modification of this channel.
        """
        return pulumi.get(self, "mutation_records")

    @mutation_records.setter
    def mutation_records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MutationRecordArgs']]]]):
        pulumi.set(self, "mutation_records", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        """
        return pulumi.get(self, "user_labels")

    @user_labels.setter
    def user_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "user_labels", value)

    @property
    @pulumi.getter(name="verificationStatus")
    def verification_status(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
        """
        return pulumi.get(self, "verification_status")

    @verification_status.setter
    def verification_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verification_status", value)


class NotificationChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mutation_records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MutationRecordArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channel_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 verification_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new notification channel, representing a single notification endpoint such as an email address, SMS number, or PagerDuty service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MutationRecordArgs']] creation_record: Record of the creation of this channel.
        :param pulumi.Input[str] description: An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        :param pulumi.Input[str] display_name: An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        :param pulumi.Input[bool] enabled: Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MutationRecordArgs']]]] mutation_records: Records of the modification of this channel.
        :param pulumi.Input[str] name: The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
        :param pulumi.Input[str] type: The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        :param pulumi.Input[str] verification_status: Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new notification channel, representing a single notification endpoint such as an email address, SMS number, or PagerDuty service.

        :param str resource_name: The name of the resource.
        :param NotificationChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mutation_records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MutationRecordArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channel_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 verification_status: Optional[pulumi.Input[str]] = None,
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
            __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

            __props__.__dict__["creation_record"] = creation_record
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["labels"] = labels
            __props__.__dict__["mutation_records"] = mutation_records
            __props__.__dict__["name"] = name
            if notification_channel_id is None and not opts.urn:
                raise TypeError("Missing required property 'notification_channel_id'")
            __props__.__dict__["notification_channel_id"] = notification_channel_id
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["type"] = type
            __props__.__dict__["user_labels"] = user_labels
            __props__.__dict__["verification_status"] = verification_status
        super(NotificationChannel, __self__).__init__(
            'google-native:monitoring/v3:NotificationChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NotificationChannel':
        """
        Get an existing NotificationChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

        __props__.__dict__["creation_record"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["mutation_records"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["user_labels"] = None
        __props__.__dict__["verification_status"] = None
        return NotificationChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationRecord")
    def creation_record(self) -> pulumi.Output['outputs.MutationRecordResponse']:
        """
        Record of the creation of this channel.
        """
        return pulumi.get(self, "creation_record")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor.labels of the NotificationChannelDescriptor corresponding to the type field.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mutationRecords")
    def mutation_records(self) -> pulumi.Output[Sequence['outputs.MutationRecordResponse']]:
        """
        Records of the modification of this channel.
        """
        return pulumi.get(self, "mutation_records")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full REST resource name for this channel. The format is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        """
        return pulumi.get(self, "user_labels")

    @property
    @pulumi.getter(name="verificationStatus")
    def verification_status(self) -> pulumi.Output[str]:
        """
        Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
        """
        return pulumi.get(self, "verification_status")

