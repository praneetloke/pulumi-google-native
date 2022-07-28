// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIKeys.V2.Inputs
{

    /// <summary>
    /// Identifier of an Android application for key use.
    /// </summary>
    public sealed class V2AndroidApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The package name of the application.
        /// </summary>
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        /// <summary>
        /// The SHA1 fingerprint of the application. For example, both sha1 formats are acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format is the latter.
        /// </summary>
        [Input("sha1Fingerprint")]
        public Input<string>? Sha1Fingerprint { get; set; }

        public V2AndroidApplicationArgs()
        {
        }
        public static new V2AndroidApplicationArgs Empty => new V2AndroidApplicationArgs();
    }
}
