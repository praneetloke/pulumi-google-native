# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApiConfigHandlerAuthFailAction',
    'ApiConfigHandlerLogin',
    'ApiConfigHandlerSecurityLevel',
    'AppDatabaseType',
    'AppServingStatus',
    'EndpointsApiServiceRolloutStrategy',
    'ErrorHandlerErrorCode',
    'IngressRuleAction',
    'SslSettingsSslManagementType',
    'UrlMapAuthFailAction',
    'UrlMapLogin',
    'UrlMapRedirectHttpResponseCode',
    'UrlMapSecurityLevel',
    'VersionInboundServicesItem',
    'VersionServingStatus',
    'VpcAccessConnectorEgressSetting',
]


class ApiConfigHandlerAuthFailAction(str, Enum):
    """
    Action to take when users access resources that require authentication. Defaults to redirect.
    """
    AUTH_FAIL_ACTION_UNSPECIFIED = "AUTH_FAIL_ACTION_UNSPECIFIED"
    """Not specified. AUTH_FAIL_ACTION_REDIRECT is assumed."""
    AUTH_FAIL_ACTION_REDIRECT = "AUTH_FAIL_ACTION_REDIRECT"
    """Redirects user to "accounts.google.com". The user is redirected back to the application URL after signing in or creating an account."""
    AUTH_FAIL_ACTION_UNAUTHORIZED = "AUTH_FAIL_ACTION_UNAUTHORIZED"
    """Rejects request with a 401 HTTP status code and an error message."""


class ApiConfigHandlerLogin(str, Enum):
    """
    Level of login required to access this resource. Defaults to optional.
    """
    LOGIN_UNSPECIFIED = "LOGIN_UNSPECIFIED"
    """Not specified. LOGIN_OPTIONAL is assumed."""
    LOGIN_OPTIONAL = "LOGIN_OPTIONAL"
    """Does not require that the user is signed in."""
    LOGIN_ADMIN = "LOGIN_ADMIN"
    """If the user is not signed in, the auth_fail_action is taken. In addition, if the user is not an administrator for the application, they are given an error message regardless of auth_fail_action. If the user is an administrator, the handler proceeds."""
    LOGIN_REQUIRED = "LOGIN_REQUIRED"
    """If the user has signed in, the handler proceeds normally. Otherwise, the auth_fail_action is taken."""


class ApiConfigHandlerSecurityLevel(str, Enum):
    """
    Security (HTTPS) enforcement for this URL.
    """
    SECURE_UNSPECIFIED = "SECURE_UNSPECIFIED"
    """Not specified."""
    SECURE_DEFAULT = "SECURE_DEFAULT"
    """Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used, and respond accordingly."""
    SECURE_NEVER = "SECURE_NEVER"
    """Requests for a URL that match this handler that use HTTPS are automatically redirected to the HTTP equivalent URL."""
    SECURE_OPTIONAL = "SECURE_OPTIONAL"
    """Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly."""
    SECURE_ALWAYS = "SECURE_ALWAYS"
    """Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect."""


class AppDatabaseType(str, Enum):
    """
    The type of the Cloud Firestore or Cloud Datastore database associated with this application.
    """
    DATABASE_TYPE_UNSPECIFIED = "DATABASE_TYPE_UNSPECIFIED"
    """Database type is unspecified."""
    CLOUD_DATASTORE = "CLOUD_DATASTORE"
    """Cloud Datastore"""
    CLOUD_FIRESTORE = "CLOUD_FIRESTORE"
    """Cloud Firestore Native"""
    CLOUD_DATASTORE_COMPATIBILITY = "CLOUD_DATASTORE_COMPATIBILITY"
    """Cloud Firestore in Datastore Mode"""


class AppServingStatus(str, Enum):
    """
    Serving status of this application.
    """
    UNSPECIFIED = "UNSPECIFIED"
    """Serving status is unspecified."""
    SERVING = "SERVING"
    """Application is serving."""
    USER_DISABLED = "USER_DISABLED"
    """Application has been disabled by the user."""
    SYSTEM_DISABLED = "SYSTEM_DISABLED"
    """Application has been disabled by the system."""


class EndpointsApiServiceRolloutStrategy(str, Enum):
    """
    Endpoints rollout strategy. If FIXED, config_id must be specified. If MANAGED, config_id must be omitted.
    """
    UNSPECIFIED_ROLLOUT_STRATEGY = "UNSPECIFIED_ROLLOUT_STRATEGY"
    """Not specified. Defaults to FIXED."""
    FIXED = "FIXED"
    """Endpoints service configuration ID will be fixed to the configuration ID specified by config_id."""
    MANAGED = "MANAGED"
    """Endpoints service configuration ID will be updated with each rollout."""


class ErrorHandlerErrorCode(str, Enum):
    """
    Error condition this handler applies to.
    """
    ERROR_CODE_UNSPECIFIED = "ERROR_CODE_UNSPECIFIED"
    """Not specified. ERROR_CODE_DEFAULT is assumed."""
    ERROR_CODE_DEFAULT = "ERROR_CODE_DEFAULT"
    """All other error types."""
    ERROR_CODE_OVER_QUOTA = "ERROR_CODE_OVER_QUOTA"
    """Application has exceeded a resource quota."""
    ERROR_CODE_DOS_API_DENIAL = "ERROR_CODE_DOS_API_DENIAL"
    """Client blocked by the application's Denial of Service protection configuration."""
    ERROR_CODE_TIMEOUT = "ERROR_CODE_TIMEOUT"
    """Deadline reached before the application responds."""


class IngressRuleAction(str, Enum):
    """
    The action to take on matched requests.
    """
    UNSPECIFIED_ACTION = "UNSPECIFIED_ACTION"
    ALLOW = "ALLOW"
    """Matching requests are allowed."""
    DENY = "DENY"
    """Matching requests are denied."""


class SslSettingsSslManagementType(str, Enum):
    """
    SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned. If MANUAL, certificate_id must be manually specified in order to configure SSL for this domain.
    """
    SSL_MANAGEMENT_TYPE_UNSPECIFIED = "SSL_MANAGEMENT_TYPE_UNSPECIFIED"
    """Defaults to AUTOMATIC."""
    AUTOMATIC = "AUTOMATIC"
    """SSL support for this domain is configured automatically. The mapped SSL certificate will be automatically renewed."""
    MANUAL = "MANUAL"
    """SSL support for this domain is configured manually by the user. Either the domain has no SSL support or a user-obtained SSL certificate has been explictly mapped to this domain."""


class UrlMapAuthFailAction(str, Enum):
    """
    Action to take when users access resources that require authentication. Defaults to redirect.
    """
    AUTH_FAIL_ACTION_UNSPECIFIED = "AUTH_FAIL_ACTION_UNSPECIFIED"
    """Not specified. AUTH_FAIL_ACTION_REDIRECT is assumed."""
    AUTH_FAIL_ACTION_REDIRECT = "AUTH_FAIL_ACTION_REDIRECT"
    """Redirects user to "accounts.google.com". The user is redirected back to the application URL after signing in or creating an account."""
    AUTH_FAIL_ACTION_UNAUTHORIZED = "AUTH_FAIL_ACTION_UNAUTHORIZED"
    """Rejects request with a 401 HTTP status code and an error message."""


class UrlMapLogin(str, Enum):
    """
    Level of login required to access this resource. Not supported for Node.js in the App Engine standard environment.
    """
    LOGIN_UNSPECIFIED = "LOGIN_UNSPECIFIED"
    """Not specified. LOGIN_OPTIONAL is assumed."""
    LOGIN_OPTIONAL = "LOGIN_OPTIONAL"
    """Does not require that the user is signed in."""
    LOGIN_ADMIN = "LOGIN_ADMIN"
    """If the user is not signed in, the auth_fail_action is taken. In addition, if the user is not an administrator for the application, they are given an error message regardless of auth_fail_action. If the user is an administrator, the handler proceeds."""
    LOGIN_REQUIRED = "LOGIN_REQUIRED"
    """If the user has signed in, the handler proceeds normally. Otherwise, the auth_fail_action is taken."""


class UrlMapRedirectHttpResponseCode(str, Enum):
    """
    30x code to use when performing redirects for the secure field. Defaults to 302.
    """
    REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED = "REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED"
    """Not specified. 302 is assumed."""
    REDIRECT_HTTP_RESPONSE_CODE301 = "REDIRECT_HTTP_RESPONSE_CODE_301"
    """301 Moved Permanently code."""
    REDIRECT_HTTP_RESPONSE_CODE302 = "REDIRECT_HTTP_RESPONSE_CODE_302"
    """302 Moved Temporarily code."""
    REDIRECT_HTTP_RESPONSE_CODE303 = "REDIRECT_HTTP_RESPONSE_CODE_303"
    """303 See Other code."""
    REDIRECT_HTTP_RESPONSE_CODE307 = "REDIRECT_HTTP_RESPONSE_CODE_307"
    """307 Temporary Redirect code."""


class UrlMapSecurityLevel(str, Enum):
    """
    Security (HTTPS) enforcement for this URL.
    """
    SECURE_UNSPECIFIED = "SECURE_UNSPECIFIED"
    """Not specified."""
    SECURE_DEFAULT = "SECURE_DEFAULT"
    """Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used, and respond accordingly."""
    SECURE_NEVER = "SECURE_NEVER"
    """Requests for a URL that match this handler that use HTTPS are automatically redirected to the HTTP equivalent URL."""
    SECURE_OPTIONAL = "SECURE_OPTIONAL"
    """Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly."""
    SECURE_ALWAYS = "SECURE_ALWAYS"
    """Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect."""


class VersionInboundServicesItem(str, Enum):
    INBOUND_SERVICE_UNSPECIFIED = "INBOUND_SERVICE_UNSPECIFIED"
    """Not specified."""
    INBOUND_SERVICE_MAIL = "INBOUND_SERVICE_MAIL"
    """Allows an application to receive mail."""
    INBOUND_SERVICE_MAIL_BOUNCE = "INBOUND_SERVICE_MAIL_BOUNCE"
    """Allows an application to receive email-bound notifications."""
    INBOUND_SERVICE_XMPP_ERROR = "INBOUND_SERVICE_XMPP_ERROR"
    """Allows an application to receive error stanzas."""
    INBOUND_SERVICE_XMPP_MESSAGE = "INBOUND_SERVICE_XMPP_MESSAGE"
    """Allows an application to receive instant messages."""
    INBOUND_SERVICE_XMPP_SUBSCRIBE = "INBOUND_SERVICE_XMPP_SUBSCRIBE"
    """Allows an application to receive user subscription POSTs."""
    INBOUND_SERVICE_XMPP_PRESENCE = "INBOUND_SERVICE_XMPP_PRESENCE"
    """Allows an application to receive a user's chat presence."""
    INBOUND_SERVICE_CHANNEL_PRESENCE = "INBOUND_SERVICE_CHANNEL_PRESENCE"
    """Registers an application for notifications when a client connects or disconnects from a channel."""
    INBOUND_SERVICE_WARMUP = "INBOUND_SERVICE_WARMUP"
    """Enables warmup requests."""


class VersionServingStatus(str, Enum):
    """
    Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
    """
    SERVING_STATUS_UNSPECIFIED = "SERVING_STATUS_UNSPECIFIED"
    """Not specified."""
    SERVING = "SERVING"
    """Currently serving. Instances are created according to the scaling settings of the version."""
    STOPPED = "STOPPED"
    """Disabled. No instances will be created and the scaling settings are ignored until the state of the version changes to SERVING."""


class VpcAccessConnectorEgressSetting(str, Enum):
    """
    The egress setting for the connector, controlling what traffic is diverted through it.
    """
    EGRESS_SETTING_UNSPECIFIED = "EGRESS_SETTING_UNSPECIFIED"
    ALL_TRAFFIC = "ALL_TRAFFIC"
    """Force the use of VPC Access for all egress traffic from the function."""
    PRIVATE_IP_RANGES = "PRIVATE_IP_RANGES"
    """Use the VPC Access Connector for private IP space from RFC1918."""