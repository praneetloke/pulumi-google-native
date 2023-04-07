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
    'GetImageResult',
    'AwaitableGetImageResult',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class GetImageResult:
    def __init__(__self__, architecture=None, archive_size_bytes=None, creation_timestamp=None, deprecated=None, description=None, disk_size_gb=None, family=None, guest_os_features=None, image_encryption_key=None, kind=None, label_fingerprint=None, labels=None, license_codes=None, licenses=None, locked=None, name=None, raw_disk=None, rollout_override=None, satisfies_pzs=None, self_link=None, shielded_instance_initial_state=None, source_disk=None, source_disk_encryption_key=None, source_disk_id=None, source_image=None, source_image_encryption_key=None, source_image_id=None, source_snapshot=None, source_snapshot_encryption_key=None, source_snapshot_id=None, source_type=None, status=None, storage_locations=None, user_licenses=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if archive_size_bytes and not isinstance(archive_size_bytes, str):
            raise TypeError("Expected argument 'archive_size_bytes' to be a str")
        pulumi.set(__self__, "archive_size_bytes", archive_size_bytes)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if deprecated and not isinstance(deprecated, dict):
            raise TypeError("Expected argument 'deprecated' to be a dict")
        pulumi.set(__self__, "deprecated", deprecated)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size_gb and not isinstance(disk_size_gb, str):
            raise TypeError("Expected argument 'disk_size_gb' to be a str")
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if family and not isinstance(family, str):
            raise TypeError("Expected argument 'family' to be a str")
        pulumi.set(__self__, "family", family)
        if guest_os_features and not isinstance(guest_os_features, list):
            raise TypeError("Expected argument 'guest_os_features' to be a list")
        pulumi.set(__self__, "guest_os_features", guest_os_features)
        if image_encryption_key and not isinstance(image_encryption_key, dict):
            raise TypeError("Expected argument 'image_encryption_key' to be a dict")
        pulumi.set(__self__, "image_encryption_key", image_encryption_key)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if license_codes and not isinstance(license_codes, list):
            raise TypeError("Expected argument 'license_codes' to be a list")
        pulumi.set(__self__, "license_codes", license_codes)
        if licenses and not isinstance(licenses, list):
            raise TypeError("Expected argument 'licenses' to be a list")
        pulumi.set(__self__, "licenses", licenses)
        if locked and not isinstance(locked, bool):
            raise TypeError("Expected argument 'locked' to be a bool")
        pulumi.set(__self__, "locked", locked)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if raw_disk and not isinstance(raw_disk, dict):
            raise TypeError("Expected argument 'raw_disk' to be a dict")
        pulumi.set(__self__, "raw_disk", raw_disk)
        if rollout_override and not isinstance(rollout_override, dict):
            raise TypeError("Expected argument 'rollout_override' to be a dict")
        pulumi.set(__self__, "rollout_override", rollout_override)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if shielded_instance_initial_state and not isinstance(shielded_instance_initial_state, dict):
            raise TypeError("Expected argument 'shielded_instance_initial_state' to be a dict")
        pulumi.set(__self__, "shielded_instance_initial_state", shielded_instance_initial_state)
        if source_disk and not isinstance(source_disk, str):
            raise TypeError("Expected argument 'source_disk' to be a str")
        pulumi.set(__self__, "source_disk", source_disk)
        if source_disk_encryption_key and not isinstance(source_disk_encryption_key, dict):
            raise TypeError("Expected argument 'source_disk_encryption_key' to be a dict")
        pulumi.set(__self__, "source_disk_encryption_key", source_disk_encryption_key)
        if source_disk_id and not isinstance(source_disk_id, str):
            raise TypeError("Expected argument 'source_disk_id' to be a str")
        pulumi.set(__self__, "source_disk_id", source_disk_id)
        if source_image and not isinstance(source_image, str):
            raise TypeError("Expected argument 'source_image' to be a str")
        pulumi.set(__self__, "source_image", source_image)
        if source_image_encryption_key and not isinstance(source_image_encryption_key, dict):
            raise TypeError("Expected argument 'source_image_encryption_key' to be a dict")
        pulumi.set(__self__, "source_image_encryption_key", source_image_encryption_key)
        if source_image_id and not isinstance(source_image_id, str):
            raise TypeError("Expected argument 'source_image_id' to be a str")
        pulumi.set(__self__, "source_image_id", source_image_id)
        if source_snapshot and not isinstance(source_snapshot, str):
            raise TypeError("Expected argument 'source_snapshot' to be a str")
        pulumi.set(__self__, "source_snapshot", source_snapshot)
        if source_snapshot_encryption_key and not isinstance(source_snapshot_encryption_key, dict):
            raise TypeError("Expected argument 'source_snapshot_encryption_key' to be a dict")
        pulumi.set(__self__, "source_snapshot_encryption_key", source_snapshot_encryption_key)
        if source_snapshot_id and not isinstance(source_snapshot_id, str):
            raise TypeError("Expected argument 'source_snapshot_id' to be a str")
        pulumi.set(__self__, "source_snapshot_id", source_snapshot_id)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_locations and not isinstance(storage_locations, list):
            raise TypeError("Expected argument 'storage_locations' to be a list")
        pulumi.set(__self__, "storage_locations", storage_locations)
        if user_licenses and not isinstance(user_licenses, list):
            raise TypeError("Expected argument 'user_licenses' to be a list")
        pulumi.set(__self__, "user_licenses", user_licenses)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        """
        The architecture of the image. Valid values are ARM64 or X86_64.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="archiveSizeBytes")
    def archive_size_bytes(self) -> str:
        """
        Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
        """
        return pulumi.get(self, "archive_size_bytes")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def deprecated(self) -> 'outputs.DeprecationStatusResponse':
        """
        The deprecation status associated with this image.
        """
        return pulumi.get(self, "deprecated")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> str:
        """
        Size of the image when restored onto a persistent disk (in GB).
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter
    def family(self) -> str:
        """
        The name of the image family to which this image belongs. The image family name can be from a publicly managed image family provided by Compute Engine, or from a custom image family you create. For example, centos-stream-9 is a publicly available image family. For more information, see Image family best practices. When creating disks, you can specify an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter(name="guestOsFeatures")
    def guest_os_features(self) -> Sequence['outputs.GuestOsFeatureResponse']:
        """
        A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
        """
        return pulumi.get(self, "guest_os_features")

    @property
    @pulumi.getter(name="imageEncryptionKey")
    def image_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
        """
        return pulumi.get(self, "image_encryption_key")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#image for images.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels to apply to this image. These can be later modified by the setLabels method.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="licenseCodes")
    def license_codes(self) -> Sequence[str]:
        """
        Integer license codes indicating which licenses are attached to this image.
        """
        return pulumi.get(self, "license_codes")

    @property
    @pulumi.getter
    def licenses(self) -> Sequence[str]:
        """
        Any applicable license URI.
        """
        return pulumi.get(self, "licenses")

    @property
    @pulumi.getter
    def locked(self) -> bool:
        """
        A flag for marketplace VM disk created from the image, which is designed for marketplace VM disk to prevent the proprietary data on the disk from being accessed unwantedly. The flag will be inherited by the disk created from the image. The disk with locked flag set to true will be prohibited from performing the operations below: - R/W or R/O disk attach - Disk detach, if disk is created via create-on-create - Create images - Create snapshots - Create disk clone (create disk from the current disk) The image with the locked field set to true will be prohibited from performing the operations below: - Create images from the current image - Update the locked field for the current image The instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Secondary disk attach - Create instant snapshot - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true 
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rawDisk")
    def raw_disk(self) -> 'outputs.ImageRawDiskResponse':
        """
        The parameters of the raw disk image.
        """
        return pulumi.get(self, "raw_disk")

    @property
    @pulumi.getter(name="rolloutOverride")
    def rollout_override(self) -> 'outputs.RolloutPolicyResponse':
        """
        A rollout policy to apply to this image. When specified, the rollout policy overrides per-zone references to the image via the associated image family. The rollout policy restricts the zones where this image is accessible when using a zonal image family reference. When the rollout policy does not include the user specified zone, or if the zone is rolled out, this image is accessible. The rollout policy for this image is read-only, except for allowlisted users. This field might not be configured. To view the latest non-deprecated image in a specific zone, use the imageFamilyViews.get method.
        """
        return pulumi.get(self, "rollout_override")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="shieldedInstanceInitialState")
    def shielded_instance_initial_state(self) -> 'outputs.InitialStateConfigResponse':
        """
        Set the secure boot keys of shielded instance.
        """
        return pulumi.get(self, "shielded_instance_initial_state")

    @property
    @pulumi.getter(name="sourceDisk")
    def source_disk(self) -> str:
        """
        URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        """
        return pulumi.get(self, "source_disk")

    @property
    @pulumi.getter(name="sourceDiskEncryptionKey")
    def source_disk_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        """
        return pulumi.get(self, "source_disk_encryption_key")

    @property
    @pulumi.getter(name="sourceDiskId")
    def source_disk_id(self) -> str:
        """
        The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
        """
        return pulumi.get(self, "source_disk_id")

    @property
    @pulumi.getter(name="sourceImage")
    def source_image(self) -> str:
        """
        URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        """
        return pulumi.get(self, "source_image")

    @property
    @pulumi.getter(name="sourceImageEncryptionKey")
    def source_image_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
        """
        return pulumi.get(self, "source_image_encryption_key")

    @property
    @pulumi.getter(name="sourceImageId")
    def source_image_id(self) -> str:
        """
        The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
        """
        return pulumi.get(self, "source_image_id")

    @property
    @pulumi.getter(name="sourceSnapshot")
    def source_snapshot(self) -> str:
        """
        URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        """
        return pulumi.get(self, "source_snapshot")

    @property
    @pulumi.getter(name="sourceSnapshotEncryptionKey")
    def source_snapshot_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
        """
        return pulumi.get(self, "source_snapshot_encryption_key")

    @property
    @pulumi.getter(name="sourceSnapshotId")
    def source_snapshot_id(self) -> str:
        """
        The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
        """
        return pulumi.get(self, "source_snapshot_id")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        The type of the image used to create this disk. The default and only valid value is RAW.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageLocations")
    def storage_locations(self) -> Sequence[str]:
        """
        Cloud Storage bucket storage location of the image (regional or multi-regional).
        """
        return pulumi.get(self, "storage_locations")

    @property
    @pulumi.getter(name="userLicenses")
    def user_licenses(self) -> Sequence[str]:
        """
        A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
        """
        return pulumi.get(self, "user_licenses")


class AwaitableGetImageResult(GetImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageResult(
            architecture=self.architecture,
            archive_size_bytes=self.archive_size_bytes,
            creation_timestamp=self.creation_timestamp,
            deprecated=self.deprecated,
            description=self.description,
            disk_size_gb=self.disk_size_gb,
            family=self.family,
            guest_os_features=self.guest_os_features,
            image_encryption_key=self.image_encryption_key,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            license_codes=self.license_codes,
            licenses=self.licenses,
            locked=self.locked,
            name=self.name,
            raw_disk=self.raw_disk,
            rollout_override=self.rollout_override,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            shielded_instance_initial_state=self.shielded_instance_initial_state,
            source_disk=self.source_disk,
            source_disk_encryption_key=self.source_disk_encryption_key,
            source_disk_id=self.source_disk_id,
            source_image=self.source_image,
            source_image_encryption_key=self.source_image_encryption_key,
            source_image_id=self.source_image_id,
            source_snapshot=self.source_snapshot,
            source_snapshot_encryption_key=self.source_snapshot_encryption_key,
            source_snapshot_id=self.source_snapshot_id,
            source_type=self.source_type,
            status=self.status,
            storage_locations=self.storage_locations,
            user_licenses=self.user_licenses)


def get_image(image: Optional[str] = None,
              project: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageResult:
    """
    Returns the specified image.
    """
    __args__ = dict()
    __args__['image'] = image
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getImage', __args__, opts=opts, typ=GetImageResult).value

    return AwaitableGetImageResult(
        architecture=__ret__.architecture,
        archive_size_bytes=__ret__.archive_size_bytes,
        creation_timestamp=__ret__.creation_timestamp,
        deprecated=__ret__.deprecated,
        description=__ret__.description,
        disk_size_gb=__ret__.disk_size_gb,
        family=__ret__.family,
        guest_os_features=__ret__.guest_os_features,
        image_encryption_key=__ret__.image_encryption_key,
        kind=__ret__.kind,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        license_codes=__ret__.license_codes,
        licenses=__ret__.licenses,
        locked=__ret__.locked,
        name=__ret__.name,
        raw_disk=__ret__.raw_disk,
        rollout_override=__ret__.rollout_override,
        satisfies_pzs=__ret__.satisfies_pzs,
        self_link=__ret__.self_link,
        shielded_instance_initial_state=__ret__.shielded_instance_initial_state,
        source_disk=__ret__.source_disk,
        source_disk_encryption_key=__ret__.source_disk_encryption_key,
        source_disk_id=__ret__.source_disk_id,
        source_image=__ret__.source_image,
        source_image_encryption_key=__ret__.source_image_encryption_key,
        source_image_id=__ret__.source_image_id,
        source_snapshot=__ret__.source_snapshot,
        source_snapshot_encryption_key=__ret__.source_snapshot_encryption_key,
        source_snapshot_id=__ret__.source_snapshot_id,
        source_type=__ret__.source_type,
        status=__ret__.status,
        storage_locations=__ret__.storage_locations,
        user_licenses=__ret__.user_licenses)


@_utilities.lift_output_func(get_image)
def get_image_output(image: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageResult]:
    """
    Returns the specified image.
    """
    ...
