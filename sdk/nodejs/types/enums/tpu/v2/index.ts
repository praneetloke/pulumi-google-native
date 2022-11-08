// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AttachedDiskMode = {
    /**
     * The disk mode is not known/set.
     */
    DiskModeUnspecified: "DISK_MODE_UNSPECIFIED",
    /**
     * Attaches the disk in read-write mode. Only one TPU node can attach a disk in read-write mode at a time.
     */
    ReadWrite: "READ_WRITE",
    /**
     * Attaches the disk in read-only mode. Multiple TPU nodes can attach a disk in read-only mode at a time.
     */
    ReadOnly: "READ_ONLY",
} as const;

/**
 * The mode in which to attach this disk. If not specified, the default is READ_WRITE mode. Only applicable to data_disks.
 */
export type AttachedDiskMode = (typeof AttachedDiskMode)[keyof typeof AttachedDiskMode];

export const NodeHealth = {
    /**
     * Health status is unknown: not initialized or failed to retrieve.
     */
    HealthUnspecified: "HEALTH_UNSPECIFIED",
    /**
     * The resource is healthy.
     */
    Healthy: "HEALTHY",
    /**
     * The resource is unresponsive.
     */
    Timeout: "TIMEOUT",
    /**
     * The in-guest ML stack is unhealthy.
     */
    UnhealthyTensorflow: "UNHEALTHY_TENSORFLOW",
    /**
     * The node is under maintenance/priority boost caused rescheduling and will resume running once rescheduled.
     */
    UnhealthyMaintenance: "UNHEALTHY_MAINTENANCE",
} as const;

/**
 * The health status of the TPU node.
 */
export type NodeHealth = (typeof NodeHealth)[keyof typeof NodeHealth];
