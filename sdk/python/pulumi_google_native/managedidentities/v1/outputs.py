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
    'BindingResponse',
    'ExprResponse',
    'TrustResponse',
]

@pulumi.output_type
class BindingResponse(dict):
    """
    Associates `members`, or principals, with a `role`.
    """
    def __init__(__self__, *,
                 condition: 'outputs.ExprResponse',
                 members: Sequence[str],
                 role: str):
        """
        Associates `members`, or principals, with a `role`.
        :param 'ExprResponse' condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param Sequence[str] members: Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        :param str role: Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.ExprResponse':
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")


@pulumi.output_type
class ExprResponse(dict):
    """
    Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
    """
    def __init__(__self__, *,
                 description: str,
                 expression: str,
                 location: str,
                 title: str):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param str description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param str title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")


@pulumi.output_type
class TrustResponse(dict):
    """
    Represents a relationship between two domains. This allows a controller in one domain to authenticate a user in another domain. If the trust is being changed, it will be placed into the UPDATING state, which indicates that the resource is being reconciled. At this point, Get will reflect an intermediate state.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "lastTrustHeartbeatTime":
            suggest = "last_trust_heartbeat_time"
        elif key == "selectiveAuthentication":
            suggest = "selective_authentication"
        elif key == "stateDescription":
            suggest = "state_description"
        elif key == "targetDnsIpAddresses":
            suggest = "target_dns_ip_addresses"
        elif key == "targetDomainName":
            suggest = "target_domain_name"
        elif key == "trustDirection":
            suggest = "trust_direction"
        elif key == "trustHandshakeSecret":
            suggest = "trust_handshake_secret"
        elif key == "trustType":
            suggest = "trust_type"
        elif key == "updateTime":
            suggest = "update_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrustResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrustResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrustResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 create_time: str,
                 last_trust_heartbeat_time: str,
                 selective_authentication: bool,
                 state: str,
                 state_description: str,
                 target_dns_ip_addresses: Sequence[str],
                 target_domain_name: str,
                 trust_direction: str,
                 trust_handshake_secret: str,
                 trust_type: str,
                 update_time: str):
        """
        Represents a relationship between two domains. This allows a controller in one domain to authenticate a user in another domain. If the trust is being changed, it will be placed into the UPDATING state, which indicates that the resource is being reconciled. At this point, Get will reflect an intermediate state.
        :param str create_time: The time the instance was created.
        :param str last_trust_heartbeat_time: The last heartbeat time when the trust was known to be connected.
        :param bool selective_authentication: Optional. The trust authentication type, which decides whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        :param str state: The current state of the trust.
        :param str state_description: Additional information about the current state of the trust, if available.
        :param Sequence[str] target_dns_ip_addresses: The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        :param str target_domain_name: The fully qualified target domain name which will be in trust with the current domain.
        :param str trust_direction: The trust direction, which decides if the current domain is trusted, trusting, or both.
        :param str trust_handshake_secret: The trust secret used for the handshake with the target domain. This will not be stored.
        :param str trust_type: The type of trust represented by the trust resource.
        :param str update_time: The last update time.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "last_trust_heartbeat_time", last_trust_heartbeat_time)
        pulumi.set(__self__, "selective_authentication", selective_authentication)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "state_description", state_description)
        pulumi.set(__self__, "target_dns_ip_addresses", target_dns_ip_addresses)
        pulumi.set(__self__, "target_domain_name", target_domain_name)
        pulumi.set(__self__, "trust_direction", trust_direction)
        pulumi.set(__self__, "trust_handshake_secret", trust_handshake_secret)
        pulumi.set(__self__, "trust_type", trust_type)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="lastTrustHeartbeatTime")
    def last_trust_heartbeat_time(self) -> str:
        """
        The last heartbeat time when the trust was known to be connected.
        """
        return pulumi.get(self, "last_trust_heartbeat_time")

    @property
    @pulumi.getter(name="selectiveAuthentication")
    def selective_authentication(self) -> bool:
        """
        Optional. The trust authentication type, which decides whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
        """
        return pulumi.get(self, "selective_authentication")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the trust.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateDescription")
    def state_description(self) -> str:
        """
        Additional information about the current state of the trust, if available.
        """
        return pulumi.get(self, "state_description")

    @property
    @pulumi.getter(name="targetDnsIpAddresses")
    def target_dns_ip_addresses(self) -> Sequence[str]:
        """
        The target DNS server IP addresses which can resolve the remote domain involved in the trust.
        """
        return pulumi.get(self, "target_dns_ip_addresses")

    @property
    @pulumi.getter(name="targetDomainName")
    def target_domain_name(self) -> str:
        """
        The fully qualified target domain name which will be in trust with the current domain.
        """
        return pulumi.get(self, "target_domain_name")

    @property
    @pulumi.getter(name="trustDirection")
    def trust_direction(self) -> str:
        """
        The trust direction, which decides if the current domain is trusted, trusting, or both.
        """
        return pulumi.get(self, "trust_direction")

    @property
    @pulumi.getter(name="trustHandshakeSecret")
    def trust_handshake_secret(self) -> str:
        """
        The trust secret used for the handshake with the target domain. This will not be stored.
        """
        return pulumi.get(self, "trust_handshake_secret")

    @property
    @pulumi.getter(name="trustType")
    def trust_type(self) -> str:
        """
        The type of trust represented by the trust resource.
        """
        return pulumi.get(self, "trust_type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update time.
        """
        return pulumi.get(self, "update_time")


