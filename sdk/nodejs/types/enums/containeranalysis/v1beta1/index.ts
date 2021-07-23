// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AliasContextKind = {
    /**
     * Unknown.
     */
    KindUnspecified: "KIND_UNSPECIFIED",
    /**
     * Git tag.
     */
    Fixed: "FIXED",
    /**
     * Git branch.
     */
    Movable: "MOVABLE",
    /**
     * Used to specify non-standard aliases. For example, if a Git repo has a ref named "refs/foo/bar".
     */
    Other: "OTHER",
} as const;

/**
 * The alias kind.
 */
export type AliasContextKind = (typeof AliasContextKind)[keyof typeof AliasContextKind];

export const BuildSignatureKeyType = {
    /**
     * `KeyType` is not set.
     */
    KeyTypeUnspecified: "KEY_TYPE_UNSPECIFIED",
    /**
     * `PGP ASCII Armored` public key.
     */
    PgpAsciiArmored: "PGP_ASCII_ARMORED",
    /**
     * `PKIX PEM` public key.
     */
    PkixPem: "PKIX_PEM",
} as const;

/**
 * The type of the key, either stored in `public_key` or referenced in `key_id`.
 */
export type BuildSignatureKeyType = (typeof BuildSignatureKeyType)[keyof typeof BuildSignatureKeyType];

export const CVSSv3AttackComplexity = {
    AttackComplexityUnspecified: "ATTACK_COMPLEXITY_UNSPECIFIED",
    AttackComplexityLow: "ATTACK_COMPLEXITY_LOW",
    AttackComplexityHigh: "ATTACK_COMPLEXITY_HIGH",
} as const;

export type CVSSv3AttackComplexity = (typeof CVSSv3AttackComplexity)[keyof typeof CVSSv3AttackComplexity];

export const CVSSv3AttackVector = {
    AttackVectorUnspecified: "ATTACK_VECTOR_UNSPECIFIED",
    AttackVectorNetwork: "ATTACK_VECTOR_NETWORK",
    AttackVectorAdjacent: "ATTACK_VECTOR_ADJACENT",
    AttackVectorLocal: "ATTACK_VECTOR_LOCAL",
    AttackVectorPhysical: "ATTACK_VECTOR_PHYSICAL",
} as const;

/**
 * Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments.
 */
export type CVSSv3AttackVector = (typeof CVSSv3AttackVector)[keyof typeof CVSSv3AttackVector];

export const CVSSv3AvailabilityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3AvailabilityImpact = (typeof CVSSv3AvailabilityImpact)[keyof typeof CVSSv3AvailabilityImpact];

export const CVSSv3ConfidentialityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3ConfidentialityImpact = (typeof CVSSv3ConfidentialityImpact)[keyof typeof CVSSv3ConfidentialityImpact];

export const CVSSv3IntegrityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3IntegrityImpact = (typeof CVSSv3IntegrityImpact)[keyof typeof CVSSv3IntegrityImpact];

export const CVSSv3PrivilegesRequired = {
    PrivilegesRequiredUnspecified: "PRIVILEGES_REQUIRED_UNSPECIFIED",
    PrivilegesRequiredNone: "PRIVILEGES_REQUIRED_NONE",
    PrivilegesRequiredLow: "PRIVILEGES_REQUIRED_LOW",
    PrivilegesRequiredHigh: "PRIVILEGES_REQUIRED_HIGH",
} as const;

export type CVSSv3PrivilegesRequired = (typeof CVSSv3PrivilegesRequired)[keyof typeof CVSSv3PrivilegesRequired];

export const CVSSv3Scope = {
    ScopeUnspecified: "SCOPE_UNSPECIFIED",
    ScopeUnchanged: "SCOPE_UNCHANGED",
    ScopeChanged: "SCOPE_CHANGED",
} as const;

export type CVSSv3Scope = (typeof CVSSv3Scope)[keyof typeof CVSSv3Scope];

export const CVSSv3UserInteraction = {
    UserInteractionUnspecified: "USER_INTERACTION_UNSPECIFIED",
    UserInteractionNone: "USER_INTERACTION_NONE",
    UserInteractionRequired: "USER_INTERACTION_REQUIRED",
} as const;

export type CVSSv3UserInteraction = (typeof CVSSv3UserInteraction)[keyof typeof CVSSv3UserInteraction];

export const DeploymentPlatform = {
    /**
     * Unknown.
     */
    PlatformUnspecified: "PLATFORM_UNSPECIFIED",
    /**
     * Google Container Engine.
     */
    Gke: "GKE",
    /**
     * Google App Engine: Flexible Environment.
     */
    Flex: "FLEX",
    /**
     * Custom user-defined platform.
     */
    Custom: "CUSTOM",
} as const;

/**
 * Platform hosting this deployment.
 */
export type DeploymentPlatform = (typeof DeploymentPlatform)[keyof typeof DeploymentPlatform];

export const DiscoveredAnalysisStatus = {
    /**
     * Unknown.
     */
    AnalysisStatusUnspecified: "ANALYSIS_STATUS_UNSPECIFIED",
    /**
     * Resource is known but no action has been taken yet.
     */
    Pending: "PENDING",
    /**
     * Resource is being analyzed.
     */
    Scanning: "SCANNING",
    /**
     * Analysis has finished successfully.
     */
    FinishedSuccess: "FINISHED_SUCCESS",
    /**
     * Analysis has finished unsuccessfully, the analysis itself is in a bad state.
     */
    FinishedFailed: "FINISHED_FAILED",
    /**
     * The resource is known not to be supported
     */
    FinishedUnsupported: "FINISHED_UNSUPPORTED",
} as const;

/**
 * The status of discovery for the resource.
 */
export type DiscoveredAnalysisStatus = (typeof DiscoveredAnalysisStatus)[keyof typeof DiscoveredAnalysisStatus];

export const DiscoveredContinuousAnalysis = {
    /**
     * Unknown.
     */
    ContinuousAnalysisUnspecified: "CONTINUOUS_ANALYSIS_UNSPECIFIED",
    /**
     * The resource is continuously analyzed.
     */
    Active: "ACTIVE",
    /**
     * The resource is ignored for continuous analysis.
     */
    Inactive: "INACTIVE",
} as const;

/**
 * Whether the resource is continuously analyzed.
 */
export type DiscoveredContinuousAnalysis = (typeof DiscoveredContinuousAnalysis)[keyof typeof DiscoveredContinuousAnalysis];

export const DiscoveryAnalysisKind = {
    /**
     * Default value. This value is unused.
     */
    NoteKindUnspecified: "NOTE_KIND_UNSPECIFIED",
    /**
     * The note and occurrence represent a package vulnerability.
     */
    Vulnerability: "VULNERABILITY",
    /**
     * The note and occurrence assert build provenance.
     */
    Build: "BUILD",
    /**
     * This represents an image basis relationship.
     */
    Image: "IMAGE",
    /**
     * This represents a package installed via a package manager.
     */
    Package: "PACKAGE",
    /**
     * The note and occurrence track deployment events.
     */
    Deployment: "DEPLOYMENT",
    /**
     * The note and occurrence track the initial discovery status of a resource.
     */
    Discovery: "DISCOVERY",
    /**
     * This represents a logical "role" that can attest to artifacts.
     */
    Attestation: "ATTESTATION",
    /**
     * This represents an in-toto link.
     */
    Intoto: "INTOTO",
} as const;

/**
 * Required. Immutable. The kind of analysis that is handled by this discovery.
 */
export type DiscoveryAnalysisKind = (typeof DiscoveryAnalysisKind)[keyof typeof DiscoveryAnalysisKind];

export const DistributionArchitecture = {
    /**
     * Unknown architecture.
     */
    ArchitectureUnspecified: "ARCHITECTURE_UNSPECIFIED",
    /**
     * X86 architecture.
     */
    X86: "X86",
    /**
     * X64 architecture.
     */
    X64: "X64",
} as const;

/**
 * The CPU architecture for which packages in this distribution channel were built.
 */
export type DistributionArchitecture = (typeof DistributionArchitecture)[keyof typeof DistributionArchitecture];

export const GenericSignedAttestationContentType = {
    /**
     * `ContentType` is not set.
     */
    ContentTypeUnspecified: "CONTENT_TYPE_UNSPECIFIED",
    /**
     * Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted in `plaintext` is a JSON blob conforming to the linked schema.
     */
    SimpleSigningJson: "SIMPLE_SIGNING_JSON",
} as const;

/**
 * Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
 */
export type GenericSignedAttestationContentType = (typeof GenericSignedAttestationContentType)[keyof typeof GenericSignedAttestationContentType];

export const GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity = {
    /**
     * Unknown.
     */
    SeverityUnspecified: "SEVERITY_UNSPECIFIED",
    /**
     * Minimal severity.
     */
    Minimal: "MINIMAL",
    /**
     * Low severity.
     */
    Low: "LOW",
    /**
     * Medium severity.
     */
    Medium: "MEDIUM",
    /**
     * High severity.
     */
    High: "HIGH",
    /**
     * Critical severity.
     */
    Critical: "CRITICAL",
} as const;

/**
 * The distro assigned severity for this vulnerability when it is available, and note provider assigned severity when distro has not yet assigned a severity for this vulnerability.
 */
export type GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity = (typeof GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity)[keyof typeof GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity];

export const LayerDirective = {
    /**
     * Default value for unsupported/missing directive.
     */
    DirectiveUnspecified: "DIRECTIVE_UNSPECIFIED",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Maintainer: "MAINTAINER",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Run: "RUN",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Cmd: "CMD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Label: "LABEL",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Expose: "EXPOSE",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Env: "ENV",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Add: "ADD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Copy: "COPY",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Entrypoint: "ENTRYPOINT",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Volume: "VOLUME",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    User: "USER",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Workdir: "WORKDIR",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Arg: "ARG",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Onbuild: "ONBUILD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Stopsignal: "STOPSIGNAL",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Healthcheck: "HEALTHCHECK",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Shell: "SHELL",
} as const;

/**
 * Required. The recovered Dockerfile directive used to construct this layer.
 */
export type LayerDirective = (typeof LayerDirective)[keyof typeof LayerDirective];

export const PgpSignedAttestationContentType = {
    /**
     * `ContentType` is not set.
     */
    ContentTypeUnspecified: "CONTENT_TYPE_UNSPECIFIED",
    /**
     * Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted from `signature` is a JSON blob conforming to the linked schema.
     */
    SimpleSigningJson: "SIMPLE_SIGNING_JSON",
} as const;

/**
 * Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
 */
export type PgpSignedAttestationContentType = (typeof PgpSignedAttestationContentType)[keyof typeof PgpSignedAttestationContentType];

export const VersionKind = {
    /**
     * Unknown.
     */
    VersionKindUnspecified: "VERSION_KIND_UNSPECIFIED",
    /**
     * A standard package version.
     */
    Normal: "NORMAL",
    /**
     * A special version representing negative infinity.
     */
    Minimum: "MINIMUM",
    /**
     * A special version representing positive infinity.
     */
    Maximum: "MAXIMUM",
} as const;

/**
 * Required. Distinguishes between sentinel MIN/MAX versions and normal versions.
 */
export type VersionKind = (typeof VersionKind)[keyof typeof VersionKind];

export const VulnerabilitySeverity = {
    /**
     * Unknown.
     */
    SeverityUnspecified: "SEVERITY_UNSPECIFIED",
    /**
     * Minimal severity.
     */
    Minimal: "MINIMAL",
    /**
     * Low severity.
     */
    Low: "LOW",
    /**
     * Medium severity.
     */
    Medium: "MEDIUM",
    /**
     * High severity.
     */
    High: "HIGH",
    /**
     * Critical severity.
     */
    Critical: "CRITICAL",
} as const;

/**
 * Note provider assigned impact of the vulnerability.
 */
export type VulnerabilitySeverity = (typeof VulnerabilitySeverity)[keyof typeof VulnerabilitySeverity];