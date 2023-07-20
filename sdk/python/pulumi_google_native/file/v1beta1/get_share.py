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
    'GetShareResult',
    'AwaitableGetShareResult',
    'get_share',
    'get_share_output',
]

@pulumi.output_type
class GetShareResult:
    def __init__(__self__, backup=None, capacity_gb=None, create_time=None, description=None, labels=None, mount_name=None, name=None, nfs_export_options=None, state=None):
        if backup and not isinstance(backup, str):
            raise TypeError("Expected argument 'backup' to be a str")
        pulumi.set(__self__, "backup", backup)
        if capacity_gb and not isinstance(capacity_gb, str):
            raise TypeError("Expected argument 'capacity_gb' to be a str")
        pulumi.set(__self__, "capacity_gb", capacity_gb)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if mount_name and not isinstance(mount_name, str):
            raise TypeError("Expected argument 'mount_name' to be a str")
        pulumi.set(__self__, "mount_name", mount_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nfs_export_options and not isinstance(nfs_export_options, list):
            raise TypeError("Expected argument 'nfs_export_options' to be a list")
        pulumi.set(__self__, "nfs_export_options", nfs_export_options)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def backup(self) -> str:
        """
        Immutable. Full name of the Cloud Filestore Backup resource that this Share is restored from, in the format of projects/{project_id}/locations/{location_id}/backups/{backup_id}. Empty, if the Share is created from scratch and not restored from a backup.
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter(name="capacityGb")
    def capacity_gb(self) -> str:
        """
        File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes. Must be greater than 0.
        """
        return pulumi.get(self, "capacity_gb")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the share was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the share with 2048 characters or less. Requests with longer descriptions will be rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mountName")
    def mount_name(self) -> str:
        """
        The mount name of the share. Must be 63 characters or less and consist of uppercase or lowercase letters, numbers, and underscores.
        """
        return pulumi.get(self, "mount_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the share, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}/shares/{share_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nfsExportOptions")
    def nfs_export_options(self) -> Sequence['outputs.NfsExportOptionsResponse']:
        """
        Nfs Export Options. There is a limit of 10 export options per file share.
        """
        return pulumi.get(self, "nfs_export_options")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The share state.
        """
        return pulumi.get(self, "state")


class AwaitableGetShareResult(GetShareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetShareResult(
            backup=self.backup,
            capacity_gb=self.capacity_gb,
            create_time=self.create_time,
            description=self.description,
            labels=self.labels,
            mount_name=self.mount_name,
            name=self.name,
            nfs_export_options=self.nfs_export_options,
            state=self.state)


def get_share(instance_id: Optional[str] = None,
              location: Optional[str] = None,
              project: Optional[str] = None,
              share_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetShareResult:
    """
    Gets the details of a specific share.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['shareId'] = share_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:file/v1beta1:getShare', __args__, opts=opts, typ=GetShareResult).value

    return AwaitableGetShareResult(
        backup=pulumi.get(__ret__, 'backup'),
        capacity_gb=pulumi.get(__ret__, 'capacity_gb'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        labels=pulumi.get(__ret__, 'labels'),
        mount_name=pulumi.get(__ret__, 'mount_name'),
        name=pulumi.get(__ret__, 'name'),
        nfs_export_options=pulumi.get(__ret__, 'nfs_export_options'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_share)
def get_share_output(instance_id: Optional[pulumi.Input[str]] = None,
                     location: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     share_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetShareResult]:
    """
    Gets the details of a specific share.
    """
    ...
