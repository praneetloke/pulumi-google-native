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
    'GetDiskResult',
    'AwaitableGetDiskResult',
    'get_disk',
    'get_disk_output',
]

@pulumi.output_type
class GetDiskResult:
    def __init__(__self__, architecture=None, async_primary_disk=None, async_secondary_disks=None, creation_timestamp=None, description=None, disk_encryption_key=None, guest_os_features=None, kind=None, label_fingerprint=None, labels=None, last_attach_timestamp=None, last_detach_timestamp=None, license_codes=None, licenses=None, location_hint=None, name=None, options=None, params=None, physical_block_size_bytes=None, provisioned_iops=None, provisioned_throughput=None, region=None, replica_zones=None, resource_policies=None, resource_status=None, satisfies_pzs=None, self_link=None, size_gb=None, source_consistency_group_policy=None, source_consistency_group_policy_id=None, source_disk=None, source_disk_id=None, source_image=None, source_image_encryption_key=None, source_image_id=None, source_snapshot=None, source_snapshot_encryption_key=None, source_snapshot_id=None, source_storage_object=None, status=None, type=None, users=None, zone=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if async_primary_disk and not isinstance(async_primary_disk, dict):
            raise TypeError("Expected argument 'async_primary_disk' to be a dict")
        pulumi.set(__self__, "async_primary_disk", async_primary_disk)
        if async_secondary_disks and not isinstance(async_secondary_disks, dict):
            raise TypeError("Expected argument 'async_secondary_disks' to be a dict")
        pulumi.set(__self__, "async_secondary_disks", async_secondary_disks)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_encryption_key and not isinstance(disk_encryption_key, dict):
            raise TypeError("Expected argument 'disk_encryption_key' to be a dict")
        pulumi.set(__self__, "disk_encryption_key", disk_encryption_key)
        if guest_os_features and not isinstance(guest_os_features, list):
            raise TypeError("Expected argument 'guest_os_features' to be a list")
        pulumi.set(__self__, "guest_os_features", guest_os_features)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_attach_timestamp and not isinstance(last_attach_timestamp, str):
            raise TypeError("Expected argument 'last_attach_timestamp' to be a str")
        pulumi.set(__self__, "last_attach_timestamp", last_attach_timestamp)
        if last_detach_timestamp and not isinstance(last_detach_timestamp, str):
            raise TypeError("Expected argument 'last_detach_timestamp' to be a str")
        pulumi.set(__self__, "last_detach_timestamp", last_detach_timestamp)
        if license_codes and not isinstance(license_codes, list):
            raise TypeError("Expected argument 'license_codes' to be a list")
        pulumi.set(__self__, "license_codes", license_codes)
        if licenses and not isinstance(licenses, list):
            raise TypeError("Expected argument 'licenses' to be a list")
        pulumi.set(__self__, "licenses", licenses)
        if location_hint and not isinstance(location_hint, str):
            raise TypeError("Expected argument 'location_hint' to be a str")
        pulumi.set(__self__, "location_hint", location_hint)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, str):
            raise TypeError("Expected argument 'options' to be a str")
        pulumi.set(__self__, "options", options)
        if params and not isinstance(params, dict):
            raise TypeError("Expected argument 'params' to be a dict")
        pulumi.set(__self__, "params", params)
        if physical_block_size_bytes and not isinstance(physical_block_size_bytes, str):
            raise TypeError("Expected argument 'physical_block_size_bytes' to be a str")
        pulumi.set(__self__, "physical_block_size_bytes", physical_block_size_bytes)
        if provisioned_iops and not isinstance(provisioned_iops, str):
            raise TypeError("Expected argument 'provisioned_iops' to be a str")
        pulumi.set(__self__, "provisioned_iops", provisioned_iops)
        if provisioned_throughput and not isinstance(provisioned_throughput, str):
            raise TypeError("Expected argument 'provisioned_throughput' to be a str")
        pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if replica_zones and not isinstance(replica_zones, list):
            raise TypeError("Expected argument 'replica_zones' to be a list")
        pulumi.set(__self__, "replica_zones", replica_zones)
        if resource_policies and not isinstance(resource_policies, list):
            raise TypeError("Expected argument 'resource_policies' to be a list")
        pulumi.set(__self__, "resource_policies", resource_policies)
        if resource_status and not isinstance(resource_status, dict):
            raise TypeError("Expected argument 'resource_status' to be a dict")
        pulumi.set(__self__, "resource_status", resource_status)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if size_gb and not isinstance(size_gb, str):
            raise TypeError("Expected argument 'size_gb' to be a str")
        pulumi.set(__self__, "size_gb", size_gb)
        if source_consistency_group_policy and not isinstance(source_consistency_group_policy, str):
            raise TypeError("Expected argument 'source_consistency_group_policy' to be a str")
        pulumi.set(__self__, "source_consistency_group_policy", source_consistency_group_policy)
        if source_consistency_group_policy_id and not isinstance(source_consistency_group_policy_id, str):
            raise TypeError("Expected argument 'source_consistency_group_policy_id' to be a str")
        pulumi.set(__self__, "source_consistency_group_policy_id", source_consistency_group_policy_id)
        if source_disk and not isinstance(source_disk, str):
            raise TypeError("Expected argument 'source_disk' to be a str")
        pulumi.set(__self__, "source_disk", source_disk)
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
        if source_storage_object and not isinstance(source_storage_object, str):
            raise TypeError("Expected argument 'source_storage_object' to be a str")
        pulumi.set(__self__, "source_storage_object", source_storage_object)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        """
        The architecture of the disk. Valid values are ARM64 or X86_64.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="asyncPrimaryDisk")
    def async_primary_disk(self) -> 'outputs.DiskAsyncReplicationResponse':
        """
        Disk asynchronously replicated into this disk.
        """
        return pulumi.get(self, "async_primary_disk")

    @property
    @pulumi.getter(name="asyncSecondaryDisks")
    def async_secondary_disks(self) -> 'outputs.DiskAsyncReplicationListResponse':
        """
        A list of disks this disk is asynchronously replicated to.
        """
        return pulumi.get(self, "async_secondary_disks")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
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
    @pulumi.getter(name="diskEncryptionKey")
    def disk_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
        """
        return pulumi.get(self, "disk_encryption_key")

    @property
    @pulumi.getter(name="guestOsFeatures")
    def guest_os_features(self) -> Sequence['outputs.GuestOsFeatureResponse']:
        """
        A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        """
        return pulumi.get(self, "guest_os_features")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#disk for disks.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels to apply to this disk. These can be later modified by the setLabels method.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastAttachTimestamp")
    def last_attach_timestamp(self) -> str:
        """
        Last attach timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_attach_timestamp")

    @property
    @pulumi.getter(name="lastDetachTimestamp")
    def last_detach_timestamp(self) -> str:
        """
        Last detach timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_detach_timestamp")

    @property
    @pulumi.getter(name="licenseCodes")
    def license_codes(self) -> Sequence[str]:
        """
        Integer license codes indicating which licenses are attached to this disk.
        """
        return pulumi.get(self, "license_codes")

    @property
    @pulumi.getter
    def licenses(self) -> Sequence[str]:
        """
        A list of publicly visible licenses. Reserved for Google's use.
        """
        return pulumi.get(self, "licenses")

    @property
    @pulumi.getter(name="locationHint")
    def location_hint(self) -> str:
        """
        An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
        """
        return pulumi.get(self, "location_hint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> str:
        """
        Internal use only.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def params(self) -> 'outputs.DiskParamsResponse':
        """
        Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
        """
        return pulumi.get(self, "params")

    @property
    @pulumi.getter(name="physicalBlockSizeBytes")
    def physical_block_size_bytes(self) -> str:
        """
        Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
        """
        return pulumi.get(self, "physical_block_size_bytes")

    @property
    @pulumi.getter(name="provisionedIops")
    def provisioned_iops(self) -> str:
        """
        Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
        """
        return pulumi.get(self, "provisioned_iops")

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> str:
        """
        Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must be between 1 and 7,124.
        """
        return pulumi.get(self, "provisioned_throughput")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="replicaZones")
    def replica_zones(self) -> Sequence[str]:
        """
        URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
        """
        return pulumi.get(self, "replica_zones")

    @property
    @pulumi.getter(name="resourcePolicies")
    def resource_policies(self) -> Sequence[str]:
        """
        Resource policies applied to this disk for automatic snapshot creations.
        """
        return pulumi.get(self, "resource_policies")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> 'outputs.DiskResourceStatusResponse':
        """
        Status information for the disk resource.
        """
        return pulumi.get(self, "resource_status")

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
        Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> str:
        """
        Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are greater than 0.
        """
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter(name="sourceConsistencyGroupPolicy")
    def source_consistency_group_policy(self) -> str:
        """
        URL of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
        """
        return pulumi.get(self, "source_consistency_group_policy")

    @property
    @pulumi.getter(name="sourceConsistencyGroupPolicyId")
    def source_consistency_group_policy_id(self) -> str:
        """
        ID of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
        """
        return pulumi.get(self, "source_consistency_group_policy_id")

    @property
    @pulumi.getter(name="sourceDisk")
    def source_disk(self) -> str:
        """
        The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
        """
        return pulumi.get(self, "source_disk")

    @property
    @pulumi.getter(name="sourceDiskId")
    def source_disk_id(self) -> str:
        """
        The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
        """
        return pulumi.get(self, "source_disk_id")

    @property
    @pulumi.getter(name="sourceImage")
    def source_image(self) -> str:
        """
        The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
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
        The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
        """
        return pulumi.get(self, "source_image_id")

    @property
    @pulumi.getter(name="sourceSnapshot")
    def source_snapshot(self) -> str:
        """
        The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
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
        The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        """
        return pulumi.get(self, "source_snapshot_id")

    @property
    @pulumi.getter(name="sourceStorageObject")
    def source_storage_object(self) -> str:
        """
        The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
        """
        return pulumi.get(self, "source_storage_object")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting. 
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def users(self) -> Sequence[str]:
        """
        Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
        """
        return pulumi.get(self, "users")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "zone")


class AwaitableGetDiskResult(GetDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiskResult(
            architecture=self.architecture,
            async_primary_disk=self.async_primary_disk,
            async_secondary_disks=self.async_secondary_disks,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            disk_encryption_key=self.disk_encryption_key,
            guest_os_features=self.guest_os_features,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            last_attach_timestamp=self.last_attach_timestamp,
            last_detach_timestamp=self.last_detach_timestamp,
            license_codes=self.license_codes,
            licenses=self.licenses,
            location_hint=self.location_hint,
            name=self.name,
            options=self.options,
            params=self.params,
            physical_block_size_bytes=self.physical_block_size_bytes,
            provisioned_iops=self.provisioned_iops,
            provisioned_throughput=self.provisioned_throughput,
            region=self.region,
            replica_zones=self.replica_zones,
            resource_policies=self.resource_policies,
            resource_status=self.resource_status,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            size_gb=self.size_gb,
            source_consistency_group_policy=self.source_consistency_group_policy,
            source_consistency_group_policy_id=self.source_consistency_group_policy_id,
            source_disk=self.source_disk,
            source_disk_id=self.source_disk_id,
            source_image=self.source_image,
            source_image_encryption_key=self.source_image_encryption_key,
            source_image_id=self.source_image_id,
            source_snapshot=self.source_snapshot,
            source_snapshot_encryption_key=self.source_snapshot_encryption_key,
            source_snapshot_id=self.source_snapshot_id,
            source_storage_object=self.source_storage_object,
            status=self.status,
            type=self.type,
            users=self.users,
            zone=self.zone)


def get_disk(disk: Optional[str] = None,
             project: Optional[str] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiskResult:
    """
    Returns the specified persistent disk.
    """
    __args__ = dict()
    __args__['disk'] = disk
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getDisk', __args__, opts=opts, typ=GetDiskResult).value

    return AwaitableGetDiskResult(
        architecture=pulumi.get(__ret__, 'architecture'),
        async_primary_disk=pulumi.get(__ret__, 'async_primary_disk'),
        async_secondary_disks=pulumi.get(__ret__, 'async_secondary_disks'),
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        description=pulumi.get(__ret__, 'description'),
        disk_encryption_key=pulumi.get(__ret__, 'disk_encryption_key'),
        guest_os_features=pulumi.get(__ret__, 'guest_os_features'),
        kind=pulumi.get(__ret__, 'kind'),
        label_fingerprint=pulumi.get(__ret__, 'label_fingerprint'),
        labels=pulumi.get(__ret__, 'labels'),
        last_attach_timestamp=pulumi.get(__ret__, 'last_attach_timestamp'),
        last_detach_timestamp=pulumi.get(__ret__, 'last_detach_timestamp'),
        license_codes=pulumi.get(__ret__, 'license_codes'),
        licenses=pulumi.get(__ret__, 'licenses'),
        location_hint=pulumi.get(__ret__, 'location_hint'),
        name=pulumi.get(__ret__, 'name'),
        options=pulumi.get(__ret__, 'options'),
        params=pulumi.get(__ret__, 'params'),
        physical_block_size_bytes=pulumi.get(__ret__, 'physical_block_size_bytes'),
        provisioned_iops=pulumi.get(__ret__, 'provisioned_iops'),
        provisioned_throughput=pulumi.get(__ret__, 'provisioned_throughput'),
        region=pulumi.get(__ret__, 'region'),
        replica_zones=pulumi.get(__ret__, 'replica_zones'),
        resource_policies=pulumi.get(__ret__, 'resource_policies'),
        resource_status=pulumi.get(__ret__, 'resource_status'),
        satisfies_pzs=pulumi.get(__ret__, 'satisfies_pzs'),
        self_link=pulumi.get(__ret__, 'self_link'),
        size_gb=pulumi.get(__ret__, 'size_gb'),
        source_consistency_group_policy=pulumi.get(__ret__, 'source_consistency_group_policy'),
        source_consistency_group_policy_id=pulumi.get(__ret__, 'source_consistency_group_policy_id'),
        source_disk=pulumi.get(__ret__, 'source_disk'),
        source_disk_id=pulumi.get(__ret__, 'source_disk_id'),
        source_image=pulumi.get(__ret__, 'source_image'),
        source_image_encryption_key=pulumi.get(__ret__, 'source_image_encryption_key'),
        source_image_id=pulumi.get(__ret__, 'source_image_id'),
        source_snapshot=pulumi.get(__ret__, 'source_snapshot'),
        source_snapshot_encryption_key=pulumi.get(__ret__, 'source_snapshot_encryption_key'),
        source_snapshot_id=pulumi.get(__ret__, 'source_snapshot_id'),
        source_storage_object=pulumi.get(__ret__, 'source_storage_object'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'),
        users=pulumi.get(__ret__, 'users'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_disk)
def get_disk_output(disk: Optional[pulumi.Input[str]] = None,
                    project: Optional[pulumi.Input[Optional[str]]] = None,
                    zone: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDiskResult]:
    """
    Returns the specified persistent disk.
    """
    ...
