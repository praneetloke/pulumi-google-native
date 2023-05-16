// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// [Input Only] Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance. This field is persisted and returned for instanceTemplate and not returned in the context of instance. This property is mutually exclusive with the source property; you can only define one or the other, but not both.
    /// </summary>
    public sealed class AttachedDiskInitializeParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The architecture of the attached disk. Valid values are arm64 or x86_64.
        /// </summary>
        [Input("architecture")]
        public Input<Pulumi.GoogleNative.Compute.V1.AttachedDiskInitializeParamsArchitecture>? Architecture { get; set; }

        /// <summary>
        /// An optional description. Provide this property when creating the disk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the disk name. If not specified, the default is to use the name of the instance. If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Specifies the size of the disk in base-2 GB. The size must be at least 10 GB. If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage. If you do not specify a sourceImage, the default disk size is 500 GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// Specifies the disk type to use to create the instance. If not specified, the default is pd-standard, specified using the full URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/pd-standard For a full list of acceptable values, see Persistent disk types. If you specify this field when creating a VM, you can provide either the full or partial URL. For example, the following values are valid: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /diskTypes/diskType - projects/project/zones/zone/diskTypes/diskType - zones/zone/diskTypes/diskType If you specify this field when creating or updating an instance template or all-instances configuration, specify the type of the disk, not the URL. For example: pd-standard.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this disk. These can be later modified by the disks.setLabels method. This field is only applicable for persistent disks.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// A list of publicly visible licenses. Reserved for Google's use.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Specifies which action to take on instance update with this disk. Default is to use the existing disk.
        /// </summary>
        [Input("onUpdateAction")]
        public Input<Pulumi.GoogleNative.Compute.V1.AttachedDiskInitializeParamsOnUpdateAction>? OnUpdateAction { get; set; }

        /// <summary>
        /// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
        /// </summary>
        [Input("provisionedIops")]
        public Input<string>? ProvisionedIops { get; set; }

        /// <summary>
        /// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must be between 1 and 7,124.
        /// </summary>
        [Input("provisionedThroughput")]
        public Input<string>? ProvisionedThroughput { get; set; }

        [Input("replicaZones")]
        private InputList<string>? _replicaZones;

        /// <summary>
        /// Required for each regional disk associated with the instance. Specify the URLs of the zones where the disk should be replicated to. You must provide exactly two replica zones, and one zone must be the same as the instance zone. You can't use this option with boot disks.
        /// </summary>
        public InputList<string> ReplicaZones
        {
            get => _replicaZones ?? (_replicaZones = new InputList<string>());
            set => _replicaZones = value;
        }

        [Input("resourceManagerTags")]
        private InputMap<string>? _resourceManagerTags;

        /// <summary>
        /// Resource manager tags to be bound to the disk. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT &amp; PATCH) when empty.
        /// </summary>
        public InputMap<string> ResourceManagerTags
        {
            get => _resourceManagerTags ?? (_resourceManagerTags = new InputMap<string>());
            set => _resourceManagerTags = value;
        }

        [Input("resourcePolicies")]
        private InputList<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations. Specified using the full or partial URL. For instance template, specify only the resource policy name.
        /// </summary>
        public InputList<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// The source image to create this disk. When creating a new instance, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required except for local SSD. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family If the source image is deleted later, this field will not be set.
        /// </summary>
        [Input("sourceImage")]
        public Input<string>? SourceImage { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key. InstanceTemplate and InstancePropertiesPatch do not store customer-supplied encryption keys, so you cannot create disks for instances in a managed instance group if the source images are encrypted with your own keys.
        /// </summary>
        [Input("sourceImageEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceImageEncryptionKey { get; set; }

        /// <summary>
        /// The source snapshot to create this disk. When creating a new instance, one of initializeParams.sourceSnapshot or initializeParams.sourceImage or disks.source is required except for local SSD. To create a disk with a snapshot that you created, specify the snapshot name in the following format: global/snapshots/my-backup If the source snapshot is deleted later, this field will not be set.
        /// </summary>
        [Input("sourceSnapshot")]
        public Input<string>? SourceSnapshot { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot.
        /// </summary>
        [Input("sourceSnapshotEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? SourceSnapshotEncryptionKey { get; set; }

        public AttachedDiskInitializeParamsArgs()
        {
        }
        public static new AttachedDiskInitializeParamsArgs Empty => new AttachedDiskInitializeParamsArgs();
    }
}
