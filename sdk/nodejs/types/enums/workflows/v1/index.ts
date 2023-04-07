// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const WorkflowCallLogLevel = {
    /**
     * No call logging level specified.
     */
    CallLogLevelUnspecified: "CALL_LOG_LEVEL_UNSPECIFIED",
    /**
     * Log all call steps within workflows, all call returns, and all exceptions raised.
     */
    LogAllCalls: "LOG_ALL_CALLS",
    /**
     * Log only exceptions that are raised from call steps within workflows.
     */
    LogErrorsOnly: "LOG_ERRORS_ONLY",
    /**
     * Explicitly log nothing.
     */
    LogNone: "LOG_NONE",
} as const;

/**
 * Optional. Describes the level of platform logging to apply to calls and call responses during executions of this workflow. If both the workflow and the execution specify a logging level, the execution level takes precedence.
 */
export type WorkflowCallLogLevel = (typeof WorkflowCallLogLevel)[keyof typeof WorkflowCallLogLevel];
