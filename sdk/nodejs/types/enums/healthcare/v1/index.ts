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
     * The archived state is currently not being used.
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

export const FhirStoreComplexDataTypeReferenceParsing = {
    /**
     * No parsing behavior specified. This is the same as DISABLED for backwards compatibility.
     */
    ComplexDataTypeReferenceParsingUnspecified: "COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED",
    /**
     * References in complex data types are ignored.
     */
    Disabled: "DISABLED",
    /**
     * References in complex data types are parsed.
     */
    Enabled: "ENABLED",
} as const;

/**
 * Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
 */
export type FhirStoreComplexDataTypeReferenceParsing = (typeof FhirStoreComplexDataTypeReferenceParsing)[keyof typeof FhirStoreComplexDataTypeReferenceParsing];

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

export const ParserConfigVersion = {
    /**
     * Unspecified parser version, equivalent to V1.
     */
    ParserVersionUnspecified: "PARSER_VERSION_UNSPECIFIED",
    /**
     * The `parsed_data` includes every given non-empty message field except the Field Separator (MSH-1) field. As a result, the parsed MSH segment starts with the MSH-2 field and the field numbers are off-by-one with respect to the HL7 standard.
     */
    V1: "V1",
    /**
     * The `parsed_data` includes every given non-empty message field.
     */
    V2: "V2",
} as const;

/**
 * Immutable. Determines the version of both the default parser to be used when `schema` is not given, as well as the schematized parser used when `schema` is specified. This field is immutable after HL7v2 store creation.
 */
export type ParserConfigVersion = (typeof ParserConfigVersion)[keyof typeof ParserConfigVersion];

export const SchemaConfigSchemaType = {
    /**
     * No schema type specified. This type is unsupported.
     */
    SchemaTypeUnspecified: "SCHEMA_TYPE_UNSPECIFIED",
    /**
     * Analytics schema defined by the FHIR community. See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md. BigQuery only allows a maximum of 10,000 columns per table. Due to this limitation, the server will not generate schemas for fields of type `Resource`, which can hold any resource type. The affected fields are `Parameters.parameter.resource`, `Bundle.entry.resource`, and `Bundle.entry.response.outcome`.
     */
    Analytics: "ANALYTICS",
    /**
     * Analytics V2, similar to schema defined by the FHIR community, with added support for extensions with one or more occurrences and contained resources in stringified JSON.
     */
    AnalyticsV2: "ANALYTICS_V2",
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

export const SchemaPackageUnexpectedSegmentHandling = {
    /**
     * Unspecified handling mode, equivalent to FAIL.
     */
    UnexpectedSegmentHandlingModeUnspecified: "UNEXPECTED_SEGMENT_HANDLING_MODE_UNSPECIFIED",
    /**
     * Unexpected segments fail to parse and return an error.
     */
    Fail: "FAIL",
    /**
     * Unexpected segments do not fail, but are omitted from the output.
     */
    Skip: "SKIP",
    /**
     * Unexpected segments do not fail, but are parsed in place and added to the current group. If a segment has a type definition, it is used, otherwise it is parsed as VARIES.
     */
    Parse: "PARSE",
} as const;

/**
 * Determines how unexpected segments (segments not matched to the schema) are handled.
 */
export type SchemaPackageUnexpectedSegmentHandling = (typeof SchemaPackageUnexpectedSegmentHandling)[keyof typeof SchemaPackageUnexpectedSegmentHandling];

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
