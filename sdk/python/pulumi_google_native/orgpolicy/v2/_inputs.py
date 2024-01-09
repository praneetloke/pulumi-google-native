# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'GoogleCloudOrgpolicyV2AlternatePolicySpecArgs',
    'GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs',
    'GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs',
    'GoogleCloudOrgpolicyV2PolicySpecArgs',
    'GoogleTypeExprArgs',
]

@pulumi.input_type
class GoogleCloudOrgpolicyV2AlternatePolicySpecArgs:
    def __init__(__self__, *,
                 launch: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecArgs']] = None):
        """
        Similar to PolicySpec but with an extra 'launch' field for launch reference. The PolicySpec here is specific for dry-run/darklaunch.
        :param pulumi.Input[str] launch: Reference to the launch that will be used while audit logging and to control the launch. Should be set only in the alternate policy.
        :param pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecArgs'] spec: Specify constraint for configurations of Google Cloud resources.
        """
        if launch is not None:
            pulumi.set(__self__, "launch", launch)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def launch(self) -> Optional[pulumi.Input[str]]:
        """
        Reference to the launch that will be used while audit logging and to control the launch. Should be set only in the alternate policy.
        """
        return pulumi.get(self, "launch")

    @launch.setter
    def launch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecArgs']]:
        """
        Specify constraint for configurations of Google Cloud resources.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs:
    def __init__(__self__, *,
                 allowed_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 denied_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        A message that holds specific allowed and denied values. This message can define specific values and subtrees of the Resource Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that are allowed or denied. This is achieved by using the `under:` and optional `is:` prefixes. The `under:` prefix is used to denote resource subtree values. The `is:` prefix is used to denote specific values, and is required only if the value contains a ":". Values prefixed with "is:" are treated the same as values with no prefix. Ancestry subtrees must be in one of the following formats: - `projects/` (for example, `projects/tokyo-rain-123`) - `folders/` (for example, `folders/1234`) - `organizations/` (for example, `organizations/1234`) The `supports_under` field of the associated `Constraint` defines whether ancestry prefixes can be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_values: List of values allowed at this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] denied_values: List of values denied at this resource.
        """
        if allowed_values is not None:
            pulumi.set(__self__, "allowed_values", allowed_values)
        if denied_values is not None:
            pulumi.set(__self__, "denied_values", denied_values)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values allowed at this resource.
        """
        return pulumi.get(self, "allowed_values")

    @allowed_values.setter
    def allowed_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_values", value)

    @property
    @pulumi.getter(name="deniedValues")
    def denied_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values denied at this resource.
        """
        return pulumi.get(self, "denied_values")

    @denied_values.setter
    def denied_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "denied_values", value)


@pulumi.input_type
class GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs:
    def __init__(__self__, *,
                 allow_all: Optional[pulumi.Input[bool]] = None,
                 condition: Optional[pulumi.Input['GoogleTypeExprArgs']] = None,
                 deny_all: Optional[pulumi.Input[bool]] = None,
                 enforce: Optional[pulumi.Input[bool]] = None,
                 values: Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs']] = None):
        """
        A rule used to express this policy.
        :param pulumi.Input[bool] allow_all: Setting this to true means that all values are allowed. This field can be set only in policies for list constraints.
        :param pulumi.Input['GoogleTypeExprArgs'] condition: A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        :param pulumi.Input[bool] deny_all: Setting this to true means that all values are denied. This field can be set only in policies for list constraints.
        :param pulumi.Input[bool] enforce: If `true`, then the policy is enforced. If `false`, then any configuration is acceptable. This field can be set only in policies for boolean constraints.
        :param pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs'] values: List of values to be used for this policy rule. This field can be set only in policies for list constraints.
        """
        if allow_all is not None:
            pulumi.set(__self__, "allow_all", allow_all)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if deny_all is not None:
            pulumi.set(__self__, "deny_all", deny_all)
        if enforce is not None:
            pulumi.set(__self__, "enforce", enforce)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="allowAll")
    def allow_all(self) -> Optional[pulumi.Input[bool]]:
        """
        Setting this to true means that all values are allowed. This field can be set only in policies for list constraints.
        """
        return pulumi.get(self, "allow_all")

    @allow_all.setter
    def allow_all(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_all", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['GoogleTypeExprArgs']]:
        """
        A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['GoogleTypeExprArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="denyAll")
    def deny_all(self) -> Optional[pulumi.Input[bool]]:
        """
        Setting this to true means that all values are denied. This field can be set only in policies for list constraints.
        """
        return pulumi.get(self, "deny_all")

    @deny_all.setter
    def deny_all(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deny_all", value)

    @property
    @pulumi.getter
    def enforce(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, then the policy is enforced. If `false`, then any configuration is acceptable. This field can be set only in policies for boolean constraints.
        """
        return pulumi.get(self, "enforce")

    @enforce.setter
    def enforce(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs']]:
        """
        List of values to be used for this policy rule. This field can be set only in policies for list constraints.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs']]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GoogleCloudOrgpolicyV2PolicySpecArgs:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 inherit_from_parent: Optional[pulumi.Input[bool]] = None,
                 reset: Optional[pulumi.Input[bool]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs']]]] = None):
        """
        Defines a Google Cloud policy specification which is used to specify constraints for configurations of Google Cloud resources.
        :param pulumi.Input[str] etag: An opaque tag indicating the current version of the policySpec, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policySpec to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        :param pulumi.Input[bool] inherit_from_parent: Determines the inheritance behavior for this policy. If `inherit_from_parent` is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
        :param pulumi.Input[bool] reset: Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs']]] rules: In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set `enforced` to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if inherit_from_parent is not None:
            pulumi.set(__self__, "inherit_from_parent", inherit_from_parent)
        if reset is not None:
            pulumi.set(__self__, "reset", reset)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        An opaque tag indicating the current version of the policySpec, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policySpec to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines the inheritance behavior for this policy. If `inherit_from_parent` is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
        """
        return pulumi.get(self, "inherit_from_parent")

    @inherit_from_parent.setter
    def inherit_from_parent(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "inherit_from_parent", value)

    @property
    @pulumi.getter
    def reset(self) -> Optional[pulumi.Input[bool]]:
        """
        Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        """
        return pulumi.get(self, "reset")

    @reset.setter
    def reset(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs']]]]:
        """
        In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set `enforced` to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class GoogleTypeExprArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param pulumi.Input[str] description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


