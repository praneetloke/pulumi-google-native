# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    def __init__(__self__, bigquery_date_sharded_spec=None, bigquery_table_spec=None, business_context=None, data_source=None, data_source_connection_spec=None, database_table_spec=None, description=None, display_name=None, fully_qualified_name=None, gcs_fileset_spec=None, integrated_system=None, labels=None, linked_resource=None, name=None, personal_details=None, routine_spec=None, schema=None, source_system_timestamps=None, type=None, usage_signal=None, user_specified_system=None, user_specified_type=None):
        if bigquery_date_sharded_spec and not isinstance(bigquery_date_sharded_spec, dict):
            raise TypeError("Expected argument 'bigquery_date_sharded_spec' to be a dict")
        pulumi.set(__self__, "bigquery_date_sharded_spec", bigquery_date_sharded_spec)
        if bigquery_table_spec and not isinstance(bigquery_table_spec, dict):
            raise TypeError("Expected argument 'bigquery_table_spec' to be a dict")
        pulumi.set(__self__, "bigquery_table_spec", bigquery_table_spec)
        if business_context and not isinstance(business_context, dict):
            raise TypeError("Expected argument 'business_context' to be a dict")
        pulumi.set(__self__, "business_context", business_context)
        if data_source and not isinstance(data_source, dict):
            raise TypeError("Expected argument 'data_source' to be a dict")
        pulumi.set(__self__, "data_source", data_source)
        if data_source_connection_spec and not isinstance(data_source_connection_spec, dict):
            raise TypeError("Expected argument 'data_source_connection_spec' to be a dict")
        pulumi.set(__self__, "data_source_connection_spec", data_source_connection_spec)
        if database_table_spec and not isinstance(database_table_spec, dict):
            raise TypeError("Expected argument 'database_table_spec' to be a dict")
        pulumi.set(__self__, "database_table_spec", database_table_spec)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if fully_qualified_name and not isinstance(fully_qualified_name, str):
            raise TypeError("Expected argument 'fully_qualified_name' to be a str")
        pulumi.set(__self__, "fully_qualified_name", fully_qualified_name)
        if gcs_fileset_spec and not isinstance(gcs_fileset_spec, dict):
            raise TypeError("Expected argument 'gcs_fileset_spec' to be a dict")
        pulumi.set(__self__, "gcs_fileset_spec", gcs_fileset_spec)
        if integrated_system and not isinstance(integrated_system, str):
            raise TypeError("Expected argument 'integrated_system' to be a str")
        pulumi.set(__self__, "integrated_system", integrated_system)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if linked_resource and not isinstance(linked_resource, str):
            raise TypeError("Expected argument 'linked_resource' to be a str")
        pulumi.set(__self__, "linked_resource", linked_resource)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if personal_details and not isinstance(personal_details, dict):
            raise TypeError("Expected argument 'personal_details' to be a dict")
        pulumi.set(__self__, "personal_details", personal_details)
        if routine_spec and not isinstance(routine_spec, dict):
            raise TypeError("Expected argument 'routine_spec' to be a dict")
        pulumi.set(__self__, "routine_spec", routine_spec)
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
    def bigquery_date_sharded_spec(self) -> 'outputs.GoogleCloudDatacatalogV1BigQueryDateShardedSpecResponse':
        """
        Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
        """
        return pulumi.get(self, "bigquery_date_sharded_spec")

    @property
    @pulumi.getter(name="bigqueryTableSpec")
    def bigquery_table_spec(self) -> 'outputs.GoogleCloudDatacatalogV1BigQueryTableSpecResponse':
        """
        Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
        """
        return pulumi.get(self, "bigquery_table_spec")

    @property
    @pulumi.getter(name="businessContext")
    def business_context(self) -> 'outputs.GoogleCloudDatacatalogV1BusinessContextResponse':
        """
        Business Context of the entry.
        """
        return pulumi.get(self, "business_context")

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> 'outputs.GoogleCloudDatacatalogV1DataSourceResponse':
        """
        Physical location of the entry.
        """
        return pulumi.get(self, "data_source")

    @property
    @pulumi.getter(name="dataSourceConnectionSpec")
    def data_source_connection_spec(self) -> 'outputs.GoogleCloudDatacatalogV1DataSourceConnectionSpecResponse':
        """
        Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
        """
        return pulumi.get(self, "data_source_connection_spec")

    @property
    @pulumi.getter(name="databaseTableSpec")
    def database_table_spec(self) -> 'outputs.GoogleCloudDatacatalogV1DatabaseTableSpecResponse':
        """
        Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
        """
        return pulumi.get(self, "database_table_spec")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fullyQualifiedName")
    def fully_qualified_name(self) -> str:
        """
        Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
        """
        return pulumi.get(self, "fully_qualified_name")

    @property
    @pulumi.getter(name="gcsFilesetSpec")
    def gcs_fileset_spec(self) -> 'outputs.GoogleCloudDatacatalogV1GcsFilesetSpecResponse':
        """
        Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
        """
        return pulumi.get(self, "gcs_fileset_spec")

    @property
    @pulumi.getter(name="integratedSystem")
    def integrated_system(self) -> str:
        """
        Indicates the entry's source system that Data Catalog integrates with, such as BigQuery, Pub/Sub, or Dataproc Metastore.
        """
        return pulumi.get(self, "integrated_system")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="linkedResource")
    def linked_resource(self) -> str:
        """
        The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
        """
        return pulumi.get(self, "linked_resource")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of an entry in URL format. Note: The entry itself and its child resources might not be stored in the location specified in its name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="personalDetails")
    def personal_details(self) -> 'outputs.GoogleCloudDatacatalogV1PersonalDetailsResponse':
        """
        Additional information related to the entry. Private to the current user.
        """
        return pulumi.get(self, "personal_details")

    @property
    @pulumi.getter(name="routineSpec")
    def routine_spec(self) -> 'outputs.GoogleCloudDatacatalogV1RoutineSpecResponse':
        """
        Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
        """
        return pulumi.get(self, "routine_spec")

    @property
    @pulumi.getter
    def schema(self) -> 'outputs.GoogleCloudDatacatalogV1SchemaResponse':
        """
        Schema of the entry. An entry might not have any schema attached to it.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="sourceSystemTimestamps")
    def source_system_timestamps(self) -> 'outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse':
        """
        Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
        """
        return pulumi.get(self, "source_system_timestamps")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageSignal")
    def usage_signal(self) -> 'outputs.GoogleCloudDatacatalogV1UsageSignalResponse':
        """
        Resource usage statistics.
        """
        return pulumi.get(self, "usage_signal")

    @property
    @pulumi.getter(name="userSpecifiedSystem")
    def user_specified_system(self) -> str:
        """
        Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
        """
        return pulumi.get(self, "user_specified_system")

    @property
    @pulumi.getter(name="userSpecifiedType")
    def user_specified_type(self) -> str:
        """
        Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
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
            business_context=self.business_context,
            data_source=self.data_source,
            data_source_connection_spec=self.data_source_connection_spec,
            database_table_spec=self.database_table_spec,
            description=self.description,
            display_name=self.display_name,
            fully_qualified_name=self.fully_qualified_name,
            gcs_fileset_spec=self.gcs_fileset_spec,
            integrated_system=self.integrated_system,
            labels=self.labels,
            linked_resource=self.linked_resource,
            name=self.name,
            personal_details=self.personal_details,
            routine_spec=self.routine_spec,
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
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:datacatalog/v1:getEntry', __args__, opts=opts, typ=GetEntryResult).value

    return AwaitableGetEntryResult(
        bigquery_date_sharded_spec=__ret__.bigquery_date_sharded_spec,
        bigquery_table_spec=__ret__.bigquery_table_spec,
        business_context=__ret__.business_context,
        data_source=__ret__.data_source,
        data_source_connection_spec=__ret__.data_source_connection_spec,
        database_table_spec=__ret__.database_table_spec,
        description=__ret__.description,
        display_name=__ret__.display_name,
        fully_qualified_name=__ret__.fully_qualified_name,
        gcs_fileset_spec=__ret__.gcs_fileset_spec,
        integrated_system=__ret__.integrated_system,
        labels=__ret__.labels,
        linked_resource=__ret__.linked_resource,
        name=__ret__.name,
        personal_details=__ret__.personal_details,
        routine_spec=__ret__.routine_spec,
        schema=__ret__.schema,
        source_system_timestamps=__ret__.source_system_timestamps,
        type=__ret__.type,
        usage_signal=__ret__.usage_signal,
        user_specified_system=__ret__.user_specified_system,
        user_specified_type=__ret__.user_specified_type)


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
