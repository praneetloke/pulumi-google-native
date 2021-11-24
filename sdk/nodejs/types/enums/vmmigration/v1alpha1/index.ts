// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ComputeEngineTargetDefaultsDiskType = {
    /**
     * An unspecified disk type. Will be used as STANDARD.
     */
    ComputeEngineDiskTypeUnspecified: "COMPUTE_ENGINE_DISK_TYPE_UNSPECIFIED",
    /**
     * A Standard disk type.
     */
    ComputeEngineDiskTypeStandard: "COMPUTE_ENGINE_DISK_TYPE_STANDARD",
    /**
     * SSD hard disk type.
     */
    ComputeEngineDiskTypeSsd: "COMPUTE_ENGINE_DISK_TYPE_SSD",
    /**
     * An alternative to SSD persistent disks that balance performance and cost.
     */
    ComputeEngineDiskTypeBalanced: "COMPUTE_ENGINE_DISK_TYPE_BALANCED",
} as const;

/**
 * The disk type to use in the VM.
 */
export type ComputeEngineTargetDefaultsDiskType = (typeof ComputeEngineTargetDefaultsDiskType)[keyof typeof ComputeEngineTargetDefaultsDiskType];

export const ComputeEngineTargetDefaultsLicenseType = {
    /**
     * The license type is the default for the OS.
     */
    ComputeEngineLicenseTypeDefault: "COMPUTE_ENGINE_LICENSE_TYPE_DEFAULT",
    /**
     * The license type is Pay As You Go license type.
     */
    ComputeEngineLicenseTypePayg: "COMPUTE_ENGINE_LICENSE_TYPE_PAYG",
    /**
     * The license type is Bring Your Own License type.
     */
    ComputeEngineLicenseTypeByol: "COMPUTE_ENGINE_LICENSE_TYPE_BYOL",
} as const;

/**
 * The license type to use in OS adaptation.
 */
export type ComputeEngineTargetDefaultsLicenseType = (typeof ComputeEngineTargetDefaultsLicenseType)[keyof typeof ComputeEngineTargetDefaultsLicenseType];

export const ComputeSchedulingOnHostMaintenance = {
    /**
     * An unknown, unexpected behavior.
     */
    OnHostMaintenanceUnspecified: "ON_HOST_MAINTENANCE_UNSPECIFIED",
    /**
     * Terminate the instance when the host machine undergoes maintenance.
     */
    Terminate: "TERMINATE",
    /**
     * Migrate the instance when the host machine undergoes maintenance.
     */
    Migrate: "MIGRATE",
} as const;

/**
 * How the instance should behave when the host machine undergoes maintenance that may temporarily impact instance performance.
 */
export type ComputeSchedulingOnHostMaintenance = (typeof ComputeSchedulingOnHostMaintenance)[keyof typeof ComputeSchedulingOnHostMaintenance];

export const ComputeSchedulingRestartType = {
    /**
     * Unspecified behavior. This will use the default.
     */
    RestartTypeUnspecified: "RESTART_TYPE_UNSPECIFIED",
    /**
     * The Instance should be automatically restarted whenever it is terminated by Compute Engine.
     */
    AutomaticRestart: "AUTOMATIC_RESTART",
    /**
     * The Instance isn't automatically restarted whenever it is terminated by Compute Engine.
     */
    NoAutomaticRestart: "NO_AUTOMATIC_RESTART",
} as const;

/**
 * Whether the Instance should be automatically restarted whenever it is terminated by Compute Engine (not terminated by user). This configuration is identical to `automaticRestart` field in Compute Engine create instance under scheduling. It was changed to an enum (instead of a boolean) to match the default value in Compute Engine which is automatic restart.
 */
export type ComputeSchedulingRestartType = (typeof ComputeSchedulingRestartType)[keyof typeof ComputeSchedulingRestartType];

export const SchedulingNodeAffinityOperator = {
    /**
     * An unknown, unexpected behavior.
     */
    OperatorUnspecified: "OPERATOR_UNSPECIFIED",
    /**
     * The node resource group should be in these resources affinity.
     */
    In: "IN",
    /**
     * The node resource group should not be in these resources affinity.
     */
    NotIn: "NOT_IN",
} as const;

/**
 * The operator to use for the node resources specified in the `values` parameter.
 */
export type SchedulingNodeAffinityOperator = (typeof SchedulingNodeAffinityOperator)[keyof typeof SchedulingNodeAffinityOperator];

export const UtilizationReportTimeFrame = {
    /**
     * The time frame was not specified and will default to WEEK.
     */
    TimeFrameUnspecified: "TIME_FRAME_UNSPECIFIED",
    /**
     * One week.
     */
    Week: "WEEK",
    /**
     * One month.
     */
    Month: "MONTH",
    /**
     * One year.
     */
    Year: "YEAR",
} as const;

/**
 * Time frame of the report.
 */
export type UtilizationReportTimeFrame = (typeof UtilizationReportTimeFrame)[keyof typeof UtilizationReportTimeFrame];

export const VmwareVmDetailsPowerState = {
    /**
     * Power state is not specified.
     */
    PowerStateUnspecified: "POWER_STATE_UNSPECIFIED",
    /**
     * The VM is turned ON.
     */
    On: "ON",
    /**
     * The VM is turned OFF.
     */
    Off: "OFF",
    /**
     * The VM is suspended. This is similar to hibernation or sleep mode.
     */
    Suspended: "SUSPENDED",
} as const;

/**
 * The power state of the VM at the moment list was taken.
 */
export type VmwareVmDetailsPowerState = (typeof VmwareVmDetailsPowerState)[keyof typeof VmwareVmDetailsPowerState];