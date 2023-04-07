# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'AuthConfigAuthType',
    'SslConfigClientCertType',
    'SslConfigServerCertType',
    'SslConfigTrustModel',
    'SslConfigType',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class AuthConfigAuthType(str, Enum):
    """
    The type of authentication configured.
    """
    AUTH_TYPE_UNSPECIFIED = "AUTH_TYPE_UNSPECIFIED"
    """
    Authentication type not specified.
    """
    USER_PASSWORD = "USER_PASSWORD"
    """
    Username and Password Authentication.
    """
    OAUTH2_JWT_BEARER = "OAUTH2_JWT_BEARER"
    """
    JSON Web Token (JWT) Profile for Oauth 2.0 Authorization Grant based authentication
    """
    OAUTH2_CLIENT_CREDENTIALS = "OAUTH2_CLIENT_CREDENTIALS"
    """
    Oauth 2.0 Client Credentials Grant Authentication
    """
    SSH_PUBLIC_KEY = "SSH_PUBLIC_KEY"
    """
    SSH Public Key Authentication
    """
    OAUTH2_AUTH_CODE_FLOW = "OAUTH2_AUTH_CODE_FLOW"
    """
    Oauth 2.0 Authorization Code Flow
    """


class SslConfigClientCertType(str, Enum):
    """
    Type of Client Cert (PEM/JKS/.. etc.)
    """
    CERT_TYPE_UNSPECIFIED = "CERT_TYPE_UNSPECIFIED"
    """
    Cert type unspecified.
    """
    PEM = "PEM"
    """
    Privacy Enhanced Mail (PEM) Type
    """


class SslConfigServerCertType(str, Enum):
    """
    Type of Server Cert (PEM/JKS/.. etc.)
    """
    CERT_TYPE_UNSPECIFIED = "CERT_TYPE_UNSPECIFIED"
    """
    Cert type unspecified.
    """
    PEM = "PEM"
    """
    Privacy Enhanced Mail (PEM) Type
    """


class SslConfigTrustModel(str, Enum):
    """
    Trust Model of the SSL connection
    """
    PUBLIC = "PUBLIC"
    """
    Public Trust Model. Takes the Default Java trust store.
    """
    PRIVATE = "PRIVATE"
    """
    Private Trust Model. Takes custom/private trust store.
    """
    INSECURE = "INSECURE"
    """
    Insecure Trust Model. Accept all certificates.
    """


class SslConfigType(str, Enum):
    """
    Controls the ssl type for the given connector version.
    """
    SSL_TYPE_UNSPECIFIED = "SSL_TYPE_UNSPECIFIED"
    """
    No SSL configuration required.
    """
    TLS = "TLS"
    """
    TLS Handshake
    """
    MTLS = "MTLS"
    """
    mutual TLS (MTLS) Handshake
    """
