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
    'GetSavedQueryResult',
    'AwaitableGetSavedQueryResult',
    'get_saved_query',
    'get_saved_query_output',
]

@pulumi.output_type
class GetSavedQueryResult:
    def __init__(__self__, content=None, create_time=None, creator=None, description=None, labels=None, last_update_time=None, last_updater=None, name=None):
        if content and not isinstance(content, dict):
            raise TypeError("Expected argument 'content' to be a dict")
        pulumi.set(__self__, "content", content)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if creator and not isinstance(creator, str):
            raise TypeError("Expected argument 'creator' to be a str")
        pulumi.set(__self__, "creator", creator)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_update_time and not isinstance(last_update_time, str):
            raise TypeError("Expected argument 'last_update_time' to be a str")
        pulumi.set(__self__, "last_update_time", last_update_time)
        if last_updater and not isinstance(last_updater, str):
            raise TypeError("Expected argument 'last_updater' to be a str")
        pulumi.set(__self__, "last_updater", last_updater)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> 'outputs.QueryContentResponse':
        """
        The query content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of this saved query.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def creator(self) -> str:
        """
        The account's email address who has created this saved query.
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of this saved query. This value should be fewer than 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> str:
        """
        The last update time of this saved query.
        """
        return pulumi.get(self, "last_update_time")

    @property
    @pulumi.getter(name="lastUpdater")
    def last_updater(self) -> str:
        """
        The account's email address who has updated this saved query most recently.
        """
        return pulumi.get(self, "last_updater")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
        """
        return pulumi.get(self, "name")


class AwaitableGetSavedQueryResult(GetSavedQueryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSavedQueryResult(
            content=self.content,
            create_time=self.create_time,
            creator=self.creator,
            description=self.description,
            labels=self.labels,
            last_update_time=self.last_update_time,
            last_updater=self.last_updater,
            name=self.name)


def get_saved_query(saved_query_id: Optional[str] = None,
                    v1_id: Optional[str] = None,
                    v1_id1: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSavedQueryResult:
    """
    Gets details about a saved query.
    """
    __args__ = dict()
    __args__['savedQueryId'] = saved_query_id
    __args__['v1Id'] = v1_id
    __args__['v1Id1'] = v1_id1
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudasset/v1:getSavedQuery', __args__, opts=opts, typ=GetSavedQueryResult).value

    return AwaitableGetSavedQueryResult(
        content=pulumi.get(__ret__, 'content'),
        create_time=pulumi.get(__ret__, 'create_time'),
        creator=pulumi.get(__ret__, 'creator'),
        description=pulumi.get(__ret__, 'description'),
        labels=pulumi.get(__ret__, 'labels'),
        last_update_time=pulumi.get(__ret__, 'last_update_time'),
        last_updater=pulumi.get(__ret__, 'last_updater'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_saved_query)
def get_saved_query_output(saved_query_id: Optional[pulumi.Input[str]] = None,
                           v1_id: Optional[pulumi.Input[str]] = None,
                           v1_id1: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSavedQueryResult]:
    """
    Gets details about a saved query.
    """
    ...
