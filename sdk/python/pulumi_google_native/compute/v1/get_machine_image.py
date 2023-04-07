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
    'GetMachineImageResult',
    'AwaitableGetMachineImageResult',
    'get_machine_image',
    'get_machine_image_output',
]

@pulumi.output_type
class GetMachineImageResult:
    def __init__(__self__, creation_timestamp=None, description=None, guest_flush=None, instance_properties=None, kind=None, machine_image_encryption_key=None, name=None, satisfies_pzs=None, saved_disks=None, self_link=None, source_disk_encryption_keys=None, source_instance=None, source_instance_properties=None, status=None, storage_locations=None, total_storage_bytes=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if guest_flush and not isinstance(guest_flush, bool):
            raise TypeError("Expected argument 'guest_flush' to be a bool")
        pulumi.set(__self__, "guest_flush", guest_flush)
        if instance_properties and not isinstance(instance_properties, dict):
            raise TypeError("Expected argument 'instance_properties' to be a dict")
        pulumi.set(__self__, "instance_properties", instance_properties)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if machine_image_encryption_key and not isinstance(machine_image_encryption_key, dict):
            raise TypeError("Expected argument 'machine_image_encryption_key' to be a dict")
        pulumi.set(__self__, "machine_image_encryption_key", machine_image_encryption_key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if saved_disks and not isinstance(saved_disks, list):
            raise TypeError("Expected argument 'saved_disks' to be a list")
        pulumi.set(__self__, "saved_disks", saved_disks)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if source_disk_encryption_keys and not isinstance(source_disk_encryption_keys, list):
            raise TypeError("Expected argument 'source_disk_encryption_keys' to be a list")
        pulumi.set(__self__, "source_disk_encryption_keys", source_disk_encryption_keys)
        if source_instance and not isinstance(source_instance, str):
            raise TypeError("Expected argument 'source_instance' to be a str")
        pulumi.set(__self__, "source_instance", source_instance)
        if source_instance_properties and not isinstance(source_instance_properties, dict):
            raise TypeError("Expected argument 'source_instance_properties' to be a dict")
        pulumi.set(__self__, "source_instance_properties", source_instance_properties)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_locations and not isinstance(storage_locations, list):
            raise TypeError("Expected argument 'storage_locations' to be a list")
        pulumi.set(__self__, "storage_locations", storage_locations)
        if total_storage_bytes and not isinstance(total_storage_bytes, str):
            raise TypeError("Expected argument 'total_storage_bytes' to be a str")
        pulumi.set(__self__, "total_storage_bytes", total_storage_bytes)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        The creation timestamp for this machine image in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="guestFlush")
    def guest_flush(self) -> bool:
        """
        [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
        """
        return pulumi.get(self, "guest_flush")

    @property
    @pulumi.getter(name="instanceProperties")
    def instance_properties(self) -> 'outputs.InstancePropertiesResponse':
        """
        Properties of source instance
        """
        return pulumi.get(self, "instance_properties")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The resource type, which is always compute#machineImage for machine image.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="machineImageEncryptionKey")
    def machine_image_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
        """
        return pulumi.get(self, "machine_image_encryption_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="savedDisks")
    def saved_disks(self) -> Sequence['outputs.SavedDiskResponse']:
        """
        An array of Machine Image specific properties for disks attached to the source instance
        """
        return pulumi.get(self, "saved_disks")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URL for this machine image. The server defines this URL.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sourceDiskEncryptionKeys")
    def source_disk_encryption_keys(self) -> Sequence['outputs.SourceDiskEncryptionKeyResponse']:
        """
        [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
        """
        return pulumi.get(self, "source_disk_encryption_keys")

    @property
    @pulumi.getter(name="sourceInstance")
    def source_instance(self) -> str:
        """
        The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
        """
        return pulumi.get(self, "source_instance")

    @property
    @pulumi.getter(name="sourceInstanceProperties")
    def source_instance_properties(self) -> 'outputs.SourceInstancePropertiesResponse':
        """
        DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
        """
        return pulumi.get(self, "source_instance_properties")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageLocations")
    def storage_locations(self) -> Sequence[str]:
        """
        The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        """
        return pulumi.get(self, "storage_locations")

    @property
    @pulumi.getter(name="totalStorageBytes")
    def total_storage_bytes(self) -> str:
        """
        Total size of the storage used by the machine image.
        """
        return pulumi.get(self, "total_storage_bytes")


class AwaitableGetMachineImageResult(GetMachineImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMachineImageResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            guest_flush=self.guest_flush,
            instance_properties=self.instance_properties,
            kind=self.kind,
            machine_image_encryption_key=self.machine_image_encryption_key,
            name=self.name,
            satisfies_pzs=self.satisfies_pzs,
            saved_disks=self.saved_disks,
            self_link=self.self_link,
            source_disk_encryption_keys=self.source_disk_encryption_keys,
            source_instance=self.source_instance,
            source_instance_properties=self.source_instance_properties,
            status=self.status,
            storage_locations=self.storage_locations,
            total_storage_bytes=self.total_storage_bytes)


def get_machine_image(machine_image: Optional[str] = None,
                      project: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMachineImageResult:
    """
    Returns the specified machine image.
    """
    __args__ = dict()
    __args__['machineImage'] = machine_image
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getMachineImage', __args__, opts=opts, typ=GetMachineImageResult).value

    return AwaitableGetMachineImageResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        guest_flush=__ret__.guest_flush,
        instance_properties=__ret__.instance_properties,
        kind=__ret__.kind,
        machine_image_encryption_key=__ret__.machine_image_encryption_key,
        name=__ret__.name,
        satisfies_pzs=__ret__.satisfies_pzs,
        saved_disks=__ret__.saved_disks,
        self_link=__ret__.self_link,
        source_disk_encryption_keys=__ret__.source_disk_encryption_keys,
        source_instance=__ret__.source_instance,
        source_instance_properties=__ret__.source_instance_properties,
        status=__ret__.status,
        storage_locations=__ret__.storage_locations,
        total_storage_bytes=__ret__.total_storage_bytes)


@_utilities.lift_output_func(get_machine_image)
def get_machine_image_output(machine_image: Optional[pulumi.Input[str]] = None,
                             project: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMachineImageResult]:
    """
    Returns the specified machine image.
    """
    ...
