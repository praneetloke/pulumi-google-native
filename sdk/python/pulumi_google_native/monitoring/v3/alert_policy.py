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

__all__ = ['AlertPolicyArgs', 'AlertPolicy']

@pulumi.input_type
class AlertPolicyArgs:
    def __init__(__self__, *,
                 alert_strategy: Optional[pulumi.Input['AlertStrategyArgs']] = None,
                 combiner: Optional[pulumi.Input['AlertPolicyCombiner']] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionArgs']]]] = None,
                 creation_record: Optional[pulumi.Input['MutationRecordArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 documentation: Optional[pulumi.Input['DocumentationArgs']] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mutation_record: Optional[pulumi.Input['MutationRecordArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input['AlertPolicySeverity']] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 validity: Optional[pulumi.Input['StatusArgs']] = None):
        """
        The set of arguments for constructing a AlertPolicy resource.
        :param pulumi.Input['AlertStrategyArgs'] alert_strategy: Control over how this alert policy's notification channels are notified.
        :param pulumi.Input['AlertPolicyCombiner'] combiner: How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        :param pulumi.Input[Sequence[pulumi.Input['ConditionArgs']]] conditions: A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition. If condition_monitoring_query_language is present, it must be the only condition.
        :param pulumi.Input['MutationRecordArgs'] creation_record: A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.The convention for the display_name of a PrometheusQueryLanguageCondition is "{rule group name}/{alert name}", where the {rule group name} and {alert name} should be taken from the corresponding Prometheus configuration file. This convention is not enforced. In any case the display_name is not a unique key of the AlertPolicy.
        :param pulumi.Input['DocumentationArgs'] documentation: Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        :param pulumi.Input['MutationRecordArgs'] mutation_record: A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        :param pulumi.Input[str] name: Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Cloud Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_channels: Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        :param pulumi.Input['AlertPolicySeverity'] severity: Optional. The severity of an alert policy indicates how important incidents generated by that policy are. The severity level will be displayed on the Incident detail page and in notifications.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.Note that Prometheus {alert name} is a valid Prometheus label names (https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels), whereas Prometheus {rule group} is an unrestricted UTF-8 string. This means that they cannot be stored as-is in user labels, because they may contain characters that are not allowed in user-label values.
        :param pulumi.Input['StatusArgs'] validity: Read-only description of how the alert policy is invalid. This field is only set when the alert policy is invalid. An invalid alert policy will not generate incidents.
        """
        if alert_strategy is not None:
            pulumi.set(__self__, "alert_strategy", alert_strategy)
        if combiner is not None:
            pulumi.set(__self__, "combiner", combiner)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if creation_record is not None:
            pulumi.set(__self__, "creation_record", creation_record)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if documentation is not None:
            pulumi.set(__self__, "documentation", documentation)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if mutation_record is not None:
            pulumi.set(__self__, "mutation_record", mutation_record)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_channels is not None:
            pulumi.set(__self__, "notification_channels", notification_channels)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if user_labels is not None:
            pulumi.set(__self__, "user_labels", user_labels)
        if validity is not None:
            pulumi.set(__self__, "validity", validity)

    @property
    @pulumi.getter(name="alertStrategy")
    def alert_strategy(self) -> Optional[pulumi.Input['AlertStrategyArgs']]:
        """
        Control over how this alert policy's notification channels are notified.
        """
        return pulumi.get(self, "alert_strategy")

    @alert_strategy.setter
    def alert_strategy(self, value: Optional[pulumi.Input['AlertStrategyArgs']]):
        pulumi.set(self, "alert_strategy", value)

    @property
    @pulumi.getter
    def combiner(self) -> Optional[pulumi.Input['AlertPolicyCombiner']]:
        """
        How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        """
        return pulumi.get(self, "combiner")

    @combiner.setter
    def combiner(self, value: Optional[pulumi.Input['AlertPolicyCombiner']]):
        pulumi.set(self, "combiner", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConditionArgs']]]]:
        """
        A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition. If condition_monitoring_query_language is present, it must be the only condition.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="creationRecord")
    def creation_record(self) -> Optional[pulumi.Input['MutationRecordArgs']]:
        """
        A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        """
        return pulumi.get(self, "creation_record")

    @creation_record.setter
    def creation_record(self, value: Optional[pulumi.Input['MutationRecordArgs']]):
        pulumi.set(self, "creation_record", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.The convention for the display_name of a PrometheusQueryLanguageCondition is "{rule group name}/{alert name}", where the {rule group name} and {alert name} should be taken from the corresponding Prometheus configuration file. This convention is not enforced. In any case the display_name is not a unique key of the AlertPolicy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def documentation(self) -> Optional[pulumi.Input['DocumentationArgs']]:
        """
        Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        """
        return pulumi.get(self, "documentation")

    @documentation.setter
    def documentation(self, value: Optional[pulumi.Input['DocumentationArgs']]):
        pulumi.set(self, "documentation", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="mutationRecord")
    def mutation_record(self) -> Optional[pulumi.Input['MutationRecordArgs']]:
        """
        A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        """
        return pulumi.get(self, "mutation_record")

    @mutation_record.setter
    def mutation_record(self, value: Optional[pulumi.Input['MutationRecordArgs']]):
        pulumi.set(self, "mutation_record", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Cloud Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        """
        return pulumi.get(self, "notification_channels")

    @notification_channels.setter
    def notification_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notification_channels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input['AlertPolicySeverity']]:
        """
        Optional. The severity of an alert policy indicates how important incidents generated by that policy are. The severity level will be displayed on the Incident detail page and in notifications.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input['AlertPolicySeverity']]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.Note that Prometheus {alert name} is a valid Prometheus label names (https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels), whereas Prometheus {rule group} is an unrestricted UTF-8 string. This means that they cannot be stored as-is in user labels, because they may contain characters that are not allowed in user-label values.
        """
        return pulumi.get(self, "user_labels")

    @user_labels.setter
    def user_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "user_labels", value)

    @property
    @pulumi.getter
    def validity(self) -> Optional[pulumi.Input['StatusArgs']]:
        """
        Read-only description of how the alert policy is invalid. This field is only set when the alert policy is invalid. An invalid alert policy will not generate incidents.
        """
        return pulumi.get(self, "validity")

    @validity.setter
    def validity(self, value: Optional[pulumi.Input['StatusArgs']]):
        pulumi.set(self, "validity", value)


class AlertPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_strategy: Optional[pulumi.Input[pulumi.InputType['AlertStrategyArgs']]] = None,
                 combiner: Optional[pulumi.Input['AlertPolicyCombiner']] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionArgs']]]]] = None,
                 creation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 documentation: Optional[pulumi.Input[pulumi.InputType['DocumentationArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mutation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input['AlertPolicySeverity']] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 validity: Optional[pulumi.Input[pulumi.InputType['StatusArgs']]] = None,
                 __props__=None):
        """
        Creates a new alerting policy.Design your application to single-thread API calls that modify the state of alerting policies in a single project. This includes calls to CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AlertStrategyArgs']] alert_strategy: Control over how this alert policy's notification channels are notified.
        :param pulumi.Input['AlertPolicyCombiner'] combiner: How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionArgs']]]] conditions: A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition. If condition_monitoring_query_language is present, it must be the only condition.
        :param pulumi.Input[pulumi.InputType['MutationRecordArgs']] creation_record: A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        :param pulumi.Input[str] display_name: A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.The convention for the display_name of a PrometheusQueryLanguageCondition is "{rule group name}/{alert name}", where the {rule group name} and {alert name} should be taken from the corresponding Prometheus configuration file. This convention is not enforced. In any case the display_name is not a unique key of the AlertPolicy.
        :param pulumi.Input[pulumi.InputType['DocumentationArgs']] documentation: Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        :param pulumi.Input[bool] enabled: Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        :param pulumi.Input[pulumi.InputType['MutationRecordArgs']] mutation_record: A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        :param pulumi.Input[str] name: Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Cloud Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_channels: Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        :param pulumi.Input['AlertPolicySeverity'] severity: Optional. The severity of an alert policy indicates how important incidents generated by that policy are. The severity level will be displayed on the Incident detail page and in notifications.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.Note that Prometheus {alert name} is a valid Prometheus label names (https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels), whereas Prometheus {rule group} is an unrestricted UTF-8 string. This means that they cannot be stored as-is in user labels, because they may contain characters that are not allowed in user-label values.
        :param pulumi.Input[pulumi.InputType['StatusArgs']] validity: Read-only description of how the alert policy is invalid. This field is only set when the alert policy is invalid. An invalid alert policy will not generate incidents.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AlertPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new alerting policy.Design your application to single-thread API calls that modify the state of alerting policies in a single project. This includes calls to CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.

        :param str resource_name: The name of the resource.
        :param AlertPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlertPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_strategy: Optional[pulumi.Input[pulumi.InputType['AlertStrategyArgs']]] = None,
                 combiner: Optional[pulumi.Input['AlertPolicyCombiner']] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionArgs']]]]] = None,
                 creation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 documentation: Optional[pulumi.Input[pulumi.InputType['DocumentationArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mutation_record: Optional[pulumi.Input[pulumi.InputType['MutationRecordArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input['AlertPolicySeverity']] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 validity: Optional[pulumi.Input[pulumi.InputType['StatusArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlertPolicyArgs.__new__(AlertPolicyArgs)

            __props__.__dict__["alert_strategy"] = alert_strategy
            __props__.__dict__["combiner"] = combiner
            __props__.__dict__["conditions"] = conditions
            __props__.__dict__["creation_record"] = creation_record
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["documentation"] = documentation
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["mutation_record"] = mutation_record
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_channels"] = notification_channels
            __props__.__dict__["project"] = project
            __props__.__dict__["severity"] = severity
            __props__.__dict__["user_labels"] = user_labels
            __props__.__dict__["validity"] = validity
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(AlertPolicy, __self__).__init__(
            'google-native:monitoring/v3:AlertPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AlertPolicy':
        """
        Get an existing AlertPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AlertPolicyArgs.__new__(AlertPolicyArgs)

        __props__.__dict__["alert_strategy"] = None
        __props__.__dict__["combiner"] = None
        __props__.__dict__["conditions"] = None
        __props__.__dict__["creation_record"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["documentation"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["mutation_record"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notification_channels"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["severity"] = None
        __props__.__dict__["user_labels"] = None
        __props__.__dict__["validity"] = None
        return AlertPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertStrategy")
    def alert_strategy(self) -> pulumi.Output['outputs.AlertStrategyResponse']:
        """
        Control over how this alert policy's notification channels are notified.
        """
        return pulumi.get(self, "alert_strategy")

    @property
    @pulumi.getter
    def combiner(self) -> pulumi.Output[str]:
        """
        How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
        """
        return pulumi.get(self, "combiner")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Sequence['outputs.ConditionResponse']]:
        """
        A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition. If condition_monitoring_query_language is present, it must be the only condition.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="creationRecord")
    def creation_record(self) -> pulumi.Output['outputs.MutationRecordResponse']:
        """
        A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
        """
        return pulumi.get(self, "creation_record")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.The convention for the display_name of a PrometheusQueryLanguageCondition is "{rule group name}/{alert name}", where the {rule group name} and {alert name} should be taken from the corresponding Prometheus configuration file. This convention is not enforced. In any case the display_name is not a unique key of the AlertPolicy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def documentation(self) -> pulumi.Output['outputs.DocumentationResponse']:
        """
        Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
        """
        return pulumi.get(self, "documentation")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="mutationRecord")
    def mutation_record(self) -> pulumi.Output['outputs.MutationRecordResponse']:
        """
        A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
        """
        return pulumi.get(self, "mutation_record")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Cloud Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationChannels")
    def notification_channels(self) -> pulumi.Output[Sequence[str]]:
        """
        Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
        """
        return pulumi.get(self, "notification_channels")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Optional. The severity of an alert policy indicates how important incidents generated by that policy are. The severity level will be displayed on the Incident detail page and in notifications.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.Note that Prometheus {alert name} is a valid Prometheus label names (https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels), whereas Prometheus {rule group} is an unrestricted UTF-8 string. This means that they cannot be stored as-is in user labels, because they may contain characters that are not allowed in user-label values.
        """
        return pulumi.get(self, "user_labels")

    @property
    @pulumi.getter
    def validity(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        Read-only description of how the alert policy is invalid. This field is only set when the alert policy is invalid. An invalid alert policy will not generate incidents.
        """
        return pulumi.get(self, "validity")

