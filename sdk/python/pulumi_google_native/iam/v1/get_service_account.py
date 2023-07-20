# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetServiceAccountResult',
    'AwaitableGetServiceAccountResult',
    'get_service_account',
    'get_service_account_output',
]

@pulumi.output_type
class GetServiceAccountResult:
    def __init__(__self__, description=None, disabled=None, display_name=None, email=None, etag=None, name=None, oauth2_client_id=None, project=None, unique_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if oauth2_client_id and not isinstance(oauth2_client_id, str):
            raise TypeError("Expected argument 'oauth2_client_id' to be a str")
        pulumi.set(__self__, "oauth2_client_id", oauth2_client_id)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        pulumi.set(__self__, "unique_id", unique_id)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Whether the service account is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email address of the service account.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Deprecated. Do not use.
        """
        warnings.warn("""Deprecated. Do not use.""", DeprecationWarning)
        pulumi.log.warn("""etag is deprecated: Deprecated. Do not use.""")

        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to access the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oauth2ClientId")
    def oauth2_client_id(self) -> str:
        """
        The OAuth 2.0 client ID for the service account.
        """
        return pulumi.get(self, "oauth2_client_id")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project that owns the service account.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> str:
        """
        The unique, stable numeric ID for the service account. Each service account retains its unique ID even if you delete the service account. For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account.
        """
        return pulumi.get(self, "unique_id")


class AwaitableGetServiceAccountResult(GetServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountResult(
            description=self.description,
            disabled=self.disabled,
            display_name=self.display_name,
            email=self.email,
            etag=self.etag,
            name=self.name,
            oauth2_client_id=self.oauth2_client_id,
            project=self.project,
            unique_id=self.unique_id)


def get_service_account(project: Optional[str] = None,
                        service_account_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountResult:
    """
    Gets a ServiceAccount.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['serviceAccountId'] = service_account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:iam/v1:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountResult).value

    return AwaitableGetServiceAccountResult(
        description=pulumi.get(__ret__, 'description'),
        disabled=pulumi.get(__ret__, 'disabled'),
        display_name=pulumi.get(__ret__, 'display_name'),
        email=pulumi.get(__ret__, 'email'),
        etag=pulumi.get(__ret__, 'etag'),
        name=pulumi.get(__ret__, 'name'),
        oauth2_client_id=pulumi.get(__ret__, 'oauth2_client_id'),
        project=pulumi.get(__ret__, 'project'),
        unique_id=pulumi.get(__ret__, 'unique_id'))


@_utilities.lift_output_func(get_service_account)
def get_service_account_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                               service_account_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceAccountResult]:
    """
    Gets a ServiceAccount.
    """
    ...
