// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EntitySystem = {
    /**
     * Storage system unspecified.
     */
    StorageSystemUnspecified: "STORAGE_SYSTEM_UNSPECIFIED",
    /**
     * The entity data is contained within a Cloud Storage bucket.
     */
    CloudStorage: "CLOUD_STORAGE",
    /**
     * The entity data is contained within a BigQuery dataset.
     */
    Bigquery: "BIGQUERY",
} as const;

/**
 * Required. Immutable. Identifies the storage system of the entity data.
 */
export type EntitySystem = (typeof EntitySystem)[keyof typeof EntitySystem];

export const EntityType = {
    /**
     * Type unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Structured and semi-structured data.
     */
    Table: "TABLE",
    /**
     * Unstructured data.
     */
    Fileset: "FILESET",
} as const;

/**
 * Required. Immutable. The type of entity.
 */
export type EntityType = (typeof EntityType)[keyof typeof EntityType];

export const GoogleCloudDataplexV1AssetResourceSpecType = {
    /**
     * Type not specified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Cloud Storage bucket.
     */
    StorageBucket: "STORAGE_BUCKET",
    /**
     * BigQuery dataset.
     */
    BigqueryDataset: "BIGQUERY_DATASET",
} as const;

/**
 * Required. Immutable. Type of resource.
 */
export type GoogleCloudDataplexV1AssetResourceSpecType = (typeof GoogleCloudDataplexV1AssetResourceSpecType)[keyof typeof GoogleCloudDataplexV1AssetResourceSpecType];

export const GoogleCloudDataplexV1ContentNotebookKernelType = {
    /**
     * Kernel Type unspecified.
     */
    KernelTypeUnspecified: "KERNEL_TYPE_UNSPECIFIED",
    /**
     * Python 3 Kernel.
     */
    Python3: "PYTHON3",
} as const;

/**
 * Required. Kernel Type of the notebook.
 */
export type GoogleCloudDataplexV1ContentNotebookKernelType = (typeof GoogleCloudDataplexV1ContentNotebookKernelType)[keyof typeof GoogleCloudDataplexV1ContentNotebookKernelType];

export const GoogleCloudDataplexV1ContentSqlScriptEngine = {
    /**
     * Value was unspecified.
     */
    QueryEngineUnspecified: "QUERY_ENGINE_UNSPECIFIED",
    /**
     * Spark SQL Query.
     */
    Spark: "SPARK",
} as const;

/**
 * Required. Query Engine to be used for the Sql Query.
 */
export type GoogleCloudDataplexV1ContentSqlScriptEngine = (typeof GoogleCloudDataplexV1ContentSqlScriptEngine)[keyof typeof GoogleCloudDataplexV1ContentSqlScriptEngine];

export const GoogleCloudDataplexV1SchemaPartitionFieldType = {
    /**
     * SchemaType unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Boolean field.
     */
    Boolean: "BOOLEAN",
    /**
     * Single byte numeric field.
     */
    Byte: "BYTE",
    /**
     * 16-bit numeric field.
     */
    Int16: "INT16",
    /**
     * 32-bit numeric field.
     */
    Int32: "INT32",
    /**
     * 64-bit numeric field.
     */
    Int64: "INT64",
    /**
     * Floating point numeric field.
     */
    Float: "FLOAT",
    /**
     * Double precision numeric field.
     */
    Double: "DOUBLE",
    /**
     * Real value numeric field.
     */
    Decimal: "DECIMAL",
    /**
     * Sequence of characters field.
     */
    String: "STRING",
    /**
     * Sequence of bytes field.
     */
    Binary: "BINARY",
    /**
     * Date and time field.
     */
    Timestamp: "TIMESTAMP",
    /**
     * Date field.
     */
    Date: "DATE",
    /**
     * Time field.
     */
    Time: "TIME",
    /**
     * Structured field. Nested fields that define the structure of the map. If all nested fields are nullable, this field represents a union.
     */
    Record: "RECORD",
    /**
     * Null field that does not have values.
     */
    Null: "NULL",
} as const;

/**
 * Required. Immutable. The type of field.
 */
export type GoogleCloudDataplexV1SchemaPartitionFieldType = (typeof GoogleCloudDataplexV1SchemaPartitionFieldType)[keyof typeof GoogleCloudDataplexV1SchemaPartitionFieldType];

export const GoogleCloudDataplexV1SchemaPartitionStyle = {
    /**
     * PartitionStyle unspecified
     */
    PartitionStyleUnspecified: "PARTITION_STYLE_UNSPECIFIED",
    /**
     * Partitions are hive-compatible. Examples: gs://bucket/path/to/table/dt=2019-10-31/lang=en, gs://bucket/path/to/table/dt=2019-10-31/lang=en/late.
     */
    HiveCompatible: "HIVE_COMPATIBLE",
} as const;

/**
 * Optional. The structure of paths containing partition data within the entity.
 */
export type GoogleCloudDataplexV1SchemaPartitionStyle = (typeof GoogleCloudDataplexV1SchemaPartitionStyle)[keyof typeof GoogleCloudDataplexV1SchemaPartitionStyle];

export const GoogleCloudDataplexV1SchemaSchemaFieldMode = {
    /**
     * Mode unspecified.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * The field has required semantics.
     */
    Required: "REQUIRED",
    /**
     * The field has optional semantics, and may be null.
     */
    Nullable: "NULLABLE",
    /**
     * The field has repeated (0 or more) semantics, and is a list of values.
     */
    Repeated: "REPEATED",
} as const;

/**
 * Required. Additional field semantics.
 */
export type GoogleCloudDataplexV1SchemaSchemaFieldMode = (typeof GoogleCloudDataplexV1SchemaSchemaFieldMode)[keyof typeof GoogleCloudDataplexV1SchemaSchemaFieldMode];

export const GoogleCloudDataplexV1SchemaSchemaFieldType = {
    /**
     * SchemaType unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Boolean field.
     */
    Boolean: "BOOLEAN",
    /**
     * Single byte numeric field.
     */
    Byte: "BYTE",
    /**
     * 16-bit numeric field.
     */
    Int16: "INT16",
    /**
     * 32-bit numeric field.
     */
    Int32: "INT32",
    /**
     * 64-bit numeric field.
     */
    Int64: "INT64",
    /**
     * Floating point numeric field.
     */
    Float: "FLOAT",
    /**
     * Double precision numeric field.
     */
    Double: "DOUBLE",
    /**
     * Real value numeric field.
     */
    Decimal: "DECIMAL",
    /**
     * Sequence of characters field.
     */
    String: "STRING",
    /**
     * Sequence of bytes field.
     */
    Binary: "BINARY",
    /**
     * Date and time field.
     */
    Timestamp: "TIMESTAMP",
    /**
     * Date field.
     */
    Date: "DATE",
    /**
     * Time field.
     */
    Time: "TIME",
    /**
     * Structured field. Nested fields that define the structure of the map. If all nested fields are nullable, this field represents a union.
     */
    Record: "RECORD",
    /**
     * Null field that does not have values.
     */
    Null: "NULL",
} as const;

/**
 * Required. The type of field.
 */
export type GoogleCloudDataplexV1SchemaSchemaFieldType = (typeof GoogleCloudDataplexV1SchemaSchemaFieldType)[keyof typeof GoogleCloudDataplexV1SchemaSchemaFieldType];

export const GoogleCloudDataplexV1StorageFormatCompressionFormat = {
    /**
     * CompressionFormat unspecified. Implies uncompressed data.
     */
    CompressionFormatUnspecified: "COMPRESSION_FORMAT_UNSPECIFIED",
    /**
     * GZip compressed set of files.
     */
    Gzip: "GZIP",
    /**
     * BZip2 compressed set of files.
     */
    Bzip2: "BZIP2",
} as const;

/**
 * Optional. The compression type associated with the stored data. If unspecified, the data is uncompressed.
 */
export type GoogleCloudDataplexV1StorageFormatCompressionFormat = (typeof GoogleCloudDataplexV1StorageFormatCompressionFormat)[keyof typeof GoogleCloudDataplexV1StorageFormatCompressionFormat];

export const GoogleCloudDataplexV1TaskTriggerSpecType = {
    /**
     * Unspecified trigger type.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * The task runs one-time shortly after Task Creation.
     */
    OnDemand: "ON_DEMAND",
    /**
     * The task is scheduled to run periodically.
     */
    Recurring: "RECURRING",
} as const;

/**
 * Required. Immutable. Trigger type of the user-specified Task.
 */
export type GoogleCloudDataplexV1TaskTriggerSpecType = (typeof GoogleCloudDataplexV1TaskTriggerSpecType)[keyof typeof GoogleCloudDataplexV1TaskTriggerSpecType];

export const GoogleCloudDataplexV1ZoneResourceSpecLocationType = {
    /**
     * Unspecified location type.
     */
    LocationTypeUnspecified: "LOCATION_TYPE_UNSPECIFIED",
    /**
     * Resources that are associated with a single region.
     */
    SingleRegion: "SINGLE_REGION",
    /**
     * Resources that are associated with a multi-region location.
     */
    MultiRegion: "MULTI_REGION",
} as const;

/**
 * Required. Immutable. The location type of the resources that are allowed to be attached to the assets within this zone.
 */
export type GoogleCloudDataplexV1ZoneResourceSpecLocationType = (typeof GoogleCloudDataplexV1ZoneResourceSpecLocationType)[keyof typeof GoogleCloudDataplexV1ZoneResourceSpecLocationType];

export const GoogleIamV1AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type GoogleIamV1AuditLogConfigLogType = (typeof GoogleIamV1AuditLogConfigLogType)[keyof typeof GoogleIamV1AuditLogConfigLogType];

export const ZoneType = {
    /**
     * Zone type not specified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * A zone that contains data that needs further processing before it is considered generally ready for consumption and analytics workloads.
     */
    Raw: "RAW",
    /**
     * A zone that contains data that is considered to be ready for broader consumption and analytics workloads. Curated structured data stored in Cloud Storage must conform to certain file formats (parquet, avro and orc) and organized in a hive-compatible directory layout.
     */
    Curated: "CURATED",
} as const;

/**
 * Required. Immutable. The type of the zone.
 */
export type ZoneType = (typeof ZoneType)[keyof typeof ZoneType];