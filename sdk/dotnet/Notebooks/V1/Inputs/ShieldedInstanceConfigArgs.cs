// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Inputs
{

    /// <summary>
    /// A set of Shielded Instance options. Check [Images using supported Shielded VM features] Not all combinations are valid.
    /// </summary>
    public sealed class ShieldedInstanceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the boot integrity of the instance. The attestation is performed against the integrity policy baseline. This baseline is initially derived from the implicitly trusted boot image when the instance is created. Enabled by default.
        /// </summary>
        [Input("enableIntegrityMonitoring")]
        public Input<bool>? EnableIntegrityMonitoring { get; set; }

        /// <summary>
        /// Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs authentic software by verifying the digital signature of all boot components, and halting the boot process if signature verification fails. Disabled by default.
        /// </summary>
        [Input("enableSecureBoot")]
        public Input<bool>? EnableSecureBoot { get; set; }

        /// <summary>
        /// Defines whether the instance has the vTPM enabled. Enabled by default.
        /// </summary>
        [Input("enableVtpm")]
        public Input<bool>? EnableVtpm { get; set; }

        public ShieldedInstanceConfigArgs()
        {
        }
        public static new ShieldedInstanceConfigArgs Empty => new ShieldedInstanceConfigArgs();
    }
}
