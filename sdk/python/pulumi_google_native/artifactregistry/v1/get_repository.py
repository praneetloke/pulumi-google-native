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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
    'get_repository_output',
]

@pulumi.output_type
class GetRepositoryResult:
    def __init__(__self__, create_time=None, description=None, docker_config=None, format=None, kms_key_name=None, labels=None, maven_config=None, mode=None, name=None, remote_repository_config=None, satisfies_pzs=None, size_bytes=None, update_time=None, virtual_repository_config=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if docker_config and not isinstance(docker_config, dict):
            raise TypeError("Expected argument 'docker_config' to be a dict")
        pulumi.set(__self__, "docker_config", docker_config)
        if format and not isinstance(format, str):
            raise TypeError("Expected argument 'format' to be a str")
        pulumi.set(__self__, "format", format)
        if kms_key_name and not isinstance(kms_key_name, str):
            raise TypeError("Expected argument 'kms_key_name' to be a str")
        pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if maven_config and not isinstance(maven_config, dict):
            raise TypeError("Expected argument 'maven_config' to be a dict")
        pulumi.set(__self__, "maven_config", maven_config)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if remote_repository_config and not isinstance(remote_repository_config, dict):
            raise TypeError("Expected argument 'remote_repository_config' to be a dict")
        pulumi.set(__self__, "remote_repository_config", remote_repository_config)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if size_bytes and not isinstance(size_bytes, str):
            raise TypeError("Expected argument 'size_bytes' to be a str")
        pulumi.set(__self__, "size_bytes", size_bytes)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if virtual_repository_config and not isinstance(virtual_repository_config, dict):
            raise TypeError("Expected argument 'virtual_repository_config' to be a dict")
        pulumi.set(__self__, "virtual_repository_config", virtual_repository_config)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the repository was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The user-provided description of the repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dockerConfig")
    def docker_config(self) -> 'outputs.DockerRepositoryConfigResponse':
        """
        Docker repository config contains repository level configuration for the repositories of docker type.
        """
        return pulumi.get(self, "docker_config")

    @property
    @pulumi.getter
    def format(self) -> str:
        """
        The format of packages that are stored in the repository.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mavenConfig")
    def maven_config(self) -> 'outputs.MavenRepositoryConfigResponse':
        """
        Maven repository config contains repository level configuration for the repositories of maven type.
        """
        return pulumi.get(self, "maven_config")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        The mode of the repository.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteRepositoryConfig")
    def remote_repository_config(self) -> 'outputs.RemoteRepositoryConfigResponse':
        """
        Configuration specific for a Remote Repository.
        """
        return pulumi.get(self, "remote_repository_config")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        If set, the repository satisfies physical zone separation.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="sizeBytes")
    def size_bytes(self) -> str:
        """
        The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
        """
        return pulumi.get(self, "size_bytes")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the repository was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="virtualRepositoryConfig")
    def virtual_repository_config(self) -> 'outputs.VirtualRepositoryConfigResponse':
        """
        Configuration specific for a Virtual Repository.
        """
        return pulumi.get(self, "virtual_repository_config")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            create_time=self.create_time,
            description=self.description,
            docker_config=self.docker_config,
            format=self.format,
            kms_key_name=self.kms_key_name,
            labels=self.labels,
            maven_config=self.maven_config,
            mode=self.mode,
            name=self.name,
            remote_repository_config=self.remote_repository_config,
            satisfies_pzs=self.satisfies_pzs,
            size_bytes=self.size_bytes,
            update_time=self.update_time,
            virtual_repository_config=self.virtual_repository_config)


def get_repository(location: Optional[str] = None,
                   project: Optional[str] = None,
                   repository_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    Gets a repository.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:artifactregistry/v1:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        docker_config=pulumi.get(__ret__, 'docker_config'),
        format=pulumi.get(__ret__, 'format'),
        kms_key_name=pulumi.get(__ret__, 'kms_key_name'),
        labels=pulumi.get(__ret__, 'labels'),
        maven_config=pulumi.get(__ret__, 'maven_config'),
        mode=pulumi.get(__ret__, 'mode'),
        name=pulumi.get(__ret__, 'name'),
        remote_repository_config=pulumi.get(__ret__, 'remote_repository_config'),
        satisfies_pzs=pulumi.get(__ret__, 'satisfies_pzs'),
        size_bytes=pulumi.get(__ret__, 'size_bytes'),
        update_time=pulumi.get(__ret__, 'update_time'),
        virtual_repository_config=pulumi.get(__ret__, 'virtual_repository_config'))


@_utilities.lift_output_func(get_repository)
def get_repository_output(location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          repository_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryResult]:
    """
    Gets a repository.
    """
    ...
