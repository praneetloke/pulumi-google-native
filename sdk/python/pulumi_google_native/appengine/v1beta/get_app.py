# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetAppResult',
    'AwaitableGetAppResult',
    'get_app',
    'get_app_output',
]

@pulumi.output_type
class GetAppResult:
    def __init__(__self__, auth_domain=None, code_bucket=None, database_type=None, default_bucket=None, default_cookie_expiration=None, default_hostname=None, dispatch_rules=None, feature_settings=None, gcr_domain=None, iap=None, location=None, name=None, service_account=None, serving_status=None):
        if auth_domain and not isinstance(auth_domain, str):
            raise TypeError("Expected argument 'auth_domain' to be a str")
        pulumi.set(__self__, "auth_domain", auth_domain)
        if code_bucket and not isinstance(code_bucket, str):
            raise TypeError("Expected argument 'code_bucket' to be a str")
        pulumi.set(__self__, "code_bucket", code_bucket)
        if database_type and not isinstance(database_type, str):
            raise TypeError("Expected argument 'database_type' to be a str")
        pulumi.set(__self__, "database_type", database_type)
        if default_bucket and not isinstance(default_bucket, str):
            raise TypeError("Expected argument 'default_bucket' to be a str")
        pulumi.set(__self__, "default_bucket", default_bucket)
        if default_cookie_expiration and not isinstance(default_cookie_expiration, str):
            raise TypeError("Expected argument 'default_cookie_expiration' to be a str")
        pulumi.set(__self__, "default_cookie_expiration", default_cookie_expiration)
        if default_hostname and not isinstance(default_hostname, str):
            raise TypeError("Expected argument 'default_hostname' to be a str")
        pulumi.set(__self__, "default_hostname", default_hostname)
        if dispatch_rules and not isinstance(dispatch_rules, list):
            raise TypeError("Expected argument 'dispatch_rules' to be a list")
        pulumi.set(__self__, "dispatch_rules", dispatch_rules)
        if feature_settings and not isinstance(feature_settings, dict):
            raise TypeError("Expected argument 'feature_settings' to be a dict")
        pulumi.set(__self__, "feature_settings", feature_settings)
        if gcr_domain and not isinstance(gcr_domain, str):
            raise TypeError("Expected argument 'gcr_domain' to be a str")
        pulumi.set(__self__, "gcr_domain", gcr_domain)
        if iap and not isinstance(iap, dict):
            raise TypeError("Expected argument 'iap' to be a dict")
        pulumi.set(__self__, "iap", iap)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if serving_status and not isinstance(serving_status, str):
            raise TypeError("Expected argument 'serving_status' to be a str")
        pulumi.set(__self__, "serving_status", serving_status)

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> str:
        """
        Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        """
        return pulumi.get(self, "auth_domain")

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> str:
        """
        Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
        """
        return pulumi.get(self, "code_bucket")

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> str:
        """
        The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        """
        return pulumi.get(self, "database_type")

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> str:
        """
        Google Cloud Storage bucket that can be used by this application to store content.
        """
        return pulumi.get(self, "default_bucket")

    @property
    @pulumi.getter(name="defaultCookieExpiration")
    def default_cookie_expiration(self) -> str:
        """
        Cookie expiration policy for this application.
        """
        return pulumi.get(self, "default_cookie_expiration")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> str:
        """
        Hostname used to reach this application, as resolved by App Engine.
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter(name="dispatchRules")
    def dispatch_rules(self) -> Sequence['outputs.UrlDispatchRuleResponse']:
        """
        HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        """
        return pulumi.get(self, "dispatch_rules")

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> 'outputs.FeatureSettingsResponse':
        """
        The feature specific settings to be used in the application.
        """
        return pulumi.get(self, "feature_settings")

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> str:
        """
        The Google Container Registry domain used for storing managed build docker images for this application.
        """
        return pulumi.get(self, "gcr_domain")

    @property
    @pulumi.getter
    def iap(self) -> 'outputs.IdentityAwareProxyResponse':
        return pulumi.get(self, "iap")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Full path to the Application resource in the API. Example: apps/myapp.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> str:
        """
        Serving status of this application.
        """
        return pulumi.get(self, "serving_status")


class AwaitableGetAppResult(GetAppResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppResult(
            auth_domain=self.auth_domain,
            code_bucket=self.code_bucket,
            database_type=self.database_type,
            default_bucket=self.default_bucket,
            default_cookie_expiration=self.default_cookie_expiration,
            default_hostname=self.default_hostname,
            dispatch_rules=self.dispatch_rules,
            feature_settings=self.feature_settings,
            gcr_domain=self.gcr_domain,
            iap=self.iap,
            location=self.location,
            name=self.name,
            service_account=self.service_account,
            serving_status=self.serving_status)


def get_app(app_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppResult:
    """
    Gets information about an application.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:appengine/v1beta:getApp', __args__, opts=opts, typ=GetAppResult).value

    return AwaitableGetAppResult(
        auth_domain=pulumi.get(__ret__, 'auth_domain'),
        code_bucket=pulumi.get(__ret__, 'code_bucket'),
        database_type=pulumi.get(__ret__, 'database_type'),
        default_bucket=pulumi.get(__ret__, 'default_bucket'),
        default_cookie_expiration=pulumi.get(__ret__, 'default_cookie_expiration'),
        default_hostname=pulumi.get(__ret__, 'default_hostname'),
        dispatch_rules=pulumi.get(__ret__, 'dispatch_rules'),
        feature_settings=pulumi.get(__ret__, 'feature_settings'),
        gcr_domain=pulumi.get(__ret__, 'gcr_domain'),
        iap=pulumi.get(__ret__, 'iap'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        service_account=pulumi.get(__ret__, 'service_account'),
        serving_status=pulumi.get(__ret__, 'serving_status'))


@_utilities.lift_output_func(get_app)
def get_app_output(app_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppResult]:
    """
    Gets information about an application.
    """
    ...
