// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const GooglePrivacyDlpV2BigQueryOptionsSampleMethod = {
    SampleMethodUnspecified: "SAMPLE_METHOD_UNSPECIFIED",
    /**
     * Scan groups of rows in the order BigQuery provides (default). Multiple groups of rows may be scanned in parallel, so results may not appear in the same order the rows are read.
     */
    Top: "TOP",
    /**
     * Randomly pick groups of rows to scan.
     */
    RandomStart: "RANDOM_START",
} as const;

export type GooglePrivacyDlpV2BigQueryOptionsSampleMethod = (typeof GooglePrivacyDlpV2BigQueryOptionsSampleMethod)[keyof typeof GooglePrivacyDlpV2BigQueryOptionsSampleMethod];

export const GooglePrivacyDlpV2CharsToIgnoreCommonCharactersToIgnore = {
    /**
     * Unused.
     */
    CommonCharsToIgnoreUnspecified: "COMMON_CHARS_TO_IGNORE_UNSPECIFIED",
    /**
     * 0-9
     */
    Numeric: "NUMERIC",
    /**
     * A-Z
     */
    AlphaUpperCase: "ALPHA_UPPER_CASE",
    /**
     * a-z
     */
    AlphaLowerCase: "ALPHA_LOWER_CASE",
    /**
     * US Punctuation, one of !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
     */
    Punctuation: "PUNCTUATION",
    /**
     * Whitespace character, one of [ \t\n\x0B\f\r]
     */
    Whitespace: "WHITESPACE",
} as const;

/**
 * Common characters to not transform when masking. Useful to avoid removing punctuation.
 */
export type GooglePrivacyDlpV2CharsToIgnoreCommonCharactersToIgnore = (typeof GooglePrivacyDlpV2CharsToIgnoreCommonCharactersToIgnore)[keyof typeof GooglePrivacyDlpV2CharsToIgnoreCommonCharactersToIgnore];

export const GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem = {
    /**
     * Includes all files.
     */
    FileTypeUnspecified: "FILE_TYPE_UNSPECIFIED",
    /**
     * Includes all file extensions not covered by another entry. Binary scanning attempts to convert the content of the file to utf_8 to scan the file. If you wish to avoid this fall back, specify one or more of the other FileType's in your storage scan.
     */
    BinaryFile: "BINARY_FILE",
    /**
     * Included file extensions: asc,asp, aspx, brf, c, cc,cfm, cgi, cpp, csv, cxx, c++, cs, css, dart, dat, dot, eml,, epbub, ged, go, h, hh, hpp, hxx, h++, hs, html, htm, mkd, markdown, m, ml, mli, perl, pl, plist, pm, php, phtml, pht, properties, py, pyw, rb, rbw, rs, rss, rc, scala, sh, sql, swift, tex, shtml, shtm, xhtml, lhs, ics, ini, java, js, json, kix, kml, ocaml, md, txt, text, tsv, vb, vcard, vcs, wml, xcodeproj, xml, xsl, xsd, yml, yaml.
     */
    TextFile: "TEXT_FILE",
    /**
     * Included file extensions: bmp, gif, jpg, jpeg, jpe, png. bytes_limit_per_file has no effect on image files. Image inspection is restricted to 'global', 'us', 'asia', and 'europe'.
     */
    Image: "IMAGE",
    /**
     * Word files >30 MB will be scanned as binary files. Included file extensions: docx, dotx, docm, dotm
     */
    Word: "WORD",
    /**
     * PDF files >30 MB will be scanned as binary files. Included file extensions: pdf
     */
    Pdf: "PDF",
    /**
     * Included file extensions: avro
     */
    Avro: "AVRO",
    /**
     * Included file extensions: csv
     */
    Csv: "CSV",
    /**
     * Included file extensions: tsv
     */
    Tsv: "TSV",
    /**
     * Powerpoint files >30 MB will be scanned as binary files. Included file extensions: pptx, pptm, potx, potm, pot
     */
    Powerpoint: "POWERPOINT",
    /**
     * Excel files >30 MB will be scanned as binary files. Included file extensions: xlsx, xlsm, xltx, xltm
     */
    Excel: "EXCEL",
} as const;

export type GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem = (typeof GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem)[keyof typeof GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem];

export const GooglePrivacyDlpV2CloudStorageOptionsSampleMethod = {
    SampleMethodUnspecified: "SAMPLE_METHOD_UNSPECIFIED",
    /**
     * Scan from the top (default).
     */
    Top: "TOP",
    /**
     * For each file larger than bytes_limit_per_file, randomly pick the offset to start scanning. The scanned bytes are contiguous.
     */
    RandomStart: "RANDOM_START",
} as const;

export type GooglePrivacyDlpV2CloudStorageOptionsSampleMethod = (typeof GooglePrivacyDlpV2CloudStorageOptionsSampleMethod)[keyof typeof GooglePrivacyDlpV2CloudStorageOptionsSampleMethod];

export const GooglePrivacyDlpV2ConditionOperator = {
    /**
     * Unused
     */
    RelationalOperatorUnspecified: "RELATIONAL_OPERATOR_UNSPECIFIED",
    /**
     * Equal. Attempts to match even with incompatible types.
     */
    EqualTo: "EQUAL_TO",
    /**
     * Not equal to. Attempts to match even with incompatible types.
     */
    NotEqualTo: "NOT_EQUAL_TO",
    /**
     * Greater than.
     */
    GreaterThan: "GREATER_THAN",
    /**
     * Less than.
     */
    LessThan: "LESS_THAN",
    /**
     * Greater than or equals.
     */
    GreaterThanOrEquals: "GREATER_THAN_OR_EQUALS",
    /**
     * Less than or equals.
     */
    LessThanOrEquals: "LESS_THAN_OR_EQUALS",
    /**
     * Exists
     */
    Exists: "EXISTS",
} as const;

/**
 * Required. Operator used to compare the field or infoType to the value.
 */
export type GooglePrivacyDlpV2ConditionOperator = (typeof GooglePrivacyDlpV2ConditionOperator)[keyof typeof GooglePrivacyDlpV2ConditionOperator];

export const GooglePrivacyDlpV2CryptoReplaceFfxFpeConfigCommonAlphabet = {
    /**
     * Unused.
     */
    FfxCommonNativeAlphabetUnspecified: "FFX_COMMON_NATIVE_ALPHABET_UNSPECIFIED",
    /**
     * `[0-9]` (radix of 10)
     */
    Numeric: "NUMERIC",
    /**
     * `[0-9A-F]` (radix of 16)
     */
    Hexadecimal: "HEXADECIMAL",
    /**
     * `[0-9A-Z]` (radix of 36)
     */
    UpperCaseAlphaNumeric: "UPPER_CASE_ALPHA_NUMERIC",
    /**
     * `[0-9A-Za-z]` (radix of 62)
     */
    AlphaNumeric: "ALPHA_NUMERIC",
} as const;

/**
 * Common alphabets.
 */
export type GooglePrivacyDlpV2CryptoReplaceFfxFpeConfigCommonAlphabet = (typeof GooglePrivacyDlpV2CryptoReplaceFfxFpeConfigCommonAlphabet)[keyof typeof GooglePrivacyDlpV2CryptoReplaceFfxFpeConfigCommonAlphabet];

export const GooglePrivacyDlpV2CustomInfoTypeExclusionType = {
    /**
     * A finding of this custom info type will not be excluded from results.
     */
    ExclusionTypeUnspecified: "EXCLUSION_TYPE_UNSPECIFIED",
    /**
     * A finding of this custom info type will be excluded from final results, but can still affect rule execution.
     */
    ExclusionTypeExclude: "EXCLUSION_TYPE_EXCLUDE",
} as const;

/**
 * If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching.
 */
export type GooglePrivacyDlpV2CustomInfoTypeExclusionType = (typeof GooglePrivacyDlpV2CustomInfoTypeExclusionType)[keyof typeof GooglePrivacyDlpV2CustomInfoTypeExclusionType];

export const GooglePrivacyDlpV2CustomInfoTypeLikelihood = {
    /**
     * Default value; same as POSSIBLE.
     */
    LikelihoodUnspecified: "LIKELIHOOD_UNSPECIFIED",
    /**
     * Highest chance of a false positive.
     */
    VeryUnlikely: "VERY_UNLIKELY",
    /**
     * High chance of a false positive.
     */
    Unlikely: "UNLIKELY",
    /**
     * Some matching signals. The default value.
     */
    Possible: "POSSIBLE",
    /**
     * Low chance of a false positive.
     */
    Likely: "LIKELY",
    /**
     * Confidence level is high. Lowest chance of a false positive.
     */
    VeryLikely: "VERY_LIKELY",
} as const;

/**
 * Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria specified by the rule. Defaults to `VERY_LIKELY` if not specified.
 */
export type GooglePrivacyDlpV2CustomInfoTypeLikelihood = (typeof GooglePrivacyDlpV2CustomInfoTypeLikelihood)[keyof typeof GooglePrivacyDlpV2CustomInfoTypeLikelihood];

export const GooglePrivacyDlpV2DeidentifyFileTypesToTransformItem = {
    /**
     * Includes all files.
     */
    FileTypeUnspecified: "FILE_TYPE_UNSPECIFIED",
    /**
     * Includes all file extensions not covered by another entry. Binary scanning attempts to convert the content of the file to utf_8 to scan the file. If you wish to avoid this fall back, specify one or more of the other FileType's in your storage scan.
     */
    BinaryFile: "BINARY_FILE",
    /**
     * Included file extensions: asc,asp, aspx, brf, c, cc,cfm, cgi, cpp, csv, cxx, c++, cs, css, dart, dat, dot, eml,, epbub, ged, go, h, hh, hpp, hxx, h++, hs, html, htm, mkd, markdown, m, ml, mli, perl, pl, plist, pm, php, phtml, pht, properties, py, pyw, rb, rbw, rs, rss, rc, scala, sh, sql, swift, tex, shtml, shtm, xhtml, lhs, ics, ini, java, js, json, kix, kml, ocaml, md, txt, text, tsv, vb, vcard, vcs, wml, xcodeproj, xml, xsl, xsd, yml, yaml.
     */
    TextFile: "TEXT_FILE",
    /**
     * Included file extensions: bmp, gif, jpg, jpeg, jpe, png. bytes_limit_per_file has no effect on image files. Image inspection is restricted to 'global', 'us', 'asia', and 'europe'.
     */
    Image: "IMAGE",
    /**
     * Word files >30 MB will be scanned as binary files. Included file extensions: docx, dotx, docm, dotm
     */
    Word: "WORD",
    /**
     * PDF files >30 MB will be scanned as binary files. Included file extensions: pdf
     */
    Pdf: "PDF",
    /**
     * Included file extensions: avro
     */
    Avro: "AVRO",
    /**
     * Included file extensions: csv
     */
    Csv: "CSV",
    /**
     * Included file extensions: tsv
     */
    Tsv: "TSV",
    /**
     * Powerpoint files >30 MB will be scanned as binary files. Included file extensions: pptx, pptm, potx, potm, pot
     */
    Powerpoint: "POWERPOINT",
    /**
     * Excel files >30 MB will be scanned as binary files. Included file extensions: xlsx, xlsm, xltx, xltm
     */
    Excel: "EXCEL",
} as const;

export type GooglePrivacyDlpV2DeidentifyFileTypesToTransformItem = (typeof GooglePrivacyDlpV2DeidentifyFileTypesToTransformItem)[keyof typeof GooglePrivacyDlpV2DeidentifyFileTypesToTransformItem];

export const GooglePrivacyDlpV2ExclusionRuleMatchingType = {
    /**
     * Invalid.
     */
    MatchingTypeUnspecified: "MATCHING_TYPE_UNSPECIFIED",
    /**
     * Full match. - Dictionary: join of Dictionary results matched complete finding quote - Regex: all regex matches fill a finding quote start to end - Exclude info type: completely inside affecting info types findings
     */
    MatchingTypeFullMatch: "MATCHING_TYPE_FULL_MATCH",
    /**
     * Partial match. - Dictionary: at least one of the tokens in the finding matches - Regex: substring of the finding matches - Exclude info type: intersects with affecting info types findings
     */
    MatchingTypePartialMatch: "MATCHING_TYPE_PARTIAL_MATCH",
    /**
     * Inverse match. - Dictionary: no tokens in the finding match the dictionary - Regex: finding doesn't match the regex - Exclude info type: no intersection with affecting info types findings
     */
    MatchingTypeInverseMatch: "MATCHING_TYPE_INVERSE_MATCH",
} as const;

/**
 * How the rule is applied, see MatchingType documentation for details.
 */
export type GooglePrivacyDlpV2ExclusionRuleMatchingType = (typeof GooglePrivacyDlpV2ExclusionRuleMatchingType)[keyof typeof GooglePrivacyDlpV2ExclusionRuleMatchingType];

export const GooglePrivacyDlpV2ExpressionsLogicalOperator = {
    /**
     * Unused
     */
    LogicalOperatorUnspecified: "LOGICAL_OPERATOR_UNSPECIFIED",
    /**
     * Conditional AND
     */
    And: "AND",
} as const;

/**
 * The operator to apply to the result of conditions. Default and currently only supported value is `AND`.
 */
export type GooglePrivacyDlpV2ExpressionsLogicalOperator = (typeof GooglePrivacyDlpV2ExpressionsLogicalOperator)[keyof typeof GooglePrivacyDlpV2ExpressionsLogicalOperator];

export const GooglePrivacyDlpV2InspectConfigContentOptionsItem = {
    /**
     * Includes entire content of a file or a data stream.
     */
    ContentUnspecified: "CONTENT_UNSPECIFIED",
    /**
     * Text content within the data, excluding any metadata.
     */
    ContentText: "CONTENT_TEXT",
    /**
     * Images found in the data.
     */
    ContentImage: "CONTENT_IMAGE",
} as const;

export type GooglePrivacyDlpV2InspectConfigContentOptionsItem = (typeof GooglePrivacyDlpV2InspectConfigContentOptionsItem)[keyof typeof GooglePrivacyDlpV2InspectConfigContentOptionsItem];

export const GooglePrivacyDlpV2InspectConfigMinLikelihood = {
    /**
     * Default value; same as POSSIBLE.
     */
    LikelihoodUnspecified: "LIKELIHOOD_UNSPECIFIED",
    /**
     * Highest chance of a false positive.
     */
    VeryUnlikely: "VERY_UNLIKELY",
    /**
     * High chance of a false positive.
     */
    Unlikely: "UNLIKELY",
    /**
     * Some matching signals. The default value.
     */
    Possible: "POSSIBLE",
    /**
     * Low chance of a false positive.
     */
    Likely: "LIKELY",
    /**
     * Confidence level is high. Lowest chance of a false positive.
     */
    VeryLikely: "VERY_LIKELY",
} as const;

/**
 * Only returns findings equal or above this threshold. The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood to learn more.
 */
export type GooglePrivacyDlpV2InspectConfigMinLikelihood = (typeof GooglePrivacyDlpV2InspectConfigMinLikelihood)[keyof typeof GooglePrivacyDlpV2InspectConfigMinLikelihood];

export const GooglePrivacyDlpV2LikelihoodAdjustmentFixedLikelihood = {
    /**
     * Default value; same as POSSIBLE.
     */
    LikelihoodUnspecified: "LIKELIHOOD_UNSPECIFIED",
    /**
     * Highest chance of a false positive.
     */
    VeryUnlikely: "VERY_UNLIKELY",
    /**
     * High chance of a false positive.
     */
    Unlikely: "UNLIKELY",
    /**
     * Some matching signals. The default value.
     */
    Possible: "POSSIBLE",
    /**
     * Low chance of a false positive.
     */
    Likely: "LIKELY",
    /**
     * Confidence level is high. Lowest chance of a false positive.
     */
    VeryLikely: "VERY_LIKELY",
} as const;

/**
 * Set the likelihood of a finding to a fixed value.
 */
export type GooglePrivacyDlpV2LikelihoodAdjustmentFixedLikelihood = (typeof GooglePrivacyDlpV2LikelihoodAdjustmentFixedLikelihood)[keyof typeof GooglePrivacyDlpV2LikelihoodAdjustmentFixedLikelihood];

export const GooglePrivacyDlpV2OutputStorageConfigOutputSchema = {
    /**
     * Unused.
     */
    OutputSchemaUnspecified: "OUTPUT_SCHEMA_UNSPECIFIED",
    /**
     * Basic schema including only `info_type`, `quote`, `certainty`, and `timestamp`.
     */
    BasicColumns: "BASIC_COLUMNS",
    /**
     * Schema tailored to findings from scanning Cloud Storage.
     */
    GcsColumns: "GCS_COLUMNS",
    /**
     * Schema tailored to findings from scanning Google Datastore.
     */
    DatastoreColumns: "DATASTORE_COLUMNS",
    /**
     * Schema tailored to findings from scanning Google BigQuery.
     */
    BigQueryColumns: "BIG_QUERY_COLUMNS",
    /**
     * Schema containing all columns.
     */
    AllColumns: "ALL_COLUMNS",
} as const;

/**
 * Schema used for writing the findings for Inspect jobs. This field is only used for Inspect and must be unspecified for Risk jobs. Columns are derived from the `Finding` object. If appending to an existing table, any columns from the predefined schema that are missing will be added. No columns in the existing table will be deleted. If unspecified, then all available columns will be used for a new table or an (existing) table with no schema, and no changes will be made to an existing table that has a schema. Only for use with external storage.
 */
export type GooglePrivacyDlpV2OutputStorageConfigOutputSchema = (typeof GooglePrivacyDlpV2OutputStorageConfigOutputSchema)[keyof typeof GooglePrivacyDlpV2OutputStorageConfigOutputSchema];

export const GooglePrivacyDlpV2SensitivityScoreScore = {
    /**
     * Unused.
     */
    SensitivityScoreUnspecified: "SENSITIVITY_SCORE_UNSPECIFIED",
    /**
     * No sensitive information detected. Limited access.
     */
    SensitivityLow: "SENSITIVITY_LOW",
    /**
     * Medium risk - PII, potentially sensitive data, or fields with free-text data that are at higher risk of having intermittent sensitive data. Consider limiting access.
     */
    SensitivityModerate: "SENSITIVITY_MODERATE",
    /**
     * High risk – SPII may be present. Exfiltration of data may lead to user data loss. Re-identification of users may be possible. Consider limiting usage and or removing SPII.
     */
    SensitivityHigh: "SENSITIVITY_HIGH",
} as const;

/**
 * The score applied to the resource.
 */
export type GooglePrivacyDlpV2SensitivityScoreScore = (typeof GooglePrivacyDlpV2SensitivityScoreScore)[keyof typeof GooglePrivacyDlpV2SensitivityScoreScore];

export const GooglePrivacyDlpV2TimePartConfigPartToExtract = {
    /**
     * Unused
     */
    TimePartUnspecified: "TIME_PART_UNSPECIFIED",
    /**
     * [0-9999]
     */
    Year: "YEAR",
    /**
     * [1-12]
     */
    Month: "MONTH",
    /**
     * [1-31]
     */
    DayOfMonth: "DAY_OF_MONTH",
    /**
     * [1-7]
     */
    DayOfWeek: "DAY_OF_WEEK",
    /**
     * [1-53]
     */
    WeekOfYear: "WEEK_OF_YEAR",
    /**
     * [0-23]
     */
    HourOfDay: "HOUR_OF_DAY",
} as const;

/**
 * The part of the time to keep.
 */
export type GooglePrivacyDlpV2TimePartConfigPartToExtract = (typeof GooglePrivacyDlpV2TimePartConfigPartToExtract)[keyof typeof GooglePrivacyDlpV2TimePartConfigPartToExtract];

export const GooglePrivacyDlpV2ValueDayOfWeekValue = {
    /**
     * The day of the week is unspecified.
     */
    DayOfWeekUnspecified: "DAY_OF_WEEK_UNSPECIFIED",
    /**
     * Monday
     */
    Monday: "MONDAY",
    /**
     * Tuesday
     */
    Tuesday: "TUESDAY",
    /**
     * Wednesday
     */
    Wednesday: "WEDNESDAY",
    /**
     * Thursday
     */
    Thursday: "THURSDAY",
    /**
     * Friday
     */
    Friday: "FRIDAY",
    /**
     * Saturday
     */
    Saturday: "SATURDAY",
    /**
     * Sunday
     */
    Sunday: "SUNDAY",
} as const;

/**
 * day of week
 */
export type GooglePrivacyDlpV2ValueDayOfWeekValue = (typeof GooglePrivacyDlpV2ValueDayOfWeekValue)[keyof typeof GooglePrivacyDlpV2ValueDayOfWeekValue];

export const JobTriggerStatus = {
    /**
     * Unused.
     */
    StatusUnspecified: "STATUS_UNSPECIFIED",
    /**
     * Trigger is healthy.
     */
    Healthy: "HEALTHY",
    /**
     * Trigger is temporarily paused.
     */
    Paused: "PAUSED",
    /**
     * Trigger is cancelled and can not be resumed.
     */
    Cancelled: "CANCELLED",
} as const;

/**
 * Required. A status for this trigger.
 */
export type JobTriggerStatus = (typeof JobTriggerStatus)[keyof typeof JobTriggerStatus];

export const OrganizationJobTriggerStatus = {
    /**
     * Unused.
     */
    StatusUnspecified: "STATUS_UNSPECIFIED",
    /**
     * Trigger is healthy.
     */
    Healthy: "HEALTHY",
    /**
     * Trigger is temporarily paused.
     */
    Paused: "PAUSED",
    /**
     * Trigger is cancelled and can not be resumed.
     */
    Cancelled: "CANCELLED",
} as const;

/**
 * Required. A status for this trigger.
 */
export type OrganizationJobTriggerStatus = (typeof OrganizationJobTriggerStatus)[keyof typeof OrganizationJobTriggerStatus];
