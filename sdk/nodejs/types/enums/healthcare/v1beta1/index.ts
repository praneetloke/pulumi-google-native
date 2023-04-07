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

export const DicomConfigFilterProfile = {
    /**
     * No tag filtration profile provided. Same as KEEP_ALL_PROFILE.
     */
    TagFilterProfileUnspecified: "TAG_FILTER_PROFILE_UNSPECIFIED",
    /**
     * Keep only the tags required to produce valid DICOM objects.
     */
    MinimalKeepListProfile: "MINIMAL_KEEP_LIST_PROFILE",
    /**
     * Remove tags based on DICOM Standard's Attribute Confidentiality Basic Profile (DICOM Standard Edition 2018e) http://dicom.nema.org/medical/dicom/2018e/output/chtml/part15/chapter_E.html.
     */
    AttributeConfidentialityBasicProfile: "ATTRIBUTE_CONFIDENTIALITY_BASIC_PROFILE",
    /**
     * Keep all tags.
     */
    KeepAllProfile: "KEEP_ALL_PROFILE",
    /**
     * Inspect within tag contents and replace sensitive text. The process can be configured using the TextConfig. Applies to all tags with the following Value Representation names: AE, LO, LT, PN, SH, ST, UC, UT, DA, DT, AS
     */
    DeidentifyTagContents: "DEIDENTIFY_TAG_CONTENTS",
} as const;

/**
 * Tag filtering profile that determines which tags to keep/remove.
 */
export type DicomConfigFilterProfile = (typeof DicomConfigFilterProfile)[keyof typeof DicomConfigFilterProfile];

export const DicomTagConfigProfileType = {
    /**
     * No profile provided. Same as `ATTRIBUTE_CONFIDENTIALITY_BASIC_PROFILE`.
     */
    ProfileTypeUnspecified: "PROFILE_TYPE_UNSPECIFIED",
    /**
     * Keep only the tags required to produce valid DICOM objects.
     */
    MinimalKeepListProfile: "MINIMAL_KEEP_LIST_PROFILE",
    /**
     * Remove tags based on DICOM Standard's [Attribute Confidentiality Basic Profile (DICOM Standard Edition 2018e)](http://dicom.nema.org/medical/dicom/2018e/output/chtml/part15/chapter_E.html).
     */
    AttributeConfidentialityBasicProfile: "ATTRIBUTE_CONFIDENTIALITY_BASIC_PROFILE",
    /**
     * Keep all tags.
     */
    KeepAllProfile: "KEEP_ALL_PROFILE",
    /**
     * Inspect tag contents and replace sensitive text. The process can be configured using the TextConfig. Applies to all tags with the following [Value Representations] (http://dicom.nema.org/medical/dicom/2018e/output/chtml/part05/sect_6.2.html#table_6.2-1): AE, LO, LT, PN, SH, ST, UC, UT, DA, DT, AS
     */
    DeidentifyTagContents: "DEIDENTIFY_TAG_CONTENTS",
} as const;

/**
 * Base profile type for handling DICOM tags.
 */
export type DicomTagConfigProfileType = (typeof DicomTagConfigProfileType)[keyof typeof DicomTagConfigProfileType];

export const FhirFieldConfigProfileType = {
    /**
     * No profile provided. Same as `BASIC`.
     */
    ProfileTypeUnspecified: "PROFILE_TYPE_UNSPECIFIED",
    /**
     * `Keep` all fields.
     */
    KeepAll: "KEEP_ALL",
    /**
     * Transforms known HIPAA 18 fields and cleans known unstructured text fields.
     */
    Basic: "BASIC",
    /**
     * Cleans all supported tags. Applies to types: Code, Date, DateTime, Decimal, HumanName, Id, LanguageCode, Markdown, Oid, String, Uri, Uuid, Xhtml
     */
    CleanAll: "CLEAN_ALL",
} as const;

/**
 * Base profile type for handling FHIR fields.
 */
export type FhirFieldConfigProfileType = (typeof FhirFieldConfigProfileType)[keyof typeof FhirFieldConfigProfileType];

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
     * VERSION_UNSPECIFIED is treated as STU3 to accommodate the existing FHIR stores.
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

export const FieldMetadataAction = {
    /**
     * No action specified.
     */
    ActionUnspecified: "ACTION_UNSPECIFIED",
    /**
     * Transform the entire field based on transformations specified in TextConfig. When the specified transformation cannot be applied to a field, RedactConfig is used. For example, a Crypto Hash transformation can't be applied to a FHIR Date field.
     */
    Transform: "TRANSFORM",
    /**
     * Inspect and transform any found PHI. When `AnnotationConfig` is provided, annotations of PHI will be generated, except for Date and Datetime.
     */
    InspectAndTransform: "INSPECT_AND_TRANSFORM",
    /**
     * Do not transform.
     */
    DoNotTransform: "DO_NOT_TRANSFORM",
} as const;

/**
 * Deidentify action for one field.
 */
export type FieldMetadataAction = (typeof FieldMetadataAction)[keyof typeof FieldMetadataAction];

export const GoogleCloudHealthcareV1beta1DicomBigQueryDestinationWriteDisposition = {
    /**
     * Default behavior is the same as WRITE_EMPTY.
     */
    WriteDispositionUnspecified: "WRITE_DISPOSITION_UNSPECIFIED",
    /**
     * Only export data if the destination table is empty.
     */
    WriteEmpty: "WRITE_EMPTY",
    /**
     * Erase all existing data in the destination table before writing the instances.
     */
    WriteTruncate: "WRITE_TRUNCATE",
    /**
     * Append data to the destination table.
     */
    WriteAppend: "WRITE_APPEND",
} as const;

/**
 * Determines whether the existing table in the destination is to be overwritten or appended to. If a write_disposition is specified, the `force` parameter is ignored.
 */
export type GoogleCloudHealthcareV1beta1DicomBigQueryDestinationWriteDisposition = (typeof GoogleCloudHealthcareV1beta1DicomBigQueryDestinationWriteDisposition)[keyof typeof GoogleCloudHealthcareV1beta1DicomBigQueryDestinationWriteDisposition];

export const GoogleCloudHealthcareV1beta1FhirBigQueryDestinationWriteDisposition = {
    /**
     * Default behavior is the same as WRITE_EMPTY.
     */
    WriteDispositionUnspecified: "WRITE_DISPOSITION_UNSPECIFIED",
    /**
     * Only export data if the destination tables are empty.
     */
    WriteEmpty: "WRITE_EMPTY",
    /**
     * Erase all existing data in the destination tables before writing the FHIR resources.
     */
    WriteTruncate: "WRITE_TRUNCATE",
    /**
     * Append data to the destination tables.
     */
    WriteAppend: "WRITE_APPEND",
} as const;

/**
 * Determines if existing data in the destination dataset is overwritten, appended to, or not written if the tables contain data. If a write_disposition is specified, the `force` parameter is ignored.
 */
export type GoogleCloudHealthcareV1beta1FhirBigQueryDestinationWriteDisposition = (typeof GoogleCloudHealthcareV1beta1FhirBigQueryDestinationWriteDisposition)[keyof typeof GoogleCloudHealthcareV1beta1FhirBigQueryDestinationWriteDisposition];

export const ImageConfigTextRedactionMode = {
    /**
     * No text redaction specified. Same as REDACT_NO_TEXT.
     */
    TextRedactionModeUnspecified: "TEXT_REDACTION_MODE_UNSPECIFIED",
    /**
     * Redact all text.
     */
    RedactAllText: "REDACT_ALL_TEXT",
    /**
     * Redact sensitive text. Uses the set of [Default DICOM InfoTypes](https://cloud.google.com/healthcare-api/docs/how-tos/dicom-deidentify#default_dicom_infotypes).
     */
    RedactSensitiveText: "REDACT_SENSITIVE_TEXT",
    /**
     * Do not redact text.
     */
    RedactNoText: "REDACT_NO_TEXT",
    /**
     * This mode is like `REDACT_SENSITIVE_TEXT` with the addition of the [Clean Descriptors Option] (https://dicom.nema.org/medical/dicom/2018e/output/chtml/part15/sect_E.3.5.html) enabled: When cleaning text, the process attempts to transform phrases matching any of the tags marked for removal (action codes D, Z, X, and U) in the [Basic Profile] (https://dicom.nema.org/medical/dicom/2018e/output/chtml/part15/chapter_E.html). These contextual phrases are replaced with the token "[CTX]". This mode uses an additional InfoType during inspection.
     */
    RedactSensitiveTextCleanDescriptors: "REDACT_SENSITIVE_TEXT_CLEAN_DESCRIPTORS",
} as const;

/**
 * Determines how to redact text from image.
 */
export type ImageConfigTextRedactionMode = (typeof ImageConfigTextRedactionMode)[keyof typeof ImageConfigTextRedactionMode];

export const OptionsPrimaryIds = {
    /**
     * No value provided. Default to the behavior specified by the base profile.
     */
    PrimaryIdsOptionUnspecified: "PRIMARY_IDS_OPTION_UNSPECIFIED",
    /**
     * Keep primary IDs.
     */
    Keep: "KEEP",
    /**
     * Regenerate primary IDs.
     */
    Regen: "REGEN",
} as const;

/**
 * Set `Action` for [`StudyInstanceUID`, `SeriesInstanceUID`, `SOPInstanceUID`, and `MediaStorageSOPInstanceUID`](http://dicom.nema.org/medical/dicom/2018e/output/chtml/part06/chapter_6.html).
 */
export type OptionsPrimaryIds = (typeof OptionsPrimaryIds)[keyof typeof OptionsPrimaryIds];

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
    /**
     * This version is the same as V2, with the following change. The `parsed_data` contains unescaped escaped field separators, component separators, sub-component separators, repetition separators, escape characters, and truncation characters. If `schema` is specified, the schematized parser uses improved parsing heuristics compared to previous versions.
     */
    V3: "V3",
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
     * A data-driven schema generated from the fields present in the FHIR data being exported, with no additional simplification. This type cannot be used for streaming to BigQuery.
     */
    Lossless: "LOSSLESS",
    /**
     * Analytics schema defined by the FHIR community. See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md. BigQuery only allows a maximum of 10,000 columns per table. Due to this limitation, the server will not generate schemas for fields of type `Resource`, which can hold any resource type. The affected fields are `Parameters.parameter.resource`, `Bundle.entry.resource`, and `Bundle.entry.response.outcome`.
     */
    Analytics: "ANALYTICS",
    /**
     * Analytics V2, similar to schema defined by the FHIR community, with added support for extensions with one or more occurrences and contained resources in stringified JSON. Analytics V2 uses more space in the destination table than Analytics V1.
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

export const TextConfigProfileType = {
    /**
     * Same as BASIC.
     */
    ProfileTypeUnspecified: "PROFILE_TYPE_UNSPECIFIED",
    /**
     * Empty profile which does not perform any transformations.
     */
    Empty: "EMPTY",
    /**
     * Basic profile applies: DATE -> DateShift Default -> ReplaceWithInfoType
     */
    Basic: "BASIC",
} as const;

/**
 * Base profile type for text transformation.
 */
export type TextConfigProfileType = (typeof TextConfigProfileType)[keyof typeof TextConfigProfileType];

export const TimePartitioningType = {
    /**
     * Default unknown time.
     */
    PartitionTypeUnspecified: "PARTITION_TYPE_UNSPECIFIED",
    /**
     * Data partitioned by hour.
     */
    Hour: "HOUR",
    /**
     * Data partitioned by day.
     */
    Day: "DAY",
    /**
     * Data partitioned by month.
     */
    Month: "MONTH",
    /**
     * Data partitioned by year.
     */
    Year: "YEAR",
} as const;

/**
 * Type of partitioning.
 */
export type TimePartitioningType = (typeof TimePartitioningType)[keyof typeof TimePartitioningType];

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
