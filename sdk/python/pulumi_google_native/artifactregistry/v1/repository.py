# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input['RepositoryFormat']] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maven_config: Optional[pulumi.Input['MavenRepositoryConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] create_time: The time when the repository was created.
        :param pulumi.Input[str] description: The user-provided description of the repository.
        :param pulumi.Input['RepositoryFormat'] format: The format of packages that are stored in the repository.
        :param pulumi.Input[str] kms_key_name: The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        :param pulumi.Input['MavenRepositoryConfigArgs'] maven_config: Maven repository config contains repository level configuration for the repositories of maven type.
        :param pulumi.Input[str] name: The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
        :param pulumi.Input[str] update_time: The time when the repository was last updated.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if maven_config is not None:
            pulumi.set(__self__, "maven_config", maven_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the repository was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The user-provided description of the repository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input['RepositoryFormat']]:
        """
        The format of packages that are stored in the repository.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input['RepositoryFormat']]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="mavenConfig")
    def maven_config(self) -> Optional[pulumi.Input['MavenRepositoryConfigArgs']]:
        """
        Maven repository config contains repository level configuration for the repositories of maven type.
        """
        return pulumi.get(self, "maven_config")

    @maven_config.setter
    def maven_config(self, value: Optional[pulumi.Input['MavenRepositoryConfigArgs']]):
        pulumi.set(self, "maven_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
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
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the repository was last updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input['RepositoryFormat']] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maven_config: Optional[pulumi.Input[pulumi.InputType['MavenRepositoryConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a repository. The returned Operation will finish once the repository has been created. Its response will be the created Repository.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time when the repository was created.
        :param pulumi.Input[str] description: The user-provided description of the repository.
        :param pulumi.Input['RepositoryFormat'] format: The format of packages that are stored in the repository.
        :param pulumi.Input[str] kms_key_name: The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        :param pulumi.Input[pulumi.InputType['MavenRepositoryConfigArgs']] maven_config: Maven repository config contains repository level configuration for the repositories of maven type.
        :param pulumi.Input[str] name: The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
        :param pulumi.Input[str] update_time: The time when the repository was last updated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RepositoryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a repository. The returned Operation will finish once the repository has been created. Its response will be the created Repository.

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input['RepositoryFormat']] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maven_config: Optional[pulumi.Input[pulumi.InputType['MavenRepositoryConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["description"] = description
            __props__.__dict__["format"] = format
            __props__.__dict__["kms_key_name"] = kms_key_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["maven_config"] = maven_config
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["repository_id"] = repository_id
            __props__.__dict__["update_time"] = update_time
        super(Repository, __self__).__init__(
            'google-native:artifactregistry/v1:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RepositoryArgs.__new__(RepositoryArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["format"] = None
        __props__.__dict__["kms_key_name"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["maven_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["update_time"] = None
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the repository was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The user-provided description of the repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        """
        The format of packages that are stored in the repository.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> pulumi.Output[str]:
        """
        The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mavenConfig")
    def maven_config(self) -> pulumi.Output['outputs.MavenRepositoryConfigResponse']:
        """
        Maven repository config contains repository level configuration for the repositories of maven type.
        """
        return pulumi.get(self, "maven_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the repository was last updated.
        """
        return pulumi.get(self, "update_time")

