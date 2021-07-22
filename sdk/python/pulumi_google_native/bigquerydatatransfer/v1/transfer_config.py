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

__all__ = ['TransferConfigArgs', 'TransferConfig']

@pulumi.input_type
class TransferConfigArgs:
    def __init__(__self__, *,
                 authorization_code: Optional[pulumi.Input[str]] = None,
                 data_refresh_window_days: Optional[pulumi.Input[int]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 destination_dataset_id: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_preferences: Optional[pulumi.Input['EmailPreferencesArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_pubsub_topic: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 schedule_options: Optional[pulumi.Input['ScheduleOptionsArgs']] = None,
                 service_account_name: Optional[pulumi.Input[str]] = None,
                 version_info: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransferConfig resource.
        :param pulumi.Input[int] data_refresh_window_days: The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
        :param pulumi.Input[str] data_source_id: Data source id. Cannot be changed once data transfer is created.
        :param pulumi.Input[str] destination_dataset_id: The BigQuery target dataset id.
        :param pulumi.Input[bool] disabled: Is this config disabled. When set to true, no runs are scheduled for a given transfer.
        :param pulumi.Input[str] display_name: User specified display name for the data transfer.
        :param pulumi.Input['EmailPreferencesArgs'] email_preferences: Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
        :param pulumi.Input[str] name: The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
        :param pulumi.Input[str] notification_pubsub_topic: Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] params: Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
        :param pulumi.Input[str] schedule: Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
        :param pulumi.Input['ScheduleOptionsArgs'] schedule_options: Options customizing the data transfer schedule.
        """
        if authorization_code is not None:
            pulumi.set(__self__, "authorization_code", authorization_code)
        if data_refresh_window_days is not None:
            pulumi.set(__self__, "data_refresh_window_days", data_refresh_window_days)
        if data_source_id is not None:
            pulumi.set(__self__, "data_source_id", data_source_id)
        if destination_dataset_id is not None:
            pulumi.set(__self__, "destination_dataset_id", destination_dataset_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email_preferences is not None:
            pulumi.set(__self__, "email_preferences", email_preferences)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_pubsub_topic is not None:
            pulumi.set(__self__, "notification_pubsub_topic", notification_pubsub_topic)
        if params is not None:
            pulumi.set(__self__, "params", params)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if schedule_options is not None:
            pulumi.set(__self__, "schedule_options", schedule_options)
        if service_account_name is not None:
            pulumi.set(__self__, "service_account_name", service_account_name)
        if version_info is not None:
            pulumi.set(__self__, "version_info", version_info)

    @property
    @pulumi.getter(name="authorizationCode")
    def authorization_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorization_code")

    @authorization_code.setter
    def authorization_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_code", value)

    @property
    @pulumi.getter(name="dataRefreshWindowDays")
    def data_refresh_window_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
        """
        return pulumi.get(self, "data_refresh_window_days")

    @data_refresh_window_days.setter
    def data_refresh_window_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_refresh_window_days", value)

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> Optional[pulumi.Input[str]]:
        """
        Data source id. Cannot be changed once data transfer is created.
        """
        return pulumi.get(self, "data_source_id")

    @data_source_id.setter
    def data_source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_id", value)

    @property
    @pulumi.getter(name="destinationDatasetId")
    def destination_dataset_id(self) -> Optional[pulumi.Input[str]]:
        """
        The BigQuery target dataset id.
        """
        return pulumi.get(self, "destination_dataset_id")

    @destination_dataset_id.setter
    def destination_dataset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_dataset_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this config disabled. When set to true, no runs are scheduled for a given transfer.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User specified display name for the data transfer.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="emailPreferences")
    def email_preferences(self) -> Optional[pulumi.Input['EmailPreferencesArgs']]:
        """
        Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
        """
        return pulumi.get(self, "email_preferences")

    @email_preferences.setter
    def email_preferences(self, value: Optional[pulumi.Input['EmailPreferencesArgs']]):
        pulumi.set(self, "email_preferences", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationPubsubTopic")
    def notification_pubsub_topic(self) -> Optional[pulumi.Input[str]]:
        """
        Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
        """
        return pulumi.get(self, "notification_pubsub_topic")

    @notification_pubsub_topic.setter
    def notification_pubsub_topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_pubsub_topic", value)

    @property
    @pulumi.getter
    def params(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
        """
        return pulumi.get(self, "params")

    @params.setter
    def params(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "params", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="scheduleOptions")
    def schedule_options(self) -> Optional[pulumi.Input['ScheduleOptionsArgs']]:
        """
        Options customizing the data transfer schedule.
        """
        return pulumi.get(self, "schedule_options")

    @schedule_options.setter
    def schedule_options(self, value: Optional[pulumi.Input['ScheduleOptionsArgs']]):
        pulumi.set(self, "schedule_options", value)

    @property
    @pulumi.getter(name="serviceAccountName")
    def service_account_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_account_name")

    @service_account_name.setter
    def service_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_name", value)

    @property
    @pulumi.getter(name="versionInfo")
    def version_info(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version_info")

    @version_info.setter
    def version_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_info", value)


class TransferConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_code: Optional[pulumi.Input[str]] = None,
                 data_refresh_window_days: Optional[pulumi.Input[int]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 destination_dataset_id: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_preferences: Optional[pulumi.Input[pulumi.InputType['EmailPreferencesArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_pubsub_topic: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 schedule_options: Optional[pulumi.Input[pulumi.InputType['ScheduleOptionsArgs']]] = None,
                 service_account_name: Optional[pulumi.Input[str]] = None,
                 version_info: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new data transfer configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_refresh_window_days: The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
        :param pulumi.Input[str] data_source_id: Data source id. Cannot be changed once data transfer is created.
        :param pulumi.Input[str] destination_dataset_id: The BigQuery target dataset id.
        :param pulumi.Input[bool] disabled: Is this config disabled. When set to true, no runs are scheduled for a given transfer.
        :param pulumi.Input[str] display_name: User specified display name for the data transfer.
        :param pulumi.Input[pulumi.InputType['EmailPreferencesArgs']] email_preferences: Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
        :param pulumi.Input[str] name: The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
        :param pulumi.Input[str] notification_pubsub_topic: Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] params: Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
        :param pulumi.Input[str] schedule: Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
        :param pulumi.Input[pulumi.InputType['ScheduleOptionsArgs']] schedule_options: Options customizing the data transfer schedule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TransferConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new data transfer configuration.

        :param str resource_name: The name of the resource.
        :param TransferConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransferConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_code: Optional[pulumi.Input[str]] = None,
                 data_refresh_window_days: Optional[pulumi.Input[int]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 destination_dataset_id: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_preferences: Optional[pulumi.Input[pulumi.InputType['EmailPreferencesArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_pubsub_topic: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 schedule_options: Optional[pulumi.Input[pulumi.InputType['ScheduleOptionsArgs']]] = None,
                 service_account_name: Optional[pulumi.Input[str]] = None,
                 version_info: Optional[pulumi.Input[str]] = None,
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
            __props__ = TransferConfigArgs.__new__(TransferConfigArgs)

            __props__.__dict__["authorization_code"] = authorization_code
            __props__.__dict__["data_refresh_window_days"] = data_refresh_window_days
            __props__.__dict__["data_source_id"] = data_source_id
            __props__.__dict__["destination_dataset_id"] = destination_dataset_id
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["email_preferences"] = email_preferences
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_pubsub_topic"] = notification_pubsub_topic
            __props__.__dict__["params"] = params
            __props__.__dict__["project"] = project
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["schedule_options"] = schedule_options
            __props__.__dict__["service_account_name"] = service_account_name
            __props__.__dict__["version_info"] = version_info
            __props__.__dict__["dataset_region"] = None
            __props__.__dict__["next_run_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(TransferConfig, __self__).__init__(
            'google-native:bigquerydatatransfer/v1:TransferConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TransferConfig':
        """
        Get an existing TransferConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TransferConfigArgs.__new__(TransferConfigArgs)

        __props__.__dict__["data_refresh_window_days"] = None
        __props__.__dict__["data_source_id"] = None
        __props__.__dict__["dataset_region"] = None
        __props__.__dict__["destination_dataset_id"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["email_preferences"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["next_run_time"] = None
        __props__.__dict__["notification_pubsub_topic"] = None
        __props__.__dict__["params"] = None
        __props__.__dict__["schedule"] = None
        __props__.__dict__["schedule_options"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return TransferConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataRefreshWindowDays")
    def data_refresh_window_days(self) -> pulumi.Output[int]:
        """
        The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
        """
        return pulumi.get(self, "data_refresh_window_days")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> pulumi.Output[str]:
        """
        Data source id. Cannot be changed once data transfer is created.
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter(name="datasetRegion")
    def dataset_region(self) -> pulumi.Output[str]:
        """
        Region in which BigQuery dataset is located.
        """
        return pulumi.get(self, "dataset_region")

    @property
    @pulumi.getter(name="destinationDatasetId")
    def destination_dataset_id(self) -> pulumi.Output[str]:
        """
        The BigQuery target dataset id.
        """
        return pulumi.get(self, "destination_dataset_id")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Is this config disabled. When set to true, no runs are scheduled for a given transfer.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        User specified display name for the data transfer.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="emailPreferences")
    def email_preferences(self) -> pulumi.Output['outputs.EmailPreferencesResponse']:
        """
        Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
        """
        return pulumi.get(self, "email_preferences")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextRunTime")
    def next_run_time(self) -> pulumi.Output[str]:
        """
        Next time when data transfer will run.
        """
        return pulumi.get(self, "next_run_time")

    @property
    @pulumi.getter(name="notificationPubsubTopic")
    def notification_pubsub_topic(self) -> pulumi.Output[str]:
        """
        Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
        """
        return pulumi.get(self, "notification_pubsub_topic")

    @property
    @pulumi.getter
    def params(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
        """
        return pulumi.get(self, "params")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: the granularity should be at least 8 hours, or less frequent.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="scheduleOptions")
    def schedule_options(self) -> pulumi.Output['outputs.ScheduleOptionsResponse']:
        """
        Options customizing the data transfer schedule.
        """
        return pulumi.get(self, "schedule_options")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the most recently updated transfer run.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Data transfer modification time. Ignored by server on input.
        """
        return pulumi.get(self, "update_time")

