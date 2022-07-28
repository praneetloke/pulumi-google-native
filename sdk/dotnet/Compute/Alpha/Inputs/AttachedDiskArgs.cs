// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// An instance-attached disk resource.
    /// </summary>
    public sealed class AttachedDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the disk will be auto-deleted when the instance is deleted (but not when the disk is detached from the instance).
        /// </summary>
        [Input("autoDelete")]
        public Input<bool>? AutoDelete { get; set; }

        /// <summary>
        /// Indicates that this is a boot disk. The virtual machine will use the first partition of the disk for its root filesystem.
        /// </summary>
        [Input("boot")]
        public Input<bool>? Boot { get; set; }

        /// <summary>
        /// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance. This name can be used to reference the device for mounting, resizing, and so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine. This field is only applicable for persistent disks.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group.
        /// </summary>
        [Input("diskEncryptionKey")]
        public Input<Inputs.CustomerEncryptionKeyArgs>? DiskEncryptionKey { get; set; }

        /// <summary>
        /// The size of the disk in GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// [Input Only] Whether to force attach the regional disk even if it's currently attached to another instance. If you try to force attach a zonal disk to an instance, you will receive an error.
        /// </summary>
        [Input("forceAttach")]
        public Input<bool>? ForceAttach { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.GuestOsFeatureArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        public InputList<Inputs.GuestOsFeatureArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.GuestOsFeatureArgs>());
            set => _guestOsFeatures = value;
        }

        /// <summary>
        /// [Input Only] Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance. This property is mutually exclusive with the source property; you can only define one or the other, but not both.
        /// </summary>
        [Input("initializeParams")]
        public Input<Inputs.AttachedDiskInitializeParamsArgs>? InitializeParams { get; set; }

        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Persistent disks must always use SCSI and the request will fail if you attempt to attach a persistent disk in any other format than SCSI. Local SSDs can use either NVME or SCSI. For performance characteristics of SCSI over NVMe, see Local SSD performance.
        /// </summary>
        [Input("interface")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.AttachedDiskInterface>? Interface { get; set; }

        /// <summary>
        /// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.AttachedDiskMode>? Mode { get; set; }

        /// <summary>
        /// For LocalSSD disks on VM Instances in STOPPED or SUSPENDED state, this field is set to PRESERVED if the LocalSSD data has been saved to a persistent location by customer request. (see the discard_local_ssd option on Stop/Suspend). Read-only in the api.
        /// </summary>
        [Input("savedState")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.AttachedDiskSavedState>? SavedState { get; set; }

        /// <summary>
        /// Specifies a valid partial or full URL to an existing Persistent Disk resource. When creating a new instance, one of initializeParams.sourceImage or initializeParams.sourceSnapshot or disks.source is required except for local SSD. If desired, you can also attach existing non-root persistent disks using this property. This field is only applicable for persistent disks. Note that for InstanceTemplate, specify the disk name for zonal disk, and the URL for regional disk.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Specifies the type of the disk, either SCRATCH or PERSISTENT. If not specified, the default is PERSISTENT.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.AttachedDiskType>? Type { get; set; }

        public AttachedDiskArgs()
        {
        }
        public static new AttachedDiskArgs Empty => new AttachedDiskArgs();
    }
}
