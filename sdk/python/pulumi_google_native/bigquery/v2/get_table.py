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
    'GetTableResult',
    'AwaitableGetTableResult',
    'get_table',
    'get_table_output',
]

@pulumi.output_type
class GetTableResult:
    def __init__(__self__, clone_definition=None, clustering=None, creation_time=None, default_collation=None, description=None, encryption_configuration=None, etag=None, expiration_time=None, external_data_configuration=None, friendly_name=None, kind=None, labels=None, last_modified_time=None, location=None, materialized_view=None, model=None, num_active_logical_bytes=None, num_active_physical_bytes=None, num_bytes=None, num_long_term_bytes=None, num_long_term_logical_bytes=None, num_long_term_physical_bytes=None, num_partitions=None, num_physical_bytes=None, num_rows=None, num_time_travel_physical_bytes=None, num_total_logical_bytes=None, num_total_physical_bytes=None, range_partitioning=None, require_partition_filter=None, schema=None, self_link=None, snapshot_definition=None, streaming_buffer=None, table_reference=None, time_partitioning=None, type=None, view=None):
        if clone_definition and not isinstance(clone_definition, dict):
            raise TypeError("Expected argument 'clone_definition' to be a dict")
        pulumi.set(__self__, "clone_definition", clone_definition)
        if clustering and not isinstance(clustering, dict):
            raise TypeError("Expected argument 'clustering' to be a dict")
        pulumi.set(__self__, "clustering", clustering)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if default_collation and not isinstance(default_collation, str):
            raise TypeError("Expected argument 'default_collation' to be a str")
        pulumi.set(__self__, "default_collation", default_collation)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if encryption_configuration and not isinstance(encryption_configuration, dict):
            raise TypeError("Expected argument 'encryption_configuration' to be a dict")
        pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if expiration_time and not isinstance(expiration_time, str):
            raise TypeError("Expected argument 'expiration_time' to be a str")
        pulumi.set(__self__, "expiration_time", expiration_time)
        if external_data_configuration and not isinstance(external_data_configuration, dict):
            raise TypeError("Expected argument 'external_data_configuration' to be a dict")
        pulumi.set(__self__, "external_data_configuration", external_data_configuration)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if materialized_view and not isinstance(materialized_view, dict):
            raise TypeError("Expected argument 'materialized_view' to be a dict")
        pulumi.set(__self__, "materialized_view", materialized_view)
        if model and not isinstance(model, dict):
            raise TypeError("Expected argument 'model' to be a dict")
        pulumi.set(__self__, "model", model)
        if num_active_logical_bytes and not isinstance(num_active_logical_bytes, str):
            raise TypeError("Expected argument 'num_active_logical_bytes' to be a str")
        pulumi.set(__self__, "num_active_logical_bytes", num_active_logical_bytes)
        if num_active_physical_bytes and not isinstance(num_active_physical_bytes, str):
            raise TypeError("Expected argument 'num_active_physical_bytes' to be a str")
        pulumi.set(__self__, "num_active_physical_bytes", num_active_physical_bytes)
        if num_bytes and not isinstance(num_bytes, str):
            raise TypeError("Expected argument 'num_bytes' to be a str")
        pulumi.set(__self__, "num_bytes", num_bytes)
        if num_long_term_bytes and not isinstance(num_long_term_bytes, str):
            raise TypeError("Expected argument 'num_long_term_bytes' to be a str")
        pulumi.set(__self__, "num_long_term_bytes", num_long_term_bytes)
        if num_long_term_logical_bytes and not isinstance(num_long_term_logical_bytes, str):
            raise TypeError("Expected argument 'num_long_term_logical_bytes' to be a str")
        pulumi.set(__self__, "num_long_term_logical_bytes", num_long_term_logical_bytes)
        if num_long_term_physical_bytes and not isinstance(num_long_term_physical_bytes, str):
            raise TypeError("Expected argument 'num_long_term_physical_bytes' to be a str")
        pulumi.set(__self__, "num_long_term_physical_bytes", num_long_term_physical_bytes)
        if num_partitions and not isinstance(num_partitions, str):
            raise TypeError("Expected argument 'num_partitions' to be a str")
        pulumi.set(__self__, "num_partitions", num_partitions)
        if num_physical_bytes and not isinstance(num_physical_bytes, str):
            raise TypeError("Expected argument 'num_physical_bytes' to be a str")
        pulumi.set(__self__, "num_physical_bytes", num_physical_bytes)
        if num_rows and not isinstance(num_rows, str):
            raise TypeError("Expected argument 'num_rows' to be a str")
        pulumi.set(__self__, "num_rows", num_rows)
        if num_time_travel_physical_bytes and not isinstance(num_time_travel_physical_bytes, str):
            raise TypeError("Expected argument 'num_time_travel_physical_bytes' to be a str")
        pulumi.set(__self__, "num_time_travel_physical_bytes", num_time_travel_physical_bytes)
        if num_total_logical_bytes and not isinstance(num_total_logical_bytes, str):
            raise TypeError("Expected argument 'num_total_logical_bytes' to be a str")
        pulumi.set(__self__, "num_total_logical_bytes", num_total_logical_bytes)
        if num_total_physical_bytes and not isinstance(num_total_physical_bytes, str):
            raise TypeError("Expected argument 'num_total_physical_bytes' to be a str")
        pulumi.set(__self__, "num_total_physical_bytes", num_total_physical_bytes)
        if range_partitioning and not isinstance(range_partitioning, dict):
            raise TypeError("Expected argument 'range_partitioning' to be a dict")
        pulumi.set(__self__, "range_partitioning", range_partitioning)
        if require_partition_filter and not isinstance(require_partition_filter, bool):
            raise TypeError("Expected argument 'require_partition_filter' to be a bool")
        pulumi.set(__self__, "require_partition_filter", require_partition_filter)
        if schema and not isinstance(schema, dict):
            raise TypeError("Expected argument 'schema' to be a dict")
        pulumi.set(__self__, "schema", schema)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if snapshot_definition and not isinstance(snapshot_definition, dict):
            raise TypeError("Expected argument 'snapshot_definition' to be a dict")
        pulumi.set(__self__, "snapshot_definition", snapshot_definition)
        if streaming_buffer and not isinstance(streaming_buffer, dict):
            raise TypeError("Expected argument 'streaming_buffer' to be a dict")
        pulumi.set(__self__, "streaming_buffer", streaming_buffer)
        if table_reference and not isinstance(table_reference, dict):
            raise TypeError("Expected argument 'table_reference' to be a dict")
        pulumi.set(__self__, "table_reference", table_reference)
        if time_partitioning and not isinstance(time_partitioning, dict):
            raise TypeError("Expected argument 'time_partitioning' to be a dict")
        pulumi.set(__self__, "time_partitioning", time_partitioning)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if view and not isinstance(view, dict):
            raise TypeError("Expected argument 'view' to be a dict")
        pulumi.set(__self__, "view", view)

    @property
    @pulumi.getter(name="cloneDefinition")
    def clone_definition(self) -> 'outputs.CloneDefinitionResponse':
        """
        Clone definition.
        """
        return pulumi.get(self, "clone_definition")

    @property
    @pulumi.getter
    def clustering(self) -> 'outputs.ClusteringResponse':
        """
        [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        """
        return pulumi.get(self, "clustering")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The time when this table was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="defaultCollation")
    def default_collation(self) -> str:
        """
        The default collation of the table.
        """
        return pulumi.get(self, "default_collation")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        [Optional] A user-friendly description of this table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> 'outputs.EncryptionConfigurationResponse':
        """
        Custom encryption configuration (e.g., Cloud KMS keys).
        """
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A hash of the table metadata. Used to ensure there were no concurrent modifications to the resource when attempting an update. Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> str:
        """
        [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter(name="externalDataConfiguration")
    def external_data_configuration(self) -> 'outputs.ExternalDataConfigurationResponse':
        """
        [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        """
        return pulumi.get(self, "external_data_configuration")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> str:
        """
        [Optional] A descriptive name for this table.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        The time when this table was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geographic location where the table resides. This value is inherited from the dataset.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="materializedView")
    def materialized_view(self) -> 'outputs.MaterializedViewDefinitionResponse':
        """
        [Optional] Materialized view definition.
        """
        return pulumi.get(self, "materialized_view")

    @property
    @pulumi.getter
    def model(self) -> 'outputs.ModelDefinitionResponse':
        """
        [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter(name="numActiveLogicalBytes")
    def num_active_logical_bytes(self) -> str:
        """
        Number of logical bytes that are less than 90 days old.
        """
        return pulumi.get(self, "num_active_logical_bytes")

    @property
    @pulumi.getter(name="numActivePhysicalBytes")
    def num_active_physical_bytes(self) -> str:
        """
        Number of physical bytes less than 90 days old. This data is not kept in real time, and might be delayed by a few seconds to a few minutes.
        """
        return pulumi.get(self, "num_active_physical_bytes")

    @property
    @pulumi.getter(name="numBytes")
    def num_bytes(self) -> str:
        """
        The size of this table in bytes, excluding any data in the streaming buffer.
        """
        return pulumi.get(self, "num_bytes")

    @property
    @pulumi.getter(name="numLongTermBytes")
    def num_long_term_bytes(self) -> str:
        """
        The number of bytes in the table that are considered "long-term storage".
        """
        return pulumi.get(self, "num_long_term_bytes")

    @property
    @pulumi.getter(name="numLongTermLogicalBytes")
    def num_long_term_logical_bytes(self) -> str:
        """
        Number of logical bytes that are more than 90 days old.
        """
        return pulumi.get(self, "num_long_term_logical_bytes")

    @property
    @pulumi.getter(name="numLongTermPhysicalBytes")
    def num_long_term_physical_bytes(self) -> str:
        """
        Number of physical bytes more than 90 days old. This data is not kept in real time, and might be delayed by a few seconds to a few minutes.
        """
        return pulumi.get(self, "num_long_term_physical_bytes")

    @property
    @pulumi.getter(name="numPartitions")
    def num_partitions(self) -> str:
        """
        The number of partitions present in the table or materialized view. This data is not kept in real time, and might be delayed by a few seconds to a few minutes.
        """
        return pulumi.get(self, "num_partitions")

    @property
    @pulumi.getter(name="numPhysicalBytes")
    def num_physical_bytes(self) -> str:
        """
        [TrustedTester] The physical size of this table in bytes, excluding any data in the streaming buffer. This includes compression and storage used for time travel.
        """
        return pulumi.get(self, "num_physical_bytes")

    @property
    @pulumi.getter(name="numRows")
    def num_rows(self) -> str:
        """
        The number of rows of data in this table, excluding any data in the streaming buffer.
        """
        return pulumi.get(self, "num_rows")

    @property
    @pulumi.getter(name="numTimeTravelPhysicalBytes")
    def num_time_travel_physical_bytes(self) -> str:
        """
        Number of physical bytes used by time travel storage (deleted or changed data). This data is not kept in real time, and might be delayed by a few seconds to a few minutes.
        """
        return pulumi.get(self, "num_time_travel_physical_bytes")

    @property
    @pulumi.getter(name="numTotalLogicalBytes")
    def num_total_logical_bytes(self) -> str:
        """
        Total number of logical bytes in the table or materialized view.
        """
        return pulumi.get(self, "num_total_logical_bytes")

    @property
    @pulumi.getter(name="numTotalPhysicalBytes")
    def num_total_physical_bytes(self) -> str:
        """
        The physical size of this table in bytes. This also includes storage used for time travel. This data is not kept in real time, and might be delayed by a few seconds to a few minutes.
        """
        return pulumi.get(self, "num_total_physical_bytes")

    @property
    @pulumi.getter(name="rangePartitioning")
    def range_partitioning(self) -> 'outputs.RangePartitioningResponse':
        """
        [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "range_partitioning")

    @property
    @pulumi.getter(name="requirePartitionFilter")
    def require_partition_filter(self) -> bool:
        """
        [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        """
        return pulumi.get(self, "require_partition_filter")

    @property
    @pulumi.getter
    def schema(self) -> 'outputs.TableSchemaResponse':
        """
        [Optional] Describes the schema of this table.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        A URL that can be used to access this resource again.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="snapshotDefinition")
    def snapshot_definition(self) -> 'outputs.SnapshotDefinitionResponse':
        """
        Snapshot definition.
        """
        return pulumi.get(self, "snapshot_definition")

    @property
    @pulumi.getter(name="streamingBuffer")
    def streaming_buffer(self) -> 'outputs.StreamingbufferResponse':
        """
        Contains information regarding this table's streaming buffer, if one is present. This field will be absent if the table is not being streamed to or if there is no data in the streaming buffer.
        """
        return pulumi.get(self, "streaming_buffer")

    @property
    @pulumi.getter(name="tableReference")
    def table_reference(self) -> 'outputs.TableReferenceResponse':
        """
        [Required] Reference describing the ID of this table.
        """
        return pulumi.get(self, "table_reference")

    @property
    @pulumi.getter(name="timePartitioning")
    def time_partitioning(self) -> 'outputs.TimePartitioningResponse':
        """
        Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "time_partitioning")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Describes the table type. The following values are supported: TABLE: A normal BigQuery table. VIEW: A virtual table defined by a SQL query. SNAPSHOT: An immutable, read-only table that is a copy of another table. [TrustedTester] MATERIALIZED_VIEW: SQL query whose result is persisted. EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage. The default value is TABLE.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def view(self) -> 'outputs.ViewDefinitionResponse':
        """
        [Optional] The view definition.
        """
        return pulumi.get(self, "view")


class AwaitableGetTableResult(GetTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTableResult(
            clone_definition=self.clone_definition,
            clustering=self.clustering,
            creation_time=self.creation_time,
            default_collation=self.default_collation,
            description=self.description,
            encryption_configuration=self.encryption_configuration,
            etag=self.etag,
            expiration_time=self.expiration_time,
            external_data_configuration=self.external_data_configuration,
            friendly_name=self.friendly_name,
            kind=self.kind,
            labels=self.labels,
            last_modified_time=self.last_modified_time,
            location=self.location,
            materialized_view=self.materialized_view,
            model=self.model,
            num_active_logical_bytes=self.num_active_logical_bytes,
            num_active_physical_bytes=self.num_active_physical_bytes,
            num_bytes=self.num_bytes,
            num_long_term_bytes=self.num_long_term_bytes,
            num_long_term_logical_bytes=self.num_long_term_logical_bytes,
            num_long_term_physical_bytes=self.num_long_term_physical_bytes,
            num_partitions=self.num_partitions,
            num_physical_bytes=self.num_physical_bytes,
            num_rows=self.num_rows,
            num_time_travel_physical_bytes=self.num_time_travel_physical_bytes,
            num_total_logical_bytes=self.num_total_logical_bytes,
            num_total_physical_bytes=self.num_total_physical_bytes,
            range_partitioning=self.range_partitioning,
            require_partition_filter=self.require_partition_filter,
            schema=self.schema,
            self_link=self.self_link,
            snapshot_definition=self.snapshot_definition,
            streaming_buffer=self.streaming_buffer,
            table_reference=self.table_reference,
            time_partitioning=self.time_partitioning,
            type=self.type,
            view=self.view)


def get_table(dataset_id: Optional[str] = None,
              project: Optional[str] = None,
              selected_fields: Optional[str] = None,
              table_id: Optional[str] = None,
              view: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTableResult:
    """
    Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.
    """
    __args__ = dict()
    __args__['datasetId'] = dataset_id
    __args__['project'] = project
    __args__['selectedFields'] = selected_fields
    __args__['tableId'] = table_id
    __args__['view'] = view
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:bigquery/v2:getTable', __args__, opts=opts, typ=GetTableResult).value

    return AwaitableGetTableResult(
        clone_definition=__ret__.clone_definition,
        clustering=__ret__.clustering,
        creation_time=__ret__.creation_time,
        default_collation=__ret__.default_collation,
        description=__ret__.description,
        encryption_configuration=__ret__.encryption_configuration,
        etag=__ret__.etag,
        expiration_time=__ret__.expiration_time,
        external_data_configuration=__ret__.external_data_configuration,
        friendly_name=__ret__.friendly_name,
        kind=__ret__.kind,
        labels=__ret__.labels,
        last_modified_time=__ret__.last_modified_time,
        location=__ret__.location,
        materialized_view=__ret__.materialized_view,
        model=__ret__.model,
        num_active_logical_bytes=__ret__.num_active_logical_bytes,
        num_active_physical_bytes=__ret__.num_active_physical_bytes,
        num_bytes=__ret__.num_bytes,
        num_long_term_bytes=__ret__.num_long_term_bytes,
        num_long_term_logical_bytes=__ret__.num_long_term_logical_bytes,
        num_long_term_physical_bytes=__ret__.num_long_term_physical_bytes,
        num_partitions=__ret__.num_partitions,
        num_physical_bytes=__ret__.num_physical_bytes,
        num_rows=__ret__.num_rows,
        num_time_travel_physical_bytes=__ret__.num_time_travel_physical_bytes,
        num_total_logical_bytes=__ret__.num_total_logical_bytes,
        num_total_physical_bytes=__ret__.num_total_physical_bytes,
        range_partitioning=__ret__.range_partitioning,
        require_partition_filter=__ret__.require_partition_filter,
        schema=__ret__.schema,
        self_link=__ret__.self_link,
        snapshot_definition=__ret__.snapshot_definition,
        streaming_buffer=__ret__.streaming_buffer,
        table_reference=__ret__.table_reference,
        time_partitioning=__ret__.time_partitioning,
        type=__ret__.type,
        view=__ret__.view)


@_utilities.lift_output_func(get_table)
def get_table_output(dataset_id: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     selected_fields: Optional[pulumi.Input[Optional[str]]] = None,
                     table_id: Optional[pulumi.Input[str]] = None,
                     view: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTableResult]:
    """
    Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.
    """
    ...
