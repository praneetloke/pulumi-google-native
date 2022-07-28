// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataFusion.V1.Inputs
{

    /// <summary>
    /// The Data Fusion version. This proto message stores information about certain Data Fusion version, which is used for Data Fusion version upgrade.
    /// </summary>
    public sealed class VersionArgs : global::Pulumi.ResourceArgs
    {
        [Input("availableFeatures")]
        private InputList<string>? _availableFeatures;

        /// <summary>
        /// Represents a list of available feature names for a given version.
        /// </summary>
        public InputList<string> AvailableFeatures
        {
            get => _availableFeatures ?? (_availableFeatures = new InputList<string>());
            set => _availableFeatures = value;
        }

        /// <summary>
        /// Whether this is currently the default version for Cloud Data Fusion
        /// </summary>
        [Input("defaultVersion")]
        public Input<bool>? DefaultVersion { get; set; }

        /// <summary>
        /// Type represents the release availability of the version
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.DataFusion.V1.VersionType>? Type { get; set; }

        /// <summary>
        /// The version number of the Data Fusion instance, such as '6.0.1.0'.
        /// </summary>
        [Input("versionNumber")]
        public Input<string>? VersionNumber { get; set; }

        public VersionArgs()
        {
        }
        public static new VersionArgs Empty => new VersionArgs();
    }
}
