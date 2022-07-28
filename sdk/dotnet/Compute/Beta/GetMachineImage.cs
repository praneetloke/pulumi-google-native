// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetMachineImage
    {
        /// <summary>
        /// Returns the specified machine image. Gets a list of available machine images by making a list() request.
        /// </summary>
        public static Task<GetMachineImageResult> InvokeAsync(GetMachineImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMachineImageResult>("google-native:compute/beta:getMachineImage", args ?? new GetMachineImageArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified machine image. Gets a list of available machine images by making a list() request.
        /// </summary>
        public static Output<GetMachineImageResult> Invoke(GetMachineImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMachineImageResult>("google-native:compute/beta:getMachineImage", args ?? new GetMachineImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineImageArgs : global::Pulumi.InvokeArgs
    {
        [Input("machineImage", required: true)]
        public string MachineImage { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetMachineImageArgs()
        {
        }
        public static new GetMachineImageArgs Empty => new GetMachineImageArgs();
    }

    public sealed class GetMachineImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("machineImage", required: true)]
        public Input<string> MachineImage { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetMachineImageInvokeArgs()
        {
        }
        public static new GetMachineImageInvokeArgs Empty => new GetMachineImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineImageResult
    {
        /// <summary>
        /// The creation timestamp for this machine image in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// </summary>
        public readonly bool GuestFlush;
        /// <summary>
        /// Properties of source instance
        /// </summary>
        public readonly Outputs.InstancePropertiesResponse InstanceProperties;
        /// <summary>
        /// The resource type, which is always compute#machineImage for machine image.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse MachineImageEncryptionKey;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// An array of Machine Image specific properties for disks attached to the source instance
        /// </summary>
        public readonly ImmutableArray<Outputs.SavedDiskResponse> SavedDisks;
        /// <summary>
        /// The URL for this machine image. The server defines this URL.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly ImmutableArray<Outputs.SourceDiskEncryptionKeyResponse> SourceDiskEncryptionKeys;
        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
        /// </summary>
        public readonly string SourceInstance;
        /// <summary>
        /// DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
        /// </summary>
        public readonly Outputs.SourceInstancePropertiesResponse SourceInstanceProperties;
        /// <summary>
        /// The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        /// </summary>
        public readonly ImmutableArray<string> StorageLocations;
        /// <summary>
        /// Total size of the storage used by the machine image.
        /// </summary>
        public readonly string TotalStorageBytes;

        [OutputConstructor]
        private GetMachineImageResult(
            string creationTimestamp,

            string description,

            bool guestFlush,

            Outputs.InstancePropertiesResponse instanceProperties,

            string kind,

            Outputs.CustomerEncryptionKeyResponse machineImageEncryptionKey,

            string name,

            bool satisfiesPzs,

            ImmutableArray<Outputs.SavedDiskResponse> savedDisks,

            string selfLink,

            ImmutableArray<Outputs.SourceDiskEncryptionKeyResponse> sourceDiskEncryptionKeys,

            string sourceInstance,

            Outputs.SourceInstancePropertiesResponse sourceInstanceProperties,

            string status,

            ImmutableArray<string> storageLocations,

            string totalStorageBytes)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            GuestFlush = guestFlush;
            InstanceProperties = instanceProperties;
            Kind = kind;
            MachineImageEncryptionKey = machineImageEncryptionKey;
            Name = name;
            SatisfiesPzs = satisfiesPzs;
            SavedDisks = savedDisks;
            SelfLink = selfLink;
            SourceDiskEncryptionKeys = sourceDiskEncryptionKeys;
            SourceInstance = sourceInstance;
            SourceInstanceProperties = sourceInstanceProperties;
            Status = status;
            StorageLocations = storageLocations;
            TotalStorageBytes = totalStorageBytes;
        }
    }
}
