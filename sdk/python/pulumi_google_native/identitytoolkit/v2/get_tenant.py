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
    'GetTenantResult',
    'AwaitableGetTenantResult',
    'get_tenant',
    'get_tenant_output',
]

@pulumi.output_type
class GetTenantResult:
    def __init__(__self__, allow_password_signup=None, autodelete_anonymous_users=None, client=None, disable_auth=None, display_name=None, email_privacy_config=None, enable_anonymous_user=None, enable_email_link_signin=None, hash_config=None, inheritance=None, mfa_config=None, monitoring=None, name=None, password_policy_config=None, recaptcha_config=None, sms_region_config=None, test_phone_numbers=None):
        if allow_password_signup and not isinstance(allow_password_signup, bool):
            raise TypeError("Expected argument 'allow_password_signup' to be a bool")
        pulumi.set(__self__, "allow_password_signup", allow_password_signup)
        if autodelete_anonymous_users and not isinstance(autodelete_anonymous_users, bool):
            raise TypeError("Expected argument 'autodelete_anonymous_users' to be a bool")
        pulumi.set(__self__, "autodelete_anonymous_users", autodelete_anonymous_users)
        if client and not isinstance(client, dict):
            raise TypeError("Expected argument 'client' to be a dict")
        pulumi.set(__self__, "client", client)
        if disable_auth and not isinstance(disable_auth, bool):
            raise TypeError("Expected argument 'disable_auth' to be a bool")
        pulumi.set(__self__, "disable_auth", disable_auth)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if email_privacy_config and not isinstance(email_privacy_config, dict):
            raise TypeError("Expected argument 'email_privacy_config' to be a dict")
        pulumi.set(__self__, "email_privacy_config", email_privacy_config)
        if enable_anonymous_user and not isinstance(enable_anonymous_user, bool):
            raise TypeError("Expected argument 'enable_anonymous_user' to be a bool")
        pulumi.set(__self__, "enable_anonymous_user", enable_anonymous_user)
        if enable_email_link_signin and not isinstance(enable_email_link_signin, bool):
            raise TypeError("Expected argument 'enable_email_link_signin' to be a bool")
        pulumi.set(__self__, "enable_email_link_signin", enable_email_link_signin)
        if hash_config and not isinstance(hash_config, dict):
            raise TypeError("Expected argument 'hash_config' to be a dict")
        pulumi.set(__self__, "hash_config", hash_config)
        if inheritance and not isinstance(inheritance, dict):
            raise TypeError("Expected argument 'inheritance' to be a dict")
        pulumi.set(__self__, "inheritance", inheritance)
        if mfa_config and not isinstance(mfa_config, dict):
            raise TypeError("Expected argument 'mfa_config' to be a dict")
        pulumi.set(__self__, "mfa_config", mfa_config)
        if monitoring and not isinstance(monitoring, dict):
            raise TypeError("Expected argument 'monitoring' to be a dict")
        pulumi.set(__self__, "monitoring", monitoring)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password_policy_config and not isinstance(password_policy_config, dict):
            raise TypeError("Expected argument 'password_policy_config' to be a dict")
        pulumi.set(__self__, "password_policy_config", password_policy_config)
        if recaptcha_config and not isinstance(recaptcha_config, dict):
            raise TypeError("Expected argument 'recaptcha_config' to be a dict")
        pulumi.set(__self__, "recaptcha_config", recaptcha_config)
        if sms_region_config and not isinstance(sms_region_config, dict):
            raise TypeError("Expected argument 'sms_region_config' to be a dict")
        pulumi.set(__self__, "sms_region_config", sms_region_config)
        if test_phone_numbers and not isinstance(test_phone_numbers, dict):
            raise TypeError("Expected argument 'test_phone_numbers' to be a dict")
        pulumi.set(__self__, "test_phone_numbers", test_phone_numbers)

    @property
    @pulumi.getter(name="allowPasswordSignup")
    def allow_password_signup(self) -> bool:
        """
        Whether to allow email/password user authentication.
        """
        return pulumi.get(self, "allow_password_signup")

    @property
    @pulumi.getter(name="autodeleteAnonymousUsers")
    def autodelete_anonymous_users(self) -> bool:
        """
        Whether anonymous users will be auto-deleted after a period of 30 days.
        """
        return pulumi.get(self, "autodelete_anonymous_users")

    @property
    @pulumi.getter
    def client(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponse':
        """
        Options related to how clients making requests on behalf of a project should be configured.
        """
        return pulumi.get(self, "client")

    @property
    @pulumi.getter(name="disableAuth")
    def disable_auth(self) -> bool:
        """
        Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
        """
        return pulumi.get(self, "disable_auth")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name of the tenant.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="emailPrivacyConfig")
    def email_privacy_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponse':
        """
        Configuration for settings related to email privacy and public visibility.
        """
        return pulumi.get(self, "email_privacy_config")

    @property
    @pulumi.getter(name="enableAnonymousUser")
    def enable_anonymous_user(self) -> bool:
        """
        Whether to enable anonymous user authentication.
        """
        return pulumi.get(self, "enable_anonymous_user")

    @property
    @pulumi.getter(name="enableEmailLinkSignin")
    def enable_email_link_signin(self) -> bool:
        """
        Whether to enable email link user authentication.
        """
        return pulumi.get(self, "enable_email_link_signin")

    @property
    @pulumi.getter(name="hashConfig")
    def hash_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2HashConfigResponse':
        """
        Hash config information of a tenant for display on Pantheon. This can only be displayed on Pantheon to avoid the sensitive information to get accidentally leaked. Only returned in GetTenant response to restrict reading of this information. Requires firebaseauth.configs.getHashConfig permission on the agent project for returning this field.
        """
        return pulumi.get(self, "hash_config")

    @property
    @pulumi.getter
    def inheritance(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2InheritanceResponse':
        """
        Specify the settings that the tenant could inherit.
        """
        return pulumi.get(self, "inheritance")

    @property
    @pulumi.getter(name="mfaConfig")
    def mfa_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponse':
        """
        The tenant-level configuration of MFA options.
        """
        return pulumi.get(self, "mfa_config")

    @property
    @pulumi.getter
    def monitoring(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponse':
        """
        Configuration related to monitoring project activity.
        """
        return pulumi.get(self, "monitoring")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of a tenant. For example: "projects/{project-id}/tenants/{tenant-id}"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passwordPolicyConfig")
    def password_policy_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponse':
        """
        The tenant-level password policy config
        """
        return pulumi.get(self, "password_policy_config")

    @property
    @pulumi.getter(name="recaptchaConfig")
    def recaptcha_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponse':
        """
        The tenant-level reCAPTCHA config.
        """
        return pulumi.get(self, "recaptcha_config")

    @property
    @pulumi.getter(name="smsRegionConfig")
    def sms_region_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponse':
        """
        Configures which regions are enabled for SMS verification code sending.
        """
        return pulumi.get(self, "sms_region_config")

    @property
    @pulumi.getter(name="testPhoneNumbers")
    def test_phone_numbers(self) -> Mapping[str, str]:
        """
        A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
        """
        return pulumi.get(self, "test_phone_numbers")


class AwaitableGetTenantResult(GetTenantResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTenantResult(
            allow_password_signup=self.allow_password_signup,
            autodelete_anonymous_users=self.autodelete_anonymous_users,
            client=self.client,
            disable_auth=self.disable_auth,
            display_name=self.display_name,
            email_privacy_config=self.email_privacy_config,
            enable_anonymous_user=self.enable_anonymous_user,
            enable_email_link_signin=self.enable_email_link_signin,
            hash_config=self.hash_config,
            inheritance=self.inheritance,
            mfa_config=self.mfa_config,
            monitoring=self.monitoring,
            name=self.name,
            password_policy_config=self.password_policy_config,
            recaptcha_config=self.recaptcha_config,
            sms_region_config=self.sms_region_config,
            test_phone_numbers=self.test_phone_numbers)


def get_tenant(project: Optional[str] = None,
               tenant_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTenantResult:
    """
    Get a tenant. Requires read permission on the Tenant resource.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:identitytoolkit/v2:getTenant', __args__, opts=opts, typ=GetTenantResult).value

    return AwaitableGetTenantResult(
        allow_password_signup=pulumi.get(__ret__, 'allow_password_signup'),
        autodelete_anonymous_users=pulumi.get(__ret__, 'autodelete_anonymous_users'),
        client=pulumi.get(__ret__, 'client'),
        disable_auth=pulumi.get(__ret__, 'disable_auth'),
        display_name=pulumi.get(__ret__, 'display_name'),
        email_privacy_config=pulumi.get(__ret__, 'email_privacy_config'),
        enable_anonymous_user=pulumi.get(__ret__, 'enable_anonymous_user'),
        enable_email_link_signin=pulumi.get(__ret__, 'enable_email_link_signin'),
        hash_config=pulumi.get(__ret__, 'hash_config'),
        inheritance=pulumi.get(__ret__, 'inheritance'),
        mfa_config=pulumi.get(__ret__, 'mfa_config'),
        monitoring=pulumi.get(__ret__, 'monitoring'),
        name=pulumi.get(__ret__, 'name'),
        password_policy_config=pulumi.get(__ret__, 'password_policy_config'),
        recaptcha_config=pulumi.get(__ret__, 'recaptcha_config'),
        sms_region_config=pulumi.get(__ret__, 'sms_region_config'),
        test_phone_numbers=pulumi.get(__ret__, 'test_phone_numbers'))


@_utilities.lift_output_func(get_tenant)
def get_tenant_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                      tenant_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTenantResult]:
    """
    Get a tenant. Requires read permission on the Tenant resource.
    """
    ...
