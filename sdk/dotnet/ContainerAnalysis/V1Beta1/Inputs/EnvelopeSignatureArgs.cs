// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    public sealed class EnvelopeSignatureArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyid")]
        public Input<string>? Keyid { get; set; }

        [Input("sig")]
        public Input<string>? Sig { get; set; }

        public EnvelopeSignatureArgs()
        {
        }
        public static new EnvelopeSignatureArgs Empty => new EnvelopeSignatureArgs();
    }
}
