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
from ._enums import *
from ._inputs import *

__all__ = ['RegionSecurityPolicyArgs', 'RegionSecurityPolicy']

@pulumi.input_type
class RegionSecurityPolicyArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 adaptive_protection_config: Optional[pulumi.Input['SecurityPolicyAdaptiveProtectionConfigArgs']] = None,
                 advanced_options_config: Optional[pulumi.Input['SecurityPolicyAdvancedOptionsConfigArgs']] = None,
                 ddos_protection_config: Optional[pulumi.Input['SecurityPolicyDdosProtectionConfigArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_options_config: Optional[pulumi.Input['SecurityPolicyRecaptchaOptionsConfigArgs']] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityPolicyRuleArgs']]]] = None,
                 type: Optional[pulumi.Input['RegionSecurityPolicyType']] = None):
        """
        The set of arguments for constructing a RegionSecurityPolicy resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[Sequence[pulumi.Input['SecurityPolicyRuleArgs']]] rules: A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        :param pulumi.Input['RegionSecurityPolicyType'] type: The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        """
        pulumi.set(__self__, "region", region)
        if adaptive_protection_config is not None:
            pulumi.set(__self__, "adaptive_protection_config", adaptive_protection_config)
        if advanced_options_config is not None:
            pulumi.set(__self__, "advanced_options_config", advanced_options_config)
        if ddos_protection_config is not None:
            pulumi.set(__self__, "ddos_protection_config", ddos_protection_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if recaptcha_options_config is not None:
            pulumi.set(__self__, "recaptcha_options_config", recaptcha_options_config)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="adaptiveProtectionConfig")
    def adaptive_protection_config(self) -> Optional[pulumi.Input['SecurityPolicyAdaptiveProtectionConfigArgs']]:
        return pulumi.get(self, "adaptive_protection_config")

    @adaptive_protection_config.setter
    def adaptive_protection_config(self, value: Optional[pulumi.Input['SecurityPolicyAdaptiveProtectionConfigArgs']]):
        pulumi.set(self, "adaptive_protection_config", value)

    @property
    @pulumi.getter(name="advancedOptionsConfig")
    def advanced_options_config(self) -> Optional[pulumi.Input['SecurityPolicyAdvancedOptionsConfigArgs']]:
        return pulumi.get(self, "advanced_options_config")

    @advanced_options_config.setter
    def advanced_options_config(self, value: Optional[pulumi.Input['SecurityPolicyAdvancedOptionsConfigArgs']]):
        pulumi.set(self, "advanced_options_config", value)

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> Optional[pulumi.Input['SecurityPolicyDdosProtectionConfigArgs']]:
        return pulumi.get(self, "ddos_protection_config")

    @ddos_protection_config.setter
    def ddos_protection_config(self, value: Optional[pulumi.Input['SecurityPolicyDdosProtectionConfigArgs']]):
        pulumi.set(self, "ddos_protection_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="recaptchaOptionsConfig")
    def recaptcha_options_config(self) -> Optional[pulumi.Input['SecurityPolicyRecaptchaOptionsConfigArgs']]:
        return pulumi.get(self, "recaptcha_options_config")

    @recaptcha_options_config.setter
    def recaptcha_options_config(self, value: Optional[pulumi.Input['SecurityPolicyRecaptchaOptionsConfigArgs']]):
        pulumi.set(self, "recaptcha_options_config", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityPolicyRuleArgs']]]]:
        """
        A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['RegionSecurityPolicyType']]:
        """
        The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['RegionSecurityPolicyType']]):
        pulumi.set(self, "type", value)


class RegionSecurityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adaptive_protection_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyAdaptiveProtectionConfigArgs']]] = None,
                 advanced_options_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyAdvancedOptionsConfigArgs']]] = None,
                 ddos_protection_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyDdosProtectionConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_options_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyRecaptchaOptionsConfigArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityPolicyRuleArgs']]]]] = None,
                 type: Optional[pulumi.Input['RegionSecurityPolicyType']] = None,
                 __props__=None):
        """
        Creates a new policy in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityPolicyRuleArgs']]]] rules: A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        :param pulumi.Input['RegionSecurityPolicyType'] type: The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionSecurityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new policy in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param RegionSecurityPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionSecurityPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adaptive_protection_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyAdaptiveProtectionConfigArgs']]] = None,
                 advanced_options_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyAdvancedOptionsConfigArgs']]] = None,
                 ddos_protection_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyDdosProtectionConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_options_config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyRecaptchaOptionsConfigArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityPolicyRuleArgs']]]]] = None,
                 type: Optional[pulumi.Input['RegionSecurityPolicyType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegionSecurityPolicyArgs.__new__(RegionSecurityPolicyArgs)

            __props__.__dict__["adaptive_protection_config"] = adaptive_protection_config
            __props__.__dict__["advanced_options_config"] = advanced_options_config
            __props__.__dict__["ddos_protection_config"] = ddos_protection_config
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["recaptcha_options_config"] = recaptcha_options_config
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["rules"] = rules
            __props__.__dict__["type"] = type
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["self_link"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project", "region"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(RegionSecurityPolicy, __self__).__init__(
            'google-native:compute/v1:RegionSecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionSecurityPolicy':
        """
        Get an existing RegionSecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionSecurityPolicyArgs.__new__(RegionSecurityPolicyArgs)

        __props__.__dict__["adaptive_protection_config"] = None
        __props__.__dict__["advanced_options_config"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["ddos_protection_config"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["recaptcha_options_config"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["type"] = None
        return RegionSecurityPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adaptiveProtectionConfig")
    def adaptive_protection_config(self) -> pulumi.Output['outputs.SecurityPolicyAdaptiveProtectionConfigResponse']:
        return pulumi.get(self, "adaptive_protection_config")

    @property
    @pulumi.getter(name="advancedOptionsConfig")
    def advanced_options_config(self) -> pulumi.Output['outputs.SecurityPolicyAdvancedOptionsConfigResponse']:
        return pulumi.get(self, "advanced_options_config")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> pulumi.Output['outputs.SecurityPolicyDdosProtectionConfigResponse']:
        return pulumi.get(self, "ddos_protection_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output only] Type of the resource. Always compute#securityPolicyfor security policies
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="recaptchaOptionsConfig")
    def recaptcha_options_config(self) -> pulumi.Output['outputs.SecurityPolicyRecaptchaOptionsConfigResponse']:
        return pulumi.get(self, "recaptcha_options_config")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.SecurityPolicyRuleResponse']]:
        """
        A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "type")

