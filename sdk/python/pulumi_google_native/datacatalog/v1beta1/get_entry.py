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
    'GetEntryResult',
    'AwaitableGetEntryResult',
    'get_entry',
    'get_entry_output',
]

@pulumi.output_type
class GetEntryResult:
    def __init__(__self__, bigquery_date_sharded_spec=None, bigquery_table_spec=None, description=None, display_name=None, gcs_fileset_spec=None, integrated_system=None, linked_resource=None, name=None, schema=None, source_system_timestamps=None, type=None, usage_signal=None, user_specified_system=None, user_specified_type=None):
        if bigquery_date_sharded_spec and not isinstance(bigquery_date_sharded_spec, dict):
            raise TypeError("Expected argument 'bigquery_date_sharded_spec' to be a dict")
        pulumi.set(__self__, "bigquery_date_sharded_spec", bigquery_date_sharded_spec)
        if bigquery_table_spec and not isinstance(bigquery_table_spec, dict):
            raise TypeError("Expected argument 'bigquery_table_spec' to be a dict")
        pulumi.set(__self__, "bigquery_table_spec", bigquery_table_spec)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if gcs_fileset_spec and not isinstance(gcs_fileset_spec, dict):
            raise TypeError("Expected argument 'gcs_fileset_spec' to be a dict")
        pulumi.set(__self__, "gcs_fileset_spec", gcs_fileset_spec)
        if integrated_system and not isinstance(integrated_system, str):
            raise TypeError("Expected argument 'integrated_system' to be a str")
        pulumi.set(__self__, "integrated_system", integrated_system)
        if linked_resource and not isinstance(linked_resource, str):
            raise TypeError("Expected argument 'linked_resource' to be a str")
        pulumi.set(__self__, "linked_resource", linked_resource)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if schema and not isinstance(schema, dict):
            raise TypeError("Expected argument 'schema' to be a dict")
        pulumi.set(__self__, "schema", schema)
        if source_system_timestamps and not isinstance(source_system_timestamps, dict):
            raise TypeError("Expected argument 'source_system_timestamps' to be a dict")
        pulumi.set(__self__, "source_system_timestamps", source_system_timestamps)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if usage_signal and not isinstance(usage_signal, dict):
            raise TypeError("Expected argument 'usage_signal' to be a dict")
        pulumi.set(__self__, "usage_signal", usage_signal)
        if user_specified_system and not isinstance(user_specified_system, str):
            raise TypeError("Expected argument 'user_specified_system' to be a str")
        pulumi.set(__self__, "user_specified_system", user_specified_system)
        if user_specified_type and not isinstance(user_specified_type, str):
            raise TypeError("Expected argument 'user_specified_type' to be a str")
        pulumi.set(__self__, "user_specified_type", user_specified_type)

    @property
    @pulumi.getter(name="bigqueryDateShardedSpec")
    def bigquery_date_sharded_spec(self) -> 'outputs.GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse':
        """
        Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        """
        return pulumi.get(self, "bigquery_date_sharded_spec")

    @property
    @pulumi.getter(name="bigqueryTableSpec")
    def bigquery_table_spec(self) -> 'outputs.GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse':
        """
        Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        """
        return pulumi.get(self, "bigquery_table_spec")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gcsFilesetSpec")
    def gcs_fileset_spec(self) -> 'outputs.GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse':
        """
        Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        """
        return pulumi.get(self, "gcs_fileset_spec")

    @property
    @pulumi.getter(name="integratedSystem")
    def integrated_system(self) -> str:
        """
        This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        """
        return pulumi.get(self, "integrated_system")

    @property
    @pulumi.getter(name="linkedResource")
    def linked_resource(self) -> str:
        """
        The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        """
        return pulumi.get(self, "linked_resource")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The Data Catalog resource name of the entry in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id} Note that this Entry and its child resources may not actually be stored in the location in this name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schema(self) -> 'outputs.GoogleCloudDatacatalogV1beta1SchemaResponse':
        """
        Schema of the entry. An entry might not have any schema attached to it.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="sourceSystemTimestamps")
    def source_system_timestamps(self) -> 'outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse':
        """
        Timestamps about the underlying resource, not about this Data Catalog entry. Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty timestamp.
        """
        return pulumi.get(self, "source_system_timestamps")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the entry. Only used for Entries with types in the EntryType enum.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageSignal")
    def usage_signal(self) -> 'outputs.GoogleCloudDatacatalogV1beta1UsageSignalResponse':
        """
        Statistics on the usage level of the resource.
        """
        return pulumi.get(self, "usage_signal")

    @property
    @pulumi.getter(name="userSpecifiedSystem")
    def user_specified_system(self) -> str:
        """
        This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        """
        return pulumi.get(self, "user_specified_system")

    @property
    @pulumi.getter(name="userSpecifiedType")
    def user_specified_type(self) -> str:
        """
        Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        """
        return pulumi.get(self, "user_specified_type")


class AwaitableGetEntryResult(GetEntryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntryResult(
            bigquery_date_sharded_spec=self.bigquery_date_sharded_spec,
            bigquery_table_spec=self.bigquery_table_spec,
            description=self.description,
            display_name=self.display_name,
            gcs_fileset_spec=self.gcs_fileset_spec,
            integrated_system=self.integrated_system,
            linked_resource=self.linked_resource,
            name=self.name,
            schema=self.schema,
            source_system_timestamps=self.source_system_timestamps,
            type=self.type,
            usage_signal=self.usage_signal,
            user_specified_system=self.user_specified_system,
            user_specified_type=self.user_specified_type)


def get_entry(entry_group_id: Optional[str] = None,
              entry_id: Optional[str] = None,
              location: Optional[str] = None,
              project: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntryResult:
    """
    Gets an entry.
    """
    __args__ = dict()
    __args__['entryGroupId'] = entry_group_id
    __args__['entryId'] = entry_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datacatalog/v1beta1:getEntry', __args__, opts=opts, typ=GetEntryResult).value

    return AwaitableGetEntryResult(
        bigquery_date_sharded_spec=pulumi.get(__ret__, 'bigquery_date_sharded_spec'),
        bigquery_table_spec=pulumi.get(__ret__, 'bigquery_table_spec'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        gcs_fileset_spec=pulumi.get(__ret__, 'gcs_fileset_spec'),
        integrated_system=pulumi.get(__ret__, 'integrated_system'),
        linked_resource=pulumi.get(__ret__, 'linked_resource'),
        name=pulumi.get(__ret__, 'name'),
        schema=pulumi.get(__ret__, 'schema'),
        source_system_timestamps=pulumi.get(__ret__, 'source_system_timestamps'),
        type=pulumi.get(__ret__, 'type'),
        usage_signal=pulumi.get(__ret__, 'usage_signal'),
        user_specified_system=pulumi.get(__ret__, 'user_specified_system'),
        user_specified_type=pulumi.get(__ret__, 'user_specified_type'))


@_utilities.lift_output_func(get_entry)
def get_entry_output(entry_group_id: Optional[pulumi.Input[str]] = None,
                     entry_id: Optional[pulumi.Input[str]] = None,
                     location: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEntryResult]:
    """
    Gets an entry.
    """
    ...
