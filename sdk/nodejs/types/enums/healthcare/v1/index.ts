// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AttributeDefinitionCategory = {
    /**
     * No category specified. This option is invalid.
     */
    CategoryUnspecified: "CATEGORY_UNSPECIFIED",
    /**
     * Specify this category when this attribute describes the properties of resources. For example, data anonymity or data type.
     */
    Resource: "RESOURCE",
    /**
     * Specify this category when this attribute describes the properties of requests. For example, requester's role or requester's organization.
     */
    Request: "REQUEST",
} as const;

/**
 * Required. The category of the attribute. The value of this field cannot be changed after creation.
 */
export type AttributeDefinitionCategory = (typeof AttributeDefinitionCategory)[keyof typeof AttributeDefinitionCategory];

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

export const ConsentState = {
    /**
     * No state specified. Treated as ACTIVE only at the time of resource creation.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The Consent is active and is considered when evaluating a user's consent on resources.
     */
    Active: "ACTIVE",
    /**
     * When a Consent is updated, the current version is archived and a new one is created with its state set to the updated Consent's previous state.
     */
    Archived: "ARCHIVED",
    /**
     * A revoked Consent is not considered when evaluating a user's consent on resources.
     */
    Revoked: "REVOKED",
    /**
     * A draft Consent is not considered when evaluating a user's consent on resources unless explicitly specified.
     */
    Draft: "DRAFT",
    /**
     * When a draft Consent is rejected by a user, it is set to a rejected state. A rejected Consent is not considered when evaluating a user's consent on resources.
     */
    Rejected: "REJECTED",
} as const;

/**
 * Required. Indicates the current state of this Consent.
 */
export type ConsentState = (typeof ConsentState)[keyof typeof ConsentState];

export const FhirStoreVersion = {
    /**
     * Users must specify a version on store creation or an error is returned.
     */
    VersionUnspecified: "VERSION_UNSPECIFIED",
    /**
     * Draft Standard for Trial Use, [Release 2](https://www.hl7.org/fhir/DSTU2)
     */
    Dstu2: "DSTU2",
    /**
     * Standard for Trial Use, [Release 3](https://www.hl7.org/fhir/STU3)
     */
    Stu3: "STU3",
    /**
     * [Release 4](https://www.hl7.org/fhir/R4)
     */
    R4: "R4",
} as const;

/**
 * Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
 */
export type FhirStoreVersion = (typeof FhirStoreVersion)[keyof typeof FhirStoreVersion];

export const GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition = {
    /**
     * Default behavior is the same as WRITE_EMPTY.
     */
    WriteDispositionUnspecified: "WRITE_DISPOSITION_UNSPECIFIED",
    /**
     * Only export data if the destination tables are empty.
     */
    WriteEmpty: "WRITE_EMPTY",
    /**
     * Erase all existing data in the tables before writing the instances.
     */
    WriteTruncate: "WRITE_TRUNCATE",
    /**
     * Append data to the existing tables.
     */
    WriteAppend: "WRITE_APPEND",
} as const;

/**
 * Determines if existing data in the destination dataset is overwritten, appended to, or not written if the tables contain data. If a write_disposition is specified, the `force` parameter is ignored.
 */
export type GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition = (typeof GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition)[keyof typeof GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition];

export const SchemaConfigSchemaType = {
    /**
     * No schema type specified. This type is unsupported.
     */
    SchemaTypeUnspecified: "SCHEMA_TYPE_UNSPECIFIED",
    /**
     * Analytics schema defined by the FHIR community. See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md. BigQuery only allows a maximum of 10,000 columns per table. Due to this limitation, the server will not generate schemas for fields of type `Resource`, which can hold any resource type. The affected fields are `Parameters.parameter.resource`, `Bundle.entry.resource`, and `Bundle.entry.response.outcome`.
     */
    Analytics: "ANALYTICS",
} as const;

/**
 * Specifies the output schema type. Schema type is required.
 */
export type SchemaConfigSchemaType = (typeof SchemaConfigSchemaType)[keyof typeof SchemaConfigSchemaType];

export const SchemaPackageSchematizedParsingType = {
    /**
     * Unspecified schematized parsing type, equivalent to `SOFT_FAIL`.
     */
    SchematizedParsingTypeUnspecified: "SCHEMATIZED_PARSING_TYPE_UNSPECIFIED",
    /**
     * Messages that fail to parse are still stored and ACKed but a parser error is stored in place of the schematized data.
     */
    SoftFail: "SOFT_FAIL",
    /**
     * Messages that fail to parse are rejected from ingestion/insertion and return an error code.
     */
    HardFail: "HARD_FAIL",
} as const;

/**
 * Determines how messages that fail to parse are handled.
 */
export type SchemaPackageSchematizedParsingType = (typeof SchemaPackageSchematizedParsingType)[keyof typeof SchemaPackageSchematizedParsingType];

export const TypePrimitive = {
    /**
     * Not a primitive.
     */
    PrimitiveUnspecified: "PRIMITIVE_UNSPECIFIED",
    /**
     * String primitive.
     */
    String: "STRING",
    /**
     * Element that can have unschematized children.
     */
    Varies: "VARIES",
    /**
     * Like STRING, but all delimiters below this element are ignored.
     */
    UnescapedString: "UNESCAPED_STRING",
} as const;

/**
 * If this is a primitive type then this field is the type of the primitive For example, STRING. Leave unspecified for composite types.
 */
export type TypePrimitive = (typeof TypePrimitive)[keyof typeof TypePrimitive];