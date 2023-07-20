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
    'GetSinkResult',
    'AwaitableGetSinkResult',
    'get_sink',
    'get_sink_output',
]

@pulumi.output_type
class GetSinkResult:
    def __init__(__self__, bigquery_options=None, create_time=None, description=None, destination=None, disabled=None, exclusions=None, filter=None, include_children=None, name=None, output_version_format=None, update_time=None, writer_identity=None):
        if bigquery_options and not isinstance(bigquery_options, dict):
            raise TypeError("Expected argument 'bigquery_options' to be a dict")
        pulumi.set(__self__, "bigquery_options", bigquery_options)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if exclusions and not isinstance(exclusions, list):
            raise TypeError("Expected argument 'exclusions' to be a list")
        pulumi.set(__self__, "exclusions", exclusions)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if include_children and not isinstance(include_children, bool):
            raise TypeError("Expected argument 'include_children' to be a bool")
        pulumi.set(__self__, "include_children", include_children)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_version_format and not isinstance(output_version_format, str):
            raise TypeError("Expected argument 'output_version_format' to be a str")
        pulumi.set(__self__, "output_version_format", output_version_format)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if writer_identity and not isinstance(writer_identity, str):
            raise TypeError("Expected argument 'writer_identity' to be a str")
        pulumi.set(__self__, "writer_identity", writer_identity)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> 'outputs.BigQueryOptionsResponse':
        """
        Optional. Options that affect sinks exporting data to BigQuery.
        """
        return pulumi.get(self, "bigquery_options")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of the sink.This field may not be present for older sinks.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A description of this sink.The maximum length of the description is 8000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> str:
        """
        The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Optional. If set to true, then this sink is disabled and it does not export any log entries.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def exclusions(self) -> Sequence['outputs.LogExclusionResponse']:
        """
        Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
        """
        return pulumi.get(self, "exclusions")

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="includeChildren")
    def include_children(self) -> bool:
        """
        Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
        """
        return pulumi.get(self, "include_children")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputVersionFormat")
    def output_version_format(self) -> str:
        """
        Deprecated. This field is unused.
        """
        warnings.warn("""Deprecated. This field is unused.""", DeprecationWarning)
        pulumi.log.warn("""output_version_format is deprecated: Deprecated. This field is unused.""")

        return pulumi.get(self, "output_version_format")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of the sink.This field may not be present for older sinks.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="writerIdentity")
    def writer_identity(self) -> str:
        """
        An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is either set by specifying custom_writer_identity or set automatically by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink cannot have a writer_identity and no additional permissions are required.
        """
        return pulumi.get(self, "writer_identity")


class AwaitableGetSinkResult(GetSinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSinkResult(
            bigquery_options=self.bigquery_options,
            create_time=self.create_time,
            description=self.description,
            destination=self.destination,
            disabled=self.disabled,
            exclusions=self.exclusions,
            filter=self.filter,
            include_children=self.include_children,
            name=self.name,
            output_version_format=self.output_version_format,
            update_time=self.update_time,
            writer_identity=self.writer_identity)


def get_sink(project: Optional[str] = None,
             sink_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSinkResult:
    """
    Gets a sink.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['sinkId'] = sink_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:logging/v2:getSink', __args__, opts=opts, typ=GetSinkResult).value

    return AwaitableGetSinkResult(
        bigquery_options=pulumi.get(__ret__, 'bigquery_options'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        destination=pulumi.get(__ret__, 'destination'),
        disabled=pulumi.get(__ret__, 'disabled'),
        exclusions=pulumi.get(__ret__, 'exclusions'),
        filter=pulumi.get(__ret__, 'filter'),
        include_children=pulumi.get(__ret__, 'include_children'),
        name=pulumi.get(__ret__, 'name'),
        output_version_format=pulumi.get(__ret__, 'output_version_format'),
        update_time=pulumi.get(__ret__, 'update_time'),
        writer_identity=pulumi.get(__ret__, 'writer_identity'))


@_utilities.lift_output_func(get_sink)
def get_sink_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                    sink_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSinkResult]:
    """
    Gets a sink.
    """
    ...
