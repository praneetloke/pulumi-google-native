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

export const AuthorizationLoggingOptionsPermissionType = {
    /**
     * Default. Should not be used.
     */
    PermissionTypeUnspecified: "PERMISSION_TYPE_UNSPECIFIED",
    /**
     * A read of admin (meta) data.
     */
    AdminRead: "ADMIN_READ",
    /**
     * A write of admin (meta) data.
     */
    AdminWrite: "ADMIN_WRITE",
    /**
     * A read of standard data.
     */
    DataRead: "DATA_READ",
    /**
     * A write of standard data.
     */
    DataWrite: "DATA_WRITE",
} as const;

/**
 * The type of the permission that was checked.
 */
export type AuthorizationLoggingOptionsPermissionType = (typeof AuthorizationLoggingOptionsPermissionType)[keyof typeof AuthorizationLoggingOptionsPermissionType];

export const CloudAuditOptionsLogName = {
    /**
     * Default. Should not be used.
     */
    UnspecifiedLogName: "UNSPECIFIED_LOG_NAME",
    /**
     * Corresponds to "cloudaudit.googleapis.com/activity"
     */
    AdminActivity: "ADMIN_ACTIVITY",
    /**
     * Corresponds to "cloudaudit.googleapis.com/data_access"
     */
    DataAccess: "DATA_ACCESS",
} as const;

/**
 * The log_name to populate in the Cloud Audit Record.
 */
export type CloudAuditOptionsLogName = (typeof CloudAuditOptionsLogName)[keyof typeof CloudAuditOptionsLogName];

export const ConditionIam = {
    /**
     * Default non-attribute.
     */
    NoAttr: "NO_ATTR",
    /**
     * Either principal or (if present) authority selector.
     */
    Authority: "AUTHORITY",
    /**
     * The principal (even if an authority selector is present), which must only be used for attribution, not authorization.
     */
    Attribution: "ATTRIBUTION",
    /**
     * Any of the security realms in the IAMContext (go/security-realms). When used with IN, the condition indicates "any of the request's realms match one of the given values; with NOT_IN, "none of the realms match any of the given values". Note that a value can be: - 'self' (i.e., allow connections from clients that are in the same security realm, which is currently but not guaranteed to be campus-sized) - 'self:metro' (i.e., clients that are in the same metro) - 'self:cloud-region' (i.e., allow connections from clients that are in the same cloud region) - 'guardians' (i.e., allow connections from its guardian realms. See go/security-realms-glossary#guardian for more information.) - a realm (e.g., 'campus-abc') - a realm group (e.g., 'realms-for-borg-cell-xx', see: go/realm-groups) A match is determined by a realm group membership check performed by a RealmAclRep object (go/realm-acl-howto). It is not permitted to grant access based on the *absence* of a realm, so realm conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
     */
    SecurityRealm: "SECURITY_REALM",
    /**
     * An approver (distinct from the requester) that has authorized this request. When used with IN, the condition indicates that one of the approvers associated with the request matches the specified principal, or is a member of the specified group. Approvers can only grant additional access, and are thus only used in a strictly positive context (e.g. ALLOW/IN or DENY/NOT_IN).
     */
    Approver: "APPROVER",
    /**
     * What types of justifications have been supplied with this request. String values should match enum names from security.credentials.JustificationType, e.g. "MANUAL_STRING". It is not permitted to grant access based on the *absence* of a justification, so justification conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN). Multiple justifications, e.g., a Buganizer ID and a manually-entered reason, are normal and supported.
     */
    JustificationType: "JUSTIFICATION_TYPE",
    /**
     * What type of credentials have been supplied with this request. String values should match enum names from security_loas_l2.CredentialsType - currently, only CREDS_TYPE_EMERGENCY is supported. It is not permitted to grant access based on the *absence* of a credentials type, so the conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
     */
    CredentialsType: "CREDENTIALS_TYPE",
    /**
     * EXPERIMENTAL -- DO NOT USE. The conditions can only be used in a "positive" context (e.g., ALLOW/IN or DENY/NOT_IN).
     */
    CredsAssertion: "CREDS_ASSERTION",
} as const;

/**
 * Trusted attributes supplied by the IAM system.
 */
export type ConditionIam = (typeof ConditionIam)[keyof typeof ConditionIam];

export const ConditionOp = {
    /**
     * Default no-op.
     */
    NoOp: "NO_OP",
    /**
     * DEPRECATED. Use IN instead.
     */
    Equals: "EQUALS",
    /**
     * DEPRECATED. Use NOT_IN instead.
     */
    NotEquals: "NOT_EQUALS",
    /**
     * The condition is true if the subject (or any element of it if it is a set) matches any of the supplied values.
     */
    In: "IN",
    /**
     * The condition is true if the subject (or every element of it if it is a set) matches none of the supplied values.
     */
    NotIn: "NOT_IN",
    /**
     * Subject is discharged
     */
    Discharged: "DISCHARGED",
} as const;

/**
 * An operator to apply the subject with.
 */
export type ConditionOp = (typeof ConditionOp)[keyof typeof ConditionOp];

export const ConditionSys = {
    /**
     * Default non-attribute type
     */
    NoAttr: "NO_ATTR",
    /**
     * Region of the resource
     */
    Region: "REGION",
    /**
     * Service name
     */
    Service: "SERVICE",
    /**
     * Resource name
     */
    Name: "NAME",
    /**
     * IP address of the caller
     */
    Ip: "IP",
} as const;

/**
 * Trusted attributes supplied by any service that owns resources and uses the IAM system for access control.
 */
export type ConditionSys = (typeof ConditionSys)[keyof typeof ConditionSys];

export const DataAccessOptionsLogMode = {
    /**
     * Client is not required to write a partial Gin log immediately after the authorization check. If client chooses to write one and it fails, client may either fail open (allow the operation to continue) or fail closed (handle as a DENY outcome).
     */
    LogModeUnspecified: "LOG_MODE_UNSPECIFIED",
    /**
     * The application's operation in the context of which this authorization check is being made may only be performed if it is successfully logged to Gin. For instance, the authorization library may satisfy this obligation by emitting a partial log entry at authorization check time and only returning ALLOW to the application if it succeeds. If a matching Rule has this directive, but the client has not indicated that it will honor such requirements, then the IAM check will result in authorization failure by setting CheckPolicyResponse.success=false.
     */
    LogFailClosed: "LOG_FAIL_CLOSED",
} as const;

export type DataAccessOptionsLogMode = (typeof DataAccessOptionsLogMode)[keyof typeof DataAccessOptionsLogMode];

export const GameServerClusterAllocationPriority = {
    /**
     * The default allocation priority. `PRIORITY_UNSPECIFIED` is the lowest possible priority.
     */
    PriorityUnspecified: "PRIORITY_UNSPECIFIED",
    /**
     * Priority 1, the highest priority.
     */
    P1: "P1",
    /**
     * Priority 2.
     */
    P2: "P2",
    /**
     * Priority 3.
     */
    P3: "P3",
    /**
     * Priority 4.
     */
    P4: "P4",
} as const;

/**
 * Optional. The allocation priority assigned to the game server cluster. Game server clusters receive new game server allocations based on the relative allocation priorites set for each cluster, if the realm is configured for multicluster allocation.
 */
export type GameServerClusterAllocationPriority = (typeof GameServerClusterAllocationPriority)[keyof typeof GameServerClusterAllocationPriority];

export const RuleAction = {
    /**
     * Default no action.
     */
    NoAction: "NO_ACTION",
    /**
     * Matching 'Entries' grant access.
     */
    Allow: "ALLOW",
    /**
     * Matching 'Entries' grant access and the caller promises to log the request per the returned log_configs.
     */
    AllowWithLog: "ALLOW_WITH_LOG",
    /**
     * Matching 'Entries' deny access.
     */
    Deny: "DENY",
    /**
     * Matching 'Entries' deny access and the caller promises to log the request per the returned log_configs.
     */
    DenyWithLog: "DENY_WITH_LOG",
    /**
     * Matching 'Entries' tell IAM.Check callers to generate logs.
     */
    Log: "LOG",
} as const;

/**
 * Required
 */
export type RuleAction = (typeof RuleAction)[keyof typeof RuleAction];