# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AttributeDefinitionCategory',
    'AuditLogConfigLogType',
    'ConsentState',
    'DicomConfigFilterProfile',
    'FhirStoreComplexDataTypeReferenceParsing',
    'FhirStoreVersion',
    'FieldMetadataAction',
    'GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition',
    'ImageConfigTextRedactionMode',
    'ParserConfigVersion',
    'SchemaConfigSchemaType',
    'SchemaPackageSchematizedParsingType',
    'SchemaPackageUnexpectedSegmentHandling',
    'TimePartitioningType',
    'TypePrimitive',
]


class AttributeDefinitionCategory(str, Enum):
    """
    Required. The category of the attribute. The value of this field cannot be changed after creation.
    """
    CATEGORY_UNSPECIFIED = "CATEGORY_UNSPECIFIED"
    """
    No category specified. This option is invalid.
    """
    RESOURCE = "RESOURCE"
    """
    Specify this category when this attribute describes the properties of resources. For example, data anonymity or data type.
    """
    REQUEST = "REQUEST"
    """
    Specify this category when this attribute describes the properties of requests. For example, requester's role or requester's organization.
    """


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class ConsentState(str, Enum):
    """
    Required. Indicates the current state of this Consent.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    No state specified. Treated as ACTIVE only at the time of resource creation.
    """
    ACTIVE = "ACTIVE"
    """
    The Consent is active and is considered when evaluating a user's consent on resources.
    """
    ARCHIVED = "ARCHIVED"
    """
    The archived state is currently not being used.
    """
    REVOKED = "REVOKED"
    """
    A revoked Consent is not considered when evaluating a user's consent on resources.
    """
    DRAFT = "DRAFT"
    """
    A draft Consent is not considered when evaluating a user's consent on resources unless explicitly specified.
    """
    REJECTED = "REJECTED"
    """
    When a draft Consent is rejected by a user, it is set to a rejected state. A rejected Consent is not considered when evaluating a user's consent on resources.
    """


class DicomConfigFilterProfile(str, Enum):
    """
    Tag filtering profile that determines which tags to keep/remove.
    """
    TAG_FILTER_PROFILE_UNSPECIFIED = "TAG_FILTER_PROFILE_UNSPECIFIED"
    """
    No tag filtration profile provided. Same as KEEP_ALL_PROFILE.
    """
    MINIMAL_KEEP_LIST_PROFILE = "MINIMAL_KEEP_LIST_PROFILE"
    """
    Keep only tags required to produce valid DICOM.
    """
    ATTRIBUTE_CONFIDENTIALITY_BASIC_PROFILE = "ATTRIBUTE_CONFIDENTIALITY_BASIC_PROFILE"
    """
    Remove tags based on DICOM Standard's Attribute Confidentiality Basic Profile (DICOM Standard Edition 2018e) http://dicom.nema.org/medical/dicom/2018e/output/chtml/part15/chapter_E.html.
    """
    KEEP_ALL_PROFILE = "KEEP_ALL_PROFILE"
    """
    Keep all tags.
    """
    DEIDENTIFY_TAG_CONTENTS = "DEIDENTIFY_TAG_CONTENTS"
    """
    Inspects within tag contents and replaces sensitive text. The process can be configured using the TextConfig. Applies to all tags with the following Value Representation names: AE, LO, LT, PN, SH, ST, UC, UT, DA, DT, AS
    """


class FhirStoreComplexDataTypeReferenceParsing(str, Enum):
    """
    Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
    """
    COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED = "COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED"
    """
    No parsing behavior specified. This is the same as DISABLED for backwards compatibility.
    """
    DISABLED = "DISABLED"
    """
    References in complex data types are ignored.
    """
    ENABLED = "ENABLED"
    """
    References in complex data types are parsed.
    """


class FhirStoreVersion(str, Enum):
    """
    Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
    """
    VERSION_UNSPECIFIED = "VERSION_UNSPECIFIED"
    """
    Users must specify a version on store creation or an error is returned.
    """
    DSTU2 = "DSTU2"
    """
    Draft Standard for Trial Use, [Release 2](https://www.hl7.org/fhir/DSTU2)
    """
    STU3 = "STU3"
    """
    Standard for Trial Use, [Release 3](https://www.hl7.org/fhir/STU3)
    """
    R4 = "R4"
    """
    [Release 4](https://www.hl7.org/fhir/R4)
    """


class FieldMetadataAction(str, Enum):
    """
    Deidentify action for one field.
    """
    ACTION_UNSPECIFIED = "ACTION_UNSPECIFIED"
    """
    No action specified.
    """
    TRANSFORM = "TRANSFORM"
    """
    Transform the entire field.
    """
    INSPECT_AND_TRANSFORM = "INSPECT_AND_TRANSFORM"
    """
    Inspect and transform any found PHI.
    """
    DO_NOT_TRANSFORM = "DO_NOT_TRANSFORM"
    """
    Do not transform.
    """


class GoogleCloudHealthcareV1FhirBigQueryDestinationWriteDisposition(str, Enum):
    """
    Determines if existing data in the destination dataset is overwritten, appended to, or not written if the tables contain data. If a write_disposition is specified, the `force` parameter is ignored.
    """
    WRITE_DISPOSITION_UNSPECIFIED = "WRITE_DISPOSITION_UNSPECIFIED"
    """
    Default behavior is the same as WRITE_EMPTY.
    """
    WRITE_EMPTY = "WRITE_EMPTY"
    """
    Only export data if the destination tables are empty.
    """
    WRITE_TRUNCATE = "WRITE_TRUNCATE"
    """
    Erase all existing data in the destination tables before writing the FHIR resources.
    """
    WRITE_APPEND = "WRITE_APPEND"
    """
    Append data to the destination tables.
    """


class ImageConfigTextRedactionMode(str, Enum):
    """
    Determines how to redact text from image.
    """
    TEXT_REDACTION_MODE_UNSPECIFIED = "TEXT_REDACTION_MODE_UNSPECIFIED"
    """
    No text redaction specified. Same as REDACT_NO_TEXT.
    """
    REDACT_ALL_TEXT = "REDACT_ALL_TEXT"
    """
    Redact all text.
    """
    REDACT_SENSITIVE_TEXT = "REDACT_SENSITIVE_TEXT"
    """
    Redact sensitive text. Uses the set of [Default DICOM InfoTypes](https://cloud.google.com/healthcare-api/docs/how-tos/dicom-deidentify#default_dicom_infotypes).
    """
    REDACT_NO_TEXT = "REDACT_NO_TEXT"
    """
    Do not redact text.
    """


class ParserConfigVersion(str, Enum):
    """
    Immutable. Determines the version of both the default parser to be used when `schema` is not given, as well as the schematized parser used when `schema` is specified. This field is immutable after HL7v2 store creation.
    """
    PARSER_VERSION_UNSPECIFIED = "PARSER_VERSION_UNSPECIFIED"
    """
    Unspecified parser version, equivalent to V1.
    """
    V1 = "V1"
    """
    The `parsed_data` includes every given non-empty message field except the Field Separator (MSH-1) field. As a result, the parsed MSH segment starts with the MSH-2 field and the field numbers are off-by-one with respect to the HL7 standard.
    """
    V2 = "V2"
    """
    The `parsed_data` includes every given non-empty message field.
    """
    V3 = "V3"
    """
    This version is the same as V2, with the following change. The `parsed_data` contains unescaped escaped field separators, component separators, sub-component separators, repetition separators, escape characters, and truncation characters. If `schema` is specified, the schematized parser uses improved parsing heuristics compared to previous versions.
    """


class SchemaConfigSchemaType(str, Enum):
    """
    Specifies the output schema type. Schema type is required.
    """
    SCHEMA_TYPE_UNSPECIFIED = "SCHEMA_TYPE_UNSPECIFIED"
    """
    No schema type specified. This type is unsupported.
    """
    ANALYTICS = "ANALYTICS"
    """
    Analytics schema defined by the FHIR community. See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md. BigQuery only allows a maximum of 10,000 columns per table. Due to this limitation, the server will not generate schemas for fields of type `Resource`, which can hold any resource type. The affected fields are `Parameters.parameter.resource`, `Bundle.entry.resource`, and `Bundle.entry.response.outcome`.
    """
    ANALYTICS_V2 = "ANALYTICS_V2"
    """
    Analytics V2, similar to schema defined by the FHIR community, with added support for extensions with one or more occurrences and contained resources in stringified JSON. Analytics V2 uses more space in the destination table than Analytics V1.
    """


class SchemaPackageSchematizedParsingType(str, Enum):
    """
    Determines how messages that fail to parse are handled.
    """
    SCHEMATIZED_PARSING_TYPE_UNSPECIFIED = "SCHEMATIZED_PARSING_TYPE_UNSPECIFIED"
    """
    Unspecified schematized parsing type, equivalent to `SOFT_FAIL`.
    """
    SOFT_FAIL = "SOFT_FAIL"
    """
    Messages that fail to parse are still stored and ACKed but a parser error is stored in place of the schematized data.
    """
    HARD_FAIL = "HARD_FAIL"
    """
    Messages that fail to parse are rejected from ingestion/insertion and return an error code.
    """


class SchemaPackageUnexpectedSegmentHandling(str, Enum):
    """
    Determines how unexpected segments (segments not matched to the schema) are handled.
    """
    UNEXPECTED_SEGMENT_HANDLING_MODE_UNSPECIFIED = "UNEXPECTED_SEGMENT_HANDLING_MODE_UNSPECIFIED"
    """
    Unspecified handling mode, equivalent to FAIL.
    """
    FAIL = "FAIL"
    """
    Unexpected segments fail to parse and return an error.
    """
    SKIP = "SKIP"
    """
    Unexpected segments do not fail, but are omitted from the output.
    """
    PARSE = "PARSE"
    """
    Unexpected segments do not fail, but are parsed in place and added to the current group. If a segment has a type definition, it is used, otherwise it is parsed as VARIES.
    """


class TimePartitioningType(str, Enum):
    """
    Type of partitioning.
    """
    PARTITION_TYPE_UNSPECIFIED = "PARTITION_TYPE_UNSPECIFIED"
    """
    Default unknown time.
    """
    HOUR = "HOUR"
    """
    Data partitioned by hour.
    """
    DAY = "DAY"
    """
    Data partitioned by day.
    """
    MONTH = "MONTH"
    """
    Data partitioned by month.
    """
    YEAR = "YEAR"
    """
    Data partitioned by year.
    """


class TypePrimitive(str, Enum):
    """
    If this is a primitive type then this field is the type of the primitive For example, STRING. Leave unspecified for composite types.
    """
    PRIMITIVE_UNSPECIFIED = "PRIMITIVE_UNSPECIFIED"
    """
    Not a primitive.
    """
    STRING = "STRING"
    """
    String primitive.
    """
    VARIES = "VARIES"
    """
    Element that can have unschematized children.
    """
    UNESCAPED_STRING = "UNESCAPED_STRING"
    """
    Like STRING, but all delimiters below this element are ignored.
    """
