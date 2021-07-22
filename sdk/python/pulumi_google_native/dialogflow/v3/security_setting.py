# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = ['SecuritySettingArgs', 'SecuritySetting']

@pulumi.input_type
class SecuritySettingArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 inspect_template: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purge_data_types: Optional[pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]]] = None,
                 redaction_scope: Optional[pulumi.Input['SecuritySettingRedactionScope']] = None,
                 redaction_strategy: Optional[pulumi.Input['SecuritySettingRedactionStrategy']] = None,
                 retention_window_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SecuritySetting resource.
        :param pulumi.Input[str] display_name: The human-readable name of the security settings, unique within the location.
        :param pulumi.Input[str] inspect_template: [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//inspectTemplates/` OR `projects//locations//inspectTemplates/` OR `organizations//inspectTemplates/`
        :param pulumi.Input[str] name: Resource name of the settings. Format: `projects//locations//securitySettings/`.
        :param pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]] purge_data_types: List of types of data to remove when retention settings triggers purge.
        :param pulumi.Input['SecuritySettingRedactionScope'] redaction_scope: Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
        :param pulumi.Input['SecuritySettingRedactionStrategy'] redaction_strategy: Strategy that defines how we do redaction.
        :param pulumi.Input[int] retention_window_days: Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
        """
        pulumi.set(__self__, "display_name", display_name)
        if inspect_template is not None:
            pulumi.set(__self__, "inspect_template", inspect_template)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if purge_data_types is not None:
            pulumi.set(__self__, "purge_data_types", purge_data_types)
        if redaction_scope is not None:
            pulumi.set(__self__, "redaction_scope", redaction_scope)
        if redaction_strategy is not None:
            pulumi.set(__self__, "redaction_strategy", redaction_strategy)
        if retention_window_days is not None:
            pulumi.set(__self__, "retention_window_days", retention_window_days)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The human-readable name of the security settings, unique within the location.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="inspectTemplate")
    def inspect_template(self) -> Optional[pulumi.Input[str]]:
        """
        [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//inspectTemplates/` OR `projects//locations//inspectTemplates/` OR `organizations//inspectTemplates/`
        """
        return pulumi.get(self, "inspect_template")

    @inspect_template.setter
    def inspect_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inspect_template", value)

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
        Resource name of the settings. Format: `projects//locations//securitySettings/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="purgeDataTypes")
    def purge_data_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]]]:
        """
        List of types of data to remove when retention settings triggers purge.
        """
        return pulumi.get(self, "purge_data_types")

    @purge_data_types.setter
    def purge_data_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]]]):
        pulumi.set(self, "purge_data_types", value)

    @property
    @pulumi.getter(name="redactionScope")
    def redaction_scope(self) -> Optional[pulumi.Input['SecuritySettingRedactionScope']]:
        """
        Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
        """
        return pulumi.get(self, "redaction_scope")

    @redaction_scope.setter
    def redaction_scope(self, value: Optional[pulumi.Input['SecuritySettingRedactionScope']]):
        pulumi.set(self, "redaction_scope", value)

    @property
    @pulumi.getter(name="redactionStrategy")
    def redaction_strategy(self) -> Optional[pulumi.Input['SecuritySettingRedactionStrategy']]:
        """
        Strategy that defines how we do redaction.
        """
        return pulumi.get(self, "redaction_strategy")

    @redaction_strategy.setter
    def redaction_strategy(self, value: Optional[pulumi.Input['SecuritySettingRedactionStrategy']]):
        pulumi.set(self, "redaction_strategy", value)

    @property
    @pulumi.getter(name="retentionWindowDays")
    def retention_window_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
        """
        return pulumi.get(self, "retention_window_days")

    @retention_window_days.setter
    def retention_window_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_window_days", value)


class SecuritySetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_template: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purge_data_types: Optional[pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]]] = None,
                 redaction_scope: Optional[pulumi.Input['SecuritySettingRedactionScope']] = None,
                 redaction_strategy: Optional[pulumi.Input['SecuritySettingRedactionStrategy']] = None,
                 retention_window_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create security settings in the specified location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The human-readable name of the security settings, unique within the location.
        :param pulumi.Input[str] inspect_template: [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//inspectTemplates/` OR `projects//locations//inspectTemplates/` OR `organizations//inspectTemplates/`
        :param pulumi.Input[str] name: Resource name of the settings. Format: `projects//locations//securitySettings/`.
        :param pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]] purge_data_types: List of types of data to remove when retention settings triggers purge.
        :param pulumi.Input['SecuritySettingRedactionScope'] redaction_scope: Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
        :param pulumi.Input['SecuritySettingRedactionStrategy'] redaction_strategy: Strategy that defines how we do redaction.
        :param pulumi.Input[int] retention_window_days: Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecuritySettingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create security settings in the specified location.

        :param str resource_name: The name of the resource.
        :param SecuritySettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecuritySettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_template: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purge_data_types: Optional[pulumi.Input[Sequence[pulumi.Input['SecuritySettingPurgeDataTypesItem']]]] = None,
                 redaction_scope: Optional[pulumi.Input['SecuritySettingRedactionScope']] = None,
                 redaction_strategy: Optional[pulumi.Input['SecuritySettingRedactionStrategy']] = None,
                 retention_window_days: Optional[pulumi.Input[int]] = None,
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
            __props__ = SecuritySettingArgs.__new__(SecuritySettingArgs)

            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["inspect_template"] = inspect_template
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["purge_data_types"] = purge_data_types
            __props__.__dict__["redaction_scope"] = redaction_scope
            __props__.__dict__["redaction_strategy"] = redaction_strategy
            __props__.__dict__["retention_window_days"] = retention_window_days
        super(SecuritySetting, __self__).__init__(
            'google-native:dialogflow/v3:SecuritySetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecuritySetting':
        """
        Get an existing SecuritySetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecuritySettingArgs.__new__(SecuritySettingArgs)

        __props__.__dict__["display_name"] = None
        __props__.__dict__["inspect_template"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["purge_data_types"] = None
        __props__.__dict__["redaction_scope"] = None
        __props__.__dict__["redaction_strategy"] = None
        __props__.__dict__["retention_window_days"] = None
        return SecuritySetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The human-readable name of the security settings, unique within the location.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="inspectTemplate")
    def inspect_template(self) -> pulumi.Output[str]:
        """
        [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//inspectTemplates/` OR `projects//locations//inspectTemplates/` OR `organizations//inspectTemplates/`
        """
        return pulumi.get(self, "inspect_template")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the settings. Format: `projects//locations//securitySettings/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="purgeDataTypes")
    def purge_data_types(self) -> pulumi.Output[Sequence[str]]:
        """
        List of types of data to remove when retention settings triggers purge.
        """
        return pulumi.get(self, "purge_data_types")

    @property
    @pulumi.getter(name="redactionScope")
    def redaction_scope(self) -> pulumi.Output[str]:
        """
        Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
        """
        return pulumi.get(self, "redaction_scope")

    @property
    @pulumi.getter(name="redactionStrategy")
    def redaction_strategy(self) -> pulumi.Output[str]:
        """
        Strategy that defines how we do redaction.
        """
        return pulumi.get(self, "redaction_strategy")

    @property
    @pulumi.getter(name="retentionWindowDays")
    def retention_window_days(self) -> pulumi.Output[int]:
        """
        Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
        """
        return pulumi.get(self, "retention_window_days")

