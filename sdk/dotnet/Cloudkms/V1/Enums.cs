// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Cloudkms.V1
{
    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct AuditLogConfigLogType : IEquatable<AuditLogConfigLogType>
    {
        private readonly string _value;

        private AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static AuditLogConfigLogType LogTypeUnspecified { get; } = new AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static AuditLogConfigLogType AdminRead { get; } = new AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static AuditLogConfigLogType DataWrite { get; } = new AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static AuditLogConfigLogType DataRead { get; } = new AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(AuditLogConfigLogType left, AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(AuditLogConfigLogType left, AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuditLogConfigLogType other && Equals(other);
        public bool Equals(AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Immutable. The immutable purpose of this CryptoKey.
    /// </summary>
    [EnumType]
    public readonly struct CryptoKeyPurpose : IEquatable<CryptoKeyPurpose>
    {
        private readonly string _value;

        private CryptoKeyPurpose(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CryptoKeyPurpose CryptoKeyPurposeUnspecified { get; } = new CryptoKeyPurpose("CRYPTO_KEY_PURPOSE_UNSPECIFIED");
        /// <summary>
        /// CryptoKeys with this purpose may be used with Encrypt and Decrypt.
        /// </summary>
        public static CryptoKeyPurpose EncryptDecrypt { get; } = new CryptoKeyPurpose("ENCRYPT_DECRYPT");
        /// <summary>
        /// CryptoKeys with this purpose may be used with AsymmetricSign and GetPublicKey.
        /// </summary>
        public static CryptoKeyPurpose AsymmetricSign { get; } = new CryptoKeyPurpose("ASYMMETRIC_SIGN");
        /// <summary>
        /// CryptoKeys with this purpose may be used with AsymmetricDecrypt and GetPublicKey.
        /// </summary>
        public static CryptoKeyPurpose AsymmetricDecrypt { get; } = new CryptoKeyPurpose("ASYMMETRIC_DECRYPT");
        /// <summary>
        /// CryptoKeys with this purpose may be used with MacSign.
        /// </summary>
        public static CryptoKeyPurpose Mac { get; } = new CryptoKeyPurpose("MAC");

        public static bool operator ==(CryptoKeyPurpose left, CryptoKeyPurpose right) => left.Equals(right);
        public static bool operator !=(CryptoKeyPurpose left, CryptoKeyPurpose right) => !left.Equals(right);

        public static explicit operator string(CryptoKeyPurpose value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CryptoKeyPurpose other && Equals(other);
        public bool Equals(CryptoKeyPurpose other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the CryptoKeyVersion.
    /// </summary>
    [EnumType]
    public readonly struct CryptoKeyVersionState : IEquatable<CryptoKeyVersionState>
    {
        private readonly string _value;

        private CryptoKeyVersionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CryptoKeyVersionState CryptoKeyVersionStateUnspecified { get; } = new CryptoKeyVersionState("CRYPTO_KEY_VERSION_STATE_UNSPECIFIED");
        /// <summary>
        /// This version is still being generated. It may not be used, enabled, disabled, or destroyed yet. Cloud KMS will automatically mark this version ENABLED as soon as the version is ready.
        /// </summary>
        public static CryptoKeyVersionState PendingGeneration { get; } = new CryptoKeyVersionState("PENDING_GENERATION");
        /// <summary>
        /// This version may be used for cryptographic operations.
        /// </summary>
        public static CryptoKeyVersionState Enabled { get; } = new CryptoKeyVersionState("ENABLED");
        /// <summary>
        /// This version may not be used, but the key material is still available, and the version can be placed back into the ENABLED state.
        /// </summary>
        public static CryptoKeyVersionState Disabled { get; } = new CryptoKeyVersionState("DISABLED");
        /// <summary>
        /// This version is destroyed, and the key material is no longer stored. This version may only become ENABLED again if this version is reimport_eligible and the original key material is reimported with a call to KeyManagementService.ImportCryptoKeyVersion.
        /// </summary>
        public static CryptoKeyVersionState Destroyed { get; } = new CryptoKeyVersionState("DESTROYED");
        /// <summary>
        /// This version is scheduled for destruction, and will be destroyed soon. Call RestoreCryptoKeyVersion to put it back into the DISABLED state.
        /// </summary>
        public static CryptoKeyVersionState DestroyScheduled { get; } = new CryptoKeyVersionState("DESTROY_SCHEDULED");
        /// <summary>
        /// This version is still being imported. It may not be used, enabled, disabled, or destroyed yet. Cloud KMS will automatically mark this version ENABLED as soon as the version is ready.
        /// </summary>
        public static CryptoKeyVersionState PendingImport { get; } = new CryptoKeyVersionState("PENDING_IMPORT");
        /// <summary>
        /// This version was not imported successfully. It may not be used, enabled, disabled, or destroyed. The submitted key material has been discarded. Additional details can be found in CryptoKeyVersion.import_failure_reason.
        /// </summary>
        public static CryptoKeyVersionState ImportFailed { get; } = new CryptoKeyVersionState("IMPORT_FAILED");
        /// <summary>
        /// This version was not generated successfully. It may not be used, enabled, disabled, or destroyed. Additional details can be found in CryptoKeyVersion.generation_failure_reason.
        /// </summary>
        public static CryptoKeyVersionState GenerationFailed { get; } = new CryptoKeyVersionState("GENERATION_FAILED");
        /// <summary>
        /// This version was destroyed, and it may not be used or enabled again. Cloud KMS is waiting for the corresponding key material residing in an external key manager to be destroyed.
        /// </summary>
        public static CryptoKeyVersionState PendingExternalDestruction { get; } = new CryptoKeyVersionState("PENDING_EXTERNAL_DESTRUCTION");
        /// <summary>
        /// This version was destroyed, and it may not be used or enabled again. However, Cloud KMS could not confirm that the corresponding key material residing in an external key manager was destroyed. Additional details can be found in CryptoKeyVersion.external_destruction_failure_reason.
        /// </summary>
        public static CryptoKeyVersionState ExternalDestructionFailed { get; } = new CryptoKeyVersionState("EXTERNAL_DESTRUCTION_FAILED");

        public static bool operator ==(CryptoKeyVersionState left, CryptoKeyVersionState right) => left.Equals(right);
        public static bool operator !=(CryptoKeyVersionState left, CryptoKeyVersionState right) => !left.Equals(right);

        public static explicit operator string(CryptoKeyVersionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CryptoKeyVersionState other && Equals(other);
        public bool Equals(CryptoKeyVersionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Algorithm to use when creating a CryptoKeyVersion based on this template. For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKey.purpose is ENCRYPT_DECRYPT.
    /// </summary>
    [EnumType]
    public readonly struct CryptoKeyVersionTemplateAlgorithm : IEquatable<CryptoKeyVersionTemplateAlgorithm>
    {
        private readonly string _value;

        private CryptoKeyVersionTemplateAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm CryptoKeyVersionAlgorithmUnspecified { get; } = new CryptoKeyVersionTemplateAlgorithm("CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED");
        /// <summary>
        /// Creates symmetric encryption keys.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm GoogleSymmetricEncryption { get; } = new CryptoKeyVersionTemplateAlgorithm("GOOGLE_SYMMETRIC_ENCRYPTION");
        /// <summary>
        /// RSASSA-PSS 2048 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPss2048Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PSS_2048_SHA256");
        /// <summary>
        /// RSASSA-PSS 3072 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPss3072Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PSS_3072_SHA256");
        /// <summary>
        /// RSASSA-PSS 4096 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPss4096Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PSS_4096_SHA256");
        /// <summary>
        /// RSASSA-PSS 4096 bit key with a SHA512 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPss4096Sha512 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PSS_4096_SHA512");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 with a 2048 bit key and a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPkcs12048Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PKCS1_2048_SHA256");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 with a 3072 bit key and a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPkcs13072Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PKCS1_3072_SHA256");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPkcs14096Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PKCS1_4096_SHA256");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA512 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignPkcs14096Sha512 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_PKCS1_4096_SHA512");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 signing without encoding, with a 2048 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignRawPkcs12048 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_RAW_PKCS1_2048");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 signing without encoding, with a 3072 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignRawPkcs13072 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_RAW_PKCS1_3072");
        /// <summary>
        /// RSASSA-PKCS1-v1_5 signing without encoding, with a 4096 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaSignRawPkcs14096 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_SIGN_RAW_PKCS1_4096");
        /// <summary>
        /// RSAES-OAEP 2048 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep2048Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_2048_SHA256");
        /// <summary>
        /// RSAES-OAEP 3072 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep3072Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_3072_SHA256");
        /// <summary>
        /// RSAES-OAEP 4096 bit key with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep4096Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_4096_SHA256");
        /// <summary>
        /// RSAES-OAEP 4096 bit key with a SHA512 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep4096Sha512 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_4096_SHA512");
        /// <summary>
        /// RSAES-OAEP 2048 bit key with a SHA1 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep2048Sha1 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_2048_SHA1");
        /// <summary>
        /// RSAES-OAEP 3072 bit key with a SHA1 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep3072Sha1 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_3072_SHA1");
        /// <summary>
        /// RSAES-OAEP 4096 bit key with a SHA1 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm RsaDecryptOaep4096Sha1 { get; } = new CryptoKeyVersionTemplateAlgorithm("RSA_DECRYPT_OAEP_4096_SHA1");
        /// <summary>
        /// ECDSA on the NIST P-256 curve with a SHA256 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm EcSignP256Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("EC_SIGN_P256_SHA256");
        /// <summary>
        /// ECDSA on the NIST P-384 curve with a SHA384 digest.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm EcSignP384Sha384 { get; } = new CryptoKeyVersionTemplateAlgorithm("EC_SIGN_P384_SHA384");
        /// <summary>
        /// ECDSA on the non-NIST secp256k1 curve. This curve is only supported for HSM protection level.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm EcSignSecp256k1Sha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("EC_SIGN_SECP256K1_SHA256");
        /// <summary>
        /// HMAC-SHA256 signing with a 256 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm HmacSha256 { get; } = new CryptoKeyVersionTemplateAlgorithm("HMAC_SHA256");
        /// <summary>
        /// HMAC-SHA1 signing with a 160 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm HmacSha1 { get; } = new CryptoKeyVersionTemplateAlgorithm("HMAC_SHA1");
        /// <summary>
        /// HMAC-SHA384 signing with a 384 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm HmacSha384 { get; } = new CryptoKeyVersionTemplateAlgorithm("HMAC_SHA384");
        /// <summary>
        /// HMAC-SHA512 signing with a 512 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm HmacSha512 { get; } = new CryptoKeyVersionTemplateAlgorithm("HMAC_SHA512");
        /// <summary>
        /// HMAC-SHA224 signing with a 224 bit key.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm HmacSha224 { get; } = new CryptoKeyVersionTemplateAlgorithm("HMAC_SHA224");
        /// <summary>
        /// Algorithm representing symmetric encryption by an external key manager.
        /// </summary>
        public static CryptoKeyVersionTemplateAlgorithm ExternalSymmetricEncryption { get; } = new CryptoKeyVersionTemplateAlgorithm("EXTERNAL_SYMMETRIC_ENCRYPTION");

        public static bool operator ==(CryptoKeyVersionTemplateAlgorithm left, CryptoKeyVersionTemplateAlgorithm right) => left.Equals(right);
        public static bool operator !=(CryptoKeyVersionTemplateAlgorithm left, CryptoKeyVersionTemplateAlgorithm right) => !left.Equals(right);

        public static explicit operator string(CryptoKeyVersionTemplateAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CryptoKeyVersionTemplateAlgorithm other && Equals(other);
        public bool Equals(CryptoKeyVersionTemplateAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// ProtectionLevel to use when creating a CryptoKeyVersion based on this template. Immutable. Defaults to SOFTWARE.
    /// </summary>
    [EnumType]
    public readonly struct CryptoKeyVersionTemplateProtectionLevel : IEquatable<CryptoKeyVersionTemplateProtectionLevel>
    {
        private readonly string _value;

        private CryptoKeyVersionTemplateProtectionLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static CryptoKeyVersionTemplateProtectionLevel ProtectionLevelUnspecified { get; } = new CryptoKeyVersionTemplateProtectionLevel("PROTECTION_LEVEL_UNSPECIFIED");
        /// <summary>
        /// Crypto operations are performed in software.
        /// </summary>
        public static CryptoKeyVersionTemplateProtectionLevel Software { get; } = new CryptoKeyVersionTemplateProtectionLevel("SOFTWARE");
        /// <summary>
        /// Crypto operations are performed in a Hardware Security Module.
        /// </summary>
        public static CryptoKeyVersionTemplateProtectionLevel Hsm { get; } = new CryptoKeyVersionTemplateProtectionLevel("HSM");
        /// <summary>
        /// Crypto operations are performed by an external key manager.
        /// </summary>
        public static CryptoKeyVersionTemplateProtectionLevel External { get; } = new CryptoKeyVersionTemplateProtectionLevel("EXTERNAL");
        /// <summary>
        /// Crypto operations are performed in an EKM-over-VPC backend.
        /// </summary>
        public static CryptoKeyVersionTemplateProtectionLevel ExternalVpc { get; } = new CryptoKeyVersionTemplateProtectionLevel("EXTERNAL_VPC");

        public static bool operator ==(CryptoKeyVersionTemplateProtectionLevel left, CryptoKeyVersionTemplateProtectionLevel right) => left.Equals(right);
        public static bool operator !=(CryptoKeyVersionTemplateProtectionLevel left, CryptoKeyVersionTemplateProtectionLevel right) => !left.Equals(right);

        public static explicit operator string(CryptoKeyVersionTemplateProtectionLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CryptoKeyVersionTemplateProtectionLevel other && Equals(other);
        public bool Equals(CryptoKeyVersionTemplateProtectionLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL.
    /// </summary>
    [EnumType]
    public readonly struct EkmConnectionKeyManagementMode : IEquatable<EkmConnectionKeyManagementMode>
    {
        private readonly string _value;

        private EkmConnectionKeyManagementMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static EkmConnectionKeyManagementMode KeyManagementModeUnspecified { get; } = new EkmConnectionKeyManagementMode("KEY_MANAGEMENT_MODE_UNSPECIFIED");
        /// <summary>
        /// EKM-side key management operations on CryptoKeys created with this EkmConnection must be initiated from the EKM directly and cannot be performed from Cloud KMS. This means that: * When creating a CryptoKeyVersion associated with this EkmConnection, the caller must supply the key path of pre-existing external key material that will be linked to the CryptoKeyVersion. * Destruction of external key material cannot be requested via the Cloud KMS API and must be performed directly in the EKM. * Automatic rotation of key material is not supported.
        /// </summary>
        public static EkmConnectionKeyManagementMode Manual { get; } = new EkmConnectionKeyManagementMode("MANUAL");
        /// <summary>
        /// All CryptoKeys created with this EkmConnection use EKM-side key management operations initiated from Cloud KMS. This means that: * When a CryptoKeyVersion associated with this EkmConnection is created, the EKM automatically generates new key material and a new key path. The caller cannot supply the key path of pre-existing external key material. * Destruction of external key material associated with this EkmConnection can be requested by calling DestroyCryptoKeyVersion. * Automatic rotation of key material is supported.
        /// </summary>
        public static EkmConnectionKeyManagementMode CloudKms { get; } = new EkmConnectionKeyManagementMode("CLOUD_KMS");

        public static bool operator ==(EkmConnectionKeyManagementMode left, EkmConnectionKeyManagementMode right) => left.Equals(right);
        public static bool operator !=(EkmConnectionKeyManagementMode left, EkmConnectionKeyManagementMode right) => !left.Equals(right);

        public static explicit operator string(EkmConnectionKeyManagementMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EkmConnectionKeyManagementMode other && Equals(other);
        public bool Equals(EkmConnectionKeyManagementMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Immutable. The wrapping method to be used for incoming key material.
    /// </summary>
    [EnumType]
    public readonly struct ImportJobImportMethod : IEquatable<ImportJobImportMethod>
    {
        private readonly string _value;

        private ImportJobImportMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static ImportJobImportMethod ImportMethodUnspecified { get; } = new ImportJobImportMethod("IMPORT_METHOD_UNSPECIFIED");
        /// <summary>
        /// This ImportMethod represents the CKM_RSA_AES_KEY_WRAP key wrapping scheme defined in the PKCS #11 standard. In summary, this involves wrapping the raw key with an ephemeral AES key, and wrapping the ephemeral AES key with a 3072 bit RSA key. For more details, see [RSA AES key wrap mechanism](http://docs.oasis-open.org/pkcs11/pkcs11-curr/v2.40/cos01/pkcs11-curr-v2.40-cos01.html#_Toc408226908).
        /// </summary>
        public static ImportJobImportMethod RsaOaep3072Sha1Aes256 { get; } = new ImportJobImportMethod("RSA_OAEP_3072_SHA1_AES_256");
        /// <summary>
        /// This ImportMethod represents the CKM_RSA_AES_KEY_WRAP key wrapping scheme defined in the PKCS #11 standard. In summary, this involves wrapping the raw key with an ephemeral AES key, and wrapping the ephemeral AES key with a 4096 bit RSA key. For more details, see [RSA AES key wrap mechanism](http://docs.oasis-open.org/pkcs11/pkcs11-curr/v2.40/cos01/pkcs11-curr-v2.40-cos01.html#_Toc408226908).
        /// </summary>
        public static ImportJobImportMethod RsaOaep4096Sha1Aes256 { get; } = new ImportJobImportMethod("RSA_OAEP_4096_SHA1_AES_256");
        /// <summary>
        /// This ImportMethod represents the CKM_RSA_AES_KEY_WRAP key wrapping scheme defined in the PKCS #11 standard. In summary, this involves wrapping the raw key with an ephemeral AES key, and wrapping the ephemeral AES key with a 3072 bit RSA key. For more details, see [RSA AES key wrap mechanism](http://docs.oasis-open.org/pkcs11/pkcs11-curr/v2.40/cos01/pkcs11-curr-v2.40-cos01.html#_Toc408226908).
        /// </summary>
        public static ImportJobImportMethod RsaOaep3072Sha256Aes256 { get; } = new ImportJobImportMethod("RSA_OAEP_3072_SHA256_AES_256");
        /// <summary>
        /// This ImportMethod represents the CKM_RSA_AES_KEY_WRAP key wrapping scheme defined in the PKCS #11 standard. In summary, this involves wrapping the raw key with an ephemeral AES key, and wrapping the ephemeral AES key with a 4096 bit RSA key. For more details, see [RSA AES key wrap mechanism](http://docs.oasis-open.org/pkcs11/pkcs11-curr/v2.40/cos01/pkcs11-curr-v2.40-cos01.html#_Toc408226908).
        /// </summary>
        public static ImportJobImportMethod RsaOaep4096Sha256Aes256 { get; } = new ImportJobImportMethod("RSA_OAEP_4096_SHA256_AES_256");
        /// <summary>
        /// This ImportMethod represents RSAES-OAEP with a 3072 bit RSA key. The key material to be imported is wrapped directly with the RSA key. Due to technical limitations of RSA wrapping, this method cannot be used to wrap RSA keys for import.
        /// </summary>
        public static ImportJobImportMethod RsaOaep3072Sha256 { get; } = new ImportJobImportMethod("RSA_OAEP_3072_SHA256");
        /// <summary>
        /// This ImportMethod represents RSAES-OAEP with a 4096 bit RSA key. The key material to be imported is wrapped directly with the RSA key. Due to technical limitations of RSA wrapping, this method cannot be used to wrap RSA keys for import.
        /// </summary>
        public static ImportJobImportMethod RsaOaep4096Sha256 { get; } = new ImportJobImportMethod("RSA_OAEP_4096_SHA256");

        public static bool operator ==(ImportJobImportMethod left, ImportJobImportMethod right) => left.Equals(right);
        public static bool operator !=(ImportJobImportMethod left, ImportJobImportMethod right) => !left.Equals(right);

        public static explicit operator string(ImportJobImportMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ImportJobImportMethod other && Equals(other);
        public bool Equals(ImportJobImportMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
    /// </summary>
    [EnumType]
    public readonly struct ImportJobProtectionLevel : IEquatable<ImportJobProtectionLevel>
    {
        private readonly string _value;

        private ImportJobProtectionLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not specified.
        /// </summary>
        public static ImportJobProtectionLevel ProtectionLevelUnspecified { get; } = new ImportJobProtectionLevel("PROTECTION_LEVEL_UNSPECIFIED");
        /// <summary>
        /// Crypto operations are performed in software.
        /// </summary>
        public static ImportJobProtectionLevel Software { get; } = new ImportJobProtectionLevel("SOFTWARE");
        /// <summary>
        /// Crypto operations are performed in a Hardware Security Module.
        /// </summary>
        public static ImportJobProtectionLevel Hsm { get; } = new ImportJobProtectionLevel("HSM");
        /// <summary>
        /// Crypto operations are performed by an external key manager.
        /// </summary>
        public static ImportJobProtectionLevel External { get; } = new ImportJobProtectionLevel("EXTERNAL");
        /// <summary>
        /// Crypto operations are performed in an EKM-over-VPC backend.
        /// </summary>
        public static ImportJobProtectionLevel ExternalVpc { get; } = new ImportJobProtectionLevel("EXTERNAL_VPC");

        public static bool operator ==(ImportJobProtectionLevel left, ImportJobProtectionLevel right) => left.Equals(right);
        public static bool operator !=(ImportJobProtectionLevel left, ImportJobProtectionLevel right) => !left.Equals(right);

        public static explicit operator string(ImportJobProtectionLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ImportJobProtectionLevel other && Equals(other);
        public bool Equals(ImportJobProtectionLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
