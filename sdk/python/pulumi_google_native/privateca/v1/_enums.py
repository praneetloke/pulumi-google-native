# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'CaPoolTier',
    'CertificateAuthorityType',
    'CertificateExtensionConstraintsKnownExtensionsItem',
    'CertificateSubjectMode',
    'EcKeyTypeSignatureAlgorithm',
    'KeyVersionSpecAlgorithm',
    'PublicKeyFormat',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """Default case. Should never be this."""
    ADMIN_READ = "ADMIN_READ"
    """Admin reads. Example: CloudIAM getIamPolicy"""
    DATA_WRITE = "DATA_WRITE"
    """Data writes. Example: CloudSQL Users create"""
    DATA_READ = "DATA_READ"
    """Data reads. Example: CloudSQL Users list"""


class CaPoolTier(str, Enum):
    """
    Required. Immutable. The Tier of this CaPool.
    """
    TIER_UNSPECIFIED = "TIER_UNSPECIFIED"
    """Not specified."""
    ENTERPRISE = "ENTERPRISE"
    """Enterprise tier."""
    DEVOPS = "DEVOPS"
    """DevOps tier."""


class CertificateAuthorityType(str, Enum):
    """
    Required. Immutable. The Type of this CertificateAuthority.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """Not specified."""
    SELF_SIGNED = "SELF_SIGNED"
    """Self-signed CA."""
    SUBORDINATE = "SUBORDINATE"
    """Subordinate CA. Could be issued by a Private CA CertificateAuthority or an unmanaged CA."""


class CertificateExtensionConstraintsKnownExtensionsItem(str, Enum):
    KNOWN_CERTIFICATE_EXTENSION_UNSPECIFIED = "KNOWN_CERTIFICATE_EXTENSION_UNSPECIFIED"
    """Not specified."""
    BASE_KEY_USAGE = "BASE_KEY_USAGE"
    """Refers to a certificate's Key Usage extension, as described in [RFC 5280 section 4.2.1.3](https://tools.ietf.org/html/rfc5280#section-4.2.1.3). This corresponds to the KeyUsage.base_key_usage field."""
    EXTENDED_KEY_USAGE = "EXTENDED_KEY_USAGE"
    """Refers to a certificate's Extended Key Usage extension, as described in [RFC 5280 section 4.2.1.12](https://tools.ietf.org/html/rfc5280#section-4.2.1.12). This corresponds to the KeyUsage.extended_key_usage message."""
    CA_OPTIONS = "CA_OPTIONS"
    """Refers to a certificate's Basic Constraints extension, as described in [RFC 5280 section 4.2.1.9](https://tools.ietf.org/html/rfc5280#section-4.2.1.9). This corresponds to the X509Parameters.ca_options field."""
    POLICY_IDS = "POLICY_IDS"
    """Refers to a certificate's Policy object identifiers, as described in [RFC 5280 section 4.2.1.4](https://tools.ietf.org/html/rfc5280#section-4.2.1.4). This corresponds to the X509Parameters.policy_ids field."""
    AIA_OCSP_SERVERS = "AIA_OCSP_SERVERS"
    """Refers to OCSP servers in a certificate's Authority Information Access extension, as described in [RFC 5280 section 4.2.2.1](https://tools.ietf.org/html/rfc5280#section-4.2.2.1), This corresponds to the X509Parameters.aia_ocsp_servers field."""


class CertificateSubjectMode(str, Enum):
    """
    Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
    """
    SUBJECT_REQUEST_MODE_UNSPECIFIED = "SUBJECT_REQUEST_MODE_UNSPECIFIED"
    """Not specified."""
    DEFAULT = "DEFAULT"
    """The default mode used in most cases. Indicates that the certificate's Subject and/or SubjectAltNames are specified in the certificate request. This mode requires the caller to have the `privateca.certificates.create` permission."""
    REFLECTED_SPIFFE = "REFLECTED_SPIFFE"
    """A mode reserved for special cases. Indicates that the certificate should have one or more SPIFFE SubjectAltNames set by the service based on the caller's identity. This mode will ignore any explicitly specified Subject and/or SubjectAltNames in the certificate request. This mode requires the caller to have the `privateca.certificates.createForSelf` permission."""


class EcKeyTypeSignatureAlgorithm(str, Enum):
    """
    Optional. A signature algorithm that must be used. If this is omitted, any EC-based signature algorithm will be allowed.
    """
    EC_SIGNATURE_ALGORITHM_UNSPECIFIED = "EC_SIGNATURE_ALGORITHM_UNSPECIFIED"
    """Not specified. Signifies that any signature algorithm may be used."""
    ECDSA_P256 = "ECDSA_P256"
    """Refers to the Elliptic Curve Digital Signature Algorithm over the NIST P-256 curve."""
    ECDSA_P384 = "ECDSA_P384"
    """Refers to the Elliptic Curve Digital Signature Algorithm over the NIST P-384 curve."""
    EDDSA25519 = "EDDSA_25519"
    """Refers to the Edwards-curve Digital Signature Algorithm over curve 25519, as described in RFC 8410."""


class KeyVersionSpecAlgorithm(str, Enum):
    """
    The algorithm to use for creating a managed Cloud KMS key for a for a simplified experience. All managed keys will be have their ProtectionLevel as `HSM`.
    """
    SIGN_HASH_ALGORITHM_UNSPECIFIED = "SIGN_HASH_ALGORITHM_UNSPECIFIED"
    """Not specified."""
    RSA_PSS2048_SHA256 = "RSA_PSS_2048_SHA256"
    """maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PSS_2048_SHA256"""
    RSA_PSS3072_SHA256 = "RSA_PSS_3072_SHA256"
    """maps to CryptoKeyVersionAlgorithm. RSA_SIGN_PSS_3072_SHA256"""
    RSA_PSS4096_SHA256 = "RSA_PSS_4096_SHA256"
    """maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PSS_4096_SHA256"""
    RSA_PKCS12048_SHA256 = "RSA_PKCS1_2048_SHA256"
    """maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_2048_SHA256"""
    RSA_PKCS13072_SHA256 = "RSA_PKCS1_3072_SHA256"
    """maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_3072_SHA256"""
    RSA_PKCS14096_SHA256 = "RSA_PKCS1_4096_SHA256"
    """maps to CryptoKeyVersionAlgorithm.RSA_SIGN_PKCS1_4096_SHA256"""
    EC_P256_SHA256 = "EC_P256_SHA256"
    """maps to CryptoKeyVersionAlgorithm.EC_SIGN_P256_SHA256"""
    EC_P384_SHA384 = "EC_P384_SHA384"
    """maps to CryptoKeyVersionAlgorithm.EC_SIGN_P384_SHA384"""


class PublicKeyFormat(str, Enum):
    """
    Required. The format of the public key.
    """
    KEY_FORMAT_UNSPECIFIED = "KEY_FORMAT_UNSPECIFIED"
    """Default unspecified value."""
    PEM = "PEM"
    """The key is PEM-encoded as defined in [RFC 7468](https://tools.ietf.org/html/rfc7468). It can be any of the following: a PEM-encoded PKCS#1/RFC 3447 RSAPublicKey structure, an RFC 5280 [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) or a PEM-encoded X.509 certificate signing request (CSR). If a [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) is specified, it can contain a A PEM-encoded PKCS#1/RFC 3447 RSAPublicKey or a NIST P-256/secp256r1/prime256v1 or P-384 key. If a CSR is specified, it will used solely for the purpose of extracting the public key. When generated by the service, it will always be an RFC 5280 [SubjectPublicKeyInfo](https://tools.ietf.org/html/rfc5280#section-4.1) structure containing an algorithm identifier and a key."""