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
    /// Apt patching is completed by executing `apt-get update &amp;&amp; apt-get upgrade`. Additional options can be set to control how this is executed.
    /// </summary>
    public sealed class AptSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludes")]
        private InputList<string>? _excludes;

        /// <summary>
        /// List of packages to exclude from update. These packages will be excluded
        /// </summary>
        public InputList<string> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<string>());
            set => _excludes = value;
        }

        [Input("exclusivePackages")]
        private InputList<string>? _exclusivePackages;

        /// <summary>
        /// An exclusive list of packages to be updated. These are the only packages that will be updated. If these packages are not installed, they will be ignored. This field cannot be specified with any other patch configuration fields.
        /// </summary>
        public InputList<string> ExclusivePackages
        {
            get => _exclusivePackages ?? (_exclusivePackages = new InputList<string>());
            set => _exclusivePackages = value;
        }

        /// <summary>
        /// By changing the type to DIST, the patching is performed using `apt-get dist-upgrade` instead.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.OSConfig.V1Beta.AptSettingsType>? Type { get; set; }

        public AptSettingsArgs()
        {
        }
        public static new AptSettingsArgs Empty => new AptSettingsArgs();
    }
}
