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
    'GetRegionSecurityPolicyResult',
    'AwaitableGetRegionSecurityPolicyResult',
    'get_region_security_policy',
    'get_region_security_policy_output',
]

@pulumi.output_type
class GetRegionSecurityPolicyResult:
    def __init__(__self__, adaptive_protection_config=None, advanced_options_config=None, creation_timestamp=None, ddos_protection_config=None, description=None, fingerprint=None, kind=None, name=None, recaptcha_options_config=None, region=None, rules=None, self_link=None, type=None):
        if adaptive_protection_config and not isinstance(adaptive_protection_config, dict):
            raise TypeError("Expected argument 'adaptive_protection_config' to be a dict")
        pulumi.set(__self__, "adaptive_protection_config", adaptive_protection_config)
        if advanced_options_config and not isinstance(advanced_options_config, dict):
            raise TypeError("Expected argument 'advanced_options_config' to be a dict")
        pulumi.set(__self__, "advanced_options_config", advanced_options_config)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if ddos_protection_config and not isinstance(ddos_protection_config, dict):
            raise TypeError("Expected argument 'ddos_protection_config' to be a dict")
        pulumi.set(__self__, "ddos_protection_config", ddos_protection_config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if recaptcha_options_config and not isinstance(recaptcha_options_config, dict):
            raise TypeError("Expected argument 'recaptcha_options_config' to be a dict")
        pulumi.set(__self__, "recaptcha_options_config", recaptcha_options_config)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="adaptiveProtectionConfig")
    def adaptive_protection_config(self) -> 'outputs.SecurityPolicyAdaptiveProtectionConfigResponse':
        return pulumi.get(self, "adaptive_protection_config")

    @property
    @pulumi.getter(name="advancedOptionsConfig")
    def advanced_options_config(self) -> 'outputs.SecurityPolicyAdvancedOptionsConfigResponse':
        return pulumi.get(self, "advanced_options_config")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> 'outputs.SecurityPolicyDdosProtectionConfigResponse':
        return pulumi.get(self, "ddos_protection_config")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        [Output only] Type of the resource. Always compute#securityPolicyfor security policies
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recaptchaOptionsConfig")
    def recaptcha_options_config(self) -> 'outputs.SecurityPolicyRecaptchaOptionsConfigResponse':
        return pulumi.get(self, "recaptcha_options_config")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional security policy resides. This field is not applicable to global security policies.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.SecurityPolicyRuleResponse']:
        """
        A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        """
        return pulumi.get(self, "type")


class AwaitableGetRegionSecurityPolicyResult(GetRegionSecurityPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionSecurityPolicyResult(
            adaptive_protection_config=self.adaptive_protection_config,
            advanced_options_config=self.advanced_options_config,
            creation_timestamp=self.creation_timestamp,
            ddos_protection_config=self.ddos_protection_config,
            description=self.description,
            fingerprint=self.fingerprint,
            kind=self.kind,
            name=self.name,
            recaptcha_options_config=self.recaptcha_options_config,
            region=self.region,
            rules=self.rules,
            self_link=self.self_link,
            type=self.type)


def get_region_security_policy(project: Optional[str] = None,
                               region: Optional[str] = None,
                               security_policy: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionSecurityPolicyResult:
    """
    List all of the ordered rules present in a single specified policy.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['securityPolicy'] = security_policy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getRegionSecurityPolicy', __args__, opts=opts, typ=GetRegionSecurityPolicyResult).value

    return AwaitableGetRegionSecurityPolicyResult(
        adaptive_protection_config=__ret__.adaptive_protection_config,
        advanced_options_config=__ret__.advanced_options_config,
        creation_timestamp=__ret__.creation_timestamp,
        ddos_protection_config=__ret__.ddos_protection_config,
        description=__ret__.description,
        fingerprint=__ret__.fingerprint,
        kind=__ret__.kind,
        name=__ret__.name,
        recaptcha_options_config=__ret__.recaptcha_options_config,
        region=__ret__.region,
        rules=__ret__.rules,
        self_link=__ret__.self_link,
        type=__ret__.type)


@_utilities.lift_output_func(get_region_security_policy)
def get_region_security_policy_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                      region: Optional[pulumi.Input[str]] = None,
                                      security_policy: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionSecurityPolicyResult]:
    """
    List all of the ordered rules present in a single specified policy.
    """
    ...
