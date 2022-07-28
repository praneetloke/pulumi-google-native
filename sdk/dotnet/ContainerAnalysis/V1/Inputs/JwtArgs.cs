// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    public sealed class JwtArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compact encoding of a JWS, which is always three base64 encoded strings joined by periods. For details, see: https://tools.ietf.org/html/rfc7515.html#section-3.1
        /// </summary>
        [Input("compactJwt")]
        public Input<string>? CompactJwt { get; set; }

        public JwtArgs()
        {
        }
        public static new JwtArgs Empty => new JwtArgs();
    }
}
