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
    'GetBucketAccessControlResult',
    'AwaitableGetBucketAccessControlResult',
    'get_bucket_access_control',
    'get_bucket_access_control_output',
]

@pulumi.output_type
class GetBucketAccessControlResult:
    def __init__(__self__, bucket=None, domain=None, email=None, entity=None, entity_id=None, etag=None, kind=None, project_team=None, role=None, self_link=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if entity and not isinstance(entity, str):
            raise TypeError("Expected argument 'entity' to be a str")
        pulumi.set(__self__, "entity", entity)
        if entity_id and not isinstance(entity_id, str):
            raise TypeError("Expected argument 'entity_id' to be a str")
        pulumi.set(__self__, "entity_id", entity_id)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if project_team and not isinstance(project_team, dict):
            raise TypeError("Expected argument 'project_team' to be a dict")
        pulumi.set(__self__, "project_team", project_team)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The domain associated with the entity, if any.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email address associated with the entity, if any.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def entity(self) -> str:
        """
        The entity holding the permission, in one of the following forms: 
        - user-userId 
        - user-email 
        - group-groupId 
        - group-email 
        - domain-domain 
        - project-team-projectId 
        - allUsers 
        - allAuthenticatedUsers Examples: 
        - The user liz@example.com would be user-liz@example.com. 
        - The group example@googlegroups.com would be group-example@googlegroups.com. 
        - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
        """
        return pulumi.get(self, "entity")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> str:
        """
        The ID for the entity, if any.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        HTTP 1.1 Entity tag for the access-control entry.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="projectTeam")
    def project_team(self) -> 'outputs.BucketAccessControlProjectTeamResponse':
        """
        The project team associated with the entity, if any.
        """
        return pulumi.get(self, "project_team")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The access permission for the entity.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The link to this access-control entry.
        """
        return pulumi.get(self, "self_link")


class AwaitableGetBucketAccessControlResult(GetBucketAccessControlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketAccessControlResult(
            bucket=self.bucket,
            domain=self.domain,
            email=self.email,
            entity=self.entity,
            entity_id=self.entity_id,
            etag=self.etag,
            kind=self.kind,
            project_team=self.project_team,
            role=self.role,
            self_link=self.self_link)


def get_bucket_access_control(bucket: Optional[str] = None,
                              entity: Optional[str] = None,
                              user_project: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketAccessControlResult:
    """
    Returns the ACL entry for the specified entity on the specified bucket.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['entity'] = entity
    __args__['userProject'] = user_project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:storage/v1:getBucketAccessControl', __args__, opts=opts, typ=GetBucketAccessControlResult).value

    return AwaitableGetBucketAccessControlResult(
        bucket=pulumi.get(__ret__, 'bucket'),
        domain=pulumi.get(__ret__, 'domain'),
        email=pulumi.get(__ret__, 'email'),
        entity=pulumi.get(__ret__, 'entity'),
        entity_id=pulumi.get(__ret__, 'entity_id'),
        etag=pulumi.get(__ret__, 'etag'),
        kind=pulumi.get(__ret__, 'kind'),
        project_team=pulumi.get(__ret__, 'project_team'),
        role=pulumi.get(__ret__, 'role'),
        self_link=pulumi.get(__ret__, 'self_link'))


@_utilities.lift_output_func(get_bucket_access_control)
def get_bucket_access_control_output(bucket: Optional[pulumi.Input[str]] = None,
                                     entity: Optional[pulumi.Input[str]] = None,
                                     user_project: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBucketAccessControlResult]:
    """
    Returns the ACL entry for the specified entity on the specified bucket.
    """
    ...
