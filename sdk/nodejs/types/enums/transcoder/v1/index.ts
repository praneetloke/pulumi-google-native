// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AnimationFadeFadeType = {
    /**
     * The fade type is not specified.
     */
    FadeTypeUnspecified: "FADE_TYPE_UNSPECIFIED",
    /**
     * Fade the overlay object into view.
     */
    FadeIn: "FADE_IN",
    /**
     * Fade the overlay object out of view.
     */
    FadeOut: "FADE_OUT",
} as const;

/**
 * Required. Type of fade animation: `FADE_IN` or `FADE_OUT`.
 */
export type AnimationFadeFadeType = (typeof AnimationFadeFadeType)[keyof typeof AnimationFadeFadeType];

export const JobMode = {
    /**
     * The job processing mode is not specified.
     */
    ProcessingModeUnspecified: "PROCESSING_MODE_UNSPECIFIED",
    /**
     * The job processing mode is interactive mode. Interactive job will either be ran or rejected if quota does not allow for it.
     */
    ProcessingModeInteractive: "PROCESSING_MODE_INTERACTIVE",
    /**
     * The job processing mode is batch mode. Batch mode allows queuing of jobs.
     */
    ProcessingModeBatch: "PROCESSING_MODE_BATCH",
} as const;

/**
 * The processing mode of the job. The default is `PROCESSING_MODE_INTERACTIVE`.
 */
export type JobMode = (typeof JobMode)[keyof typeof JobMode];

export const ManifestType = {
    /**
     * The manifest type is not specified.
     */
    ManifestTypeUnspecified: "MANIFEST_TYPE_UNSPECIFIED",
    /**
     * Create `HLS` manifest. The corresponding file extension is `.m3u8`.
     */
    Hls: "HLS",
    /**
     * Create `DASH` manifest. The corresponding file extension is `.mpd`.
     */
    Dash: "DASH",
} as const;

/**
 * Required. Type of the manifest, can be `HLS` or `DASH`.
 */
export type ManifestType = (typeof ManifestType)[keyof typeof ManifestType];
