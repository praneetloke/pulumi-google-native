// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Returns the specified snapshot resource. Returns INVALID_ARGUMENT if called for a non-boot volume.
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("google-native:baremetalsolution/v2:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified snapshot resource. Returns INVALID_ARGUMENT if called for a non-boot volume.
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("google-native:baremetalsolution/v2:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("snapshotId", required: true)]
        public string SnapshotId { get; set; } = null!;

        [Input("volumeId", required: true)]
        public string VolumeId { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// The creation time of the snapshot.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the volume which this snapshot belongs to.
        /// </summary>
        public readonly string StorageVolume;
        /// <summary>
        /// The type of the snapshot which indicates whether it was scheduled or manual/ad-hoc.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSnapshotResult(
            string createTime,

            string description,

            string name,

            string storageVolume,

            string type)
        {
            CreateTime = createTime;
            Description = description;
            Name = name;
            StorageVolume = storageVolume;
            Type = type;
        }
    }
}
