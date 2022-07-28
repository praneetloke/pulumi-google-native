// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Represents a single Zypper package repository. This repository is added to a repo file that is stored at `/etc/zypp/repos.d/google_osconfig.repo`.
    /// </summary>
    public sealed class ZypperRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the repository directory.
        /// </summary>
        [Input("baseUrl", required: true)]
        public Input<string> BaseUrl { get; set; } = null!;

        /// <summary>
        /// The display name of the repository.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("gpgKeys")]
        private InputList<string>? _gpgKeys;

        /// <summary>
        /// URIs of GPG keys.
        /// </summary>
        public InputList<string> GpgKeys
        {
            get => _gpgKeys ?? (_gpgKeys = new InputList<string>());
            set => _gpgKeys = value;
        }

        /// <summary>
        /// A one word, unique name for this repository. This is the `repo id` in the zypper config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for guest policy conflicts.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public ZypperRepositoryArgs()
        {
        }
        public static new ZypperRepositoryArgs Empty => new ZypperRepositoryArgs();
    }
}
