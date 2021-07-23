# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetReservationIamPolicyResult',
    'AwaitableGetReservationIamPolicyResult',
    'get_reservation_iam_policy',
]

@pulumi.output_type
class GetReservationIamPolicyResult:
    def __init__(__self__, audit_configs=None, bindings=None, etag=None, iam_owned=None, rules=None, version=None):
        if audit_configs and not isinstance(audit_configs, list):
            raise TypeError("Expected argument 'audit_configs' to be a list")
        pulumi.set(__self__, "audit_configs", audit_configs)
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        pulumi.set(__self__, "bindings", bindings)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if iam_owned and not isinstance(iam_owned, bool):
            raise TypeError("Expected argument 'iam_owned' to be a bool")
        pulumi.set(__self__, "iam_owned", iam_owned)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="auditConfigs")
    def audit_configs(self) -> Sequence['outputs.AuditConfigResponse']:
        """
        Specifies cloud audit logging configuration for this policy.
        """
        return pulumi.get(self, "audit_configs")

    @property
    @pulumi.getter
    def bindings(self) -> Sequence['outputs.BindingResponse']:
        """
        Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
        """
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="iamOwned")
    def iam_owned(self) -> bool:
        """
        This is deprecated and has no effect. Do not use.
        """
        return pulumi.get(self, "iam_owned")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.RuleResponse']:
        """
        This is deprecated and has no effect. Do not use.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def version(self) -> int:
        """
        Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "version")


class AwaitableGetReservationIamPolicyResult(GetReservationIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReservationIamPolicyResult(
            audit_configs=self.audit_configs,
            bindings=self.bindings,
            etag=self.etag,
            iam_owned=self.iam_owned,
            rules=self.rules,
            version=self.version)


def get_reservation_iam_policy(options_requested_policy_version: Optional[str] = None,
                               project: Optional[str] = None,
                               resource: Optional[str] = None,
                               zone: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReservationIamPolicyResult:
    """
    Gets the access control policy for a resource. May be empty if no such policy or resource exists.
    """
    __args__ = dict()
    __args__['optionsRequestedPolicyVersion'] = options_requested_policy_version
    __args__['project'] = project
    __args__['resource'] = resource
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getReservationIamPolicy', __args__, opts=opts, typ=GetReservationIamPolicyResult).value

    return AwaitableGetReservationIamPolicyResult(
        audit_configs=__ret__.audit_configs,
        bindings=__ret__.bindings,
        etag=__ret__.etag,
        iam_owned=__ret__.iam_owned,
        rules=__ret__.rules,
        version=__ret__.version)