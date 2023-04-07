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
    'GetWorkforcePoolProviderResult',
    'AwaitableGetWorkforcePoolProviderResult',
    'get_workforce_pool_provider',
    'get_workforce_pool_provider_output',
]

@pulumi.output_type
class GetWorkforcePoolProviderResult:
    def __init__(__self__, attribute_condition=None, attribute_mapping=None, description=None, disabled=None, display_name=None, name=None, oidc=None, saml=None, state=None):
        if attribute_condition and not isinstance(attribute_condition, str):
            raise TypeError("Expected argument 'attribute_condition' to be a str")
        pulumi.set(__self__, "attribute_condition", attribute_condition)
        if attribute_mapping and not isinstance(attribute_mapping, dict):
            raise TypeError("Expected argument 'attribute_mapping' to be a dict")
        pulumi.set(__self__, "attribute_mapping", attribute_mapping)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if oidc and not isinstance(oidc, dict):
            raise TypeError("Expected argument 'oidc' to be a dict")
        pulumi.set(__self__, "oidc", oidc)
        if saml and not isinstance(saml, dict):
            raise TypeError("Expected argument 'saml' to be a dict")
        pulumi.set(__self__, "saml", saml)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> str:
        """
        A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        """
        return pulumi.get(self, "attribute_condition")

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> Mapping[str, str]:
        """
        Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/locations/global/workforcePools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 4KB. For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        """
        return pulumi.get(self, "attribute_mapping")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A user-specified description of the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        A user-specified display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the provider. Format: `locations/{location}/workforcePools/{workforce_pool_id}/providers/{provider_id}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def oidc(self) -> 'outputs.GoogleIamAdminV1WorkforcePoolProviderOidcResponse':
        """
        An OpenId Connect 1.0 identity provider configuration.
        """
        return pulumi.get(self, "oidc")

    @property
    @pulumi.getter
    def saml(self) -> 'outputs.GoogleIamAdminV1WorkforcePoolProviderSamlResponse':
        """
        A SAML identity provider configuration.
        """
        return pulumi.get(self, "saml")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the provider.
        """
        return pulumi.get(self, "state")


class AwaitableGetWorkforcePoolProviderResult(GetWorkforcePoolProviderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkforcePoolProviderResult(
            attribute_condition=self.attribute_condition,
            attribute_mapping=self.attribute_mapping,
            description=self.description,
            disabled=self.disabled,
            display_name=self.display_name,
            name=self.name,
            oidc=self.oidc,
            saml=self.saml,
            state=self.state)


def get_workforce_pool_provider(location: Optional[str] = None,
                                provider_id: Optional[str] = None,
                                workforce_pool_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkforcePoolProviderResult:
    """
    Gets an individual WorkforcePoolProvider.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['providerId'] = provider_id
    __args__['workforcePoolId'] = workforce_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:iam/v1:getWorkforcePoolProvider', __args__, opts=opts, typ=GetWorkforcePoolProviderResult).value

    return AwaitableGetWorkforcePoolProviderResult(
        attribute_condition=__ret__.attribute_condition,
        attribute_mapping=__ret__.attribute_mapping,
        description=__ret__.description,
        disabled=__ret__.disabled,
        display_name=__ret__.display_name,
        name=__ret__.name,
        oidc=__ret__.oidc,
        saml=__ret__.saml,
        state=__ret__.state)


@_utilities.lift_output_func(get_workforce_pool_provider)
def get_workforce_pool_provider_output(location: Optional[pulumi.Input[str]] = None,
                                       provider_id: Optional[pulumi.Input[str]] = None,
                                       workforce_pool_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkforcePoolProviderResult]:
    """
    Gets an individual WorkforcePoolProvider.
    """
    ...
