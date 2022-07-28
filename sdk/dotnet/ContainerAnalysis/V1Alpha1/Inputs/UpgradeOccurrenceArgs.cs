// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// An Upgrade Occurrence represents that a specific resource_url could install a specific upgrade. This presence is supplied via local sources (i.e. it is present in the mirror and the running system has noticed its availability).
    /// </summary>
    public sealed class UpgradeOccurrenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Metadata about the upgrade for available for the specific operating system for the resource_url. This allows efficient filtering, as well as making it easier to use the occurrence.
        /// </summary>
        [Input("distribution")]
        public Input<Inputs.UpgradeDistributionArgs>? Distribution { get; set; }

        /// <summary>
        /// Required - The package this Upgrade is for.
        /// </summary>
        [Input("package")]
        public Input<string>? Package { get; set; }

        /// <summary>
        /// Required - The version of the package in a machine + human readable form.
        /// </summary>
        [Input("parsedVersion")]
        public Input<Inputs.VersionArgs>? ParsedVersion { get; set; }

        public UpgradeOccurrenceArgs()
        {
        }
        public static new UpgradeOccurrenceArgs Empty => new UpgradeOccurrenceArgs();
    }
}
