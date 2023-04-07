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
from ._inputs import *

__all__ = ['WorkforcePoolProviderArgs', 'WorkforcePoolProvider']

@pulumi.input_type
class WorkforcePoolProviderArgs:
    def __init__(__self__, *,
                 attribute_mapping: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 workforce_pool_id: pulumi.Input[str],
                 workforce_pool_provider_id: pulumi.Input[str],
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']] = None,
                 saml: Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']] = None):
        """
        The set of arguments for constructing a WorkforcePoolProvider resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/locations/global/workforcePools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 4KB. For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        :param pulumi.Input[str] workforce_pool_provider_id: Required. The ID for the provider, which becomes the final component of the resource name. This value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix `gcp-` is reserved for use by Google, and may not be specified.
        :param pulumi.Input[str] attribute_condition: A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        :param pulumi.Input[str] description: A user-specified description of the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A user-specified display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderOidcArgs'] oidc: An OpenId Connect 1.0 identity provider configuration.
        :param pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderSamlArgs'] saml: A SAML identity provider configuration.
        """
        pulumi.set(__self__, "attribute_mapping", attribute_mapping)
        pulumi.set(__self__, "workforce_pool_id", workforce_pool_id)
        pulumi.set(__self__, "workforce_pool_provider_id", workforce_pool_provider_id)
        if attribute_condition is not None:
            pulumi.set(__self__, "attribute_condition", attribute_condition)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if oidc is not None:
            pulumi.set(__self__, "oidc", oidc)
        if saml is not None:
            pulumi.set(__self__, "saml", saml)

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/locations/global/workforcePools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 4KB. For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        """
        return pulumi.get(self, "attribute_mapping")

    @attribute_mapping.setter
    def attribute_mapping(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "attribute_mapping", value)

    @property
    @pulumi.getter(name="workforcePoolId")
    def workforce_pool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workforce_pool_id")

    @workforce_pool_id.setter
    def workforce_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workforce_pool_id", value)

    @property
    @pulumi.getter(name="workforcePoolProviderId")
    def workforce_pool_provider_id(self) -> pulumi.Input[str]:
        """
        Required. The ID for the provider, which becomes the final component of the resource name. This value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix `gcp-` is reserved for use by Google, and may not be specified.
        """
        return pulumi.get(self, "workforce_pool_provider_id")

    @workforce_pool_provider_id.setter
    def workforce_pool_provider_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workforce_pool_provider_id", value)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> Optional[pulumi.Input[str]]:
        """
        A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        """
        return pulumi.get(self, "attribute_condition")

    @attribute_condition.setter
    def attribute_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_condition", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-specified description of the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-specified display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def oidc(self) -> Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']]:
        """
        An OpenId Connect 1.0 identity provider configuration.
        """
        return pulumi.get(self, "oidc")

    @oidc.setter
    def oidc(self, value: Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']]):
        pulumi.set(self, "oidc", value)

    @property
    @pulumi.getter
    def saml(self) -> Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']]:
        """
        A SAML identity provider configuration.
        """
        return pulumi.get(self, "saml")

    @saml.setter
    def saml(self, value: Optional[pulumi.Input['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']]):
        pulumi.set(self, "saml", value)


class WorkforcePoolProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']]] = None,
                 saml: Optional[pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']]] = None,
                 workforce_pool_id: Optional[pulumi.Input[str]] = None,
                 workforce_pool_provider_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new WorkforcePoolProvider in a WorkforcePool. You cannot reuse the name of a deleted provider until 30 days after deletion.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_condition: A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/locations/global/workforcePools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 4KB. For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        :param pulumi.Input[str] description: A user-specified description of the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A user-specified display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']] oidc: An OpenId Connect 1.0 identity provider configuration.
        :param pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']] saml: A SAML identity provider configuration.
        :param pulumi.Input[str] workforce_pool_provider_id: Required. The ID for the provider, which becomes the final component of the resource name. This value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix `gcp-` is reserved for use by Google, and may not be specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkforcePoolProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new WorkforcePoolProvider in a WorkforcePool. You cannot reuse the name of a deleted provider until 30 days after deletion.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param WorkforcePoolProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkforcePoolProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderOidcArgs']]] = None,
                 saml: Optional[pulumi.Input[pulumi.InputType['GoogleIamAdminV1WorkforcePoolProviderSamlArgs']]] = None,
                 workforce_pool_id: Optional[pulumi.Input[str]] = None,
                 workforce_pool_provider_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkforcePoolProviderArgs.__new__(WorkforcePoolProviderArgs)

            __props__.__dict__["attribute_condition"] = attribute_condition
            if attribute_mapping is None and not opts.urn:
                raise TypeError("Missing required property 'attribute_mapping'")
            __props__.__dict__["attribute_mapping"] = attribute_mapping
            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["location"] = location
            __props__.__dict__["oidc"] = oidc
            __props__.__dict__["saml"] = saml
            if workforce_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'workforce_pool_id'")
            __props__.__dict__["workforce_pool_id"] = workforce_pool_id
            if workforce_pool_provider_id is None and not opts.urn:
                raise TypeError("Missing required property 'workforce_pool_provider_id'")
            __props__.__dict__["workforce_pool_provider_id"] = workforce_pool_provider_id
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "workforce_pool_id", "workforce_pool_provider_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(WorkforcePoolProvider, __self__).__init__(
            'google-native:iam/v1:WorkforcePoolProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkforcePoolProvider':
        """
        Get an existing WorkforcePoolProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkforcePoolProviderArgs.__new__(WorkforcePoolProviderArgs)

        __props__.__dict__["attribute_condition"] = None
        __props__.__dict__["attribute_mapping"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["oidc"] = None
        __props__.__dict__["saml"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["workforce_pool_id"] = None
        __props__.__dict__["workforce_pool_provider_id"] = None
        return WorkforcePoolProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> pulumi.Output[str]:
        """
        A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        """
        return pulumi.get(self, "attribute_condition")

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/locations/global/workforcePools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 4KB. For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        """
        return pulumi.get(self, "attribute_mapping")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A user-specified description of the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A user-specified display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the provider. Format: `locations/{location}/workforcePools/{workforce_pool_id}/providers/{provider_id}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def oidc(self) -> pulumi.Output['outputs.GoogleIamAdminV1WorkforcePoolProviderOidcResponse']:
        """
        An OpenId Connect 1.0 identity provider configuration.
        """
        return pulumi.get(self, "oidc")

    @property
    @pulumi.getter
    def saml(self) -> pulumi.Output['outputs.GoogleIamAdminV1WorkforcePoolProviderSamlResponse']:
        """
        A SAML identity provider configuration.
        """
        return pulumi.get(self, "saml")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the provider.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workforcePoolId")
    def workforce_pool_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workforce_pool_id")

    @property
    @pulumi.getter(name="workforcePoolProviderId")
    def workforce_pool_provider_id(self) -> pulumi.Output[str]:
        """
        Required. The ID for the provider, which becomes the final component of the resource name. This value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix `gcp-` is reserved for use by Google, and may not be specified.
        """
        return pulumi.get(self, "workforce_pool_provider_id")

