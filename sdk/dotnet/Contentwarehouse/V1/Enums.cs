// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Contentwarehouse.V1
{
    /// <summary>
    /// Indicates the category (image, audio, video etc.) of the original content.
    /// </summary>
    [EnumType]
    public readonly struct DocumentContentCategory : IEquatable<DocumentContentCategory>
    {
        private readonly string _value;

        private DocumentContentCategory(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// No category is specified.
        /// </summary>
        public static DocumentContentCategory ContentCategoryUnspecified { get; } = new DocumentContentCategory("CONTENT_CATEGORY_UNSPECIFIED");
        /// <summary>
        /// Content is of image type.
        /// </summary>
        public static DocumentContentCategory ContentCategoryImage { get; } = new DocumentContentCategory("CONTENT_CATEGORY_IMAGE");
        /// <summary>
        /// Content is of audio type.
        /// </summary>
        public static DocumentContentCategory ContentCategoryAudio { get; } = new DocumentContentCategory("CONTENT_CATEGORY_AUDIO");
        /// <summary>
        /// Content is of video type.
        /// </summary>
        public static DocumentContentCategory ContentCategoryVideo { get; } = new DocumentContentCategory("CONTENT_CATEGORY_VIDEO");

        public static bool operator ==(DocumentContentCategory left, DocumentContentCategory right) => left.Equals(right);
        public static bool operator !=(DocumentContentCategory left, DocumentContentCategory right) => !left.Equals(right);

        public static explicit operator string(DocumentContentCategory value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentContentCategory other && Equals(other);
        public bool Equals(DocumentContentCategory other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This is used when DocAI was not used to load the document and parsing/ extracting is needed for the inline_raw_document. For example, if inline_raw_document is the byte representation of a PDF file, then this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
    /// </summary>
    [EnumType]
    public readonly struct DocumentRawDocumentFileType : IEquatable<DocumentRawDocumentFileType>
    {
        private readonly string _value;

        private DocumentRawDocumentFileType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// No raw document specified or it is non-parsable
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypeUnspecified { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Adobe PDF format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypePdf { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_PDF");
        /// <summary>
        /// Microsoft Word format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypeDocx { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_DOCX");
        /// <summary>
        /// Microsoft Excel format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypeXlsx { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_XLSX");
        /// <summary>
        /// Microsoft Powerpoint format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypePptx { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_PPTX");
        /// <summary>
        /// UTF-8 encoded text format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypeText { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_TEXT");
        /// <summary>
        /// TIFF or TIF image file format
        /// </summary>
        public static DocumentRawDocumentFileType RawDocumentFileTypeTiff { get; } = new DocumentRawDocumentFileType("RAW_DOCUMENT_FILE_TYPE_TIFF");

        public static bool operator ==(DocumentRawDocumentFileType left, DocumentRawDocumentFileType right) => left.Equals(right);
        public static bool operator !=(DocumentRawDocumentFileType left, DocumentRawDocumentFileType right) => !left.Equals(right);

        public static explicit operator string(DocumentRawDocumentFileType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DocumentRawDocumentFileType other && Equals(other);
        public bool Equals(DocumentRawDocumentFileType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Identifies the type of operation.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContentwarehouseV1AccessControlActionOperationType : IEquatable<GoogleCloudContentwarehouseV1AccessControlActionOperationType>
    {
        private readonly string _value;

        private GoogleCloudContentwarehouseV1AccessControlActionOperationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The unknown operation type.
        /// </summary>
        public static GoogleCloudContentwarehouseV1AccessControlActionOperationType Unknown { get; } = new GoogleCloudContentwarehouseV1AccessControlActionOperationType("UNKNOWN");
        /// <summary>
        /// Adds newly given policy bindings in the existing bindings list.
        /// </summary>
        public static GoogleCloudContentwarehouseV1AccessControlActionOperationType AddPolicyBinding { get; } = new GoogleCloudContentwarehouseV1AccessControlActionOperationType("ADD_POLICY_BINDING");
        /// <summary>
        /// Removes newly given policy bindings from the existing bindings list.
        /// </summary>
        public static GoogleCloudContentwarehouseV1AccessControlActionOperationType RemovePolicyBinding { get; } = new GoogleCloudContentwarehouseV1AccessControlActionOperationType("REMOVE_POLICY_BINDING");
        /// <summary>
        /// Replaces existing policy bindings with the given policy binding list
        /// </summary>
        public static GoogleCloudContentwarehouseV1AccessControlActionOperationType ReplacePolicyBinding { get; } = new GoogleCloudContentwarehouseV1AccessControlActionOperationType("REPLACE_POLICY_BINDING");

        public static bool operator ==(GoogleCloudContentwarehouseV1AccessControlActionOperationType left, GoogleCloudContentwarehouseV1AccessControlActionOperationType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContentwarehouseV1AccessControlActionOperationType left, GoogleCloudContentwarehouseV1AccessControlActionOperationType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContentwarehouseV1AccessControlActionOperationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContentwarehouseV1AccessControlActionOperationType other && Equals(other);
        public bool Equals(GoogleCloudContentwarehouseV1AccessControlActionOperationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The retrieval importance of the property during search.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance : IEquatable<GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance>
    {
        private readonly string _value;

        private GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// No importance specified. Default medium importance.
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance RetrievalImportanceUnspecified { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("RETRIEVAL_IMPORTANCE_UNSPECIFIED");
        /// <summary>
        /// Highest importance.
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance Highest { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("HIGHEST");
        /// <summary>
        /// Higher importance.
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance Higher { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("HIGHER");
        /// <summary>
        /// High importance.
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance High { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("HIGH");
        /// <summary>
        /// Medium importance.
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance Medium { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("MEDIUM");
        /// <summary>
        /// Low importance (negative).
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance Low { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("LOW");
        /// <summary>
        /// Lowest importance (negative).
        /// </summary>
        public static GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance Lowest { get; } = new GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance("LOWEST");

        public static bool operator ==(GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance left, GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance left, GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance other && Equals(other);
        public bool Equals(GoogleCloudContentwarehouseV1PropertyDefinitionRetrievalImportance other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Identifies the trigger type for running the policy.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContentwarehouseV1RuleTriggerType : IEquatable<GoogleCloudContentwarehouseV1RuleTriggerType>
    {
        private readonly string _value;

        private GoogleCloudContentwarehouseV1RuleTriggerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Trigger for unknown action.
        /// </summary>
        public static GoogleCloudContentwarehouseV1RuleTriggerType Unknown { get; } = new GoogleCloudContentwarehouseV1RuleTriggerType("UNKNOWN");
        /// <summary>
        /// Trigger for create document action.
        /// </summary>
        public static GoogleCloudContentwarehouseV1RuleTriggerType OnCreate { get; } = new GoogleCloudContentwarehouseV1RuleTriggerType("ON_CREATE");
        /// <summary>
        /// Trigger for update document action.
        /// </summary>
        public static GoogleCloudContentwarehouseV1RuleTriggerType OnUpdate { get; } = new GoogleCloudContentwarehouseV1RuleTriggerType("ON_UPDATE");

        public static bool operator ==(GoogleCloudContentwarehouseV1RuleTriggerType left, GoogleCloudContentwarehouseV1RuleTriggerType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContentwarehouseV1RuleTriggerType left, GoogleCloudContentwarehouseV1RuleTriggerType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContentwarehouseV1RuleTriggerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContentwarehouseV1RuleTriggerType other && Equals(other);
        public bool Equals(GoogleCloudContentwarehouseV1RuleTriggerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type for update.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContentwarehouseV1UpdateOptionsUpdateType : IEquatable<GoogleCloudContentwarehouseV1UpdateOptionsUpdateType>
    {
        private readonly string _value;

        private GoogleCloudContentwarehouseV1UpdateOptionsUpdateType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Defaults to full replace behavior, ie. FULL_REPLACE.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeUnspecified { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Fully replace all the fields (including previously linked raw document). Any field masks will be ignored.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeReplace { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_REPLACE");
        /// <summary>
        /// Merge the fields into the existing entities.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeMerge { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_MERGE");
        /// <summary>
        /// Inserts the properties by names.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeInsertPropertiesByNames { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_INSERT_PROPERTIES_BY_NAMES");
        /// <summary>
        /// Replace the properties by names.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeReplacePropertiesByNames { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_REPLACE_PROPERTIES_BY_NAMES");
        /// <summary>
        /// Delete the properties by names.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeDeletePropertiesByNames { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_DELETE_PROPERTIES_BY_NAMES");
        /// <summary>
        /// For each of the property, replaces the property if the it exists, otherwise inserts a new property. And for the rest of the fields, merge them based on update mask and merge fields options.
        /// </summary>
        public static GoogleCloudContentwarehouseV1UpdateOptionsUpdateType UpdateTypeMergeAndReplaceOrInsertPropertiesByNames { get; } = new GoogleCloudContentwarehouseV1UpdateOptionsUpdateType("UPDATE_TYPE_MERGE_AND_REPLACE_OR_INSERT_PROPERTIES_BY_NAMES");

        public static bool operator ==(GoogleCloudContentwarehouseV1UpdateOptionsUpdateType left, GoogleCloudContentwarehouseV1UpdateOptionsUpdateType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContentwarehouseV1UpdateOptionsUpdateType left, GoogleCloudContentwarehouseV1UpdateOptionsUpdateType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContentwarehouseV1UpdateOptionsUpdateType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContentwarehouseV1UpdateOptionsUpdateType other && Equals(other);
        public bool Equals(GoogleCloudContentwarehouseV1UpdateOptionsUpdateType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. The type of the layout element that is being referenced if any.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType : IEquatable<GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType>
    {
        private readonly string _value;

        private GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Layout Unspecified.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType LayoutTypeUnspecified { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("LAYOUT_TYPE_UNSPECIFIED");
        /// <summary>
        /// References a Page.blocks element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType Block { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("BLOCK");
        /// <summary>
        /// References a Page.paragraphs element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType Paragraph { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("PARAGRAPH");
        /// <summary>
        /// References a Page.lines element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType Line { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("LINE");
        /// <summary>
        /// References a Page.tokens element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType Token { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("TOKEN");
        /// <summary>
        /// References a Page.visual_elements element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType VisualElement { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("VISUAL_ELEMENT");
        /// <summary>
        /// Refrrences a Page.tables element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType Table { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("TABLE");
        /// <summary>
        /// References a Page.form_fields element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType FormField { get; } = new GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType("FORM_FIELD");

        public static bool operator ==(GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType left, GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType left, GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType other && Equals(other);
        public bool Equals(GoogleCloudDocumentaiV1DocumentPageAnchorPageRefLayoutType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Detected orientation for the Layout.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDocumentaiV1DocumentPageLayoutOrientation : IEquatable<GoogleCloudDocumentaiV1DocumentPageLayoutOrientation>
    {
        private readonly string _value;

        private GoogleCloudDocumentaiV1DocumentPageLayoutOrientation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified orientation.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageLayoutOrientation OrientationUnspecified { get; } = new GoogleCloudDocumentaiV1DocumentPageLayoutOrientation("ORIENTATION_UNSPECIFIED");
        /// <summary>
        /// Orientation is aligned with page up.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageLayoutOrientation PageUp { get; } = new GoogleCloudDocumentaiV1DocumentPageLayoutOrientation("PAGE_UP");
        /// <summary>
        /// Orientation is aligned with page right. Turn the head 90 degrees clockwise from upright to read.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageLayoutOrientation PageRight { get; } = new GoogleCloudDocumentaiV1DocumentPageLayoutOrientation("PAGE_RIGHT");
        /// <summary>
        /// Orientation is aligned with page down. Turn the head 180 degrees from upright to read.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageLayoutOrientation PageDown { get; } = new GoogleCloudDocumentaiV1DocumentPageLayoutOrientation("PAGE_DOWN");
        /// <summary>
        /// Orientation is aligned with page left. Turn the head 90 degrees counterclockwise from upright to read.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageLayoutOrientation PageLeft { get; } = new GoogleCloudDocumentaiV1DocumentPageLayoutOrientation("PAGE_LEFT");

        public static bool operator ==(GoogleCloudDocumentaiV1DocumentPageLayoutOrientation left, GoogleCloudDocumentaiV1DocumentPageLayoutOrientation right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDocumentaiV1DocumentPageLayoutOrientation left, GoogleCloudDocumentaiV1DocumentPageLayoutOrientation right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDocumentaiV1DocumentPageLayoutOrientation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDocumentaiV1DocumentPageLayoutOrientation other && Equals(other);
        public bool Equals(GoogleCloudDocumentaiV1DocumentPageLayoutOrientation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Detected break type.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType : IEquatable<GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType>
    {
        private readonly string _value;

        private GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified break type.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType TypeUnspecified { get; } = new GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType("TYPE_UNSPECIFIED");
        /// <summary>
        /// A single whitespace.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType Space { get; } = new GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType("SPACE");
        /// <summary>
        /// A wider whitespace.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType WideSpace { get; } = new GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType("WIDE_SPACE");
        /// <summary>
        /// A hyphen that indicates that a token has been split across lines.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType Hyphen { get; } = new GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType("HYPHEN");

        public static bool operator ==(GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType left, GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType left, GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType other && Equals(other);
        public bool Equals(GoogleCloudDocumentaiV1DocumentPageTokenDetectedBreakType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of provenance operation.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudDocumentaiV1DocumentProvenanceType : IEquatable<GoogleCloudDocumentaiV1DocumentProvenanceType>
    {
        private readonly string _value;

        private GoogleCloudDocumentaiV1DocumentProvenanceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Operation type unspecified. If no operation is specified a provenance entry is simply used to match against a `parent`.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType OperationTypeUnspecified { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("OPERATION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Add an element.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType Add { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("ADD");
        /// <summary>
        /// Remove an element identified by `parent`.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType Remove { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("REMOVE");
        /// <summary>
        /// Updates any fields within the given provenance scope of the message. It overwrites the fields rather than replacing them. Use this when you want to update a field value of an entity without also updating all the child properties.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType Update { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("UPDATE");
        /// <summary>
        /// Currently unused. Replace an element identified by `parent`.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType Replace { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("REPLACE");
        /// <summary>
        /// Deprecated. Request human review for the element identified by `parent`.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType EvalRequested { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("EVAL_REQUESTED");
        /// <summary>
        /// Deprecated. Element is reviewed and approved at human review, confidence will be set to 1.0.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType EvalApproved { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("EVAL_APPROVED");
        /// <summary>
        /// Deprecated. Element is skipped in the validation process.
        /// </summary>
        public static GoogleCloudDocumentaiV1DocumentProvenanceType EvalSkipped { get; } = new GoogleCloudDocumentaiV1DocumentProvenanceType("EVAL_SKIPPED");

        public static bool operator ==(GoogleCloudDocumentaiV1DocumentProvenanceType left, GoogleCloudDocumentaiV1DocumentProvenanceType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudDocumentaiV1DocumentProvenanceType left, GoogleCloudDocumentaiV1DocumentProvenanceType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudDocumentaiV1DocumentProvenanceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudDocumentaiV1DocumentProvenanceType other && Equals(other);
        public bool Equals(GoogleCloudDocumentaiV1DocumentProvenanceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct GoogleIamV1AuditLogConfigLogType : IEquatable<GoogleIamV1AuditLogConfigLogType>
    {
        private readonly string _value;

        private GoogleIamV1AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static GoogleIamV1AuditLogConfigLogType LogTypeUnspecified { get; } = new GoogleIamV1AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static GoogleIamV1AuditLogConfigLogType AdminRead { get; } = new GoogleIamV1AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static GoogleIamV1AuditLogConfigLogType DataWrite { get; } = new GoogleIamV1AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static GoogleIamV1AuditLogConfigLogType DataRead { get; } = new GoogleIamV1AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(GoogleIamV1AuditLogConfigLogType left, GoogleIamV1AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(GoogleIamV1AuditLogConfigLogType left, GoogleIamV1AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(GoogleIamV1AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleIamV1AuditLogConfigLogType other && Equals(other);
        public bool Equals(GoogleIamV1AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
