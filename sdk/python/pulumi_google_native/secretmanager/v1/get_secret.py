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
    'GetSecretResult',
    'AwaitableGetSecretResult',
    'get_secret',
    'get_secret_output',
]

@pulumi.output_type
class GetSecretResult:
    def __init__(__self__, create_time=None, etag=None, expire_time=None, labels=None, name=None, replication=None, rotation=None, topics=None, ttl=None, version_aliases=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if replication and not isinstance(replication, dict):
            raise TypeError("Expected argument 'replication' to be a dict")
        pulumi.set(__self__, "replication", replication)
        if rotation and not isinstance(rotation, dict):
            raise TypeError("Expected argument 'rotation' to be a dict")
        pulumi.set(__self__, "rotation", rotation)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)
        if ttl and not isinstance(ttl, str):
            raise TypeError("Expected argument 'ttl' to be a str")
        pulumi.set(__self__, "ttl", ttl)
        if version_aliases and not isinstance(version_aliases, dict):
            raise TypeError("Expected argument 'version_aliases' to be a dict")
        pulumi.set(__self__, "version_aliases", version_aliases)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time at which the Secret was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Optional. Etag of the currently stored Secret.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        Optional. Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `\p{Ll}\p{Lo}{0,62}` Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}` No more than 64 labels can be assigned to a given resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the Secret in the format `projects/*/secrets/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def replication(self) -> 'outputs.ReplicationResponse':
        """
        Immutable. The replication policy of the secret data attached to the Secret. The replication policy cannot be changed after the Secret has been created.
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter
    def rotation(self) -> 'outputs.RotationResponse':
        """
        Optional. Rotation policy attached to the Secret. May be excluded if there is no rotation policy.
        """
        return pulumi.get(self, "rotation")

    @property
    @pulumi.getter
    def topics(self) -> Sequence['outputs.TopicResponse']:
        """
        Optional. A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter
    def ttl(self) -> str:
        """
        Input only. The TTL for the Secret.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="versionAliases")
    def version_aliases(self) -> Mapping[str, str]:
        """
        Optional. Mapping from version alias to version name. A version alias is a string with a maximum length of 63 characters and can contain uppercase and lowercase letters, numerals, and the hyphen (`-`) and underscore ('_') characters. An alias string must start with a letter and cannot be the string 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret. Version-Alias pairs will be viewable via GetSecret and modifiable via UpdateSecret. At launch Access by Allias will only be supported on GetSecretVersion and AccessSecretVersion.
        """
        return pulumi.get(self, "version_aliases")


class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            create_time=self.create_time,
            etag=self.etag,
            expire_time=self.expire_time,
            labels=self.labels,
            name=self.name,
            replication=self.replication,
            rotation=self.rotation,
            topics=self.topics,
            ttl=self.ttl,
            version_aliases=self.version_aliases)


def get_secret(project: Optional[str] = None,
               secret_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
    """
    Gets metadata for a given Secret.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['secretId'] = secret_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:secretmanager/v1:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        create_time=__ret__.create_time,
        etag=__ret__.etag,
        expire_time=__ret__.expire_time,
        labels=__ret__.labels,
        name=__ret__.name,
        replication=__ret__.replication,
        rotation=__ret__.rotation,
        topics=__ret__.topics,
        ttl=__ret__.ttl,
        version_aliases=__ret__.version_aliases)


@_utilities.lift_output_func(get_secret)
def get_secret_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                      secret_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretResult]:
    """
    Gets metadata for a given Secret.
    """
    ...
