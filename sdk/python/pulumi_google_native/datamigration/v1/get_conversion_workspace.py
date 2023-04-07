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
    'GetConversionWorkspaceResult',
    'AwaitableGetConversionWorkspaceResult',
    'get_conversion_workspace',
    'get_conversion_workspace_output',
]

@pulumi.output_type
class GetConversionWorkspaceResult:
    def __init__(__self__, create_time=None, destination=None, display_name=None, global_settings=None, has_uncommitted_changes=None, latest_commit_id=None, latest_commit_time=None, name=None, source=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if destination and not isinstance(destination, dict):
            raise TypeError("Expected argument 'destination' to be a dict")
        pulumi.set(__self__, "destination", destination)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if global_settings and not isinstance(global_settings, dict):
            raise TypeError("Expected argument 'global_settings' to be a dict")
        pulumi.set(__self__, "global_settings", global_settings)
        if has_uncommitted_changes and not isinstance(has_uncommitted_changes, bool):
            raise TypeError("Expected argument 'has_uncommitted_changes' to be a bool")
        pulumi.set(__self__, "has_uncommitted_changes", has_uncommitted_changes)
        if latest_commit_id and not isinstance(latest_commit_id, str):
            raise TypeError("Expected argument 'latest_commit_id' to be a str")
        pulumi.set(__self__, "latest_commit_id", latest_commit_id)
        if latest_commit_time and not isinstance(latest_commit_time, str):
            raise TypeError("Expected argument 'latest_commit_time' to be a str")
        pulumi.set(__self__, "latest_commit_time", latest_commit_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the workspace resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.DatabaseEngineInfoResponse':
        """
        The destination engine details.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the workspace.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="globalSettings")
    def global_settings(self) -> Mapping[str, str]:
        """
        A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
        """
        return pulumi.get(self, "global_settings")

    @property
    @pulumi.getter(name="hasUncommittedChanges")
    def has_uncommitted_changes(self) -> bool:
        """
        Whether the workspace has uncommitted changes (changes which were made after the workspace was committed).
        """
        return pulumi.get(self, "has_uncommitted_changes")

    @property
    @pulumi.getter(name="latestCommitId")
    def latest_commit_id(self) -> str:
        """
        The latest commit ID.
        """
        return pulumi.get(self, "latest_commit_id")

    @property
    @pulumi.getter(name="latestCommitTime")
    def latest_commit_time(self) -> str:
        """
        The timestamp when the workspace was committed.
        """
        return pulumi.get(self, "latest_commit_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Full name of the workspace resource, in the form of: projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> 'outputs.DatabaseEngineInfoResponse':
        """
        The source engine details.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp when the workspace resource was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetConversionWorkspaceResult(GetConversionWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConversionWorkspaceResult(
            create_time=self.create_time,
            destination=self.destination,
            display_name=self.display_name,
            global_settings=self.global_settings,
            has_uncommitted_changes=self.has_uncommitted_changes,
            latest_commit_id=self.latest_commit_id,
            latest_commit_time=self.latest_commit_time,
            name=self.name,
            source=self.source,
            update_time=self.update_time)


def get_conversion_workspace(conversion_workspace_id: Optional[str] = None,
                             location: Optional[str] = None,
                             project: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConversionWorkspaceResult:
    """
    Gets details of a single conversion workspace.
    """
    __args__ = dict()
    __args__['conversionWorkspaceId'] = conversion_workspace_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datamigration/v1:getConversionWorkspace', __args__, opts=opts, typ=GetConversionWorkspaceResult).value

    return AwaitableGetConversionWorkspaceResult(
        create_time=__ret__.create_time,
        destination=__ret__.destination,
        display_name=__ret__.display_name,
        global_settings=__ret__.global_settings,
        has_uncommitted_changes=__ret__.has_uncommitted_changes,
        latest_commit_id=__ret__.latest_commit_id,
        latest_commit_time=__ret__.latest_commit_time,
        name=__ret__.name,
        source=__ret__.source,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_conversion_workspace)
def get_conversion_workspace_output(conversion_workspace_id: Optional[pulumi.Input[str]] = None,
                                    location: Optional[pulumi.Input[str]] = None,
                                    project: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConversionWorkspaceResult]:
    """
    Gets details of a single conversion workspace.
    """
    ...
