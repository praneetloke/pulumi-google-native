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

export const DataMaskingPolicyPredefinedExpression = {
    /**
     * Default, unspecified predefined expression. No masking will take place since no expression is specified.
     */
    PredefinedExpressionUnspecified: "PREDEFINED_EXPRESSION_UNSPECIFIED",
    /**
     * Masking expression to replace data with SHA-256 hash.
     */
    Sha256: "SHA256",
    /**
     * Masking expression to replace data with NULLs.
     */
    AlwaysNull: "ALWAYS_NULL",
    /**
     * Masking expression to replace data with their default masking values. The default masking values for each type listed as below: * STRING: "" * BYTES: b'' * INTEGER: 0 * FLOAT: 0.0 * NUMERIC: 0 * BOOLEAN: FALSE * TIMESTAMP: 1970-01-01 00:00:00 UTC * DATE: 1970-01-01 * TIME: 00:00:00 * DATETIME: 1970-01-01T00:00:00 * GEOGRAPHY: POINT(0 0) * BIGNUMERIC: 0 * ARRAY: [] * STRUCT: NOT_APPLICABLE * JSON: NULL
     */
    DefaultMaskingValue: "DEFAULT_MASKING_VALUE",
    /**
     * Masking expression shows the last four characters of text. The masking behavior is as follows: * If text length > 4 characters: Replace text with XXXXX, append last four characters of original text. * If text length <= 4 characters: Apply SHA-256 hash.
     */
    LastFourCharacters: "LAST_FOUR_CHARACTERS",
    /**
     * Masking expression shows the first four characters of text. The masking behavior is as follows: * If text length > 4 characters: Replace text with XXXXX, prepend first four characters of original text. * If text length <= 4 characters: Apply SHA-256 hash.
     */
    FirstFourCharacters: "FIRST_FOUR_CHARACTERS",
    /**
     * Masking expression for email addresses. The masking behavior is as follows: * Syntax-valid email address: Replace username with XXXXX. For example, cloudysanfrancisco@gmail.com becomes XXXXX@gmail.com. * Syntax-invalid email address: Apply SHA-256 hash. For more information, see Email mask.
     */
    EmailMask: "EMAIL_MASK",
    /**
     * Masking expression to only show the *year* of `Date`, `DateTime` and `TimeStamp`. For example, with the year 2076: * DATE : 2076-01-01 * DATETIME : 2076-01-01T00:00:00 * TIMESTAMP : 2076-01-01 00:00:00 UTC Truncation occurs according to the UTC time zone. To change this, adjust the default time zone using the `time_zone` system variable. For more information, see the System variables reference.
     */
    DateYearMask: "DATE_YEAR_MASK",
} as const;

/**
 * A predefined masking expression.
 */
export type DataMaskingPolicyPredefinedExpression = (typeof DataMaskingPolicyPredefinedExpression)[keyof typeof DataMaskingPolicyPredefinedExpression];

export const DataPolicyDataPolicyType = {
    /**
     * Default value for the data policy type. This should not be used.
     */
    DataPolicyTypeUnspecified: "DATA_POLICY_TYPE_UNSPECIFIED",
    /**
     * Used to create a data policy for column-level security, without data masking.
     */
    ColumnLevelSecurityPolicy: "COLUMN_LEVEL_SECURITY_POLICY",
    /**
     * Used to create a data policy for data masking.
     */
    DataMaskingPolicy: "DATA_MASKING_POLICY",
} as const;

/**
 * Type of data policy.
 */
export type DataPolicyDataPolicyType = (typeof DataPolicyDataPolicyType)[keyof typeof DataPolicyDataPolicyType];