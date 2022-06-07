// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Alpha
{
    [EnumType]
    public readonly struct ScanConfigTargetPlatformsItem : IEquatable<ScanConfigTargetPlatformsItem>
    {
        private readonly string _value;

        private ScanConfigTargetPlatformsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The target platform is unknown. Requests with this enum value will be rejected with INVALID_ARGUMENT error.
        /// </summary>
        public static ScanConfigTargetPlatformsItem TargetPlatformUnspecified { get; } = new ScanConfigTargetPlatformsItem("TARGET_PLATFORM_UNSPECIFIED");
        /// <summary>
        /// Google App Engine service.
        /// </summary>
        public static ScanConfigTargetPlatformsItem AppEngine { get; } = new ScanConfigTargetPlatformsItem("APP_ENGINE");
        /// <summary>
        /// Google Compute Engine service.
        /// </summary>
        public static ScanConfigTargetPlatformsItem Compute { get; } = new ScanConfigTargetPlatformsItem("COMPUTE");
        /// <summary>
        /// Google Cloud Run service.
        /// </summary>
        public static ScanConfigTargetPlatformsItem CloudRun { get; } = new ScanConfigTargetPlatformsItem("CLOUD_RUN");
        /// <summary>
        /// Google Cloud Function service.
        /// </summary>
        public static ScanConfigTargetPlatformsItem CloudFunctions { get; } = new ScanConfigTargetPlatformsItem("CLOUD_FUNCTIONS");

        public static bool operator ==(ScanConfigTargetPlatformsItem left, ScanConfigTargetPlatformsItem right) => left.Equals(right);
        public static bool operator !=(ScanConfigTargetPlatformsItem left, ScanConfigTargetPlatformsItem right) => !left.Equals(right);

        public static explicit operator string(ScanConfigTargetPlatformsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScanConfigTargetPlatformsItem other && Equals(other);
        public bool Equals(ScanConfigTargetPlatformsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The user agent used during scanning.
    /// </summary>
    [EnumType]
    public readonly struct ScanConfigUserAgent : IEquatable<ScanConfigUserAgent>
    {
        private readonly string _value;

        private ScanConfigUserAgent(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The user agent is unknown. Service will default to CHROME_LINUX.
        /// </summary>
        public static ScanConfigUserAgent UserAgentUnspecified { get; } = new ScanConfigUserAgent("USER_AGENT_UNSPECIFIED");
        /// <summary>
        /// Chrome on Linux. This is the service default if unspecified.
        /// </summary>
        public static ScanConfigUserAgent ChromeLinux { get; } = new ScanConfigUserAgent("CHROME_LINUX");
        /// <summary>
        /// Chrome on Android.
        /// </summary>
        public static ScanConfigUserAgent ChromeAndroid { get; } = new ScanConfigUserAgent("CHROME_ANDROID");
        /// <summary>
        /// Safari on IPhone.
        /// </summary>
        public static ScanConfigUserAgent SafariIphone { get; } = new ScanConfigUserAgent("SAFARI_IPHONE");

        public static bool operator ==(ScanConfigUserAgent left, ScanConfigUserAgent right) => left.Equals(right);
        public static bool operator !=(ScanConfigUserAgent left, ScanConfigUserAgent right) => !left.Equals(right);

        public static explicit operator string(ScanConfigUserAgent value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScanConfigUserAgent other && Equals(other);
        public bool Equals(ScanConfigUserAgent other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The execution state of the ScanRun.
    /// </summary>
    [EnumType]
    public readonly struct ScanRunExecutionState : IEquatable<ScanRunExecutionState>
    {
        private readonly string _value;

        private ScanRunExecutionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Represents an invalid state caused by internal server error. This value should never be returned.
        /// </summary>
        public static ScanRunExecutionState ExecutionStateUnspecified { get; } = new ScanRunExecutionState("EXECUTION_STATE_UNSPECIFIED");
        /// <summary>
        /// The scan is waiting in the queue.
        /// </summary>
        public static ScanRunExecutionState Queued { get; } = new ScanRunExecutionState("QUEUED");
        /// <summary>
        /// The scan is in progress.
        /// </summary>
        public static ScanRunExecutionState Scanning { get; } = new ScanRunExecutionState("SCANNING");
        /// <summary>
        /// The scan is either finished or stopped by user.
        /// </summary>
        public static ScanRunExecutionState Finished { get; } = new ScanRunExecutionState("FINISHED");

        public static bool operator ==(ScanRunExecutionState left, ScanRunExecutionState right) => left.Equals(right);
        public static bool operator !=(ScanRunExecutionState left, ScanRunExecutionState right) => !left.Equals(right);

        public static explicit operator string(ScanRunExecutionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScanRunExecutionState other && Equals(other);
        public bool Equals(ScanRunExecutionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The result state of the ScanRun. This field is only available after the execution state reaches "FINISHED".
    /// </summary>
    [EnumType]
    public readonly struct ScanRunResultState : IEquatable<ScanRunResultState>
    {
        private readonly string _value;

        private ScanRunResultState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value. This value is returned when the ScanRun is not yet finished.
        /// </summary>
        public static ScanRunResultState ResultStateUnspecified { get; } = new ScanRunResultState("RESULT_STATE_UNSPECIFIED");
        /// <summary>
        /// The scan finished without errors.
        /// </summary>
        public static ScanRunResultState Success { get; } = new ScanRunResultState("SUCCESS");
        /// <summary>
        /// The scan finished with errors.
        /// </summary>
        public static ScanRunResultState Error { get; } = new ScanRunResultState("ERROR");
        /// <summary>
        /// The scan was terminated by user.
        /// </summary>
        public static ScanRunResultState Killed { get; } = new ScanRunResultState("KILLED");

        public static bool operator ==(ScanRunResultState left, ScanRunResultState right) => left.Equals(right);
        public static bool operator !=(ScanRunResultState left, ScanRunResultState right) => !left.Equals(right);

        public static explicit operator string(ScanRunResultState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScanRunResultState other && Equals(other);
        public bool Equals(ScanRunResultState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
