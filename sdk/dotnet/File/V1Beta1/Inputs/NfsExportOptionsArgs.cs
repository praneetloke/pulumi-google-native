// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.File.V1Beta1.Inputs
{

    /// <summary>
    /// NFS export options specifications.
    /// </summary>
    public sealed class NfsExportOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
        /// </summary>
        [Input("accessMode")]
        public Input<Pulumi.GoogleNative.File.V1Beta1.NfsExportOptionsAccessMode>? AccessMode { get; set; }

        /// <summary>
        /// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
        /// </summary>
        [Input("anonGid")]
        public Input<string>? AnonGid { get; set; }

        /// <summary>
        /// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
        /// </summary>
        [Input("anonUid")]
        public Input<string>? AnonUid { get; set; }

        [Input("ipRanges")]
        private InputList<string>? _ipRanges;

        /// <summary>
        /// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        /// </summary>
        public InputList<string> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<string>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
        /// </summary>
        [Input("squashMode")]
        public Input<Pulumi.GoogleNative.File.V1Beta1.NfsExportOptionsSquashMode>? SquashMode { get; set; }

        public NfsExportOptionsArgs()
        {
        }
        public static new NfsExportOptionsArgs Empty => new NfsExportOptionsArgs();
    }
}
