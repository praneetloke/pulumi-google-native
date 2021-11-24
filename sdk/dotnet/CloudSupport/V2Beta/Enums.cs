// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CloudSupport.V2Beta
{
    /// <summary>
    /// The severity of this case.
    /// </summary>
    [EnumType]
    public readonly struct CaseSeverity : IEquatable<CaseSeverity>
    {
        private readonly string _value;

        private CaseSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Severity is undefined or has not been set yet.
        /// </summary>
        public static CaseSeverity SeverityUnspecified { get; } = new CaseSeverity("SEVERITY_UNSPECIFIED");
        /// <summary>
        /// Extreme impact on a production service. Service is hard down.
        /// </summary>
        public static CaseSeverity S0 { get; } = new CaseSeverity("S0");
        /// <summary>
        /// Critical impact on a production service. Service is currently unusable.
        /// </summary>
        public static CaseSeverity S1 { get; } = new CaseSeverity("S1");
        /// <summary>
        /// Severe impact on a production service. Service is usable but greatly impaired.
        /// </summary>
        public static CaseSeverity S2 { get; } = new CaseSeverity("S2");
        /// <summary>
        /// Medium impact on a production service. Service is available, but moderately impaired.
        /// </summary>
        public static CaseSeverity S3 { get; } = new CaseSeverity("S3");
        /// <summary>
        /// General questions or minor issues. Production service is fully available.
        /// </summary>
        public static CaseSeverity S4 { get; } = new CaseSeverity("S4");

        public static bool operator ==(CaseSeverity left, CaseSeverity right) => left.Equals(right);
        public static bool operator !=(CaseSeverity left, CaseSeverity right) => !left.Equals(right);

        public static explicit operator string(CaseSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CaseSeverity other && Equals(other);
        public bool Equals(CaseSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}