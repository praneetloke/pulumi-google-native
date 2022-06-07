# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['BackupArgs', 'Backup']

@pulumi.input_type
class BackupArgs:
    def __init__(__self__, *,
                 backup_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 source_file_share: Optional[pulumi.Input[str]] = None,
                 source_instance: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Backup resource.
        :param pulumi.Input[str] backup_id: Required. The ID to use for the backup. The ID must be unique within the specified project and location. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
        :param pulumi.Input[str] description: A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        :param pulumi.Input[str] kms_key_name: Immutable. KMS key name used for data encryption.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] source_file_share: Name of the file share in the source Cloud Filestore instance that the backup is created from.
        :param pulumi.Input[str] source_instance: The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
        """
        pulumi.set(__self__, "backup_id", backup_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if source_file_share is not None:
            pulumi.set(__self__, "source_file_share", source_file_share)
        if source_instance is not None:
            pulumi.set(__self__, "source_instance", source_instance)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Input[str]:
        """
        Required. The ID to use for the backup. The ID must be unique within the specified project and location. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
        """
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. KMS key name used for data encryption.
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user provided metadata.
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
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sourceFileShare")
    def source_file_share(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the file share in the source Cloud Filestore instance that the backup is created from.
        """
        return pulumi.get(self, "source_file_share")

    @source_file_share.setter
    def source_file_share(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_file_share", value)

    @property
    @pulumi.getter(name="sourceInstance")
    def source_instance(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
        """
        return pulumi.get(self, "source_instance")

    @source_instance.setter
    def source_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_instance", value)


class Backup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 source_file_share: Optional[pulumi.Input[str]] = None,
                 source_instance: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a backup.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_id: Required. The ID to use for the backup. The ID must be unique within the specified project and location. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
        :param pulumi.Input[str] description: A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        :param pulumi.Input[str] kms_key_name: Immutable. KMS key name used for data encryption.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] source_file_share: Name of the file share in the source Cloud Filestore instance that the backup is created from.
        :param pulumi.Input[str] source_instance: The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backup.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param BackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 source_file_share: Optional[pulumi.Input[str]] = None,
                 source_instance: Optional[pulumi.Input[str]] = None,
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
            __props__ = BackupArgs.__new__(BackupArgs)

            if backup_id is None and not opts.urn:
                raise TypeError("Missing required property 'backup_id'")
            __props__.__dict__["backup_id"] = backup_id
            __props__.__dict__["description"] = description
            __props__.__dict__["kms_key_name"] = kms_key_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["source_file_share"] = source_file_share
            __props__.__dict__["source_instance"] = source_instance
            __props__.__dict__["capacity_gb"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["download_bytes"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["satisfies_pzs"] = None
            __props__.__dict__["source_instance_tier"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["storage_bytes"] = None
        super(Backup, __self__).__init__(
            'google-native:file/v1beta1:Backup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Backup':
        """
        Get an existing Backup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BackupArgs.__new__(BackupArgs)

        __props__.__dict__["capacity_gb"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["download_bytes"] = None
        __props__.__dict__["kms_key_name"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["satisfies_pzs"] = None
        __props__.__dict__["source_file_share"] = None
        __props__.__dict__["source_instance"] = None
        __props__.__dict__["source_instance_tier"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["storage_bytes"] = None
        return Backup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="capacityGb")
    def capacity_gb(self) -> pulumi.Output[str]:
        """
        Capacity of the source file share when the backup was created.
        """
        return pulumi.get(self, "capacity_gb")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the backup was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadBytes")
    def download_bytes(self) -> pulumi.Output[str]:
        """
        Amount of bytes that will be downloaded if the backup is restored
        """
        return pulumi.get(self, "download_bytes")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> pulumi.Output[str]:
        """
        Immutable. KMS key name used for data encryption.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the backup, in the format `projects/{project_id}/locations/{location_id}/backups/{backup_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> pulumi.Output[bool]:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="sourceFileShare")
    def source_file_share(self) -> pulumi.Output[str]:
        """
        Name of the file share in the source Cloud Filestore instance that the backup is created from.
        """
        return pulumi.get(self, "source_file_share")

    @property
    @pulumi.getter(name="sourceInstance")
    def source_instance(self) -> pulumi.Output[str]:
        """
        The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
        """
        return pulumi.get(self, "source_instance")

    @property
    @pulumi.getter(name="sourceInstanceTier")
    def source_instance_tier(self) -> pulumi.Output[str]:
        """
        The service tier of the source Cloud Filestore instance that this backup is created from.
        """
        return pulumi.get(self, "source_instance_tier")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The backup state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageBytes")
    def storage_bytes(self) -> pulumi.Output[str]:
        """
        The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
        """
        return pulumi.get(self, "storage_bytes")

