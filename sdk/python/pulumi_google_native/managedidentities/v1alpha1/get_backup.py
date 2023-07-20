# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetBackupResult',
    'AwaitableGetBackupResult',
    'get_backup',
    'get_backup_output',
]

@pulumi.output_type
class GetBackupResult:
    def __init__(__self__, create_time=None, description=None, labels=None, name=None, state=None, status_message=None, type=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the backups was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A short description of the backup.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the backup.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        Additional information about the current status of this backup, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Indicates whether it’s an on-demand backup or scheduled.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Last update time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetBackupResult(GetBackupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupResult(
            create_time=self.create_time,
            description=self.description,
            labels=self.labels,
            name=self.name,
            state=self.state,
            status_message=self.status_message,
            type=self.type,
            update_time=self.update_time)


def get_backup(backup_id: Optional[str] = None,
               domain_id: Optional[str] = None,
               project: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupResult:
    """
    Gets details of a single Backup.
    """
    __args__ = dict()
    __args__['backupId'] = backup_id
    __args__['domainId'] = domain_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:managedidentities/v1alpha1:getBackup', __args__, opts=opts, typ=GetBackupResult).value

    return AwaitableGetBackupResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        status_message=pulumi.get(__ret__, 'status_message'),
        type=pulumi.get(__ret__, 'type'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_backup)
def get_backup_output(backup_id: Optional[pulumi.Input[str]] = None,
                      domain_id: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupResult]:
    """
    Gets details of a single Backup.
    """
    ...
