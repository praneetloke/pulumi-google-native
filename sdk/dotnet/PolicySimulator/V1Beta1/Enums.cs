// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.PolicySimulator.V1Beta1
{
    /// <summary>
    /// The logs to use as input for the Replay.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource : IEquatable<GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource>
    {
        private readonly string _value;

        private GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An unspecified log source. If the log source is unspecified, the Replay defaults to using `RECENT_ACCESSES`.
        /// </summary>
        public static GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource LogSourceUnspecified { get; } = new GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource("LOG_SOURCE_UNSPECIFIED");
        /// <summary>
        /// All access logs from the last 90 days. These logs may not include logs from the most recent 7 days.
        /// </summary>
        public static GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource RecentAccesses { get; } = new GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource("RECENT_ACCESSES");

        public static bool operator ==(GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource left, GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource right) => left.Equals(right);
        public static bool operator !=(GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource left, GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource other && Equals(other);
        public bool Equals(GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
