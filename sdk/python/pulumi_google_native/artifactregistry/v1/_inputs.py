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
    'BindingArgs',
    'DockerRepositoryConfigArgs',
    'DockerRepositoryArgs',
    'ExprArgs',
    'MavenRepositoryConfigArgs',
    'MavenRepositoryArgs',
    'NpmRepositoryArgs',
    'PythonRepositoryArgs',
    'RemoteRepositoryConfigArgs',
    'UpstreamPolicyArgs',
    'VirtualRepositoryConfigArgs',
]

@pulumi.input_type
class BindingArgs:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['ExprArgs']] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Associates `members`, or principals, with a `role`.
        :param pulumi.Input['ExprArgs'] condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
        :param pulumi.Input[str] role: Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['ExprArgs']]:
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['ExprArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class DockerRepositoryConfigArgs:
    def __init__(__self__, *,
                 immutable_tags: Optional[pulumi.Input[bool]] = None):
        """
        DockerRepositoryConfig is docker related repository details. Provides additional configuration details for repositories of the docker format type.
        :param pulumi.Input[bool] immutable_tags: The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
        """
        if immutable_tags is not None:
            pulumi.set(__self__, "immutable_tags", immutable_tags)

    @property
    @pulumi.getter(name="immutableTags")
    def immutable_tags(self) -> Optional[pulumi.Input[bool]]:
        """
        The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
        """
        return pulumi.get(self, "immutable_tags")

    @immutable_tags.setter
    def immutable_tags(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "immutable_tags", value)


@pulumi.input_type
class DockerRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['DockerRepositoryPublicRepository']] = None):
        """
        Configuration for a Docker remote repository.
        :param pulumi.Input['DockerRepositoryPublicRepository'] public_repository: One of the publicly available Docker repositories supported by Artifact Registry.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['DockerRepositoryPublicRepository']]:
        """
        One of the publicly available Docker repositories supported by Artifact Registry.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['DockerRepositoryPublicRepository']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class ExprArgs:
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


@pulumi.input_type
class MavenRepositoryConfigArgs:
    def __init__(__self__, *,
                 allow_snapshot_overwrites: Optional[pulumi.Input[bool]] = None,
                 version_policy: Optional[pulumi.Input['MavenRepositoryConfigVersionPolicy']] = None):
        """
        MavenRepositoryConfig is maven related repository details. Provides additional configuration details for repositories of the maven format type.
        :param pulumi.Input[bool] allow_snapshot_overwrites: The repository with this flag will allow publishing the same snapshot versions.
        :param pulumi.Input['MavenRepositoryConfigVersionPolicy'] version_policy: Version policy defines the versions that the registry will accept.
        """
        if allow_snapshot_overwrites is not None:
            pulumi.set(__self__, "allow_snapshot_overwrites", allow_snapshot_overwrites)
        if version_policy is not None:
            pulumi.set(__self__, "version_policy", version_policy)

    @property
    @pulumi.getter(name="allowSnapshotOverwrites")
    def allow_snapshot_overwrites(self) -> Optional[pulumi.Input[bool]]:
        """
        The repository with this flag will allow publishing the same snapshot versions.
        """
        return pulumi.get(self, "allow_snapshot_overwrites")

    @allow_snapshot_overwrites.setter
    def allow_snapshot_overwrites(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_snapshot_overwrites", value)

    @property
    @pulumi.getter(name="versionPolicy")
    def version_policy(self) -> Optional[pulumi.Input['MavenRepositoryConfigVersionPolicy']]:
        """
        Version policy defines the versions that the registry will accept.
        """
        return pulumi.get(self, "version_policy")

    @version_policy.setter
    def version_policy(self, value: Optional[pulumi.Input['MavenRepositoryConfigVersionPolicy']]):
        pulumi.set(self, "version_policy", value)


@pulumi.input_type
class MavenRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['MavenRepositoryPublicRepository']] = None):
        """
        Configuration for a Maven remote repository.
        :param pulumi.Input['MavenRepositoryPublicRepository'] public_repository: One of the publicly available Maven repositories supported by Artifact Registry.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['MavenRepositoryPublicRepository']]:
        """
        One of the publicly available Maven repositories supported by Artifact Registry.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['MavenRepositoryPublicRepository']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class NpmRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['NpmRepositoryPublicRepository']] = None):
        """
        Configuration for a Npm remote repository.
        :param pulumi.Input['NpmRepositoryPublicRepository'] public_repository: One of the publicly available Npm repositories supported by Artifact Registry.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['NpmRepositoryPublicRepository']]:
        """
        One of the publicly available Npm repositories supported by Artifact Registry.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['NpmRepositoryPublicRepository']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class PythonRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['PythonRepositoryPublicRepository']] = None):
        """
        Configuration for a Python remote repository.
        :param pulumi.Input['PythonRepositoryPublicRepository'] public_repository: One of the publicly available Python repositories supported by Artifact Registry.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['PythonRepositoryPublicRepository']]:
        """
        One of the publicly available Python repositories supported by Artifact Registry.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['PythonRepositoryPublicRepository']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RemoteRepositoryConfigArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 docker_repository: Optional[pulumi.Input['DockerRepositoryArgs']] = None,
                 maven_repository: Optional[pulumi.Input['MavenRepositoryArgs']] = None,
                 npm_repository: Optional[pulumi.Input['NpmRepositoryArgs']] = None,
                 python_repository: Optional[pulumi.Input['PythonRepositoryArgs']] = None):
        """
        Remote repository configuration.
        :param pulumi.Input[str] description: The description of the remote source.
        :param pulumi.Input['DockerRepositoryArgs'] docker_repository: Specific settings for a Docker remote repository.
        :param pulumi.Input['MavenRepositoryArgs'] maven_repository: Specific settings for a Maven remote repository.
        :param pulumi.Input['NpmRepositoryArgs'] npm_repository: Specific settings for an Npm remote repository.
        :param pulumi.Input['PythonRepositoryArgs'] python_repository: Specific settings for a Python remote repository.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if docker_repository is not None:
            pulumi.set(__self__, "docker_repository", docker_repository)
        if maven_repository is not None:
            pulumi.set(__self__, "maven_repository", maven_repository)
        if npm_repository is not None:
            pulumi.set(__self__, "npm_repository", npm_repository)
        if python_repository is not None:
            pulumi.set(__self__, "python_repository", python_repository)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the remote source.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dockerRepository")
    def docker_repository(self) -> Optional[pulumi.Input['DockerRepositoryArgs']]:
        """
        Specific settings for a Docker remote repository.
        """
        return pulumi.get(self, "docker_repository")

    @docker_repository.setter
    def docker_repository(self, value: Optional[pulumi.Input['DockerRepositoryArgs']]):
        pulumi.set(self, "docker_repository", value)

    @property
    @pulumi.getter(name="mavenRepository")
    def maven_repository(self) -> Optional[pulumi.Input['MavenRepositoryArgs']]:
        """
        Specific settings for a Maven remote repository.
        """
        return pulumi.get(self, "maven_repository")

    @maven_repository.setter
    def maven_repository(self, value: Optional[pulumi.Input['MavenRepositoryArgs']]):
        pulumi.set(self, "maven_repository", value)

    @property
    @pulumi.getter(name="npmRepository")
    def npm_repository(self) -> Optional[pulumi.Input['NpmRepositoryArgs']]:
        """
        Specific settings for an Npm remote repository.
        """
        return pulumi.get(self, "npm_repository")

    @npm_repository.setter
    def npm_repository(self, value: Optional[pulumi.Input['NpmRepositoryArgs']]):
        pulumi.set(self, "npm_repository", value)

    @property
    @pulumi.getter(name="pythonRepository")
    def python_repository(self) -> Optional[pulumi.Input['PythonRepositoryArgs']]:
        """
        Specific settings for a Python remote repository.
        """
        return pulumi.get(self, "python_repository")

    @python_repository.setter
    def python_repository(self, value: Optional[pulumi.Input['PythonRepositoryArgs']]):
        pulumi.set(self, "python_repository", value)


@pulumi.input_type
class UpstreamPolicyArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 repository: Optional[pulumi.Input[str]] = None):
        """
        Artifact policy configuration for the repository contents.
        :param pulumi.Input[str] id: The user-provided ID of the upstream policy.
        :param pulumi.Input[int] priority: Entries with a greater priority value take precedence in the pull order.
        :param pulumi.Input[str] repository: A reference to the repository resource, for example: "projects/p1/locations/us-central1/repositories/repo1".
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The user-provided ID of the upstream policy.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Entries with a greater priority value take precedence in the pull order.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        A reference to the repository resource, for example: "projects/p1/locations/us-central1/repositories/repo1".
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)


@pulumi.input_type
class VirtualRepositoryConfigArgs:
    def __init__(__self__, *,
                 upstream_policies: Optional[pulumi.Input[Sequence[pulumi.Input['UpstreamPolicyArgs']]]] = None):
        """
        Virtual repository configuration.
        :param pulumi.Input[Sequence[pulumi.Input['UpstreamPolicyArgs']]] upstream_policies: Policies that configure the upstream artifacts distributed by the Virtual Repository. Upstream policies cannot be set on a standard repository.
        """
        if upstream_policies is not None:
            pulumi.set(__self__, "upstream_policies", upstream_policies)

    @property
    @pulumi.getter(name="upstreamPolicies")
    def upstream_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UpstreamPolicyArgs']]]]:
        """
        Policies that configure the upstream artifacts distributed by the Virtual Repository. Upstream policies cannot be set on a standard repository.
        """
        return pulumi.get(self, "upstream_policies")

    @upstream_policies.setter
    def upstream_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UpstreamPolicyArgs']]]]):
        pulumi.set(self, "upstream_policies", value)


