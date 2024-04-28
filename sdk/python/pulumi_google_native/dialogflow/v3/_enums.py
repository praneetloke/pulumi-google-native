# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'EntityTypeAutoExpansionMode',
    'EntityTypeKind',
    'ExperimentState',
    'GoogleCloudDialogflowCxV3DataStoreConnectionDataStoreType',
    'GoogleCloudDialogflowCxV3ExperimentResultMetricCountType',
    'GoogleCloudDialogflowCxV3ExperimentResultMetricType',
    'GoogleCloudDialogflowCxV3InputAudioConfigAudioEncoding',
    'GoogleCloudDialogflowCxV3InputAudioConfigModelVariant',
    'GoogleCloudDialogflowCxV3NluSettingsModelTrainingMode',
    'GoogleCloudDialogflowCxV3NluSettingsModelType',
    'GoogleCloudDialogflowCxV3ResponseMessageResponseType',
    'GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsAudioFormat',
    'GoogleCloudDialogflowCxV3TestCaseResultTestResult',
    'GoogleCloudDialogflowCxV3VoiceSelectionParamsSsmlGender',
    'GoogleCloudDialogflowCxV3WebhookGenericWebServiceHttpMethod',
    'GoogleCloudDialogflowCxV3WebhookGenericWebServiceWebhookType',
    'SecuritySettingPurgeDataTypesItem',
    'SecuritySettingRedactionScope',
    'SecuritySettingRedactionStrategy',
    'SecuritySettingRetentionStrategy',
    'SessionEntityTypeEntityOverrideMode',
]


class EntityTypeAutoExpansionMode(str, Enum):
    """
    Indicates whether the entity type can be automatically expanded.
    """
    AUTO_EXPANSION_MODE_UNSPECIFIED = "AUTO_EXPANSION_MODE_UNSPECIFIED"
    """
    Auto expansion disabled for the entity.
    """
    AUTO_EXPANSION_MODE_DEFAULT = "AUTO_EXPANSION_MODE_DEFAULT"
    """
    Allows an agent to recognize values that have not been explicitly listed in the entity.
    """


class EntityTypeKind(str, Enum):
    """
    Required. Indicates the kind of entity type.
    """
    KIND_UNSPECIFIED = "KIND_UNSPECIFIED"
    """
    Not specified. This value should be never used.
    """
    KIND_MAP = "KIND_MAP"
    """
    Map entity types allow mapping of a group of synonyms to a canonical value.
    """
    KIND_LIST = "KIND_LIST"
    """
    List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
    """
    KIND_REGEXP = "KIND_REGEXP"
    """
    Regexp entity types allow to specify regular expressions in entries values.
    """


class ExperimentState(str, Enum):
    """
    The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
    """
    STATE_UNSPECIFIED = "STATE_UNSPECIFIED"
    """
    State unspecified.
    """
    DRAFT = "DRAFT"
    """
    The experiment is created but not started yet.
    """
    RUNNING = "RUNNING"
    """
    The experiment is running.
    """
    DONE = "DONE"
    """
    The experiment is done.
    """
    ROLLOUT_FAILED = "ROLLOUT_FAILED"
    """
    The experiment with auto-rollout enabled has failed.
    """


class GoogleCloudDialogflowCxV3DataStoreConnectionDataStoreType(str, Enum):
    """
    The type of the connected data store.
    """
    DATA_STORE_TYPE_UNSPECIFIED = "DATA_STORE_TYPE_UNSPECIFIED"
    """
    Not specified. This value indicates that the data store type is not specified, so it will not be used during search.
    """
    PUBLIC_WEB = "PUBLIC_WEB"
    """
    A data store that contains public web content.
    """
    UNSTRUCTURED = "UNSTRUCTURED"
    """
    A data store that contains unstructured private data.
    """
    STRUCTURED = "STRUCTURED"
    """
    A data store that contains structured data (for example FAQ).
    """


class GoogleCloudDialogflowCxV3ExperimentResultMetricCountType(str, Enum):
    """
    Count-based metric type. Only one of type or count_type is specified in each Metric.
    """
    COUNT_TYPE_UNSPECIFIED = "COUNT_TYPE_UNSPECIFIED"
    """
    Count type unspecified.
    """
    TOTAL_NO_MATCH_COUNT = "TOTAL_NO_MATCH_COUNT"
    """
    Total number of occurrences of a 'NO_MATCH'.
    """
    TOTAL_TURN_COUNT = "TOTAL_TURN_COUNT"
    """
    Total number of turn counts.
    """
    AVERAGE_TURN_COUNT = "AVERAGE_TURN_COUNT"
    """
    Average turn count in a session.
    """


class GoogleCloudDialogflowCxV3ExperimentResultMetricType(str, Enum):
    """
    Ratio-based metric type. Only one of type or count_type is specified in each Metric.
    """
    METRIC_UNSPECIFIED = "METRIC_UNSPECIFIED"
    """
    Metric unspecified.
    """
    CONTAINED_SESSION_NO_CALLBACK_RATE = "CONTAINED_SESSION_NO_CALLBACK_RATE"
    """
    Percentage of contained sessions without user calling back in 24 hours.
    """
    LIVE_AGENT_HANDOFF_RATE = "LIVE_AGENT_HANDOFF_RATE"
    """
    Percentage of sessions that were handed to a human agent.
    """
    CALLBACK_SESSION_RATE = "CALLBACK_SESSION_RATE"
    """
    Percentage of sessions with the same user calling back.
    """
    ABANDONED_SESSION_RATE = "ABANDONED_SESSION_RATE"
    """
    Percentage of sessions where user hung up.
    """
    SESSION_END_RATE = "SESSION_END_RATE"
    """
    Percentage of sessions reached Dialogflow 'END_PAGE' or 'END_SESSION'.
    """


class GoogleCloudDialogflowCxV3InputAudioConfigAudioEncoding(str, Enum):
    """
    Required. Audio encoding of the audio content to process.
    """
    AUDIO_ENCODING_UNSPECIFIED = "AUDIO_ENCODING_UNSPECIFIED"
    """
    Not specified.
    """
    AUDIO_ENCODING_LINEAR16 = "AUDIO_ENCODING_LINEAR_16"
    """
    Uncompressed 16-bit signed little-endian samples (Linear PCM).
    """
    AUDIO_ENCODING_FLAC = "AUDIO_ENCODING_FLAC"
    """
    [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless Audio Codec) is the recommended encoding because it is lossless (therefore recognition is not compromised) and requires only about half the bandwidth of `LINEAR16`. `FLAC` stream encoding supports 16-bit and 24-bit samples, however, not all fields in `STREAMINFO` are supported.
    """
    AUDIO_ENCODING_MULAW = "AUDIO_ENCODING_MULAW"
    """
    8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
    """
    AUDIO_ENCODING_AMR = "AUDIO_ENCODING_AMR"
    """
    Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz` must be 8000.
    """
    AUDIO_ENCODING_AMR_WB = "AUDIO_ENCODING_AMR_WB"
    """
    Adaptive Multi-Rate Wideband codec. `sample_rate_hertz` must be 16000.
    """
    AUDIO_ENCODING_OGG_OPUS = "AUDIO_ENCODING_OGG_OPUS"
    """
    Opus encoded audio frames in Ogg container ([OggOpus](https://wiki.xiph.org/OggOpus)). `sample_rate_hertz` must be 16000.
    """
    AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE = "AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE"
    """
    Although the use of lossy encodings is not recommended, if a very low bitrate encoding is required, `OGG_OPUS` is highly preferred over Speex encoding. The [Speex](https://speex.org/) encoding supported by Dialogflow API has a header byte in each block, as in MIME type `audio/x-speex-with-header-byte`. It is a variant of the RTP Speex encoding defined in [RFC 5574](https://tools.ietf.org/html/rfc5574). The stream is a sequence of blocks, one block per RTP packet. Each block starts with a byte containing the length of the block, in bytes, followed by one or more frames of Speex data, padded to an integral number of bytes (octets) as specified in RFC 5574. In other words, each RTP header is replaced with a single byte containing the block length. Only Speex wideband is supported. `sample_rate_hertz` must be 16000.
    """


class GoogleCloudDialogflowCxV3InputAudioConfigModelVariant(str, Enum):
    """
    Optional. Which variant of the Speech model to use.
    """
    SPEECH_MODEL_VARIANT_UNSPECIFIED = "SPEECH_MODEL_VARIANT_UNSPECIFIED"
    """
    No model variant specified. In this case Dialogflow defaults to USE_BEST_AVAILABLE.
    """
    USE_BEST_AVAILABLE = "USE_BEST_AVAILABLE"
    """
    Use the best available variant of the Speech model that the caller is eligible for. Please see the [Dialogflow docs](https://cloud.google.com/dialogflow/docs/data-logging) for how to make your project eligible for enhanced models.
    """
    USE_STANDARD = "USE_STANDARD"
    """
    Use standard model variant even if an enhanced model is available. See the [Cloud Speech documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models) for details about enhanced models.
    """
    USE_ENHANCED = "USE_ENHANCED"
    """
    Use an enhanced model variant: * If an enhanced variant does not exist for the given model and request language, Dialogflow falls back to the standard variant. The [Cloud Speech documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models) describes which models have enhanced variants. * If the API caller isn't eligible for enhanced models, Dialogflow returns an error. Please see the [Dialogflow docs](https://cloud.google.com/dialogflow/docs/data-logging) for how to make your project eligible.
    """


class GoogleCloudDialogflowCxV3NluSettingsModelTrainingMode(str, Enum):
    """
    Indicates NLU model training mode.
    """
    MODEL_TRAINING_MODE_UNSPECIFIED = "MODEL_TRAINING_MODE_UNSPECIFIED"
    """
    Not specified. `MODEL_TRAINING_MODE_AUTOMATIC` will be used.
    """
    MODEL_TRAINING_MODE_AUTOMATIC = "MODEL_TRAINING_MODE_AUTOMATIC"
    """
    NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
    """
    MODEL_TRAINING_MODE_MANUAL = "MODEL_TRAINING_MODE_MANUAL"
    """
    User needs to manually trigger NLU model training. Best for large flows whose models take long time to train.
    """


class GoogleCloudDialogflowCxV3NluSettingsModelType(str, Enum):
    """
    Indicates the type of NLU model.
    """
    MODEL_TYPE_UNSPECIFIED = "MODEL_TYPE_UNSPECIFIED"
    """
    Not specified. `MODEL_TYPE_STANDARD` will be used.
    """
    MODEL_TYPE_STANDARD = "MODEL_TYPE_STANDARD"
    """
    Use standard NLU model.
    """
    MODEL_TYPE_ADVANCED = "MODEL_TYPE_ADVANCED"
    """
    Use advanced NLU model.
    """


class GoogleCloudDialogflowCxV3ResponseMessageResponseType(str, Enum):
    """
    Response type.
    """
    RESPONSE_TYPE_UNSPECIFIED = "RESPONSE_TYPE_UNSPECIFIED"
    """
    Not specified.
    """
    ENTRY_PROMPT = "ENTRY_PROMPT"
    """
    The response is from an entry prompt in the page.
    """
    PARAMETER_PROMPT = "PARAMETER_PROMPT"
    """
    The response is from form-filling prompt in the page.
    """
    HANDLER_PROMPT = "HANDLER_PROMPT"
    """
    The response is from a transition route or an event handler in the page or flow or transition route group.
    """


class GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsAudioFormat(str, Enum):
    """
    File format for exported audio file. Currently only in telephony recordings.
    """
    AUDIO_FORMAT_UNSPECIFIED = "AUDIO_FORMAT_UNSPECIFIED"
    """
    Unspecified. Do not use.
    """
    MULAW = "MULAW"
    """
    G.711 mu-law PCM with 8kHz sample rate.
    """
    MP3 = "MP3"
    """
    MP3 file format.
    """
    OGG = "OGG"
    """
    OGG Vorbis.
    """


class GoogleCloudDialogflowCxV3TestCaseResultTestResult(str, Enum):
    """
    Whether the test case passed in the agent environment.
    """
    TEST_RESULT_UNSPECIFIED = "TEST_RESULT_UNSPECIFIED"
    """
    Not specified. Should never be used.
    """
    PASSED = "PASSED"
    """
    The test passed.
    """
    FAILED = "FAILED"
    """
    The test did not pass.
    """


class GoogleCloudDialogflowCxV3VoiceSelectionParamsSsmlGender(str, Enum):
    """
    Optional. The preferred gender of the voice. If not set, the service will choose a voice based on the other parameters such as language_code and name. Note that this is only a preference, not requirement. If a voice of the appropriate gender is not available, the synthesizer substitutes a voice with a different gender rather than failing the request.
    """
    SSML_VOICE_GENDER_UNSPECIFIED = "SSML_VOICE_GENDER_UNSPECIFIED"
    """
    An unspecified gender, which means that the client doesn't care which gender the selected voice will have.
    """
    SSML_VOICE_GENDER_MALE = "SSML_VOICE_GENDER_MALE"
    """
    A male voice.
    """
    SSML_VOICE_GENDER_FEMALE = "SSML_VOICE_GENDER_FEMALE"
    """
    A female voice.
    """
    SSML_VOICE_GENDER_NEUTRAL = "SSML_VOICE_GENDER_NEUTRAL"
    """
    A gender-neutral voice.
    """


class GoogleCloudDialogflowCxV3WebhookGenericWebServiceHttpMethod(str, Enum):
    """
    Optional. HTTP method for the flexible webhook calls. Standard webhook always uses POST.
    """
    HTTP_METHOD_UNSPECIFIED = "HTTP_METHOD_UNSPECIFIED"
    """
    HTTP method not specified.
    """
    POST = "POST"
    """
    HTTP POST Method.
    """
    GET = "GET"
    """
    HTTP GET Method.
    """
    HEAD = "HEAD"
    """
    HTTP HEAD Method.
    """
    PUT = "PUT"
    """
    HTTP PUT Method.
    """
    DELETE = "DELETE"
    """
    HTTP DELETE Method.
    """
    PATCH = "PATCH"
    """
    HTTP PATCH Method.
    """
    OPTIONS = "OPTIONS"
    """
    HTTP OPTIONS Method.
    """


class GoogleCloudDialogflowCxV3WebhookGenericWebServiceWebhookType(str, Enum):
    """
    Optional. Type of the webhook.
    """
    WEBHOOK_TYPE_UNSPECIFIED = "WEBHOOK_TYPE_UNSPECIFIED"
    """
    Default value. This value is unused.
    """
    STANDARD = "STANDARD"
    """
    Represents a standard webhook.
    """
    FLEXIBLE = "FLEXIBLE"
    """
    Represents a flexible webhook.
    """


class SecuritySettingPurgeDataTypesItem(str, Enum):
    PURGE_DATA_TYPE_UNSPECIFIED = "PURGE_DATA_TYPE_UNSPECIFIED"
    """
    Unspecified. Do not use.
    """
    DIALOGFLOW_HISTORY = "DIALOGFLOW_HISTORY"
    """
    Dialogflow history. This does not include Cloud logging, which is owned by the user - not Dialogflow.
    """


class SecuritySettingRedactionScope(str, Enum):
    """
    Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
    """
    REDACTION_SCOPE_UNSPECIFIED = "REDACTION_SCOPE_UNSPECIFIED"
    """
    Don't redact any kind of data.
    """
    REDACT_DISK_STORAGE = "REDACT_DISK_STORAGE"
    """
    On data to be written to disk or similar devices that are capable of holding data even if power is disconnected. This includes data that are temporarily saved on disk.
    """


class SecuritySettingRedactionStrategy(str, Enum):
    """
    Strategy that defines how we do redaction.
    """
    REDACTION_STRATEGY_UNSPECIFIED = "REDACTION_STRATEGY_UNSPECIFIED"
    """
    Do not redact.
    """
    REDACT_WITH_SERVICE = "REDACT_WITH_SERVICE"
    """
    Call redaction service to clean up the data to be persisted.
    """


class SecuritySettingRetentionStrategy(str, Enum):
    """
    Specifies the retention behavior defined by SecuritySettings.RetentionStrategy.
    """
    RETENTION_STRATEGY_UNSPECIFIED = "RETENTION_STRATEGY_UNSPECIFIED"
    """
    Retains the persisted data with Dialogflow's internal default 365d TTLs.
    """
    REMOVE_AFTER_CONVERSATION = "REMOVE_AFTER_CONVERSATION"
    """
    Removes data when the conversation ends. If there is no Conversation explicitly established, a default conversation ends when the corresponding Dialogflow session ends.
    """


class SessionEntityTypeEntityOverrideMode(str, Enum):
    """
    Required. Indicates whether the additional data should override or supplement the custom entity type definition.
    """
    ENTITY_OVERRIDE_MODE_UNSPECIFIED = "ENTITY_OVERRIDE_MODE_UNSPECIFIED"
    """
    Not specified. This value should be never used.
    """
    ENTITY_OVERRIDE_MODE_OVERRIDE = "ENTITY_OVERRIDE_MODE_OVERRIDE"
    """
    The collection of session entities overrides the collection of entities in the corresponding custom entity type.
    """
    ENTITY_OVERRIDE_MODE_SUPPLEMENT = "ENTITY_OVERRIDE_MODE_SUPPLEMENT"
    """
    The collection of session entities extends the collection of entities in the corresponding custom entity type. Note: Even in this override mode calls to `ListSessionEntityTypes`, `GetSessionEntityType`, `CreateSessionEntityType` and `UpdateSessionEntityType` only return the additional entities added in this session entity type. If you want to get the supplemented list, please call EntityTypes.GetEntityType on the custom entity type and merge.
    """
