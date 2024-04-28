# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'BackendMetastoreMetastoreType',
    'DatabaseDumpDatabaseType',
    'DatabaseDumpType',
    'HiveMetastoreConfigEndpointProtocol',
    'MaintenanceWindowDayOfWeek',
    'ScalingConfigInstanceSize',
    'ServiceDatabaseType',
    'ServiceReleaseChannel',
    'ServiceTier',
    'TelemetryConfigLogFormat',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class BackendMetastoreMetastoreType(str, Enum):
    """
    The type of the backend metastore.
    """
    METASTORE_TYPE_UNSPECIFIED = "METASTORE_TYPE_UNSPECIFIED"
    """
    The metastore type is not set.
    """
    DATAPLEX = "DATAPLEX"
    """
    The backend metastore is Dataplex.
    """
    BIGQUERY = "BIGQUERY"
    """
    The backend metastore is BigQuery.
    """
    DATAPROC_METASTORE = "DATAPROC_METASTORE"
    """
    The backend metastore is Dataproc Metastore.
    """


class DatabaseDumpDatabaseType(str, Enum):
    """
    The type of the database.
    """
    DATABASE_TYPE_UNSPECIFIED = "DATABASE_TYPE_UNSPECIFIED"
    """
    The type of the source database is unknown.
    """
    MYSQL = "MYSQL"
    """
    The type of the source database is MySQL.
    """


class DatabaseDumpType(str, Enum):
    """
    Optional. The type of the database dump. If unspecified, defaults to MYSQL.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    The type of the database dump is unknown.
    """
    MYSQL = "MYSQL"
    """
    Database dump is a MySQL dump file.
    """
    AVRO = "AVRO"
    """
    Database dump contains Avro files.
    """


class HiveMetastoreConfigEndpointProtocol(str, Enum):
    """
    The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
    """
    ENDPOINT_PROTOCOL_UNSPECIFIED = "ENDPOINT_PROTOCOL_UNSPECIFIED"
    """
    The protocol is not set.
    """
    THRIFT = "THRIFT"
    """
    Use the legacy Apache Thrift protocol for the metastore service endpoint.
    """
    GRPC = "GRPC"
    """
    Use the modernized gRPC protocol for the metastore service endpoint.
    """


class MaintenanceWindowDayOfWeek(str, Enum):
    """
    The day of week, when the window starts.
    """
    DAY_OF_WEEK_UNSPECIFIED = "DAY_OF_WEEK_UNSPECIFIED"
    """
    The day of the week is unspecified.
    """
    MONDAY = "MONDAY"
    """
    Monday
    """
    TUESDAY = "TUESDAY"
    """
    Tuesday
    """
    WEDNESDAY = "WEDNESDAY"
    """
    Wednesday
    """
    THURSDAY = "THURSDAY"
    """
    Thursday
    """
    FRIDAY = "FRIDAY"
    """
    Friday
    """
    SATURDAY = "SATURDAY"
    """
    Saturday
    """
    SUNDAY = "SUNDAY"
    """
    Sunday
    """


class ScalingConfigInstanceSize(str, Enum):
    """
    An enum of readable instance sizes, with each instance size mapping to a float value (e.g. InstanceSize.EXTRA_SMALL = scaling_factor(0.1))
    """
    INSTANCE_SIZE_UNSPECIFIED = "INSTANCE_SIZE_UNSPECIFIED"
    """
    Unspecified instance size
    """
    EXTRA_SMALL = "EXTRA_SMALL"
    """
    Extra small instance size, maps to a scaling factor of 0.1.
    """
    SMALL = "SMALL"
    """
    Small instance size, maps to a scaling factor of 0.5.
    """
    MEDIUM = "MEDIUM"
    """
    Medium instance size, maps to a scaling factor of 1.0.
    """
    LARGE = "LARGE"
    """
    Large instance size, maps to a scaling factor of 3.0.
    """
    EXTRA_LARGE = "EXTRA_LARGE"
    """
    Extra large instance size, maps to a scaling factor of 6.0.
    """


class ServiceDatabaseType(str, Enum):
    """
    Immutable. The database type that the Metastore service stores its data.
    """
    DATABASE_TYPE_UNSPECIFIED = "DATABASE_TYPE_UNSPECIFIED"
    """
    The DATABASE_TYPE is not set.
    """
    MYSQL = "MYSQL"
    """
    MySQL is used to persist the metastore data.
    """
    SPANNER = "SPANNER"
    """
    Spanner is used to persist the metastore data.
    """


class ServiceReleaseChannel(str, Enum):
    """
    Immutable. The release channel of the service. If unspecified, defaults to STABLE.
    """
    RELEASE_CHANNEL_UNSPECIFIED = "RELEASE_CHANNEL_UNSPECIFIED"
    """
    Release channel is not specified.
    """
    CANARY = "CANARY"
    """
    The CANARY release channel contains the newest features, which may be unstable and subject to unresolved issues with no known workarounds. Services using the CANARY release channel are not subject to any SLAs.
    """
    STABLE = "STABLE"
    """
    The STABLE release channel contains features that are considered stable and have been validated for production use.
    """


class ServiceTier(str, Enum):
    """
    The tier of the service.
    """
    TIER_UNSPECIFIED = "TIER_UNSPECIFIED"
    """
    The tier is not set.
    """
    DEVELOPER = "DEVELOPER"
    """
    The developer tier provides limited scalability and no fault tolerance. Good for low-cost proof-of-concept.
    """
    ENTERPRISE = "ENTERPRISE"
    """
    The enterprise tier provides multi-zone high availability, and sufficient scalability for enterprise-level Dataproc Metastore workloads.
    """


class TelemetryConfigLogFormat(str, Enum):
    """
    The output format of the Dataproc Metastore service's logs.
    """
    LOG_FORMAT_UNSPECIFIED = "LOG_FORMAT_UNSPECIFIED"
    """
    The LOG_FORMAT is not set.
    """
    LEGACY = "LEGACY"
    """
    Logging output uses the legacy textPayload format.
    """
    JSON = "JSON"
    """
    Logging output uses the jsonPayload format.
    """
