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

export const PubsubConfigMessageFormat = {
    /**
     * Unspecified.
     */
    MessageFormatUnspecified: "MESSAGE_FORMAT_UNSPECIFIED",
    /**
     * The message payload is a serialized protocol buffer of SourceRepoEvent.
     */
    Protobuf: "PROTOBUF",
    /**
     * The message payload is a JSON string of SourceRepoEvent.
     */
    Json: "JSON",
} as const;

/**
 * The format of the Cloud Pub/Sub messages.
 */
export type PubsubConfigMessageFormat = (typeof PubsubConfigMessageFormat)[keyof typeof PubsubConfigMessageFormat];
