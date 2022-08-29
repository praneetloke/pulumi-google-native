// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuditLogConfigLogType = {
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
export type AuditLogConfigLogType = (typeof AuditLogConfigLogType)[keyof typeof AuditLogConfigLogType];

export const ClusterDefaultStorageType = {
    /**
     * The user did not specify a storage type.
     */
    StorageTypeUnspecified: "STORAGE_TYPE_UNSPECIFIED",
    /**
     * Flash (SSD) storage should be used.
     */
    Ssd: "SSD",
    /**
     * Magnetic drive (HDD) storage should be used.
     */
    Hdd: "HDD",
} as const;

/**
 * Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
 */
export type ClusterDefaultStorageType = (typeof ClusterDefaultStorageType)[keyof typeof ClusterDefaultStorageType];

export const InstanceType = {
    /**
     * The type of the instance is unspecified. If set when creating an instance, a `PRODUCTION` instance will be created. If set when updating an instance, the type will be left unchanged.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * An instance meant for production use. `serve_nodes` must be set on the cluster.
     */
    Production: "PRODUCTION",
    /**
     * DEPRECATED: Prefer PRODUCTION for all use cases, as it no longer enforces a higher minimum node count than DEVELOPMENT.
     */
    Development: "DEVELOPMENT",
} as const;

/**
 * The type of the instance. Defaults to `PRODUCTION`.
 */
export type InstanceType = (typeof InstanceType)[keyof typeof InstanceType];

export const TableGranularity = {
    /**
     * The user did not specify a granularity. Should not be returned. When specified during table creation, MILLIS will be used.
     */
    TimestampGranularityUnspecified: "TIMESTAMP_GRANULARITY_UNSPECIFIED",
    /**
     * The table keeps data versioned at a granularity of 1ms.
     */
    Millis: "MILLIS",
} as const;

/**
 * Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
 */
export type TableGranularity = (typeof TableGranularity)[keyof typeof TableGranularity];
