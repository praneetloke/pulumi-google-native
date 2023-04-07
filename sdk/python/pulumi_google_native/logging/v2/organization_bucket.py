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
from ._enums import *
from ._inputs import *

__all__ = ['OrganizationBucketArgs', 'OrganizationBucket']

@pulumi.input_type
class OrganizationBucketArgs:
    def __init__(__self__, *,
                 bucket_id: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 analytics_enabled: Optional[pulumi.Input[bool]] = None,
                 cmek_settings: Optional[pulumi.Input['CmekSettingsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input['IndexConfigArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 restricted_fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a OrganizationBucket resource.
        :param pulumi.Input[str] bucket_id: Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
        :param pulumi.Input[bool] analytics_enabled: Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
        :param pulumi.Input['CmekSettingsArgs'] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[Sequence[pulumi.Input['IndexConfigArgs']]] index_configs: A list of indexed fields and related configuration data.
        :param pulumi.Input[bool] locked: Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] restricted_fields: Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        pulumi.set(__self__, "bucket_id", bucket_id)
        pulumi.set(__self__, "organization_id", organization_id)
        if analytics_enabled is not None:
            pulumi.set(__self__, "analytics_enabled", analytics_enabled)
        if cmek_settings is not None:
            pulumi.set(__self__, "cmek_settings", cmek_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if index_configs is not None:
            pulumi.set(__self__, "index_configs", index_configs)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if restricted_fields is not None:
            pulumi.set(__self__, "restricted_fields", restricted_fields)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Input[str]:
        """
        Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
        """
        return pulumi.get(self, "bucket_id")

    @bucket_id.setter
    def bucket_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_id", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="analyticsEnabled")
    def analytics_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
        """
        return pulumi.get(self, "analytics_enabled")

    @analytics_enabled.setter
    def analytics_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "analytics_enabled", value)

    @property
    @pulumi.getter(name="cmekSettings")
    def cmek_settings(self) -> Optional[pulumi.Input['CmekSettingsArgs']]:
        """
        The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
        """
        return pulumi.get(self, "cmek_settings")

    @cmek_settings.setter
    def cmek_settings(self, value: Optional[pulumi.Input['CmekSettingsArgs']]):
        pulumi.set(self, "cmek_settings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="indexConfigs")
    def index_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IndexConfigArgs']]]]:
        """
        A list of indexed fields and related configuration data.
        """
        return pulumi.get(self, "index_configs")

    @index_configs.setter
    def index_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IndexConfigArgs']]]]):
        pulumi.set(self, "index_configs", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter(name="restrictedFields")
    def restricted_fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
        """
        return pulumi.get(self, "restricted_fields")

    @restricted_fields.setter
    def restricted_fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "restricted_fields", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


class OrganizationBucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analytics_enabled: Optional[pulumi.Input[bool]] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 cmek_settings: Optional[pulumi.Input[pulumi.InputType['CmekSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 restricted_fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a log bucket that can be used to store log entries. After a bucket has been created, the bucket's location cannot be changed.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] analytics_enabled: Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
        :param pulumi.Input[str] bucket_id: Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
        :param pulumi.Input[pulumi.InputType['CmekSettingsArgs']] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexConfigArgs']]]] index_configs: A list of indexed fields and related configuration data.
        :param pulumi.Input[bool] locked: Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] restricted_fields: Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationBucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a log bucket that can be used to store log entries. After a bucket has been created, the bucket's location cannot be changed.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param OrganizationBucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationBucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analytics_enabled: Optional[pulumi.Input[bool]] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 cmek_settings: Optional[pulumi.Input[pulumi.InputType['CmekSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IndexConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 restricted_fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationBucketArgs.__new__(OrganizationBucketArgs)

            __props__.__dict__["analytics_enabled"] = analytics_enabled
            if bucket_id is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_id'")
            __props__.__dict__["bucket_id"] = bucket_id
            __props__.__dict__["cmek_settings"] = cmek_settings
            __props__.__dict__["description"] = description
            __props__.__dict__["index_configs"] = index_configs
            __props__.__dict__["location"] = location
            __props__.__dict__["locked"] = locked
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["restricted_fields"] = restricted_fields
            __props__.__dict__["retention_days"] = retention_days
            __props__.__dict__["create_time"] = None
            __props__.__dict__["lifecycle_state"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["bucket_id", "location", "organization_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(OrganizationBucket, __self__).__init__(
            'google-native:logging/v2:OrganizationBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationBucket':
        """
        Get an existing OrganizationBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationBucketArgs.__new__(OrganizationBucketArgs)

        __props__.__dict__["analytics_enabled"] = None
        __props__.__dict__["bucket_id"] = None
        __props__.__dict__["cmek_settings"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["index_configs"] = None
        __props__.__dict__["lifecycle_state"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["locked"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization_id"] = None
        __props__.__dict__["restricted_fields"] = None
        __props__.__dict__["retention_days"] = None
        __props__.__dict__["update_time"] = None
        return OrganizationBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="analyticsEnabled")
    def analytics_enabled(self) -> pulumi.Output[bool]:
        """
        Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
        """
        return pulumi.get(self, "analytics_enabled")

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Output[str]:
        """
        Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
        """
        return pulumi.get(self, "bucket_id")

    @property
    @pulumi.getter(name="cmekSettings")
    def cmek_settings(self) -> pulumi.Output['outputs.CmekSettingsResponse']:
        """
        The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
        """
        return pulumi.get(self, "cmek_settings")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of the bucket. This is not set for any of the default buckets.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="indexConfigs")
    def index_configs(self) -> pulumi.Output[Sequence['outputs.IndexConfigResponse']]:
        """
        A list of indexed fields and related configuration data.
        """
        return pulumi.get(self, "index_configs")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> pulumi.Output[str]:
        """
        The bucket lifecycle state.
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def locked(self) -> pulumi.Output[bool]:
        """
        Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="restrictedFields")
    def restricted_fields(self) -> pulumi.Output[Sequence[str]]:
        """
        Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
        """
        return pulumi.get(self, "restricted_fields")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[int]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of the bucket.
        """
        return pulumi.get(self, "update_time")

